package resolve

import (
	"testing"

	"github.com/hiroshiasayadev-prog/brewprint/internal/rawyaml"
	"github.com/hiroshiasayadev-prog/brewprint/internal/semantic"
)

func TestTaskReturnSourceSemanticSourcePreserved(t *testing.T) {
	project, diagnostics := Build(returnSourceProject([]rawyaml.Task{
		{ID: "run", Main: true, Returns: &rawyaml.Return{Name: "result", Model: "user", Source: "produce"}},
		{ID: "produce", Returns: &rawyaml.Return{Name: "user", Model: "user"}},
	}, nil, nil))
	if len(diagnostics) != 0 {
		t.Fatalf("diagnostics = %#v, want none", diagnostics)
	}
	run := project.TasksByQID["shop.task.run"]
	if run == nil || run.Returns == nil || run.Returns.Source != "produce" {
		t.Fatalf("semantic return source = %#v, want produce", run)
	}
	if run.Returns.SourceRef.Kind != semantic.FlowSourceNode || run.Returns.SourceRef.Node != "shop/task/run.yaml#produce" {
		t.Fatalf("semantic return SourceRef = %#v, want same-file private produce node", run.Returns.SourceRef)
	}
}

func TestTaskReturnSourceOK(t *testing.T) {
	cases := []struct {
		name  string
		tasks []rawyaml.Task
		flow  []rawyaml.FlowEntry
	}{
		{
			name: "node id",
			tasks: []rawyaml.Task{
				{ID: "run", Main: true, Returns: &rawyaml.Return{Name: "result", Model: "user", Source: "produce"}},
				{ID: "produce", Returns: &rawyaml.Return{Name: "user", Model: "user"}},
			},
		},
		{
			name: "initialized source",
			tasks: []rawyaml.Task{
				{ID: "run", Main: true, Initializes: []rawyaml.Initialize{{Name: "cached_user", Model: "user"}}, Returns: &rawyaml.Return{Name: "result", Model: "user", Source: "cached_user"}},
			},
		},
		{
			name: "$params source",
			tasks: []rawyaml.Task{
				{ID: "run", Main: true, Params: []rawyaml.Param{{Name: "input", Model: "user"}}, Returns: &rawyaml.Return{Name: "result", Model: "user", Source: "$params.input"}},
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, diagnostics := Build(returnSourceProject(tc.tasks, tc.flow, nil))
			if len(diagnostics) != 0 {
				t.Fatalf("diagnostics = %#v, want none", diagnostics)
			}
		})
	}
}

func TestTaskReturnSourceCollectedAssetOKWithForwardVisibility(t *testing.T) {
	project, diagnostics := Build(returnSourceProject([]rawyaml.Task{
		{ID: "run", Main: true, Params: []rawyaml.Param{{Name: "items", Model: "list<cart_item>"}}, Returns: &rawyaml.Return{Name: "items", Model: "list<cart_item>", Source: "validated_items"}},
		{ID: "noop", Params: []rawyaml.Param{{Name: "items", Model: "list<cart_item>"}}, Returns: &rawyaml.Return{Name: "items", Model: "list<cart_item>"}},
		{ID: "validate_item", Params: []rawyaml.Param{{Name: "item", Model: "cart_item"}}, Returns: &rawyaml.Return{Name: "item", Model: "cart_item"}},
	}, []rawyaml.FlowEntry{
		{Step: "noop", Params: map[string]string{"items": "$params.items"}},
		{Foreach: "validate_item", Over: "$params.items", Params: map[string]string{"item": "$item"}, Returns: "validated_items"},
	}, nil))
	if len(diagnostics) != 0 {
		t.Fatalf("diagnostics = %#v, want none", diagnostics)
	}
	run := project.TasksByQID["shop.task.run"]
	if run == nil || run.Returns == nil || run.Returns.SourceRef.Raw != "validated_items" || run.Returns.SourceRef.TypeRef == nil {
		t.Fatalf("return SourceRef = %#v, want collected source", run.Returns)
	}
}

func TestTaskReturnSourceInvalidAndUnresolved(t *testing.T) {
	cases := []struct {
		name string
		task rawyaml.Task
		code string
	}{
		{name: "$item", task: rawyaml.Task{ID: "run", Main: true, Returns: &rawyaml.Return{Name: "result", Model: "user", Source: "$item"}}, code: diagnosticInvalidReturnSource},
		{name: "typo", task: rawyaml.Task{ID: "run", Main: true, Returns: &rawyaml.Return{Name: "result", Model: "user", Source: "typo_source"}}, code: diagnosticUnresolvedReturnSource},
		{name: "node without returns", task: rawyaml.Task{ID: "run", Main: true, Returns: &rawyaml.Return{Name: "result", Model: "user", Source: "consume"}}, code: diagnosticInvalidReturnSource},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, diagnostics := Build(returnSourceProject([]rawyaml.Task{tc.task, rawyaml.Task{ID: "consume", Params: []rawyaml.Param{{Name: "input", Model: "user"}}}}, nil, nil))
			assertDiagnosticCode(t, diagnostics, tc.code)
			assertNoDiagnosticCode(t, diagnostics, diagnosticIncompatibleReturnType)
		})
	}
}

func TestTaskReturnSourceIncompatibleType(t *testing.T) {
	_, diagnostics := Build(returnSourceProject([]rawyaml.Task{
		{ID: "run", Main: true, Returns: &rawyaml.Return{Name: "result", Model: "order", Source: "produce"}},
		{ID: "produce", Returns: &rawyaml.Return{Name: "user", Model: "user"}},
	}, nil, nil))
	assertDiagnosticCode(t, diagnostics, diagnosticIncompatibleReturnType)
}

func TestTaskReturnSourceTypeSuppression(t *testing.T) {
	_, sourceDiagnostics := Build(returnSourceProject([]rawyaml.Task{
		{ID: "run", Main: true, Returns: &rawyaml.Return{Name: "result", Model: "order", Source: "produce"}},
		{ID: "produce", Returns: &rawyaml.Return{Name: "missing", Model: "missing_model"}},
	}, nil, nil))
	assertDiagnosticCode(t, sourceDiagnostics, diagnosticUnresolvedModel)
	assertNoDiagnosticCode(t, sourceDiagnostics, diagnosticIncompatibleReturnType)

	_, targetDiagnostics := Build(returnSourceProject([]rawyaml.Task{
		{ID: "run", Main: true, Returns: &rawyaml.Return{Name: "result", Model: "missing_model", Source: "produce"}},
		{ID: "produce", Returns: &rawyaml.Return{Name: "user", Model: "user"}},
	}, nil, nil))
	assertDiagnosticCode(t, targetDiagnostics, diagnosticUnresolvedModel)
	assertNoDiagnosticCode(t, targetDiagnostics, diagnosticIncompatibleReturnType)
}

func TestInitializedSourceFlowWiringOK(t *testing.T) {
	flow := []rawyaml.FlowEntry{
		{Step: "consume", Params: map[string]string{"input": "cached_user"}},
		{Branch: "route", Params: map[string]string{"input": "cached_user"}, Cases: []rawyaml.BranchCase{{Label: "ok", Step: "consume", Params: map[string]string{"input": "cached_user"}}}},
		{Fork: "fan_out", Branches: []rawyaml.ForkBranch{{Steps: []rawyaml.FlowStep{{Step: "left", Params: map[string]string{"input": "cached_user"}}}}}, Join: "aggregate"},
		{Foreach: "consume", Over: "cached_users", Params: map[string]string{"input": "$item"}},
	}
	_, diagnostics := Build(returnSourceProject([]rawyaml.Task{
		{ID: "run", Main: true, Initializes: []rawyaml.Initialize{{Name: "cached_user", Model: "user"}, {Name: "cached_users", Model: "user_list"}}},
		{ID: "consume", Params: []rawyaml.Param{{Name: "input", Model: "user"}}, Returns: &rawyaml.Return{Name: "result", Model: "user"}},
		{ID: "left", Params: []rawyaml.Param{{Name: "input", Model: "user"}}, Returns: &rawyaml.Return{Name: "left_result", Model: "user"}},
	}, flow, []rawyaml.Model{{ID: "user_list", Kind: "list", Element: "user"}}))
	if len(diagnostics) != 0 {
		t.Fatalf("diagnostics = %#v, want none", diagnostics)
	}
}

func TestInitializedSourceDuplicateFlowSource(t *testing.T) {
	_, nodeDiagnostics := Build(returnSourceProject([]rawyaml.Task{
		{ID: "run", Main: true, Initializes: []rawyaml.Initialize{{Name: "consume", Model: "user"}}},
		{ID: "consume", Returns: &rawyaml.Return{Name: "result", Model: "user"}},
	}, []rawyaml.FlowEntry{{Step: "consume"}}, nil))
	assertDiagnosticCode(t, nodeDiagnostics, diagnosticDuplicateFlowSource)

	_, foreachDiagnostics := Build(returnSourceProject([]rawyaml.Task{
		{ID: "run", Main: true, Params: []rawyaml.Param{{Name: "users", Model: "list<user>"}}, Initializes: []rawyaml.Initialize{{Name: "collected_users", Model: "user_list"}}},
		{ID: "consume", Params: []rawyaml.Param{{Name: "input", Model: "user"}}, Returns: &rawyaml.Return{Name: "result", Model: "user"}},
	}, []rawyaml.FlowEntry{{Foreach: "consume", Over: "$params.users", Params: map[string]string{"input": "$item"}, Returns: "collected_users"}}, []rawyaml.Model{{ID: "user_list", Kind: "list", Element: "user"}}))
	assertDiagnosticCode(t, foreachDiagnostics, diagnosticDuplicateFlowSource)
}

func TestTaskReturnNameAndInitializedSourceNameDoNotConflict(t *testing.T) {
	_, diagnostics := Build(returnSourceProject([]rawyaml.Task{
		{ID: "run", Main: true, Initializes: []rawyaml.Initialize{{Name: "result", Model: "user"}}, Returns: &rawyaml.Return{Name: "result", Model: "user", Source: "result"}},
	}, nil, nil))
	assertNoDiagnosticCode(t, diagnostics, diagnosticDuplicateFlowSource)
}

func TestInitializedSourceUnwrittenReferenceIsValid(t *testing.T) {
	_, diagnostics := Build(returnSourceProject([]rawyaml.Task{
		{ID: "run", Main: true, Initializes: []rawyaml.Initialize{{Name: "cached_user", Model: "user"}}},
		{ID: "consume", Params: []rawyaml.Param{{Name: "input", Model: "user"}}, Returns: &rawyaml.Return{Name: "result", Model: "user"}},
	}, []rawyaml.FlowEntry{{Step: "consume", Params: map[string]string{"input": "cached_user"}}}, nil))
	if len(diagnostics) != 0 {
		t.Fatalf("diagnostics = %#v, want none", diagnostics)
	}
}

func TestInitializedSourceWiringTypeCompatibilityAndSuppression(t *testing.T) {
	_, incompatibleDiagnostics := Build(returnSourceProject([]rawyaml.Task{
		{ID: "run", Main: true, Initializes: []rawyaml.Initialize{{Name: "cached_user", Model: "user"}}},
		{ID: "consume", Params: []rawyaml.Param{{Name: "input", Model: "order"}}, Returns: &rawyaml.Return{Name: "result", Model: "order"}},
	}, []rawyaml.FlowEntry{{Step: "consume", Params: map[string]string{"input": "cached_user"}}}, nil))
	assertDiagnosticCode(t, incompatibleDiagnostics, diagnosticIncompatibleWiringType)

	_, unresolvedDiagnostics := Build(returnSourceProject([]rawyaml.Task{
		{ID: "run", Main: true, Initializes: []rawyaml.Initialize{{Name: "cached_user", Model: "missing_model"}}},
		{ID: "consume", Params: []rawyaml.Param{{Name: "input", Model: "order"}}, Returns: &rawyaml.Return{Name: "result", Model: "order"}},
	}, []rawyaml.FlowEntry{{Step: "consume", Params: map[string]string{"input": "cached_user"}}}, nil))
	assertDiagnosticCode(t, unresolvedDiagnostics, diagnosticUnresolvedModel)
	assertNoDiagnosticCode(t, unresolvedDiagnostics, diagnosticIncompatibleWiringType)
}

func TestInitializedModelDoesNotAcceptInlineContainer(t *testing.T) {
	_, diagnostics := Build(returnSourceProject([]rawyaml.Task{
		{ID: "run", Main: true, Initializes: []rawyaml.Initialize{{Name: "cached_users", Model: "list<user>"}}},
	}, nil, nil))
	assertDiagnosticCode(t, diagnostics, diagnosticInvalidTypeRef)
}

func TestJoinParamImplicitReturnNameStillWorks(t *testing.T) {
	_, diagnostics := Build(returnSourceProject([]rawyaml.Task{
		{ID: "run", Main: true},
		{ID: "left", Returns: &rawyaml.Return{Name: "left_result", Model: "user"}},
	}, []rawyaml.FlowEntry{{Fork: "fan_out", Branches: []rawyaml.ForkBranch{{Steps: []rawyaml.FlowStep{{Step: "left"}}}}, Join: "aggregate"}}, nil))
	if len(diagnostics) != 0 {
		t.Fatalf("diagnostics = %#v, want none", diagnostics)
	}
}

func returnSourceProject(tasks []rawyaml.Task, flow []rawyaml.FlowEntry, extraModels []rawyaml.Model) *rawyaml.Project {
	models := []rawyaml.Model{
		{ID: "user", Kind: "struct", Fields: []rawyaml.ModelField{{Name: "id", Type: "str"}}},
		{ID: "order", Kind: "struct", Fields: []rawyaml.ModelField{{Name: "id", Type: "str"}}},
		{ID: "cart_item", Kind: "struct", Fields: []rawyaml.ModelField{{Name: "id", Type: "str"}}},
	}
	models = append(models, extraModels...)
	return &rawyaml.Project{Files: []rawyaml.File{{
		ID:   "shop/task/run.yaml",
		Kind: rawyaml.FileKindNode,
		NodeFile: &rawyaml.NodeFile{
			Models:   models,
			Tasks:    tasks,
			Branches: []rawyaml.ControlNode{{ID: "route", Params: []rawyaml.Param{{Name: "input", Model: "user"}}}},
			Forks:    []rawyaml.ControlNode{{ID: "fan_out"}},
			Joins:    []rawyaml.ControlNode{{ID: "aggregate", Params: []rawyaml.Param{{Name: "left_result", Model: "user"}}, Returns: &rawyaml.Return{Name: "joined", Model: "user"}}},
			Flow:     flow,
		},
	}}}
}
