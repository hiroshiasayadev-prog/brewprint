# PRODUCT-TASK-SPEC-023-04: Decide post-Investigation cancellation contract

- **id**: PRODUCT-TASK-SPEC-023-04
- **status**: done
- **date**: 2026-07-03
- **work_item**: PRODUCT-WORK-SPEC-023
- **task_type**: decision
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-023-02
- **outputs**:
  - PRODUCT-TASK-SPEC-023-04

## Goal

Resolve the exact cancellation contract gaps identified by PRODUCT-INV-SPEC-010.

## Work

- Decide substantive body requirements for cancelled Tasks.
- Decide substantive body requirements for cancelled Work Items.
- Decide the semantic owner and writable boundary for cancellation propagation.
- Prevent a cancellation executor from being cancelled as part of its own target operation.
- Decide post-Evidence responsibility-validator invocation for cancelled Tasks.
- Decide the violation disposition needed before cancellation finalization.
- Preserve all terminal decisions from T01.
- Record concise reasons, canonical targets, and provisional ADR routes.

This Task must not:

- reopen T01 decisions;
- perform further impact Investigation;
- change the Task graph;
- author ADR, Specification, skill, checklist, or implementation content;
- perform review, correction, synchronization, implementation, stage, or commit work.

### Decision ledger

| ID | topic | status | depends on | decision summary | reason | canonical target | ADR route |
|---|---|---|---|---|---|---|---|
| J-001 | Cancelled Task substantive sections | `decided` | none | A cancelled Task requires substantive `## Goal`, `## Done condition`, and `## Evidence`. `## Work`, `## Verification`, and `## Implementation contract` are not required solely because of cancellation. Existing substantive content remains historical and must not be erased. | Cancellation must preserve the intended outcome, the unmet completion boundary, and the intentional-stop Evidence without pretending that planned work or verification necessarily existed. | Task authoring | `candidate` |
| J-002 | Cancelled Work Item substantive sections | `decided` | none | A cancelled Work Item requires substantive `## Goal`, `## Boundary`, `## Completion Condition`, and `## Evidence`. `## Impact Scope`, `## Task flow`, and `## Task Candidates` are not required solely because of cancellation. Existing substantive content remains historical and must not be erased. | Cancellation must preserve the intended resolution, ownership boundary, unmet completion boundary, and propagation Evidence without requiring planning sections that may never have been completed. | Work Item authoring | `candidate` |
| J-003 | Cancellation propagation owner and write boundary | `decided` | J-001, J-002 | Cancellation is one atomic lifecycle operation, not a Task responsibility. Cancelling a Task updates the target Task and directly dependent Tasks. Cancelling a Work Item updates the Work Item, every owned unfinished Task, directly dependent Tasks outside the target set, any referencing `work_item_execution` Task, and direct dependents of that execution Task. The affected set is fixed before writing, and the operation is all-or-nothing. | A Task inside the target Work Item would cancel itself during propagation. An external atomic lifecycle operation avoids that paradox and prevents partially cancelled graph state. Concrete transaction or command mechanics remain outside W023. | Work Item and Task authoring; workflow support | `candidate` |
| J-004 | Post-cancellation-Evidence validator invocation | `decided` | J-001 | Do not invoke the post-Evidence semantic responsibility validator when a Task becomes `cancelled`. Keep the second invocation only before `done`. Cancellation readiness is governed by the cancelled-state section and Evidence contract. | A cancelled Task has no completed outcome to validate, may legitimately lack substantive Work or Verification, and can be cancelled through Work Item propagation. Requiring human exception handling during propagation would add no useful responsibility judgment. | Task authoring; responsibility-boundary validator | `candidate` |
| J-005 | Validator violation disposition before cancellation | `decided` | J-004 | Not applicable. Cancellation does not invoke the post-Evidence validator, so no validator-violation disposition exists in the cancellation route. | J-004 removes the prerequisite event that could produce this judgment. | Task authoring; responsibility-boundary validator | `not_required` |

### Current cursor

- Decision: none
- Loop state: `decision_complete`
- At most one decision is `in_discussion`.

## Done condition

- J-001 through J-005 are `decided`, `deferred`, or validly `blocked`.
- Cancelled-state body readiness is explicit for both record kinds.
- Cancellation propagation has one coherent owner and no self-cancellation paradox.
- Validator invocation and violation disposition are explicit.
- Canonical targets are sufficient for ADR routing and authoring graph coordination.
- No canonical artifact is authored.

## Verification

- Confirm every explicit user answer is persisted before advancing.
- Confirm at most one decision is `in_discussion`.
- Confirm T01 decisions remain unchanged.
- Confirm PRODUCT-INV-SPEC-010 findings J-001 through J-005 are fully dispositioned.
- Confirm no graph change, canonical authoring, review, synchronization, implementation, stage, or commit occurs.

## Evidence

- PRODUCT-INV-SPEC-010 concluded that canonical authoring is not ready.
- The missing judgments concern body readiness, propagation ownership, and validator workflow use.
- Exact checklist criteria require no direct change under the current responsibility-only contract.
- The user decided J-001 on 2026-07-03: cancelled Tasks require substantive Goal, Done condition, and Evidence only.
- The user decided J-002 on 2026-07-03: cancelled Work Items require substantive Goal, Boundary, Completion Condition, and Evidence.
- The user decided J-003 on 2026-07-03: cancellation is an external atomic lifecycle operation with a precomputed affected set and all-or-nothing writes.
- The user decided J-004 on 2026-07-03: cancelled Tasks do not receive the post-Evidence responsibility-validator invocation.
- J-005 is not applicable because no cancellation-time validator result exists.
- J-001 through J-005 are terminal.
- Result: `PASS`.
- DRMCP is non-operational. Filesystem authoring is the required fallback.
