# Reference: Existing artifact ownership

- **id**: `spec:product.concepts.namespace_model.existing_artifacts`
- **status**: draft
- **date**: 2026-06-24
- **parent**: `spec:product.concepts.namespace_model`

## What this is

Defines the V01-ADR-096 ownership decision for existing artifacts, their effective app namespace attribution for navigation and grouping, and the ownership policy for new artifacts.

## Historical ownership decision

Per V01-ADR-096, all existing artifacts are treated as owned by the `PRODUCT` namespace. Per-app namespace attribution migration is not performed.

This is a historical judgment applied to artifacts created before the artifact ID grammar was adopted. It does not change issued artifact IDs. Issued-ID retention and resolvability are defined in `spec:product.concepts.namespace_model.legacy_id_compatibility`.

## Effective attribution

For navigation, grouping, and attribution resolution in UI and MCP, each domain-prefix group maps to an effective app namespace. This mapping is a logical projection and is not a retroactive ownership reassignment.

| domain prefix | example artifacts | effective app namespace |
|---|---|---|
| `MCP` | `V01-REQ-MCP-*` / `V01-WORK-MCP-*` / `V01-TASK-MCP-*` | `DRMCP` |
| `DATA` | `V01-REQ-DATA-*` / `V01-WORK-DATA-*` / `V01-INV-DATA-*` | `BPDSL` |
| `RESOLVE` | `V01-REQ-RESOLVE-*` / `V01-WORK-RESOLVE-*` | `BPDSL` |
| `SELFHOST` | `V01-REQ-SELFHOST-*` / `V01-WORK-SELFHOST-*` | Cross-app verification activity |
| `PRODUCT` | `V01-REQ-PRODUCT-*` / `V01-WORK-PRODUCT-*` / `V01-INV-PRODUCT-*` | `PRODUCT` |

Tools may associate an effective app namespace label with a legacy ID for navigation and grouping. This does not create a new ID or alias. For example, `V01-REQ-MCP-013` remains the issued ID and has an effective app attribution of `DRMCP`. `SELFHOST` is cross-app; no single app attribution applies.

## New-artifact ownership

For new sequential records (ADR, investigation, requirement, work item, task), use the `<APP_NAMESPACE>-...` form when the owning app namespace is confirmed, or the `PRODUCT` namespace when the concern is cross-app or attribution is unclear.

New PRODUCT-owned artifacts use the canonical app-aware artifact ID form. The domain segment identifies the actual PRODUCT concern:

| example | description |
|---|---|
| `PRODUCT-REQ-SPEC-001` | New requirement in the PRODUCT `SPEC` domain. |
| `PRODUCT-WORK-NAMESPACE-001` | New work item in the PRODUCT `NAMESPACE` domain. |
| `PRODUCT-TASK-NAMESPACE-001-01` | New task under the corresponding PRODUCT work item. |

For the canonical artifact ID grammar, see `spec:product.concepts.namespace_model.artifact_id_grammar`.
