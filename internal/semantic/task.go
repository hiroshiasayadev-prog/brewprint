package semantic

type Task struct {
	BaseNode
	Params      []Param
	Returns     *Return
	Reads       []StoreRef
	Writes      []StoreRef
	Initializes []InitializedStore
	Endpoint    bool
	Method      string
	Path        string
}

type Param struct {
	Name      string
	Model     QualifiedID
	ModelName string
	Note      string
}

type Return struct {
	Name      string
	Model     QualifiedID
	ModelName string
	Asset     *Asset
}

type InitializedStore struct {
	Name  string
	Model QualifiedID
	Note  string
	Store *Store
}

type StoreRef struct {
	Name        string
	Store       QualifiedID
	FilePrivate bool
}
