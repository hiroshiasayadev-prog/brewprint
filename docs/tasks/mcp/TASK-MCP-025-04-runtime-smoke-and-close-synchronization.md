# TASK-MCP-025-04: Runtime smoke and close synchronization

- **id**: TASK-MCP-025-04
- **status**: not_started
- **date**: 2026-06-07
- **work_item**: WORK-MCP-025
- **source_requirement**: REQ-MCP-024
- **estimate**: 0.5d
- **depends_on**:
  - TASK-MCP-025-03
- **outputs**:
  - runtime smoke evidence
  - REQ-MCP-024 / REQ-MCP-028 / WORK-MCP-025 / task close synchronization

## Goal

Run runtime smoke for the new structured create diagnostic behavior, then synchronize workflow artifact statuses when all acceptance criteria are met.

## Work

- Update `tmp.py` to add smoke cases for REQ-MCP-024 and REQ-MCP-028:
  - `propose_record_create` with missing required fields → verify batch diagnostic lists all missing fields.
  - `propose_record_create` with invalid status → verify `invalid_metadata_value` with `allowed_values`.
  - `propose_record_create` in `report_required_follow_up` mode → verify clarity diagnostic.
- Run smoke against the repo-local MCP server.
- Record smoke results as Evidence.
- Synchronize final statuses for `TASK-MCP-025-*`, `WORK-MCP-025`, `REQ-MCP-024`, and `REQ-MCP-028`.

## Done condition

- Runtime smoke passes for all three new diagnostic behaviors.
- Test commands and results are recorded as Evidence.
- `REQ-MCP-024` status set to `accepted`.
- `REQ-MCP-028` status set to `accepted`.
- `WORK-MCP-025` status set to `done`.
- All tasks synchronized to `done`.

## Verification

- Run `go test ./...` after smoke.
- Run `validate_records` for affected artifacts.
