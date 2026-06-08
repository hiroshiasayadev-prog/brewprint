# V01-TASK-MCP-014-04: Runtime smoke close evidence and status synchronization

- **id**: V01-TASK-MCP-014-04
- **status**: done
- **date**: 2026-06-02
- **work_item**: V01-WORK-MCP-014
- **source_requirement**: V01-REQ-MCP-014
- **estimate**: 0.5d-1d
- **depends_on**:
  - V01-TASK-MCP-014-03
- **outputs**:
  - Runtime smoke evidence for propose_record_create fields plus body contract
  - Updated TASK/WORK/REQ close evidence and statuses

## Goal

## Work

## Done condition

## Verification

## Evidence
Runtime smoke was run on 2026-06-02 through the actual stdio JSON-RPC MCP server path:

- Command path: `go run ./cmd/design-records-mcp --root .`
- Smoke assertions used `initialize`, `tools/list`, and `tools/call` for `propose_record_create`.
- No `accept_proposed_write` call was made during smoke; create proposals remained retained no-write proposals.

Runtime smoke results:

| Case | Result |
|---|---|
| `tools/list` schema | pass; `propose_record_create.required` was `kind`, `id`, `title`; `fields` was not required; `body` and `body_cache_id` were exposed |
| fields-only create | pass; proposal succeeded without `fields.id`; diff contained MCP-generated H1 and metadata |
| fields plus body create | pass; section-only body succeeded; diff contained MCP-generated H1, metadata with resolved ID, and caller content sections |
| server-side `REQ-MCP-new` plus fields plus body | pass; resolved to `V01-REQ-MCP-016`; generated H1 and metadata used the resolved ID; caller body did not need the resolved ID |
| exact ID plus matching `fields.id` | pass; compatibility path succeeded |
| exact ID plus mismatching `fields.id` | pass; rejected with `invalid_request` diagnostic |
| `REQ-MCP-new` plus `fields.id` | pass; rejected with `invalid_request` diagnostic |
| body plus `body_cache_id` | pass; rejected with `invalid_body_source` diagnostic |
| fields plus `body_cache_id` | pass; rejected with `invalid_request` diagnostic in V01-REQ-MCP-014 scope |
| body-only legacy full-record create | pass; proposal succeeded as compatibility mode, not the preferred path |
| fields plus body containing stale full-record input | pass; rejected with `invalid_request`; no duplicate H1/metadata was silently generated |

Verification commands run on 2026-06-02:

- `git status --short` completed; unrelated dirty files were present and were not staged or reverted.
- `git diff -- internal/designrecords internal/designrecordsmcp docs/tasks/mcp/TASK-MCP-014-04-runtime-smoke-close-evidence-and-status-synchronization.md docs/work-items/mcp/WORK-MCP-014-normalize-propose-record-create-id-fields-body-contract.md docs/requirements/mcp/REQ-MCP-014-propose-record-create-id-fields-body.md` completed; prior implementation diffs were visible before close edits.
- `go test ./internal/designrecords ./internal/designrecordsmcp` passed.
- `go test ./...` passed.
- Design Records MCP `validate_records(kind=requirement, id_range=V01-REQ-MCP-014..V01-REQ-MCP-014)` passed with no diagnostics.
- Design Records MCP `validate_records(kind=work_item, id_range=V01-WORK-MCP-014..V01-WORK-MCP-014)` passed with no diagnostics.
- Design Records MCP `validate_records(kind=task, id_range=V01-TASK-MCP-014-01..V01-TASK-MCP-014-04)` passed with no diagnostics.

V01-REQ-MCP-015 retry/cache expansion remains out of scope for this task. In particular, `fields + body_cache_id` remains invalid here.
