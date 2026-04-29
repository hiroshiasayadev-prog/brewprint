package resolve

import (
	"fmt"
	"strings"

	"github.com/hiroshiasayadev-prog/brewprint/internal/semantic"
)

const (
	diagnosticSemanticValidation        = "semantic_validation"
	diagnosticInvalidModelID            = "invalid_model_id"
	diagnosticUnresolvedModel           = "unresolved_model"
	diagnosticUnresolvedFieldType       = "unresolved_field_type"
	diagnosticUnresolvedFK              = "unresolved_fk"
	diagnosticUnresolvedStore           = "unresolved_store"
	diagnosticInvalidEndpoint           = "invalid_endpoint"
	diagnosticInvalidStoreKind          = "invalid_store_kind"
	diagnosticInvalidModelKind          = "invalid_model_kind"
	diagnosticDuplicateModelField       = "duplicate_model_field"
	diagnosticDuplicatePrimaryKey       = "duplicate_primary_key"
	diagnosticMissingRequiredField      = "missing_required_field"
	diagnosticDuplicateNode             = "duplicate_node"
	diagnosticDuplicateMainNode         = "duplicate_main_node"
	diagnosticDuplicateActor            = "duplicate_actor"
	diagnosticDuplicateInitializedStore = "duplicate_initialized_store"
	diagnosticUnsupportedFlowEntry      = "unsupported_flow_entry"
	diagnosticUnresolvedFlowTask        = "unresolved_flow_task"
	diagnosticUnresolvedFlowNode        = "unresolved_flow_node"
	diagnosticInvalidFlowBranch         = "invalid_flow_branch"
	diagnosticUnmatchedJoinParam        = "unmatched_join_param"
	diagnosticUnresolvedTransitionState = "unresolved_transition_state"
	diagnosticUnresolvedTransitionEvent = "unresolved_transition_event"
	diagnosticDuplicateTransition       = "duplicate_transition"
	diagnosticMissingTransitionGuard    = "missing_transition_guard"
	diagnosticDuplicateView             = "duplicate_view"
	diagnosticInvalidViewDefinition     = "invalid_view_definition"
	diagnosticDuplicateViewModule       = "duplicate_view_module"
	diagnosticUnresolvedSequenceStep    = "unresolved_sequence_step"
	diagnosticNonContinuousSequence     = "non_continuous_sequence"
)

var primitiveTypes = map[string]struct{}{
	"str":      {},
	"int":      {},
	"float":    {},
	"bool":     {},
	"bytes":    {},
	"datetime": {},
	"any":      {},
}

var validStoreKinds = map[string]struct{}{
	"db":         {},
	"session":    {},
	"collection": {},
	"context":    {},
}

var validModelKinds = map[string]struct{}{
	"struct": {},
	"list":   {},
	"dict":   {},
}

var validHTTPMethods = map[string]struct{}{
	"GET":    {},
	"POST":   {},
	"PUT":    {},
	"DELETE": {},
	"PATCH":  {},
}

func validateProject(project *semantic.Project, symbols *symbolTable) {
	validateRequiredNodeIDs(project, symbols)
	validateModelDefinitions(project, symbols)
	validateStoreDefinitions(project, symbols)
	validateTaskDefinitions(project, symbols)
	validateControlDefinitions(project, symbols)
	validateEventDefinitions(project, symbols)
}

func validateRequiredNodeIDs(project *semantic.Project, symbols *symbolTable) {
	for _, node := range project.NodesByQID {
		if node.GetID() == "" {
			symbols.addDiagnosticCode(semantic.SeverityError, diagnosticMissingRequiredField, node.GetFileID(), "node id is required: "+string(node.GetKind()))
		}
	}
}

func validateModelDefinitions(project *semantic.Project, symbols *symbolTable) {
	for _, model := range project.ModelsByQID {
		if isPrimitive(model.ID) {
			symbols.addDiagnosticCode(semantic.SeverityError, diagnosticInvalidModelID, model.FileID, "model id uses primitive reserved word: "+model.ID)
		}
		if model.Kind == "" {
			symbols.addDiagnosticCode(semantic.SeverityError, diagnosticMissingRequiredField, model.FileID, "model kind is required: "+model.QID.String())
		} else if _, ok := validModelKinds[model.Kind]; !ok {
			symbols.addDiagnosticCode(semantic.SeverityError, diagnosticInvalidModelKind, model.FileID, "invalid model kind: "+model.Kind)
		}
		seenFields := map[string]struct{}{}
		pkCount := 0
		module := moduleForFileID(model.FileID)
		for _, field := range model.Fields {
			fieldLabel := field.Name
			if fieldLabel == "" {
				fieldLabel = "<unnamed>"
				symbols.addDiagnosticCode(semantic.SeverityError, diagnosticMissingRequiredField, model.FileID, "model field name is required: "+model.QID.String())
			}
			if field.Name != "" {
				if _, exists := seenFields[field.Name]; exists {
					symbols.addDiagnosticCode(semantic.SeverityError, diagnosticDuplicateModelField, model.FileID, "duplicate model field: "+model.QID.String()+"."+field.Name)
				}
				seenFields[field.Name] = struct{}{}
			}
			if field.PK {
				pkCount++
			}
			if field.Type == "" {
				symbols.addDiagnosticCode(semantic.SeverityError, diagnosticMissingRequiredField, model.FileID, "model field type is required: "+model.QID.String()+"."+fieldLabel)
			} else if !modelOrPrimitiveExists(project, module, field.Type) {
				symbols.addDiagnosticCode(semantic.SeverityError, diagnosticUnresolvedFieldType, model.FileID, "unresolved field type: "+model.QID.String()+"."+fieldLabel+" -> "+field.Type)
			}
			if field.FK != "" {
				if !fieldExists(project, module, field.FK) {
					symbols.addDiagnosticCode(semantic.SeverityError, diagnosticUnresolvedFK, model.FileID, "unresolved fk: "+model.QID.String()+"."+field.Name+" -> "+field.FK)
				}
			}
		}
		if pkCount > 1 {
			symbols.addDiagnosticCode(semantic.SeverityError, diagnosticDuplicatePrimaryKey, model.FileID, "model has multiple primary keys: "+model.QID.String())
		}
	}
}

func validateStoreDefinitions(project *semantic.Project, symbols *symbolTable) {
	for _, store := range project.StoresByQID {
		validateStore(project, symbols, store)
	}
	for _, storesByName := range project.StoresByFileLocal {
		for _, store := range storesByName {
			validateStore(project, symbols, store)
		}
	}
}

func validateStore(project *semantic.Project, symbols *symbolTable, store *semantic.Store) {
	if store.StoreKind == "" {
		symbols.addDiagnosticCode(semantic.SeverityError, diagnosticMissingRequiredField, store.FileID, "store kind is required: "+store.QID.String())
	} else if _, ok := validStoreKinds[store.StoreKind]; !ok {
		symbols.addDiagnosticCode(semantic.SeverityError, diagnosticInvalidStoreKind, store.FileID, "invalid store kind: "+store.StoreKind)
	}
	if store.OfName != "" && !modelExists(project, store.Of) {
		symbols.addDiagnosticCode(semantic.SeverityError, diagnosticUnresolvedModel, store.FileID, "unresolved store model: "+store.OfName)
	}
}

func validateTaskDefinitions(project *semantic.Project, symbols *symbolTable) {
	for _, task := range project.TasksByQID {
		validateParams(project, symbols, task.FileID, "task params", task.Params)
		if task.Returns != nil {
			validateReturn(project, symbols, task.FileID, task.QID.String(), "task return", task.Returns)
		}
		for _, init := range task.Initializes {
			if init.Name == "" {
				symbols.addDiagnosticCode(semantic.SeverityError, diagnosticMissingRequiredField, task.FileID, "initialized store name is required: "+task.QID.String())
			}
			if init.Model == "" {
				symbols.addDiagnosticCode(semantic.SeverityError, diagnosticMissingRequiredField, task.FileID, "initialized store model is required: "+task.QID.String()+"."+init.Name)
			} else if !modelExists(project, init.Model) {
				symbols.addDiagnosticCode(semantic.SeverityError, diagnosticUnresolvedModel, task.FileID, "unresolved initialized store model: "+init.Model.String())
			}
		}
		for _, read := range task.Reads {
			if read.Store == "" {
				symbols.addDiagnosticCode(semantic.SeverityError, diagnosticUnresolvedStore, task.FileID, "unresolved read store: "+read.Name)
			}
		}
		for _, write := range task.Writes {
			if write.Store == "" {
				symbols.addDiagnosticCode(semantic.SeverityError, diagnosticUnresolvedStore, task.FileID, "unresolved write store: "+write.Name)
			}
		}
		if task.Endpoint {
			if _, ok := validHTTPMethods[task.Method]; !ok {
				symbols.addDiagnosticCode(semantic.SeverityError, diagnosticInvalidEndpoint, task.FileID, "endpoint task has invalid or missing method: "+task.QID.String())
			}
			if strings.Contains(task.Path, "/") {
				symbols.addDiagnosticCode(semantic.SeverityError, diagnosticInvalidEndpoint, task.FileID, "endpoint path must be a single segment: "+task.Path)
			}
		}
	}
}

func validateControlDefinitions(project *semantic.Project, symbols *symbolTable) {
	for _, branch := range project.BranchesByQID {
		validateParams(project, symbols, branch.FileID, "branch params", branch.Params)
	}
	for _, fork := range project.ForksByQID {
		validateParams(project, symbols, fork.FileID, "fork params", fork.Params)
	}
	for _, join := range project.JoinsByQID {
		validateParams(project, symbols, join.FileID, "join params", join.Params)
		if join.Returns != nil {
			validateReturn(project, symbols, join.FileID, join.QID.String(), "join return", join.Returns)
		}
	}
}

func validateEventDefinitions(project *semantic.Project, symbols *symbolTable) {
	for _, event := range project.EventsByQID {
		if event.PayloadName != "" && !modelExists(project, event.PayloadModel) {
			symbols.addDiagnosticCode(semantic.SeverityError, diagnosticUnresolvedModel, event.FileID, "unresolved event payload model: "+event.PayloadName)
		}
		if event.WatchesName != "" {
			if _, ok := project.StoresByQID[event.Watches]; !ok {
				symbols.addDiagnosticCode(semantic.SeverityError, diagnosticUnresolvedStore, event.FileID, "unresolved event watched store: "+event.WatchesName)
			}
		}
	}
}

func validateParams(project *semantic.Project, symbols *symbolTable, fileID semantic.FileID, context string, params []semantic.Param) {
	for _, param := range params {
		module := moduleForFileID(fileID)
		if param.Name == "" {
			symbols.addDiagnosticCode(semantic.SeverityError, diagnosticMissingRequiredField, fileID, fmt.Sprintf("%s name is required", context))
		}
		if param.ModelName == "" {
			symbols.addDiagnosticCode(semantic.SeverityError, diagnosticMissingRequiredField, fileID, fmt.Sprintf("%s model is required: %s", context, param.Name))
		} else if !modelOrPrimitiveExists(project, module, param.ModelName) {
			symbols.addDiagnosticCode(semantic.SeverityError, diagnosticUnresolvedModel, fileID, fmt.Sprintf("unresolved %s model: %s", context, param.ModelName))
		}
	}
}

func validateReturn(project *semantic.Project, symbols *symbolTable, fileID semantic.FileID, owner string, context string, ret *semantic.Return) {
	if ret.Name == "" {
		symbols.addDiagnosticCode(semantic.SeverityError, diagnosticMissingRequiredField, fileID, context+" name is required: "+owner)
	}
	if ret.ModelName == "" {
		symbols.addDiagnosticCode(semantic.SeverityError, diagnosticMissingRequiredField, fileID, context+" model is required: "+owner+"."+ret.Name)
	} else if !modelExists(project, ret.Model) {
		symbols.addDiagnosticCode(semantic.SeverityError, diagnosticUnresolvedModel, fileID, "unresolved "+context+" model: "+ret.ModelName)
	}
}

func modelOrPrimitiveExists(project *semantic.Project, module string, raw string) bool {
	if raw == "" {
		return false
	}
	if isPrimitive(raw) {
		return true
	}
	return modelExists(project, resolveModelQID(module, raw))
}

func modelExists(project *semantic.Project, qid semantic.QualifiedID) bool {
	if qid == "" {
		return false
	}
	_, ok := project.ModelsByQID[qid]
	return ok
}

func fieldExists(project *semantic.Project, module string, raw string) bool {
	modelQID, fieldName := resolveFK(module, raw)
	if modelQID == "" || fieldName == "" {
		return false
	}
	model := project.ModelsByQID[modelQID]
	if model == nil {
		return false
	}
	for _, field := range model.Fields {
		if field.Name == fieldName {
			return true
		}
	}
	return false
}

func isPrimitive(raw string) bool {
	_, ok := primitiveTypes[raw]
	return ok
}
