# Contract: Requirement H1-adjacent metadata

- **id**: `spec:drcli.design_records_cli.artifacts.requirement.h1_adjacent_metadata`
- **status**: draft
- **date**: 2026-10-07
- **parent**: `spec:drcli.design_records_cli.artifacts.requirement`
- **contract_class**: `format`
- **usdm_covers**:
  - usdm:product.design_records.authoring_semantics.workflow_record_authoring#R003,#R004

## What this is

Defines the H1-adjacent metadata fields for Requirement records.

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
| `status` | `optional` | `scalar` | `string` | Any string. |
| `date` | `optional` | `scalar` | `string` | Any string. |
| `source_refs` | `optional` | `indented_list` | `ref` | Permitted record kinds: `spec`, `decision`, `investigation`, `requirement`, `work_item`, and `task`. |

Only the fields listed in this table may appear in Requirement metadata.
A Requirement record does not persist an `id` field.
If `id` appears, it is prohibited unlisted metadata and has no identity authority.

## Related specs

| ref | relation |
|---|---|
| `spec:drcli.design_records_cli.artifacts.base.definitions.h1_adjacent_metadata` | Shared H1-adjacent metadata notation, value forms, and value types. |
| `spec:drcli.design_records_cli.artifacts.requirement.identity_and_structure` | Defines the Requirement public ID grammar represented by the H1 prefix. |
| `spec:product.design_records.authoring_standards.requirement_authoring` | Product authority for the Requirement metadata field allow-list. |
| `spec:product.design_records.traceability.artifact_refs` | Product authority for canonical reference forms allowed in `source_refs`. |
