# DRMCP-TASK-MCP-001-08: Track current and legacy fixture baseline

- **id**: DRMCP-TASK-MCP-001-08
- **status**: in_progress
- **date**: 2026-06-28
- **work_item**: DRMCP-WORK-MCP-001
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 0.5d coordination
- **depends_on**:
  - DRMCP-TASK-MCP-001-02
  - DRMCP-TASK-MCP-001-03
  - DRMCP-TASK-MCP-001-04
  - DRMCP-TASK-MCP-001-05
  - DRMCP-TASK-MCP-001-06
  - DRMCP-TASK-MCP-001-07
- **outputs**:
  - DRMCP-WORK-MCP-008

## Goal

Accept the current-format and configured legacy-fallback fixture baseline.

## Work

- Track `DRMCP-WORK-MCP-008` as the exact child Work Item selected by T01.
- Delegate fixture design, fixture creation, and fixture-local verification to that child Work Item.
- Require separate fixture roots or explicit separation for current and legacy records.
- Require accepted and rejected identity, configuration, relation, and source-format cases.
- Track the child Work Item through fixture review and `done`.
- Record the child Work Item ID and accepted evidence here.

This Task does not create fixtures or fixture tests.
All fixture authoring and local verification belong to the selected child Work Item.

## Done condition

- The selected child Work Item is `done`.
- Fixtures cover current app-aware sequential IDs and path-derived spec refs.
- Fixtures cover accepted V01 sequential families under configured fallback.
- Fixtures reject `V01-SPEC-*`, app-prefixless IDs, path inputs, and fuzzy repair.
- Fixtures cover current-only operation without legacy roots.
- Current and legacy fixtures cannot leak into each other's normal scope.
- The child review has no blocking or major findings.
- The exact child Work Item ID and evidence pointer are recorded here.

## Verification

- Review the fixture matrix and representative fixture files.
- Confirm every required outcome in `DRMCP-REQ-MCP-001` has a fixture case.
- Confirm that this Task contains no direct fixture implementation evidence.

## Evidence

Selected child Work Item: `DRMCP-WORK-MCP-008`.

The child fixture-baseline lifecycle began on 2026-06-28.
`DRMCP-TASK-MCP-008-01` now supplies the initial coverage matrix, bounded existing-test inventory, candidate fixture-root and manifest proposal, and future owner allocation.

No fixture file, production implementation, or existing test was changed as hub evidence.
T08 remains `in_progress` until the selected child Work Item reaches `done` and its accepted evidence is recorded here.
