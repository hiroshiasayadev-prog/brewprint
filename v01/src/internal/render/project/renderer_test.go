package project

import (
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"testing"

	"github.com/hiroshiasayadev-prog/brewprint/v01/src/internal/rawyaml"
	"github.com/hiroshiasayadev-prog/brewprint/v01/src/internal/resolve"
	"github.com/hiroshiasayadev-prog/brewprint/v01/src/internal/semantic"
	"github.com/hiroshiasayadev-prog/brewprint/v01/src/internal/source"
)

func TestRenderUC001Manifest(t *testing.T) {
	raw, semanticProject := loadUC001(t)

	files, diagnostics, err := Render(raw, semanticProject)
	if err != nil {
		t.Fatalf("render project: %v", err)
	}
	if len(diagnostics) != 0 {
		t.Fatalf("render diagnostics = %#v", diagnostics)
	}

	var got []string
	for _, file := range files {
		if strings.TrimSpace(file.Content) == "" {
			t.Fatalf("rendered file %s has empty content", file.Path)
		}
		got = append(got, file.Path)
	}
	sort.Strings(got)

	want := []string{
		"_cross/api.md",
		"_cross/er.md",
		"_preview/wireframe.html",
		"auth/dag-login.md",
		"auth/index.md",
		"auth/model-credential.md",
		"auth/model-login_form.md",
		"auth/model-login_log.md",
		"auth/model-request_context.md",
		"auth/model-token.md",
		"auth/state-auth.md",
		"auth/wireframe-auth-loading.html",
		"auth/wireframe-auth-login_screen.html",
		"catalog/dag-get_items.md",
		"catalog/index.md",
		"catalog/model-item.md",
		"catalog/model-item_list.md",
		"catalog/state-inventory.md",
		"commerce/dag-add_to_cart.md",
		"commerce/dag-checkout.md",
		"commerce/dag-process_order.md",
		"commerce/dag-process_payment.md",
		"commerce/dag-validate_cart.md",
		"commerce/index.md",
		"commerce/model-address.md",
		"commerce/model-cart.md",
		"commerce/model-cart_item.md",
		"commerce/model-cart_item_list.md",
		"commerce/model-order.md",
		"commerce/model-order_item.md",
		"commerce/model-payment_event.md",
		"commerce/seq-checkout_flow.md",
		"commerce/seq-payment_webhook_flow.md",
		"commerce/state-order.md",
		"commerce/wireframe-order-cart.html",
		"commerce/wireframe-order-checkout_screen.html",
		"index.md",
	}
	sort.Strings(want)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("render paths = %#v, want %#v", got, want)
	}
}

func TestRenderUC001IndexesMatchFixtures(t *testing.T) {
	raw, semanticProject := loadUC001(t)

	files, diagnostics, err := Render(raw, semanticProject)
	if err != nil {
		t.Fatalf("render project: %v", err)
	}
	if len(diagnostics) != 0 {
		t.Fatalf("render diagnostics = %#v", diagnostics)
	}

	byPath := map[string]string{}
	for _, file := range files {
		byPath[file.Path] = file.Content
	}

	for _, path := range []string{
		"index.md",
		"auth/index.md",
		"catalog/index.md",
		"commerce/index.md",
	} {
		want, err := os.ReadFile(filepath.Join("../../../docs/uc/001-ec-checkout-flow/renders", filepath.FromSlash(path)))
		if err != nil {
			t.Fatalf("read fixture %s: %v", path, err)
		}
		if got := byPath[path]; got != string(want) {
			t.Fatalf("rendered index %s = %q, want %q", path, got, string(want))
		}
	}
}

func TestHumanizeProjectDir(t *testing.T) {
	if got := humanizeProjectDir("001-ec-checkout-flow"); got != "EC Checkout Flow" {
		t.Fatalf("humanizeProjectDir = %q, want EC Checkout Flow", got)
	}
	if got := humanizeProjectDir("sample-api"); got != "Sample API" {
		t.Fatalf("humanizeProjectDir = %q, want Sample API", got)
	}
}

func TestWriteCreatesNestedOutput(t *testing.T) {
	outRoot := t.TempDir()
	files := []File{{Path: "group/nested.md", Content: "# nested\n", Source: "test source"}}

	if err := Write(outRoot, files); err != nil {
		t.Fatalf("write renders: %v", err)
	}
	content, err := os.ReadFile(filepath.Join(outRoot, "group", "nested.md"))
	if err != nil {
		t.Fatalf("read written render: %v", err)
	}
	if string(content) != "# nested\n" {
		t.Fatalf("written content = %q", string(content))
	}
}

func TestValidateFilePathCollisionsReportsSourceObjects(t *testing.T) {
	files := []File{
		{Path: "commerce/dag-process_payment.md", Content: "first", Source: "task payment.task.process_payment (payment/task/process_payment.yaml)"},
		{Path: "commerce/dag-process_payment.md", Content: "second", Source: "task payment.webhooks.task.process_payment (payment/webhooks/task/process_payment.yaml)"},
	}

	err := validateFilePathCollisions(files)
	if err == nil {
		t.Fatalf("validateFilePathCollisions returned nil error")
	}
	message := err.Error()
	for _, want := range []string{
		"render output path collision",
		"commerce/dag-process_payment.md",
		"payment/task/process_payment.yaml",
		"payment/webhooks/task/process_payment.yaml",
	} {
		if !strings.Contains(message, want) {
			t.Fatalf("collision error missing %q: %s", want, message)
		}
	}
}

func TestWriteRejectsDuplicatePathsBeforeOverwrite(t *testing.T) {
	outRoot := t.TempDir()
	files := []File{
		{Path: "commerce/dag-process_payment.md", Content: "first\n", Source: "first source"},
		{Path: "commerce/dag-process_payment.md", Content: "second\n", Source: "second source"},
	}

	err := Write(outRoot, files)
	if err == nil {
		t.Fatalf("Write with duplicate paths returned nil error")
	}
	if !strings.Contains(err.Error(), "first source") || !strings.Contains(err.Error(), "second source") {
		t.Fatalf("duplicate path error does not include both sources: %v", err)
	}
	if _, statErr := os.Stat(filepath.Join(outRoot, "commerce", "dag-process_payment.md")); !os.IsNotExist(statErr) {
		t.Fatalf("duplicate path write created output file: %v", statErr)
	}
}

func TestCleanOutRootRemovesExistingDirectory(t *testing.T) {
	outRoot := filepath.Join(t.TempDir(), "renders")
	stalePath := filepath.Join(outRoot, "stale.md")
	if err := os.MkdirAll(filepath.Dir(stalePath), 0o755); err != nil {
		t.Fatalf("mkdir stale dir: %v", err)
	}
	if err := os.WriteFile(stalePath, []byte("stale"), 0o644); err != nil {
		t.Fatalf("write stale file: %v", err)
	}

	if err := CleanOutRoot(outRoot); err != nil {
		t.Fatalf("clean out root: %v", err)
	}
	if _, err := os.Stat(stalePath); !os.IsNotExist(err) {
		t.Fatalf("stale file still exists after clean: %v", err)
	}
	if info, err := os.Stat(outRoot); err != nil || !info.IsDir() {
		t.Fatalf("out root not recreated as directory: info=%v err=%v", info, err)
	}
}

func TestCleanOutRootRejectsUnsafePaths(t *testing.T) {
	for _, path := range []string{"", ".", string(filepath.Separator), filepath.Join("..", "renders")} {
		if err := CleanOutRoot(path); err == nil {
			t.Fatalf("CleanOutRoot(%q) returned nil error", path)
		}
	}
}

func loadUC001(t *testing.T) (rawProject *rawyaml.Project, semanticProject *semantic.Project) {
	t.Helper()
	yamlRoot := filepath.FromSlash("../../../docs/uc/001-ec-checkout-flow/yaml")
	loader := source.Loader{}
	raw, err := loader.Load(yamlRoot)
	if err != nil {
		t.Fatalf("load yaml root: %v", err)
	}
	project, diagnostics := resolve.Build(raw)
	for _, diagnostic := range diagnostics {
		if diagnostic.Severity == semantic.SeverityError {
			t.Fatalf("semantic diagnostic: %s: %s", diagnostic.FileID, diagnostic.Message)
		}
	}
	return raw, project
}
