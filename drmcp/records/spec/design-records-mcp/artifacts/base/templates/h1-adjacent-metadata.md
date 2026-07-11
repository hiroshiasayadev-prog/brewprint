# Contract: H1-adjacent metadata template

- **id**: `spec:drmcp.design_records_mcp.artifacts.base.templates.h1_adjacent_metadata`
- **status**: draft
- **date**: 2026-07-10
- **parent**: `spec:drmcp.design_records_mcp.artifacts.base.templates`
- **contract_class**: `format`

## What this is

Provides the template for an artifact-specific `h1-adjacent-metadata.md` Specification.
A field whose presence depends on another field or artifact-specific rule uses `reference` and names a separate Specification that defines the presence conditions and field-specific constraints.
`value format` lists permitted artifact kinds and concise implementation-checkable constraints directly.
Use a Specification ref for value format only when one field accepts a complex format spanning multiple value types.

## Template

````markdown
# Contract: <ARTIFACT_NAME> H1-adjacent metadata

- **id**: `<SPEC_REF>`
- **status**: draft
- **date**: `<YYYY-MM-DD>`
- **parent**: `<PARENT_SPEC_REF>`
- **contract_class**: `format`

## What this is

Defines the H1-adjacent metadata fields for `<ARTIFACT_NAME>` records.
A field whose presence is rule-dependent uses `reference` and names the Specification that defines its presence conditions and field-specific constraints.

## Fields

| field | requirement | form | value type | value format |
|---|---|---|---|---|
| `<FIELD_NAME>` | `<mandatory-or-optional-or-reference>` | `<scalar-or-inline_list-or-indented_list>` | `<string-or-ref-or-ref_or_literal>` | `<EXPLICIT_VALUE_FORMAT_OR_COMPLEX_FORMAT_SPEC_REF>` |

## Related specs

| ref | relation |
|---|---|
| `spec:drmcp.design_records_mcp.artifacts.base.definitions.h1_adjacent_metadata` | Shared H1-adjacent metadata notation, value forms, and value types. |
| `<PRODUCT_AUTHORITY_REF>` | Product authority consumed by this artifact Specification. |
````
