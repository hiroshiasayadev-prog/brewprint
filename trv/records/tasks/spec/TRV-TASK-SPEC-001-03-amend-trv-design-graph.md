# TRV-TASK-SPEC-001-03: Remove unnecessary TRV impact-Investigation route

- **id**: TRV-TASK-SPEC-001-03
- **status**: done
- **date**: 2026-07-02
- **work_item**: TRV-WORK-SPEC-001
- **task_type**: coordination
- **estimate**: 0.25d
- **depends_on**:
  - TRV-TASK-SPEC-001-02
- **outputs**:
  - TRV-WORK-SPEC-001
  - TRV-TASK-SPEC-001-03
  - TRV-TASK-SPEC-001-04

## Goal

Remove the unnecessary standalone impact-Investigation route from the TRV new-application design graph.

## Work

- Treat TRV as a new application with no existing production implementation to impact-analyze.
- Cancel the reserved `TRV-INV-SPEC-001` route without creating an Investigation record.
- Amend TRV-WORK-SPEC-001 so repository alignment is handled inside downstream graph coordination.
- Amend T04 to consume terminal T02 decisions directly.
- Keep only a bounded alignment check for namespace authority, PRODUCT semantic references, checklist assets, existing ADR reuse, shared writers, and review order.
- Preserve actual conflict routing when the bounded check discovers a concrete contradiction requiring a decision.

This Task must not:

- reopen or revise terminal T02 decisions;
- author an Investigation, Requirement, ADR, or Specification;
- perform implementation planning or production implementation;
- create implementation Tasks or an executor prompt;
- issue a review verdict;
- perform stage or commit work.

## Done condition

- No TRV impact Investigation is required or created.
- TRV-INV-SPEC-001 is no longer reserved by the active graph.
- T04 remains the sole downstream graph-coordination owner and consumes T02 directly through this graph amendment.
- Repository alignment is bounded to information needed for authoring ownership, ADR routing, shared-writer order, and review order.
- Production implementation and current DRMCP integration remain outside the graph.

## Verification

- Confirm T03 uses `coordination`, not `investigation`.
- Confirm no TRV Investigation record exists.
- Confirm Work Item and T04 no longer require concluded Investigation Evidence.
- Confirm no design decision or canonical artifact body was changed.
- Confirm no implementation Task or executor prompt was created.

## Evidence

- T02 completed the app-local decision ledger.
- The user rejected the mandatory impact-Investigation premise because TRV is a new application.
- The previous T03 contract was an over-application of an existing-artifact design-convergence pattern.
- The active route now uses a lightweight repository-alignment check inside T04.
- Result: `PASS`.
