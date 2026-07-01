# PRODUCT-TASK-SPEC-018-01: Decide design convergence workflow

- **id**: PRODUCT-TASK-SPEC-018-01
- **status**: done
- **date**: 2026-07-01
- **work_item**: PRODUCT-WORK-SPEC-018
- **task_type**: decision
- **estimate**: 1d
- **depends_on**: []
- **outputs**:
  - PRODUCT-TASK-SPEC-018-01

## Goal

Decide the complete design-convergence workflow contract from design-topic intake through reviewed closure.

## Work

- Identify the workflow scope, start condition, and completion condition.
- Decide each responsibility boundary and conditional return route.
- Decide Requirement, Work Item, and Task amendment boundaries.
- Decide shared-writer, review, correction, and synchronization boundaries.
- Persist every accepted item as D-001 through D-023.

This Task does not author ADRs, Specifications, or skill content.

## Done condition

- D-001 through D-023 are all `decided`.
- No decision remains `open`, `in_discussion`, or `blocked`.
- The complete decision set is ready for ADR routing and boundary partitioning.

## Verification

- Confirm exactly 23 decision entries exist.
- Confirm each entry has status `decided`.
- Confirm the decision set covers workflow entry, responsibility routing, review, correction, and closure.
- Confirm no ADR, Specification, or skill authoring is included in this Task.

## Evidence

### Loop state

- Status: `decision_complete`
- Current decision: `none`
- Target skill: `skills/design-convergence-workflow/`
- Replacement target: `skills/design-decision-workflow/`

The following entries are the completed decision ledger owned by this Task.
The temporary file `skills/design-convergence-workflow/decision-ledger.md` contains the same decision content and may be removed after downstream workflow authoring no longer needs it.

### D-001 — Responsibility scope

Status: `decided`

The replacement skill owns the complete design workflow from initial design-topic input through reviewed design closure. It is not limited to post-decision convergence.

### D-002 — Replacement policy

Status: `decided`

`skills/design-decision-workflow/` will be fully replaced and deleted after the successor skill is ready. No deprecated stub will remain.

### D-003 — Skill name

Status: `decided`

The successor skill path is:

```text
skills/design-convergence-workflow/
```

### D-004 — Workflow start

Status: `decided`

The workflow starts immediately after a design topic is raised, before decision inventory and interactive questioning.

### D-005 — Workflow completion

Status: `decided`

The workflow completes only after:

- integrated independent review passes;
- required findings are corrected and independently closed;
- originating artifacts, execution graph, lifecycle, Evidence, and relations are synchronized.

Production implementation is out of scope.

### D-006 — Task responsibility split

Status: `decided`

The workflow distinguishes these responsibility units:

1. decision inventory and interactive decision loop;
2. decision-impact and conflict investigation;
3. conflict resolution and originating-artifact reconciliation decision;
4. conditional execution-graph amendment;
5. ADR routing and ADR-boundary partitioning;
6. conditional ADR authoring;
7. Specification and originating-artifact authoring;
8. integrated independent review;
9. named-finding correction;
10. independent finding-closure review;
11. lifecycle, Evidence, and relation synchronization.

Decision-impact investigation and integrated independent review are mandatory. Other units are conditional when their owned outcome is unnecessary.

### D-007 — Impact-investigation artifact

Status: `decided`

Decision-impact and conflict investigation uses an `investigation` Task that owns a formal Investigation record. The Investigation records affected artifacts, conflicts, graph-change candidates, and shared-writer candidates.

### D-008 — Conflict-resolution decision boundary

Status: `decided`

The reconciliation `decision` Task selects resolution policy for discovered inconsistencies, including:

- Requirement preservation, amendment, split, or replacement;
- Work Item continuation, amendment, split, or creation;
- prior-decision preservation, reconsideration, or supersession;
- ADR and Specification change targets;
- whether graph amendment is required;
- shared-writer ordering policy.

It does not itself add Tasks, change dependencies, assign writers, or author canonical artifacts.

### D-009 — Conditional coordination insertion

Status: `decided`

A `coordination` Task is required only when the existing execution graph must change, for example:

- a required responsibility has no owning Task;
- a Task must be added, removed, or split;
- dependencies, blockers, owners, writer order, review order, or release conditions must change;
- work must be split into another Work Item.

`coordination` changes the execution graph. It does not choose the design resolution or author the originating artifact.

Normative examples:

1. No inconsistency and the existing graph is sufficient: add no Task.
2. The current decision already fixed the artifact amendment, but no Task owns the write: add only the required `authoring` Task and adjust dependencies.
3. The originating artifact disposition remains unresolved: add a reconciliation `decision` Task, then a conditional `authoring` Task, and block downstream work until the decision is complete.

Task insertion is determined by unowned judgment or required graph change, not merely by detecting an inconsistency.

### D-010 — Return-to-decision convergence rule

Status: `decided`

Any workflow phase that discovers an unresolved design choice must stop downstream progression and route the issue back to a `decision` owner.

- If an active, incomplete decision Task already owns the choice and the graph remains valid, return directly to that Task.
- If the earlier decision Task is already complete, do not reopen or substantively rewrite it. Run `coordination` to create a new decision Task for the reconsideration.
- If no Task owns the choice or dependencies, blockers, writers, review order, release conditions, or Work Item boundaries must change, run `coordination` first to create or repair the route.

Examples:

- Authoring exposes multiple materially different interpretations: stop authoring and return to decision; add coordination only when the route or owner is missing.
- Review exposes a new design choice: do not send it to correction; return to decision, with coordination when the graph must change.
- Synchronization discovers that new work or judgment is required: stop synchronization and return to coordination or decision rather than introducing the change there.

When the current decision already notices an inconsistency and also fixes its disposition, no additional decision Task is required. Existing authoring may proceed if it owns the affected artifact; otherwise coordination adds only the missing authoring route.

### D-011 — Originating-artifact mismatch classification

Status: `decided`

Classify the relationship between a decision result and its originating artifact as follows:

1. `consistent_refinement`: the decision concretizes the originating artifact without changing its meaning or boundary; no originating-artifact change is required.
2. `stale_representation`: the intended disposition is already decided, but the originating artifact wording or references are stale; update it through `authoring`.
3. `semantic_conflict`: the decision and originating artifact disagree materially, and the disposition remains unresolved; return to a `decision` owner.
4. `workflow_graph_drift`: the existing Tasks, dependencies, ownership, writer order, review order, release conditions, or Work Item boundary no longer support the accepted route; use `coordination` to amend the graph.

`semantic_conflict` and `workflow_graph_drift` may occur together. The classification determines the required responsibility route rather than prescribing a fixed Task sequence.

### D-012 — Existing Requirement amendment boundary

Status: `decided`

Use the following disposition rules for an originating Requirement:

- Amend the existing Requirement when the motivating problem and purpose remain the same, and the change is clarification or a boundary adjustment within the same Required Outcome. The current Work Item completion meaning must remain coherent.
- Create a new Requirement when the motivating problem, Required Outcome, or material Scope / Excluded Scope changes enough that the new request could be accepted or rejected independently from the original.
- Split a follow-up Requirement when the original Requirement remains completable as written and the decision process reveals an adjacent additional need.
- Reconsider the decision rather than editing the Requirement when the Requirement remains valid and the current decision is the element that departed from it.

In-place amendment is for making the same requirement accurate. A materially different request receives a separate Requirement identity.

### D-013 — Work Item continuation boundary

Status: `decided`

Continue the current Work Item when:

- it still resolves the same Requirement;
- its Goal and Completion Conditions retain their meaning;
- convergence requires only Task additions, dependency changes, or ordering changes inside the same delivery boundary.

Create or split another Work Item when:

- a new Requirement is introduced;
- the new scope has an independent completion judgment;
- ownership, release timing, or primary deliverables separate from the current Work Item;
- an adjacent design topic can proceed independently after the current Work Item closes.

A larger Task graph alone does not require a new Work Item. Split only when the resolution identity or completion boundary becomes independent.

### D-014 — Existing Task amendment versus new Task

Status: `decided`

Amend an existing Task when:

- its `task_type` remains unchanged;
- its owned outcome and completion judgment remain the same;
- only the target artifacts, dependencies, or bounded procedure expand;
- the Task still preserves a single responsibility.

Add a separate Task when:

- a different `task_type` is required;
- a distinct outcome or completion judgment appears;
- independence must be preserved, such as authoring versus review or correction versus finding-closure review;
- the existing Task would otherwise own multiple responsibilities.

The rule is: extend the existing Task for the same responsibility; add a new Task for a new responsibility.

### D-015 — Shared-writer serialization

Status: `decided`

When multiple Tasks write the same artifact or section:

- do not allow concurrent writes;
- use `coordination` to persist a deterministic writer order;
- require each later writer to read and preserve the accepted semantics written by earlier writers;
- place integrated review after the final writer so the combined result is reviewed;
- return to investigation and decision, with coordination when needed, if a later writer must weaken, remove, or reinterpret earlier accepted semantics.

The shared-writer route is sequential authoring followed by one integrated review of the final combined state.

### D-016 — Integrated review boundary

Status: `decided`

Use one final integrated review Task per Work Item after all required decisions, coordination, ADR authoring, originating-artifact authoring, Specification authoring, and shared-writer sequencing are complete.

The integrated review inspects:

- the originating Requirement;
- the Work Item Goal, Boundary, and Completion Conditions;
- decided outcomes and the supporting Investigation;
- required ADRs;
- the final Specification state;
- added or amended Task graph structure;
- the final combined state after all shared writers.

Do not create multiple partial integrated reviews inside one Work Item. When a design boundary is independently closable, split it into another Work Item and give each Work Item its own final integrated review.

### D-017 — Review finding return route

Status: `decided`

Route integrated-review findings as follows:

- Use `correction` when the accepted decision is clear and the defect is only an incorrect or incomplete projection, stale text, missing reference, or contradictory wording that requires no new design choice.
- Return to `decision` when resolving the finding requires a new choice, exposes conflict among accepted decisions, or leaves unresolved whether the Requirement or the decision should change.
- Run `coordination` first when the required correction or decision Task has no owner, downstream work must be blocked, or dependencies and writer order must be amended.

Do not place unresolved design judgment inside a correction Task.

### D-018 — Pre-authoring convergence gate

Status: `decided`

By default, do not create a separate `verification` Task before ADR and Specification authoring. Treat convergence as a completion or release condition on the owning `coordination` Task and downstream authoring Tasks.

The normal gate requires:

- no unresolved decision remains;
- the originating Requirement and Work Item disposition is fixed;
- every required Task exists in the graph;
- dependencies and writer order are fixed;
- authoring targets are identified.

Create a separate `verification` Task only when the gate spans multiple Work Items, depends on objective machine-checkable results, requires an independently owned verification record, or must be performed by a verifier distinct from the coordination owner.

### D-019 — Decision Task completion and revision boundary

Status: `decided`

A `decision` Task completes when every owned item is `decided`, `deferred`, or validly `blocked` and its Done condition is satisfied.

After completion:

- the Task does not track whether ADR routing, ADR authoring, Specification synchronization, review, or Work Item closure has completed;
- `decided` remains `decided`; there is no downstream `recorded` state;
- downstream Tasks record their own results, references, and Evidence without writing progress back into the completed decision Task.

When a later change would alter the selected option, rationale, responsibility boundary, scope, or canonical target:

- do not substantively rewrite the completed decision Task;
- create a new `decision` Task that identifies the earlier decision, records the reason for reconsideration, and owns the revised decision;
- route the revised decision through the required ADR, authoring, review, and synchronization Tasks.

Meaning-preserving editorial or broken-reference correction may update the original Task when it does not alter the decision or its completion judgment.

This decision constrains D-014: substantive revision of a completed decision is a new decision responsibility and therefore requires a new Task.

### D-020 — Conditional Task materialization

Status: `decided`

Do not create speculative correction or finding-closure review Tasks in the initial graph.

The Work Item records the conditional route, but the concrete Tasks are materialized only when integrated review returns named findings:

- a `PASS` verdict proceeds directly to closure synchronization;
- a `NEEDS REVISION` verdict routes through `coordination` to create the exact correction and independent finding-closure review Tasks required by the actual findings;
- outputs, writable boundaries, dependencies, Done conditions, and review scope are derived from the named findings rather than predicted in advance;
- do not create placeholder or synthetic no-op Tasks for branches that were not taken.

This avoids stale Task contracts when the actual finding targets or required corrections differ from the initial prediction.

### D-021 — ADR routing and boundary partitioning

Status: `decided`

ADR routing is a separate responsibility from ADR authoring.

Routing owns:

- classifying each decided item as ADR `required`, `covered`, `not_required`, or `blocked`;
- deciding whether several decision items form one coherent ADR boundary or must be split across multiple ADRs;
- preventing both oversized omnibus ADRs and mechanically fragmented one-row-per-ADR authoring;
- selecting create, amend, reuse, or supersede disposition;
- mapping exact decision IDs to each routed ADR boundary and affected Specification targets.

Routing does not author ADR body content. Conditional ADR authoring consumes the completed routing result and writes the routed ADR set through one or more separate `authoring` Tasks, split only where artifact ownership or completion boundaries require it.

### D-022 — New decision after completed review

Status: `decided`

When review or later work exposes a new design judgment after the original decision, authoring, or review Tasks are complete:

- create a new decision Task rather than reopening the completed decision Task;
- create new authoring Tasks for the revised ADR, Specification, Requirement, or other canonical targets as required;
- create a new integrated review Task for the revised combined state;
- retain completed authoring and review Tasks as historical Evidence and inputs rather than changing their outcomes.

A finding that requires no new judgment remains on the correction and independent finding-closure review route.

### D-023 — Closure-synchronization write boundary

Status: `decided`

A closure-synchronization Task may write only:

- its own status and Evidence;
- the parent Work Item status and closure Evidence;
- exact mechanically derivable lifecycle or relation targets explicitly named by its contract and writable boundary.

It must not:

- rewrite a completed decision Task or its decision entries;
- alter authoring Task Evidence or outputs;
- alter a review verdict or finding set;
- create or amend the Task graph;
- introduce new design judgment;
- author or correct canonical design content.

When closure discovers missing work, graph change, or unresolved judgment, it stops and routes the issue to `coordination`, `decision`, or the appropriate correction owner.

### Decision closure

- D-001 through D-023 are `decided`.
- No open decisions remain.
- ADR routing is owned by `PRODUCT-TASK-SPEC-018-02`.
