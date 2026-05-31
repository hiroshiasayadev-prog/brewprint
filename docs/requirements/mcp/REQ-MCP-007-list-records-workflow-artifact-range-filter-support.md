# REQ-MCP-007: list_records workflow artifact range filter support

- **id**: REQ-MCP-007
- **status**: accepted
- **date**: 2026-06-01
- **source_refs**:
  - SPEC-design-records-mcp-tools
  - WORK-DATA-004
- **work_items**:
  - WORK-MCP-007

## Problem

`list_records` exposes an `id_range` request field, but the current Design Records MCP contract limits `id_range` to `ADR-NNN` decision records only.

This is valid under the current MVP contract, but it is confusing during workflow artifact navigation because requirement, work item, and task IDs are also numbered and commonly inspected as local ranges.

For example, checking whether `WORK-DATA-004` is already occupied should be expressible as a workflow artifact query rather than falling back to filesystem directory listing or exact-ID probing.

## Need

Design Records MCP should support an explicit and safe way to list workflow artifacts by ID range or equivalent domain-scoped sequence filter.

The contract must avoid ambiguous cross-domain comparisons such as `WORK-DATA-*` versus `WORK-MCP-*`, and must not silently reinterpret unsupported ranges.

## Candidate direction

One acceptable direction is to extend `list_records.id_range` beyond decision records only when both range endpoints belong to the same ID family and domain.

Examples of potentially valid ranges:

- `REQ-MCP-001` .. `REQ-MCP-010`
- `WORK-DATA-001` .. `WORK-DATA-010`
- `TASK-DATA-004-01` .. `TASK-DATA-004-99`

Examples that should remain invalid unless a later contract explicitly defines ordering:

- `WORK-DATA-001` .. `WORK-MCP-010`
- `REQ-DATA-001` .. `TASK-DATA-010`
- `SPEC-*` ranges
- mixed workflow artifact kinds

An alternative direction is to keep `id_range` decision-only and introduce a more explicit workflow artifact range/filter field. The chosen direction should be decided by a follow-up work item.

## Scope

This requirement captures the usability and contract gap only.

It does not implement the range filter, update the MCP tool schema, update tests, or change existing `list_records` behavior.

## Acceptance expectation

A follow-up work item should decide and implement one of the following:

- extend `list_records.id_range` to workflow artifact IDs with explicit same-family ordering rules, or
- rename/split the current decision-only range contract and introduce a separate workflow artifact navigation filter.

The final behavior should make workflow artifact range navigation possible without relying on filesystem listing.

## Close evidence

- Implemented by `WORK-MCP-007`.
- `list_records.id_range` now supports `REQ-<DOMAIN>-NNN`, `WORK-<DOMAIN>-NNN`, and `TASK-<DOMAIN>-NNN-MM` in addition to `ADR-NNN`.
- Workflow artifact ranges are same-family and same-domain; task ranges are additionally same-work-sequence.
- `validate_records.id_range` uses the same range selection contract.
- Mixed-domain, mixed-family, mixed-task-work-sequence, malformed, and unsupported family ranges return `invalid_id_range`.
- `go test ./internal/designrecords ./internal/designrecordsmcp` passed on Windows PowerShell.
- Runtime MCP verification passed for valid workflow ranges, invalid workflow ranges, `validate_records` task range, and ADR range compatibility.
