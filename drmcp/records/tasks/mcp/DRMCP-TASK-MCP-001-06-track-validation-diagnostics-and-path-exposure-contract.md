# DRMCP-TASK-MCP-001-06: Track validation, diagnostics, and path-exposure contract correction

- **id**: DRMCP-TASK-MCP-001-06
- **status**: in_progress
- **date**: 2026-06-26
- **work_item**: DRMCP-WORK-MCP-001
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 0.5d coordination
- **depends_on**:
  - DRMCP-TASK-MCP-001-03
  - DRMCP-TASK-MCP-001-05
- **outputs**:
  - DRMCP-WORK-MCP-006

## Goal

Accept the corrected validation, diagnostic, and path-exposure contract gate.

## Work

- Track `DRMCP-WORK-MCP-006` as the exact child Work Item selected by T01.
- Delegate cross-namespace validation, current-to-legacy relation validation, diagnostic representation, and path-exposure contracts to that child Work Item.
- Confirm that the child Work Item separates semantic invalidity from DRMCP diagnostic representation.
- Track the child Work Item through contract review and `done`.
- Record the child Work Item ID and accepted evidence here.

This Task does not modify validation, diagnostic, or response contracts.
All detailed contract work belongs to the selected child Work Item.

## Done condition

- The selected child Work Item is `done`.
- Current cross-namespace and configured current-to-legacy relation behavior is explicit.
- Unsupported, unresolved, duplicate, disabled-fallback, and source-format diagnostics are machine-readable.
- Legacy archive records are excluded from normal current repository validation.
- Normal listing and retrieval responses hide physical paths.
- Path exposure is limited to explicit patch, diagnostic, or debug surfaces.
- The child review has no blocking or major findings.
- The exact child Work Item ID and evidence pointer are recorded here.

## Verification

- Review the final validation, diagnostic, and response-boundary contracts.
- Confirm that normal response schemas expose canonical identity rather than physical paths.
- Confirm that this Task contains no direct contract implementation evidence.

## Evidence

Selected child Work Item: `DRMCP-WORK-MCP-006`.

`DRMCP-TASK-MCP-006-01` opened on 2026-06-28 and moved `DRMCP-WORK-MCP-006` to `in_progress`.

`DRMCP-TASK-MCP-006-01` closed on 2026-06-28 after re-review returned `PASS`, `F-MIN-01` was closed, and scoped `git diff --check` passed with only a non-blocking LF-to-CRLF warning.

The accepted T01 baseline supplies the authority matrix, W003-W005 ownership inputs, contradiction inventory, candidate manifests, and T02-T05 split.

`DRMCP-TASK-MCP-006-02` closed on 2026-06-28 after defining the current repository and relation-validation execution contract.

T02 established repository-wide, app-scoped, and exact-current-ref validation; W003-retained subject selection; current cross-root and configured legacy relation lookup; request, execution, and diagnostic boundaries; and the normal validation wrapper.

The T02 limited re-review returned `PASS`, closed `F-MIN-01`, and left no blocking, major, or minor findings. Advisory `A-01` is delegated to T03.

T06 remains `in_progress` until the child Work Item reaches `done`, its final independent review has no blocking or major finding, and the gate evidence is accepted here.
