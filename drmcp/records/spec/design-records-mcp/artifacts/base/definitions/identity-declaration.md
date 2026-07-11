# Reference: Identity declaration

- **id**: `spec:drmcp.design_records_mcp.artifacts.base.definitions.identity_declaration`
- **status**: draft
- **date**: 2026-07-10
- **parent**: `spec:drmcp.design_records_mcp.artifacts.base`

## What this is

Defines the shared identity declaration forms used by DRMCP artifact Specifications.

This Specification defines the common form for each record structure and the artifact-specific values that each artifact Specification must declare.
Product Specifications remain authoritative for canonical identity and reference semantics.

## Current contract

Each artifact Specification must declare its identity by using the form for its selected record structure.

| record structure | identity form |
|---|---|
| `sequential` | `<APP_NAMESPACE>-<ARTIFACT_KIND>-<DOMAIN_NAMESPACE>-<ARTIFACT_SPECIFIC_SEGMENTS...>` |
| `tree` | `<RECORD_KIND>:<APP_NAMESPACE>(.<PATH_SEGMENT>)*` |

## Sequential identity declaration

The sequential identity form contains these shared segments in this order:

```text
<APP_NAMESPACE>-<ARTIFACT_KIND>-<DOMAIN_NAMESPACE>-<ARTIFACT_SPECIFIC_SEGMENTS...>
```

The shared form requires:

- `<APP_NAMESPACE>` as the first segment.
- `<ARTIFACT_KIND>` after the app namespace.
- `<DOMAIN_NAMESPACE>` after the artifact kind.
- Hyphen separators between segments.
- Artifact-specific segments only after the domain namespace.

Each sequential artifact Specification must declare:

| declaration | requirement |
|---|---|
| artifact kind | Declare the literal value used for `<ARTIFACT_KIND>`. |
| artifact-specific segments | Declare every segment after `<DOMAIN_NAMESPACE>`. |
| segment order | Declare the complete order of artifact-specific segments. |
| segment role | State the role of each artifact-specific segment. |
| sequence segment | Identify every artifact-specific segment that receives a sequential value. |

Sequence allocation scope is declared under the sequential rules defined by `spec:drmcp.design_records_mcp.artifacts.base.definitions.record_structure`.

## Tree identity declaration

The tree identity form is:

```text
<RECORD_KIND>:<APP_NAMESPACE>(.<PATH_SEGMENT>)*
```

The shared form requires:

- A `<RECORD_KIND>:` prefix.
- `<APP_NAMESPACE>` as the first identity segment.
- Dot-separated path segments after the app namespace.
- `<RECORD_KIND>:<APP_NAMESPACE>` as the identity root.
- No additional path segment for an `index.md` file that represents its containing directory node.
- The non-index file stem as the final path segment for a leaf record.

Each tree artifact Specification must declare only the literal value used for `<RECORD_KIND>`.

The identity root is derived as `<RECORD_KIND>:<APP_NAMESPACE>`.
The source root is the artifact-kind directory under `<APP_NAMESPACE>/records/`.
Directory names and non-index Markdown file stems map to `<PATH_SEGMENT>` values under the shared tree rules above.
These fixed tree mappings are not redeclared by each artifact Specification.

## Boundary

| concern | owner |
|---|---|
| Shared sequential and tree identity declaration forms | This Specification. |
| Artifact-specific kind values and sequential segments | The artifact-specific identity Specification. |
| Fixed tree identity root, source root, and path mapping | This Specification. |
| Sequence allocation scope | `spec:drmcp.design_records_mcp.artifacts.base.definitions.record_structure` and the artifact-specific Specification. |
| Relation-dependent segment agreement | The artifact-specific relation or identity-validation Specification. |
| Canonical identity and reference semantics | Product authority. |
| Parser behavior, invalid identity handling, duplicate handling, and diagnostics | Other DRMCP Specifications. |

## Related specs

| ref | relation |
|---|---|
| `spec:drmcp.design_records_mcp.artifacts.base.definitions.record_structure` | Defines supported record structures and sequential allocation declarations. |
| `spec:product.design_records.namespace_model.artifact_id_grammar` | Product authority for sequential artifact ID grammar. |
| `spec:product.design_records.spec_format.spec_id_as_ref` | Product authority for path-derived Specification identity. |
| `spec:product.design_records.traceability.artifact_refs` | Product authority for record kinds and canonical reference forms. |
