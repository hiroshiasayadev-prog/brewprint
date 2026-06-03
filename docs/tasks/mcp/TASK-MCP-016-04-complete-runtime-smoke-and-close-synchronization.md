# TASK-MCP-016-04: Complete runtime smoke and close synchronization

- **id**: TASK-MCP-016-04
- **status**: done
- **date**: 2026-06-03
- **work_item**: WORK-MCP-016
- **source_requirement**: REQ-MCP-017
- **estimate**: 0.5d-1d
- **depends_on**:
  - TASK-MCP-016-03
- **outputs**:
  - Runtime smoke evidence and close synchronization for REQ-MCP-017 and WORK-MCP-016

## Goal

Verify the implemented validation through runtime smoke and close the workflow artifacts when evidence is complete.

## Work

- Run targeted Go tests for design record validation and MCP tool behavior.
- Run runtime smoke through `go run ./cmd/design-records-mcp --root .` or the current equivalent command.
- Verify that validation reports the required section diagnostic for close-state empty section cases.
- Update task/work/requirement evidence and statuses after implementation is verified.

## Done condition

This task is done when runtime evidence is recorded and the requirement/work item can be closed consistently.

## Verification

- `go test ./internal/designrecords ./internal/designrecordsmcp`
- `go test ./...` if the implementation scope or touched files justify a full run
- MCP runtime smoke for the validation path
- `validate_records` for `REQ-MCP-017`, `WORK-MCP-016`, and the `TASK-MCP-016-*` task set created for this work item

## Evidence
Completed on 2026-06-03.

Runtime smoke summary:

- `go test ./internal/designrecords ./internal/designrecordsmcp` passed.
- `go test ./...` passed.
- MCP runtime smoke through `go run ./cmd/design-records-mcp --root .` passed using newline-delimited JSON-RPC requests.
- `initialize` returned server info for `brewprint-design-records-mcp` version `0.1.0`.
- `validate_records` returned `ok:true` / `diagnostics:null` for:
  - `TASK-MCP-016-01..TASK-MCP-016-04`
  - `WORK-MCP-016`
  - `REQ-MCP-017`
- MCP stderr was empty.
- MCP process exit code was `0`.
- Smoke script counted `ok:true` responses: `3`.
- Final smoke output: `runtime smoke PASS`.

Protocol note:

- The first smoke attempt used `Content-Length` framed input and produced JSON-RPC parse errors because this MCP runtime reads newline-delimited JSON-RPC from stdin.
- The successful smoke used newline-delimited JSON-RPC and is the accepted runtime evidence.

Close synchronization:

- `TASK-MCP-016-01` and `TASK-MCP-016-02` had already completed the policy/spec contract work.
- `TASK-MCP-016-03` completed implementation and regression tests.
- This task records runtime smoke evidence and enables `WORK-MCP-016` / `REQ-MCP-017` close synchronization.
