package token

import (
	"github.com/egoholic/spec/rawsig"
)

var (
	ColonToken               = NewRuneToken(':')
	SemicolonToken           = NewRuneToken(';')
	OpeningBracketToken      = NewRuneToken('[')
	ClosingBracketToken      = NewRuneToken(']')
	OpeningCurlyBracketToken = NewRuneToken('{')
	ClosingCurlyBracketToken = NewRuneToken('}')
	SpaceToken               = NewVariantToken([]Token{NewRuneToken(' '), NewRuneToken('\n'), NewRuneToken('\t'), NewRuneToken('\t')})

	StringToken           = NewTokenListFromString("string")
	IntToken              = NewTokenListFromString("int")
	FloatToken            = NewTokenListFromString("float")
	BoolToken             = NewTokenListFromString("bool")
	AnyPrimitiveTypeToken = NewVariantToken([]Token{StringToken, IntToken, FloatToken, BoolToken})

	AnyStandardTypeToken *VariantToken

	SliceToken  = NewTokenListFromSlice([]Token{OpeningBracketToken, ClosingBracketToken, AnyStandardTypeToken})
	MapToken    = NewTokenListFromSlice([]Token{OpeningBracketToken, AnyPrimitiveTypeToken, ClosingBracketToken, AnyStandardTypeToken})
	StructToken = NewTokenListFromSlice([]Token{OpeningCurlyBracketToken, ClosingCurlyBracketToken})
)

func init() {
	AnyStandardTypeToken = NewVariantToken([]Token{StringToken, IntToken, FloatToken, BoolToken, SliceToken, MapToken, StructToken})
}

type Token interface {
	Matches(*rawsig.RawSignature) bool
}

type VariantToken []Token

func NewVariantToken(tokens []Token) *VariantToken {
	vt := VariantToken(tokens)
	return &vt
}

func (token *VariantToken) Matches(rawSig *rawsig.RawSignature) bool {
	tokens := []Token(*token)
	for _, token := range tokens {
		if token.Matches(rawSig) {
			return true
		}
	}
	return false
}

type KeyValuePairToken struct {
	key Token
	sep Token
	val Token
}

func NewKeyValuePairToken() *KeyValuePairToken {

}

type RuneToken rune

func NewRuneToken(r rune) *RuneToken {
	t := RuneToken(r)
	return &t
}

func (token *RuneToken) Matches(rawSig *rawsig.RawSignature) bool {
	runeT, err := rawSig.Next()
	if err != nil {
		return false
	}
	return runeT == rune(*token)
}

type TokenList struct {
	value    Token
	nextElem *TokenList
}

func NewTokenList(token Token) *TokenList {
	return &TokenList{token, nil}
}

func NewTokenListFromString(str string) *TokenList {
	var (
		runes     = []rune(str)
		tokenList *TokenList
	)
	for i := len(runes) - 1; i >= 0; i-- {
		tokenList = tokenList.Append(NewRuneToken(runes[i]))
	}
	return tokenList
}

func NewTokenListFromSlice(tokens []Token) *TokenList {
	var tokenList *TokenList
	for i := len(tokens) - 1; i >= 0; i-- {
		tokenList = tokenList.Append(tokens[i])
	}
	return tokenList
}

func (tokenList *TokenList) Append(token Token) *TokenList {
	return &TokenList{token, tokenList}
}

func (tokenList *TokenList) Next() *TokenList {
	return tokenList.nextElem
}

func (tokenList *TokenList) Matches(rawSig *rawsig.RawSignature) bool {
	var tl *TokenList
	for tl = tokenList.Next(); tl != nil; tl = tokenList.Next() {
		if !tl.Matches(rawSig) {
			return false
		}
	}
	return true
}
