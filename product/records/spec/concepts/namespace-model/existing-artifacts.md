# Reference: Existing artifact ownership

- **id**: `spec:product.concepts.namespace_model.existing_artifacts`
- **status**: draft
- **date**: 2026-06-22
- **parent**: `spec:product.concepts.namespace_model`

## What this is

Defines the namespace ownership of existing artifacts under the V01-ADR-096 decision, and the policy for new artifact creation going forward.

## Existing artifact ownership

Per V01-ADR-096, all existing artifacts are treated as owned by the `PRODUCT` namespace. Per-app namespace attribution migration is not performed.

| existing ID prefix | effective app namespace | notes |
|---|---|---|
| `REQ-MCP-*` / `WORK-MCP-*` / `TASK-MCP-*` | DRMCP | Artifacts from the single-app era |
| `REQ-DATA-*` / `WORK-DATA-*` / `TASK-DATA-*` | BPDSL | Artifacts from the single-app era |
| `REQ-RESOLVE-*` / `WORK-RESOLVE-*` | BPDSL | Artifacts from the single-app era |
| `REQ-SELFHOST-*` / `WORK-SELFHOST-*` | Cross-app verification activity | Self-hosting is not a domain of a specific app; it is a verification activity applicable to any app |

For new artifacts, use `<APP_NAMESPACE>-...` form when the owning app namespace is confirmed, or the `PRODUCT` namespace when cross-app or attribution is unknown. For v2 ID grammar details, see `spec:product.concepts.namespace_model.v2_grammar`.
