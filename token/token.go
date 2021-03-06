package token

import (
	"strings"

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
	LatinLetterToken         = LowLatinLetterToken.Join(HighLatinLetterToken)
	DigitToken               = NewVariantTokenFromRunes([]rune("0123456789"))
	DigitOrLatinLetterToken  = LatinLetterToken.Join(DigitToken)
	StringToken              = NewTokenListFromString("string")
	IntToken                 = NewTokenListFromString("int")
	FloatToken               = NewTokenListFromString("float")
	BoolToken                = NewTokenListFromString("bool")
	AnyPrimitiveTypeToken    = NewVariantToken([]Token{StringToken, IntToken, FloatToken, BoolToken})

	// Words could start with ONLY a low or high latin letter.
	WordToken            = NewWordToken(LatinLetterToken, DigitOrLatinLetterToken, nil)
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
	String() string
}

type multyRuneToken struct {
	runes []rune
}

func NewMultyRuneToken(runes []rune) *multyRuneToken {
	return &multyRuneToken{runes}
}

func (token *multyRuneToken) Matches(rawSig *rawsig.RawSignature) bool {
	for _, r := range token.runes {
		rsr, err := rawSig.Next()
		if err != nil {
			return false
		}
		if rsr != r {
			return false
		}
	}
	return true
}

func (token *multyRuneToken) String() string {
	return string(token.runes)
}

type variantToken struct {
	options []Token
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
	joined := []Token{}

	for i := len(token2.options) - 1; i >= 0; i-- {
		joined = append(joined, token2.options[i])
	}
	for i := len(token.options) - 1; i >= 0; i-- {
		joined = append(joined, token.options[i])
	}

	return &variantToken{joined}
}

func (token *variantToken) Matches(rawSig *rawsig.RawSignature) bool {
	for _, tokenOption := range token.options {
		if tokenOption == nil {
			return false
		}
		rawSig.Reset()
		if tokenOption.Matches(rawSig) {
			return true
		}
	}
	return false
}

func (token *variantToken) String() string {
	builder := &strings.Builder{}
	builder.WriteString("(one-of ")
	for _, tokenOption := range token.options {
		builder.WriteString(tokenOption.String())
		builder.WriteRune(' ')
	}
	builder.WriteRune(')')
	return builder.String()
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

func (token *keyValuePairToken) String() string {
	builder := &strings.Builder{}
	builder.WriteString(token.key.String())
	builder.WriteString(token.sep.String())
	builder.WriteString(token.val.String())
	return builder.String()
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

func (token *keyValuePairsSetToken) String() string {
	builder := &strings.Builder{}
	for _, pairToken := range token.pairs {
		builder.WriteString(pairToken.String())
	}
	return builder.String()
}

type anyRuneToken struct{}

func NewAnyRuneToken() *anyRuneToken {
	return &anyRuneToken{}
}

func (token *anyRuneToken) Matches(rawSig *rawsig.RawSignature) bool {
	_, err := rawSig.Next()
	return err == nil
}

func (token *anyRuneToken) String() string {
	return "<any-rune>"
}

type runeToken struct {
	r rune
}

func NewRuneToken(r rune) *runeToken {
	return &runeToken{r}
}

func (token *runeToken) String() string {
	return string(token.r)
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

func (token *wordToken) Matches(rawSig *rawsig.RawSignature) bool {
	if token.prefix != nil {
		if !token.prefix.Matches(rawSig) {
			return false
		}
	}

	if !token.root.Matches(rawSig) {
		return false
	}

	if token.suffix != nil {
		if !token.suffix.Matches(rawSig) {
			return false
		}
	}

	return true
}

func (token *wordToken) String() string {
	builder := &strings.Builder{}
	if token.prefix != nil {
		builder.WriteString(token.prefix.String())
	}
	builder.WriteString(token.root.String())
	if token.suffix != nil {
		builder.WriteString(token.suffix.String())
	}

	return builder.String()
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

func (token *tokenList) Next() *tokenList {
	return token.nextElem
}

func (token *tokenList) String() string {
	builder := &strings.Builder{}
	builder.WriteRune('(')
	builder.WriteString(token.value.String())
	builder.WriteString(token.nextElem.String())
	builder.WriteRune(')')
	return builder.String()
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
