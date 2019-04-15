package parser

import (
	"errors"

	"github.com/egoholic/spec/rawsig"
	"github.com/egoholic/spec/sig"
)

var (
	INT_TOKEN        = []rune("int")
	STRING_TOKEN     = []rune("string")
	FLOAT_TOKEN      = []rune("float")
	BOOL_TOKEN       = []rune("bool")
	PRIMITIVE_TOKENS = [][]rune{INT_TOKEN, STRING_TOKEN, FLOAT_TOKEN, BOOL_TOKEN}
)

func ParsePrimitive(rawSig *rawsig.RawSignature) (signature sig.Signature, err error) {
	var (
		matchedRunes []rune
		ok           bool
	)

	for _, token := range PRIMITIVE_TOKENS {
		for _, tr := range token {
			rsr, err := rawSig.Next()
			if err != nil {
				rawSig.Reset()
				break
			}
			if rsr != tr {
				rawSig.Reset()
				break
			}
		}
		matchedRunes, ok = rawSig.Commit()
		if ok {
			signature = sig.NewPrimitiveSignature(string(matchedRunes))
			return
		}
	}
	err = errors.New("not a primitive signature")
	return
}
