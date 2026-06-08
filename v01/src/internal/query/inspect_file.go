package query

import "github.com/hiroshiasayadev-prog/brewprint/v01/src/internal/semantic"

func (s *Service) fileMembers(fileID semantic.FileID, kind string) map[string]any {
	members := map[string]any{}
	if nodes := s.fileNodeRefs(fileID); len(nodes) > 0 {
		members["nodes"] = nodes
	}
	if mainNode := s.fileMainNode(fileID); mainNode.ID != "" {
		members["main_node"] = mainNode
	}
	if flow := s.project.FlowByFile[fileID]; len(flow) > 0 {
		members["flow"] = map[string]any{"entries": flowEntriesSummary(flow), "schema_status": "draft"}
	}
	if transitions := s.project.TransitionsByFile[fileID]; len(transitions) > 0 {
		members["transitions"] = transitionRefsFromSemantic(transitions)
	}
	if kind == "state_file" || len(s.project.TransitionsByFile[fileID]) > 0 {
		members["states"] = s.fileNodesByKind(fileID, semantic.NodeKindState)
		members["events"] = s.fileNodesByKind(fileID, semantic.NodeKindEvent)
		members["wireframes"] = s.fileWireframePresence(fileID)
	}
	if kind == "sequence_diagram" {
		for _, scenario := range s.project.ScenariosByID {
			if scenario.FileID == fileID {
				members["view"] = scenarioObjectRef(scenario)
				members["state_file"] = scenario.StateFile.String()
				members["steps"] = scenarioStepRefs(scenario)
			}
		}
	}
	if kind == "api_table" {
		for _, view := range s.project.APIViewsByID {
			if view.FileID == fileID {
				members["view"] = apiViewObjectRef(view)
				members["http_root_path"] = view.HTTPRootPath
				members["modules"] = signatureForAPIView(view)["modules"]
			}
		}
	}
	if kind == "er_diagram" {
		for _, view := range s.project.ERViewsByID {
			if view.FileID == fileID {
				members["view"] = erViewObjectRef(view)
				members["modules"] = signatureForERView(view)["modules"]
			}
		}
	}
	if kind == "render_index" {
		members["groups"] = s.project.RenderGroups
	}
	return members
}

func (s *Service) fileNodeRefs(fileID semantic.FileID) []ObjectRef {
	out := make([]ObjectRef, 0, len(s.project.NodesByFile[fileID]))
	for _, node := range s.project.NodesByFile[fileID] {
		if model, ok := node.(*semantic.Model); ok && model.FilePrivate {
			continue
		}
		out = append(out, listObjectRef(objectRef(node)))
	}
	return out
}

func (s *Service) fileMainNode(fileID semantic.FileID) ObjectRef {
	qid := s.project.MainNodeByFile[fileID]
	if qid == "" {
		return ObjectRef{}
	}
	if node := s.project.NodesByQID[qid]; node != nil {
		return listObjectRef(objectRef(node))
	}
	return ObjectRef{}
}

func (s *Service) fileNodesByKind(fileID semantic.FileID, kind semantic.NodeKind) []ObjectRef {
	var out []ObjectRef
	for _, node := range s.project.NodesByFile[fileID] {
		if model, ok := node.(*semantic.Model); ok && model.FilePrivate {
			continue
		}
		if node.GetKind() == kind {
			out = append(out, listObjectRef(objectRef(node)))
		}
	}
	return out
}

func (s *Service) isStateFile(fileID semantic.FileID) bool {
	if len(s.project.TransitionsByFile[fileID]) > 0 {
		return true
	}
	for _, node := range s.project.NodesByFile[fileID] {
		if node.GetKind() == semantic.NodeKindState || node.GetKind() == semantic.NodeKindEvent {
			return true
		}
	}
	return false
}

func (s *Service) fileWireframePresence(fileID semantic.FileID) []map[string]any {
	var out []map[string]any
	for _, node := range s.project.NodesByFile[fileID] {
		state, ok := node.(*semantic.State)
		if !ok {
			continue
		}
		out = append(out, map[string]any{"state": state.QID.String(), "present": state.Wireframe != nil})
	}
	return out
}

func transitionRefsFromSemantic(transitions []semantic.Transition) []TransitionRef {
	out := make([]TransitionRef, 0, len(transitions))
	for _, transition := range transitions {
		out = append(out, transitionRefFromSemantic(transition))
	}
	sortTransitionRefs(out)
	return out
}
