package native

import (
	"errors"
	"strings"
	"testing"
)

func TestNewRegistryIndexesBindings(t *testing.T) {
	binding := NewBinding(
		Function{ID: "runtime.test.run", Signature: "() -> string"},
		"test.run",
		func(Call, []any) (any, error) { return "done", nil },
	)

	registry, err := NewRegistry([]Binding{binding})
	if err != nil {
		t.Fatalf("NewRegistry() error = %v", err)
	}
	if registry.Len() != 1 {
		t.Fatalf("Registry.Len() = %d, want 1", registry.Len())
	}

	byFunctionID, exists := registry.ByFunctionID("runtime.test.run")
	if !exists {
		t.Fatal("ByFunctionID() did not find runtime.test.run")
	}
	if byFunctionID.BindingID != "test.run" {
		t.Fatalf("ByFunctionID() binding ID = %q, want %q", byFunctionID.BindingID, "test.run")
	}

	byBindingID, exists := registry.ByBindingID("test.run")
	if !exists {
		t.Fatal("ByBindingID() did not find test.run")
	}
	if byBindingID.Function.ID != "runtime.test.run" {
		t.Fatalf("ByBindingID() function ID = %q, want %q", byBindingID.Function.ID, "runtime.test.run")
	}

	if _, exists := registry.ByFunctionID("runtime.test.missing"); exists {
		t.Fatal("ByFunctionID() found an unknown function")
	}
	if _, exists := registry.ByBindingID("test.missing"); exists {
		t.Fatal("ByBindingID() found an unknown binding ID")
	}
}

func TestRegistryInvokeDispatchesUniformInvoker(t *testing.T) {
	function := Function{ID: "runtime.test.echo", Signature: "(value: string) -> string"}
	var receivedCall Call
	binding := NewBinding(
		function,
		"test.echo",
		func(call Call, arguments []any) (any, error) {
			receivedCall = call
			if len(arguments) != 1 {
				t.Fatalf("invoker argument count = %d, want 1", len(arguments))
			}
			return arguments[0], nil
		},
	)
	registry, err := NewRegistry([]Binding{binding})
	if err != nil {
		t.Fatalf("NewRegistry() error = %v", err)
	}

	result, err := registry.Invoke(function, []any{"hello"})
	if err != nil {
		t.Fatalf("Registry.Invoke() error = %v", err)
	}
	if result != "hello" {
		t.Fatalf("Registry.Invoke() result = %#v, want %q", result, "hello")
	}
	if receivedCall != registry {
		t.Fatalf("invoker call = %#v, want registry", receivedCall)
	}

	external := &recordingCall{}
	result, err = registry.Dispatch(external, function, []any{"again"})
	if err != nil {
		t.Fatalf("Registry.Dispatch() error = %v", err)
	}
	if result != "again" {
		t.Fatalf("Registry.Dispatch() result = %#v, want %q", result, "again")
	}
	if receivedCall != external {
		t.Fatalf("dispatch invoker call = %#v, want external call", receivedCall)
	}
}

func TestRegistryInvokeRejectsUnavailableOrStaleFunction(t *testing.T) {
	function := Function{ID: "runtime.test.run", Signature: "() -> void"}
	registry, err := NewRegistry([]Binding{
		NewBinding(function, "test.run", func(Call, []any) (any, error) { return nil, nil }),
	})
	if err != nil {
		t.Fatalf("NewRegistry() error = %v", err)
	}

	tests := []struct {
		name     string
		function Function
		contains string
	}{
		{
			name:     "missing binding",
			function: Function{ID: "runtime.test.missing", Signature: "() -> void"},
			contains: "has no native binding",
		},
		{
			name:     "signature mismatch",
			function: Function{ID: function.ID, Signature: "() -> string"},
			contains: "signature mismatch",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, invokeErr := registry.Invoke(test.function, nil)
			if invokeErr == nil {
				t.Fatalf("Registry.Invoke() result = %#v, want error", result)
			}
			if !strings.Contains(invokeErr.Error(), test.contains) {
				t.Fatalf("Registry.Invoke() error = %q, want substring %q", invokeErr, test.contains)
			}
		})
	}
}

func TestRegistryInvokePropagatesInvokerError(t *testing.T) {
	expected := errors.New("native failed")
	function := Function{ID: "runtime.test.fail", Signature: "() -> void"}
	registry, err := NewRegistry([]Binding{
		NewBinding(function, "test.fail", func(Call, []any) (any, error) { return nil, expected }),
	})
	if err != nil {
		t.Fatalf("NewRegistry() error = %v", err)
	}

	_, err = registry.Invoke(function, nil)
	if !errors.Is(err, expected) {
		t.Fatalf("Registry.Invoke() error = %v, want %v", err, expected)
	}
}

func TestNewRegistryRejectsInvalidBindings(t *testing.T) {
	invoker := Invoker(func(Call, []any) (any, error) { return nil, nil })
	validFunction := Function{ID: "runtime.test.run", Signature: "() -> void"}

	tests := []struct {
		name     string
		bindings []Binding
		contains string
	}{
		{
			name: "empty function ID",
			bindings: []Binding{
				NewBinding(Function{}, "test.run", invoker),
			},
			contains: "empty function ID",
		},
		{
			name: "function ID surrounding whitespace",
			bindings: []Binding{
				NewBinding(Function{ID: " runtime.test.run ", Signature: "() -> void"}, "test.run", invoker),
			},
			contains: "function ID",
		},
		{
			name: "empty signature",
			bindings: []Binding{
				NewBinding(Function{ID: "runtime.test.run"}, "test.run", invoker),
			},
			contains: "empty signature",
		},
		{
			name: "signature surrounding whitespace",
			bindings: []Binding{
				NewBinding(Function{ID: "runtime.test.run", Signature: " () -> void "}, "test.run", invoker),
			},
			contains: "signature",
		},
		{
			name: "empty binding ID",
			bindings: []Binding{
				NewBinding(validFunction, "", invoker),
			},
			contains: "empty binding ID",
		},
		{
			name: "binding ID surrounding whitespace",
			bindings: []Binding{
				NewBinding(validFunction, " test.run ", invoker),
			},
			contains: "binding ID",
		},
		{
			name: "nil invoker",
			bindings: []Binding{
				NewBinding(validFunction, "test.run", nil),
			},
			contains: "nil invoker",
		},
		{
			name: "duplicate function ID",
			bindings: []Binding{
				NewBinding(validFunction, "test.run.first", invoker),
				NewBinding(validFunction, "test.run.second", invoker),
			},
			contains: "duplicate native binding for function",
		},
		{
			name: "duplicate binding ID",
			bindings: []Binding{
				NewBinding(validFunction, "test.run", invoker),
				NewBinding(
					Function{ID: "runtime.test.run_again", Signature: "() -> void"},
					"test.run",
					invoker,
				),
			},
			contains: "duplicate native binding ID",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			registry, err := NewRegistry(test.bindings)
			if err == nil {
				t.Fatalf("NewRegistry() registry = %#v, want error", registry)
			}
			if !strings.Contains(err.Error(), test.contains) {
				t.Fatalf("NewRegistry() error = %q, want substring %q", err, test.contains)
			}
		})
	}
}

type recordingCall struct{}

func (*recordingCall) Invoke(Function, []any) (any, error) {
	return nil, nil
}

func TestNilRegistryOperationsAreSafe(t *testing.T) {
	var registry *Registry
	if registry.Len() != 0 {
		t.Fatalf("nil Registry.Len() = %d, want 0", registry.Len())
	}
	if _, exists := registry.ByFunctionID("runtime.test.run"); exists {
		t.Fatal("nil Registry.ByFunctionID() found a binding")
	}
	if _, exists := registry.ByBindingID("test.run"); exists {
		t.Fatal("nil Registry.ByBindingID() found a binding")
	}
	if _, err := registry.Invoke(Function{ID: "runtime.test.run"}, nil); err == nil {
		t.Fatal("nil Registry.Invoke() error = nil, want error")
	}
}
