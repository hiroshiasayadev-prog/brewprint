# WORK-MCP-029: Reflect workflow status-gated required section rules in authoring guides

- **id**: WORK-MCP-029
- **status**: done
- **date**: 2026-06-07
- **source_requirement**: REQ-MCP-032
- **impact_refs**:
  - SPEC-design-records-mcp-tools
- **tasks**:
  - TASK-MCP-029-01

## Goal

Make the authoring guides and MCP tools spec visible enough that agents can know valid workflow statuses and status-gated required narrative sections before attempting status transitions.

## Boundary

In scope:

- Update authoring guidance for requirement, work item, and task records.
- Add or expose a kind-by-kind valid status matrix in the MCP tools spec.
- Keep the existing validation semantics unchanged.
- Prepare the work for Codex review and close the workflow artifacts after review passes.

Out of scope:

- Changing runtime validation behavior.
- Adding new authoring transaction operations.
- Reworking the full authoring guide structure.

## Impact Scope

- `docs/guides/requirement-authoring.md`
- `docs/guides/work-item-authoring.md`
- `docs/guides/task-authoring.md`
- `docs/spec/design-records-mcp/tools.md`
- Related workflow artifacts for this requirement/work item/task closure.

## Task flow

1. Update guidance and spec visibility.
2. Run documentation review with Codex.
3. Apply any review fixes.
4. Close task, work item, and requirement when evidence is available.

## Task Candidates

- Update authoring guides and tools spec.
- Review the documentation consistency.
- Close synchronization.

## Completion Condition

- Authoring guides clearly state the status-gated required section rules for requirement, work item, and task artifacts.
- Work item guide includes `## Evidence` in the canonical format.
- MCP tools spec exposes kind-by-kind valid statuses, or an equivalent visible contract.
- Review passes with no blocking findings.

## Evidence

Implementation, review, and close synchronization completed.

Implemented documentation visibility updates:

- Requirement authoring guidance states that `status: accepted` requires non-empty `Requirement` and `Required Outcome` sections.
- Work item authoring guidance includes `## Evidence` in the canonical format and states that `status: done` requires non-empty `Goal`, `Boundary`, and `Evidence` sections.
- Task authoring guidance states that `status: done` requires non-empty `Goal`, `Work`, `Done condition`, `Verification`, and `Evidence` sections.
- `docs/spec/design-records-mcp/tools.md` exposes the kind-by-kind workflow status matrix.

Review result:

- Verdict: PASS_WITH_MINOR.
- Blocking findings: none.
- Minor finding 1: `tools.md` `last_updated` was stale; fixed to `2026-06-07`.
- Minor finding 2: `REQ-MCP-032` Evidence wording read as current-state evidence after the guide update; changed to historical wording with `Before this change,`.
- Non-issues confirmed:
  - Guide status-gated section wording matches implementation behavior.
  - The status matrix matches runtime status vocabulary.
  - Work item guide canonical format includes `## Evidence` and aligns with required narrative section policy.

Close verification:

- `TASK-MCP-029-01` updated to `done` with review and patch evidence.
- `WORK-MCP-029` updated to `done` with implementation, review, and patch evidence.
- Review passed with no blocking findings after minor patch application.
