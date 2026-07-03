# TRV-TASK-SPEC-005-03: Coordinate contract Specification authoring route

- **id**: TRV-TASK-SPEC-005-03
- **status**: not_started
- **date**: 2026-07-03
- **work_item**: TRV-WORK-SPEC-005
- **task_type**: coordination
- **estimate**: 0.5d
- **depends_on**:
  - TRV-TASK-SPEC-005-02
- **outputs**:
  - TRV-WORK-SPEC-005
  - TRV-TASK-SPEC-005-03

## Goal

Materialize the exact ADR routing, contract authoring, integrated review, and closure graph from the decided Specification placement.

## Work

- Consume terminal T02 placement and ownership decisions without changing them.
- Create exact ADR-routing ownership when the decision identifies durable choices.
- Create bounded ADR and Specification authoring Tasks for the decided artifact set.
- Serialize shared writers.
- Create one integrated independent review after the final writer.
- Create one verdict-gated closure synchronization owner.
- Keep correction and finding-closure review Tasks conditional on named findings.

This Task must not decide placement, author canonical content, review or repair findings, perform closure, detailed design, implementation, stage, or commit work.

## Done condition

- Every decided contract artifact has one exact owner.
- Writer and review order are deterministic.
- One integrated review follows the final writer.
- One closure owner follows the accepted review route.
- No speculative finding Task exists.

## Verification

- Confirm T02 precedes every authoring owner.
- Confirm artifact outputs match the decided topic tree and placement.
- Confirm shared writers are serialized.
- Confirm no canonical content or implementation work changed.

## Evidence

- T02 is the required placement and ownership decision dependency.
- Coordination has not started.
