# V01-TASK-MCP-014-03: Implement fields plus body create schema validation rendering and tests

- **id**: V01-TASK-MCP-014-03
- **status**: done
- **date**: 2026-06-02
- **work_item**: V01-WORK-MCP-014
- **source_requirement**: V01-REQ-MCP-014
- **estimate**: 1d-2d
- **depends_on**:
  - V01-TASK-MCP-014-02
- **outputs**:
  - Updated propose_record_create schema exposure
  - Updated authoring validation and rendering
  - Regression tests for fields plus body and invalid body source cases

## Goal

## Work

## Done condition

## Verification

## Evidence
- Changed `internal/designrecordsmcp/tools.go` so `propose_record_create` still exposes `kind`, `id`, `domain`, `parent_id`, `title`, `fields`, `body`, `body_cache_id`, and `reciprocal_update_mode`, while schema-required inputs are now only `kind`, `id`, and `title`.
- Changed `internal/designrecords/authoring.go` so structured create allows `fields + body`, rejects `fields + body_cache_id`, preserves legacy full-record `body` / `body_cache_id` create modes, renders H1 and metadata from the resolved target ID, and appends caller section-only body content.
- Added structured-body validation that rejects caller bodies with a leading H1, leading YAML metadata, leading bullet metadata block, metadata `id`, or a guessed resolved ID when the top-level ID uses `new`.
- Updated `internal/designrecords/authoring_test.go` for fields-only without `fields.id`, `fields + body`, resolved `REQ-MCP-new` rendering, matching and mismatching `fields.id`, invalid body source combinations, section-body rejection cases, and legacy body/body-cache create compatibility.
- Updated `internal/designrecordsmcp/tools_call_test.go` to verify `fields` is no longer listed in the `propose_record_create` required schema inputs.
- Ran `gofmt -w internal/designrecords/authoring.go internal/designrecords/authoring_test.go internal/designrecordsmcp/tools.go internal/designrecordsmcp/tools_call_test.go`.
- `go test ./internal/designrecords ./internal/designrecordsmcp` passed on 2026-06-02.
- `go test ./...` passed on 2026-06-02.
- Remaining work intentionally left to V01-TASK-MCP-014-04: runtime smoke evidence and TASK/WORK/REQ close synchronization.
- V01-REQ-MCP-015 retry/cache expansion was not implemented here; `fields + body_cache_id` remains invalid for this task boundary.
