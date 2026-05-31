# TASK-DATA-003-04: UC-002 model response helper candidate review

- **id**: TASK-DATA-003-04
- **status**: todo
- **date**: 2026-05-31
- **work_item**: WORK-DATA-003
- **source_requirement**: REQ-DATA-002
- **estimate**: 0.5d-1d
- **depends_on**:
  - TASK-DATA-003-03
- **outputs**:
  - UC-002 model response helper candidate classification
  - Boundary note delegating actual migration to WORK-DATA-004 or later work
  - WORK-DATA-003 close-readiness input

## Goal

Review UC-002 model response helper-shape candidates after model-file render support is available, and classify what should be migrated later without performing the migration in this task.

This task is the WORK-DATA-003 close-readiness review for UC-002 migration scope. Actual YAML / fixture / golden migration is delegated to WORK-DATA-004 or later follow-up work so that WORK-DATA-003 does not mix model-file render exposure with later helper policy work.

## Work

- Review the relevant UC-002 model response helper-shape candidates recorded by prior investigations / requirements.
- Classify candidates into:
  - model-file helper migration candidates
  - later helper policy candidates that must wait for REQ-DATA-003 / WORK-DATA-004
  - candidates that should remain unchanged
- Confirm model-file render support from TASK-DATA-003-03 removes the render-blocker for model-file helper candidates.
- Record why actual migration is delegated out of WORK-DATA-003.
- Update WORK-DATA-003 close-readiness evidence if needed.

## Included Scope

- Review and classification only.
- UC-002 migration boundary clarification.
- WORK-DATA-003 close-readiness input.
- Delegation note for actual migration to WORK-DATA-004 or later follow-up work.

## Excluded Scope

- UC-002 YAML migration.
- Fixture / golden regeneration for UC-002.
- Renderer implementation.
- REQ-DATA-003 implementation.
- WORK-DATA-004 implementation or status correction.
- WORK-DATA-002 reopening.
- Tagged union rendering / ADR-073 implementation.

## Done Condition

- UC-002 model response helper-shape candidates are classified.
- Actual migration is explicitly delegated out of WORK-DATA-003.
- Any dependency on REQ-DATA-003 / WORK-DATA-004 is identified without implementing it.
- WORK-DATA-003 has enough evidence to decide whether it can close after model-file render exposure and candidate classification.
- No UC-002 YAML, fixture / golden, renderer, or later helper policy implementation is performed.

## Verification

- Confirm no UC-002 YAML / render output files were changed by this task.
- Confirm TASK-DATA-003-03 model-file render support is already done.
- Confirm REQ-DATA-003 / WORK-DATA-004 scope is not implemented here.
- Confirm WORK-DATA-002 remains closed.

## Evidence

Pending review.
