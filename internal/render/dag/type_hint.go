package dag

import (
	"strings"

	"github.com/hiroshiasayadev-prog/brewprint/internal/semantic"
)

// assetTypeHint returns the top-level DAG label hint for a TypeRef.
// Returns "" when ref is nil or the kind is unknown.
func assetTypeHint(ref *semantic.TypeRef) string {
	if ref == nil {
		return ""
	}
	switch ref.Kind {
	case semantic.TypeRefPrimitive:
		return ref.Name
	case semantic.TypeRefNamedModel:
		return namedModelLocalID(ref)
	case semantic.TypeRefList:
		return "list"
	case semantic.TypeRefDict:
		return "dict"
	default:
		return ""
	}
}

// namedModelLocalID extracts the local ID from a named-model TypeRef.
// Uses ref.Model (resolved QID) first; falls back to ref.Name.
// Public QIDs use dot-separators ("auth.model.login_form" → "login_form").
// Private QIDs use hash-separator ("auth/task/login.yaml#login_form" → "login_form").
func namedModelLocalID(ref *semantic.TypeRef) string {
	if ref == nil || ref.Kind != semantic.TypeRefNamedModel {
		return ""
	}
	qid := string(ref.Model)
	if qid == "" {
		qid = ref.Name
	}
	// Private model QIDs: fileID#localID
	if idx := strings.LastIndex(qid, "#"); idx >= 0 {
		return qid[idx+1:]
	}
	// Public model QIDs: module.model.localID
	if idx := strings.LastIndex(qid, "."); idx >= 0 {
		return qid[idx+1:]
	}
	return qid
}

// calcAssetHint returns the hint string for a DAG asset node label.
// Returns "" when the hint cannot be determined, the TypeRef is invalid/unresolved,
// or the named model local ID is ambiguous in the DAG render scope.
func calcAssetHint(ref *semantic.TypeRef, ambiguous map[string]bool) string {
	hint := assetTypeHint(ref)
	if hint == "" {
		return ""
	}
	if ref != nil && ref.Kind == semantic.TypeRefNamedModel && ambiguous[hint] {
		return ""
	}
	return hint
}

// collectFlowAssetTypeRefs gathers all TypeRefs for assets that appear in a flow DAG render.
// Covers: params boundary, task/join returns, foreach collected assets.
func collectFlowAssetTypeRefs(project *semantic.Project, main *semantic.Task) []*semantic.TypeRef {
	var refs []*semantic.TypeRef

	for _, param := range main.Params {
		refs = append(refs, param.TypeRef)
	}

	flows := project.FlowByFile[main.FileID]
	collectedSources := project.FlowCollectedSourcesByFile[main.FileID]

	for _, entry := range flows {
		switch entry.Kind {
		case semantic.FlowKindStep:
			if task := project.TasksByQID[entry.Step.Task]; task != nil && task.Returns != nil {
				refs = append(refs, task.Returns.TypeRef)
			}
		case semantic.FlowKindForeach:
			if entry.Foreach.Returns != "" && collectedSources != nil {
				if cs := collectedSources[entry.Foreach.Returns]; cs != nil {
					refs = append(refs, cs.TypeRef)
				}
			}
		case semantic.FlowKindFork:
			for _, branch := range entry.Fork.Branches {
				for _, step := range branch.Steps {
					if task := project.TasksByQID[step.Task]; task != nil && task.Returns != nil {
						refs = append(refs, task.Returns.TypeRef)
					}
				}
			}
			if join := project.JoinsByQID[entry.Fork.Join]; join != nil && join.Returns != nil {
				refs = append(refs, join.Returns.TypeRef)
			}
		case semantic.FlowKindBranch:
			for _, c := range entry.Branch.Cases {
				if task := project.TasksByQID[c.Step.Task]; task != nil && task.Returns != nil {
					refs = append(refs, task.Returns.TypeRef)
				}
			}
		}
	}

	return refs
}

// computeAmbiguousHints returns the set of named-model local IDs that map to more than one
// distinct QID within the DAG render scope (same local ID, different identities).
func computeAmbiguousHints(refs []*semantic.TypeRef) map[string]bool {
	localIDToQIDs := map[string]map[semantic.QualifiedID]bool{}
	for _, ref := range refs {
		if ref == nil || ref.Kind != semantic.TypeRefNamedModel {
			continue
		}
		localID := namedModelLocalID(ref)
		if localID == "" {
			continue
		}
		if localIDToQIDs[localID] == nil {
			localIDToQIDs[localID] = map[semantic.QualifiedID]bool{}
		}
		if ref.Model != "" {
			localIDToQIDs[localID][ref.Model] = true
		}
	}
	ambiguous := map[string]bool{}
	for localID, qids := range localIDToQIDs {
		if len(qids) > 1 {
			ambiguous[localID] = true
		}
	}
	return ambiguous
}
