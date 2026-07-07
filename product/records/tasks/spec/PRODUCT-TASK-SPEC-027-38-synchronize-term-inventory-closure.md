# PRODUCT-TASK-SPEC-027-38: Synchronize term-inventory closure

- **id**: PRODUCT-TASK-SPEC-027-38
- **status**: done
- **date**: 2026-07-04
- **work_item**: PRODUCT-WORK-SPEC-027
- **task_type**: synchronization
- **estimate**: 0.25d
- **depends_on**:
  - PRODUCT-TASK-SPEC-027-37
- **outputs**:
  - PRODUCT-REQ-SPEC-013
  - PRODUCT-WORK-SPEC-027

## Goal

Synchronize Requirement and Work Item lifecycle, closure Evidence, and graph state from the concluded term-inventory Investigation.

## Work

- Confirm PRODUCT-TASK-SPEC-027-37 is `done` and PRODUCT-INV-SPEC-011 is `concluded`.
- Synchronize PRODUCT-REQ-SPEC-013 lifecycle and completion Evidence from the accepted Investigation result.
- Synchronize PRODUCT-WORK-SPEC-027 lifecycle and final closure Evidence.
- Confirm Work Item ownership relations and dependency graph remain consistent.
- Record only mechanically derivable closure state.

This Task must not:

- start new research;
- correct findings;
- make semantic judgments;
- classify or define terms;
- alter observations JSONL or coverage JSON;
- change the Task graph;
- stage or commit changes.

## Done condition

PRODUCT-REQ-SPEC-013, PRODUCT-WORK-SPEC-027, and the completed Task graph express the same accepted closure result without substantive research or design changes.

## Verification

- Confirm T37 is the only dependency.
- Confirm Requirement and Work Item lifecycle states match the concluded Investigation result.
- Confirm all Task ownership and dependency relations are consistent.
- Confirm no JSONL, coverage, Investigation findings, stage, or commit change occurred.

## Evidence

- PRODUCT-TASK-SPEC-027-37 owns final validation and Investigation conclusion.
- DRMCP is non-operational. Filesystem authoring was used for Design Record reads and writes.
- PRODUCT-TASK-SPEC-027-37 is `done`.
- PRODUCT-INV-SPEC-011 is `concluded`.
- PRODUCT-REQ-SPEC-013 remains `accepted`; its Evidence records the final factual corpus result.
- PRODUCT-WORK-SPEC-027 owns the completion judgment for PRODUCT-REQ-SPEC-013.
- PRODUCT-WORK-SPEC-027 changed from `in_progress` to `done`.
- PRODUCT-WORK-SPEC-027 records a PASS result for every Completion Condition.
- The Work Item lists exactly PRODUCT-TASK-SPEC-027-01 through PRODUCT-TASK-SPEC-027-38.
- All 38 Task records reference PRODUCT-WORK-SPEC-027 through `work_item`.
- PRODUCT-TASK-SPEC-027-38 depends only on PRODUCT-TASK-SPEC-027-37.
- After this synchronization, all 38 owned Tasks are `done`.
- No Task graph, JSONL, coverage, Investigation finding, canonical design content, stage, or commit changed.
