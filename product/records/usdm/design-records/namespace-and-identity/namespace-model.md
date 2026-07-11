# USDM requirement: Namespace model

- **id**: `usdm:product.design_records.namespace_and_identity.namespace_model`
- **status**: draft
- **date**: 2026-07-09
- **kind**: requirement
- **parent**: `usdm:product.design_records.namespace_and_identity`

## What this is

Defines app/domain namespace boundaries for sequential workflow artifact IDs.

This record includes app/domain ID-space boundaries and subdomain as non-ID metadata.
This record does not include current Brewprint registry facts, V01 compatibility, repository placement, or authoring-tool advisory behavior.

## Requirements: Namespace model
> source: spec:product.design_records.namespace_model

| id | requirement | notes |
|---|---|---|
| R001 | New sequential workflow artifact IDs MUST include an app namespace segment. | Applies to REQ, WORK, INV, ADR, and TASK records. |
| R002 | New sequential workflow artifact IDs MUST include a domain namespace segment. | Applies to REQ, WORK, INV, ADR, and TASK records. |
| R003 | Implementations MUST treat each `<app namespace, domain namespace>` pair as a separate workflow artifact namespace. |  |
| R004 | Implementations MUST NOT treat a domain namespace as globally unique across all app namespaces. |  |
| R005 | The app namespace segment MUST identify the app, subsystem, or product-level scope that owns the workflow artifact. |  |
| R006 | The domain namespace segment MUST identify the owner-local ID bucket for related workflow artifacts. |  |

## Requirements: Subdomain model
> source: spec:product.design_records.namespace_model.subdomain_model

| id | requirement | notes |
|---|---|---|
| R007 | Subdomain metadata MUST NOT be used as a public ID segment. |  |
| R008 | Subdomain metadata MUST NOT affect sequence allocation. |  |
| R009 | Missing subdomain metadata MUST NOT make a workflow artifact invalid. |  |
