# Overview: Repository layout model

- **id**: `spec:product.concepts.repository_layout`
- **status**: draft
- **date**: 2026-06-23
- **parent**: `root`

## What this is

Defines the repository layout contract for Brewprint app namespaces.

This concept owns the meaning and requiredness of `records/`, `dsl/`, and `src/`, the target DSL-to-source relationship, and design-record placement below each app namespace's `records/` root.

The current Brewprint repository inventory is documented separately by `spec:product.brewprint.layout`.

## Current contract

Top-level app namespace directories separate app ownership before artifact kind or implementation technology.

An app namespace may contain these concern-based directories:

```text
<app>/
  records/
  dsl/
  src/
```

| directory | contract |
|---|---|
| `records/` | Stores app-owned design records, workflow artifacts, specifications, and authoring guidance. Required for an active app namespace participating in Brewprint design governance. |
| `dsl/` | Stores BPDSL definitions used as an implementation source. Present only when the app owns DSL definitions. |
| `src/` | Stores runtime or implementation source. Present only when the app owns implementation code. |

A records-only app namespace is valid.

### DSL-to-source relationship

For an implementation-bearing app with an operational DSL pipeline, the long-term target is:

```text
dsl/ -> generated src/
```

Until BPDSL can express and generate the required implementation, handwritten `src/` derived from current specifications and internal design is an allowed bootstrap state.

Handwritten bootstrap source does not replace `records/spec/` as the canonical authority for current design contracts.

### Records structure

The `records/` root uses a kind-first structure. Sequential domain-scoped artifacts use domain subdirectories.

```text
records/
  adr/<domain>/
  spec/
  investigations/<domain>/
  requirements/<domain>/
  work-items/<domain>/
  tasks/<domain>/
  guides/
```

| directory | purpose |
|---|---|
| `adr/<domain>/` | Architecture decision records grouped by domain. |
| `spec/` | App-owned specifications organized by their canonical topic tree. |
| `investigations/<domain>/` | Investigation records grouped by domain. |
| `requirements/<domain>/` | Requirement records grouped by domain. |
| `work-items/<domain>/` | Work items grouped by domain. |
| `tasks/<domain>/` | Tasks grouped by domain. |
| `guides/` | Auxiliary authoring guides not represented as spec records. |

Physical paths are repository locations. Canonical identity comes from public record IDs or path-derived `spec:` references.

## Rules

- `records/`, `dsl/`, and `src/` are concern-based; this contract does not require all three in every app namespace.
- `records/` is required for an active app namespace participating in design governance.
- `dsl/` and `src/` are optional unless the app owns those concerns.
- The target implementation flow for an operational DSL pipeline is `dsl/` to generated `src/`.
- Direct spec-to-handwritten-`src/` implementation is a permitted bootstrap state while DSL support is insufficient.
- Record placement follows the kind and domain structure defined in this document.
- Namespace identity, ID grammar, and sequence allocation are owned by `spec:product.concepts.namespace_model`.
- Tool-specific index inclusion filters are outside this concept.
- Migration timing and mechanics require explicit migration work.

## Boundary

Owned by this concept:

- App namespace directory concerns: `records/`, `dsl/`, and `src/`.
- Records-only app namespaces.
- The target DSL-to-generated-source relationship.
- The handwritten source bootstrap allowance.
- Kind-first and domain-scoped design-record placement.
- Compatibility path policy expressed by child specs.

Not owned by this concept:

- Which app namespaces currently exist or which directories currently exist in each app.
- Active app and domain namespace definitions.
- Canonical ID grammar and sequence allocation.
- BPDSL language and generator contracts.
- App-specific structure below `dsl/` or `src/`.
- Tool-specific discovery inclusion filters.
- Actual migration execution.

## Topics

| title | kind | ref | summary |
|---|---|---|---|
| Record discovery paths | Reference | `spec:product.concepts.repository_layout.record_discovery_paths` | Kind-level path patterns for new records and legacy compatibility discovery. |

## Related specs

| ref | relation |
|---|---|
| `spec:product.brewprint.layout` | Current Brewprint repository inventory and observed app namespace states. |
| `spec:product.concepts.namespace_model` | App and domain namespace semantics. |
| `spec:product.concepts.namespace_model.artifact_id_grammar` | Record ID grammar and sequence allocation scope. |
| `spec:product.concepts.project_artifact_model.design_flow` | Design-to-DSL-to-source flow and bootstrap relationship. |
| `spec:product.concepts.spec_format` | Path-derived identity and structure for spec records. |

## Sources

- V01-ADR-095: Coupling boundary between YAML DSL and Design Records MCP.
- V01-ADR-097: App namespace-first repository directory layout.
- PRODUCT-TASK-SPEC-005-18: Relocation of record discovery path conventions from DRMCP.
