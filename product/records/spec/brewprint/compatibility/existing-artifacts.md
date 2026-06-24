# Reference: Existing artifact compatibility

- **id**: `spec:product.brewprint.compatibility.existing_artifacts`
- **status**: draft
- **date**: 2026-06-24
- **parent**: `spec:product.brewprint.compatibility`

## What this is

Records Brewprint historical ownership decisions, effective attribution, and issued-ID migration state for existing artifacts.

## Historical ownership decision

Per V01-ADR-096, all existing artifacts are treated as owned by the `PRODUCT` namespace.
Per-app namespace attribution migration is not performed.

This decision applies to artifacts created before the app-aware artifact ID grammar was adopted.
It does not change issued artifact IDs.

## Effective attribution

Effective attribution maps legacy domain-prefix groups to a current profile label for navigation and grouping.
It is not retroactive ownership reassignment.
It does not create a new ID or alias.

| domain prefix | example artifacts | effective attribution |
|---|---|---|
| `MCP` | `V01-REQ-MCP-*` / `V01-WORK-MCP-*` / `V01-TASK-MCP-*` | `DRMCP` |
| `DATA` | `V01-REQ-DATA-*` / `V01-WORK-DATA-*` / `V01-INV-DATA-*` | `BPDSL` |
| `RESOLVE` | `V01-REQ-RESOLVE-*` / `V01-WORK-RESOLVE-*` | `BPDSL` |
| `SELFHOST` | `V01-REQ-SELFHOST-*` / `V01-WORK-SELFHOST-*` | Cross-app verification activity. |
| `PRODUCT` | `V01-REQ-PRODUCT-*` / `V01-WORK-PRODUCT-*` / `V01-INV-PRODUCT-*` | `PRODUCT` |

`V01-REQ-MCP-013` remains the issued ID.
Its effective attribution label is `DRMCP`.

`SELFHOST` is cross-app.
No single app attribution applies.

## Migration state

| item | state |
|---|---|
| Issued V01 IDs | Retained unchanged. |
| Per-app migration of existing artifacts | Not performed. |
| Effective attribution | Logical compatibility projection only. |
| New ID or alias creation | Not created by effective attribution. |

## Boundary

| content | treatment |
|---|---|
| New-artifact ownership selection | Generic Design Records semantics in `spec:product.design_records.namespace_model.existing_artifacts`. |
| UI, MCP, projection, display, or grouping behavior | DRMCP app-local specifications. |
| Legacy ID family retention | `spec:product.brewprint.compatibility.legacy_id_compatibility`. |

## Related specs

| ref | relation |
|---|---|
| `spec:product.brewprint.compatibility` | Parent compatibility overview. |
| `spec:product.brewprint.namespaces.domain_catalog` | Current and legacy-effective Brewprint domain profile. |
| `spec:product.design_records.namespace_model.existing_artifacts` | Generic new-artifact ownership rule. |
