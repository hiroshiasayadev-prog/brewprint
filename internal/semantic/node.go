package semantic

type NodeKind string

const (
	NodeKindTask   NodeKind = "task"
	NodeKindModel  NodeKind = "model"
	NodeKindStore  NodeKind = "store"
	NodeKindActor  NodeKind = "actor"
	NodeKindState  NodeKind = "state"
	NodeKindEvent  NodeKind = "event"
	NodeKindBranch NodeKind = "branch"
	NodeKindFork   NodeKind = "fork"
	NodeKindJoin   NodeKind = "join"
)

type Node interface {
	GetQID() QualifiedID
	GetFileID() FileID
	GetID() string
	GetKind() NodeKind
	IsMain() bool
}

type BaseNode struct {
	QID    QualifiedID
	FileID FileID
	ID     string
	Kind   NodeKind
	Main   bool
	Note   string
}

func (n BaseNode) GetQID() QualifiedID { return n.QID }
func (n BaseNode) GetFileID() FileID   { return n.FileID }
func (n BaseNode) GetID() string       { return n.ID }
func (n BaseNode) GetKind() NodeKind   { return n.Kind }
func (n BaseNode) IsMain() bool        { return n.Main }
