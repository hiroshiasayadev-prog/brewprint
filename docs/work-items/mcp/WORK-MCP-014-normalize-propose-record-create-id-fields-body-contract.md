# WORK-MCP-014: Normalize propose_record_create id/fields/body contract

- **id**: WORK-MCP-014
- **status**: done
- **date**: 2026-06-02
- **source_requirement**: REQ-MCP-014
- **impact_refs**:
  - SPEC-design-records-mcp-tools
- **tasks**:
  - TASK-MCP-014-01
  - TASK-MCP-014-02
  - TASK-MCP-014-03
  - TASK-MCP-014-04

## Goal

## Boundary

## Evidence
WORK-MCP-014 is closed on 2026-06-02 because all owned tasks have completed and the runtime contract was verified.

Completed task evidence:

- `TASK-MCP-014-01` completed the contract inventory and gap classification.
- `TASK-MCP-014-02` updated `SPEC-design-records-mcp-tools` and authoring guidance for the `fields + body` create contract.
- `TASK-MCP-014-03` implemented schema, validation, rendering, and regression tests for the normalized contract.
- `TASK-MCP-014-04` completed runtime JSON-RPC smoke evidence, tests, validation, and close synchronization.

Close verification:

- Runtime smoke through `go run ./cmd/design-records-mcp --root .` passed all requested `propose_record_create` cases.
- `go test ./internal/designrecords ./internal/designrecordsmcp` passed on 2026-06-02.
- `go test ./...` passed on 2026-06-02.
- Design Records MCP validation for `REQ-MCP-014`, `WORK-MCP-014`, and `TASK-MCP-014-01..TASK-MCP-014-04` passed with no diagnostics.

Remaining out-of-scope work:

- REQ-MCP-015 cache/retry expansion remains separate.
- `fields + body_cache_id` remains invalid in this REQ-MCP-014 / WORK-MCP-014 scope.
