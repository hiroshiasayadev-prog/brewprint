# Reference: Legacy ID compatibility

- **id**: `spec:product.brewprint.compatibility.legacy_id_compatibility`
- **status**: draft
- **date**: 2026-06-24
- **parent**: `spec:product.brewprint.compatibility`

## What this is

Defines the Brewprint compatibility policy for V01-epoch artifact IDs.

## Accepted legacy families

The following public-ID families were issued before the app-aware artifact ID grammar was adopted.
All are accepted as valid resolvable references.

| family | example |
|---|---|
| `V01-ADR-*` | `V01-ADR-096` |
| `V01-SPEC-*` | `V01-SPEC-design-records-mcp-schema` |
| `V01-INV-*` | `V01-INV-MCP-001` |
| `V01-REQ-*` | `V01-REQ-PRODUCT-001` |
| `V01-WORK-*` | `V01-WORK-MCP-004` |
| `V01-TASK-*` | `V01-TASK-MCP-003-01` |

## Retention policy

Issued legacy IDs remain unchanged and resolvable.
No V01 artifact is renamed as part of this policy.

Legacy public IDs are not the canonical form for new sequential records.
New sequential records use the grammar defined in `spec:product.design_records.namespace_model.artifact_id_grammar`.

New specs use path-derived `spec:` refs.
Spec identity is defined in `spec:product.design_records.spec_format.spec_id_as_ref`.

## Spec identity note

The `V01-SPEC-*` family provides compatibility-only indexed IDs for specs from the V01 epoch.
These IDs do not constitute canonical spec identity.

Canonical spec identity is path-derived.

## Related specs

| ref | relation |
|---|---|
| `spec:product.brewprint.compatibility` | Parent compatibility overview. |
| `spec:product.brewprint.compatibility.existing_artifacts` | Historical ownership and effective attribution. |
| `spec:product.design_records.namespace_model.artifact_id_grammar` | Canonical grammar for new workflow artifact IDs. |
| `spec:product.design_records.spec_format.spec_id_as_ref` | Canonical spec identity rules. |
