package rawyaml

type ControlNode struct {
	ID      string  `yaml:"id"`
	Type    string  `yaml:"type"`
	Main    bool    `yaml:"main"`
	Note    string  `yaml:"note"`
	Params  []Param `yaml:"params"`
	Returns *Return `yaml:"returns"`
}

type FlowEntry struct {
	Step     string            `yaml:"step"`
	Params   map[string]string `yaml:"params"`
	Foreach  string            `yaml:"foreach"`
	Over     string            `yaml:"over"`
	Mode     string            `yaml:"mode"`
	Returns  string            `yaml:"returns"`
	Fork     string            `yaml:"fork"`
	Branches []ForkBranch      `yaml:"branches"`
	Join     string            `yaml:"join"`
	Branch   string            `yaml:"branch"`
	Cases    []BranchCase      `yaml:"cases"`
}

type ForkBranch struct {
	Steps []FlowStep `yaml:"steps"`
}

type FlowStep struct {
	Step   string            `yaml:"step"`
	Params map[string]string `yaml:"params"`
}

type BranchCase struct {
	Label  string            `yaml:"label"`
	Step   string            `yaml:"step"`
	Params map[string]string `yaml:"params"`
}
