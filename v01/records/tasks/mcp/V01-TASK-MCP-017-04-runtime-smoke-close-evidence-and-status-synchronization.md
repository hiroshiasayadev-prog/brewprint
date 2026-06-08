# V01-TASK-MCP-017-04: Runtime smoke close evidence and status synchronization

- **id**: V01-TASK-MCP-017-04
- **status**: done
- **date**: 2026-06-03
- **work_item**: V01-WORK-MCP-017
- **source_requirement**: V01-REQ-MCP-018
- **estimate**: 0.5d-1d
- **depends_on**:
  - V01-TASK-MCP-017-03
- **outputs**:
  - Runtime smoke evidence for exact ID gap warning behavior
  - Updated TASK WORK REQ close evidence and statuses

## Goal

Verify the exact ID gap-warning behavior through runtime smoke and close V01-REQ-MCP-018 / V01-WORK-MCP-017 when all owned work is complete.

## Work

- Run targeted tests and broad MCP package tests.
- Run runtime smoke through the Design Records MCP server for risky exact ID create and normal new placeholder create.
- Confirm warning feedback is non-blocking.
- Synchronize task, work item, and requirement evidence and statuses.

## Done condition

Runtime smoke and tests pass, and V01-REQ-MCP-018 is accepted with close evidence linking the completed work item.

## Verification

- go test ./internal/designrecords ./internal/designrecordsmcp
- go test ./...
- Runtime JSON-RPC smoke through go run ./cmd/design-records-mcp --root .
- Design Records MCP validation for V01-REQ-MCP-018, V01-WORK-MCP-017, and owned TASK records

## Evidence
Runtime smoke and close synchronization evidence recorded from the 2026-06-03 PowerShell run.

Tests:

- `go test ./internal/designrecords ./internal/designrecordsmcp` — PASS
- `go test -run "TestExactIDSequenceGapWarning|TestToolsCallProposeExactWorkIDGapDiagnostic" -v ./internal/designrecords ./internal/designrecordsmcp` — PASS

Runtime JSON-RPC smoke through `go run ./cmd/design-records-mcp --root .`:

- `initialize` succeeded.
- Exact ID create probe: `V01-WORK-MCP-999`
  - `proposal_created: true`
  - top-level `diagnostics` contained `exact_id_sequence_gap`
  - diagnostic severity was `info`
  - `validation.ok: true`
  - `validation.diagnostics: []`
  - proposal was not accepted.
- New placeholder create probe: `WORK-MCP-new`
  - `proposal_created: true`
  - resolved to `V01-WORK-MCP-019`
  - top-level `diagnostics: []`
  - no `exact_id_sequence_gap`
  - proposal was not accepted.

Post-smoke repository check:

- `git status --short` showed no smoke-created `V01-WORK-MCP-999` or `V01-WORK-MCP-019` files.
- Existing dirty worktree entries were pre-existing task/work/spec/implementation changes, not runtime-smoke writes.

Close synchronization:

- V01-TASK-MCP-017-01, V01-TASK-MCP-017-02, V01-TASK-MCP-017-03, and V01-TASK-MCP-017-04 are complete.
- V01-WORK-MCP-017 can be closed as done.
- V01-REQ-MCP-018 can be accepted because the required non-blocking exact ID gap warning behavior is implemented, tested, documented, and runtime-smoke verified.
