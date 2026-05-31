package resolve

import (
	"testing"

	"github.com/hiroshiasayadev-prog/brewprint/internal/rawyaml"
	"github.com/hiroshiasayadev-prog/brewprint/internal/semantic"
)

func TestBuildValidationDiagnostics(t *testing.T) {
	_, diagnostics := Build(&rawyaml.Project{Files: []rawyaml.File{
		{
			ID:   "auth/model/user.yaml",
			Kind: rawyaml.FileKindNode,
			NodeFile: &rawyaml.NodeFile{Models: []rawyaml.Model{
				{
					ID:   "str",
					Kind: "struct",
					Fields: []rawyaml.ModelField{
						{Name: "id", Type: "str", PK: true},
						{Name: "id", Type: "missing_field_type", PK: true},
						{Name: "role_id", Type: "str", FK: "role.missing_id"},
					},
				},
			}},
		},
		{
			ID:   "auth/store/user_db.yaml",
			Kind: rawyaml.FileKindNode,
			NodeFile: &rawyaml.NodeFile{Stores: []rawyaml.Store{
				{ID: "user_db", Kind: "cache", Of: "missing_store_model"},
			}},
		},
		{
			ID:   "auth/task/login.yaml",
			Kind: rawyaml.FileKindNode,
			NodeFile: &rawyaml.NodeFile{Tasks: []rawyaml.Task{
				{
					ID:       "login",
					Main:     true,
					Endpoint: true,
					Method:   "FETCH",
					Path:     "auth/login",
					Params: []rawyaml.Param{
						{Name: "request", Model: "missing_param_model"},
					},
					Returns: &rawyaml.Return{Name: "result", Model: "missing_return_model"},
					Reads:   []string{"missing_read_store"},
					Writes:  []string{"missing_write_store"},
					Initializes: []rawyaml.Initialize{
						{Name: "tmp", Model: "missing_init_model"},
					},
				},
			}},
		},
	}})

	assertDiagnosticCode(t, diagnostics, diagnosticInvalidModelID)
	assertDiagnosticCode(t, diagnostics, diagnosticUnresolvedFieldType)
	assertDiagnosticCode(t, diagnostics, diagnosticUnresolvedFK)
	assertDiagnosticCode(t, diagnostics, diagnosticDuplicateModelField)
	assertDiagnosticCode(t, diagnostics, diagnosticDuplicatePrimaryKey)
	assertDiagnosticCode(t, diagnostics, diagnosticInvalidStoreKind)
	assertDiagnosticCode(t, diagnostics, diagnosticUnresolvedModel)
	assertDiagnosticCode(t, diagnostics, diagnosticUnresolvedStore)
	assertDiagnosticCode(t, diagnostics, diagnosticInvalidEndpoint)
}

func TestBuildEnumModelDiagnostics(t *testing.T) {
	_, diagnostics := Build(&rawyaml.Project{Files: []rawyaml.File{
		{
			ID:   "shop/model/status.yaml",
			Kind: rawyaml.FileKindNode,
			NodeFile: &rawyaml.NodeFile{Models: []rawyaml.Model{
				{ID: "missing_values", Kind: "enum"},
				{ID: "empty_value", Kind: "enum", Values: []string{"active", ""}},
				{ID: "duplicate_value", Kind: "enum", Values: []string{"active", "active"}},
				{
					ID:      "wrong_shape",
					Kind:    "enum",
					Values:  []string{"active"},
					Fields:  []rawyaml.ModelField{{Name: "id", Type: "str"}},
					Element: "str",
					Value:   "str",
				},
			}},
		},
	}})

	assertDiagnosticCode(t, diagnostics, diagnosticInvalidEnumModel)
	assertDiagnosticCode(t, diagnostics, diagnosticDuplicateEnumValue)
}

func TestBuildValidEnumModel(t *testing.T) {
	project, diagnostics := Build(&rawyaml.Project{Files: []rawyaml.File{
		{
			ID:   "shop/model/status.yaml",
			Kind: rawyaml.FileKindNode,
			NodeFile: &rawyaml.NodeFile{Models: []rawyaml.Model{{
				ID:     "status",
				Kind:   "enum",
				Values: []string{"active", "inactive"},
			}}},
		},
		{
			ID:   "shop/model/user.yaml",
			Kind: rawyaml.FileKindNode,
			NodeFile: &rawyaml.NodeFile{Models: []rawyaml.Model{{
				ID:   "user",
				Kind: "struct",
				Fields: []rawyaml.ModelField{{
					Name: "status",
					Type: "status",
				}},
			}}},
		},
	}})
	if len(diagnostics) != 0 {
		t.Fatalf("diagnostics = %#v, want none", diagnostics)
	}
	model := project.ModelsByQID["shop.model.status"]
	if model == nil || model.Kind != "enum" || len(model.Values) != 2 || model.Values[0] != "active" {
		t.Fatalf("enum model not built as expected: %#v", model)
	}
}

func TestEnumTypeRefsUseNominalCompatibility(t *testing.T) {
	project, diagnostics := Build(&rawyaml.Project{Files: []rawyaml.File{
		{
			ID:   "shop/model/status.yaml",
			Kind: rawyaml.FileKindNode,
			NodeFile: &rawyaml.NodeFile{Models: []rawyaml.Model{
				{ID: "status", Kind: "enum", Values: []string{"active", "inactive"}},
				{ID: "severity", Kind: "enum", Values: []string{"info", "warning"}},
			}},
		},
	}})
	if len(diagnostics) != 0 {
		t.Fatalf("diagnostics = %#v, want none", diagnostics)
	}
	status := mustParseTypeRefForTest(t, "status")
	severity := mustParseTypeRefForTest(t, "severity")
	str := mustParseTypeRefForTest(t, "str")
	any := mustParseTypeRefForTest(t, "any")
	if !typeRefsCompatible(project, status, mustParseTypeRefForTest(t, "status")) {
		t.Fatalf("same enum type should be compatible")
	}
	if typeRefsCompatible(project, status, severity) {
		t.Fatalf("different enum models should not be compatible")
	}
	if typeRefsCompatible(project, status, str) || typeRefsCompatible(project, str, status) {
		t.Fatalf("enum and str should not be implicitly compatible")
	}
	if !typeRefsCompatible(project, any, status) || !typeRefsCompatible(project, status, any) {
		t.Fatalf("any should remain compatible with enum")
	}
}

func TestBuildMissingRequiredFieldDiagnostics(t *testing.T) {
	_, diagnostics := Build(&rawyaml.Project{Files: []rawyaml.File{
		{
			ID:   "auth/model/broken.yaml",
			Kind: rawyaml.FileKindNode,
			NodeFile: &rawyaml.NodeFile{Models: []rawyaml.Model{
				{
					Fields: []rawyaml.ModelField{{}},
				},
			}},
		},
		{
			ID:   "auth/store/session.yaml",
			Kind: rawyaml.FileKindNode,
			NodeFile: &rawyaml.NodeFile{Stores: []rawyaml.Store{
				{ID: "session_store"},
			}},
		},
		{
			ID:   "auth/task/login.yaml",
			Kind: rawyaml.FileKindNode,
			NodeFile: &rawyaml.NodeFile{Tasks: []rawyaml.Task{
				{
					Params:      []rawyaml.Param{{}},
					Returns:     &rawyaml.Return{},
					Initializes: []rawyaml.Initialize{{}},
				},
			}},
		},
	}})

	if count := countDiagnosticCode(diagnostics, diagnosticMissingRequiredField); count < 8 {
		t.Fatalf("missing_required_field count = %d, want at least 8: %#v", count, diagnostics)
	}
}

func TestBuildFlowDiagnosticCodes(t *testing.T) {
	_, diagnostics := Build(&rawyaml.Project{Files: []rawyaml.File{
		{
			ID:   "auth/task/login.yaml",
			Kind: rawyaml.FileKindNode,
			NodeFile: &rawyaml.NodeFile{Flow: []rawyaml.FlowEntry{
				{},
				{Step: "missing_task"},
			}},
		},
	}})

	assertDiagnostic(t, diagnostics, diagnosticUnsupportedFlowEntry, semantic.SeverityWarning)
	assertDiagnostic(t, diagnostics, diagnosticUnresolvedFlowTask, semantic.SeverityError)
}

func TestBuildTransitionDiagnosticCodes(t *testing.T) {
	_, diagnostics := Build(&rawyaml.Project{Files: []rawyaml.File{
		{
			ID:   "auth/state.yaml",
			Kind: rawyaml.FileKindNode,
			NodeFile: &rawyaml.NodeFile{
				States: []rawyaml.State{
					{ID: "idle"},
					{ID: "processing"},
				},
				Events: []rawyaml.Event{
					{ID: "submitted"},
				},
				Transitions: []rawyaml.Transition{
					{From: "missing_from", On: "missing_event", To: "missing_to"},
					{From: "idle", On: "submitted", To: "processing"},
					{From: "idle", On: "submitted", To: "processing"},
					{From: "idle", On: "submitted", To: "processing", Guard: "paid"},
				},
			},
		},
	}})

	assertDiagnosticCode(t, diagnostics, diagnosticUnresolvedTransitionState)
	assertDiagnosticCode(t, diagnostics, diagnosticUnresolvedTransitionEvent)
	assertDiagnosticCode(t, diagnostics, diagnosticDuplicateTransition)
	assertDiagnosticCode(t, diagnostics, diagnosticMissingTransitionGuard)
}

func TestBuildViewDiagnosticCodes(t *testing.T) {
	_, diagnostics := Build(&rawyaml.Project{Files: []rawyaml.File{
		{
			ID:   "views/api_missing.yaml",
			Kind: rawyaml.FileKindView,
			APIView: &rawyaml.APIView{
				ID:           "public_api",
				HTTPRootPath: "/api",
				Modules: []rawyaml.APIViewModule{
					{Module: "auth"},
					{Module: "auth"},
				},
			},
		},
		{
			ID:   "views/api_duplicate.yaml",
			Kind: rawyaml.FileKindView,
			APIView: &rawyaml.APIView{
				ID:           "public_api",
				HTTPRootPath: "/api",
				Modules:      []rawyaml.APIViewModule{{Module: "order"}},
			},
		},
		{
			ID:     "views/er_invalid.yaml",
			Kind:   rawyaml.FileKindView,
			ERView: &rawyaml.ERView{ID: "broken_er", Modules: []rawyaml.ERViewModule{{}}},
		},
	}})

	assertDiagnosticCode(t, diagnostics, diagnosticDuplicateView)
	assertDiagnosticCode(t, diagnostics, diagnosticDuplicateViewModule)
	assertDiagnosticCode(t, diagnostics, diagnosticInvalidViewDefinition)
}

func TestBuildScenarioDiagnosticCodes(t *testing.T) {
	_, diagnostics := Build(&rawyaml.Project{Files: []rawyaml.File{
		{
			ID:   "auth/state.yaml",
			Kind: rawyaml.FileKindNode,
			NodeFile: &rawyaml.NodeFile{
				States: []rawyaml.State{{ID: "idle"}, {ID: "done"}, {ID: "other"}},
				Events: []rawyaml.Event{{ID: "submitted"}, {ID: "finished"}},
				Transitions: []rawyaml.Transition{
					{From: "idle", On: "submitted", To: "done"},
					{From: "other", On: "finished", To: "done"},
				},
			},
		},
		{
			ID:   "views/scenario.yaml",
			Kind: rawyaml.FileKindView,
			SequenceScenario: &rawyaml.SequenceScenario{
				ID:        "checkout",
				StateFile: "auth/state.yaml",
				Steps: []rawyaml.SequenceStep{
					{FromState: "missing", Via: "missing_event"},
					{FromState: "idle", Via: "submitted"},
					{FromState: "other", Via: "finished"},
				},
			},
		},
	}})

	assertDiagnosticCode(t, diagnostics, diagnosticUnresolvedSequenceStep)
	assertDiagnosticCode(t, diagnostics, diagnosticNonContinuousSequence)
}

func TestSortedDiagnostics(t *testing.T) {
	diagnostics := []semantic.Diagnostic{
		{Severity: semantic.SeverityWarning, Code: "z_warning", FileID: "b.yaml", Message: "warning"},
		{Severity: semantic.SeverityError, Code: "z_error", FileID: "b.yaml", Message: "error b"},
		{Severity: semantic.SeverityError, Code: "a_error", FileID: "b.yaml", Message: "error b a"},
		{Severity: semantic.SeverityError, Code: "z_error", FileID: "a.yaml", Message: "error a"},
	}

	sorted := sortedDiagnostics(diagnostics)
	wantCodes := []string{"z_error", "a_error", "z_error", "z_warning"}
	wantFiles := []semantic.FileID{"a.yaml", "b.yaml", "b.yaml", "b.yaml"}
	for i := range wantCodes {
		if sorted[i].Code != wantCodes[i] || sorted[i].FileID != wantFiles[i] {
			t.Fatalf("sorted[%d] = %#v, want code=%s file=%s", i, sorted[i], wantCodes[i], wantFiles[i])
		}
	}
}

func assertDiagnosticCode(t *testing.T, diagnostics []semantic.Diagnostic, code string) {
	t.Helper()
	assertDiagnostic(t, diagnostics, code, semantic.SeverityError)
}

func assertDiagnostic(t *testing.T, diagnostics []semantic.Diagnostic, code string, severity semantic.Severity) {
	t.Helper()
	for _, diagnostic := range diagnostics {
		if diagnostic.Code == code {
			if diagnostic.Severity != severity {
				t.Fatalf("diagnostic %s severity = %s, want %s", code, diagnostic.Severity, severity)
			}
			return
		}
	}
	t.Fatalf("diagnostic code %s not found in %#v", code, diagnostics)
}

func countDiagnosticCode(diagnostics []semantic.Diagnostic, code string) int {
	count := 0
	for _, diagnostic := range diagnostics {
		if diagnostic.Code == code {
			count++
		}
	}
	return count
}
