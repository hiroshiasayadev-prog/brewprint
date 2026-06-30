# Reference: Record model

- **id**: `spec:drmcp.design_records_mcp.schema.record_model`
- **status**: draft
- **date**: 2026-06-30
- **parent**: `spec:drmcp.design_records_mcp.schema.overview`

## What this is

Defines how DRMCP retains discovered current sources and builds addressable active-index records.

## Current contract

### Source retention

Every discovered current candidate retains source provenance:

| value | meaning |
|---|---|
| `app_namespace` | Configured app namespace for the current records root. |
| `kind` | Record kind selected by the discovery path. |
| `path` | Repository-relative Markdown source path. |

The path is operational source data. Normal retrieval path exposure is owned by `DRMCP-WORK-MCP-004`; repair diagnostics are owned by `DRMCP-WORK-MCP-006`.

### Index behavior

| source state | active-index behavior | validation behavior |
|---|---|---|
| One canonical ID is determined and all required content is valid. | Create a normal addressable record. | Validate normally. |
| One canonical ID is determined but H1, metadata, or document content is invalid. | Keep the source addressable under that canonical ID. Preserve only values actually parsed; do not invent missing values. | Report the invalid source with provenance. |
| No canonical ID can be determined. | Do not create an addressable record. | Retain the source path for validation and repair diagnostics. |
| Multiple sources produce the same canonical ID. | Create no winner for that ID. Filesystem order must not choose one. | Retain every conflicting source path and report the conflict. |

Unrelated valid records remain addressable when another source is invalid or conflicted.

The canonical ID rules are defined by `spec:drmcp.design_records_mcp.schema.id_normalization`.
The common and kind-specific parsed fields are defined by `spec:drmcp.design_records_mcp.schema.fields`.

### Field availability

A valid record supplies every field required by PRODUCT authority for its kind.
An invalid but addressable record may have missing or invalid fields.

- Canonical `id` and record `kind` exist for every addressable record.
- A missing field remains missing.
- A parsed invalid value remains available for validation; it is not replaced with a default.
- DRMCP does not synthesize title text, lifecycle status, dates, or relation values from filenames or unrelated content.

### Internal state boundary

The runtime keeps these states separate:

| state | contract |
|---|---|
| Source state | Raw source data and provenance. |
| Parsed state | Parsed fields and document structure, including invalid present values. |
| Addressable index entry | One uniquely addressable canonical identity. |
| Invalid but addressable state | One canonical identity with invalid or missing non-identity content. |
| Identity conflict | Every source claiming one canonical identity, with no winner. |
| Validation finding | Transport-neutral finding data. |
| Public projection | Operation-specific result created by an application use case. |

No nullable record type spans source loading, parsing, indexing, validation, application, and MCP transport.

Concrete Go struct names are not prescribed here. `spec:drmcp.implementation` owns the internal layer and package architecture.

Public omission, null, warning, heading, and body representation are owned by `DRMCP-WORK-MCP-004`.
Exact diagnostic identifiers, severity, and source-location fields are owned by `DRMCP-WORK-MCP-006`.

## Related specs

| ref | relation |
|---|---|
| `spec:drmcp.design_records_mcp.schema.discovery` | Candidate discovery, invalid current spec behavior, and duplicate identity boundary. |
| `spec:drmcp.design_records_mcp.schema.fields` | Common and kind-specific parsed field vocabulary. |
| `spec:drmcp.design_records_mcp.schema.id_normalization` | Canonical identity mapping. |
| `spec:drmcp.design_records_mcp.namespace_scanning` | Configured current roots and active-index construction. |
| `spec:drmcp.implementation` | Internal state separation and package ownership. |

## Sources

- `DRMCP-TASK-MCP-003-02`: Duplicate identity and active-index separation decisions.
- `DRMCP-TASK-MCP-003-03`: Current spec invalid-source and identity decisions.
- `DRMCP-TASK-MCP-003-04`: Shared record integration decisions.
- `DRMCP-ADR-MCP-004`: Internal state and operation-contract separation.
