# Reference: Current Brewprint repository layout

- **id**: `spec:product.brewprint.layout`
- **status**: draft
- **date**: 2026-06-23
- **parent**: `spec:product.brewprint`

## What this is

Records the current Brewprint repository layout.

This file is an inventory reference, not the normative layout contract. Directory meaning, requiredness, record placement, and the DSL-to-source target are defined by `spec:product.concepts.repository_layout`.

## Current app namespace layout

```text
product/
  records/

drmcp/
  records/
  src/

bpdsl/
  records/
```

| app namespace | directories currently present | current state |
|---|---|---|
| `product/` | `records/` | Records-only product and cross-app governance namespace. |
| `drmcp/` | `records/`, `src/` | Design records plus handwritten implementation source. No app-local `dsl/` yet. |
| `bpdsl/` | `records/` | Namespaced design records are present. App-local DSL and source placement are not established yet. |

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

Do not define new required directories or generation semantics here. Change `spec:product.concepts.repository_layout` when the product contract changes.

## Related specs

| ref | relation |
|---|---|
| `spec:product.concepts.repository_layout` | Normative `records/`, `dsl/`, `src/`, record placement, and DSL-to-source contract. |
| `spec:product.concepts.namespace_model` | Active app and domain namespace semantics. |
