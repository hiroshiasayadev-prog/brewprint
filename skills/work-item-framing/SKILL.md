# Work Item framing workflow

## Purpose

Establish how one Requirement will be handled before downstream design or execution planning begins.

The workflow aligns the user's Desired Outcome with the Requirement's Required Outcome.
It concludes with one explicit source disposition.
When the disposition is `proceed`, it also fixes an actionable downstream Work Item contract.

Repository authoring standards remain authoritative for Requirement, Work Item, and Task shape.

## Use this skill when

Use this workflow before creating a repository-persistent Work Item for downstream design or execution.

Use it when source handling, Requirement identity, downstream Goal, Boundary, Completion Condition, unknown handling, or workflow route remains undecided.

## Do not use this skill when

Do not use this workflow to:

- migrate an existing Work Item;
- repair or retroactively legitimize an existing unframed Work Item;
- perform downstream design, implementation, or execution;
- define lifecycle semantics such as `cancelled`;
- perform an independent framing review.

Start a new framing route from the direct Requirement when future work still matters.

## Required companions

| responsibility | companion |
|---|---|
| one-question judgment, persistence, resume, and Task materialization | `interactive-framing-loop.md` |
| source disposition, unknown handling, and downstream route selection | `framing-routing.md` |

Read `skills/design-convergence-workflow/SKILL.md` only after framing selects design convergence.

## Entry contract

Repository-persistent framing starts from one Requirement.

| input state | required entry action |
|---|---|
| Existing Requirement | Read its Problem, Evidence, Required Outcome, and exclusions. Treat the Problem as accepted input unless a material conflict appears. |
| No Requirement | Identify the Problem and Desired Outcome sufficiently to capture a Requirement. Create the Requirement before starting the framing Work Item. |

Do not create a framing Work Item from an unpersisted raw topic.

## Initial graph and ownership

Create one framing Work Item whose Goal is to decide how the Requirement will be handled.

The initial Task graph contains exactly one `decision` Task.
Do not create speculative Investigation, authoring, decomposition, coordination, review, or synchronization Tasks.

The framing Work Item owns its Goal, Boundary, completion condition, and Task graph overview.

The framing `decision` Task owns:

- the decision inventory and cursor;
- explicit user answers;
- Desired Outcome and Required Outcome alignment;
- source disposition;
- conditional downstream Work Item contract;
- unknown handling;
- selected downstream route;
- direct materialization of uniquely determined same-Work-Item Tasks.

This workflow does not define a `framing` Task type.

## Normal flow

```text
Requirement
  -> framing Work Item
  -> one framing decision Task
     -> align outcomes
     -> decide source disposition
     -> fix downstream Work Item contract when proceeding
     -> select unknown and downstream routes
     -> materialize only required Tasks
  -> complete required framing Tasks
  -> direct framing Work Item closure
```

Return to the active incomplete decision Task after a conditional Investigation.
Do not reopen a completed decision Task.
A later changed judgment requires a new decision Task.

## Proceed contract

When the disposition is `proceed`, fix:

- Goal;
- Boundary;
- Completion Condition;
- direct material source;
- unknown handling;
- initial downstream route.

Fix only the detail required to establish Work Item identity and completion meaning.
Do not absorb downstream design into framing.

## Conditional Task materialization

The active framing `decision` Task may create and register an additional Task directly only when:

- the new Task belongs to the same framing Work Item;
- the selected decision uniquely determines its Task type, primary outcome, and dependency route;
- no separate graph judgment remains;
- no completed Task is substantively changed;
- no external Work Item graph is changed.

Update the parent Work Item Task list and flow in the same persisted decision step.
Do not create a `coordination` Task only to materialize a route already fixed by the active decision.

Use `coordination` for graph repair, alternative graph design, shared-writer order, cross-Work-Item change, or release-order judgment.
Use `work_item_decomposition` when the accepted route creates or splits a downstream Work Item.

## Investigation boundary

Formal Investigation is conditional.

Create an `investigation` Task only when one bounded research question needs durable Evidence, uncertainty, and options before framing can conclude.

A small repository read may remain inside framing.
Limited downstream research may remain in the accepted downstream route when it does not change Goal or completion meaning.
The user may also defer the topic.

## Review and closure

Framing has no mandatory independent review.
Direct user judgment is the framing acceptance boundary.

Do not create a synchronization Task for simple closure.
The framing workflow may set its Work Item to `done` when:

- the source disposition and reason are explicit;
- the required next route is persisted;
- a required proceed contract is complete;
- every materialized framing Task is complete;
- no required framing judgment remains hidden.

## Design-convergence handoff

Design convergence is one conditional downstream route.

Start `skills/design-convergence-workflow/` only when framing selects it.
Provide the accepted Goal, Boundary, Completion Condition, direct source, unknown handling, Investigation decision, and initial route.

## Completion checklist

Verify:

- one explicit source disposition exists;
- Desired Outcome and Required Outcome alignment is recorded;
- any Requirement identity action has an exact owner;
- the proceed contract exists when required;
- formal Investigation is present only when selected;
- no speculative Task exists;
- direct Task materialization stayed inside the bounded exception;
- no existing Work Item migration or repair occurred;
- the exact downstream route is persisted.
