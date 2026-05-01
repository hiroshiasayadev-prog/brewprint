package query

import (
	"fmt"
	"sort"
	"strings"

	"github.com/hiroshiasayadev-prog/brewprint/internal/semantic"
)

const (
	impactKindFlowStepTask       = "flow_step_task"
	impactKindSequenceStepAction = "sequence_step_action"
	sourceLocationUnavailable    = "source_location_unavailable"
)

type flowImpactHit struct {
	fileID      semantic.FileID
	owner       ObjectRef
	taskID      semantic.QualifiedID
	rawTaskID   string
	location    *SourceLocation
	description string
}

func (h flowImpactHit) line() int {
	if h.location == nil {
		return 0
	}
	return h.location.Line
}

func (s *Service) collectTaskAnalyzeImpacts(req AnalyzeImpactRequest, targetKey semantic.ObjectKey, target ObjectRef) ([]ImpactEntry, []semantic.Diagnostic) {
	if target.Object != "node" || target.Kind != "task" {
		return nil, nil
	}
	if req.Change.Kind != AnalyzeImpactChangeRename && req.Change.Kind != AnalyzeImpactChangeRemove && req.Change.Kind != AnalyzeImpactChangeContract {
		return nil, nil
	}

	taskID := taskQIDFromObjectKey(targetKey)
	if taskID == "" {
		return nil, nil
	}

	collector := &taskImpactCollector{
		service:      s,
		request:      req,
		target:       target,
		targetKey:    targetKey,
		taskID:       taskID,
		taskIDString: taskID.String(),
		seen:         map[string]struct{}{},
	}
	collector.collectTransitionActions()
	collector.collectFlowSteps()
	collector.collectSequenceSteps()
	collector.finish()
	return collector.impacts, collector.diagnostics
}

type taskImpactCollector struct {
	service      *Service
	request      AnalyzeImpactRequest
	target       ObjectRef
	targetKey    semantic.ObjectKey
	taskID       semantic.QualifiedID
	taskIDString string
	impacts      []ImpactEntry
	diagnostics  []semantic.Diagnostic
	seen         map[string]struct{}
}

func (c *taskImpactCollector) addImpact(entry ImpactEntry) {
	if !c.service.impactInScope(entry, c.target, c.request.ScopeModules) {
		return
	}
	if entry.Source != nil && entry.Source.File != "" && (entry.Source.Line == 0 || entry.Source.Column == 0) {
		entry.Fixability = "unknown"
		entry.SuggestedFixes = nil
		c.addSourceLocationUnavailableDiagnostic(entry.Source.File)
	}
	key := impactDedupeKey(entry)
	if _, ok := c.seen[key]; ok {
		return
	}
	c.seen[key] = struct{}{}
	c.impacts = append(c.impacts, entry)
}

func (c *taskImpactCollector) addSourceLocationUnavailableDiagnostic(file string) {
	for _, diagnostic := range c.diagnostics {
		if diagnostic.Code == sourceLocationUnavailable && diagnostic.FileID.String() == file {
			return
		}
	}
	c.diagnostics = append(c.diagnostics, semantic.Diagnostic{
		Severity: semantic.SeverityWarning,
		Code:     sourceLocationUnavailable,
		FileID:   semantic.FileID(file),
		Message:  "source line/column is unavailable; impact fixability was downgraded",
	})
}

func (c *taskImpactCollector) finish() {
	sort.Slice(c.impacts, func(i, j int) bool {
		if c.impacts[i].Kind != c.impacts[j].Kind {
			return c.impacts[i].Kind < c.impacts[j].Kind
		}
		if c.impacts[i].Object.File != c.impacts[j].Object.File {
			return c.impacts[i].Object.File < c.impacts[j].Object.File
		}
		if c.impacts[i].Object.ID != c.impacts[j].Object.ID {
			return c.impacts[i].Object.ID < c.impacts[j].Object.ID
		}
		return sourceSortKey(c.impacts[i].Source) < sourceSortKey(c.impacts[j].Source)
	})
	sort.Slice(c.diagnostics, func(i, j int) bool {
		if c.diagnostics[i].Code != c.diagnostics[j].Code {
			return c.diagnostics[i].Code < c.diagnostics[j].Code
		}
		return c.diagnostics[i].FileID < c.diagnostics[j].FileID
	})
}

func (c *taskImpactCollector) collectTransitionActions() {
	refs := append([]semantic.Reference{}, c.service.project.ReferencesByTarget[c.targetKey]...)
	sort.Slice(refs, func(i, j int) bool {
		if refs[i].Kind != refs[j].Kind {
			return refs[i].Kind < refs[j].Kind
		}
		return refs[i].From.ID < refs[j].From.ID
	})
	for _, ref := range refs {
		if ref.Kind != semantic.ReferenceKindTransitionAction {
			continue
		}
		object := objectRefFromReferenceEndpoint(ref.From)
		source := c.sourceForTransitionAction(ref.From.ID)
		c.addImpact(c.makeTaskReferenceImpact(
			"transition_action",
			object,
			[]string{string(semantic.ReferenceKindTransitionAction)},
			source,
			fmt.Sprintf("transition %s が task %s を action として呼び出している", object.ID, c.taskIDString),
			fmt.Sprintf("transition action reference を task %s の変更に合わせて見直す", c.taskIDString),
		))
	}
}

func (c *taskImpactCollector) collectFlowSteps() {
	var fileIDs []string
	for fileID := range c.service.project.FlowByFile {
		fileIDs = append(fileIDs, fileID.String())
	}
	sort.Strings(fileIDs)

	for _, rawFileID := range fileIDs {
		fileID := semantic.FileID(rawFileID)
		entries := c.service.project.FlowByFile[fileID]
		owner := c.flowOwnerObject(fileID)
		for _, entry := range entries {
			for _, hit := range c.flowHitsForEntry(fileID, owner, entry, "flow") {
				c.addImpact(c.makeTaskReferenceImpact(
					impactKindFlowStepTask,
					hit.owner,
					[]string{},
					hit.location,
					fmt.Sprintf("%s の flow step が task %s を呼び出している", hit.owner.ID, c.taskIDString),
					fmt.Sprintf("flow step task reference を task %s の変更に合わせて見直す", c.taskIDString),
				))
			}
		}
	}
}

func (c *taskImpactCollector) flowHitsForEntry(fileID semantic.FileID, owner ObjectRef, entry semantic.FlowEntry, path string) []flowImpactHit {
	var hits []flowImpactHit
	switch entry.Kind {
	case semantic.FlowKindStep:
		hits = append(hits, c.flowHitForStep(fileID, owner, entry.Step, path+".step")...)
	case semantic.FlowKindForeach:
		if entry.Foreach.Task == c.taskID {
			hits = append(hits, flowImpactHit{
				fileID:      fileID,
				owner:       owner,
				taskID:      entry.Foreach.Task,
				rawTaskID:   entry.Foreach.TaskID,
				location:    c.sourceForFlowTask(fileID, entry.Foreach.TaskID),
				description: path + ".foreach",
			})
		}
	case semantic.FlowKindFork:
		for branchIndex, branch := range entry.Fork.Branches {
			for stepIndex, step := range branch.Steps {
				hits = append(hits, c.flowHitForStep(fileID, owner, step, fmt.Sprintf("%s.fork.branch[%d].step[%d]", path, branchIndex+1, stepIndex+1))...)
			}
		}
	case semantic.FlowKindBranch:
		for caseIndex, branchCase := range entry.Branch.Cases {
			hits = append(hits, c.flowHitForStep(fileID, owner, branchCase.Step, fmt.Sprintf("%s.branch.case[%d]", path, caseIndex+1))...)
		}
	}
	return hits
}

func (c *taskImpactCollector) flowHitForStep(fileID semantic.FileID, owner ObjectRef, step semantic.StepFlow, path string) []flowImpactHit {
	if step.Task != c.taskID {
		return nil
	}
	return []flowImpactHit{{
		fileID:      fileID,
		owner:       owner,
		taskID:      step.Task,
		rawTaskID:   step.TaskID,
		location:    c.sourceForFlowTask(fileID, step.TaskID),
		description: path,
	}}
}

func (c *taskImpactCollector) collectSequenceSteps() {
	var scenarios []*semantic.SequenceScenario
	for _, scenario := range c.service.project.ScenariosByID {
		scenarios = append(scenarios, scenario)
	}
	sort.Slice(scenarios, func(i, j int) bool { return scenarios[i].ID < scenarios[j].ID })

	for _, scenario := range scenarios {
		for stepIndex, step := range scenario.Steps {
			if step.Transition.ActionTask != c.taskID {
				continue
			}
			object := scenarioObjectRef(scenario)
			source := c.sourceForScenarioStep(scenario, stepIndex+1)
			c.addImpact(c.makeTaskReferenceImpact(
				impactKindSequenceStepAction,
				object,
				[]string{string(semantic.ReferenceKindScenarioStepTransition), string(semantic.ReferenceKindTransitionAction)},
				source,
				fmt.Sprintf("sequence scenario %s step %d が action task %s を含む transition を辿っている", scenario.ID, stepIndex+1, c.taskIDString),
				fmt.Sprintf("scenario step action usage を task %s の変更に合わせて見直す", c.taskIDString),
			))
		}
	}
}

func (c *taskImpactCollector) makeTaskReferenceImpact(kind string, object ObjectRef, via []string, source *SourceLocation, reason string, action string) ImpactEntry {
	impact := ImpactEntry{
		Kind:              kind,
		Object:            object,
		Reason:            reason,
		Via:               via,
		Source:            source,
		RecommendedAction: action,
	}

	switch c.request.Change.Kind {
	case AnalyzeImpactChangeRename:
		impact.Severity = "breaking"
		impact.Fixability = "suggested"
		if sourceLocationPrecise(source) {
			fix := SuggestedFix{
				Kind:       "replace_reference",
				Confidence: "medium",
				From:       c.taskIDString,
				To:         c.request.Change.NewID,
				Source:     source,
			}
			gate := c.service.mechanicalJudgementGate(mechanicalJudgementInput{
				From:                    c.taskIDString,
				To:                      c.request.Change.NewID,
				Source:                  source,
				YAMLStructurePreserving: true,
				ReferenceStable:         false,
			})
			if gate.Mechanical {
				impact.Fixability = "mechanical"
				fix.Confidence = "high"
			}
			impact.SuggestedFixes = []SuggestedFix{fix}
		}
	case AnalyzeImpactChangeRemove:
		impact.Severity = "breaking"
		impact.Fixability = "manual_review"
	case AnalyzeImpactChangeContract:
		impact.Severity = "warning"
		impact.Fixability = "manual_review"
	}
	return impact
}

func (c *taskImpactCollector) sourceForTransitionAction(transitionID string) *SourceLocation {
	transition, err := c.service.transitionBySelector(Selector{Object: "transition", ID: transitionID})
	if err != nil {
		return nil
	}
	block := c.service.findTransitionSource(transition)
	if block.text == "" {
		return sourceLocationFromBlock(transition.FileID, sourceBlock{})
	}
	lines := splitSourceLines(block.text)
	for i, line := range lines {
		key, value, ok := yamlKeyValue(line)
		if ok && key == "action" && value == c.taskIDString {
			absoluteLine := block.startLine + i
			column := indentOf(line) + 1
			return &SourceLocation{File: transition.FileID.String(), Line: absoluteLine, Column: column, EndLine: absoluteLine, EndColumn: column + len(strings.TrimSpace(line))}
		}
	}
	return sourceLocationFromBlock(transition.FileID, block)
}

func (c *taskImpactCollector) sourceForFlowTask(fileID semantic.FileID, rawTaskID string) *SourceLocation {
	content, ok := c.service.fileContent(fileID)
	if !ok {
		return sourceLocationFromBlock(fileID, sourceBlock{})
	}
	lines := splitSourceLines(content)
	start, end := topLevelSectionRange(lines, "flow")
	if start < 0 {
		return sourceLocationFromBlock(fileID, sourceBlock{})
	}
	for i := start; i < end; i++ {
		key, value, ok := yamlKeyValue(lines[i])
		if !ok {
			continue
		}
		if (key == "step" || key == "foreach") && value == rawTaskID {
			column := indentOf(lines[i]) + 1
			return &SourceLocation{File: fileID.String(), Line: i + 1, Column: column, EndLine: i + 1, EndColumn: column + len(strings.TrimSpace(lines[i]))}
		}
	}
	return sourceLocationFromBlock(fileID, sourceBlock{})
}

func (c *taskImpactCollector) sourceForScenarioStep(scenario *semantic.SequenceScenario, stepIndex int) *SourceLocation {
	if scenario == nil {
		return nil
	}
	content, ok := c.service.fileContent(scenario.FileID)
	if !ok {
		return sourceLocationFromBlock(scenario.FileID, sourceBlock{})
	}
	lines := splitSourceLines(content)
	start, end := topLevelSectionRange(lines, "steps")
	if start < 0 {
		return sourceLocationFromBlock(scenario.FileID, sourceBlock{})
	}
	current := 0
	for i := start; i < end; i++ {
		trimmed := strings.TrimSpace(lines[i])
		if !strings.HasPrefix(trimmed, "- ") && trimmed != "-" {
			continue
		}
		itemEnd := nextSequenceItemOrSectionEnd(lines, i+1, end, indentOf(lines[i]))
		current++
		if current == stepIndex {
			block := makeBlock(lines, i, itemEnd)
			return sourceLocationFromBlock(scenario.FileID, block)
		}
		i = itemEnd - 1
	}
	return sourceLocationFromBlock(scenario.FileID, sourceBlock{})
}

func (c *taskImpactCollector) flowOwnerObject(fileID semantic.FileID) ObjectRef {
	if mainQID := c.service.project.MainNodeByFile[fileID]; mainQID != "" {
		if node := c.service.project.NodesByQID[mainQID]; node != nil {
			return objectRef(node)
		}
	}
	return fileObjectRef(fileID, c.service.fileKind(fileID))
}

func taskQIDFromObjectKey(key semantic.ObjectKey) semantic.QualifiedID {
	raw := string(key)
	if raw == "" || strings.Contains(raw, ":") {
		return ""
	}
	return semantic.QualifiedID(raw)
}

func impactDedupeKey(entry ImpactEntry) string {
	return strings.Join([]string{
		entry.Kind,
		entry.Object.Object,
		entry.Object.Kind,
		entry.Object.ID,
		strings.Join(entry.Via, ">"),
		sourceSortKey(entry.Source),
	}, "\x00")
}
