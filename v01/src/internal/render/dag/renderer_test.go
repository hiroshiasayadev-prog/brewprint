package dag

import (
	"path/filepath"
	"testing"

	"github.com/hiroshiasayadev-prog/brewprint/v01/src/internal/resolve"
	"github.com/hiroshiasayadev-prog/brewprint/v01/src/internal/semantic"
	"github.com/hiroshiasayadev-prog/brewprint/v01/src/internal/source"
	"github.com/hiroshiasayadev-prog/brewprint/v01/src/internal/testutil/golden"
)

func TestRenderDAGGolden(t *testing.T) {
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
	assertLoginResolved(t, project)
	assertM2Resolved(t, project)

	cases := []struct {
		name     string
		fileID   semantic.FileID
		goldenMD string
	}{
		{
			name:     "login",
			fileID:   semantic.FileID("auth/task/login.yaml"),
			goldenMD: filepath.FromSlash("../../../docs/uc/001-ec-checkout-flow/renders/auth/dag-login.md"),
		},
		{
			name:     "add_to_cart",
			fileID:   semantic.FileID("cart/task/add_to_cart.yaml"),
			goldenMD: filepath.FromSlash("../../../docs/uc/001-ec-checkout-flow/renders/commerce/dag-add_to_cart.md"),
		},
		{
			name:     "validate_cart",
			fileID:   semantic.FileID("cart/task/validate_cart.yaml"),
			goldenMD: filepath.FromSlash("../../../docs/uc/001-ec-checkout-flow/renders/commerce/dag-validate_cart.md"),
		},
		{
			name:     "checkout",
			fileID:   semantic.FileID("order/task/checkout.yaml"),
			goldenMD: filepath.FromSlash("../../../docs/uc/001-ec-checkout-flow/renders/commerce/dag-checkout.md"),
		},
		{
			name:     "process_order",
			fileID:   semantic.FileID("order/task/process_order.yaml"),
			goldenMD: filepath.FromSlash("../../../docs/uc/001-ec-checkout-flow/renders/commerce/dag-process_order.md"),
		},
		{
			name:     "process_payment",
			fileID:   semantic.FileID("payment/webhooks/task/process_payment.yaml"),
			goldenMD: filepath.FromSlash("../../../docs/uc/001-ec-checkout-flow/renders/commerce/dag-process_payment.md"),
		},
		{
			name:     "get_items",
			fileID:   semantic.FileID("catalog/task/get_items.yaml"),
			goldenMD: filepath.FromSlash("../../../docs/uc/001-ec-checkout-flow/renders/catalog/dag-get_items.md"),
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := RenderFile(project, tc.fileID)
			if err != nil {
				t.Fatalf("render dag: %v", err)
			}
			golden.AssertEqualFile(t, tc.goldenMD, actual)
		})
	}
}

func assertLoginResolved(t *testing.T, project *semantic.Project) {
	t.Helper()
	task, ok := project.TasksByQID[semantic.QualifiedID("auth.task.login")]
	if !ok {
		t.Fatalf("auth.task.login not found")
	}
	if len(task.Params) != 1 || task.Params[0].Model != semantic.QualifiedID("auth.model.login_form") {
		t.Fatalf("form param not resolved to auth.model.login_form: %#v", task.Params)
	}
	if task.Returns == nil || task.Returns.Model != semantic.QualifiedID("auth.model.token") {
		t.Fatalf("auth_token return not resolved to auth.model.token: %#v", task.Returns)
	}
	assertStoreRef(t, task.Reads, "user_db", semantic.QualifiedID("auth.store.user_db"), false)
	assertStoreRef(t, task.Reads, "request_context_store", semantic.QualifiedID("auth.store.request_context_store"), false)
	assertStoreRef(t, task.Writes, "session_store", semantic.QualifiedID("auth.store.session_store"), false)
	assertStoreRef(t, task.Writes, "login_log_db", semantic.QualifiedID("auth.task.login.store.login_log_db"), true)
	if _, exists := project.StoresByQID[semantic.QualifiedID("auth.task.login.store.login_log_db")]; exists {
		t.Fatalf("file-private initialized store leaked into StoresByQID")
	}
}

func assertM2Resolved(t *testing.T, project *semantic.Project) {
	t.Helper()
	if _, ok := project.BranchesByQID[semantic.QualifiedID("order/task/process_order.yaml#route_by_inventory")]; !ok {
		t.Fatalf("order/task/process_order.yaml#route_by_inventory not found")
	}
	if _, ok := project.ForksByQID[semantic.QualifiedID("order/task/checkout.yaml#parallel_processing")]; !ok {
		t.Fatalf("order/task/checkout.yaml#parallel_processing not found")
	}
	if _, ok := project.JoinsByQID[semantic.QualifiedID("order/task/checkout.yaml#finalize_checkout")]; !ok {
		t.Fatalf("order/task/checkout.yaml#finalize_checkout not found")
	}
	validateFlow := project.FlowByFile[semantic.FileID("cart/task/validate_cart.yaml")]
	if len(validateFlow) != 1 || validateFlow[0].Kind != semantic.FlowKindForeach {
		t.Fatalf("validate_cart flow not resolved as foreach: %#v", validateFlow)
	}
	checkoutFlow := project.FlowByFile[semantic.FileID("order/task/checkout.yaml")]
	if len(checkoutFlow) != 2 || checkoutFlow[1].Kind != semantic.FlowKindFork {
		t.Fatalf("checkout flow not resolved as step + fork: %#v", checkoutFlow)
	}
	processFlow := project.FlowByFile[semantic.FileID("order/task/process_order.yaml")]
	if len(processFlow) != 2 || processFlow[1].Kind != semantic.FlowKindBranch {
		t.Fatalf("process_order flow not resolved as step + branch: %#v", processFlow)
	}
}

func assertStoreRef(t *testing.T, refs []semantic.StoreRef, name string, qid semantic.QualifiedID, private bool) {
	t.Helper()
	for _, ref := range refs {
		if ref.Name == name {
			if ref.Store != qid || ref.FilePrivate != private {
				t.Fatalf("store ref %s = (%s, private=%v), want (%s, private=%v)", name, ref.Store, ref.FilePrivate, qid, private)
			}
			return
		}
	}
	t.Fatalf("store ref %s not found in %#v", name, refs)
}
