package query

import (
	"sort"

	"github.com/hiroshiasayadev-prog/brewprint/internal/semantic"
)

func (s *Service) Inspect(req InspectRequest) (InspectResponse, error) {
	if req.Detail != "" && req.Detail != "brief" && req.Detail != "normal" && req.Detail != "full" {
		return InspectResponse{}, errUnsupportedDetail(req.Detail)
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
		subTasks = append(subTasks, map[string]any{
			"object":   "node",
			"kind":     "task",
			"id":       sub.QID.String(),
			"file":     sub.FileID.String(),
			"local_id": sub.ID,
			"label":    sub.ID,
			"signature": Signature{
				"reads":  storeIDs(sub.Reads),
				"writes": storeIDs(sub.Writes),
			},
			"source": sourceMap(sub.FileID),
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
