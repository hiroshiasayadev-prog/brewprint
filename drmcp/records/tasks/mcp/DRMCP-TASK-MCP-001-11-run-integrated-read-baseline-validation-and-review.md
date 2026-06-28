# DRMCP-TASK-MCP-001-11: Run integrated read-baseline validation and review

- **id**: DRMCP-TASK-MCP-001-11
- **status**: not_started
- **date**: 2026-06-26
- **work_item**: DRMCP-WORK-MCP-001
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 2d
- **depends_on**:
  - DRMCP-TASK-MCP-001-07
  - DRMCP-TASK-MCP-001-08
  - DRMCP-TASK-MCP-001-09
  - DRMCP-TASK-MCP-001-10
- **outputs**:
  - DRMCP-WORK-MCP-001

## Goal

Validate the integrated current read baseline and obtain an independent review verdict.

## Work

- Confirm that every lifecycle-tracking Task records a `done` child Work Item and accepted evidence.
- Run the full scoped DRMCP validation and automated test set.
- Verify current-only behavior without `legacy_roots`.
- Verify configured fallback and legacy leakage prevention.
- Compare implementation behavior with corrected contracts and fixtures.
- Request an independent contract and implementation review.
- Route every substantive correction to the owning child Work Item.
- Reopen the affected lifecycle gate until the child Work Item closes the finding.
- Record final validation and review evidence in this Task and the parent Work Item.

This Task may perform integrated verification only.
This Task must not absorb corrective contract or implementation work from a child Work Item.

## Done condition

- All required scoped validation and automated tests pass.
- Current-only and configured-fallback scenarios match `DRMCP-REQ-MCP-001`.
- No legacy record leaks into normal listing, current validation, or authoring targets.
- Physical paths remain absent from normal listing and retrieval responses.
- Every substantive finding is closed in its owning child Work Item.
- Independent review reports no blocking or major findings.
- Final validation commands, results, review verdict, and residual limitations are recorded.

## Verification

- Re-run the accepted validation and test commands after all routed findings close.
- Inspect the final review against `DRMCP-ADR-MCP-001` and `DRMCP-REQ-MCP-001`.
- Confirm that corrective diffs reside in child Work Item evidence rather than this Task.

## Evidence

Pending child Work Item completion.
