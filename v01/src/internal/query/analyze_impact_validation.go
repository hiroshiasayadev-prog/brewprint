package query

import (
	"encoding/json"
	"fmt"
	"sort"
)

const (
	AnalyzeImpactChangeRename                 = "rename"
	AnalyzeImpactChangeRemove                 = "remove"
	AnalyzeImpactChangeType                   = "change_type"
	AnalyzeImpactChangeContract               = "change_contract"
	AnalyzeImpactChangeTransitionTarget       = "change_transition_target"
	AnalyzeImpactChangeAdd                    = "add"
	analyzeImpactInvalidChangePayloadFragment = "invalid change payload"
)

func (c *AnalyzeImpactChange) UnmarshalJSON(data []byte) error {
	type alias AnalyzeImpactChange
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	var decoded alias
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	decoded.rawKeys = map[string]struct{}{}
	for key := range raw {
		decoded.rawKeys[key] = struct{}{}
	}
	*c = AnalyzeImpactChange(decoded)
	return nil
}

func validateAnalyzeImpactChange(change AnalyzeImpactChange) error {
	if change.Kind == "" {
		return invalidAnalyzeImpactChange("change.kind is required")
	}

	allowed := []string{"kind"}
	switch change.Kind {
	case AnalyzeImpactChangeRename:
		if change.NewID == "" {
			return invalidAnalyzeImpactChange("rename.new_id is required")
		}
		allowed = append(allowed, "new_id")
	case AnalyzeImpactChangeRemove:
		// no payload
	case AnalyzeImpactChangeType:
		if change.NewType == "" {
			return invalidAnalyzeImpactChange("change_type.new_type is required")
		}
		allowed = append(allowed, "new_type")
	case AnalyzeImpactChangeContract:
		allowed = append(allowed, "note")
	case AnalyzeImpactChangeTransitionTarget:
		if change.NewTo == "" && change.NewAction == "" {
			return invalidAnalyzeImpactChange("change_transition_target.new_to or new_action is required")
		}
		allowed = append(allowed, "new_to", "new_action")
	case AnalyzeImpactChangeAdd:
		if change.AddedID == "" {
			return invalidAnalyzeImpactChange("add.added_id is required")
		}
		allowed = append(allowed, "added_id")
	default:
		return invalidAnalyzeImpactChange("unsupported change.kind: " + change.Kind)
	}

	if len(change.rawKeys) > 0 {
		allowedSet := map[string]struct{}{}
		for _, key := range allowed {
			allowedSet[key] = struct{}{}
		}
		var extras []string
		for key := range change.rawKeys {
			if _, ok := allowedSet[key]; !ok {
				extras = append(extras, key)
			}
		}
		if len(extras) > 0 {
			sort.Strings(extras)
			return invalidAnalyzeImpactChange("unexpected payload for " + change.Kind + ": " + joinStrings(extras, ", "))
		}
	}

	return nil
}

func invalidAnalyzeImpactChange(message string) error {
	return fmt.Errorf("%s: %s", analyzeImpactInvalidChangePayloadFragment, message)
}

func joinStrings(values []string, sep string) string {
	if len(values) == 0 {
		return ""
	}
	out := values[0]
	for _, value := range values[1:] {
		out += sep + value
	}
	return out
}
