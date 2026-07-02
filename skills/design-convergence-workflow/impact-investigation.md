# Impact investigation

## Purpose

Own one bounded Investigation record that determines how the accepted decision set affects repository authority, workflow structure, and shared writers.

This phase is conditional.
Use it when framing or a later design decision selects one bounded research question that needs durable Evidence, uncertainty, and options.

## Investigation Task contract

Use `task_type: investigation`.
The Task owns one formal Investigation record for one bounded research question.

The Investigation records:

- directly affected Requirements, ADRs, Specifications, Work Items, Tasks, skills, and implementation-facing contracts;
- contradictions and stale representations;
- unresolved semantic conflicts;
- execution-graph change candidates;
- shared-writer candidates;
- existing authority that already covers the decision;
- uncertainty and missing evidence;
- options that require later user judgment.

It does not adopt a design choice, amend the graph, author canonical content, or perform review.

## Read boundary

Before investigating:

1. Read the framing handoff, parent Work Item, and completed decision Task.
2. Read exact accepted authority named by the Work Item or decision targets.
3. Read existing ADRs that may cover or conflict with the decision.
4. Read only adjacent artifacts required to confirm a scoped relation or contradiction.

Do not perform repository-wide traversal when exact targets are known.

## Investigation questions

For every decided item, determine:

- Which canonical artifact currently owns the relevant rule?
- Does the decision refine, contradict, or leave that artifact unchanged?
- Does an accepted ADR already cover the durable choice?
- Does the current Work Item still have one coherent completion boundary?
- Does an existing Task own every required downstream responsibility?
- Must dependencies, blockers, writers, review order, or release conditions change?
- Will several Tasks write the same artifact or section?
- Is any apparent conflict actually stale wording with an already decided disposition?
- Does the decision expose an adjacent independently completable topic?

## Required output

The Investigation must provide a scoped inventory with at least:

| item | required content |
|---|---|
| affected artifact | Public ID or active `spec:` ref, plus section when needed. |
| observed state | Relevant current rule or absence. |
| relation to decision | refinement, stale representation, semantic conflict, or graph drift candidate. |
| decision IDs | Exact source decisions. |
| required judgment | None, reconciliation decision, ADR routing, or other named owner. |
| graph candidate | Missing Task, dependency, blocker, writer order, review order, or split candidate. |
| shared writer | Exact artifact or section and candidate writers. |
| evidence | File, section, or scoped diff supporting the result. |

Record uncertainty explicitly.
Do not turn a hypothesis into a finding without evidence.

## Conflict classification preparation

The Investigation proposes, but does not finally adopt, one or more of:

- `consistent_refinement`;
- `stale_representation`;
- `semantic_conflict`;
- `workflow_graph_drift`.

`convergence-routing.md` owns the final disposition when judgment is required.

A semantic conflict and graph drift may coexist.

## Graph-change candidates

Record a graph-change candidate when:

- no Task owns a required responsibility;
- an existing Task would own multiple primary outcomes;
- a completed Task would need substantive revision;
- a dependency or blocker no longer reflects the accepted route;
- shared writers lack deterministic order;
- review order or release conditions are incomplete;
- an independent topic needs another Work Item.

Do not create or amend Tasks in the Investigation phase.

## Shared-writer candidates

For each artifact or section with several candidate writers, record:

- exact target;
- candidate writer Tasks;
- semantic dependency between writers;
- whether the writes can be separated by section;
- required writer order candidate;
- risk if writes occur concurrently.

Coordination owns the final serialization.

## Completion condition

The Investigation is complete when:

- every decided item has a scoped impact result;
- every material conflict has evidence and a proposed classification;
- every missing owner or graph defect is recorded as a candidate;
- every shared writer is identified;
- every uncertainty has a named blocker or next owner;
- no design option is silently selected.

## Stop conditions

Stop and report the exact blocker when:

- required authority cannot be located within the scoped namespace;
- the intended decision cannot be reconstructed from the completed decision Task;
- an artifact relation is ambiguous and affects the route;
- repository state differs materially from the Work Item assumptions;
- investigation would require changing an artifact.

## Handoff

The next owner receives:

- the completed Investigation record;
- exact decision IDs;
- proposed mismatch classes;
- originating-artifact disposition candidates;
- graph-change and shared-writer candidates;
- unresolved evidence gaps.

The next phase is reconciliation decision when judgment remains, otherwise graph coordination or ADR routing as required.
