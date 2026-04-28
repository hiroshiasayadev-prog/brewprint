package rawyaml

type APIView struct {
	As           string          `yaml:"as"`
	ID           string          `yaml:"id"`
	Note         string          `yaml:"note"`
	HTTPRootPath string          `yaml:"http_root_path"`
	Modules      []APIViewModule `yaml:"modules"`
}

type APIViewModule struct {
	Module            string `yaml:"module"`
	IncludeSubmodules bool   `yaml:"include_submodules"`
}
