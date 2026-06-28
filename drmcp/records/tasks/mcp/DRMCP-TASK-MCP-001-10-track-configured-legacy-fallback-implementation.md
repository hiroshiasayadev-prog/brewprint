# DRMCP-TASK-MCP-001-10: Track configured legacy-fallback implementation

- **id**: DRMCP-TASK-MCP-001-10
- **status**: not_started
- **date**: 2026-06-26
- **work_item**: DRMCP-WORK-MCP-001
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 0.5d coordination
- **depends_on**:
  - DRMCP-TASK-MCP-001-02
  - DRMCP-TASK-MCP-001-05
  - DRMCP-TASK-MCP-001-08
  - DRMCP-TASK-MCP-001-09
- **outputs**:
  - DRMCP-WORK-MCP-010

## Goal

Accept the configured read-only legacy archive fallback implementation gate.

## Work

- Track `DRMCP-WORK-MCP-010` as the exact child Work Item selected by T01.
- Delegate legacy-root loading, legacy index construction, exact legacy retrieval, fallback resolution, relation validation, and leakage tests to that child Work Item.
- Require implementation against the corrected resolver contract and accepted legacy fixtures.
- Require no automatic `v01/` discovery or current-index merging.
- Track the child Work Item through implementation review and `done`.
- Record the child Work Item ID and accepted evidence here.

This Task does not modify source code, fixtures, or tests.
All implementation and local verification belong to the selected child Work Item.

## Done condition

- The selected child Work Item is `done`.
- Legacy fallback activates only for configured `legacy_roots`.
- The legacy index remains separate and read-only.
- Exact retrieval and fallback accept only approved V01 sequential families.
- Legacy records do not leak into normal listing, current validation, or authoring targets.
- Rejected identity and configuration cases pass their tests.
- The child review has no blocking or major findings.
- The exact child Work Item ID and evidence pointer are recorded here.

## Verification

- Review the child implementation evidence, test results, and leakage checks.
- Confirm that disabled-fallback behavior works without legacy roots.
- Confirm that this Task contains no direct implementation evidence beyond the child evidence pointer.

## Evidence

Selected child Work Item: `DRMCP-WORK-MCP-010`.

T10 remains `not_started` until the child Work Item begins and later reaches `done`.
