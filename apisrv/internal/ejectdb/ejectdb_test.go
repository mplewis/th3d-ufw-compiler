package ejectdb_test

import (
	"testing"

	"../structs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "."
)

func TestParse(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Parse Suite")
}

var _ = Describe("EjectDB", func() {
	It("ejects results when new ones are inserted", func() {
		db := NewEjectDB(2)
		Ω(len(db.Results)).Should(Equal(0))

		db.Put("foo", structs.CompileResult{IntelHex: "hex_for_foo"})
		db.Put("bar", structs.CompileResult{IntelHex: "hex_for_bar"})
		db.Put("baz", structs.CompileResult{IntelHex: "hex_for_baz"})
		db.Put("quux", structs.CompileResult{IntelHex: "hex_for_quux"})
		Ω(len(db.Results)).Should(Equal(2))
		Ω(db.Results["baz"].IntelHex).Should(Equal("hex_for_baz"))
		Ω(db.Results["quux"].IntelHex).Should(Equal("hex_for_quux"))
	})
})
