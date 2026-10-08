# Contract: Investigation H1-adjacent metadata

- **id**: `spec:drcli.design_records_cli.artifacts.investigation.h1_adjacent_metadata`
- **status**: draft
- **date**: 2026-10-07
- **parent**: `spec:drcli.design_records_cli.artifacts.investigation`
- **contract_class**: `format`
- **usdm_covers**:
  - usdm:product.design_records.authoring_semantics.workflow_record_authoring#R003-R004
  - usdm:product.design_records.namespace_and_identity.workflow_artifact_identity#R014
  - usdm:product.design_records.traceability_and_relations.metadata_relation_schema#R003-R006

## What this is

Defines the H1-adjacent metadata fields for Investigation records.

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
| `status` | `mandatory` | `scalar` | `string` | One of `investigating`, `concluded`, or `superseded`. |
| `date` | `mandatory` | `scalar` | `string` | Strict `YYYY-MM-DD`. |
| `trigger` | `mandatory` | `scalar` | `string` | Non-empty string. |
| `scope` | `mandatory` | `scalar` | `string` | Non-empty string of one or two sentences. |
| `non_scope` | `mandatory` | `scalar` | `string` | Non-empty string of one or two sentences. |
| `source_refs` | `mandatory` | `indented_list` | `ref` | `spec`, `decision`, `investigation`, `requirement`, or `work_item`; an empty list is allowed; empty list items are prohibited. |
| `follow_up_candidates` | `mandatory` | `indented_list` | `string` | Non-empty human-readable candidate text or canonical-reference text for `spec`, `decision`, `investigation`, `requirement`, or `work_item`; Task ID text is prohibited; an empty list is allowed; empty list items are prohibited. |
| `supersedes` | `optional` | `indented_list` | `ref` | `investigation` only. |
| `related_requirements` | `optional` | `indented_list` | `ref` | `requirement` only. |
| `related_work_items` | `optional` | `indented_list` | `ref` | `work_item` only. |
| `related_adrs` | `optional` | `indented_list` | `ref` | `decision` only. |
| `related_specs` | `optional` | `indented_list` | `ref` | `spec` only. |
| `related_internal_design` | `optional` | `indented_list` | `string` | Non-empty string. |
| `related_coverage` | `optional` | `indented_list` | `string` | Non-empty string. |
| `follow_up_results` | `optional` | `indented_list` | `ref` | `spec`, `decision`, `investigation`, `requirement`, or `work_item`. |

Only the fields listed in this table may appear in Investigation metadata.
An Investigation record does not persist an `id` field.
If `id` appears, it is prohibited unlisted metadata and has no identity authority.

## Related specs

| ref | relation |
|---|---|
| `spec:drcli.design_records_cli.artifacts.base.definitions.h1_adjacent_metadata` | Shared H1-adjacent metadata notation, requirement values, value forms, and value types. |
| `spec:product.design_records.authoring_standards.investigation_authoring` | Product authority for Investigation metadata fields, lifecycle values, and field constraints. |
| `spec:product.design_records.traceability.artifact_refs` | Product authority for current canonical reference forms. |
| `spec:product.design_records.traceability.metadata_schema` | Product authority for Investigation relation fields and Task-ID exclusion. |
