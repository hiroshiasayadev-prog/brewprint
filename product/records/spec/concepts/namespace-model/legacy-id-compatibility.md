# Reference: Legacy ID compatibility

- **id**: `spec:product.concepts.namespace_model.legacy_id_compatibility`
- **status**: draft
- **date**: 2026-06-24
- **parent**: `spec:product.concepts.namespace_model`

## What this is

Defines the compatibility policy for V01-epoch artifact IDs: which legacy public-ID families are accepted, that their issued IDs remain unchanged and resolvable, and that they are not canonical for new records.

## Accepted legacy families

The following public-ID families were issued before the artifact ID grammar was adopted. All are accepted as valid resolvable references.

| family | example |
|---|---|
| `V01-ADR-*` | `V01-ADR-096` |
| `V01-SPEC-*` | `V01-SPEC-design-records-mcp-schema` |
| `V01-INV-*` | `V01-INV-MCP-001` |
| `V01-REQ-*` | `V01-REQ-PRODUCT-001` |
| `V01-WORK-*` | `V01-WORK-MCP-004` |
| `V01-TASK-*` | `V01-TASK-MCP-003-01` |

## Retention policy

Issued legacy IDs remain unchanged and resolvable. No V01-* artifact is renamed as part of this policy.

Legacy public IDs are not the canonical form for new sequential records. New sequential records (ADR, investigation, requirement, work item, and task) must use the grammar defined in `spec:product.concepts.namespace_model.artifact_id_grammar`. New specs use path-derived `spec:` refs; see `spec:product.concepts.spec_format.spec_id_as_ref`.

## Spec identity note

The `V01-SPEC-*` family provides compatibility-only indexed IDs for specs from the V01 epoch. These IDs do not constitute canonical spec identity. Canonical spec identity is path-derived; see `spec:product.concepts.spec_format.spec_id_as_ref`.

## Related specs

| ref | relation |
|---|---|
| `spec:product.concepts.namespace_model` | Parent overview. |
| `spec:product.concepts.namespace_model.artifact_id_grammar` | Canonical grammar for new records. |
| `spec:product.concepts.namespace_model.existing_artifacts` | Attribution policy for existing artifacts under V01-ADR-096. |
| `spec:product.concepts.spec_format.spec_id_as_ref` | Canonical spec identity rules. |
