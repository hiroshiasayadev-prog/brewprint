# PRODUCT-TASK-SPEC-019-10: Decompose checklist authoring and validator implementation Work Items

- **id**: PRODUCT-TASK-SPEC-019-10
- **status**: done
- **date**: 2026-07-01
- **work_item**: PRODUCT-WORK-SPEC-019
- **task_type**: work_item_decomposition
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-019-09
- **outputs**:
  - PRODUCT-TASK-SPEC-019-10
  - PRODUCT-WORK-SPEC-019
  - PRODUCT-WORK-SPEC-020
  - PRODUCT-WORK-SPEC-021

## Goal

Create separate downstream Work Items for Task responsibility-boundary checklist authoring and temporary standalone validator implementation.

## Work

- Create PRODUCT-WORK-SPEC-020 for checklist authoring.
- Create PRODUCT-WORK-SPEC-021 for standalone validator implementation.
- Give each Work Item one distinct Goal and completion boundary.
- Add this Task as a direct material source to both Work Items.
- Record only coarse downstream routing in W019.
- Leave each child Work Item's internal Task graph undecided.

This Task must not:

- change W019 Task dependencies or release conditions;
- author checklist content;
- implement the validator;
- create child Tasks;
- perform ADR routing, canonical Specification authoring, review, synchronization, stage, or commit work.

## Done condition

- PRODUCT-WORK-SPEC-020 exists with a checklist-authoring boundary.
- PRODUCT-WORK-SPEC-021 exists with a standalone implementation boundary.
- Both Work Items include PRODUCT-TASK-SPEC-019-10 in `source_refs`.
- The Work Item responsibilities do not overlap.
- W019 records the coarse W020 to W021 downstream route.
- No child Task or child-owned deliverable is authored.

## Verification

- Confirm both Work Item IDs and paths follow PRODUCT SPEC namespace rules.
- Confirm both Work Items contain all canonical sections.
- Confirm checklist wording and storage belong only to W020.
- Confirm executable implementation belongs only to W021.
- Confirm W021 consumes accepted W020 checklist artifacts.
- Confirm no Task-graph, checklist, implementation, ADR, Specification, review, synchronization, stage, or commit work occurred.

## Evidence

- T09 fixed the two independent Work Item identities and release order.
- PRODUCT-WORK-SPEC-020 and PRODUCT-WORK-SPEC-021 were created.
- W019 records the downstream routing without duplicating child-internal Task graphs.
- No child Task or child-owned deliverable was created.
- Result: `PASS`.
