package runtime

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/hiroshiasayadev-prog/brewprint/puppydsl/src/native"
	"gopkg.in/yaml.v3"
)

var zeroArgumentCallPattern = regexp.MustCompile(`^([a-z][a-z0-9_]*(?:\.[a-z_][a-z0-9_]*)+)\(\)$`)

// Invoke resolves one declared function and executes its complete structured
// context and implementation boundary.
func (execution *Execution) Invoke(function native.Function, arguments []any) (result any, err error) {
	if execution == nil || execution.program == nil {
		return nil, fmt.Errorf("invoke PuppyDSL function %q through nil execution", function.ID)
	}
	declaration, exists := execution.program.functions[function.ID]
	if !exists {
		return nil, fmt.Errorf("PuppyDSL function %q is not declared", function.ID)
	}
	if declaration.Function.Signature != function.Signature {
		return nil, fmt.Errorf(
			"PuppyDSL function %q signature mismatch: requested %q; declared %q",
			function.ID,
			function.Signature,
			declaration.Function.Signature,
		)
	}
	if _, active := execution.activeFunctions[function.ID]; active {
		return nil, fmt.Errorf("recursive PuppyDSL invocation is prohibited for function %q", function.ID)
	}
	execution.activeFunctions[function.ID] = struct{}{}
	defer delete(execution.activeFunctions, function.ID)

	opened, openErr := execution.openContexts(declaration)
	if openErr != nil {
		return nil, openErr
	}
	defer func() {
		if closeErr := execution.closeContexts(opened); closeErr != nil {
			err = errors.Join(err, closeErr)
		}
	}()

	switch declaration.Implementation {
	case implementationNative:
		return execution.program.nativeRegistry.Dispatch(execution, declaration.Function, arguments)
	case implementationSteps:
		return execution.invokeSteps(declaration, arguments)
	case implementationStub:
		return nil, fmt.Errorf("PuppyDSL function %q is a non-executable stub", function.ID)
	default:
		return nil, fmt.Errorf("PuppyDSL function %q has unknown implementation form %q", function.ID, declaration.Implementation)
	}
}

type openedContext struct {
	id    string
	scope ContextScope
}

func (execution *Execution) openContexts(declaration functionDeclaration) ([]openedContext, error) {
	if len(declaration.Opens) == 0 {
		return nil, nil
	}
	if execution.program.contextManager == nil {
		return nil, fmt.Errorf(
			"PuppyDSL function %q opens contexts but no context manager was supplied",
			declaration.Function.ID,
		)
	}

	opened := make([]openedContext, 0, len(declaration.Opens))
	for _, contextID := range declaration.Opens {
		if _, active := execution.activeContexts[contextID]; active {
			closeErr := execution.closeContexts(opened)
			return nil, errors.Join(
				fmt.Errorf("PuppyDSL function %q reopens active context %q", declaration.Function.ID, contextID),
				closeErr,
			)
		}
		scope, err := execution.program.contextManager.Open(contextID)
		if err != nil {
			closeErr := execution.closeContexts(opened)
			return nil, errors.Join(
				fmt.Errorf("open context %q for PuppyDSL function %q: %w", contextID, declaration.Function.ID, err),
				closeErr,
			)
		}
		if scope == nil {
			closeErr := execution.closeContexts(opened)
			return nil, errors.Join(
				fmt.Errorf("context manager returned nil scope for context %q", contextID),
				closeErr,
			)
		}
		execution.activeContexts[contextID] = struct{}{}
		opened = append(opened, openedContext{id: contextID, scope: scope})
	}
	return opened, nil
}

func (execution *Execution) closeContexts(opened []openedContext) error {
	var closeErr error
	for index := len(opened) - 1; index >= 0; index-- {
		current := opened[index]
		delete(execution.activeContexts, current.id)
		if err := current.scope.Close(); err != nil {
			closeErr = errors.Join(closeErr, fmt.Errorf("close context %q: %w", current.id, err))
		}
	}
	return closeErr
}

func (execution *Execution) invokeSteps(declaration functionDeclaration, arguments []any) (any, error) {
	inputs, output, ok := splitSignature(declaration.Function.Signature)
	if !ok {
		return nil, fmt.Errorf("PuppyDSL function %q has an unsupported signature %q", declaration.Function.ID, declaration.Function.Signature)
	}
	if inputs != "" {
		return nil, fmt.Errorf(
			"initial steps executor supports only zero-input functions; %q declares %q",
			declaration.Function.ID,
			inputs,
		)
	}
	if len(arguments) != 0 {
		return nil, native.ArgumentCountError(declaration.Function, 0, len(arguments))
	}
	if output != "void" {
		return nil, fmt.Errorf(
			"initial steps executor supports only void functions; %q returns %q",
			declaration.Function.ID,
			output,
		)
	}
	if declaration.ReturnBinding != "" {
		return nil, fmt.Errorf(
			"initial steps executor does not support return bindings; %q returns %q",
			declaration.Function.ID,
			declaration.ReturnBinding,
		)
	}

	for _, step := range declaration.Steps {
		if step.Name == "_" || !strings.HasPrefix(step.Name, "_") {
			return nil, fmt.Errorf(
				"initial steps executor requires a named discard step for void calls; function %q step %q",
				declaration.Function.ID,
				step.Name,
			)
		}
		if step.Expression.Kind != yaml.ScalarNode {
			return nil, fmt.Errorf(
				"initial steps executor supports only scalar direct calls; function %q step %q uses YAML node kind %d",
				declaration.Function.ID,
				step.Name,
				step.Expression.Kind,
			)
		}
		targetID, parseErr := parseZeroArgumentCall(step.Expression.Value)
		if parseErr != nil {
			return nil, fmt.Errorf(
				"execute PuppyDSL function %q step %q: %w",
				declaration.Function.ID,
				step.Name,
				parseErr,
			)
		}
		target, exists := execution.program.functions[targetID]
		if !exists {
			return nil, fmt.Errorf(
				"execute PuppyDSL function %q step %q: target function %q is not declared",
				declaration.Function.ID,
				step.Name,
				targetID,
			)
		}
		targetInputs, targetOutput, signatureOK := splitSignature(target.Function.Signature)
		if !signatureOK || targetInputs != "" || targetOutput != "void" {
			return nil, fmt.Errorf(
				"execute PuppyDSL function %q step %q: target %q must have signature () -> void; found %q",
				declaration.Function.ID,
				step.Name,
				targetID,
				target.Function.Signature,
			)
		}
		if _, invokeErr := execution.Invoke(target.Function, nil); invokeErr != nil {
			return nil, fmt.Errorf(
				"execute PuppyDSL function %q step %q: %w",
				declaration.Function.ID,
				step.Name,
				invokeErr,
			)
		}
	}
	return nil, nil
}

func parseZeroArgumentCall(expression string) (string, error) {
	normalized := strings.Join(strings.Fields(expression), " ")
	matches := zeroArgumentCallPattern.FindStringSubmatch(normalized)
	if matches == nil {
		return "", fmt.Errorf("initial invocation parser supports only qualified zero-argument calls; found %q", normalized)
	}
	return matches[1], nil
}

func splitSignature(signature string) (inputs string, output string, ok bool) {
	arrow := strings.LastIndex(signature, ") -> ")
	if !strings.HasPrefix(signature, "(") || arrow < 1 {
		return "", "", false
	}
	inputs = strings.TrimSpace(signature[1:arrow])
	output = strings.TrimSpace(signature[arrow+5:])
	if output == "" {
		return "", "", false
	}
	return inputs, output, true
}
