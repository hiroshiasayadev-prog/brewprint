# PRODUCT-TASK-SPEC-023-01: Decide cancelled lifecycle contract

- **id**: PRODUCT-TASK-SPEC-023-01
- **status**: done
- **date**: 2026-07-03
- **work_item**: PRODUCT-WORK-SPEC-023
- **task_type**: decision
- **estimate**: 0.5d
- **depends_on**: []
- **outputs**:
  - PRODUCT-TASK-SPEC-023-01

## Goal

Fix one bounded semantic contract for terminal Work Item and Task cancellation.

## Work

- Decide whether an unfinished Task may be cancelled independently.
- Decide how independent Task cancellation affects its parent Work Item.
- Decide how a cancelled prerequisite affects dependent Tasks.
- Decide cancellation propagation coherence for a Work Item and its owned Tasks.
- Decide required cancellation Evidence.
- Decide the effect of a cancelled child Work Item on `work_item_execution`.
- Decide post-cancellation mutation and terminality boundaries.
- Preserve all scope exclusions fixed by PRODUCT-REQ-SPEC-009.
- Record concise workflow Evidence and provisional canonical targets.

This Task must not:

- investigate repository-wide impact;
- author ADR, Specification, workflow-support, or validation content;
- change the Task graph;
- perform review, correction, synchronization, implementation, stage, or commit work.

### Decision ledger

| ID | topic | status | depends on | decision summary | reason | canonical target | ADR route |
|---|---|---|---|---|---|---|---|
| D-001 | Independent Task cancellation | `decided` | none | A Task may move independently from `not_started`, `in_progress`, or `blocked` to `cancelled`. A `done` or `cancelled` Task cannot be cancelled. | Independent Task cancellation is required to represent intentionally abandoned work without cancelling the parent Work Item. | Task authoring | `candidate` |
| D-002 | Work Item cancellation eligibility | `decided` | none | A Work Item may move from `not_started`, `in_progress`, or `blocked` to `cancelled`. A `done` or `cancelled` Work Item cannot be cancelled. | Cancellation occurs before completion and `cancelled` is terminal. | Work Item authoring | `candidate` |
| D-003 | Work Item-to-Task propagation | `decided` | D-002 | Cancelling a Work Item changes each owned `not_started`, `in_progress`, or `blocked` Task to `cancelled`. Owned `done` Tasks remain unchanged. | PRODUCT-REQ-SPEC-009 fixes this propagation contract. | Work Item and Task authoring | `candidate` |
| D-004 | Parent Work Item after independent Task cancellation | `decided` | D-001 | Cancelling one Task does not change the parent Work Item status. The parent may later become `done` only when its current Completion Condition remains satisfied without that Task. Otherwise the graph or Completion Condition requires a separate valid change, or the Work Item must be explicitly cancelled. | Task cancellation does not imply abandonment of the parent outcome. Parent completion remains governed by the parent Completion Condition. | Work Item and Task authoring | `candidate` |
| D-005 | Cancelled prerequisite effect | `decided` | D-001 | A cancelled prerequisite does not satisfy `depends_on`. Any directly dependent Task that is `not_started` moves to `blocked`; an already `blocked` dependent remains `blocked`. No automatic cancellation propagates through Task dependencies. | Dependency failure may be repaired by changing the graph, so automatic cancellation would discard a potentially recoverable Task. | Task authoring | `candidate` |
| D-006 | Propagation coherence | `decided` | D-003 | A `cancelled` Work Item may own only Tasks in `done` or `cancelled`. The cancellation transition must propagate to every owned `not_started`, `in_progress`, or `blocked` Task before the resulting state is valid. | PRODUCT-REQ-SPEC-009 requires complete propagation across every owned unfinished Task while preserving `done` Tasks. | Work Item and Task authoring | `candidate` |
| D-007 | Cancellation Evidence | `decided` | D-001, D-002 | Use the existing `## Evidence` section. A cancelled Task records the intentional-stop reason, the decision or change that made the Task unnecessary, and any dependent Tasks moved to `blocked`. A cancelled Work Item records the unfinished Goal disposition, propagation results, preserved `done` Tasks, and Tasks changed to `cancelled`. Do not add cancellation-specific metadata or sections. | Existing Evidence sections already own lifecycle reasons and execution results. The cancellation contract needs substantive content, not another writable state surface. | Work Item and Task authoring | `candidate` |
| D-008 | Cancelled child Work Item effect | `decided` | D-001 | When the referenced child Work Item becomes `cancelled`, its `work_item_execution` Task becomes `cancelled` and records the child cancellation in Evidence. Dependent Tasks are then handled by D-005. The parent Work Item status does not change automatically. | The execution Task's only owned outcome becomes unattainable when the referenced child is intentionally terminated. `blocked` would incorrectly imply recoverability of the same child completion boundary. | Task and Work Item authoring | `candidate` |
| D-009 | Terminality and later mutation | `decided` | D-001, D-002 | `cancelled` is irreversible. Only meaning-preserving editorial correction, broken-reference repair, factual Evidence correction, and mechanical relation repair are allowed. Resuming work requires a new Task or Work Item that cites the cancelled record as a source when materially relevant. | Cancelled work usually changes scope or procedure. A new record preserves the failed execution history and gives the revised work its own completion judgment. | Work Item and Task authoring | `candidate` |
| D-010 | Explicit exclusions | `decided` | none | Perform no existing-record migration, descendant Work Item cancellation, framing design, or concrete implementation design. | PRODUCT-REQ-SPEC-009 fixes these exclusions. | Work Item boundary | `not_required` |
| D-011 | Provisional canonical boundary | `decided` | none | PRODUCT owns the semantic lifecycle in Work Item and Task authoring authority. T02 determines additional direct projections and ADR coverage. | PRODUCT-REQ-SPEC-009 assigns canonical meaning and propagation to PRODUCT while deferring mechanics. | Work Item and Task authoring | `unknown` |

### Current cursor

- Decision: none
- Loop state: `decision_complete`
- At most one decision is `in_discussion`.

## Done condition

- D-001 through D-011 are `decided`, `deferred`, or validly `blocked`.
- Direct Work Item and Task cancellation semantics are explicit.
- Parent, dependency, propagation, Evidence, and `work_item_execution` effects are explicit.
- Terminality and post-cancellation boundaries are explicit.
- Provisional canonical targets are sufficient for the mandatory T02 Investigation.
- No canonical artifact is authored.

## Verification

- Confirm every user answer is persisted before advancing the cursor.
- Confirm at most one decision is `in_discussion`.
- Confirm PRODUCT-REQ-SPEC-009 fixed outcomes are not reopened.
- Confirm unresolved impact facts remain assigned to T02.
- Confirm no authoring, graph change, review, synchronization, implementation, stage, or commit occurs.

## Evidence

- PRODUCT-REQ-SPEC-009 fixes the status name, terminal meaning, Work Item propagation set, preservation of `done` Tasks, and excluded scope.
- Current Work Item and Task authoring authority has no cancellation status or transition rules.
- The user decided D-001 on 2026-07-03: an unfinished Task may be cancelled independently.
- The user decided D-004 on 2026-07-03: independent Task cancellation does not change the parent Work Item status.
- The user decided D-005 on 2026-07-03: a cancelled prerequisite blocks dependents but does not cancel them.
- D-006 was resolved from PRODUCT-REQ-SPEC-009: a cancelled Work Item may own only `done` or `cancelled` Tasks.
- The user decided D-007 on 2026-07-03: cancellation Evidence uses the existing `## Evidence` section with substantive reason and propagation details.
- The user decided D-008 on 2026-07-03: cancelling a child Work Item cancels its `work_item_execution` Task and blocks dependents through D-005.
- The user decided D-009 on 2026-07-03: cancelled records are irreversible, and materially resumed work uses a new Task or Work Item.
- D-001 through D-011 are terminal.
- Result: `PASS`.
- DRMCP is non-operational. Filesystem authoring is the required fallback.
