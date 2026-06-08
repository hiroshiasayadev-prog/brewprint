# V01-TASK-MCP-016-02: Update validation diagnostic contract for required narrative sections

- **id**: V01-TASK-MCP-016-02
- **status**: done
- **date**: 2026-06-03
- **work_item**: V01-WORK-MCP-016
- **source_requirement**: V01-REQ-MCP-017
- **estimate**: 0.5d-1d
- **depends_on**:
  - V01-TASK-MCP-016-01
- **outputs**:
  - Spec and diagnostic contract update for empty required narrative section validation

## Goal

Update the public validation/diagnostic contract for close-state required narrative section checks.

## Work

- Update the Design Records MCP validation tool contract to describe status-sensitive required narrative section validation.
- Define the diagnostic category, severity, and payload fields for an empty required narrative section.
- Confirm whether the diagnostic catalog in `docs/spec/diagnostics.md` needs an entry or whether the MCP tools spec fully owns this diagnostic.
- Keep the contract limited to non-empty body validation, not narrative quality assessment.

## Done condition

This task is done when the spec/diagnostic contract is explicit enough for implementation and tests to target.

## Verification

- Review the updated spec text against the policy matrix from `V01-TASK-MCP-016-01`.
- Confirm diagnostics include record ID, section name, status, severity, and actionable message.

## Evidence
Completed on 2026-06-03.

Spec updates:

- Updated `SPEC-design-records-mcp-tools` / `docs/spec/design-records-mcp/tools.md` `validate_records` contract.
- Updated `SPEC-design-records-mcp-schema` / `docs/spec/design-records-mcp/schema.md` diagnostic category contract.
- Updated both spec files' `last_updated` front matter to `2026-06-03`.

Contract decisions recorded:

- `validate_records` now includes status-gated required narrative section validation for workflow artifacts.
- The validation is limited to required heading presence and section body non-empty checks.
- Narrative quality, sufficiency, and semantic correctness are explicitly not judged.
- Required section body is non-empty only when the heading line is excluded, surrounding whitespace is trimmed, and at least one non-whitespace character remains.

Diagnostic categories:

- `missing_required_section`: required narrative section heading is absent for the gated status.
- `empty_required_section`: required narrative section heading exists, but body is empty or whitespace-only for the gated status.
- Both categories have `error` severity.

Required diagnostic fields:

- Standard diagnostic fields remain `category`, `severity`, `record_id`, `path`, and `message` where applicable.
- Required section diagnostics additionally require `section` and `status`.

Policy matrix carried forward from `V01-TASK-MCP-016-01`:

| artifact kind | gated status | required non-empty narrative sections |
|---|---|---|
| `work_item` | `done` | `Goal`, `Boundary`, `Evidence` |
| `task` | `done` | `Goal`, `Work`, `Done condition`, `Verification`, `Evidence` |
| `requirement` | `accepted` | `Requirement`, `Required Outcome` |

REQ accepted clarification:

- `requirement` `accepted` is an adoption-readiness gate, not a close/completion state.
- `Evidence`, `Boundary`, and `Explicitly Excluded Scope` are intentionally not required non-empty for `REQ accepted`.

Scope correction:

- Confirmed `docs/spec/design-records-mcp/schema.md` is the Design Records MCP diagnostic catalog target.
- Did not update `docs/spec/diagnostics.md`, because it belongs to the brewprint DSL semantic diagnostics catalog.
