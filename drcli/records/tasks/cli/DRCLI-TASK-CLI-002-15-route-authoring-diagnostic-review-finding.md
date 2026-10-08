# DRCLI-TASK-CLI-002-15: Route authoring diagnostic review finding

- **status**: done
- **date**: 2026-10-08
- **work_item**: DRCLI-WORK-CLI-002
- **task_type**: coordination
- **estimate**: 0.5d
- **depends_on**:
  - DRCLI-TASK-CLI-002-10
- **outputs**:
  - DRCLI-TASK-CLI-002-16
  - DRCLI-TASK-CLI-002-17
  - DRCLI-TASK-CLI-002-11
  - DRCLI-WORK-CLI-002

## Goal

Route T10 finding DRCLI-T10-F-MAJ-001 into one bounded correction and one independent finding-closure review.

## Work

- Preserve T10 as completed review Evidence.
- Create one correction Task for the authoring machine-readable diagnostic gap.
- Create one dependent independent closure-review Task.
- Make T11 closure synchronization depend on successful finding closure rather than the initial T10 result alone.
- Do not create new decision work because T10 explicitly states no new user judgment is required.

## Done condition

- DRCLI-T10-F-MAJ-001 has one exact correction owner.
- A later independent reviewer owns finding closure.
- T11 cannot run before closure review.
- No Specification content is changed by this coordination Task.

## Verification

- Confirm the correction boundary covers only the affected authoring diagnostics/output surface.
- Confirm the correction cannot alter accepted direct-write, retry-cache, proposal, validation, or write-guard semantics.
- Confirm the closure reviewer is read-only.

## Evidence

- T10 completed with verdict NEEDS REVISION and exactly one major finding.
- DRCLI-T10-F-MAJ-001 requires no new user judgment and explicitly routes to correction.
- T16 and T17 are the minimal finding-driven route.
