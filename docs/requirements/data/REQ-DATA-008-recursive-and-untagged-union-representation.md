# REQ-DATA-008: Recursive and untagged-union representation

- **id**: REQ-DATA-008
- **status**: captured
- **date**: 2026-06-01
- **source_refs**:
  - REQ-DATA-002
  - WORK-DATA-009
  - TASK-DATA-009-03
  - TASK-DATA-009-04
- **work_items**:
  - WORK-DATA-015

## Requirement

The project needs explicit representation support for recursive structures and untagged unions that are not covered by the tagged / discriminated union successor scope.

This requirement owns the `recursive / union structure` candidate bucket from `TASK-DATA-009-03`: N-009 and N-044.

## Evidence

`TASK-DATA-009-03` classified this bucket as a type-system expressiveness gap: an untagged union list and recursive `ObjectRef.parent`. `REQ-DATA-004` / `WORK-DATA-010` cover tagged union and discriminator payload support only, so silently expanding that work to recursive or untagged semantics would blur the successor boundary.

`TASK-DATA-009-04` created this requirement as the distinct owner for recursive and untagged-union representation.

## Required Outcome

- Decide whether and how recursive references and untagged union shapes are represented in brewprint data models.
- Decide whether this remains separate from ADR-073 or requires an explicit future ADR-073 successor / broadening decision.
- Identify any future spec, diagnostic, YAML, render, and fixture evidence updates required after the contract is accepted.
- Keep the candidate bucket traceable back to `REQ-DATA-002`, `WORK-DATA-009`, `TASK-DATA-009-03`, and `TASK-DATA-009-04`.

## Explicitly Excluded Scope

- Performing parser, renderer, validator, MCP, or other implementation changes in this requirement.
- Performing UC-002 YAML migration in this requirement.
- Regenerating fixtures or golden outputs in this requirement.
- Implementing ADR-073 tagged union support, ADR-074 DAG TypeRef hint support, or ADR-078 / ADR-079 / ADR-080 MCP identity support.
- Reopening M15, WORK-DATA-001, WORK-DATA-002, WORK-DATA-003, WORK-DATA-004, WORK-DATA-005, WORK-DATA-006, WORK-DATA-007, WORK-DATA-008, WORK-DATA-009, or WORK-DATA-010.

## Boundary

This requirement captures recursive and untagged-union representation needs. It does not decide the final schema, implementation sequence, or UC-002 migration set; those are owned by `WORK-DATA-015` and later task artifacts.
