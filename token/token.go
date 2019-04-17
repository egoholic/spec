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

	AnyStandardTypeToken *variantToken

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
type variantToken struct {
	variants []Token
}

func NewVariantToken(tokens []Token) *variantToken {
	return &variantToken{tokens}
}

func NewVariantTokenFromRunes(runes []rune) *variantToken {
	var tokens []Token
	for _, r := range runes {
		tokens = append(tokens, NewRuneToken(r))
	}

	return &variantToken{tokens}
}

func (token *variantToken) Join(token2 *variantToken) *variantToken {
	var joined = make([]Token, 10)
	for _, t := range token.variants {
		joined = append(joined, t)
	}

	for _, t := range token2.variants {
		joined = append(joined, t)
	}

	return &variantToken{joined}
}

func (token *variantToken) Matches(rawSig *rawsig.RawSignature) bool {
	for _, token := range token.variants {
		if token.Matches(rawSig) {
			return true
		}
	}
	return false
}

type keyValuePairToken struct {
	key Token
	sep Token
	val Token
}

func NewKeyValuePairToken(key, sep, val Token) *keyValuePairToken {
	return &keyValuePairToken{key, sep, val}
}

func (token *keyValuePairToken) Matches(rawSig *rawsig.RawSignature) bool {
	return token.key.Matches(rawSig) && token.sep.Matches(rawSig) && token.val.Matches(rawSig)
}

func (token *keyValuePairToken) MatchesWithoutKey(rawSig *rawsig.RawSignature) bool {
	return token.sep.Matches(rawSig) && token.val.Matches(rawSig)
}

type keyValuePairsSetToken struct {
	pairs map[Token]*keyValuePairToken
}

func NewKeyValuePairsSetToken(pairs map[Token]*keyValuePairToken) *keyValuePairsSetToken {
	return &keyValuePairsSetToken{pairs}
}

// This naive implementation assumes that key-value pairs are ordered and order matters during comparing.
// So that:
//     a:1, b:2 and b:2, a:1 are different tokens.
// Also this implementation does not support subset match.
// So that:
//     a:1 and a:1, b:2 could not match.
//
func (token *keyValuePairsSetToken) Matches(rawSig *rawsig.RawSignature) bool {
	for _, pairToken := range token.pairs {
		if !pairToken.Matches(rawSig) {
			return false
		}
	}
	return true
}

type anyRuneToken struct{}

func NewAnyRuneToken() *anyRuneToken {
	return &anyRuneToken{}
}

func (token *anyRuneToken) Matches(rawSig *rawsig.RawSignature) bool {
	_, err := rawSig.Next()
	return err == nil
}

type runeToken struct {
	r rune
}

func NewRuneToken(r rune) *runeToken {
	return &runeToken{r}
}

func (token *runeToken) Matches(rawSig *rawsig.RawSignature) bool {
	r, err := rawSig.Next()
	if err != nil {
		return false
	}
	return r == token.r
}

type wordToken struct {
	prefix Token
	root   Token
	suffix Token
}

func NewWordToken(prefix, root, suffix Token) *wordToken {
	return &wordToken{prefix, root, suffix}
}

type tokenList struct {
	value    Token
	nextElem *tokenList
}

func NewTokenList(token Token) *tokenList {
	return &tokenList{token, nil}
}

func NewTokenListFromString(str string) *tokenList {
	var (
		runes     = []rune(str)
		tokenList *tokenList
	)
	for i := len(runes) - 1; i >= 0; i-- {
		tokenList = tokenList.Append(NewRuneToken(runes[i]))
	}
	return tokenList
}

func NewTokenListFromSlice(tokens []Token) *tokenList {
	var tokenList *tokenList
	for i := len(tokens) - 1; i >= 0; i-- {
		tokenList = tokenList.Append(tokens[i])
	}
	return tokenList
}

func (tailToken *tokenList) Append(headToken Token) *tokenList {
	return &tokenList{headToken, tailToken}
}

func (tokenList *tokenList) Next() *tokenList {
	return tokenList.nextElem
}

func (token *tokenList) Matches(rawSig *rawsig.RawSignature) bool {
	var tl *tokenList
	for tl = token.Next(); tl != nil; tl = token.Next() {
		if !tl.Matches(rawSig) {
			return false
		}
	}
	return true
}
