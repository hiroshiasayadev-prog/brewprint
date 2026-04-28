package semantic

type SequenceScenario struct {
	FileID    FileID
	ID        string
	Title     string
	StateFile FileID
	Steps     []SequenceStep
}

type SequenceStep struct {
	FromState  string
	Via        string
	Guard      string
	Transition Transition
}
