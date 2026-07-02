# PRODUCT-WORK-SPEC-020: Author Task responsibility-boundary checklists

- **id**: PRODUCT-WORK-SPEC-020
- **status**: done
- **date**: 2026-07-01
- **source_refs**:
  - PRODUCT-TASK-SPEC-019-10
- **impact_refs**:
  - spec:product.responsibility_boundary_validator
  - spec:product.design_records.authoring_standards.task_authoring
- **tasks**:
  - PRODUCT-TASK-SPEC-020-01
  - PRODUCT-TASK-SPEC-020-02
  - PRODUCT-TASK-SPEC-020-03
  - PRODUCT-TASK-SPEC-020-04
  - PRODUCT-TASK-SPEC-020-05
  - PRODUCT-TASK-SPEC-020-06

## Goal

Produce one reviewed checklist artifact set for semantic Task responsibility-boundary validation.

The artifact set must provide common criteria and Task-type-specific criteria for every canonical Task type.

## Boundary

This Work Item owns:

- exact checklist item wording;
- common responsibility-boundary criteria;
- Task-type-specific criteria;
- checklist composition and selection projection from the accepted validator contract;
- checklist storage format and repository placement;
- checklist artifact authoring;
- independent integrated review;
- finding-driven correction and closure review when required;
- lifecycle, Evidence, and relation synchronization.

This Work Item does not own:

- the W019 validator product contract;
- changes to PRODUCT-REQ-SPEC-007;
- temporary standalone validator implementation;
- language model, provider, runtime, retry, timeout, or implementation error policy;
- DRMCP integration;
- automatic Task correction;
- migration of existing Task records.

## Impact Scope

| target | impact |
|---|---|
| `spec:product.responsibility_boundary_validator` | Supplies checklist composition, criterion, rationale, and result semantics. |
| `spec:product.design_records.authoring_standards.task_authoring` | Supplies the canonical Task types, owned outcomes, completion judgments, and prohibited overlaps. |
| Checklist artifact set | Receives the exact common and Task-type-specific checklist content and storage contract. |
| `PRODUCT-WORK-SPEC-021` | Consumes the accepted checklist artifact set during implementation. |

## Task flow

```text
PRODUCT-TASK-SPEC-019-22 early-release graph coordination
  -> PRODUCT-TASK-SPEC-020-01 internal graph coordination
  -> PRODUCT-TASK-SPEC-020-02 checklist artifact contract decision
  -> PRODUCT-TASK-SPEC-020-03 checklist coverage and placement Investigation
  -> PRODUCT-TASK-SPEC-022-02 work_item_execution canonical authority
  -> PRODUCT-TASK-SPEC-022-03 W020 checklist release coordination
  -> PRODUCT-TASK-SPEC-020-04 checklist artifact authoring
  -> accepted W019 integrated review route
  -> PRODUCT-TASK-SPEC-020-05 integrated independent checklist review
     -> PASS: PRODUCT-TASK-SPEC-020-06 closure synchronization
     -> NEEDS REVISION: finding-specific coordination, correction, and independent closure review
```

W020 decision, Investigation, and authoring may proceed before W019 closure.
W020 integrated review waits for accepted W019 integrated review.
PRODUCT-WORK-SPEC-021 remains blocked until W019 closure and accepted W020 review.

## Task Candidates

| task | task type | responsibility | dependency |
|---|---|---|---|
| `PRODUCT-TASK-SPEC-020-01` | `coordination` | Materialize the W020 internal decision, Investigation, authoring, review, and closure graph. | PRODUCT-TASK-SPEC-019-22 |
| `PRODUCT-TASK-SPEC-020-02` | `decision` | Fix checklist source format, placement, partitioning, and criterion schema. | T01 |
| `PRODUCT-TASK-SPEC-020-03` | `investigation` | Check every canonical Task type and affected authority for checklist coverage, placement effects, and conflict. | T02 |
| `PRODUCT-TASK-SPEC-020-04` | `authoring` | Author the complete common and Task-type-specific checklist artifact set. | T03 and PRODUCT-TASK-SPEC-022-03 |
| `PRODUCT-TASK-SPEC-020-05` | `review` | Independently review checklist completeness, type coverage, and contract alignment. | T04 and accepted W019 review |
| `PRODUCT-TASK-SPEC-020-06` | `synchronization` | Synchronize accepted lifecycle, Evidence, and relations. | accepted T05 route |

Correction and finding-closure review Tasks are created only after named findings exist.

## Completion Condition

- One accepted checklist storage and placement contract exists.
- Common responsibility-boundary criteria are authored.
- Every canonical Task type has an exact checklist projection.
- Each criterion supports one binary judgment and one concise reason.
- Checklist content remains limited to Task responsibility boundaries.
- Checklist artifacts align with the accepted W019 validator contract.
- Checklist artifacts are consumable by PRODUCT-WORK-SPEC-021 without content inference.
- One integrated independent review returns `PASS`, or every required finding is independently closed.
- Lifecycle, Evidence, and relations express the accepted result.
- No validator implementation or DRMCP integration is performed.

## Evidence

- PRODUCT-TASK-SPEC-019-09 classified checklist creation as authoring rather than implementation.
- PRODUCT-TASK-SPEC-019-10 created this independent completion boundary.
- W019 excludes exact checklist wording and storage format from its own boundary.
- PRODUCT-WORK-SPEC-021 is the separate implementation owner.
- PRODUCT-TASK-SPEC-019-21 permits W020 decision, Investigation, and authoring before W019 closure.
- PRODUCT-TASK-SPEC-019-22 materialized the revised release graph and W020 internal coordination owner.
- PRODUCT-TASK-SPEC-020-01 materialized T02 through T06 and reserved PRODUCT-INV-SPEC-008.
- PRODUCT-TASK-SPEC-022-02 added the accepted `work_item_execution` authority.
- PRODUCT-TASK-SPEC-022-03 released T04 and preserved T04 as the sole checklist writer.
- PRODUCT-TASK-SPEC-020-04 completed the eleven-type checklist artifact set.
- PRODUCT-TASK-SPEC-020-05 returned final `PASS` after the factual compactness correction.
- PRODUCT-TASK-SPEC-020-06 synchronized W020 closure with every Completion Condition evaluated as `PASS`.
- PRODUCT-WORK-SPEC-021 may consume the accepted checklist artifact set.
- No validator implementation, DRMCP integration, stage, or commit work was performed by W020.
