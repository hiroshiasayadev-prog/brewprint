package sequence

import (
	"path/filepath"
	"testing"

	"github.com/hiroshiasayadev-prog/brewprint/v01/src/internal/resolve"
	"github.com/hiroshiasayadev-prog/brewprint/v01/src/internal/semantic"
	"github.com/hiroshiasayadev-prog/brewprint/v01/src/internal/source"
	"github.com/hiroshiasayadev-prog/brewprint/v01/src/internal/testutil/golden"
)

func TestRenderSequenceGolden(t *testing.T) {
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
		name       string
		scenarioID string
		goldenMD   string
	}{
		{
			name:       "checkout_flow",
			scenarioID: "checkout_flow",
			goldenMD:   filepath.FromSlash("../../../docs/uc/001-ec-checkout-flow/renders/commerce/seq-checkout_flow.md"),
		},
		{
			name:       "payment_webhook_flow",
			scenarioID: "payment_webhook_flow",
			goldenMD:   filepath.FromSlash("../../../docs/uc/001-ec-checkout-flow/renders/commerce/seq-payment_webhook_flow.md"),
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := RenderScenario(project, tc.scenarioID)
			if err != nil {
				t.Fatalf("render sequence: %v", err)
			}
			golden.AssertEqualFile(t, tc.goldenMD, actual)
		})
	}
}
