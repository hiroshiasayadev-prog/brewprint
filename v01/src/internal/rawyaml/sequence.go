package rawyaml

type SequenceScenario struct {
	As        string         `yaml:"as"`
	ID        string         `yaml:"id"`
	Title     string         `yaml:"title"`
	StateFile string         `yaml:"state_file"`
	Steps     []SequenceStep `yaml:"steps"`
}

type SequenceStep struct {
	FromState string `yaml:"from_state"`
	Via       string `yaml:"via"`
	Guard     string `yaml:"guard"`
}
