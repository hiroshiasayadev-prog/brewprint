# REQ-DATA-006: Request option and response behavior constraints

- **id**: REQ-DATA-006
- **status**: captured
- **date**: 2026-06-01
- **source_refs**:
  - REQ-DATA-002
  - WORK-DATA-009
  - TASK-DATA-009-03
  - TASK-DATA-009-04
- **work_items**:
  - WORK-DATA-013

## Requirement

The project needs an explicit contract for request option and response behavior constraints that are currently preserved as UC-002 notes instead of machine-readable structure or specified behavior.

This requirement owns the `numeric / default behavior` candidate bucket from `TASK-DATA-009-03`: N-011, N-017, N-022, N-024, N-025, and N-028.

## Evidence

`TASK-DATA-009-03` classified this bucket as behavior-oriented debt: numeric ranges, defaults, omitted-value behavior, unknown-value behavior, fallback branches, and cross-response grouping. Existing helper-shape, enum minimum, tagged-union, DAG TypeRef hint, and MCP identity work do not own that contract.

`TASK-DATA-009-04` created this requirement as the successor owner rather than folding behavior constraints into remaining UC-002 cleanup planning.

## Required Outcome

- Decide the contract boundary for numeric ranges, defaults, omitted request values, unknown values, fallback responses, and cross-response behavior grouping.
- Identify the specs, diagnostics, YAML surfaces, and fixture evidence that would need future updates after the contract is accepted.
- Keep the candidate bucket traceable back to `REQ-DATA-002`, `WORK-DATA-009`, `TASK-DATA-009-03`, and `TASK-DATA-009-04`.

## Explicitly Excluded Scope

- Performing parser, renderer, validator, MCP, or other implementation changes in this requirement.
- Performing UC-002 YAML migration in this requirement.
- Regenerating fixtures or golden outputs in this requirement.
- Implementing ADR-073 tagged union support, ADR-074 DAG TypeRef hint support, or ADR-078 / ADR-079 / ADR-080 MCP identity support.
- Reopening M15, WORK-DATA-001, WORK-DATA-002, WORK-DATA-003, WORK-DATA-004, WORK-DATA-005, WORK-DATA-006, WORK-DATA-007, WORK-DATA-008, WORK-DATA-009, or WORK-DATA-010.

## Boundary

This requirement captures the need for behavior constraints. It does not decide the final schema, implementation sequence, or UC-002 migration set; those are owned by `WORK-DATA-013` and later task artifacts.
