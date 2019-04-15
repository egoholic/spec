package parser

import (
	"fmt"

	"github.com/egoholic/spec/rawsig"
	"github.com/egoholic/spec/sig"
)

// A parser is a functions from raw signature to signature.
type Parser func(*rawsig.RawSignature, *ParserCombinator) (sig.Signature, error)

// Parser combinator is responsible for organizing parsing functions into an extensible parsing system.
type ParserCombinator struct {
	parsers []Parser
}

func NewParserCombinator(parsers []Parser) *ParserCombinator {
	return &ParserCombinator{parsers}
}

func (parserCombinator *ParserCombinator) Parse(rawSig *rawsig.RawSignature) (signature sig.Signature, err error) {
	for _, parser := range parserCombinator.parsers {
		signature, err = parser(rawSig, parserCombinator)
		if err != nil {
			break
		}
		return
	}
	err = fmt.Errorf("can't parse signature `%s`", rawSig.String())
	return
}
