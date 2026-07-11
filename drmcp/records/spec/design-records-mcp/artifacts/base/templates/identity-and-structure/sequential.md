# Contract: Sequential identity and structure template

- **id**: `spec:drmcp.design_records_mcp.artifacts.base.templates.identity_and_structure.sequential`
- **status**: draft
- **date**: 2026-07-11
- **parent**: `spec:drmcp.design_records_mcp.artifacts.base.templates.identity_and_structure`
- **contract_class**: `format`

## What this is

Provides the template for an artifact-specific `identity-and-structure.md` Specification that uses the `sequential` record structure.
The app namespace and domain namespace are supplied by the record location and are not listed in the artifact-specific segment table or replaced with artifact-specific literals.

`allocation scope` states the boundary within which a sequence value is allocated.
Use `domain` when the sequence is allocated within the current app namespace, artifact kind, and domain namespace.
When an artifact uses a different allocation scope, declare that artifact-specific scope explicitly.
Use `-` for segments that are not independently allocated.

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

- **structure**: `sequential`
- **artifact kind**: `<ARTIFACT_KIND_LITERAL>`

## Identity form

```text
<APP_NAMESPACE>-<ARTIFACT_KIND>-<DOMAIN_NAMESPACE>-<ARTIFACT_SPECIFIC_SEGMENTS...>
```

## Artifact-specific segments

| segment | role | format | sequence | allocation scope |
|---|---|---|---|---|
| `<SEGMENT_NAME>` | `<SEGMENT_ROLE>` | `<SEGMENT_FORMAT>` | `<yes-or-no>` | `<domain-or-artifact-specific-scope-or-dash>` |

## Related specs

| ref | relation |
|---|---|
| `spec:drmcp.design_records_mcp.artifacts.base.definitions.record_structure` | Shared sequential record-structure rules. |
| `spec:drmcp.design_records_mcp.artifacts.base.definitions.identity_declaration` | Shared sequential identity form. |
| `<PRODUCT_AUTHORITY_REF>` | Product authority consumed by this artifact Specification. |
````
