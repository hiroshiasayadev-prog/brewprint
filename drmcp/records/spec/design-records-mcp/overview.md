# Overview: Design Records MCP

- **id**: `spec:drmcp.design_records_mcp.overview`
- **status**: draft
- **date**: 2026-06-17
- **parent**: `-`

## What this is

Design Records MCP is an auxiliary MCP server that supports operation of brewprint design records and workflow artifacts through machine-readable metadata, MCP query tools, and validation. Targets ADR, spec, investigation, requirement, work item, and task records.

Primary objectives:

- Build a record index for ADR / spec / investigation / requirement / work item / task.
- Return structured record ID / kind / status / path and kind-specific metadata.
- Return detail representations for a selected set of record IDs in one batch.
- Detect basic metadata inconsistencies in records.
- Resolve semantic/artifact refs between docs artifacts and enable broken-reference checking.
- Allow an LLM in a separate session to narrow down which design records to read before reading their full body.

Design Records MCP does not replace spec-first documentation practice. The sole source of truth for current specifications remains `drmcp/records/spec/**`; ADRs are the authoritative record of design decisions. Design Records MCP is a query/validation layer that makes those relationships mechanically traversable.

> Source: V01-ADR-076

## Current contract

### Record scope

DRMCP indexes the following record kinds. Paths are relative to `<records_root>`; `records_root` and `namespace_prefix` are derived from the app namespace directory (see `spec:drmcp.design_records_mcp.namespace_scanning`).

| kind | discovery path | public ID example (`namespace_prefix = V01-`) |
|---|---|---|
| `decision` | `<records_root>/adr/<domain>/<namespace_prefix>ADR-*-*.md` | `V01-ADR-076` |
| `spec` | `<records_root>/spec/**/*.md` (files with `design_record.id` + `design_record.kind` only) | `V01-SPEC-design-records-mcp-overview` |
| `investigation` | `<records_root>/investigations/<domain>/<namespace_prefix>INV-*-*.md` | `V01-INV-MCP-001` |
| `requirement` | `<records_root>/requirements/<domain>/<namespace_prefix>REQ-*-*.md` | `V01-REQ-MCP-001` |
| `work_item` | `<records_root>/work-items/<domain>/<namespace_prefix>WORK-*-*.md` | `V01-WORK-DRMCP-001` |
| `task` | `<records_root>/tasks/<domain>/<namespace_prefix>TASK-*-*.md` | `V01-TASK-MCP-001-01` |

New ADRs use the domain-subdirectory path shown above. Existing flat V01 ADRs remain discoverable through the compatibility path `<records_root>/adr/<namespace_prefix>ADR-*.md`.

Existing specs without a `design_record` block are not indexed; no `missing_design_record` diagnostic is issued for them. Legacy M-series task records are excluded from `task` discovery. The record kind set is not a closed enumeration — additional artifact kinds may be added by subsequent decisions. UC docs and impl notes are excluded from record kind indexing in MVP.

> Source: V01-ADR-076, V01-ADR-087, V01-ADR-091, V01-ADR-092

### Tool boundary

Design Records MCP operates on a separate data source from the existing brewprint MCP. For the authoritative cross-app artifact model governing this boundary, see `spec:product.design_records.artifact_model`.

| MCP | data source | primary responsibility |
|---|---|---|
| brewprint MCP | `ResolvedProject` built from brewprint YAML | semantic object query / inspect / impact analysis |
| Design Records MCP | bullet metadata in `adr/`, `investigations/`, `requirements/`, `work-items/`, `tasks/`; YAML front matter in `design_record`-bearing `spec/` files | design record / workflow artifact index / read / validation; semantic/artifact ref resolve per traceability spec |

> Source: V01-ADR-076

## Topics

| title | kind | ref | summary |
|---|---|---|---|
| Responsibility boundary | Reference | `spec:drmcp.design_records_mcp.responsibility_boundary` | Boundary against existing brewprint MCP and against general-purpose filesystem tools. |
| Resolver responsibility | Reference | `spec:drmcp.design_records_mcp.resolver` | Canonical reference model this MCP implements; resolver input/output scope and MVP required inputs. |
| Namespace scanning | Reference | `spec:drmcp.design_records_mcp.namespace_scanning` | Multi-root scan behavior, namespace_prefix derivation, and kind-level prefix application. |
| MVP scope | Reference | `spec:drmcp.design_records_mcp.mvp_scope` | P0/P1 tool set and items explicitly outside MVP. |
| Schema | Overview | `spec:drmcp.design_records_mcp.schema.overview` | Record data model, metadata grammar, field definitions, ID normalization, discovery, and authoring schema. |
| Tools | Overview | `spec:drmcp.design_records_mcp.tools.overview` | Full MCP tool set: read/navigation tools, authoring transaction tools, and shared response conventions. |
