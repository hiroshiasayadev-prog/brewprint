package query

import (
	"fmt"
	"sort"
	"strings"

	"github.com/hiroshiasayadev-prog/brewprint/v01/src/internal/semantic"
)

const (
	defaultReferenceTreeMaxNodes = 200
	defaultReferenceTreeMaxEdges = 500
	maxReferenceTreeDepth        = 4
)

func (s *Service) GetReferenceTree(req GetReferenceTreeRequest) (GetReferenceTreeResponse, error) {
	if err := s.ensureSelectorSupported(toolGetReferenceTree, req.Selector); err != nil {
		return GetReferenceTreeResponse{}, err
	}

	rootKey, root, err := s.referenceTarget(req.Selector)
	if err != nil {
		return GetReferenceTreeResponse{}, err
	}
	if !isReferenceTreeSupportedRoot(root) {
		return GetReferenceTreeResponse{}, fmt.Errorf("unsupported object for reference tree: %s:%s", root.Object, root.Kind)
	}

	direction := req.Direction
	if direction != string(semantic.ReferenceDirectionOut) && direction != string(semantic.ReferenceDirectionIn) && direction != string(semantic.ReferenceDirectionBoth) {
		return GetReferenceTreeResponse{}, fmt.Errorf("unsupported direction: %s", direction)
	}
	if req.Depth < 0 || req.Depth > maxReferenceTreeDepth {
		return GetReferenceTreeResponse{}, fmt.Errorf("invalid depth: %d", req.Depth)
	}

	maxNodes := req.MaxNodes
	if maxNodes <= 0 {
		maxNodes = defaultReferenceTreeMaxNodes
	}
	maxEdges := req.MaxEdges
	if maxEdges <= 0 {
		maxEdges = defaultReferenceTreeMaxEdges
	}

	kindFilter := map[string]struct{}{}
	for _, kind := range req.Kinds {
		kindFilter[kind] = struct{}{}
	}

	type queueItem struct {
		key   semantic.ObjectKey
		depth int
		via   []string
	}

	nodes := []ReferenceTreeNode{{Object: root, Depth: 0, Via: []string{}}}
	edges := []ReferenceTreeEdge{}
	visited := map[semantic.ObjectKey]struct{}{rootKey: struct{}{}}
	queue := []queueItem{{key: rootKey, depth: 0, via: []string{}}}
	truncatedReasons := []string{}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if current.depth >= req.Depth {
			continue
		}

		candidates := s.referenceTreeCandidatesForRoot(current.key, rootKey, root, direction, kindFilter)
		for _, candidate := range candidates {
			nextDepth := current.depth + 1
			if _, seen := visited[candidate.nextKey]; !seen && len(nodes) >= maxNodes {
				truncatedReasons = appendUniqueString(truncatedReasons, "max_nodes")
				queue = nil
				break
			}
			if len(edges) >= maxEdges {
				truncatedReasons = appendUniqueString(truncatedReasons, "max_edges")
				queue = nil
				break
			}

			edges = append(edges, ReferenceTreeEdge{
				Kind:      string(candidate.ref.Kind),
				Direction: candidate.direction,
				From:      endpointFromSemantic(candidate.ref.From),
				To:        endpointFromSemantic(candidate.ref.To),
				Depth:     nextDepth,
			})

			if _, seen := visited[candidate.nextKey]; seen {
				continue
			}
			visited[candidate.nextKey] = struct{}{}
			via := append(append([]string{}, current.via...), string(candidate.ref.Kind))
			nodes = append(nodes, ReferenceTreeNode{
				Object: s.objectRefForReferenceKey(candidate.nextKey, candidate.nextEndpoint),
				Depth:  nextDepth,
				Via:    via,
			})
			queue = append(queue, queueItem{key: candidate.nextKey, depth: nextDepth, via: via})
		}
	}

	return GetReferenceTreeResponse{
		Root:             root,
		Direction:        direction,
		Depth:            req.Depth,
		Nodes:            nodes,
		Edges:            edges,
		Truncated:        len(truncatedReasons) > 0,
		TruncatedReasons: truncatedReasons,
		Diagnostics:      []semantic.Diagnostic{},
	}, nil
}

func isReferenceTreeSupportedRoot(root ObjectRef) bool {
	switch root.Object {
	case "node", "transition", "field", "asset":
		return true
	case "view":
		return root.Kind == "sequence_diagram"
	case "file":
		return root.Kind == "node" || root.Kind == "state_file"
	default:
		return false
	}
}

func (s *Service) referenceTreeCandidatesForRoot(key, rootKey semantic.ObjectKey, root ObjectRef, direction string, kindFilter map[string]struct{}) []referenceTreeCandidate {
	if key == rootKey && root.Object == "file" {
		switch root.Kind {
		case "node":
			return s.nodeFileReferenceTreeCandidates(semantic.FileID(root.ID), direction, kindFilter)
		case "state_file":
			return s.stateFileReferenceTreeCandidates(semantic.FileID(root.ID), direction, kindFilter)
		}
	}
	return s.referenceTreeCandidates(key, direction, kindFilter)
}

func (s *Service) nodeFileReferenceTreeCandidates(fileID semantic.FileID, direction string, kindFilter map[string]struct{}) []referenceTreeCandidate {
	var out []referenceTreeCandidate
	for _, node := range s.project.NodesByFile[fileID] {
		key := semantic.NodeObjectKey(node.GetQID())
		out = append(out, s.referenceTreeCandidatesFromIndexes(
			filterNodeFileReferences(s.project.ReferencesBySource[key]),
			filterNodeFileReferences(s.project.ReferencesByTarget[key]),
			direction,
			kindFilter,
		)...)
	}
	sortReferenceTreeCandidates(out)
	return out
}

func (s *Service) stateFileReferenceTreeCandidates(fileID semantic.FileID, direction string, kindFilter map[string]struct{}) []referenceTreeCandidate {
	var out []referenceTreeCandidate
	for _, key := range stateFileReferenceKeys(s.project, fileID) {
		out = append(out, s.referenceTreeCandidatesFromIndexes(
			s.project.ReferencesBySource[key],
			s.project.ReferencesByTarget[key],
			direction,
			kindFilter,
		)...)
	}
	sortReferenceTreeCandidates(out)
	return out
}

type referenceTreeCandidate struct {
	ref          semantic.Reference
	direction    string
	nextKey      semantic.ObjectKey
	nextEndpoint semantic.ReferenceEndpoint
}

func (s *Service) referenceTreeCandidates(key semantic.ObjectKey, direction string, kindFilter map[string]struct{}) []referenceTreeCandidate {
	out := s.referenceTreeCandidatesFromIndexes(s.project.ReferencesBySource[key], s.project.ReferencesByTarget[key], direction, kindFilter)
	sortReferenceTreeCandidates(out)
	return out
}

func (s *Service) referenceTreeCandidatesFromIndexes(sourceRefs, targetRefs []semantic.Reference, direction string, kindFilter map[string]struct{}) []referenceTreeCandidate {
	out := []referenceTreeCandidate{}
	if direction == string(semantic.ReferenceDirectionOut) || direction == string(semantic.ReferenceDirectionBoth) {
		for _, ref := range sourceRefs {
			if !referenceKindAllowed(ref.Kind, kindFilter) {
				continue
			}
			out = append(out, referenceTreeCandidate{
				ref:          ref,
				direction:    string(semantic.ReferenceDirectionOut),
				nextKey:      ref.TargetKey,
				nextEndpoint: ref.To,
			})
		}
	}
	if direction == string(semantic.ReferenceDirectionIn) || direction == string(semantic.ReferenceDirectionBoth) {
		for _, ref := range targetRefs {
			if !referenceKindAllowed(ref.Kind, kindFilter) {
				continue
			}
			out = append(out, referenceTreeCandidate{
				ref:          ref,
				direction:    string(semantic.ReferenceDirectionIn),
				nextKey:      ref.SourceKey,
				nextEndpoint: ref.From,
			})
		}
	}
	return out
}

func sortReferenceTreeCandidates(out []referenceTreeCandidate) {
	sort.Slice(out, func(i, j int) bool {
		if out[i].direction != out[j].direction {
			return out[i].direction < out[j].direction
		}
		if out[i].ref.Kind != out[j].ref.Kind {
			return out[i].ref.Kind < out[j].ref.Kind
		}
		if out[i].ref.From.ID != out[j].ref.From.ID {
			return out[i].ref.From.ID < out[j].ref.From.ID
		}
		return out[i].ref.To.ID < out[j].ref.To.ID
	})
}

func referenceKindAllowed(kind semantic.ReferenceKind, kindFilter map[string]struct{}) bool {
	if len(kindFilter) == 0 {
		return true
	}
	_, ok := kindFilter[string(kind)]
	return ok
}

func (s *Service) objectRefForReferenceKey(key semantic.ObjectKey, endpoint semantic.ReferenceEndpoint) ObjectRef {
	rawKey := string(key)
	if strings.HasPrefix(rawKey, "scenario:") {
		id := strings.TrimPrefix(rawKey, "scenario:")
		if scenario := s.project.ScenariosByID[id]; scenario != nil {
			return scenarioObjectRef(scenario)
		}
	}
	if strings.HasPrefix(rawKey, "file:") {
		fileID := semantic.FileID(strings.TrimPrefix(rawKey, "file:"))
		kind := s.fileKind(fileID)
		if kind == "file" && s.isStateFile(fileID) {
			kind = "state_file"
		}
		return fileObjectRef(fileID, kind)
	}
	if !strings.Contains(rawKey, ":") {
		if node := s.project.NodesByQID[semantic.QualifiedID(rawKey)]; node != nil {
			return objectRef(node)
		}
	}
	return objectRefFromReferenceEndpoint(endpoint)
}

func objectRefFromReferenceEndpoint(endpoint semantic.ReferenceEndpoint) ObjectRef {
	object := endpoint.Object
	kind := endpoint.Kind
	if object == "model_field" {
		object = "field"
	}
	if object == "" {
		object = kind
	}
	if kind == "" {
		kind = object
	}
	ref := ObjectRef{
		Object:      object,
		Kind:        kind,
		ID:          endpoint.ID,
		QualifiedID: endpoint.QualifiedID.String(),
		Label:       endpoint.Name,
		File:        endpoint.File.String(),
		LocalID:     endpoint.LocalID,
	}
	if ref.File != "" {
		ref.Source = sourceMap(semantic.FileID(ref.File))
	}
	return ref
}

func appendUniqueString(values []string, value string) []string {
	for _, existing := range values {
		if existing == value {
			return values
		}
	}
	return append(values, value)
}
