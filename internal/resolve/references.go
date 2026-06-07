package resolve

import (
	"strings"

	"github.com/hiroshiasayadev-prog/brewprint/internal/semantic"
)

func buildReferences(project *semantic.Project) {
	seenTasks := map[*semantic.Task]struct{}{}
	for _, task := range project.TasksByQID {
		if _, ok := seenTasks[task]; ok {
			continue
		}
		seenTasks[task] = struct{}{}
		for _, param := range task.Params {
			addReference(project, semantic.Reference{
				Kind:      semantic.ReferenceKindParamModel,
				SourceKey: semantic.NodeObjectKey(task.QID),
				TargetKey: modelOrPrimitiveKey(project, param.Model, param.ModelName),
				From:      nodeEndpoint(task),
				To:        modelOrPrimitiveEndpoint(project, param.Model, param.ModelName),
			})
		}
		if task.Returns != nil {
			addReference(project, semantic.Reference{
				Kind:      semantic.ReferenceKindReturnModel,
				SourceKey: semantic.NodeObjectKey(task.QID),
				TargetKey: modelOrPrimitiveKey(project, task.Returns.Model, task.Returns.ModelName),
				From:      nodeEndpoint(task),
				To:        modelOrPrimitiveEndpoint(project, task.Returns.Model, task.Returns.ModelName),
			})
			addReference(project, semantic.Reference{
				Kind:      semantic.ReferenceKindProducesAsset,
				SourceKey: semantic.NodeObjectKey(task.QID),
				TargetKey: semantic.AssetObjectKey(task.Returns.Asset),
				From:      nodeEndpoint(task),
				To:        assetEndpoint(task.Returns.Asset),
			})
		}
		for _, read := range task.Reads {
			if read.Store == "" {
				continue
			}
			addReference(project, semantic.Reference{
				Kind:      semantic.ReferenceKindReads,
				SourceKey: semantic.NodeObjectKey(task.QID),
				TargetKey: semantic.NodeObjectKey(read.Store),
				From:      nodeEndpoint(task),
				To:        storeEndpoint(project, read.Store),
			})
		}
		for _, write := range task.Writes {
			if write.Store == "" {
				continue
			}
			addReference(project, semantic.Reference{
				Kind:      semantic.ReferenceKindWrites,
				SourceKey: semantic.NodeObjectKey(task.QID),
				TargetKey: semantic.NodeObjectKey(write.Store),
				From:      nodeEndpoint(task),
				To:        storeEndpoint(project, write.Store),
			})
		}
	}

	seenBranches := map[*semantic.Branch]struct{}{}
	for _, branch := range project.BranchesByQID {
		if _, ok := seenBranches[branch]; ok {
			continue
		}
		seenBranches[branch] = struct{}{}
		for _, param := range branch.Params {
			addReference(project, semantic.Reference{
				Kind:      semantic.ReferenceKindParamModel,
				SourceKey: semantic.NodeObjectKey(branch.QID),
				TargetKey: modelOrPrimitiveKey(project, param.Model, param.ModelName),
				From:      nodeEndpoint(branch),
				To:        modelOrPrimitiveEndpoint(project, param.Model, param.ModelName),
			})
		}
	}

	seenJoins := map[*semantic.Join]struct{}{}
	for _, join := range project.JoinsByQID {
		if _, ok := seenJoins[join]; ok {
			continue
		}
		seenJoins[join] = struct{}{}
		for _, param := range join.Params {
			addReference(project, semantic.Reference{
				Kind:      semantic.ReferenceKindParamModel,
				SourceKey: semantic.NodeObjectKey(join.QID),
				TargetKey: modelOrPrimitiveKey(project, param.Model, param.ModelName),
				From:      nodeEndpoint(join),
				To:        modelOrPrimitiveEndpoint(project, param.Model, param.ModelName),
			})
		}
		if join.Returns != nil {
			addReference(project, semantic.Reference{
				Kind:      semantic.ReferenceKindReturnModel,
				SourceKey: semantic.NodeObjectKey(join.QID),
				TargetKey: modelOrPrimitiveKey(project, join.Returns.Model, join.Returns.ModelName),
				From:      nodeEndpoint(join),
				To:        modelOrPrimitiveEndpoint(project, join.Returns.Model, join.Returns.ModelName),
			})
			addReference(project, semantic.Reference{
				Kind:      semantic.ReferenceKindProducesAsset,
				SourceKey: semantic.NodeObjectKey(join.QID),
				TargetKey: semantic.AssetObjectKey(join.Returns.Asset),
				From:      nodeEndpoint(join),
				To:        assetEndpoint(join.Returns.Asset),
			})
		}
	}

	for _, store := range project.StoresByQID {
		if store.Of == "" {
			continue
		}
		addReference(project, semantic.Reference{
			Kind:      semantic.ReferenceKindStoreOf,
			SourceKey: semantic.NodeObjectKey(store.QID),
			TargetKey: modelOrPrimitiveKey(project, store.Of, store.OfName),
			From:      nodeEndpoint(store),
			To:        modelOrPrimitiveEndpoint(project, store.Of, store.OfName),
		})
	}
	for _, storesByName := range project.StoresByFileLocal {
		for _, store := range storesByName {
			if store.Of == "" {
				continue
			}
			addReference(project, semantic.Reference{
				Kind:      semantic.ReferenceKindStoreOf,
				SourceKey: semantic.NodeObjectKey(store.QID),
				TargetKey: modelOrPrimitiveKey(project, store.Of, store.OfName),
				From:      nodeEndpoint(store),
				To:        modelOrPrimitiveEndpoint(project, store.Of, store.OfName),
			})
		}
	}

	for _, event := range project.EventsByQID {
		if event.PayloadModel != "" || event.PayloadName != "" {
			addReference(project, semantic.Reference{
				Kind:      semantic.ReferenceKindEventPayload,
				SourceKey: semantic.NodeObjectKey(event.QID),
				TargetKey: modelOrPrimitiveKey(project, event.PayloadModel, event.PayloadName),
				From:      nodeEndpoint(event),
				To:        modelOrPrimitiveEndpoint(project, event.PayloadModel, event.PayloadName),
			})
		}
		if event.Actor != "" {
			addReference(project, semantic.Reference{
				Kind:      semantic.ReferenceKindEventActor,
				SourceKey: semantic.NodeObjectKey(event.QID),
				TargetKey: semantic.NodeObjectKey(semantic.QualifiedID(event.Actor)),
				From:      nodeEndpoint(event),
				To:        actorEndpoint(project, event.Actor),
			})
		}
		if event.Watches != "" {
			addReference(project, semantic.Reference{
				Kind:      semantic.ReferenceKindEventWatches,
				SourceKey: semantic.NodeObjectKey(event.QID),
				TargetKey: semantic.NodeObjectKey(event.Watches),
				From:      nodeEndpoint(event),
				To:        storeEndpoint(project, event.Watches),
			})
		}
	}

	for _, transitions := range project.TransitionsByFile {
		for _, transition := range transitions {
			transitionKey := semantic.TransitionObjectKey(transition)
			transitionEndpoint := transitionEndpoint(transition)
			if transition.FromState != "" {
				addReference(project, semantic.Reference{
					Kind:      semantic.ReferenceKindTransitionFrom,
					SourceKey: transitionKey,
					TargetKey: semantic.NodeObjectKey(transition.FromState),
					From:      transitionEndpoint,
					To:        nodeEndpoint(project.StatesByQID[transition.FromState]),
				})
			}
			if transition.Event != "" {
				addReference(project, semantic.Reference{
					Kind:      semantic.ReferenceKindTransitionEvent,
					SourceKey: transitionKey,
					TargetKey: semantic.NodeObjectKey(transition.Event),
					From:      transitionEndpoint,
					To:        nodeEndpoint(project.EventsByQID[transition.Event]),
				})
			}
			if transition.ToState != "" {
				addReference(project, semantic.Reference{
					Kind:      semantic.ReferenceKindTransitionTo,
					SourceKey: transitionKey,
					TargetKey: semantic.NodeObjectKey(transition.ToState),
					From:      transitionEndpoint,
					To:        nodeEndpoint(project.StatesByQID[transition.ToState]),
				})
			}
			if transition.ActionTask != "" {
				if task := project.TasksByQID[transition.ActionTask]; task != nil {
					addReference(project, semantic.Reference{
						Kind:      semantic.ReferenceKindTransitionAction,
						SourceKey: transitionKey,
						TargetKey: semantic.NodeObjectKey(transition.ActionTask),
						From:      transitionEndpoint,
						To:        nodeEndpoint(task),
					})
				}
			}
		}
	}

	for _, scenario := range project.ScenariosByID {
		scenarioKey := semantic.ScenarioObjectKey(scenario.ID)
		addReference(project, semantic.Reference{
			Kind:      semantic.ReferenceKindScenarioStateFile,
			SourceKey: scenarioKey,
			TargetKey: semantic.StateFileObjectKey(scenario.StateFile),
			From:      scenarioEndpoint(scenario),
			To:        stateFileEndpoint(scenario.StateFile),
		})
		for i, step := range scenario.Steps {
			transition := step.Transition
			addReference(project, semantic.Reference{
				Kind:      semantic.ReferenceKindScenarioStepTransition,
				SourceKey: scenarioKey,
				TargetKey: semantic.TransitionObjectKey(transition),
				From:      scenarioStepEndpoint(scenario, i+1),
				To:        transitionEndpoint(transition),
			})
		}
	}

	buildAssetConsumerReferences(project)

	for _, model := range project.ModelsByQID {
		module := moduleForFileID(model.FileID)
		for _, field := range model.Fields {
			fieldKey := semantic.ModelFieldObjectKey(model.QID, field.Name)
			fieldFrom := modelFieldEndpoint(model, field.Name)
			if field.Type != "" {
				targetQID := resolveModelQID(module, field.Type)
				addReference(project, semantic.Reference{
					Kind:      semantic.ReferenceKindFieldType,
					SourceKey: fieldKey,
					TargetKey: modelOrPrimitiveKey(project, targetQID, field.Type),
					From:      fieldFrom,
					To:        modelOrPrimitiveEndpoint(project, targetQID, field.Type),
				})
			}
			if field.FK != "" {
				fkModel, fkField := resolveFK(module, field.FK)
				addReference(project, semantic.Reference{
					Kind:      semantic.ReferenceKindFieldFK,
					SourceKey: fieldKey,
					TargetKey: semantic.ModelFieldObjectKey(fkModel, fkField),
					From:      fieldFrom,
					To:        modelFieldTargetEndpoint(fkModel, fkField),
				})
			}
		}
	}
}

func addReference(project *semantic.Project, ref semantic.Reference) {
	if ref.SourceKey == "" || ref.TargetKey == "" {
		return
	}
	project.ReferencesBySource[ref.SourceKey] = append(project.ReferencesBySource[ref.SourceKey], ref)
	project.ReferencesByTarget[ref.TargetKey] = append(project.ReferencesByTarget[ref.TargetKey], ref)
}

func nodeEndpoint(node semantic.Node) semantic.ReferenceEndpoint {
	if node == nil {
		return semantic.ReferenceEndpoint{}
	}
	if isPrivateSubNode(node) {
		return semantic.ReferenceEndpoint{
			Object:  "node",
			Kind:    string(node.GetKind()),
			ID:      semantic.PrivateNodeID(node.GetFileID(), node.GetID()),
			Name:    node.GetID(),
			File:    node.GetFileID(),
			LocalID: node.GetID(),
		}
	}
	return semantic.ReferenceEndpoint{
		Object:      "node",
		Kind:        string(node.GetKind()),
		ID:          node.GetQID().String(),
		QualifiedID: node.GetQID(),
		Name:        node.GetID(),
		File:        node.GetFileID(),
	}
}

func isPrivateSubNode(node semantic.Node) bool {
	if model, ok := node.(*semantic.Model); ok {
		return model.FilePrivate
	}
	if node == nil || node.IsMain() {
		return false
	}
	switch node.GetKind() {
	case semantic.NodeKindTask, semantic.NodeKindBranch, semantic.NodeKindFork, semantic.NodeKindJoin:
		return true
	default:
		return false
	}
}

func storeEndpoint(project *semantic.Project, qid semantic.QualifiedID) semantic.ReferenceEndpoint {
	if store := project.StoresByQID[qid]; store != nil {
		return nodeEndpoint(store)
	}
	for _, storesByName := range project.StoresByFileLocal {
		for _, store := range storesByName {
			if store.QID == qid {
				return nodeEndpoint(store)
			}
		}
	}
	return semantic.ReferenceEndpoint{Object: "node", Kind: "store", ID: qid.String(), QualifiedID: qid}
}

func actorEndpoint(project *semantic.Project, actorID string) semantic.ReferenceEndpoint {
	if actor := project.ActorsByQID[semantic.QualifiedID(actorID)]; actor != nil {
		return nodeEndpoint(actor)
	}
	return semantic.ReferenceEndpoint{Object: "node", Kind: "actor", ID: actorID, QualifiedID: semantic.QualifiedID(actorID), Name: actorID}
}

func scenarioEndpoint(scenario *semantic.SequenceScenario) semantic.ReferenceEndpoint {
	if scenario == nil {
		return semantic.ReferenceEndpoint{}
	}
	return semantic.ReferenceEndpoint{
		Object: "view",
		Kind:   "sequence_diagram",
		ID:     scenario.ID,
		Name:   scenario.Title,
		File:   scenario.FileID,
	}
}

func scenarioStepEndpoint(scenario *semantic.SequenceScenario, index int) semantic.ReferenceEndpoint {
	if scenario == nil {
		return semantic.ReferenceEndpoint{}
	}
	return semantic.ReferenceEndpoint{
		Object:  "scenario_step",
		Kind:    "sequence_step",
		ID:      string(semantic.ScenarioStepObjectKey(scenario.ID, index)),
		Name:    scenario.ID,
		File:    scenario.FileID,
		LocalID: scenario.ID,
	}
}

func stateFileEndpoint(fileID semantic.FileID) semantic.ReferenceEndpoint {
	return semantic.ReferenceEndpoint{
		Object: "file",
		Kind:   "state_file",
		ID:     fileID.String(),
		File:   fileID,
	}
}

func transitionEndpoint(transition semantic.Transition) semantic.ReferenceEndpoint {
	return semantic.ReferenceEndpoint{
		Object:    "transition",
		Kind:      "transition",
		ID:        semantic.TransitionID(transition),
		File:      transition.FileID,
		LocalID:   transition.From + ":" + transition.On,
		StateFile: transition.FileID,
		From:      transition.From,
		On:        transition.On,
		To:        transition.To,
		Guard:     transition.Guard,
		Action:    transition.ActionTask,
	}
}

func assetEndpoint(asset *semantic.Asset) semantic.ReferenceEndpoint {
	if asset == nil {
		return semantic.ReferenceEndpoint{}
	}
	return semantic.ReferenceEndpoint{
		Object:    "asset",
		Kind:      "asset",
		ID:        semantic.AssetID(asset.ProducedBy, asset.Name),
		Name:      asset.Name,
		Producer:  asset.ProducedBy,
		Model:     asset.Model,
		ScopeFile: asset.FileID,
	}
}

func buildAssetConsumerReferences(project *semantic.Project) {
	for _, entries := range project.FlowByFile {
		for _, entry := range entries {
			switch entry.Kind {
			case semantic.FlowKindStep:
				addAssetConsumerReferences(project, entry.Step.Task, entry.Step.Params)
			case semantic.FlowKindForeach:
				addAssetConsumerReference(project, entry.Foreach.Task, entry.Foreach.Over)
				addAssetConsumerReferences(project, entry.Foreach.Task, entry.Foreach.Params)
			case semantic.FlowKindFork:
				for _, branch := range entry.Fork.Branches {
					for _, step := range branch.Steps {
						addAssetConsumerReferences(project, step.Task, step.Params)
					}
				}
				addAssetConsumerReferences(project, entry.Fork.Join, entry.Fork.JoinParams)
			case semantic.FlowKindBranch:
				addAssetConsumerReferences(project, entry.Branch.Branch, entry.Branch.Params)
				for _, branchCase := range entry.Branch.Cases {
					addAssetConsumerReferences(project, branchCase.Step.Task, branchCase.Step.Params)
				}
			}
		}
	}
}

func addAssetConsumerReferences(project *semantic.Project, consumerID semantic.QualifiedID, wirings []semantic.ParamWiring) {
	for _, wiring := range wirings {
		addAssetConsumerReference(project, consumerID, wiring.Source)
	}
}

func addAssetConsumerReference(project *semantic.Project, consumerID semantic.QualifiedID, source semantic.FlowSource) {
	if source.Kind != semantic.FlowSourceNode || source.Node == "" || source.AssetName == "" || consumerID == "" {
		return
	}
	asset := assetByProducerAndName(project, source.Node, source.AssetName)
	consumer := project.NodesByID[consumerID]
	if asset == nil || consumer == nil {
		return
	}
	addReference(project, semantic.Reference{
		Kind:      semantic.ReferenceKindConsumesAsset,
		SourceKey: semantic.AssetObjectKey(asset),
		TargetKey: semantic.NodeObjectKey(consumer.GetQID()),
		From:      assetEndpoint(asset),
		To:        nodeEndpoint(consumer),
	})
}

func assetByProducerAndName(project *semantic.Project, producer semantic.QualifiedID, name string) *semantic.Asset {
	if task := project.TasksByQID[producer]; task != nil && task.Returns != nil && task.Returns.Asset != nil && task.Returns.Asset.Name == name {
		return task.Returns.Asset
	}
	if join := project.JoinsByQID[producer]; join != nil && join.Returns != nil && join.Returns.Asset != nil && join.Returns.Asset.Name == name {
		return join.Returns.Asset
	}
	for _, task := range project.TasksByQID {
		if task.Returns != nil && task.Returns.Asset != nil && task.Returns.Asset.ProducedBy == producer && task.Returns.Asset.Name == name {
			return task.Returns.Asset
		}
	}
	for _, join := range project.JoinsByQID {
		if join.Returns != nil && join.Returns.Asset != nil && join.Returns.Asset.ProducedBy == producer && join.Returns.Asset.Name == name {
			return join.Returns.Asset
		}
	}
	return nil
}

func modelFieldEndpoint(model *semantic.Model, fieldName string) semantic.ReferenceEndpoint {
	return semantic.ReferenceEndpoint{
		Object:      "field",
		Kind:        "field",
		ID:          model.QID.String() + "." + fieldName,
		QualifiedID: model.QID,
		Name:        fieldName,
		File:        model.FileID,
	}
}

func modelFieldTargetEndpoint(model semantic.QualifiedID, fieldName string) semantic.ReferenceEndpoint {
	return semantic.ReferenceEndpoint{
		Object:      "field",
		Kind:        "field",
		ID:          model.String() + "." + fieldName,
		QualifiedID: model,
		Name:        fieldName,
	}
}

func modelOrPrimitiveKey(project *semantic.Project, qid semantic.QualifiedID, raw string) semantic.ObjectKey {
	if qid != "" {
		if project.ModelsByQID[qid] != nil {
			return semantic.NodeObjectKey(qid)
		}
	}
	if raw == "" && qid != "" {
		raw = qid.String()
	}
	return semantic.PrimitiveObjectKey(raw)
}

func modelOrPrimitiveEndpoint(project *semantic.Project, qid semantic.QualifiedID, raw string) semantic.ReferenceEndpoint {
	if qid != "" {
		if model := project.ModelsByQID[qid]; model != nil {
			return nodeEndpoint(model)
		}
	}
	if raw == "" && qid != "" {
		raw = qid.String()
	}
	return semantic.ReferenceEndpoint{Object: "primitive", Kind: "primitive", ID: raw, Name: raw}
}

func resolveFK(module string, fk string) (semantic.QualifiedID, string) {
	parts := strings.Split(fk, ".")
	if len(parts) < 2 {
		return "", fk
	}
	fieldName := parts[len(parts)-1]
	modelParts := parts[:len(parts)-1]
	for _, part := range modelParts {
		if part == "model" {
			return semantic.QualifiedID(strings.Join(modelParts, ".")), fieldName
		}
	}
	if len(modelParts) == 1 && module != "" {
		return semantic.QualifiedID(module + ".model." + modelParts[0]), fieldName
	}
	return semantic.QualifiedID(strings.Join(modelParts, ".")), fieldName
}
