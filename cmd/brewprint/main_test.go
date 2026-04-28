package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/hiroshiasayadev-prog/brewprint/internal/semantic"
)

func TestRunMCPStdio(t *testing.T) {
	yamlRoot := filepath.FromSlash("../../docs/uc/001-ec-checkout-flow/yaml")
	input := strings.Join([]string{
		`{"jsonrpc":"2.0","id":1,"method":"initialize","params":{}}`,
		`{"jsonrpc":"2.0","id":2,"method":"tools/list"}`,
		`{"jsonrpc":"2.0","id":3,"method":"tools/call","params":{"name":"get_signature","arguments":{"selector":{"id":"auth.task.login"}}}}`,
	}, "\n") + "\n"

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	if err := run([]string{"mcp", "--yaml-root", yamlRoot}, strings.NewReader(input), &stdout, &stderr); err != nil {
		t.Fatalf("run mcp: %v\nstderr=%s", err, stderr.String())
	}

	lines := strings.Split(strings.TrimSpace(stdout.String()), "\n")
	if len(lines) != 3 {
		t.Fatalf("response lines len = %d, want 3\n%s", len(lines), stdout.String())
	}
	for i, line := range lines {
		var response map[string]any
		if err := json.Unmarshal([]byte(line), &response); err != nil {
			t.Fatalf("line %d is not json: %v\n%s", i+1, err, line)
		}
		if response["jsonrpc"] != "2.0" {
			t.Fatalf("line %d jsonrpc = %#v, want 2.0", i+1, response["jsonrpc"])
		}
		if _, ok := response["error"]; ok {
			t.Fatalf("line %d unexpected error: %#v", i+1, response)
		}
	}
}

func TestRunValidate(t *testing.T) {
	yamlRoot := filepath.FromSlash("../../docs/uc/001-ec-checkout-flow/yaml")

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	if err := run([]string{"validate", "--yaml-root", yamlRoot}, strings.NewReader(""), &stdout, &stderr); err != nil {
		t.Fatalf("run validate: %v\nstderr=%s", err, stderr.String())
	}
	if got := strings.TrimSpace(stdout.String()); got != "ok" {
		t.Fatalf("validate stdout = %q, want ok", got)
	}
}

func TestRunValidateInvalidProject(t *testing.T) {
	yamlRoot := invalidValidateYAMLRoot(t)

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	err := run([]string{"validate", "--yaml-root", yamlRoot}, strings.NewReader(""), &stdout, &stderr)
	if err == nil {
		t.Fatalf("run validate invalid project returned nil error")
	}
	if !strings.Contains(err.Error(), "validation failed: 1 error(s), 0 warning(s)") {
		t.Fatalf("validate error = %v", err)
	}
	want := "error unresolved_model auth/task/login.yaml: unresolved task params model: missing_model"
	if got := strings.TrimSpace(stdout.String()); got != want {
		t.Fatalf("validate stdout = %q, want %q", got, want)
	}
}

func TestRunValidateJSON(t *testing.T) {
	yamlRoot := filepath.FromSlash("../../docs/uc/001-ec-checkout-flow/yaml")

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	if err := run([]string{"validate", "--yaml-root", yamlRoot, "--format", "json"}, strings.NewReader(""), &stdout, &stderr); err != nil {
		t.Fatalf("run validate json: %v\nstderr=%s", err, stderr.String())
	}
	var out validateOutput
	if err := json.Unmarshal(stdout.Bytes(), &out); err != nil {
		t.Fatalf("unmarshal validate json: %v\n%s", err, stdout.String())
	}
	if out.ErrorCount != 0 || out.WarningCount != 0 || len(out.Diagnostics) != 0 {
		t.Fatalf("validate json output = %#v", out)
	}
}

func TestRunValidateInvalidProjectJSON(t *testing.T) {
	yamlRoot := invalidValidateYAMLRoot(t)

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	err := run([]string{"validate", "--yaml-root", yamlRoot, "--format", "json"}, strings.NewReader(""), &stdout, &stderr)
	if err == nil {
		t.Fatalf("run validate invalid project json returned nil error")
	}
	var out validateOutput
	if jsonErr := json.Unmarshal(stdout.Bytes(), &out); jsonErr != nil {
		t.Fatalf("unmarshal validate invalid json: %v\n%s", jsonErr, stdout.String())
	}
	if out.ErrorCount != 1 || out.WarningCount != 0 || len(out.Diagnostics) != 1 {
		t.Fatalf("validate invalid json output = %#v", out)
	}
	diagnostic := out.Diagnostics[0]
	if diagnostic.Code != "unresolved_model" || diagnostic.FileID != "auth/task/login.yaml" {
		t.Fatalf("validate invalid json diagnostic = %#v", diagnostic)
	}
}

func TestRunValidateWarningProject(t *testing.T) {
	yamlRoot := warningValidateYAMLRoot(t)

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	if err := run([]string{"validate", "--yaml-root", yamlRoot}, strings.NewReader(""), &stdout, &stderr); err != nil {
		t.Fatalf("run validate warning project: %v\nstderr=%s", err, stderr.String())
	}
	want := "warning unsupported_flow_entry auth/task/login.yaml: unsupported empty flow entry"
	if got := strings.TrimSpace(stdout.String()); got != want {
		t.Fatalf("validate warning stdout = %q, want %q", got, want)
	}
}

func TestRunValidateWarningProjectJSON(t *testing.T) {
	yamlRoot := warningValidateYAMLRoot(t)

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	if err := run([]string{"validate", "--yaml-root", yamlRoot, "--format", "json"}, strings.NewReader(""), &stdout, &stderr); err != nil {
		t.Fatalf("run validate warning project json: %v\nstderr=%s", err, stderr.String())
	}
	var out validateOutput
	if err := json.Unmarshal(stdout.Bytes(), &out); err != nil {
		t.Fatalf("unmarshal validate warning json: %v\n%s", err, stdout.String())
	}
	if out.ErrorCount != 0 || out.WarningCount != 1 || len(out.Diagnostics) != 1 {
		t.Fatalf("validate warning json output = %#v", out)
	}
	diagnostic := out.Diagnostics[0]
	if diagnostic.Severity != semantic.SeverityWarning || diagnostic.Code != "unsupported_flow_entry" {
		t.Fatalf("validate warning json diagnostic = %#v", diagnostic)
	}
}

func TestRunValidateMissingRequiredFieldsJSON(t *testing.T) {
	yamlRoot := missingRequiredValidateYAMLRoot(t)

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	err := run([]string{"validate", "--yaml-root", yamlRoot, "--format", "json"}, strings.NewReader(""), &stdout, &stderr)
	if err == nil {
		t.Fatalf("run validate missing required json returned nil error")
	}
	var out validateOutput
	if jsonErr := json.Unmarshal(stdout.Bytes(), &out); jsonErr != nil {
		t.Fatalf("unmarshal validate missing required json: %v\n%s", jsonErr, stdout.String())
	}
	if out.ErrorCount == 0 || len(out.Diagnostics) == 0 {
		t.Fatalf("validate missing required json output = %#v", out)
	}
	if !hasDiagnosticCode(out.Diagnostics, "missing_required_field") {
		t.Fatalf("missing_required_field diagnostic not found: %#v", out.Diagnostics)
	}
}

func TestFormatDiagnostic(t *testing.T) {
	diagnostic := semantic.Diagnostic{
		Severity: semantic.SeverityError,
		Code:     "unresolved_model",
		FileID:   "order/state.yaml",
		Message:  "unresolved event payload model: payment_event",
	}
	want := "error unresolved_model order/state.yaml: unresolved event payload model: payment_event"
	if got := formatDiagnostic(diagnostic); got != want {
		t.Fatalf("formatDiagnostic = %q, want %q", got, want)
	}
}

func invalidValidateYAMLRoot(t *testing.T) string {
	t.Helper()
	yamlRoot := t.TempDir()
	writeTestYAML(t, yamlRoot, filepath.FromSlash("auth/task/login.yaml"), `nodes:
  - id: login
    type: task
    params:
      - name: credentials
        model: missing_model
`)
	return yamlRoot
}

func warningValidateYAMLRoot(t *testing.T) string {
	t.Helper()
	yamlRoot := t.TempDir()
	writeTestYAML(t, yamlRoot, filepath.FromSlash("auth/task/login.yaml"), `nodes: []
flow:
  - {}
`)
	return yamlRoot
}

func missingRequiredValidateYAMLRoot(t *testing.T) string {
	t.Helper()
	yamlRoot := t.TempDir()
	writeTestYAML(t, yamlRoot, filepath.FromSlash("auth/model/broken.yaml"), `nodes:
  - id: broken
    type: model
    fields:
      - name: id
`)
	return yamlRoot
}

func hasDiagnosticCode(diagnostics []semantic.Diagnostic, code string) bool {
	for _, diagnostic := range diagnostics {
		if diagnostic.Code == code {
			return true
		}
	}
	return false
}

func writeTestYAML(t *testing.T, root string, rel string, content string) {
	t.Helper()
	path := filepath.Join(root, rel)
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("mkdir test yaml dir: %v", err)
	}
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("write test yaml: %v", err)
	}
}

func TestRunErrors(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	if err := run(nil, strings.NewReader(""), &stdout, &stderr); err == nil {
		t.Fatalf("run without args returned nil error")
	}
	if err := run([]string{"mcp"}, strings.NewReader(""), &stdout, &stderr); err == nil {
		t.Fatalf("run mcp without yaml-root returned nil error")
	}
	if err := run([]string{"validate"}, strings.NewReader(""), &stdout, &stderr); err == nil {
		t.Fatalf("run validate without yaml-root returned nil error")
	}
	if err := run([]string{"validate", "--yaml-root", "unused", "--format", "yaml"}, strings.NewReader(""), &stdout, &stderr); err == nil {
		t.Fatalf("run validate with unsupported format returned nil error")
	}
	if err := run([]string{"missing"}, strings.NewReader(""), &stdout, &stderr); err == nil {
		t.Fatalf("run unknown command returned nil error")
	}
}
