package rawsig

import "errors"

type RawSignature struct {
	rawSig         string
	runes          []rune
	length         int
	matchingCursor int
	matchedCursor  int
}

func New(rawSig string) *RawSignature {
	runes := []rune(rawSig)
	return &RawSignature{rawSig, runes, len(runes), 0, 0}
}

func (rawSig *RawSignature) String() string {
	return rawSig.rawSig
}

func (rawSig *RawSignature) HasNext() bool {
	return rawSig.matchingCursor < rawSig.length
}

func (rawSig *RawSignature) Next() (r rune, err error) {
	if !rawSig.HasNext() {
		err = errors.New("there is no next rune")
		return
	}
	r = rawSig.runes[rawSig.matchingCursor]
	rawSig.matchingCursor++

	return
}

func (rawSig *RawSignature) Reset() {
	rawSig.matchingCursor = rawSig.matchedCursor
}

func (rawSig *RawSignature) Commit() (runes []rune, ok bool) {
	runes = rawSig.runes[rawSig.matchedCursor:rawSig.matchingCursor]
	ok = rawSig.matchedCursor < rawSig.matchingCursor
	rawSig.matchedCursor = rawSig.matchingCursor
	return
}
