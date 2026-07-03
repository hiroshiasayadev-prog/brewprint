# TRV-TASK-SPEC-002-16: Coordinate revised architecture finding route

- **id**: TRV-TASK-SPEC-002-16
- **status**: done
- **date**: 2026-07-02
- **work_item**: TRV-WORK-SPEC-002
- **task_type**: coordination
- **estimate**: 0.25d
- **depends_on**:
  - TRV-TASK-SPEC-002-14
- **outputs**:
  - TRV-WORK-SPEC-002
  - TRV-TASK-SPEC-002-15
  - TRV-TASK-SPEC-002-16
  - TRV-TASK-SPEC-002-17
  - TRV-TASK-SPEC-002-18

## Goal

Materialize the correction, independent finding-closure review, and closure-blocking route required by the T14 findings.

## Work

- Consume T14 F-MAJ-01, F-MAJ-02, and F-MIN-01 without changing their required outcomes.
- Classify all three findings as projection defects that require no new user judgment.
- Group the findings into one correction Task because they share one Specification writer, one completion judgment, and one review boundary.
- Create one later independent review Task that decides each finding as `CLOSED` or `OPEN`.
- Block T15 until the finding-closure review completes with every required finding closed.
- Update the W002 Task graph and Task inventory with the exact post-review route.

This Task must not:

- repair a finding;
- change T09 decisions, T10 routing, T14 findings, or reviewed artifacts;
- issue a review verdict;
- synchronize lifecycle or ADR migration;
- design W003 or W004;
- perform implementation, stage, or commit work.

## Done condition

- T17 owns the complete bounded repair for F-MAJ-01, F-MAJ-02, and F-MIN-01.
- T18 owns independent finding closure after T17.
- T15 depends on T18 and remains blocked until every required finding is independently closed.
- W002 records the exact revised route without changing historical completed Tasks.

## Verification

- Confirm every T14 finding has one repair owner and one later independent closure owner.
- Confirm T17 and T18 have distinct Task types, completion judgments, and owners.
- Confirm T15 cannot begin from the T14 `NEEDS REVISION` verdict alone.
- Confirm no reviewed Specification or completed Task changed.

## Evidence

- T14 returned `NEEDS REVISION` with F-MAJ-01, F-MAJ-02, and F-MIN-01.
- Every finding states `New user judgment required: no` and routes to correction.
- T17 was created as one bounded correction owner for the three directly related architecture-Specification defects.
- T18 was created as the independent finding-closure review owner.
- T15 was changed from a direct T14 dependency to a T18 dependency and marked blocked.
- W002 was extended with T16 through T18 and the exact correction route.
- DRMCP authoring is non-operational under the current agent authoring policy. Filesystem fallback was used.
- Result: `PASS`.
