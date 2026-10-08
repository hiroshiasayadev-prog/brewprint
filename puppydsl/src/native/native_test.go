package native

import (
	"strings"
	"testing"
)

func TestArgumentCountError(t *testing.T) {
	err := ArgumentCountError(Function{ID: "runtime.test.run"}, 2, 1)
	for _, expected := range []string{"runtime.test.run", "received 1 arguments", "expected 2"} {
		if !strings.Contains(err.Error(), expected) {
			t.Fatalf("ArgumentCountError() = %q, want substring %q", err, expected)
		}
	}
}

func TestInputConvertsGeneratedArgumentType(t *testing.T) {
	function := Function{ID: "runtime.test.echo"}
	value, err := Input[string](function, []any{"hello"}, 0, "value")
	if err != nil {
		t.Fatalf("Input[string]() error = %v", err)
	}
	if value != "hello" {
		t.Fatalf("Input[string]() = %q, want %q", value, "hello")
	}

	if _, err := Input[int64](function, []any{"hello"}, 0, "value"); err == nil {
		t.Fatal("Input[int64]() error = nil, want type error")
	}
	if _, err := Input[string](function, nil, 0, "value"); err == nil {
		t.Fatal("Input[string]() error = nil, want missing argument error")
	}
}

func TestOutputConvertsGeneratedResultType(t *testing.T) {
	function := Function{ID: "runtime.test.echo"}
	value, err := Output[string](function, "hello")
	if err != nil {
		t.Fatalf("Output[string]() error = %v", err)
	}
	if value != "hello" {
		t.Fatalf("Output[string]() = %q, want %q", value, "hello")
	}
	if _, err := Output[int64](function, "hello"); err == nil {
		t.Fatal("Output[int64]() error = nil, want type error")
	}
}
