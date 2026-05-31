package resolve

import (
	"testing"

	"github.com/hiroshiasayadev-prog/brewprint/internal/rawyaml"
	"github.com/hiroshiasayadev-prog/brewprint/internal/semantic"
)

func TestFlowWiringTypesStepOK(t *testing.T) {
	cases := []struct {
		name        string
		sourceModel string
		targetModel string
		extraModels []rawyaml.Model
	}{
		{name: "user to user", sourceModel: "user", targetModel: "user"},
		{name: "any to user", sourceModel: "any", targetModel: "user"},
		{name: "list any to list user", sourceModel: "list<any>", targetModel: "list<user>"},
		{name: "named list to list user", sourceModel: "user_list", targetModel: "list<user>", extraModels: []rawyaml.Model{{ID: "user_list", Kind: "list", Element: "user"}}},
		{name: "named dict to dict config", sourceModel: "config_map", targetModel: "dict<config>", extraModels: []rawyaml.Model{{ID: "config_map", Kind: "dict", Value: "config"}}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, diagnostics := Build(flowWiringProject(tc.sourceModel, tc.targetModel, tc.extraModels, []rawyaml.FlowEntry{{Step: "consume", Params: map[string]string{"input": "produce"}}}))
			assertNoDiagnosticCode(t, diagnostics, diagnosticIncompatibleWiringType)
		})
	}
}

func TestFlowWiringTypesControlOK(t *testing.T) {
	_, branchDiagnostics := Build(flowWiringProject("user", "user", nil, []rawyaml.FlowEntry{{
		Branch: "route",
		Params: map[string]string{"input": "produce"},
		Cases:  []rawyaml.BranchCase{{Label: "ok", Step: "consume", Params: map[string]string{"input": "produce"}}},
	}}))
	if len(branchDiagnostics) != 0 {
		t.Fatalf("branch diagnostics = %#v, want none", branchDiagnostics)
	}

	_, forkDiagnostics := Build(flowWiringProject("user", "user", nil, []rawyaml.FlowEntry{{
		Fork:     "fan_out",
		Branches: []rawyaml.ForkBranch{{Steps: []rawyaml.FlowStep{{Step: "consume", Params: map[string]string{"input": "produce"}}}}},
		Join:     "aggregate",
	}}))
	if len(forkDiagnostics) != 0 {
		t.Fatalf("fork diagnostics = %#v, want none", forkDiagnostics)
	}
}

func TestFlowWiringTypesForeachOK(t *testing.T) {
	cases := []struct {
		name      string
		overModel string
		extra     []rawyaml.Model
	}{
		{name: "list user", overModel: "list<user>"},
		{name: "named list user", overModel: "user_list", extra: []rawyaml.Model{{ID: "user_list", Kind: "list", Element: "user"}}},
		{name: "any", overModel: "any"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, diagnostics := Build(flowWiringProject(tc.overModel, "user", tc.extra, []rawyaml.FlowEntry{{
				Foreach: "consume",
				Over:    "produce",
				Params:  map[string]string{"input": "$item"},
			}}))
			if len(diagnostics) != 0 {
				t.Fatalf("diagnostics = %#v, want none", diagnostics)
			}
		})
	}
}

func TestFlowWiringTypesParamSource(t *testing.T) {
	_, diagnostics := Build(flowWiringProject("user", "user", nil, []rawyaml.FlowEntry{{Step: "consume", Params: map[string]string{"input": "$params.input"}}}))
	if len(diagnostics) != 0 {
		t.Fatalf("$params source diagnostics = %#v, want none", diagnostics)
	}

	_, missingParamDiagnostics := Build(flowWiringProject("user", "user", nil, []rawyaml.FlowEntry{{Step: "consume", Params: map[string]string{"input": "$params.missing"}}}))
	assertDiagnosticCode(t, missingParamDiagnostics, diagnosticUnresolvedWiringSource)
	assertNoDiagnosticCode(t, missingParamDiagnostics, diagnosticIncompatibleWiringType)
}

func TestFlowWiringTypesNG(t *testing.T) {
	cases := []struct {
		name        string
		sourceModel string
		targetModel string
		extraModels []rawyaml.Model
	}{
		{name: "user to order", sourceModel: "user", targetModel: "order"},
		{name: "str to int", sourceModel: "str", targetModel: "int"},
		{name: "list user to list order", sourceModel: "list<user>", targetModel: "list<order>"},
		{name: "dict config to dict user", sourceModel: "dict<config>", targetModel: "dict<user>"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, diagnostics := Build(flowWiringProject(tc.sourceModel, tc.targetModel, tc.extraModels, []rawyaml.FlowEntry{{Step: "consume", Params: map[string]string{"input": "produce"}}}))
			assertDiagnosticCode(t, diagnostics, diagnosticIncompatibleWiringType)
		})
	}
}

func TestFlowWiringTypesInvalidSources(t *testing.T) {
	_, branchSourceDiagnostics := Build(flowWiringProject("user", "user", nil, []rawyaml.FlowEntry{{Step: "consume", Params: map[string]string{"input": "route"}}}))
	assertDiagnosticCode(t, branchSourceDiagnostics, diagnosticInvalidWiringSource)
	assertNoDiagnosticCode(t, branchSourceDiagnostics, diagnosticIncompatibleWiringType)

	_, itemOutsideDiagnostics := Build(flowWiringProject("user", "user", nil, []rawyaml.FlowEntry{{Step: "consume", Params: map[string]string{"input": "$item"}}}))
	assertDiagnosticCode(t, itemOutsideDiagnostics, diagnosticInvalidWiringSource)
	assertNoDiagnosticCode(t, itemOutsideDiagnostics, diagnosticIncompatibleWiringType)
}

func TestFlowWiringTypesForeachOverInvalidAndIncompatible(t *testing.T) {
	_, invalidOverDiagnostics := Build(flowWiringProject("user", "user", nil, []rawyaml.FlowEntry{{
		Foreach: "consume",
		Over:    "produce",
		Params:  map[string]string{"input": "$item"},
	}}))
	assertDiagnosticCode(t, invalidOverDiagnostics, diagnosticInvalidForeachOverType)
	assertNoDiagnosticCode(t, invalidOverDiagnostics, diagnosticIncompatibleWiringType)

	_, incompatibleItemDiagnostics := Build(flowWiringProject("list<order>", "user", nil, []rawyaml.FlowEntry{{
		Foreach: "consume",
		Over:    "produce",
		Params:  map[string]string{"input": "$item"},
	}}))
	assertDiagnosticCode(t, incompatibleItemDiagnostics, diagnosticIncompatibleWiringType)
}

func TestFlowWiringTypeSuppression(t *testing.T) {
	_, unresolvedSourceDiagnostics := Build(flowWiringProject("user", "order", nil, []rawyaml.FlowEntry{{Step: "consume", Params: map[string]string{"input": "missing_source"}}}))
	assertDiagnosticCode(t, unresolvedSourceDiagnostics, diagnosticUnresolvedWiringSource)
	assertNoDiagnosticCode(t, unresolvedSourceDiagnostics, diagnosticIncompatibleWiringType)

	_, unresolvedTargetDiagnostics := Build(flowWiringProject("user", "missing_model", nil, []rawyaml.FlowEntry{{Step: "consume", Params: map[string]string{"input": "produce"}}}))
	assertDiagnosticCode(t, unresolvedTargetDiagnostics, diagnosticUnresolvedModel)
	assertNoDiagnosticCode(t, unresolvedTargetDiagnostics, diagnosticIncompatibleWiringType)

	_, invalidOverDiagnostics := Build(flowWiringProject("user", "order", nil, []rawyaml.FlowEntry{{
		Foreach: "consume",
		Over:    "produce",
		Params:  map[string]string{"input": "$item"},
	}}))
	assertDiagnosticCode(t, invalidOverDiagnostics, diagnosticInvalidForeachOverType)
	assertNoDiagnosticCode(t, invalidOverDiagnostics, diagnosticIncompatibleWiringType)
}

func flowWiringProject(sourceModel string, targetModel string, extraModels []rawyaml.Model, flow []rawyaml.FlowEntry) *rawyaml.Project {
	models := []rawyaml.Model{
		{ID: "user", Kind: "struct", Fields: []rawyaml.ModelField{{Name: "id", Type: "str"}}},
		{ID: "order", Kind: "struct", Fields: []rawyaml.ModelField{{Name: "id", Type: "str"}}},
		{ID: "config", Kind: "struct", Fields: []rawyaml.ModelField{{Name: "id", Type: "str"}}},
	}
	models = append(models, extraModels...)
	return &rawyaml.Project{Files: []rawyaml.File{
		{
			ID:   "shop/model/types.yaml",
			Kind: rawyaml.FileKindNode,
			NodeFile: &rawyaml.NodeFile{
				Models: models,
			},
		},
		{
			ID:   "shop/task/run.yaml",
			Kind: rawyaml.FileKindNode,
			NodeFile: &rawyaml.NodeFile{
				Tasks: []rawyaml.Task{
					{ID: "run", Main: true, Params: []rawyaml.Param{{Name: "input", Model: sourceModel}}},
					{ID: "produce", Returns: &rawyaml.Return{Name: "input", Model: sourceModel}},
					{ID: "consume", Params: []rawyaml.Param{{Name: "input", Model: targetModel}}, Returns: &rawyaml.Return{Name: "input", Model: targetModel}},
				},
				Branches: []rawyaml.ControlNode{{ID: "route", Params: []rawyaml.Param{{Name: "input", Model: targetModel}}}},
				Forks:    []rawyaml.ControlNode{{ID: "fan_out"}},
				Joins:    []rawyaml.ControlNode{{ID: "aggregate", Params: []rawyaml.Param{{Name: "input", Model: sourceModel}}, Returns: &rawyaml.Return{Name: "input", Model: sourceModel}}},
				Flow:     flow,
			},
		},
	}}
}

func assertNoDiagnosticCode(t *testing.T, diagnostics []semantic.Diagnostic, code string) {
	t.Helper()
	for _, diagnostic := range diagnostics {
		if diagnostic.Code == code {
			t.Fatalf("diagnostic code %s found unexpectedly in %#v", code, diagnostics)
		}
	}
}
