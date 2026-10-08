# Contract: Requirement source

- **id**: `spec:drcli.design_records_cli.artifacts.requirement.source`
- **status**: draft
- **date**: 2026-10-07
- **parent**: `spec:drcli.design_records_cli.artifacts.requirement`
- **contract_class**: `format`
- **usdm_covers**:
  - usdm:product.design_records.authoring_semantics.workflow_record_authoring#R001-R002,#R005-R006
  - usdm:product.design_records.namespace_and_identity.workflow_artifact_identity#R008

## What this is

Defines the source-document shape for Requirement records.

## Current contract

This Specification defines the DRCLI-local projection and declarations for this topic.
Where this Specification cites PRODUCT authority, the cited PRODUCT Specification remains authoritative for PRODUCT-owned semantics.

## Rules

- DRCLI-local declarations must preserve cited PRODUCT-owned semantics.
- A DRCLI-local declaration must not create an independent replacement for cited PRODUCT authority.

## Validation rules

- A DRCLI interpretation that conflicts with cited PRODUCT authority is nonconforming.
- Validation of PRODUCT-owned meaning uses the cited PRODUCT authority.

## H1 prefix

- **type**: `identity`

### Value

```text
<APP_NAMESPACE>-REQ-<DOMAIN_NAMESPACE>-<SEQUENCE>
```

## Identity authority

The complete H1 prefix is the sole source-internal identity authority.
Requirement metadata and the file name do not supply, repair, complete, normalize, infer, or replace the identity.
The file-name public-ID prefix is validated separately as a conformance projection.

## H2 heading policy

- **unlisted headings**: `allowed`

## H2 headings

| heading | condition | format reference |
|---|---|---|
| `## Requirement` | `always` | `-` |
| `## Evidence` | `always` | `-` |
| `## Required Outcome` | `status = accepted: always;`<br/>`status != accepted: optional;` | `-` |

## Related specs

| ref | relation |
|---|---|
| `spec:drcli.design_records_cli.artifacts.base.definitions.source` | Shared source-document rules. |
| `spec:drcli.design_records_cli.artifacts.requirement.identity_and_structure` | Defines the public Requirement ID used as the H1 prefix. |
| `spec:drcli.design_records_cli.artifacts.requirement.h1_adjacent_metadata` | Defines the metadata block placed after H1. |
| `spec:product.design_records.authoring_standards.requirement_authoring` | Product authority for Requirement source headings and presence conditions. |
