package query

import (
	"sort"
	"strings"

	"github.com/hiroshiasayadev-prog/brewprint/v01/src/internal/semantic"
)

func (s *Service) ListObjects(req ListObjectsRequest) (ListObjectsResponse, error) {
	if err := s.requireProject(); err != nil {
		return ListObjectsResponse{}, err
	}
	req, err := s.validateListObjectsRequest(req)
	if err != nil {
		return ListObjectsResponse{}, err
	}

	var objects []ObjectRef
	objects = append(objects, s.listNodeObjects(req)...)
	objects = append(objects, s.listViewObjects(req)...)
	objects = append(objects, s.listTransitionObjects(req)...)
	objects = append(objects, s.listFieldObjects(req)...)

	sort.Slice(objects, func(i, j int) bool {
		if objects[i].Object != objects[j].Object {
			return objects[i].Object < objects[j].Object
		}
		if objects[i].Kind != objects[j].Kind {
			return objects[i].Kind < objects[j].Kind
		}
		if objects[i].File != objects[j].File {
			return objects[i].File < objects[j].File
		}
		return objects[i].ID < objects[j].ID
	})

	return ListObjectsResponse{Objects: objects, Diagnostics: []semantic.Diagnostic{}}, nil
}

func (s *Service) listNodeObjects(req ListObjectsRequest) []ObjectRef {
	if req.Object != "" && req.Object != "node" {
		return nil
	}
	out := []ObjectRef{}
	for _, node := range s.project.NodesByQID {
		ref := objectRef(node)
		if listObjectMatches(ref, req) {
			out = append(out, listObjectRef(ref))
		}
	}
	return out
}

func (s *Service) listViewObjects(req ListObjectsRequest) []ObjectRef {
	if req.Object != "" && req.Object != "view" {
		return nil
	}
	out := []ObjectRef{}
	for _, scenario := range s.project.ScenariosByID {
		ref := scenarioObjectRef(scenario)
		if listObjectMatches(ref, req) {
			out = append(out, listObjectRef(ref))
		}
	}
	for _, view := range s.project.APIViewsByID {
		ref := ObjectRef{Object: "view", Kind: "api_table", ID: view.ID, Label: view.ID, File: view.FileID.String()}
		if listObjectMatches(ref, req) {
			out = append(out, listObjectRef(ref))
		}
	}
	for _, view := range s.project.ERViewsByID {
		ref := ObjectRef{Object: "view", Kind: "er_diagram", ID: view.ID, Label: view.ID, File: view.FileID.String()}
		if listObjectMatches(ref, req) {
			out = append(out, listObjectRef(ref))
		}
	}
	return out
}

func (s *Service) listTransitionObjects(req ListObjectsRequest) []ObjectRef {
	if req.Object != "" && req.Object != "transition" {
		return nil
	}
	out := []ObjectRef{}
	for _, transitions := range s.project.TransitionsByFile {
		for _, transition := range transitions {
			ref := transitionObjectRef(transition)
			if listObjectMatches(ref, req) {
				out = append(out, listObjectRef(ref))
			}
		}
	}
	return out
}

func (s *Service) listFieldObjects(req ListObjectsRequest) []ObjectRef {
	if req.Object != "" && req.Object != "field" {
		return nil
	}
	out := []ObjectRef{}
	for _, model := range s.project.ModelsByQID {
		for _, field := range model.Fields {
			ref := fieldObjectRef(model, field)
			if listObjectMatches(ref, req) {
				out = append(out, listObjectRef(ref))
			}
		}
	}
	return out
}

func listObjectMatches(ref ObjectRef, req ListObjectsRequest) bool {
	if req.Object != "" && ref.Object != req.Object {
		return false
	}
	if req.Kind != "" && ref.Kind != req.Kind {
		return false
	}
	if req.File != "" && ref.File != req.File {
		return false
	}
	if req.Module != "" && moduleForObjectRef(ref) != req.Module {
		return false
	}
	return true
}

func listObjectRef(ref ObjectRef) ObjectRef {
	ref.Module = moduleForObjectRef(ref)
	if ref.File != "" {
		ref.Source = sourceMap(semantic.FileID(ref.File))
	}
	return ref
}

func moduleForObjectRef(ref ObjectRef) string {
	if ref.QualifiedID != "" {
		return moduleForQualifiedID(ref.QualifiedID)
	}
	if ref.File != "" {
		return moduleForFileID(semantic.FileID(ref.File))
	}
	return ""
}

func moduleForQualifiedID(id string) string {
	parts := strings.Split(id, ".")
	for i, part := range parts {
		if isKindSegment(part) {
			return strings.Join(parts[:i], ".")
		}
	}
	return ""
}

func isKindSegment(part string) bool {
	kinds := map[string]struct{}{
		"task": {}, "model": {}, "store": {}, "state": {}, "event": {},
		"branch": {}, "fork": {}, "join": {},
	}
	_, ok := kinds[part]
	return ok
}
