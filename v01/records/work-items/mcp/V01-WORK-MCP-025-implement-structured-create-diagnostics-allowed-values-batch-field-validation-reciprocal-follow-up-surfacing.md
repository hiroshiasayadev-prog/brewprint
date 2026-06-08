# V01-WORK-MCP-025: Implement structured create diagnostics: allowed_values, batch field validation, reciprocal follow-up surfacing

- **id**: V01-WORK-MCP-025
- **status**: done
- **date**: 2026-06-07
- **source_requirement**: V01-REQ-MCP-024
- **impact_refs**:
  - V01-REQ-MCP-028
  - SPEC-design-records-mcp-tools
- **tasks**:
  - V01-TASK-MCP-025-01
  - V01-TASK-MCP-025-02
  - V01-TASK-MCP-025-03
  - V01-TASK-MCP-025-04

## Goal

Extend `propose_record_create` to surface all detectable input problems in a single proposal response and add structured repair guidance to diagnostic output.

Primary driver: V01-REQ-MCP-024 (structured `allowed_values` and `repair_suggestion` in diagnostics) and V01-REQ-MCP-028 (batch required field validation, reciprocal follow-up surfacing at proposal time, `include_required` mode clarity).

## Boundary

In scope:

- Add `allowed_values`, `required_fields`, `target_kind`, and `repair_suggestion` fields to the `Diagnostic` struct and its JSON serialization.
- Change required field validation in `renderWorkflowMetadata` / `renderADRMetadata` to collect all missing fields in a single pass, not stop on the first missing field.
- Add proposal-time status value validation for create requests, emitting `invalid_metadata_value` with `allowed_values` when the status is not in the allowed set for the kind.
- Add a proposal-level diagnostic when `reciprocal_update_mode: report_required_follow_up` is used, identifying that `include_required` is the safe mode for accept.
- Update `SPEC-design-records-mcp-tools` to document the new diagnostic fields and batch validation behavior.
- Add regression tests for batch validation, `allowed_values` format, and reciprocal follow-up clarity.

Out of scope:

- Multi-record atomic transactions (V01-REQ-MCP-025).
- Automatically accepting repair suggestions.
- Changing canonical status vocabularies (already addressed by V01-ADR-094 / V01-REQ-MCP-026).
- `propose_record_update` diagnostic changes (a separate concern).
- Removing `accept_proposed_write` guard checks.

## Impact Scope

- `internal/designrecords/types.go`: add `AllowedValues`, `RequiredFields`, `TargetKind`, `RepairSuggestion` to `Diagnostic` struct and `MarshalJSON`.
- `internal/designrecords/authoring.go`: batch field validation in `renderWorkflowMetadata`/`renderADRMetadata`; status value validation; reciprocal follow-up clarity diagnostic.
- `internal/designrecords/authoring_test.go`: regression tests for new behavior.
- `internal/designrecordsmcp/tools_call_test.go`: MCP boundary tests for new diagnostic JSON shapes.
- `docs/spec/design-records-mcp/tools.md`: contract update for new diagnostic fields.

## Task flow

V01-TASK-MCP-025-01 then V01-TASK-MCP-025-02 then V01-TASK-MCP-025-03 then V01-TASK-MCP-025-04

## Task Candidates

- `V01-TASK-MCP-025-01`: Investigate current validation and diagnostic paths for `propose_record_create`.
- `V01-TASK-MCP-025-02`: Update `SPEC-design-records-mcp-tools` for structured diagnostics and batch validation contract.
- `V01-TASK-MCP-025-03`: Implement structured diagnostics, batch field validation, and regression tests.
- `V01-TASK-MCP-025-04`: Runtime smoke and close synchronization.

## Completion Condition

- `propose_record_create` with multiple missing fields returns all missing field names in a single diagnostic.
- `invalid_metadata_value` diagnostics for status include `allowed_values`.
- `missing_required_metadata` batch diagnostic includes `required_fields` and `target_kind`.
- `report_required_follow_up` mode proposals include a proposal-level diagnostic identifying `include_required` as the safe accept mode.
- `repair_suggestion` is included in deterministic repair cases.
- All regression tests pass.
- Runtime smoke confirms new behavior end-to-end.
- V01-REQ-MCP-024, V01-REQ-MCP-028, this work item, and all tasks are synchronized to final statuses.

## Evidence

Verdict: PASS. All completion conditions met.

Files changed:

- `internal/designrecords/types.go`: added `AllowedValues`, `RequiredFields`, `TargetKind`, `RepairSuggestion` fields to `Diagnostic` struct and `diagnosticJSON`; updated `MarshalJSON`.
- `internal/designrecords/authoring.go`: added helper functions `requiredCreateFieldNames`, `allowedStatusValuesForKind`, `validateCreateFieldsBatch`, `validateCreateStatusForCreate`, `reciprocalFollowUpModeRequiredDiagnostic`; changed `requiredReciprocalUpdates` to return 4 values (added `[]Diagnostic`); updated `prepareCreate` to call batch field validation and status validation before rendering and to handle 4-value return.
- `internal/designrecords/authoring_test.go`: fixed test fixtures from `implementation_pending`/`todo` to `not_started`; added `WORK_missing_multiple_required_fields_batch_diagnostic` and `WORK_single_missing_required_field_batch_diagnostic` subtests; added `TestProposeRecordCreateStatusDiagnostic` and `TestProposeRecordCreateReciprocalFollowUpMode`.
- `docs/spec/design-records-mcp/tools.md`: added `missing_required_metadata_batch`, `reciprocal_follow_up_mode_required`, `no_op_update` diagnostic categories; added authoring diagnostic additional fields table; added JSON examples.
- `tmp.py`: added smoke cases [4] (batch missing-field) and [5] (invalid status / allowed_values).

`go test ./...`: passed (all packages). Runtime smoke: all 5 cases PASS.
