# Overview: PuppyDSL functions

- **id**: `spec:puppydsl.language.functions`
- **status**: draft
- **date**: 2026-07-25
- **parent**: `spec:puppydsl.language`

## What this is

Defines typed callable declarations and the common contract exposed by native, steps-defined, and stub functions.

## Current contract

Every function declaration contains:

- one stable function ID;
- one non-empty semantic `detail`;
- optional ordered generic parameters;
- one typed signature;
- optional context `opens` and `requires` declarations;
- exactly one implementation form.

A function signature declares named inputs and exactly one output.

```yaml
functions:
  common.text.is_empty:
    detail: Determines whether a string contains no characters.
    signature: >-
      (value: string) -> boolean
    implementation:
      native: text.is_empty
```

Every signature is authored as a YAML folded scalar using `>-`.

## Signature model

Conceptual grammar:

```text
(<input-name>: <type>, ...) -> <output-type>
```

| subject | contract |
|---|---|
| Input identity | Each input has one unique lower-snake name. |
| Input type | One resolved type expression. |
| Output | Exactly one resolved type expression. |
| Multiple related outputs | Use one declared transparent record. |
| No information result | Use `void`. |
| No continuing result path | Use `never`. |

## Generic functions

A generic function declares ordered type parameters in its declaration key.

```yaml
functions:
  common.collection.compare_sets<T>:
    detail: Compares two typed collections as sets.
    generic:
      T: equatable
    signature: >-
      (
        left: list<T>,
        right: list<T>
      ) -> SetComparison<T>
    implementation:
      native: collection.compare_sets
```

The optional `generic` mapping lists only constrained parameters.

## Context declarations

A function may declare:

| field | meaning |
|---|---|
| `opens` | Ordered context IDs opened for the complete invocation. |
| `requires` | Context-provided values required by the function. |
| `restart_on` | Event IDs that trigger bounded complete-invocation restart. |
| `max_restarts` | Finite restart limit when `restart_on` is present. |

Context and restart validity are defined by their owning language topics.

## Callable identity

Ordinary functions use qualified IDs. The designated entry function uses the reserved unqualified ID `main`.

All implementation forms expose the same signature and context contract. Call syntax does not depend on implementation form.

A leading underscore on the terminal function segment marks an implementation-oriented convention. The core language does not enforce visibility restrictions.

## Non-goals

- Generated caller or binder APIs.
- Native-language package or symbol mapping.
- Multiple returns.
- Overloading by signature.
- Dynamic function construction or repeated string-based dispatch semantics.

## Topics

| title | kind | ref | summary |
|---|---|---|---|
| Implementation forms | Reference | `spec:puppydsl.language.functions.implementation_forms` | Native, steps, stub, `void`, `never`, and executable-profile rules. |
| Steps and bindings | Reference | `spec:puppydsl.language.functions.steps_and_bindings` | Ordered steps, immutable locals, discard labels, returns, and lexical scope. |

## Related specs

| ref | relation |
|---|---|
| `spec:puppydsl.language.source_model.identifiers` | Defines FunctionID and input-name spelling. |
| `spec:puppydsl.language.type_system.type_expressions` | Defines signature type syntax. |
| `spec:puppydsl.language.contexts` | Defines `opens` and `requires`. |
| `spec:puppydsl.language.events` | Defines restart declarations. |
