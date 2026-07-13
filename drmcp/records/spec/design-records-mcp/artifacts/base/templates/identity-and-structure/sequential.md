# Contract: Sequential identity and structure template

- **id**: `spec:drmcp.design_records_mcp.artifacts.base.templates.identity_and_structure.sequential`
- **status**: draft
- **date**: 2026-07-13
- **parent**: `spec:drmcp.design_records_mcp.artifacts.base.templates.identity_and_structure`
- **contract_class**: `format`

## What this is

Provides the identity, record-structure, and source-placement template for an artifact that uses the `sequential` record structure.
The app namespace and domain namespace are supplied by the record location and are not listed in the artifact-specific segment table or replaced with artifact-specific literals.
The artifact directory is one literal directory name relative to `<APP_NAMESPACE>/records/`.

`sequence` is `yes` for every segment that contains a sequence value, including an inherited sequence value.
Rows with `sequence: yes` appear in the priority order used to construct the sequential-listing ordering tuple.
The row order must remain consistent with the artifact-specific segment order in the identity form.

`allocation scope` states the boundary within which a sequence value is allocated.
Use `domain` when the sequence is allocated within the current app namespace, artifact kind, and domain namespace.
When an artifact uses a different allocation scope, declare that artifact-specific scope explicitly.
Use `-` for a sequence-bearing segment that the artifact does not allocate independently.

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

## Source placement

- **artifact directory**: `<ARTIFACT_DIRECTORY_LITERAL>`

## Identity form

```text
<APP_NAMESPACE>-<ARTIFACT_KIND>-<DOMAIN_NAMESPACE>-<ARTIFACT_SPECIFIC_SEGMENTS...>
```

## Identity authority

The complete public ID in the H1 prefix is the canonical identity authority for an `<ARTIFACT_NAME>` source.
It must agree with the configured app namespace, the `<ARTIFACT_KIND_LITERAL>` kind selected by this source corpus, and the physical domain directory.
H1-adjacent metadata and the file name do not supply or repair identity.
The file-name public-ID prefix is a nonidentity conformance projection of the H1 public ID.

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
