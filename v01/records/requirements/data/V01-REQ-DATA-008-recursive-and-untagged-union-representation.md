# V01-REQ-DATA-008: Recursive and untagged-union representation

- **id**: V01-REQ-DATA-008
- **status**: accepted
- **date**: 2026-06-01
- **source_refs**:
  - V01-REQ-DATA-002
  - V01-WORK-DATA-009
  - V01-TASK-DATA-009-03
  - V01-TASK-DATA-009-04
- **work_items**:
  - V01-WORK-DATA-015

## Requirement

The project needs explicit representation support for recursive structures and untagged unions that are not covered by the tagged / discriminated union successor scope.

This requirement owns the `recursive / union structure` candidate bucket from `V01-TASK-DATA-009-03`: N-009 and N-044.

## Evidence

`V01-TASK-DATA-009-03` classified this bucket as a type-system expressiveness gap: an untagged union list and recursive `ObjectRef.parent`. `V01-REQ-DATA-004` / `V01-WORK-DATA-010` cover tagged union and discriminator payload support only, so silently expanding that work to recursive or untagged semantics would blur the successor boundary.

`V01-TASK-DATA-009-04` created this requirement as the distinct owner for recursive and untagged-union representation.

Accepted resolution on 2026-06-08:

- Recursive structures are represented through named model TypeRef references only.
- Inline recursive shapes are not introduced.
- Untagged union / general `oneOf` / `anyOf` / scalar union support is not introduced.
- Untagged-like machine-readable surfaces use tagged union envelope models.
- UC-002 N-044 was migrated to `object_ref.parent: object_ref`.
- UC-002 N-009 was migrated to `diagnostic.related: list<diagnostic_related>` using a tagged union envelope.

## Required Outcome

- Decide whether and how recursive references and untagged union shapes are represented in brewprint data models.
- Decide whether this remains separate from V01-ADR-073 or requires an explicit future V01-ADR-073 successor / broadening decision.
- Identify any future spec, diagnostic, YAML, render, and fixture evidence updates required after the contract is accepted.
- Keep the candidate bucket traceable back to `V01-REQ-DATA-002`, `V01-WORK-DATA-009`, `V01-TASK-DATA-009-03`, and `V01-TASK-DATA-009-04`.

## Explicitly Excluded Scope

- Performing parser, renderer, validator, MCP, or other implementation changes in this requirement.
- Performing UC-002 YAML migration in this requirement.
- Regenerating fixtures or golden outputs in this requirement.
- Implementing V01-ADR-073 tagged union support, V01-ADR-074 DAG TypeRef hint support, or V01-ADR-078 / V01-ADR-079 / V01-ADR-080 MCP identity support.
- Reopening M15, V01-WORK-DATA-001, V01-WORK-DATA-002, V01-WORK-DATA-003, V01-WORK-DATA-004, V01-WORK-DATA-005, V01-WORK-DATA-006, V01-WORK-DATA-007, V01-WORK-DATA-008, V01-WORK-DATA-009, or V01-WORK-DATA-010.

## Boundary

This requirement is accepted.

Final boundary:

- Recursive named model references are supported.
- Inline recursive shapes remain unsupported.
- Untagged union / general `oneOf` / `anyOf` / scalar union support remains unsupported.
- Tagged union envelope models are the selected replacement pattern for untagged-like machine-readable surfaces.

The accepted schema and UC-002 cleanup are owned by `V01-WORK-DATA-015` and its completed tasks.
