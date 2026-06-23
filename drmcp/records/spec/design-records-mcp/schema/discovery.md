# Reference: Discovery

- **id**: `spec:drmcp.design_records_mcp.schema.discovery`
- **status**: draft
- **date**: 2026-06-22
- **parent**: `spec:drmcp.design_records_mcp.schema.overview`

## What this is

Defines the DRMCP-specific index inclusion conditions per record kind. Kind-level path-pattern conventions have been relocated to `spec:product.concepts.repository_layout.record_discovery_paths` (Phase 2 relocation per PRODUCT-WORK-SPEC-004).

## Current contract

Path-pattern conventions by kind are defined in `spec:product.concepts.repository_layout.record_discovery_paths`. New `decision` records use a domain subdirectory like other domain-scoped sequential artifacts: `<records_root>/adr/<domain>/<namespace_prefix>ADR-*-*.md`. Existing flat ADR records remain supported through the compatibility pattern `<records_root>/adr/<namespace_prefix>ADR-*.md`.

DRMCP-specific index inclusion conditions:

| kind | inclusion condition |
|---|---|
| `spec` | File must carry both `design_record.id` and `design_record.kind` in YAML front matter. Files without these fields are excluded without a `missing_design_record` diagnostic. |
| `spec` | Files whose `design_record.kind` is not `spec` are excluded without a diagnostic. |
| all other kinds | No additional inclusion filter beyond the path pattern. |

Additional notes:

- This spec does not restrict adding other record kinds in future.

> Source: V01-ADR-076 §bootstrap方針, V01-ADR-077 §validate_records の責務, V01-ADR-092 §1
