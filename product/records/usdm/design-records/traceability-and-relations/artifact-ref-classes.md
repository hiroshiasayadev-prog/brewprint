# USDM requirement: Canonical record kinds and reference forms

- **id**: `usdm:product.design_records.traceability_and_relations.artifact_ref_classes`
- **status**: draft
- **date**: 2026-07-10
- **kind**: requirement
- **parent**: `usdm:product.design_records.traceability_and_relations`

## What this is

This record defines requirements for Product-defined current record kinds, canonical reference forms, and field-use boundaries in Design Records traceability.

This record does not define field-specific allowed target sets or Brewprint legacy compatibility mappings.

## Requirements: Canonical record kinds and references
> source: spec:product.design_records.traceability.artifact_refs

| id | requirement | notes |
|---|---|---|
| R001 | Current Design Record reference values must use a Product-defined canonical reference form. | Field contracts may further restrict which canonical reference forms are allowed in each field. |
| R002 | Product-defined current canonical reference forms must be limited to path-derived `spec:` refs and complete public IDs for decision, investigation, requirement, work item, and task records. | Legacy issued IDs are compatibility inputs, not current canonical reference forms. |
| R003 | New and migrated Specification records must not receive new `SPEC-*` public IDs as canonical spec refs. |  |
| R004 | Legacy issued IDs must be treated as compatibility inputs only when Brewprint compatibility records preserve them. |  |
| R005 | Persisted metadata fields must permit a canonical reference form before that form may appear in the field. | This separates the global reference-form set from field-specific allowed target sets. |
| R006 | A traceability relation or reference value must be treated as invalid when the value uses an unrecognized reference form or a canonical reference form not permitted by the field contract. |  |

## Requirements: Artifact ID grammar
> source: spec:product.design_records.namespace_model.artifact_id_grammar

| id | requirement | notes |
|---|---|---|
| R007 | Public ID refs must use the complete public ID of the target record. | Applies to decision, investigation, requirement, work item, and task records. |
| R008 | Bare grammar fragments such as `REQ-*`, `WORK-*`, and `TASK-*` must not be treated as canonical external refs. |  |
