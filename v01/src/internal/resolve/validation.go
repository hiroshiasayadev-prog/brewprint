package resolve

import (
	"fmt"
	"strings"

	"github.com/hiroshiasayadev-prog/brewprint/v01/src/internal/semantic"
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
	diagnosticInvalidTypeRef            = "invalid_type_ref"
	diagnosticOpaqueTypeRef             = "opaque_type_ref"
	diagnosticInvalidEnumModel          = "invalid_enum_model"
	diagnosticDuplicateEnumValue        = "duplicate_enum_value"
	diagnosticInvalidTaggedUnionModel   = "invalid_tagged_union_model"
	diagnosticDuplicateVariantTag       = "duplicate_variant_tag"
	diagnosticInvalidVariantField       = "invalid_variant_field"
	diagnosticDuplicateModelID          = "duplicate_model_id"
	diagnosticInvalidPrivateModelRef    = "invalid_private_model_reference"
	diagnosticDuplicateNode             = "duplicate_node"
	diagnosticDuplicateSubNode          = "duplicate_sub_node"
	diagnosticDuplicateMainNode         = "duplicate_main_node"
	diagnosticDuplicateActor            = "duplicate_actor"
	diagnosticDuplicateInitializedStore = "duplicate_initialized_store"
	diagnosticUnsupportedFlowEntry      = "unsupported_flow_entry"
	diagnosticUnresolvedFlowTask        = "unresolved_flow_task"
	diagnosticUnresolvedFlowNode        = "unresolved_flow_node"
	diagnosticInvalidFlowBranch         = "invalid_flow_branch"
	diagnosticUnmatchedJoinParam        = "unmatched_join_param"
	diagnosticIncompatibleWiringType    = "incompatible_wiring_type"
	diagnosticInvalidWiringSource       = "invalid_wiring_source"
	diagnosticUnresolvedWiringSource    = "unresolved_wiring_source"
	diagnosticInvalidForeachOverType    = "invalid_foreach_over_type"
	diagnosticInvalidForeachReturns     = "invalid_foreach_returns"
	diagnosticDuplicateFlowSource       = "duplicate_flow_source"
	diagnosticUnresolvedReturnSource    = "unresolved_return_source"
	diagnosticInvalidReturnSource       = "invalid_return_source"
	diagnosticIncompatibleReturnType    = "incompatible_return_type"
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
	"struct":       {},
	"list":         {},
	"dict":         {},
	"enum":         {},
	"tagged_union": {},
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
	validateModelIDCollisions(project, symbols)
	validateModelDefinitions(project, symbols)
	validateStoreDefinitions(project, symbols)
	validateTaskDefinitions(project, symbols)
	validateControlDefinitions(project, symbols)
	validateEventDefinitions(project, symbols)
	validateFlowWiringTypes(project, symbols)
}

func validateRequiredNodeIDs(project *semantic.Project, symbols *symbolTable) {
	for _, nodes := range project.NodesByFile {
		for _, node := range nodes {
			if node.GetID() == "" {
				symbols.addDiagnosticCode(semantic.SeverityError, diagnosticMissingRequiredField, node.GetFileID(), "node id is required: "+string(node.GetKind()))
			}
		}
	}
}

func validateModelIDCollisions(project *semantic.Project, symbols *symbolTable) {
	for fileID, byName := range project.PrivateModelsByFile {
		module := moduleForFileID(fileID)
		for id := range byName {
			publicQID := resolveModelQID(module, id)
			if project.ModelsByQID[publicQID] != nil {
				symbols.addDiagnosticCode(semantic.SeverityError, diagnosticDuplicateModelID, fileID, "private model id conflicts with public model in same module: "+id)
			}
			for _, node := range project.NodesByFile[fileID] {
				if node.GetID() != id {
					continue
				}
				if _, ok := node.(*semantic.Model); ok {
					continue
				}
				symbols.addDiagnosticCode(semantic.SeverityError, diagnosticDuplicateModelID, fileID, "private model id conflicts with node local id in same file: "+id)
				break
			}
		}
	}
}

func validateModelDefinitions(project *semantic.Project, symbols *symbolTable) {
	for _, model := range allModels(project) {
		if isTaskFileID(model.FileID) && model.Main {
			symbols.addDiagnosticCode(semantic.SeverityError, diagnosticSemanticValidation, model.FileID, "task file model must not be main: "+model.ID)
		}
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
		if model.Kind == "list" {
			if model.Element == "" {
				symbols.addDiagnosticCode(semantic.SeverityError, diagnosticMissingRequiredField, model.FileID, "model element is required: "+model.QID.String())
			} else {
				validateTypeRef(project, symbols, model.FileID, module, model.Element, "model.element "+model.QID.String(), diagnosticUnresolvedModel, "unresolved model element: "+model.Element)
			}
		}
		if model.Kind == "dict" {
			if model.Value == "" {
				symbols.addDiagnosticCode(semantic.SeverityError, diagnosticMissingRequiredField, model.FileID, "model value is required: "+model.QID.String())
			} else {
				validateTypeRef(project, symbols, model.FileID, module, model.Value, "model.value "+model.QID.String(), diagnosticUnresolvedModel, "unresolved model value: "+model.Value)
			}
		}
		if model.Kind == "enum" {
			validateEnumModel(symbols, model)
		}
		if model.Kind == "tagged_union" {
			validateTaggedUnionModel(project, symbols, model)
		}
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
			} else {
				validateTypeRef(project, symbols, model.FileID, module, field.Type, "fields[].type "+model.QID.String()+"."+fieldLabel, diagnosticUnresolvedFieldType, "unresolved field type: "+model.QID.String()+"."+fieldLabel+" -> "+field.Type)
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

func validateEnumModel(symbols *symbolTable, model *semantic.Model) {
	if len(model.Fields) > 0 || model.Element != "" || model.Value != "" {
		symbols.addDiagnosticCode(semantic.SeverityError, diagnosticInvalidEnumModel, model.FileID, "enum model must not define fields, element, or value: "+model.QID.String())
	}
	if len(model.Values) == 0 {
		symbols.addDiagnosticCode(semantic.SeverityError, diagnosticInvalidEnumModel, model.FileID, "enum model values are required and must be non-empty: "+model.QID.String())
		return
	}
	seen := map[string]struct{}{}
	for _, value := range model.Values {
		if value == "" {
			symbols.addDiagnosticCode(semantic.SeverityError, diagnosticInvalidEnumModel, model.FileID, "enum model values must not contain empty string: "+model.QID.String())
			continue
		}
		if _, exists := seen[value]; exists {
			symbols.addDiagnosticCode(semantic.SeverityError, diagnosticDuplicateEnumValue, model.FileID, "duplicate enum value: "+model.QID.String()+"."+value)
			continue
		}
		seen[value] = struct{}{}
	}
}

func validateTaggedUnionModel(project *semantic.Project, symbols *symbolTable, model *semantic.Model) {
	if len(model.Fields) > 0 || model.Element != "" || model.Value != "" || len(model.Values) > 0 {
		symbols.addDiagnosticCode(semantic.SeverityError, diagnosticInvalidTaggedUnionModel, model.FileID, "tagged union model must not define fields, element, value, or values: "+model.QID.String())
	}
	if model.Discriminator == "" {
		symbols.addDiagnosticCode(semantic.SeverityError, diagnosticInvalidTaggedUnionModel, model.FileID, "tagged union discriminator is required: "+model.QID.String())
	} else if strings.Contains(model.Discriminator, ".") {
		symbols.addDiagnosticCode(semantic.SeverityError, diagnosticInvalidTaggedUnionModel, model.FileID, "tagged union discriminator must not be a dot path: "+model.QID.String())
	}
	if len(model.Variants) == 0 {
		symbols.addDiagnosticCode(semantic.SeverityError, diagnosticInvalidTaggedUnionModel, model.FileID, "tagged union variants are required and must be non-empty: "+model.QID.String())
		return
	}
	module := moduleForFileID(model.FileID)
	seenTags := map[string]struct{}{}
	for _, variant := range model.Variants {
		if variant.Tag == "" {
			symbols.addDiagnosticCode(semantic.SeverityError, diagnosticInvalidTaggedUnionModel, model.FileID, "tagged union variant tag is required and must be non-empty: "+model.QID.String())
			continue
		}
		if _, exists := seenTags[variant.Tag]; exists {
			symbols.addDiagnosticCode(semantic.SeverityError, diagnosticDuplicateVariantTag, model.FileID, "duplicate variant tag: "+model.QID.String()+"."+variant.Tag)
		} else {
			seenTags[variant.Tag] = struct{}{}
		}
		if variant.Fields == nil {
			symbols.addDiagnosticCode(semantic.SeverityError, diagnosticInvalidVariantField, model.FileID, "tagged union variant fields are required: "+model.QID.String()+"."+variant.Tag)
			continue
		}
		seenFieldNames := map[string]struct{}{}
		for _, field := range variant.Fields {
			fieldLabel := field.Name
			if fieldLabel == "" {
				fieldLabel = "<unnamed>"
				symbols.addDiagnosticCode(semantic.SeverityError, diagnosticInvalidVariantField, model.FileID, "variant field name is required: "+model.QID.String()+"."+variant.Tag)
			}
			if field.Name != "" {
				if model.Discriminator != "" && field.Name == model.Discriminator {
					symbols.addDiagnosticCode(semantic.SeverityError, diagnosticInvalidVariantField, model.FileID, "variant field must not repeat discriminator: "+model.QID.String()+"."+variant.Tag+"."+field.Name)
				}
				if _, exists := seenFieldNames[field.Name]; exists {
					symbols.addDiagnosticCode(semantic.SeverityError, diagnosticDuplicateModelField, model.FileID, "duplicate variant payload field: "+model.QID.String()+"."+variant.Tag+"."+field.Name)
				}
				seenFieldNames[field.Name] = struct{}{}
			}
			if field.PK || field.FK != "" || field.Unique {
				symbols.addDiagnosticCode(semantic.SeverityError, diagnosticInvalidVariantField, model.FileID, "variant field must not use pk, fk, or unique: "+model.QID.String()+"."+variant.Tag+"."+fieldLabel)
			}
			if field.Type == "" {
				symbols.addDiagnosticCode(semantic.SeverityError, diagnosticInvalidVariantField, model.FileID, "variant field type is required: "+model.QID.String()+"."+variant.Tag+"."+fieldLabel)
			} else {
				validateTypeRef(project, symbols, model.FileID, module, field.Type, "variants[].fields[].type "+model.QID.String()+"."+variant.Tag+"."+fieldLabel, diagnosticUnresolvedFieldType, "unresolved variant field type: "+model.QID.String()+"."+variant.Tag+"."+fieldLabel+" -> "+field.Type)
			}
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
	seen := map[*semantic.Task]struct{}{}
	for _, task := range project.TasksByQID {
		if _, ok := seen[task]; ok {
			continue
		}
		seen[task] = struct{}{}
		validateParams(project, symbols, task.FileID, "task params", task.Params)
		if task.Returns != nil {
			validateReturn(project, symbols, task.FileID, task.QID.String(), "task return", task.Returns)
			validateTaskReturnSource(project, symbols, task)
		}
		for _, init := range task.Initializes {
			if init.Name == "" {
				symbols.addDiagnosticCode(semantic.SeverityError, diagnosticMissingRequiredField, task.FileID, "initialized store name is required: "+task.QID.String())
			}
			if init.ModelName == "" {
				symbols.addDiagnosticCode(semantic.SeverityError, diagnosticMissingRequiredField, task.FileID, "initialized store model is required: "+task.QID.String()+"."+init.Name)
			} else if _, err := parseInitializedModelRef(init.ModelName, moduleForFileID(task.FileID)); err != nil {
				symbols.addDiagnosticCode(semantic.SeverityError, diagnosticInvalidTypeRef, task.FileID, fmt.Sprintf("invalid initialized store model %q at %s.%s: %v", init.ModelName, task.QID.String(), init.Name, err))
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
	seenBranches := map[*semantic.Branch]struct{}{}
	for _, branch := range project.BranchesByQID {
		if _, ok := seenBranches[branch]; ok {
			continue
		}
		seenBranches[branch] = struct{}{}
		validateParams(project, symbols, branch.FileID, "branch params", branch.Params)
	}
	seenForks := map[*semantic.Fork]struct{}{}
	for _, fork := range project.ForksByQID {
		if _, ok := seenForks[fork]; ok {
			continue
		}
		seenForks[fork] = struct{}{}
		validateParams(project, symbols, fork.FileID, "fork params", fork.Params)
	}
	seenJoins := map[*semantic.Join]struct{}{}
	for _, join := range project.JoinsByQID {
		if _, ok := seenJoins[join]; ok {
			continue
		}
		seenJoins[join] = struct{}{}
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
		} else {
			if validateTypeRef(project, symbols, fileID, module, param.ModelName, context+" param "+param.Name, diagnosticUnresolvedModel, fmt.Sprintf("unresolved %s model: %s", context, param.ModelName)) {
				validateParamPrivateModelReference(project, symbols, fileID, context, param)
			}
		}
	}
}

func validateParamPrivateModelReference(project *semantic.Project, symbols *symbolTable, fileID semantic.FileID, context string, param semantic.Param) {
	model := firstTaskFilePrivateModelReference(project, fileID, param.TypeRef)
	if model == nil || privateModelIdentityInvalid(symbols, fileID, model.ID) {
		return
	}
	paramName := param.Name
	if paramName == "" {
		paramName = "<unnamed>"
	}
	symbols.addDiagnosticCode(
		semantic.SeverityError,
		diagnosticInvalidPrivateModelRef,
		fileID,
		fmt.Sprintf("invalid private model reference at %s param %s params[].model: %s resolves to private model %s", context, paramName, param.ModelName, model.QID),
	)
}

func firstTaskFilePrivateModelReference(project *semantic.Project, fileID semantic.FileID, ref *semantic.TypeRef) *semantic.Model {
	if ref == nil {
		return nil
	}
	switch ref.Kind {
	case semantic.TypeRefNamedModel:
		model := privateModelByQID(project, ref.Model)
		if model != nil && model.FileID == fileID && isTaskFileID(model.FileID) {
			return model
		}
	case semantic.TypeRefList:
		return firstTaskFilePrivateModelReference(project, fileID, ref.Elem)
	case semantic.TypeRefDict:
		return firstTaskFilePrivateModelReference(project, fileID, ref.Value)
	}
	return nil
}

func privateModelIdentityInvalid(symbols *symbolTable, fileID semantic.FileID, id string) bool {
	if id == "" {
		return false
	}
	for _, diagnostic := range symbols.diags {
		if diagnostic.Code == diagnosticDuplicateModelID && diagnostic.FileID == fileID && strings.HasSuffix(diagnostic.Message, ": "+id) {
			return true
		}
	}
	return false
}

func validateReturn(project *semantic.Project, symbols *symbolTable, fileID semantic.FileID, owner string, context string, ret *semantic.Return) {
	if ret.Name == "" {
		symbols.addDiagnosticCode(semantic.SeverityError, diagnosticMissingRequiredField, fileID, context+" name is required: "+owner)
	}
	if ret.ModelName == "" {
		symbols.addDiagnosticCode(semantic.SeverityError, diagnosticMissingRequiredField, fileID, context+" model is required: "+owner+"."+ret.Name)
	} else {
		module := moduleForFileID(fileID)
		validateTypeRef(project, symbols, fileID, module, ret.ModelName, context+" "+owner+"."+ret.Name, diagnosticUnresolvedModel, "unresolved "+context+" model: "+ret.ModelName)
	}
}

func validateTypeRef(project *semantic.Project, symbols *symbolTable, fileID semantic.FileID, module string, raw string, position string, unresolvedCode string, unresolvedMessage string) bool {
	ref, err := parseTypeRef(raw, module)
	if err != nil {
		symbols.addDiagnosticCode(semantic.SeverityError, diagnosticInvalidTypeRef, fileID, fmt.Sprintf("invalid TypeRef %q at %s: %v", raw, position, err))
		return false
	}
	if typeRefHasOpaqueAnyContainer(ref) {
		symbols.addDiagnosticCode(semantic.SeverityWarning, diagnosticOpaqueTypeRef, fileID, fmt.Sprintf("opaque container TypeRef %q at %s; consider introducing a named model for the shape", raw, position))
	}
	resolveScopedTypeRef(project, fileID, ref)
	if !typeRefModelsExist(project, ref) {
		symbols.addDiagnosticCode(semantic.SeverityError, unresolvedCode, fileID, unresolvedMessage)
		return false
	}
	return true
}

func parseInitializedModelRef(raw string, module string) (*semantic.TypeRef, error) {
	ref, err := parseTypeRef(raw, module)
	if err != nil {
		return nil, err
	}
	if ref.Kind != semantic.TypeRefNamedModel {
		return nil, fmt.Errorf("initialized store model must be a model id")
	}
	return ref, nil
}

func modelExists(project *semantic.Project, qid semantic.QualifiedID) bool {
	if qid == "" {
		return false
	}
	return modelByQID(project, qid) != nil
}

func fieldExists(project *semantic.Project, module string, raw string) bool {
	modelQID, fieldName := resolveFK(module, raw)
	if modelQID == "" || fieldName == "" {
		return false
	}
	model := modelByQID(project, modelQID)
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

func validateFlowWiringTypes(project *semantic.Project, symbols *symbolTable) {
	for fileID, entries := range project.FlowByFile {
		validateDuplicateFlowSources(project, symbols, fileID, entries)
		visibleCollected := map[string]*semantic.FlowCollectedSource{}
		for _, entry := range entries {
			switch entry.Kind {
			case semantic.FlowKindStep:
				validateFlowParamWirings(project, symbols, fileID, entry.Step.Params, paramsForTask(project, entry.Step.Task), "step.params "+entry.Step.TaskID, nil, visibleCollected, "")
			case semantic.FlowKindBranch:
				validateFlowParamWirings(project, symbols, fileID, entry.Branch.Params, paramsForBranch(project, entry.Branch.Branch), "branch.params "+entry.Branch.BranchID, nil, visibleCollected, "")
				for _, branchCase := range entry.Branch.Cases {
					validateFlowParamWirings(project, symbols, fileID, branchCase.Step.Params, paramsForTask(project, branchCase.Step.Task), "branch.cases[].params "+entry.Branch.BranchID+"."+branchCase.Label, nil, visibleCollected, "")
				}
			case semantic.FlowKindFork:
				for branchIndex, branch := range entry.Fork.Branches {
					for _, step := range branch.Steps {
						validateFlowParamWirings(project, symbols, fileID, step.Params, paramsForTask(project, step.Task), fmt.Sprintf("fork.branches[%d].steps[].params %s", branchIndex, step.TaskID), nil, visibleCollected, "")
					}
				}
				validateFlowParamWirings(project, symbols, fileID, entry.Fork.JoinParams, paramsForJoin(project, entry.Fork.Join), "join.params "+entry.Fork.JoinID, nil, visibleCollected, "")
			case semantic.FlowKindForeach:
				if entry.Foreach.Returns != "" {
					validateForeachReturns(project, symbols, fileID, entry.Foreach)
				}
				itemType, itemOK := resolveForeachItemType(project, symbols, fileID, entry.Foreach, visibleCollected)
				validateFlowParamWirings(project, symbols, fileID, entry.Foreach.Params, paramsForTask(project, entry.Foreach.Task), "foreach.params "+entry.Foreach.TaskID, itemTypeResolver(itemType, itemOK), visibleCollected, entry.Foreach.Returns)
				if entry.Foreach.Returns != "" {
					if collected := project.FlowCollectedSourcesByFile[fileID][entry.Foreach.Returns]; collected != nil {
						visibleCollected[entry.Foreach.Returns] = collected
					}
				}
			}
		}
	}
}

func validateDuplicateFlowSources(project *semantic.Project, symbols *symbolTable, fileID semantic.FileID, entries []semantic.FlowEntry) {
	nodeIDs := map[string]struct{}{}
	for _, node := range project.NodesByFile[fileID] {
		if node.GetID() != "" {
			nodeIDs[node.GetID()] = struct{}{}
		}
	}
	seenReturns := map[string]struct{}{}
	for _, entry := range entries {
		if entry.Kind != semantic.FlowKindForeach || entry.Foreach.Returns == "" {
			continue
		}
		name := entry.Foreach.Returns
		if _, exists := nodeIDs[name]; exists {
			symbols.addDiagnosticCode(semantic.SeverityError, diagnosticDuplicateFlowSource, fileID, "duplicate flow source: foreach.returns conflicts with node id: "+name)
		}
		if _, exists := seenReturns[name]; exists {
			symbols.addDiagnosticCode(semantic.SeverityError, diagnosticDuplicateFlowSource, fileID, "duplicate flow source: foreach.returns conflicts with another foreach.returns: "+name)
		}
		seenReturns[name] = struct{}{}
	}
	if main := mainTaskForFile(project, fileID); main != nil {
		seenInits := map[string]struct{}{}
		for _, init := range main.Initializes {
			if init.Name == "" {
				continue
			}
			if _, exists := nodeIDs[init.Name]; exists {
				symbols.addDiagnosticCode(semantic.SeverityError, diagnosticDuplicateFlowSource, fileID, "duplicate flow source: initializes[].name conflicts with node id: "+init.Name)
			}
			if _, exists := seenReturns[init.Name]; exists {
				symbols.addDiagnosticCode(semantic.SeverityError, diagnosticDuplicateFlowSource, fileID, "duplicate flow source: initializes[].name conflicts with foreach.returns: "+init.Name)
			}
			if _, exists := seenInits[init.Name]; exists {
				symbols.addDiagnosticCode(semantic.SeverityError, diagnosticDuplicateFlowSource, fileID, "duplicate flow source: initializes[].name conflicts with another initializes[].name: "+init.Name)
			}
			seenInits[init.Name] = struct{}{}
		}
	}
}

func validateForeachReturns(project *semantic.Project, symbols *symbolTable, fileID semantic.FileID, foreach semantic.ForeachFlow) {
	if task := project.TasksByQID[foreach.Task]; task == nil || task.Returns == nil {
		symbols.addDiagnosticCode(semantic.SeverityError, diagnosticInvalidForeachReturns, fileID, "invalid foreach.returns for "+foreach.TaskID+": apply task has no returns")
	}
	for _, wiring := range foreach.Params {
		if wiring.Source.Raw == foreach.Returns {
			symbols.addDiagnosticCode(semantic.SeverityError, diagnosticInvalidForeachReturns, fileID, "invalid foreach.returns for "+foreach.TaskID+": foreach.params references its own returns: "+foreach.Returns)
		}
	}
}

func validateFlowParamWirings(project *semantic.Project, symbols *symbolTable, fileID semantic.FileID, wirings []semantic.ParamWiring, targetParams []semantic.Param, position string, itemResolver func() (*semantic.TypeRef, bool), visibleCollected map[string]*semantic.FlowCollectedSource, invalidSelfReturn string) {
	if len(wirings) == 0 || len(targetParams) == 0 {
		return
	}
	targetByName := map[string]semantic.Param{}
	for _, param := range targetParams {
		targetByName[param.Name] = param
	}
	for _, wiring := range wirings {
		if invalidSelfReturn != "" && wiring.Source.Raw == invalidSelfReturn {
			continue
		}
		target, ok := targetByName[wiring.TargetParam]
		if !ok || !typeRefResolved(project, target.TypeRef) {
			continue
		}
		source, ok := resolveWiringSourceTypeRef(project, symbols, fileID, wiring.Source, position+"."+wiring.TargetParam, itemResolver, visibleCollected)
		if !ok || !typeRefResolved(project, source) {
			continue
		}
		if !typeRefsCompatible(project, source, target.TypeRef) {
			symbols.addDiagnosticCode(semantic.SeverityError, diagnosticIncompatibleWiringType, fileID, fmt.Sprintf("incompatible wiring type at %s.%s: source %s is not compatible with target %s", position, wiring.TargetParam, source.String(), target.TypeRef.String()))
		}
	}
}

func resolveWiringSourceTypeRef(project *semantic.Project, symbols *symbolTable, fileID semantic.FileID, source semantic.FlowSource, position string, itemResolver func() (*semantic.TypeRef, bool), visibleCollected map[string]*semantic.FlowCollectedSource) (*semantic.TypeRef, bool) {
	switch source.Kind {
	case semantic.FlowSourceParam:
		ref, ok := mainParamTypeRef(project, fileID, source.ParamName)
		if !ok {
			symbols.addDiagnosticCode(semantic.SeverityError, diagnosticUnresolvedWiringSource, fileID, "unresolved wiring source at "+position+": "+source.Raw)
			return nil, false
		}
		return ref, true
	case semantic.FlowSourceItem:
		if itemResolver == nil {
			symbols.addDiagnosticCode(semantic.SeverityError, diagnosticInvalidWiringSource, fileID, "invalid wiring source at "+position+": $item is only valid in foreach.params")
			return nil, false
		}
		return itemResolver()
	case semantic.FlowSourceInitialized:
		return source.TypeRef, true
	case semantic.FlowSourceNode:
		if source.Node == "" {
			if collected := visibleCollected[source.Raw]; collected != nil {
				return collected.TypeRef, true
			}
			if init, ok := initializedSource(project, fileID, source.Raw); ok {
				return init.TypeRef, true
			}
			symbols.addDiagnosticCode(semantic.SeverityError, diagnosticUnresolvedWiringSource, fileID, "unresolved wiring source at "+position+": "+source.Raw)
			return nil, false
		}
		if task := project.TasksByQID[source.Node]; task != nil {
			if task.Returns == nil {
				symbols.addDiagnosticCode(semantic.SeverityError, diagnosticInvalidWiringSource, fileID, "invalid wiring source at "+position+": node has no returns: "+source.Raw)
				return nil, false
			}
			return task.Returns.TypeRef, true
		}
		if join := project.JoinsByQID[source.Node]; join != nil {
			if join.Returns == nil {
				symbols.addDiagnosticCode(semantic.SeverityError, diagnosticInvalidWiringSource, fileID, "invalid wiring source at "+position+": node has no returns: "+source.Raw)
				return nil, false
			}
			return join.Returns.TypeRef, true
		}
		if project.NodesByID[source.Node] != nil {
			symbols.addDiagnosticCode(semantic.SeverityError, diagnosticInvalidWiringSource, fileID, "invalid wiring source at "+position+": node is not a task or join: "+source.Raw)
			return nil, false
		}
		symbols.addDiagnosticCode(semantic.SeverityError, diagnosticUnresolvedWiringSource, fileID, "unresolved wiring source at "+position+": "+source.Raw)
		return nil, false
	default:
		symbols.addDiagnosticCode(semantic.SeverityError, diagnosticUnresolvedWiringSource, fileID, "unresolved wiring source at "+position+": "+source.Raw)
		return nil, false
	}
}

func resolveForeachItemType(project *semantic.Project, symbols *symbolTable, fileID semantic.FileID, foreach semantic.ForeachFlow, visibleCollected map[string]*semantic.FlowCollectedSource) (*semantic.TypeRef, bool) {
	overType, ok := resolveWiringSourceTypeRef(project, symbols, fileID, foreach.Over, "foreach.over "+foreach.TaskID, nil, visibleCollected)
	if !ok || !typeRefResolved(project, overType) {
		return nil, false
	}
	if typeRefIsAny(overType) {
		return overType, true
	}
	normalized := normalizeContainerTypeRef(project, overType)
	if normalized == nil || normalized.Kind != semantic.TypeRefList || normalized.Elem == nil || !typeRefResolved(project, normalized.Elem) {
		symbols.addDiagnosticCode(semantic.SeverityError, diagnosticInvalidForeachOverType, fileID, fmt.Sprintf("invalid foreach.over type for %s: %s is not list<T>", foreach.TaskID, overType.String()))
		return nil, false
	}
	return normalized.Elem, true
}

func itemTypeResolver(itemType *semantic.TypeRef, ok bool) func() (*semantic.TypeRef, bool) {
	return func() (*semantic.TypeRef, bool) {
		if !ok {
			return nil, false
		}
		return itemType, true
	}
}

func mainParamTypeRef(project *semantic.Project, fileID semantic.FileID, name string) (*semantic.TypeRef, bool) {
	mainTask := mainTaskForFile(project, fileID)
	if mainTask == nil {
		return nil, false
	}
	for _, param := range mainTask.Params {
		if param.Name == name {
			return param.TypeRef, true
		}
	}
	return nil, false
}

func mainTaskForFile(project *semantic.Project, fileID semantic.FileID) *semantic.Task {
	if project == nil {
		return nil
	}
	return project.TasksByQID[project.MainNodeByFile[fileID]]
}

func initializedSource(project *semantic.Project, fileID semantic.FileID, name string) (semantic.InitializedStore, bool) {
	mainTask := mainTaskForFile(project, fileID)
	if mainTask == nil {
		return semantic.InitializedStore{}, false
	}
	for _, init := range mainTask.Initializes {
		if init.Name == name {
			return init, true
		}
	}
	return semantic.InitializedStore{}, false
}

func validateTaskReturnSource(project *semantic.Project, symbols *symbolTable, task *semantic.Task) {
	if task == nil || task.Returns == nil || task.Returns.Source == "" {
		return
	}
	position := "returns.source " + task.QID.String() + "." + task.Returns.Name
	source, sourceType, ok := resolveReturnSource(project, symbols, task.FileID, task.Returns.Source, position)
	task.Returns.SourceRef = source
	if !ok || !typeRefResolved(project, sourceType) || !typeRefResolved(project, task.Returns.TypeRef) {
		return
	}
	if !typeRefsCompatible(project, sourceType, task.Returns.TypeRef) {
		symbols.addDiagnosticCode(semantic.SeverityError, diagnosticIncompatibleReturnType, task.FileID, fmt.Sprintf("incompatible return type at %s: source %s is not compatible with target %s", position, sourceType.String(), task.Returns.TypeRef.String()))
	}
}

func resolveReturnSource(project *semantic.Project, symbols *symbolTable, fileID semantic.FileID, raw string, position string) (semantic.FlowSource, *semantic.TypeRef, bool) {
	if raw == "$item" {
		symbols.addDiagnosticCode(semantic.SeverityError, diagnosticInvalidReturnSource, fileID, "invalid return source at "+position+": $item is not valid in returns.source")
		return semantic.FlowSource{Kind: semantic.FlowSourceItem, Raw: raw}, nil, false
	}
	if strings.HasPrefix(raw, "$params.") {
		paramName := strings.TrimPrefix(raw, "$params.")
		ref, ok := mainParamTypeRef(project, fileID, paramName)
		source := semantic.FlowSource{Kind: semantic.FlowSourceParam, Raw: raw, ParamName: paramName, TypeRef: ref}
		if !ok {
			symbols.addDiagnosticCode(semantic.SeverityError, diagnosticUnresolvedReturnSource, fileID, "unresolved return source at "+position+": "+raw)
			return source, nil, false
		}
		return source, ref, true
	}
	qid := resolveAnyNodeQID(project, fileID, raw)
	if qid != "" {
		source := semantic.FlowSource{Kind: semantic.FlowSourceNode, Raw: raw, Node: qid}
		if task := project.TasksByQID[qid]; task != nil {
			if task.Returns == nil {
				symbols.addDiagnosticCode(semantic.SeverityError, diagnosticInvalidReturnSource, fileID, "invalid return source at "+position+": node has no returns: "+raw)
				return source, nil, false
			}
			source.AssetName = task.Returns.Name
			source.TypeRef = task.Returns.TypeRef
			return source, task.Returns.TypeRef, true
		}
		if join := project.JoinsByQID[qid]; join != nil {
			if join.Returns == nil {
				symbols.addDiagnosticCode(semantic.SeverityError, diagnosticInvalidReturnSource, fileID, "invalid return source at "+position+": node has no returns: "+raw)
				return source, nil, false
			}
			source.AssetName = join.Returns.Name
			source.TypeRef = join.Returns.TypeRef
			return source, join.Returns.TypeRef, true
		}
		if project.NodesByID[qid] != nil {
			symbols.addDiagnosticCode(semantic.SeverityError, diagnosticInvalidReturnSource, fileID, "invalid return source at "+position+": node is not a task or join: "+raw)
			return source, nil, false
		}
	}
	if collected := project.FlowCollectedSourcesByFile[fileID][raw]; collected != nil {
		source := semantic.FlowSource{Kind: semantic.FlowSourceNode, Raw: raw, AssetName: raw, TypeRef: collected.TypeRef}
		return source, collected.TypeRef, true
	}
	if init, ok := initializedSource(project, fileID, raw); ok {
		source := semantic.FlowSource{Kind: semantic.FlowSourceInitialized, Raw: raw, TypeRef: init.TypeRef}
		return source, init.TypeRef, true
	}
	symbols.addDiagnosticCode(semantic.SeverityError, diagnosticUnresolvedReturnSource, fileID, "unresolved return source at "+position+": "+raw)
	return semantic.FlowSource{Kind: semantic.FlowSourceNode, Raw: raw}, nil, false
}

func typeRefResolved(project *semantic.Project, ref *semantic.TypeRef) bool {
	if ref == nil {
		return false
	}
	switch ref.Kind {
	case semantic.TypeRefPrimitive:
		return true
	case semantic.TypeRefNamedModel:
		model := modelByQID(project, ref.Model)
		if model == nil {
			return false
		}
		switch model.Kind {
		case "list":
			return typeRefResolved(project, model.ElementRef)
		case "dict":
			return typeRefResolved(project, model.ValueRef)
		default:
			return true
		}
	case semantic.TypeRefList:
		return typeRefResolved(project, ref.Elem)
	case semantic.TypeRefDict:
		return typeRefResolved(project, ref.Value)
	default:
		return false
	}
}

func paramsForTask(project *semantic.Project, qid semantic.QualifiedID) []semantic.Param {
	if task := project.TasksByQID[qid]; task != nil {
		return task.Params
	}
	return nil
}

func paramsForBranch(project *semantic.Project, qid semantic.QualifiedID) []semantic.Param {
	if branch := project.BranchesByQID[qid]; branch != nil {
		return branch.Params
	}
	return nil
}

func paramsForJoin(project *semantic.Project, qid semantic.QualifiedID) []semantic.Param {
	if join := project.JoinsByQID[qid]; join != nil {
		return join.Params
	}
	return nil
}
