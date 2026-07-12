# Contract: Decision identity and structure

- **id**: `spec:drmcp.design_records_mcp.artifacts.decision.identity_and_structure`
- **status**: draft
- **date**: 2026-07-12
- **parent**: `spec:drmcp.design_records_mcp.artifacts.decision`
- **contract_class**: `format`
- **usdm_covers**:
  - usdm:product.design_records.namespace_and_identity.workflow_artifact_identity#R001,#R003,#R006,#R010

## What this is

Defines the identity, record structure, and source placement for Decision records.

## Record structure

- **structure**: `sequential`
- **artifact kind**: `ADR`

## Source placement

- **artifact directory**: `adr`

## Identity form

```text
<APP_NAMESPACE>-ADR-<DOMAIN_NAMESPACE>-<SEQUENCE>
```

## Artifact-specific segments

| segment | role | format | sequence | allocation scope |
|---|---|---|---|---|
| `<SEQUENCE>` | ADR sequence number. | Three-digit, zero-padded decimal. | `yes` | `domain` |

## Related specs

| ref | relation |
|---|---|
| `spec:drmcp.design_records_mcp.artifacts.base.definitions.record_structure` | Shared sequential record-structure rules. |
| `spec:drmcp.design_records_mcp.artifacts.base.definitions.identity_declaration` | Shared sequential identity form. |
| `spec:product.design_records.authoring_standards.adr_authoring` | Product authority for ADR identity and standard placement. |
| `spec:product.design_records.namespace_model.artifact_id_grammar` | Product authority for ADR ID grammar, sequence format, allocation scope, and subdomain exclusion. |
| `spec:product.design_records.traceability.artifact_refs` | Product authority for the Decision record kind and canonical reference form. |
| `spec:product.design_records.repository_layout.record_discovery_paths` | Product authority for ADR domain-scoped placement. |
