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
	LowLatinLetterToken      = NewVariantTokenFromRunes([]rune("abcdefghijklmnopqrstuvwxyz"))
	HighLatinLetterToken     = NewVariantTokenFromRunes([]rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ"))
	DigitToken               = NewVariantTokenFromRunes([]rune("0123456789"))
	StringToken              = NewTokenListFromString("string")
	IntToken                 = NewTokenListFromString("int")
	FloatToken               = NewTokenListFromString("float")
	BoolToken                = NewTokenListFromString("bool")
	AnyPrimitiveTypeToken    = NewVariantToken([]Token{StringToken, IntToken, FloatToken, BoolToken})

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

type ExceptToken struct {
	token Token
}

func NewExceptToken(token Token) *ExceptToken {
	return &ExceptToken{token}
}

func (token *ExceptToken) Matches(rawSig *rawsig.RawSignature) bool {
	return !token.token.Matches(rawSig)
}

type VariantToken struct {
	variants []Token
}

func NewVariantToken(tokens []Token) *VariantToken {
	return &VariantToken{tokens}
}

func NewVariantTokenFromRunes(runes []rune) *VariantToken {
	var tokens []Token
	for _, r := range runes {
		tokens = append(tokens, NewRuneToken(r))
	}

	return &VariantToken{tokens}
}

func (token *VariantToken) Join(token2 *VariantToken) *VariantToken {
	var joined = make([]Token, 10)
	for _, t := range token.variants {
		joined = append(joined, t)
	}

	for _, t := range token2.variants {
		joined = append(joined, t)
	}

	return &VariantToken{joined}
}

func (token *VariantToken) Matches(rawSig *rawsig.RawSignature) bool {
	for _, token := range token.variants {
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

func NewKeyValuePairToken(key, sep, val Token) *KeyValuePairToken {
	return &KeyValuePairToken{key, sep, val}
}

func (token *KeyValuePairToken) Matches(rawSig *rawsig.RawSignature) bool {
	return token.key.Matches(rawSig) && token.sep.Matches(rawSig) && token.val.Matches(rawSig)
}

func (token *KeyValuePairToken) MatchesWithoutKey(rawSig *rawsig.RawSignature) bool {
	return token.sep.Matches(rawSig) && token.val.Matches(rawSig)
}

type KeyValuePairsSetToken struct {
	pairs map[Token]*KeyValuePairToken
}

func NewKeyValuePairsSetToken(pairs map[Token]*KeyValuePairToken) *KeyValuePairsSetToken {
	return &KeyValuePairsSetToken{pairs}
}

// This naive implementation assumes that key-value pairs are ordered and order matters during comparing.
// So that:
//     a:1, b:2 and b:2, a:1 are different tokens.
// Also this implementation does not support subset match.
// So that:
//     a:1 and a:1, b:2 could not match.
//
func (token *KeyValuePairsSetToken) Matches(rawSig *rawsig.RawSignature) bool {
	for _, pairToken := range token.pairs {
		if !pairToken.Matches(rawSig) {
			return false
		}
	}
	return true
}

type AnyRuneToken struct{}

func NewAnyRuneToken() *AnyRuneToken {
	return &AnyRuneToken{}
}

func (token *AnyRuneToken) Matches(rawSig *rawsig.RawSignature) bool {
	_, err := rawSig.Next()
	return err == nil
}

type RuneToken struct {
	r rune
}

func NewRuneToken(r rune) *RuneToken {
	return &RuneToken{r}
}

func (token *RuneToken) Matches(rawSig *rawsig.RawSignature) bool {
	r, err := rawSig.Next()
	if err != nil {
		return false
	}
	return r == token.r
}

type WordToken struct {
	tokenList *TokenList
}

func NewWordToken(word string) *WordToken {
	runes := []rune(word)

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
