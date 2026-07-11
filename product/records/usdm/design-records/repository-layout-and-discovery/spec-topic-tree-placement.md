# USDM requirement: Spec topic tree placement

- **id**: `usdm:product.design_records.repository_layout_and_discovery.spec_topic_tree_placement`
- **status**: draft
- **date**: 2026-07-09
- **kind**: requirement
- **parent**: `usdm:product.design_records.repository_layout_and_discovery`

## What this is

Requirements for canonical topic tree placement of Specification records.

## Requirements: spec:product.design_records.repository_layout

| id | requirement | notes |
|---|---|---|
| R001 | The implementation must treat Specification records as placed in the canonical topic tree under `records/spec/`. | Specifications use topic tree placement, not sequential artifact domain placement. |
| R002 | The implementation must treat the directory and file structure under `records/spec/` as the physical representation of the Specification topic tree. | Detailed path-derived `spec:` ref derivation rules are covered by `physical-path-boundary`. |
