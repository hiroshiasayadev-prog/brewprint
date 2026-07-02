# PRODUCT-TASK-SPEC-019-09: Decide checklist and implementation Work Item split

- **id**: PRODUCT-TASK-SPEC-019-09
- **status**: done
- **date**: 2026-07-01
- **work_item**: PRODUCT-WORK-SPEC-019
- **task_type**: decision
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-019-08
- **outputs**:
  - PRODUCT-TASK-SPEC-019-09

## Goal

Decide the downstream Work Item boundaries for checklist authoring and temporary standalone validator implementation.

## Work

- Preserve T07 as the completed prior reconciliation checkpoint.
- Decide whether checklist authoring belongs to the implementation Work Item.
- Fix the checklist-authoring Work Item boundary.
- Fix the implementation Work Item boundary.
- Fix coarse release order between W019 and the two downstream Work Items.
- State the decomposition input without creating Work Items.

This Task must not:

- rewrite T07;
- create or edit Work Items;
- author checklist content;
- implement the validator;
- perform ADR routing, canonical authoring, review, synchronization, stage, or commit work.

### Decision ledger

| item | status | decision |
|---|---|---|
| D-001 | `decided` | Checklist authoring and standalone validator implementation use separate downstream Work Items. |

### D-001 decision

- Checklist boundary: One downstream Work Item owns exact checklist wording, common criteria, Task-type-specific criteria, checklist storage format, checklist placement, authoring, and independent review.
- Implementation boundary: One separate downstream Work Item owns the temporary standalone validator implementation and consumes the accepted checklist artifacts.
- W019 boundary: W019 continues to own the validator product contract, Requirement reconciliation, canonical Specification projection, integrated review, and closure.
- Release order: Checklist authoring starts after W019 design closure.
- Release order: Standalone implementation starts after W019 design closure and accepted checklist authoring.
- Independence: Checklist artifacts have an independent authoring and review completion judgment from executable implementation.
- Reason: Checklist content is normative authoring input. Runtime delivery has a separate owner, verification boundary, and release decision.
- Decomposition target: Create one checklist-authoring Work Item and one standalone-validator implementation Work Item.
- ADR route: Not selected by this Task.

## Done condition

- The checklist-authoring boundary is explicit.
- The implementation boundary is explicit.
- The two Work Items have distinguishable completion judgments.
- W019 continuation and coarse release order are explicit.
- No Work Item is authored by this Task.

## Verification

- Confirm exactly one decision item exists and is `decided`.
- Confirm checklist authoring is not classified as implementation.
- Confirm implementation consumes accepted checklist artifacts.
- Confirm T07 remains unchanged.
- Confirm no Work Item, checklist, implementation, ADR, Specification, review, synchronization, stage, or commit work occurred.

## Evidence

- The user explicitly classified checklist creation as authoring.
- The user required checklist authoring and implementation to use separate Work Items.
- T07 remains completed historical Evidence.
- D-001 provides fixed input to T10.
