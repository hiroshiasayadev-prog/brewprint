# PRODUCT-TASK-SPEC-018-03: Amend decision checkpoint ADR

- **id**: PRODUCT-TASK-SPEC-018-03
- **status**: done
- **date**: 2026-07-01
- **work_item**: PRODUCT-WORK-SPEC-018
- **task_type**: authoring
- **estimate**: 0.25d
- **depends_on**:
  - PRODUCT-TASK-SPEC-018-02
- **outputs**:
  - PRODUCT-TASK-SPEC-018-03
  - PRODUCT-ADR-SPEC-006

## Goal

Clarify `PRODUCT-ADR-SPEC-006` without changing its checkpoint-versus-canonical-state decision.

## Work

- Preserve the accepted ADR status and original decision boundary.
- Remove the stale completed-decision-Task writeback consequence.
- Assign downstream references and Evidence to downstream Tasks.
- Add ADR routing and ADR-boundary partitioning responsibility.
- Preserve ADR routing as separate from ADR authoring.

## Done condition

- `PRODUCT-ADR-SPEC-006` retains its original selected alternative.
- The ADR no longer requires ADR references or downstream progress in the completed decision Task.
- The ADR explicitly covers routing classification, partitioning, and disposition.
- No new ADR, Specification, or skill file is authored by this Task.

## Verification

- Confirm ADR status remains `accepted`.
- Confirm ADR date remains unchanged for the meaning-preserving clarification.
- Confirm `supersedes` remains empty.
- Confirm the canonical ownership table remains intact.
- Confirm only `PRODUCT-ADR-SPEC-006` and this Task are outputs.

## Evidence

- Input route: `PRODUCT-TASK-SPEC-018-02`, B-007.
- Decisions projected: D-019 and D-021 from `PRODUCT-TASK-SPEC-018-01`.
- Removed consequence: completed decision Tasks reference authored ADRs after authoring.
- Added consequence: downstream authoring, Specification, review, and closure Tasks own their own references and Evidence.
- Added routing scope: ADR-boundary partitioning and create, amend, reuse, or supersede disposition.
- Result: clarification amendment complete; no supersession required.
