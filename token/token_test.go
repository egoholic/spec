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
						rawSig := rawsig.New("alberta")
						runeToken := NewRuneToken('a')
						Expect(runeToken.Matches(rawSig)).To(BeTrue())
					})
				})
			})

			Describe(".String()", func() {
				It("returns string", func() {
					runeToken := NewRuneToken('a')
					Expect(runeToken.String()).To(Equal("a"))
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

				Describe(".String()", func() {
					It("returns string", func() {
						anyRuneToken := NewAnyRuneToken()
						Expect(anyRuneToken.String()).To(Equal("<any-rune>"))
					})
				})
			})
		})

		Describe("variantToken", func() {
			Describe(".Matches()", func() {
				Context("when matches", func() {
					It("returns true", func() {
						rawSig1 := rawsig.New("alberta")
						rawSig2 := rawsig.New("zero")
						rawSig3 := rawsig.New("kyiv")
						token1 := NewRuneToken('a')
						token2 := NewRuneToken('z')
						varToken := NewVariantToken([]Token{token1, token2})
						Expect(varToken.Matches(rawSig1)).To(BeTrue())
						Expect(varToken.Matches(rawSig2)).To(BeTrue())
						Expect(varToken.Matches(rawSig3)).To(BeFalse())
					})
				})
			})

			Describe(".String()", func() {
				It("returns string", func() {
					token1 := NewRuneToken('a')
					token2 := NewRuneToken('z')
					varToken := NewVariantToken([]Token{token1, token2})
					Expect(varToken.String()).To(Equal("(one-of a z )"))
				})
			})

			Describe(".Join()", func() {
				It("returns new variant token with options from both", func() {
					token1 := NewRuneToken('a')
					token2 := NewRuneToken('b')
					token3 := NewRuneToken('c')
					token4 := NewRuneToken('d')
					varToken1 := NewVariantToken([]Token{token1, token2})
					varToken2 := NewVariantToken([]Token{token3, token4})
					varToken3 := varToken1.Join(varToken2)
					Expect(varToken1.Matches(rawsig.New("a"))).To(BeTrue())
					Expect(varToken1.Matches(rawsig.New("b"))).To(BeTrue())
					Expect(varToken1.Matches(rawsig.New("c"))).To(BeFalse())
					Expect(varToken1.Matches(rawsig.New("d"))).To(BeFalse())

					Expect(varToken2.Matches(rawsig.New("a"))).To(BeFalse())
					Expect(varToken2.Matches(rawsig.New("b"))).To(BeFalse())
					Expect(varToken2.Matches(rawsig.New("c"))).To(BeTrue())
					Expect(varToken2.Matches(rawsig.New("d"))).To(BeTrue())

					Expect(varToken3.Matches(rawsig.New("a"))).To(BeTrue())
					Expect(varToken3.Matches(rawsig.New("b"))).To(BeTrue())
					Expect(varToken3.Matches(rawsig.New("c"))).To(BeTrue())
					Expect(varToken3.Matches(rawsig.New("d"))).To(BeTrue())
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
