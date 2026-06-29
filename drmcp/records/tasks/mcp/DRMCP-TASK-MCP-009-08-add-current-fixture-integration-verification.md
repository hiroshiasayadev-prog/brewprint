# DRMCP-TASK-MCP-009-08: Add current fixture integration verification

- **id**: DRMCP-TASK-MCP-009-08
- **status**: not_started
- **date**: 2026-06-29
- **work_item**: DRMCP-WORK-MCP-009
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 2d
- **depends_on**:
  - DRMCP-TASK-MCP-009-05
  - DRMCP-TASK-MCP-009-06
  - DRMCP-TASK-MCP-009-07
- **outputs**:
  - drmcp/src/internal/designrecords/current_read_fixture_test.go
  - drmcp/src/internal/designrecords/authoring_test.go
  - drmcp/src/internal/designrecords/authoring_guidance_test.go

## Goal

Verify the complete current-only read path against accepted W008 fixtures without configured legacy roots.

Prove package integration, public current operations, path hiding, and authoring non-regression before independent review.
Migrate only protected authoring test setup and fixture identities to the final explicit current-root contract without changing authoring source or assertion semantics.

## Work

- Add one package-local integration test file that reads the accepted fixture manifest and current arrangements.
- Cover W009-owned current cases C01 through C16.
- Cover W009-owned rejection and isolation cases R02 through R08, R10, R14, and R20.
- Exercise configuration, parsing, active indexing, listing, exact retrieval, current resolution, current validation, diagnostics, and path hiding.
- Run without configured legacy roots.
- Treat the complete W008 fixture tree as read-only.
- Confirm T05, T06, and T07 are merged and their owned tests pass before starting fixture integration.
- Migrate `authoring_test.go` and `authoring_guidance_test.go` from auto-discovered or shape-invalid roots, non-app identities, and YAML-front-matter spec setup to the final current-only test shape.
- Preserve the tested authoring behavior, proposal assertions, diagnostics, and write semantics; change only setup, fixture paths/content, and canonical identities required by the final read contracts.
- Run complete affected-package tests after the three P3 branches merge and the test-only migration is complete.
- Verify authoring tests pass without authoring production changes.

C17 and R21 through R23 remain W-SPEC-owned.
All L cases and other legacy-owned R cases remain W010-owned.

### Execution slices

| slice | owner model | parallel group | dependency | exact file boundary or inventory method | allowed changes | prohibited changes | commands | expected evidence | escalation condition |
|---|---|---|---|---|---|---|---|---|---|
| S08A integration oracle and test migration | Sonnet | P4 | T05, T06, and T07 accepted and merged; their owned tests pass | Add `current_read_fixture_test.go`; modify only `authoring_test.go` and `authoring_guidance_test.go`; fixtures read-only. | New integration tests and test-only migration of config, paths, content, and canonical identities. | Production files, fixture bytes, authoring behavior changes, legacy behavior, W-SPEC validators, or weakened assertions. | `gofmt -w`; targeted authoring tests; targeted integration test; full affected-package tests. | Case matrix, authoring migration diff, command outputs, fixture read-only proof, and P3 owned-test PASS pointers. | Any required production or authoring semantic correction; reopen the owning T02-T07 Task or stop for reviewed boundary change. |
| S08B complete verification | Haiku | P4 | S08A complete | The three T08 test files plus final T02-T07 changed-file manifests; read-only. | No file changes. | Test repair, production correction, assertion weakening, or case reinterpretation. | `gofmt -d`; targeted authoring and integration tests; full package tests; optional race test when supported; scoped Git check. | Raw outputs, case count, authoring non-regression result, affected-package pass results, and exact changed paths. | Any failure, missing case, fixture modification, assertion weakening, or legacy execution; escalate to Sonnet. |

## Done condition

- One package-local integration test consumes accepted fixture paths rather than redefining fixture data inline.
- C01 through C16 have current-runtime coverage appropriate to their W009 ownership.
- R02 through R08, R10, R14, and R20 have rejection, conflict, invalid-root, unresolved, or path-hiding coverage.
- No L case, legacy fallback, legacy validation subject, or legacy active-index behavior executes.
- C17 and R21 through R23 are not claimed as W009 coverage.
- T05, T06, and T07 are integrated before this Task starts.
- T05, T06, and T07 owned tests pass before this Task starts.
- Normal list, retrieval, and resolver outputs contain no physical path.
- Current-only operation succeeds with legacy roots omitted.
- Fixture bytes remain unchanged.
- `authoring_test.go` and `authoring_guidance_test.go` use explicit current roots, app-aware canonical identities, and current spec fixtures without weakening authoring behavior assertions.
- Full `designrecords` tests pass, including authoring tests.
- Every package added by the T05 catalog boundary passes its full test command.
- Only `current_read_fixture_test.go`, `authoring_test.go`, and `authoring_guidance_test.go` change in T08.

## Verification

Run from repository root:

```powershell
gofmt -d drmcp/src/internal/designrecords/current_read_fixture_test.go
gofmt -d drmcp/src/internal/designrecords/authoring_test.go drmcp/src/internal/designrecords/authoring_guidance_test.go
go test ./drmcp/src/internal/designrecords -run '^(TestAuthoring.*|TestPropose.*|TestBodyCache.*|TestExactID.*|TestReplaceNamedSection.*|TestMultiOp.*|TestDiffMode.*|TestListAuthoringGuides|TestGetAuthoringGuidance.*|TestBuildIndexIgnoresAuthoringGuides)' -count=1
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

- the exact T05/T06/T07 owned-test evidence used as the P3 start gate;
- the exact W008 case-to-test matrix;
- fixture root and configuration used by each integration group;
- proof that `legacy_roots` is omitted;
- targeted, full package, and optional race outputs;
- fixture scoped status and whitespace result;
- scoped authoring test-migration diff and proof that authoring production files are unchanged;
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
    - path: drmcp/src/internal/designrecords/authoring_test.go
      tests: []
    - path: drmcp/src/internal/designrecords/authoring_guidance_test.go
      tests: []
  future_canonicalization:
    internal_design_ref: pending
    bpdsl_ref: pending
```

Populate `tests` with real Go test function names before Task closure.
Keep `implementation` empty because T08 owns no production implementation.
