package rawyaml

type Model struct {
	ID            string         `yaml:"id"`
	Type          string         `yaml:"type"`
	Main          bool           `yaml:"main"`
	Kind          string         `yaml:"kind"`
	Fields        []ModelField   `yaml:"fields"`
	Element       string         `yaml:"element"`
	Value         string         `yaml:"value"`
	Values        []string       `yaml:"values"`
	Discriminator string         `yaml:"discriminator"`
	Variants      []ModelVariant `yaml:"variants"`
	Note          string         `yaml:"note"`
}

type ModelField struct {
	Name   string `yaml:"name"`
	Type   string `yaml:"type"`
	PK     bool   `yaml:"pk"`
	FK     string `yaml:"fk"`
	Unique bool   `yaml:"unique"`
	Note   string `yaml:"note"`
}

// ModelVariant represents a single variant in a tagged_union model.
// Fields is nil when the "fields" key is absent from the YAML (error), and
// a non-nil empty slice when "fields: []" (payload-less variant).
type ModelVariant struct {
	Tag    string       `yaml:"tag"`
	Fields []ModelField `yaml:"fields"`
}
