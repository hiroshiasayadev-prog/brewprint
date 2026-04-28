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

func TestRunRender(t *testing.T) {
	yamlRoot := filepath.FromSlash("../../docs/uc/001-ec-checkout-flow/yaml")
	outRoot := t.TempDir()

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	if err := run([]string{"render", "--yaml-root", yamlRoot, "--out", outRoot}, strings.NewReader(""), &stdout, &stderr); err != nil {
		t.Fatalf("run render: %v\nstderr=%s", err, stderr.String())
	}
	if got := strings.TrimSpace(stdout.String()); got != "rendered 23 file(s)" {
		t.Fatalf("render stdout = %q", got)
	}
	if got := strings.TrimSpace(stderr.String()); got != "" {
		t.Fatalf("render stderr = %q", got)
	}

	for _, rel := range []string{
		"index.md",
		"auth/index.md",
		"commerce/index.md",
		"catalog/index.md",
		"commerce/dag-add_to_cart.md",
		"catalog/dag-get_items.md",
		"auth/dag-login.md",
		"commerce/dag-process_payment.md",
		"commerce/seq-checkout_flow.md",
		"_cross/api.md",
		"_preview/wireframe.html",
	} {
		if _, err := os.Stat(filepath.Join(outRoot, filepath.FromSlash(rel))); err != nil {
			t.Fatalf("rendered index missing %s: %v", rel, err)
		}
	}
}

func TestRunRenderCleanRemovesStaleFiles(t *testing.T) {
	yamlRoot := filepath.FromSlash("../../docs/uc/001-ec-checkout-flow/yaml")
	outRoot := t.TempDir()
	stalePath := filepath.Join(outRoot, "stale.md")
	if err := os.WriteFile(stalePath, []byte("stale"), 0o644); err != nil {
		t.Fatalf("write stale file: %v", err)
	}

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	if err := run([]string{"render", "--yaml-root", yamlRoot, "--out", outRoot, "--clean"}, strings.NewReader(""), &stdout, &stderr); err != nil {
		t.Fatalf("run render clean: %v\nstderr=%s", err, stderr.String())
	}
	if got := strings.TrimSpace(stdout.String()); got != "rendered 23 file(s)" {
		t.Fatalf("render clean stdout = %q", got)
	}
	if _, err := os.Stat(stalePath); !os.IsNotExist(err) {
		t.Fatalf("stale file still exists after --clean: %v", err)
	}
	if _, err := os.Stat(filepath.Join(outRoot, "index.md")); err != nil {
		t.Fatalf("rendered index missing after --clean: %v", err)
	}
}

func TestRunRenderCleanRejectsOutRootContainingYAMLRoot(t *testing.T) {
	yamlRoot := filepath.FromSlash("../../docs/uc/001-ec-checkout-flow/yaml")
	outRoot := filepath.Dir(yamlRoot)

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	err := run([]string{"render", "--yaml-root", yamlRoot, "--out", outRoot, "--clean"}, strings.NewReader(""), &stdout, &stderr)
	if err == nil {
		t.Fatalf("run render clean with unsafe out root returned nil error")
	}
	if !strings.Contains(err.Error(), "--clean out root must not contain yaml root") {
		t.Fatalf("render clean unsafe error = %v", err)
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

func assertRenderedFileEquals(t *testing.T, outRoot string, rel string, goldenPath string) {
	t.Helper()
	actualPath := filepath.Join(outRoot, filepath.FromSlash(rel))
	actual, err := os.ReadFile(actualPath)
	if err != nil {
		t.Fatalf("read rendered file %s: %v", rel, err)
	}
	golden, err := os.ReadFile(goldenPath)
	if err != nil {
		t.Fatalf("read golden file %s: %v", goldenPath, err)
	}
	if string(actual) != string(golden) {
		t.Fatalf("rendered file %s did not match golden", rel)
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
	if err := run([]string{"render"}, strings.NewReader(""), &stdout, &stderr); err == nil {
		t.Fatalf("run render without args returned nil error")
	}
	if err := run([]string{"validate", "--yaml-root", "unused", "--format", "yaml"}, strings.NewReader(""), &stdout, &stderr); err == nil {
		t.Fatalf("run validate with unsupported format returned nil error")
	}
	if err := run([]string{"missing"}, strings.NewReader(""), &stdout, &stderr); err == nil {
		t.Fatalf("run unknown command returned nil error")
	}
}
