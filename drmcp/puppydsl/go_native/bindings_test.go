package gonative

import (
	"strings"
	"testing"

	puppymcp "github.com/hiroshiasayadev-prog/brewprint/drmcp/puppydsl/go_native/puppygen/drmcp_runtime/mcp"
)

func TestBindingsIncludesMCPServe(t *testing.T) {
	bindings := Bindings()
	if len(bindings) != 1 {
		t.Fatalf("Bindings() count = %d, want 1", len(bindings))
	}

	binding := bindings[0]
	if binding.Function.ID != "drmcp_runtime.mcp.serve" {
		t.Fatalf("binding function ID = %q, want %q", binding.Function.ID, "drmcp_runtime.mcp.serve")
	}
	if binding.BindingID != "mcp.serve" {
		t.Fatalf("binding ID = %q, want %q", binding.BindingID, "mcp.serve")
	}
	if binding.Invoker == nil {
		t.Fatal("binding invoker is nil")
	}
}

func TestBindingsBuildNativeRegistry(t *testing.T) {
	registry, err := NewRegistry()
	if err != nil {
		t.Fatalf("NewRegistry() error = %v", err)
	}
	if registry.Len() != 1 {
		t.Fatalf("Registry.Len() = %d, want 1", registry.Len())
	}

	byFunctionID, exists := registry.ByFunctionID("drmcp_runtime.mcp.serve")
	if !exists {
		t.Fatal("registry does not contain function drmcp_runtime.mcp.serve")
	}
	if byFunctionID.BindingID != "mcp.serve" {
		t.Fatalf("function lookup binding ID = %q, want %q", byFunctionID.BindingID, "mcp.serve")
	}

	byBindingID, exists := registry.ByBindingID("mcp.serve")
	if !exists {
		t.Fatal("registry does not contain binding mcp.serve")
	}
	if byBindingID.Function.ID != "drmcp_runtime.mcp.serve" {
		t.Fatalf(
			"binding lookup function ID = %q, want %q",
			byBindingID.Function.ID,
			"drmcp_runtime.mcp.serve",
		)
	}
}

func TestRegistryInvokesMCPServeThroughGeneratedFacade(t *testing.T) {
	registry, err := NewRegistry()
	if err != nil {
		t.Fatalf("NewRegistry() error = %v", err)
	}
	if err := puppymcp.Serve(registry); err != nil {
		t.Fatalf("puppymcp.Serve(registry) error = %v", err)
	}
}

func TestGeneratedMCPServeInvokerRejectsArguments(t *testing.T) {
	registry, err := NewRegistry()
	if err != nil {
		t.Fatalf("NewRegistry() error = %v", err)
	}
	binding, exists := registry.ByFunctionID("drmcp_runtime.mcp.serve")
	if !exists {
		t.Fatal("registry does not contain function drmcp_runtime.mcp.serve")
	}

	_, err = binding.Invoker(registry, []any{"unexpected"})
	if err == nil {
		t.Fatal("Serve invoker error = nil, want argument-count error")
	}
	if !strings.Contains(err.Error(), "received 1 arguments; expected 0") {
		t.Fatalf("Serve invoker error = %q, want argument-count diagnostic", err)
	}
}
