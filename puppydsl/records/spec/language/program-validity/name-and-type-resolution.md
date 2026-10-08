# Reference: PuppyDSL name and type resolution

- **id**: `spec:puppydsl.language.program_validity.name_and_type_resolution`
- **status**: draft
- **date**: 2026-07-25
- **parent**: `spec:puppydsl.language.program_validity`

## What this is

Defines declaration registries and static resolution for types, functions, contexts, provided values, events, enum members, signatures, and expression targets.

## Registries

One source set constructs these logical registries:

| registry | identity |
|---|---|
| Type registry | Type ID. |
| Callable registry | Ordinary function ID and generated event callable ID. |
| Context registry | Context ID. |
| Provided-value registry | `<context-id>.<provided-key>`. |
| Event registry | Event ID. |

One identity appears at most once in its registry.

Ordinary function IDs and generated event callable IDs share one callable namespace.

## Declaration resolution

A declaration resolves every referenced identity before execution.

| reference position | required target |
|---|---|
| Record field, union variant, signature input or output | Type expression. |
| Function call | Callable declaration. |
| Enum literal or enum switch case | Declared enum and member. |
| `opens` | Context declaration. |
| `requires` | Provided-value declaration. |
| Provider argument | Same-context or route-valid ancestor provided value. |
| `invalidated_by`, `restart_on` | Event declaration. |

Unknown identities are program errors.

## Type-expression resolution

The resolver:

- distinguishes built-in types, declared types, and active type parameters;
- validates generic argument counts;
- validates dictionary key eligibility;
- validates generic constraints;
- rejects `never` in prohibited positions;
- preserves exact nominal identity.

A generic parameter resolves only within its declaration.

## Signature resolution

Every function signature must:

- use the required folded YAML scalar style;
- parse completely;
- contain unique input names;
- resolve every input and output type;
- declare exactly one output.

A function declaration key and optional `generic` mapping must agree on parameter identity and constraints.

## Invocation resolution

A direct invocation resolves its function target before argument checking.

After target resolution, validation checks:

- positional or named mode;
- argument count and names;
- lexical binding references;
- literal types;
- generic substitutions;
- assignability to signature inputs;
- static output type.

## Constructor and access resolution

Record construction resolves one concrete transparent-record target and its complete field set.

Dictionary construction resolves exact `K` and `V` types.

An access chain resolves every field or dictionary postfix in order. Each intermediate resolved type determines whether the next postfix is valid.

## Context resolution

`requires` and provider references use full provided-value IDs.

A resolved provided value does not by itself prove route availability. Route availability belongs to flow analysis.

## Event callable resolution

An event declaration contributes one callable contract derived from its payload.

The generated callable is resolved like an ordinary typed call, but it has no ordinary implementation form.

## Validation rules

- Reject duplicate declaration identities.
- Reject conflicts in the shared callable namespace.
- Reject unknown types, functions, contexts, provided values, events, or enum members.
- Reject malformed or partially parsed type expressions and signatures.
- Reject generic parameter or constraint mismatches.
- Reject invalid call argument binding or assignability.
- Reject invalid constructor targets, fields, keys, or values.
- Reject invalid access-chain postfixes.
- Evaluate compatibility only after every prerequisite identity and type resolves.

## Related specs

| ref | relation |
|---|---|
| `spec:puppydsl.language.program_validity` | Parent validity overview. |
| `spec:puppydsl.language.source_model.identifiers` | Defines registry identity spelling. |
| `spec:puppydsl.language.type_system.type_expressions` | Defines type references. |
| `spec:puppydsl.language.expressions.invocation` | Defines call binding. |
| `spec:puppydsl.language.contexts` | Defines provided-value identity. |
| `spec:puppydsl.language.events` | Defines generated event callables. |
