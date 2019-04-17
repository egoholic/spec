package token_test

import (
	"github.com/egoholic/spec/rawsig"
	. "github.com/egoholic/spec/token"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("token", func() {
	Describe("tokens", func() {
		Describe("runeToken", func() {
			Describe(".Matches()", func() {
				Context("when matches", func() {
					It("returns true", func() {
						r := 'a'
						rawSig := rawsig.New("alberta")
						runeToken := NewRuneToken(r)
						Expect(runeToken.Matches(rawSig)).To(BeTrue())
					})
				})
			})
		})

		Describe("anyRuneToken", func() {
			Describe(".Matches()", func() {
				Context("when matches", func() {
					It("returns true", func() {
						rawSig := rawsig.New("alberta")
						anyRuneToken := NewAnyRuneToken()
						Expect(anyRuneToken.Matches(rawSig)).To(BeTrue())
					})
				})
			})
		})

		Describe("variantToken", func() {
			Describe(".Matches()", func() {
				Context("when matches", func() {
					It("returns true", func() {
						rawSig := rawsig.New("alberta")
						token1 := NewRuneToken('a')
						token2 := NewRuneToken('z')
						varToken := NewVariantToken([]Token{token1, token2})
						Expect(varToken.Matches(rawSig)).To(BeTrue())
					})
				})
			})
		})

		Describe("variantToken", func() {
			Describe(".Matches()", func() {
				Context("when matches", func() {
					It("returns true", func() {
						rawSig := rawsig.New("alberta")
						token1 := NewRuneToken('a')
						token2 := NewRuneToken('z')
						varToken := NewVariantToken([]Token{token1, token2})
						Expect(varToken.Matches(rawSig)).To(BeTrue())
					})
				})
			})
		})
		// Describe("wordToken", func() {
		// 	Context("matching", func() {
		// 		Describe(".Matches()", func() {
		// 			Context("when matches", func() {
		// 				It("returns true", func() {
		// 					rawSig := rawsig.New("alberta")
		// 					wordToken := NewWordToken(nil, NewTokenListFromString("albert"), nil)
		// 					Expect(wordToken.Matches(rawSig)).To(BeTrue())
		// 				})
		// 			})
		// 		})
		// 	})
		// })
	})
})
