package native

import "fmt"

// Function identifies one generated PuppyDSL callable contract.
type Function struct {
	ID        string
	Signature string
}

// Call invokes one already-declared PuppyDSL function through the runtime.
type Call interface {
	Invoke(function Function, arguments []any) (any, error)
}

// Invoker is the uniform runtime entry point generated around one typed native
// implementation.
type Invoker func(call Call, arguments []any) (any, error)

// Binding associates one generated PuppyDSL native declaration with its uniform
// runtime invoker.
type Binding struct {
	Function  Function
	BindingID string
	Invoker   Invoker
}

// NewBinding constructs one native binding descriptor.
func NewBinding(function Function, bindingID string, invoker Invoker) Binding {
	return Binding{
		Function:  function,
		BindingID: bindingID,
		Invoker:   invoker,
	}
}

// ArgumentCountError reports a generated native invocation with the wrong number
// of runtime arguments.
func ArgumentCountError(function Function, expected, actual int) error {
	return fmt.Errorf(
		"puppydsl function %q received %d arguments; expected %d",
		function.ID,
		actual,
		expected,
	)
}

// Input converts one runtime argument to the concrete Go input type declared by
// a generated native implementation.
func Input[T any](function Function, arguments []any, index int, name string) (T, error) {
	var zero T
	if index < 0 || index >= len(arguments) {
		return zero, fmt.Errorf(
			"puppydsl function %q has no argument %q at index %d",
			function.ID,
			name,
			index,
		)
	}

	value, ok := arguments[index].(T)
	if !ok {
		return zero, fmt.Errorf(
			"puppydsl function %q argument %q at index %d has Go type %T; expected generated input type",
			function.ID,
			name,
			index,
			arguments[index],
		)
	}
	return value, nil
}

// Output converts one runtime result to the generated facade's declared Go type.
func Output[T any](function Function, value any) (T, error) {
	result, ok := value.(T)
	if ok {
		return result, nil
	}

	var zero T
	return zero, fmt.Errorf(
		"puppydsl function %q returned %T; expected generated output type",
		function.ID,
		value,
	)
}
