package model

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/hiroshiasayadev-prog/brewprint/internal/resolve"
	"github.com/hiroshiasayadev-prog/brewprint/internal/semantic"
	"github.com/hiroshiasayadev-prog/brewprint/internal/source"
	"github.com/hiroshiasayadev-prog/brewprint/internal/testutil/golden"
)

func TestRenderModelGolden(t *testing.T) {
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

	cases := []struct {
		name     string
		fileID   semantic.FileID
		goldenMD string
	}{
		{
			name:     "struct with private helpers",
			fileID:   semantic.FileID("auth/model/login_form.yaml"),
			goldenMD: filepath.FromSlash("../../../docs/uc/001-ec-checkout-flow/renders/auth/model-login_form.md"),
		},
		{
			name:     "dict",
			fileID:   semantic.FileID("auth/model/request_context.yaml"),
			goldenMD: filepath.FromSlash("../../../docs/uc/001-ec-checkout-flow/renders/auth/model-request_context.md"),
		},
		{
			name:     "list",
			fileID:   semantic.FileID("cart/model/cart_item_list.yaml"),
			goldenMD: filepath.FromSlash("../../../docs/uc/001-ec-checkout-flow/renders/commerce/model-cart_item_list.md"),
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := RenderFile(project, tc.fileID)
			if err != nil {
				t.Fatalf("render model: %v", err)
			}
			golden.AssertEqualFile(t, tc.goldenMD, actual)
		})
	}
}

func TestRenderTaggedUnionModelGolden(t *testing.T) {
	yamlRoot := filepath.FromSlash("../../../docs/uc/002-brewprint-self-hosting/yaml")
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

	actual, err := RenderFile(project, "mcp/model/analyze_impact_change.yaml")
	if err != nil {
		t.Fatalf("render model: %v", err)
	}
	golden.AssertEqualFile(t, filepath.FromSlash("../../../docs/uc/002-brewprint-self-hosting/renders/mcp/model-analyze_impact_change.md"), actual)
}

func TestRenderModelPrivateHelpersIncludeMinimumKinds(t *testing.T) {
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

	actual, err := RenderFile(project, "auth/model/login_form.yaml")
	if err != nil {
		t.Fatalf("render model: %v", err)
	}
	for _, want := range []string{
		"## Public model",
		"### Fields",
		"## Private models",
		"| login_factor | struct | kind: str<br/>value: str | login_form 内だけで使う追加認証要素。 |",
		"| login_form_status | enum | draft<br/>submitted<br/>rejected | login_form 内だけで使う入力状態。 |",
		"| login_factor_list | list | element: login_factor | login_form 内だけで使う追加認証要素リスト。 |",
		"| login_metadata | dict | value: str | login_form 内だけで使う送信メタデータ。 |",
	} {
		if !strings.Contains(actual, want) {
			t.Fatalf("model render missing %q:\n%s", want, actual)
		}
	}
}
