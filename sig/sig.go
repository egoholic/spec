package sig

import "fmt"

// Every concrete signature type should correspond the interface above.
type Signature interface {
	Token() string
	Title() string
}

// Signature of a whole endpoint.
type EndpointSignature struct {
	token     string
	inputSig  Signature
	outputSig Signature
}

func NewEndpointSignature(token string, inputSig Signature, outputSig Signature) *EndpointSignature {
	return &EndpointSignature{token, inputSig, outputSig}
}
func (sig *EndpointSignature) Token() string {
	return sig.token
}

// Signature of a struct data type. Tuples and records could be represented as structs as well.
type StructSignature struct {
	token  string
	fields map[string]Signature
}

func NewStructSignature(token string, fields map[string]Signature) *StructSignature {
	return &StructSignature{token, fields}
}
func (sig *StructSignature) Token() string {
	return sig.token
}

// Signature of map (associative array, dictionary, etc).
type MapSignature struct {
	token    string
	keySig   Signature
	valueSig Signature
}

func NewMapSignature(token string, keySig Signature, valueSig Signature) *MapSignature {
	return &MapSignature{token, keySig, valueSig}
}
func (sig *MapSignature) Token() string {
	return sig.token
}
func (sig *MapSignature) Title() string {
	return ""
}

// Signature of slice (array, list, as well).
type SliceSignature struct {
	token    string
	valueSig Signature
}

func NewSliceSignature(token string, valueSig Signature) *SliceSignature {
	return &SliceSignature{token, valueSig}
}
func (sig *SliceSignature) Token() string {
	return sig.token
}

func (sig *SliceSignature) Title() string {
	return fmt.Sprintf("%s%s", sig.Token(), sig.valueSig.Title())
}

// Signature of primitives (string, int, float, bool, etc.)
type PrimitiveSignature struct {
	token string
}

func NewPrimitiveSignature(token string) *PrimitiveSignature {
	return &PrimitiveSignature{token}
}

func (sig *PrimitiveSignature) Token() string {
	return sig.token
}

func (sig *PrimitiveSignature) Title() string {
	return sig.token
}
