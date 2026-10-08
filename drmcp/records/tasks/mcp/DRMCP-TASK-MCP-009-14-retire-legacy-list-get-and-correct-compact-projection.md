# DRMCP-TASK-MCP-009-14: Retire legacy list get and correct compact projection

- **id**: DRMCP-TASK-MCP-009-14
- **status**: blocked
- **date**: 2026-06-29
- **work_item**: DRMCP-WORK-MCP-009
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 1d
- **depends_on**:
  - DRMCP-TASK-MCP-009-12
- **outputs**:
  - drmcp/src/internal/designrecords/tools.go
  - drmcp/src/internal/designrecords/list_records_test.go
  - drmcp/src/internal/designrecords/get_records_test.go

## Goal

Remove obsolete package list and get behavior and legacy tests.

Correct the compact current listing projection without changing protected shared types or compatibility helpers.

## Work

### Read

- This Task.
- `drmcp/src/internal/designrecords/tools.go`.
- `drmcp/src/internal/designrecords/list_records_test.go`.
- `drmcp/src/internal/designrecords/get_records_test.go`.

### Change

Only change:

- `drmcp/src/internal/designrecords/tools.go`.
- `drmcp/src/internal/designrecords/list_records_test.go`.
- `drmcp/src/internal/designrecords/get_records_test.go`.

### Implement

Remove these obsolete production symbols from `tools.go`:

- `type listRecordsScope`.
- `func ListRecords`.
- `func GetRecords`.
- `func newListRecordsScope`.
- `func (s listRecordsScope) selectRecord`.

In `currentListedRecord`, remove the `Kind` assignment.
The accepted compact shape is `ref`, `title`, `status`, and `date`.
Do not modify `CurrentListedRecord` in protected `types.go`; its existing `omitempty` tag makes the zero-value field absent from JSON.

Remove these legacy tests from `list_records_test.go`:

- `TestLegacyListingBasicFiltersAndResponseShape`.
- `TestLegacyListingStatusAndIDFilters`.
- `TestLegacyListingIDRangeFilter`.
- `TestLegacyListingWorkflowIDRangeFilter`.
- `TestLegacyListingRequestErrors`.
- `TestLegacyListingSortOrderAndLimit`.
- `TestLegacyListingRepositoryBootstrapQueries`.

Remove these legacy-only helpers from `list_records_test.go`:

- `buildListRecordsTestIndex`.
- `assertListRecordsErrorCode`.
- `listedRecordIDs`.
- `findListedRecord`.
- `latestDecisionRecordID`.

Retain:

- `TestListRecordsCurrentCompactFiltersDefaultsAndNoPath`.
- `TestListRecordsCurrentStatusOrderLimitAndHasMore`.
- `TestListRecordsCurrentRejectsObsoleteAndInvalidInputs`.
- `intPtr`.
- `currentReadTestIndex`.

In `TestListRecordsCurrentCompactFiltersDefaultsAndNoPath`, change the exact forbidden-key assertion to:

```go
for _, forbidden := range []string{"id", "kind", "path", "decision", "requirement", "headings", "body"} {
```

Do not add a new test for this clause.

Remove these legacy tests from `get_records_test.go`:

- `TestGetRecordsMixedKindsPartialResultsAndDuplicates`.
- `TestGetRecordsAllMissingIsNormalExactLookupResponse`.
- `TestGetRecordsRequestErrors`.

Retain:

- `TestGetRecordsCurrentExactBatchSuccessWarningsAndNoPath`.
- `TestGetRecordsCurrentRequestErrors`.
- `TestNormalReadGetRecordsPathIsNotSerialized`.
- `hasOperationWarning`.
- `jsonContainsKey`.
- `jsonValueContainsKey`.

### Test

Run from repository root:

```powershell
gofmt -w drmcp/src/internal/designrecords/tools.go drmcp/src/internal/designrecords/list_records_test.go drmcp/src/internal/designrecords/get_records_test.go

go test ./drmcp/src/internal/designrecords -run '^TestListRecordsCurrent' -count=1

go test ./drmcp/src/internal/designrecords -run '^(TestGetRecordsCurrent|TestNormalReadGetRecords)' -count=1
```

Run this exact scoped absence check:

```powershell
$paths = @(
  'drmcp/src/internal/designrecords/tools.go',
  'drmcp/src/internal/designrecords/list_records_test.go',
  'drmcp/src/internal/designrecords/get_records_test.go'
)

$removedNames = @(
  'listRecordsScope',
  'ListRecords',
  'GetRecords',
  'newListRecordsScope',
  'TestLegacyListingBasicFiltersAndResponseShape',
  'TestLegacyListingStatusAndIDFilters',
  'TestLegacyListingIDRangeFilter',
  'TestLegacyListingWorkflowIDRangeFilter',
  'TestLegacyListingRequestErrors',
  'TestLegacyListingSortOrderAndLimit',
  'TestLegacyListingRepositoryBootstrapQueries',
  'buildListRecordsTestIndex',
  'assertListRecordsErrorCode',
  'listedRecordIDs',
  'findListedRecord',
  'latestDecisionRecordID',
  'TestGetRecordsMixedKindsPartialResultsAndDuplicates',
  'TestGetRecordsAllMissingIsNormalExactLookupResponse',
  'TestGetRecordsRequestErrors'
)

$pattern = (($removedNames | ForEach-Object { [regex]::Escape($_) }) -join '|')
$matches = git grep -n -w -E -e $pattern -- $paths

if ($LASTEXITCODE -eq 0) {
  $matches
  throw 'Removed legacy symbols, tests, or helpers are still present.'
}

if ($LASTEXITCODE -ne 1) {
  throw "git grep failed with exit code $LASTEXITCODE"
}

$removedReceiverText = 'func (s listRecordsScope) selectRecord'
$receiverMatches = git grep -n -F -e $removedReceiverText -- $paths

if ($LASTEXITCODE -eq 0) {
  $receiverMatches
  throw 'Removed listRecordsScope selectRecord method is still present.'
}

if ($LASTEXITCODE -ne 1) {
  throw "git grep failed with exit code $LASTEXITCODE"
}
```

Exit semantics:

- `0`: forbidden match found; `FAIL`.
- `1`: no forbidden match; `PASS`.
- `2` or greater: command failure; `BLOCKED`.

Run this exact retained-test and compact-assertion presence check:

```powershell
$requiredPatterns = @(
  'func TestListRecordsCurrentCompactFiltersDefaultsAndNoPath',
  'func TestListRecordsCurrentStatusOrderLimitAndHasMore',
  'func TestListRecordsCurrentRejectsObsoleteAndInvalidInputs',
  '"id", "kind", "path", "decision", "requirement", "headings", "body"',
  'func TestGetRecordsCurrentExactBatchSuccessWarningsAndNoPath',
  'func TestGetRecordsCurrentRequestErrors',
  'func TestNormalReadGetRecordsPathIsNotSerialized'
)

foreach ($requiredPattern in $requiredPatterns) {
  git grep -n -F -e $requiredPattern -- $paths
  if ($LASTEXITCODE -eq 1) {
    throw "Required retained test or assertion is absent: $requiredPattern"
  }
  if ($LASTEXITCODE -ne 0) {
    throw "git grep failed with exit code $LASTEXITCODE for: $requiredPattern"
  }
}
```

T15 owns correction-scoped aggregate verification. T08 owns the first accepted full `designrecords` package PASS.

The three scoped edits and static checks completed. Package compilation is blocked by stale retired API calls in `resolve_reference_test.go`, outside this Task's writer boundary. T19 accepted the repair graph and T20 released T22 as the temporary test-only repair owner.

### Stop

Stop as `BLOCKED` without wider exploration when:

- any file outside `Change` must be modified;
- a protected type or compatibility helper must be changed;
- current behavior cannot be preserved within the exact boundary;
- a failure belongs to T13, T06, T07, T08, or an unmapped owner;
- an accepted contract must change;
- the exact verification commands cannot run.

### Prohibited operations

Do not change:

- `drmcp/src/internal/designrecords/types.go`.
- `drmcp/src/internal/designrecords/id_range.go`.
- `drmcp/src/internal/designrecords/validation.go`.
- `drmcp/src/internal/designrecords/validation_test.go`.
- `drmcp/src/internal/designrecords/authoring.go`.
- `drmcp/src/internal/designrecords/resolver.go`.
- `drmcp/src/internal/designrecords/resolve_reference_test.go`.
- any adapter file.
- any Design Record, fixture, schema, or manifest.

Do not remove `id_range.go` or its compatibility surface.
Do not perform unrelated cleanup, independent review, stage, or commit.

### Output

1. Result.
2. Changed files.
3. Focused commands and exit status.
4. Removed-symbol and retained-test evidence.
5. Blocker.
6. T15 readiness.

## Done condition

- The five obsolete package symbols are absent.
- The compact current listing JSON omits `kind`.
- The seven legacy listing tests and five legacy-only helpers are absent.
- The three legacy get tests are absent.
- All named current tests and shared helpers remain.
- Protected files remain unchanged.
- Both focused test commands pass.
- The correction is ready for T15 aggregate verification.

## Verification

Use the exact commands in `## Work`.

Use `git.inspect_diff` with:

- `cwd`: `C:\Users\imved\projects\brewprint-t05`;
- `scope`: `both`;
- `include_untracked`: `true`;
- `paths`:
  - `drmcp/src/internal/designrecords/tools.go`;
  - `drmcp/src/internal/designrecords/list_records_test.go`;
  - `drmcp/src/internal/designrecords/get_records_test.go`.

Use `git.inspect_worktree` with the same three paths and whitespace checking enabled.

Scoped Git rules:

- Do not inspect or infer repository-wide clean status.
- Do not report pre-existing changes outside the three paths as findings.
- Inspect staged, unstaged, and untracked state only for the three paths.
- Check whitespace only for the three paths.
- An LF-to-CRLF conversion warning alone is not a failure.
- If a textual patch cannot be inspected, record that limitation. Do not claim the patch was verified.

Confirm through the scoped correction diff that T14 did not modify `types.go`, `id_range.go`, `validation.go`, `validation_test.go`, `authoring.go`, `resolver.go`, or `resolve_reference_test.go`.

## Evidence

Execution result: `BLOCKED` after all scoped edits and static checks completed.

```text
release state: released
released by: DRMCP-TASK-MCP-009-12
execution mode: parallel
status: blocked
writer ownership:
  - drmcp/src/internal/designrecords/tools.go
  - drmcp/src/internal/designrecords/list_records_test.go
  - drmcp/src/internal/designrecords/get_records_test.go
parallel writer overlap: none
obsolete production symbols: absent
legacy tests and helpers: absent
compact forbidden-key assertion: present
current retained tests: present
gofmt: PASS
scoped whitespace: PASS
focused tests: package compile blocked
stage or commit: none
blocker owner: DRMCP-TASK-MCP-009-22
repair release state: released
released by: DRMCP-TASK-MCP-009-20
post-repair gate: T15 waits for T23 completion
```

Compile blocker:

```text
drmcp/src/internal/designrecords/resolve_reference_test.go
TestResolveReferenceUsesSemanticRefsFromNonRecordSpec
undefined symbol: ListRecords
undefined symbol: GetRecord
```

The stale calls are outside the T14 writer boundary and inside the accepted T06 test path.
T14 did not modify `resolve_reference_test.go`, production resolver behavior, Design Records, stage, or commit.

T13 owns only `drmcp/src/internal/designrecordsmcp/tools.go`.
T14 does not consume T13 output and may execute in parallel after release.

Finding ownership:

- B-01: legacy package behavior retirement.
- M-01: compact projection correction.
- M-04: obsolete package list and get retirement clause.
