package semantic

type Project struct {
	NodesByQID     map[QualifiedID]Node
	NodesByFile    map[FileID][]Node
	MainNodeByFile map[FileID]QualifiedID

	TasksByQID    map[QualifiedID]*Task
	ModelsByQID   map[QualifiedID]*Model
	BranchesByQID map[QualifiedID]*Branch
	ForksByQID    map[QualifiedID]*Fork
	JoinsByQID    map[QualifiedID]*Join

	StoresByQID       map[QualifiedID]*Store
	StoresByFileLocal map[FileID]map[string]*Store

	ActorsByQID map[QualifiedID]*Actor

	TasksReadingStore map[QualifiedID][]QualifiedID
	TasksWritingStore map[QualifiedID][]QualifiedID

	FlowByFile map[FileID][]FlowEntry
}

func NewProject() *Project {
	return &Project{
		NodesByQID:        map[QualifiedID]Node{},
		NodesByFile:       map[FileID][]Node{},
		MainNodeByFile:    map[FileID]QualifiedID{},
		TasksByQID:        map[QualifiedID]*Task{},
		ModelsByQID:       map[QualifiedID]*Model{},
		BranchesByQID:     map[QualifiedID]*Branch{},
		ForksByQID:        map[QualifiedID]*Fork{},
		JoinsByQID:        map[QualifiedID]*Join{},
		StoresByQID:       map[QualifiedID]*Store{},
		StoresByFileLocal: map[FileID]map[string]*Store{},
		ActorsByQID:       map[QualifiedID]*Actor{},
		TasksReadingStore: map[QualifiedID][]QualifiedID{},
		TasksWritingStore: map[QualifiedID][]QualifiedID{},
		FlowByFile:        map[FileID][]FlowEntry{},
	}
}
