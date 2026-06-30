# DRMCP-TASK-MCP-001-11: Run integrated read-baseline validation and review

- **id**: DRMCP-TASK-MCP-001-11
- **status**: not_started
- **date**: 2026-06-30
- **work_item**: DRMCP-WORK-MCP-001
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 2d
- **depends_on**:
  - DRMCP-TASK-MCP-001-07
  - DRMCP-TASK-MCP-001-08
  - DRMCP-TASK-MCP-001-09
  - DRMCP-TASK-MCP-001-10
  - DRMCP-TASK-MCP-001-16
  - DRMCP-TASK-MCP-001-17
- **outputs**:
  - DRMCP-WORK-MCP-001

## Goal

Run the sole integrated verification gate for the complete rebuilt read baseline.

Obtain an independent integrated review verdict.

## Work

- Confirm T09 records reviewed W012 completion.
- Confirm T10 records reviewed W010 completion.
- Confirm T16 records reviewed W-SPEC-001 completion.
- Confirm T17 records reviewed W-SPEC-002 completion.
- Verify the W012 current runtime.
- Verify W-SPEC-001 per-file detectors.
- Verify W-SPEC-002 Topics graph validation.
- Verify W010 configured legacy fallback.
- Verify current-only, configured-fallback, and non-leakage behavior.
- Verify accepted diagnostic and path-exposure contracts.
- Run full affected-package tests.
- Request an independent integrated review.
- Route every failure to its owning Work Item.
- Record final verification and review evidence here and in W001.

This Task is the sole integrated verification owner.
This Task must not repair implementation failures.

## Done condition

- W012, W010, W-SPEC-001, and W-SPEC-002 are reviewed and `done`.
- Full affected-package tests pass.
- Current-only behavior matches the accepted contracts.
- Configured-fallback behavior matches the accepted contracts.
- No legacy record leaks into normal listing, current validation, or authoring targets.
- Accepted diagnostic and path-exposure contracts pass.
- Every failure is closed by its owning Work Item.
- Independent integrated review reports no blocking or major findings.
- Final commands, results, review verdict, and residual limitations are recorded.

## Verification

- Run the accepted full affected-package test commands after all owning Work Items close routed failures.
- Confirm W012, W010, W-SPEC-001, and W-SPEC-002 evidence is complete.
- Confirm current-only, configured-fallback, non-leakage, diagnostic, and path-exposure behavior.
- Confirm T11 contains no corrective implementation diff.
- Confirm each failure is routed to its owning Work Item.

## Evidence

Pending T09, T10, T16, and T17 completion.

T11 remains the sole integrated read-baseline verification owner.
