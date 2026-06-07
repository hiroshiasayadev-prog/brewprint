# TASK-MCP-029-01: Update authoring guides and tools spec for status-gated validation visibility

- **id**: TASK-MCP-029-01
- **status**: done
- **date**: 2026-06-07
- **work_item**: WORK-MCP-029
- **source_requirement**: REQ-MCP-032
- **estimate**: 0.5d
- **depends_on**:
- **outputs**:
  - docs/guides/requirement-authoring.md
  - docs/guides/work-item-authoring.md
  - docs/guides/task-authoring.md
  - docs/spec/design-records-mcp/tools.md

## Goal

Update the relevant authoring guides and MCP tools spec so valid statuses and status-gated required narrative sections are visible before agents attempt workflow artifact status transitions.

## Work

- Add `## Evidence` to the work item authoring format.
- State the `work_item status: done` required section gate.
- State the `task status: done` required section gate.
- State the `requirement status: accepted` required section gate.
- Add a kind-by-kind valid status matrix to the MCP tools spec.
- Prepare the resulting diff for Codex review.

## Done condition

- Documentation reflects the existing MCP validation behavior without changing semantics.
- The updated docs remove the observed ambiguity around metadata-only transition to `done` / `accepted`.
- Review has no blocking findings.

## Verification

- Static documentation review.
- Codex review of the changed guides/spec.
- `validate_records` after workflow artifact updates.

## Evidence

Documentation update and review completed.

Review result:

- Verdict: PASS_WITH_MINOR.
- Blocking findings: none.
- Minor finding 1: `docs/spec/design-records-mcp/tools.md` front matter `last_updated` was stale after adding the status matrix; updated from `2026-06-05` to `2026-06-07`.
- Minor finding 2: `REQ-MCP-032` Evidence used present-tense wording that conflicted with the updated guide state; changed the evidence sentence to start with `Before this change,`.
- Non-issues confirmed by review:
  - Guide status-gated section wording matches implementation behavior.
  - `tools.md` status matrix matches `statusAllowedForKind` vocabulary.
  - Work item guide now includes `## Evidence` in canonical format.
  - Individual validation for `REQ-MCP-032`, `WORK-MCP-029`, and `TASK-MCP-029-01` was reported clean before close synchronization.

Applied close patch:

- `docs/spec/design-records-mcp/tools.md`: `last_updated: 2026-06-07`.
- `docs/requirements/mcp/REQ-MCP-032-expose-workflow-status-gated-required-sections-in-authoring-guidance.md`: changed stale evidence wording to historical framing.

Close verification:

- Review had no blocking findings.
- Workflow close synchronization updates this task to `done` after applying the minor patch.
