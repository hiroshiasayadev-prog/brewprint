# DRCLI-TASK-CLI-002-01: Coordinate DRCLI Specification authoring graph

- **status**: done
- **date**: 2026-10-07
- **work_item**: DRCLI-WORK-CLI-002
- **task_type**: coordination
- **estimate**: 0.5d
- **depends_on**: []
- **outputs**:
  - DRCLI-WORK-CLI-002
  - DRCLI-TASK-CLI-002-02
  - DRCLI-TASK-CLI-002-03
  - DRCLI-TASK-CLI-002-04
  - DRCLI-TASK-CLI-002-05
  - DRCLI-TASK-CLI-002-06
  - DRCLI-TASK-CLI-002-07
  - DRCLI-TASK-CLI-002-08
  - DRCLI-TASK-CLI-002-09
  - DRCLI-TASK-CLI-002-10
  - DRCLI-TASK-CLI-002-11

## Goal

Create one executable, writer-disjoint Task graph for DRCLI Specification convergence.

## Work

- Split the Specification by non-overlapping path ownership.
- Separate DRCLI-specific design judgment from mechanical DRMCP adaptation.
- Fix dependency order for CLI-surface decisions, parallel authoring, integration, independent review, and closure.
- Assign one writer per path family.
- Keep correction Tasks conditional on actual review findings.

## Done condition

- Every required Task exists.
- Parallel authoring writers have disjoint path families.
- T08 depends on the CLI-specific decision Task.
- T09 depends on all authoring writers.
- T10 is independent and read-only.
- T11 owns only accepted lifecycle/Evidence synchronization.
- No speculative correction Task exists.

## Verification

- Compare the writer map against every Task Change boundary.
- Confirm the graph is acyclic.
- Confirm no parallel Task writes root indexes.
- Confirm implementation is absent.

## Evidence

- DRCLI-WORK-CLI-002 records the accepted writer map and dependency flow.
- T02 through T11 were materialized from this coordination outcome.
