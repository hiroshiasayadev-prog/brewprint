# Overview: PuppyDSL control flow

- **id**: `spec:puppydsl.language.control_flow`
- **status**: draft
- **date**: 2026-07-25
- **parent**: `spec:puppydsl.language`

## What this is

Defines bounded structured control flow for boolean branching, exhaustive elimination, collection iteration, and iteration skipping.

## Current contract

PuppyDSL control flow contains three structured step forms:

| form | purpose |
|---|---|
| `if` | Branch on one previously bound boolean. |
| `switch` | Exhaustively eliminate one union, enum, or optional value. |
| `for` | Iterate one finite immutable list or dictionary. |

Control-flow nodes follow the common lexical-scope and result rules from `spec:puppydsl.language.functions.steps_and_bindings`.

A control-flow step uses an ordinary key for an information-bearing result and a discard key for `void` or `never`.

## Shared result rules

| completion shape | complete result |
|---|---|
| At least one path continues and no continuing path returns a value | `void`. |
| Every path terminates through `never` | `never`. |
| Continuing paths return values | Every continuing path returns the same exact type. |
| A path terminates through `never` | The path does not participate in result-type agreement. |
| A path uses `continue` inside `for` | The path contributes no iteration output. |

PuppyDSL does not infer a union or optional to reconcile different branch result types.

## Common lexical scope

- Child scopes may read active ancestor bindings.
- Child locals are unavailable after the child scope completes.
- A child scope cannot shadow or reassign an active ancestor binding.
- Sibling child scopes may reuse the same local name.
- Only the complete information-bearing control-flow result enters the enclosing scope.

## Non-goals

- `while`, unbounded loops, or general recursion.
- `break` or labelled `continue`.
- Open-ended switch over strings, integers, or named scalars.
- Inline comparisons, arithmetic, boolean operators, or type tests.
- Exception handling or recovery.

## Topics

| title | kind | ref | summary |
|---|---|---|---|
| Conditional execution | Reference | `spec:puppydsl.language.control_flow.conditional` | Boolean conditions, branch completion, omitted `else`, and branch result typing. |
| Exhaustive switch | Reference | `spec:puppydsl.language.control_flow.switch` | Union, enum, and optional cases, narrowing, and exhaustive result rules. |
| Bounded iteration | Reference | `spec:puppydsl.language.control_flow.iteration` | List and dictionary iteration, collection results, ordering, and `continue`. |

## Related specs

| ref | relation |
|---|---|
| `spec:puppydsl.language.functions.steps_and_bindings` | Defines shared scope and binding behavior. |
| `spec:puppydsl.language.type_system.assignability` | Defines result compatibility. |
| `spec:puppydsl.language.program_validity.flow_analysis` | Validates path completion and reachability. |
