# PRODUCT-TASK-SPEC-021-13: Create TRV app-local design Work Item

- **id**: PRODUCT-TASK-SPEC-021-13
- **status**: done
- **date**: 2026-07-02
- **work_item**: PRODUCT-WORK-SPEC-021
- **task_type**: work_item_decomposition
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-021-12
- **outputs**:
  - PRODUCT-WORK-SPEC-021
  - PRODUCT-TASK-SPEC-021-13
  - TRV-WORK-SPEC-001

## Goal

Create the independent TRV app-local design Work Item selected by T06.

## Work

- Require a `PASS` review, independent closure of every required T12 finding, or explicit user acceptance of an exact mechanical-only correction.
- Create `TRV-WORK-SPEC-001` with title `Define Task Responsibility Validator app-local design`.
- Place the record under `trv/records/work-items/spec/`.
- Set direct `source_refs` to this decomposition Task and `spec:product.responsibility_boundary_validator`.
- Define the accepted app-local Requirement, decision, ADR, Specification, integrated review, and closure boundary.
- Exclude implementation, executor Tasks, and current DRMCP integration.
- State that a separate TRV implementation Work Item follows design closure.
- Record only the coarse successor handoff in W021.
- Leave the child-internal Task graph to the child Work Item.

This Task must not:

- change the accepted child identity or completion boundary;
- create child Tasks, Requirements, ADRs, Specifications, or implementation artifacts;
- change W021 Task dependencies or release order;
- wait for or claim completion of the child Work Item;
- create a `work_item_execution` relation;
- perform review, synchronization, implementation, stage, or commit work.

## Done condition

- `TRV-WORK-SPEC-001` exists at the active TRV Work Item path.
- The child has one coherent app-local design Goal and Boundary.
- The child completion boundary includes Requirement, decisions, ADRs, Specifications, integrated review, and lifecycle closure.
- The child excludes implementation and current DRMCP integration.
- The child names a separate later implementation Work Item.
- Direct material sources follow the Work Item authoring rules.
- W021 records only the coarse independent-successor handoff.
- No child Task graph or completion tracking is created.

## Verification

- Confirm `TRV-WORK-SPEC-001` ID, path, metadata, and sections follow active TRV namespace rules.
- Confirm the child lists `PRODUCT-TASK-SPEC-021-13` in `source_refs`.
- Confirm the child responsibility and completion boundary match T06 D-004.
- Confirm T12 has an accepted review route or an explicit user-accepted mechanical-correction exception.
- Confirm W021 contains no child-internal Task list or `work_item_execution` Task.
- Confirm only declared outputs changed.

## Evidence

- T06 selected the exact child identity and completion boundary.
- T07 created this decomposition owner after integrated review.
- T12 returned `NEEDS REVISION` with one Minor finding, F-MIN-01.
- F-MIN-01 concerned only the W021 `impact_refs` and `## Impact Scope` projection.
- The exact mechanical correction added `spec:trv` and `TRV-WORK-SPEC-001` to `impact_refs` and removed the read-only PRODUCT-WORK-SPEC-020 checklist row from `## Impact Scope`.
- The correction changed no decision, semantic contract, ownership boundary, completion condition, Task graph, lifecycle, or release behavior.
- The user explicitly accepted the corrected projection without another review and authorized continuation to W021 closure.
- DRMCP is non-operational, so filesystem authoring was used.
- Created `TRV-WORK-SPEC-001` under `trv/records/work-items/spec/`.
- The child directly references this decomposition Task and `spec:product.responsibility_boundary_validator`.
- The child owns TRV Requirement, decision, ADR, Specification, integrated review, and design closure.
- The child excludes implementation and current DRMCP integration.
- The child names a separate later TRV implementation Work Item without fixing its identity prematurely.
- W021 records only the independent-successor handoff and does not track child completion.
- No child Task, Requirement, ADR, Specification, implementation artifact, `work_item_execution`, review, synchronization, stage, or commit was created.

### Verification result

- Child ID: `TRV-WORK-SPEC-001`.
- Child path: valid under the active TRV namespace.
- Direct source Task: present.
- PRODUCT semantic contract source: present.
- App-local design completion boundary: present.
- Implementation exclusion: present.
- Current DRMCP integration exclusion: present.
- Child Task graph: not materialized.
- Parent completion tracking: absent.
