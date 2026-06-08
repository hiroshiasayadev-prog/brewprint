package rawyaml

type RenderIndex struct {
	Groups []RenderGroup `yaml:"groups"`
}

type RenderGroup struct {
	ID      string   `yaml:"id"`
	Label   string   `yaml:"label"`
	Modules []string `yaml:"modules"`
}
