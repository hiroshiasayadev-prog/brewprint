# Reference: PuppyDSL identifiers

- **id**: `spec:puppydsl.language.source_model.identifiers`
- **status**: draft
- **date**: 2026-07-25
- **parent**: `spec:puppydsl.language.source_model`

## What this is

Defines identifier spelling and qualified identity for PuppyDSL declarations and lexical bindings.

## Identifier forms

| form | grammar | uses |
|---|---|---|
| Type ID | `[A-Z][A-Za-z0-9]*` | Declared types. |
| Lower-snake ID | `[a-z][a-z0-9]*(?:_[a-z0-9]+)*` | Namespace segments, ordinary function terminals, contexts, events, fields, parameters, provided values, locals, and enum members. |
| Implementation-oriented function terminal | `_[a-z][a-z0-9]*(?:_[a-z0-9]+)*` | Function terminal segments that expose an implementation-oriented convention. |

A qualified function or event ID uses one or more lower-snake namespace segments followed by one terminal segment.

```text
<namespace>(.<namespace>)*.<terminal>
```

Examples:

```text
common.text.trim
repository.changed
operation.validation._compose_checks
```

The designated entry function is the reserved unqualified ID `main`.

## Rules

| subject | rule |
|---|---|
| Hyphen | Prohibited in declaration and lexical identifiers. |
| Type parameter | Uses the Type ID form, usually one uppercase identifier such as `T`. |
| Generic declaration key | Appends `<T, U>` to the declared type or function identity. |
| Qualified enum member | Uses `<EnumType>.<member>`. |
| File stem | Application-owned `.puppy.yaml` stems use lower-snake spelling. |
| Directory segment | Application-owned source directory segments use lower-snake spelling. |
| Declaration-kind directory | Organizes files but does not add a function namespace segment. |

The identifier restrictions do not apply to quoted strings, filesystem paths carried as data, or ordinary prose.

## Reserved identities

| identity | reservation |
|---|---|
| `main` | Reserved unqualified callable ID for the designated entry function. The reservation does not create a general keyword in unrelated identifier positions. |
| `present` | Optional-switch case word. It is not an optional value or enum member by language definition. |
| `absent` | Optional-switch case word. It is not an optional value or enum member by language definition. |
| `self` | Destructor argument binding for the provided value being destroyed. |

## Validation rules

- Reject identifiers that violate their required form.
- Reject duplicate identities within one applicable registry or lexical scope.
- Reject an ordinary function ID that conflicts with a generated event callable ID.
- Reject `main` when declared outside `main.puppy.yaml`.
- Reject source file stems or application-owned directory segments that violate lower-snake spelling.

## Related specs

| ref | relation |
|---|---|
| `spec:puppydsl.language.source_model` | Parent source model. |
| `spec:puppydsl.language.functions` | Function identity and generic declaration use. |
| `spec:puppydsl.language.events` | Event identity and generated callable namespace. |
