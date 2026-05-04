package semantic

type FlowKind string

const (
	FlowKindStep    FlowKind = "step"
	FlowKindForeach FlowKind = "foreach"
	FlowKindFork    FlowKind = "fork"
	FlowKindBranch  FlowKind = "branch"
)

type FlowEntry struct {
	Kind    FlowKind
	Step    StepFlow
	Foreach ForeachFlow
	Fork    ForkFlow
	Branch  BranchFlow
}

type StepFlow struct {
	Task   QualifiedID
	TaskID string
	Params []ParamWiring
}

type ForeachFlow struct {
	Task    QualifiedID
	TaskID  string
	Over    FlowSource
	Mode    string
	Params  []ParamWiring
	Returns string
}

type FlowCollectedSource struct {
	Name              string
	FileID            FileID
	ProducedByForeach QualifiedID
	ProducedByTaskID  string
	TypeRef           *TypeRef
}

type ForkFlow struct {
	Fork       QualifiedID
	ForkID     string
	Branches   []ForkBranchFlow
	Join       QualifiedID
	JoinID     string
	JoinParams []ParamWiring
}

type ForkBranchFlow struct {
	Steps []StepFlow
}

type BranchFlow struct {
	Branch   QualifiedID
	BranchID string
	Params   []ParamWiring
	Cases    []BranchCaseFlow
}

type BranchCaseFlow struct {
	Label string
	Step  StepFlow
}

type ParamWiring struct {
	TargetParam string
	Source      FlowSource
}

type FlowSourceKind string

const (
	FlowSourceParam FlowSourceKind = "param"
	FlowSourceItem  FlowSourceKind = "item"
	FlowSourceNode  FlowSourceKind = "node"
)

type FlowSource struct {
	Kind      FlowSourceKind
	Raw       string
	ParamName string
	Node      QualifiedID
	AssetName string
}
