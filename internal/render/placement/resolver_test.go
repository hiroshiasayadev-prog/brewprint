package placement

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/hiroshiasayadev-prog/brewprint/internal/rawyaml"
	"github.com/hiroshiasayadev-prog/brewprint/internal/resolve"
	"github.com/hiroshiasayadev-prog/brewprint/internal/semantic"
	"github.com/hiroshiasayadev-prog/brewprint/internal/source"
)

func TestResolverUC001(t *testing.T) {
	rawProject, semanticProject := loadUC001(t)
	resolver, diagnostics := NewResolver(rawProject, semanticProject)
	assertNoErrors(t, diagnostics)

	assertGroup(t, resolver, semantic.FileID("auth/task/login.yaml"), "auth")
	assertGroup(t, resolver, semantic.FileID("cart/task/validate_cart.yaml"), "commerce")
	assertGroup(t, resolver, semantic.FileID("order/task/checkout.yaml"), "commerce")
	assertGroup(t, resolver, semantic.FileID("payment/webhooks/task/process_payment.yaml"), "commerce")
	assertGroup(t, resolver, semantic.FileID("catalog/task/get_items.yaml"), "catalog")
	assertGroup(t, resolver, semantic.FileID("inventory/state.yaml"), "catalog")

	assertDAGPath(t, resolver, semanticProject, "auth.task.login", "auth/dag-login.md")
	assertDAGPath(t, resolver, semanticProject, "cart.task.validate_cart", "commerce/dag-validate_cart.md")
	assertDAGPath(t, resolver, semanticProject, "order.task.checkout", "commerce/dag-checkout.md")
	assertDAGPath(t, resolver, semanticProject, "order.task.process_order", "commerce/dag-process_order.md")

	master := resolver.MasterIndexMarkdown("UC-001", semanticProject)
	for _, want := range []string{"[認証](auth/index.md)", "[商取引](commerce/index.md)", "[カタログ](catalog/index.md)", "*(cross)*", "*(preview)*"} {
		if !strings.Contains(master, want) {
			t.Fatalf("master index missing %q:\n%s", want, master)
		}
	}

	commerce := resolver.GroupIndexMarkdown("commerce", semanticProject)
	for _, want := range []string{"dag-validate_cart.md", "dag-checkout.md", "dag-process_order.md"} {
		if !strings.Contains(commerce, want) {
			t.Fatalf("commerce index missing %q:\n%s", want, commerce)
		}
	}
}

func TestResolverValidation(t *testing.T) {
	project := semantic.NewProject()
	project.NodesByQID[semantic.QualifiedID("auth.task.login")] = &semantic.Task{BaseNode: semantic.BaseNode{QID: "auth.task.login", FileID: "auth/task/login.yaml", ID: "login", Kind: semantic.NodeKindTask, Main: true}}
	project.NodesByQID[semantic.QualifiedID("cart.task.validate_cart")] = &semantic.Task{BaseNode: semantic.BaseNode{QID: "cart.task.validate_cart", FileID: "cart/task/validate_cart.yaml", ID: "validate_cart", Kind: semantic.NodeKindTask, Main: true}}

	raw := &rawyaml.Project{Files: []rawyaml.File{{
		ID:   "render_index.yaml",
		Kind: rawyaml.FileKindRenderIndex,
		RenderIndex: &rawyaml.RenderIndex{Groups: []rawyaml.RenderGroup{
			{ID: "_bad", Modules: []string{"auth"}},
			{ID: "Bad", Modules: []string{"auth"}},
			{ID: "commerce", Modules: []string{"cart/sub"}},
		}},
	}}}

	_, diagnostics := NewResolver(raw, project)
	assertHasDiagnostic(t, diagnostics, SeverityError, "must not start with underscore")
	assertHasDiagnostic(t, diagnostics, SeverityError, "must match [a-z0-9_]+")
	assertHasDiagnostic(t, diagnostics, SeverityError, "belongs to multiple groups")
	assertHasDiagnostic(t, diagnostics, SeverityError, "top-level module")
}

func loadUC001(t *testing.T) (*rawyaml.Project, *semantic.Project) {
	t.Helper()
	yamlRoot := filepath.FromSlash("../../../docs/uc/001-ec-checkout-flow/yaml")
	loader := source.Loader{}
	rawProject, err := loader.Load(yamlRoot)
	if err != nil {
		t.Fatalf("load yaml root: %v", err)
	}
	semanticProject, diagnostics := resolve.Build(rawProject)
	for _, diagnostic := range diagnostics {
		if diagnostic.Severity == semantic.SeverityError {
			t.Fatalf("semantic diagnostic: %s: %s", diagnostic.FileID, diagnostic.Message)
		}
	}
	return rawProject, semanticProject
}

func assertGroup(t *testing.T, resolver *Resolver, fileID semantic.FileID, want string) {
	t.Helper()
	got, ok := resolver.GroupForFile(fileID)
	if !ok || got != want {
		t.Fatalf("GroupForFile(%s) = (%s, %v), want (%s, true)", fileID, got, ok, want)
	}
}

func assertDAGPath(t *testing.T, resolver *Resolver, project *semantic.Project, qid string, want string) {
	t.Helper()
	task := project.TasksByQID[semantic.QualifiedID(qid)]
	if task == nil {
		t.Fatalf("task not found: %s", qid)
	}
	got, err := resolver.DAGPath(task)
	if err != nil {
		t.Fatalf("DAGPath(%s): %v", qid, err)
	}
	if got != want {
		t.Fatalf("DAGPath(%s) = %s, want %s", qid, got, want)
	}
}

func assertNoErrors(t *testing.T, diagnostics []Diagnostic) {
	t.Helper()
	for _, diagnostic := range diagnostics {
		if diagnostic.Severity == SeverityError {
			t.Fatalf("unexpected error diagnostic: %s", diagnostic.Message)
		}
	}
}

func assertHasDiagnostic(t *testing.T, diagnostics []Diagnostic, severity Severity, contains string) {
	t.Helper()
	for _, diagnostic := range diagnostics {
		if diagnostic.Severity == severity && strings.Contains(diagnostic.Message, contains) {
			return
		}
	}
	t.Fatalf("diagnostic not found severity=%s contains=%q in %#v", severity, contains, diagnostics)
}
