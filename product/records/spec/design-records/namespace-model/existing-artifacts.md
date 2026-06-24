# Reference: New artifact namespace ownership

- **id**: `spec:product.design_records.namespace_model.existing_artifacts`
- **status**: draft
- **date**: 2026-06-24
- **parent**: `spec:product.design_records.namespace_model`

## What this is

Defines the app-independent rule for selecting the namespace of new Design Records artifacts.

Historical Brewprint ownership decisions, effective attribution, and V01 migration state are compatibility facts.
They are recorded in `spec:product.brewprint.compatibility.existing_artifacts`.

## Current contract

For new sequential records, use the app-aware artifact ID form when the owning app namespace is confirmed.
Use the product-level namespace when the concern is cross-app or attribution is unclear.

| condition | namespace selection |
|---|---|
| A single app namespace owns the concern. | Use the owning app namespace. |
| A product-level policy or governance concern spans apps. | Use the product-level namespace. |
| Ownership is unclear at authoring time. | Use the product-level namespace until an accepted decision assigns ownership. |
| The record is a spec. | Use the path-derived `spec:` identity contract instead of a sequential workflow ID. |

## Rules

- New workflow artifacts use the canonical artifact ID grammar.
- The app namespace segment identifies the owner.
- The domain segment identifies the owner-local concern.
- Effective attribution for legacy IDs does not create a new ID or alias.

## Boundary

| content | owner |
|---|---|
| Generic new-artifact ownership selection | This spec. |
| Canonical artifact ID grammar | `spec:product.design_records.namespace_model.artifact_id_grammar`. |
| Current Brewprint domain assignments | `spec:product.brewprint.namespaces.domain_catalog`. |
| Brewprint V01 ownership and effective attribution | `spec:product.brewprint.compatibility.existing_artifacts`. |
| UI, MCP, projection, or tool display behavior | DRMCP app-local specifications. |

## Related specs

| ref | relation |
|---|---|
| `spec:product.design_records.namespace_model` | Parent namespace model overview. |
| `spec:product.design_records.namespace_model.artifact_id_grammar` | Canonical grammar for new workflow artifact IDs. |
| `spec:product.brewprint.compatibility.existing_artifacts` | Brewprint historical ownership and effective attribution. |
