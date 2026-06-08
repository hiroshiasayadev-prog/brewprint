package semantic

type TypeRefKind string

const (
	TypeRefPrimitive  TypeRefKind = "primitive"
	TypeRefNamedModel TypeRefKind = "named_model"
	TypeRefList       TypeRefKind = "list"
	TypeRefDict       TypeRefKind = "dict"
)

type TypeRef struct {
	Kind  TypeRefKind
	Raw   string
	Name  string
	Model QualifiedID
	Elem  *TypeRef
	Value *TypeRef
}

func (ref *TypeRef) String() string {
	if ref == nil {
		return ""
	}
	if ref.Raw != "" {
		return ref.Raw
	}
	return ref.Name
}
