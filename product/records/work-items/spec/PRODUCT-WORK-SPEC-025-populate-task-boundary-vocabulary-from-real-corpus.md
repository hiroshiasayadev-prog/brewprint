# PRODUCT-WORK-SPEC-025: Populate task-boundary-vocabulary from real corpus

- **id**: PRODUCT-WORK-SPEC-025
- **status**: done
- **date**: 2026-07-03
- **source_refs**:
  - PRODUCT-REQ-SPEC-011
- **impact_refs**:
  - spec:product.design_records.authoring_standards.task_authoring
  - skills/task-boundary-vocabulary/
- **tasks**:
  - PRODUCT-TASK-SPEC-025-01
  - PRODUCT-TASK-SPEC-025-02
  - PRODUCT-TASK-SPEC-025-03
  - PRODUCT-TASK-SPEC-025-04
  - PRODUCT-TASK-SPEC-025-05
  - PRODUCT-TASK-SPEC-025-06
  - PRODUCT-TASK-SPEC-025-07
  - PRODUCT-TASK-SPEC-025-08
  - PRODUCT-TASK-SPEC-025-09
  - PRODUCT-TASK-SPEC-025-10
  - PRODUCT-TASK-SPEC-025-11

## Goal

Populate `skills/task-boundary-vocabulary/` with corpus-verified boundary-vocabulary entries beyond `decision.md`, using a parallel chappy-log-extraction-then-Claude-review pipeline instead of a heavyweight formal-Investigation route.

## Boundary

This Work Item owns:

- the Investigation-Task lightweight Evidence exception in `spec:product.design_records.authoring_standards.task_authoring`;
- five parallel, corpus-range-scoped chappy-executed log-extraction Tasks, each recording raw phrase / source Task ID / vocabulary target / effective meaning entries directly in its own Evidence;
- one Claude-executed review Task that reconciles the five logs, verifies entries against source Tasks, clusters them, and writes confirmed entries into `skills/task-boundary-vocabulary/` task_type files;
- one Claude-executed authoring Task that reorganizes the same five logs' full raw entries (not just the boundary-violation subset) into a general, per-task_type vocabulary reference, as a separate artifact from the boundary-violation dictionary.

This Work Item does not own:

- TRV application reimplementation or automated-validator re-evaluation;
- a full rewrite of existing Task records;
- ASD-STE100 adoption;
- production lint implementation;
- a fixed 12-field collection schema or parallel formal Investigation Tasks per responsibility cluster;
- canonical-term reconciliation (a separate later Work Item, after roughly 30 log entries accumulate).

## Impact Scope

| target | impact |
|---|---|
| `spec:product.design_records.authoring_standards.task_authoring` | Add the Investigation-Task lightweight Evidence exception. |
| `skills/task-boundary-vocabulary/` | Add corpus-verified entries to at least one task_type file beyond `decision.md`. |
| `skills/task-vocabulary-reference/` | New artifact: general per-task_type vocabulary reference built from the same corpus scan. |

## Task flow

```text
PRODUCT-TASK-SPEC-025-01 vocabulary-collection decision
  -> PRODUCT-TASK-SPEC-025-02 author lightweight investigation-Evidence exception
     -> PRODUCT-TASK-SPEC-025-05 coordinate shared-writer split
        -> PRODUCT-TASK-SPEC-025-06 extract W001-W007 (investigation, Evidence-only)
        -> PRODUCT-TASK-SPEC-025-07 extract W009-W012 (investigation, Evidence-only)
        -> PRODUCT-TASK-SPEC-025-08 extract W013-W017 (investigation, Evidence-only)
        -> PRODUCT-TASK-SPEC-025-09 extract W018-W019 (investigation, Evidence-only)
        -> PRODUCT-TASK-SPEC-025-10 extract W020-W025 (investigation, Evidence-only)
           -> PRODUCT-TASK-SPEC-025-04 reconcile, review, cluster, and author vocabulary entries
```

`PRODUCT-TASK-SPEC-025-03` (the original single sequential-session extraction Task) is `cancelled`; `PRODUCT-TASK-SPEC-025-05` replaced it with the five parallel Tasks above to remove a shared-writer conflict.

T01 directly materialized T02. T05 (coordination) then replaced the originally materialized T03 with T06–T10 and re-pointed T04.

## Task Candidates

| task | task type | responsibility | dependency |
|---|---|---|---|
| `PRODUCT-TASK-SPEC-025-01` | `decision` | Fix the vocabulary-collection Requirement handling, Task graph, and lightweight-exception decision. | none |
| `PRODUCT-TASK-SPEC-025-02` | `authoring` | Add the Investigation-Task lightweight Evidence exception to task-authoring standards. | T01 |
| `PRODUCT-TASK-SPEC-025-03` | `investigation` | Cancelled: replaced by T05–T10 to remove a shared-writer conflict. | T02 |
| `PRODUCT-TASK-SPEC-025-05` | `coordination` | Replace T03 with five parallel-safe investigation Tasks and re-point T04. | T01 |
| `PRODUCT-TASK-SPEC-025-06` | `investigation` | Extract boundary-vocabulary entries from `PRODUCT-WORK-SPEC-001`–`007`. | T02 |
| `PRODUCT-TASK-SPEC-025-07` | `investigation` | Extract boundary-vocabulary entries from `PRODUCT-WORK-SPEC-009`–`012`. | T02 |
| `PRODUCT-TASK-SPEC-025-08` | `investigation` | Extract boundary-vocabulary entries from `PRODUCT-WORK-SPEC-013`–`017`. | T02 |
| `PRODUCT-TASK-SPEC-025-09` | `investigation` | Extract boundary-vocabulary entries from `PRODUCT-WORK-SPEC-018`–`019`. | T02 |
| `PRODUCT-TASK-SPEC-025-10` | `investigation` | Extract boundary-vocabulary entries from `PRODUCT-WORK-SPEC-020`–`025` (excluding this Work Item's own Tasks). | T02 |
| `PRODUCT-TASK-SPEC-025-04` | `authoring` | Reconcile the five logs; verify, cluster, and write confirmed entries into `skills/task-boundary-vocabulary/`. | T06, T07, T08, T09, T10 |
| `PRODUCT-TASK-SPEC-025-11` | `authoring` | Build a general per-task_type vocabulary reference from the same five logs' full raw entries. | T06, T07, T08, T09, T10 |

No additional Task is materialized speculatively.

## Completion Condition

- T01, T02, T04, T05, T06, T07, T08, T09, T10, and T11 are `done`. T03 is `cancelled`.
- `skills/task-boundary-vocabulary/` contains corpus-verified entries in at least one task_type file beyond `decision.md`.
- `skills/task-vocabulary-reference/` contains a general per-task_type vocabulary reference.
- Reconciliation of canonical term definitions remains excluded from this Work Item.

## Evidence

- PRODUCT-REQ-SPEC-011 is accepted and defines the vocabulary-collection requirement.
- The user redirected the app namespace from TRV to PRODUCT; TRV application work is deferred.
- The user required a per-owner Task split between chappy's raw extraction and Claude's review, rather than one combined authoring Task.
- The user identified that `investigation` task_type otherwise requires a full Investigation record, which reproduces the heavyweight route already rejected for this work.
- The user directed adding a lightweight Evidence-only exception to `spec:product.design_records.authoring_standards.task_authoring` instead of relabeling chappy's Task as `authoring`.
- T01 records the complete accepted decision set and directly materialized T02.
- T02 added the accepted exception to `spec:product.design_records.authoring_standards.task_authoring`.
- chappy identified that a single-session, single-Evidence T03 exceeded reliable read depth for a 179-Task corpus; T03 was first amended to a coverage-ledger, 6-session sequential design.
- The user then required a true Task split instead of sequential sessions, to allow parallel execution without a shared-writer conflict.
- T05 (coordination) cancelled T03, created T06–T10 (one investigation Task per disjoint corpus range, each with its own Evidence write target), and re-pointed T04 to depend on all five.
- The user reviewed T04's output and found the promoted-entry count (1 new pattern, plus reinforcement of an existing one) thin relative to the 219 raw findings and 177 scanned Tasks. Verification confirmed the raw extraction stage was not thin; the small file output reflects T04's promotion filter (only cross-type disguises, not normal in-type vocabulary), by the boundary-violation dictionary's own design.
- The user clarified they also wanted a general, per-task_type vocabulary reference (normal in-type phrasing) as a separate artifact, built from the same already-collected T06–T10 raw logs. `PRODUCT-REQ-SPEC-011` and this Work Item were amended to include it; `PRODUCT-TASK-SPEC-025-11` was added.
- T11 built `skills/task-vocabulary-reference/` (SKILL.md + 4 cluster files) from the same 219 raw findings, deduplicating repeated phrasing. It grouped by the `vocabulary target` field rather than strictly by source Task type, since T06/T08 lack a per-entry type field and re-reading source Tasks was out of scope — recorded as a deviation in T11's own Evidence.
- DRMCP is non-operational. Filesystem authoring is the required fallback.
- T04 reconciled all five logs: 177 scanned Tasks against 177 actual `status: done` Tasks in scope (181 files minus 4 confirmed non-done exclusions), no unresolved gap or overlap. `skills/task-boundary-vocabulary/correction.md` was created with one corpus-verified entry (5 corroborating instances); `decision.md`'s existing "Fix" entry was strengthened from 2 to 12 corpus citations. Every one of the 219 logged findings received a recorded conclusion (reflected, isolated as Open question, or explicitly matched-no-action), satisfying the D-006 exception's closure condition.
- T01, T02, T04, T05, T06, T07, T08, T09, T10, and T11 are all `done`; T03 is `cancelled`. Completion Condition satisfied.
