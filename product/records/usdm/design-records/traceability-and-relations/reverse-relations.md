# USDM requirement: Reverse relations

- **id**: `usdm:product.design_records.traceability_and_relations.reverse_relations`
- **status**: draft
- **date**: 2026-07-10
- **kind**: requirement
- **parent**: `usdm:product.design_records.traceability_and_relations`

## What this is

This record defines requirements for deriving Requirement-to-Work Item relations from Work Item metadata.

## Requirements: spec:product.design_records.traceability.metadata_schema

| id | requirement | notes |
|---|---|---|
| R001 | The authoritative source for Requirement-to-Work Item relations must be the Work Item `source_refs` field. |  |
| R002 | A Work Item reference field on a Requirement record must not establish a Requirement-to-Work Item relation by itself. |  |
| R003 | Requirement records must not be required to contain a Work Item reference field. |  |
| R004 | When Work Item A lists Work Item B in `source_refs`, Work Item A must not treat Work Item B's Requirement refs as its own Requirement refs. |  |
