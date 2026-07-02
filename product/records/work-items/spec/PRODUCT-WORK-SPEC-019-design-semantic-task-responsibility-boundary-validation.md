# PRODUCT-WORK-SPEC-019: Design semantic Task responsibility-boundary validation

- **id**: PRODUCT-WORK-SPEC-019
- **status**: done
- **date**: 2026-07-01
- **source_refs**:
  - PRODUCT-REQ-SPEC-007
- **impact_refs**:
  - PRODUCT-REQ-SPEC-005
  - PRODUCT-ADR-SPEC-001
  - spec:product.design_records.authoring_standards.task_authoring
  - spec:product.responsibility_boundary_validator
  - PRODUCT-WORK-SPEC-020
  - PRODUCT-WORK-SPEC-021
- **tasks**:
  - PRODUCT-TASK-SPEC-019-01
  - PRODUCT-TASK-SPEC-019-02
  - PRODUCT-TASK-SPEC-019-03
  - PRODUCT-TASK-SPEC-019-04
  - PRODUCT-TASK-SPEC-019-05
  - PRODUCT-TASK-SPEC-019-06
  - PRODUCT-TASK-SPEC-019-07
  - PRODUCT-TASK-SPEC-019-08
  - PRODUCT-TASK-SPEC-019-09
  - PRODUCT-TASK-SPEC-019-10
  - PRODUCT-TASK-SPEC-019-11
  - PRODUCT-TASK-SPEC-019-12
  - PRODUCT-TASK-SPEC-019-13
  - PRODUCT-TASK-SPEC-019-14
  - PRODUCT-TASK-SPEC-019-15
  - PRODUCT-TASK-SPEC-019-16
  - PRODUCT-TASK-SPEC-019-17
  - PRODUCT-TASK-SPEC-019-18
  - PRODUCT-TASK-SPEC-019-19
  - PRODUCT-TASK-SPEC-019-20
  - PRODUCT-TASK-SPEC-019-21
  - PRODUCT-TASK-SPEC-019-22

## Goal

Define and independently validate the semantic contract for a temporary standalone Task responsibility-boundary validation tool required by `PRODUCT-REQ-SPEC-007`.

The contract must produce criterion-level binary judgments with Task-local supporting reasons without claiming complete executor readiness.

The temporary tool is independent from DRMCP. Future DRMCP integration is a separate design and delivery boundary.

## Boundary

This Work Item owns:

- the decision inventory and interactive decision loop for the semantic validation contract;
- the semantic composition and selection rules for responsibility-boundary criteria;
- criterion-level binary judgment semantics;
- Task-local rationale and evidence semantics for each criterion judgment;
- overall compliance aggregation semantics;
- the boundary between semantic non-compliance and validator execution failure;
- the temporary standalone tool boundary;
- mandatory impact and conflict investigation;
- Requirement and Specification reconciliation when needed;
- ADR routing and conditional ADR authoring;
- canonical PRODUCT Specification authoring;
- one integrated independent review;
- finding-driven correction and independent finding-closure review when required;
- lifecycle, Evidence, and relation synchronization.

This Work Item does not own:

- concrete language model, provider, or runtime selection;
- concrete MCP request, response, or error schemas;
- exact checklist item wording;
- checklist storage format;
- DRMCP integration;
- changes to existing DRMCP tools, Specifications, diagnostics, or implementation;
- production implementation;
- complete executor-readiness validation;
- multi-Task or Work Item graph validation;
- automatic Task correction or rewriting;
- migration of existing Task records.

## Impact Scope

| target | impact |
|---|---|
| `PRODUCT-REQ-SPEC-007` | Supplies the accepted validation need and compliance-claim boundary. |
| `PRODUCT-REQ-SPEC-005` | Supplies the canonical typed single-responsibility Task contract being evaluated. |
| `PRODUCT-ADR-SPEC-001` | Receives the bounded PRODUCT-root ownership amendment required for the standalone validator area. |
| `spec:product.design_records.authoring_standards.task_authoring` | Supplies Task-type outcomes, completion judgments, prohibited overlaps, and section-alignment rules. |
| `spec:product.responsibility_boundary_validator` | Will own current normative criterion, aggregation, rationale, failure-boundary, and standalone-tool semantics. |
| `PRODUCT-WORK-SPEC-020` | Owns exact checklist wording, storage format, placement, authoring, and independent review. |
| `PRODUCT-WORK-SPEC-021` | Owns temporary standalone validator implementation after W019 design closure and accepted W020 checklist authoring. |
| Future DRMCP integration | Separate future Requirement or Work Item. W019 does not modify or constrain the current DRMCP tool surface. |

## Task flow

```text
PRODUCT-TASK-SPEC-019-01 initial decision inventory and interactive decision loop
  -> PRODUCT-TASK-SPEC-019-02 graph coordination for corrected tool boundary
  -> PRODUCT-TASK-SPEC-019-03 successor decision for temporary standalone scope
  -> PRODUCT-TASK-SPEC-019-04 mandatory Investigation route coordination
  -> PRODUCT-TASK-SPEC-019-05 impact and conflict Investigation
  -> PRODUCT-TASK-SPEC-019-06 reconciliation-route coordination
  -> PRODUCT-TASK-SPEC-019-07 post-Investigation reconciliation decision
  -> PRODUCT-TASK-SPEC-019-08 successor split-route coordination
  -> PRODUCT-TASK-SPEC-019-09 checklist-versus-implementation Work Item decision
  -> PRODUCT-TASK-SPEC-019-10 Work Item decomposition
  -> PRODUCT-TASK-SPEC-019-11 post-decomposition route coordination
  -> PRODUCT-TASK-SPEC-019-12 ADR routing and boundary partitioning
  -> PRODUCT-TASK-SPEC-019-13 post-routing authoring and review coordination
  -> PRODUCT-TASK-SPEC-019-14 author ADR-015 through ADR-017
  -> PRODUCT-TASK-SPEC-019-15 resolve PRODUCT-REQ-SPEC-007 amendment disposition as no amendment
  -> PRODUCT-TASK-SPEC-019-16 align PRODUCT root ownership and author validator Specification plus parent registration
  -> PRODUCT-TASK-SPEC-019-17 author narrow task_authoring usage rule
  -> PRODUCT-TASK-SPEC-019-18 integrated independent review
     -> PASS: PRODUCT-TASK-SPEC-019-19 closure synchronization
     -> NEEDS REVISION:
        -> explicit user acceptance of every named finding as a non-blocking workflow exception: PRODUCT-TASK-SPEC-019-19 closure synchronization
        -> otherwise finding-specific coordination
           -> correction or successor decision
           -> independent finding-closure review or new integrated review
           -> closure synchronization

PRODUCT-TASK-SPEC-019-20 early-release reconsideration coordination
  -> PRODUCT-TASK-SPEC-019-21 successor release-order decision
  -> PRODUCT-TASK-SPEC-019-22 cross-Work-Item graph coordination
     -> PRODUCT-WORK-SPEC-020 decision, Investigation, and checklist authoring may proceed early
     -> PRODUCT-WORK-SPEC-020 integrated review waits for accepted W019 review
     -> PRODUCT-WORK-SPEC-021 waits for W019 closure and accepted W020 review
```

`PRODUCT-TASK-SPEC-019-01` preserves the initial completed decision checkpoint.
`PRODUCT-TASK-SPEC-019-02` and `PRODUCT-TASK-SPEC-019-03` preserve the later correction from a presumed DRMCP integration boundary to a temporary standalone tool boundary.
`PRODUCT-TASK-SPEC-019-04` materializes the mandatory Investigation owner without performing Investigation work.
`PRODUCT-TASK-SPEC-019-05` owns one formal Investigation record for the corrected standalone-tool boundary.
`PRODUCT-TASK-SPEC-019-06` materializes the missing W019 reconciliation owner without performing the decision loop.
`PRODUCT-TASK-SPEC-019-07` owns the six original W019 reconciliation judgments.
`PRODUCT-TASK-SPEC-019-08` preserves T07 and materializes the successor decision and decomposition route.
`PRODUCT-TASK-SPEC-019-09` fixes checklist authoring and standalone implementation as separate Work Items.
`PRODUCT-TASK-SPEC-019-10` creates PRODUCT-WORK-SPEC-020 and PRODUCT-WORK-SPEC-021.
`PRODUCT-TASK-SPEC-019-11` materializes the ADR-routing and post-routing coordination owners.
`PRODUCT-TASK-SPEC-019-12` owns complete ADR routing and boundary partitioning.
`PRODUCT-TASK-SPEC-019-13` materialized T14 through T19.
`PRODUCT-TASK-SPEC-019-14` through `PRODUCT-TASK-SPEC-019-17` own the ordered authoring route.
`PRODUCT-TASK-SPEC-019-18` owns the one integrated independent review.
`PRODUCT-TASK-SPEC-019-19` owns verdict-gated closure synchronization.
W018 T11 J-001 is fixed input for MC-002 and is not reopened by W019.
PRODUCT-TASK-SPEC-018-19 independently closed F-BLK-01 and F-MAJ-01.
The shared `task_authoring` writer is therefore released after T16.
W019 closure does not wait for W020 or W021 completion.
PRODUCT-TASK-SPEC-019-21 supersedes only the current W020 start timing from T09.
W020 decision, Investigation, and checklist authoring may proceed before W019 closure.
W020 integrated review waits for accepted W019 integrated review.
W021 starts only after W019 design closure and accepted W020 review.
Correction and finding-closure review Tasks are not created before named findings exist.

## Task Candidates

| task | task type | responsibility | dependency |
|---|---|---|---|
| `PRODUCT-TASK-SPEC-019-01` | `decision` | Preserve the initial semantic responsibility-boundary validation decisions. | none |
| `PRODUCT-TASK-SPEC-019-02` | `coordination` | Correct the Work Item graph after the temporary standalone tool boundary was clarified. | T01 |
| `PRODUCT-TASK-SPEC-019-03` | `decision` | Fix the temporary standalone tool boundary and defer DRMCP integration. | T02 |
| `PRODUCT-TASK-SPEC-019-04` | `coordination` | Materialize the mandatory impact Investigation route and reserve its bounded output. | T03 |
| `PRODUCT-TASK-SPEC-019-05` | `investigation` | Inventory affected PRODUCT authority, temporary-tool obligations, conflicts, graph candidates, shared writers, and uncertainty. | T04 |
| `PRODUCT-TASK-SPEC-019-06` | `coordination` | Materialize the missing post-Investigation reconciliation decision owner. | T05 |
| `PRODUCT-TASK-SPEC-019-07` | `decision` | Resolve the six original W019 reconciliation judgments without selecting an ADR route. | T06 |
| `PRODUCT-TASK-SPEC-019-08` | `coordination` | Materialize the successor decision and Work Item decomposition route without rewriting T07. | T07 |
| `PRODUCT-TASK-SPEC-019-09` | `decision` | Fix checklist authoring and standalone implementation as separate downstream Work Item boundaries. | T08 |
| `PRODUCT-TASK-SPEC-019-10` | `work_item_decomposition` | Create PRODUCT-WORK-SPEC-020 and PRODUCT-WORK-SPEC-021 with non-overlapping responsibilities. | T09 |
| `PRODUCT-TASK-SPEC-019-11` | `coordination` | Materialize T12 ADR routing and T13 post-routing graph ownership without speculative authoring Tasks. | T10 |
| `PRODUCT-TASK-SPEC-019-12` | `decision` | Classify every terminal W019 decision and partition coherent ADR boundaries. | T11 |
| `PRODUCT-TASK-SPEC-019-13` | `coordination` | Materialize exact authoring, shared-writer, integrated-review, and closure Tasks from the completed T12 route. | T12 |
| `PRODUCT-TASK-SPEC-019-14` | `authoring` | Author PRODUCT-ADR-SPEC-015 through PRODUCT-ADR-SPEC-017. | T13 |
| `PRODUCT-TASK-SPEC-019-15` | `authoring` | Resolve PRODUCT-REQ-SPEC-007 amendment disposition and preserve the Requirement unchanged when downstream policy remains ADR/Specification-owned. | T14 |
| `PRODUCT-TASK-SPEC-019-16` | `authoring` | Align PRODUCT-ADR-SPEC-001, author the standalone validator Specification, and register it directly under `spec:product`. | T15 |
| `PRODUCT-TASK-SPEC-019-17` | `authoring` | Add the narrow validator ownership and usage rule to `task_authoring`. | T16 and W018 T19 |
| `PRODUCT-TASK-SPEC-019-18` | `review` | Independently review the final combined W019 state. | T14 through T17 |
| `PRODUCT-TASK-SPEC-019-19` | `synchronization` | Propagate the accepted review route into W019 lifecycle, Evidence, relations, and closure. | T18 or later accepted finding-closure route |
| `PRODUCT-TASK-SPEC-019-20` | `coordination` | Materialize a successor decision and graph route for W020 early-start reconsideration. | T09 and T15 |
| `PRODUCT-TASK-SPEC-019-21` | `decision` | Permit W020 decision, Investigation, and authoring before W019 closure while gating W020 review. | T20 |
| `PRODUCT-TASK-SPEC-019-22` | `coordination` | Apply the revised cross-Work-Item release route and create the W020 internal graph owner. | T21 |

## Completion Condition

- Every owned decision is `decided`, `deferred`, or validly `blocked`.
- The mandatory impact and conflict Investigation is complete.
- Criterion composition and `task_type` selection semantics are canonical.
- Every evaluated criterion returns one binary judgment.
- Every criterion judgment returns Task-local supporting evidence or a concise reason.
- Overall compliance aggregation is canonical.
- Missing Task-local evidence is handled without external inference.
- Semantic non-compliance is distinct from validator execution failure.
- The compliance claim remains limited to one Task record and its responsibility boundary.
- The accepted product shape is a temporary standalone tool, not a DRMCP integration.
- Its canonical Specification is a direct child of `spec:product`, not a child of `spec:product.design_records`.
- Exact checklist wording, storage format, placement, authoring, and checklist review belong to PRODUCT-WORK-SPEC-020.
- Concrete app namespace, source placement, model, provider, runtime, retry, implementation errors, and executable delivery belong to PRODUCT-WORK-SPEC-021.
- Current DRMCP tools, Specifications, diagnostics, and implementation remain unchanged.
- Required ADRs are accepted.
- Canonical PRODUCT Specifications contain the final accepted contract.
- One integrated independent review returns `PASS`, every required finding is independently closed, or the user explicitly accepts every named finding as a non-blocking workflow exception.
- Lifecycle, Evidence, and relations express the accepted result.
- No production implementation or existing-record migration is performed.

## Evidence

- `PRODUCT-REQ-SPEC-007` is accepted and requires semantic Task responsibility-boundary validation.
- The Requirement requires an overall machine-readable result.
- The Requirement requires one binary judgment for each evaluated criterion.
- The Requirement requires Task-local supporting evidence or a concise reason for each criterion judgment.
- `PRODUCT-REQ-SPEC-005` and the canonical Task authoring Specification define the responsibility model being evaluated.
- `PRODUCT-TASK-SPEC-019-01` owns the initial decision inventory and interactive decision loop.
- The user later clarified that the validator is a temporary standalone tool and that DRMCP integration belongs to future work.
- `PRODUCT-TASK-SPEC-019-02` preserves the graph correction without reopening T01.
- `PRODUCT-TASK-SPEC-019-03` owns the corrected standalone-tool decision boundary.
- `PRODUCT-TASK-SPEC-019-04` owns materialization of the mandatory Investigation route.
- `PRODUCT-TASK-SPEC-019-05` owns one bounded formal Investigation.
- `PRODUCT-INV-SPEC-006` is already owned by W018; W019 reserves `PRODUCT-INV-SPEC-007` instead.
- `PRODUCT-TASK-SPEC-019-05` and concluded `PRODUCT-INV-SPEC-007` triggered the post-Investigation reconciliation route.
- `PRODUCT-TASK-SPEC-019-06` materialized T07 without executing its decision loop.
- `PRODUCT-TASK-SPEC-019-07` owns exactly R-001 through R-006 as the completed original reconciliation checkpoint.
- The user later classified checklist creation as authoring and required a separate Work Item from implementation.
- `PRODUCT-TASK-SPEC-019-08` materialized the successor route without rewriting T07.
- `PRODUCT-TASK-SPEC-019-09` fixed the two independent Work Item boundaries.
- `PRODUCT-TASK-SPEC-019-10` created PRODUCT-WORK-SPEC-020 and PRODUCT-WORK-SPEC-021.
- `PRODUCT-TASK-SPEC-019-11` created T12 and T13 without pre-creating conditional ADR or finding Tasks.
- `PRODUCT-TASK-SPEC-019-12` completed the three-ADR routing ledger.
- `PRODUCT-TASK-SPEC-019-13` materialized T14 through T19.
- `PRODUCT-TASK-SPEC-019-14` completed ADR authoring.
- `PRODUCT-TASK-SPEC-019-15` completed with a no-amendment disposition for PRODUCT-REQ-SPEC-007.
- `PRODUCT-TASK-SPEC-019-16` is the next released owner.
- PRODUCT-WORK-SPEC-020 owns checklist authoring and independent checklist review.
- PRODUCT-TASK-SPEC-019-21 permits W020 decision, Investigation, and authoring before W019 closure.
- PRODUCT-TASK-SPEC-019-22 applies the revised release graph without reopening T09.
- PRODUCT-WORK-SPEC-021 owns temporary standalone validator implementation and consumes accepted W020 checklist artifacts.
- MC-001 requires W019 reconciliation before canonical authoring.
- MC-002 consumes `PRODUCT-TASK-SPEC-018-11` J-001 as fixed authority.
- W019 does not reopen the accepted `work_item_decomposition` and `coordination` split.
- T11 completed the required post-T10 coordination.
- ADR routing is materialized as T12.
- Exact authoring and review graph materialization is complete through T13.
- PRODUCT-TASK-SPEC-018-19 closed the external W018 findings.
- Shared `task_authoring` writing remains serialized after T16 and the completed W018 T19 gate.
- T18 directly reviewed the current combined state through T22 despite the incomplete persisted dependency order.
- T18 returned `NEEDS REVISION` with F-MAJ-01 through F-MAJ-03.
- The user explicitly accepted all three findings as non-blocking workflow exceptions because the reviewed semantic design and canonical projection were correct and no unreviewed graph state remained.
- T19 preserved the T18 verdict and finding set, recorded the user disposition, and synchronized W019 closure.
- Closure remains recorded through T19.
