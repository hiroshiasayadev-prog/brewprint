package query

import (
	"path/filepath"
	"testing"

	"github.com/hiroshiasayadev-prog/brewprint/internal/resolve"
	"github.com/hiroshiasayadev-prog/brewprint/internal/semantic"
	"github.com/hiroshiasayadev-prog/brewprint/internal/source"
)

func TestQueryServiceUC001(t *testing.T) {
	service := newUC001Service(t)

	t.Run("GetSignature", func(t *testing.T) {
		login, err := service.GetSignature(GetSignatureRequest{Selector: Selector{ID: "auth.task.login"}})
		if err != nil {
			t.Fatalf("GetSignature login: %v", err)
		}
		if login.Object.Kind != "task" {
			t.Fatalf("login kind = %s, want task", login.Object.Kind)
		}
		if login.Signature["main"] != true {
			t.Fatalf("login main = %#v, want true", login.Signature["main"])
		}
		params := login.Signature["params"].([]ParamSignature)
		if len(params) != 1 || params[0].Model != "auth.model.login_form" {
			t.Fatalf("login params = %#v", params)
		}
		ret := login.Signature["returns"].(*ReturnSignature)
		if ret.Name != "auth_token" || ret.Model != "auth.model.token" || ret.Asset == nil {
			t.Fatalf("login returns = %#v", ret)
		}
		endpoint := login.Signature["endpoint"].(EndpointSignature)
		if endpoint.Method != "POST" || endpoint.LeafPath != "login" {
			t.Fatalf("login endpoint = %#v", endpoint)
		}

		model, err := service.GetSignature(GetSignatureRequest{Selector: Selector{ID: "auth.model.login_form"}})
		if err != nil {
			t.Fatalf("GetSignature model: %v", err)
		}
		fields := model.Signature["fields"].([]FieldSignature)
		if len(fields) == 0 {
			t.Fatalf("login_form fields is empty")
		}

		store, err := service.GetSignature(GetSignatureRequest{Selector: Selector{ID: "auth.store.user_db"}})
		if err != nil {
			t.Fatalf("GetSignature store: %v", err)
		}
		if store.Signature["store_kind"] != "db" || store.Signature["of"] != "auth.model.credential" {
			t.Fatalf("user_db signature = %#v", store.Signature)
		}

		join, err := service.GetSignature(GetSignatureRequest{Selector: Selector{ID: "order.join.finalize_checkout"}})
		if err != nil {
			t.Fatalf("GetSignature join: %v", err)
		}
		joinParams := join.Signature["params"].([]ParamSignature)
		if len(joinParams) != 2 {
			t.Fatalf("join params = %#v", joinParams)
		}
	})

	t.Run("GetReferences", func(t *testing.T) {
		loginRefs, err := service.GetReferences(GetReferencesRequest{Selector: Selector{ID: "auth.task.login"}})
		if err != nil {
			t.Fatalf("GetReferences login: %v", err)
		}
		assertHasReference(t, loginRefs.References, "param_model", "out", "auth.task.login", "auth.model.login_form")
		assertHasReference(t, loginRefs.References, "return_model", "out", "auth.task.login", "auth.model.token")
		assertHasReference(t, loginRefs.References, "produces_asset", "out", "auth.task.login", "")
		assertHasReference(t, loginRefs.References, "reads", "out", "auth.task.login", "auth.store.user_db")
		assertHasReference(t, loginRefs.References, "writes", "out", "auth.task.login", "auth.store.session_store")

		storeRefs, err := service.GetReferences(GetReferencesRequest{Selector: Selector{ID: "auth.store.user_db"}, Direction: "in"})
		if err != nil {
			t.Fatalf("GetReferences store: %v", err)
		}
		assertHasReference(t, storeRefs.References, "reads", "in", "auth.task.login", "auth.store.user_db")

		modelRefs, err := service.GetReferences(GetReferencesRequest{Selector: Selector{ID: "auth.model.login_form"}, Direction: "in"})
		if err != nil {
			t.Fatalf("GetReferences model: %v", err)
		}
		assertHasReference(t, modelRefs.References, "param_model", "in", "auth.task.login", "auth.model.login_form")

		readRefs, err := service.GetReferences(GetReferencesRequest{Selector: Selector{ID: "auth.task.login"}, Kinds: []string{"reads"}})
		if err != nil {
			t.Fatalf("GetReferences reads filter: %v", err)
		}
		if len(readRefs.References) != 2 {
			t.Fatalf("reads filtered refs len = %d, want 2: %#v", len(readRefs.References), readRefs.References)
		}
		for _, ref := range readRefs.References {
			if ref.Kind != "reads" {
				t.Fatalf("non-read ref in reads filter: %#v", ref)
			}
		}
	})

	t.Run("ListEndpoints", func(t *testing.T) {
		endpoints, err := service.ListEndpoints(ListEndpointsRequest{APITableID: "ec_api"})
		if err != nil {
			t.Fatalf("ListEndpoints ec_api: %v", err)
		}
		if len(endpoints.Tables) != 1 {
			t.Fatalf("endpoint tables len = %d, want 1: %#v", len(endpoints.Tables), endpoints.Tables)
		}
		table := endpoints.Tables[0]
		if table.ID != "ec_api" || table.HTTPRootPath != "/api" {
			t.Fatalf("endpoint table = %#v", table)
		}
		if len(table.Sections) != 5 {
			t.Fatalf("endpoint sections len = %d, want 5: %#v", len(table.Sections), table.Sections)
		}
		assertHasEndpoint(t, table.Sections, "auth", "auth.task.login", "POST", "/api/login")
		assertHasEndpoint(t, table.Sections, "catalog", "catalog.task.get_items", "GET", "/api/catalog_items")
		assertHasEndpoint(t, table.Sections, "payment.webhooks", "payment.webhooks.task.process_payment", "POST", "/api/stripe")
	})

	t.Run("Inspect", func(t *testing.T) {
		login, err := service.Inspect(InspectRequest{Selector: Selector{ID: "auth.task.login"}})
		if err != nil {
			t.Fatalf("Inspect login: %v", err)
		}
		assets := login.Members["assets"].([]AssetRef)
		if len(assets) != 1 || assets[0].Name != "auth_token" {
			t.Fatalf("login assets = %#v", assets)
		}
		if len(login.References) == 0 {
			t.Fatalf("login references is empty")
		}

		checkout, err := service.Inspect(InspectRequest{Selector: Selector{ID: "order.task.checkout"}})
		if err != nil {
			t.Fatalf("Inspect checkout: %v", err)
		}
		if _, ok := checkout.Members["flow"]; !ok {
			t.Fatalf("checkout flow member missing: %#v", checkout.Members)
		}
		if _, ok := checkout.Members["sub_tasks"]; !ok {
			t.Fatalf("checkout sub_tasks member missing: %#v", checkout.Members)
		}

		model, err := service.Inspect(InspectRequest{Selector: Selector{ID: "auth.model.login_form"}})
		if err != nil {
			t.Fatalf("Inspect model: %v", err)
		}
		if _, ok := model.Members["fields"]; !ok {
			t.Fatalf("model fields member missing: %#v", model.Members)
		}

		store, err := service.Inspect(InspectRequest{Selector: Selector{ID: "auth.store.user_db"}})
		if err != nil {
			t.Fatalf("Inspect store: %v", err)
		}
		if _, ok := store.Members["model"]; !ok {
			t.Fatalf("store model member missing: %#v", store.Members)
		}
		assertHasReference(t, store.References, "reads", "in", "auth.task.login", "auth.store.user_db")
	})
}

func newUC001Service(t *testing.T) *Service {
	t.Helper()
	yamlRoot := filepath.FromSlash("../../docs/uc/001-ec-checkout-flow/yaml")
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
	return NewService(project)
}

func assertHasEndpoint(t *testing.T, sections []EndpointSection, module, task, method, path string) {
	t.Helper()
	for _, section := range sections {
		if section.Module != module {
			continue
		}
		for _, endpoint := range section.Endpoints {
			if endpoint.Task == task && endpoint.Method == method && endpoint.Path == path {
				return
			}
		}
	}
	t.Fatalf("endpoint not found module=%s task=%s method=%s path=%s in %#v", module, task, method, path, sections)
}

func assertHasReference(t *testing.T, refs []Reference, kind, direction, fromID, toID string) {
	t.Helper()
	for _, ref := range refs {
		if ref.Kind != kind || ref.Direction != direction {
			continue
		}
		if fromID != "" && ref.From.ID != fromID {
			continue
		}
		if toID != "" && ref.To.ID != toID {
			continue
		}
		return
	}
	t.Fatalf("reference not found kind=%s direction=%s from=%s to=%s in %#v", kind, direction, fromID, toID, refs)
}
