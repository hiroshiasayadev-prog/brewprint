package query

import (
	"sort"

	"github.com/hiroshiasayadev-prog/brewprint/internal/semantic"
)

func (s *Service) Inspect(req InspectRequest) (InspectResponse, error) {
	if req.Detail != "" && req.Detail != "brief" && req.Detail != "normal" && req.Detail != "full" {
		return InspectResponse{}, errUnsupportedDetail(req.Detail)
	}

	if s.isFileSelector(req.Selector) {
		return s.inspectFile(req)
	}
	if s.isAPIViewSelector(req.Selector) {
		return s.inspectAPIView(req)
	}
	if s.isERViewSelector(req.Selector) {
		return s.inspectERView(req)
	}
	if s.isScenarioSelector(req.Selector) {
		return s.inspectScenario(req)
	}
	if s.isTransitionSelector(req.Selector) {
		return s.inspectTransition(req)
	}
	if s.isFieldSelector(req.Selector) {
		return s.inspectField(req)
	}
	if s.isAssetSelector(req.Selector) {
		return s.inspectAsset(req)
	}
	if s.isPrivateSubNodeSelector(req.Selector) {
		return s.inspectPrivateSubNode(req)
	}

	node, err := s.nodeByID(req.Selector.ID)
	if err != nil {
		return InspectResponse{}, err
	}
	sig, doc, err := s.signatureForNode(node)
	if err != nil {
		return InspectResponse{}, err
	}
	refs, err := s.GetReferences(GetReferencesRequest{Selector: req.Selector, Direction: string(semantic.ReferenceDirectionBoth)})
	if err != nil {
		return InspectResponse{}, err
	}

	members := map[string]any{}
	switch n := node.(type) {
	case *semantic.Task:
		members = s.taskMembers(n)
	case *semantic.Model:
		members["fields"] = fieldSignatures(n.Fields)
	case *semantic.Store:
		members = s.storeMembers(n)
	case *semantic.State:
		members = s.stateMembers(n)
	case *semantic.Event:
		members = s.eventMembers(n)
	default:
		return InspectResponse{}, errUnsupportedInspect(node.GetKind())
	}

	return InspectResponse{
		Object:      objectRef(node),
		Signature:   sig,
		Doc:         doc,
		Source:      sourceMap(node.GetFileID()),
		Members:     members,
		References:  refs.References,
		Diagnostics: []semantic.Diagnostic{},
	}, nil
}

func errUnsupportedInspect(kind semantic.NodeKind) error {
	return &unsupportedInspectError{kind: kind}
}

func errUnsupportedDetail(detail string) error {
	return &unsupportedDetailError{detail: detail}
}

type unsupportedInspectError struct{ kind semantic.NodeKind }

type unsupportedDetailError struct{ detail string }

func (e *unsupportedInspectError) Error() string {
	return "unsupported inspect kind: " + string(e.kind)
}

func (e *unsupportedDetailError) Error() string {
	return "unsupported detail: " + e.detail
}

func (s *Service) taskMembers(task *semantic.Task) map[string]any {
	members := map[string]any{}
	if task.Returns != nil && task.Returns.Asset != nil {
		members["assets"] = []AssetRef{*assetRef(task.Returns.Asset)}
	}

	subTasks := make([]map[string]any, 0)
	for _, node := range s.project.NodesByFile[task.FileID] {
		sub, ok := node.(*semantic.Task)
		if !ok || sub.QID == task.QID || sub.Main {
			continue
		}
		subRef := objectRef(sub)
		subTasks = append(subTasks, map[string]any{
			"object":   subRef.Object,
			"kind":     subRef.Kind,
			"id":       subRef.ID,
			"file":     subRef.File,
			"local_id": subRef.LocalID,
			"label":    subRef.Label,
			"signature": Signature{
				"reads":  storeIDs(sub.Reads),
				"writes": storeIDs(sub.Writes),
			},
			"source": subRef.Source,
		})
	}
	sort.Slice(subTasks, func(i, j int) bool { return subTasks[i]["id"].(string) < subTasks[j]["id"].(string) })
	if len(subTasks) > 0 {
		members["sub_tasks"] = subTasks
	}

	if flow := s.project.FlowByFile[task.FileID]; len(flow) > 0 {
		members["flow"] = map[string]any{
			"file":          task.FileID.String(),
			"entries":       flowEntriesSummary(flow),
			"schema_status": "draft",
		}
	}
	if transitions := transitionRefsForTask(s.project, task.QID); len(transitions) > 0 {
		members["action_transitions"] = transitions
	}
	return members
}

func (s *Service) storeMembers(store *semantic.Store) map[string]any {
	members := map[string]any{}
	if model := s.project.ModelsByQID[store.Of]; model != nil {
		members["model"] = map[string]any{
			"object": "node",
			"kind":   "model",
			"id":     model.QID.String(),
			"fields": fieldSignatures(model.Fields),
		}
	}
	return members
}

func (s *Service) stateMembers(state *semantic.State) map[string]any {
	members := map[string]any{}
	var incoming []TransitionRef
	var outgoing []TransitionRef
	for _, transitions := range s.project.TransitionsByFile {
		for _, transition := range transitions {
			if transition.ToState == state.QID {
				incoming = append(incoming, transitionRefFromSemantic(transition))
			}
			if transition.FromState == state.QID {
				outgoing = append(outgoing, transitionRefFromSemantic(transition))
			}
		}
	}
	sortTransitionRefs(incoming)
	sortTransitionRefs(outgoing)
	if len(incoming) > 0 {
		members["incoming_transitions"] = incoming
	}
	if len(outgoing) > 0 {
		members["outgoing_transitions"] = outgoing
	}
	return members
}

func (s *Service) inspectField(req InspectRequest) (InspectResponse, error) {
	model, field, err := s.modelFieldBySelector(req.Selector)
	if err != nil {
		return InspectResponse{}, err
	}
	refs, err := s.GetReferences(GetReferencesRequest{Selector: req.Selector, Direction: string(semantic.ReferenceDirectionBoth)})
	if err != nil {
		return InspectResponse{}, err
	}
	return InspectResponse{
		Object:      fieldObjectRef(model, field),
		Signature:   signatureForField(field),
		Doc:         field.Note,
		Source:      sourceMap(model.FileID),
		Members:     s.fieldMembers(model, field),
		References:  refs.References,
		Diagnostics: []semantic.Diagnostic{},
	}, nil
}

func (s *Service) fieldMembers(model *semantic.Model, field semantic.ModelField) map[string]any {
	members := map[string]any{"model": objectRef(model)}
	if field.Type != "" {
		members["type"] = field.Type
	}
	if field.FK != "" {
		members["fk"] = field.FK
	}
	return members
}

func (s *Service) inspectTransition(req InspectRequest) (InspectResponse, error) {
	transition, err := s.transitionBySelector(req.Selector)
	if err != nil {
		return InspectResponse{}, err
	}
	refs, err := s.GetReferences(GetReferencesRequest{Selector: req.Selector, Direction: string(semantic.ReferenceDirectionBoth)})
	if err != nil {
		return InspectResponse{}, err
	}
	return InspectResponse{
		Object:      transitionObjectRef(transition),
		Signature:   signatureForTransition(transition),
		Doc:         transition.Note,
		Source:      sourceMap(transition.FileID),
		Members:     s.transitionMembers(transition),
		References:  refs.References,
		Diagnostics: []semantic.Diagnostic{},
	}, nil
}

func (s *Service) transitionMembers(transition semantic.Transition) map[string]any {
	members := map[string]any{}
	if state := s.project.StatesByQID[transition.FromState]; state != nil {
		members["from_state"] = objectRef(state)
	}
	if event := s.project.EventsByQID[transition.Event]; event != nil {
		members["event"] = objectRef(event)
	}
	if state := s.project.StatesByQID[transition.ToState]; state != nil {
		members["to_state"] = objectRef(state)
	}
	if task := s.project.TasksByQID[transition.ActionTask]; task != nil {
		members["action_task"] = objectRef(task)
	}
	return members
}

func (s *Service) inspectScenario(req InspectRequest) (InspectResponse, error) {
	scenario, err := s.scenarioBySelector(req.Selector)
	if err != nil {
		return InspectResponse{}, err
	}
	refs, err := s.GetReferences(GetReferencesRequest{Selector: req.Selector, Direction: string(semantic.ReferenceDirectionBoth)})
	if err != nil {
		return InspectResponse{}, err
	}
	return InspectResponse{
		Object:      scenarioObjectRef(scenario),
		Signature:   signatureForScenario(scenario),
		Source:      sourceMap(scenario.FileID),
		Members:     map[string]any{"steps": scenarioStepRefs(scenario)},
		References:  refs.References,
		Diagnostics: []semantic.Diagnostic{},
	}, nil
}

func scenarioStepRefs(scenario *semantic.SequenceScenario) []ScenarioStepRef {
	out := make([]ScenarioStepRef, 0, len(scenario.Steps))
	for i, step := range scenario.Steps {
		transition := transitionRefFromSemantic(step.Transition)
		var action *string
		if transition.Action != "" {
			value := transition.Action
			action = &value
		}
		out = append(out, ScenarioStepRef{
			Index:           i + 1,
			FromState:       step.FromState,
			Via:             step.Via,
			Guard:           step.Guard,
			GuardExactMatch: step.Guard == step.Transition.Guard,
			Transition:      transition,
			Action:          action,
		})
	}
	return out
}

func (s *Service) inspectAsset(req InspectRequest) (InspectResponse, error) {
	asset, err := s.assetBySelector(req.Selector)
	if err != nil {
		return InspectResponse{}, err
	}
	refs, err := s.GetReferences(GetReferencesRequest{Selector: req.Selector, Direction: string(semantic.ReferenceDirectionBoth)})
	if err != nil {
		return InspectResponse{}, err
	}
	return InspectResponse{
		Object:      assetObjectRef(asset),
		Signature:   signatureForAsset(asset),
		Source:      sourceMap(asset.FileID),
		References:  refs.References,
		Diagnostics: []semantic.Diagnostic{},
	}, nil
}

func (s *Service) inspectPrivateSubNode(req InspectRequest) (InspectResponse, error) {
	node, err := s.privateSubNodeBySelector(req.Selector)
	if err != nil {
		return InspectResponse{}, err
	}
	sig, doc, err := s.signatureForNode(node)
	if err != nil {
		return InspectResponse{}, err
	}
	refs, err := s.GetReferences(GetReferencesRequest{Selector: req.Selector, Direction: string(semantic.ReferenceDirectionBoth)})
	if err != nil {
		return InspectResponse{}, err
	}
	return InspectResponse{
		Object:      objectRef(node),
		Signature:   sig,
		Doc:         doc,
		Source:      sourceMap(node.GetFileID()),
		References:  refs.References,
		Diagnostics: []semantic.Diagnostic{},
	}, nil
}

func (s *Service) inspectFile(req InspectRequest) (InspectResponse, error) {
	fileID, err := s.fileBySelector(req.Selector)
	if err != nil {
		return InspectResponse{}, err
	}
	kind := s.fileKind(fileID)
	members := s.fileMembers(fileID, kind)
	return InspectResponse{
		Object:      fileObjectRef(fileID, kind),
		Signature:   Signature{"kind": kind},
		Source:      sourceMap(fileID),
		Members:     members,
		Diagnostics: []semantic.Diagnostic{},
	}, nil
}

func (s *Service) eventMembers(event *semantic.Event) map[string]any {
	members := map[string]any{}
	var triggering []TransitionRef
	for _, transitions := range s.project.TransitionsByFile {
		for _, transition := range transitions {
			if transition.Event == event.QID {
				triggering = append(triggering, transitionRefFromSemantic(transition))
			}
		}
	}
	sortTransitionRefs(triggering)
	if len(triggering) > 0 {
		members["triggering_transitions"] = triggering
	}
	if hints := sequenceHintsForEvent(event); len(hints) > 0 {
		members["sequence_hints"] = hints
	}
	return members
}

func (s *Service) fileKind(fileID semantic.FileID) string {
	if source, ok := s.project.SourceFilesByID[fileID]; ok {
		if source.Kind == "node" && s.isStateFile(fileID) {
			return "state_file"
		}
		if source.Kind == "view" && source.ViewAs != "" {
			return source.ViewAs
		}
		return source.Kind
	}
	if s.isStateFile(fileID) {
		return "state_file"
	}
	return "file"
}

func sequenceHintsForEvent(event *semantic.Event) map[string]any {
	switch event.Source {
	case "external":
		return map[string]any{"advisory": true, "participant": "Actor", "actor": event.Actor, "message_label_source": "METHOD path"}
	case "ui":
		return map[string]any{"advisory": true, "participant": "User", "message_label_source": "event id"}
	case "internal":
		return map[string]any{"advisory": true, "participant": "Task", "message_label_source": "event id"}
	case "er":
		return map[string]any{"advisory": true, "participant": "DB", "message_label_source": "watched store"}
	default:
		return nil
	}
}

func transitionRefsForTask(project *semantic.Project, taskID semantic.QualifiedID) []TransitionRef {
	if project == nil {
		return nil
	}
	refs := project.ActionsByTask[taskID]
	out := make([]TransitionRef, 0, len(refs))
	for _, ref := range refs {
		out = append(out, transitionRefFromSemantic(ref.Transition))
	}
	sortTransitionRefs(out)
	return out
}

func sortTransitionRefs(refs []TransitionRef) {
	sort.Slice(refs, func(i, j int) bool { return refs[i].ID < refs[j].ID })
}

func transitionRefFromSemantic(transition semantic.Transition) TransitionRef {
	return TransitionRef{
		Object:    "transition",
		Kind:      "transition",
		ID:        transitionID(transition),
		StateFile: transition.FileID.String(),
		From:      transition.From,
		On:        transition.On,
		To:        transition.To,
		Guard:     transition.Guard,
		Action:    transition.ActionTask.String(),
	}
}

func transitionID(transition semantic.Transition) string {
	return semantic.TransitionID(transition)
}

func flowEntriesSummary(entries []semantic.FlowEntry) []map[string]any {
	out := make([]map[string]any, 0, len(entries))
	for _, entry := range entries {
		summary := map[string]any{"kind": string(entry.Kind)}
		switch entry.Kind {
		case semantic.FlowKindStep:
			summary["step"] = entry.Step.TaskID
		case semantic.FlowKindForeach:
			summary["foreach"] = entry.Foreach.TaskID
			summary["mode"] = entry.Foreach.Mode
		case semantic.FlowKindFork:
			summary["fork"] = entry.Fork.ForkID
			summary["join"] = entry.Fork.JoinID
		case semantic.FlowKindBranch:
			summary["branch"] = entry.Branch.BranchID
		}
		out = append(out, summary)
	}
	return out
}
