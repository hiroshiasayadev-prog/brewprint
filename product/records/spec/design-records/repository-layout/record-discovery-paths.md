# Reference: Record discovery paths

- **id**: `spec:product.design_records.repository_layout.record_discovery_paths`
- **status**: draft
- **date**: 2026-06-24
- **parent**: `spec:product.design_records.repository_layout`

## What this is

Defines app-independent path-pattern conventions for locating record files within a `records_root` by record kind.
It does not define DRMCP discovery filters or namespace-prefix derivation behavior.

## Current contract

| kind | path pattern |
|---|---|
| `decision` | `<records_root>/adr/*/<record_prefix>ADR-*-*.md` |
| `spec` | `<records_root>/spec/**/*.md` |
| `investigation` | `<records_root>/investigations/*/<record_prefix>INV-*-*.md` |
| `requirement` | `<records_root>/requirements/*/<record_prefix>REQ-*-*.md` |
| `work_item` | `<records_root>/work-items/*/<record_prefix>WORK-*-*.md` |
| `task` | `<records_root>/tasks/*/<record_prefix>TASK-*-*.md` |

New ADRs use the domain-subdirectory pattern.
Existing flat ADR records remain compatible through `<records_root>/adr/<record_prefix>ADR-*.md`.

`<record_prefix>` is an abstract filename placeholder for the app-aware record ID prefix.
This contract does not define how a tool derives or filters that prefix.

## Rules

- Path patterns describe repository placement, not tool indexing behavior.
- Path patterns use a `records_root` supplied by the caller or implementation context.
- Spec discovery uses the topic tree under `<records_root>/spec/`.
- Sequential record discovery uses the kind and domain subdirectories defined by the repository layout model.
- Tool-specific inclusion filters are outside this contract.
- Tool-specific namespace-prefix derivation is outside this contract.

## DRMCP boundary

The following implementation-specific concerns remain outside PRODUCT normative text.

| implementation-specific concern | app-local owner |
|---|---|
| Namespace-prefix derivation from `records_root`. | `spec:drmcp.design_records_mcp.namespace_scanning`. |
| DRMCP-specific index inclusion conditions. | `spec:drmcp.design_records_mcp.schema.discovery`. |
| DRMCP-specific discovery provenance. | DRMCP app-local specifications, if retained. |

## Related specs

| ref | relation |
|---|---|
| `spec:product.design_records.repository_layout` | Parent repository-layout overview. |

## Sources

- V01-ADR-076 section bootstrap policy.
- V01-ADR-092 section 1.
