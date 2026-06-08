# V01-TASK-MCP-011-04: Regression tests and runtime smoke

- **id**: V01-TASK-MCP-011-04
- **status**: done
- **date**: 2026-06-02
- **work_item**: V01-WORK-MCP-011
- **source_requirement**: V01-REQ-MCP-011
- **estimate**: 0.5d
- **depends_on**:
  - V01-TASK-MCP-011-03
- **outputs**:
  - Go regression test results for Design Records MCP authoring behavior
  - JSON-RPC runtime smoke results for authoring transaction tool surface
  - Close verification evidence for V01-WORK-MCP-011, V01-REQ-MCP-011, and V01-REQ-MCP-012

## Goal

Verify the V01-TASK-MCP-011-03 implementation through regression tests, runtime MCP tool-surface smoke, proposal cleanup, and close-readiness checks.

## Work

- Run targeted authoring regression tests for `internal/designrecords` and `internal/designrecordsmcp`.
- Run full package tests for `internal/designrecords` and `internal/designrecordsmcp`.
- Start Design Records MCP from the repository root and exercise representative `propose_record_create`, `propose_record_update`, and `accept_proposed_write` JSON-RPC calls.
- Confirm create input normalization, ID/domain normalization, proposal-local validation diagnostics, and accept-time pre-write validation behavior.
- Do not fix unrelated repository validation errors.
- Record whether additional edge-case tests or response-shape review are needed.
- Update `V01-WORK-MCP-011`, `V01-REQ-MCP-011`, and `V01-REQ-MCP-012` close evidence if verification supports closing.

## Done condition

- Required Go regression commands pass, or failures are recorded with cause.
- Runtime smoke confirms the V01-TASK-MCP-011-03 behavior through the MCP JSON-RPC/tool surface.
- Unrelated `INV-DOCS` / `TASK-MCP-005` / other repository diagnostics do not pollute proposal-local validation diagnostics.
- Runtime-created proposals are not left as repository changes.
- Close status and relation metadata are synchronized for affected workflow artifacts where scope is clear.

## Verification

Required checks:

```powershell
go test ./internal/designrecords ./internal/designrecordsmcp -run Authoring
go test ./internal/designrecords ./internal/designrecordsmcp
```

Required record validation:

```text
validate_records(kind=task, id_range=V01-TASK-MCP-011-01..V01-TASK-MCP-011-04)
validate_records(kind=work_item, id_range=V01-WORK-MCP-011..V01-WORK-MCP-011)
validate_records(kind=requirement, id_range=V01-REQ-MCP-011..V01-REQ-MCP-012)
validate_records(kind=spec)
```

Final repository checks:

```powershell
git status --short
git diff --check
```

## Evidence

2026-06-02: Go regression verification passed.

- `go test ./internal/designrecords ./internal/designrecordsmcp -run Authoring`
  - `internal/designrecords`: pass
  - `internal/designrecordsmcp`: pass
- `go test ./internal/designrecords ./internal/designrecordsmcp`
  - `internal/designrecords`: pass
  - `internal/designrecordsmcp`: pass

2026-06-02: JSON-RPC runtime smoke was run against `go run ./cmd/design-records-mcp --root <repo root>`.

Confirmed behavior:

- fields-only `propose_record_create` succeeded for `V01-REQ-MCP-960` with proposal-local `validation.ok: true` and no diagnostics.
- structured create without `fields.id` succeeded for `V01-REQ-MCP-961`; top-level ID rendered as metadata ID.
- lowercase `domain: "mcp"` with `id: "V01-REQ-MCP-962"` succeeded; canonical target domain was `MCP` and path domain was `docs/requirements/mcp`.
- `body` + `fields` was rejected with `invalid_request` before proposal creation.
- `body_cache_id` + `fields` was rejected with `invalid_request` before proposal creation.
- `body` + `body_cache_id` remained rejected with `invalid_body_source`.
- top-level ID mismatch with `fields.id` was rejected with `invalid_request`.
- `REQ-MCP-new` placeholder create with supplied `fields.id` was rejected with `invalid_request`.
- Proposal-local validation diagnostics for successful create/update proposals were empty; unrelated existing `INV-DOCS`, `TASK-MCP-005`, and other repository validation diagnostics did not appear in proposal-local `validation.diagnostics`.
- `accept_proposed_write` was smoke-tested on an update proposal and returned `written: true`, `state: accepted`, `validation.ok: true`, and no diagnostics. The temporary accepted write was restored after the smoke.

Proposal cleanup:

- Create smoke proposals were not accepted and did not write repository files.
- The runtime server used process-local proposal storage; after process exit, no repository proposal artifacts remained.
- The single accepted update proposal was used only to verify accept-time pre-write validation and was restored immediately.

Edge-case judgment:

- No additional multi-file reciprocal proposal test is required to close this work. Existing regression coverage already checks reciprocal task-create proposal generation and proposal-local filtering of unrelated diagnostics; no runtime failure showed related-record diagnostics being suppressed.
- No `repository_health` or alternate response-shape review is required now because V01-TASK-MCP-011-03 did not add a separate repository-health field.
- No `V01-TASK-MCP-011-05` was created. If later changes add repository-health reporting or expose affected related-record diagnostics differently, a focused follow-up can be considered then.
