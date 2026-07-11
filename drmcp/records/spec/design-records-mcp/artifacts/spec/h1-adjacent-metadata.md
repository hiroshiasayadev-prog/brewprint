# Contract: Specification H1-adjacent metadata

- **id**: `spec:drmcp.design_records_mcp.artifacts.spec.h1_adjacent_metadata`
- **status**: draft
- **date**: 2026-07-11
- **parent**: `spec:drmcp.design_records_mcp.artifacts.spec`
- **contract_class**: `format`
- **usdm_covers**:
  - usdm:product.design_records.spec_document_format.metadata_shape#R002-R005

## What this is

Defines the H1-adjacent metadata fields for Specification records.
A field whose presence is rule-dependent uses `reference` and names the Specification that defines its presence conditions and field-specific constraints.

## Fields

| field | requirement | form | value type | value format |
|---|---|---|---|---|
| `id` | `mandatory` | `scalar` | `ref` | Permitted record kind: `spec`; `spec:<APP_NAMESPACE>(.<PATH_SEGMENT>)*`; the value matches the path-derived canonical Specification ref of the source file. |
| `status` | `mandatory` | `scalar` | `string` | Non-empty string. |
| `date` | `mandatory` | `scalar` | `string` | Strict `YYYY-MM-DD`. |
| `parent` | `mandatory` | `scalar` | `ref_or_literal` | Permitted record kind: `spec`; permitted literals: `root` and `-`. |
| `contract_class` | `reference` | `scalar` | `string` | `spec:product.design_records.spec_format.document_shape`; required when the H1 prefix is `Contract` and prohibited otherwise; one of `interface` or `format`. |
| `usdm_covers` | `optional` | `indented_list` | `ref` | Permitted record kind: `usdm`; each item is a full USDM requirement ID or a compact row-list expression using `#RNNN`, `RNNN`, or ascending `RNNN-RNNN` row tokens after one USDM record ID; one or more items; empty child items and duplicate expanded requirement IDs are prohibited. |

## Related specs

| ref | relation |
|---|---|
| `spec:drmcp.design_records_mcp.artifacts.base.definitions.h1_adjacent_metadata` | Shared H1-adjacent metadata notation, requirement values, value forms, and value types. |
| `spec:product.design_records.spec_format.document_shape` | Product authority for Specification metadata fields and `contract_class` conditions. |
| `spec:product.design_records.spec_format.spec_id_as_ref` | Product authority for Specification `id` and `parent` values. |
| `spec:product.design_records.usdm.coverage_format` | Product authority for optional `usdm_covers` metadata. |
