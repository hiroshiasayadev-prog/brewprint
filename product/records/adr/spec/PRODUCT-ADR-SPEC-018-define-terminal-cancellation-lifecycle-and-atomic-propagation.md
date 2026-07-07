# PRODUCT-ADR-SPEC-018: Define terminal cancellation lifecycle and atomic propagation

- **status**: accepted
- **date**: 2026-07-03
- **depends_on**: []
- **supersedes**: []
- **migrated_to_spec**: null

## Context

Work Items and Tasks currently terminate successfully only through `done`.

That model cannot represent work that is intentionally abandoned before its completion condition is satisfied.
Using `blocked` for intentional termination incorrectly implies that resumption is still expected.
Deleting or rewriting the record would erase the execution history.

Cancellation also affects several connected records:

- a Work Item owns Tasks;
- Tasks may depend on other Tasks;
- a child Work Item may be represented by a `work_item_execution` Task;
- partial propagation would leave the graph in a contradictory state.

The lifecycle therefore needs a terminal unsuccessful outcome and one coherent propagation boundary.

## Decision

Add terminal `cancelled` status to Work Items and Tasks.

A Work Item or Task may move to `cancelled` only from:

- `not_started`;
- `in_progress`;
- `blocked`.

A `done` or already `cancelled` record cannot be cancelled.
No lifecycle transition leaves `cancelled`.

### Independent Task cancellation

A Task may be cancelled independently of its parent Work Item.

Independent Task cancellation does not change the parent Work Item status.
The parent may later become `done` only when its current Completion Condition remains satisfied without the cancelled Task.
Otherwise the graph or Completion Condition requires a separate valid change, or the Work Item must itself be cancelled.

A cancelled prerequisite does not satisfy `depends_on`.
Every directly dependent `not_started` Task becomes `blocked`.
An already `blocked` direct dependent remains `blocked` and records the dependency failure.
Cancellation does not automatically propagate through Task dependencies.

### Work Item cancellation

Cancelling a Work Item changes every owned Task in `not_started`, `in_progress`, or `blocked` to `cancelled`.
Owned `done` Tasks remain `done`.

A valid cancelled Work Item therefore owns only `done` or `cancelled` Tasks.

Work Item cancellation does not cancel child Work Items or transitive descendants.
It affects parent-graph execution records that directly represent the cancelled Work Item according to the accepted `work_item_execution` contract.

### Evidence

Use the existing `## Evidence` section.

A cancelled Task records:

- why the Task intentionally stopped before completion;
- the decision or change that made the Task unnecessary;
- directly dependent Tasks moved to or retained in `blocked`.

A cancelled Work Item records:

- why its Goal ended without completion;
- the owned Tasks preserved as `done`;
- the owned Tasks changed to `cancelled`;
- directly affected external dependency or execution records.

A status-only statement such as “cancelled” is insufficient Evidence.
No cancellation-specific metadata field or body section is added.

### Atomic lifecycle operation

Cancellation is one lifecycle operation, not a Task responsibility.

Before writing, the operation determines the complete affected set.
It then applies every required lifecycle and Evidence update atomically.
The result is either the complete valid cancelled state or no state change.
Partial best-effort cancellation is invalid.

For direct Task cancellation, the affected set includes:

- the target Task;
- every directly dependent Task whose state or Evidence must change.

For Work Item cancellation, the affected set includes:

- the target Work Item;
- every owned unfinished Task;
- directly dependent Tasks outside the owned cancellation set;
- every `work_item_execution` Task directly representing the target Work Item;
- direct dependents of those execution Tasks.

The concrete transaction, command, API, diagnostic, and implementation mechanism is outside this decision.

## Rationale

A distinct terminal status keeps intentional abandonment separate from temporary blocking and successful completion.

Independent Task cancellation is necessary because one execution unit may become unnecessary while the parent outcome remains valid.
Avoiding upward status propagation prevents one discarded Task from silently abandoning an entire Work Item.

Blocking rather than cancelling direct dependents preserves recoverable work.
A dependency can later be replaced or removed through an explicit graph change.

Work Item-to-Task propagation keeps the owned graph coherent.
Preserving completed Tasks retains valid historical results.

An external atomic lifecycle operation avoids a self-cancellation paradox in which a Task inside the target Work Item attempts to cancel its own execution container.
Atomicity also prevents mixed intermediate states that no authoring or review rule could interpret safely.

## Rejected alternatives

| alternative | rejection reason |
|---|---|
| Use `blocked` for intentionally abandoned work. | `blocked` means progress may resume after a dependency or decision is resolved. |
| Permit Task cancellation only through parent Work Item cancellation. | Individual Tasks may become unnecessary while the parent outcome remains valid. |
| Automatically cancel the parent when one Task is cancelled. | One Task does not necessarily determine the parent Completion Condition. |
| Automatically cancel every downstream dependent Task. | Direct dependents may remain recoverable after graph repair. |
| Cancel child Work Items and transitive descendants. | Child Work Item identity and ownership require separate explicit judgment. |
| Execute cancellation through an owned synchronization or coordination Task. | The executor could be included in the affected cancellation set and cancel itself. |
| Apply best-effort writes and repair partial state later. | Partial cancellation creates contradictory lifecycle and dependency state. |
| Add cancellation-specific metadata fields. | Existing Evidence sections already own lifecycle reasons and effects. |

## Consequences

- Work Item and Task authoring Specifications must add `cancelled`, eligible transitions, terminality, Evidence, dependency effects, and cancelled-state section readiness.
- Work Item cancellation must be validated as one coherent state across the Work Item and every owned Task.
- Direct dependents of cancelled Tasks become blocked rather than cancelled.
- `work_item_execution` authority must define the cancelled-child branch.
- Workflow support must distinguish successful `done` closure from unsuccessful cancellation.
- Concrete DRMCP transaction and command design requires separate downstream work.
- Existing-record migration is not required by this decision.

## Evidence

- PRODUCT-REQ-SPEC-009 requires terminal Work Item and Task cancellation and owned-Task propagation.
- PRODUCT-TASK-SPEC-023-01: D-001 through D-007.
- PRODUCT-INV-SPEC-010: canonical lifecycle, propagation-owner, and partial-state impact findings.
- PRODUCT-TASK-SPEC-023-04: J-003 atomic lifecycle operation.
- PRODUCT-TASK-SPEC-023-06: B-001 ADR routing boundary.
