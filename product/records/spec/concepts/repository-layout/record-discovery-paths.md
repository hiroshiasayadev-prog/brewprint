# Reference: Record discovery paths

- **id**: `spec:product.concepts.repository_layout.record_discovery_paths`
- **status**: draft
- **date**: 2026-06-22
- **parent**: `spec:product.concepts.repository_layout`

## What this is

Defines the path-pattern conventions for discovering record files within a `records_root` by kind. Relocated from `spec:drmcp.design_records_mcp.schema.discovery` (path-pattern content only; DRMCP kind-filter behavior stays in `spec:drmcp.design_records_mcp.schema.discovery`).

## Current contract

| kind | path pattern |
|---|---|
| `decision` | `<records_root>/adr/<namespace_prefix>ADR-*.md` |
| `spec` | `<records_root>/spec/**/*.md` |
| `investigation` | `<records_root>/investigations/*/<namespace_prefix>INV-*-*.md` |
| `requirement` | `<records_root>/requirements/*/<namespace_prefix>REQ-*-*.md` |
| `work_item` | `<records_root>/work-items/*/<namespace_prefix>WORK-*-*.md` |
| `task` | `<records_root>/tasks/*/<namespace_prefix>TASK-*-*.md` |

`<namespace_prefix>` is derived from the `records_root` per `spec:product.concepts.namespace_model.v1_namespace_algorithm`.

DRMCP-specific index inclusion conditions (e.g. whether a spec file is included based on its front-matter fields) are defined in `spec:drmcp.design_records_mcp.schema.discovery`.

> Source: V01-ADR-076 §bootstrap方針, V01-ADR-092 §1
