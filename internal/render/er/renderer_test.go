package er

import (
	"path/filepath"
	"testing"

	"github.com/hiroshiasayadev-prog/brewprint/internal/resolve"
	"github.com/hiroshiasayadev-prog/brewprint/internal/semantic"
	"github.com/hiroshiasayadev-prog/brewprint/internal/source"
	"github.com/hiroshiasayadev-prog/brewprint/internal/testutil/golden"
)

func TestRenderERGolden(t *testing.T) {
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
	assertERViewResolved(t, project)

	actual, err := RenderView(project, "ec_er")
	if err != nil {
		t.Fatalf("render ER: %v", err)
	}
	golden.AssertEqualFile(t, filepath.FromSlash("../../../docs/uc/001-ec-checkout-flow/renders/_cross/er.md"), actual)
}

func assertERViewResolved(t *testing.T, project *semantic.Project) {
	t.Helper()
	view := project.ERViewsByID["ec_er"]
	if view == nil {
		t.Fatalf("ec_er view not found")
	}
	wantModules := []string{"auth", "catalog", "cart", "order", "payment"}
	if len(view.Modules) != len(wantModules) {
		t.Fatalf("ec_er modules = %#v, want %v", view.Modules, wantModules)
	}
	for i, want := range wantModules {
		if view.Modules[i].Module != want {
			t.Fatalf("ec_er module[%d] = %s, want %s", i, view.Modules[i].Module, want)
		}
	}
}
