# Contract: Task source

- **id**: `spec:drmcp.design_records_mcp.artifacts.task.source`
- **status**: draft
- **date**: 2026-07-11
- **parent**: `spec:drmcp.design_records_mcp.artifacts.task`
- **contract_class**: `format`
- **usdm_covers**:
  - usdm:product.design_records.authoring_semantics.workflow_record_authoring#R006-R008

## What this is

Defines the source-document shape for Task records.

## H1 prefix

- **type**: `identity`

### Value

```text
<APP_NAMESPACE>-TASK-<DOMAIN_NAMESPACE>-<WORK_SEQUENCE>-<TASK_SEQUENCE>
```

## H2 heading policy

- **unlisted headings**: `prohibited`

## H2 headings

| heading | condition | format reference |
|---|---|---|
| `## Goal` | `always` | `-` |
| `## Work` | `always` | `-` |
| `## Implementation contract` | task_type = implementation: always;<br/>task_type != implementation: prohibited; | `spec:product.design_records.authoring_standards.task_authoring` |
| `## Done condition` | `always` | `-` |
| `## Verification` | `always` | `-` |
| `## Evidence` | `always` | `-` |

## Related specs

| ref | relation |
|---|---|
| `spec:drmcp.design_records_mcp.artifacts.base.definitions.source` | Shared H1 shape, metadata placement, H1 prefix declaration, and H2 declaration rules. |
| `spec:drmcp.design_records_mcp.artifacts.task.identity_and_structure` | Defines the canonical Task identity represented by the H1 prefix. |
| `spec:drmcp.design_records_mcp.artifacts.task.h1_adjacent_metadata` | Defines the Task metadata fields placed after H1. |
| `spec:product.design_records.authoring_standards.task_authoring` | Product authority for Task source headings, presence conditions, and the Implementation contract table format. |
