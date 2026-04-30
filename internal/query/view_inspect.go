package query

import (
	"fmt"
	"sort"
	"strings"

	"github.com/hiroshiasayadev-prog/brewprint/internal/semantic"
)

type apiInspectModule struct {
	Module            string `json:"module"`
	IncludeSubmodules bool   `json:"include_submodules"`
	EndpointCount     int    `json:"endpoint_count"`
}

type apiInspectEndpoint struct {
	Module   string `json:"module"`
	Task     string `json:"task"`
	Method   string `json:"method"`
	Path     string `json:"path"`
	LeafPath string `json:"leaf_path"`
}

type erInspectModule struct {
	Module     string `json:"module"`
	StoreCount int    `json:"store_count"`
	ModelCount int    `json:"model_count"`
}

type erInspectFKRelation struct {
	FromModel   string `json:"from_model"`
	FromField   string `json:"from_field"`
	ToModel     string `json:"to_model"`
	ToField     string `json:"to_field"`
	FK          string `json:"fk"`
	Cardinality string `json:"cardinality"`
	Unique      bool   `json:"unique,omitempty"`
}

type erInspectExcludedRefsSummary struct {
	Count int                    `json:"count"`
	Refs  []erInspectExcludedRef `json:"refs,omitempty"`
}

type erInspectExcludedRef struct {
	FromModel   string `json:"from_model"`
	FromField   string `json:"from_field"`
	FK          string `json:"fk"`
	TargetModel string `json:"target_model,omitempty"`
	Reason      string `json:"reason"`
}

type erInspectSummary struct {
	Modules   []erInspectModule
	Stores    []ObjectRef
	Models    []ObjectRef
	Relations []erInspectFKRelation
	Excluded  erInspectExcludedRefsSummary
}

func (s *Service) apiViewBySelector(selector Selector) (*semantic.APIView, error) {
	if err := s.requireProject(); err != nil {
		return nil, err
	}
	if selector.ID == "" {
		return nil, fmt.Errorf("selector.id is required")
	}
	if selector.Object != "" && selector.Object != "view" {
		return nil, fmt.Errorf("unsupported selector object for API view: %s", selector.Object)
	}
	if selector.Kind != "" && selector.Kind != "api_table" {
		return nil, fmt.Errorf("unsupported selector kind for API view: %s", selector.Kind)
	}
	view := s.project.APIViewsByID[selector.ID]
	if view == nil {
		return nil, fmt.Errorf("object not found: %s", selector.ID)
	}
	return view, nil
}

func (s *Service) isAPIViewSelector(selector Selector) bool {
	if selector.Kind == "api_table" {
		return true
	}
	if selector.Object != "view" {
		return false
	}
	if selector.Kind != "" {
		return selector.Kind == "api_table"
	}
	if s == nil || s.project == nil || selector.ID == "" {
		return false
	}
	return s.project.APIViewsByID[selector.ID] != nil
}

func (s *Service) erViewBySelector(selector Selector) (*semantic.ERView, error) {
	if err := s.requireProject(); err != nil {
		return nil, err
	}
	if selector.ID == "" {
		return nil, fmt.Errorf("selector.id is required")
	}
	if selector.Object != "" && selector.Object != "view" {
		return nil, fmt.Errorf("unsupported selector object for ER view: %s", selector.Object)
	}
	if selector.Kind != "" && selector.Kind != "er_diagram" {
		return nil, fmt.Errorf("unsupported selector kind for ER view: %s", selector.Kind)
	}
	view := s.project.ERViewsByID[selector.ID]
	if view == nil {
		return nil, fmt.Errorf("object not found: %s", selector.ID)
	}
	return view, nil
}

func (s *Service) isERViewSelector(selector Selector) bool {
	if selector.Kind == "er_diagram" {
		return true
	}
	if selector.Object != "view" {
		return false
	}
	if selector.Kind != "" {
		return selector.Kind == "er_diagram"
	}
	if s == nil || s.project == nil || selector.ID == "" {
		return false
	}
	return s.project.ERViewsByID[selector.ID] != nil
}

func apiViewObjectRef(view *semantic.APIView) ObjectRef {
	if view == nil {
		return ObjectRef{}
	}
	return ObjectRef{
		Object: "view",
		Kind:   "api_table",
		ID:     view.ID,
		Label:  view.ID,
		File:   view.FileID.String(),
	}
}

func erViewObjectRef(view *semantic.ERView) ObjectRef {
	if view == nil {
		return ObjectRef{}
	}
	return ObjectRef{
		Object: "view",
		Kind:   "er_diagram",
		ID:     view.ID,
		Label:  view.ID,
		File:   view.FileID.String(),
	}
}

func signatureForAPIView(view *semantic.APIView) Signature {
	modules := make([]map[string]any, 0, len(view.Modules))
	for _, module := range view.Modules {
		modules = append(modules, map[string]any{
			"module":             module.Module,
			"include_submodules": module.IncludeSubmodules,
		})
	}
	return Signature{
		"id":             view.ID,
		"http_root_path": view.HTTPRootPath,
		"modules":        modules,
	}
}

func signatureForERView(view *semantic.ERView) Signature {
	modules := make([]map[string]any, 0, len(view.Modules))
	for _, module := range view.Modules {
		modules = append(modules, map[string]any{"module": module.Module})
	}
	return Signature{
		"id":      view.ID,
		"modules": modules,
	}
}

func (s *Service) inspectAPIView(req InspectRequest) (InspectResponse, error) {
	view, err := s.apiViewBySelector(req.Selector)
	if err != nil {
		return InspectResponse{}, err
	}
	table := s.endpointTable(view)
	return InspectResponse{
		Object:    apiViewObjectRef(view),
		Signature: signatureForAPIView(view),
		Doc:       view.Note,
		Source:    sourceMap(view.FileID),
		Members: map[string]any{
			"modules":             apiInspectModules(view, table),
			"sections":            table.Sections,
			"collected_endpoints": apiInspectEndpoints(table),
		},
		Diagnostics: []semantic.Diagnostic{},
	}, nil
}

func apiInspectModules(view *semantic.APIView, table EndpointTable) []apiInspectModule {
	endpointCounts := map[string]int{}
	for _, section := range table.Sections {
		endpointCounts[section.Module] = len(section.Endpoints)
	}
	modules := make([]apiInspectModule, 0, len(view.Modules))
	for _, module := range view.Modules {
		modules = append(modules, apiInspectModule{
			Module:            module.Module,
			IncludeSubmodules: module.IncludeSubmodules,
			EndpointCount:     endpointCounts[module.Module],
		})
	}
	return modules
}

func apiInspectEndpoints(table EndpointTable) []apiInspectEndpoint {
	var endpoints []apiInspectEndpoint
	for _, section := range table.Sections {
		for _, endpoint := range section.Endpoints {
			endpoints = append(endpoints, apiInspectEndpoint{
				Module:   section.Module,
				Task:     endpoint.Task,
				Method:   endpoint.Method,
				Path:     endpoint.Path,
				LeafPath: endpoint.LeafPath,
			})
		}
	}
	return endpoints
}

func (s *Service) inspectERView(req InspectRequest) (InspectResponse, error) {
	view, err := s.erViewBySelector(req.Selector)
	if err != nil {
		return InspectResponse{}, err
	}
	summary := s.erInspectSummary(view)
	return InspectResponse{
		Object:    erViewObjectRef(view),
		Signature: signatureForERView(view),
		Doc:       view.Note,
		Source:    sourceMap(view.FileID),
		Members: map[string]any{
			"modules":               summary.Modules,
			"included_stores":       summary.Stores,
			"included_models":       summary.Models,
			"fk_relations":          summary.Relations,
			"excluded_refs_summary": summary.Excluded,
		},
		Diagnostics: []semantic.Diagnostic{},
	}, nil
}

func (s *Service) erInspectSummary(view *semantic.ERView) erInspectSummary {
	stores := s.erStoresForView(view)
	models := s.erModelsFromStores(stores)
	includedModels := map[semantic.QualifiedID]struct{}{}
	for _, model := range models {
		includedModels[model.QID] = struct{}{}
	}
	relations, excluded := erInspectRelations(models, includedModels)

	storeRefs := make([]ObjectRef, 0, len(stores))
	for _, store := range stores {
		storeRefs = append(storeRefs, objectRef(store))
	}
	modelRefs := make([]ObjectRef, 0, len(models))
	for _, model := range models {
		modelRefs = append(modelRefs, objectRef(model))
	}

	return erInspectSummary{
		Modules:   s.erInspectModules(view),
		Stores:    storeRefs,
		Models:    modelRefs,
		Relations: relations,
		Excluded:  excluded,
	}
}

func (s *Service) erInspectModules(view *semantic.ERView) []erInspectModule {
	modules := make([]erInspectModule, 0, len(view.Modules))
	for _, viewModule := range view.Modules {
		stores := s.erDBStoresForModule(viewModule.Module)
		models := s.erModelsFromStores(stores)
		modules = append(modules, erInspectModule{
			Module:     viewModule.Module,
			StoreCount: len(stores),
			ModelCount: len(models),
		})
	}
	return modules
}

func (s *Service) erStoresForView(view *semantic.ERView) []*semantic.Store {
	var stores []*semantic.Store
	for _, viewModule := range view.Modules {
		stores = append(stores, s.erDBStoresForModule(viewModule.Module)...)
	}
	return stores
}

func (s *Service) erDBStoresForModule(module string) []*semantic.Store {
	var stores []*semantic.Store
	for _, store := range s.project.StoresByQID {
		if store.StoreKind != "db" || moduleForFileID(store.FileID) != module {
			continue
		}
		stores = append(stores, store)
	}
	sort.Slice(stores, func(i, j int) bool {
		if stores[i].Of != stores[j].Of {
			return stores[i].Of < stores[j].Of
		}
		return stores[i].QID < stores[j].QID
	})
	return stores
}

func (s *Service) erModelsFromStores(stores []*semantic.Store) []*semantic.Model {
	seen := map[semantic.QualifiedID]struct{}{}
	var models []*semantic.Model
	for _, store := range stores {
		model := s.project.ModelsByQID[store.Of]
		if model == nil || model.Kind != "struct" {
			continue
		}
		if _, ok := seen[model.QID]; ok {
			continue
		}
		seen[model.QID] = struct{}{}
		models = append(models, model)
	}
	sort.Slice(models, func(i, j int) bool { return models[i].QID < models[j].QID })
	return models
}

func erInspectRelations(models []*semantic.Model, includedModels map[semantic.QualifiedID]struct{}) ([]erInspectFKRelation, erInspectExcludedRefsSummary) {
	seenRelations := map[string]struct{}{}
	var relations []erInspectFKRelation
	var excluded []erInspectExcludedRef
	for _, model := range models {
		for _, field := range model.Fields {
			if field.FK == "" {
				continue
			}
			targetModel := erModelQIDFromFK(model, field.FK)
			if _, ok := includedModels[targetModel]; !ok {
				excluded = append(excluded, erInspectExcludedRef{
					FromModel:   model.QID.String(),
					FromField:   field.Name,
					FK:          field.FK,
					TargetModel: targetModel.String(),
					Reason:      "target_model_not_in_view",
				})
				continue
			}
			cardinality := "many_to_one"
			if field.Unique {
				cardinality = "one_to_one"
			}
			key := model.QID.String() + "\x00" + field.Name + "\x00" + targetModel.String() + "\x00" + fmt.Sprint(field.Unique)
			if _, ok := seenRelations[key]; ok {
				continue
			}
			seenRelations[key] = struct{}{}
			relations = append(relations, erInspectFKRelation{
				FromModel:   model.QID.String(),
				FromField:   field.Name,
				ToModel:     targetModel.String(),
				ToField:     erFieldNameFromFK(field.FK),
				FK:          field.FK,
				Cardinality: cardinality,
				Unique:      field.Unique,
			})
		}
	}
	sort.Slice(relations, func(i, j int) bool {
		if relations[i].FromModel != relations[j].FromModel {
			return relations[i].FromModel < relations[j].FromModel
		}
		if relations[i].FromField != relations[j].FromField {
			return relations[i].FromField < relations[j].FromField
		}
		return relations[i].ToModel < relations[j].ToModel
	})
	sort.Slice(excluded, func(i, j int) bool {
		if excluded[i].FromModel != excluded[j].FromModel {
			return excluded[i].FromModel < excluded[j].FromModel
		}
		return excluded[i].FromField < excluded[j].FromField
	})
	return relations, erInspectExcludedRefsSummary{Count: len(excluded), Refs: excluded}
}

func erModelQIDFromFK(source *semantic.Model, fk string) semantic.QualifiedID {
	parts := strings.Split(fk, ".")
	for i, part := range parts {
		if part == "model" && i+1 < len(parts) {
			return semantic.QualifiedID(strings.Join(parts[:i+2], "."))
		}
	}
	if len(parts) == 2 && source != nil {
		module := moduleForFileID(source.FileID)
		if module != "" {
			return semantic.QualifiedID(module + ".model." + parts[0])
		}
	}
	if len(parts) >= 2 {
		return semantic.QualifiedID(strings.Join(parts[:len(parts)-1], "."))
	}
	return ""
}

func erFieldNameFromFK(fk string) string {
	parts := strings.Split(fk, ".")
	if len(parts) == 0 {
		return ""
	}
	return parts[len(parts)-1]
}
