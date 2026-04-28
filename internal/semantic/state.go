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

type TransitionKey struct {
	StateFile FileID
	FromState QualifiedID
	Event     QualifiedID
	Guard     string
}

type TransitionEventKey struct {
	StateFile FileID
	FromState QualifiedID
	Event     QualifiedID
}

type TransitionRef struct {
	Key        TransitionKey
	Transition Transition
}
