package token

import (
	"github.com/egoholic/spec/rawsig"
)

type Token interface {
	Matches(*rawsig.RawSignature) bool
}

type SingleSymbolToken rune

func (token *SingleSymbolToken) Matches(rawSig *rawsig.RawSignature) bool {
	next, err := rawSig.Next()
	if err != nil {
		return false
	}
	return rune(*token) == next
}

type PrimitiveToken []rune
func (token *PrimitiveToken) Matches(rawSig *rawsig.RawSignature) bool {
  for _, tokenRune := range *token {
		rawSigRune, err := rawSig.Next()
		if err != nil {
			return false
		}
		if tokenRune != rawSigRune {
			return false
		}
	}
	return true
}

type BinaryToken struct {
	token1 Token
	token2 Token
}

func (token *BinaryToken) Matches(rawSig *rawsig.RawSignature) bool {
  return token.token1.Matches(rawSig) && token.token2.Matches(rawSig)
}

type TernaryToken struct {
	token1 Token
	token2 Token
	token3 Token
}

type TernaryTokenSetUniqueByFirstToken struct {
	tokens map[string]*TernaryToken
}

func (token *TernaryToken) Matches(rawSig *rawsig.RawSignature) bool {
  return token.token1.Matches(rawSig) && token.token2.Matches(rawSig) && token.token3.Matches(rawSig)
}

var (
	IntToken        = PrimitiveToken("int")
	StringToken     = PrimitiveToken("string")
	FloatToken      = PrimitiveToken("float")
	BoolToken       = PrimitiveToken("bool")
	PrimitiveTokens = []PrimitiveToken{IntToken, StringToken, FloatToken, BoolToken}
	SliceToken      = BinaryToken{PrimitiveToken("[]"), Token}
	MapToken        = BinaryToken{PrimitiveToken("map["), Token, SingleSymbolToken(']'), Token}
)

func init() {
	var KeyValuePairToken = TernaryToken(Token, SingleSymbolToken(':'), Token)
	var StructToken =
}
