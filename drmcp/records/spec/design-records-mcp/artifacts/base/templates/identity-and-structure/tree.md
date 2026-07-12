# Contract: Tree identity and structure template

- **id**: `spec:drmcp.design_records_mcp.artifacts.base.templates.identity_and_structure.tree`
- **status**: draft
- **date**: 2026-07-12
- **parent**: `spec:drmcp.design_records_mcp.artifacts.base.templates.identity_and_structure`
- **contract_class**: `format`

## What this is

Provides the identity, record-structure, and source-placement template for an artifact that uses the `tree` record structure.
The artifact directory is one literal directory name relative to `<APP_NAMESPACE>/records/`.

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
| `spec:drmcp.design_records_mcp.artifacts.base.definitions.record_structure` | Shared tree record-structure rules. |
| `spec:drmcp.design_records_mcp.artifacts.base.definitions.identity_declaration` | Shared tree identity form. |
| `<PRODUCT_AUTHORITY_REF>` | Product authority consumed by this artifact Specification. |
````
