package render_test

import (
	"errors"
	"io/ioutil"
	"math"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "."
)

func TestParse(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Parse Suite")
}

var _ = Context("recording written responses", func() {
	var rr httptest.ResponseRecorder
	BeforeEach(func() {
		rr = *httptest.NewRecorder()
	})

	Describe("ValidResult", func() {
		Context("with serializable data", func() {
			data := map[string]string{"foo": "bar"}

			It("writes the expected data", func() {
				ValidResult(&rr, data)
				r := rr.Result()
				body, _ := ioutil.ReadAll(r.Body)
				Ω(r.StatusCode).Should(Equal(http.StatusOK))
				Ω(r.Header.Get("Content-Type")).Should(Equal("application/json"))
				Ω(string(body)).Should(Equal(`{"foo":"bar"}`))
			})
		})

		Context("with unserializable data", func() {
			data := math.Inf(0)

			It("writes the error", func() {
				ValidResult(&rr, data)
				r := rr.Result()
				body, _ := ioutil.ReadAll(r.Body)
				Ω(r.StatusCode).Should(Equal(http.StatusInternalServerError))
				Ω(r.Header.Get("Content-Type")).Should(Equal("application/json"))
				Ω(string(body)).Should(Equal(`{"error":"*json.UnsupportedValueError: json: unsupported value: +Inf"}`))
			})
		})
	})

	Describe("ClientError", func() {
		error := errors.New("User age cannot be negative")

		It("writes the expected data", func() {
			ClientError(&rr, error)
			r := rr.Result()
			body, _ := ioutil.ReadAll(r.Body)
			Ω(r.StatusCode).Should(Equal(http.StatusBadRequest))
			Ω(r.Header.Get("Content-Type")).Should(Equal("application/json"))
			Ω(string(body)).Should(Equal(`{"error":"*errors.errorString: User age cannot be negative"}`))
		})
	})

	Describe("ServerError", func() {
		error := errors.New("Out of disk space")

		It("writes the expected data", func() {
			ServerError(&rr, error)
			r := rr.Result()
			body, _ := ioutil.ReadAll(r.Body)
			Ω(r.StatusCode).Should(Equal(http.StatusInternalServerError))
			Ω(r.Header.Get("Content-Type")).Should(Equal("application/json"))
			Ω(string(body)).Should(Equal(`{"error":"*errors.errorString: Out of disk space"}`))
		})
	})
})
