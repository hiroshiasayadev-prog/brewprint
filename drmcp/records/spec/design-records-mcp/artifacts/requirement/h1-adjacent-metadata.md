# Contract: Requirement H1-adjacent metadata

- **id**: `spec:drmcp.design_records_mcp.artifacts.requirement.h1_adjacent_metadata`
- **status**: draft
- **date**: 2026-07-11
- **parent**: `spec:drmcp.design_records_mcp.artifacts.requirement`
- **contract_class**: `format`
- **usdm_covers**:
  - usdm:product.design_records.authoring_semantics.workflow_record_authoring#R003,#R004

## What this is

Defines the H1-adjacent metadata fields for Requirement records.

## Fields

| field | requirement | form | value type | value format |
|---|---|---|---|---|
| `id` | `mandatory` | `scalar` | `ref` | Permitted record kind: `requirement`; `<APP_NAMESPACE>-REQ-<DOMAIN_NAMESPACE>-<SEQUENCE>`; `<SEQUENCE>` is a three-digit, zero-padded decimal. |
| `status` | `optional` | `scalar` | `string` | Any string. |
| `date` | `optional` | `scalar` | `string` | Any string. |
| `source_refs` | `optional` | `indented_list` | `ref` | Permitted record kinds: `spec`, `decision`, `investigation`, `requirement`, `work_item`, and `task`. |

Only the fields listed in this table may appear in Requirement metadata.

## Related specs

| ref | relation |
|---|---|
| `spec:drmcp.design_records_mcp.artifacts.base.definitions.h1_adjacent_metadata` | Shared H1-adjacent metadata notation, value forms, and value types. |
| `spec:drmcp.design_records_mcp.artifacts.requirement.identity_and_structure` | Defines the Requirement public ID used by `id`. |
| `spec:product.design_records.authoring_standards.requirement_authoring` | Product authority for the Requirement metadata field allow-list. |
| `spec:product.design_records.traceability.artifact_refs` | Product authority for canonical reference forms allowed in `source_refs`. |
