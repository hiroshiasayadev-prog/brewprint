package semantic

type FileID string

type QualifiedID string

func (id FileID) String() string {
	return string(id)
}

func (id QualifiedID) String() string {
	return string(id)
}
