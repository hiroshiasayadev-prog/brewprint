# Overview: Brewprint DSL language

- **id**: `spec:bpdsl.dsl.overview`
- **status**: draft
- **date**: 2026-06-16
- **parent**: `spec:bpdsl.overview`

## What this is

Entry point for the brewprint DSL language spec. Covers the YAML design language contract: node kinds, edge syntax, file classification, name resolution, type references, project layout, and the design philosophy behind the language.

## Current contract

The same YAML file is the single source of truth for all derived artifacts — diagrams, MCP responses, and implementation context.

### Diagram layers

| layer | views | source elements |
|---|---|---|
| Application | Sequence Diagram, State Diagram, API Table | `actor`, `event`, `state`, `task` (endpoint=true) |
| Processing | DAG | `task`, `asset`, `branch`, `fork`, `join` + flow edges |
| Data | ER Diagram | `store` (kind=db) + model fields |

Layer dependency direction: Application → Processing → Data. Processing layer nodes do not reference Application layer nodes.

### Cross-edges

Cross-layer edges declared on `task` nodes:

| kind | direction | meaning |
|---|---|---|
| `write` | task → store | Task updates a store. |
| `read` | task → store | Task reads from a store. |

### Propagation direction

Forward (event-driven): event trigger → state transition → task execution → asset produced → store written.

Reverse (data-driven): store change → UI update.

### Responsibility split

| layer | responsibility |
|---|---|
| DAG / UML | Happy-path structure |
| impl design | Exceptions, concurrency, transactions |
| Implementation | Code-level guarantees |

Content that cannot be expressed as diagram structure uses `note` fields. Note content is not machine-validated.

### Scope

In scope: node definitions, edge wiring, view rendering, MCP query contract.

Out of scope: code generation, concrete style / visual layout, concurrency / rollback / bidirectional sync (happy-path only).

## Non-goals

Non-functional attributes (`retry`, `idempotent`, `async`) as first-class YAML fields are deferred to post-dogfood evaluation.

## Topics

| title | kind | ref | summary |
|---|---|---|---|
| DSL design philosophy | Concept | `spec:bpdsl.dsl.design_philosophy` | AI-first design, YAML as single source of truth, static verifiability, and happy-path boundary. |
| Node definitions | Overview | `spec:bpdsl.dsl.nodes.overview` | Node kind catalog — processing, data, and application nodes. |
| Edge definitions | Overview | `spec:bpdsl.dsl.edges.overview` | Data flow, state transition, and cross-edge syntax. |
| YAML file types | Reference | `spec:bpdsl.dsl.file_types` | File classification algorithm and `as:` value catalog. |
| Name resolution | Reference | `spec:bpdsl.dsl.naming` | QualifiedID format, sentinel parsing, module scope, and FK resolution. |
| TypeRef | Reference | `spec:bpdsl.dsl.type_ref` | TypeRef syntax, named model resolution, and type compatibility rules. |
| Project layout | Reference | `spec:bpdsl.dsl.project_layout` | Project directory structure, render output placement, and render_index.yaml schema. |
| Diagnostics | Reference | `spec:bpdsl.dsl.diagnostics` | Validation diagnostic codes, severities, and CLI output formats. |
