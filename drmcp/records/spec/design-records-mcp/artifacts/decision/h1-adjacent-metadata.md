# Contract: Decision H1-adjacent metadata

- **id**: `spec:drmcp.design_records_mcp.artifacts.decision.h1_adjacent_metadata`
- **status**: draft
- **date**: 2026-07-11
- **parent**: `spec:drmcp.design_records_mcp.artifacts.decision`
- **contract_class**: `format`
- **usdm_covers**:
  - usdm:product.design_records.authoring_semantics.workflow_record_authoring#R003,#R005

## What this is

Defines the H1-adjacent metadata fields for Decision records.

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

## Related specs

| ref | relation |
|---|---|
| `spec:drmcp.design_records_mcp.artifacts.base.definitions.h1_adjacent_metadata` | Shared H1-adjacent metadata notation, requirement values, value forms, and value types. |
| `spec:product.design_records.authoring_standards.adr_authoring` | Product authority for Decision metadata fields, lifecycle values, dates, and field-specific reference targets. |
| `spec:product.design_records.traceability.artifact_refs` | Product authority for the Decision canonical reference form. |
