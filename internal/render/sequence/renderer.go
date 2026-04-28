package sequence

import (
	"fmt"
	"strings"

	"github.com/hiroshiasayadev-prog/brewprint/internal/semantic"
)

type dbOperation struct {
	Step    int
	Task    string
	SubTask string
	Store   string
	Op      string
}

func RenderScenario(project *semantic.Project, scenarioID string) (string, error) {
	if project == nil {
		return "", fmt.Errorf("project is nil")
	}
	scenario := project.ScenariosByID[scenarioID]
	if scenario == nil {
		return "", fmt.Errorf("sequence scenario not found: %s", scenarioID)
	}
	return render(project, scenario), nil
}

func render(project *semantic.Project, scenario *semantic.SequenceScenario) string {
	messages := []string{}
	ops := []dbOperation{}
	hasUI := false
	hasActor := false
	actorID := ""
	actorLabel := ""
	hasAPI := false
	apiID := ""
	hasDB := false

	for i, step := range scenario.Steps {
		stepIndex := i + 1
		transition := step.Transition
		event := project.EventsByQID[transition.Event]
		task := project.TasksByQID[transition.ActionTask]
		source := ""
		if event != nil {
			source = event.Source
		}

		dbOps := collectDBOperations(project, stepIndex, task)
		if len(dbOps) > 0 {
			hasDB = true
			ops = append(ops, dbOps...)
		}

		switch source {
		case "ui":
			hasUI = true
			if task == nil {
				messages = append(messages, fmt.Sprintf("UI->>UI: %d. %s", stepIndex, transition.On))
				continue
			}
			hasAPI = true
			apiID = task.QID.String()
			messages = append(messages, fmt.Sprintf("UI->>API: %d. %s", stepIndex, requestLabel(task)))
			messages = append(messages, dbMessages(stepIndex, dbOps)...)
			messages = append(messages, fmt.Sprintf("API-->>UI: %d. %s", stepIndex, responseLabel(task)))
		case "external":
			hasActor = true
			if event != nil {
				actorID = event.Actor
				actorLabel = titleCase(event.Actor)
			}
			if task != nil {
				hasAPI = true
				apiID = task.QID.String()
				messages = append(messages, fmt.Sprintf("%s->>API: %d. %s", actorLabel, stepIndex, requestLabel(task)))
				messages = append(messages, dbMessages(stepIndex, dbOps)...)
				messages = append(messages, fmt.Sprintf("API-->>%s: %d. %s", actorLabel, stepIndex, responseLabel(task)))
			}
		case "er":
			if event != nil && event.Watches != "" {
				if store := project.StoresByQID[event.Watches]; store != nil && store.StoreKind == "db" {
					hasDB = true
					messages = append(messages, fmt.Sprintf("DB->>API: %d. %s", stepIndex, transition.On))
				} else {
					messages = append(messages, fmt.Sprintf("API->>API: %d. %s", stepIndex, transition.On))
				}
			} else {
				messages = append(messages, fmt.Sprintf("API->>API: %d. %s", stepIndex, transition.On))
			}
			if task != nil {
				hasAPI = true
				apiID = task.QID.String()
				messages = append(messages, dbMessages(stepIndex, dbOps)...)
			}
		case "internal":
			messages = append(messages, fmt.Sprintf("API->>API: %d. %s", stepIndex, transition.On))
			if task != nil {
				hasAPI = true
				apiID = task.QID.String()
				messages = append(messages, dbMessages(stepIndex, dbOps)...)
			}
		default:
			messages = append(messages, fmt.Sprintf("API->>API: %d. %s", stepIndex, transition.On))
		}
	}

	title := scenario.Title
	if title == "" {
		title = scenario.ID
	}

	var b strings.Builder
	b.WriteString("# " + title + "\n\n")
	b.WriteString("```mermaid\n")
	b.WriteString("sequenceDiagram\n")
	if hasActor {
		if actorLabel == "" {
			actorLabel = actorID
		}
		b.WriteString("  participant " + actorLabel + " as " + actorID + "\n")
	}
	if hasUI {
		b.WriteString("  participant UI\n")
	}
	if hasAPI {
		b.WriteString("  participant API as " + apiID + "\n")
	}
	if hasDB {
		b.WriteString("  participant DB\n")
	}
	b.WriteString("\n")
	for _, message := range messages {
		b.WriteString("  " + message + "\n")
	}
	b.WriteString("```\n\n")
	b.WriteString("## DB操作\n\n")
	b.WriteString("| step | task | sub_task | store | 操作 |\n")
	b.WriteString("|---|---|---|---|---|\n")
	for _, op := range ops {
		b.WriteString(fmt.Sprintf("| %d | %s | %s | %s | %s |\n", op.Step, op.Task, op.SubTask, op.Store, op.Op))
	}
	b.WriteString("\n")
	return b.String()
}

func requestLabel(task *semantic.Task) string {
	return strings.TrimSpace(task.Method + " /" + task.Path)
}

func responseLabel(task *semantic.Task) string {
	if task.Returns != nil && task.Returns.Name != "" {
		return task.Returns.Name
	}
	return "200 OK"
}

func dbMessages(stepIndex int, ops []dbOperation) []string {
	seen := map[string]struct{}{}
	for _, op := range ops {
		seen[op.Op] = struct{}{}
	}
	var out []string
	for _, op := range []string{"reads", "writes"} {
		if _, ok := seen[op]; ok {
			out = append(out, fmt.Sprintf("API->>DB: %d. %s", stepIndex, op))
		}
	}
	return out
}

func collectDBOperations(project *semantic.Project, stepIndex int, task *semantic.Task) []dbOperation {
	if project == nil || task == nil {
		return nil
	}
	var tasks []*semantic.Task
	for _, node := range project.NodesByFile[task.FileID] {
		if t, ok := node.(*semantic.Task); ok {
			tasks = append(tasks, t)
		}
	}
	var ops []dbOperation
	for _, t := range tasks {
		subTask := "-"
		if t.QID != task.QID {
			subTask = t.ID
		}
		taskOps := append(dbOpsForRefs(project, stepIndex, task.QID.String(), subTask, t.Reads, "reads"), dbOpsForRefs(project, stepIndex, task.QID.String(), subTask, t.Writes, "writes")...)
		ops = append(ops, orderDBOperations(taskOps)...)
	}
	return ops
}

func dbOpsForRefs(project *semantic.Project, stepIndex int, taskID, subTask string, refs []semantic.StoreRef, op string) []dbOperation {
	var out []dbOperation
	for _, ref := range refs {
		store := storeByQID(project, ref.Store)
		if store == nil || store.StoreKind != "db" {
			continue
		}
		out = append(out, dbOperation{
			Step:    stepIndex,
			Task:    taskID,
			SubTask: subTask,
			Store:   storeLabel(store, ref.Name),
			Op:      op,
		})
	}
	return out
}

func orderDBOperations(in []dbOperation) []dbOperation {
	if len(in) < 2 {
		return in
	}
	out := append([]dbOperation(nil), in...)
	for i := 0; i < len(out); i++ {
		for j := i + 1; j < len(out); j++ {
			if dbOperationLess(out[j], out[i]) {
				out[i], out[j] = out[j], out[i]
			}
		}
	}
	return out
}

func dbOperationLess(a, b dbOperation) bool {
	aLocal := !strings.Contains(a.Store, ".")
	bLocal := !strings.Contains(b.Store, ".")
	if aLocal != bLocal {
		return aLocal
	}
	if a.Store != b.Store {
		return a.Store < b.Store
	}
	if a.Op == b.Op {
		return false
	}
	return a.Op == "reads"
}

func storeByQID(project *semantic.Project, qid semantic.QualifiedID) *semantic.Store {
	if store := project.StoresByQID[qid]; store != nil {
		return store
	}
	for _, storesByName := range project.StoresByFileLocal {
		for _, store := range storesByName {
			if store.QID == qid {
				return store
			}
		}
	}
	return nil
}

func storeLabel(store *semantic.Store, fallback string) string {
	if fallback != "" {
		return fallback
	}
	return store.QID.String()
}

func titleCase(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}
