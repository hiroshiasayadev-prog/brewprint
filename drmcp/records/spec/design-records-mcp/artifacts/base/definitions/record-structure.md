# Reference: Record structure

- **id**: `spec:drmcp.design_records_mcp.artifacts.base.definitions.record_structure`
- **status**: draft
- **date**: 2026-07-10
- **parent**: `spec:drmcp.design_records_mcp.artifacts.base`

## What this is

Defines the record-structure choices used by DRMCP artifact Specifications.

This Specification defines shared structure vocabulary and the additional declarations required from each artifact kind.
It does not redefine Product-owned identity, relation, or canonical-reference semantics.

## Current contract

Each artifact kind must select exactly one supported record structure.

| structure | meaning |
|---|---|
| `sequential` | Records are placed by domain and use one or more sequence segments in their public IDs. |
| `tree` | The directory and file hierarchy represents the logical record tree. |

## Sequential structure

A sequential structure follows these rules:

- The artifact kind directory contains one domain-directory level.
- Record files are placed directly under the domain directory.
- Directories below the domain directory are not part of the standard sequential structure.
- Directories below the domain directory do not add public ID segments.

An artifact kind that selects `sequential` must declare:

| declaration | requirement |
|---|---|
| sequence segment | Identify every public ID segment that receives a sequential value. |
| allocation scope | Identify the boundary within which each declared sequence segment is allocated. |

Use `domain` when a sequence is allocated within the current app namespace, artifact kind, and domain namespace.
When an artifact uses a different allocation scope, the artifact-specific Specification declares that scope explicitly.
Use `-` for a segment that is not independently allocated.

The artifact-specific declaration does not define relation-dependent segment agreement or inherited-segment rules.
Those rules belong to the artifact-specific relation or identity-validation Specification.

This base definition does not prescribe sequence names, segment count, width, or artifact-specific exception scopes.
The artifact-specific Specification consumes the applicable Product authority for those values.

## Tree structure

A tree structure follows these rules:

- The directory and file hierarchy represents the logical record tree.
- An `index.md` file represents its containing directory node.
- A non-index Markdown file represents a leaf node.

An artifact kind that selects `tree` has no additional structure-specific declaration.
Root identity, source root, and path-to-identity mapping follow the shared tree identity rules and are not redeclared by each artifact Specification.

## Boundary

| concern | owner |
|---|---|
| Record-structure vocabulary | This Specification. |
| Required heading and table shape for artifact declarations | Base template Specifications. |
| Exact public ID grammar and canonical-reference semantics | Product authority. |
| Artifact-specific identity mapping | The artifact-specific identity Specification. |
| Relation-dependent segment agreement | The artifact-specific relation or identity-validation Specification. |
| Parser, discovery, indexing, and validation orchestration | Other DRMCP Specifications. |

## Related specs

| ref | relation |
|---|---|
| `spec:product.design_records.repository_layout.record_discovery_paths` | Product authority for kind-specific placement patterns. |
| `spec:product.design_records.namespace_model.artifact_id_grammar` | Product authority for sequential artifact ID grammar and allocation scopes. |
| `spec:product.design_records.spec_format.spec_id_as_ref` | Product authority for path-derived Specification identity. |
| `spec:product.design_records.traceability.artifact_refs` | Product authority for record kinds and canonical reference forms. |
