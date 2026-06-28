# DRMCP-TASK-MCP-009-08: Add current fixture integration verification

- **id**: DRMCP-TASK-MCP-009-08
- **status**: not_started
- **date**: 2026-06-28
- **work_item**: DRMCP-WORK-MCP-009
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 1.5d
- **depends_on**:
  - DRMCP-TASK-MCP-009-05
  - DRMCP-TASK-MCP-009-06
  - DRMCP-TASK-MCP-009-07
- **outputs**:
  - drmcp/src/internal/designrecords/current_read_fixture_test.go

## Goal

Verify the complete current-only read path against accepted W008 fixtures without configured legacy roots.

Prove package integration, public current operations, path hiding, and authoring non-regression before independent review.

## Work

- Add one package-local integration test file that reads the accepted fixture manifest and current arrangements.
- Cover W009-owned current cases C01 through C16.
- Cover W009-owned rejection and isolation cases R02 through R08, R10, R14, and R20.
- Exercise configuration, parsing, active indexing, listing, exact retrieval, current resolution, current validation, diagnostics, and path hiding.
- Run without configured legacy roots.
- Treat the complete W008 fixture tree as read-only.
- Confirm T05, T06, and T07 are merged and the P3 integration full-package gate passes before starting fixture integration.
- Run complete affected-package tests after the three P3 branches merge.
- Verify authoring tests still pass without authoring source changes.

C17 and R21 through R23 remain W-SPEC-owned.
All L cases and other legacy-owned R cases remain W010-owned.

### Execution slices

| slice | owner model | parallel group | dependency | exact file boundary or inventory method | allowed changes | prohibited changes | commands | expected evidence | escalation condition |
|---|---|---|---|---|---|---|---|---|---|
| S08A integration oracle and test | Sonnet | P4 | T05, T06, and T07 accepted, merged, and passing the P3 integration full-package gate | Add `current_read_fixture_test.go` only. Read `testdata/read-baseline/**` without modification. | New integration tests and local test-only helpers in the one new file. | Production files, existing tests, fixture bytes, legacy behavior, W-SPEC validators, authoring files. | `gofmt -w`; targeted integration test; full affected-package tests. | Exact case-to-assertion matrix, command outputs, fixture read-only proof, P3 integration PASS pointer. | Any case requires production correction; reopen the owning T02-T07 Task instead of editing production here. |
| S08B complete verification | Haiku | P4 | S08A complete | New test file plus final T02-T07 changed-file manifests; read-only. | No file changes. | Test repair, production correction, or case reinterpretation. | `gofmt -d`; targeted integration test; full package tests; optional race test when supported; scoped Git check. | Raw outputs, case count, affected-package pass results, exact changed paths. | Any failure, missing case, fixture modification, or legacy execution; escalate to Sonnet. |

## Done condition

- One package-local integration test consumes accepted fixture paths rather than redefining fixture data inline.
- C01 through C16 have current-runtime coverage appropriate to their W009 ownership.
- R02 through R08, R10, R14, and R20 have rejection, conflict, invalid-root, unresolved, or path-hiding coverage.
- No L case, legacy fallback, legacy validation subject, or legacy active-index behavior executes.
- C17 and R21 through R23 are not claimed as W009 coverage.
- T05, T06, and T07 are integrated before this Task starts.
- The P3 integration full-package gate passes before this Task starts.
- Normal list, retrieval, and resolver outputs contain no physical path.
- Current-only operation succeeds with legacy roots omitted.
- Fixture bytes remain unchanged.
- Full `designrecords` tests pass, including authoring tests.
- Every package added by the T05 catalog boundary passes its full test command.
- Only the new integration test file changes in T08.

## Verification

Run from repository root:

```powershell
gofmt -d drmcp/src/internal/designrecords/current_read_fixture_test.go
go test ./drmcp/src/internal/designrecords -run 'TestCurrentReadFixtureBaseline' -count=1
go test ./drmcp/src/internal/designrecords -count=1
```

Also run the full package command for every MCP adapter package recorded by T05.
When the environment supports the race detector, additionally run:

```powershell
go test -race ./drmcp/src/internal/designrecords -run 'TestCurrentReadFixtureBaseline' -count=1
```

A missing race-detector prerequisite is an accurately recorded limitation, not a fabricated pass.

## Evidence

Record:

- the exact T05/T06/T07 integration evidence used as the P3 start gate;
- the exact W008 case-to-test matrix;
- fixture root and configuration used by each integration group;
- proof that `legacy_roots` is omitted;
- targeted, full package, and optional race outputs;
- fixture scoped status and whitespace result;
- authoring non-regression result;
- any reopened upstream Task and its accepted correction evidence.

### Provisional implementation mapping

```yaml
implementation_mapping:
  status: provisional
  contract_refs:
    - DRMCP-REQ-MCP-001
    - DRMCP-ADR-MCP-001
    - DRMCP-WORK-MCP-003
    - DRMCP-WORK-MCP-004
    - DRMCP-WORK-MCP-005
    - DRMCP-WORK-MCP-006
    - DRMCP-WORK-MCP-007
    - DRMCP-WORK-MCP-008
  fixture_cases:
    - C01
    - C02
    - C03
    - C04
    - C05
    - C06
    - C07
    - C08
    - C09
    - C10
    - C11
    - C12
    - C13
    - C14
    - C15
    - C16
    - R02
    - R03
    - R04
    - R05
    - R06
    - R07
    - R08
    - R10
    - R14
    - R20
  implementation: []
  verification:
    - path: drmcp/src/internal/designrecords/current_read_fixture_test.go
      tests: []
  future_canonicalization:
    internal_design_ref: pending
    bpdsl_ref: pending
```

Populate `tests` with real Go test function names before Task closure.
Keep `implementation` empty because T08 owns no production implementation.
