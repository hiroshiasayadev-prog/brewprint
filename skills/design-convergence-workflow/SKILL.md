# Design convergence workflow

## Purpose

This skill governs repository-persistent design work from an accepted framing handoff through reviewed design closure.

It prevents:

- design choices remaining only in chat;
- implementation planning starting before design closure;
- unresolved conflicts being hidden inside authoring;
- graph changes being performed without a coordination owner;
- ADRs and Specifications diverging;
- authors approving their own work;
- completed workflow records being rewritten to represent later progress;
- closure synchronization absorbing correction or design judgment.

This skill coordinates the workflow. Repository authoring standards remain authoritative for each artifact kind.

## Use this skill when

Use this workflow when repository work requires one or more of the following:

- design decision inventory inside an accepted Work Item boundary;
- decision inventory or interactive user judgment;
- impact and conflict investigation;
- Requirement, Work Item, or Task reconciliation;
- execution-graph amendment;
- ADR routing or ADR authoring;
- Specification or originating-artifact synchronization;
- integrated independent design review;
- finding correction and independent closure review;
- lifecycle, Evidence, and relation synchronization.

Resume the same workflow when later authoring, review, implementation planning, or closure exposes a missing design judgment.

## Do not use this skill when

Do not start this workflow for:

- an explanation or recommendation that requires no repository persistence;
- mechanical execution of an already accepted executor-ready implementation Task;
- one bounded editorial correction with no unresolved judgment, graph change, or review route;
- independent review whose complete contract already exists;
- mechanical synchronization whose exact inputs and writable boundary are already fixed.

When those activities reveal unresolved design judgment, return to this workflow.

## Required companions

Read the companion for every active phase:

| phase | companion |
|---|---|
| decision inventory, questioning, persistence, resume | `interactive-decision-loop.md` |
| impact and conflict discovery | `impact-investigation.md` |
| mismatch classification and artifact disposition | `convergence-routing.md` |
| Task graph, dependency, blocker, writer, or review-order changes | `graph-coordination.md` |
| parent-to-child Work Item creation or split | `work-item-decomposition.md` |
| parent-graph tracking of one already-created child Work Item | `work-item-execution.md` |
| ADR classification and coherent ADR boundaries | `adr-routing.md` |
| ADR, Specification, or originating-artifact writing | `design-authoring.md` |
| integrated review, finding routing, correction, and closure review | `design-review-gate.md` |
| lifecycle, Evidence, relation, and Work Item closure | `closure-synchronization.md` |

Re-read the relevant companion when responsibility, target artifact, writer, dependency, verdict, or lifecycle owner changes.

## End-to-end boundary

The workflow starts only after Work Item framing selects design convergence.
It starts from an accepted Goal, Boundary, Completion Condition, direct source, unknown-handling decision, and initial route.

The workflow completes only after:

- required decisions are terminal for the Work Item;
- required ADR and canonical authoring is complete;
- one integrated independent review returns `PASS`, or every required finding is independently closed;
- originating artifacts, graph state, lifecycle, Evidence, and relations are synchronized.

Production implementation is outside this workflow.
Implementation planning begins after design closure.

## Canonical ownership

| concern | owner |
|---|---|
| design topic, completion boundary, and Task graph overview | Work Item |
| decision inventory, cursor, explicit answers, and concise workflow Evidence | `decision` Task |
| bounded research question, facts, uncertainty, conflicts, and candidates | Investigation record owned by an `investigation` Task |
| Task graph, dependency, blocker, writer order, review order, or release change | `coordination` Task |
| parent-to-child Work Item creation or split | `work_item_decomposition` Task |
| one already-created child Work Item represented in the parent graph | `work_item_execution` Task |
| durable choice, alternatives, rationale, consequences, and supersession | ADR |
| current normative behavior, structure, boundary, and constraint | Specification |
| one bounded decided artifact set | `authoring` Task |
| independent verdict and finding set | `review` Task |
| repair of named findings | `correction` Task and affected artifacts |
| independent finding disposition | later `review` Task |
| mechanically derived lifecycle, Evidence, relation, and closure propagation | `synchronization` Task |

A Task ledger is workflow Evidence. It is not canonical design state.

## Responsibility architecture

Use these responsibility units:

1. decision inventory and interactive decision loop;
2. conditional decision-impact and conflict investigation;
3. conflict resolution and originating-artifact reconciliation decision;
4. conditional execution-graph amendment;
5. conditional Work Item decomposition;
6. conditional Work Item execution tracking;
7. ADR routing and ADR-boundary partitioning;
8. conditional ADR authoring;
9. Specification and originating-artifact authoring;
10. mandatory integrated independent review;
11. named-finding correction;
12. independent finding-closure review;
13. lifecycle, Evidence, and relation synchronization.

Integrated independent review is mandatory.
Investigation and the other units are conditional when their owned outcome is unnecessary.

Do not combine distinct primary outcomes or completion judgments merely because one session can perform several phases.

## Normal flow

```text
accepted framing handoff
  -> bounded Work Item and decision owner
  -> interactive decision loop
  -> impact and conflict Investigation when required
  -> reconciliation decision when needed
  -> graph coordination when needed
  -> Work Item decomposition when a child boundary is accepted
  -> Work Item execution tracking when the parent graph must wait for one child
  -> ADR routing and boundary partitioning
  -> ADR authoring when required
  -> Specification and originating-artifact authoring
  -> one integrated independent review
     -> PASS: closure synchronization
     -> NEEDS REVISION: finding-specific coordination
        -> correction
        -> independent finding-closure review
        -> closure synchronization
```

The order may return to an earlier responsibility when a later phase exposes unresolved judgment or graph drift.

## Pre-authoring convergence gate

Do not begin ADR or canonical authoring until:

- no unresolved decision remains for the authoring boundary;
- Requirement and Work Item disposition is fixed;
- every required Task exists;
- dependencies and shared-writer order are fixed;
- exact authoring targets are identified.

Treat this as a completion or release condition on existing owners.
Create a separate `verification` Task only when the gate requires independent ownership, machine-checkable aggregation, cross-Work-Item evaluation, or reusable release Evidence.

## Return-to-decision rule

Any phase that discovers unresolved design judgment must stop downstream progression.

- Return directly to an active incomplete decision Task when it already owns the choice and the graph remains valid.
- When the earlier decision Task is complete, use coordination to create a new decision Task.
- Use coordination first when ownership, dependency, blocker, writer order, review order, or release conditions must change.
- Use `work_item_decomposition` when an accepted decision creates or splits a child Work Item.

Do not reopen or substantively rewrite a completed decision Task.
Meaning-preserving editorial or broken-reference correction is allowed only when it does not change the decision or completion judgment.

## Originating-artifact mismatch classes

Use `convergence-routing.md` to classify:

- `consistent_refinement`;
- `stale_representation`;
- `semantic_conflict`;
- `workflow_graph_drift`.

The classification selects the responsibility route.
It does not prescribe one fixed Task sequence for every topic.

## Shared writers

Never allow concurrent writes to the same artifact or section.

Use coordination to persist deterministic writer order.
Each later writer must preserve accepted semantics from earlier writers.
A later writer that must weaken, remove, or reinterpret accepted semantics returns to investigation and decision.

Place integrated review after the final writer.

## Integrated review boundary

Use one final integrated review Task per Work Item after all required decisions, coordination, ADR authoring, originating-artifact authoring, Specification authoring, and shared-writer sequencing are complete.

Split another Work Item when a design boundary has an independent completion judgment.
Each Work Item then receives its own integrated review.

## Conditional finding route

Do not create speculative correction or finding-closure review Tasks in the initial graph.

- `PASS` proceeds directly to closure synchronization.
- `NEEDS REVISION` routes through coordination to create Tasks derived from named findings.

Correction repairs named findings.
A later independent review closes or keeps them open.
The correction author does not close their own findings.

## Completed-record preservation

Completed decision, authoring, and review Tasks remain historical Evidence.
Later work uses new Tasks rather than changing their outcomes.

Downstream ADR, Specification, review, and closure progress is not written into a completed decision Task.
A completed decision remains `decided`; there is no downstream `recorded` state.

## Stop conditions

Stop and report the exact route when:

- required authority or selected Investigation Evidence is missing;
- the user answer allows materially different interpretations;
- accepted authorities conflict and no disposition is decided;
- the canonical target cannot be identified safely;
- an accepted ADR requires a material reversal without supersession authority;
- the current graph lacks an owner or valid dependency route;
- an author would need to introduce new judgment;
- review prerequisites are incomplete;
- closure would need to repair content, alter a verdict, or change the graph.

Do not infer a design answer to avoid stopping.

## Completion checklist

Before declaring design closure, verify:

- the Work Item has one coherent resolution and completion boundary;
- all required decisions are `decided`, `deferred`, or validly `blocked`;
- every selected Investigation exists and is complete;
- originating-artifact disposition and graph route are fixed;
- every required child execution relation uses one `work_item_execution` Task and one `work_item_ref`;
- every decision has an ADR routing outcome;
- required ADRs exist at the required lifecycle state;
- current Specifications and originating artifacts reflect accepted design;
- shared writers completed in deterministic order;
- one integrated independent review evaluated the final combined state;
- required findings are independently closed;
- closure synchronization changed only its declared writable targets;
- implementation planning can proceed without hidden design judgment.
