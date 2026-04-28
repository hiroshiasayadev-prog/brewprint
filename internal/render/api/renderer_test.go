package api

import (
	"path/filepath"
	"testing"

	"github.com/hiroshiasayadev-prog/brewprint/internal/resolve"
	"github.com/hiroshiasayadev-prog/brewprint/internal/semantic"
	"github.com/hiroshiasayadev-prog/brewprint/internal/source"
	"github.com/hiroshiasayadev-prog/brewprint/internal/testutil/golden"
)

func TestRenderAPIGolden(t *testing.T) {
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
	assertAPIViewResolved(t, project)

	actual, err := RenderView(project, "ec_api")
	if err != nil {
		t.Fatalf("render API: %v", err)
	}
	golden.AssertEqualFile(t, filepath.FromSlash("../../../docs/uc/001-ec-checkout-flow/renders/_cross/api.md"), actual)
}

func assertAPIViewResolved(t *testing.T, project *semantic.Project) {
	t.Helper()
	view := project.APIViewsByID["ec_api"]
	if view == nil {
		t.Fatalf("ec_api view not found")
	}
	if view.HTTPRootPath != "/api" {
		t.Fatalf("ec_api http_root_path = %s, want /api", view.HTTPRootPath)
	}
	wantModules := []string{"auth", "catalog", "cart", "order", "payment.webhooks"}
	if len(view.Modules) != len(wantModules) {
		t.Fatalf("ec_api modules = %#v, want %v", view.Modules, wantModules)
	}
	for i, want := range wantModules {
		if view.Modules[i].Module != want {
			t.Fatalf("ec_api module[%d] = %s, want %s", i, view.Modules[i].Module, want)
		}
		if view.Modules[i].IncludeSubmodules {
			t.Fatalf("ec_api module[%d] include_submodules = true, want false", i)
		}
	}
}
