# Interactive framing loop

## Purpose

Run one-question-at-a-time framing through one active `decision` Task.

The Task owns workflow judgment and resume state.
It does not own canonical Requirement authoring, formal Investigation authoring, downstream Work Item creation, or implementation.

## Startup

Read:

- the framing Work Item;
- the source Requirement;
- exact accepted authority needed for the current question;
- `framing-routing.md`;
- canonical Requirement, Work Item, and Task authoring standards.

Resolve repository facts before asking the user.
Do not ask the user to restate text already available in the Requirement.

## Decision inventory

Start with these common items:

| item | required judgment |
|---|---|
| Desired Outcome | What result the user intends from the source Requirement. |
| Outcome alignment | Whether Desired Outcome matches Required Outcome. |
| Source disposition | How the Requirement will be handled. |
| Unknown handling | Whether uncertainty is resolved now, investigated, carried as limited research, deferred, or blocked. |
| Downstream route | What responsibility follows framing. |

Add these items only when the disposition is `proceed`:

| item | required judgment |
|---|---|
| Goal | The downstream Work Item outcome. |
| Boundary | Owned and excluded responsibility. |
| Completion Condition | Observable completion meaning. |
| Direct source | Material Requirement provenance. |
| Initial route | Design convergence, decomposition, implementation planning, mechanical authoring, or another explicit route. |

Do not ask proceed-only questions after rejection, deferral, or another terminal non-proceed disposition.

## One-question rule

Ask exactly one unresolved judgment per user turn.

Before asking:

1. read the latest Task and Requirement state;
2. resolve repository-answerable facts;
3. identify the earliest unblocked decision;
4. present the options and material consequences;
5. ask only for the user-owned judgment.

Do not batch independent decisions.
Do not repeat a question already answered in chat or persisted Evidence.

## Persistence rule

Persist each accepted answer before advancing the cursor.

For each decision item record:

- item ID;
- status;
- selected disposition or answer;
- concise reason;
- dependency;
- canonical target or downstream owner;
- required Task materialization, when any.

At most one item may be `in_discussion`.

Use these item states:

- `open`;
- `in_discussion`;
- `decided`;
- `deferred`;
- `blocked`.

The decision Task completes when every owned item is `decided`, `deferred`, or validly `blocked`.

## Outcome alignment

For an existing Requirement:

- read the Problem and Required Outcome;
- treat the Problem as accepted unless a material conflict appears;
- state the interpreted Required Outcome concisely;
- ask whether the Desired Outcome matches it;
- route any mismatch before defining a downstream Work Item.

Do not silently rewrite the Requirement to match a proposed Work Item.

## Conditional Investigation

When the current framing judgment needs durable research:

1. decide the bounded research question;
2. directly materialize one `investigation` Task when the exception conditions hold;
3. set the current framing item to `blocked` on that Investigation;
4. keep the framing decision Task incomplete;
5. resume the same decision Task after Investigation completion.

A small repository read does not require an Investigation Task.
Limited research that does not alter Goal or completion meaning may be delegated to the downstream Work Item.

## Conditional Task materialization

The active framing decision may create a same-Work-Item Task only when the accepted answer uniquely fixes:

- one Task type;
- one primary outcome;
- one completion judgment;
- exact dependencies;
- exact output or bounded output family.

In the same persisted step:

- create the Task;
- add it to the parent Work Item `tasks` list;
- update Task flow or candidates;
- record the decision item that required it.

Use `coordination` instead when several valid graphs remain or graph repair is independently judged.
Use `work_item_decomposition` to create a downstream Work Item.

## Resume rule

On resume:

- read the persisted cursor;
- read outputs from Tasks that blocked the decision;
- verify that the current item still owns the judgment;
- continue from the earliest unblocked item.

Do not reconstruct state from chat when repository state exists.
Do not reopen a completed decision Task.

## Handoff

A handoff includes:

- framing Work Item ID;
- active decision Task ID;
- source Requirement ID;
- current item and state;
- decisions already persisted;
- exact blocker or next question;
- materialized Tasks and dependencies;
- prohibited scope.

## Completion

Before marking the decision Task `done`, confirm:

- every owned item is terminal;
- one source disposition is explicit;
- the proceed contract is complete when required;
- unknown handling is explicit;
- the downstream route is explicit;
- every required Task is materialized;
- no speculative Task was created;
- no canonical artifact authoring was absorbed.
