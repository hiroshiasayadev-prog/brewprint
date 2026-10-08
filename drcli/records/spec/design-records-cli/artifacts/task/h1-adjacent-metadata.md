# Contract: Task H1-adjacent metadata

- **id**: `spec:drcli.design_records_cli.artifacts.task.h1_adjacent_metadata`
- **status**: draft
- **date**: 2026-10-07
- **parent**: `spec:drcli.design_records_cli.artifacts.task`
- **contract_class**: `format`
- **usdm_covers**:
  - usdm:product.design_records.authoring_semantics.workflow_record_authoring#R003-R004,#R009
  - usdm:product.design_records.traceability_and_relations.workflow_relations#R009

## What this is

Defines the H1-adjacent metadata fields for Task records.
A field whose presence is rule-dependent uses `reference` and names the Specification that defines its presence conditions and field-specific constraints.

## Current contract

This Specification defines the DRCLI-local projection and declarations for this topic.
Where this Specification cites PRODUCT authority, the cited PRODUCT Specification remains authoritative for PRODUCT-owned semantics.

## Rules

- DRCLI-local declarations must preserve cited PRODUCT-owned semantics.
- A DRCLI-local declaration must not create an independent replacement for cited PRODUCT authority.

## Validation rules

- A DRCLI interpretation that conflicts with cited PRODUCT authority is nonconforming.
- Validation of PRODUCT-owned meaning uses the cited PRODUCT authority.

## Fields

| field | requirement | form | value type | value format |
|---|---|---|---|---|
| `status` | `mandatory` | `scalar` | `string` | One of `not_started`, `in_progress`, `blocked`, `done`, or `cancelled`. |
| `date` | `mandatory` | `scalar` | `string` | Strict `YYYY-MM-DD`. |
| `work_item` | `mandatory` | `scalar` | `ref` | Permitted artifact kind: `work_item`. |
| `work_item_ref` | `reference` | `scalar` | `ref` | `spec:product.design_records.authoring_standards.task_authoring` |
| `task_type` | `mandatory` | `scalar` | `string` | One of `investigation`, `decision`, `authoring`, `implementation`, `review`, `correction`, `verification`, `coordination`, `work_item_decomposition`, `work_item_execution`, or `synchronization`. |
| `estimate` | `mandatory` | `scalar` | `string` | Non-empty string. |
| `depends_on` | `mandatory` | `indented_list` | `ref` | Permitted artifact kind: `task`; zero or more items; empty child items are prohibited. |
| `outputs` | `mandatory` | `indented_list` | `ref_or_literal` | Permitted artifact kinds: `spec`, `decision`, `investigation`, `requirement`, `work_item`, `task`; otherwise a non-empty string; zero or more items; empty child items are prohibited. |

Only the fields listed in this table may appear in Task metadata.
A Task record does not persist an `id` field.
If `id` appears, it is prohibited unlisted metadata and has no identity authority.
Task metadata must not contain `source_requirement`, `source_refs`, or any other source-provenance field.

## Related specs

| ref | relation |
|---|---|
| `spec:drcli.design_records_cli.artifacts.base.definitions.h1_adjacent_metadata` | Shared H1-adjacent metadata notation, requirement values, value forms, and value types. |
| `spec:drcli.design_records_cli.artifacts.task.identity_and_structure` | Defines the Task public ID grammar represented by the H1 prefix. |
| `spec:product.design_records.authoring_standards.task_authoring` | Product authority for Task metadata fields, values, conditional presence, and field constraints. |
| `spec:product.design_records.traceability.metadata_schema` | Product authority for persisted Task relation fields. |
| `spec:product.design_records.traceability.artifact_refs` | Product authority for permitted record kinds and canonical reference forms. |
