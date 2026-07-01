# PRODUCT-TASK-SPEC-016-04: Run Specification-conflict decision loop

- **id**: PRODUCT-TASK-SPEC-016-04
- **status**: done
- **date**: 2026-06-30
- **work_item**: PRODUCT-WORK-SPEC-016
- **source_requirement**: PRODUCT-REQ-SPEC-005
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-016-03
- **outputs**:
  - PRODUCT-TASK-SPEC-016-04

## Goal

Persist one disposition at a time for material Specification conflicts found by T03.

Finish with every W016-owned conflict decided, deferred, or blocked.

## Work

- Initialize a conflict decision ledger from T03.
- Ask exactly one conflict disposition per user turn.
- Check each answer against REQ-005, T02 decisions, and W017 ownership.
- Persist the answer and verify the scoped diff before advancing.
- Route W017-owned questions to W017 without deciding them.

## Done condition

- Every W016-owned conflict has a durable disposition.
- Every shared conflict has an explicit writer-order or dependency disposition.
- Every W017-owned conflict remains routed to W017.
- ADR routing and Specification synchronization can proceed without hidden conflict decisions.

## Verification

- Confirm at most one conflict decision is `in_discussion`.
- Confirm every explicit answer is persisted before cursor advancement.
- Confirm no ADR or Specification file changed.
- Confirm no W017-owned relation decision was adopted in W016.

## Evidence

### Loop state

Loop status: done

Current conflict: none

Inventory source: PRODUCT-TASK-SPEC-016-03

### Conflict decision ledger

| ID | Topic | Status | T02 authority | Decision summary | ADR route | Canonical target |
|---|---|---|---|---|---|---|
| C-001 | Persisted `task_type` field and closed value set | decided | D-001, D-002 | Add required scalar `task_type` to the canonical Task metadata contract. Permit exactly one value from the closed set `investigation`, `decision`, `authoring`, `implementation`, `review`, `correction`, `verification`, `coordination`, and `synchronization`. Require the field on create, allow it as an update target, and defer parser and diagnostic details to downstream DRMCP. | candidate | `spec:product.design_records.authoring_standards.task_authoring` |
| C-002 | Type-specific outcome and overlap contracts | decided | D-003 through D-011 | Represent all nine Task-type contracts in one canonical type-contract table within `task_authoring`. Each row defines the primary outcome, completion judgment, prohibited overlaps, and type-specific notes. Put longer rules after the table as type-specific supplements. Do not treat T02 as the canonical contract source. | candidate | `spec:product.design_records.authoring_standards.task_authoring` |
| C-003 | Mandatory single-responsibility cohesion rule | decided | D-017, D-018 | Make responsibility cohesion mandatory. A Task must be split when it owns more than one primary outcome, completion judgment, acceptance or verification boundary, owner, release decision, or Task type. Keep the 0.5d-to-3d effort range as advisory guidance rather than a responsibility rule. Do not split mechanically when one primary outcome and completion judgment remain shared. | candidate | `spec:product.design_records.authoring_standards.task_authoring` |
| C-004 | Conditional `## Implementation contract` section | decided | D-006, D-017 | Require `## Implementation contract` only for `task_type: implementation` and prohibit it for all other Task types. Place it after `## Work` and before `## Done condition`. Represent each target with required columns `target`, `required change`, `acceptance criterion`, and `verification`. Require substantive values before `done`; permit `TBD` only before completion. `Done condition` and `Verification` may not introduce acceptance requirements absent from this section. | candidate | `spec:product.design_records.authoring_standards.task_authoring` |
| C-005 | Goal, Work, Done condition, and Verification alignment | decided | D-017 | Require all four common sections to serve one `task_type`-aligned primary outcome. `Goal` declares the outcome. `Work` contains only actions needed to produce it. `Done condition` defines one completion judgment for it. `Verification` confirms that judgment and adds no new acceptance requirement. Supporting actions from another type are allowed only when they own no separate deliverable or completion judgment. | candidate | `spec:product.design_records.authoring_standards.task_authoring` |
| C-006 | Decision ledger versus mandatory ADR wording | decided | D-004, D-012 | A `decision` Task persists the selected option, concise reason, and canonical target in its ledger, then completes. ADR need is evaluated by a downstream routing Task. Only when routed does a separate `authoring` Task create, update, or supersede the ADR. A `decision` Task does not modify ADR body content. | not_required | `spec:product.design_records.authoring_standards.task_authoring` |
| C-007 | Workflow rationale versus canonical ADR rationale | decided | D-004, D-012 | Limit Task decision-ledger rationale to the selected option, rejected-option identifiers, a one-to-three-sentence reason, canonical target, and ADR-routing state. When an ADR is required, the ADR owns background, detailed alternatives, trade-offs, consequences, supersession, and durable rationale. After ADR authoring, the Task retains workflow history but points to the ADR as canonical rationale. | not_required | `spec:product.design_records.authoring_standards.task_authoring` |
| C-008 | Review evidence versus verification outcome | decided | D-013 | Treat checks needed only to complete the current Task as supporting evidence in that Task's `## Verification`. Create a separate `verification` Task when the verification itself owns an independent pass/fail judgment, requires a separate owner or session, aggregates multiple Tasks or an integrated state, gates release or the next phase, routes failures, or evaluates post-completion state. `review` judges soundness and defects; `verification` confirms predefined criteria. | candidate | `spec:product.design_records.authoring_standards.task_authoring` |
| C-009 | Correction versus finding-closure review | decided | D-008, D-014 | A `correction` Task may implement the repair and record self-check Evidence, but it may not independently close the finding it repairs. A separate `review` Task owns finding closure, even when the same person executes both in separate Tasks or sessions. A formally recorded finding follows this rule regardless of size; a minor issue discovered and fixed within the original Task may remain inside that Task when no independent finding was opened. | candidate | `spec:product.design_records.authoring_standards.task_authoring` |
| C-010 | Coordination versus synchronization | decided | D-010, D-011, D-015 | `coordination` owns changes to execution graph, owner, dependency, blocker, release order, and next-step structure. `synchronization` only propagates already accepted state across related records. A synchronization Task must stop and route back to `coordination` or `decision` when a new Task, dependency, owner, release judgment, or choice among propagation methods is required. | candidate | `spec:product.design_records.authoring_standards.task_authoring` |
| C-011 | Parent coordination and child Work Item detail duplication | decided | D-010, D-015 | A coordinating parent Work Item may record each child Work Item ID, one-line purpose, responsibility boundary, coarse inter-child routing, and parent-level completion state. It must not duplicate child Task lists, internal Task order, procedures, detailed dependencies, release conditions, or child-local next-step decisions. Those remain owned by each child Work Item. | candidate | `spec:product.design_records.authoring_standards.work_item_authoring` |
| C-012 | Implementation-time judgment boundary | decided | D-006, D-016 | An `implementation` Task may choose only observable-equivalent local details that do not change accepted contract, external behavior, acceptance criteria, validation semantics, diagnostics, persistence shape, fallback order, or Specification meaning. Any choice that affects those boundaries, resolves a Specification conflict, or produces materially different user-visible behavior must stop the implementation Task and route to `decision`. | candidate | `spec:product.design_records.authoring_standards.task_authoring` |
| C-013 | Decision Task ledger versus canonical artifact ownership | decided | D-004, D-012 | Clarify that a Task does not own canonical design state. A `decision` Task may own temporary and historical decision-workflow state, including questions, options, answers, concise reasons, routing state, and cursor position, until the result is routed to an ADR or synchronized into a Specification. ADRs and Specifications remain the canonical design-state owners. | candidate | `spec:product.design_records.artifact_model.artifact_responsibility_matrix` |
| C-014 | Task checkpoint versus durable ADR recording | decided | D-004, D-012 | Every explicit decision is first checkpointed in the `decision` Task ledger so the workflow can be resumed safely. A later ADR-routing step determines whether durable design rationale is required. Only routed decisions receive a separate ADR-authoring Task; non-routed decisions synchronize directly into the relevant Specification. After ADR authoring, the Task retains history and references the ADR. | candidate | `spec:product.design_records.authoring_standards.artifact_boundary` |

### Recorded dispositions

- C-001: The canonical Task metadata contract must persist exactly one required scalar `task_type`. The allowed values are the nine T02 values. Create requires the field. Update may change it. PRODUCT owns the field semantics and allowed values. DRMCP parser, validation, and diagnostic behavior remain downstream.
- C-002: `task_authoring` must contain one type-contract table covering all nine values. Each row defines primary outcome, completion judgment, prohibited overlaps, and type-specific notes. Longer type rules may follow as supplements. The T02 ledger remains workflow Evidence rather than the canonical Specification.
- C-003: Single-responsibility cohesion is mandatory. Split a Task when it owns multiple primary outcomes, completion judgments, acceptance or verification boundaries, owners, release decisions, or Task types. Retain the 0.5d-to-3d effort range as advisory guidance. Do not split solely because several files or independently editable locations serve one shared outcome and completion judgment.
- C-004: `## Implementation contract` is required only for `task_type: implementation` and prohibited for all other types. It appears after `## Work` and before `## Done condition`. Each target maps to `target`, `required change`, `acceptance criterion`, and `verification`. All values must be substantive before `done`; `TBD` is allowed only before completion. `Done condition` and `Verification` must not add requirements absent from this contract.
- C-005: `## Goal`, `## Work`, `## Done condition`, and `## Verification` must all serve one primary outcome matching `task_type`. Goal declares the outcome. Work contains only actions needed to produce it. Done condition defines one completion judgment. Verification confirms that judgment without adding new acceptance requirements. Supporting actions from another type are allowed only when they own no separate deliverable or completion judgment.
- C-006: A `decision` Task records the selected option, concise reason, and canonical target in its ledger and then completes. ADR need is evaluated later by an ADR-routing Task. Only routed decisions receive a separate `authoring` Task that creates, updates, or supersedes the ADR. The `decision` Task itself must not author ADR body content.
- C-007: Task Evidence may retain only workflow-level rationale: selected option, rejected-option identifiers, a one-to-three-sentence reason, canonical target, and ADR-routing state. When an ADR is required, the ADR owns detailed background, alternatives, trade-offs, consequences, supersession, and durable rationale. The Task remains historical Evidence and points to the ADR after authoring.
- C-008: Verification that exists only to complete the current Task remains supporting Evidence in that Task. A separate `verification` Task is required when verification owns an independent pass/fail judgment, separate owner or session, aggregate or integrated scope, release or next-phase gate, failure routing, or post-completion evaluation. `review` judges soundness and defects; `verification` confirms predefined criteria.
- C-009: A `correction` Task implements a repair and may record self-check Evidence, but a separate `review` Task must independently close any formally recorded finding. The same person may perform both only as separate Tasks or sessions. An unrecorded minor issue found and fixed inside the original Task does not require a separate closure Task.
- C-010: `coordination` changes the execution graph or its judgments, including Task structure, owner, dependency, blocker, release order, or next steps. `synchronization` only propagates already accepted state. Synchronization must stop and route to `coordination` or `decision` when any new graph or release judgment is required.
- C-011: A coordinating parent Work Item may summarize child IDs, one-line purposes, responsibility boundaries, coarse inter-child routing, and parent-level completion state. It must not duplicate child Task graphs, procedures, detailed dependencies, release conditions, or child-local next-step decisions; those remain owned by each child Work Item.
- C-012: An `implementation` Task owns only observable-equivalent local choices that leave accepted contract, external behavior, acceptance criteria, validation semantics, diagnostics, persistence shape, fallback order, and Specification meaning unchanged. Any contract-affecting or materially user-visible choice, including Specification conflict resolution, must stop implementation and route to `decision`.
- C-013: A Task does not own canonical design state. A `decision` Task may own temporary and historical decision-workflow state, including questions, options, answers, concise reasons, routing state, and cursor position, until the result is routed to an ADR or synchronized into a Specification. ADRs and Specifications remain canonical design-state owners.
- C-014: Every explicit decision is first checkpointed in the `decision` Task ledger. ADR routing then determines whether durable design rationale is required. Routed decisions receive a separate ADR-authoring Task; non-routed decisions synchronize directly into the relevant Specification. After ADR authoring, the Task remains historical Evidence and references the ADR.

### Routed and excluded scope

- T03 R-001 through R-005 remain routed to PRODUCT-WORK-SPEC-017.
- T03 O-001 through O-004 remain outside T04.
- No W017-owned source-relation decision is adopted in this loop.
- No ADR or Specification authoring begins in T04.

### Closure verification

- T03 is `done`.
- C-001 through C-014 are represented once and durably recorded as `decided`.
- No conflict remains `open` or `in_discussion`.
- T03 R-001 through R-005 remain routed to PRODUCT-WORK-SPEC-017.
- No ADR or Specification authoring was performed in T04.
- T04 is complete and ready for downstream ADR routing.
