package runtime

import (
	"errors"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/hiroshiasayadev-prog/brewprint/puppydsl/src/native"
)

func TestExecutionRunsZeroArgumentStepsThroughNativeRegistry(t *testing.T) {
	root := t.TempDir()
	writeTestSource(t, filepath.Join(root, "main.puppy.yaml"), `# responsibility: Defines the test entry.
# excludes: Other functions.

functions:
  main:
    detail: Runs the test native function.
    opens:
      - outer
      - inner
    signature: >-
      () -> void
    implementation:
      steps:
        _served: test.serve()
`)
	writeTestSource(t, filepath.Join(root, "contexts.puppy.yaml"), `# responsibility: Defines test contexts.
# excludes: Context values.

contexts:
  outer:
    detail: Owns the outer test lifetime.
  inner:
    detail: Owns the inner test lifetime.
`)
	writeTestSource(t, filepath.Join(root, "native.puppy.yaml"), `# responsibility: Defines the test native function.
# excludes: Entry composition.

functions:
  test.serve:
    detail: Records one native invocation.
    signature: >-
      () -> void
    implementation:
      native: test.serve
`)

	var events []string
	binding := native.NewBinding(
		native.Function{ID: "test.serve", Signature: "() -> void"},
		"test.serve",
		func(call native.Call, arguments []any) (any, error) {
			events = append(events, "invoke:test.serve")
			if _, ok := call.(*Execution); !ok {
				t.Fatalf("native invoker call = %T, want *runtime.Execution", call)
			}
			if len(arguments) != 0 {
				t.Fatalf("native invoker arguments = %#v, want none", arguments)
			}
			return nil, nil
		},
	)
	registry, err := native.NewRegistry([]native.Binding{binding})
	if err != nil {
		t.Fatalf("native.NewRegistry() error = %v", err)
	}
	manager := &recordingContextManager{events: &events}
	program, err := Load(root, Options{
		NativeRegistry: registry,
		ContextManager: manager,
	})
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	execution := program.NewExecution()
	_, err = execution.Invoke(native.Function{ID: "main", Signature: "() -> void"}, nil)
	if err != nil {
		t.Fatalf("Execution.Invoke(main) error = %v", err)
	}

	want := []string{
		"open:outer",
		"open:inner",
		"invoke:test.serve",
		"close:inner",
		"close:outer",
	}
	if !reflect.DeepEqual(events, want) {
		t.Fatalf("execution events = %#v, want %#v", events, want)
	}
}

func TestLoadRejectsFingerprintMismatch(t *testing.T) {
	root := t.TempDir()
	writeTestSource(t, filepath.Join(root, "main.puppy.yaml"), `# responsibility: Defines the test entry.
# excludes: Other functions.

functions:
  main:
    detail: Completes without information.
    signature: >-
      () -> void
    implementation:
      steps:
        _served: test.serve()
`)
	writeTestSource(t, filepath.Join(root, "native.puppy.yaml"), `# responsibility: Defines the test native function.
# excludes: Entry composition.

functions:
  test.serve:
    detail: Completes without information.
    signature: >-
      () -> void
    implementation:
      native: test.serve
`)

	registry, err := native.NewRegistry([]native.Binding{
		native.NewBinding(
			native.Function{ID: "test.serve", Signature: "() -> void"},
			"test.serve",
			func(native.Call, []any) (any, error) { return nil, nil },
		),
	})
	if err != nil {
		t.Fatalf("native.NewRegistry() error = %v", err)
	}
	program, err := Load(root, Options{NativeRegistry: registry})
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}
	if program.SourceFingerprint() == "" {
		t.Fatal("SourceFingerprint() is empty")
	}

	_, err = Load(root, Options{
		NativeRegistry:            registry,
		ExpectedSourceFingerprint: "stale",
	})
	if err == nil || !strings.Contains(err.Error(), "fingerprint mismatch") {
		t.Fatalf("Load() error = %v, want fingerprint mismatch", err)
	}
}

func TestExecutionClosesContextsAfterNativeFailure(t *testing.T) {
	root := t.TempDir()
	writeTestSource(t, filepath.Join(root, "main.puppy.yaml"), `# responsibility: Defines the test entry.
# excludes: Other functions.

functions:
  main:
    detail: Runs the failing native function.
    opens:
      - process
    signature: >-
      () -> void
    implementation:
      steps:
        _failed: test.fail()
`)
	writeTestSource(t, filepath.Join(root, "contexts.puppy.yaml"), `# responsibility: Defines the test context.
# excludes: Context values.

contexts:
  process:
    detail: Owns the test lifetime.
`)
	writeTestSource(t, filepath.Join(root, "native.puppy.yaml"), `# responsibility: Defines the failing native function.
# excludes: Entry composition.

functions:
  test.fail:
    detail: Reports one native failure.
    signature: >-
      () -> void
    implementation:
      native: test.fail
`)

	expected := errors.New("native failed")
	registry, err := native.NewRegistry([]native.Binding{
		native.NewBinding(
			native.Function{ID: "test.fail", Signature: "() -> void"},
			"test.fail",
			func(native.Call, []any) (any, error) { return nil, expected },
		),
	})
	if err != nil {
		t.Fatalf("native.NewRegistry() error = %v", err)
	}
	var events []string
	program, err := Load(root, Options{
		NativeRegistry: registry,
		ContextManager: &recordingContextManager{events: &events},
	})
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	_, err = program.NewExecution().Invoke(native.Function{ID: "main", Signature: "() -> void"}, nil)
	if !errors.Is(err, expected) {
		t.Fatalf("Execution.Invoke(main) error = %v, want %v", err, expected)
	}
	want := []string{"open:process", "close:process"}
	if !reflect.DeepEqual(events, want) {
		t.Fatalf("execution events = %#v, want %#v", events, want)
	}
}

type recordingContextManager struct {
	events *[]string
}

func (manager *recordingContextManager) Open(contextID string) (ContextScope, error) {
	*manager.events = append(*manager.events, "open:"+contextID)
	return &recordingContextScope{id: contextID, events: manager.events}, nil
}

type recordingContextScope struct {
	id     string
	events *[]string
}

func (scope *recordingContextScope) Close() error {
	*scope.events = append(*scope.events, "close:"+scope.id)
	return nil
}

func writeTestSource(t *testing.T, path, content string) {
	t.Helper()
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("os.WriteFile(%s) error = %v", path, err)
	}
}
