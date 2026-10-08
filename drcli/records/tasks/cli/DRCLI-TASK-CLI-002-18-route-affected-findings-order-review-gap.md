# DRCLI-TASK-CLI-002-18: Route affected-findings ordering review gap

- **status**: done
- **date**: 2026-10-08
- **work_item**: DRCLI-WORK-CLI-002
- **task_type**: coordination
- **estimate**: 0.5d
- **depends_on**:
  - DRCLI-TASK-CLI-002-17
- **outputs**:
  - DRCLI-TASK-CLI-002-19
  - DRCLI-TASK-CLI-002-20
  - DRCLI-TASK-CLI-002-11
  - DRCLI-WORK-CLI-002

## Goal

Route the single outstanding T17 finding DRCLI-T17-F-MIN-001 to one bounded contract correction and a separate independent closure review.

## Work

- Preserve T10 and T17 review decisions as historical evidence.
- Materialize one bounded correction Task to define deterministic affected-record group ordering in the DRCLI authoring diagnostic catalog, without changing PRODUCT finding order.
- Materialize one dependent read-only independent finding-closure review.
- Ensure T11 closure synchronization waits for the new independent closure review and all earlier required findings.
- Register the new Tasks and finding-driven dependency route in the parent Work Item.
- Perform no production implementation, Git operation, or Specification edit.

## Done condition

- DRCLI-T17-F-MIN-001 has a bounded correction owner and independent review gate.
- T11 is dependent on the later review, never on T17 alone.
- The parent Work Item reflects the complete finding-driven route and remains open.
- No user-owned new decision is invented.

## Verification

- T17 verdict is NEEDS REVISION and DRCLI-T10-F-MAJ-001 remains OPEN.
- The remaining concern is only order of affected candidate-record groups; 42 error identities/exit mappings already passed.
- Correction and review writer boundaries are disjoint.

## Evidence

- T17 identified exactly one required minor correction, DRCLI-T17-F-MIN-001.
- T17 explicitly confirmed the major identity/exit-category problem was corrected but could not close the original major finding while deterministic affected-record-group order remained unspecified.
- T17 says no new user judgment is needed.
- Created T19 correction and T20 independent closure-review tasks.
- Registered the additional finding-driven graph and T11 dependency; no DRCLI Specification or DRMCP content was changed by this Task.
