# Reference: Resolver responsibility

- **id**: `spec:drmcp.design_records_mcp.resolver`
- **status**: draft
- **date**: 2026-06-22
- **parent**: `spec:drmcp.design_records_mcp.overview`

## What this is

Defines DRMCP-specific resolver behavior: which tool exposes the resolver, which inputs return `unsupported`, and the lookup-source vs. record-kind boundary. The canonical reference model (supported input forms, MVP exclusions, validation scope) is owned by `spec:product.design_records.traceability.resolve_and_validation`.

## Current contract

The canonical reference model — supported resolver inputs, workflow artifact ref handling, MVP exclusions, and validation scope — is defined in `spec:product.design_records.traceability.resolve_and_validation`.

### Public tool and unsupported inputs

The resolver's public tool name is `resolve_reference` (see `spec:drmcp.design_records_mcp.tools.resolve_reference`).

| input form | resolver response |
|---|---|
| `internal-design:` / `coverage:` / `COV-*` | `unsupported` |
| Physical file paths | `unsupported` |
| Unrecognized ID forms | `unsupported` |

The `yaml:` reserved prefix public resolver input and direct query response behavior are not defined in MVP.

### Lookup source vs. record kind boundary

The artifacts that the resolver reads as a lookup source and the artifacts that `list_records` / `get_record` expose as record kinds are not required to be the same set.

> Source: V01-ADR-087 §4, V01-ADR-088, V01-ADR-092 §3–§7
