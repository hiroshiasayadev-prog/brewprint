package semantic

type ObjectKey string

type ReferenceKind string

const (
	ReferenceKindParamModel    ReferenceKind = "param_model"
	ReferenceKindReturnModel   ReferenceKind = "return_model"
	ReferenceKindProducesAsset ReferenceKind = "produces_asset"
	ReferenceKindReads         ReferenceKind = "reads"
	ReferenceKindWrites        ReferenceKind = "writes"
	ReferenceKindStoreOf       ReferenceKind = "store_of"
	ReferenceKindFieldType     ReferenceKind = "field_type"
	ReferenceKindFieldFK       ReferenceKind = "field_fk"
)

type ReferenceDirection string

const (
	ReferenceDirectionOut  ReferenceDirection = "out"
	ReferenceDirectionIn   ReferenceDirection = "in"
	ReferenceDirectionBoth ReferenceDirection = "both"
)

type ReferenceEndpoint struct {
	Object      string
	Kind        string
	ID          string
	QualifiedID QualifiedID
	Name        string
	Producer    QualifiedID
	Model       QualifiedID
	ScopeFile   FileID
	File        FileID
	LocalID     string
}

type Reference struct {
	Kind      ReferenceKind
	SourceKey ObjectKey
	TargetKey ObjectKey
	From      ReferenceEndpoint
	To        ReferenceEndpoint
}

func NodeObjectKey(qid QualifiedID) ObjectKey {
	return ObjectKey(qid.String())
}

func AssetObjectKey(asset *Asset) ObjectKey {
	if asset == nil {
		return ""
	}
	return ObjectKey("asset:" + asset.ProducedBy.String() + ":" + asset.Name)
}

func ModelFieldObjectKey(model QualifiedID, fieldName string) ObjectKey {
	if model == "" || fieldName == "" {
		return ""
	}
	return ObjectKey("field:" + model.String() + ":" + fieldName)
}

func PrimitiveObjectKey(name string) ObjectKey {
	if name == "" {
		return ""
	}
	return ObjectKey("primitive:" + name)
}
