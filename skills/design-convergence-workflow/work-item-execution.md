# Work Item execution

## Purpose

Represent one already-created child Work Item as one execution unit in the parent Work Item Task graph.

This phase tracks child completion or cancellation without duplicating child-owned work.
It does not create the child Work Item or change the parent graph.

## Inputs

Read:

- the parent Work Item;
- the existing child Work Item;
- the `work_item_execution` Task;
- the accepted parent-level dependency and release route;
- Task and Work Item authoring standards.

Do not begin before the child Work Item exists and the parent graph names it through `work_item_ref`.

## Owned outcome

A `work_item_execution` Task owns one already-created child Work Item as one parent-graph execution unit.

The Task must:

- use exactly one scalar `work_item_ref`;
- reference a Work Item different from its parent `work_item`;
- preserve the child Work Item as owner of its internal Task graph and work;
- record the child status needed for its own completion Evidence;
- remain incomplete while the child is `not_started` or `in_progress`;
- use `blocked` when the child blocks the parent execution route and the Task contract permits it;
- become `done` only after the child Work Item is `done`;
- become `cancelled` when the child Work Item is `cancelled`, record that terminal child state, and leave the parent Work Item status unchanged.

The Task must not:

- create, split, or redefine the child Work Item;
- change Tasks, dependencies, blockers, owners, writer order, review order, or release order;
- duplicate child Tasks, procedures, deliverables, decisions, or review Evidence;
- perform child-owned implementation or authoring;
- issue an independent review verdict;
- repair findings;
- propagate unrelated lifecycle or relation state.

## Parent and child boundaries

The parent Work Item owns the execution Task and its place in the parent graph.

The child Work Item owns:

- its Goal and Boundary;
- its internal Task graph;
- its decisions and Investigation records;
- its deliverables and verification;
- its integrated review and closure Evidence.

A completed execution Task satisfies only its own Done condition.
The parent Work Item still evaluates its complete Completion Condition.

## Relation contract

| field | value |
|---|---|
| parent ownership | Task `work_item` and parent Work Item `tasks` |
| child execution target | Task scalar `work_item_ref` |
| cardinality | exactly one child Work Item |
| child reverse field | none |
| target state before Task completion | `done` |
| target state that cancels the execution Task | `cancelled` |

The child target must use the same app namespace and domain as the execution Task.
The child target must already exist.

## Completion condition

Work Item execution is complete when:

- `work_item_ref` identifies one existing child Work Item;
- the referenced child Work Item is `done`;
- Task Evidence records the referenced child and observed status;
- the Task duplicates no child-owned execution detail;
- the Task Done condition is satisfied.

## Cancellation outcome

When the referenced child Work Item becomes `cancelled`:

- the `work_item_execution` Task becomes `cancelled` through the atomic cancellation lifecycle operation;
- Task Evidence records the referenced child and observed `cancelled` status;
- every direct dependent Task follows the cancelled-prerequisite rule and becomes or remains `blocked`;
- the parent Work Item status does not change automatically;
- no child Work Item or transitive descendant is cancelled by this relation.

The execution Task does not perform or own the propagation operation.

## Verification

Confirm:

- the parent owns the execution Task;
- `work_item_ref` differs from the parent `work_item`;
- the referenced Work Item exists;
- the referenced Work Item is `done` before the Task is `done`;
- a referenced `cancelled` child produces a `cancelled` execution Task with Evidence and blocked direct dependents;
- child cancellation does not change the parent Work Item status;
- no child Task graph or deliverable is copied into the execution Task;
- no graph change, review, correction, or synchronization was absorbed.

## Stop conditions

Stop and route to `coordination` when:

- the execution Task, dependency, blocker, owner, or release route must change;
- the target child Work Item is missing;
- more than one child must be represented;
- the parent must choose among several child targets.

Route child creation or split to `work_item_decomposition`.
Route unresolved child identity or completion semantics to `decision`.
