package types

type Kind uint8

const (
	BoolKind Kind = iota
	NumberKind
	StringKind
	BlobKind
	ValueKind
	ListKind
	MapKind
	SetKind

	StructKind

	TypeKind
)

var KindToString = map[Kind]string{
	BoolKind:   "bool",
	NumberKind: "number",
	StringKind: "string",
	BlobKind:   "blob",
	ValueKind:  "value",
	ListKind:   "list",
	MapKind:    "map",
	SetKind:    "set",
	StructKind: "struct",
	TypeKind:   "type",
}

// toString return the name of kind.
func (k Kind) toString() string {
	return KindToString[k]
}
