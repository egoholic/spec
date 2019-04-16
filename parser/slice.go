package parser

import (
	"errors"

	"github.com/egoholic/spec/rawsig"
	"github.com/egoholic/spec/sig"
)

var (
	SLICE_TOKEN = []rune("[]")
)

func ParseSlice(rawSig *rawsig.RawSignature, parserCombinator *ParserCombinator) (signature sig.Signature, err error) {
	var tr, rsr rune
	for _, tr = range SLICE_TOKEN {
		rsr, err = rawSig.Next()
		if err != nil {
			rawSig.Reset()
			return
		}
		if tr != rsr {
			rawSig.Reset()
			err = errors.New("not a slice signature")
			return
		}
	}
	matchedRunes, ok := rawSig.Commit()
	if !ok {
		rawSig.Reset()
		err = errors.New("not a slice signature")
		return
	}

	nestedSignature, err := parserCombinator.Parse(rawSig)
	if err != nil {
		rawSig.Reset()
		return
	}
	signature = sig.NewSliceSignature(string(matchedRunes), nestedSignature)
	return
}
