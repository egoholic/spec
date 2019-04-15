package parser

import (
	"github.com/egoholic/spec/rawsig"
	"github.com/egoholic/spec/sig"
)

// A parser is a functions from raw signature to signature.
type Parser func(rawsig.RawSignature) (sig.Signature, error)

// Parser combinator is responsible for organizing parsing functions into an extensible parsing system.
type ParserCombinator struct {
	parsers []Parser
}

func NewParserCombinator(parsers []Parser) *ParserCombinator {
	return &ParserCombinator{parsers}
}

func (parser *ParserCombinator) Parse() sig.Signature {
	return nil
}
