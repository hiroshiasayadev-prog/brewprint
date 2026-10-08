# Overview: PuppyDSL language

- **id**: `spec:puppydsl.language`
- **status**: draft
- **date**: 2026-07-25
- **parent**: `spec:puppydsl.overview`

## What this is

Defines the design-facing PuppyDSL language. The specification covers source discovery, declarations, type semantics, expressions, functions, structured control flow, contexts, events, and whole-program validity.

## Current contract

A PuppyDSL application is one source set beneath a selected application root.

```text
.puppy.yaml sources
  -> source-envelope and YAML parsing
  -> declaration registration
  -> name and type resolution
  -> expression and control-flow analysis
  -> validated typed program
  -> runtime execution
```

The language is statically typed and nominal by default. Bindings are immutable. Calls and control-flow nodes use declared semantic functions and bounded constructs instead of arbitrary operators or reflection.

Expected domain outcomes are ordinary typed values. Unexpected execution failures propagate outside the PuppyDSL value model.

## Non-goals

- Generator APIs, generated package layout, or native representation wrappers.
- Application-specific module dispatch or manifest schemas.
- Mutable variables, arbitrary operators, lambdas, reflection, dynamic objects, or exception recovery.
- Unbounded loops or YAML-defined recursion.
- Diagnostic transport formats or validator CLI behavior.

## Topic map

```text
source model
  -> type system
  -> expressions
  -> functions
  -> control flow
  -> contexts and events
  -> whole-program validity
```

## Topics

| title | kind | ref | summary |
|---|---|---|---|
| Source model | Overview | `spec:puppydsl.language.source_model` | Source discovery, declaration files, entry declaration, and source envelopes. |
| Type system | Overview | `spec:puppydsl.language.type_system` | Type expressions, declared types, generic composites, and assignability. |
| Expressions | Overview | `spec:puppydsl.language.expressions` | Calls, literals, constructors, binding references, and access chains. |
| Functions | Overview | `spec:puppydsl.language.functions` | Function declarations, implementation forms, steps, and immutable bindings. |
| Control flow | Overview | `spec:puppydsl.language.control_flow` | Conditional execution, exhaustive switch, bounded iteration, and continue. |
| Contexts | Concept | `spec:puppydsl.language.contexts` | Structured context lifetimes, providers, requirements, caching, and invalidation. |
| Events | Concept | `spec:puppydsl.language.events` | Declared event identity, typed dispatch, context invalidation, and bounded restart. |
| Program validity | Overview | `spec:puppydsl.language.program_validity` | Resolution, scope, type, path, graph, context, and restart validity rules. |

## Related specs

| ref | relation |
|---|---|
| `spec:puppydsl.overview` | Parent PuppyDSL specification entry point. |
