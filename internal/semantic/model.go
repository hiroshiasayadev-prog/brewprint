package semantic

type Model struct {
	BaseNode
	Kind    string
	Fields  []ModelField
	Element string
	Value   string
}

type ModelField struct {
	Name   string
	Type   string
	PK     bool
	FK     string
	Unique bool
	Note   string
}
