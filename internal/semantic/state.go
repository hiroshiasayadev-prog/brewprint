package semantic

type State struct {
	BaseNode
	Initial   bool
	Final     bool
	Wireframe *WireframeElement
}

type Event struct {
	BaseNode
	Source       string
	Actor        string
	PayloadModel QualifiedID
	PayloadName  string
	Watches      QualifiedID
	WatchesName  string
}

type Transition struct {
	FileID     FileID
	From       string
	On         string
	To         string
	Action     string
	Guard      string
	Note       string
	FromState  QualifiedID
	Event      QualifiedID
	ToState    QualifiedID
	ActionTask QualifiedID
}
