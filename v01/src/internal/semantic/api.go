package semantic

type APIView struct {
	FileID       FileID
	ID           string
	Note         string
	HTTPRootPath string
	Modules      []APIViewModule
}

type APIViewModule struct {
	Module            string
	IncludeSubmodules bool
}
