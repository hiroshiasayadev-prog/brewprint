package semantic

type Store struct {
	BaseNode
	StoreKind   string
	Of          QualifiedID
	OfName      string
	FilePrivate bool
	LocalName   string
}
