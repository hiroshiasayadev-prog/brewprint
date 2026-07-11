# USDM requirement: Semantic ref boundary

- **id**: `usdm:product.design_records.traceability_and_relations.semantic_ref_boundary`
- **status**: draft
- **date**: 2026-07-10
- **kind**: requirement
- **parent**: `usdm:product.design_records.traceability_and_relations`

## What this is

This record defines requirements for canonical Specification refs in Design Records traceability.

## Requirements: spec:product.design_records.spec_format.spec_id_as_ref

| id | requirement | notes |
|---|---|---|
| R001 | New and migrated Specification records must use the path-derived document-level `spec:` ref as the canonical spec identity. |  |
| R002 | The canonical spec ref must be written in the H1-adjacent `id` metadata marker. |  |
| R003 | For new and migrated Specification records, each canonical spec ref must map to exactly one spec path. |  |
| R004 | Moving or renaming a new or migrated Specification record must change the canonical spec ref unless a later compatibility contract explicitly preserves the old ref. |  |
