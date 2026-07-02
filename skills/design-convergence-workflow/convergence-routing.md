# Convergence routing

## Purpose

Decide how accepted decisions reconcile with originating Requirements, Work Items, Tasks, ADRs, Specifications, and current workflow state.

This phase selects the semantic disposition.
It does not amend the execution graph or author canonical artifacts.

## Inputs

Read:

- parent Work Item;
- completed decision Task;
- completed impact Investigation;
- directly governing Requirements, ADRs, and Specifications;
- relevant authoring standards;
- current Task graph when graph drift is a candidate.

Treat accepted decisions as fixed inputs unless direct authority conflict requires reconsideration.

## Mismatch classes

Classify each decision-to-originating-artifact relationship as:

| class | meaning | required route |
|---|---|---|
| `consistent_refinement` | The decision concretizes the source without changing its meaning or boundary. | No originating-artifact amendment. Continue to downstream routing. |
| `stale_representation` | The intended disposition is already decided, but wording, references, or scope representation is stale. | Author the exact correction through an `authoring` Task. |
| `semantic_conflict` | The decision and source disagree materially, and the correct disposition remains unresolved. | Return to a `decision` owner. |
| `workflow_graph_drift` | Current Tasks, dependencies, owners, writer order, review order, release conditions, or Work Item boundary no longer support the accepted route. | Use `coordination`. |

`semantic_conflict` and `workflow_graph_drift` may occur together.

## Reconciliation decision ownership

Use a `decision` Task when judgment remains about:

- Requirement preservation, amendment, split, replacement, or follow-up;
- Work Item continuation, amendment, split, or creation;
- prior-decision preservation, reconsideration, or supersession;
- ADR and Specification change targets;
- whether graph amendment is required;
- shared-writer ordering policy;
- whether an adjacent topic is independently completable.

The decision Task does not:

- add or remove Tasks;
- change dependencies or blockers;
- assign writers;
- author ADRs, Specifications, Requirements, or Work Items;
- perform review or synchronization.

## Requirement identity boundary

Amend an existing Requirement when:

- the motivating problem and purpose remain the same;
- the change clarifies or adjusts a boundary within the same Required Outcome;
- the current Work Item completion meaning remains coherent.

Create a new Requirement when:

- the motivating problem changes;
- the Required Outcome changes materially;
- material scope can be accepted or rejected independently.

Create a follow-up Requirement when the original remains completable and an adjacent need appears.

Reconsider the decision when the Requirement remains valid and the current decision is the conflicting element.

## Work Item identity boundary

Continue the current Work Item when:

- it still resolves the same direct material sources;
- its Goal and Completion Conditions retain their meaning;
- only Tasks, dependencies, or ordering change inside the same delivery boundary.

Create or split another Work Item when:

- a new Requirement is introduced;
- the new scope has an independent completion judgment;
- ownership, release timing, or primary deliverables separate;
- an adjacent topic can proceed independently after current closure.

The reconciliation decision selects the Work Item disposition.
A later `work_item_decomposition` Task creates or splits the child Work Item.
A later `work_item_execution` Task represents the existing child in the parent graph when parent completion must wait for it.
Task count alone does not require a new Work Item.

## Existing Task amendment boundary

An existing incomplete Task may be amended when:

- `task_type` remains unchanged;
- owned outcome and completion judgment remain the same;
- only target artifacts, dependencies, or bounded procedure expand;
- one primary responsibility remains.

Create a separate Task when:

- a different `task_type` is required;
- a distinct outcome or completion judgment appears;
- independence must be preserved;
- the existing Task would own several responsibilities.

Do not substantively amend a completed decision Task.
A changed selected option, rationale, responsibility boundary, scope, or canonical target requires a new decision Task.

## Return routes

### No inconsistency

When the decision is a `consistent_refinement` and the graph is sufficient, add no reconciliation or coordination Task.
Proceed to ADR routing or canonical authoring.

### Decided stale representation

When the artifact disposition is already fixed:

- use an existing `authoring` Task when it owns the target;
- otherwise use coordination to add only the missing authoring route.

Do not create another decision Task merely because text is stale.

### Unresolved semantic conflict

Stop downstream progression.

- Return to an active incomplete decision Task when it owns the choice.
- When the earlier decision Task is complete, use coordination to create a new decision Task.
- Block affected downstream work until the conflict is decided.

### Graph drift

Hand exact Task-graph requirements to `graph-coordination.md`.
Do not change the graph in the reconciliation decision Task.

### Accepted Work Item creation or split

Hand the fixed identity and completion boundary to `work-item-decomposition.md`.
Do not create or split the Work Item in the reconciliation or coordination Task.

### Accepted child Work Item execution relation

Hand one already-created child Work Item to `work-item-execution.md` when the parent graph must wait for child completion.
Use one `work_item_execution` Task and one scalar `work_item_ref`.
Do not duplicate child Tasks, procedures, or completion Evidence in the parent Task.
Use coordination when the execution Task, dependency, blocker, or release route must be added or changed.

## Shared-writer disposition

Decide whether several writes:

- are independent by artifact or section;
- require deterministic serialization;
- require a later writer to preserve earlier accepted semantics;
- reveal a semantic conflict that must return to decision.

Coordination persists the final writer order.

## Completion condition

Routing is complete when:

- every investigated mismatch has one classification;
- Requirement and Work Item identity dispositions are fixed;
- every required new judgment has an owner;
- every required graph change is stated without being performed here;
- every canonical authoring target is known or blocked for a named reason;
- shared-writer policy is decided;
- adjacent independent scope is split or explicitly deferred.

## Stop conditions

Stop when:

- accepted authority conflicts and precedence cannot resolve it;
- the user has not selected among materially different dispositions;
- an existing completed decision would need substantive rewriting;
- the correct Requirement or Work Item identity cannot be determined;
- artifact ownership is ambiguous;
- the Investigation lacks evidence required for disposition.

Do not let authoring, coordination, correction, or synchronization absorb unresolved reconciliation judgment.
