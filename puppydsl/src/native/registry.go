package native

import (
	"fmt"
	"strings"
)

// Registry indexes validated native bindings by both PuppyDSL function ID and
// native binding ID.
type Registry struct {
	byFunctionID map[string]Binding
	byBindingID  map[string]Binding
}

// NewRegistry validates and indexes one complete native binding set.
func NewRegistry(bindings []Binding) (*Registry, error) {
	registry := &Registry{
		byFunctionID: make(map[string]Binding, len(bindings)),
		byBindingID:  make(map[string]Binding, len(bindings)),
	}

	for index, binding := range bindings {
		functionID := binding.Function.ID
		trimmedFunctionID := strings.TrimSpace(functionID)
		if trimmedFunctionID == "" {
			return nil, fmt.Errorf("native binding at index %d has an empty function ID", index)
		}
		if trimmedFunctionID != functionID {
			return nil, fmt.Errorf("native binding at index %d has function ID %q with surrounding whitespace", index, functionID)
		}

		signature := binding.Function.Signature
		trimmedSignature := strings.TrimSpace(signature)
		if trimmedSignature == "" {
			return nil, fmt.Errorf("native binding for function %q has an empty signature", functionID)
		}
		if trimmedSignature != signature {
			return nil, fmt.Errorf("native binding for function %q has signature %q with surrounding whitespace", functionID, signature)
		}

		bindingID := binding.BindingID
		trimmedBindingID := strings.TrimSpace(bindingID)
		if trimmedBindingID == "" {
			return nil, fmt.Errorf("native binding for function %q has an empty binding ID", functionID)
		}
		if trimmedBindingID != bindingID {
			return nil, fmt.Errorf("native binding for function %q has binding ID %q with surrounding whitespace", functionID, bindingID)
		}
		if binding.Invoker == nil {
			return nil, fmt.Errorf("native binding %q for function %q has a nil invoker", bindingID, functionID)
		}

		if previous, exists := registry.byFunctionID[functionID]; exists {
			return nil, fmt.Errorf(
				"duplicate native binding for function %q: binding IDs %q and %q",
				functionID,
				previous.BindingID,
				bindingID,
			)
		}
		if previous, exists := registry.byBindingID[bindingID]; exists {
			return nil, fmt.Errorf(
				"duplicate native binding ID %q: functions %q and %q",
				bindingID,
				previous.Function.ID,
				functionID,
			)
		}

		registry.byFunctionID[functionID] = binding
		registry.byBindingID[bindingID] = binding
	}

	return registry, nil
}

// Len returns the number of registered native bindings.
func (registry *Registry) Len() int {
	if registry == nil {
		return 0
	}
	return len(registry.byFunctionID)
}

// ByFunctionID returns the native binding registered for one PuppyDSL function.
func (registry *Registry) ByFunctionID(functionID string) (Binding, bool) {
	if registry == nil {
		return Binding{}, false
	}
	binding, exists := registry.byFunctionID[functionID]
	return binding, exists
}

// ByBindingID returns the native binding registered under one native binding ID.
func (registry *Registry) ByBindingID(bindingID string) (Binding, bool) {
	if registry == nil {
		return Binding{}, false
	}
	binding, exists := registry.byBindingID[bindingID]
	return binding, exists
}

// Invoke resolves one generated function descriptor and dispatches it through
// the binding's uniform invoker. Registry satisfies Call for native-only call
// chains.
func (registry *Registry) Invoke(function Function, arguments []any) (any, error) {
	return registry.Dispatch(registry, function, arguments)
}

// Dispatch invokes one registered native binding while preserving the supplied
// mixed-runtime Call boundary for nested generated facade calls.
func (registry *Registry) Dispatch(call Call, function Function, arguments []any) (any, error) {
	if registry == nil {
		return nil, fmt.Errorf("invoke PuppyDSL function %q through nil native registry", function.ID)
	}
	if call == nil {
		return nil, fmt.Errorf("invoke PuppyDSL function %q with nil runtime call boundary", function.ID)
	}

	binding, exists := registry.byFunctionID[function.ID]
	if !exists {
		return nil, fmt.Errorf("PuppyDSL function %q has no native binding", function.ID)
	}
	if binding.Function.Signature != function.Signature {
		return nil, fmt.Errorf(
			"PuppyDSL function %q signature mismatch: requested %q; registered %q",
			function.ID,
			function.Signature,
			binding.Function.Signature,
		)
	}
	return binding.Invoker(call, arguments)
}
