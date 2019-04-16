package parser_test

import (
	. "github.com/egoholic/spec/parser"
	"github.com/egoholic/spec/rawsig"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("parser", func() {
	Context("parsing", func() {
		Describe("ParsePrimitive()", func() {
			Context("when correct raw signature", func() {
				Context("when string", func() {
					It("returns string signature", func() {
						rawSig := rawsig.New("string")
						parserCombinator := NewParserCombinator([]Parser{ParsePrimitive})
						signature, err := ParsePrimitive(rawSig, parserCombinator)
						Expect(err).ShouldNot(HaveOccurred())
						Expect(signature.Title()).To(Equal("string"))
						Expect(signature.Token()).To(Equal("string"))
					})
				})

				Context("when int", func() {
					It("returns int signature", func() {
						rawSig := rawsig.New("int")
						parserCombinator := NewParserCombinator([]Parser{ParsePrimitive})
						signature, err := ParsePrimitive(rawSig, parserCombinator)
						Expect(err).ShouldNot(HaveOccurred())
						Expect(signature.Title()).To(Equal("int"))
						Expect(signature.Token()).To(Equal("int"))
					})
				})

				Context("when float", func() {
					It("returns float signature", func() {
						rawSig := rawsig.New("float")
						parserCombinator := NewParserCombinator([]Parser{ParsePrimitive})
						signature, err := ParsePrimitive(rawSig, parserCombinator)
						Expect(err).ShouldNot(HaveOccurred())
						Expect(signature.Title()).To(Equal("float"))
						Expect(signature.Token()).To(Equal("float"))

					})
				})

				Context("when bool", func() {
					It("returns bool signature", func() {
						rawSig := rawsig.New("bool")
						parserCombinator := NewParserCombinator([]Parser{ParsePrimitive})
						signature, err := ParsePrimitive(rawSig, parserCombinator)
						Expect(err).ShouldNot(HaveOccurred())
						Expect(signature.Title()).To(Equal("bool"))
						Expect(signature.Token()).To(Equal("bool"))
					})
				})
			})

			Context("when wrong raw signature", func() {

			})
		})
	})
})
