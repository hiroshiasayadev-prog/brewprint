# USDM requirement: Workflow relations

- **id**: `usdm:product.design_records.traceability_and_relations.workflow_relations`
- **status**: draft
- **date**: 2026-07-10
- **kind**: requirement
- **parent**: `usdm:product.design_records.traceability_and_relations`

## What this is

This record defines requirements for reading and validating workflow relation metadata.

## Requirements: Trace metadata schema
> source: spec:product.design_records.traceability.metadata_schema

| id | requirement | notes |
|---|---|---|
| R001 | The implementation must read Work Item `source_refs` as the Work Item source relation. |  |
| R002 | Work Item `source_refs` must be present and non-empty. |  |
| R003 | The order of Work Item `source_refs` entries must have no semantic meaning. |  |
| R004 | Duplicate canonical refs inside one Work Item `source_refs` field must be treated as invalid. |  |
| R005 | A Work Item must not include its own canonical identity in `source_refs`. |  |
| R006 | The implementation must read Work Item `tasks` as Task ownership relation values. |  |
| R007 | The implementation must read Task `work_item` as the owning Work Item relation value. |  |
| R008 | The implementation must read Task `depends_on` as Task dependency relation values. |  |
| R009 | Task records must not contain a source relation field. |  |
| R010 | Workflow relations must be determined only from defined metadata fields. | This excludes relation inference from paths, file names, parent directories, and ID string structure. |
| R011 | Work Item `source_refs` must allow Product-defined canonical reference forms for spec, decision, investigation, requirement, work item, and task records. | Compatibility inputs remain separate from the current field contract unless compatibility migration rules explicitly allow them. |
