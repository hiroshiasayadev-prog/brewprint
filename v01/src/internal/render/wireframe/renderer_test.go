package wireframe

import (
	"path/filepath"
	"testing"

	"github.com/hiroshiasayadev-prog/brewprint/v01/src/internal/resolve"
	"github.com/hiroshiasayadev-prog/brewprint/v01/src/internal/semantic"
	"github.com/hiroshiasayadev-prog/brewprint/v01/src/internal/source"
	"github.com/hiroshiasayadev-prog/brewprint/v01/src/internal/testutil/golden"
)

func TestRenderWireframeGolden(t *testing.T) {
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
	assertWireframesResolved(t, project)

	cases := []struct {
		name     string
		stateQID semantic.QualifiedID
		golden   string
	}{
		{
			name:     "auth_login_screen",
			stateQID: semantic.QualifiedID("auth.state.login_screen"),
			golden:   filepath.FromSlash("../../../docs/uc/001-ec-checkout-flow/renders/auth/wireframe-auth-login_screen.html"),
		},
		{
			name:     "auth_loading",
			stateQID: semantic.QualifiedID("auth.state.loading"),
			golden:   filepath.FromSlash("../../../docs/uc/001-ec-checkout-flow/renders/auth/wireframe-auth-loading.html"),
		},
		{
			name:     "order_cart",
			stateQID: semantic.QualifiedID("order.state.cart"),
			golden:   filepath.FromSlash("../../../docs/uc/001-ec-checkout-flow/renders/commerce/wireframe-order-cart.html"),
		},
		{
			name:     "order_checkout_screen",
			stateQID: semantic.QualifiedID("order.state.checkout_screen"),
			golden:   filepath.FromSlash("../../../docs/uc/001-ec-checkout-flow/renders/commerce/wireframe-order-checkout_screen.html"),
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := RenderState(project, tc.stateQID)
			if err != nil {
				t.Fatalf("render wireframe: %v", err)
			}
			golden.AssertEqualFile(t, tc.golden, actual)
		})
	}
}

func TestRenderWireframePreviewGolden(t *testing.T) {
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

	actual := RenderPreview(project, "EC Checkout Flow Wireframe Preview")
	golden.AssertEqualFile(t, filepath.FromSlash("../../../docs/uc/001-ec-checkout-flow/renders/_preview/wireframe.html"), actual)
}

func assertWireframesResolved(t *testing.T, project *semantic.Project) {
	t.Helper()
	for _, qid := range []semantic.QualifiedID{
		"auth.state.login_screen",
		"auth.state.loading",
		"order.state.cart",
		"order.state.checkout_screen",
	} {
		state := project.StatesByQID[qid]
		if state == nil {
			t.Fatalf("state not found: %s", qid)
		}
		if state.Wireframe == nil {
			t.Fatalf("state has no wireframe: %s", qid)
		}
	}
}
