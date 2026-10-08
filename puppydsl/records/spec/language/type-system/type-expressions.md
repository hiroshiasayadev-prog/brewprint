# Reference: PuppyDSL type expressions

- **id**: `spec:puppydsl.language.type_system.type_expressions`
- **status**: draft
- **date**: 2026-07-25
- **parent**: `spec:puppydsl.language.type_system`

## What this is

Defines the type-reference syntax used by declarations, signatures, fields, generic arguments, context values, and event payloads.

## Type expression forms

| form | examples | meaning |
|---|---|---|
| Primitive | `string`, `boolean`, `integer` | Built-in scalar type. |
| Control result | `void`, `never` | Function completion type. |
| Declared type | `ArtifactKind`, `SourceLocation` | One nominal declared type. |
| Type parameter | `T` | Parameter bound by the surrounding generic declaration. |
| List | `list<T>` | Ordered immutable collection. |
| Optional | `optional<T>` | Present `T` or absence. |
| Dictionary | `dict<K, V>` | Immutable unique-key association. |
| Generic declared type | `SetComparison<T>` | One concrete instantiation of a generic declared type. |

Conceptual grammar:

```text
TypeExpression := Primitive
                | ControlResult
                | TypeID
                | TypeParameter
                | list<TypeExpression>
                | optional<TypeExpression>
                | dict<TypeExpression, TypeExpression>
                | TypeID<TypeExpression, ...>
```

Whitespace may appear around punctuation. Type identity does not depend on formatting whitespace.

## Generic declaration parameters

A generic type or function declares ordered parameters in its declaration key.

```yaml
types:
  Pair<T, U>:
    detail: Preserves two typed values.
    record:
      first: T
      second: U
```

A `generic` mapping lists only constrained parameters.

```yaml
generic:
  T: equatable
```

The initial constraint vocabulary contains only `equatable`. Unconstrained parameters are omitted from `generic`.

## Dictionary key type

`dict<K, V>` accepts these key types:

| key category | valid |
|---|---|
| `string`, `integer`, `boolean` | yes |
| Named scalar backed by a valid primitive | yes |
| Declared enum | yes |
| List, optional, dictionary, record, opaque, union | no |
| `void`, `never` | no |

Key eligibility is a built-in dictionary rule. It is not a user-authored generic constraint.

## `void` placement

`void` is valid as:

- a function output;
- an event payload that carries no data;
- a generic type argument when the enclosing composite remains information-bearing.

The initial language contract does not assign additional special conversion or absence semantics to `void`.

## `never` placement

`never` is valid only as a function output.

`never` is prohibited as:

- a function input;
- a record field;
- a union variant;
- a generic type argument;
- a context-provided value;
- an event payload.

## Validation rules

- Reject unknown declared type IDs.
- Reject unknown type parameters.
- Reject a generic instantiation with the wrong argument count.
- Reject an unsupported generic constraint.
- Reject a type argument that does not satisfy its declared constraint.
- Reject an invalid dictionary key type.
- Reject `never` outside a function output.
- Reject malformed nesting, missing delimiters, or trailing unparsed tokens.

## Related specs

| ref | relation |
|---|---|
| `spec:puppydsl.language.type_system` | Parent type-system overview. |
| `spec:puppydsl.language.type_system.declared_types` | Defines the declared type forms named by type expressions. |
| `spec:puppydsl.language.type_system.assignability` | Defines compatibility between resolved type expressions. |
