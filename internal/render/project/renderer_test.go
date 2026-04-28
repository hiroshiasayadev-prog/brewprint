package project

import (
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"testing"

	"github.com/hiroshiasayadev-prog/brewprint/internal/rawyaml"
	"github.com/hiroshiasayadev-prog/brewprint/internal/resolve"
	"github.com/hiroshiasayadev-prog/brewprint/internal/semantic"
	"github.com/hiroshiasayadev-prog/brewprint/internal/source"
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
		"auth/state-auth.md",
		"auth/wireframe-auth-loading.html",
		"auth/wireframe-auth-login_screen.html",
		"catalog/dag-get_items.md",
		"catalog/index.md",
		"catalog/state-inventory.md",
		"commerce/dag-add_to_cart.md",
		"commerce/dag-checkout.md",
		"commerce/dag-process_order.md",
		"commerce/dag-process_payment.md",
		"commerce/dag-validate_cart.md",
		"commerce/index.md",
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

func TestWriteCreatesNestedOutput(t *testing.T) {
	outRoot := t.TempDir()
	files := []File{{Path: "group/nested.md", Content: "# nested\n"}}

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
