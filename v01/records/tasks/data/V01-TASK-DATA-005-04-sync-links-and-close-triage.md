# V01-TASK-DATA-005-04: Sync links and close triage

- **id**: V01-TASK-DATA-005-04
- **status**: done
- **date**: 2026-06-01
- **work_item**: V01-WORK-DATA-005
- **source_requirement**: V01-REQ-DATA-002
- **estimate**: 0.5d
- **depends_on**:
  - V01-TASK-DATA-005-03
- **outputs**:
  - Source requirement / successor artifact link synchronization
  - V01-WORK-DATA-005 close evidence

## Goal

Synchronize requirement and work item links after the successor split is decided, then prepare V01-WORK-DATA-005 for close.

## Work

- Update source requirement links only where the successor split requires it.
- Confirm newly created or selected successor artifacts are discoverable by ID-as-ref.
- Record close evidence for V01-WORK-DATA-005.
- Confirm no implementation or completed work item scope was reopened.

## Done Condition

- V01-REQ-DATA-002 and any other affected requirements point to the correct successor work items.
- V01-WORK-DATA-005 has close evidence showing that all deferred items have next actions.
- No deferred item remains orphaned only in historical notes.
- No implementation, fixture, golden, renderer, validator, parser, or MCP tool behavior is changed.

## Verification

- Use Design Records MCP to retrieve the affected requirements and work items.
- Confirm V01-WORK-DATA-005 task relation is visible.
- Confirm completed work item boundaries remain closed.

## Evidence

Completed on 2026-06-01.

### Link synchronization

- `V01-REQ-DATA-002.work_items` was synchronized to include:
  - `V01-WORK-DATA-002`
  - `V01-WORK-DATA-003`
  - `V01-WORK-DATA-005`
  - `V01-WORK-DATA-006`
  - `V01-WORK-DATA-008`
  - `V01-WORK-DATA-009`
- `V01-REQ-DATA-004.work_items` points to `V01-WORK-DATA-010` for tagged union / discriminator payload support.
- `V01-REQ-DATA-005.work_items` points to `V01-WORK-DATA-007` for DAG asset TypeRef hint render support.
- MCP semantic identity / state machine identity uses the already existing `V01-REQ-MCP-004` / `V01-WORK-MCP-004`; no duplicate MCP identity requirement or work item is created.
- `V01-REQ-MCP-007` / `V01-WORK-MCP-007` remain the separate list_records workflow artifact range filter support artifacts.

### Close input for V01-WORK-DATA-005

All deferred buckets now have successor handling:

- V01-ADR-073 tagged union / discriminator payload: `V01-REQ-DATA-004` / `V01-WORK-DATA-010`.
- V01-ADR-074 DAG asset TypeRef hint: `V01-REQ-DATA-005` / `V01-WORK-DATA-007`.
- V01-ADR-078 / V01-ADR-079 / V01-ADR-080 MCP semantic identity / state machine identity: existing `V01-REQ-MCP-004` / `V01-WORK-MCP-004`.
- UC-002 duplicate task QID / unresolved flow task issue: `V01-WORK-DATA-008` under `V01-REQ-DATA-002`.
- Remaining UC-002 notes retreat debt: `V01-WORK-DATA-009` under `V01-REQ-DATA-002`.

### Verification result

- Design Records MCP retrieval confirmed `V01-REQ-DATA-004`, `V01-WORK-DATA-010`, `V01-REQ-MCP-004`, `V01-WORK-MCP-004`, `V01-REQ-MCP-007`, and `V01-WORK-MCP-007` resolve to the intended records.
- Requirement relation validation passed after creating `V01-REQ-DATA-004`, `V01-REQ-DATA-005`, and confirming existing MCP identity artifacts.
- Work item relation validation initially reported `V01-REQ-DATA-002.work_items` was missing `V01-WORK-DATA-006`, `V01-WORK-DATA-008`, and `V01-WORK-DATA-009`; this task synchronized those links.
- No implementation, fixture, golden, renderer, validator, parser, MCP tool schema, UC-002 YAML, or render output was changed.
- V01-WORK-DATA-002, V01-WORK-DATA-003, and V01-WORK-DATA-004 were not reopened.
