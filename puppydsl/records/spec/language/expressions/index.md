# Overview: PuppyDSL expressions

- **id**: `spec:puppydsl.language.expressions`
- **status**: draft
- **date**: 2026-07-25
- **parent**: `spec:puppydsl.language`

## What this is

Defines the restricted expression forms permitted in steps, providers, destructors, and typed value positions.

## Current contract

PuppyDSL expressions describe semantic calls and typed value movement. PuppyDSL does not provide a general operator grammar.

The initial expression vocabulary is:

| expression | result |
|---|---|
| Literal | Primitive or enum value. |
| Binding reference | Existing typed value. |
| Function invocation | Declared function output. |
| Binding-origin access chain | Projected record field or required dictionary value. |
| Transparent-record constructor | One exact record value. |
| Dictionary constructor | One exact immutable dictionary value. |

Control flow is not an expression string. `if`, `switch`, and `for` are structured step nodes.

## Restricted surface

The language excludes:

- arithmetic and comparison operators;
- inline boolean composition;
- nested calls in argument or constructor positions;
- list indexing;
- lambdas;
- arbitrary runtime type assertions;
- null and optional literals;
- mutable update expressions.

A meaningful computation that is not one supported expression form must be exposed as a declared function.

## Topics

| title | kind | ref | summary |
|---|---|---|---|
| Invocation | Reference | `spec:puppydsl.language.expressions.invocation` | Function targets, argument modes, literal tokens, binding references, and call validation. |
| Construction and access | Reference | `spec:puppydsl.language.expressions.construction_and_access` | Record and dictionary construction, field projection, and required-key lookup. |

## Related specs

| ref | relation |
|---|---|
| `spec:puppydsl.language.type_system` | Supplies expression input and output types. |
| `spec:puppydsl.language.functions.steps_and_bindings` | Binds expression results to immutable step locals. |
| `spec:puppydsl.language.control_flow` | Defines structured control-flow nodes outside expression strings. |
