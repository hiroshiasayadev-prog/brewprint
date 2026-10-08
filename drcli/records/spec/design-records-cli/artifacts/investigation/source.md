# Contract: Investigation source

- **id**: `spec:drcli.design_records_cli.artifacts.investigation.source`
- **status**: draft
- **date**: 2026-10-07
- **parent**: `spec:drcli.design_records_cli.artifacts.investigation`
- **contract_class**: `format`
- **usdm_covers**:
  - usdm:product.design_records.authoring_semantics.workflow_record_authoring#R005-R006
  - usdm:product.design_records.namespace_and_identity.workflow_artifact_identity#R008

## What this is

Defines the source-document shape for Investigation records.

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
<APP_NAMESPACE>-INV-<DOMAIN_NAMESPACE>-<SEQUENCE>
```

## Identity authority

The complete H1 prefix is the sole source-internal identity authority.
Investigation metadata and the file name do not supply, repair, complete, normalize, infer, or replace the identity.
The file-name public-ID prefix is validated separately as a conformance projection.

## H2 heading policy

- **unlisted headings**: `prohibited`

## H2 headings

| heading | condition | format reference |
|---|---|---|
| `## Investigation scope` | `always` | `-` |
| `## Out of scope` | `always` | `-` |
| `## Background` | `always` | `-` |
| `## What was investigated` | `always` | `-` |
| `## Findings` | `always` | `-` |
| `## Cross-cutting observations` | `always` | `-` |
| `## Follow-up judgment candidates` | `always` | `-` |
| `## Recommendation` | `always` | `-` |
| `## Follow-up artifact candidates` | `always` | `-` |
| `## Open questions` | `always` | `-` |

## Related specs

| ref | relation |
|---|---|
| `spec:drcli.design_records_cli.artifacts.base.definitions.source` | Shared H1 shape, metadata placement, and H2 declaration rules. |
| `spec:drcli.design_records_cli.artifacts.investigation.identity_and_structure` | Defines the Investigation identity represented by the H1 prefix. |
| `spec:drcli.design_records_cli.artifacts.investigation.h1_adjacent_metadata` | Defines the Investigation metadata fields placed after H1. |
| `spec:product.design_records.authoring_standards.investigation_authoring` | Product authority for Investigation source headings and presence conditions. |
