package rawyaml

type FileKind string

const (
	FileKindNode        FileKind = "node"
	FileKindView        FileKind = "view"
	FileKindRenderIndex FileKind = "render_index"
	FileKindUnsupported FileKind = "unsupported"
)

type Project struct {
	Root  string
	Files []File
}

type File struct {
	ID               string
	Path             string
	Content          string
	Kind             FileKind
	ViewAs           string
	NodeFile         *NodeFile
	RenderIndex      *RenderIndex
	SequenceScenario *SequenceScenario
	ERView           *ERView
	APIView          *APIView
}

type NodeFile struct {
	Tasks       []Task
	Models      []Model
	Stores      []Store
	Actors      []Actor
	States      []State
	Events      []Event
	Branches    []ControlNode
	Forks       []ControlNode
	Joins       []ControlNode
	Flow        []FlowEntry
	Transitions []Transition
	Unsupported []CommonNode
}

type CommonNode struct {
	ID   string `yaml:"id"`
	Type string `yaml:"type"`
	Main bool   `yaml:"main"`
	Note string `yaml:"note"`
}
