# DRMCP-TASK-MCP-001-09: Track current read implementation

- **id**: DRMCP-TASK-MCP-001-09
- **status**: in_progress
- **date**: 2026-06-28
- **work_item**: DRMCP-WORK-MCP-001
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 0.5d coordination
- **depends_on**:
  - DRMCP-TASK-MCP-001-03
  - DRMCP-TASK-MCP-001-04
  - DRMCP-TASK-MCP-001-05
  - DRMCP-TASK-MCP-001-06
  - DRMCP-TASK-MCP-001-08
- **outputs**:
  - DRMCP-WORK-MCP-009

## Goal

Accept the corrected current-format read implementation gate.

## Work

- Track `DRMCP-WORK-MCP-009` as the exact child Work Item selected by T01.
- Delegate active index, current spec handling, query, exact retrieval, resolver, validation, diagnostics, and current-only tests to that child Work Item.
- Require implementation against corrected contracts and accepted fixtures.
- Reject implementation work that preserves legacy behavior only because existing tests pass.
- Track the child Work Item through implementation review and `done`.
- Record the child Work Item ID and accepted evidence here.

This Task does not modify source code, fixtures, or tests.
All implementation and local verification belong to the selected child Work Item.

## Done condition

- The selected child Work Item is `done`.
- Current active roots build one deterministic active index.
- Current specs use the accepted metadata and identity format.
- Query, retrieval, resolution, validation, diagnostics, and path hiding match corrected contracts.
- Current-only tests pass without configured legacy roots.
- The child review has no blocking or major findings.
- The exact child Work Item ID and evidence pointer are recorded here.

## Verification

- Review the child implementation evidence, test results, and contract traceability.
- Confirm that no legacy archive record appears in normal current operations.
- Confirm that this Task contains no direct implementation evidence beyond the child evidence pointer.

## Evidence

Selected child Work Item: `DRMCP-WORK-MCP-009`.

Child Work Item planning began on 2026-06-28.

- `DRMCP-WORK-MCP-009` is `in_progress`;
- `DRMCP-TASK-MCP-009-01` owns the bounded code-to-contract inventory and Claude Code execution graph;
- T02 through T09 are the formal child implementation, verification, and closure Tasks;
- production implementation has not started;
- this hub Task remains `in_progress` until W009 reaches reviewed `done`.
