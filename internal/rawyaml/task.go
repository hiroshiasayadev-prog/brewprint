package rawyaml

type Task struct {
	ID          string       `yaml:"id"`
	Type        string       `yaml:"type"`
	Main        bool         `yaml:"main"`
	Note        string       `yaml:"note"`
	Params      []Param      `yaml:"params"`
	Returns     *Return      `yaml:"returns"`
	Reads       []string     `yaml:"reads"`
	Writes      []string     `yaml:"writes"`
	Initializes []Initialize `yaml:"initializes"`
	Endpoint    bool         `yaml:"endpoint"`
	Method      string       `yaml:"method"`
	Path        string       `yaml:"path"`
}

type Param struct {
	Name  string `yaml:"name"`
	Model string `yaml:"model"`
	Note  string `yaml:"note"`
}

type Return struct {
	Name   string `yaml:"name"`
	Model  string `yaml:"model"`
	Source string `yaml:"source"`
}

type Initialize struct {
	Name  string `yaml:"name"`
	Model string `yaml:"model"`
	Note  string `yaml:"note"`
}
