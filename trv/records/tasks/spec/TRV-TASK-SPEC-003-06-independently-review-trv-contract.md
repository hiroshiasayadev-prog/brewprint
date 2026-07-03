# TRV-TASK-SPEC-003-06: Independently review TRV contract

- **id**: TRV-TASK-SPEC-003-06
- **status**: blocked
- **date**: 2026-07-03
- **work_item**: TRV-WORK-SPEC-003
- **task_type**: review
- **estimate**: 0.5d
- **depends_on**:
  - TRV-TASK-SPEC-003-04
  - TRV-TASK-SPEC-003-05
- **outputs**:
  - TRV-TASK-SPEC-003-06

## Goal

Issue one independent integrated verdict for the final TRV application-contract state.

## Work

- Review W003, T01 through T05, routed decisions, ADRs, contract Specifications, `spec:trv`, and closed W002 authority.
- Verify external interface, Task input, result envelopes, caller ownership, and compatibility.
- Verify W004-owned detail and current DRMCP integration remain excluded.
- Return PASS, NEEDS REVISION, NOT READY, or BLOCKED with exact findings.

This Task must not edit reviewed artifacts, repair findings, change the graph, perform closure, or start detailed design or implementation.

## Done condition

- One independent verdict covers the complete W003 state.
- Every material decision has a complete ADR and Specification trace.
- Findings name exact severity, target, outcome, and owner.

## Verification

- Confirm reviewer independence from T04 and T05.
- Confirm reviewed artifacts are stable.
- Confirm no reviewed artifact changed.

## Evidence

- T03 created this review owner before the Specification-placement gap was discovered.
- W003 has no complete contract Specification set to review.
- The user moved the complete contract review boundary to TRV-WORK-SPEC-005.
- This Task remains blocked and must not issue a verdict for the incomplete W003 state.
- Replacement review ownership will be materialized inside W005 after placement is decided.
