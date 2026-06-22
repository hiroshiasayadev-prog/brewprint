# Reference: Responsibility boundary

- **id**: `spec:drmcp.design_records_mcp.responsibility_boundary`
- **status**: draft
- **date**: 2026-06-17
- **parent**: `spec:drmcp.design_records_mcp.overview`

## What this is

Defines the boundary between Design Records MCP, the existing brewprint MCP, and general-purpose filesystem tools. Design Records MCP operates on a separate data source from the existing brewprint MCP.

## Current contract

### Boundary against existing brewprint MCP

| MCP | data source | primary responsibility |
|---|---|---|
| brewprint MCP | `ResolvedProject` built from brewprint YAML | semantic object query / inspect / impact analysis |
| Design Records MCP | bullet metadata in `adr/`, `investigations/`, `requirements/`, `work-items/`, `tasks/`; YAML front matter in `design_record`-bearing `spec/` files | design record / workflow artifact index / read / validation; semantic/artifact ref resolve per traceability spec |

Design Records MCP is designed to launch and validate independently from the existing brewprint MCP. It does not extend the existing `QueryService` responsibilities into docs management. This spec lives in `drmcp/records/spec/design-records-mcp/**`, not mixed with `bpdsl/records/spec/mcp/tools/**`.

> Source: V01-ADR-076 §既存brewprint MCPとの関係

### Boundary against filesystem tools

Design Records MCP is not a substitute for general-purpose filesystem tools.

**What Design Records MCP handles:**

- Retrieve metadata / path / headings / raw body from one or multiple explicitly specified record IDs.
- Return a structured list of records.
- Validate basic metadata inconsistencies in records.
- Suggest the next ADR number and recommended path.

**What Design Records MCP does not handle:**

- Arbitrary file read/write.
- General Markdown editing.
- Automatic ADR body generation or update.
- Automatic rewriting of commit hashes.
- Git operations.

> Source: V01-ADR-077 §filesystemとの責務境界, V01-ADR-090 §決定
