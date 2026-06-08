package query

import (
	"fmt"
	"sort"
	"strings"

	"github.com/hiroshiasayadev-prog/brewprint/v01/src/internal/semantic"
)

type ListEndpointsRequest struct {
	APITableID string `json:"api_table_id"`
}

type ListEndpointsResponse struct {
	Tables      []EndpointTable       `json:"tables"`
	Diagnostics []semantic.Diagnostic `json:"diagnostics"`
}

type EndpointTable struct {
	ID           string            `json:"id"`
	HTTPRootPath string            `json:"http_root_path"`
	Sections     []EndpointSection `json:"sections"`
}

type EndpointSection struct {
	Module            string     `json:"module"`
	IncludeSubmodules bool       `json:"include_submodules"`
	Endpoints         []Endpoint `json:"endpoints"`
}

type Endpoint struct {
	Method   string            `json:"method"`
	Path     string            `json:"path"`
	LeafPath string            `json:"leaf_path"`
	Task     string            `json:"task"`
	Params   []string          `json:"params,omitempty"`
	Returns  string            `json:"returns,omitempty"`
	Source   map[string]string `json:"source,omitempty"`
}

func (s *Service) ListEndpoints(req ListEndpointsRequest) (ListEndpointsResponse, error) {
	if err := s.requireProject(); err != nil {
		return ListEndpointsResponse{}, err
	}

	views := s.endpointViews(req.APITableID)
	if req.APITableID != "" && len(views) == 0 {
		return ListEndpointsResponse{}, fmt.Errorf("API view not found: %s", req.APITableID)
	}

	out := ListEndpointsResponse{Diagnostics: []semantic.Diagnostic{}}
	for _, view := range views {
		out.Tables = append(out.Tables, s.endpointTable(view))
	}
	return out, nil
}

func (s *Service) endpointViews(id string) []*semantic.APIView {
	var views []*semantic.APIView
	if id != "" {
		if view := s.project.APIViewsByID[id]; view != nil {
			views = append(views, view)
		}
		return views
	}
	for _, view := range s.project.APIViewsByID {
		views = append(views, view)
	}
	sort.Slice(views, func(i, j int) bool { return views[i].ID < views[j].ID })
	return views
}

func (s *Service) endpointTable(view *semantic.APIView) EndpointTable {
	table := EndpointTable{ID: view.ID, HTTPRootPath: view.HTTPRootPath}
	for _, module := range view.Modules {
		section := EndpointSection{
			Module:            module.Module,
			IncludeSubmodules: module.IncludeSubmodules,
			Endpoints:         s.endpointsForModule(view, module),
		}
		if len(section.Endpoints) == 0 {
			continue
		}
		table.Sections = append(table.Sections, section)
	}
	return table
}

func (s *Service) endpointsForModule(view *semantic.APIView, viewModule semantic.APIViewModule) []Endpoint {
	var endpoints []Endpoint
	for _, task := range s.project.TasksByQID {
		if !task.Endpoint {
			continue
		}
		module := moduleForFileID(task.FileID)
		if !endpointModuleIncluded(module, viewModule) {
			continue
		}
		leaf := endpointLeafPath(task)
		routeID := endpointRelativeID(viewModule.Module, module, leaf)
		endpoints = append(endpoints, Endpoint{
			Method:   task.Method,
			Path:     endpointFullPath(view.HTTPRootPath, routeID),
			LeafPath: leaf,
			Task:     task.QID.String(),
			Params:   endpointParams(task),
			Returns:  endpointReturns(task),
			Source:   sourceMap(task.FileID),
		})
	}
	sort.Slice(endpoints, func(i, j int) bool { return endpoints[i].Task < endpoints[j].Task })
	return endpoints
}

func endpointModuleIncluded(module string, viewModule semantic.APIViewModule) bool {
	if module == viewModule.Module {
		return true
	}
	return viewModule.IncludeSubmodules && strings.HasPrefix(module, viewModule.Module+".")
}

func endpointLeafPath(task *semantic.Task) string {
	if task.Path != "" {
		return task.Path
	}
	return task.ID
}

func endpointRelativeID(sectionModule, taskModule, leaf string) string {
	if taskModule == sectionModule {
		return leaf
	}
	prefix := strings.TrimPrefix(taskModule, sectionModule+".")
	if prefix != taskModule {
		return strings.ReplaceAll(prefix, ".", "/") + "/" + leaf
	}
	return strings.ReplaceAll(taskModule, ".", "/") + "/" + leaf
}

func endpointFullPath(root, relative string) string {
	root = strings.TrimRight(root, "/")
	relative = strings.TrimLeft(relative, "/")
	if root == "" || root == "/" {
		return "/" + relative
	}
	return root + "/" + relative
}

func endpointParams(task *semantic.Task) []string {
	params := make([]string, 0, len(task.Params))
	for _, param := range task.Params {
		params = append(params, param.Model.String())
	}
	return params
}

func endpointReturns(task *semantic.Task) string {
	if task.Returns == nil {
		return ""
	}
	return task.Returns.Model.String()
}

func moduleForFileID(fileID semantic.FileID) string {
	parts := strings.Split(fileID.String(), "/")
	if len(parts) == 0 {
		return ""
	}
	for i, part := range parts[:len(parts)-1] {
		if part == "model" || part == "store" || part == "task" || part == "state" || part == "event" || part == "branch" || part == "fork" || part == "join" {
			return strings.Join(parts[:i], ".")
		}
	}
	if len(parts) == 1 {
		return ""
	}
	return strings.Join(parts[:len(parts)-1], ".")
}
