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
Normalize the `propose_record_create` contract so callers can create design records without duplicating or guessing record identity.

This work item completes `REQ-MCP-014` by making top-level `id` the canonical create target input, allowing `fields + body` create proposals, and ensuring MCP-generated H1 / metadata use the resolved target ID while caller-supplied `body` contains only content sections.

## Boundary
This work item owns the `propose_record_create` contract cleanup for schema exposure, public spec/guidance, implementation, tests, and runtime smoke evidence.

In scope:

- `propose_record_create` schema and validation
- `fields.id` compatibility and rejection rules
- `fields + body` section-only body create mode
- MCP-generated H1 and metadata from resolved ID
- legacy `body-only` / `body_cache_id-only` full-record mode separation
- runtime smoke and close evidence for `REQ-MCP-014`

Out of scope:

- `propose_record_update` contract changes
- `REQ-MCP-015` cache/retry expansion
- making `fields + body_cache_id` valid
- unrelated DATA / UC / ADR changes
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
