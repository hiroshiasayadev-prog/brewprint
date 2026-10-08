# Overview: PuppyDSL

- **id**: `spec:puppydsl.overview`
- **status**: draft
- **date**: 2026-07-25
- **parent**: `root`

## What this is

Entry point for PuppyDSL specifications. PuppyDSL is a typed YAML execution language for declaring semantic functions and composing them without exposing implementation-level operations.

## Current contract

PuppyDSL source declares typed values, callable functions, structured control flow, contexts, and runtime events.

The language preserves these boundaries:

- PuppyDSL source owns design-facing declarations and compositions.
- A validator parses and proves program validity before execution.
- A runtime executes only a validated program.
- Native integrations implement declared functions without changing their PuppyDSL contracts.
- Code generation and application-specific module contracts are separate specification areas.

## Non-goals

- Code-generation APIs or generated source layout.
- Native-language representation binding.
- Application-specific manifests, modules, artifacts, references, or operation contracts.
- General-purpose scripting features such as mutable variables, unbounded loops, reflection, or exception handling.

## Topics

| title | kind | ref | summary |
|---|---|---|---|
| PuppyDSL language | Overview | `spec:puppydsl.language` | Source, type, expression, function, control-flow, context, event, and program-validity contracts. |

## Related specs

| ref | relation |
|---|---|
| `spec:product.design_records.namespace_model` | Generic app namespace model used by this specification tree. |
