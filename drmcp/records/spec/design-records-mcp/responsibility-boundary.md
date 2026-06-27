# Reference: Responsibility boundary

- **id**: `spec:drmcp.design_records_mcp.responsibility_boundary`
- **status**: draft
- **date**: 2026-06-27
- **parent**: `spec:drmcp.design_records_mcp.overview`

## What this is

Defines the boundary between Design Records MCP, the existing brewprint MCP, and general-purpose filesystem tools. Design Records MCP operates on a separate data source from the existing brewprint MCP.

## Current contract

### Boundary against existing brewprint MCP

| MCP | data source | primary responsibility |
|---|---|---|
| brewprint MCP | `ResolvedProject` built from brewprint YAML | Semantic object query, inspection, and impact analysis. |
| Design Records MCP | Current Markdown records under configured app roots, plus optional configured legacy archive roots in a separate index | Design record discovery, indexing, listing, exact retrieval, reference resolution, validation, and authoring under their owning contracts. |

Design Records MCP operates independently from the existing brewprint MCP. It does not extend the existing `QueryService` into Design Records management.

PRODUCT namespace, repository-layout, authoring, traceability, and spec-format contracts own semantic rules. DRMCP owns concrete parsing, indexing, operation, validation, and response behavior.

### DRMCP read responsibility split

| owner | responsibility |
|---|---|
| `DRMCP-WORK-MCP-003` | Configured current-root discovery, source parsing, canonical identity, active-index construction, normalized record model, invalid-source retention, and duplicate-conflict state. |
| `DRMCP-WORK-MCP-004` | `list_records` request and compact result behavior; `get_records` exact request and successful-record projection; batch ordering; partial-success triggers; request-wide body inclusion; normal path hiding. |
| `DRMCP-WORK-MCP-005` | Resolver invocation, current-first resolution, configured legacy fallback, fallback ordering, and current/legacy resolution orchestration. |
| `DRMCP-WORK-MCP-006` | Warning and diagnostic schema, category names, severity, shared fields, source-location representation, validation semantics, and exceptional path exposure. |

`list_records` does not perform exact retrieval or resolution. It queries the active index for one app namespace, supported sequential kind, and domain.

`get_records` is the sole exact-retrieval operation. A single retrieval uses one `refs` element. Exact retrieval does not invoke the resolver.

### Boundary against filesystem tools

Design Records MCP is not a substitute for general-purpose filesystem tools.

Design Records MCP handles:

- compact current-record listing through `list_records`;
- exact current and configured legacy retrieval through `get_records`;
- normalized metadata, headings, and request-wide optional body projection;
- reference resolution under the resolver contract;
- validation and machine-readable diagnostics;
- source provenance retained internally for validation and repair.

Design Records MCP does not handle:

- arbitrary file read or write;
- general Markdown editing outside authoring transaction contracts;
- automatic record-body generation;
- Git operations;
- physical-path projection in normal list or exact-retrieval records.

Physical paths may appear only through W006-owned diagnostic, patch, debug, or emergency surfaces.
