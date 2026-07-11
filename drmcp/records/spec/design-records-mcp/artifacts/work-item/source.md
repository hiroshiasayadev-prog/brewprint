# Contract: Work Item source

- **id**: `spec:drmcp.design_records_mcp.artifacts.work_item.source`
- **status**: draft
- **date**: 2026-07-11
- **parent**: `spec:drmcp.design_records_mcp.artifacts.work_item`
- **contract_class**: `format`
- **usdm_covers**:
  - usdm:product.design_records.authoring_semantics.workflow_record_authoring#R001-R002,#R006

## What this is

Defines the source-document shape for Work Item records.

## H1 prefix

- **type**: `identity`

### Value

```text
<APP_NAMESPACE>-WORK-<DOMAIN_NAMESPACE>-<SEQUENCE>
```

## H2 heading policy

- **unlisted headings**: `prohibited`

## H2 headings

| heading | condition | format reference |
|---|---|---|
| `## Goal` | `always` | `-` |
| `## Boundary` | `always` | `-` |
| `## Impact Scope` | `always` | `-` |
| `## Task flow` | `always` | `-` |
| `## Task Candidates` | `always` | `-` |
| `## Completion Condition` | `always` | `-` |
| `## Evidence` | `always` | `-` |

## Related specs

| ref | relation |
|---|---|
| `spec:drmcp.design_records_mcp.artifacts.base.definitions.source` | Shared H1 shape, metadata placement, and H2 declaration rules. |
| `spec:drmcp.design_records_mcp.artifacts.work_item.identity_and_structure` | Defines the Work Item identity represented by the H1 prefix. |
| `spec:drmcp.design_records_mcp.artifacts.work_item.h1_adjacent_metadata` | Defines the Work Item metadata fields placed after H1. |
| `spec:product.design_records.authoring_standards.work_item_authoring` | Product authority for Work Item source headings and presence conditions. |
