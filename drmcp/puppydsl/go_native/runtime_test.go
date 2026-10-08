package gonative

import (
	"path/filepath"
	"reflect"
	"testing"

	puppygen "github.com/hiroshiasayadev-prog/brewprint/drmcp/puppydsl/go_native/puppygen"
	puppyruntime "github.com/hiroshiasayadev-prog/brewprint/puppydsl/src/runtime"
)

func TestMainRunsStepsThroughNativeRegistry(t *testing.T) {
	registry, err := NewRegistry()
	if err != nil {
		t.Fatalf("NewRegistry() error = %v", err)
	}
	manager := &integrationContextManager{}
	program, err := puppyruntime.Load(filepath.Join(".."), puppyruntime.Options{
		NativeRegistry:            registry,
		ContextManager:            manager,
		ExpectedSourceFingerprint: puppygen.SourceFingerprint,
	})
	if err != nil {
		t.Fatalf("runtime.Load() error = %v", err)
	}

	if err := puppygen.Main(program.NewExecution()); err != nil {
		t.Fatalf("puppygen.Main() error = %v", err)
	}

	want := []string{
		"open:launch",
		"open:repository_configuration",
		"open:artifact_modules",
		"close:artifact_modules",
		"close:repository_configuration",
		"close:launch",
	}
	if !reflect.DeepEqual(manager.events, want) {
		t.Fatalf("main context events = %#v, want %#v", manager.events, want)
	}
}

type integrationContextManager struct {
	events []string
}

func (manager *integrationContextManager) Open(contextID string) (puppyruntime.ContextScope, error) {
	manager.events = append(manager.events, "open:"+contextID)
	return &integrationContextScope{manager: manager, contextID: contextID}, nil
}

type integrationContextScope struct {
	manager   *integrationContextManager
	contextID string
}

func (scope *integrationContextScope) Close() error {
	scope.manager.events = append(scope.manager.events, "close:"+scope.contextID)
	return nil
}
