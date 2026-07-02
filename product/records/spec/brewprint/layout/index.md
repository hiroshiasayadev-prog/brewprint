# Reference: Current Brewprint repository layout

- **id**: `spec:product.brewprint.layout`
- **status**: draft
- **date**: 2026-07-02
- **parent**: `spec:product.brewprint`

## What this is

Records the current Brewprint repository layout.

This file is an inventory reference, not the normative layout contract. Design Records directory meaning, requiredness, and record placement are defined by `spec:product.design_records.repository_layout`. The preserved DSL-to-source target is recorded in `spec:product.bpdsl.repository_implementation_flow`.

## Current app namespace layout

```text
product/
  records/

drmcp/
  records/
  src/

bpdsl/
  records/

trv/
  records/
```

| app namespace | directories currently present | current state |
|---|---|---|
| `product/` | `records/` | Records-only product and cross-app governance namespace. |
| `drmcp/` | `records/`, `src/` | Design records plus handwritten implementation source. No app-local `dsl/` yet. |
| `bpdsl/` | `records/` | Namespaced design records are present. App-local DSL and source placement are not established yet. |
| `trv/` | `records/` | Records-only Task Responsibility Validator namespace. App-local design remains pending. |

## Other current repository areas

| path | current role |
|---|---|
| `v01/` | Read-only legacy snapshot and compatibility source. |
| `scripts/` | Repository scripts and migration or validation helpers. |
| `tools/` | Repository-local supporting tools. |
| `bin/` | Built or runnable helper binaries. |
| root Go module files | Current repository-level Go build bootstrap. |

Root-level temporary, compatibility, or development files may exist. Their presence does not establish new layout rules.

## Maintenance rule

Update this reference when the observed repository tree changes materially.

Do not define new required directories or generation semantics here. Change `spec:product.design_records.repository_layout` when the Design Records placement contract changes. BPDSL implementation-flow material remains under `spec:product.bpdsl.repository_implementation_flow` until BPDSL migration review.

## Related specs

| ref | relation |
|---|---|
| `spec:product.design_records.repository_layout` | Normative Design Records `records/` placement contract. |
| `spec:product.bpdsl.repository_implementation_flow` | Temporary preservation of the former `dsl/`, `src/`, and DSL-to-source model. |
| `spec:product.design_records.namespace_model` | App-independent app and domain namespace semantics. |
