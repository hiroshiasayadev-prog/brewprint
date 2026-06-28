# DRMCP-TASK-MCP-009-04: Build deterministic current active index

- **id**: DRMCP-TASK-MCP-009-04
- **status**: not_started
- **date**: 2026-06-28
- **work_item**: DRMCP-WORK-MCP-009
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 1.5d
- **depends_on**:
  - DRMCP-TASK-MCP-009-02
  - DRMCP-TASK-MCP-009-03
- **outputs**:
  - drmcp/src/internal/designrecords/index.go
  - drmcp/src/internal/designrecords/index_test.go
  - drmcp/src/internal/designrecords/tools.go
  - drmcp/src/internal/designrecords/resolver.go
  - drmcp/src/internal/designrecords/resolve_reference_test.go
  - drmcp/src/internal/designrecords/validation.go
  - drmcp/src/internal/designrecords/validation_test.go

## Goal

Integrate the accepted T02 and T03 branches into one compileable current-read foundation.

Build one deterministic active index from explicit current roots only.
Represent duplicate current canonical identities as conflict groups with no arbitrary winner.
Separate public read-operation file ownership before P3 starts.

## Work

- Merge the accepted T02 and T03 branch states.
- Discover current sequential records only in accepted kind and domain locations.
- Discover current spec Markdown recursively without following symlinks.
- Preserve stable app namespace and portable source provenance internally.
- Build addressable-source, active-record, and conflict-group views from T03 types.
- Sort all current index views deterministically.
- Fail complete index construction when a mandatory current root is invalid.
- Keep future legacy archive indexing structurally separate and unimplemented.
- Remove semantic-alias indexing and mixed current/V01 assumptions.
- Move or isolate the public `ResolveReference` entrypoint under `resolver.go`.
- Move or isolate the public `ValidateRecords` entrypoint under `validation.go`.
- Leave `tools.go` responsible for list/get public operations only.
- Freeze the shared current API used by T05 through T07.
- Run the T02/T03/T04 integrated package compile and full-package test.
- Record any expected transitional full-package failures by test name and owning T05-T07 Task.

T04 may adjust direct tests only when required by the active-index implementation or public operation ownership split.
T04 must not change resolver semantics, validation semantics, list/get behavior, or stale behavior only to force a full-package PASS.

### Execution slices

| slice | owner model | parallel group | dependency | exact file boundary or inventory method | allowed changes | prohibited changes | commands | expected evidence | escalation condition |
|---|---|---|---|---|---|---|---|---|---|
| S04A serial foundation integration | Sonnet | P2 | T02 and T03 targeted acceptance | `index.go`, `index_test.go`, `tools.go`, `resolver.go`, `resolve_reference_test.go`, `validation.go`, `validation_test.go`. Test edits are limited to compile or registration adjustment caused by the public operation split. | Merge T02/T03, replace current discovery and index assembly, move or isolate public operation entrypoints, and freeze shared APIs. | Config redesign, shared type redesign, parser redesign, resolver semantics, validation semantics, list/get semantics, authoring, fixtures, legacy index, stale-behavior compatibility hacks. | `gofmt -w`; targeted index tests; package compile; full package test. | Deterministic ordering, duplicate no-winner, public operation ownership map, frozen API summary, compile output, full-package output or expected T05-T07 transitional failures. | Any missing shared type, parser output, unresolved ownership, or API mismatch; return to T02/T03 rather than crossing boundaries. |
| S04B mechanical verification | Haiku | P2 | S04A complete | Final S04 boundary; read-only. | No changes. | Index redesign, operation split repair, resolver or validation semantics, or test repair. | `gofmt -d`; targeted and repeated index tests; package compile; full package test; scoped Git check. | Raw outputs, exact changed paths, non-overlap proof for T05-T07. | Any compile failure, unexpected full-package failure, nondeterministic repeat result, or unresolved ownership; escalate to Sonnet. |

## Done condition

- Only explicit current roots feed the active index.
- Current sequential and spec sources are discovered by accepted placement rules.
- Symlinks are not followed.
- Invalid mandatory current roots fail complete index construction.
- Valid empty roots produce an empty current contribution.
- Duplicate canonical identities create a deterministic conflict group and no winner.
- Stable ordering is independent of filesystem enumeration order.
- No legacy source is loaded or inserted into the active index.
- Public `ResolveReference` ownership is in `resolver.go` or a dedicated wrapper outside `tools.go`.
- Public `ValidateRecords` ownership is in `validation.go` or a dedicated wrapper outside `tools.go`.
- `tools.go` owns list/get public operations only after T04.
- T05, T06, and T07 can start without sharing a writable file.
- The merged T02/T03/T04 state compiles.
- The full-package test is run.
- Any remaining full-package failure is explicitly limited to named T05-T07 stale-operation tests.
- Unexpected failures, compile failures, unresolved ownership, or API mismatch block T04 closure.
- W008 cases C08, C12, and R07 are covered.
- Only the S04 boundary files change.

## Verification

Run from repository root:

```powershell
gofmt -d `
  drmcp/src/internal/designrecords/index.go `
  drmcp/src/internal/designrecords/index_test.go `
  drmcp/src/internal/designrecords/tools.go `
  drmcp/src/internal/designrecords/resolver.go `
  drmcp/src/internal/designrecords/resolve_reference_test.go `
  drmcp/src/internal/designrecords/validation.go `
  drmcp/src/internal/designrecords/validation_test.go
go test ./drmcp/src/internal/designrecords -run 'Test(BuildIndex|CurrentActiveIndex|DuplicateCurrent|IndexDetermin)' -count=1
go test ./drmcp/src/internal/designrecords -run 'Test(BuildIndex|CurrentActiveIndex|DuplicateCurrent|IndexDetermin)' -count=10
go test ./drmcp/src/internal/designrecords -run '^$' -count=1
go test ./drmcp/src/internal/designrecords -count=1
```

If the final full-package command fails only in stale operations owned by T05 through T07, record each failing test and owning Task.
Any compile failure, unexpected failure, ownership gap, or API mismatch prevents T04 closure.

## Evidence

Record:

- exact discovery roots and source kinds;
- deterministic ordering proof from repeated tests;
- duplicate conflict-group output and no-winner proof;
- invalid and empty root behavior;
- current-only index proof;
- public operation ownership split proof;
- frozen shared API boundary for T05 through T07;
- targeted, repeated, package compile, and full-package outputs;
- expected transitional failure list, or explicit statement that no transitional failure remains.

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
    - DRMCP-WORK-MCP-008
  fixture_cases:
    - C08
    - C12
    - R07
  implementation:
    - path: drmcp/src/internal/designrecords/index.go
      symbols: []
    - path: drmcp/src/internal/designrecords/tools.go
      symbols: []
    - path: drmcp/src/internal/designrecords/resolver.go
      symbols: []
    - path: drmcp/src/internal/designrecords/validation.go
      symbols: []
  verification:
    - path: drmcp/src/internal/designrecords/index_test.go
      tests: []
    - path: drmcp/src/internal/designrecords/resolve_reference_test.go
      tests: []
    - path: drmcp/src/internal/designrecords/validation_test.go
      tests: []
  future_canonicalization:
    internal_design_ref: pending
    bpdsl_ref: pending
```

Populate `symbols` and `tests` with real names before Task closure.
Remove any path that does not contain the final contract-significant implementation or verification.
