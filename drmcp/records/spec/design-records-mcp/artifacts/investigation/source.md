# Contract: Investigation source

- **id**: `spec:drmcp.design_records_mcp.artifacts.investigation.source`
- **status**: draft
- **date**: 2026-07-11
- **parent**: `spec:drmcp.design_records_mcp.artifacts.investigation`
- **contract_class**: `format`
- **usdm_covers**:
  - usdm:product.design_records.authoring_semantics.workflow_record_authoring#R006

## What this is

Defines the source-document shape for Investigation records.

## H1 prefix

- **type**: `identity`

### Value

```text
<APP_NAMESPACE>-INV-<DOMAIN_NAMESPACE>-<SEQUENCE>
```

## H2 heading policy

- **unlisted headings**: `prohibited`

## H2 headings

| heading | condition | format reference |
|---|---|---|
| `## Investigation scope` | `always` | `-` |
| `## Out of scope` | `always` | `-` |
| `## Background` | `always` | `-` |
| `## What was investigated` | `always` | `-` |
| `## Findings` | `always` | `-` |
| `## Cross-cutting observations` | `always` | `-` |
| `## Follow-up judgment candidates` | `always` | `-` |
| `## Recommendation` | `always` | `-` |
| `## Follow-up artifact candidates` | `always` | `-` |
| `## Open questions` | `always` | `-` |

## Related specs

| ref | relation |
|---|---|
| `spec:drmcp.design_records_mcp.artifacts.base.definitions.source` | Shared H1 shape, metadata placement, and H2 declaration rules. |
| `spec:drmcp.design_records_mcp.artifacts.investigation.identity_and_structure` | Defines the Investigation identity represented by the H1 prefix. |
| `spec:drmcp.design_records_mcp.artifacts.investigation.h1_adjacent_metadata` | Defines the Investigation metadata fields placed after H1. |
| `spec:product.design_records.authoring_standards.investigation_authoring` | Product authority for Investigation source headings and presence conditions. |
