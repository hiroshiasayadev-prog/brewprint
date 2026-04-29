package query

import (
	"fmt"
	"sort"

	"github.com/hiroshiasayadev-prog/brewprint/internal/semantic"
)

type Service struct {
	project *semantic.Project
}

func NewService(project *semantic.Project) *Service {
	return &Service{project: project}
}

func (s *Service) requireProject() error {
	if s == nil || s.project == nil {
		return fmt.Errorf("query service has no project")
	}
	return nil
}

func (s *Service) nodeByID(id string) (semantic.Node, error) {
	if err := s.requireProject(); err != nil {
		return nil, err
	}
	if id == "" {
		return nil, fmt.Errorf("selector.id is required")
	}
	node := s.project.NodesByQID[semantic.QualifiedID(id)]
	if node == nil {
		return nil, fmt.Errorf("object not found: %s", id)
	}
	return node, nil
}

func (s *Service) scenarioBySelector(selector Selector) (*semantic.SequenceScenario, error) {
	if err := s.requireProject(); err != nil {
		return nil, err
	}
	if selector.ID == "" {
		return nil, fmt.Errorf("selector.id is required")
	}
	if selector.Object != "" && selector.Object != "view" {
		return nil, fmt.Errorf("unsupported selector object for scenario: %s", selector.Object)
	}
	if selector.Kind != "" && selector.Kind != "sequence_diagram" {
		return nil, fmt.Errorf("unsupported selector kind for scenario: %s", selector.Kind)
	}
	scenario := s.project.ScenariosByID[selector.ID]
	if scenario == nil {
		return nil, fmt.Errorf("object not found: %s", selector.ID)
	}
	return scenario, nil
}

func (s *Service) isScenarioSelector(selector Selector) bool {
	if selector.Object == "view" || selector.Kind == "sequence_diagram" {
		return true
	}
	if s == nil || s.project == nil || selector.ID == "" {
		return false
	}
	return s.project.ScenariosByID[selector.ID] != nil
}

func (s *Service) transitionBySelector(selector Selector) (semantic.Transition, error) {
	if err := s.requireProject(); err != nil {
		return semantic.Transition{}, err
	}
	if selector.ID == "" {
		return semantic.Transition{}, fmt.Errorf("selector.id is required")
	}
	if selector.Object != "" && selector.Object != "transition" {
		return semantic.Transition{}, fmt.Errorf("unsupported selector object for transition: %s", selector.Object)
	}
	if selector.Kind != "" && selector.Kind != "transition" {
		return semantic.Transition{}, fmt.Errorf("unsupported selector kind for transition: %s", selector.Kind)
	}
	for _, transitions := range s.project.TransitionsByFile {
		for _, transition := range transitions {
			if semantic.TransitionID(transition) == selector.ID {
				return transition, nil
			}
		}
	}
	return semantic.Transition{}, fmt.Errorf("object not found: %s", selector.ID)
}

func (s *Service) isTransitionSelector(selector Selector) bool {
	return selector.Object == "transition" || selector.Kind == "transition"
}

func (s *Service) modelFieldBySelector(selector Selector) (*semantic.Model, semantic.ModelField, error) {
	if err := s.requireProject(); err != nil {
		return nil, semantic.ModelField{}, err
	}
	modelID := selector.ID
	fieldName := selector.LocalID
	if selector.File != "" && selector.LocalID != "" {
		modelID = selector.File
		fieldName = selector.LocalID
	}
	if modelID == "" || fieldName == "" {
		return nil, semantic.ModelField{}, fmt.Errorf("selector.id and selector.local_id are required for field")
	}
	if selector.Object != "" && selector.Object != "field" {
		return nil, semantic.ModelField{}, fmt.Errorf("unsupported selector object for field: %s", selector.Object)
	}
	if selector.Kind != "" && selector.Kind != "field" {
		return nil, semantic.ModelField{}, fmt.Errorf("unsupported selector kind for field: %s", selector.Kind)
	}
	model := s.project.ModelsByQID[semantic.QualifiedID(modelID)]
	if model == nil {
		return nil, semantic.ModelField{}, fmt.Errorf("object not found: %s", modelID)
	}
	for _, field := range model.Fields {
		if field.Name == fieldName {
			return model, field, nil
		}
	}
	return nil, semantic.ModelField{}, fmt.Errorf("object not found: %s.%s", modelID, fieldName)
}

func (s *Service) isFieldSelector(selector Selector) bool {
	return selector.Object == "field" || selector.Kind == "field"
}

func (s *Service) fileBySelector(selector Selector) (semantic.FileID, error) {
	if err := s.requireProject(); err != nil {
		return "", err
	}
	id := selector.ID
	if id == "" {
		id = selector.File
	}
	if id == "" {
		return "", fmt.Errorf("selector.id is required")
	}
	if selector.Object != "" && selector.Object != "file" {
		return "", fmt.Errorf("unsupported selector object for file: %s", selector.Object)
	}
	if selector.Kind != "" && selector.Kind != "state_file" {
		return "", fmt.Errorf("unsupported selector kind for file: %s", selector.Kind)
	}
	fileID := semantic.FileID(id)
	if _, ok := s.project.TransitionsByFile[fileID]; ok {
		return fileID, nil
	}
	if _, ok := s.project.NodesByFile[fileID]; ok {
		return fileID, nil
	}
	for _, scenario := range s.project.ScenariosByID {
		if scenario.StateFile == fileID || scenario.FileID == fileID {
			return fileID, nil
		}
	}
	return "", fmt.Errorf("object not found: %s", id)
}

func (s *Service) isFileSelector(selector Selector) bool {
	return selector.Object == "file" || selector.Kind == "state_file"
}

func objectRef(node semantic.Node) ObjectRef {
	return ObjectRef{
		Object:      "node",
		Kind:        string(node.GetKind()),
		ID:          node.GetQID().String(),
		QualifiedID: node.GetQID().String(),
		Label:       node.GetID(),
		File:        node.GetFileID().String(),
	}
}

func scenarioObjectRef(scenario *semantic.SequenceScenario) ObjectRef {
	if scenario == nil {
		return ObjectRef{}
	}
	return ObjectRef{
		Object: "view",
		Kind:   "sequence_diagram",
		ID:     scenario.ID,
		Label:  scenario.Title,
		File:   scenario.FileID.String(),
	}
}

func transitionObjectRef(transition semantic.Transition) ObjectRef {
	return ObjectRef{
		Object:  "transition",
		Kind:    "transition",
		ID:      semantic.TransitionID(transition),
		File:    transition.FileID.String(),
		LocalID: transition.From + ":" + transition.On,
	}
}

func fieldObjectRef(model *semantic.Model, field semantic.ModelField) ObjectRef {
	if model == nil {
		return ObjectRef{}
	}
	return ObjectRef{
		Object:      "field",
		Kind:        "field",
		ID:          model.QID.String() + "." + field.Name,
		QualifiedID: model.QID.String(),
		Label:       field.Name,
		File:        model.FileID.String(),
		LocalID:     field.Name,
	}
}

func fileObjectRef(fileID semantic.FileID, kind string) ObjectRef {
	if kind == "" {
		kind = "file"
	}
	return ObjectRef{
		Object: "file",
		Kind:   kind,
		ID:     fileID.String(),
		File:   fileID.String(),
	}
}

func sourceMap(fileID semantic.FileID) map[string]string {
	if fileID == "" {
		return nil
	}
	return map[string]string{"file": fileID.String()}
}

func assetRef(asset *semantic.Asset) *AssetRef {
	if asset == nil {
		return nil
	}
	return &AssetRef{
		Object:    "asset",
		Name:      asset.Name,
		Producer:  asset.ProducedBy.String(),
		Model:     asset.Model.String(),
		ScopeFile: asset.FileID.String(),
	}
}

func paramSignatures(params []semantic.Param) []ParamSignature {
	out := make([]ParamSignature, 0, len(params))
	for _, param := range params {
		out = append(out, ParamSignature{Name: param.Name, Model: param.Model.String(), Doc: param.Note})
	}
	return out
}

func returnSignature(ret *semantic.Return) *ReturnSignature {
	if ret == nil {
		return nil
	}
	return &ReturnSignature{Name: ret.Name, Model: ret.Model.String(), Asset: assetRef(ret.Asset)}
}

func storeIDs(refs []semantic.StoreRef) []string {
	out := make([]string, 0, len(refs))
	for _, ref := range refs {
		if ref.Store == "" {
			continue
		}
		out = append(out, ref.Store.String())
	}
	sort.Strings(out)
	return out
}

func fieldSignatures(fields []semantic.ModelField) []FieldSignature {
	out := make([]FieldSignature, 0, len(fields))
	for _, field := range fields {
		out = append(out, FieldSignature{
			Name:   field.Name,
			Type:   field.Type,
			PK:     field.PK,
			FK:     field.FK,
			Unique: field.Unique,
			Doc:    field.Note,
		})
	}
	return out
}
