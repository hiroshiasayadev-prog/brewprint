# TASK-MCP-025-01: Investigate current validation and diagnostic paths for propose_record_create

- **id**: TASK-MCP-025-01
- **status**: not_started
- **date**: 2026-06-07
- **work_item**: WORK-MCP-025
- **source_requirement**: REQ-MCP-024
- **estimate**: 0.5d
- **depends_on**:
- **outputs**:
  - gap analysis: missing-field validation, status value validation, reciprocal follow-up diagnostic paths
  - implementation recommendations for TASK-MCP-025-03

## Goal

Document the current behavior of `propose_record_create` validation paths that are in scope for REQ-MCP-024 and REQ-MCP-028, and identify what changes are needed.

## Work

- Trace `renderWorkflowMetadata` / `renderADRMetadata` to confirm first-missing-field-only behavior.
- Confirm `Diagnostic` struct fields currently available vs. fields required by REQ-MCP-024 (`allowed_values`, `required_fields`, `target_kind`, `repair_suggestion`).
- Identify where status value validation could be inserted for create requests.
- Confirm `report_required_follow_up` mode behavior and where a proposal-level clarity diagnostic would be added.
- Record findings as Evidence and produce implementation recommendations.

## Done condition

- Gap analysis covers all four scope items: batch field validation, `allowed_values`, reciprocal follow-up clarity, `repair_suggestion`.
- Implementation recommendations identify specific functions to change.
- Evidence is recorded in this task.

## Verification

- Cross-check findings against REQ-MCP-024 and REQ-MCP-028 acceptance criteria.
