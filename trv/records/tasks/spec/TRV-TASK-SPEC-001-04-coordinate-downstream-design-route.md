# TRV-TASK-SPEC-001-04: Coordinate downstream TRV design route

- **id**: TRV-TASK-SPEC-001-04
- **status**: done
- **date**: 2026-07-02
- **work_item**: TRV-WORK-SPEC-001
- **task_type**: coordination
- **estimate**: 0.5d
- **depends_on**:
  - TRV-TASK-SPEC-001-03
- **outputs**:
  - TRV-WORK-SPEC-001
  - TRV-TASK-SPEC-001-04
  - TRV-TASK-SPEC-001-05
  - TRV-TASK-SPEC-001-06
  - TRV-TASK-SPEC-001-07

## Goal

Materialize one downstream design-convergence graph from terminal T02 decisions after a bounded repository-alignment check.

## Work

- Consume terminal T02 decisions without reopening them.
- Confirm the active TRV namespace, authoring guides, PRODUCT semantic authority, and checklist locations.
- Confirm no concrete contradiction requires a reconciliation decision before decomposition.
- Create T05 as the accepted TRV Requirement authoring owner.
- Create T06 as the parent-to-child Work Item decomposition owner.
- Fix exact child identities in T06:
  - TRV-WORK-SPEC-002 for application architecture;
  - TRV-WORK-SPEC-003 for the external and application contract;
  - TRV-WORK-SPEC-004 for implementation-ready detailed Specifications.
- Create T07 as the post-decomposition graph owner.
- Require T07 to create one child-local bootstrap owner and one parent `work_item_execution` Task per child after the child Work Items exist.
- Require the parent route to execute W002, then W003, then W004.
- Require a final integrated parent review and verdict-gated closure synchronization after child completion.
- Preserve the later implementation-readiness gate without creating an implementation Work Item.
- Keep conflict, correction, and finding-closure branches unmaterialized until concrete findings exist.

This Task must not:

- create an Investigation record;
- make a TRV design or reconciliation decision;
- author or amend a Requirement, ADR, or Specification body;
- perform implementation planning or production implementation;
- create an implementation Work Item, implementation Task, or executor prompt;
- issue a review verdict or repair findings;
- perform lifecycle closure, stage, or commit work.

## Done condition

- T05 owns accepted Requirement authoring.
- T06 owns creation of the three fixed child Work Items.
- T07 owns post-decomposition parent execution tracking and child bootstrap graph materialization.
- Dependencies form T04 to T05 to T06 to T07.
- The accepted route requires W002 before W003 and W003 before W004.
- Child-internal canonical authoring, ADR routing, review, and closure remain child-owned.
- Conditional conflict and finding work remains unmaterialized until concrete findings exist.
- No standalone Investigation, production implementation, implementation Work Item, executor prompt, or current DRMCP integration route exists.

## Verification

- Confirm T05, T06, and T07 each have one canonical `task_type`, primary outcome, and completion judgment.
- Confirm Task IDs and paths follow active TRV SPEC rules.
- Confirm TRV-WORK-SPEC-001 lists T05 through T07.
- Confirm dependencies are acyclic and form T04 to T05 to T06 to T07.
- Confirm T06 creates Work Items but no child Tasks.
- Confirm T07 changes graphs only after the child Work Items exist.
- Confirm no Investigation dependency or `TRV-INV-SPEC-001` reference remains in the active route.
- Confirm no implementation Work Item, implementation Task, executor prompt, or current DRMCP integration route is created.

## Evidence

- T02 concluded the app-local decision ledger.
- T03 removed the unnecessary mandatory impact-Investigation route for this new application.
- No TRV Investigation record exists.
- Repository alignment found no contradiction requiring a new decision before decomposition.
- T05 was created for `TRV-REQ-SPEC-001` authoring.
- T06 was created for TRV-WORK-SPEC-002 through TRV-WORK-SPEC-004 decomposition.
- T07 was created for post-decomposition execution tracking and child graph bootstrap.
- No child Work Item, child Task, implementation Work Item, or canonical design artifact was created by T04.
- Result: `PASS`.
