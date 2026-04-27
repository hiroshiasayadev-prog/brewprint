package semantic

type Project struct {
	NodesByQID     map[QualifiedID]Node
	NodesByFile    map[FileID][]Node
	MainNodeByFile map[FileID]QualifiedID

	TasksByQID  map[QualifiedID]*Task
	ModelsByQID map[QualifiedID]*Model

	StoresByQID       map[QualifiedID]*Store
	StoresByFileLocal map[FileID]map[string]*Store

	ActorsByQID map[QualifiedID]*Actor

	TasksReadingStore map[QualifiedID][]QualifiedID
	TasksWritingStore map[QualifiedID][]QualifiedID
}

func NewProject() *Project {
	return &Project{
		NodesByQID:        map[QualifiedID]Node{},
		NodesByFile:       map[FileID][]Node{},
		MainNodeByFile:    map[FileID]QualifiedID{},
		TasksByQID:        map[QualifiedID]*Task{},
		ModelsByQID:       map[QualifiedID]*Model{},
		StoresByQID:       map[QualifiedID]*Store{},
		StoresByFileLocal: map[FileID]map[string]*Store{},
		ActorsByQID:       map[QualifiedID]*Actor{},
		TasksReadingStore: map[QualifiedID][]QualifiedID{},
		TasksWritingStore: map[QualifiedID][]QualifiedID{},
	}
}
