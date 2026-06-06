# WORK-MCP-025: Implement structured create diagnostics: allowed_values, batch field validation, reciprocal follow-up surfacing

- **id**: WORK-MCP-025
- **status**: not_started
- **date**: 2026-06-07
- **source_requirement**: REQ-MCP-024
- **impact_refs**:
  - REQ-MCP-028
  - SPEC-design-records-mcp-tools
- **tasks**:
  - TASK-MCP-025-01
  - TASK-MCP-025-02
  - TASK-MCP-025-03
  - TASK-MCP-025-04

## Goal

Extend `propose_record_create` to surface all detectable input problems in a single proposal response and add structured repair guidance to diagnostic output.

Primary driver: REQ-MCP-024 (structured `allowed_values` and `repair_suggestion` in diagnostics) and REQ-MCP-028 (batch required field validation, reciprocal follow-up surfacing at proposal time, `include_required` mode clarity).

## Boundary

In scope:

- Add `allowed_values`, `required_fields`, `target_kind`, and `repair_suggestion` fields to the `Diagnostic` struct and its JSON serialization.
- Change required field validation in `renderWorkflowMetadata` / `renderADRMetadata` to collect all missing fields in a single pass, not stop on the first missing field.
- Add proposal-time status value validation for create requests, emitting `invalid_metadata_value` with `allowed_values` when the status is not in the allowed set for the kind.
- Add a proposal-level diagnostic when `reciprocal_update_mode: report_required_follow_up` is used, identifying that `include_required` is the safe mode for accept.
- Update `SPEC-design-records-mcp-tools` to document the new diagnostic fields and batch validation behavior.
- Add regression tests for batch validation, `allowed_values` format, and reciprocal follow-up clarity.

Out of scope:

- Multi-record atomic transactions (REQ-MCP-025).
- Automatically accepting repair suggestions.
- Changing canonical status vocabularies (already addressed by ADR-094 / REQ-MCP-026).
- `propose_record_update` diagnostic changes (a separate concern).
- Removing `accept_proposed_write` guard checks.

## Impact Scope

- `internal/designrecords/types.go`: add `AllowedValues`, `RequiredFields`, `TargetKind`, `RepairSuggestion` to `Diagnostic` struct and `MarshalJSON`.
- `internal/designrecords/authoring.go`: batch field validation in `renderWorkflowMetadata`/`renderADRMetadata`; status value validation; reciprocal follow-up clarity diagnostic.
- `internal/designrecords/authoring_test.go`: regression tests for new behavior.
- `internal/designrecordsmcp/tools_call_test.go`: MCP boundary tests for new diagnostic JSON shapes.
- `docs/spec/design-records-mcp/tools.md`: contract update for new diagnostic fields.

## Task flow

TASK-MCP-025-01 then TASK-MCP-025-02 then TASK-MCP-025-03 then TASK-MCP-025-04

## Task Candidates

- `TASK-MCP-025-01`: Investigate current validation and diagnostic paths for `propose_record_create`.
- `TASK-MCP-025-02`: Update `SPEC-design-records-mcp-tools` for structured diagnostics and batch validation contract.
- `TASK-MCP-025-03`: Implement structured diagnostics, batch field validation, and regression tests.
- `TASK-MCP-025-04`: Runtime smoke and close synchronization.

## Completion Condition

- `propose_record_create` with multiple missing fields returns all missing field names in a single diagnostic.
- `invalid_metadata_value` diagnostics for status include `allowed_values`.
- `missing_required_metadata` batch diagnostic includes `required_fields` and `target_kind`.
- `report_required_follow_up` mode proposals include a proposal-level diagnostic identifying `include_required` as the safe accept mode.
- `repair_suggestion` is included in deterministic repair cases.
- All regression tests pass.
- Runtime smoke confirms new behavior end-to-end.
- REQ-MCP-024, REQ-MCP-028, this work item, and all tasks are synchronized to final statuses.
