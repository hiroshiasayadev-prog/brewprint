# Contract: Task identity and structure

- **id**: `spec:drcli.design_records_cli.artifacts.task.identity_and_structure`
- **status**: draft
- **date**: 2026-10-08
- **parent**: `spec:drcli.design_records_cli.artifacts.task`
- **contract_class**: `format`
- **usdm_covers**:
  - usdm:product.design_records.namespace_and_identity.workflow_artifact_identity#R002,#R005,#R007-R008,#R015-R017

## What this is

Defines the identity, record structure, and source placement for Task records.

## Current contract

This Specification defines the DRCLI-local projection and declarations for this topic.
Where this Specification cites PRODUCT authority, the cited PRODUCT Specification remains authoritative for PRODUCT-owned semantics.

## Rules

- DRCLI-local declarations must preserve cited PRODUCT-owned semantics.
- A DRCLI-local declaration must not create an independent replacement for cited PRODUCT authority.

## Validation rules

- A DRCLI interpretation that conflicts with cited PRODUCT authority is nonconforming.
- Validation of PRODUCT-owned meaning uses the cited PRODUCT authority.

## Record structure

- **structure**: `sequential`
- **artifact kind**: `TASK`

## Source placement

- **artifact directory**: `tasks`

## Identity form

```text
<APP_NAMESPACE>-TASK-<DOMAIN_NAMESPACE>-<WORK_SEQUENCE>-<TASK_SEQUENCE>
```

## Identity authority

The complete public ID in the H1 prefix is the canonical identity authority for a Task source.
It must agree with the resolved app namespace, the `TASK` kind selected by this source corpus, and the physical domain directory.
H1-adjacent metadata and the file name do not supply or repair identity.
The file-name public-ID prefix is a nonidentity conformance projection of the H1 public ID.

## Artifact-specific segments

| segment | role | format | sequence | allocation scope |
|---|---|---|---|---|
| `<WORK_SEQUENCE>` | Parent Work Item sequence. | Three-digit, zero-padded decimal. | `yes` | `-` |
| `<TASK_SEQUENCE>` | Task sequence number. | Two-digit, zero-padded decimal. | `yes` | `parent work item` |

## Related specs

| ref | relation |
|---|---|
| `spec:drcli.design_records_cli.artifacts.base.definitions.record_structure` | Shared sequential record-structure rules. |
| `spec:drcli.design_records_cli.artifacts.base.definitions.identity_declaration` | Shared sequential identity form. |
| `spec:product.design_records.authoring_standards.task_authoring` | Product authority for Task identity. |
| `spec:product.design_records.namespace_model.artifact_id_grammar` | Product authority for Task ID grammar, segment formats, and Task sequence allocation scope. |
| `spec:product.design_records.repository_layout.record_discovery_paths` | Product authority for Task domain-scoped placement. |
