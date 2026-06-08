# V01-WORK-MCP-012: propose failure body_cache_id support

- **id**: V01-WORK-MCP-012
- **status**: done
- **date**: 2026-06-02
- **source_requirement**: V01-REQ-MCP-015
- **impact_refs**:
  - V01-ADR-093
  - SPEC-design-records-mcp-tools
- **tasks**:
  - V01-TASK-MCP-012-01
  - V01-TASK-MCP-012-02
  - V01-TASK-MCP-012-03

## Goal

## Boundary

## Evidence
V01-WORK-MCP-012 is closed on 2026-06-03 because V01-REQ-MCP-015 body-cache failure behavior is covered by reproduction, regression tests, and runtime smoke.

Completed task evidence:

- `V01-TASK-MCP-012-01` reproduced and classified failed-propose body-cache behavior. The originally suspected update-side missing-cache gap did not reproduce; update no-match and ambiguous failures already returned `body_cache`.
- `V01-TASK-MCP-012-02` added regression coverage for the classification and confirmed no implementation change was required for the V01-REQ-MCP-015 body-cache return behavior.
- `V01-TASK-MCP-012-03` completed runtime smoke through `go run ./cmd/design-records-mcp --root .` and confirmed JSON-RPC `tools/call` responses include `body_cache` for representative failed-propose cases after body receipt.

Close verification:

- `go test ./internal/designrecords ./internal/designrecordsmcp` passed.
- `go test ./...` passed across all 17 packages.
- Runtime smoke confirmed `propose_record_update` named-section no-match with submitted `body` returned `proposal_created:false`, `section_selector_no_match`, and `body_cache`.
- Runtime smoke confirmed `propose_record_create` `fields + body` with stale full-record body returned `proposal_created:false`, `invalid_request`, and `body_cache`.

Boundary kept for follow-up:

- `fields + body_cache_id` as a valid retry form for `fields + body` create changes the `propose_record_create` input contract and remains outside V01-WORK-MCP-012.
- Legacy full-record create deprecation / fields-required create contract tightening remains outside V01-WORK-MCP-012 and should be handled by a separate requirement.
