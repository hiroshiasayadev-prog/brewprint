# TASK-MCP-025-03: Implement structured diagnostics, batch field validation, and regression tests

- **id**: TASK-MCP-025-03
- **status**: not_started
- **date**: 2026-06-07
- **work_item**: WORK-MCP-025
- **source_requirement**: REQ-MCP-024
- **estimate**: 1.5d
- **depends_on**:
  - TASK-MCP-025-01
  - TASK-MCP-025-02
- **outputs**:
  - implementation patch
  - regression tests for batch validation
  - regression tests for allowed_values diagnostic
  - regression tests for reciprocal follow-up clarity

## Goal

Implement the structured create diagnostic behavior required by REQ-MCP-024 and REQ-MCP-028.

## Work

- Add `AllowedValues []string`, `RequiredFields []string`, `TargetKind string`, `RepairSuggestion map[string]any` to `Diagnostic` struct and `MarshalJSON` in `types.go`.
- Change `renderWorkflowMetadata` and `renderADRMetadata` to collect all missing fields in a single pass instead of returning on first missing field.
- Add a `validateCreateFields` pre-render step in `prepareCreate` that collects all missing fields and returns a single structured `missing_required_metadata_batch` diagnostic (or equivalent) with `required_fields` and `target_kind`.
- Add proposal-time status value validation for create requests: check the `status` field in `fields` against the allowed set for the kind, emit `invalid_metadata_value` with `allowed_values` when invalid.
- In `requiredReciprocalUpdates` / `prepareCreate`, add a proposal-level diagnostic when `reciprocal_update_mode` is `report_required_follow_up`, identifying that `include_required` is the safe accept mode.
- Add regression tests:
  - `propose_record_create` with multiple missing fields → single diagnostic listing all missing fields.
  - `propose_record_create` with invalid status → `invalid_metadata_value` with `allowed_values`.
  - `propose_record_create` in `report_required_follow_up` mode → proposal-level clarity diagnostic.
  - Existing single-missing-field tests still pass.
  - MCP boundary test for new diagnostic JSON shapes.

## Done condition

- All new fields serialize correctly to JSON (no empty arrays emitted for nil).
- Batch missing-field diagnostic names all missing fields in one response.
- `allowed_values` is populated for `invalid_metadata_value` status failures on create.
- `report_required_follow_up` mode emits clarity diagnostic.
- `go test ./internal/designrecords ./internal/designrecordsmcp` passes.
- `go test ./...` passes.

## Verification

- Run targeted tests, then full package tests, then `go test ./...`.
- Confirm no regression in existing create/update/validation paths.
