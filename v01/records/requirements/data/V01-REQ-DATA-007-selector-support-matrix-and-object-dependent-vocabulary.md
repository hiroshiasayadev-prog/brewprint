# V01-REQ-DATA-007: Selector support matrix and object-dependent vocabulary

- **id**: V01-REQ-DATA-007
- **status**: captured
- **date**: 2026-06-01
- **source_refs**:
  - V01-REQ-DATA-002
  - V01-WORK-DATA-009
  - V01-TASK-DATA-009-03
  - V01-TASK-DATA-009-04
- **work_items**:
  - V01-WORK-DATA-014

## Requirement

The project needs an explicit selector support matrix and object-dependent vocabulary contract so UC-002 selector behavior is not left as broad prose notes.

This requirement owns the `selector matrix / support matrix` candidate bucket from `V01-TASK-DATA-009-03`: N-020, N-031, N-037, N-040, and N-042.

## Evidence

`V01-TASK-DATA-009-03` classified this bucket as needing a support-matrix contract before YAML cleanup is meaningful. The bucket is adjacent to selector and ObjectRef behavior and is not owned by helper-shape migration, tagged-union support, DAG TypeRef hint support, or MCP identity work.

`V01-TASK-DATA-009-04` created this requirement as a separate successor so selector support rules do not remain only in close notes.

## Required Outcome

- Decide how selector support matrices and object-dependent kind vocabulary should be represented and validated.
- Identify any future spec, diagnostic, YAML, and fixture evidence updates required after the contract is accepted.
- Keep the candidate bucket traceable back to `V01-REQ-DATA-002`, `V01-WORK-DATA-009`, `V01-TASK-DATA-009-03`, and `V01-TASK-DATA-009-04`.

## Explicitly Excluded Scope

- Performing parser, renderer, validator, MCP, or other implementation changes in this requirement.
- Performing UC-002 YAML migration in this requirement.
- Regenerating fixtures or golden outputs in this requirement.
- Implementing V01-ADR-073 tagged union support, V01-ADR-074 DAG TypeRef hint support, or V01-ADR-078 / V01-ADR-079 / V01-ADR-080 MCP identity support.
- Reopening M15, V01-WORK-DATA-001, V01-WORK-DATA-002, V01-WORK-DATA-003, V01-WORK-DATA-004, V01-WORK-DATA-005, V01-WORK-DATA-006, V01-WORK-DATA-007, V01-WORK-DATA-008, V01-WORK-DATA-009, or V01-WORK-DATA-010.

## Boundary

This requirement captures the need for selector support matrix and object-dependent vocabulary handling. It does not decide the final schema, implementation sequence, or UC-002 migration set; those are owned by `V01-WORK-DATA-014` and later task artifacts.
