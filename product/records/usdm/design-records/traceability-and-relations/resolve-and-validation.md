# USDM requirement: Resolve and validation

- **id**: `usdm:product.design_records.traceability_and_relations.resolve_and_validation`
- **status**: draft
- **date**: 2026-07-10
- **kind**: requirement
- **parent**: `usdm:product.design_records.traceability_and_relations`

## What this is

This record defines requirements for traceability lookup sources and invalid traceability states.

## Requirements: spec:product.design_records.traceability.resolve_and_validation

| id | requirement | notes |
|---|---|---|
| R001 | A Specification canonical ref must be treated as a valid spec identity only when the H1-adjacent `id` matches the path-derived ref. |  |
| R002 | A new or migrated Specification record must be treated as invalid when its H1-adjacent `id` differs from the path-derived ref. |  |
| R003 | ADR, investigation, requirement, work item, and task identities must be registered from each record's complete public ID. |  |
| R004 | Legacy issued IDs must be registered only as compatibility inputs from mappings preserved by Brewprint compatibility records. |  |
| R005 | A Work Item must be treated as invalid when `source_refs` is missing or empty. |  |
| R006 | A Work Item `source_refs` entry must be treated as invalid when the entry is unresolved, unrecognized, noncanonical, or not permitted by the `source_refs` field contract. |  |
| R007 | Work Item and Task ownership must be treated as invalid when Work Item `tasks` and Task `work_item` do not agree. |  |
| R008 | A Task dependency must be treated as invalid when Task `depends_on` points to a missing Task ID. |  |
| R009 | A cycle formed through Work Item `source_refs` must be treated as invalid. |  |
| R010 | When a Work Item `source_refs` entry points to a Task ID, cycle validation must treat the entry as a reference to that Task's owning Work Item. |  |
| R011 | Work Item `tasks`, Task `work_item`, and Task `depends_on` must not create edges for Work Item `source_refs` cycle validation. |  |
| R012 | Multiple new or migrated Specification paths with the same canonical `spec:` ref must be treated as duplicate spec identity. |  |
| R013 | Multiple records with the same complete public ID must be treated as duplicate record identity. |  |
| R014 | A legacy issued ID that maps to multiple targets through compatibility records must be treated as an ambiguous compatibility ID. |  |
| R015 | ADR, investigation, requirement, work item, and task records must be treated as invalid when the resolved public ID does not match the grammar for the record kind. | Validates persisted record identity after read. |
