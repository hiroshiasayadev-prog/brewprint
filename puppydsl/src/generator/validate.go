package generator

import (
	"fmt"
	"sort"
	"strings"
)

var primitiveGoTypes = map[string]string{
	"string":  "string",
	"boolean": "bool",
	"integer": "int64",
	"float":   "float64",
}

func validateProgram(loaded program) error {
	types := make(map[string]typeDeclaration, len(loaded.Types))
	for _, declaration := range loaded.Types {
		if goExportedName(declaration.ID) != declaration.ID {
			return fmt.Errorf("%s: type %s is not directly Go-mappable PascalCase", declaration.Source, declaration.ID)
		}
		types[declaration.ID] = declaration
	}

	for _, declaration := range loaded.Types {
		switch declaration.Kind {
		case typeKindScalar:
			if _, supported := primitiveGoTypes[declaration.Scalar]; !supported {
				return fmt.Errorf("%s: scalar %s uses unsupported representation %s", declaration.Source, declaration.ID, declaration.Scalar)
			}
		case typeKindEnum:
			if len(declaration.Members) == 0 {
				return fmt.Errorf("%s: enum %s has no members", declaration.Source, declaration.ID)
			}
			seen := map[string]string{}
			for _, member := range declaration.Members {
				goName := declaration.ID + goExportedName(member)
				if previous, exists := seen[goName]; exists {
					return fmt.Errorf("%s: enum %s members %s and %s collide as %s", declaration.Source, declaration.ID, previous, member, goName)
				}
				seen[goName] = member
			}
		case typeKindRecord:
			seen := map[string]string{}
			for _, field := range declaration.Fields {
				goName := goExportedName(field.Name)
				if previous, exists := seen[goName]; exists {
					return fmt.Errorf("%s: record %s fields %s and %s collide as %s", declaration.Source, declaration.ID, previous, field.Name, goName)
				}
				seen[goName] = field.Name
				if err := validateTypeExpression(field.Type, types, false); err != nil {
					return fmt.Errorf("%s: record %s field %s: %w", declaration.Source, declaration.ID, field.Name, err)
				}
			}
		case typeKindUnion:
			if len(declaration.Variants) == 0 {
				return fmt.Errorf("%s: union %s has no variants", declaration.Source, declaration.ID)
			}
			seen := map[string]struct{}{}
			for _, variant := range declaration.Variants {
				if _, exists := seen[variant]; exists {
					return fmt.Errorf("%s: union %s repeats variant %s", declaration.Source, declaration.ID, variant)
				}
				seen[variant] = struct{}{}
				variantType, exists := types[variant]
				if !exists {
					return fmt.Errorf("%s: union %s references unknown variant %s", declaration.Source, declaration.ID, variant)
				}
				if variantType.Kind == typeKindUnion {
					return fmt.Errorf("%s: union %s cannot directly contain union variant %s", declaration.Source, declaration.ID, variant)
				}
			}
		case typeKindOpaque:
			// Opaque values have no generated design-facing fields.
		default:
			return fmt.Errorf("%s: type %s has unsupported kind %s", declaration.Source, declaration.ID, declaration.Kind)
		}
	}

	packageSymbols := map[string]map[string]string{}
	bindingIDs := map[string]string{}
	for _, declaration := range loaded.Functions {
		packagePath := strings.Join(declaration.Namespace, "/")
		if len(declaration.Namespace) > 0 {
			packageName := declaration.Namespace[len(declaration.Namespace)-1]
			if _, keyword := goKeywords[packageName]; keyword {
				return fmt.Errorf("%s: function namespace %s maps to Go keyword package %s", declaration.Source, packagePath, packageName)
			}
		}
		symbols := packageSymbols[packagePath]
		if symbols == nil {
			symbols = map[string]string{}
			packageSymbols[packagePath] = symbols
		}
		generatedSymbols := []string{declaration.GoName, goLocalName(declaration.GoName) + "Function"}
		if declaration.Implementation == implementationNative {
			generatedSymbols = append(generatedSymbols, "Bind"+declaration.GoName, declaration.GoName+"Implementation")
		}
		for _, symbol := range generatedSymbols {
			if previous, exists := symbols[symbol]; exists {
				return fmt.Errorf("functions %s and %s collide as generated Go symbol %s in package %s", previous, declaration.ID, symbol, packagePath)
			}
			symbols[symbol] = declaration.ID
		}

		inputNames := map[string]string{}
		for _, input := range declaration.Signature.Inputs {
			goName := goLocalName(input.Name)
			if _, reserved := generatedFunctionReservedNames[goName]; reserved {
				return fmt.Errorf("%s: function %s input %s collides with generated Go name %s", declaration.Source, declaration.ID, input.Name, goName)
			}
			if previous, exists := inputNames[goName]; exists {
				return fmt.Errorf("%s: function %s inputs %s and %s collide as %s", declaration.Source, declaration.ID, previous, input.Name, goName)
			}
			inputNames[goName] = input.Name
			if err := validateTypeExpression(input.Type, types, false); err != nil {
				return fmt.Errorf("%s: function %s input %s: %w", declaration.Source, declaration.ID, input.Name, err)
			}
		}
		if err := validateTypeExpression(declaration.Signature.Output, types, true); err != nil {
			return fmt.Errorf("%s: function %s output: %w", declaration.Source, declaration.ID, err)
		}
		if declaration.Implementation == implementationNative {
			if previous, exists := bindingIDs[declaration.NativeBindingID]; exists {
				return fmt.Errorf("native binding ID %s is used by both %s and %s", declaration.NativeBindingID, previous, declaration.ID)
			}
			bindingIDs[declaration.NativeBindingID] = declaration.ID
		}
	}
	return nil
}

func validateTypeExpression(expression typeExpression, types map[string]typeDeclaration, allowControlResult bool) error {
	if _, primitive := primitiveGoTypes[expression.Name]; primitive {
		if len(expression.Arguments) != 0 {
			return fmt.Errorf("primitive %s cannot have type arguments", expression.Name)
		}
		return nil
	}
	if expression.Name == "void" || expression.Name == "never" {
		if !allowControlResult {
			return fmt.Errorf("%s is valid only as a function output", expression.Name)
		}
		if len(expression.Arguments) != 0 {
			return fmt.Errorf("%s cannot have type arguments", expression.Name)
		}
		return nil
	}

	switch expression.Name {
	case "list", "optional":
		if len(expression.Arguments) != 1 {
			return fmt.Errorf("%s requires exactly one type argument", expression.Name)
		}
		return validateTypeExpression(expression.Arguments[0], types, false)
	case "dict":
		if len(expression.Arguments) != 2 {
			return fmt.Errorf("dict requires exactly two type arguments")
		}
		if err := validateDictionaryKeyType(expression.Arguments[0], types); err != nil {
			return fmt.Errorf("dict key: %w", err)
		}
		return validateTypeExpression(expression.Arguments[1], types, false)
	}

	if len(expression.Arguments) != 0 {
		return fmt.Errorf("generic named type %s is not implemented", expression.Name)
	}
	if _, exists := types[expression.Name]; !exists {
		return fmt.Errorf("unknown type %s", expression.Name)
	}
	return nil
}

func validateDictionaryKeyType(expression typeExpression, types map[string]typeDeclaration) error {
	if len(expression.Arguments) != 0 {
		return fmt.Errorf("%s is not a supported dictionary key type", expression.Name)
	}
	switch expression.Name {
	case "string", "integer", "boolean":
		return nil
	}
	declaration, exists := types[expression.Name]
	if !exists {
		return fmt.Errorf("unknown type %s", expression.Name)
	}
	if declaration.Kind == typeKindEnum {
		return nil
	}
	if declaration.Kind != typeKindScalar {
		return fmt.Errorf("%s is not a primitive-backed named scalar or enum", expression.Name)
	}
	switch declaration.Scalar {
	case "string", "integer", "boolean":
		return nil
	default:
		return fmt.Errorf("%s uses unsupported dictionary-key scalar %s", expression.Name, declaration.Scalar)
	}
}

func goExportedName(identifier string) string {
	parts := strings.Split(identifier, "_")
	var builder strings.Builder
	for _, part := range parts {
		if part == "" {
			continue
		}
		builder.WriteString(strings.ToUpper(part[:1]))
		builder.WriteString(part[1:])
	}
	return builder.String()
}

func goLocalName(identifier string) string {
	exported := goExportedName(identifier)
	if exported == "" {
		return "value"
	}
	name := strings.ToLower(exported[:1]) + exported[1:]
	if _, keyword := goKeywords[name]; keyword {
		return name + "Value"
	}
	return name
}

var generatedFunctionReservedNames = map[string]struct{}{
	"call": {}, "native": {}, "types": {}, "value": {}, "err": {}, "zero": {},
}

var goKeywords = map[string]struct{}{
	"break": {}, "default": {}, "func": {}, "interface": {}, "select": {},
	"case": {}, "defer": {}, "go": {}, "map": {}, "struct": {},
	"chan": {}, "else": {}, "goto": {}, "package": {}, "switch": {},
	"const": {}, "fallthrough": {}, "if": {}, "range": {}, "type": {},
	"continue": {}, "for": {}, "import": {}, "return": {}, "var": {},
}

func sortedUnionMembership(types []typeDeclaration) map[string][]string {
	membership := map[string][]string{}
	for _, declaration := range types {
		if declaration.Kind != typeKindUnion {
			continue
		}
		for _, variant := range declaration.Variants {
			membership[variant] = append(membership[variant], declaration.ID)
		}
	}
	for variant := range membership {
		sort.Strings(membership[variant])
	}
	return membership
}
