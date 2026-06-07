# REQ-MCP-030: authoring create includes required parent reciprocal updates by default

- **id**: REQ-MCP-030
- **status**: accepted
- **date**: 2026-06-07
- **source_refs**:
  - REQ-MCP-028
  - WORK-MCP-026
- **work_items**:
  - WORK-MCP-027

## Requirement

`propose_record_create` for workflow child records, especially task create with `parent_id` / `fields.work_item` pointing to an existing work item, should include required parent reciprocal updates in the proposal diff by default.

For task creation, when the parent work item can be resolved unambiguously, the proposal should include both:

- the new task file create diff
- the parent work item `tasks` append diff

This should prevent the current wasteful flow where a proposal is first created in report-only mode, accept fails with `required_follow_up_not_satisfied`, the proposal is discarded, and the same create is proposed again with `reciprocal_update_mode: include_required`.

## Evidence

Observed authoring flow for `TASK-MCP-026-01..05` showed that report-only reciprocal handling produced proposals whose validation/accept path required follow-up updates. The agent then discarded all proposals and re-proposed them with `include_required`, even though `parent_id` and `fields.work_item` already made the required parent update mechanically knowable.

The desired behavior is not immediate repository mutation. The parent update must be included only as part of the retained proposal diff and applied atomically by `accept_proposed_write`.

## Required Outcome

- By default, task create with `parent_id` / `fields.work_item` pointing to an existing work item includes the parent work item `tasks` append in the same proposal diff.
- The automatically included reciprocal update is reported as an `info` diagnostic, for example `reciprocal_update_included`.
- Required follow-up updates already included in the proposal diff are not also emitted as unsatisfied `required_follow_up_updates`.
- Explicit `reciprocal_update_mode: report_required_follow_up` remains available and preserves report-only behavior.
- `accept_proposed_write` applies the child create and parent reciprocal update atomically.
- Unsafe or ambiguous cases remain blocking or report-only rather than being silently repaired.

## Explicitly Excluded Scope

- Do not automatically repair mismatched `parent_id` and `fields.work_item`.
- Do not create missing parent records.
- Do not append to non-work-item parents.
- Do not duplicate a task ID already present in the parent work item `tasks` list.
- Do not change repository files during `propose_record_create`; only the retained proposal diff changes.

## Boundary

This requirement concerns Design Records MCP authoring transaction behavior for create proposals and reciprocal workflow relations. It does not define general workflow validation semantics outside authoring proposals, and it does not replace explicit relation validation.
