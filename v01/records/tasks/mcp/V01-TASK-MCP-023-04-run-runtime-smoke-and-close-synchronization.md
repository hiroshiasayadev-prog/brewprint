# V01-TASK-MCP-023-04: Run runtime smoke and close synchronization

- **id**: V01-TASK-MCP-023-04
- **status**: done
- **date**: 2026-06-07
- **work_item**: V01-WORK-MCP-023
- **source_requirement**: V01-REQ-MCP-027
- **estimate**: 0.5d
- **depends_on**:
  - V01-TASK-MCP-023-03
- **outputs**:
  - runtime smoke evidence
  - V01-REQ-MCP-027 / V01-WORK-MCP-023 / task close synchronization

## Goal

Run runtime smoke for accurate update diffs and no-op update handling, then synchronize workflow artifact statuses when all acceptance criteria are met.

## Work

- Run the relevant Go test set after implementation is complete.
- Run MCP runtime smoke for a real `propose_record_update` metadata change and confirm the diff is a bounded modify diff, not a whole-file addition.
- Run MCP runtime smoke for a no-op `propose_record_update` and confirm no retained proposal is created.
- Confirm existing proposal accept guards still behave as expected for real update proposals.
- Record evidence in this task.
- Synchronize final statuses for `TASK-MCP-023-*`, `V01-WORK-MCP-023`, and `V01-REQ-MCP-027` when acceptance criteria are met.

## Done condition

- Runtime smoke passes for real update diff output.
- Runtime smoke passes for no-op update handling.
- Test commands and results are recorded.
- Related workflow artifacts are status-synchronized.
- No unexpected repository-wide validation errors are introduced by this work.

## Verification

- Run targeted tests and MCP runtime smoke.
- Run Design Records validation for the affected workflow artifacts.

## Evidence

Verdict: PASS.

Repository files were changed only for close synchronization after runtime smoke passed.

Files changed for close synchronization:

- `docs/tasks/mcp/TASK-MCP-023-04-run-runtime-smoke-and-close-synchronization.md`
- `docs/work-items/mcp/WORK-MCP-023-fix-propose-record-update-diff-output-and-no-op-detection.md`
- `docs/requirements/mcp/REQ-MCP-027-propose-record-update-git-style-accurate-diff-and-no-op-detection.md`

Tests run and results:

- `git status --short`: dirty worktree existed before this task, including unrelated `CLAUDE.md`, implementation files from V01-TASK-MCP-023-03, untracked workflow docs, and `tmp.py`.
- `go test ./internal/designrecords -run "Diff|NoOp|Authoring|Metadata|Section" -v`: passed.
- `go test ./internal/designrecordsmcp -run "NoOp|AuthoringTransaction" -v`: passed.
- `go test ./internal/designrecords ./internal/designrecordsmcp`: passed.
- `go test ./...`: passed.

Runtime smoke setup:

- Created a temporary root under `%TEMP%` and copied the current repository `docs/` tree into it.
- Launched the repo-local server with `go run ./cmd/design-records-mcp --root <temp-root>`.
- Drove the server over newline-delimited JSON-RPC using `initialize`, `notifications/initialized`, and `tools/call` requests.
- The accepted runtime smoke write modified only the temporary copied docs tree, not the real repository.

Runtime smoke request/response summary:

- Real `propose_record_update` `metadata_fields_replace` for `V01-TASK-MCP-023-04` status `todo` to `in_progress` returned `proposal_created: true` with a retained `proposal_id`.
- The real update `diff.text` was a git-style modify diff containing `diff --git a/`, `index`, `--- a/`, `+++ b/`, `@@`, the old status line, and the new status line.
- The real update diff was bounded: 13 total diff lines and one added content line. It did not render the whole target file as only `+` lines.
- No-op `propose_record_update` `metadata_fields_replace` for `V01-TASK-MCP-023-03` status `done` returned `proposal_created: false`, no `proposal_id`, no `diff`, `validation.ok: true`, and a top-level `no_op_update` diagnostic with severity `info`.
- Optional `named_section_replace` no-op for `V01-TASK-MCP-023-04` `Evidence`, using a replacement body that started with the same `## Evidence` heading, returned `proposal_created: false`, no `proposal_id`, preserved `section_replacement_body_heading_stripped`, and included `no_op_update`.
- A real update proposal for `V01-TASK-MCP-023-04` status `blocked` was accepted in the temp root; `accept_proposed_write` returned `written: true`, and `get_record` confirmed the target metadata changed to `blocked`.
- BaseHash stale protection was not bypassed: accepting an older retained proposal for the same temp-root target after the newer write returned `written: false` with `proposal_stale`.
- The no-op update could not be accepted because the no-op response contained no `proposal_id`.

Close synchronization:

- `V01-TASK-MCP-023-04` status set to `done`.
- `V01-WORK-MCP-023` status set to `done`.
- `V01-REQ-MCP-027` status set to `accepted`.
- `V01-REQ-MCP-029` was not touched.

Design Records validation after close synchronization:

- `validate_records` for `V01-TASK-MCP-023-04`: passed.
- `validate_records` for `V01-WORK-MCP-023`: passed.
- `validate_records` for `V01-REQ-MCP-027`: passed.
- `validate_records` for `V01-TASK-MCP-023-01..04`: passed.

Boundary notes:

- The installed Design Records MCP app produced an old whole-file addition diff for a dry-run `V01-WORK-MCP-023` status proposal, conflicting with the repo-local CLI smoke result. That app proposal was discarded and not accepted.
- `V01-WORK-MCP-023` needed an `Evidence` section before status `done`; without it, proposal-local validation reported `missing_required_section` for `Evidence`.
