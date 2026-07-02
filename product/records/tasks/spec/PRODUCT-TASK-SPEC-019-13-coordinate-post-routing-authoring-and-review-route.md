# PRODUCT-TASK-SPEC-019-13: Coordinate post-routing authoring and review route

- **id**: PRODUCT-TASK-SPEC-019-13
- **status**: done
- **date**: 2026-07-01
- **work_item**: PRODUCT-WORK-SPEC-019
- **task_type**: coordination
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-019-12
- **outputs**:
  - PRODUCT-TASK-SPEC-019-13
  - PRODUCT-TASK-SPEC-019-14
  - PRODUCT-TASK-SPEC-019-15
  - PRODUCT-TASK-SPEC-019-16
  - PRODUCT-TASK-SPEC-019-17
  - PRODUCT-TASK-SPEC-019-18
  - PRODUCT-TASK-SPEC-019-19
  - PRODUCT-WORK-SPEC-019

## Goal

Materialize the exact W019 authoring, shared-writer, integrated-review, and closure route selected by T12.

## Work

- Consume the completed T12 ADR-routing ledger without reopening its decisions.
- Create T14 as the authoring owner for PRODUCT-ADR-SPEC-015 through PRODUCT-ADR-SPEC-017.
- Create T15 as the PRODUCT-REQ-SPEC-007 amendment-disposition owner.
- Create T16 as the PRODUCT-root ownership-alignment, dedicated validator Specification, and parent-registration owner.
- Create T17 as the separate `task_authoring` usage-rule writer.
- Create T18 as the one integrated independent W019 review owner.
- Create T19 as the verdict-gated W019 closure-synchronization owner.
- Serialize T14 through T17 in canonical authoring order.
- Require T17 to depend on independent W018 finding closure.
- Keep W020 and W021 outside the W019 internal Task graph.
- Keep correction and finding-closure review Tasks conditional on named W019 findings.

This Task must not:

- change ADR-routing decisions;
- author ADR, Requirement, Specification, checklist, or implementation content;
- repair W018 findings;
- perform integrated review, correction, synchronization, stage, or commit work.

## Done condition

- T14 through T19 exist with one responsibility and completion judgment each.
- T14 owns exactly the three new ADRs selected by T12.
- T15 owns only the Requirement amendment disposition and may complete with a no-amendment result.
- T16 owns only the bounded PRODUCT-ADR-SPEC-001 root alignment, validator Specification, and direct PRODUCT-parent registration.
- T17 owns only the narrow shared `task_authoring` update.
- T18 is the only W019 integrated review owner.
- T19 is the only W019 closure-synchronization owner.
- Writer, review, and release order are deterministic.
- No speculative correction or finding-closure review Task exists.

## Verification

- Confirm every materialized Task ID, type, dependency, output, and parent Work Item.
- Confirm T14 through T17 follow ADR, Requirement, Specification, then shared-writer order.
- Confirm T17 depends on PRODUCT-TASK-SPEC-018-19.
- Confirm PRODUCT-TASK-SPEC-018-19 reports `PASS` and closes F-BLK-01 and F-MAJ-01.
- Confirm T18 depends on every W019 authoring Task.
- Confirm T19 remains verdict-gated.
- Confirm W020 and W021 are represented only through coarse downstream routing.
- Confirm no ADR, Requirement, Specification, checklist, implementation, review, correction, synchronization, stage, or commit work occurred.

## Evidence

- T12 selected PRODUCT-ADR-SPEC-015 through PRODUCT-ADR-SPEC-017 and exact canonical targets.
- PRODUCT-TASK-SPEC-018-19 is `done`, reports `PASS`, and closes F-BLK-01 and F-MAJ-01.
- T14 through T19 were materialized.
- Authoring order is T14 -> T15 -> T16 -> T17.
- Integrated review follows every writer through T18.
- Closure remains conditional through T19.
- No child-owned deliverable or canonical artifact was authored by this coordination Task.
- Result: `PASS`.
