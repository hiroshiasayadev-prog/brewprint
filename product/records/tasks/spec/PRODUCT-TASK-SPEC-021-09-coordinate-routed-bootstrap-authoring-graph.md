# PRODUCT-TASK-SPEC-021-09: Coordinate routed bootstrap authoring graph

- **id**: PRODUCT-TASK-SPEC-021-09
- **status**: done
- **date**: 2026-07-02
- **work_item**: PRODUCT-WORK-SPEC-021
- **task_type**: coordination
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-021-08
- **outputs**:
  - PRODUCT-WORK-SPEC-021
  - PRODUCT-TASK-SPEC-021-09
  - PRODUCT-TASK-SPEC-021-10
  - PRODUCT-TASK-SPEC-021-16

## Goal

Materialize the exact ADR-authoring route from T08 and release the serialized canonical authoring chain.

## Work

- Consume the terminal T08 routing result without changing it.
- Create exact `authoring` Tasks only for ADR work classified as required.
- Do not create an ADR-authoring Task when every route is `covered` or `not_required`.
- Add every required ADR writer to W021 ownership.
- Serialize required ADR writers before T10.
- Amend T10 dependencies to include the final routed ADR writer.
- Preserve T10 through T14 responsibility and completion boundaries.
- Update this Task outputs and W021 Task flow with exact created Task IDs.

This Task must not:

- choose or change an ADR disposition;
- author or modify ADR, namespace, Specification, or app-local content;
- change T06 or T08 outcomes;
- create correction or finding-closure review Tasks before named findings exist;
- create implementation or `work_item_execution` Tasks;
- perform review, synchronization, stage, or commit work.

## Done condition

- Every required ADR boundary has one exact authoring owner.
- No unnecessary or no-op ADR-authoring Task exists.
- Required ADR writers have deterministic dependencies and writer order.
- T10 depends on T09 and every required prior ADR writer.
- W021 lists every created Task.
- T10 through T14 remain released only through the accepted routed order.

## Verification

- Confirm every created Task has one `authoring` responsibility and one completion judgment.
- Confirm every created Task uses the next available W021 Task sequence.
- Confirm Work Item `tasks` and Task `work_item` relations agree.
- Confirm T10 cannot begin before required ADR authoring completes.
- Confirm no canonical artifact content changed.
- Confirm no speculative correction, implementation, or child-execution Task exists.

## Evidence

- T07 created this conditional graph-coordination owner.
- T08 completed corrected ADR routing with one required authoring boundary and no blockers.
- T08 B-001 requires a non-material amendment to `PRODUCT-ADR-SPEC-016`.
- The erroneous extra ADR-authoring owner was removed because workflow-artifact correction is not ADR authoring.
- `PRODUCT-TASK-SPEC-021-16` is the sole ADR-authoring owner.
- T16 now depends directly on T09.
- T10 now depends on T09 and T16.
- W021 ownership and Task flow were corrected accordingly.
- No other ADR-authoring Task exists.
- No ADR, Specification, namespace profile, Work Item deliverable, review, synchronization, implementation, stage, or commit work was performed.

### Writer and release order

```text
PRODUCT-TASK-SPEC-021-09
  -> PRODUCT-TASK-SPEC-021-16 amend PRODUCT-ADR-SPEC-016
  -> PRODUCT-TASK-SPEC-021-10 author TRV namespace profile
```

### Verification result

- Required ADR boundaries with owners: 1 of 1.
- Unnecessary ADR-authoring Tasks: 0.
- ADR writer order: deterministic.
- T10 release dependencies: T09 and T16.
- Work Item to Task ownership: bidirectionally represented.
- Circular dependencies: none.
- Speculative correction Tasks: none.
- Implementation Tasks: none.
- `work_item_execution` Tasks: none.
