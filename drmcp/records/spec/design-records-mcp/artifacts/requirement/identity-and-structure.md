# Contract: Requirement identity and structure

- **id**: `spec:drmcp.design_records_mcp.artifacts.requirement.identity_and_structure`
- **status**: draft
- **date**: 2026-07-13
- **parent**: `spec:drmcp.design_records_mcp.artifacts.requirement`
- **contract_class**: `format`
- **usdm_covers**:
  - usdm:product.design_records.namespace_and_identity.workflow_artifact_identity#R001,#R003,#R006,#R008,#R010,#R015-R017

## What this is

Defines the identity, record structure, and source placement for Requirement records.

## Record structure

- **structure**: `sequential`
- **artifact kind**: `REQ`

## Source placement

- **artifact directory**: `requirements`

## Identity form

```text
<APP_NAMESPACE>-REQ-<DOMAIN_NAMESPACE>-<SEQUENCE>
```

## Identity authority

The complete public ID in the H1 prefix is the canonical identity authority for a Requirement source.
It must agree with the configured app namespace, the `REQ` kind selected by this source corpus, and the physical domain directory.
H1-adjacent metadata and the file name do not supply or repair identity.
The file-name public-ID prefix is a nonidentity conformance projection of the H1 public ID.

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
