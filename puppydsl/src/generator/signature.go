package generator

import (
	"fmt"
	"strings"
	"unicode"
)

func parseFunctionSignature(source string) (functionSignature, error) {
	compact := strings.Map(func(character rune) rune {
		if unicode.IsSpace(character) {
			return -1
		}
		return character
	}, source)

	parser := signatureParser{text: compact}
	if err := parser.expect('('); err != nil {
		return functionSignature{}, err
	}

	var inputs []functionInput
	if !parser.consume(')') {
		for {
			name, err := parser.identifier()
			if err != nil {
				return functionSignature{}, fmt.Errorf("parse input name: %w", err)
			}
			if err := parser.expect(':'); err != nil {
				return functionSignature{}, fmt.Errorf("parse input %q: %w", name, err)
			}
			typeValue, err := parser.typeExpression()
			if err != nil {
				return functionSignature{}, fmt.Errorf("parse input %q type: %w", name, err)
			}
			inputs = append(inputs, functionInput{Name: name, Type: typeValue})

			if parser.consume(')') {
				break
			}
			if err := parser.expect(','); err != nil {
				return functionSignature{}, err
			}
		}
	}

	if !parser.consumeString("->") {
		return functionSignature{}, parser.errorf("expected ->")
	}
	output, err := parser.typeExpression()
	if err != nil {
		return functionSignature{}, fmt.Errorf("parse output type: %w", err)
	}
	if parser.position != len(parser.text) {
		return functionSignature{}, parser.errorf("unexpected trailing input")
	}

	seen := make(map[string]struct{}, len(inputs))
	for _, input := range inputs {
		if _, exists := seen[input.Name]; exists {
			return functionSignature{}, fmt.Errorf("duplicate signature input %q", input.Name)
		}
		seen[input.Name] = struct{}{}
	}

	return functionSignature{Inputs: inputs, Output: output}, nil
}

func parseTypeExpressionSource(source string) (typeExpression, error) {
	compact := strings.Map(func(character rune) rune {
		if unicode.IsSpace(character) {
			return -1
		}
		return character
	}, source)
	parser := signatureParser{text: compact}
	expression, err := parser.typeExpression()
	if err != nil {
		return typeExpression{}, err
	}
	if parser.position != len(parser.text) {
		return typeExpression{}, parser.errorf("unexpected trailing input")
	}
	return expression, nil
}

type signatureParser struct {
	text     string
	position int
}

func (parser *signatureParser) typeExpression() (typeExpression, error) {
	name, err := parser.identifier()
	if err != nil {
		return typeExpression{}, err
	}

	expression := typeExpression{Name: name}
	if !parser.consume('<') {
		return expression, nil
	}

	for {
		argument, parseErr := parser.typeExpression()
		if parseErr != nil {
			return typeExpression{}, parseErr
		}
		expression.Arguments = append(expression.Arguments, argument)

		if parser.consume('>') {
			break
		}
		if err := parser.expect(','); err != nil {
			return typeExpression{}, err
		}
	}
	return expression, nil
}

func (parser *signatureParser) identifier() (string, error) {
	start := parser.position
	for parser.position < len(parser.text) {
		character := parser.text[parser.position]
		if !isIdentifierCharacter(character) {
			break
		}
		parser.position++
	}
	if start == parser.position {
		return "", parser.errorf("expected identifier")
	}
	return parser.text[start:parser.position], nil
}

func isIdentifierCharacter(character byte) bool {
	return character >= 'a' && character <= 'z' ||
		character >= 'A' && character <= 'Z' ||
		character >= '0' && character <= '9' ||
		character == '_'
}

func (parser *signatureParser) expect(expected byte) error {
	if parser.consume(expected) {
		return nil
	}
	return parser.errorf("expected %q", expected)
}

func (parser *signatureParser) consume(expected byte) bool {
	if parser.position >= len(parser.text) || parser.text[parser.position] != expected {
		return false
	}
	parser.position++
	return true
}

func (parser *signatureParser) consumeString(expected string) bool {
	if !strings.HasPrefix(parser.text[parser.position:], expected) {
		return false
	}
	parser.position += len(expected)
	return true
}

func (parser *signatureParser) errorf(format string, arguments ...any) error {
	return fmt.Errorf("%s at offset %d in %q", fmt.Sprintf(format, arguments...), parser.position, parser.text)
}
