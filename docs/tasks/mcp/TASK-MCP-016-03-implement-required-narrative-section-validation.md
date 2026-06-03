# TASK-MCP-016-03: Implement required narrative section validation

- **id**: TASK-MCP-016-03
- **status**: done
- **date**: 2026-06-03
- **work_item**: WORK-MCP-016
- **source_requirement**: REQ-MCP-017
- **estimate**: 1d-2d
- **depends_on**:
  - TASK-MCP-016-01
  - TASK-MCP-016-02
- **outputs**:
  - Implementation and regression tests for close-state required narrative section validation

## Goal

Implement required narrative section validation for workflow artifacts in close states.

## Work

- Add validation logic for missing or whitespace-only required section bodies.
- Apply the policy from `TASK-MCP-016-01`.
- Emit diagnostics matching `TASK-MCP-016-02`.
- Add regression tests for `WORK done`, `TASK done`, and `REQ accepted` cases.
- Keep non-close states limited to existing metadata and relation validation unless the policy task says otherwise.

## Done condition

This task is done when implementation and regression tests cover the required section validation behavior.

## Verification

- Run targeted Go tests for design record validation and MCP tool behavior.
- Confirm diagnostic fields match the documented contract.

## Evidence

Completed on 2026-06-03.

Files changed:

- `internal/designrecords/types.go`: added `DiagnosticMissingRequiredSection` and `DiagnosticEmptyRequiredSection` constants; added `Section string` and `Status string` fields to `Diagnostic`; updated `MarshalJSON` to include those fields.
- `internal/designrecords/validation.go`: added `requiredSectionPolicyFor`, `requiredNarrativeSectionDiagnostics`, `findHeadingLevel`, and `extractSectionBody` helpers; wired `requiredNarrativeSectionDiagnostics` into `recordDiagnostics`.
- `internal/designrecords/validation_test.go`: updated existing fixtures in `TestValidateRecordsOKWhenNoErrorDiagnostics` and `writeWorkflowHappyPathFixture` to include required narrative sections for `REQ accepted`; added `TestValidateRecordsRequiredNarrativeSections` covering all policy matrix cases and diagnostic field assertions; added `assertRequiredSectionDiagnostic` and `assertNoRequiredSectionDiagnosticForRecord` helpers.
- `internal/designrecordsmcp/tools_call_test.go`: updated `toolsCallTestIndex()` fixtures for `REQ-MCP-003` (accepted) and `TASK-MCP-003-01` (done) to include required narrative sections and headings.

Diagnostic categories implemented:

- `missing_required_section` (severity: error) — required narrative section heading absent for gated status.
- `empty_required_section` (severity: error) — required narrative section heading present but body empty or whitespace-only.

Both diagnostics include required fields: `category`, `severity`, `record_id`, `path`, `message`, `section`, `status`.

Policy matrix implemented:

| artifact kind | gated status | required non-empty narrative sections |
|---|---|---|
| `work_item` | `done` | `Goal`, `Boundary`, `Evidence` |
| `task` | `done` | `Goal`, `Work`, `Done condition`, `Verification`, `Evidence` |
| `requirement` | `accepted` | `Requirement`, `Required Outcome` |

`Evidence`, `Boundary`, and `Explicitly Excluded Scope` are NOT required for `REQ accepted`.

Test cases added/updated:

- `WORK status: done` — missing Goal, empty Goal, empty Boundary, empty Evidence, valid with all three non-empty.
- `TASK status: done` — each of Goal / Work / Done condition / Verification / Evidence individually enforced.
- `REQ status: accepted` — Requirement required, Required Outcome required, Evidence not required, Boundary not required, Explicitly Excluded Scope not required.
- Non-gated statuses — `WORK not_started`, `TASK todo`, `REQ captured` do not emit section diagnostics.
- Diagnostic field assertions — `missing_required_section` and `empty_required_section` include `record_id`, `path`, `section`, `status`, `severity`, `message`.

Commands run:

```
go test ./internal/designrecords ./internal/designrecordsmcp
```

Result: both packages pass (2.001s, 0.526s).

```
go test ./...
```

Result: all packages pass. No regressions.

Runtime smoke and final close synchronization are intentionally deferred to `TASK-MCP-016-04`.
