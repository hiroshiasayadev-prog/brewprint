# PRODUCT-TASK-SPEC-023-08: Author cancelled lifecycle ADRs

- **id**: PRODUCT-TASK-SPEC-023-08
- **status**: done
- **date**: 2026-07-03
- **work_item**: PRODUCT-WORK-SPEC-023
- **task_type**: authoring
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-023-06
  - PRODUCT-TASK-SPEC-023-07
- **outputs**:
  - PRODUCT-ADR-SPEC-018
  - PRODUCT-ADR-SPEC-005
  - PRODUCT-ADR-SPEC-010
  - PRODUCT-ADR-SPEC-011
  - PRODUCT-ADR-SPEC-014
  - PRODUCT-ADR-SPEC-017
  - PRODUCT-TASK-SPEC-023-08

## Goal

Create and amend the exact ADR set routed by T06 so every durable cancellation decision has historically honest authority.

## Work

- Create PRODUCT-ADR-SPEC-018 from T06 B-001.
- Amend PRODUCT-ADR-SPEC-005 and PRODUCT-ADR-SPEC-010 for T06 B-002.
- Amend PRODUCT-ADR-SPEC-011 and PRODUCT-ADR-SPEC-014 for T06 B-003.
- Amend PRODUCT-ADR-SPEC-017 for T06 B-004.
- Preserve all existing selected architectures and record only non-material amendments where routed.
- Keep current normative Specification detail out of ADR bodies.
- Record exact changed ADRs and verification Evidence.

This Task must not:

- reopen T01, T04, or T06 decisions;
- author Specifications or workflow-support files;
- change the Task graph;
- perform independent review, correction, synchronization, implementation, stage, or commit work.

## Done condition

- PRODUCT-ADR-SPEC-018 exists and expresses B-001.
- Five existing ADR amendments express B-002 through B-004 without concealing reversal.
- ADR metadata and required sections satisfy ADR authoring rules.
- Alternatives, rationale, and consequences remain bounded to routed decisions.
- T09 has exact accepted ADR input.

## Verification

- Confirm only the six routed ADRs and this Task changed.
- Confirm no supersession was introduced.
- Confirm amendment dates change only where decision meaning changed.
- Confirm Specification prose was not duplicated into ADRs.
- Confirm no review, synchronization, implementation, stage, or commit occurred.

## Evidence

- Created accepted PRODUCT-ADR-SPEC-018 for terminal cancellation lifecycle and atomic propagation.
- Amended PRODUCT-ADR-SPEC-005 and PRODUCT-ADR-SPEC-010 for the cancelled-child `work_item_execution` branch.
- Amended PRODUCT-ADR-SPEC-011 and PRODUCT-ADR-SPEC-014 for irreversible cancelled history and new-record resumption.
- Amended PRODUCT-ADR-SPEC-017 so post-Evidence validation applies before `done`, not cancellation.
- All five amendments preserve their selected architecture and record non-material extensions only.
- No ADR supersession was introduced.
- No Specification, workflow-support, Task graph, review, synchronization, implementation, stage, or commit work occurred.
- DRMCP is non-operational. Filesystem authoring was the required fallback.
- Result: `PASS`.
