package sig

type Signature interface {
}

type EndpointSignature struct {
	accessibleVia string
}
type StructSignature struct {
}
type MapSignature struct {
}
type SliceSignature struct {
}
type PrimitiveSignature struct {
	title string
}
