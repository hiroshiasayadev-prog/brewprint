# PRODUCT-WORK-SPEC-018: Design convergence workflow successor

- **id**: PRODUCT-WORK-SPEC-018
- **status**: done
- **date**: 2026-07-01
- **source_refs**:
  - PRODUCT-REQ-SPEC-005
  - PRODUCT-REQ-SPEC-006
  - PRODUCT-ADR-SPEC-004
  - PRODUCT-ADR-SPEC-005
  - PRODUCT-ADR-SPEC-006
- **impact_refs**:
  - spec:product.design_records.authoring_standards.requirement_authoring
  - spec:product.design_records.authoring_standards.adr_authoring
  - spec:product.design_records.authoring_standards.work_item_authoring
  - spec:product.design_records.authoring_standards.task_authoring
  - spec:product.design_records.authoring_standards.artifact_boundary
  - spec:product.design_records.artifact_model.artifact_responsibility_matrix
- **tasks**:
  - PRODUCT-TASK-SPEC-018-01
  - PRODUCT-TASK-SPEC-018-02
  - PRODUCT-TASK-SPEC-018-03
  - PRODUCT-TASK-SPEC-018-04
  - PRODUCT-TASK-SPEC-018-05
  - PRODUCT-TASK-SPEC-018-06
  - PRODUCT-TASK-SPEC-018-07
  - PRODUCT-TASK-SPEC-018-08
  - PRODUCT-TASK-SPEC-018-09
  - PRODUCT-TASK-SPEC-018-10
  - PRODUCT-TASK-SPEC-018-11
  - PRODUCT-TASK-SPEC-018-12
  - PRODUCT-TASK-SPEC-018-13
  - PRODUCT-TASK-SPEC-018-14
  - PRODUCT-TASK-SPEC-018-15
  - PRODUCT-TASK-SPEC-018-16
  - PRODUCT-TASK-SPEC-018-17
  - PRODUCT-TASK-SPEC-018-18
  - PRODUCT-TASK-SPEC-018-19
  - PRODUCT-TASK-SPEC-018-20

## Goal

Replace the current design-decision workflow skill with one reviewed design-convergence workflow.

The successor must cover design-topic intake through lifecycle, Evidence, and relation synchronization.

## Boundary

This Work Item owns:

- mandatory formal decision-impact Investigation;
- post-Investigation reconciliation and graph repair;
- ADR routing for the decided workflow design;
- coherent ADR-boundary partitioning;
- required ADR creation or amendment;
- successor skill authoring and activation;
- direct canonical Specification projection;
- one integrated independent review;
- finding-driven correction and closure review when required;
- lifecycle, Evidence, and relation synchronization.

This Work Item does not own:

- production implementation;
- DRMCP implementation;
- revision of completed decision entries D-001 through D-023;
- speculative correction or finding-closure Tasks before named findings exist.

## Impact Scope

| target | impact |
|---|---|
| `skills/design-convergence-workflow/` | Becomes the complete successor workflow authority. |
| `skills/design-decision-workflow/` | Remains unchanged until the successor is ready, then is removed without a deprecated stub. |
| `prompt_chappy.md` | Must reference the successor only after replacement authoring is complete. |
| `spec:product.design_records.authoring_standards.requirement_authoring` | Receives the accepted Requirement amendment-versus-new identity boundary. |
| `spec:product.design_records.authoring_standards.adr_authoring` | Receives the accepted amendment-versus-supersession materiality boundary. |
| `spec:product.design_records.authoring_standards.work_item_authoring` | Receives Work Item continuation, review, and conditional graph rules. |
| `spec:product.design_records.authoring_standards.task_authoring` | Receives completed-decision non-writeback and post-completion Task boundaries. |
| `spec:product.design_records.authoring_standards.artifact_boundary` | Receives the final routing and historical-checkpoint boundary. |
| `spec:product.design_records.artifact_model.artifact_responsibility_matrix` | Receives the final completed-record and closure-write ownership boundary. |
| `PRODUCT-ADR-SPEC-006` | Requires a meaning-preserving clarification amendment. |

## Task flow

```text
PRODUCT-TASK-SPEC-018-01 completed decision ledger
  -> PRODUCT-TASK-SPEC-018-02 ADR routing and boundary partitioning
  -> PRODUCT-TASK-SPEC-018-03 PRODUCT-ADR-SPEC-006 clarification amendment
  -> PRODUCT-TASK-SPEC-018-04 new routed ADR-set authoring
  -> PRODUCT-TASK-SPEC-018-05 successor skill authoring and activation
  -> PRODUCT-TASK-SPEC-018-06 canonical Specification synchronization
  -> PRODUCT-TASK-SPEC-018-07 prerequisite gate
     -> NOT READY P-BLK-01: PRODUCT-TASK-SPEC-018-09 graph coordination
        -> PRODUCT-TASK-SPEC-018-10 formal impact Investigation
        -> PRODUCT-TASK-SPEC-018-11 post-Investigation reconciliation decision
        -> PRODUCT-TASK-SPEC-018-12 ADR-routing revalidation
        -> PRODUCT-TASK-SPEC-018-13 bounded authoring-route coordination
        -> PRODUCT-TASK-SPEC-018-14 Work Item decomposition Task-type authoring
        -> resume PRODUCT-TASK-SPEC-018-07 integrated review
           -> NEEDS REVISION F-BLK-01 / F-MAJ-01: PRODUCT-TASK-SPEC-018-15 finding-route coordination
              -> PRODUCT-TASK-SPEC-018-16 ADR materiality decision
                 -> PRODUCT-TASK-SPEC-018-17 ADR amendment-boundary authoring
              -> PRODUCT-TASK-SPEC-018-18 activation correction
              -> PRODUCT-TASK-SPEC-018-19 independent finding-closure review
                 -> CLOSED: PRODUCT-TASK-SPEC-018-20 closure-route coordination
                    -> PRODUCT-TASK-SPEC-018-08 closure synchronization
                 -> OPEN: coordination derives the next exact finding route
```

`PRODUCT-TASK-SPEC-018-01` through `PRODUCT-TASK-SPEC-018-20` exist in the final graph.

## Task Candidates

| task | task type | responsibility | dependency |
|---|---|---|---|
| `PRODUCT-TASK-SPEC-018-01` | `decision` | Own D-001 through D-023 as the completed decision ledger. | none |
| `PRODUCT-TASK-SPEC-018-02` | `decision` | Classify D-001 through D-023 and partition coherent ADR boundaries. | T01 |
| `PRODUCT-TASK-SPEC-018-03` | `authoring` | Amend `PRODUCT-ADR-SPEC-006` without changing its core decision. | T02 |
| `PRODUCT-TASK-SPEC-018-04` | `authoring` | Author the six new ADRs selected by T02. | T03 |
| `PRODUCT-TASK-SPEC-018-05` | `authoring` | Author and activate the successor skill, update the instruction pointer, and remove the old skill only after replacement readiness. | T04 |
| `PRODUCT-TASK-SPEC-018-06` | `authoring` | Project accepted workflow rules into the exact canonical Specification targets. | T04 |
| `PRODUCT-TASK-SPEC-018-07` | `review` | Independently review the final ADR, skill, Specification, and Task-graph state. The first attempt is blocked by prerequisite `P-BLK-01`. | T05 and T06; the resumed route additionally requires T12 through T14. |
| `PRODUCT-TASK-SPEC-018-08` | `synchronization` | Synchronize lifecycle, Evidence, relations, and Work Item closure after accepted finding closure. | T19 and T20 |
| `PRODUCT-TASK-SPEC-018-09` | `coordination` | Repair the missing mandatory Investigation and reconvergence ownership route. | T06. T07 `P-BLK-01` is trigger Evidence, not a dependency. |
| `PRODUCT-TASK-SPEC-018-10` | `investigation` | Produce one formal impact and conflict Investigation for D-001 through D-023. | T09 |
| `PRODUCT-TASK-SPEC-018-11` | `decision` | Decide the post-Investigation reconciliation and graph disposition. | T10 |
| `PRODUCT-TASK-SPEC-018-12` | `decision` | Revalidate T02 ADR routing and ADR boundaries against the formal Investigation. | T11 |
| `PRODUCT-TASK-SPEC-018-13` | `coordination` | Materialize one bounded T14 authoring route and release T07 behind the T14 dependency. | T12 |
| `PRODUCT-TASK-SPEC-018-14` | `authoring` | Add `work_item_decomposition` across ADR-004, ADR-005, ADR-010, authoring Specifications, and workflow support. | T13 |
| `PRODUCT-TASK-SPEC-018-15` | `coordination` | Materialize the exact F-BLK-01 and F-MAJ-01 repair route. | T07 |
| `PRODUCT-TASK-SPEC-018-16` | `decision` | Decide whether the responsibility extraction is a material ownership or architecture change. | T15 |
| `PRODUCT-TASK-SPEC-018-17` | `authoring` | Project the accepted ADR amendment materiality boundary into ADR-006, `adr_authoring`, and `adr-routing.md`. | T16 |
| `PRODUCT-TASK-SPEC-018-18` | `correction` | Repair the missing decomposition companion activation and ownership summary. | T15 |
| `PRODUCT-TASK-SPEC-018-19` | `review` | Independently decide closure of F-BLK-01 and F-MAJ-01. | T17 and T18 |
| `PRODUCT-TASK-SPEC-018-20` | `coordination` | Materialize T08 as the accepted closure-synchronization owner. | T19 |

The accepted finding route and closure route are complete.

## Completion Condition

- One formal Investigation records affected artifacts, semantic conflicts, graph-change candidates, shared-writer candidates, uncertainty, and Evidence for D-001 through D-023.
- Post-Investigation reconciliation, ADR-routing revalidation, and bounded successor authoring complete before integrated review resumes.
- Every D-001 through D-023 decision has a reviewed routing trace.
- Every required ADR is accepted.
- `PRODUCT-ADR-SPEC-006` preserves its core decision and contains no stale decision-Task writeback consequence.
- The successor skill covers the complete decided workflow.
- The old skill is removed only after successor activation is complete.
- Every normative workflow rule is projected into the exact canonical Specification target.
- One integrated independent review returns `PASS`, or every required finding is independently closed.
- Lifecycle, Evidence, and relations express the same accepted result.
- No production implementation is performed.

## Evidence

- `PRODUCT-TASK-SPEC-018-01` owns D-001 through D-023 as the completed decision ledger.
- `skills/design-convergence-workflow/decision-ledger.md` remains a temporary duplicate and may be deleted after downstream authoring no longer needs it.
- No earlier Design Record Task owned ADR routing for this successor workflow.
- `PRODUCT-TASK-SPEC-018-02` is the routing owner.
- `PRODUCT-TASK-SPEC-018-03` completed the `PRODUCT-ADR-SPEC-006` clarification amendment.
- `PRODUCT-TASK-SPEC-018-04` authored `PRODUCT-ADR-SPEC-009` through `PRODUCT-ADR-SPEC-014`.
- `PRODUCT-TASK-SPEC-018-05` authored and activated `skills/design-convergence-workflow/`, updated `prompt_chappy.md`, and removed the old skill from the repository path.
- Existing `PRODUCT-ADR-SPEC-004`, `PRODUCT-ADR-SPEC-005`, and `PRODUCT-ADR-SPEC-006` were checked for coverage and conflict.
- `PRODUCT-TASK-SPEC-018-06` completed canonical Specification synchronization.
- Changed Specification targets: `requirement_authoring`, `work_item_authoring`, `task_authoring`, `artifact_boundary`, and `artifact_responsibility_matrix`.
- `PRODUCT-TASK-SPEC-018-07` returned `NOT READY` because mandatory Investigation ownership and Evidence were absent.
- T07 prerequisite `P-BLK-01` routes the missing-owner repair to `coordination` without immediate user judgment.
- `PRODUCT-TASK-SPEC-018-09` completed the missing-owner graph repair.
- `PRODUCT-TASK-SPEC-018-10` owns the formal `PRODUCT-INV-SPEC-006` Investigation.
- `PRODUCT-TASK-SPEC-018-11` owns post-Investigation reconciliation.
- `PRODUCT-TASK-SPEC-018-12` owns the completed ADR-routing revalidation.
- `PRODUCT-TASK-SPEC-018-13` materialized one bounded authoring route and aligned T07.
- `PRODUCT-TASK-SPEC-018-14` completed the ADR, Specification, and workflow-support projection.
- No migration or follow-up Work Item exists.
- `PRODUCT-TASK-SPEC-018-07` completed integrated review with `NEEDS REVISION`.
- T07 finding F-BLK-01 requires an explicit ADR materiality decision.
- T07 finding F-MAJ-01 requires a bounded activation correction.
- `PRODUCT-TASK-SPEC-018-15` materialized T16 through T19.
- `PRODUCT-TASK-SPEC-018-16` decided that the responsibility extraction is non-material and preserves `amend` for ADR-004, ADR-005, and ADR-010.
- `PRODUCT-TASK-SPEC-018-17` aligned ADR-006, `adr_authoring`, and `adr-routing.md`.
- `PRODUCT-TASK-SPEC-018-18` corrected the active decomposition companion pointer and ownership summary.
- `PRODUCT-TASK-SPEC-018-19` independently closed F-BLK-01 and F-MAJ-01 with no direct regression.
- `PRODUCT-TASK-SPEC-018-20` materialized T08 as the closure-synchronization owner.
- `PRODUCT-TASK-SPEC-018-08` synchronized the accepted review route and closure state.
- Every Completion Condition passed.
- W018 is closed as `done`.
- No production implementation, stage, or commit was performed.
