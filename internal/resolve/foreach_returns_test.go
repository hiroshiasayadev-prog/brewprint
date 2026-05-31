package resolve

import (
	"testing"

	"github.com/hiroshiasayadev-prog/brewprint/internal/rawyaml"
	"github.com/hiroshiasayadev-prog/brewprint/internal/semantic"
)

func TestForeachReturnsCollectedSourceOK(t *testing.T) {
	project, diagnostics := Build(foreachReturnsProject("cart_item", "list<cart_item>", []rawyaml.FlowEntry{
		{Foreach: "validate_item", Over: "$params.items", Params: map[string]string{"item": "$item"}, Returns: "validated_items"},
		{Step: "summarize", Params: map[string]string{"items": "validated_items"}},
	}, nil))
	if len(diagnostics) != 0 {
		t.Fatalf("diagnostics = %#v, want none", diagnostics)
	}
	collected := project.FlowCollectedSourcesByFile["shop/task/run.yaml"]["validated_items"]
	if collected == nil || collected.TypeRef == nil || collected.TypeRef.String() != "list<cart_item>" {
		t.Fatalf("collected TypeRef = %#v, want list<cart_item>", collected)
	}
}

func TestForeachReturnsAnyCollectedSourceOK(t *testing.T) {
	project, diagnostics := Build(foreachReturnsProject("any", "list<any>", []rawyaml.FlowEntry{
		{Foreach: "validate_item", Over: "$params.items", Params: map[string]string{"item": "$item"}, Returns: "validated_items"},
		{Step: "summarize", Params: map[string]string{"items": "validated_items"}},
	}, nil))
	assertNoErrorDiagnostics(t, diagnostics)
	assertDiagnostic(t, diagnostics, diagnosticOpaqueTypeRef, semantic.SeverityWarning)
	collected := project.FlowCollectedSourcesByFile["shop/task/run.yaml"]["validated_items"]
	if collected == nil || collected.TypeRef == nil || collected.TypeRef.String() != "list<any>" {
		t.Fatalf("collected TypeRef = %#v, want list<any>", collected)
	}
}

func TestForeachReturnsNamedListTargetOK(t *testing.T) {
	_, diagnostics := Build(foreachReturnsProject("cart_item", "cart_item_list", []rawyaml.FlowEntry{
		{Foreach: "validate_item", Over: "$params.items", Params: map[string]string{"item": "$item"}, Returns: "validated_items"},
		{Step: "summarize", Params: map[string]string{"items": "validated_items"}},
	}, []rawyaml.Model{{ID: "cart_item_list", Kind: "list", Element: "cart_item"}}))
	if len(diagnostics) != 0 {
		t.Fatalf("diagnostics = %#v, want none", diagnostics)
	}
}

func TestForeachReturnsCollectedSourceInControlParamsOK(t *testing.T) {
	_, branchDiagnostics := Build(foreachReturnsProject("cart_item", "list<cart_item>", []rawyaml.FlowEntry{
		{Foreach: "validate_item", Over: "$params.items", Params: map[string]string{"item": "$item"}, Returns: "validated_items"},
		{Branch: "route", Params: map[string]string{"items": "validated_items"}, Cases: []rawyaml.BranchCase{{Label: "ok", Step: "summarize", Params: map[string]string{"items": "validated_items"}}}},
	}, nil))
	if len(branchDiagnostics) != 0 {
		t.Fatalf("branch diagnostics = %#v, want none", branchDiagnostics)
	}

	_, forkDiagnostics := Build(foreachReturnsProject("cart_item", "list<cart_item>", []rawyaml.FlowEntry{
		{Foreach: "validate_item", Over: "$params.items", Params: map[string]string{"item": "$item"}, Returns: "validated_items"},
		{Fork: "fan_out", Branches: []rawyaml.ForkBranch{{Steps: []rawyaml.FlowStep{{Step: "summarize", Params: map[string]string{"items": "validated_items"}}}}}, Join: "aggregate"},
	}, nil))
	if len(forkDiagnostics) != 0 {
		t.Fatalf("fork diagnostics = %#v, want none", forkDiagnostics)
	}
}

func TestForeachReturnsCollectedSourceAsForeachOverOK(t *testing.T) {
	_, diagnostics := Build(foreachReturnsProject("cart_item", "list<cart_item>", []rawyaml.FlowEntry{
		{Foreach: "validate_item", Over: "$params.items", Params: map[string]string{"item": "$item"}, Returns: "validated_items"},
		{Foreach: "validate_item", Over: "validated_items", Params: map[string]string{"item": "$item"}, Returns: "validated_items_again"},
	}, nil))
	if len(diagnostics) != 0 {
		t.Fatalf("diagnostics = %#v, want none", diagnostics)
	}
}

func TestForeachReturnsDuplicateFlowSource(t *testing.T) {
	_, nodeDiagnostics := Build(foreachReturnsProjectWithTasks("cart_item", "list<cart_item>", []rawyaml.Task{{ID: "validated_items", Returns: &rawyaml.Return{Name: "result", Model: "cart_item"}}}, []rawyaml.FlowEntry{
		{Foreach: "validate_item", Over: "$params.items", Params: map[string]string{"item": "$item"}, Returns: "validated_items"},
	}, nil))
	assertDiagnosticCode(t, nodeDiagnostics, diagnosticDuplicateFlowSource)

	_, returnsDiagnostics := Build(foreachReturnsProject("cart_item", "list<cart_item>", []rawyaml.FlowEntry{
		{Foreach: "validate_item", Over: "$params.items", Params: map[string]string{"item": "$item"}, Returns: "validated_items"},
		{Foreach: "validate_item", Over: "$params.items", Params: map[string]string{"item": "$item"}, Returns: "validated_items"},
	}, nil))
	assertDiagnosticCode(t, returnsDiagnostics, diagnosticDuplicateFlowSource)
}

func TestForeachReturnsNameDoesNotConflictWithTaskReturnName(t *testing.T) {
	_, diagnostics := Build(foreachReturnsProject("cart_item", "list<cart_item>", []rawyaml.FlowEntry{
		{Foreach: "validate_item", Over: "$params.items", Params: map[string]string{"item": "$item"}, Returns: "validated_items"},
		{Step: "summarize", Params: map[string]string{"items": "validated_items"}},
	}, nil))
	assertNoDiagnosticCode(t, diagnostics, diagnosticDuplicateFlowSource)
}

func TestForeachReturnsInvalid(t *testing.T) {
	_, noReturnDiagnostics := Build(foreachReturnsProjectWithTasks("cart_item", "list<cart_item>", []rawyaml.Task{{ID: "validate_item", Params: []rawyaml.Param{{Name: "item", Model: "cart_item"}}}}, []rawyaml.FlowEntry{
		{Foreach: "validate_item", Over: "$params.items", Params: map[string]string{"item": "$item"}, Returns: "validated_items"},
	}, nil))
	assertDiagnosticCode(t, noReturnDiagnostics, diagnosticInvalidForeachReturns)

	_, selfRefDiagnostics := Build(foreachReturnsProject("cart_item", "list<cart_item>", []rawyaml.FlowEntry{
		{Foreach: "validate_item", Over: "$params.items", Params: map[string]string{"item": "$item", "prev": "validated_items"}, Returns: "validated_items"},
	}, nil))
	assertDiagnosticCode(t, selfRefDiagnostics, diagnosticInvalidForeachReturns)
	assertNoDiagnosticCode(t, selfRefDiagnostics, diagnosticUnresolvedWiringSource)
}

func TestInvalidForeachReturnsDoesNotRegisterCollectedSource(t *testing.T) {
	project, diagnostics := Build(foreachReturnsProjectWithTasks("cart_item", "list<cart_item>", []rawyaml.Task{{ID: "validate_item", Params: []rawyaml.Param{{Name: "item", Model: "cart_item"}}}}, []rawyaml.FlowEntry{
		{Foreach: "validate_item", Over: "$params.items", Params: map[string]string{"item": "$item"}, Returns: "validated_items"},
	}, nil))
	assertDiagnosticCode(t, diagnostics, diagnosticInvalidForeachReturns)
	if collected := project.FlowCollectedSourcesByFile["shop/task/run.yaml"]["validated_items"]; collected != nil {
		t.Fatalf("collected source registered for invalid foreach.returns: %#v", collected)
	}
}

func TestForeachReturnsOmittedDoesNotCreateCollectedSource(t *testing.T) {
	project, diagnostics := Build(foreachReturnsProject("cart_item", "list<cart_item>", []rawyaml.FlowEntry{
		{Foreach: "validate_item", Over: "$params.items", Params: map[string]string{"item": "$item"}},
	}, nil))
	if len(diagnostics) != 0 {
		t.Fatalf("diagnostics = %#v, want none", diagnostics)
	}
	if sources := project.FlowCollectedSourcesByFile["shop/task/run.yaml"]; len(sources) != 0 {
		t.Fatalf("collected sources = %#v, want none", sources)
	}
}

func TestForeachReturnsForwardAndTypoSourceUnresolved(t *testing.T) {
	_, forwardDiagnostics := Build(foreachReturnsProject("cart_item", "list<cart_item>", []rawyaml.FlowEntry{
		{Step: "summarize", Params: map[string]string{"items": "validated_items"}},
		{Foreach: "validate_item", Over: "$params.items", Params: map[string]string{"item": "$item"}, Returns: "validated_items"},
	}, nil))
	assertDiagnosticCode(t, forwardDiagnostics, diagnosticUnresolvedWiringSource)
	assertNoDiagnosticCode(t, forwardDiagnostics, diagnosticIncompatibleWiringType)

	_, typoDiagnostics := Build(foreachReturnsProject("cart_item", "list<cart_item>", []rawyaml.FlowEntry{
		{Step: "summarize", Params: map[string]string{"items": "typo_source"}},
	}, nil))
	assertDiagnosticCode(t, typoDiagnostics, diagnosticUnresolvedWiringSource)
	assertNoDiagnosticCode(t, typoDiagnostics, diagnosticIncompatibleWiringType)
}

func TestForeachReturnsInvalidSourceAndIncompatibleType(t *testing.T) {
	_, invalidDiagnostics := Build(foreachReturnsProject("cart_item", "list<cart_item>", []rawyaml.FlowEntry{{Step: "summarize", Params: map[string]string{"items": "route"}}}, nil))
	assertDiagnosticCode(t, invalidDiagnostics, diagnosticInvalidWiringSource)
	assertNoDiagnosticCode(t, invalidDiagnostics, diagnosticIncompatibleWiringType)

	_, incompatibleDiagnostics := Build(foreachReturnsProject("cart_item", "list<order>", []rawyaml.FlowEntry{
		{Foreach: "validate_item", Over: "$params.items", Params: map[string]string{"item": "$item"}, Returns: "validated_items"},
		{Step: "summarize", Params: map[string]string{"items": "validated_items"}},
	}, nil))
	assertDiagnosticCode(t, incompatibleDiagnostics, diagnosticIncompatibleWiringType)
}

func TestForeachReturnsUnresolvedReturnTypeSuppressesIncompatible(t *testing.T) {
	_, diagnostics := Build(foreachReturnsProject("missing_model", "list<order>", []rawyaml.FlowEntry{
		{Foreach: "validate_item", Over: "$params.items", Params: map[string]string{"item": "$item"}, Returns: "validated_items"},
		{Step: "summarize", Params: map[string]string{"items": "validated_items"}},
	}, nil))
	assertDiagnosticCode(t, diagnostics, diagnosticUnresolvedModel)
	assertNoDiagnosticCode(t, diagnostics, diagnosticIncompatibleWiringType)
}

func foreachReturnsProject(applyReturnModel string, targetModel string, flow []rawyaml.FlowEntry, extraModels []rawyaml.Model) *rawyaml.Project {
	return foreachReturnsProjectWithTasks(applyReturnModel, targetModel, nil, flow, extraModels)
}

func foreachReturnsProjectWithTasks(applyReturnModel string, targetModel string, extraTasks []rawyaml.Task, flow []rawyaml.FlowEntry, extraModels []rawyaml.Model) *rawyaml.Project {
	models := []rawyaml.Model{
		{ID: "cart_item", Kind: "struct", Fields: []rawyaml.ModelField{{Name: "id", Type: "str"}}},
		{ID: "order", Kind: "struct", Fields: []rawyaml.ModelField{{Name: "id", Type: "str"}}},
	}
	models = append(models, extraModels...)
	tasks := []rawyaml.Task{
		{ID: "run", Main: true, Params: []rawyaml.Param{{Name: "items", Model: "list<cart_item>"}}},
		{ID: "validate_item", Params: []rawyaml.Param{{Name: "item", Model: "cart_item"}, {Name: "prev", Model: "list<cart_item>"}}, Returns: &rawyaml.Return{Name: "validated_items", Model: applyReturnModel}},
		{ID: "summarize", Params: []rawyaml.Param{{Name: "items", Model: targetModel}}, Returns: &rawyaml.Return{Name: "summary", Model: "order"}},
	}
	for _, extra := range extraTasks {
		for i, task := range tasks {
			if task.ID == extra.ID {
				tasks[i] = extra
				extra.ID = ""
				break
			}
		}
		if extra.ID != "" {
			tasks = append(tasks, extra)
		}
	}
	return &rawyaml.Project{Files: []rawyaml.File{{
		ID:   "shop/task/run.yaml",
		Kind: rawyaml.FileKindNode,
		NodeFile: &rawyaml.NodeFile{
			Models:   models,
			Tasks:    tasks,
			Branches: []rawyaml.ControlNode{{ID: "route", Params: []rawyaml.Param{{Name: "items", Model: targetModel}}}},
			Forks:    []rawyaml.ControlNode{{ID: "fan_out"}},
			Joins:    []rawyaml.ControlNode{{ID: "aggregate", Returns: &rawyaml.Return{Name: "joined", Model: "order"}}},
			Flow:     flow,
		},
	}}}
}

func assertDiagnosticCodeCount(t *testing.T, diagnostics []semantic.Diagnostic, code string, want int) {
	t.Helper()
	if got := countDiagnosticCode(diagnostics, code); got != want {
		t.Fatalf("diagnostic code %s count = %d, want %d: %#v", code, got, want, diagnostics)
	}
}

func assertNoErrorDiagnostics(t *testing.T, diagnostics []semantic.Diagnostic) {
	t.Helper()
	for _, diagnostic := range diagnostics {
		if diagnostic.Severity == semantic.SeverityError {
			t.Fatalf("got error diagnostic: %#v in %#v", diagnostic, diagnostics)
		}
	}
}
