# Contract: H1-adjacent metadata template

- **id**: `spec:drcli.design_records_cli.artifacts.base.templates.h1_adjacent_metadata`
- **status**: draft
- **date**: 2026-10-07
- **parent**: `spec:drcli.design_records_cli.artifacts.base.templates`
- **contract_class**: `format`

## What this is

Provides the template for an artifact-specific `h1-adjacent-metadata.md` Specification.
A field whose presence depends on another field or artifact-specific rule uses `reference` and names a separate Specification that defines the presence conditions and field-specific constraints.
`value format` lists permitted artifact kinds and concise implementation-checkable constraints directly.
Use a Specification ref for value format only when one field accepts a complex format spanning multiple value types.
For a sequential artifact, do not declare `id` in the field table and include the explicit prohibited-field boundary shown below.
A tree artifact may declare visible `id` only when its artifact contract uses that field as a projection checked against path-derived identity.

## Current contract

This Specification defines the DRCLI-local projection and declarations for this topic.
Where this Specification cites PRODUCT authority, the cited PRODUCT Specification remains authoritative for PRODUCT-owned semantics.

## Rules

- DRCLI-local declarations must preserve cited PRODUCT-owned semantics.
- A DRCLI-local declaration must not create an independent replacement for cited PRODUCT authority.

## Validation rules

- A DRCLI interpretation that conflicts with cited PRODUCT authority is nonconforming.
- Validation of PRODUCT-owned meaning uses the cited PRODUCT authority.

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

<FIELD_BOUNDARY_STATEMENTS>

## Related specs

| ref | relation |
|---|---|
| `spec:drcli.design_records_cli.artifacts.base.definitions.h1_adjacent_metadata` | Shared H1-adjacent metadata notation, value forms, and value types. |
| `<PRODUCT_AUTHORITY_REF>` | Product authority consumed by this artifact Specification. |
````

For a sequential artifact, replace `<FIELD_BOUNDARY_STATEMENTS>` with:

```markdown
Only the fields listed in this table may appear in `<ARTIFACT_NAME>` metadata.
An `<ARTIFACT_NAME>` record does not persist an `id` field.
If `id` appears, it is prohibited unlisted metadata and has no identity authority.
```

For a tree artifact, replace the placeholder with its artifact-specific field boundary.
