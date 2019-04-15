package parser

import "github.com/egoholic/spec/sig"

type Parser struct{}

func New() *Parser {
	return &Parser{}
}

func (parser *Parser) Parse() sig.Signature {
	return nil
}
