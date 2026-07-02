# PRODUCT-TASK-SPEC-019-11: Coordinate post-decomposition ADR-routing route

- **id**: PRODUCT-TASK-SPEC-019-11
- **status**: done
- **date**: 2026-07-01
- **work_item**: PRODUCT-WORK-SPEC-019
- **task_type**: coordination
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-019-10
- **outputs**:
  - PRODUCT-TASK-SPEC-019-11
  - PRODUCT-TASK-SPEC-019-12
  - PRODUCT-TASK-SPEC-019-13
  - PRODUCT-WORK-SPEC-019

## Goal

Materialize one bounded ADR-routing route and one post-routing graph owner after the W019 Work Item decomposition.

## Work

- Preserve T07 and T09 as completed decision Evidence.
- Create T12 as the ADR-routing and ADR-boundary partitioning owner.
- Create T13 as the coordination owner for exact post-routing authoring, review, and closure Tasks.
- Set T12 to depend on T11.
- Set T13 to depend on T12.
- Add T11 through T13 to W019.
- Record the external W018 finding-closure gate for the later `task_authoring` writer and W019 integrated review.
- Keep ADR authoring conditional until T12 selects exact ADR dispositions.
- Keep correction and finding-closure review Tasks absent until a W019 review names findings.

This Task must not:

- select an ADR disposition;
- author ADR, Requirement, Specification, checklist, or implementation content;
- modify `task_authoring`;
- perform review, correction, synchronization, stage, or commit work.

## Done condition

- T12 exists with one ADR-routing responsibility.
- T13 exists with one post-routing graph-coordination responsibility.
- T12 depends on T11.
- T13 depends on T12.
- W019 lists T11 through T13 and records the route.
- No speculative ADR authoring, correction, or finding-closure Task exists.
- The W018 accepted finding-closure result is recorded as the release gate for the shared `task_authoring` writer and W019 integrated review.

## Verification

- Confirm T11 through T13 IDs, types, dependencies, and outputs agree.
- Confirm W019 lists T11 through T13.
- Confirm no downstream authoring Task was created before ADR routing.
- Confirm T07, T09, and T10 remain unchanged.
- Confirm no ADR, Requirement, Specification, review, correction, synchronization, stage, or commit work occurred.

## Evidence

- T10 completed the checklist-authoring and implementation Work Item decomposition.
- T12 and T13 were materialized as separate routing and graph owners.
- Current W018 repair state includes completed T15 through T17.
- W018 T15 reserves T18 correction and T19 independent finding-closure review, but those Task files are not yet present.
- The later W019 `task_authoring` writer and integrated review therefore remain gated by independent closure of W018 findings F-BLK-01 and F-MAJ-01.
- No speculative downstream authoring or finding Tasks were created.
