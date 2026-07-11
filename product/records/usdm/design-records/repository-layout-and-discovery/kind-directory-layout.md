# USDM requirement: Kind directory layout

- **id**: `usdm:product.design_records.repository_layout_and_discovery.kind_directory_layout`
- **status**: draft
- **date**: 2026-07-09
- **kind**: requirement
- **parent**: `usdm:product.design_records.repository_layout_and_discovery`

## What this is

Requirements for the kind-first directory layout directly under `records/` for Design Records.

## Requirements: spec:product.design_records.repository_layout

| id | requirement | notes |
|---|---|---|
| R001 | The implementation must treat the Design Records placement directly under `records/` as a kind-first directory layout. | The scope includes only active canonical kind directories. It includes at least `adr/`, `spec/`, `investigations/`, `requirements/`, `work-items/`, and `tasks/`. `guides/` is a stale source entry and is not part of this requirement. `usdm/` is also excluded because MVP USDM is auxiliary. |
