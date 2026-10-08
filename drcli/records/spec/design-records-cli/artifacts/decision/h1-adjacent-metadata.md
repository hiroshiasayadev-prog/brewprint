# Contract: Decision H1-adjacent metadata

- **id**: `spec:drcli.design_records_cli.artifacts.decision.h1_adjacent_metadata`
- **status**: draft
- **date**: 2026-10-07
- **parent**: `spec:drcli.design_records_cli.artifacts.decision`
- **contract_class**: `format`
- **usdm_covers**:
  - usdm:product.design_records.authoring_semantics.workflow_record_authoring#R003-R004

## What this is

Defines the H1-adjacent metadata fields for Decision records.

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
| `status` | `mandatory` | `scalar` | `string` | One of `proposed`, `accepted`, or `superseded`. |
| `date` | `mandatory` | `scalar` | `string` | Strict `YYYY-MM-DD`. |
| `depends_on` | `mandatory` | `indented_list` | `ref` | `decision` only; a field marker with no indented child items represents an empty list; empty child list items are prohibited. |
| `supersedes` | `mandatory` | `indented_list` | `ref` | `decision` only; a field marker with no indented child items represents an empty list; empty child list items are prohibited. |
| `migrated_to_spec` | `mandatory` | `scalar` | `string` | Strict `YYYY-MM-DD` or the literal `null`. |

Only the fields listed in this table may appear in Decision metadata.
A Decision record does not persist an `id` field.
If `id` appears, it is prohibited unlisted metadata and has no identity authority.

## Related specs

| ref | relation |
|---|---|
| `spec:drcli.design_records_cli.artifacts.base.definitions.h1_adjacent_metadata` | Shared H1-adjacent metadata notation, requirement values, value forms, and value types. |
| `spec:product.design_records.authoring_standards.adr_authoring` | Product authority for Decision metadata fields, lifecycle values, dates, and field-specific reference targets. |
| `spec:product.design_records.traceability.artifact_refs` | Product authority for the Decision canonical reference form. |
