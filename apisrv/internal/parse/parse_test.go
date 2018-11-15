package parse_test

import (
	"errors"
	"reflect"

	"../structs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "."
)

var _ = Describe("Validate", func() {
	Context("with valid input data", func() {
		jsonb := []byte(`{"config_header":"foo","pio_config":"bar"}`)
		expected := structs.CompileRequest{ConfigHeader: "foo", PioConfig: "bar"}

		It("validates and parses", func() {
			Ω(Validate(jsonb)).Should(Equal(expected))
		})
	})

	Context("with invalid JSON", func() {
		jsonb := []byte(`{"unparseable due to: (syntax errors)`)

		It("returns the parsing error", func() {
			_, err := Validate(jsonb)
			Ω(reflect.TypeOf(err).String()).Should(Equal("*json.SyntaxError"))
		})
	})

	Context("missing config_header", func() {
		jsonb := []byte(`{"pio_config":"bar"}`)

		It("returns the validation error", func() {
			_, err := Validate(jsonb)
			Ω(err).Should(Equal(errors.New("config_header must not be empty")))
		})
	})

	Context("missing pio_config", func() {
		jsonb := []byte(`{"config_header":"foo"}`)

		It("returns the validation error", func() {
			_, err := Validate(jsonb)
			Ω(err).Should(Equal(errors.New("pio_config must not be empty")))
		})
	})
})
