package mcp

import (
	"encoding/json"
	"path/filepath"
	"testing"

	"github.com/hiroshiasayadev-prog/brewprint/internal/query"
	"github.com/hiroshiasayadev-prog/brewprint/internal/resolve"
	"github.com/hiroshiasayadev-prog/brewprint/internal/semantic"
	"github.com/hiroshiasayadev-prog/brewprint/internal/source"
)

func TestServerTools(t *testing.T) {
	server := newUC001Server(t)
	tools := server.Tools()
	want := []string{"get_signature", "get_references", "inspect", "list_endpoints"}
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
