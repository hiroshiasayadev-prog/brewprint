# Overview: Design Records MCP

- **id**: `spec:drmcp.design_records_mcp.overview`
- **status**: draft
- **date**: 2026-06-27
- **parent**: `-`

## What this is

Design Records MCP is an auxiliary MCP server that supports operation of brewprint design records and workflow artifacts through machine-readable metadata, MCP query tools, and validation. Targets ADR, spec, investigation, requirement, work item, and task records.

Primary objectives:

- Build one active index over configured current records roots.
- Cover ADR, spec, investigation, requirement, work-item, and task records.
- Return query and retrieval projections under operation-specific contracts.
- Retain source provenance for validation and repair.
- Detect metadata and relation inconsistencies in records.
- Resolve semantic and artifact refs under the resolver contract.
- Allow an LLM to narrow the design records that require full-body reading.

Design Records MCP does not replace spec-first documentation practice. Current specifications remain authoritative in each app namespace's `records/spec/**` tree. ADRs remain the authoritative record of design decisions. Design Records MCP makes those records mechanically discoverable and traversable.

Historical source: V01-ADR-076. Current discovery and identity authority comes from `DRMCP-REQ-MCP-001`, W003 outputs, and the linked PRODUCT specifications.

## Current contract

### Record scope

DRMCP builds one active index from explicit current-root configuration. Each entry pairs one `app_namespace` with its repository-relative `<app_namespace>/records` root. DRMCP does not auto-discover app directories or derive an app namespace from a path.

| kind | candidate path under `<records_root>` | current canonical identity |
|---|---|---|
| `decision` | `adr/<domain>/<record_prefix>ADR-*-*.md`, plus the PRODUCT-defined flat ADR compatibility pattern | Complete app-aware artifact ID from H1. |
| `spec` | `spec/**/*.md` | Path-derived canonical `spec:` ref. |
| `investigation` | `investigations/<domain>/<record_prefix>INV-*-*.md` | Complete app-aware artifact ID from H1. |
| `requirement` | `requirements/<domain>/<record_prefix>REQ-*-*.md` | Complete app-aware artifact ID from H1. |
| `work_item` | `work-items/<domain>/<record_prefix>WORK-*-*.md` | Complete app-aware artifact ID from H1. |
| `task` | `tasks/<domain>/<record_prefix>TASK-*-*.md` | Complete app-aware task ID from H1. |

PRODUCT repository-layout specifications own candidate path patterns. PRODUCT namespace and spec-format specifications own canonical identity semantics. DRMCP owns configured-root loading, parser mapping, source retention, and active-index construction.

Current specs use H1-adjacent visible metadata. Their canonical identity comes from the configured app namespace and repository placement. Metadata `id` is a required consistency value. YAML front matter and legacy spec aliases are not active current metadata or identity inputs.

A candidate with one unique canonical ID remains addressable when other source content is invalid. A source without a determinable canonical ID remains validation-only. Duplicate canonical identity creates no winner. Current and optional legacy indexes remain separate operational scopes.

Sources: `DRMCP-TASK-MCP-003-02`, `DRMCP-TASK-MCP-003-03`, `DRMCP-TASK-MCP-003-04`, and the PRODUCT authorities linked from their outputs.

### Tool boundary

Design Records MCP operates on a separate data source from the existing brewprint MCP. For the authoritative cross-app artifact model governing this boundary, see `spec:product.design_records.artifact_model`.

| MCP | data source | primary responsibility |
|---|---|---|
| brewprint MCP | `ResolvedProject` built from brewprint YAML | semantic object query / inspect / impact analysis |
| Design Records MCP | Current Markdown records under explicitly configured app roots; H1-adjacent metadata; path-derived current spec identity; optional configured legacy archive sources in a separate index | Design record and workflow artifact discovery, indexing, read operations, validation execution, and reference resolution under their owning contracts |

Normal list and exact-retrieval representation is owned by `DRMCP-WORK-MCP-004`. Diagnostic source-location and exceptional physical-path exposure are owned by `DRMCP-WORK-MCP-006`.

## Topics

| title | kind | ref | summary |
|---|---|---|---|
| Responsibility boundary | Reference | `spec:drmcp.design_records_mcp.responsibility_boundary` | Boundary against existing brewprint MCP and against general-purpose filesystem tools. |
| Resolver responsibility | Reference | `spec:drmcp.design_records_mcp.resolver` | Canonical reference model this MCP implements; resolver input/output scope and MVP required inputs. |
| Namespace scanning | Reference | `spec:drmcp.design_records_mcp.namespace_scanning` | Explicit current-root configuration, app association, active-index construction, and current/legacy root separation. |
| MVP scope | Reference | `spec:drmcp.design_records_mcp.mvp_scope` | P0/P1 tool set and items explicitly outside MVP. |
| Schema | Overview | `spec:drmcp.design_records_mcp.schema.overview` | Record data model, metadata grammar, field definitions, ID normalization, discovery, and authoring schema. |
| Tools | Overview | `spec:drmcp.design_records_mcp.tools.overview` | Full MCP tool set: read/navigation tools, authoring transaction tools, and shared response conventions. |
