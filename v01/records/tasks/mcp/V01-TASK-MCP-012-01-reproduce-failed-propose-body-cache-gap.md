# V01-TASK-MCP-012-01: reproduce failed-propose body cache gap

- **id**: V01-TASK-MCP-012-01
- **status**: done
- **date**: 2026-06-02
- **work_item**: V01-WORK-MCP-012
- **source_requirement**: V01-REQ-MCP-015
- **estimate**: 0.5d-1d
- **depends_on**:
- **outputs**:
  - Reproduction notes for propose_record_create / propose_record_update failure cases
  - Classification of request-level invalid cases vs body-after-receive preparation failures

## Goal
Reproduce the failed-propose body-cache behavior reported by V01-REQ-MCP-015 and classify which failures are request-level rejections versus failures after the submitted body has been received as a content source.

## Work
- Inspect the current authoring implementation and tests around `body_cache`, `body_cache_id`, `proposal_created`, `propose_record_create`, and `propose_record_update`.
- Exercise representative failed-propose cases for update and create paths.
- Separate request-level invalid cases from failures that occur after the submitted body is available as a content source.
- Report whether the originally suspected missing-cache gap reproduces and identify the implementation locations for follow-up regression coverage.

## Done condition
- Reproduction result is recorded with enough detail to tell whether the missing-cache gap exists.
- Classification distinguishes request-level invalid cases from post-body-receipt failures.
- Follow-up work can add regression coverage without reopening the reproduction question.

## Verification
- Reviewed implementation/test locations around authoring body cache handling.
- Ran/received reproduction results for focused tests and JSON-RPC smoke during the classification pass.
- Confirmed no persistent file modifications remained from the temporary reproduction test.

## Evidence
V01-TASK-MCP-012-01 completed the reproduction/classification pass for V01-REQ-MCP-015.

Review result was `PARTIAL` because the original suspected update-side missing-cache gap did not reproduce: current implementation already returned `body_cache` for named-section no-match and ambiguous update failures.

Important finding:

- `propose_record_update` named-section no-match with submitted `body` returned `proposal_created:false` with `section_selector_no_match` and `body_cache`.
- `propose_record_update` named-section ambiguous with submitted `body` returned `proposal_created:false` with `section_selector_ambiguous` and `body_cache`.
- Create-side classification had to be interpreted after V01-WORK-MCP-014: `fields + body` is valid when body is section-only content; full-record body in `fields + body` mode is rejected after the submitted body is received and may return `body_cache`.

Commands reported from the reproduction pass:

```text
git status --short
rg -n "body_cache|body_cache_id|proposal_created|propose_record_create|propose_record_update|invalid_body_source|proposal_preparation_failed|invalid_request" internal cmd docs
go test ./internal/designrecordsmcp ./internal/designrecords
JSON-RPC smoke via go run ./cmd/design-records-mcp --root .
go test ./internal/designrecords -run TestMCP012TempAmbiguousSelectorReturnsBodyCache -count=1 -v
```

A temporary focused test file was added and removed during the investigation. No files remained modified by this reproduction work.
