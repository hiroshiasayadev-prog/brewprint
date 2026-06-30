# Interactive decision loop

This document defines the repository-persistent loop used to identify unresolved design decisions, ask the user one decision at a time, and preserve every explicit answer before advancing.

It is a required companion to `SKILL.md` for decision inventory, questioning, interruption, handoff, or resume.

## Purpose

Keep the interactive state recoverable from repository artifacts rather than chat history.

The loop must distinguish:

- facts resolvable from accepted repository authority;
- decisions requiring user judgment;
- decisions blocked by another decision or investigation;
- explicit decisions not yet synchronized to canonical ADRs or Specifications;
- decisions fully recorded in canonical artifacts.

## Persistent Task contract

The decision-loop Task must identify:

- parent design-topic Work Item;
- bounded topic and excluded scope;
- directly relevant authority;
- canonical ADR and Specification targets known at inventory time;
- decision inventory;
- dependency order;
- current cursor;
- decision status definitions;
- write ownership;
- stop conditions;
- downstream ADR, Specification, review, and closure Tasks.

Do not rely on a chat handoff as the only source of the current cursor or accepted answers.

## Decision inventory shape

Use a stable local decision ID such as `D-001` within the owning Task.

Recommended table:

```md
## Decision confirmation loop

Loop status: in_progress
Current decision: D-003

| ID | Topic | Status | Depends on | Decision summary | ADR route | Canonical target |
|---|---|---|---|---|---|---|
| D-001 | Persistence authority | decided | — | Database is canonical | required | `spec:...#Persistence` |
| D-002 | YAML support | decided | D-001 | Out of scope | not_required | `spec:...#Out-of-scope` |
| D-003 | Retention boundary | in_discussion | D-001 | — | unknown | `spec:...#History` |
| D-004 | Restore semantics | blocked | D-003 | — | candidate | `spec:...#Recovery` |
```

Each row must contain enough information to resume without reconstructing the decision from chat.

The summary is concise workflow evidence, not the complete normative contract.

## Inventory procedure

Before asking the first question:

1. Read only the directly relevant accepted Requirements, ADRs, Specifications, investigations, Work Item, and Task records.
2. Extract questions that cannot be resolved from those artifacts.
3. Resolve repository-answerable facts without asking the user.
4. Detect contradictions between accepted artifacts and classify them explicitly.
5. Add all currently known decisions to the inventory.
6. Add `Depends on` edges.
7. Mark unavailable downstream items `blocked`.
8. Assign a likely ADR route: `required`, `candidate`, `not_required`, or `unknown`.
9. Assign likely canonical targets without inventing unsupported paths or IDs.
10. Select the first unblocked `open` item and mark it `in_discussion`.

Do not require the complete future decision set to be knowable. New decisions may be appended when directly revealed by a confirmed answer, but the reason for adding them must be recorded.

Do not turn implementation questions into design decisions when accepted design already determines the answer.

## One-question rule

Ask exactly one unresolved decision per user turn.

One decision may contain tightly coupled sub-criteria only when they cannot be answered independently without creating an inconsistent state.

Do not present a checklist of unrelated questions for the user to answer in one message.

Do not ask the next decision until the current answer is persisted to the Task or the current item is explicitly marked `blocked` or `deferred`.

## Question format

Use this shape when alternatives need explanation:

```text
Decision: D-003 — <topic>

Why this must be decided:
<one or two concise sentences>

Recommendation:
<recommended choice and the decisive reason>

Alternatives:
- A: <meaning and consequence>
- B: <meaning and consequence>

Question:
<one direct question>
```

Do not manufacture alternatives when only one viable interpretation remains.

When the user already proposed a concrete choice, validate it against accepted authority and ask only the minimum confirmation needed.

When a recommendation is uncertain, state the uncertainty and the missing premise.

## Answer classification

Classify the response before writing:

### Explicit decision

The answer selects a clear behavior, boundary, ownership rule, compatibility policy, or scope disposition.

Action:

1. normalize it into concise design language;
2. check it against accepted authority and prior decisions;
3. persist it immediately;
4. mark the item `decided`;
5. update ADR route and canonical target when now known;
6. select the next unblocked item.

### Ambiguous answer

The response allows materially different interpretations.

Action:

- keep the same item `in_discussion`;
- state the exact ambiguity;
- ask one narrower follow-up for the same decision ID;
- do not advance the cursor.

### Repository-resolvable statement

The user asks whether accepted artifacts already determine the answer.

Action:

- read the relevant authority;
- record the resolved fact and source;
- ask for judgment only if a genuine choice remains.

### Deferral

The user explicitly excludes the decision from the current scope.

Action:

- mark it `deferred`;
- record the reason and destination Work Item, Task, investigation, or release when known;
- identify dependent items that become deferred or blocked;
- do not silently treat deferral as a design choice.

### Contradiction or reversal

The answer conflicts with an accepted ADR, Specification, Requirement, or earlier decision.

Action:

- keep the item active;
- name the conflicting artifact or decision;
- explain whether supersession, Requirement change, or scope correction is needed;
- obtain an explicit disposition;
- do not silently overwrite the earlier decision.

## Immediate persistence

After an explicit answer, update the decision-loop Task before asking the next question.

The write must include:

- normalized decision summary;
- status `decided`;
- dependencies affected;
- ADR route;
- expected canonical target;
- newly revealed decisions or blockers;
- current cursor;
- date or Evidence required by repository policy.

The Task write is the immediate durable checkpoint.

Canonical ADR and Specification synchronization occurs in their named downstream phases unless the Work Item explicitly assigns the same writer to perform those phases sequentially.

Do not mark the item `recorded` until required canonical synchronization and verification are complete.

## Write verification

After each decision-ledger update:

1. inspect the written section or scoped diff;
2. verify the decision ID, summary, status, dependency changes, and cursor;
3. ensure no unrelated decision was changed;
4. ensure the Task remains structurally valid;
5. only then ask the next question.

If the write partially fails:

- do not ask the next question;
- retain the accepted answer in the current response context;
- restore the Task to a consistent state;
- report the exact incomplete write target;
- do not claim durable persistence until verification succeeds.

## Cursor selection

After recording the current decision:

1. recompute blocked dependencies;
2. choose the first unblocked `open` item in the accepted priority order;
3. mark only that item `in_discussion`;
4. update `Current decision`;
5. if no unblocked item remains, set loop status according to the terminal rules below.

At most one item may be `in_discussion`.

## Loop terminal states

Use:

- `in_progress`: at least one item is `in_discussion` or an unblocked item is `open`;
- `decision_complete`: all required items are `decided`, `deferred`, or explicitly `blocked` and ready for ADR routing;
- `blocked`: no decision can advance because a named input, investigation, or authority is missing;
- `superseded`: the Task is replaced by another explicit workflow artifact.

`decision_complete` does not mean canonical design closure. ADR authoring, Specification synchronization, review, and closure remain downstream.

## Resume procedure

When resuming an interrupted loop:

1. read repository instruction sources required for the session;
2. read the parent Work Item and decision-loop Task;
3. read only the canonical targets needed to verify the current cursor and recent recorded decisions;
4. verify that no more than one item is `in_discussion`;
5. verify that `decided` and `recorded` states match repository evidence;
6. do not ask any `decided`, `recorded`, `deferred`, or validly `blocked` question again;
7. resume the declared current item;
8. if the cursor is stale, correct it to the first unblocked non-terminal item before asking a question.

Do not reconstruct accepted decisions from chat when the Task or canonical records exist.

## Adding newly revealed decisions

A confirmed answer may expose a previously unknown decision.

Add it only when:

- it materially affects the bounded design topic;
- it cannot be resolved from accepted authority;
- it is not merely an implementation choice already delegated by the design;
- its dependency and canonical target can be stated or explicitly marked unknown.

Record:

- why the item was added;
- which decision revealed it;
- its dependency edges;
- whether it blocks downstream ADR or Specification authoring.

Do not expand the Work Item into an adjacent design topic. Defer adjacent scope explicitly.

## Handoff contract

A handoff to another ChatGPT session must instruct it to:

- read the repository instruction source;
- read this skill and the decision-loop Task;
- resume from the repository cursor;
- not repeat terminal decisions;
- ask exactly one current decision;
- persist the answer before advancing;
- avoid implementation and independent review.

The handoff should not paste the full decision history when the Task already contains it.

## Anti-patterns

Do not:

- ask multiple unrelated questions at once;
- rely on the user to remember which decisions were already made;
- mark an ambiguous preference as `decided`;
- mark a Task-only summary as `recorded`;
- advance before the Task write is verified;
- change an earlier decision without explicit supersession;
- hide a contradiction by editing the question wording;
- ask for user judgment when accepted authority already answers the issue;
- treat an implementation detail as an architecture decision without cause;
- perform independent review in the loop session;
- close the parent Work Item when the loop merely reaches `decision_complete`.

## Readiness checklist

Before handing off to ADR routing, verify:

- all known required decisions have an inventory row;
- dependency edges are explicit;
- no more than one item is `in_discussion`;
- every explicit user answer is durably persisted;
- every unresolved item is `open`, `in_discussion`, or `blocked` for a named reason;
- every terminal loop item is `decided`, `deferred`, or validly `blocked`;
- ADR route is known or explicitly `unknown` for each `decided` item;
- expected Specification targets are identified or blocked for a named reason;
- no required decision exists only in chat;
- no production implementation has begun.
