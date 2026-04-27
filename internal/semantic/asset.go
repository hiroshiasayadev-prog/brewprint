package semantic

type Asset struct {
	Name       string
	Model      QualifiedID
	ModelName  string
	ProducedBy QualifiedID
	FileID     FileID
}
