package semantic

type Model struct {
	BaseNode
	FilePrivate bool
	LocalName   string
	Kind        string
	Fields      []ModelField
	Element     string
	ElementRef  *TypeRef
	Value       string
	ValueRef    *TypeRef
	Values      []string
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
