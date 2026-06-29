# DRMCP-TASK-MCP-009-04: Build deterministic current active index

- **id**: DRMCP-TASK-MCP-009-04
- **status**: done
- **date**: 2026-06-29
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
  - drmcp/src/internal/designrecords/validation.go

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
- Leave `tools.go` responsible for list/get operations plus T05-owned retired operations after resolver and validation entrypoints move out.
- Freeze the shared current API used by T05 through T07.
- Run the T02/T03/T04 integrated package compile and full-package test.
- Record expected transitional full-package failures by test name and owning T05-T07 Task or T08 authoring-test migration boundary.

T04 may adjust direct tests only when required by the active-index implementation or public operation ownership split.
Authoring tests and authoring source remain outside T04. T08 owns later test-only migration to the final current fixture and config shape.
T04 must not change resolver semantics, validation semantics, list/get behavior, or stale behavior only to force a full-package PASS.

### Execution slices

| slice | owner model | parallel group | dependency | exact file boundary or inventory method | allowed changes | prohibited changes | commands | expected evidence | escalation condition |
|---|---|---|---|---|---|---|---|---|---|
| S04A serial foundation integration | Sonnet | P2 | T02 and T03 targeted acceptance | `index.go`, `index_test.go`, `tools.go`, `resolver.go`, and `validation.go` only. | Merge T02/T03, replace current discovery and index assembly, move or isolate public operation entrypoints, and freeze shared APIs. | Config redesign, shared type redesign, parser redesign, resolver semantics, validation semantics, list/get semantics, authoring, fixtures, legacy index, stale-behavior compatibility hacks. | `gofmt -w`; targeted index tests; package compile; full package test. | Deterministic ordering, duplicate no-winner, public operation ownership map, frozen API summary, compile output, and named T05-T07 or T08 transitional-failure handoff. | Any missing shared type, parser output, ownership gap outside the T08 assignment, or API mismatch; return to T02/T03 rather than crossing boundaries. |
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
- `tools.go` no longer owns public resolver or validation entrypoints after T04.
- T05 receives exclusive writable ownership of `tools.go`, including retirement of `get_record`, `suggest_next_record`, and obsolete range behavior.
- T05, T06, and T07 can start without sharing a writable file.
- The merged T02/T03/T04 state compiles.
- The full-package test is run.
- Any remaining full-package failure is explicitly limited to named T05-T07 stale-operation tests or T08-owned authoring test setup that still assumes superseded current-root, identity, or spec-format behavior.
- Unexpected failures, compile failures, unresolved ownership outside that T08 assignment, or API mismatch block T04 closure.
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
  drmcp/src/internal/designrecords/validation.go
go test ./drmcp/src/internal/designrecords -run 'Test(BuildIndex|CurrentActiveIndex|DuplicateCurrent|IndexDetermin)' -count=1
go test ./drmcp/src/internal/designrecords -run 'Test(BuildIndex|CurrentActiveIndex|DuplicateCurrent|IndexDetermin)' -count=10
go test ./drmcp/src/internal/designrecords -run '^$' -count=1
go test ./drmcp/src/internal/designrecords -count=1
```

If the final full-package command fails only in stale operations owned by T05 through T07 or in the protected authoring test setup assigned to T08, record each failing test group and owner.
Any compile failure, unexpected failure, ownership gap outside that assignment, or API mismatch prevents T04 closure.

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
- expected transitional failure list, including the T08 authoring-test handoff, or explicit statement that no transitional failure remains.

### Implementation evidence

- Production commit: `3b842c0 feat(drmcp): build deterministic current active index`.
- T03 shared-model correction was committed separately before the T04 production commit; it adds `Index.ConflictGroups` using the already accepted `CurrentConflict` type.
- Final T04 production boundary: `index.go`, `index_test.go`, `tools.go`, `resolver.go`, and `validation.go`.
- `BuildIndex` reads explicit configured current roots only, validates every mandatory root, discovers accepted current placements, skips file and directory symlinks, excludes semantic aliases and legacy sources, creates no-winner conflict groups, and sorts every exposed internal view deterministically.
- Public `ResolveReference` ownership is isolated in `resolver.go`; public `ValidateRecords` ownership is isolated in `validation.go`; `tools.go` retains list/get entrypoints and T05-owned retired operations until T05 removes them.
- Exact T04 index tests passed with `-count=1` and `-count=10`.
- `TestCurrentActiveIndexMultiRootOrderIndependent` passed with `-count=100`.
- Package compile passed with the package compile command selecting no tests.
- `gofmt -w` completed and scoped `git diff --check` passed. LF-to-CRLF messages were line-ending conversion warnings, not whitespace findings.
- Scoped post-commit inspection of the T04 source boundary and this Task reported no remaining source changes; repository-wide cleanliness was not inferred.

### Independent implementation review disposition

The first independent implementation review returned `NEEDS REVISION` with two major findings.

- `F-MAJ-01` was accepted and returned to T03 because root `spec/index.md` did not derive `spec:<app_namespace>` and metadata `id` could become an identity fallback. The T03 corrective patch is applied and verified.
- `F-MAJ-02` correctly identified a contradiction in this Task record but proposed the wrong implementation correction. `SuggestNextRecord` remains in `tools.go` for T05 to retire under the accepted graph. T04 does not move temporary retired code into a new file. The Work, Done condition, and Evidence wording now reflect the T05 ownership boundary.

Corrective verification passed. The full-package command still fails only in the previously assigned T05, T06, T07, and T08 stale-test groups, with no new failure group introduced.

Scoped re-review completed with `PASS`. Both previous major findings are closed, and T04 is accepted for closure.

### Closure acceptance

- Re-review verdict: `PASS`.
- `F-MAJ-01`: `CLOSED`.
- `F-MAJ-02`: `CLOSED`.
- Blocking findings: none.
- Major findings: none.
- Minor findings: none.
- T03 corrective closure: accepted.
- T04 closure readiness: ready and synchronized to `done`.
- T05, T06, and T07 parallel-start readiness: ready after this closure synchronization.
- Repository-local formatter and tests were not independently rerun by the reviewer; provided execution evidence was accepted as consistent with the reviewed implementation and Task Evidence.
- Repository-wide cleanliness was not checked or inferred.

### Transitional full-package failure ownership

The full package test was run. The package compiled, and remaining failures were limited to these assigned migration groups:

- T05: stale list, single-get, retired suggestion, and repository-bootstrap tests in `list_records_test.go`, `get_record_test.go`, and `suggest_next_record_test.go`.
- T06: stale semantic-alias, legacy-bootstrap, and old request-shape tests in `resolve_reference_test.go`.
- T07: stale range, old diagnostic, semantic-alias, legacy-bootstrap, and old fixture-shape tests in `validation_test.go`.
- T08: protected authoring tests whose setup still calls `NewConfig` with an empty or shape-invalid records root, uses non-app canonical identities, or supplies YAML-front-matter specs. The affected boundary is `authoring_test.go` and `authoring_guidance_test.go`; authoring production source remains unchanged.

No failure required changing T04 index semantics, resolver semantics, validation semantics, list/get behavior, authoring behavior, or legacy behavior.

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
      symbols:
        - BuildIndex
        - validateCurrentRecordsRoot
        - separateConflictGroups
        - discoverCurrentADRRecords
        - discoverCurrentSpecRecords
        - discoverCurrentInvestigationRecords
        - discoverCurrentRequirementRecords
        - discoverCurrentWorkItemRecords
        - discoverCurrentTaskRecords
        - sortIndexDeterministically
    - path: drmcp/src/internal/designrecords/tools.go
      symbols:
        - ListRecords
        - GetRecords
    - path: drmcp/src/internal/designrecords/resolver.go
      symbols:
        - ResolveReference
    - path: drmcp/src/internal/designrecords/validation.go
      symbols:
        - ValidateRecords
  verification:
    - path: drmcp/src/internal/designrecords/index_test.go
      tests:
        - TestBuildIndexEmptyState
        - TestBuildIndexRejectsInvalidMandatoryRoot
        - TestCurrentActiveIndexMultiRootOrderIndependent
        - TestDuplicateCurrentCanonicalIdentityHasNoWinner
        - TestIndexDeterministicIssueOrdering
        - TestBuildIndexDoesNotIndexCurrentSemanticAliases
        - TestBuildIndexSkipsFileSymlinksForEveryCurrentKind
        - TestBuildIndexSkipsDirectorySymlinksForEveryCurrentKind
        - TestBuildIndexRejectsSymlinkedMandatoryRoot
  future_canonicalization:
    internal_design_ref: pending
    bpdsl_ref: pending
```

The mapping lists the final contract-significant symbols and verification tests. Paths without final T04-owned implementation or verification were removed.
