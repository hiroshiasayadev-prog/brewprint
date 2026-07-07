# PRODUCT-TASK-SPEC-023-06: Classify cancelled lifecycle ADR routing

- **id**: PRODUCT-TASK-SPEC-023-06
- **status**: done
- **date**: 2026-07-03
- **work_item**: PRODUCT-WORK-SPEC-023
- **task_type**: decision
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-023-04
  - PRODUCT-TASK-SPEC-023-05
- **outputs**:
  - PRODUCT-TASK-SPEC-023-06

## Goal

Classify every terminal cancellation decision into one complete ADR route and coherent ADR boundary set.

## Work

- Read T01, PRODUCT-INV-SPEC-010, and T04.
- Inspect accepted ADR coverage for lifecycle, dependency, Work Item execution, identity, append-only history, and validator invocation.
- Classify each T01 and T04 decision as `required`, `covered`, `not_required`, or `blocked`.
- Select `create`, `amend`, `reuse`, or `supersede` for every routed ADR boundary.
- Partition durable choices into coherent non-omnibus boundaries.
- Identify exact ADR IDs, Specification targets, workflow-support targets, and required authoring order.
- Record exact downstream authoring Task boundaries for T07.

This Task must not:

- reopen terminal decisions;
- author or amend ADR files;
- author Specifications or workflow-support files;
- change the Task graph;
- perform review, correction, synchronization, implementation, stage, or commit work.

## Done condition

- Every T01 D-item and T04 J-item has exactly one routing outcome.
- Every required choice belongs to one coherent ADR boundary.
- Every covered choice names accepted current authority.
- Every not-required choice has an exact reason and canonical target.
- Every amendment or supersession has a materiality judgment.
- Exact downstream authoring boundaries and writer order are explicit.
- No ADR body is authored.

## Verification

- Confirm all terminal decisions appear exactly once.
- Confirm only allowed routing outcomes are used.
- Confirm no durable lifecycle or atomicity choice is omitted.
- Confirm existing ADR amendments preserve historical honesty.
- Confirm exact canonical and workflow-support targets are named.
- Confirm no graph change, authoring, review, synchronization, implementation, stage, or commit occurs.

## Evidence

### Execution basis

- T01 D-001 through D-011 are terminal.
- T04 J-001 through J-005 are terminal.
- PRODUCT-INV-SPEC-010 is concluded.
- Accepted ADR authority inspected: PRODUCT-ADR-SPEC-001, 005, 010, 011, 014, and 017.
- ADR routing and amendment materiality follow `skills/design-convergence-workflow/adr-routing.md` and the ADR authoring Specification.
- DRMCP is non-operational. Filesystem authoring was the required fallback.

### Decision routing results

| decision | outcome | ADR boundary | exact ADR | disposition | reason | canonical target |
|---|---|---|---|---|---|---|
| D-001 | `required` | B-001 | PRODUCT-ADR-SPEC-018 | `create` | Independent terminal Task cancellation changes lifecycle and future workflow behavior. | Task authoring |
| D-002 | `required` | B-001 | PRODUCT-ADR-SPEC-018 | `create` | Work Item cancellation establishes a second terminal lifecycle outcome. | Work Item authoring |
| D-003 | `required` | B-001 | PRODUCT-ADR-SPEC-018 | `create` | Work Item-to-Task cancellation propagation is a durable cross-record state rule. | Work Item and Task authoring |
| D-004 | `required` | B-001 | PRODUCT-ADR-SPEC-018 | `create` | Independent Task cancellation and parent completion continuity require durable rationale. | Work Item and Task authoring |
| D-005 | `required` | B-001 | PRODUCT-ADR-SPEC-018 | `create` | Cancelled prerequisites introduce durable dependency-failure behavior. | Task authoring and graph coordination support |
| D-006 | `required` | B-001 | PRODUCT-ADR-SPEC-018 | `create` | A cancelled Work Item's valid owned-Task state is a canonical coherence constraint. | Work Item and Task authoring |
| D-007 | `required` | B-001 | PRODUCT-ADR-SPEC-018 | `create` | Cancellation Evidence distinguishes intentional termination from blocking and successful completion. | Work Item and Task authoring |
| D-008 | `required` | B-002 | PRODUCT-ADR-SPEC-005; PRODUCT-ADR-SPEC-010 | `amend` | Existing Work Item execution authority lacks the child-cancelled branch. The responsibility architecture remains unchanged. | Task and Work Item authoring; work-item-execution and graph-coordination support |
| D-009 | `required` | B-003 | PRODUCT-ADR-SPEC-011; PRODUCT-ADR-SPEC-014 | `amend` | Existing identity-continuity and append-only history rules need an explicit terminal-cancellation exception. Their selected architecture remains valid. | Work Item and Task authoring |
| D-010 | `not_required` | none | none | `reuse` | Migration, descendant cancellation, framing, and concrete implementation exclusions are bounded scope statements. | PRODUCT-WORK-SPEC-023 and canonical non-goals |
| D-011 | `covered` | existing ownership authority | PRODUCT-ADR-SPEC-001 | `reuse` | PRODUCT already owns app-independent Design Records authoring and lifecycle semantics. | PRODUCT design-records Specifications |
| J-001 | `not_required` | none | none | `reuse` | Exact cancelled-Task section readiness is direct normative projection of B-001. | Task authoring |
| J-002 | `not_required` | none | none | `reuse` | Exact cancelled-Work-Item section readiness is direct normative projection of B-001. | Work Item authoring |
| J-003 | `required` | B-001 | PRODUCT-ADR-SPEC-018 | `create` | Atomic affected-set computation and all-or-nothing propagation are durable lifecycle and failure-handling choices. | Work Item and Task authoring; lifecycle workflow support |
| J-004 | `required` | B-004 | PRODUCT-ADR-SPEC-017 | `amend` | The two-point validation architecture remains, but the final invocation is narrowed to successful completion rather than cancellation. | Task authoring and responsibility-boundary validator |
| J-005 | `not_required` | none | none | `reuse` | J-004 removes cancellation-time validator results, so no violation route exists. | Task authoring and responsibility-boundary validator |

Routing row count: 16.

| outcome | count |
|---|---:|
| `required` | 11 |
| `covered` | 1 |
| `not_required` | 4 |
| `blocked` | 0 |

### ADR boundaries

#### B-001 — Terminal cancellation lifecycle and atomic propagation

- Included decisions: D-001 through D-007 and J-003.
- Exact ADR: new PRODUCT-ADR-SPEC-018.
- Disposition: `create`.
- Bounded question: How do Work Items and Tasks terminate intentionally before completion while preserving graph coherence and history?
- Adopted scope: status eligibility, terminal meaning, parent behavior, dependency effects, propagation set, Evidence, atomic affected-set calculation, and all-or-nothing writes.
- Rejected-alternative scope: parent-only Task cancellation; automatic parent cancellation; transitive dependency cancellation; Task-owned cancellation execution; partial best-effort propagation.
- Specification targets: task authoring and work-item authoring.
- Workflow-support targets: graph coordination and closure wording where needed.

#### B-002 — Cancelled child Work Item in `work_item_execution`

- Included decision: D-008.
- Exact ADRs: PRODUCT-ADR-SPEC-005 and PRODUCT-ADR-SPEC-010.
- Disposition: `amend` both.
- Materiality: non-material amendment. The typed responsibility architecture and execution-unit ownership remain unchanged. The amendment adds the terminal failure branch that follows the new lifecycle.
- Writer order: ADR-005 before ADR-010.
- Specification target: task authoring and work-item authoring.
- Workflow-support targets: `work-item-execution.md` and `graph-coordination.md`.

#### B-003 — Irreversible cancellation history and resumed-work identity

- Included decision: D-009.
- Exact ADRs: PRODUCT-ADR-SPEC-011 and PRODUCT-ADR-SPEC-014.
- Disposition: `amend` both.
- Materiality: non-material amendment. Completion-boundary identity and append-only reconvergence remain selected. The amendment excludes terminal cancelled records from in-place continuation and extends historical preservation to them.
- Writer order: ADR-011 before ADR-014.
- Specification targets: task authoring and work-item authoring.

#### B-004 — Successful-completion-only final responsibility validation

- Included decision: J-004.
- Exact ADR: PRODUCT-ADR-SPEC-017.
- Disposition: `amend`.
- Materiality: non-material amendment. The same validator, two workflow checkpoints for authored and completed Tasks, and human-owned exceptions remain valid. The second checkpoint is clarified as `done` finalization only.
- Specification targets: task authoring and `spec:product.responsibility_boundary_validator`.
- Checklist artifacts: unchanged.

### Exact authoring boundaries for T07

#### ADR authoring boundary

- Task type: `authoring`.
- Writable ADRs: PRODUCT-ADR-SPEC-018, PRODUCT-ADR-SPEC-005, PRODUCT-ADR-SPEC-010, PRODUCT-ADR-SPEC-011, PRODUCT-ADR-SPEC-014, and PRODUCT-ADR-SPEC-017.
- Writer order: create ADR-018; amend ADR-005 then ADR-010; amend ADR-011 then ADR-014; amend ADR-017.
- Completion judgment: every durable cancellation decision has historically honest ADR authority with no reversal concealed as amendment.

#### Canonical and workflow-support authoring boundary

- Task type: `authoring`.
- Depends on: completed ADR authoring boundary.
- Writable Specifications:
  - `spec:product.design_records.authoring_standards.task_authoring`;
  - `spec:product.design_records.authoring_standards.work_item_authoring`;
  - `spec:product.responsibility_boundary_validator`.
- Writable workflow support:
  - `skills/design-convergence-workflow/work-item-execution.md`;
  - `skills/design-convergence-workflow/graph-coordination.md`;
  - `skills/design-convergence-workflow/closure-synchronization.md` only to prevent successful closure wording from including cancellation.
- Explicitly unchanged:
  - responsibility-validator checklist files;
  - Requirement content;
  - existing records;
  - DRMCP implementation.
- Completion judgment: all current normative and workflow-support text expresses the same cancellation contract and ADR authority.

#### Integrated review boundary

- One independent `review` Task after both authoring boundaries.
- Review T01, PRODUCT-INV-SPEC-010, T04, T06, all six ADR targets, all three Specification targets, three workflow-support targets, and final W023 graph state.
- Correction and finding-closure review remain conditional on named findings.

#### Closure boundary

- One `synchronization` Task may be materialized after the integrated review Task because its direct PASS dependency and writable Work Item/Task closure targets are exact.
- It must not implement the cancellation operation or modify unrelated existing records.

### Verification result

- All 16 terminal decisions appear exactly once.
- One new ADR and five in-place ADR amendments are required.
- No ADR supersession is required.
- One decision is covered by accepted ownership authority.
- Four decisions are direct projections or non-applicable scope.
- No item is blocked.
- Exact authoring, review, and closure boundaries are ready for T07.
- No ADR, Specification, workflow-support, graph, implementation, stage, or commit change occurred.
- Result: `PASS`.
