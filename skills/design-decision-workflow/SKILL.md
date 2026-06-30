# Design decision workflow

## Purpose

This skill governs repository-persistent design work that begins with unresolved decisions and ends only after accepted decisions are represented in canonical artifacts and independently reviewed.

The workflow prevents:

- design decisions remaining only in chat;
- implementation planning beginning before required decisions are frozen;
- ADRs being authored without corresponding current-state Specification updates;
- Specifications being changed without an independent design review;
- a resumed session repeating already-recorded questions;
- lifecycle state being advanced before findings are closed.

This skill governs the ChatGPT assistant that coordinates the workflow. It does not replace repository authoring standards, Design Record schemas, or implementation execution skills.

## Use this skill when

Use this skill when any of the following is true:

- a design topic contains multiple unresolved product, architecture, compatibility, ownership, data-model, API, validation, lifecycle, or scope decisions;
- the user wants unresolved decisions organized and asked one at a time;
- decisions must be persisted incrementally instead of remaining in chat;
- a Requirement, Work Item, ADR, Specification, or Task has open questions that block implementation planning;
- an investigation produced alternatives that require user judgment;
- a design topic needs an explicit sequence from decision confirmation through ADR authoring, Specification synchronization, independent review, finding correction, and closure;
- an interrupted design-decision session must resume from repository state.

## Do not use this skill when

Do not create this workflow when:

- all required decisions are already accepted and recorded;
- the work is mechanical implementation of an executor-ready Task;
- the work is only independent review of an already-frozen design;
- the user requests only an explanation, comparison, or recommendation with no repository authoring;
- the remaining questions can be answered directly from accepted canonical artifacts without user judgment;
- one isolated documentation correction can be completed without a decision inventory or conditional ADR route.

## Required companion documents

Read the companion document required by the current phase:

- `interactive-decision-loop.md` for inventory, questioning, persistence, interruption, or resume;
- `adr-routing.md` when deciding whether a confirmed decision requires an ADR, updating an ADR, or superseding an accepted ADR;
- `design-review-gate.md` when creating or evaluating the independent review, finding-correction, re-review, or closure sequence.

Re-read the relevant companion when the phase, canonical target, ADR disposition, review boundary, or lifecycle owner changes.

## Relationship to implementation execution

This workflow freezes what the accepted design is.

The Claude Code token-budget and execution-hub skills govern how an accepted design is implemented.

The required order is:

```text
Design decision workflow
  -> accepted ADRs and Specifications
  -> independent design review PASS
  -> design closure synchronization
  -> executor-ready implementation Task authoring
  -> execution-hub applicability evaluation
  -> implementation execution
```

Do not issue an implementation or finding-correction prompt for production code while a required design decision remains unresolved or while the design review gate has not passed.

## Canonical artifact topology

The default persistent structure is:

```text
Design topic Work Item
├─ scope and decision-inventory Task
├─ interactive decision-loop Task
├─ conditional ADR authoring or update Task
├─ Specification synchronization Task
├─ independent design-review Task
├─ conditional finding-correction Task
├─ conditional finding-closure re-review Task
└─ lifecycle and Evidence closure-synchronization Task
```

Small topics may combine scope inventory and the interactive loop in one Task when the decision set is bounded and no ownership ambiguity results.

Do not combine:

- decision confirmation with independent review;
- ADR authoring with independent review;
- finding correction with finding-closure re-review;
- design closure synchronization with production implementation.

## Canonical ownership

Persist concerns in the following artifacts:

| concern | canonical owner |
|---|---|
| design topic, delivery boundary, Task graph, and completion conditions | Work Item |
| unresolved-decision inventory, dependency order, cursor, and decision summaries | decision-loop Task |
| durable choice, alternatives, rationale, consequences, and supersession | ADR |
| current normative behavior, shape, constraints, and out-of-scope boundary | Specification |
| independent verdict and findings | independent design-review Task |
| correction of named findings | finding-correction Task and the affected canonical artifacts |
| finding disposition after correction | finding-closure re-review Task |
| final statuses, Evidence, exact artifact refs, and Work Item closure | closure-synchronization Task |

The Task decision ledger owns workflow state only. It is not the normative product or architecture specification.

Do not leave a durable decision only in chat or only in a Task summary.

## Work Item contract

The design topic Work Item must define:

- the bounded design topic;
- canonical Requirements, existing ADRs, Specifications, and investigations in scope;
- the Task graph and dependency order;
- the conditional ADR route;
- the independent review owner;
- correction and re-review ownership;
- deferred scope;
- completion conditions.

Default completion conditions:

- every required decision is recorded or explicitly deferred;
- every durable design choice has an accepted ADR when required;
- every normative current-state decision is represented in a Specification;
- ADRs and Specifications do not conflict;
- independent design review is `PASS`;
- blocking, major, and required minor findings are closed according to the review contract;
- lifecycle and Evidence are synchronized;
- no implementation contract depends on an unresolved design decision.

Do not duplicate complete decision content or child Task lifecycle state into the Work Item.

## Workflow phases

### 1. Scope and decision inventory

Before asking the first question:

1. Read the directly relevant accepted authority.
2. Identify contradictions, gaps, ambiguous boundaries, and missing ownership.
3. Separate repository-resolvable facts from decisions requiring user judgment.
4. Create the complete known decision inventory.
5. Record dependency edges and blocked decisions.
6. Identify likely canonical targets.
7. Classify ADR disposition as `candidate`, `not_required`, or `unknown`.
8. Select the first unblocked decision.

This phase is read-only except for the Work Item or Task records explicitly assigned to inventory authoring.

### 2. Interactive decision loop

Follow `interactive-decision-loop.md`.

The loop must:

- ask exactly one unresolved decision at a time;
- explain the consequence and provide a recommendation when useful;
- avoid asking questions already answered by accepted repository authority;
- persist each explicit answer immediately to the decision ledger;
- keep ambiguous answers on the same decision ID;
- detect conflicts with earlier accepted decisions;
- advance only after the current answer is durably recorded in the Task;
- finish with every required item `decided`, `deferred`, or `blocked`.

The loop does not independently review its own decisions and does not begin production implementation.

### 3. Conditional ADR authoring or update

Follow `adr-routing.md`.

For every `decided` item:

- determine whether an ADR is required;
- create, amend, or supersede the correct ADR when required;
- preserve one coherent decision boundary per ADR;
- record the exact ADR reference back in the decision ledger or closure Evidence;
- do not use an ADR as the current normative Specification.

Skip this phase for decisions classified `not_required`.

### 4. Specification synchronization

After required ADR authoring is complete:

- write current normative behavior into the correct Specification sections;
- add the relevant ADR relationship or reference according to repository policy;
- remove or revise stale contradictory text;
- record explicit out-of-scope boundaries;
- preserve Requirement intent;
- produce exact changed-file and section Evidence for review;
- do not mark the overall workflow complete.

An ADR explains why. A Specification defines what is currently true.

### 5. Independent design review

Follow `design-review-gate.md`.

The review must be read-only and must trace:

```text
decision inventory and accepted answers
  -> ADR when required
  -> current Specification
  -> Work Item completion conditions
```

The reviewer does not correct findings or synchronize lifecycle state.

### 6. Conditional finding correction and closure re-review

When review returns `NEEDS REVISION`:

- create or use an explicitly named correction Task;
- correct only named findings and their direct consistency effects;
- preserve already accepted decisions unless a finding demonstrates a contradiction;
- perform an independent closure re-review after correction;
- do not let the correction author close their own findings.

### 7. Closure synchronization

After independent review `PASS`, or after all required findings are independently closed:

- update decision entries to `recorded` where canonical ADR and Specification refs exist;
- preserve explicitly `deferred` and `blocked` scope with reasons;
- record review verdict and finding disposition;
- synchronize exact changed files and canonical refs into Evidence;
- confirm Work Item completion conditions;
- update only explicitly owned lifecycle records;
- do not create implementation Tasks unless that operation is separately in scope.

## Decision status model

Use these states:

| status | meaning |
|---|---|
| `open` | identified but not yet discussed |
| `in_discussion` | the current question being resolved |
| `decided` | explicit user decision persisted to the Task, but canonical ADR or Specification synchronization is not complete |
| `recorded` | required canonical artifacts are synchronized and verified |
| `blocked` | cannot proceed until a named predecessor, investigation, authority, or external decision exists |
| `deferred` | explicitly excluded from the current Work Item with a recorded destination or reason |
| `superseded` | replaced by a later explicit decision and retained for traceability |

Do not mark an item `recorded` merely because it appears in chat or in the Task ledger.

## Writer ownership

Assign one writer per phase and artifact family.

The default ownership is:

- decision-loop Task during interactive questioning: decision-loop coordinator;
- ADR files: ADR authoring Task;
- Specification files: Specification synchronization Task;
- review Task: independent reviewer;
- correction files: named correction Task;
- final Task and Work Item lifecycle fields and Evidence: closure-synchronization Task.

When one session performs more than one authoring phase, keep the phase boundaries explicit and do not perform independent review in that session.

Do not allow concurrent writes to the decision ledger, the same ADR, the same Specification section, or the same lifecycle record.

## Stop conditions

Stop and report the exact blocker when:

- a required decision depends on missing authority or investigation;
- the user answer is ambiguous or conflicts with an earlier accepted decision;
- the required canonical target cannot be identified safely;
- an accepted ADR would need to be overturned without an explicit supersession decision;
- a write would cross the declared Work Item or Task boundary;
- the review exposes a new unresolved design decision;
- lifecycle closure would claim canonical synchronization that has not occurred.

Do not broaden repository exploration or infer a design answer to avoid stopping.

## Anti-patterns

Do not:

- ask several unrelated design questions in one turn;
- ask the user questions whose answers already exist in accepted artifacts;
- keep confirmed decisions only in chat;
- treat the Task ledger as the normative Specification;
- create an ADR for every editorial clarification;
- skip an ADR for a durable trade-off merely to reduce artifact count;
- write a Specification before required durable choices are explicit;
- let an ADR replace current normative Specification text;
- let an author independently approve their own design changes;
- correct findings inside the review Task;
- batch lifecycle closure before canonical files and review Evidence exist;
- start implementation while required decisions remain `open`, `in_discussion`, or `decided`;
- reopen a `recorded` decision without a new contradiction, supersession request, or explicit user instruction.

## Readiness checklist

Before declaring the design workflow complete, verify:

- the Work Item owns a bounded design topic and complete Task graph;
- every required decision has a terminal state for this Work Item;
- every `recorded` decision points to existing canonical artifacts;
- ADR routing was evaluated for every `decided` decision;
- required ADRs are accepted or ready for the declared review gate;
- Specifications state the current normative design;
- no stale contradictory Specification text remains in scope;
- independent review inspected the decision-to-ADR-to-Specification trace;
- findings are independently dispositioned;
- closure synchronization has one named writer;
- implementation planning can proceed without hidden design work.
