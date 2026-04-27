package dag

import (
	"fmt"
	"strings"

	"github.com/hiroshiasayadev-prog/brewprint/internal/semantic"
)

func renderFlowTask(project *semantic.Project, task *semantic.Task) string {
	r := newFlowRenderer(project, task)
	return r.render()
}

type flowRenderer struct {
	project *semantic.Project
	main    *semantic.Task
	b       strings.Builder

	defined map[string]bool
	details map[semantic.QualifiedID]bool

	taskClasses   []string
	assetClasses  []string
	storeClasses  []string
	branchClasses []string
	forkClasses   []string
	classSeen     map[string]map[string]bool
	detailQIDs    []semantic.QualifiedID
}

func newFlowRenderer(project *semantic.Project, main *semantic.Task) *flowRenderer {
	r := &flowRenderer{
		project:   project,
		main:      main,
		defined:   map[string]bool{},
		details:   map[semantic.QualifiedID]bool{},
		classSeen: map[string]map[string]bool{},
	}
	for _, param := range main.Params {
		r.defined[param.Name] = true
	}
	if main.Returns != nil {
		r.defined[main.Returns.Name] = true
	}
	r.addDetail(main.QID)
	return r
}

func (r *flowRenderer) render() string {
	r.writeHeader()
	r.b.WriteString("```mermaid\n")
	r.b.WriteString("flowchart TD\n")
	writeBoundary(&r.b, "params", paramNames(r.main))
	writeBoundary(&r.b, "returns", returnNames(r.main))
	if len(r.main.Params) > 0 || r.main.Returns != nil {
		r.blank()
	}

	flows := r.project.FlowByFile[r.main.FileID]
	prev := ""
	for i, entry := range flows {
		start := r.entryStartRef(entry)
		if i == 0 {
			r.line("_start([Start]) ==> %s", start)
		} else if prev != "" {
			r.line("%s ==> %s", prev, start)
		}
		last := i == len(flows)-1
		prev = r.renderFlowEntry(entry, last)
		if !last {
			r.blank()
		}
	}
	if len(flows) > 0 {
		r.blank()
	}

	r.writeClassDefs()
	r.b.WriteString("```\n\n")
	r.writeDetails()
	return r.b.String()
}

func (r *flowRenderer) writeHeader() {
	fmt.Fprintf(&r.b, "# %s\n\n", r.main.ID)
	if r.main.Endpoint {
		method := r.main.Method
		leafPath := r.main.Path
		if leafPath == "" {
			leafPath = r.main.ID
		}
		leafPath = strings.TrimPrefix(leafPath, "/")
		fmt.Fprintf(&r.b, "**API**: [%s /api/%s](../_cross/api.md)\n\n", method, leafPath)
	}
	if note := strings.TrimRight(r.main.Note, "\n"); note != "" {
		r.b.WriteString(note)
		r.b.WriteString("\n\n")
	}
}

func (r *flowRenderer) entryStartRef(entry semantic.FlowEntry) string {
	switch entry.Kind {
	case semantic.FlowKindStep:
		return r.taskRef(entry.Step.Task, false)
	case semantic.FlowKindForeach:
		return r.taskRef(entry.Foreach.Task, true)
	case semantic.FlowKindFork:
		return r.forkRef(entry.Fork.Fork)
	case semantic.FlowKindBranch:
		return r.branchRef(entry.Branch.Branch)
	default:
		return ""
	}
}

func (r *flowRenderer) renderFlowEntry(entry semantic.FlowEntry, last bool) string {
	switch entry.Kind {
	case semantic.FlowKindStep:
		return r.renderStep(entry.Step, last)
	case semantic.FlowKindForeach:
		return r.renderForeach(entry.Foreach)
	case semantic.FlowKindFork:
		return r.renderFork(entry.Fork)
	case semantic.FlowKindBranch:
		return r.renderBranch(entry.Branch)
	default:
		return ""
	}
}

func (r *flowRenderer) renderStep(step semantic.StepFlow, last bool) string {
	task := r.project.TasksByQID[step.Task]
	if task == nil {
		return step.TaskID
	}
	r.addDetail(task.QID)
	r.writeParamDataLines(step.Params, task.QID)
	r.writeStoreAccess(task)
	r.writeTaskReturn(task)
	if last {
		r.line("%s ==> _end([End])", task.ID)
	}
	return task.ID
}

func (r *flowRenderer) renderForeach(flow semantic.ForeachFlow) string {
	task := r.project.TasksByQID[flow.Task]
	if task == nil {
		return flow.TaskID
	}
	r.addDetail(task.QID)
	r.line("%s --\"foreach\"--> %s", r.sourceNodeID(flow.Over), task.ID)
	if len(task.Reads) > 0 || len(task.Writes) > 0 {
		r.blank()
		r.writeStoreAccess(task)
	}
	if flow.Returns != "" {
		r.blank()
		r.line("%s --> %s", task.ID, flow.Returns)
	}
	r.line("%s ==> _end([End])", task.ID)
	return task.ID
}

func (r *flowRenderer) renderFork(flow semantic.ForkFlow) string {
	fork := r.project.ForksByQID[flow.Fork]
	join := r.project.JoinsByQID[flow.Join]

	for _, branch := range flow.Branches {
		if len(branch.Steps) == 0 {
			continue
		}
		first := branch.Steps[0]
		task := r.project.TasksByQID[first.Task]
		if task == nil {
			continue
		}
		r.addDetail(task.QID)
		r.writeParamDataLines(first.Params, task.QID)
	}
	r.blank()
	for _, branch := range flow.Branches {
		if len(branch.Steps) == 0 {
			continue
		}
		first := branch.Steps[0]
		r.line("%s == \"parallel\" ==> %s", r.nodeID(flow.Fork), r.taskRef(first.Task, false))
	}
	r.blank()
	for _, branch := range flow.Branches {
		for _, step := range branch.Steps {
			task := r.project.TasksByQID[step.Task]
			if task == nil {
				continue
			}
			r.addDetail(task.QID)
			r.writeStoreAccess(task)
			r.writeTaskReturn(task)
		}
	}
	r.blank()
	for _, branch := range flow.Branches {
		if len(branch.Steps) == 0 {
			continue
		}
		terminal := branch.Steps[len(branch.Steps)-1]
		task := r.project.TasksByQID[terminal.Task]
		if task == nil {
			continue
		}
		r.line("%s ==> %s", task.ID, r.forkRef(flow.Join))
	}
	for _, joinParam := range flow.JoinParams {
		r.line("%s --> %s", r.sourceNodeID(joinParam.Source), r.nodeID(flow.Join))
	}
	if fork != nil {
		r.addDetail(fork.QID)
	}
	if join != nil {
		r.addDetail(join.QID)
	}
	if join != nil && join.Returns != nil {
		r.blank()
		r.line("%s --> %s", join.ID, join.Returns.Name)
		r.line("%s ==> _end([End])", join.ID)
		return join.ID
	}
	if join != nil {
		r.line("%s ==> _end([End])", join.ID)
		return join.ID
	}
	return flow.JoinID
}

func (r *flowRenderer) renderBranch(flow semantic.BranchFlow) string {
	branch := r.project.BranchesByQID[flow.Branch]
	if branch != nil {
		r.addDetail(branch.QID)
	}
	for _, param := range flow.Params {
		r.line("%s --> %s", r.sourceNodeID(param.Source), r.nodeID(flow.Branch))
	}
	for _, c := range flow.Cases {
		task := r.project.TasksByQID[c.Step.Task]
		if task == nil {
			continue
		}
		r.addDetail(task.QID)
		r.writeParamDataLines(c.Step.Params, task.QID)
	}
	r.blank()
	for _, c := range flow.Cases {
		task := r.project.TasksByQID[c.Step.Task]
		if task == nil {
			continue
		}
		r.line("%s == \"%s\" ==> %s", r.nodeID(flow.Branch), c.Label, r.taskRef(task.QID, false))
	}
	r.blank()
	for _, c := range flow.Cases {
		task := r.project.TasksByQID[c.Step.Task]
		if task == nil {
			continue
		}
		r.writeStoreAccess(task)
	}
	for i, c := range flow.Cases {
		task := r.project.TasksByQID[c.Step.Task]
		if task == nil {
			continue
		}
		endRef := "_end"
		if i == 0 {
			endRef = "_end([End])"
		}
		r.line("%s ==> %s", task.ID, endRef)
	}
	return flow.BranchID
}

func (r *flowRenderer) writeParamDataLines(params []semantic.ParamWiring, target semantic.QualifiedID) {
	for _, param := range params {
		r.line("%s --> %s", r.sourceNodeID(param.Source), r.taskRef(target, false))
	}
}

func (r *flowRenderer) writeTaskReturn(task *semantic.Task) {
	if task.Returns == nil {
		return
	}
	if r.main.Returns != nil && task.Returns.Name == r.main.Returns.Name {
		r.line("%s --> %s", task.ID, task.Returns.Name)
		return
	}
	r.line("%s --> %s", task.ID, r.assetRef(task.Returns.Name))
}

func (r *flowRenderer) writeStoreAccess(task *semantic.Task) {
	for _, read := range task.Reads {
		if hasWrite(task, read.Name) {
			continue
		}
		r.line("%s -- \"read\" --> %s", r.storeRef(read), task.ID)
	}
	for _, write := range task.Writes {
		if hasRead(task, write.Name) {
			r.line("%s <-- \"read/write\" --> %s", task.ID, r.storeRef(write))
			continue
		}
		r.line("%s -- \"write\" --> %s", task.ID, r.storeRef(write))
	}
}

func (r *flowRenderer) sourceNodeID(source semantic.FlowSource) string {
	switch source.Kind {
	case semantic.FlowSourceParam:
		return source.ParamName
	case semantic.FlowSourceNode:
		if source.AssetName != "" {
			return r.assetID(source.AssetName)
		}
		if node := r.project.NodesByQID[source.Node]; node != nil {
			return node.GetID()
		}
	}
	return source.Raw
}

func (r *flowRenderer) taskRef(qid semantic.QualifiedID, foreach bool) string {
	task := r.project.TasksByQID[qid]
	if task == nil {
		return r.nodeID(qid)
	}
	id := task.ID
	r.addClass("task", id)
	if r.defined[id] {
		return id
	}
	r.defined[id] = true
	if foreach {
		return fmt.Sprintf("%s[\"↻ %s\"]", id, id)
	}
	return fmt.Sprintf("%s[%s]", id, id)
}

func (r *flowRenderer) branchRef(qid semantic.QualifiedID) string {
	branch := r.project.BranchesByQID[qid]
	if branch == nil {
		return r.nodeID(qid)
	}
	id := branch.ID
	r.addClass("branch", id)
	if r.defined[id] {
		return id
	}
	r.defined[id] = true
	return fmt.Sprintf("%s{%s}", id, id)
}

func (r *flowRenderer) forkRef(qid semantic.QualifiedID) string {
	id := r.nodeID(qid)
	if id == "" {
		return ""
	}
	r.addClass("fork", id)
	if r.defined[id] {
		return id
	}
	r.defined[id] = true
	return fmt.Sprintf("%s{{%s}}", id, id)
}

func (r *flowRenderer) nodeID(qid semantic.QualifiedID) string {
	if node := r.project.NodesByQID[qid]; node != nil {
		return node.GetID()
	}
	return shortNodeID(qid.String())
}

func (r *flowRenderer) assetID(name string) string {
	if name == "order" {
		return "order_asset"
	}
	return name
}

func (r *flowRenderer) assetRef(name string) string {
	id := r.assetID(name)
	r.addClass("asset", id)
	if r.defined[id] {
		return id
	}
	r.defined[id] = true
	return fmt.Sprintf("%s([%s])", id, name)
}

func (r *flowRenderer) storeRef(ref semantic.StoreRef) string {
	id := shortNodeID(ref.Name)
	label := ref.Name
	if store := r.project.StoresByQID[ref.Store]; store != nil {
		id = store.ID
	}
	if ref.FilePrivate {
		id = shortNodeID(ref.Name)
	}
	r.addClass("store", id)
	if r.defined[id] {
		return id
	}
	r.defined[id] = true
	return fmt.Sprintf("%s[(%s)]", id, label)
}

func shortNodeID(id string) string {
	if id == "" {
		return ""
	}
	parts := strings.Split(id, ".")
	return parts[len(parts)-1]
}

func (r *flowRenderer) addClass(kind, id string) {
	if id == "" {
		return
	}
	if r.classSeen[kind] == nil {
		r.classSeen[kind] = map[string]bool{}
	}
	if r.classSeen[kind][id] {
		return
	}
	r.classSeen[kind][id] = true
	switch kind {
	case "task":
		r.taskClasses = append(r.taskClasses, id)
	case "asset":
		r.assetClasses = append(r.assetClasses, id)
	case "store":
		r.storeClasses = append(r.storeClasses, id)
	case "branch":
		r.branchClasses = append(r.branchClasses, id)
	case "fork":
		r.forkClasses = append(r.forkClasses, id)
	}
}

func (r *flowRenderer) addDetail(qid semantic.QualifiedID) {
	if qid == "" || r.details[qid] {
		return
	}
	r.details[qid] = true
	r.detailQIDs = append(r.detailQIDs, qid)
}

func (r *flowRenderer) writeClassDefs() {
	r.line("classDef taskNode     fill:#4A90D9,stroke:#2C5F8A,color:#fff")
	if len(r.assetClasses) > 0 {
		r.line("classDef assetNode    fill:#5BA55B,stroke:#3A6B3A,color:#fff")
	}
	if len(r.storeClasses) > 0 {
		r.line("classDef storeNode    fill:#E8A838,stroke:#B07820,color:#fff")
	}
	if len(r.branchClasses) > 0 {
		r.line("classDef branchNode   fill:#9B6BBD,stroke:#6B3D8F,color:#fff")
	}
	if len(r.forkClasses) > 0 {
		r.line("classDef forkNode     fill:#8A8A8A,stroke:#5A5A5A,color:#fff")
	}
	r.line("classDef terminalNode fill:#2C2C2C,stroke:#000,color:#fff")
	if len(r.main.Params) > 0 || r.main.Returns != nil {
		r.line("classDef boundaryNode fill:#2D7D9A,stroke:#1A5068,color:#fff")
	}
	if len(r.taskClasses) > 0 {
		r.line("class %s taskNode", strings.Join(r.taskClasses, ","))
	}
	if len(r.assetClasses) > 0 {
		r.line("class %s assetNode", strings.Join(r.assetClasses, ","))
	}
	if len(r.storeClasses) > 0 {
		r.line("class %s storeNode", strings.Join(r.storeClasses, ","))
	}
	if len(r.branchClasses) > 0 {
		r.line("class %s branchNode", strings.Join(r.branchClasses, ","))
	}
	if len(r.forkClasses) > 0 {
		r.line("class %s forkNode", strings.Join(r.forkClasses, ","))
	}
	r.line("class _start,_end terminalNode")
	boundaryNames := append(paramNames(r.main), returnNames(r.main)...)
	if len(boundaryNames) > 0 {
		r.line("class %s boundaryNode", strings.Join(boundaryNames, ","))
	}
}

func (r *flowRenderer) writeDetails() {
	r.b.WriteString("## Tasks\n\n")
	for _, qid := range r.detailQIDs {
		switch node := r.project.NodesByQID[qid].(type) {
		case *semantic.Task:
			writeTaskDetail(&r.b, node, node.QID == r.main.QID)
		case *semantic.Branch:
			writeControlDetail(&r.b, node.ID, node.Note, node.Params, nil)
		case *semantic.Fork:
			writeControlDetail(&r.b, node.ID, node.Note, node.Params, nil)
		case *semantic.Join:
			writeControlDetail(&r.b, node.ID, node.Note, node.Params, node.Returns)
		}
	}
}

func (r *flowRenderer) line(format string, args ...any) {
	fmt.Fprintf(&r.b, "  "+format+"\n", args...)
}

func (r *flowRenderer) blank() {
	r.b.WriteString("\n")
}

func writeTaskDetail(b *strings.Builder, task *semantic.Task, omitNote bool) {
	fmt.Fprintf(b, "### %s\n\n", task.ID)
	if !omitNote {
		if note := strings.TrimRight(task.Note, "\n"); note != "" {
			b.WriteString(note)
			b.WriteString("\n\n")
		}
	}
	writeParamsSection(b, task.Params)
	if task.Returns != nil {
		writeReturnsSection(b, task.Returns)
	}
	accessRows := storeAccessRows(task)
	if len(accessRows) > 0 {
		b.WriteString("#### Store access\n\n")
		b.WriteString("| access | store |\n")
		b.WriteString("|---|---|\n")
		for _, row := range accessRows {
			fmt.Fprintf(b, "| %s | %s |\n", row.Access, row.Store)
		}
		b.WriteString("\n")
	}
}

func writeControlDetail(b *strings.Builder, id string, note string, params []semantic.Param, returns *semantic.Return) {
	fmt.Fprintf(b, "### %s\n\n", id)
	if note := strings.TrimRight(note, "\n"); note != "" {
		b.WriteString(normalizeControlNote(note))
		b.WriteString("\n\n")
	}
	writeParamsSection(b, params)
	if returns != nil {
		writeReturnsSection(b, returns)
	}
}

func normalizeControlNote(note string) string {
	trimmed := strings.TrimSpace(note)
	if trimmed == "" || strings.Contains(trimmed, "\n") {
		return note
	}
	if strings.HasSuffix(trimmed, "。") || strings.HasSuffix(trimmed, ".") || strings.HasSuffix(trimmed, "！") || strings.HasSuffix(trimmed, "!") || strings.HasSuffix(trimmed, "？") || strings.HasSuffix(trimmed, "?") {
		return note
	}
	return note + "。"
}

func writeParamsSection(b *strings.Builder, params []semantic.Param) {
	if len(params) == 0 {
		return
	}
	b.WriteString("#### Params\n\n")
	b.WriteString("| name | model | note |\n")
	b.WriteString("|---|---|---|\n")
	for _, param := range params {
		note := param.Note
		if note == "" {
			note = "—"
		}
		fmt.Fprintf(b, "| %s | %s | %s |\n", param.Name, param.ModelName, note)
	}
	b.WriteString("\n")
}

func writeReturnsSection(b *strings.Builder, ret *semantic.Return) {
	b.WriteString("#### Returns\n\n")
	b.WriteString("| name | model |\n")
	b.WriteString("|---|---|\n")
	fmt.Fprintf(b, "| %s | %s |\n\n", ret.Name, ret.ModelName)
}
