# TASK-MCP-023-03: Implement accurate update diffs, no-op detection, and regression tests

- **id**: TASK-MCP-023-03
- **status**: done
- **date**: 2026-06-07
- **work_item**: WORK-MCP-023
- **source_requirement**: REQ-MCP-027
- **estimate**: 1.5d
- **depends_on**:
  - TASK-MCP-023-01
  - TASK-MCP-023-02
- **outputs**:
  - implementation patch
  - regression tests for update diff output
  - regression tests for no-op update behavior

## Goal

Implement the accurate update diff and no-op detection behavior required by REQ-MCP-027.

## Work

- Update the authoring proposal path so existing-file update diffs are generated from current persisted content to proposed content.
- Ensure modify diffs are git-style unified diffs with bounded context.
- Add no-op detection after applying `metadata_block_replace`, `metadata_fields_replace`, and `named_section_replace` semantics.
- Ensure no-op update requests do not retain a proposal.
- Return clear no-op response fields or diagnostics according to the spec update.
- Preserve existing proposal retention, validation, and accept guard behavior for real changes.
- Add regression tests for metadata-only real changes, no-op metadata updates, and named-section update behavior.

## Done condition

- Existing-file real updates produce reviewable bounded modify diffs.
- No-op update requests produce no retained write proposal.
- Existing create/add proposal diff behavior remains valid.
- Existing accept guards and affected-record validation behavior remain intact.
- Targeted and package-level tests pass.

## Verification

- Run targeted authoring tests for diff and no-op behavior.
- Run `go test ./internal/designrecords ./internal/designrecordsmcp`.
- Add any MCP boundary tests needed for response JSON shape.

## Evidence
Verdict: PASS.

Repository files were changed for TASK-MCP-023-03 implementation only.

Files changed:

- `go.mod`
- `go.sum`
- `internal/designrecords/authoring.go`
- `internal/designrecords/types.go`
- `internal/designrecords/authoring_test.go`
- `internal/designrecordsmcp/tools_call_test.go`
- `docs/tasks/mcp/TASK-MCP-023-03-implement-accurate-update-diffs-no-op-detection-and-regression-tests.md` after this evidence/status update

Implementation summary:

- Existing-file update proposals now retain persisted base content on modify `ProposedFile` values.
- Modify proposal diffs now compare base content to proposed content and emit git-style unified diff text with `diff --git`, `index <oldhash>..<newhash> 100644`, `--- a/<path>`, `+++ b/<path>`, and bounded hunks.
- Diff hunk generation uses `github.com/pmezard/go-difflib/difflib` instead of a hand-rolled diff algorithm.
- Create/add proposal diffs still use `/dev/null` and whole-file additions.
- No-op update detection now runs after update operation semantics and diagnostics, before retained proposal persistence.
- No-op responses return `proposal_created:false`, `operation:"update"`, target identity, successful validation when no validation errors exist, and a top-level `no_op_update` info diagnostic with record ID and path.
- No-op responses omit `proposal_id` and `diff`.
- Existing accept-time stale protection remains based on `BaseHash` and was preserved.

Tests added/updated:

- Metadata fields real update asserts git-style bounded modify diff, old/new status lines, and no whole-file addition output.
- Metadata fields no-op suppresses retained proposal creation.
- Metadata block no-op suppresses retained proposal creation.
- Named section real update asserts bounded modify diff with old/new section content.
- Named section no-op after heading-safe body normalization returns no-op while preserving the heading-stripped warning diagnostic.
- Create/add diff regression asserts `/dev/null` add diff behavior remains valid.
- Multi-hunk modify diff regression asserts separate hunks are emitted.
- MCP tool-call boundary test asserts no-op JSON response omits `proposal_id` and `diff` and includes `no_op_update` info diagnostic.

Commands run and results:

- `git status --short`: dirty worktree existed before implementation, including unrelated `CLAUDE.md`, spec/work/requirement/task docs, and `tmp.py` changes.
- `rg "type ProposedFile|buildDiff|persistProposal|prepareUpdate|ProposeRecordUpdate|BaseHash|contentHash|metadata_fields_replace|replaceNamedSection|section_replacement_body_heading_stripped" internal/designrecords internal/designrecordsmcp`: located implementation and test targets.
- `go get github.com/pmezard/go-difflib@v1.0.0`: passed; added `github.com/pmezard/go-difflib v1.0.0`.
- `gofmt -w internal/designrecords/authoring.go internal/designrecords/types.go internal/designrecords/authoring_test.go internal/designrecordsmcp/tools_call_test.go`: passed.
- `go test ./internal/designrecords -run "Diff|NoOp|Authoring|Metadata|Section" -v`: passed after correcting test assertions to match bounded unified diff output.
- `go test ./internal/designrecordsmcp -run "NoOp|AuthoringTransaction" -v`: passed.
- `go test ./internal/designrecords ./internal/designrecordsmcp`: passed.
- `go test ./...`: passed.
- `go mod tidy`: passed; kept `github.com/pmezard/go-difflib v1.0.0` as a direct requirement.
- Final `go test ./internal/designrecords ./internal/designrecordsmcp`: passed.
- Final `go test ./...`: passed.
- Design Records validation for `TASK-MCP-023-03`: passed.

Known unrelated failures:

- None observed in final test runs.

Boundary notes:

- `TASK-MCP-023-04` was not closed.
- Runtime smoke close synchronization was not run.
- Unrelated DATA tasks were not modified.
- No closed scopes were reopened.
- Semantics of `metadata_block_replace`, `metadata_fields_replace`, and `named_section_replace` were not changed.
- Accept-time stale guards were preserved.
