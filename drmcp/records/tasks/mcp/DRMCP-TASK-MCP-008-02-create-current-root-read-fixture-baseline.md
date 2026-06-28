# DRMCP-TASK-MCP-008-02: Create current-root read fixture baseline

- **id**: DRMCP-TASK-MCP-008-02
- **status**: done
- **date**: 2026-06-28
- **work_item**: DRMCP-WORK-MCP-008
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 1d
- **depends_on**:
  - DRMCP-TASK-MCP-008-01
- **outputs**:
  - DRMCP-WORK-MCP-008
  - DRMCP-TASK-MCP-001-08
  - drmcp/src/internal/designrecords/testdata/read-baseline/manifest.json
  - drmcp/src/internal/designrecords/testdata/read-baseline/current/
  - drmcp/src/internal/designrecords/testdata/read-baseline/config/

## Goal

Create the shared current-root fixture baseline for T01 cases C01 through C14 and C17.

Keep the fixture package local to `drmcp/src/internal/designrecords` and exclude legacy, rejection, invalid-root, and leakage material.

## Work

- Confirm the bounded Task and package-local fixture inventory.
- Accept or replace the T01 candidate fixture root.
- Create separate PRODUCT and DRMCP current app roots.
- Create valid app-aware ADR, investigation, requirement, work-item, and task fixtures.
- Create path-derived current spec fixtures with H1-adjacent metadata.
- Create one valid Topics parent-child arrangement.
- Create fixture-local current-only and omitted-legacy configuration declarations.
- Create one authoritative manifest containing only T02-owned cases.
- Record exact current inputs, listing candidates, relation targets, validation subjects, and future runtime owners.
- Keep production implementation and existing Go tests unchanged.

## Done condition

- T01 cases C01 through C14 and C17 are materialized.
- PRODUCT and DRMCP current roots are physically separate.
- Every sequential fixture uses an app-aware canonical ID.
- Every current spec uses a path-derived `spec:` ref and H1-adjacent metadata.
- One exact cross-app relation uses a canonical public ID.
- Normal listing candidates contain current sequential records only.
- Current validation subjects contain no legacy archive source.
- The valid Topics row contains `title`, `kind`, `ref`, and `summary`.
- The manifest records the required case fields and no T03 or T04 case.
- No production implementation or existing Go test changes.
- Fixture-local structure is reviewed before this Task changes to `done`.

## Verification

- Parse `manifest.json` as JSON.
- Confirm case IDs equal C01 through C14 and C17.
- Resolve every declared T02 internal path from the accepted fixture root and confirm it exists.
- Confirm current roots are disjoint and remain below the package-local fixture root.
- Confirm sequential H1 IDs, metadata IDs where applicable, file-name prefixes, and manifest refs agree.
- Confirm spec IDs match their fixture-relative paths below each app `records/spec` root.
- Confirm the Topics child exists and its `parent` matches the declaring Index ref.
- Confirm no `legacy/`, invalid arrangement, rejection, leakage, duplicate, or unresolved fixture was created.
- Run scoped JSON, lifecycle, path, and whitespace checks externally when repository-local command execution is available.

## Evidence

### Pre-creation inventory

The exact directory `drmcp/records/tasks/mcp/` was listed once.
The only existing W008 child Task was `DRMCP-TASK-MCP-008-01`.
No `DRMCP-TASK-MCP-008-02` record existed.

The exact directory `drmcp/src/internal/designrecords/` was listed once.
No package-local `testdata` directory existed.
The candidate path `drmcp/src/internal/designrecords/testdata/read-baseline/` did not exist because its parent `testdata` directory was absent.

The T01 candidate path is accepted without change.
The path is package-local, does not create a repository top-level fixture directory, and does not collide with existing files.

### Accepted fixture boundary

T02 materializes only C01 through C14 and C17.
No legacy archive, rejection, invalid-root, duplicate, overlap, unresolved, empty-`legacy_roots`, or leakage fixture belongs to this Task.

The configuration JSON files are fixture-local arrangement declarations.
They preserve the accepted semantic fields without claiming to define the future runtime configuration serialization.

### Materialized current baseline

The accepted fixture root is `drmcp/src/internal/designrecords/testdata/read-baseline/`.

T02 created:

- one authoritative `manifest.json` containing only C01 through C14 and C17;
- separate PRODUCT and DRMCP current roots;
- one PRODUCT ADR fixture;
- one DRMCP investigation fixture;
- two DRMCP requirement fixtures;
- one DRMCP Work Item fixture;
- one DRMCP Task fixture;
- one PRODUCT Index spec and one child Overview spec;
- one current-only root arrangement declaration;
- one configuration declaration with the `legacy_roots` key omitted.

The manifest records stable case IDs, fixture classes, current root arrangements, fixture-root-relative paths, exact inputs, `legacy_roots_state`, expected classifications, intentional invalidity, structured runtime-owner lists, planned Task, and notes.
Every T02 case has `expected_classification: accepted` and `intentional_invalidity: null`.
No future T03 or T04 case is pre-registered.

The PRODUCT Index Topics row uses `title`, `kind`, `ref`, and `summary`.
The child Overview `parent` is `spec:product.fixture_baseline`, matching the declaring Index ref.
The DRMCP requirement uses exact cross-app source ref `PRODUCT-ADR-SPEC-901`.
The listing arrangement contains two DRMCP requirement candidates and keeps the PRODUCT spec exact-only.

### Filesystem static inspection

The package-local fixture tree was reread after creation.
It contains only `manifest.json`, `current/`, and `config/` at the suite root.
The current tree contains separate `product/records` and `drmcp/records` roots.
No `legacy/`, `arrangements/`, duplicate, invalid-root, overlap, rejection, or leakage directory exists.
A bounded filename search found no `V01-*` fixture file.

All eight current record and spec files plus both configuration declarations were reread through filesystem MCP.
The visible sequential IDs, file-name prefixes, workflow relations, spec metadata, Topics row, and child `parent` marker agree by static inspection.
The manifest was reread and contains exactly C01 through C14 and C17 with no future-case placeholder.
This inspection does not replace JSON parsing, PRODUCT validation, DRMCP parsing, or runtime assertions.

### Lifecycle synchronization

- This Task is `done` after corrected scoped checks and independent scoped re-review acceptance.
- `DRMCP-WORK-MCP-008` remains `in_progress` and now lists this Task.
- Hub `DRMCP-TASK-MCP-001-08` remains `in_progress` and points to this Task's materialized current baseline.
- `DRMCP-TASK-MCP-008-01` remains `done`.

### Changed-file boundary

Workflow files:

- new: `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-008-02-create-current-root-read-fixture-baseline.md`;
- modified: `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-008-current-and-legacy-read-fixture-baseline.md`;
- modified: `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-001-08-track-current-and-legacy-fixture-baseline.md`.

Fixture files are limited to `drmcp/src/internal/designrecords/testdata/read-baseline/`.
Production implementation and existing Go tests were not changed.

### External scoped verification

The user executed the recorded scoped fixture and lifecycle verification from the repository root.
The reported result was:

```text
fixture_shape=OK
lifecycle_shape=OK
```

The targeted status output matched the declared T02 boundary:

- modified hub `DRMCP-TASK-MCP-001-08`;
- modified `DRMCP-WORK-MCP-008`;
- untracked T02 Task;
- untracked `manifest.json`;
- untracked current fixture files;
- untracked current-only configuration declarations.

Tracked whitespace checks returned `tracked_exit=0` for W008 and hub T08.
Every untracked Task or fixture file returned `untracked_exit=1` through `git diff --no-index --check -- NUL <path>`.
No whitespace error and no exit code `2` or greater was reported.
LF-to-CRLF working-copy warnings were reported for the checked files and are non-blocking.

The verification establishes fixture JSON shape, declared path existence, bounded current-root structure, selected metadata and relation consistency, lifecycle shape, and whitespace readiness.
It does not establish production runtime behavior or existing Go test results.

The first independent review returned `NEEDS REVISION` with no blocking finding.
F-MAJ-01 identified the inactive `PRODUCT` / `MCP` domain assignment.
F-MAJ-02 identified repository-root-relative manifest paths instead of fixture-root-relative paths.
F-MIN-01 identified the missing structured W-SPEC-001 owner for C07.

The correction renamed and moved the PRODUCT ADR to `PRODUCT-ADR-SPEC-901` under `current/product/records/adr/spec/`, synchronized all references, converted internal manifest paths to fixture-root-relative values, and represented every `runtime_owner` as a list.
C07 now records both `DRMCP-WORK-MCP-009` and `DRMCP-WORK-SPEC-001`.

The earlier fixture and lifecycle results predate these corrected bytes.
A corrected scoped `git.inspect_worktree` run reported `result: pass` for the three workflow files and all eleven fixture files.
The workflow scope contains two modified tracked files and one untracked Task; the fixture scope contains eleven untracked files.
Whitespace checks passed for every scoped file, with LF-to-CRLF conversion warnings only.
`repository_wide_clean` was not asserted.
The user then ran the corrected verifier from the repository root and reported:

```text
fixture_shape=OK
lifecycle_shape=OK
```

This confirms the corrected manifest parses as JSON, all fixture-root-relative paths resolve, the active `PRODUCT` / `SPEC` ADR identity and synchronized relations are present, every runtime owner is structurally represented as a list, C07 contains both accepted owners, and the T02/W008/hub/T01 lifecycle remains correct.

The independent scoped re-review returned `PASS`.
F-MAJ-01, F-MAJ-02, and F-MIN-01 are `CLOSED`.
No blocking, major, or minor regression was reported.

T02 is closure-ready and now `done`.
W008 and hub T08 remain `in_progress` because T03 through T05 are not complete.

### Command execution state

The read-only `git.inspect_worktree` tool was executed only for the declared workflow and fixture paths.
No runtime operation, existing Go test, or repository-wide validation was executed by this assistant.
The corrected JSON fixture-shape and lifecycle commands were executed externally by the user and returned the results recorded above.
Repository-wide cleanliness remains unasserted.
