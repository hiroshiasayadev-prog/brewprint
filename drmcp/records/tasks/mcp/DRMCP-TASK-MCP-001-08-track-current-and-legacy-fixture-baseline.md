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
`DRMCP-TASK-MCP-008-01` supplies the accepted coverage matrix, bounded existing-test inventory, package-local fixture-root proposal, manifest schema, and future owner allocation.
T01 independent review returned `PASS`, and external scoped lifecycle and whitespace verification completed.

`DRMCP-TASK-MCP-008-02` materializes current accepted cases C01 through C14 and C17 under `drmcp/src/internal/designrecords/testdata/read-baseline/`.
The first T02 review identified three corrections. Corrected fixture-shape and lifecycle checks passed, and scoped re-review returned `PASS` with all three findings closed.
`DRMCP-TASK-MCP-008-02` is `done`; the selected child Work Item remains `in_progress` through T05.
`DRMCP-TASK-MCP-008-03` materialized configured legacy accepted cases L01 through L09.
Independent T03 review returned `PASS` with no blocking, major, or minor finding, and T03 is `done`.
Post-closure external verification returned `fixture_shape=OK` and `lifecycle_shape=OK`.
`DRMCP-TASK-MCP-008-04` materialized C15, C16, L10 through L13, and R01 through R24.
Independent T04 review returned `PASS` with no blocking, major, or minor finding, and T04 is `done`.
The review accepted the 22-file boundary, lifecycle synchronization, and scoped Git evidence; the ready-to-run Python verifier remained `NOT RUN` during review without inferred results.
After closure synchronization, external verification returned `fixture_shape=OK` and `lifecycle_shape=OK` against the final T04 state.
T04 structural evidence, review acceptance, and exclusions are recorded in `DRMCP-TASK-MCP-008-04`; fixture details are not duplicated here.
The child Work Item owns the fixture files, manifest, and structural evidence; this hub Task records only lifecycle and evidence pointers.
No production implementation or existing Go test changed as hub evidence.
T08 remains `in_progress` until the selected child Work Item reaches `done` and its accepted evidence is recorded here.
