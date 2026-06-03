# TASK-MCP-017-04: Runtime smoke close evidence and status synchronization

- **id**: TASK-MCP-017-04
- **status**: done
- **date**: 2026-06-03
- **work_item**: WORK-MCP-017
- **source_requirement**: REQ-MCP-018
- **estimate**: 0.5d-1d
- **depends_on**:
  - TASK-MCP-017-03
- **outputs**:
  - Runtime smoke evidence for exact ID gap warning behavior
  - Updated TASK WORK REQ close evidence and statuses

## Goal

Verify the exact ID gap-warning behavior through runtime smoke and close REQ-MCP-018 / WORK-MCP-017 when all owned work is complete.

## Work

- Run targeted tests and broad MCP package tests.
- Run runtime smoke through the Design Records MCP server for risky exact ID create and normal new placeholder create.
- Confirm warning feedback is non-blocking.
- Synchronize task, work item, and requirement evidence and statuses.

## Done condition

Runtime smoke and tests pass, and REQ-MCP-018 is accepted with close evidence linking the completed work item.

## Verification

- go test ./internal/designrecords ./internal/designrecordsmcp
- go test ./...
- Runtime JSON-RPC smoke through go run ./cmd/design-records-mcp --root .
- Design Records MCP validation for REQ-MCP-018, WORK-MCP-017, and owned TASK records

## Evidence
Runtime smoke and close synchronization evidence recorded from the 2026-06-03 PowerShell run.

Tests:

- `go test ./internal/designrecords ./internal/designrecordsmcp` — PASS
- `go test -run "TestExactIDSequenceGapWarning|TestToolsCallProposeExactWorkIDGapDiagnostic" -v ./internal/designrecords ./internal/designrecordsmcp` — PASS

Runtime JSON-RPC smoke through `go run ./cmd/design-records-mcp --root .`:

- `initialize` succeeded.
- Exact ID create probe: `WORK-MCP-999`
  - `proposal_created: true`
  - top-level `diagnostics` contained `exact_id_sequence_gap`
  - diagnostic severity was `info`
  - `validation.ok: true`
  - `validation.diagnostics: []`
  - proposal was not accepted.
- New placeholder create probe: `WORK-MCP-new`
  - `proposal_created: true`
  - resolved to `WORK-MCP-019`
  - top-level `diagnostics: []`
  - no `exact_id_sequence_gap`
  - proposal was not accepted.

Post-smoke repository check:

- `git status --short` showed no smoke-created `WORK-MCP-999` or `WORK-MCP-019` files.
- Existing dirty worktree entries were pre-existing task/work/spec/implementation changes, not runtime-smoke writes.

Close synchronization:

- TASK-MCP-017-01, TASK-MCP-017-02, TASK-MCP-017-03, and TASK-MCP-017-04 are complete.
- WORK-MCP-017 can be closed as done.
- REQ-MCP-018 can be accepted because the required non-blocking exact ID gap warning behavior is implemented, tested, documented, and runtime-smoke verified.
