package semantic

type Model struct {
	BaseNode
	FilePrivate   bool
	LocalName     string
	Kind          string
	Fields        []ModelField
	Element       string
	ElementRef    *TypeRef
	Value         string
	ValueRef      *TypeRef
	Values        []string
	Discriminator string
	Variants      []ModelVariant
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

// ModelVariant represents a single variant in a tagged_union model.
// Fields is nil when the "fields" key was absent from the source YAML.
// A non-nil empty slice means the variant has no payload fields (valid).
type ModelVariant struct {
	Tag    string
	Fields []ModelField
}
