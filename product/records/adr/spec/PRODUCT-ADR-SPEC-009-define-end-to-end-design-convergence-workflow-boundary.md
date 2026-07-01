# PRODUCT-ADR-SPEC-009: Define the end-to-end design convergence workflow boundary

- **status**: accepted
- **date**: 2026-07-01
- **depends_on**:
  - PRODUCT-ADR-SPEC-004
  - PRODUCT-ADR-SPEC-005
  - PRODUCT-ADR-SPEC-006
- **supersedes**: []
- **migrated_to_spec**: null

## Context

The current design-decision workflow begins with decision inventory and interactive questioning.
It does not own the earlier design-topic intake boundary.
It also leaves completion ambiguous after ADR and Specification authoring.

A complete workflow needs one declared start and one declared terminal state.
Without that boundary, investigation, graph repair, review, and closure can become optional session conventions.

## Decision

Use one design-convergence workflow from initial design-topic input through reviewed design closure.

The workflow starts immediately after a design topic is raised.
The workflow starts before decision inventory and interactive questioning.

The workflow completes only after:

- required decisions are terminal for the Work Item;
- required ADR and canonical authoring is complete;
- one integrated independent review passes;
- required findings are corrected and independently closed;
- originating artifacts, execution graph, lifecycle, Evidence, and relations are synchronized.

Production implementation is outside this workflow.
Implementation planning begins only after design closure.

## Rationale

One end-to-end boundary prevents design work from existing only in chat or ad hoc session state.

An early start captures scope, authority, and investigation before questions are asked.
A reviewed terminal state prevents authoring completion from being mistaken for accepted design closure.

Keeping production implementation outside the workflow preserves the design-versus-execution responsibility boundary.

## Rejected alternatives

| alternative | rejection reason |
|---|---|
| Start only when interactive questions begin. | Topic scoping, authority reading, and impact discovery would remain outside the governed workflow. |
| Complete after ADR and Specification authoring. | The design could close without independent review or synchronized lifecycle Evidence. |
| Include production implementation. | Design convergence and implementation execution have different outcomes, owners, and completion judgments. |
| Use separate unrelated workflows for intake and closure. | Split authority would create gaps and duplicated completion rules. |

## Consequences

- `skills/design-convergence-workflow/` must define the complete entry-to-closure sequence.
- `prompt_chappy.md` must activate the successor after replacement readiness.
- Each design Work Item must define its convergence graph and review boundary.
- Implementation Tasks remain downstream and separate.
- The old workflow skill may be removed only after the successor is ready.

## Evidence

- `PRODUCT-TASK-SPEC-018-01`: D-001, D-004, and D-005.
- `PRODUCT-TASK-SPEC-018-02`: B-001 routing boundary.
- `PRODUCT-WORK-SPEC-018`: successor workflow completion boundary.
