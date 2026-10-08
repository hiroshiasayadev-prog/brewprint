# Reference: PuppyDSL assignability

- **id**: `spec:puppydsl.language.type_system.assignability`
- **status**: draft
- **date**: 2026-07-25
- **parent**: `spec:puppydsl.language.type_system`

## What this is

Defines when a typed PuppyDSL value may be supplied to a declared target type.

## Current contract

Exact type identity is required unless this specification declares one explicit compatibility rule.

The only implicit compatibility rule is direct variant-to-union assignment.

| source type | target type | result |
|---|---|---|
| `T` | same `T` | valid |
| Direct variant `V` | Union `U` that directly lists `V` | valid |
| Union `U` | Variant `V` | invalid outside exhaustive narrowing |
| Variant of `U1` | Unrelated union `U2` | invalid unless `U2` directly lists that exact variant |
| Named scalar | Backing primitive | invalid |
| Backing primitive | Named scalar | invalid |
| Enum | `string` | invalid |
| `optional<T>` | `T` | invalid outside `present` narrowing |
| `T` | `optional<T>` | invalid without a declared producer boundary |

Variant-to-union assignment preserves the supplied variant as the union value's active variant.

## Assignment positions

The same assignability rule applies to:

- function arguments;
- function returns;
- transparent-record constructor fields;
- dictionary constructor keys and values;
- context provider arguments;
- event payload arguments;
- continuing branch and case results;
- bounded-iteration returned elements.

## Union boundaries

Direct variant-to-union assignment does not imply:

- transitive assignment through another union;
- a common union inferred from different branch result types;
- a union inferred from heterogeneous list elements;
- union-to-variant downcast;
- conversion between two unions with overlapping variants.

A union value narrows only inside an exhaustive union switch or another separately specified trusted narrowing boundary. The core language defines no general trusted narrowing boundary outside exhaustive switch.

## Optional boundaries

`optional<T>` may be passed or returned where the target type is exactly `optional<T>`.

Contained `T` is available only inside the optional switch `present` case. The `absent` case contains no value binding.

PuppyDSL defines no `null`, `none`, `some`, or implicit optional constructor.

## Generic constraints

A generic call or type instantiation substitutes one exact type for each parameter.

The `equatable` constraint accepts only types whose equality is defined by the language. The complete concrete type eligibility set remains pending.

A substitution must be consistent across every occurrence of the same type parameter.

## No implicit normalization

PuppyDSL does not use structural equivalence for records or unions.

Two declarations with identical fields or variants remain distinct unless they share the same declared type identity.

## Validation rules

- Reject every incompatible assignment position.
- Reject generic substitutions that are ambiguous or inconsistent.
- Reject a concrete argument that violates a generic constraint.
- Do not infer a union, optional, or conversion to repair an incompatible expression.
- Do not evaluate compatibility until both source and target types resolve.

## Related specs

| ref | relation |
|---|---|
| `spec:puppydsl.language.type_system` | Parent type-system overview. |
| `spec:puppydsl.language.type_system.declared_types` | Defines the nominal source and target types. |
| `spec:puppydsl.language.expressions.invocation` | Applies assignability to call arguments. |
| `spec:puppydsl.language.functions.steps_and_bindings` | Applies assignability to returns and local results. |
| `spec:puppydsl.language.control_flow` | Applies assignability to control-flow result agreement. |
