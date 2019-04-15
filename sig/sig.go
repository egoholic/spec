package sig

// Every concrete signature type should correspond the interface above.
type Signature interface {
	Title() string
}

// Signature of a whole endpoint.
type EndpointSignature struct {
	title     string
	inputSig  Signature
	outputSig Signature
}

func (sig *EndpointSignature) Title() string {
	return sig.title
}

// Signature of a struct data type. Tuples and records could be represented as structs as well.
type StructSignature struct {
	title  string
	fields map[string]Signature
}

func (sig *StructSignature) Title() string {
	return sig.title
}

// Signature of map (associative array, dictionary, etc).
type MapSignature struct {
	title    string
	keySig   Signature
	valueSig Signature
}

func (sig *MapSignature) Title() string {
	return sig.title
}

// Signature of slice (array, list, as well).
type SliceSignature struct {
	title    string
	valueSig Signature
}

func (sig *SliceSignature) Title() string {
	return sig.title
}

// Signature of primitives (string, int, float, bool, etc.)
type PrimitiveSignature struct {
	title string
}

func NewPrimitiveSignature(title string) *PrimitiveSignature {
	return &PrimitiveSignature{title}
}

func (sig *PrimitiveSignature) Title() string {
	return sig.title
}
