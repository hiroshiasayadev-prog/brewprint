package mcp

import (
	"encoding/json"
	"path/filepath"
	"strings"
	"testing"

	"github.com/hiroshiasayadev-prog/brewprint/internal/query"
	"github.com/hiroshiasayadev-prog/brewprint/internal/resolve"
	"github.com/hiroshiasayadev-prog/brewprint/internal/semantic"
	"github.com/hiroshiasayadev-prog/brewprint/internal/source"
)

func TestServerTools(t *testing.T) {
	server := newUC001Server(t)
	tools := server.Tools()
	want := []string{"list_objects", "get_signature", "get_source", "get_references", "inspect", "list_endpoints"}
	if len(tools) != len(want) {
		t.Fatalf("tools len = %d, want %d: %#v", len(tools), len(want), tools)
	}
	for i, name := range want {
		if tools[i].Name != name {
			t.Fatalf("tool[%d] = %s, want %s", i, tools[i].Name, name)
		}
	}
}

func TestServerCallTool(t *testing.T) {
	server := newUC001Server(t)

	t.Run("list_objects", func(t *testing.T) {
		envelope := call(t, server, "list_objects", `{"object":"node","kind":"task","module":"order"}`)
		if envelope.Error != nil {
			t.Fatalf("list_objects error: %#v", envelope.Error)
		}
		result := resultMap(t, envelope)
		objects := result["objects"].([]any)
		found := false
		for _, item := range objects {
			object := item.(map[string]any)
			if object["id"] != "order.task.checkout" {
				continue
			}
			found = true
			if object["module"] != "order" || object["source"].(map[string]any)["file"] != "order/task/checkout.yaml" {
				t.Fatalf("list_objects checkout object = %#v", object)
			}
		}
		if !found {
			t.Fatalf("order.task.checkout not found in %#v", objects)
		}
	})

	t.Run("get_signature", func(t *testing.T) {
		envelope := call(t, server, "get_signature", `{"selector":{"id":"auth.task.login"}}`)
		if envelope.Error != nil {
			t.Fatalf("get_signature error: %#v", envelope.Error)
		}
		result := resultMap(t, envelope)
		object := result["object"].(map[string]any)
		if object["kind"] != "task" || object["id"] != "auth.task.login" {
			t.Fatalf("get_signature object = %#v", object)
		}
	})

	t.Run("get_signature_m11_private_sub_node", func(t *testing.T) {
		envelope := call(t, server, "get_signature", `{"selector":{"object":"node","file":"order/task/checkout.yaml","local_id":"build_order"}}`)
		if envelope.Error != nil {
			t.Fatalf("get_signature private sub node error: %#v", envelope.Error)
		}
		result := resultMap(t, envelope)
		object := result["object"].(map[string]any)
		if object["id"] != "order/task/checkout.yaml#build_order" || object["kind"] != "task" || object["local_id"] != "build_order" {
			t.Fatalf("private sub node object = %#v", object)
		}
	})

	t.Run("get_signature_m11_asset", func(t *testing.T) {
		envelope := call(t, server, "get_signature", `{"selector":{"object":"asset","id":"order.task.build_order#draft_order"}}`)
		if envelope.Error != nil {
			t.Fatalf("get_signature asset error: %#v", envelope.Error)
		}
		result := resultMap(t, envelope)
		object := result["object"].(map[string]any)
		if object["object"] != "asset" || object["id"] != "order.task.build_order#draft_order" {
			t.Fatalf("asset object = %#v", object)
		}
		signature := result["signature"].(map[string]any)
		if signature["producer"] != "order.task.build_order" || signature["model"] != "order.model.order" || signature["scope_file"] != "order/task/checkout.yaml" {
			t.Fatalf("asset signature = %#v", signature)
		}
	})

	t.Run("get_signature_field", func(t *testing.T) {
		envelope := call(t, server, "get_signature", `{"selector":{"object":"field","id":"order.model.order","local_id":"id"}}`)
		if envelope.Error != nil {
			t.Fatalf("get_signature field error: %#v", envelope.Error)
		}
		result := resultMap(t, envelope)
		object := result["object"].(map[string]any)
		if object["object"] != "field" || object["id"] != "order.model.order.id" {
			t.Fatalf("field object = %#v", object)
		}
		signature := result["signature"].(map[string]any)
		if signature["name"] != "id" || signature["type"] != "str" || signature["pk"] != true {
			t.Fatalf("field signature = %#v", signature)
		}
	})

	t.Run("get_source", func(t *testing.T) {
		envelope := call(t, server, "get_source", `{"selector":{"id":"auth.task.login"}}`)
		if envelope.Error != nil {
			t.Fatalf("get_source error: %#v", envelope.Error)
		}
		result := resultMap(t, envelope)
		object := result["object"].(map[string]any)
		if object["kind"] != "task" || object["id"] != "auth.task.login" {
			t.Fatalf("get_source object = %#v", object)
		}
		source := result["source"].(map[string]any)
		if source["file"] != "auth/task/login.yaml" {
			t.Fatalf("get_source source = %#v", source)
		}
		snippet := result["snippet"].(map[string]any)
		text := snippet["text"].(string)
		if snippet["language"] != "yaml" || !strings.Contains(text, "id: login") || !strings.Contains(text, "type: task") {
			t.Fatalf("get_source snippet = %#v", snippet)
		}
	})

	t.Run("get_references", func(t *testing.T) {
		envelope := call(t, server, "get_references", `{"selector":{"id":"auth.task.login"},"kinds":["reads"]}`)
		if envelope.Error != nil {
			t.Fatalf("get_references error: %#v", envelope.Error)
		}
		result := resultMap(t, envelope)
		refs := result["references"].([]any)
		if len(refs) != 2 {
			t.Fatalf("references len = %d, want 2: %#v", len(refs), refs)
		}
	})

	t.Run("get_references_transition", func(t *testing.T) {
		envelope := call(t, server, "get_references", `{"selector":{"object":"transition","id":"order/state.yaml#processing:payment_webhook_received[payload.status == 'succeeded']"},"direction":"both"}`)
		if envelope.Error != nil {
			t.Fatalf("get_references transition error: %#v", envelope.Error)
		}
		result := resultMap(t, envelope)
		object := result["object"].(map[string]any)
		if object["object"] != "transition" || object["id"] != "order/state.yaml#processing:payment_webhook_received[payload.status == 'succeeded']" {
			t.Fatalf("transition object = %#v", object)
		}
		refs := result["references"].([]any)
		if len(refs) != 5 {
			t.Fatalf("transition references len = %d, want 5: %#v", len(refs), refs)
		}
	})

	t.Run("get_references_m11_asset", func(t *testing.T) {
		envelope := call(t, server, "get_references", `{"selector":{"object":"asset","id":"order.task.build_order#draft_order"},"direction":"out","kinds":["consumes_asset"]}`)
		if envelope.Error != nil {
			t.Fatalf("get_references asset error: %#v", envelope.Error)
		}
		result := resultMap(t, envelope)
		object := result["object"].(map[string]any)
		if object["object"] != "asset" || object["id"] != "order.task.build_order#draft_order" {
			t.Fatalf("asset references object = %#v", object)
		}
		refs := result["references"].([]any)
		if len(refs) != 2 {
			t.Fatalf("asset references len = %d, want 2: %#v", len(refs), refs)
		}
	})

	t.Run("get_references_field", func(t *testing.T) {
		envelope := call(t, server, "get_references", `{"selector":{"object":"field","id":"order.model.order","local_id":"id"},"direction":"both"}`)
		if envelope.Error != nil {
			t.Fatalf("get_references field error: %#v", envelope.Error)
		}
		result := resultMap(t, envelope)
		object := result["object"].(map[string]any)
		if object["object"] != "field" || object["id"] != "order.model.order.id" || object["local_id"] != "id" {
			t.Fatalf("field object = %#v", object)
		}
		refs := result["references"].([]any)
		if len(refs) != 3 {
			t.Fatalf("field references len = %d, want 3: %#v", len(refs), refs)
		}
	})

	t.Run("get_references_file", func(t *testing.T) {
		envelope := call(t, server, "get_references", `{"selector":{"object":"file","kind":"state_file","id":"order/state.yaml"},"direction":"in"}`)
		if envelope.Error != nil {
			t.Fatalf("get_references file error: %#v", envelope.Error)
		}
		result := resultMap(t, envelope)
		object := result["object"].(map[string]any)
		if object["object"] != "file" || object["kind"] != "state_file" || object["id"] != "order/state.yaml" {
			t.Fatalf("file object = %#v", object)
		}
		refs := result["references"].([]any)
		if len(refs) != 2 {
			t.Fatalf("file references len = %d, want 2: %#v", len(refs), refs)
		}
	})

	t.Run("inspect", func(t *testing.T) {
		envelope := call(t, server, "inspect", `{"selector":{"id":"order.task.checkout"}}`)
		if envelope.Error != nil {
			t.Fatalf("inspect error: %#v", envelope.Error)
		}
		result := resultMap(t, envelope)
		members := result["members"].(map[string]any)
		if _, ok := members["flow"]; !ok {
			t.Fatalf("inspect flow missing: %#v", members)
		}
	})

	t.Run("inspect_file", func(t *testing.T) {
		envelope := call(t, server, "inspect", `{"selector":{"object":"file","kind":"state_file","id":"order/state.yaml"}}`)
		if envelope.Error != nil {
			t.Fatalf("inspect file error: %#v", envelope.Error)
		}
		result := resultMap(t, envelope)
		object := result["object"].(map[string]any)
		if object["object"] != "file" || object["kind"] != "state_file" || object["id"] != "order/state.yaml" {
			t.Fatalf("inspect file object = %#v", object)
		}
		members := result["members"].(map[string]any)
		if len(members["states"].([]any)) != 5 || len(members["events"].([]any)) != 4 || len(members["transitions"].([]any)) != 5 {
			t.Fatalf("inspect file members = %#v", members)
		}
	})

	t.Run("inspect_scenario", func(t *testing.T) {
		envelope := call(t, server, "inspect", `{"selector":{"object":"view","kind":"sequence_diagram","id":"checkout_flow"}}`)
		if envelope.Error != nil {
			t.Fatalf("inspect scenario error: %#v", envelope.Error)
		}
		result := resultMap(t, envelope)
		object := result["object"].(map[string]any)
		if object["object"] != "view" || object["kind"] != "sequence_diagram" || object["id"] != "checkout_flow" {
			t.Fatalf("inspect scenario object = %#v", object)
		}
		members := result["members"].(map[string]any)
		steps := members["steps"].([]any)
		if len(steps) != 2 {
			t.Fatalf("inspect scenario steps len = %d, want 2: %#v", len(steps), steps)
		}
	})

	t.Run("inspect_transition", func(t *testing.T) {
		envelope := call(t, server, "inspect", `{"selector":{"object":"transition","id":"order/state.yaml#processing:payment_webhook_received[payload.status == 'succeeded']"}}`)
		if envelope.Error != nil {
			t.Fatalf("inspect transition error: %#v", envelope.Error)
		}
		result := resultMap(t, envelope)
		object := result["object"].(map[string]any)
		if object["object"] != "transition" || object["id"] != "order/state.yaml#processing:payment_webhook_received[payload.status == 'succeeded']" {
			t.Fatalf("inspect transition object = %#v", object)
		}
		signature := result["signature"].(map[string]any)
		if signature["guard"] != "payload.status == 'succeeded'" || signature["action"] != "payment.webhooks.task.process_payment" {
			t.Fatalf("inspect transition signature = %#v", signature)
		}
		members := result["members"].(map[string]any)
		if members["from_state"].(map[string]any)["id"] != "order.state.processing" || members["event"].(map[string]any)["id"] != "order.event.payment_webhook_received" || members["to_state"].(map[string]any)["id"] != "order.state.confirmed" || members["action_task"].(map[string]any)["id"] != "payment.webhooks.task.process_payment" {
			t.Fatalf("inspect transition members = %#v", members)
		}
	})

	t.Run("inspect_m11_private_sub_node", func(t *testing.T) {
		envelope := call(t, server, "inspect", `{"selector":{"object":"node","id":"order/task/checkout.yaml#build_order"}}`)
		if envelope.Error != nil {
			t.Fatalf("inspect private sub node error: %#v", envelope.Error)
		}
		result := resultMap(t, envelope)
		object := result["object"].(map[string]any)
		if object["id"] != "order/task/checkout.yaml#build_order" || object["local_id"] != "build_order" {
			t.Fatalf("inspect private sub node object = %#v", object)
		}
		refs := result["references"].([]any)
		if len(refs) == 0 {
			t.Fatalf("inspect private sub node references empty: %#v", result)
		}
	})

	t.Run("inspect_field", func(t *testing.T) {
		envelope := call(t, server, "inspect", `{"selector":{"object":"field","id":"order.model.order","local_id":"id"}}`)
		if envelope.Error != nil {
			t.Fatalf("inspect field error: %#v", envelope.Error)
		}
		result := resultMap(t, envelope)
		object := result["object"].(map[string]any)
		if object["object"] != "field" || object["id"] != "order.model.order.id" {
			t.Fatalf("inspect field object = %#v", object)
		}
		signature := result["signature"].(map[string]any)
		if signature["name"] != "id" || signature["type"] != "str" || signature["pk"] != true {
			t.Fatalf("inspect field signature = %#v", signature)
		}
		members := result["members"].(map[string]any)
		if members["model"].(map[string]any)["id"] != "order.model.order" || members["type"] != "str" {
			t.Fatalf("inspect field members = %#v", members)
		}
		refs := result["references"].([]any)
		if len(refs) != 3 {
			t.Fatalf("inspect field references len = %d, want 3: %#v", len(refs), refs)
		}
	})

	t.Run("inspect_api_table", func(t *testing.T) {
		envelope := call(t, server, "inspect", `{"selector":{"object":"view","kind":"api_table","id":"ec_api"}}`)
		if envelope.Error != nil {
			t.Fatalf("inspect API table error: %#v", envelope.Error)
		}
		result := resultMap(t, envelope)
		object := result["object"].(map[string]any)
		if object["object"] != "view" || object["kind"] != "api_table" || object["id"] != "ec_api" {
			t.Fatalf("inspect API table object = %#v", object)
		}
		signature := result["signature"].(map[string]any)
		if signature["http_root_path"] != "/api" {
			t.Fatalf("inspect API table signature = %#v", signature)
		}
		members := result["members"].(map[string]any)
		endpoints := members["collected_endpoints"].([]any)
		if len(endpoints) == 0 {
			t.Fatalf("inspect API table endpoints empty: %#v", members)
		}
	})

	t.Run("inspect_er_diagram", func(t *testing.T) {
		envelope := call(t, server, "inspect", `{"selector":{"object":"view","kind":"er_diagram","id":"ec_er"}}`)
		if envelope.Error != nil {
			t.Fatalf("inspect ER diagram error: %#v", envelope.Error)
		}
		result := resultMap(t, envelope)
		object := result["object"].(map[string]any)
		if object["object"] != "view" || object["kind"] != "er_diagram" || object["id"] != "ec_er" {
			t.Fatalf("inspect ER diagram object = %#v", object)
		}
		members := result["members"].(map[string]any)
		models := members["included_models"].([]any)
		relations := members["fk_relations"].([]any)
		if len(models) == 0 || len(relations) == 0 {
			t.Fatalf("inspect ER diagram members = %#v", members)
		}
	})

	t.Run("list_endpoints", func(t *testing.T) {
		envelope := call(t, server, "list_endpoints", `{"api_table_id":"ec_api"}`)
		if envelope.Error != nil {
			t.Fatalf("list_endpoints error: %#v", envelope.Error)
		}
		result := resultMap(t, envelope)
		tables := result["tables"].([]any)
		if len(tables) != 1 {
			t.Fatalf("tables len = %d, want 1: %#v", len(tables), tables)
		}
	})
}

func TestServerCallToolErrors(t *testing.T) {
	server := newUC001Server(t)

	t.Run("unknown_tool", func(t *testing.T) {
		envelope := call(t, server, "missing_tool", `{}`)
		if envelope.Error == nil || envelope.Error.Code != "unknown_tool" {
			t.Fatalf("unknown tool envelope = %#v", envelope)
		}
	})

	t.Run("invalid_args", func(t *testing.T) {
		envelope := call(t, server, "get_signature", `{`)
		if envelope.Error == nil || envelope.Error.Code != "invalid_args" {
			t.Fatalf("invalid args envelope = %#v", envelope)
		}
	})

	t.Run("not_found", func(t *testing.T) {
		envelope := call(t, server, "get_signature", `{"selector":{"id":"auth.task.missing"}}`)
		if envelope.Error == nil || envelope.Error.Code != "not_found" {
			t.Fatalf("not found envelope = %#v", envelope)
		}
	})
}

func call(t *testing.T, server *Server, name string, args string) Envelope {
	t.Helper()
	data := server.CallToolJSON(name, []byte(args))
	var envelope Envelope
	if err := json.Unmarshal(data, &envelope); err != nil {
		t.Fatalf("unmarshal envelope: %v\n%s", err, string(data))
	}
	return envelope
}

func resultMap(t *testing.T, envelope Envelope) map[string]any {
	t.Helper()
	data, err := json.Marshal(envelope.Result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	var out map[string]any
	if err := json.Unmarshal(data, &out); err != nil {
		t.Fatalf("unmarshal result map: %v", err)
	}
	return out
}

func newUC001Server(t *testing.T) *Server {
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
	return NewServer(query.NewService(project))
}
