package state

import (
	"path/filepath"
	"testing"

	"github.com/hiroshiasayadev-prog/brewprint/v01/src/internal/resolve"
	"github.com/hiroshiasayadev-prog/brewprint/v01/src/internal/semantic"
	"github.com/hiroshiasayadev-prog/brewprint/v01/src/internal/source"
	"github.com/hiroshiasayadev-prog/brewprint/v01/src/internal/testutil/golden"
)

func TestRenderStateGolden(t *testing.T) {
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
			name:     "auth",
			fileID:   semantic.FileID("auth/state.yaml"),
			goldenMD: filepath.FromSlash("../../../docs/uc/001-ec-checkout-flow/renders/auth/state-auth.md"),
		},
		{
			name:     "order",
			fileID:   semantic.FileID("order/state.yaml"),
			goldenMD: filepath.FromSlash("../../../docs/uc/001-ec-checkout-flow/renders/commerce/state-order.md"),
		},
		{
			name:     "inventory",
			fileID:   semantic.FileID("inventory/state.yaml"),
			goldenMD: filepath.FromSlash("../../../docs/uc/001-ec-checkout-flow/renders/catalog/state-inventory.md"),
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := RenderFile(project, tc.fileID)
			if err != nil {
				t.Fatalf("render state: %v", err)
			}
			golden.AssertEqualFile(t, tc.goldenMD, actual)
		})
	}
}
