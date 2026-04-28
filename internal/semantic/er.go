package semantic

type ERView struct {
	FileID  FileID
	ID      string
	Note    string
	Modules []ERViewModule
}

type ERViewModule struct {
	Module string
}
