# V01-TASK-MCP-025-02: Update SPEC-design-records-mcp-tools for structured diagnostics and batch validation contract

- **id**: V01-TASK-MCP-025-02
- **status**: not_started
- **date**: 2026-06-07
- **work_item**: V01-WORK-MCP-025
- **source_requirement**: V01-REQ-MCP-024
- **estimate**: 0.5d
- **depends_on**:
  - V01-TASK-MCP-025-01
- **outputs**:
  - SPEC-design-records-mcp-tools update: new diagnostic fields, batch missing-field diagnostic, reciprocal follow-up clarity diagnostic

## Goal

Update `SPEC-design-records-mcp-tools` to document the new diagnostic fields and batch validation behavior before implementation.

## Work

- Add `allowed_values`, `required_fields`, `target_kind`, and `repair_suggestion` to the diagnostic shape documentation.
- Document `missing_required_metadata_batch` (or equivalent) as the new category for batch missing field failures.
- Document `reciprocal_follow_up_mode_required` (or equivalent) as the new proposal-level diagnostic emitted in `report_required_follow_up` mode.
- Document `repair_suggestion` format and advisory-only semantics.
- Keep changes consistent with existing diagnostic section format in the spec.

## Done condition

- Spec documents all new diagnostic fields with JSON examples.
- Batch missing-field diagnostic category and shape are specified.
- Reciprocal follow-up clarity diagnostic is specified.
- `repair_suggestion` is documented as advisory.

## Verification

- Confirm spec changes are consistent with V01-REQ-MCP-024 and V01-REQ-MCP-028 acceptance criteria.
