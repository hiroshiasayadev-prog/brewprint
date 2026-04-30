package semantic

import "strconv"

type ObjectKey string

type ReferenceKind string

const (
	ReferenceKindParamModel    ReferenceKind = "param_model"
	ReferenceKindReturnModel   ReferenceKind = "return_model"
	ReferenceKindProducesAsset ReferenceKind = "produces_asset"
	ReferenceKindConsumesAsset ReferenceKind = "consumes_asset"
	ReferenceKindReads         ReferenceKind = "reads"
	ReferenceKindWrites        ReferenceKind = "writes"
	ReferenceKindStoreOf       ReferenceKind = "store_of"
	ReferenceKindFieldType        ReferenceKind = "field_type"
	ReferenceKindFieldFK          ReferenceKind = "field_fk"
	ReferenceKindTransitionEvent  ReferenceKind = "transition_event"
	ReferenceKindTransitionFrom   ReferenceKind = "transition_from"
	ReferenceKindTransitionTo     ReferenceKind = "transition_to"
	ReferenceKindTransitionAction ReferenceKind = "transition_action"
	ReferenceKindEventPayload           ReferenceKind = "event_payload"
	ReferenceKindEventActor             ReferenceKind = "event_actor"
	ReferenceKindEventWatches           ReferenceKind = "event_watches"
	ReferenceKindScenarioStateFile      ReferenceKind = "scenario_state_file"
	ReferenceKindScenarioStepTransition ReferenceKind = "scenario_step_transition"
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
	StateFile   FileID
	From        string
	On          string
	To          string
	Guard       string
	Action      QualifiedID
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

func ScenarioObjectKey(id string) ObjectKey {
	if id == "" {
		return ""
	}
	return ObjectKey("scenario:" + id)
}

func ScenarioStepObjectKey(scenarioID string, index int) ObjectKey {
	if scenarioID == "" || index <= 0 {
		return ""
	}
	return ObjectKey("scenario_step:" + scenarioID + ":" + strconv.Itoa(index))
}

func StateFileObjectKey(fileID FileID) ObjectKey {
	if fileID == "" {
		return ""
	}
	return ObjectKey("file:" + fileID.String())
}

func TransitionObjectKey(transition Transition) ObjectKey {
	if transition.FileID == "" || transition.From == "" || transition.On == "" {
		return ""
	}
	return ObjectKey("transition:" + TransitionID(transition))
}

func TransitionID(transition Transition) string {
	id := transition.FileID.String() + "#" + transition.From + ":" + transition.On
	if transition.Guard != "" {
		id += "[" + transition.Guard + "]"
	}
	return id
}
