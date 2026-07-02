# Graph coordination

## Purpose

Persist required execution-graph changes after the semantic route is decided.

Coordination changes ownership and ordering.
It does not choose design outcomes or author child deliverables.

## When coordination is required

Use a `coordination` Task when one or more of these must change:

- a required responsibility has no owning Task;
- a Task must be added, removed, split, or replaced;
- an incomplete Task needs a meaning-preserving contract amendment;
- dependencies or blockers must change;
- writer ownership or shared-writer order must change;
- review order or release conditions must change;
- downstream work must be blocked or released;
- routing to another Work Item must be represented in the Task graph;
- the parent graph must wait for one already-created child Work Item;
- a completed Task requires a new successor Task for revised work.

Do not add coordination merely because an inconsistency exists.
Use it only when the graph must change through an independently owned graph judgment.

An active framing `decision` Task may directly materialize a uniquely determined same-Work-Item Task under `skills/work-item-framing/SKILL.md`.
Do not add coordination only to repeat that already fixed materialization route.

## Inputs

Read:

- parent Work Item;
- completed decision and any selected Investigation records;
- reconciliation disposition;
- current child Task contracts;
- exact artifact and writer candidates;
- accepted Task and Work Item authoring standards.

Every graph change must trace to an accepted decision, named finding, or mechanically necessary owner repair.
The framing direct-materialization exception is not a general coordination replacement.

## Coordination outputs

Coordination may:

- create a missing Task with one primary responsibility;
- split a mixed-responsibility Task;
- amend an incomplete Task while preserving its type and completion judgment;
- add or remove dependencies and blockers;
- persist writer serialization;
- persist review and release order;
- route an accepted Work Item creation or split to a `work_item_decomposition` owner;
- create or release one `work_item_execution` Task when the parent graph must wait for one already-created child Work Item;
- materialize finding-specific correction and closure-review Tasks;
- identify exact next owners and release conditions.

Coordination must not:

- decide Requirement or Work Item identity without accepted input;
- select an ADR disposition without completed routing;
- write ADR, Specification, Requirement, skill, or implementation content;
- repair findings;
- issue an independent verdict;
- perform lifecycle closure.

## Task materialization rules

Create a Task only when it has:

- one `task_type`;
- one primary outcome;
- one completion judgment;
- one writer or independent owner;
- exact dependencies;
- exact outputs or a bounded output family;
- a clear release condition.

Extend an existing incomplete Task only when its type, outcome, and completion judgment remain unchanged.

Create a new Task when:

- a different type is required;
- a distinct outcome appears;
- independence must be preserved;
- a completed Task would otherwise be reopened;
- one Task would own several responsibilities.

## Completed-record rule

Do not reopen completed or cancelled Tasks to represent later work.

Completed decision, authoring, review, correction, and synchronization Tasks preserve their accepted outcome.
Cancelled Tasks preserve their intentional-stop outcome and require a new Task for materially resumed work.

Create successor Tasks that:

- identify the earlier completed owner;
- state why new work is required;
- depend on the relevant accepted Evidence;
- preserve earlier outcomes as historical inputs.

## Shared-writer serialization

When several Tasks write the same artifact or section:

1. prohibit concurrent writes;
2. record a deterministic writer order;
3. add explicit dependencies between writers;
4. require each later writer to read and preserve earlier accepted semantics;
5. place integrated review after the final writer;
6. return to investigation and decision if a later writer must weaken, remove, or reinterpret earlier semantics.

Separate writers by section only when the sections are semantically independent and no common acceptance judgment is required.

## Conditional finding route

Do not pre-create correction or finding-closure review Tasks.

After integrated review returns `NEEDS REVISION`:

1. read exact finding IDs and required outcomes;
2. group findings only when they share one repair owner and completion judgment;
3. create exact correction Tasks with bounded writable targets;
4. create later independent review Tasks for finding closure;
5. block closure synchronization until required findings are closed;
6. create a new decision route instead of correction when a finding requires judgment.

Do not create no-op Tasks for unused branches.

## Work Item decomposition handoff

Coordination does not create or split Work Items.

When an accepted decision requires a child Work Item:

- create or release one `work_item_decomposition` Task;
- preserve the accepted identity and completion boundary as fixed input;
- record required dependencies between the parent route and decomposition owner;
- keep child-internal deliverables outside the coordination Task.

The decomposition owner follows `work-item-decomposition.md`.

## Work Item execution handoff

Coordination may create or release a `work_item_execution` Task only when:

- the child Work Item already exists;
- the parent graph must represent the child as one execution unit;
- exactly one child Work Item ID is known;
- the Task can use `work_item_ref` without duplicating child internals.

Coordination fixes dependencies, blockers, and release order around the execution Task.
The execution Task owns no graph change.
The execution Task completes only after the referenced child Work Item is `done`.
When the referenced child Work Item becomes `cancelled`, the execution Task becomes `cancelled`, its direct dependents follow the cancelled-prerequisite rule, and the parent Work Item status remains unchanged.
The atomic lifecycle operation owns those status and Evidence writes; coordination does not perform cancellation.

Route child creation or split to `work_item_decomposition` before creating the execution relation.

## Pre-authoring release gate

Release ADR and canonical authoring only when:

- no unresolved decision remains for their scope;
- Requirement and Work Item disposition is fixed;
- every required owner exists;
- dependencies and blockers are correct;
- shared-writer order is fixed;
- exact authoring targets are known.

A separate verification Task is unnecessary unless the gate needs independent or reusable objective Evidence.

## Completion condition

Coordination is complete when:

- every required graph change is persisted;
- each Task has one responsibility and valid dependency route;
- writer and review order are deterministic;
- affected downstream work is correctly blocked or released;
- conditional branches remain abstract until triggered;
- no child-owned deliverable was authored by coordination.

## Verification

Inspect the Work Item and exact changed Task records.
Confirm:

- bidirectional Work Item–Task ownership;
- unique Task IDs and correct parent sequence;
- no missing dependencies;
- no circular writer order;
- no completed or cancelled Task was substantively rewritten;
- every `work_item_execution` Task references exactly one existing child Work Item;
- no speculative correction or review Task exists;
- downstream release conditions match the accepted route.

## Stop conditions

Stop when:

- the semantic disposition is unresolved;
- the required Task type or completion boundary is unclear;
- a graph change would exceed the current Work Item authority;
- shared-writer order requires a new design choice;
- a named finding does not define a repair outcome;
- the current Task graph differs materially from the Investigation Evidence.
