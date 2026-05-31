package rawyaml

type Model struct {
	ID      string       `yaml:"id"`
	Type    string       `yaml:"type"`
	Kind    string       `yaml:"kind"`
	Fields  []ModelField `yaml:"fields"`
	Element string       `yaml:"element"`
	Value   string       `yaml:"value"`
	Values  []string     `yaml:"values"`
	Note    string       `yaml:"note"`
}

type ModelField struct {
	Name   string `yaml:"name"`
	Type   string `yaml:"type"`
	PK     bool   `yaml:"pk"`
	FK     string `yaml:"fk"`
	Unique bool   `yaml:"unique"`
	Note   string `yaml:"note"`
}
