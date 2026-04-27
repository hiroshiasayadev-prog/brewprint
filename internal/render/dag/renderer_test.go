package dag

import (
	"path/filepath"
	"testing"

	"github.com/hiroshiasayadev-prog/brewprint/internal/resolve"
	"github.com/hiroshiasayadev-prog/brewprint/internal/semantic"
	"github.com/hiroshiasayadev-prog/brewprint/internal/source"
	"github.com/hiroshiasayadev-prog/brewprint/internal/testutil/golden"
)

func TestRenderLoginDAGGolden(t *testing.T) {
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

	actual, err := RenderFile(project, semantic.FileID("auth/task/login.yaml"))
	if err != nil {
		t.Fatalf("render login dag: %v", err)
	}

	golden.AssertEqualFile(t,
		filepath.FromSlash("../../../docs/uc/001-ec-checkout-flow/renders/auth/dag-login.md"),
		actual,
	)
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
