# Reference: MVP scope

- **id**: `spec:drmcp.design_records_mcp.mvp_scope`
- **status**: draft
- **date**: 2026-06-28
- **parent**: `spec:drmcp.design_records_mcp.overview`

## What this is

Defines the current-format read baseline tool surface and its exclusions without asserting implementation status.

## Current contract

### P0 current-format read tools

| tool | current responsibility | contract owner |
|---|---|---|
| `list_records` | Compact active-index listing for one app namespace, supported sequential kind, and domain. Specs and legacy archive records are excluded. | `DRMCP-WORK-MCP-004` |
| `get_records` | Sole exact retrieval for one through 20 current or configured legacy refs. Single retrieval uses one `refs` element. | `DRMCP-WORK-MCP-004` |
| `resolve_reference` | Current-first reference resolution and configured legacy fallback. | `DRMCP-WORK-MCP-005` |
| `validate_records` | Current-record validation and machine-readable diagnostic results. | `DRMCP-WORK-MCP-006` |

`get_record` is not part of the current tool surface. The contract does not retain a compatibility alias.

This table classifies the current-format read contract only. It does not claim that an operation is implemented or released.

Authoring-guidance and authoring-transaction tools are governed by their own contracts and delivery phases. This read-baseline table does not reclassify them.

### P1 tools

No P1 read tool is defined by the current realignment baseline.

`suggest_next_record` is removed. Namespace-aware authoring placeholders own sequential allocation.

### Read-baseline exclusions

| item | exclusion |
|---|---|
| Broad, cross-app, cross-kind, or cross-domain listing | `list_records` requires explicit app, sequential kind, and domain scope. |
| Range or exact-ID listing | Exact retrieval uses `get_records`; range listing is unsupported. |
| Spec or legacy normal listing | Specs remain exact-retrieval targets. Legacy records require configured exact or resolver behavior. |
| Resolver invocation during exact retrieval | `get_records` classifies each exact input once and does not invoke `resolve_reference`. |
| Warning and diagnostic representation or repository-validation execution in W004 | W004 owns its operation warning triggers only. Shared representation is defined by `spec:drmcp.design_records_mcp.schema.diagnostics`, validation execution by `spec:drmcp.design_records_mcp.tools.validate_records`, and semantic invalidity by PRODUCT authorities. |
| Normal successful path projection | Successful list, exact-retrieval, and resolver projections remain path-free. Portable diagnostic locations and explicit authoring paths follow their owning contracts. No current operation exposes an absolute physical path. |
| Natural-language dependency inference | Read operations use explicit parsed source content only. |
| Git history or code static analysis | Outside the Design Records read baseline. |
| UI, multi-project management, or public CLI design | Outside the MCP read contract. |
| Workflow graph analysis or derived progress | Outside the current read tool surface. |

### Source boundary

W003 owns discovery, current source parsing, canonical identity, active-index construction, normalized record fields, and invalid-source retention.

Current specs use H1-adjacent metadata and path-derived `spec:` identity. Legacy archive sources remain separate and do not enter normal listing.
