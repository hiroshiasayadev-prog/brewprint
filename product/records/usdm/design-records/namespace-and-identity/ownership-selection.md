# USDM requirement: Ownership selection

- **id**: `usdm:product.design_records.namespace_and_identity.ownership_selection`
- **status**: draft
- **date**: 2026-07-10
- **kind**: requirement
- **parent**: `usdm:product.design_records.namespace_and_identity`

## What this is

Defines Product namespace ownership rules that affect current Design Record identity.

This record includes owner-segment semantics and the Product/implementation ownership boundary.
This record does not include authoring-time namespace selection policy, legacy attribution migration, or concrete app/domain lists.

## Requirements: spec:product.design_records.namespace_model.app_namespaces

| id | requirement | notes |
|---|---|---|
| R004 | Implementations MUST NOT treat a tool implementation as the owner of the generic namespace concept. | Generic namespace concepts belong to Product semantics, not an app-local implementation. |

## Requirements: spec:product.design_records.namespace_model.existing_artifacts

| id | requirement | notes |
|---|---|---|
| R008 | Implementations MUST use path-derived `spec:` identity for spec records instead of sequential workflow IDs. |  |
| R009 | Implementations MUST treat the app namespace segment as the owner identifier. | Applies to sequential workflow artifact IDs. |
| R010 | Implementations MUST treat the domain segment as the owner-local concern identifier. | Applies to sequential workflow artifact IDs. |
