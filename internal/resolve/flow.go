package resolve

import (
	"sort"
	"strings"

	"github.com/hiroshiasayadev-prog/brewprint/internal/rawyaml"
	"github.com/hiroshiasayadev-prog/brewprint/internal/semantic"
)

func buildFlows(raw *rawyaml.Project, project *semantic.Project, symbols *symbolTable) {
	for _, file := range raw.Files {
		if file.Kind != rawyaml.FileKindNode || file.NodeFile == nil || len(file.NodeFile.Flow) == 0 {
			continue
		}
		fileID := semantic.FileID(file.ID)
		var entries []semantic.FlowEntry
		for _, rawEntry := range file.NodeFile.Flow {
			entry, ok := buildFlowEntry(project, symbols, fileID, rawEntry)
			if ok {
				entries = append(entries, entry)
			}
		}
		project.FlowByFile[fileID] = entries
	}
}

func buildFlowEntry(project *semantic.Project, symbols *symbolTable, fileID semantic.FileID, raw rawyaml.FlowEntry) (semantic.FlowEntry, bool) {
	switch {
	case raw.Step != "":
		return semantic.FlowEntry{Kind: semantic.FlowKindStep, Step: buildStepFlow(project, symbols, fileID, raw.Step, raw.Params)}, true
	case raw.Foreach != "":
		return semantic.FlowEntry{Kind: semantic.FlowKindForeach, Foreach: buildForeachFlow(project, symbols, fileID, raw)}, true
	case raw.Fork != "":
		return semantic.FlowEntry{Kind: semantic.FlowKindFork, Fork: buildForkFlow(project, symbols, fileID, raw)}, true
	case raw.Branch != "":
		return semantic.FlowEntry{Kind: semantic.FlowKindBranch, Branch: buildBranchFlow(project, symbols, fileID, raw)}, true
	default:
		symbols.addDiagnosticCode(semantic.SeverityWarning, diagnosticUnsupportedFlowEntry, fileID, "unsupported empty flow entry")
		return semantic.FlowEntry{}, false
	}
}

func buildStepFlow(project *semantic.Project, symbols *symbolTable, fileID semantic.FileID, id string, params map[string]string) semantic.StepFlow {
	qid := resolveNodeQID(project, fileID, semantic.NodeKindTask, id)
	if qid == "" {
		symbols.addDiagnosticCode(semantic.SeverityError, diagnosticUnresolvedFlowTask, fileID, "unresolved flow step: "+id)
	}
	return semantic.StepFlow{
		Task:   qid,
		TaskID: id,
		Params: buildParamWirings(project, fileID, params),
	}
}

func buildForeachFlow(project *semantic.Project, symbols *symbolTable, fileID semantic.FileID, raw rawyaml.FlowEntry) semantic.ForeachFlow {
	qid := resolveNodeQID(project, fileID, semantic.NodeKindTask, raw.Foreach)
	if qid == "" {
		symbols.addDiagnosticCode(semantic.SeverityError, diagnosticUnresolvedFlowTask, fileID, "unresolved foreach task: "+raw.Foreach)
	}
	mode := raw.Mode
	if mode == "" {
		mode = "sequential"
	}
	return semantic.ForeachFlow{
		Task:    qid,
		TaskID:  raw.Foreach,
		Over:    resolveFlowSource(project, fileID, raw.Over),
		Mode:    mode,
		Params:  buildParamWirings(project, fileID, raw.Params),
		Returns: raw.Returns,
	}
}

func buildForkFlow(project *semantic.Project, symbols *symbolTable, fileID semantic.FileID, raw rawyaml.FlowEntry) semantic.ForkFlow {
	forkQID := resolveNodeQID(project, fileID, semantic.NodeKindFork, raw.Fork)
	if forkQID == "" {
		symbols.addDiagnosticCode(semantic.SeverityError, diagnosticUnresolvedFlowNode, fileID, "unresolved fork node: "+raw.Fork)
	}
	joinQID := resolveNodeQID(project, fileID, semantic.NodeKindJoin, raw.Join)
	if joinQID == "" {
		symbols.addDiagnosticCode(semantic.SeverityError, diagnosticUnresolvedFlowNode, fileID, "unresolved join node: "+raw.Join)
	}

	flow := semantic.ForkFlow{Fork: forkQID, ForkID: raw.Fork, Join: joinQID, JoinID: raw.Join}
	for _, rawBranch := range raw.Branches {
		branch := semantic.ForkBranchFlow{}
		for _, rawStep := range rawBranch.Steps {
			if rawStep.Step == "" {
				symbols.addDiagnosticCode(semantic.SeverityError, diagnosticInvalidFlowBranch, fileID, "fork branch step is missing step id")
				continue
			}
			branch.Steps = append(branch.Steps, buildStepFlow(project, symbols, fileID, rawStep.Step, rawStep.Params))
		}
		flow.Branches = append(flow.Branches, branch)
	}
	flow.JoinParams = buildJoinParamWirings(project, symbols, fileID, flow)
	return flow
}

func buildJoinParamWirings(project *semantic.Project, symbols *symbolTable, fileID semantic.FileID, fork semantic.ForkFlow) []semantic.ParamWiring {
	join := project.JoinsByQID[fork.Join]
	if join == nil {
		return nil
	}
	terminalByReturn := map[string]semantic.StepFlow{}
	for _, branch := range fork.Branches {
		if len(branch.Steps) == 0 {
			continue
		}
		terminal := branch.Steps[len(branch.Steps)-1]
		task := project.TasksByQID[terminal.Task]
		if task == nil || task.Returns == nil {
			continue
		}
		terminalByReturn[task.Returns.Name] = terminal
	}
	var out []semantic.ParamWiring
	for _, param := range join.Params {
		terminal, ok := terminalByReturn[param.Name]
		if !ok {
			symbols.addDiagnosticCode(semantic.SeverityError, diagnosticUnmatchedJoinParam, fileID, "join param has no matching branch return: "+param.Name)
			continue
		}
		out = append(out, semantic.ParamWiring{
			TargetParam: param.Name,
			Source: semantic.FlowSource{
				Kind:      semantic.FlowSourceNode,
				Raw:       terminal.TaskID,
				Node:      terminal.Task,
				AssetName: param.Name,
			},
		})
	}
	return out
}

func buildBranchFlow(project *semantic.Project, symbols *symbolTable, fileID semantic.FileID, raw rawyaml.FlowEntry) semantic.BranchFlow {
	branchQID := resolveNodeQID(project, fileID, semantic.NodeKindBranch, raw.Branch)
	if branchQID == "" {
		symbols.addDiagnosticCode(semantic.SeverityError, diagnosticUnresolvedFlowNode, fileID, "unresolved branch node: "+raw.Branch)
	}
	flow := semantic.BranchFlow{
		Branch:   branchQID,
		BranchID: raw.Branch,
		Params:   buildParamWirings(project, fileID, raw.Params),
	}
	for _, rawCase := range raw.Cases {
		flow.Cases = append(flow.Cases, semantic.BranchCaseFlow{
			Label: rawCase.Label,
			Step:  buildStepFlow(project, symbols, fileID, rawCase.Step, rawCase.Params),
		})
	}
	return flow
}

func buildParamWirings(project *semantic.Project, fileID semantic.FileID, params map[string]string) []semantic.ParamWiring {
	if len(params) == 0 {
		return nil
	}
	keys := make([]string, 0, len(params))
	for key := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	out := make([]semantic.ParamWiring, 0, len(keys))
	for _, key := range keys {
		out = append(out, semantic.ParamWiring{TargetParam: key, Source: resolveFlowSource(project, fileID, params[key])})
	}
	return out
}

func resolveFlowSource(project *semantic.Project, fileID semantic.FileID, raw string) semantic.FlowSource {
	if strings.HasPrefix(raw, "$params.") {
		return semantic.FlowSource{Kind: semantic.FlowSourceParam, Raw: raw, ParamName: strings.TrimPrefix(raw, "$params.")}
	}
	if raw == "$item" {
		return semantic.FlowSource{Kind: semantic.FlowSourceItem, Raw: raw}
	}
	qid := resolveAnyNodeQID(project, fileID, raw)
	source := semantic.FlowSource{Kind: semantic.FlowSourceNode, Raw: raw, Node: qid}
	if task := project.TasksByQID[qid]; task != nil && task.Returns != nil {
		source.AssetName = task.Returns.Name
	}
	if join := project.JoinsByQID[qid]; join != nil && join.Returns != nil {
		source.AssetName = join.Returns.Name
	}
	return source
}

func resolveNodeQID(project *semantic.Project, fileID semantic.FileID, kind semantic.NodeKind, ref string) semantic.QualifiedID {
	if ref == "" {
		return ""
	}
	if isFullQID(ref, string(kind)) {
		qid := semantic.QualifiedID(ref)
		if node := project.NodesByQID[qid]; node != nil && node.GetKind() == kind {
			return qid
		}
		return ""
	}
	for _, node := range project.NodesByFile[fileID] {
		if node.GetKind() == kind && node.GetID() == ref {
			return node.GetQID()
		}
	}
	return ""
}

func resolveAnyNodeQID(project *semantic.Project, fileID semantic.FileID, ref string) semantic.QualifiedID {
	if ref == "" {
		return ""
	}
	if strings.Contains(ref, ".") {
		qid := semantic.QualifiedID(ref)
		if _, ok := project.NodesByQID[qid]; ok {
			return qid
		}
	}
	for _, node := range project.NodesByFile[fileID] {
		if node.GetID() == ref {
			return node.GetQID()
		}
	}
	return ""
}
