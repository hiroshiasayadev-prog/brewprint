package query

import (
	"fmt"
	"sort"

	"github.com/hiroshiasayadev-prog/brewprint/internal/semantic"
)

func (s *Service) GetReferences(req GetReferencesRequest) (GetReferencesResponse, error) {
	node, err := s.nodeByID(req.Selector.ID)
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

	key := semantic.NodeObjectKey(node.GetQID())
	var refs []Reference
	if direction == string(semantic.ReferenceDirectionOut) || direction == string(semantic.ReferenceDirectionBoth) {
		refs = append(refs, s.referencesFromIndex(s.project.ReferencesBySource[key], string(semantic.ReferenceDirectionOut), kindFilter)...)
	}
	if direction == string(semantic.ReferenceDirectionIn) || direction == string(semantic.ReferenceDirectionBoth) {
		refs = append(refs, s.referencesFromIndex(s.project.ReferencesByTarget[key], string(semantic.ReferenceDirectionIn), kindFilter)...)
	}

	sortReferences(refs)
	return GetReferencesResponse{
		Object:      objectRef(node),
		Direction:   direction,
		Depth:       1,
		References:  refs,
		Diagnostics: []semantic.Diagnostic{},
	}, nil
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
