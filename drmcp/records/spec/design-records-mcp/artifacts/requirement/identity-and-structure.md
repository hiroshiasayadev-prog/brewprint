# Contract: Requirement identity and structure

- **id**: `spec:drmcp.design_records_mcp.artifacts.requirement.identity_and_structure`
- **status**: draft
- **date**: 2026-07-11
- **parent**: `spec:drmcp.design_records_mcp.artifacts.requirement`
- **contract_class**: `format`
- **usdm_covers**:
  - usdm:product.design_records.namespace_and_identity.workflow_artifact_identity#R001,#R003,#R006,#R010

## What this is

Defines the identity and record structure for Requirement records.

## Record structure

- **structure**: `sequential`
- **artifact kind**: `REQ`

## Identity form

```text
<APP_NAMESPACE>-REQ-<DOMAIN_NAMESPACE>-<SEQUENCE>
```

## Artifact-specific segments

| segment | role | format | sequence | allocation scope |
|---|---|---|---|---|
| `<SEQUENCE>` | Requirement sequence number. | Three-digit, zero-padded decimal. | `yes` | `domain` |

## Related specs

| ref | relation |
|---|---|
| `spec:drmcp.design_records_mcp.artifacts.base.definitions.record_structure` | Shared sequential record-structure rules. |
| `spec:drmcp.design_records_mcp.artifacts.base.definitions.identity_declaration` | Shared sequential identity form. |
| `spec:product.design_records.namespace_model.artifact_id_grammar` | Product authority for Requirement ID grammar, sequence format, and allocation scope. |
| `spec:product.design_records.traceability.artifact_refs` | Product authority for the Requirement record kind and canonical reference form. |
| `spec:product.design_records.repository_layout.record_discovery_paths` | Product authority for Requirement domain-scoped placement. |
