# Contract: Task identity and structure

- **id**: `spec:drmcp.design_records_mcp.artifacts.task.identity_and_structure`
- **status**: draft
- **date**: 2026-07-12
- **parent**: `spec:drmcp.design_records_mcp.artifacts.task`
- **contract_class**: `format`
- **usdm_covers**:
  - usdm:product.design_records.namespace_and_identity.workflow_artifact_identity#R002,#R005,#R007

## What this is

Defines the identity, record structure, and source placement for Task records.

## Record structure

- **structure**: `sequential`
- **artifact kind**: `TASK`

## Source placement

- **artifact directory**: `tasks`

## Identity form

```text
<APP_NAMESPACE>-TASK-<DOMAIN_NAMESPACE>-<WORK_SEQUENCE>-<TASK_SEQUENCE>
```

## Artifact-specific segments

| segment | role | format | sequence | allocation scope |
|---|---|---|---|---|
| `<WORK_SEQUENCE>` | Parent Work Item sequence. | Three-digit, zero-padded decimal. | `no` | `-` |
| `<TASK_SEQUENCE>` | Task sequence number. | Two-digit, zero-padded decimal. | `yes` | `parent work item` |

## Related specs

| ref | relation |
|---|---|
| `spec:drmcp.design_records_mcp.artifacts.base.definitions.record_structure` | Shared sequential record-structure rules. |
| `spec:drmcp.design_records_mcp.artifacts.base.definitions.identity_declaration` | Shared sequential identity form. |
| `spec:product.design_records.authoring_standards.task_authoring` | Product authority for Task identity. |
| `spec:product.design_records.namespace_model.artifact_id_grammar` | Product authority for Task ID grammar, segment formats, and Task sequence allocation scope. |
| `spec:product.design_records.repository_layout.record_discovery_paths` | Product authority for Task domain-scoped placement. |
