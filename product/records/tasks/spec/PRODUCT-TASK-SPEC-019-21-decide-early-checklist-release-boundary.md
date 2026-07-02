# PRODUCT-TASK-SPEC-019-21: Decide early checklist release boundary

- **id**: PRODUCT-TASK-SPEC-019-21
- **status**: done
- **date**: 2026-07-01
- **work_item**: PRODUCT-WORK-SPEC-019
- **task_type**: decision
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-019-20
- **outputs**:
  - PRODUCT-TASK-SPEC-019-21

## Goal

Decide which W020 responsibilities may proceed before W019 design closure.

## Work

- Preserve T09 as the completed original release-order decision.
- Decide whether W020 contract decisions, Investigation, and checklist authoring may proceed early.
- Decide the W020 integrated-review gate.
- Preserve the accepted W021 implementation release boundary.
- Record the current release order without modifying the graph.

This Task must not:

- rewrite T09;
- amend Work Items or Tasks other than this decision ledger;
- author checklist content;
- modify ADRs or Specifications;
- implement the validator;
- perform coordination, review, correction, synchronization, stage, or commit work.

### Decision ledger

| item | status | decision |
|---|---|---|
| D-001 | `decided` | W020 design and checklist authoring may proceed before W019 closure, while W020 integrated review remains gated by accepted W019 design review. |

### D-001 decision

- Early work: W020 checklist contract decisions may start immediately.
- Early work: W020 checklist impact Investigation may start after its contract decision.
- Early work: W020 checklist artifact authoring may start after its Investigation.
- Fixed authority: Early work uses accepted PRODUCT-ADR-SPEC-015 through PRODUCT-ADR-SPEC-017 and the current Task-type contract.
- Review gate: W020 integrated review waits for W019 T18 `PASS` or complete independent closure of every W019 review finding.
- Closure gate: W020 closes only after its own accepted integrated review route.
- Implementation gate: W021 still waits for W019 design closure and accepted W020 checklist review.
- Historical preservation: T09 remains unchanged as the earlier release-order checkpoint.
- Current authority: This successor decision replaces only the current W020 start timing from T09.
- ADR route: `not_required`; existing ADRs already govern independent boundaries and review order.

## Done condition

- The early W020 work boundary is explicit.
- The W020 integrated-review gate is explicit.
- The W021 release boundary remains unchanged.
- T09 remains historical Evidence.
- No graph or artifact content is changed by this Task.

## Verification

- Confirm D-001 is `decided`.
- Confirm W020 authoring may proceed before W019 closure.
- Confirm W020 integrated review cannot accept unreviewed W019 canonical state.
- Confirm W021 remains blocked by W019 closure and accepted W020 review.
- Confirm T09 remains unchanged.

## Evidence

- The user explicitly directed checklist decisions and checklist creation to proceed where possible.
- PRODUCT-ADR-SPEC-015 already fixes checklist composition and Task-local semantic evaluation semantics.
- PRODUCT-ADR-SPEC-016 fixes the standalone ownership boundary.
- PRODUCT-ADR-SPEC-017 fixes invocation and human exception semantics.
- W019 T16 and T17 remain canonical projection writers, so W020 review waits for accepted W019 review.
