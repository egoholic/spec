package parser_test

import (
	. "github.com/egoholic/spec/parser"
	"github.com/egoholic/spec/rawsig"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("parser", func() {
	Context("parsing", func() {
		Describe("ParseSlice()", func() {
			Context("when correct raw signature", func() {
				It("returns signature", func() {
					rawSig := rawsig.New("[]string")
					parserCombinator := NewParserCombinator([]Parser{ParsePrimitive, ParseSlice})
					signature, err := ParseSlice(rawSig, parserCombinator)
					Expect(err).ShouldNot(HaveOccurred())
					Expect(signature.Token()).To(Equal("[]"))
					Expect(signature.Title()).To(Equal("[]string"))
				})
			})
		})
	})
})
