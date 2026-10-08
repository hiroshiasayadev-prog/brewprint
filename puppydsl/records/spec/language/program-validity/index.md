# Overview: PuppyDSL program validity

- **id**: `spec:puppydsl.language.program_validity`
- **status**: draft
- **date**: 2026-07-25
- **parent**: `spec:puppydsl.language`

## What this is

Defines the conditions under which one discovered PuppyDSL source set forms a valid typed program.

## Current contract

Program validity is established before execution.

```text
source set
  -> source-envelope validation
  -> declaration decoding
  -> registry construction
  -> name and type resolution
  -> expression and signature checking
  -> lexical and path analysis
  -> call, context, provider, and event graph analysis
  -> profile validity
```

A validity rule is evaluated only after its required declarations, identities, and types resolve. One invalid declaration does not determine the validity of an independent declaration.

## Validation layers

| layer | responsibility |
|---|---|
| Source | File suffix, raw header, YAML document, root, and entry file. |
| Declaration | Required fields, declaration shape, identifiers, and duplicate identities. |
| Resolution | Types, functions, contexts, provided values, events, and enum members. |
| Typing | Signatures, calls, constructors, access chains, returns, and control-flow results. |
| Lexical flow | Binding visibility, shadowing, reachability, `never`, and `continue`. |
| Program graphs | Function calls, context routes, provider dependencies, event dispatch, and restart safety. |
| Profile | Authoring, design-closure, and activation requirements. |

## Validation profiles

| profile | purpose |
|---|---|
| Authoring | Accept complete typed declarations while permitting valid stubs. |
| Design closure | Require every reachable declared design function to have an executable implementation direction completed. |
| Activation | Require all selected entry routes, native bindings, contexts, and events to be executable. |

Profiles add readiness constraints. A later profile does not weaken syntax, resolution, or type validity.

## Validity dependencies

A rule that requires unresolved information has no valid input until the prerequisite succeeds.

Examples:

- Type compatibility requires resolved source and target types.
- Argument binding requires a resolved callable signature.
- Iteration element typing requires a resolved collection type.
- Exhaustiveness requires a valid union, enum, or optional target type.

Independent declarations and independent graph regions retain their own validity judgments.

## Determinism

For one source set, selected entry-route set, and validation profile, the program validity result is deterministic.

Diagnostic object shape, diagnostic aggregation, cascade-reporting policy, ordering, source-location projection, and CLI serialization belong to validator specifications.

## Non-goals

- Diagnostic transport formats.
- Runtime execution-failure handling.
- Application-specific semantic validation.
- Code-generation conformance.

## Topics

| title | kind | ref | summary |
|---|---|---|---|
| Name and type resolution | Reference | `spec:puppydsl.language.program_validity.name_and_type_resolution` | Registry, duplicate, symbol, signature, type-expression, call, and context-name resolution. |
| Flow analysis | Reference | `spec:puppydsl.language.program_validity.flow_analysis` | Lexical reachability, path completion, call cycles, context routes, provider cycles, and restart safety. |

## Related specs

| ref | relation |
|---|---|
| `spec:puppydsl.language.source_model` | Supplies the source set and declarations. |
| `spec:puppydsl.language.type_system` | Supplies static type rules. |
| `spec:puppydsl.language.functions.implementation_forms` | Defines profile-specific stub validity. |
