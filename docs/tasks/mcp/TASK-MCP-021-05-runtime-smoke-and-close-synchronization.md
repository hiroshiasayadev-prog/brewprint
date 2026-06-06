# TASK-MCP-021-05: runtime smoke and close synchronization

- **id**: TASK-MCP-021-05
- **status**: done
- **date**: 2026-06-05
- **work_item**: WORK-MCP-021
- **source_requirement**: REQ-MCP-022
- **estimate**: 0.5d
- **depends_on**:
  - TASK-MCP-021-02
  - TASK-MCP-021-03
  - TASK-MCP-021-04
- **outputs**:
  - Runtime smoke evidence for heading-safe named_section_replace normalization
  - Close synchronization updates for tasks, work item, and requirement

## Goal

Verify the completed heading-safe normalization behavior through runtime smoke and synchronize close evidence across the related workflow artifacts.

## Work

- Run runtime smoke against Design Records MCP authoring update behavior.
- Confirm direct `body` replacement with a leading duplicate selected heading produces normalized diff and warning diagnostic.
- Confirm `body_cache_id` replacement follows the same behavior.
- Confirm non-matching and internal headings are not incorrectly stripped.
- Update task evidence and statuses.
- Update work item close evidence and status when completion conditions are met.
- Update REQ-MCP-022 status and close evidence when accepted completion is verified.

## Done condition

- Runtime smoke confirms the required behavior.
- All related tasks have final evidence.
- WORK-MCP-021 is closed only after implementation, tests, docs, and smoke evidence are complete.
- REQ-MCP-022 is synchronized with the completed work item.

## Verification

- Runtime smoke output is recorded.
- `validate_records` reports no workflow metadata errors for the touched artifacts.

## Evidence

Runtime smoke was executed by the user and passed.

Summary:

- MCP server initialized successfully.
- Direct `body` heading stripping:
  - `proposal_created: true`
  - diagnostic category: `section_replacement_body_heading_stripped`
  - severity: `warning`
  - stripped heading: `Evidence`
  - stripped level: `2`
  - `body_cache_id` was created
- `body_cache_id` heading stripping:
  - `proposal_created: true`
  - diagnostic category: `section_replacement_body_heading_stripped`
  - severity: `warning`
  - stripped heading: `Evidence`
  - stripped level: `2`
- Level mismatch:
  - `proposal_created: true`
  - diagnostics: `[]`
  - `level_mismatch_heading_preserved: true`
- All runtime smoke proposals were discarded successfully.
- Final marker: `runtime smoke PASS`.
