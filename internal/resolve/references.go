package resolve

import (
	"strings"

	"github.com/hiroshiasayadev-prog/brewprint/internal/semantic"
)

func buildReferences(project *semantic.Project) {
	for _, task := range project.TasksByQID {
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

	for _, branch := range project.BranchesByQID {
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

	for _, join := range project.JoinsByQID {
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
	return semantic.ReferenceEndpoint{
		Object:      "node",
		Kind:        string(node.GetKind()),
		ID:          node.GetQID().String(),
		QualifiedID: node.GetQID(),
		Name:        node.GetID(),
		File:        node.GetFileID(),
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
		Name:      asset.Name,
		Producer:  asset.ProducedBy,
		Model:     asset.Model,
		ScopeFile: asset.FileID,
	}
}

func modelFieldEndpoint(model *semantic.Model, fieldName string) semantic.ReferenceEndpoint {
	return semantic.ReferenceEndpoint{
		Object:      "model_field",
		Kind:        "field",
		ID:          model.QID.String() + "." + fieldName,
		QualifiedID: model.QID,
		Name:        fieldName,
		File:        model.FileID,
	}
}

func modelFieldTargetEndpoint(model semantic.QualifiedID, fieldName string) semantic.ReferenceEndpoint {
	return semantic.ReferenceEndpoint{
		Object:      "model_field",
		Kind:        "field",
		ID:          model.String() + "." + fieldName,
		QualifiedID: model,
		Name:        fieldName,
	}
}

func modelOrPrimitiveKey(project *semantic.Project, qid semantic.QualifiedID, raw string) semantic.ObjectKey {
	if qid != "" {
		if _, ok := project.ModelsByQID[qid]; ok {
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
