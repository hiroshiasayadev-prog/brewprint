package query

import (
	"fmt"

	"github.com/hiroshiasayadev-prog/brewprint/v01/src/internal/semantic"
)

const (
	toolGetSignature     = "get_signature"
	toolGetReferences    = "get_references"
	toolGetReferenceTree = "get_reference_tree"
	toolInspect          = "inspect"
)

func (s *Service) ensureSelectorSupported(tool string, selector Selector) error {
	if selector.Object == "primitive" || selector.Kind == "primitive" {
		if selector.ID == "" {
			return fmt.Errorf("selector.id is required for primitive")
		}
		return fmt.Errorf("unsupported object for %s: primitive", tool)
	}

	if selector.Object == "file" || selector.Kind == "file" || selector.Kind == "state_file" || selector.Kind == "render_index" {
		fileID, err := s.fileBySelector(selector)
		if err != nil {
			return err
		}
		kind := selector.Kind
		if kind == "" || kind == "file" {
			kind = s.fileKind(fileID)
		}
		if tool == toolGetSignature {
			return fmt.Errorf("unsupported object for %s: file:%s", tool, kind)
		}
		if tool == toolGetReferences || tool == toolGetReferenceTree {
			if kind != "node" && kind != "state_file" {
				return fmt.Errorf("unsupported object for %s: file:%s", tool, kind)
			}
		}
	}

	if selector.Object == "view" || isViewKindSelector(selector.Kind) {
		switch selector.Kind {
		case "api_table":
			if _, err := s.apiViewBySelector(selector); err != nil {
				return err
			}
			if tool != toolInspect {
				return fmt.Errorf("unsupported object for %s: view:api_table", tool)
			}
		case "er_diagram":
			if _, err := s.erViewBySelector(selector); err != nil {
				return err
			}
			if tool != toolInspect {
				return fmt.Errorf("unsupported object for %s: view:er_diagram", tool)
			}
		}
	}

	return nil
}

func isViewKindSelector(kind string) bool {
	return kind == "sequence_diagram" || kind == "api_table" || kind == "er_diagram"
}

func normalizeFieldKind(kind string) string {
	if kind == "model_field" {
		return "field"
	}
	return kind
}

func isFieldKind(kind string) bool {
	return kind == "field" || kind == "model_field"
}

func (s *Service) validateListObjectsRequest(req ListObjectsRequest) (ListObjectsRequest, error) {
	req.Kind = normalizeFieldKind(req.Kind)
	if req.Object != "" {
		if _, ok := validListObjectValues[req.Object]; !ok {
			return req, fmt.Errorf("invalid args: unknown list_objects object: %s", req.Object)
		}
	}
	if req.Kind != "" && !s.listObjectKindAllowed(req.Object, req.Kind) {
		return req, fmt.Errorf("invalid args: unknown list_objects kind: %s", req.Kind)
	}
	return req, nil
}

var validListObjectValues = map[string]struct{}{
	"node":       {},
	"view":       {},
	"transition": {},
	"field":      {},
}

func (s *Service) listObjectKindAllowed(object, kind string) bool {
	if object != "" {
		return kindAllowedForObject(object, kind)
	}
	for candidate := range validListObjectValues {
		if kindAllowedForObject(candidate, kind) {
			return true
		}
	}
	return false
}

func kindAllowedForObject(object, kind string) bool {
	switch object {
	case "node":
		switch kind {
		case "task", "model", "store", "state", "event", "actor", "branch", "fork", "join":
			return true
		}
	case "view":
		return kind == "sequence_diagram" || kind == "api_table" || kind == "er_diagram"
	case "transition":
		return kind == "transition"
	case "field":
		return kind == "field"
	}
	return false
}

func stateFileReferenceKeys(project *semantic.Project, fileID semantic.FileID) []semantic.ObjectKey {
	keys := []semantic.ObjectKey{semantic.StateFileObjectKey(fileID)}
	for _, transition := range project.TransitionsByFile[fileID] {
		keys = append(keys, semantic.TransitionObjectKey(transition))
	}
	return keys
}
