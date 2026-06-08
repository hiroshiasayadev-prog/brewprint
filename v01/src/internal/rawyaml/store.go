package rawyaml

type Store struct {
	ID   string `yaml:"id"`
	Type string `yaml:"type"`
	Kind string `yaml:"kind"`
	Of   string `yaml:"of"`
	Note string `yaml:"note"`
}
