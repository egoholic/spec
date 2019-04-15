package rawsig

type RawSignature struct {
	runes  []rune
	cursor int
}

func New(rawSig string) *RawSignature {
	return &RawSignature{[]rune(rawSig), 0}
}
