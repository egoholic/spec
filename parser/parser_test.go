package parser_test

import (
	. "github.com/egoholic/spec/parser"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("parser", func() {
	Describe("ParserCombinator", func() {
		Context("creation", func() {
			Describe("NewParserCombinator()", func() {
				It("returns parser combinator", func() {
					Expect(NewParserCombinator([]Parser{})).To(BeAssignableToTypeOf(&ParserCombinator{}))
				})
			})
		})
		Context("parsing", func() {
			Describe(".Parse()", func() {
				Context("when correct raw signature", func() {
					Context("when endpoint signature", func() {

					})

					Context("when raw struct signature", func() {

					})

					Context("when raw map signature", func() {

					})

					Context("when raw slice signature", func() {

					})

					Context("when raw primitive signature", func() {

					})
				})

				Context("when wrong raw signature", func() {

				})
			})
		})
	})
})
