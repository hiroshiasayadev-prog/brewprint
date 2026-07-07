# PRODUCT-TASK-SPEC-025-01: Decide vocabulary-collection route and lightweight investigation exception

- **id**: PRODUCT-TASK-SPEC-025-01
- **status**: done
- **date**: 2026-07-03
- **work_item**: PRODUCT-WORK-SPEC-025
- **task_type**: decision
- **estimate**: 0.5d
- **depends_on**: []
- **outputs**:
  - PRODUCT-REQ-SPEC-011
  - PRODUCT-TASK-SPEC-025-01
  - PRODUCT-TASK-SPEC-025-02
  - PRODUCT-TASK-SPEC-025-03
  - PRODUCT-TASK-SPEC-025-04
  - PRODUCT-WORK-SPEC-025

## Goal

Fix one bounded MVP contract for the TRV-pivot vocabulary-collection work and its exact 3-Task downstream route.

## Work

- Align the Problem and Desired Outcome against the TRV-INV-SPEC-004/005 findings.
- Decide source disposition and app-namespace ownership.
- Decide unknown handling.
- Decide the Task-graph shape, including per-owner Task separation between chappy's extraction and Claude's review.
- Resolve the conflict between the `investigation` task_type contract (which otherwise requires a full Investigation record) and the rejected heavyweight formal-Investigation route.
- Decide and record the accepted Investigation-Task lightweight Evidence exception wording.
- Decide the Completion Condition.
- Directly materialize the uniquely required T02, T03, and T04 Tasks.
- Update the parent Work Item Task list and flow.

This Task does not author `spec:product.design_records.authoring_standards.task_authoring` content or `skills/task-boundary-vocabulary/` content.

The direct T02/T03/T04 materialization is the accepted framing bootstrap exception: the accepted decision uniquely fixed each Task's type, primary outcome, and dependency route.

### Decision ledger

| ID | topic | status | decision summary | reason | canonical target | ADR route |
|---|---|---|---|---|---|---|
| D-001 | Problem / Desired Outcome | `decided` | Populate `skills/task-boundary-vocabulary/` from real corpus evidence, beyond `decision.md`. | TRV-INV-SPEC-004/005 found the checklist lacks concrete violation examples, not that models lack capability. | PRODUCT-REQ-SPEC-011 | `not_required` |
| D-002 | Source disposition | `decided` | `proceed`. | User confirmed. | PRODUCT-REQ-SPEC-011 | `not_required` |
| D-003 | Unknown handling | `decided` | No unresolved fact; Investigation not required for the Requirement itself. | TRV-INV-SPEC-004/005 already conclusive. | Framing decision | `not_required` |
| D-004 | App namespace | `decided` | Use `PRODUCT`, not `TRV`. | TRV application work is deferred; vocabulary collection is a PRODUCT-level authoring concern. | PRODUCT-REQ-SPEC-011, PRODUCT-WORK-SPEC-025 | `not_required` |
| D-005 | Task graph shape | `decided` | Split chappy's extraction and Claude's review into separate Tasks by owner. | User wants chappy's raw extraction independently tracked and verifiable, not folded into one combined authoring Task. | PRODUCT-WORK-SPEC-025 Task flow | `not_required` |
| D-006 | Investigation-Task lightweight exception | `decided` | Add an exception to `spec:product.design_records.authoring_standards.task_authoring` permitting an `investigation` Task to record results directly in its own Evidence, under explicit user judgment, closed by at least one downstream conclusion (including "no further action required"). | A full Investigation record is disproportionate for simple log extraction and reproduces the heavyweight route already rejected for this work. Relabeling the Task as `authoring` was rejected because it would require an artifact outside the Task's own scope. | `spec:product.design_records.authoring_standards.task_authoring` | `not_required` |
| D-007 | Task A (chappy) type | `decided` | `investigation`, using the D-006 exception. | Keeps chappy's role semantically accurate (research, not canonical authoring) without formal-record overhead. | PRODUCT-TASK-SPEC-025-03 | `not_required` |
| D-008 | Task B (Claude) type | `decided` | `authoring`. | Writing confirmed entries into `skills/task-boundary-vocabulary/` is canonical artifact authoring from already-decided inputs. | PRODUCT-TASK-SPEC-025-04 | `not_required` |
| D-009 | Completion Condition | `decided` | The Work Item is `done` when T01 through T04 are `done`. Canonical-term reconciliation is excluded and deferred to a later Work Item after roughly 30 log entries accumulate. | User confirmed. | PRODUCT-WORK-SPEC-025 | `not_required` |
| D-010 | Downstream route | `decided` | Materialize T02 (authoring, exception clause), T03 (investigation, chappy extraction), T04 (authoring, Claude review) in a linear dependency chain. | The accepted decisions uniquely fix each Task's type, outcome, and dependency; no separate graph judgment remains. | PRODUCT-WORK-SPEC-025 | `not_required` |

### Current cursor

- Decision: none
- Loop state: `decision_complete`
- Every owned decision is terminal.

## Done condition

- D-001 through D-010 are `decided`.
- The Problem, Desired Outcome, source disposition, Boundary, and proceed contract are fixed.
- The Investigation-Task lightweight Evidence exception is fixed.
- The exact T02, T03, T04 route is materialized.
- No `task-authoring.md` or `skills/task-boundary-vocabulary/` content is authored by this decision Task itself.

## Verification

- Confirm every user judgment is represented once in the ledger.
- Confirm no unresolved decision remains.
- Confirm T02, T03, and T04 each have one primary outcome and the correct dependency chain.
- Confirm this decision Task did not itself author `task-authoring.md` content or vocabulary entries.

## Evidence

- The user confirmed the Problem and Desired Outcome (D-001).
- The user confirmed `proceed` disposition (D-002).
- The user confirmed unknown handling required no Investigation (D-003).
- The user redirected the app namespace from TRV to PRODUCT (D-004).
- The user required a per-owner Task split between chappy and Claude (D-005).
- The user identified the conflict between the `investigation` task_type contract and the rejected heavyweight formal-Investigation route, and rejected relabeling the Task as `authoring` (D-006).
- The user proposed and iteratively refined the Investigation-Task lightweight Evidence exception wording; the final wording was accepted (D-006).
- The user confirmed the 3-additional-Task structure and each Task's type (D-007, D-008, D-010).
- The user confirmed the Completion Condition (D-009).
- DRMCP is non-operational. Filesystem authoring is the required fallback.
