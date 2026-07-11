# USDM requirement: Metadata relation schema

- **id**: `usdm:product.design_records.traceability_and_relations.metadata_relation_schema`
- **status**: draft
- **date**: 2026-07-10
- **kind**: requirement
- **parent**: `usdm:product.design_records.traceability_and_relations`

## What this is

This record defines requirements for metadata and table fields that supply Design Records traceability relations.

## Requirements: Trace metadata schema
> source: spec:product.design_records.traceability.metadata_schema

| id | requirement | notes |
|---|---|---|
| R001 | Traceability metadata handling must treat the H1-adjacent `parent` marker as the canonical parent ref for topic placement. |  |
| R002 | Traceability metadata handling must read the authoritative child relationship for Index and Overview specs from the visible `## Topics` table. |  |
| R003 | Investigation `source_refs` entries must use supported canonical refs and must resolve. |  |
| R004 | Investigation `follow_up_results` entries must use supported canonical refs and must resolve. |  |
| R005 | When an artifact ref is written in Investigation `follow_up_candidates`, that value must use canonical form. | Natural-language candidate descriptions are not canonical refs. |
| R006 | A not-yet-created `follow_up_candidates` target must not be invalid only because the target does not resolve yet. |  |
