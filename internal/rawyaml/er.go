package rawyaml

type ERView struct {
	As      string         `yaml:"as"`
	ID      string         `yaml:"id"`
	Note    string         `yaml:"note"`
	Modules []ERViewModule `yaml:"modules"`
}

type ERViewModule struct {
	Module string `yaml:"module"`
}
