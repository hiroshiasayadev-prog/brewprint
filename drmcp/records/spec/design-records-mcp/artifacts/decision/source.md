# Contract: Decision source

- **id**: `spec:drmcp.design_records_mcp.artifacts.decision.source`
- **status**: draft
- **date**: 2026-07-11
- **parent**: `spec:drmcp.design_records_mcp.artifacts.decision`
- **contract_class**: `format`
- **usdm_covers**:
  - usdm:product.design_records.authoring_semantics.workflow_record_authoring#R006

## What this is

Defines the source-document shape for Decision records.

## H1 prefix

- **type**: `identity`

### Value

```text
<APP_NAMESPACE>-ADR-<DOMAIN_NAMESPACE>-<SEQUENCE>
```

## H2 heading policy

- **unlisted headings**: `prohibited`

## H2 headings

| heading | condition | format reference |
|---|---|---|
| `## Context` | `always` | `-` |
| `## Decision` | `always` | `-` |
| `## Rationale` | `always` | `-` |
| `## Rejected alternatives` | `optional` | `-` |
| `## Consequences` | `always` | `-` |
| `## Evidence` | `always` | `-` |

## Related specs

| ref | relation |
|---|---|
| `spec:drmcp.design_records_mcp.artifacts.base.definitions.source` | Shared H1 shape, metadata placement, and H2 declaration rules. |
| `spec:drmcp.design_records_mcp.artifacts.decision.identity_and_structure` | Defines the canonical Decision identity represented by the H1 prefix. |
| `spec:drmcp.design_records_mcp.artifacts.decision.h1_adjacent_metadata` | Defines the Decision metadata fields placed after H1. |
| `spec:product.design_records.authoring_standards.adr_authoring` | Product authority for Decision source headings and presence conditions. |
