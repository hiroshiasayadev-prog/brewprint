# Contract: Tree identity and structure template

- **id**: `spec:drcli.design_records_cli.artifacts.base.templates.identity_and_structure.tree`
- **status**: draft
- **date**: 2026-10-08
- **parent**: `spec:drcli.design_records_cli.artifacts.base.templates.identity_and_structure`
- **contract_class**: `format`

## What this is

Provides the identity, record-structure, and source-placement template for an artifact that uses the `tree` record structure.
The artifact directory is one literal directory name relative to `<RESOLVED_RECORDS_ROOT>/`.

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
# Contract: <ARTIFACT_NAME> identity and structure

- **id**: `<SPEC_REF>`
- **status**: draft
- **date**: `<YYYY-MM-DD>`
- **parent**: `<PARENT_SPEC_REF>`
- **contract_class**: `format`

## What this is

Defines the identity and record structure for `<ARTIFACT_NAME>` records.

## Record structure

- **structure**: `tree`
- **record kind**: `<RECORD_KIND_LITERAL>`

## Source placement

- **artifact directory**: `<ARTIFACT_DIRECTORY_LITERAL>`

## Related specs

| ref | relation |
|---|---|
| `spec:drcli.design_records_cli.artifacts.base.definitions.record_structure` | Shared tree record-structure rules. |
| `spec:drcli.design_records_cli.artifacts.base.definitions.identity_declaration` | Shared tree identity form. |
| `<PRODUCT_AUTHORITY_REF>` | Product authority consumed by this artifact Specification. |
````
