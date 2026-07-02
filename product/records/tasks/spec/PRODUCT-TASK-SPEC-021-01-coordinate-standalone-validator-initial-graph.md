# PRODUCT-TASK-SPEC-021-01: Coordinate standalone validator initial graph

- **id**: PRODUCT-TASK-SPEC-021-01
- **status**: done
- **date**: 2026-07-02
- **work_item**: PRODUCT-WORK-SPEC-021
- **task_type**: coordination
- **estimate**: 0.5d
- **depends_on**: []
- **outputs**:
  - PRODUCT-WORK-SPEC-021
  - PRODUCT-TASK-SPEC-021-01
  - PRODUCT-TASK-SPEC-021-02
  - PRODUCT-TASK-SPEC-021-03
  - PRODUCT-TASK-SPEC-021-04

## Goal

Materialize one initial W021 decision, Investigation, and post-Investigation coordination route.

## Work

- Add T01 through T04 to the W021 Task graph.
- Create T02 as the implementation-boundary decision owner.
- Create T03 as the focused implementation impact Investigation owner.
- Reserve PRODUCT-INV-SPEC-009 for T03.
- Create T04 as the executor-ready implementation graph coordination owner.
- Persist the dependency chain T01 to T02 to T03 to T04.
- Keep implementation, verification, review, and closure Tasks conditional until T04.
- Keep execution-hub authoring, review, and release Tasks conditional on the later trigger evaluation.

This Task must not:

- make an implementation-boundary decision;
- author PRODUCT-INV-SPEC-009;
- perform production implementation;
- materialize executor Tasks;
- create a Claude Code implementation prompt;
- perform independent review or lifecycle synchronization;
- modify checklist, ADR, Requirement, or Specification content;
- stage or commit changes.

## Done condition

- T01 through T04 exist with one responsibility and completion judgment each.
- W021 lists T01 through T04 in order.
- Dependencies form T01 to T02 to T03 to T04.
- PRODUCT-INV-SPEC-009 is reserved only for T03.
- Implementation and execution-hub Tasks remain conditional.
- No implementation contract or Claude Code prompt is produced.

## Verification

- Confirm every Task uses one canonical `task_type`.
- Confirm every Task references PRODUCT-WORK-SPEC-021.
- Confirm Work Item `tasks` and Task `work_item` relations match.
- Confirm T01 is `done` and T02 through T04 are `not_started`.
- Confirm no implementation Task or Investigation record exists.
- Confirm no production or checklist file changed.

## Evidence

- PRODUCT-WORK-SPEC-019 is `done` through the accepted W019 closure route.
- PRODUCT-WORK-SPEC-020 is `done` after T05 returned `PASS` and T06 synchronized closure.
- The reviewed checklist artifact set is released for W021 consumption.
- W021 had `tasks: []` before this Task.
- T01 materialized only the initial decision, Investigation, and coordination route.
- PRODUCT-INV-SPEC-009 was unused and is reserved for T03.
- No implementation, Investigation authoring, executor Task, or Claude Code prompt was produced.
- DRMCP is non-operational, so filesystem authoring was used.
- Result: `PASS`.
