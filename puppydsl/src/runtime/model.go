package runtime

import (
	"github.com/hiroshiasayadev-prog/brewprint/puppydsl/src/native"
	"gopkg.in/yaml.v3"
)

// ContextManager opens one declared PuppyDSL context instance.
type ContextManager interface {
	Open(contextID string) (ContextScope, error)
}

// ContextScope closes one context instance opened by ContextManager.
type ContextScope interface {
	Close() error
}

// Options supplies the native and context boundaries required to activate one
// PuppyDSL program.
type Options struct {
	NativeRegistry            *native.Registry
	ContextManager            ContextManager
	ExpectedSourceFingerprint string
}

// Program is one loaded and bootstrap-validated PuppyDSL source set.
type Program struct {
	root              string
	sourceFingerprint string
	functions         map[string]functionDeclaration
	contexts          map[string]struct{}
	nativeRegistry    *native.Registry
	contextManager    ContextManager
}

type implementationKind string

const (
	implementationNative implementationKind = "native"
	implementationSteps  implementationKind = "steps"
	implementationStub   implementationKind = "stub"
)

type functionDeclaration struct {
	Function        native.Function
	Detail          string
	Opens           []string
	Implementation  implementationKind
	NativeBindingID string
	Steps           []stepDeclaration
	ReturnBinding   string
	Source          string
}

type stepDeclaration struct {
	Name       string
	Expression *yaml.Node
}

// SourceFingerprint returns the exact sorted .puppy.yaml source-set fingerprint
// loaded into this program.
func (program *Program) SourceFingerprint() string {
	if program == nil {
		return ""
	}
	return program.sourceFingerprint
}

// NewExecution creates one non-concurrent invocation state for the program.
func (program *Program) NewExecution() *Execution {
	return &Execution{
		program:         program,
		activeContexts:  map[string]struct{}{},
		activeFunctions: map[string]struct{}{},
	}
}

// Execution owns one active function/context chain and implements native.Call.
type Execution struct {
	program         *Program
	activeContexts  map[string]struct{}
	activeFunctions map[string]struct{}
}

var _ native.Call = (*Execution)(nil)
