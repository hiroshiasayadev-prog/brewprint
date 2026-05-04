package semantic

type Model struct {
	BaseNode
	Kind       string
	Fields     []ModelField
	Element    string
	ElementRef *TypeRef
	Value      string
	ValueRef   *TypeRef
}

type ModelField struct {
	Name    string
	Type    string
	TypeRef *TypeRef
	PK      bool
	FK      string
	Unique  bool
	Note    string
}
