# Interactive decision loop

## Purpose

Persist unresolved design decisions and explicit user answers in a resumable `decision` Task.

The loop owns workflow state only.
It does not own canonical ADR or Specification content.

## Decision Task contract

The Task must identify:

- parent Work Item;
- bounded topic and excluded scope;
- directly relevant authority;
- known canonical targets;
- decision inventory;
- dependency order;
- current cursor;
- status definitions;
- stop conditions;
- expected downstream route.

Use stable local IDs such as `D-001`.
Each entry must preserve enough information to resume without chat reconstruction.

Recommended fields:

| field | meaning |
|---|---|
| ID | Stable local decision ID. |
| Topic | The bounded judgment. |
| Status | `open`, `in_discussion`, `decided`, `blocked`, `deferred`, or `superseded`. |
| Depends on | Earlier decisions, Investigation, or authority. |
| Decision summary | Concise accepted outcome. |
| Reason | One to three sentences of workflow Evidence. |
| Canonical target | Expected ADR, Specification, Requirement, Work Item, or scope target. |
| ADR route | `candidate`, `unknown`, or a resolved downstream route. |

Detailed durable rationale belongs in an ADR when required.

## Inventory procedure

Before asking the first question:

1. Read only directly relevant accepted authority.
2. Separate repository-resolvable facts from user judgments.
3. Identify contradictions, gaps, ambiguous boundaries, and missing ownership.
4. Add every currently known judgment to the inventory.
5. Add dependency edges and named blockers.
6. Identify likely canonical targets without inventing IDs or paths.
7. Assign provisional ADR routing when useful.
8. Select the first unblocked `open` item and mark it `in_discussion`.

Newly revealed decisions may be appended when a confirmed answer exposes them.
Record which earlier decision revealed each new item.
Do not expand into an independently completable adjacent topic.

## One-question rule

Ask exactly one unresolved decision per user turn.

Tightly coupled criteria may remain one decision only when answering them separately would create an invalid intermediate state.

Do not ask the next question until the current answer is durably persisted, explicitly deferred, or validly blocked.

## Question shape

When alternatives require explanation, use:

```text
Decision: D-<nnn> — <topic>

Why this must be decided:
<brief consequence>

Recommendation:
<choice and decisive reason>

Alternatives:
- A: <meaning and consequence>
- B: <meaning and consequence>

Question:
<one direct question>
```

Do not manufacture alternatives when accepted authority leaves only one valid result.
When the user already proposes a concrete choice, validate it and ask only the minimum confirmation needed.

## Answer classification

### Explicit decision

When the answer selects a clear behavior, boundary, owner, compatibility policy, or scope disposition:

1. normalize it into concise design language;
2. compare it with accepted authority and prior decisions;
3. persist it immediately;
4. mark the item `decided`;
5. update dependencies, canonical target, and provisional ADR route;
6. select the next unblocked item.

### Ambiguous answer

When materially different interpretations remain:

- keep the same item `in_discussion`;
- state the exact ambiguity;
- ask one narrower follow-up;
- do not advance the cursor.

### Repository-resolvable statement

When accepted authority may already determine the answer:

- read that authority;
- record the resolved fact and source;
- ask for judgment only if a real choice remains.

### Deferral

When the user explicitly excludes the item:

- mark it `deferred`;
- record the reason and destination when known;
- propagate dependent blockers or deferrals;
- confirm the current design remains coherent without it.

### Contradiction or reversal

When the answer conflicts with accepted authority or an earlier decision:

- keep the item active;
- name the conflict;
- route the issue through `convergence-routing.md`;
- obtain an explicit preservation, amendment, reconsideration, or supersession disposition;
- do not silently overwrite either side.

## Immediate persistence

After each explicit answer, update the owning decision Task before advancing.

Persist:

- normalized outcome;
- status `decided`;
- concise reason;
- affected dependencies;
- expected canonical target;
- provisional ADR route;
- newly revealed decisions or blockers;
- current cursor;
- required date or Evidence.

Inspect the scoped write before asking the next question.
If persistence fails, do not advance.

## Cursor rules

After recording an item:

1. recompute blocked dependencies;
2. choose the first unblocked `open` item in the accepted priority order;
3. mark only that item `in_discussion`;
4. update the current cursor;
5. otherwise set a terminal loop state.

At most one item may be `in_discussion`.

## Loop states

| state | meaning |
|---|---|
| `in_progress` | An item is being discussed or an unblocked item remains open. |
| `decision_complete` | Every owned item is `decided`, `deferred`, or validly `blocked`. |
| `blocked` | No item can advance because named input is missing. |
| `superseded` | A later workflow artifact replaced this Task. |

`decision_complete` does not mean design closure.
Impact investigation, reconciliation, ADR routing, authoring, review, and synchronization remain downstream.

## Completed decision boundary

A completed decision Task does not track downstream progress.

Do not change a `decided` item because:

- ADR routing completed;
- an ADR was authored;
- a Specification was synchronized;
- review passed or failed;
- the Work Item closed.

There is no downstream `recorded` decision state.
Downstream Tasks own their own refs and Evidence.

When later work changes the selected option, rationale, responsibility boundary, scope, or canonical target:

- do not reopen or substantively rewrite the completed Task;
- use coordination to create a new decision Task;
- identify the earlier decision and reason for reconsideration;
- route the revised result through new downstream Tasks.

Meaning-preserving editorial and broken-reference corrections may update the original Task only when its decision and completion judgment remain unchanged.

## Resume procedure

When resuming:

1. read repository instructions;
2. read the parent Work Item and decision Task;
3. read only authority needed for the current cursor;
4. confirm at most one item is `in_discussion`;
5. confirm all explicit answers are durably persisted;
6. do not repeat terminal items;
7. repair a stale cursor to the first unblocked non-terminal item;
8. resume that item.

Repository state is authoritative over chat summaries.

## Handoff contract

A ChatGPT handoff must direct the next session to:

- read the repository instruction source;
- read this workflow and the owning decision Task;
- resume from the repository cursor;
- avoid repeating terminal decisions;
- ask exactly one current decision;
- persist before advancing;
- avoid authoring, independent review, and implementation.

Do not paste the full decision history when the Task already contains it.

## Readiness for investigation and routing

Before leaving the loop, verify:

- all known required judgments have entries;
- dependencies and blockers are explicit;
- every user answer is persisted;
- no more than one item is `in_discussion`;
- each terminal item is `decided`, `deferred`, or validly `blocked`;
- expected canonical targets are known or blocked for a named reason;
- no required decision exists only in chat;
- production implementation has not begun.
