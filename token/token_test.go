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
						token1 := NewRuneToken('a')
						token2 := NewRuneToken('z')
						varToken := NewVariantToken([]Token{token1, token2})
						Expect(varToken.Matches(rawsig.New("a"))).To(BeTrue())
						Expect(varToken.Matches(rawsig.New("z"))).To(BeTrue())

					})
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
