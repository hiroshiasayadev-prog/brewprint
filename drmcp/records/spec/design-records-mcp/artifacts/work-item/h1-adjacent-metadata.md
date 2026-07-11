# Contract: Work Item H1-adjacent metadata

- **id**: `spec:drmcp.design_records_mcp.artifacts.work_item.h1_adjacent_metadata`
- **status**: draft
- **date**: 2026-07-11
- **parent**: `spec:drmcp.design_records_mcp.artifacts.work_item`
- **contract_class**: `format`
- **usdm_covers**:
  - usdm:product.design_records.authoring_semantics.workflow_record_authoring#R003-R004,#R010
  - usdm:product.design_records.traceability_and_relations.workflow_relations#R001-R006,#R011

## What this is

Defines the H1-adjacent metadata fields for Work Item records.

## Fields

| field | requirement | form | value type | value format |
|---|---|---|---|---|
| `id` | `mandatory` | `scalar` | `ref` | Permitted record kind: `work_item`; `<APP_NAMESPACE>-WORK-<DOMAIN_NAMESPACE>-<SEQUENCE>`; `<SEQUENCE>` is a three-digit, zero-padded decimal; the value matches the Work Item public ID in H1 and the file name. |
| `status` | `mandatory` | `scalar` | `string` | One of `not_started`, `in_progress`, `blocked`, `done`, or `cancelled`. |
| `date` | `mandatory` | `scalar` | `string` | Strict `YYYY-MM-DD`. |
| `source_refs` | `mandatory` | `indented_list` | `ref` | Permitted artifact kinds: `spec`, `decision`, `investigation`, `requirement`, `work_item`, `task`; one or more items; duplicate items, empty items, and self-reference are prohibited; item order has no semantic meaning. |
| `impact_refs` | `mandatory` | `indented_list` | `ref` | Permitted artifact kinds: `spec`, `decision`, `investigation`, `requirement`, `work_item`, `task`; zero or more items; duplicate and empty items are prohibited. |
| `tasks` | `mandatory` | `indented_list` | `ref` | Permitted artifact kind: `task`; zero or more items; duplicate and empty items are prohibited. |

Only the fields listed in this table may appear in Work Item metadata.
`source_refs` is the only persisted source-provenance field for a Work Item.

## Related specs

| ref | relation |
|---|---|
| `spec:drmcp.design_records_mcp.artifacts.base.definitions.h1_adjacent_metadata` | Shared H1-adjacent metadata notation, requirement values, value forms, and value types. |
| `spec:drmcp.design_records_mcp.artifacts.work_item.identity_and_structure` | Work Item public ID format used by `id`. |
| `spec:product.design_records.authoring_standards.work_item_authoring` | Product authority for Work Item metadata fields, values, and constraints. |
| `spec:product.design_records.traceability.metadata_schema` | Product authority for persisted Work Item relation fields. |
| `spec:product.design_records.traceability.artifact_refs` | Product authority for permitted record kinds and reference forms. |
