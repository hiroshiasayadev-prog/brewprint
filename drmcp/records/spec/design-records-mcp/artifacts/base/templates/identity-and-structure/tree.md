# Contract: Tree identity and structure template

- **id**: `spec:drmcp.design_records_mcp.artifacts.base.templates.identity_and_structure.tree`
- **status**: draft
- **date**: 2026-07-10
- **parent**: `spec:drmcp.design_records_mcp.artifacts.base.templates.identity_and_structure`
- **contract_class**: `format`

## What this is

Provides the template for an artifact-specific `identity-and-structure.md` Specification that uses the `tree` record structure.

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

## Related specs

| ref | relation |
|---|---|
| `spec:drmcp.design_records_mcp.artifacts.base.definitions.record_structure` | Shared tree record-structure rules. |
| `spec:drmcp.design_records_mcp.artifacts.base.definitions.identity_declaration` | Shared tree identity form. |
| `<PRODUCT_AUTHORITY_REF>` | Product authority consumed by this artifact Specification. |
````
