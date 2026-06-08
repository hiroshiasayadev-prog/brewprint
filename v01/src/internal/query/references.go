package query

import (
	"fmt"
	"sort"

	"github.com/hiroshiasayadev-prog/brewprint/v01/src/internal/semantic"
)

func (s *Service) GetReferences(req GetReferencesRequest) (GetReferencesResponse, error) {
	if err := s.ensureSelectorSupported(toolGetReferences, req.Selector); err != nil {
		return GetReferencesResponse{}, err
	}
	key, object, err := s.referenceTarget(req.Selector)
	if err != nil {
		return GetReferencesResponse{}, err
	}
	direction := req.Direction
	if direction == "" {
		direction = string(semantic.ReferenceDirectionOut)
	}
	if direction != string(semantic.ReferenceDirectionOut) && direction != string(semantic.ReferenceDirectionIn) && direction != string(semantic.ReferenceDirectionBoth) {
		return GetReferencesResponse{}, fmt.Errorf("unsupported direction: %s", direction)
	}

	kindFilter := map[string]struct{}{}
	for _, kind := range req.Kinds {
		kindFilter[kind] = struct{}{}
	}

	if object.Object == "file" {
		return s.getFileReferences(req, object, direction, kindFilter)
	}

	var refs []Reference
	if direction == string(semantic.ReferenceDirectionOut) || direction == string(semantic.ReferenceDirectionBoth) {
		refs = append(refs, s.referencesFromIndex(s.project.ReferencesBySource[key], string(semantic.ReferenceDirectionOut), kindFilter)...)
	}
	if direction == string(semantic.ReferenceDirectionIn) || direction == string(semantic.ReferenceDirectionBoth) {
		refs = append(refs, s.referencesFromIndex(s.project.ReferencesByTarget[key], string(semantic.ReferenceDirectionIn), kindFilter)...)
	}

	sortReferences(refs)
	return GetReferencesResponse{
		Object:      object,
		Direction:   direction,
		Depth:       1,
		References:  refs,
		Diagnostics: []semantic.Diagnostic{},
	}, nil
}

func (s *Service) getFileReferences(req GetReferencesRequest, object ObjectRef, direction string, kindFilter map[string]struct{}) (GetReferencesResponse, error) {
	fileID := semantic.FileID(object.ID)
	var refs []Reference
	switch object.Kind {
	case "node":
		refs = s.nodeFileReferences(fileID, direction, kindFilter)
	case "state_file":
		refs = s.stateFileReferences(fileID, direction, kindFilter)
	default:
		return GetReferencesResponse{}, fmt.Errorf("unsupported object for references: file:%s", object.Kind)
	}
	sortReferences(refs)
	return GetReferencesResponse{
		Object:      object,
		Direction:   direction,
		Depth:       1,
		References:  refs,
		Diagnostics: []semantic.Diagnostic{},
	}, nil
}

func (s *Service) nodeFileReferences(fileID semantic.FileID, direction string, kindFilter map[string]struct{}) []Reference {
	var refs []Reference
	for _, node := range s.project.NodesByFile[fileID] {
		key := semantic.NodeObjectKey(node.GetQID())
		if direction == string(semantic.ReferenceDirectionOut) || direction == string(semantic.ReferenceDirectionBoth) {
			refs = append(refs, s.referencesFromIndex(filterNodeFileReferences(s.project.ReferencesBySource[key]), string(semantic.ReferenceDirectionOut), kindFilter)...)
		}
		if direction == string(semantic.ReferenceDirectionIn) || direction == string(semantic.ReferenceDirectionBoth) {
			refs = append(refs, s.referencesFromIndex(filterNodeFileReferences(s.project.ReferencesByTarget[key]), string(semantic.ReferenceDirectionIn), kindFilter)...)
		}
	}
	return refs
}

func filterNodeFileReferences(refs []semantic.Reference) []semantic.Reference {
	out := make([]semantic.Reference, 0, len(refs))
	for _, ref := range refs {
		switch ref.Kind {
		case semantic.ReferenceKindConsumesAsset, semantic.ReferenceKindScenarioStateFile, semantic.ReferenceKindScenarioStepTransition:
			continue
		default:
			out = append(out, ref)
		}
	}
	return out
}

func (s *Service) stateFileReferences(fileID semantic.FileID, direction string, kindFilter map[string]struct{}) []Reference {
	keys := stateFileReferenceKeys(s.project, fileID)
	var refs []Reference
	for _, key := range keys {
		if direction == string(semantic.ReferenceDirectionOut) || direction == string(semantic.ReferenceDirectionBoth) {
			refs = append(refs, s.referencesFromIndex(s.project.ReferencesBySource[key], string(semantic.ReferenceDirectionOut), kindFilter)...)
		}
		if direction == string(semantic.ReferenceDirectionIn) || direction == string(semantic.ReferenceDirectionBoth) {
			refs = append(refs, s.referencesFromIndex(s.project.ReferencesByTarget[key], string(semantic.ReferenceDirectionIn), kindFilter)...)
		}
	}
	return refs
}

func (s *Service) referenceTarget(selector Selector) (semantic.ObjectKey, ObjectRef, error) {
	if s.isScenarioSelector(selector) {
		scenario, err := s.scenarioBySelector(selector)
		if err != nil {
			return "", ObjectRef{}, err
		}
		return semantic.ScenarioObjectKey(scenario.ID), scenarioObjectRef(scenario), nil
	}
	if s.isTransitionSelector(selector) {
		transition, err := s.transitionBySelector(selector)
		if err != nil {
			return "", ObjectRef{}, err
		}
		return semantic.TransitionObjectKey(transition), transitionObjectRef(transition), nil
	}
	if s.isFieldSelector(selector) {
		model, field, err := s.modelFieldBySelector(selector)
		if err != nil {
			return "", ObjectRef{}, err
		}
		return semantic.ModelFieldObjectKey(model.QID, field.Name), fieldObjectRef(model, field), nil
	}
	if s.isAssetSelector(selector) {
		asset, err := s.assetBySelector(selector)
		if err != nil {
			return "", ObjectRef{}, err
		}
		return semantic.AssetObjectKey(asset), assetObjectRef(asset), nil
	}
	if s.isPrivateSubNodeSelector(selector) {
		node, err := s.privateSubNodeBySelector(selector)
		if err != nil {
			return "", ObjectRef{}, err
		}
		return semantic.NodeObjectKey(node.GetQID()), objectRef(node), nil
	}
	if s.isFileSelector(selector) {
		fileID, err := s.fileBySelector(selector)
		if err != nil {
			return "", ObjectRef{}, err
		}
		kind := selector.Kind
		if kind == "" && s.project.TransitionsByFile[fileID] != nil {
			kind = "state_file"
		}
		return semantic.StateFileObjectKey(fileID), fileObjectRef(fileID, kind), nil
	}
	node, err := s.nodeByID(selector.ID)
	if err != nil {
		return "", ObjectRef{}, err
	}
	return semantic.NodeObjectKey(node.GetQID()), objectRef(node), nil
}

func (s *Service) referencesFromIndex(in []semantic.Reference, direction string, kindFilter map[string]struct{}) []Reference {
	out := make([]Reference, 0, len(in))
	for _, ref := range in {
		if len(kindFilter) > 0 {
			if _, ok := kindFilter[string(ref.Kind)]; !ok {
				continue
			}
		}
		out = append(out, Reference{
			Kind:      string(ref.Kind),
			Direction: direction,
			From:      endpointFromSemantic(ref.From),
			To:        endpointFromSemantic(ref.To),
		})
	}
	return out
}

func endpointFromSemantic(endpoint semantic.ReferenceEndpoint) ReferenceEndpoint {
	return ReferenceEndpoint{
		Object:      endpoint.Object,
		Kind:        endpoint.Kind,
		ID:          endpoint.ID,
		QualifiedID: endpoint.QualifiedID.String(),
		Name:        endpoint.Name,
		Producer:    endpoint.Producer.String(),
		Model:       endpoint.Model.String(),
		ScopeFile:   endpoint.ScopeFile.String(),
		File:        endpoint.File.String(),
		LocalID:     endpoint.LocalID,
		StateFile:   endpoint.StateFile.String(),
		FromState:   endpoint.From,
		On:          endpoint.On,
		ToState:     endpoint.To,
		Guard:       endpoint.Guard,
		Action:      endpoint.Action.String(),
	}
}

func sortReferences(refs []Reference) {
	sort.Slice(refs, func(i, j int) bool {
		if refs[i].Direction != refs[j].Direction {
			return refs[i].Direction < refs[j].Direction
		}
		if refs[i].Kind != refs[j].Kind {
			return refs[i].Kind < refs[j].Kind
		}
		if refs[i].From.ID != refs[j].From.ID {
			return refs[i].From.ID < refs[j].From.ID
		}
		return refs[i].To.ID < refs[j].To.ID
	})
}
