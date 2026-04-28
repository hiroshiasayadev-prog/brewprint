package rawyaml

type State struct {
	ID        string            `yaml:"id"`
	Type      string            `yaml:"type"`
	Initial   bool              `yaml:"initial"`
	Final     bool              `yaml:"final"`
	Note      string            `yaml:"note"`
	Wireframe *WireframeElement `yaml:"wireframe"`
}

type Event struct {
	ID      string        `yaml:"id"`
	Type    string        `yaml:"type"`
	Source  string        `yaml:"source"`
	Actor   string        `yaml:"actor"`
	Payload *EventPayload `yaml:"payload"`
	Watches string        `yaml:"watches"`
	Note    string        `yaml:"note"`
}

type EventPayload struct {
	Model string `yaml:"model"`
}

type Transition struct {
	From   string `yaml:"from"`
	On     string `yaml:"on"`
	To     string `yaml:"to"`
	Action string `yaml:"action"`
	Guard  string `yaml:"guard"`
	Note   string `yaml:"note"`
}
