# Overview: PuppyDSL type system

- **id**: `spec:puppydsl.language.type_system`
- **status**: draft
- **date**: 2026-07-25
- **parent**: `spec:puppydsl.language`

## What this is

Defines PuppyDSL runtime values and static type identity. The type system owns primitive types, declared nominal types, generic composites, type expressions, and assignment compatibility.

## Current contract

PuppyDSL is statically typed. Declared types are nominal unless a rule explicitly states otherwise.

The built-in primitive types are:

```text
string
boolean
integer
```

The built-in control-result types are:

```text
void
never
```

The built-in generic composites are:

```text
list<T>
optional<T>
dict<K, V>
```

Declared type forms are:

```text
named scalar
enum
transparent record
opaque type
union
```

There is no `any`, implicit structural typing, class inheritance, runtime reflection, or general null value.

## Type categories

| category | value semantics | design-facing inspection |
|---|---|---|
| Primitive | Built-in scalar value. | Used directly. |
| Named scalar | Nominal value backed by one primitive. | No fields. |
| Enum | One member from a closed declaration. | Exhaustive switch only for member branching. |
| Transparent record | Nominal value with a declared field set. | Declared fields may be projected. |
| Opaque type | Nominal value whose representation is hidden. | No field access or construction. |
| Union | One value from a closed declared variant set. | Exhaustive switch only for variant elimination. |
| List | Ordered immutable collection. | Bounded iteration through control flow. |
| Optional | Present `T` or absence. | Exhaustive `present` / `absent` switch. |
| Dictionary | Immutable unique-key association. | Required-key lookup and bounded iteration. |

## Control-result types

`void` represents normal completion without an information-bearing result. A direct `void` result creates no local binding.

`never` represents a call path that cannot continue normally. `never` is valid only as a function output and cannot exist as an ordinary value.

## Non-goals

- Native-language representation mapping.
- Generated wrappers or code-generation contracts.
- Mutable collections.
- Pointer identity, address identity, or arbitrary reference graphs.
- Implicit conversion between nominal and primitive values.

## Topics

| title | kind | ref | summary |
|---|---|---|---|
| Type expressions | Reference | `spec:puppydsl.language.type_system.type_expressions` | Grammar and placement rules for primitive, declared, generic, and parameterized type references. |
| Declared types | Reference | `spec:puppydsl.language.type_system.declared_types` | Named scalar, enum, transparent record, opaque, union, generic, and recursive declaration rules. |
| Assignability | Reference | `spec:puppydsl.language.type_system.assignability` | Exact identity, variant-to-union assignment, generic constraints, and prohibited conversions. |

## Related specs

| ref | relation |
|---|---|
| `spec:puppydsl.language` | Parent language overview. |
| `spec:puppydsl.language.expressions` | Constructs and accesses typed values. |
| `spec:puppydsl.language.program_validity` | Resolves and validates type use across a program. |
