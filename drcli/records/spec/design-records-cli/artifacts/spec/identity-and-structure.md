# Contract: Specification identity and structure

- **id**: `spec:drcli.design_records_cli.artifacts.spec.identity_and_structure`
- **status**: draft
- **date**: 2026-10-07
- **parent**: `spec:drcli.design_records_cli.artifacts.spec`
- **contract_class**: `format`

## What this is

Defines the identity, record structure, and source placement for Specification records.

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

- **structure**: `tree`
- **record kind**: `spec`

## Source placement

- **artifact directory**: `spec`

## Related specs

| ref | relation |
|---|---|
| `spec:drcli.design_records_cli.artifacts.base.definitions.record_structure` | Shared tree record-structure rules. |
| `spec:drcli.design_records_cli.artifacts.base.definitions.identity_declaration` | Shared tree identity form and path mapping. |
| `spec:product.design_records.traceability.artifact_refs` | Product authority for the Specification record kind and canonical reference form. |
| `spec:product.design_records.spec_format.spec_id_as_ref` | Product authority for path-derived Specification identity. |
| `spec:product.design_records.repository_layout` | Product authority for Specification topic-tree placement under `records/spec/`. |
| `spec:product.design_records.repository_layout.record_discovery_paths` | Product authority for Specification source placement. |
