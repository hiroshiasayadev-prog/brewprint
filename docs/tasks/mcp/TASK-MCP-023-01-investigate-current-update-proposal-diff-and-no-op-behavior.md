# TASK-MCP-023-01: Investigate current update proposal diff and no-op behavior

- **id**: TASK-MCP-023-01
- **status**: done
- **date**: 2026-06-07
- **work_item**: WORK-MCP-023
- **source_requirement**: REQ-MCP-027
- **estimate**: 0.5d
- **depends_on**:
- **outputs**:
  - investigation notes for current diff generation path
  - identified no-op detection gap
  - implementation and spec update recommendations

## Goal

Identify where `propose_record_update` builds proposed content, creates retained proposals, and generates `diff.text`, and confirm why existing-file update proposals can appear as whole-file additions.

## Work

- Inspect the authoring update flow for `metadata_block_replace`, `metadata_fields_replace`, and `named_section_replace`.
- Locate the diff generation helper used by update proposals and compare it with create proposal behavior.
- Check whether the diff helper has access to persisted current file content for modify proposals.
- Check whether byte-equivalent proposed content is detected before proposal retention.
- Identify response-shape implications for explicit no-op results.
- Record recommended implementation boundaries for TASK-MCP-023-02 and TASK-MCP-023-03.

## Done condition

- The current code path responsible for incorrect update diffs is identified.
- The current no-op behavior is classified as missing, partial, or already present but misplaced.
- Required implementation files and likely tests are listed.
- Any spec gaps are listed for the next task.

## Verification

- Use repo-local search, targeted code inspection, and existing tests where useful.
- Prefer command-backed evidence from grep, `go test`, or a small runtime reproduction when available.

## Evidence

Verdict: PASS for investigation; needs spec revision before implementation.

Repository files were not changed during the investigation.

Findings:

- Root cause: update proposals currently retain only proposed content for diff generation. `buildDiff` renders every proposed line as an added line, so existing-file modify proposals can appear as whole-file additions.
- Responsibility gap: create and update proposal paths share a create/add-oriented diff renderer. Create proposals may validly show all lines as additions, but modify proposals must compare current persisted content with proposed content.
- No-op behavior: missing. After `metadata_block_replace`, `metadata_fields_replace`, or `named_section_replace` applies operation semantics, there is no byte-equivalence check between current persisted content and proposed content before retaining a proposal.
- Spec gap: `SPEC-design-records-mcp-tools` does not yet define git-style modify diff requirements or a no-op update response shape.
- Test gap: existing tests assert proposed content presence, but do not assert bounded modify diff structure, old/new lines, hunk headers, or no-op proposal suppression.

Relevant implementation observations:

- `prepareUpdate` reads persisted file content before applying update semantics.
- Modify proposals store `BaseHash`, which supports accept-time stale protection.
- MCP JSON dispatch appears to pass the authoring response through directly; no JSON-layer distortion was identified.

No-op classification:

- Missing after successful update operation semantics.
- Existing `proposal_created:false` cases cover invalid request/body source, body cache failures, selector failures, or error diagnostics, but not no-op update success.

Recommended next steps:

- Proceed to `TASK-MCP-023-02` before implementation.
- Define modify `diff.text` as git-style unified diff comparing persisted current content to proposed content.
- Define no-op response behavior, including `proposal_created:false`, no retained `proposal_id`, and clear no-op diagnostics or response fields.
- Then implement in `TASK-MCP-023-03` by making diff generation change-aware and adding no-op detection before proposal retention.

Suggested implementation and test targets:

- `internal/designrecords/authoring.go`: retain/access base content for modify diff generation, add no-op detection, and replace create-only diff behavior for modify proposals.
- `internal/designrecords/authoring_test.go`: add metadata real-change diff, metadata no-op, metadata block no-op, named-section real-change, named-section no-op, and create proposal regression tests.
- `internal/designrecordsmcp/tools_call_test.go`: add MCP response JSON shape coverage for no-op updates.
- `docs/spec/design-records-mcp/tools.md`: define the response contract before implementation.

Commands and results from investigation:

- `git status --short`: dirty worktree existed before investigation; relevant new REQ/WORK/TASK docs were untracked in that local view.
- `git diff -- docs/requirements/...REQ-MCP-027... docs/work-items/...WORK-MCP-023... docs/tasks/mcp/TASK-MCP-023-*.md`: no diff output in Codex view because those files were untracked.
- `rg "propose_record_update|metadata_fields_replace|diff|proposal_created|no-op|noop|Retained|retain|unified" internal docs/spec/design-records-mcp`: located authoring flow, tests, MCP schema, and spec sections.
- `go test ./internal/designrecords ./internal/designrecordsmcp`: failed due existing status-vocabulary / undefined `designrecords.RecordStatusImplementationPending` issue, not caused by this investigation.
- `validate_records` for `TASK-MCP-023-01`: ok true.
