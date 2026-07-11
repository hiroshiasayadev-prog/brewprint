# USDM requirement: Domain-scoped placement

- **id**: `usdm:product.design_records.repository_layout_and_discovery.domain_scoped_placement`
- **status**: draft
- **date**: 2026-07-09
- **kind**: requirement
- **parent**: `usdm:product.design_records.repository_layout_and_discovery`

## What this is

Requirements for domain subdirectory placement of sequential Design Records artifacts.

## Requirements: Repository layout model
> source: spec:product.design_records.repository_layout

| id | requirement | notes |
|---|---|---|
| R001 | The implementation must treat sequential Design Records artifacts as placed in a domain subdirectory under their kind directory. | The scope includes `adr/<domain>/`, `investigations/<domain>/`, `requirements/<domain>/`, `work-items/<domain>/`, and `tasks/<domain>/`. |
| R002 | The implementation must not treat `records/spec/` as subject to domain subdirectory placement. | Specification records use canonical topic tree placement, not domain placement. |
