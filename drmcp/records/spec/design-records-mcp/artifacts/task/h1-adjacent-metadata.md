# Contract: Task H1-adjacent metadata

- **id**: `spec:drmcp.design_records_mcp.artifacts.task.h1_adjacent_metadata`
- **status**: draft
- **date**: 2026-07-11
- **parent**: `spec:drmcp.design_records_mcp.artifacts.task`
- **contract_class**: `format`
- **usdm_covers**:
  - usdm:product.design_records.authoring_semantics.workflow_record_authoring#R003-R004,#R009
  - usdm:product.design_records.traceability_and_relations.workflow_relations#R009

## What this is

Defines the H1-adjacent metadata fields for Task records.
A field whose presence is rule-dependent uses `reference` and names the Specification that defines its presence conditions and field-specific constraints.

## Fields

| field | requirement | form | value type | value format |
|---|---|---|---|---|
| `id` | `mandatory` | `scalar` | `ref` | Permitted record kind: `task`; `<APP_NAMESPACE>-TASK-<DOMAIN_NAMESPACE>-<WORK_SEQUENCE>-<TASK_SEQUENCE>`; `<WORK_SEQUENCE>` is a three-digit, zero-padded decimal; `<TASK_SEQUENCE>` is a two-digit, zero-padded decimal; the value matches the Task public ID in H1 and the file name. |
| `status` | `mandatory` | `scalar` | `string` | One of `not_started`, `in_progress`, `blocked`, `done`, or `cancelled`. |
| `date` | `mandatory` | `scalar` | `string` | Strict `YYYY-MM-DD`. |
| `work_item` | `mandatory` | `scalar` | `ref` | Permitted artifact kind: `work_item`. |
| `work_item_ref` | `reference` | `scalar` | `ref` | `spec:product.design_records.authoring_standards.task_authoring` |
| `task_type` | `mandatory` | `scalar` | `string` | One of `investigation`, `decision`, `authoring`, `implementation`, `review`, `correction`, `verification`, `coordination`, `work_item_decomposition`, `work_item_execution`, or `synchronization`. |
| `estimate` | `mandatory` | `scalar` | `string` | Non-empty string. |
| `depends_on` | `mandatory` | `indented_list` | `ref` | Permitted artifact kind: `task`; zero or more items; empty child items are prohibited. |
| `outputs` | `mandatory` | `indented_list` | `ref_or_literal` | Permitted artifact kinds: `spec`, `decision`, `investigation`, `requirement`, `work_item`, `task`; otherwise a non-empty string; zero or more items; empty child items are prohibited. |

Only the fields listed in this table may appear in Task metadata.
Task metadata must not contain `source_requirement`, `source_refs`, or any other source-provenance field.

## Related specs

| ref | relation |
|---|---|
| `spec:drmcp.design_records_mcp.artifacts.base.definitions.h1_adjacent_metadata` | Shared H1-adjacent metadata notation, requirement values, value forms, and value types. |
| `spec:drmcp.design_records_mcp.artifacts.task.identity_and_structure` | Task public ID format used by `id`. |
| `spec:product.design_records.authoring_standards.task_authoring` | Product authority for Task metadata fields, values, conditional presence, and field constraints. |
| `spec:product.design_records.traceability.metadata_schema` | Product authority for persisted Task relation fields. |
| `spec:product.design_records.traceability.artifact_refs` | Product authority for permitted record kinds and canonical reference forms. |
