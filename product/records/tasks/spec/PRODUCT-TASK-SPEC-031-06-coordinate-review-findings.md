# PRODUCT-TASK-SPEC-031-06: Coordinate USDM hierarchy review findings

- **status**: done
- **date**: 2026-09-30
- **work_item**: PRODUCT-WORK-SPEC-031
- **task_type**: coordination
- **estimate**: 0.25d
- **depends_on**:
  - PRODUCT-TASK-SPEC-031-04
- **outputs**:
  - PRODUCT-TASK-SPEC-031-07
  - PRODUCT-TASK-SPEC-031-08
  - PRODUCT-TASK-SPEC-031-09

## Goal

Route REV-031-001 and REV-031-002, plus the accepted cross-app hierarchy boundary, into append-only reconvergence work.

## Work

- Preserve PRODUCT-TASK-SPEC-031-01 through 031-04 as completed history.
- Create one new decision Task for sibling allocation, coverage response shape, and cross-app row-hierarchy boundary.
- Create one authoring Task after the decision.
- Create one new independent integrated review after authoring.
- Re-route closure synchronization to the new review.

## Done condition

Every named review finding and new accepted hierarchy-boundary judgment has an explicit owner and dependency route.

## Verification

The new route is decision -> authoring -> independent review -> existing closure synchronization.

## Evidence

- REV-031-001 requires a new decision because PRODUCT-TASK-SPEC-031-01 is complete.
- REV-031-002 requires a new decision because the public response shape was not uniquely determined.
- The accepted namespace boundary is: USDM row decomposition remains within one USDM requirement record; cross-app realization uses full-ID `usdm_covers`.
- PRODUCT-TASK-SPEC-031-07 owns the new decisions.
- PRODUCT-TASK-SPEC-031-08 owns ADR/spec projection.
- PRODUCT-TASK-SPEC-031-09 owns the independent re-review.
