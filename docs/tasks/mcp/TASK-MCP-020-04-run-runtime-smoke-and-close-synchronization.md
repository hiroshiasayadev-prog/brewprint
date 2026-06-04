# TASK-MCP-020-04: Run runtime smoke and close synchronization

- **id**: TASK-MCP-020-04
- **status**: todo
- **date**: 2026-06-05
- **work_item**: WORK-MCP-020
- **source_requirement**: REQ-MCP-020
- **estimate**: 0.5d
- **depends_on**:
  - TASK-MCP-020-03
- **outputs**:
  - runtime smoke evidence
  - REQ-MCP-020 / WORK-MCP-020 / task close synchronization

## Goal

Run runtime smoke for the new metadata field replacement operation and synchronize close status for the related workflow artifacts.

## Work

- Run the relevant Go test set after implementation is complete.
- Run an MCP runtime smoke that updates only a task `status` field through the new field-level metadata replacement operation.
- Confirm unspecified metadata fields remain preserved after the update proposal / accept flow.
- Record evidence in this task.
- Synchronize final statuses for `TASK-MCP-020-*`, `WORK-MCP-020`, and `REQ-MCP-020` when acceptance criteria are met.

## Done condition

- Runtime smoke passes.
- Test commands and results are recorded.
- Related workflow artifacts are status-synchronized.
- No unexpected repository-wide validation errors are introduced by this work.

## Verification

- Run targeted tests and MCP runtime smoke.
- Run Design Records validation for the affected workflow artifacts.

## Evidence

Not started.
