package resolve

import (
	"strings"
	"testing"

	"github.com/hiroshiasayadev-prog/brewprint/internal/rawyaml"
	"github.com/hiroshiasayadev-prog/brewprint/internal/semantic"
)

func TestTaskFileHelperModelsResolveSameFileTypeRefs(t *testing.T) {
	project, diagnostics := Build(&rawyaml.Project{Files: []rawyaml.File{taskFileWithHelpers("auth/task/login.yaml", "login", "login_form", "login_token")}})
	if len(diagnostics) != 0 {
		t.Fatalf("diagnostics = %#v, want none", diagnostics)
	}

	task := project.TasksByQID["auth.task.login"]
	if task == nil {
		t.Fatalf("auth.task.login not found")
	}
	if got, want := task.Returns.Model, semantic.QualifiedID("auth/task/login.yaml#login_token"); got != want {
		t.Fatalf("return model = %s, want %s", got, want)
	}

	token := project.PrivateModelsByFile["auth/task/login.yaml"]["login_token"]
	if token == nil || len(token.Fields) != 1 {
		t.Fatalf("login_token helper not built: %#v", token)
	}
	if got, want := token.Fields[0].TypeRef.Model, semantic.QualifiedID("auth/task/login.yaml#login_credentials"); got != want {
		t.Fatalf("helper field TypeRef model = %s, want %s", got, want)
	}
	if project.ModelsByQID["auth.model.login_form"] != nil {
		t.Fatalf("task-file helper leaked as public model")
	}
}

func TestTaskFilePrivateHelperParamModelIsRejected(t *testing.T) {
	_, diagnostics := Build(&rawyaml.Project{Files: []rawyaml.File{taskFileWithPrivateParamHelper("auth/task/login.yaml")}})

	assertDiagnosticCode(t, diagnostics, diagnosticInvalidPrivateModelRef)
	assertNoDiagnosticCode(t, diagnostics, diagnosticUnresolvedModel)

	var found semantic.Diagnostic
	for _, diagnostic := range diagnostics {
		if diagnostic.Code == diagnosticInvalidPrivateModelRef {
			found = diagnostic
			break
		}
	}
	if found.FileID != "auth/task/login.yaml" {
		t.Fatalf("invalid_private_model_reference file = %s, want auth/task/login.yaml", found.FileID)
	}
	if !strings.Contains(found.Message, "params[].model") || !strings.Contains(found.Message, "login_form") || !strings.Contains(found.Message, "auth/task/login.yaml#login_form") {
		t.Fatalf("invalid_private_model_reference message does not identify params[].model and model id: %q", found.Message)
	}
}

func TestTaskFilePrivateHelperParamModelIsRejectedForControlNodes(t *testing.T) {
	_, diagnostics := Build(&rawyaml.Project{Files: []rawyaml.File{{
		ID:   "auth/task/login.yaml",
		Kind: rawyaml.FileKindNode,
		NodeFile: &rawyaml.NodeFile{
			Tasks:    []rawyaml.Task{{ID: "login", Main: true}},
			Branches: []rawyaml.ControlNode{{ID: "route", Params: []rawyaml.Param{{Name: "form", Model: "login_form"}}}},
			Forks:    []rawyaml.ControlNode{{ID: "fan_out", Params: []rawyaml.Param{{Name: "form", Model: "login_form"}}}},
			Joins:    []rawyaml.ControlNode{{ID: "aggregate", Params: []rawyaml.Param{{Name: "form", Model: "login_form"}}}},
			Models: []rawyaml.Model{{
				ID:     "login_form",
				Kind:   "struct",
				Fields: []rawyaml.ModelField{{Name: "username", Type: "str"}},
			}},
		},
	}}})

	if got := countDiagnosticCode(diagnostics, diagnosticInvalidPrivateModelRef); got != 3 {
		t.Fatalf("invalid_private_model_reference count = %d, want 3: %#v", got, diagnostics)
	}
	assertNoDiagnosticCode(t, diagnostics, diagnosticUnresolvedModel)
}

func TestTaskFilePrivateHelperReturnModelIsAccepted(t *testing.T) {
	_, diagnostics := Build(&rawyaml.Project{Files: []rawyaml.File{{
		ID:   "auth/task/login.yaml",
		Kind: rawyaml.FileKindNode,
		NodeFile: &rawyaml.NodeFile{
			Tasks: []rawyaml.Task{{
				ID:      "login",
				Main:    true,
				Returns: &rawyaml.Return{Name: "result", Model: "login_token"},
			}},
			Models: []rawyaml.Model{{
				ID:     "login_token",
				Kind:   "struct",
				Fields: []rawyaml.ModelField{{Name: "token", Type: "str"}},
			}},
		},
	}}})

	if len(diagnostics) != 0 {
		t.Fatalf("private helper returns.model diagnostics = %#v, want none", diagnostics)
	}
}

func TestTaskFilePrivateHelperUnresolvedParamStillUsesUnresolvedModel(t *testing.T) {
	_, diagnostics := Build(&rawyaml.Project{Files: []rawyaml.File{{
		ID:   "auth/task/login.yaml",
		Kind: rawyaml.FileKindNode,
		NodeFile: &rawyaml.NodeFile{Tasks: []rawyaml.Task{{
			ID:     "login",
			Main:   true,
			Params: []rawyaml.Param{{Name: "form", Model: "missing_model"}},
		}}},
	}}})

	assertDiagnosticCode(t, diagnostics, diagnosticUnresolvedModel)
	assertNoDiagnosticCode(t, diagnostics, diagnosticInvalidPrivateModelRef)
}

func TestTaskFileHelperModelConflictsWithPublicModelID(t *testing.T) {
	_, diagnostics := Build(&rawyaml.Project{Files: []rawyaml.File{
		typeRefStructModelFile("auth/model/login_form.yaml", "login_form"),
		taskFileWithHelpers("auth/task/login.yaml", "login", "login_form", "login_token"),
	}})
	assertDiagnosticCode(t, diagnostics, diagnosticDuplicateModelID)
}

func TestTaskFileHelperModelDuplicateInSameFile(t *testing.T) {
	_, diagnostics := Build(&rawyaml.Project{Files: []rawyaml.File{{
		ID:   "auth/task/login.yaml",
		Kind: rawyaml.FileKindNode,
		NodeFile: &rawyaml.NodeFile{
			Tasks: []rawyaml.Task{{ID: "login", Main: true, Params: []rawyaml.Param{{Name: "form", Model: "login_form"}}}},
			Models: []rawyaml.Model{
				{ID: "login_form", Kind: "struct", Fields: []rawyaml.ModelField{{Name: "username", Type: "str"}}},
				{ID: "login_form", Kind: "struct", Fields: []rawyaml.ModelField{{Name: "email", Type: "str"}}},
			},
		},
	}}})
	assertDiagnosticCode(t, diagnostics, diagnosticDuplicateModelID)
	assertNoDiagnosticCode(t, diagnostics, diagnosticInvalidPrivateModelRef)
}

func TestTaskFileHelperModelConflictsWithMainTaskLocalID(t *testing.T) {
	_, diagnostics := Build(&rawyaml.Project{Files: []rawyaml.File{{
		ID:   "auth/task/login.yaml",
		Kind: rawyaml.FileKindNode,
		NodeFile: &rawyaml.NodeFile{
			Tasks: []rawyaml.Task{{ID: "login", Main: true, Params: []rawyaml.Param{{Name: "payload", Model: "login"}}}},
			Models: []rawyaml.Model{{
				ID:     "login",
				Kind:   "struct",
				Fields: []rawyaml.ModelField{{Name: "username", Type: "str"}},
			}},
		},
	}}})
	assertDiagnosticCode(t, diagnostics, diagnosticDuplicateModelID)
	assertNoDiagnosticCode(t, diagnostics, diagnosticUnresolvedModel)
	assertNoDiagnosticCode(t, diagnostics, diagnosticUnresolvedFieldType)
	assertNoDiagnosticCode(t, diagnostics, diagnosticInvalidPrivateModelRef)
}

func TestTaskFileHelperModelConflictsWithPrivateSubTaskLocalID(t *testing.T) {
	_, diagnostics := Build(&rawyaml.Project{Files: []rawyaml.File{{
		ID:   "auth/task/login.yaml",
		Kind: rawyaml.FileKindNode,
		NodeFile: &rawyaml.NodeFile{
			Tasks: []rawyaml.Task{{ID: "login", Main: true}, {ID: "normalize"}},
			Models: []rawyaml.Model{{
				ID:     "normalize",
				Kind:   "struct",
				Fields: []rawyaml.ModelField{{Name: "username", Type: "str"}},
			}},
		},
	}}})
	assertDiagnosticCode(t, diagnostics, diagnosticDuplicateModelID)
	assertNoDiagnosticCode(t, diagnostics, diagnosticUnresolvedModel)
	assertNoDiagnosticCode(t, diagnostics, diagnosticUnresolvedFieldType)
}

func TestTaskFileModelMainTrueIsInvalidAndNotPrivateHelper(t *testing.T) {
	project, diagnostics := Build(&rawyaml.Project{Files: []rawyaml.File{{
		ID:   "auth/task/login.yaml",
		Kind: rawyaml.FileKindNode,
		NodeFile: &rawyaml.NodeFile{Models: []rawyaml.Model{{
			ID:     "login_form",
			Main:   true,
			Kind:   "struct",
			Fields: []rawyaml.ModelField{{Name: "username", Type: "str"}},
		}}},
	}}})
	assertDiagnosticCode(t, diagnostics, diagnosticSemanticValidation)
	if project.PrivateModelsByFile["auth/task/login.yaml"]["login_form"] != nil {
		t.Fatalf("task-file main model was registered as private helper")
	}
}

func TestPublicModelFileMainModelRemainsValid(t *testing.T) {
	project, diagnostics := Build(&rawyaml.Project{Files: []rawyaml.File{{
		ID:   "auth/model/login_form.yaml",
		Kind: rawyaml.FileKindNode,
		NodeFile: &rawyaml.NodeFile{Models: []rawyaml.Model{{
			ID:     "login_form",
			Main:   true,
			Kind:   "struct",
			Fields: []rawyaml.ModelField{{Name: "username", Type: "str"}},
		}}},
	}}})
	if len(diagnostics) != 0 {
		t.Fatalf("diagnostics = %#v, want none", diagnostics)
	}
	if project.ModelsByQID["auth.model.login_form"] == nil {
		t.Fatalf("public model main node missing from ModelsByQID")
	}
}

func TestModelFileHelperModelsResolveSameFileTypeRefs(t *testing.T) {
	project, diagnostics := Build(&rawyaml.Project{Files: []rawyaml.File{{
		ID:   "auth/model/login_form.yaml",
		Kind: rawyaml.FileKindNode,
		NodeFile: &rawyaml.NodeFile{Models: []rawyaml.Model{
			{
				ID:   "login_form",
				Main: true,
				Kind: "struct",
				Fields: []rawyaml.ModelField{
					{Name: "factors", Type: "login_factor_list"},
					{Name: "status", Type: "login_form_status"},
					{Name: "metadata", Type: "login_metadata"},
				},
			},
			{ID: "login_factor", Kind: "struct", Fields: []rawyaml.ModelField{{Name: "kind", Type: "str"}}},
			{ID: "login_form_status", Kind: "enum", Values: []string{"draft", "submitted"}},
			{ID: "login_factor_list", Kind: "list", Element: "login_factor"},
			{ID: "login_metadata", Kind: "dict", Value: "str"},
		}},
	}}})
	if len(diagnostics) != 0 {
		t.Fatalf("diagnostics = %#v, want none", diagnostics)
	}

	main := project.ModelsByQID["auth.model.login_form"]
	if main == nil {
		t.Fatalf("public main model not found")
	}
	if project.ModelsByQID["auth.model.login_factor"] != nil {
		t.Fatalf("model-file helper leaked as public model")
	}
	helper := project.PrivateModelsByFile["auth/model/login_form.yaml"]["login_factor_list"]
	if helper == nil {
		t.Fatalf("login_factor_list helper not found")
	}
	if got, want := helper.ElementRef.Model, semantic.QualifiedID("auth/model/login_form.yaml#login_factor"); got != want {
		t.Fatalf("helper element TypeRef model = %s, want %s", got, want)
	}
	if got, want := main.Fields[0].TypeRef.Model, semantic.QualifiedID("auth/model/login_form.yaml#login_factor_list"); got != want {
		t.Fatalf("main field TypeRef model = %s, want %s", got, want)
	}
}

func TestTaskFileHelperModelSameLocalIDAcrossFilesIsAllowed(t *testing.T) {
	project, diagnostics := Build(&rawyaml.Project{Files: []rawyaml.File{
		taskFileWithHelpers("auth/task/login.yaml", "login", "request", "response"),
		taskFileWithHelpers("auth/task/register.yaml", "register", "request", "response"),
	}})
	if len(diagnostics) != 0 {
		t.Fatalf("diagnostics = %#v, want none", diagnostics)
	}
	if project.PrivateModelsByFile["auth/task/login.yaml"]["request"] == nil {
		t.Fatalf("login request helper not found")
	}
	if project.PrivateModelsByFile["auth/task/register.yaml"]["request"] == nil {
		t.Fatalf("register request helper not found")
	}
}

func TestQualifiedTypeRefCannotReferenceTaskFileHelperModel(t *testing.T) {
	_, diagnostics := Build(&rawyaml.Project{Files: []rawyaml.File{
		taskFileWithHelpers("auth/task/login.yaml", "login", "login_form", "login_token"),
		{
			ID:   "auth/task/register.yaml",
			Kind: rawyaml.FileKindNode,
			NodeFile: &rawyaml.NodeFile{Tasks: []rawyaml.Task{{
				ID:     "register",
				Main:   true,
				Params: []rawyaml.Param{{Name: "form", Model: "auth.model.login_form"}},
			}}},
		},
	}})
	assertDiagnosticCode(t, diagnostics, diagnosticUnresolvedModel)
}

func taskFileWithHelpers(fileID string, taskID string, paramModel string, returnModel string) rawyaml.File {
	return rawyaml.File{
		ID:   fileID,
		Kind: rawyaml.FileKindNode,
		NodeFile: &rawyaml.NodeFile{
			Tasks: []rawyaml.Task{{
				ID:      taskID,
				Main:    true,
				Params:  []rawyaml.Param{{Name: "request_id", Model: "str"}},
				Returns: &rawyaml.Return{Name: "result", Model: returnModel},
			}},
			Models: []rawyaml.Model{
				{
					ID:     paramModel,
					Kind:   "struct",
					Fields: []rawyaml.ModelField{{Name: "username", Type: "str"}},
				},
				{
					ID:     "login_credentials",
					Kind:   "struct",
					Fields: []rawyaml.ModelField{{Name: "password", Type: "str"}},
				},
				{
					ID:     returnModel,
					Kind:   "struct",
					Fields: []rawyaml.ModelField{{Name: "credentials", Type: "login_credentials"}},
				},
			},
		},
	}
}

func taskFileWithPrivateParamHelper(fileID string) rawyaml.File {
	return rawyaml.File{
		ID:   fileID,
		Kind: rawyaml.FileKindNode,
		NodeFile: &rawyaml.NodeFile{
			Tasks: []rawyaml.Task{{
				ID:      "login",
				Main:    true,
				Params:  []rawyaml.Param{{Name: "form", Model: "login_form"}},
				Returns: &rawyaml.Return{Name: "result", Model: "str"},
			}},
			Models: []rawyaml.Model{{
				ID:     "login_form",
				Kind:   "struct",
				Fields: []rawyaml.ModelField{{Name: "username", Type: "str"}},
			}},
		},
	}
}
