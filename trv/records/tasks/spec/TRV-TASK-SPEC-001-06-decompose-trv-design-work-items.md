# TRV-TASK-SPEC-001-06: Decompose TRV design Work Items

- **id**: TRV-TASK-SPEC-001-06
- **status**: done
- **date**: 2026-07-02
- **work_item**: TRV-WORK-SPEC-001
- **task_type**: work_item_decomposition
- **estimate**: 0.5d
- **depends_on**:
  - TRV-TASK-SPEC-001-05
- **outputs**:
  - TRV-TASK-SPEC-001-06
  - TRV-WORK-SPEC-001
  - TRV-WORK-SPEC-002
  - TRV-WORK-SPEC-003
  - TRV-WORK-SPEC-004

## Goal

Create three sequential TRV design Work Items for architecture, application contract, and implementation-ready detailed Specifications.

## Work

- Create `TRV-WORK-SPEC-002` with title `Define TRV application architecture`.
- Create `TRV-WORK-SPEC-003` with title `Define TRV application contract`.
- Create `TRV-WORK-SPEC-004` with title `Define TRV implementation-ready detailed Specifications`.
- Give each child one independent completion judgment and integrated-review boundary.
- Use these direct sources for each child:
  - `TRV-TASK-SPEC-001-06`;
  - `TRV-TASK-SPEC-001-02`;
  - `TRV-REQ-SPEC-001`.
- Bound W002 to architecture style, logical components, dependency direction, provider/runtime ownership, and architecture ADR/Specification projection.
- Bound W003 to external interface, Task input, MCP identity and envelopes, caller interaction, compatibility, and contract ADR/Specification projection.
- Bound W004 to exact packages, files, symbols, interfaces, schemas, parsing, validation, retry, configuration, launcher, tests, fixtures, and verification commands.
- Record only the coarse route `W002 -> W003 -> W004` in TRV-WORK-SPEC-001.
- Add the three child Work Item IDs to the parent impact scope.
- Leave every child `tasks` list empty.

This Task must not:

- change the parent or child Task graph;
- create child Tasks;
- author child decisions, ADRs, Specifications, or review Evidence;
- create an implementation Work Item;
- wait for child completion;
- perform implementation, review, synchronization, stage, or commit work.

## Done condition

- TRV-WORK-SPEC-002, TRV-WORK-SPEC-003, and TRV-WORK-SPEC-004 exist.
- Each child has one distinguishable Goal and non-overlapping Boundary.
- Each child contains the exact decomposition Task, T02, and TRV Requirement in `source_refs`.
- The parent records only coarse child purposes and the W002 to W003 to W004 route.
- Child Task lists remain empty.
- No implementation Work Item or child-owned deliverable is created.

## Verification

- Confirm all IDs and paths follow active TRV SPEC Work Item rules.
- Confirm all canonical Work Item sections exist.
- Confirm architecture, contract, and detailed-Specification responsibilities do not overlap.
- Confirm W004 remains design-only and implementation-ready rather than implementation-owning.
- Confirm no Task graph changed.
- Confirm only the declared outputs changed.

## Evidence

- T02 D-016 fixed the three sequential design boundaries.
- T02 D-015 deferred implementation decomposition until reviewed detailed-Specification closure.
- T04 materialized this decomposition owner.
- TRV-WORK-SPEC-002 was created for application architecture.
- TRV-WORK-SPEC-003 was created for the external and application contract.
- TRV-WORK-SPEC-004 was created for implementation-ready detailed Specifications.
- Each child lists T06, T02, and TRV-REQ-SPEC-001 as direct sources.
- The parent records the coarse W002 to W003 to W004 route and child impact refs.
- Every child Task list remains empty.
- No child Task, canonical design artifact, implementation Work Item, stage, or commit was created.
- Result: `PASS`.
