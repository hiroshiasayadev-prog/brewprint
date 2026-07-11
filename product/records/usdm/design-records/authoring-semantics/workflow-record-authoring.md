# USDM requirement: Workflow record authoring shape

- **id**: `usdm:product.design_records.authoring_semantics.workflow_record_authoring`
- **status**: draft
- **date**: 2026-07-10
- **kind**: requirement
- **parent**: `usdm:product.design_records.authoring_semantics`

## What this is

This record defines DRMCP static-validation requirements for ADR, Requirement, Work Item, Task, and Investigation record shape.

## Requirements: Authoring standards
> source: spec:product.design_records.authoring_standards

| id | requirement | notes |
|---|---|---|
| R001 | ADR, Requirement, Work Item, Task, and Investigation records must contain exactly one real ATX H1 outside fenced code blocks. | Condenses the common file-shape rule from the artifact-specific authoring specs. |
| R002 | ADR, Requirement, Work Item, Task, and Investigation records must place the bullet metadata block immediately after H1 and before the first H2 body section. | Condenses the common metadata-placement rule. |
| R003 | ADR, Requirement, Work Item, Task, and Investigation metadata must contain only fields allowed by the persisted metadata schema for that artifact kind. | Implement as an artifact-kind-specific metadata allow-list. |
| R004 | Requirement, Work Item, and Task metadata must contain a persisted `id` field. | Current source specs intentionally differ from ADR and Investigation. |
| R005 | ADR and Investigation metadata must not contain a persisted `id` field. | Current source specs intentionally differ from Requirement, Work Item, and Task. |
| R006 | ADR, Requirement, Work Item, Task, and Investigation records must contain the required canonical H2 sections defined for their artifact kind. | Section presence is structural; substantive content rules are separate. |

## Requirements: Task authoring
> source: spec:product.design_records.authoring_standards.task_authoring

| id | requirement | notes |
|---|---|---|
| R007 | A non-implementation Task record must not contain a `## Implementation contract` section. | Determined from `task_type` and H2 headings. |
| R008 | An implementation Task record must contain a `## Implementation contract` section with the expected Markdown table shape. | Static validation is limited to section and table shape. |
| R009 | Task metadata must not contain `source_requirement` or `source_refs`. | Task provenance is reached through `work_item`. |

## Requirements: Work-item authoring
> source: spec:product.design_records.authoring_standards.work_item_authoring

| id | requirement | notes |
|---|---|---|
| R010 | Work Item metadata must use `source_refs` as the only persisted source-provenance field. | Enforced through the Work Item metadata allow-list. |
