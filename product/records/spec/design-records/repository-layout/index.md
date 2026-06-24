# Overview: Repository layout model

- **id**: `spec:product.design_records.repository_layout`
- **status**: draft
- **date**: 2026-06-24
- **parent**: `spec:product.design_records`

## What this is

Defines app-independent placement rules for Design Records under `<app>/records/`.
It defines record kind directories, domain-scoped sequential record placement, spec topic-tree placement, and record discovery path patterns.

The current Brewprint repository inventory is documented by `spec:product.brewprint.layout`.

## Current contract

An app namespace that participates in Brewprint design governance may be records-only.
The app namespace stores Design Records under:

```text
<app>/
  records/
```

The `records/` root uses a kind-first structure.
Sequential domain-scoped artifacts use domain subdirectories.

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

Physical paths are repository locations.
Canonical identity comes from public record IDs or path-derived `spec:` refs.

## Rules

- `records/` is the app-independent Design Records placement root.
- `records/` is required for an active app namespace participating in design governance.
- Records-only app namespaces are valid.
- Record placement follows the kind and domain structure defined in this document.
- Spec files are organized by their canonical topic tree under `records/spec/`.
- Namespace identity, ID grammar, and sequence allocation are owned by `spec:product.design_records.namespace_model`.
- Tool-specific index inclusion filters are outside this concept.
- Migration timing and mechanics require explicit migration work.

## Boundary

| owned by this concept | not owned by this concept |
|---|---|
| Placement under `<app>/records/`. | Which app namespaces currently exist. |
| Records-only app namespace validity. | Which directories currently exist in each app. |
| Kind-first Design Records placement. | Active app and domain namespace definitions. |
| Domain-scoped sequential record placement. | Canonical ID grammar and sequence allocation. |
| Spec topic-tree placement. | BPDSL language and generator contracts. |
| Physical path versus canonical identity boundary. | App-specific structure below `dsl/` or `src/`. |
| App-independent record discovery path patterns. | Tool-specific discovery inclusion filters. |

## BPDSL staging note

The previous repository-layout text defined `dsl/`, `src/`, DSL-to-source generation, and handwritten bootstrap behavior.
Those statements are preserved under temporary BPDSL staging.
They are not current generic repository-layout contracts.

| removed source material | preserved location |
|---|---|
| `<app>/dsl/` as BPDSL implementation source. | `spec:product.bpdsl.repository_implementation_flow`. |
| `<app>/src/` as generated source or handwritten bootstrap source. | `spec:product.bpdsl.repository_implementation_flow`. |
| Long-term `dsl/ -> generated src/` target. | `spec:product.bpdsl.repository_implementation_flow`. |
| Direct spec-to-handwritten-`src/` bootstrap allowance. | `spec:product.bpdsl.repository_implementation_flow`. |

## Topics

| title | kind | ref | summary |
|---|---|---|---|
| Record discovery paths | Reference | `spec:product.design_records.repository_layout.record_discovery_paths` | App-independent record path patterns. |

## Related specs

| ref | relation |
|---|---|
| `spec:product.brewprint.layout` | Current Brewprint repository inventory and observed app namespace states. |
| `spec:product.design_records.namespace_model` | App and domain namespace semantics. |
| `spec:product.design_records.spec_format` | Path-derived identity and structure for spec records. |
| `spec:product.bpdsl.repository_implementation_flow` | Temporary preservation of removed BPDSL repository implementation-flow material. |

## Sources

- V01-ADR-097: App namespace-first repository directory layout.
- PRODUCT-TASK-SPEC-005-18: Relocation of record discovery path conventions from DRMCP.
