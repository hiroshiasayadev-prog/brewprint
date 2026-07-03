# TRV-TASK-SPEC-001-01: Coordinate TRV app-local design initial graph

- **id**: TRV-TASK-SPEC-001-01
- **status**: done
- **date**: 2026-07-02
- **work_item**: TRV-WORK-SPEC-001
- **task_type**: coordination
- **estimate**: 0.5d
- **depends_on**: []
- **outputs**:
  - TRV-WORK-SPEC-001
  - TRV-TASK-SPEC-001-01
  - TRV-TASK-SPEC-001-02
  - TRV-TASK-SPEC-001-03
  - TRV-TASK-SPEC-001-04

## Goal

Materialize one initial TRV app-local design decision and downstream graph-coordination route.

## Work

- Create `trv/records/tasks/spec/`.
- Add T01 through T04 to the TRV-WORK-SPEC-001 Task graph.
- Create T02 as the TRV app-local design decision owner.
- Create T03 as the graph-amendment owner when the initial route needs correction.
- Create T04 as the downstream design graph-coordination owner.
- Persist the dependency chain T01 to T02 to T03 to T04.
- Set TRV-WORK-SPEC-001 to `in_progress`.
- Start T02 with one decision item in discussion.
- Keep Requirement authoring, ADR routing, canonical authoring, review, correction, synchronization, and implementation handoff conditional until T04.

This Task must not:

- decide any TRV app-local design outcome;
- author TRV-REQ-SPEC-001 or another Requirement;
- author an ADR or Specification;
- materialize implementation Tasks;
- create an executor prompt;
- perform implementation, review, correction, synchronization, stage, or commit work.

## Done condition

- T01 through T04 exist with one responsibility and one completion judgment each.
- TRV-WORK-SPEC-001 lists T01 through T04 in sequence.
- Dependencies form T01 to T02 to T03 to T04.
- T01 is `done`.
- T02 is `in_progress` with at most one item `in_discussion`.
- T03 and T04 are `not_started`.
- Conditional downstream Tasks remain unmaterialized.
- No Requirement, Investigation, ADR, Specification, implementation Task, or executor prompt is produced.

## Verification

- Confirm every Task ID and path follows the active TRV / SPEC namespace rules.
- Confirm every Task uses one canonical `task_type`.
- Confirm every Task references TRV-WORK-SPEC-001.
- Confirm Work Item `tasks` and Task `work_item` relations match.
- Confirm T01 alone is `done`.
- Confirm T02 alone is `in_progress`.
- Confirm T03 and T04 are `not_started`.
- Confirm no downstream deliverable exists.
- Confirm no production file changed.

## Evidence

- TRV is active under the canonical namespace and repository-layout profiles.
- TRV-WORK-SPEC-001 had `tasks: []` before this Task.
- T01 materialized only the initial decision and coordination route.
- T03 later corrected the graph by removing the unnecessary standalone impact-Investigation route; no Investigation record was created.
- T02 started with D-001 as the only `in_discussion` item.
- Requirement authoring, ADR routing, canonical authoring, review, correction, synchronization, and implementation handoff remain conditional.
- PRODUCT-owned validator semantics remain fixed inputs rather than TRV decision items.
- DRMCP is non-operational under the current agent authoring policy, so filesystem authoring was used.
- No production implementation, executor prompt, stage, or commit was performed.
- Result: `PASS`.
