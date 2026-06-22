# Reference: v1 record ID grammar

- **id**: `spec:product.concepts.namespace_model.v1_id_grammar`
- **status**: draft
- **date**: 2026-06-22
- **parent**: `spec:product.concepts.namespace_model`

## What this is

Defines the two-layer v1 record ID model: public ID (externally visible, namespace-prefixed) and bare ID (internal validation form after prefix stripping). Relocated from `spec:drmcp.design_records_mcp.schema.id_normalization` (Phase 2 relocation per PRODUCT-WORK-SPEC-004).

## Current contract

### Public ID

The ID exposed externally by all tools. Formed by concatenating `namespace_prefix` + `bare_id`.

```
public_id = namespace_prefix + bare_id
```

Public ID examples (single-root mode, `--records-root v01/records`, `namespace_prefix = V01-`):

| kind | public ID example |
|---|---|
| `decision` | `V01-ADR-076` |
| `spec` | `V01-SPEC-design-records-mcp-schema` |
| `investigation` | `V01-INV-MCP-001` |
| `requirement` | `V01-REQ-MCP-003` |
| `work_item` | `V01-WORK-MCP-003` |
| `task` | `V01-TASK-MCP-003-01` |

All tool inputs and outputs use the public ID. Bare IDs are not accepted as tool inputs.

### Bare ID grammar

The form used for internal validation after the parser strips `namespace_prefix`. Not part of the external tool contract.

| kind | bare ID grammar |
|---|---|
| `decision` | `ADR-NNN` (3-digit zero-padded) |
| `spec` | `SPEC-<slug>` |
| `investigation` | `INV-<DOMAIN>-NNN` |
| `requirement` | `REQ-<DOMAIN>-NNN` |
| `work_item` | `WORK-<DOMAIN>-NNN` |
| `task` | `TASK-<DOMAIN>-<WORK-SEQUENCE>-<TASK-SEQUENCE>` |

When bare ID grammar forms appear elsewhere in specs without explicit qualification, they refer to the internal validation rule (not the public tool contract form).

> Source: V01-ADR-097, V01-ADR-099
