# USDM requirement: Records root

- **id**: `usdm:product.design_records.repository_layout_and_discovery.records_root`
- **status**: draft
- **date**: 2026-07-09
- **kind**: requirement
- **parent**: `usdm:product.design_records.repository_layout_and_discovery`

## What this is

Requirements for the app-local Design Records placement root and record discovery base path.

## Requirements: spec:product.design_records.repository_layout

| id | requirement | notes |
|---|---|---|
| R001 | The implementation must treat `records/` under an app namespace directory as the Design Records placement root for an active app namespace that participates in design governance. | This requirement does not define the full structure of the app namespace directory. |
| R002 | The implementation must not require `src/`, `dsl/`, or any other non-Design Records directory under an app namespace directory before treating the namespace as a Design Records namespace. | This requirement does not prohibit `src/` or `dsl/`. It means that `records/` alone can be sufficient for a Design Records namespace. |

## Requirements: spec:product.design_records.repository_layout.record_discovery_paths

| id | requirement | notes |
|---|---|---|
| R003 | The implementation must treat `records/` under a valid app namespace directory as the base path for record discovery. | `records_root` corresponds to `records/` under an app namespace directory. This row does not define how an implementation determines that an app namespace directory is valid. |
