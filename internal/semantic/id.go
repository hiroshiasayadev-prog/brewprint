package semantic

type FileID string

type QualifiedID string

func (id FileID) String() string {
	return string(id)
}

func (id QualifiedID) String() string {
	return string(id)
}

func PrivateNodeID(fileID FileID, localID string) string {
	if fileID == "" || localID == "" {
		return ""
	}
	return fileID.String() + "#" + localID
}

func AssetID(producer QualifiedID, name string) string {
	if producer == "" || name == "" {
		return ""
	}
	return producer.String() + "#" + name
}
