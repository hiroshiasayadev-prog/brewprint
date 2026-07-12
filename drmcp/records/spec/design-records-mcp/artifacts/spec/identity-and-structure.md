# Contract: Specification identity and structure

- **id**: `spec:drmcp.design_records_mcp.artifacts.spec.identity_and_structure`
- **status**: draft
- **date**: 2026-07-12
- **parent**: `spec:drmcp.design_records_mcp.artifacts.spec`
- **contract_class**: `format`

## What this is

Defines the identity, record structure, and source placement for Specification records.

## Record structure

- **structure**: `tree`
- **record kind**: `spec`

## Source placement

- **artifact directory**: `spec`

## Related specs

| ref | relation |
|---|---|
| `spec:drmcp.design_records_mcp.artifacts.base.definitions.record_structure` | Shared tree record-structure rules. |
| `spec:drmcp.design_records_mcp.artifacts.base.definitions.identity_declaration` | Shared tree identity form and path mapping. |
| `spec:product.design_records.traceability.artifact_refs` | Product authority for the Specification record kind and canonical reference form. |
| `spec:product.design_records.spec_format.spec_id_as_ref` | Product authority for path-derived Specification identity. |
| `spec:product.design_records.repository_layout` | Product authority for Specification topic-tree placement under `records/spec/`. |
| `spec:product.design_records.repository_layout.record_discovery_paths` | Product authority for Specification source placement. |
