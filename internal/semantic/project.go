package semantic

type Project struct {
	NodesByQID     map[QualifiedID]Node
	NodesByFile    map[FileID][]Node
	MainNodeByFile map[FileID]QualifiedID

	TasksByQID    map[QualifiedID]*Task
	ModelsByQID   map[QualifiedID]*Model
	StatesByQID   map[QualifiedID]*State
	EventsByQID   map[QualifiedID]*Event
	BranchesByQID map[QualifiedID]*Branch
	ForksByQID    map[QualifiedID]*Fork
	JoinsByQID    map[QualifiedID]*Join

	StoresByQID       map[QualifiedID]*Store
	StoresByFileLocal map[FileID]map[string]*Store

	ActorsByQID map[QualifiedID]*Actor

	TasksReadingStore map[QualifiedID][]QualifiedID
	TasksWritingStore map[QualifiedID][]QualifiedID

	ReferencesBySource map[ObjectKey][]Reference
	ReferencesByTarget map[ObjectKey][]Reference

	FlowByFile        map[FileID][]FlowEntry
	TransitionsByFile map[FileID][]Transition

	TransitionsByStateEventGuard map[TransitionKey]TransitionRef
	TransitionsByStateEvent      map[TransitionEventKey][]TransitionRef
	ActionsByTask                map[QualifiedID][]TransitionRef

	ScenariosByID map[string]*SequenceScenario
	ERViewsByID   map[string]*ERView
	APIViewsByID  map[string]*APIView
}

func NewProject() *Project {
	return &Project{
		NodesByQID:         map[QualifiedID]Node{},
		NodesByFile:        map[FileID][]Node{},
		MainNodeByFile:     map[FileID]QualifiedID{},
		TasksByQID:         map[QualifiedID]*Task{},
		ModelsByQID:        map[QualifiedID]*Model{},
		StatesByQID:        map[QualifiedID]*State{},
		EventsByQID:        map[QualifiedID]*Event{},
		BranchesByQID:      map[QualifiedID]*Branch{},
		ForksByQID:         map[QualifiedID]*Fork{},
		JoinsByQID:         map[QualifiedID]*Join{},
		StoresByQID:        map[QualifiedID]*Store{},
		StoresByFileLocal:  map[FileID]map[string]*Store{},
		ActorsByQID:        map[QualifiedID]*Actor{},
		TasksReadingStore:  map[QualifiedID][]QualifiedID{},
		TasksWritingStore:  map[QualifiedID][]QualifiedID{},
		ReferencesBySource: map[ObjectKey][]Reference{},
		ReferencesByTarget: map[ObjectKey][]Reference{},
		FlowByFile:                    map[FileID][]FlowEntry{},
		TransitionsByFile:             map[FileID][]Transition{},
		TransitionsByStateEventGuard:  map[TransitionKey]TransitionRef{},
		TransitionsByStateEvent:       map[TransitionEventKey][]TransitionRef{},
		ActionsByTask:                 map[QualifiedID][]TransitionRef{},
		ScenariosByID:                 map[string]*SequenceScenario{},
		ERViewsByID:        map[string]*ERView{},
		APIViewsByID:       map[string]*APIView{},
	}
}
