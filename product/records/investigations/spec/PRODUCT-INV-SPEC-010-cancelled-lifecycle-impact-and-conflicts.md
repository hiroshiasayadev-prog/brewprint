# PRODUCT-INV-SPEC-010: Cancelled lifecycle impact and conflicts

- **status**: concluded
- **date**: 2026-07-03
- **trigger**: PRODUCT-WORK-SPEC-023 cancelled Work Item and Task lifecycle design
- **scope**: Decided cancellation semanticsに対するcanonical authority、workflow support、ADR coverage、validation、graph ownership、shared-writer impactの調査。
- **non_scope**: 追加判断の採用、ADRまたはSpecification authoring、Task graph変更、既存record migration、DRMCPまたはcommand実装、review、closure、stage、commit。
- **source_refs**:
  - PRODUCT-REQ-SPEC-009
  - PRODUCT-WORK-SPEC-023
  - spec:product.design_records.authoring_standards.work_item_authoring
  - spec:product.design_records.authoring_standards.task_authoring
  - PRODUCT-ADR-SPEC-005
  - PRODUCT-ADR-SPEC-010
  - PRODUCT-ADR-SPEC-011
  - PRODUCT-ADR-SPEC-014
  - PRODUCT-ADR-SPEC-017
- **follow_up_candidates**:
  - post-Investigation cancellation contract reconciliation decision
  - ADR routing and canonical authoring for cancelled lifecycle
  - later DRMCP lifecycle implementation work
- **related_requirements**:
  - PRODUCT-REQ-SPEC-009
- **related_work_items**:
  - PRODUCT-WORK-SPEC-023
- **related_adrs**:
  - PRODUCT-ADR-SPEC-005
  - PRODUCT-ADR-SPEC-010
  - PRODUCT-ADR-SPEC-011
  - PRODUCT-ADR-SPEC-014
  - PRODUCT-ADR-SPEC-017
- **related_specs**:
  - spec:product.design_records.authoring_standards.work_item_authoring
  - spec:product.design_records.authoring_standards.task_authoring
  - spec:product.responsibility_boundary_validator

## Investigation scope

This Investigation answers one bounded research question.

```text
Given PRODUCT-REQ-SPEC-009 and the terminal decisions in PRODUCT-TASK-SPEC-023-01, which current authorities and workflow consumers require amendment, which conflicts or missing judgments remain, and which downstream owners are required before canonical authoring?
```

The Investigation treats the T01 decisions as fixed inputs.
It does not reopen accepted cancellation semantics.

## Out of scope

- Existing Work Item or Task migration.
- Automatic cancellation of child Work Items or transitive descendants.
- Framing workflow design.
- Concrete DRMCP transaction, command, parser, diagnostic, or implementation behavior.
- ADR routing adoption.
- Canonical authoring.
- Task graph mutation.
- Independent review and lifecycle closure.

## Background

Current Work Item and Task authoring authority defines four statuses: `not_started`, `in_progress`, `blocked`, and `done`.

PRODUCT-REQ-SPEC-009 requires terminal `cancelled` for both kinds.
T01 additionally decides independent Task cancellation, dependent blocking, irreversible history, Work Item propagation, Evidence requirements, and `work_item_execution` cancellation.

The decided contract changes lifecycle, failure handling, dependency behavior, completion-boundary behavior, and historical preservation.
These concerns require canonical projection and durable rationale.

## What was investigated

- Work Item and Task status tables and lifecycle rules.
- Task dependency and Work Item ownership relations.
- `work_item_execution` completion behavior.
- Work Item identity continuity and append-only reconvergence ADRs.
- Design-convergence graph, execution, and closure companions.
- Task responsibility validator invocation semantics.
- Direct canonical targets, shared writers, and future implementation boundary.

Repository reads were limited to direct PRODUCT authorities and workflow-support consumers.

## Findings

### 1. Canonical lifecycle authority

| affected artifact | observed state | relation to T01 | classification | required judgment or owner | graph candidate | shared writer | evidence |
|---|---|---|---|---|---|---|---|
| `spec:product.design_records.authoring_standards.work_item_authoring` | Defines only four statuses. `done` is the only successful terminal state. No cancellation transition, propagation, Evidence, or restart rule exists. | D-002, D-003, D-004, D-006, D-007, D-009 require direct normative projection. | `consistent_refinement` with missing projection | Canonical `authoring` after ADR routing. | Authoring Task required. | Same writer as Task authoring because propagation spans both kinds. | Status lifecycle and identity sections. |
| `spec:product.design_records.authoring_standards.task_authoring` | Defines only four statuses. `depends_on` requires prerequisite completion. `work_item_execution` recognizes only child `done`, `blocked`, `not_started`, and `in_progress`. | D-001, D-005, D-007, D-008, D-009 require direct normative projection. | `consistent_refinement` with missing projection | Canonical `authoring` after ADR routing. | Authoring Task required. | Same writer as Work Item authoring. | Metadata, lifecycle, type contract, continuation, and relation sections. |
| Task `depends_on` relation | The relation is not restricted to the same Work Item. A cancelled prerequisite can therefore affect direct dependents outside the cancelled Task's parent Work Item. | D-005 applies to every direct dependent Task, not only siblings. | `consistent_refinement` | No new user judgment. Preserve D-005 as graph-wide direct dependency behavior. | None beyond canonical projection. | Task authoring writer. | Task authoring and traceability specs define complete Task IDs without same-parent restriction. |

### 2. Missing terminal-content requirements

| affected artifact | observed state | relation to T01 | classification | required judgment or owner | graph candidate | shared writer | evidence |
|---|---|---|---|---|---|---|---|
| Cancelled Task body | Task sections become substantive only when `done`. A `not_started` Task may currently retain `TBD` in every common section. D-007 requires substantive cancellation Evidence but does not define whether Goal, Work, Done condition, or Verification may remain `TBD`. | The abandoned outcome may be unrecoverable when Goal remains `TBD`. Requiring every done-gated section may also falsely imply completion work occurred. | `semantic_conflict` / missing judgment | New `decision` owner must select exact substantive-section rules for `cancelled` Tasks. | Decision Task required before authoring. | Work Item and Task authoring writer after decision. | Task file-shape and lifecycle rules. |
| Cancelled Work Item body | Work Item Goal, Boundary, and Evidence are substantive only when `done`. D-007 requires cancellation Evidence but does not define cancelled-state Goal and Boundary readiness. | A cancelled Work Item must identify the unfinished resolution boundary without pretending completion. | `semantic_conflict` / missing judgment | New `decision` owner must select exact substantive-section rules for `cancelled` Work Items. | Same decision Task. | Work Item authoring writer. | Work Item file-shape and lifecycle rules. |

### 3. Cancellation propagation responsibility

| affected artifact | observed state | relation to T01 | classification | required judgment or owner | graph candidate | shared writer | evidence |
|---|---|---|---|---|---|---|---|
| Work Item cancellation propagation | D-003 and D-006 require one coherent resulting state across the Work Item and all owned unfinished Tasks. Existing closure synchronization applies only after accepted review. A synchronization Task owned by the Work Item being cancelled would itself be an unfinished owned Task and would be cancelled during the same propagation. | The design needs an initiator and write boundary that do not create a self-cancellation paradox. | `semantic_conflict` / missing ownership judgment | New `decision` owner must define whether cancellation is an external lifecycle operation, an externally owned synchronization Task, or another bounded owner. | Decision Task required. Later authoring must state the semantic procedure without selecting concrete tool mechanics. | Work Item and Task authoring writer; workflow-support writer when directly affected. | D-003, D-006, synchronization Task contract, and closure companion. |
| Direct Task cancellation propagation | D-005 requires blocking all direct dependents. Existing Task authoring defines dependency meaning but no owner for cross-record propagation. | One cancellation can write several dependent Task records. | `workflow_graph_drift` candidate | The same new decision must define the semantic propagation owner and writable boundary. | Later graph coordination after decision. | Task authoring and workflow-support writer. | Task authoring `depends_on`; synchronization responsibility contract. |

### 4. Work Item execution and workflow support

| affected artifact | observed state | relation to T01 | classification | required judgment or owner | graph candidate | shared writer | evidence |
|---|---|---|---|---|---|---|---|
| `PRODUCT-ADR-SPEC-005` | States that `work_item_execution` becomes `done` only after child `done`; a blocked child may block it. No terminal child cancellation outcome exists. | D-008 adds an independently durable failure branch while preserving the same responsibility architecture. | `stale_representation` | ADR routing should evaluate non-material amendment. | ADR-routing Task required. | ADR writer before Specification and skill writer. | ADR Decision section. |
| `PRODUCT-ADR-SPEC-010` | Typed workflow routing says `work_item_execution` completes only after child `done`. | D-008 requires cancellation routing without changing typed phase ownership. | `stale_representation` | ADR routing should evaluate non-material amendment. | Same ADR-routing Task. | ADR writer. | ADR Decision section. |
| `skills/design-convergence-workflow/work-item-execution.md` | Defines only child `done`, `blocked`, `not_started`, and `in_progress` handling. | D-008 requires child `cancelled` to cancel the execution Task and route dependents through D-005. | `stale_representation` | Workflow-support authoring after routing. | Authoring Task required. | Same writer as canonical lifecycle projection or deterministically serialized. | Owned outcome, relation, completion, and verification sections. |
| `skills/design-convergence-workflow/graph-coordination.md` | Defines blocked and released execution routes but not cancelled prerequisites or cancelled child execution units. | D-005 and D-008 require graph consequences. | `stale_representation` | Workflow-support authoring after routing. | Authoring Task required. | Same workflow-support writer. | Work Item execution handoff and completion sections. |
| `skills/design-convergence-workflow/closure-synchronization.md` | Uses “terminal status” in a reviewed-success closure procedure and examples only `done`. Cancellation is not a reviewed-success closure. | Adding `cancelled` could make the generic phrase misleading. | `stale_representation` | Authoring should narrow successful closure wording to `done` and avoid treating cancellation as design closure. | Authoring Task required only if the exact section remains misleading after the propagation-owner decision. | Workflow-support writer. | Work Item closure and lifecycle synchronization sections. |

### 5. Historical identity and restart authority

| affected artifact | observed state | relation to T01 | classification | required judgment or owner | graph candidate | shared writer | evidence |
|---|---|---|---|---|---|---|---|
| `PRODUCT-ADR-SPEC-011` | Preserves Work Item identity while the Requirement, Goal, and Completion Conditions remain coherent. It does not exclude a terminal cancelled record. | D-009 requires a new Work Item even when later work resembles the cancelled boundary. | `stale_representation` | ADR routing should evaluate a non-material terminal-state exception amendment. | ADR-routing Task required. | ADR writer. | ADR Decision and Consequences. |
| `PRODUCT-ADR-SPEC-014` | Preserves completed decision, authoring, and review Tasks through append-only reconvergence. It does not cover cancelled Task or Work Item history. | D-009 extends the same historical-preservation rationale to unsuccessful terminal records. | `consistent_refinement` requiring durable rationale | ADR routing should evaluate amendment rather than a separate overlapping ADR when the rationale remains coherent. | ADR-routing Task required. | ADR writer. | ADR Decision and Consequences. |
| Task and Work Item authoring continuation rules | Incomplete Tasks may be amended. Work Items continue while completion identity remains coherent. Neither rule excludes `cancelled`. | D-009 requires a new record for materially resumed work. | `stale_representation` | Canonical authoring after routing. | Authoring Task required. | Task and Work Item authoring writer. | Continuation and identity sections. |

### 6. Semantic responsibility validation

| affected artifact | observed state | relation to T01 | classification | required judgment or owner | graph candidate | shared writer | evidence |
|---|---|---|---|---|---|---|---|
| `PRODUCT-ADR-SPEC-017` and `spec:product.responsibility_boundary_validator` | Require validation after Task authoring and after final Task Evidence. Current wording assigns the second call to completion or release workflow. | A cancelled Task has final Evidence but no successful completion or normal release. T01 does not decide whether cancellation finalization invokes the same validator. | `semantic_conflict` / missing judgment | New `decision` owner must decide whether cancellation triggers the post-Evidence invocation and how violations are handled before terminalization. | Decision Task required. | ADR and Specification authoring writer if invocation semantics change. | ADR-017 Decision; validator Workflow use; Task authoring usage table. |
| Exact checklist assets | Criteria evaluate responsibility structure, not lifecycle success. No criterion depends on `done` versus `cancelled`. | Cancellation does not require a new Task type or checklist file. | `consistent_refinement` with no direct change | No checklist authoring required unless the later decision changes responsibility criteria. | None. | None. | Common and type-specific checklist boundaries. |

### 7. DRMCP and implementation boundary

| affected artifact | observed state | relation to T01 | classification | required judgment or owner | graph candidate | shared writer | evidence |
|---|---|---|---|---|---|---|---|
| Current DRMCP status parsing and validation | Concrete parser, accepted-value, diagnostics, writer transaction, and multi-record update behavior are not owned by W023. | PRODUCT semantic changes will require later app-local implementation and tests. | `consistent_refinement` outside current boundary | Record a later implementation Work Item or Requirement after design closure. | No current implementation Task. | Future DRMCP owners. | PRODUCT-REQ-SPEC-009 Explicitly Excluded Scope and Boundary. |
| Existing records | Existing migration is explicitly excluded. | No current record status changes are required. | `consistent_refinement` | None. | None. | None. | PRODUCT-REQ-SPEC-009. |

## Cross-cutting observations

- The accepted status name and main transition semantics are stable.
- Canonical authoring cannot start because terminal body readiness, propagation ownership, and validator invocation remain unresolved.
- Work Item cancellation propagation cannot safely be modeled as an unfinished Task owned by the Work Item being cancelled without defining an external initiator boundary.
- Direct Task dependency effects are graph-wide because `depends_on` is not limited to one Work Item.
- `work_item_execution`, identity continuity, and append-only history contain direct stale representations.
- Exact checklist wording is not directly affected.
- Concrete DRMCP mechanics remain valid follow-up scope rather than a current design target.

## Follow-up judgment candidates

- Decide which Task sections must be substantive when a Task becomes `cancelled`.
- Decide which Work Item sections must be substantive when a Work Item becomes `cancelled`.
- Decide the semantic owner and writable boundary for direct Task cancellation and Work Item-to-Task propagation.
- Avoid a cancellation operation that is owned by an unfinished Task inside the Work Item being cancelled.
- Decide whether the post-Evidence semantic responsibility validator runs before a Task becomes `cancelled`.
- Decide whether a rejected validator violation blocks cancellation or requires an explicit exception route.

These candidates belong to one bounded post-Investigation decision Task.

## Recommendation

Create one successor `decision` Task for the unresolved terminal-content, propagation-owner, and validator-invocation contract.

After that decision completes, use a separate coordination Task to materialize:

- ADR routing;
- bounded ADR and canonical authoring;
- one integrated independent review;
- verdict-dependent correction or direct closure synchronization.

ADR routing should evaluate:

- a coherent new cancellation-lifecycle ADR boundary or an amendment-centered boundary;
- amendments to PRODUCT-ADR-SPEC-005, PRODUCT-ADR-SPEC-010, PRODUCT-ADR-SPEC-011, and PRODUCT-ADR-SPEC-014;
- PRODUCT-ADR-SPEC-017 amendment only when the validator decision changes its current completion-or-release wording.

Do not start canonical authoring before the successor decision and ADR routing are complete.

## Follow-up artifact candidates

- One post-Investigation cancellation reconciliation decision Task.
- One later graph-coordination Task after that decision.
- One ADR-routing Task.
- One bounded ADR, Work Item authoring, Task authoring, and workflow-support authoring Task or serialized writer set.
- One integrated independent review Task.
- One closure synchronization Task after an accepted review route.
- A separate future DRMCP implementation Work Item after W023 design closure.

No ADR ID or future implementation ID is reserved by this Investigation.

## Open questions

- Which common Task sections must be substantive when status is `cancelled`?
- Must a cancelled Work Item have substantive Goal and Boundary in addition to Evidence?
- What semantic operation owns multi-record cancellation propagation without cancelling its own executor?
- Does Task responsibility validation run after final cancellation Evidence and before terminal status is accepted?
- Which existing ADRs are amended, and whether one new cancellation-lifecycle ADR is required?
