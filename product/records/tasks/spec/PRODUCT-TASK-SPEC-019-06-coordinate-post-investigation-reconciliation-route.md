# PRODUCT-TASK-SPEC-019-06: Coordinate post-Investigation reconciliation route

- **id**: PRODUCT-TASK-SPEC-019-06
- **status**: done
- **date**: 2026-07-01
- **work_item**: PRODUCT-WORK-SPEC-019
- **task_type**: coordination
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-019-05
- **outputs**:
  - PRODUCT-WORK-SPEC-019
  - PRODUCT-TASK-SPEC-019-06
  - PRODUCT-TASK-SPEC-019-07

## Goal

Materialize the missing post-Investigation reconciliation decision owner for W019.

## Work

- Add T06 and T07 to the W019 Task inventory.
- Replace the abstract immediate post-T05 reconciliation step with exact T06 and T07 routing.
- Create T07 with one bounded reconciliation decision responsibility.
- Preserve downstream graph coordination, ADR routing, authoring, review, and closure as conditional candidates.
- Treat W018 T11 J-001 as the accepted authority for MC-002.
- Require T07 to consume W018 T11 J-001 without reopening the Task-type split.
- Require later W019 graph coordination to serialize shared `task_authoring` writing after the W018 canonical repair.
- Preserve T01 through T05 without substantive changes.
- Keep W019 `status: in_progress`.

This Task does not perform the T07 decision loop.
This Task does not select an ADR route or author downstream deliverables.

## Done condition

- W019 lists T06 and T07.
- T06 exists as one `coordination` Task.
- T07 exists as one `decision` Task.
- T07 depends on T06.
- T07 owns all remaining W019 reconciliation judgments and no ADR-routing judgment.
- MC-002 has no duplicate W019 decision owner.
- No downstream deliverable Task is prematurely materialized.
- No completed Task is substantively rewritten.

## Verification

- Confirm Work Item and Task ownership is bidirectional.
- Confirm T06 and T07 IDs match the W019 parent sequence.
- Confirm T06 outputs exactly W019, T06, and T07.
- Confirm T07 outputs only T07.
- Confirm the dependency route has no cycle.
- Confirm MC-002 references W018 T11 J-001 as existing authority.
- Confirm T07 does not own ADR routing.
- Confirm correction and finding-closure Tasks remain absent.
- Confirm only the three allowed paths changed in scoped inspection.
- Confirm scoped whitespace inspection passes.

## Evidence

- Accepted trigger Task: `PRODUCT-TASK-SPEC-019-05`.
- Concluded trigger Investigation: `PRODUCT-INV-SPEC-007`.
- MC-001 requires a W019 reconciliation decision before canonical authoring.
- MC-002 decision authority is `PRODUCT-TASK-SPEC-018-11`, `### J-001 decision`.
- T06 materialized T07 without executing its decision loop.
- Downstream graph coordination remains conditional.
- ADR routing remains downstream of T07.
- Any later W019 writer for `spec:product.design_records.authoring_standards.task_authoring` must follow the W018 canonical repair.
- DRMCP is non-operational under `spec:product.design_records.authoring_standards.agent_authoring_policy`.
- Filesystem authoring was used for W019, T06, and T07.
- Changed paths:
  - `product/records/work-items/spec/PRODUCT-WORK-SPEC-019-design-semantic-task-responsibility-boundary-validation.md`
  - `product/records/tasks/spec/PRODUCT-TASK-SPEC-019-06-coordinate-post-investigation-reconciliation-route.md`
  - `product/records/tasks/spec/PRODUCT-TASK-SPEC-019-07-decide-post-investigation-reconciliation.md`
- Scoped Git diff inspection covered only the three changed paths.
- Scoped whitespace inspection passed with no whitespace findings.
- Staged changes were absent for the scoped paths.
- No stage or commit was performed.
