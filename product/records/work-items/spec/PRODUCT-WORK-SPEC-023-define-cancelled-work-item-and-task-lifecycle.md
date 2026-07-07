# PRODUCT-WORK-SPEC-023: Define cancelled Work Item and Task lifecycle

- **id**: PRODUCT-WORK-SPEC-023
- **status**: done
- **date**: 2026-07-03
- **source_refs**:
  - PRODUCT-REQ-SPEC-009
- **impact_refs**:
  - spec:product.design_records.authoring_standards.work_item_authoring
  - spec:product.design_records.authoring_standards.task_authoring
  - spec:product.responsibility_boundary_validator
  - PRODUCT-ADR-SPEC-005
  - PRODUCT-ADR-SPEC-010
  - PRODUCT-ADR-SPEC-011
  - PRODUCT-ADR-SPEC-014
  - PRODUCT-ADR-SPEC-017
  - PRODUCT-ADR-SPEC-018
- **tasks**:
  - PRODUCT-TASK-SPEC-023-01
  - PRODUCT-TASK-SPEC-023-02
  - PRODUCT-TASK-SPEC-023-03
  - PRODUCT-TASK-SPEC-023-04
  - PRODUCT-TASK-SPEC-023-05
  - PRODUCT-TASK-SPEC-023-06
  - PRODUCT-TASK-SPEC-023-07
  - PRODUCT-TASK-SPEC-023-08
  - PRODUCT-TASK-SPEC-023-09
  - PRODUCT-TASK-SPEC-023-10
  - PRODUCT-TASK-SPEC-023-11

## Goal

Define one canonical `cancelled` lifecycle contract for Work Items and Tasks.

Define direct cancellation, terminal-state behavior, and Work Item-to-Task cancellation propagation without treating cancellation as successful completion.

## Boundary

This Work Item owns:

- canonical `cancelled` meaning for Work Items and Tasks;
- allowed cancellation transitions;
- direct Task cancellation semantics;
- Work Item-to-owned-Task propagation;
- effects on dependency and `work_item_execution` boundaries;
- cancellation Evidence and validation requirements;
- ADR routing and canonical Specification projection;
- one final integrated independent review;
- final lifecycle, Evidence, and relation synchronization.

This Work Item does not own:

- migration of existing Work Items or Tasks;
- automatic cancellation of child Work Items or transitive descendants;
- the framing workflow that may choose cancellation;
- concrete DRMCP, command, transaction, or implementation mechanics;
- production implementation;
- stage or commit work.

## Impact Scope

| target | impact |
|---|---|
| `spec:product.design_records.authoring_standards.work_item_authoring` | Add the Work Item status, transitions, propagation, and Evidence contract. |
| `spec:product.design_records.authoring_standards.task_authoring` | Add the Task status, transitions, dependency effects, and type-specific consequences. |
| ADR authority | Record durable lifecycle and propagation choices when routing requires an ADR. |
| Workflow support and derived validation assets | Reflect accepted cancellation semantics where the Investigation confirms direct impact. |

## Task flow

```text
PRODUCT-TASK-SPEC-023-01 cancellation contract decision
  -> PRODUCT-TASK-SPEC-023-02 impact and conflict Investigation
  -> PRODUCT-TASK-SPEC-023-03 post-Investigation graph coordination
  -> PRODUCT-TASK-SPEC-023-04 post-Investigation reconciliation decision
  -> PRODUCT-TASK-SPEC-023-05 post-reconciliation graph coordination
  -> PRODUCT-TASK-SPEC-023-06 ADR routing and boundary partitioning
  -> PRODUCT-TASK-SPEC-023-07 post-routing authoring and review coordination
  -> PRODUCT-TASK-SPEC-023-08 ADR authoring
  -> PRODUCT-TASK-SPEC-023-09 canonical and workflow-support authoring
  -> PRODUCT-TASK-SPEC-023-10 integrated independent review
  -> PRODUCT-TASK-SPEC-023-11 direct-PASS closure synchronization
```

T03 materializes only the reconciliation route required by T02 Evidence.
T05 materializes T06 ADR routing and T07 post-routing coordination after T04 resolves every Investigation judgment.
T07 materializes concrete writers only after T06 fixes exact routing outcomes.
T11 is valid only for a direct T10 `PASS`; finding-driven branches remain unmaterialized.

Correction and finding-closure review Tasks are created only after named findings exist.

## Task Candidates

| task | task type | responsibility | dependency |
|---|---|---|---|
| `PRODUCT-TASK-SPEC-023-01` | `decision` | Fix the semantic cancellation, transition, propagation, dependency, and Evidence contract. | none |
| `PRODUCT-TASK-SPEC-023-02` | `investigation` | Inventory affected authority, conflicts, graph changes, and shared writers. | T01 |
| `PRODUCT-TASK-SPEC-023-03` | `coordination` | Materialize the exact post-Investigation reconciliation route. | T02 |
| `PRODUCT-TASK-SPEC-023-04` | `decision` | Resolve cancelled-state body readiness, propagation ownership, and validator invocation. | T02 and T03 |
| `PRODUCT-TASK-SPEC-023-05` | `coordination` | Materialize the ADR-routing and post-routing coordination owners. | T04 |
| `PRODUCT-TASK-SPEC-023-06` | `decision` | Classify every cancellation decision and partition coherent ADR boundaries. | T04 and T05 |
| `PRODUCT-TASK-SPEC-023-07` | `coordination` | Materialize exact ADR, canonical authoring, review, and accepted-route closure Tasks. | T06 |
| `PRODUCT-TASK-SPEC-023-08` | `authoring` | Create ADR-018 and amend ADR-005, ADR-010, ADR-011, ADR-014, and ADR-017. | T06 and T07 |
| `PRODUCT-TASK-SPEC-023-09` | `authoring` | Update Task, Work Item, validator, and workflow-support contracts. | T08 |
| `PRODUCT-TASK-SPEC-023-10` | `review` | Independently review the final combined cancellation design. | T09 |
| `PRODUCT-TASK-SPEC-023-11` | `synchronization` | Close W023 only after direct integrated-review PASS. | T10 |

## Completion Condition

- Work Item and Task status sets include terminal `cancelled`.
- Cancellation remains distinct from `done` and does not satisfy the owned completion condition.
- Allowed direct transitions are explicit.
- Work Item cancellation changes every owned unfinished Task to `cancelled`.
- Owned Tasks already `done` remain unchanged.
- Dependency and `work_item_execution` consequences are explicit.
- Cancellation Evidence and validation requirements are explicit.
- Existing-record migration, descendant cancellation, framing, and implementation remain excluded.
- Required ADR and canonical Specification projections agree.
- One integrated independent review returns `PASS`, or every required finding is independently `CLOSED`.
- Lifecycle, Evidence, relations, and Work Item closure express the accepted result.

## Evidence

- PRODUCT-REQ-SPEC-009 is accepted and directly requires a terminal `cancelled` status for Work Items and Tasks.
- The Requirement directly fixes Work Item-to-owned-Task propagation for `not_started`, `in_progress`, and `blocked` Tasks.
- The Requirement preserves owned Tasks already `done`.
- The Requirement excludes existing-record migration, descendant Work Item cancellation, framing, and concrete implementation mechanics.
- Work Item and Task authoring Specifications now define terminal `cancelled` and its accepted lifecycle contract.
- DRMCP is non-operational. Filesystem authoring was the required fallback.
- T01 completed the initial semantic cancellation contract.
- T02 created concluded PRODUCT-INV-SPEC-010.
- The Investigation found unresolved cancelled-state body readiness, propagation ownership, and validator invocation.
- T03 created T04 as the sole reconciliation decision owner and T05 as its successor graph owner.
- T04 completed the post-Investigation cancellation contract.
- T05 created T06 as the sole ADR-routing owner and T07 as the sole post-routing graph owner.
- T06 completed ADR routing with one new ADR and five amendments.
- T07 created T08 through T11 with deterministic writer, review, and direct-PASS closure order.
- T08 created PRODUCT-ADR-SPEC-018 and amended PRODUCT-ADR-SPEC-005, 010, 011, 014, and 017.
- T09 projected the accepted contract into Task authoring, Work Item authoring, the responsibility-boundary validator, and the three affected workflow-support files.
- T10 independently reviewed the final combined design and returned `PASS` with no Blocking, Major, Minor, or closure-blocking finding.
- T10 Advisory A-001 requires no correction or additional owner.
- T11 verified every Completion Condition and synchronized the direct-PASS closure route.
- Accepted canonical Specifications:
  - `spec:product.design_records.authoring_standards.task_authoring`;
  - `spec:product.design_records.authoring_standards.work_item_authoring`;
  - `spec:product.responsibility_boundary_validator`.
- Accepted workflow-support outputs:
  - `skills/design-convergence-workflow/work-item-execution.md`;
  - `skills/design-convergence-workflow/graph-coordination.md`;
  - `skills/design-convergence-workflow/closure-synchronization.md`.
- Correction and finding-closure review Tasks were not materialized because no qualifying finding exists.
- PRODUCT-TASK-SPEC-023-11 and PRODUCT-WORK-SPEC-023 are `done`.
- Result: `PASS`.
