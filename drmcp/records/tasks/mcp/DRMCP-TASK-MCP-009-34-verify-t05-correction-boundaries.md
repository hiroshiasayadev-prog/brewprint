# DRMCP-TASK-MCP-009-34: Verify T05 correction boundaries

- **id**: DRMCP-TASK-MCP-009-34
- **status**: not_started
- **date**: 2026-06-30
- **work_item**: DRMCP-WORK-MCP-009
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 0.5d
- **depends_on**:
  - DRMCP-TASK-MCP-009-32
- **outputs**:
  - DRMCP-TASK-MCP-009-34

## Goal

Verify only the static T05 correction boundary.

Record exact absence, presence, scoped diff, protected-path, and whitespace observations without running Go tests or judging findings.

## Work

### Read

Read only:

- this Task;
- T13, T14, T21, T22, and T23 Evidence;
- the six writer paths;
- the seven protected paths.

Do not read T15, T16, T17, or sibling leaf results.
Do not explore the repository.

### Writer paths

```text
drmcp/src/internal/designrecordsmcp/tools.go
drmcp/src/internal/designrecordsmcp/tools_call_test.go
drmcp/src/internal/designrecords/tools.go
drmcp/src/internal/designrecords/list_records_test.go
drmcp/src/internal/designrecords/get_records_test.go
drmcp/src/internal/designrecords/resolve_reference_test.go
```

### Protected paths

```text
drmcp/src/internal/designrecords/types.go
drmcp/src/internal/designrecords/config.go
drmcp/src/internal/designrecords/id_range.go
drmcp/src/internal/designrecords/validation.go
drmcp/src/internal/designrecords/validation_test.go
drmcp/src/internal/designrecords/authoring.go
drmcp/src/internal/designrecords/resolver.go
```

### Absence checks

Check only the writer paths for these removed declarations:

```text
type listRecordsScope
func ListRecords(
func GetRecords(
func newListRecordsScope(
func (s listRecordsScope) selectRecord(
func idRangeSchema(
func TestLegacyListingBasicFiltersAndResponseShape(
func TestLegacyListingStatusAndIDFilters(
func TestLegacyListingIDRangeFilter(
func TestLegacyListingWorkflowIDRangeFilter(
func TestLegacyListingRequestErrors(
func TestLegacyListingSortOrderAndLimit(
func TestLegacyListingRepositoryBootstrapQueries(
func buildListRecordsTestIndex(
func assertListRecordsErrorCode(
func listedRecordIDs(
func findListedRecord(
func latestDecisionRecordID(
func TestGetRecordsMixedKindsPartialResultsAndDuplicates(
func TestGetRecordsAllMissingIsNormalExactLookupResponse(
func TestGetRecordsRequestErrors(
```

Check `resolve_reference_test.go` for any remaining `ListRecords` or `GetRecord` identifier.

Confirm these paths are absent:

```text
drmcp/src/internal/designrecords/get_record_test.go
drmcp/src/internal/designrecords/suggest_next_record_test.go
```

Record every match by exact path and line.
A match is `FAIL`.
A command or inspection failure is `BLOCKED`.

### Presence checks

Verify these exact items:

| path | required text |
|---|---|
| `drmcp/src/internal/designrecords/tools.go` | `func ListCurrentRecords(` |
| `drmcp/src/internal/designrecords/tools.go` | `func currentListedRecord(` |
| `drmcp/src/internal/designrecords/tools.go` | `func GetCurrentRecords(` |
| `drmcp/src/internal/designrecords/tools.go` | `func currentGetRecord(` |
| `drmcp/src/internal/designrecordsmcp/tools.go` | `func Tools()` |
| `drmcp/src/internal/designrecords/list_records_test.go` | `func TestListRecordsCurrentCompactFiltersDefaultsAndNoPath(` |
| `drmcp/src/internal/designrecords/list_records_test.go` | `func TestListRecordsCurrentStatusOrderLimitAndHasMore(` |
| `drmcp/src/internal/designrecords/list_records_test.go` | `func TestListRecordsCurrentRejectsObsoleteAndInvalidInputs(` |
| `drmcp/src/internal/designrecords/list_records_test.go` | `"id", "kind", "path", "decision", "requirement", "headings", "body"` |
| `drmcp/src/internal/designrecords/get_records_test.go` | `func TestGetRecordsCurrentExactBatchSuccessWarningsAndNoPath(` |
| `drmcp/src/internal/designrecords/get_records_test.go` | `func TestGetRecordsCurrentRequestErrors(` |
| `drmcp/src/internal/designrecords/get_records_test.go` | `func TestNormalReadGetRecordsPathIsNotSerialized(` |
| `drmcp/src/internal/designrecordsmcp/tools_call_test.go` | `func TestToolsCallSuccess(` |
| `drmcp/src/internal/designrecordsmcp/tools_call_test.go` | `func TestToolsListWorkflowKindEnums(` |
| `drmcp/src/internal/designrecordsmcp/tools_call_test.go` | `func TestToolsCallToolErrors(` |
| `drmcp/src/internal/designrecords/resolve_reference_test.go` | `func TestResolveReferenceUsesSemanticRefsFromNonRecordSpec(` |
| `drmcp/src/internal/designrecords/resolve_reference_test.go` | `spec:non-record.doc` |
| `drmcp/src/internal/designrecords/resolve_reference_test.go` | `spec:non-record.section` |

A missing item is `FAIL`.
An inspection failure is `BLOCKED`.

### Exact static commands

Run this PowerShell block once from repository root for declaration and stale-call absence:

```powershell
$ErrorActionPreference = 'Stop'

$writerPaths = @(
  'drmcp/src/internal/designrecordsmcp/tools.go',
  'drmcp/src/internal/designrecordsmcp/tools_call_test.go',
  'drmcp/src/internal/designrecords/tools.go',
  'drmcp/src/internal/designrecords/list_records_test.go',
  'drmcp/src/internal/designrecords/get_records_test.go',
  'drmcp/src/internal/designrecords/resolve_reference_test.go'
)

$removedPatterns = @(
  'type\s+listRecordsScope\b',
  'func\s+ListRecords\s*\(',
  'func\s+GetRecords\s*\(',
  'func\s+newListRecordsScope\s*\(',
  'func\s+\(s\s+listRecordsScope\)\s+selectRecord\s*\(',
  'func\s+idRangeSchema\s*\(',
  'func\s+TestLegacyListingBasicFiltersAndResponseShape\s*\(',
  'func\s+TestLegacyListingStatusAndIDFilters\s*\(',
  'func\s+TestLegacyListingIDRangeFilter\s*\(',
  'func\s+TestLegacyListingWorkflowIDRangeFilter\s*\(',
  'func\s+TestLegacyListingRequestErrors\s*\(',
  'func\s+TestLegacyListingSortOrderAndLimit\s*\(',
  'func\s+TestLegacyListingRepositoryBootstrapQueries\s*\(',
  'func\s+buildListRecordsTestIndex\s*\(',
  'func\s+assertListRecordsErrorCode\s*\(',
  'func\s+listedRecordIDs\s*\(',
  'func\s+findListedRecord\s*\(',
  'func\s+latestDecisionRecordID\s*\(',
  'func\s+TestGetRecordsMixedKindsPartialResultsAndDuplicates\s*\(',
  'func\s+TestGetRecordsAllMissingIsNormalExactLookupResponse\s*\(',
  'func\s+TestGetRecordsRequestErrors\s*\('
)

$failed = $false
foreach ($pattern in $removedPatterns) {
  $matches = Select-String -Path $writerPaths -Pattern $pattern
  if ($matches) {
    $matches | ForEach-Object { Write-Output "UNEXPECTED $($_.Path):$($_.LineNumber):$($_.Line)" }
    $failed = $true
  }
}

$staleCalls = Select-String `
  -Path 'drmcp/src/internal/designrecords/resolve_reference_test.go' `
  -Pattern '\b(ListRecords|GetRecord)\b'
if ($staleCalls) {
  $staleCalls | ForEach-Object { Write-Output "STALE_CALL $($_.Path):$($_.LineNumber):$($_.Line)" }
  $failed = $true
}

foreach ($removedPath in @(
  'drmcp/src/internal/designrecords/get_record_test.go',
  'drmcp/src/internal/designrecords/suggest_next_record_test.go'
)) {
  if (Test-Path -LiteralPath $removedPath) {
    Write-Output "UNEXPECTED_PATH $removedPath"
    $failed = $true
  }
}

if ($failed) { exit 1 }
exit 0
```

Run this PowerShell block once from repository root for required presence:

```powershell
$ErrorActionPreference = 'Stop'

$required = @(
  @('drmcp/src/internal/designrecords/tools.go', 'func ListCurrentRecords('),
  @('drmcp/src/internal/designrecords/tools.go', 'func currentListedRecord('),
  @('drmcp/src/internal/designrecords/tools.go', 'func GetCurrentRecords('),
  @('drmcp/src/internal/designrecords/tools.go', 'func currentGetRecord('),
  @('drmcp/src/internal/designrecordsmcp/tools.go', 'func Tools()'),
  @('drmcp/src/internal/designrecords/list_records_test.go', 'func TestListRecordsCurrentCompactFiltersDefaultsAndNoPath('),
  @('drmcp/src/internal/designrecords/list_records_test.go', 'func TestListRecordsCurrentStatusOrderLimitAndHasMore('),
  @('drmcp/src/internal/designrecords/list_records_test.go', 'func TestListRecordsCurrentRejectsObsoleteAndInvalidInputs('),
  @('drmcp/src/internal/designrecords/list_records_test.go', '"id", "kind", "path", "decision", "requirement", "headings", "body"'),
  @('drmcp/src/internal/designrecords/get_records_test.go', 'func TestGetRecordsCurrentExactBatchSuccessWarningsAndNoPath('),
  @('drmcp/src/internal/designrecords/get_records_test.go', 'func TestGetRecordsCurrentRequestErrors('),
  @('drmcp/src/internal/designrecords/get_records_test.go', 'func TestNormalReadGetRecordsPathIsNotSerialized('),
  @('drmcp/src/internal/designrecordsmcp/tools_call_test.go', 'func TestToolsCallSuccess('),
  @('drmcp/src/internal/designrecordsmcp/tools_call_test.go', 'func TestToolsListWorkflowKindEnums('),
  @('drmcp/src/internal/designrecordsmcp/tools_call_test.go', 'func TestToolsCallToolErrors('),
  @('drmcp/src/internal/designrecords/resolve_reference_test.go', 'func TestResolveReferenceUsesSemanticRefsFromNonRecordSpec('),
  @('drmcp/src/internal/designrecords/resolve_reference_test.go', 'spec:non-record.doc'),
  @('drmcp/src/internal/designrecords/resolve_reference_test.go', 'spec:non-record.section')
)

$failed = $false
foreach ($item in $required) {
  $path = $item[0]
  $text = $item[1]
  $match = Select-String -Path $path -SimpleMatch -Pattern $text
  if (-not $match) {
    Write-Output "MISSING $path :: $text"
    $failed = $true
  }
}

if ($failed) { exit 1 }
exit 0
```

Interpret exit `0` as static-command PASS, exit `1` as frozen-item mismatch, and any other execution failure as `BLOCKED`.

### Scoped Git and whitespace checks

Use `git.inspect_diff` with:

- `cwd`: `C:\Users\imved\projects\brewprint-t05`;
- `paths`: the six writer paths;
- `scope`: `both`;
- `include_untracked`: `true`.

Use `git.inspect_worktree` twice:

- first with the six writer paths;
- second with the seven protected paths;
- `include_untracked`: `true`;
- `check_whitespace`: `true`.

Compare protected-path state only with the scoped pre-execution Evidence in T13, T14, T21, T22, and T23.
Do not infer repository-wide cleanliness.
Do not treat an LF-to-CRLF warning alone as failure.

Confirm the complete T22 diff contains only:

- deletion of the `ListRecords` assertion block;
- deletion of the `GetRecord` assertion block.

Confirm the T24-T26 current-root migration is absent.
Do not execute or judge surviving resolver assertions.

### Result

Return:

- `PASS` when every absence, presence, diff, protected-path, and whitespace check passes;
- `FAIL` when a frozen boundary item mismatches;
- `BLOCKED` when the evidence cannot be inspected mechanically.

Do not route a failure to an owner.
Do not accept a mapping or close a finding.

### Prohibited operations

- Do not modify files.
- Do not run Go tests, formatter, generator, or full-package verification.
- Do not update Design Records.
- Do not perform repository-wide traversal or cleanliness claims.
- Do not stage or commit.

### Output

1. Result: `PASS`, `FAIL`, or `BLOCKED`.
2. Absence results.
3. Presence results.
4. T22 exact-diff result.
5. Writer-path diff result.
6. Protected-path result.
7. Whitespace result.
8. Inspection limitation or blocker.
9. T36 readiness.

## Done condition

- Every frozen absence and presence item has an observation.
- The six writer paths have direct diff evidence.
- The seven protected paths have scoped state and whitespace evidence.
- T22 exact-diff and rollback observations exist.
- No Go command or ownership judgment occurs.
- No file changes occur.

## Verification

Use only the exact paths and checks in `## Work`.

## Evidence

Execution pending reviewed release by T32.
