# Reference: PuppyDSL invocation

- **id**: `spec:puppydsl.language.expressions.invocation`
- **status**: draft
- **date**: 2026-07-25
- **parent**: `spec:puppydsl.language.expressions`

## What this is

Defines direct function-call syntax, argument binding, literal tokens, and lexical binding references used by PuppyDSL expressions.

## Invocation form

A direct call uses one declared function ID followed by a parenthesized argument list.

```yaml
normalized: common.text.trim(body)
```

A long call may use a YAML folded scalar.

```yaml
comparison: >-
  common.collection.compare_sets(
    left: declared_values,
    right: observed_values
  )
```

The YAML parser produces one scalar string. A dedicated invocation parser interprets the call target and argument tokens.

The current contract fixes call shape and token meaning. Exact lexer rules for escaping, whitespace normalization, and parser safety limits remain pending.

## Argument modes

One call uses either positional or named arguments.

| mode | example | binding rule |
|---|---|---|
| Positional | `common.text.trim(body)` | Source order maps to signature input order. |
| Named | `example.compare(left: expected, right: actual)` | Each argument name maps to one signature input. |
| Mixed | `example.compare(expected, right: actual)` | Invalid. |

Named arguments must supply every required input exactly once. Unknown, missing, or duplicate arguments are invalid.

## Invocation tokens

| token form | static meaning |
|---|---|
| Single-quoted or double-quoted token | `string` literal. |
| Unquoted decimal integer | `integer` literal. |
| Exact `true` or `false` | `boolean` literal. |
| `<EnumType>.<member>` | One declared enum value. |
| Unquoted identifier or qualified identifier | Visible binding reference. |

A quoted token is always a string literal. Quoted text does not resolve as a binding, integer, boolean, or enum member.

## Binding references

A binding reference resolves from the active lexical environment.

Possible sources are:

- function inputs;
- context-provided values declared through `requires`;
- ordinary locals bound by earlier steps;
- active ancestor-scope bindings;
- current iteration bindings;
- a switch target under case-local narrowing.

A discard step label is not a binding and cannot appear as an argument.

## Optional and null syntax

The invocation grammar contains no optional constructor or untyped null literal.

These spellings are invalid:

```text
null
Null
NULL
~
none
some(value)
```

An `optional<T>` value originates from a declared typed boundary and is eliminated through exhaustive switch.

## Nested expression restriction

A call argument is exactly one literal or one visible binding reference.

These argument forms are invalid:

- nested function call;
- field projection;
- dictionary lookup;
- record or dictionary constructor;
- arithmetic, comparison, or boolean expression.

The caller binds a required intermediate value in an earlier step.

## Type binding

Each argument must be assignable to its declared signature input.

Generic calls infer one exact substitution from the supplied argument types. Every occurrence of one type parameter uses the same substitution.

## Validation rules

| invalid condition | required result |
|---|---|
| Unknown function ID | Reject the invocation. |
| Mixed positional and named arguments | Reject the invocation. |
| Wrong positional count | Reject the invocation. |
| Missing, unknown, or duplicate named argument | Reject the invocation. |
| Unknown binding reference | Reject the invocation. |
| Unknown enum type or member | Reject the invocation. |
| Argument type is not assignable | Reject the invocation. |
| Generic substitution is inconsistent or violates a constraint | Reject the invocation. |
| Nested expression appears as an argument | Reject the invocation. |
| Null or optional-literal spelling appears | Reject the invocation. |

Argument compatibility is evaluated only after the callable, binding, enum member, and required types resolve.

## Related specs

| ref | relation |
|---|---|
| `spec:puppydsl.language.expressions` | Parent expression overview. |
| `spec:puppydsl.language.type_system.assignability` | Defines argument compatibility. |
| `spec:puppydsl.language.functions` | Defines function signatures and callable identity. |
| `spec:puppydsl.language.functions.steps_and_bindings` | Defines lexical binding visibility. |
