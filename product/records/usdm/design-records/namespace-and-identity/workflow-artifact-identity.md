# USDM requirement: Workflow artifact identity

- **id**: `usdm:product.design_records.namespace_and_identity.workflow_artifact_identity`
- **status**: draft
- **date**: 2026-07-10
- **kind**: requirement
- **parent**: `usdm:product.design_records.namespace_and_identity`

## What this is

Defines public ID requirements for REQ, WORK, INV, ADR, and TASK records.

This record includes ID grammar, sequence scope, and complete public IDs as canonical reference forms.
This record does not include spec IDs, legacy ID compatibility, physical paths, or field-specific allowed target sets.

## Requirements: Artifact ID grammar
> source: spec:product.design_records.namespace_model.artifact_id_grammar

| id | requirement | notes |
|---|---|---|
| R001 | Implementations MUST use `<APP_NAMESPACE>-<ARTIFACT_KIND>-<DOMAIN_NAMESPACE>-<SEQUENCE>` for new REQ, WORK, INV, and ADR IDs. |  |
| R002 | Implementations MUST use `<APP_NAMESPACE>-TASK-<DOMAIN_NAMESPACE>-<WORK_SEQUENCE>-<TASK_SEQUENCE>` for new TASK IDs. |  |
| R003 | Implementations MUST use 3-digit zero-padded sequence numbers for REQ, WORK, INV, and ADR IDs. |  |
| R004 | Implementations MUST use the parent Work Item sequence as the TASK `WORK_SEQUENCE` segment. |  |
| R005 | Implementations MUST use 2-digit zero-padded sequence numbers for TASK `TASK_SEQUENCE`. |  |
| R006 | Implementations MUST scope REQ, WORK, INV, and ADR sequence allocation by app namespace, artifact kind, and domain namespace. |  |
| R007 | Implementations MUST scope TASK sequence allocation by parent Work Item. |  |
| R008 | Implementations MUST use the complete public ID as the canonical reference form for REQ, WORK, INV, ADR, and TASK records. |  |
| R009 | Implementations MUST NOT treat bare forms such as `REQ-*`, `WORK-*`, or `TASK-*` as canonical external references. |  |
| R010 | Implementations MUST NOT include subdomain segments in artifact IDs. |  |

## Requirements: Canonical record kinds and references
> source: spec:product.design_records.traceability.artifact_refs

| id | requirement | notes |
|---|---|---|
| R011 | Implementations MUST treat decision, investigation, requirement, work item, and task record refs as complete public IDs. |  |
| R012 | Implementations MUST treat new and migrated specs as path-derived `spec:` refs, not `SPEC-*` public IDs. |  |
| R013 | Resolvers MUST support Task public IDs as direct resolver inputs. | Field-specific metadata contracts decide whether Task public IDs are allowed in persisted relation fields. |
| R014 | Investigation canonical-reference fields MUST NOT accept Task public IDs. |  |
