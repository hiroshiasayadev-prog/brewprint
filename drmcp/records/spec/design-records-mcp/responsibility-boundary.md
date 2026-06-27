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
| brewprint MCP | `ResolvedProject` built from brewprint YAML | semantic object query / inspect / impact analysis |
| Design Records MCP | Current Markdown records under explicitly configured app roots; H1-adjacent metadata; path-derived current spec identity; optional configured legacy archive sources in a separate index | Design record and workflow artifact discovery, indexing, read operations, validation execution, and reference resolution under their owning contracts |

Design Records MCP is designed to launch and validate independently from the existing brewprint MCP. It does not extend the existing `QueryService` responsibilities into docs management. This spec lives in `drmcp/records/spec/design-records-mcp/**`, not mixed with `bpdsl/records/spec/mcp/tools/**`.

PRODUCT namespace, repository-layout, and spec-format specifications own current identity and source semantics. DRMCP owns configured-root loading, parser mapping, index construction, validation execution, and MCP operation behavior.

Sources: `DRMCP-REQ-MCP-001`, `DRMCP-TASK-MCP-003-02`, `DRMCP-TASK-MCP-003-03`, and `DRMCP-TASK-MCP-003-04`.

### Boundary against filesystem tools

Design Records MCP is not a substitute for general-purpose filesystem tools.

**What Design Records MCP handles:**

- Retrieve normalized metadata, headings, and optional raw body through the owning read-operation contracts.
- Return structured record projections through the owning query contract.
- Validate metadata, identity, and relation inconsistencies.
- Retain repository-relative source provenance for validation and repair.
- Suggest the next ADR number and recommended authoring location.

**What Design Records MCP does not handle:**

- Arbitrary file read/write.
- General Markdown editing.
- Automatic ADR body generation or update.
- Automatic rewriting of commit hashes.
- Git operations.

Normal list and exact-retrieval response representation is owned by `DRMCP-WORK-MCP-004`. Diagnostic source-location fields and exceptional physical-path exposure are owned by `DRMCP-WORK-MCP-006`.

Historical sources: V01-ADR-077 and V01-ADR-090.
