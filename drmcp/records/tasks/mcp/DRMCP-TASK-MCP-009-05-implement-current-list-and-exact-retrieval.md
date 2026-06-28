# DRMCP-TASK-MCP-009-05: Implement current list and exact retrieval

- **id**: DRMCP-TASK-MCP-009-05
- **status**: not_started
- **date**: 2026-06-28
- **work_item**: DRMCP-WORK-MCP-009
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 2d
- **depends_on**:
  - DRMCP-TASK-MCP-009-04
- **outputs**:
  - drmcp/src/internal/designrecords/tools.go
  - drmcp/src/internal/designrecords/id_range.go
  - drmcp/src/internal/designrecords/list_records_test.go
  - drmcp/src/internal/designrecords/get_records_test.go
  - drmcp/src/internal/designrecords/get_record_test.go
  - drmcp/src/internal/designrecords/suggest_next_record_test.go
  - drmcp/src/internal/designrecordsmcp/tools.go
  - drmcp/src/internal/designrecordsmcp/tools_call.go
  - drmcp/src/internal/designrecordsmcp/tools_call_test.go
  - drmcp/src/internal/designrecordsmcp/jsonrpc_test.go
  - drmcp/src/cmd/design-records-mcp/main_test.go

## Goal

Implement compact active-current listing and ordered exact current retrieval.

Retire `get_record`, `suggest_next_record`, and obsolete range behavior from the public read surface.

## Work

- Implement `list_records` with app namespace, kind, domain, optional status, default descending order, default limit 20, and limit range 1 through 100.
- Return only compact `ref`, `title`, `status`, and `date` fields plus `has_more` and warnings.
- Exclude specs from normal listing.
- Implement `get_records` with `refs`, ordered deduplication, 1 through 20 inputs, partial success, successful records only, and top-level warnings.
- Return normalized metadata, headings, and optional body without physical paths.
- Classify malformed, unsupported, unresolved, and duplicate requested refs without repair.
- Remove the package-local retired operations and ID-range behavior.
- Inventory the bounded MCP registration surface before editing any file outside the target package.
- Add each direct retired-operation registration file to this Task's evidence before editing it.
- Keep resolver semantics in T06 and validation semantics in T07.

### Execution slices

| slice | owner model | parallel group | dependency | exact file boundary or inventory method | allowed changes | prohibited changes | commands | expected evidence | escalation condition |
|---|---|---|---|---|---|---|---|---|---|
| S05A package operations | Sonnet | P3 | T04 accepted | Exact initial files: `tools.go`, `id_range.go`, `list_records_test.go`, `get_records_test.go`, `get_record_test.go`, `suggest_next_record_test.go`. Files may be deleted when the accepted surface removes them. | Implement list/get, remove retired functions and obsolete tests/ranges. | `types.go`, config, parser, index, resolver, validation, authoring, fixtures. | `gofmt -w`; targeted read-tool tests; own-branch targeted acceptance. | Request/response examples, warning cases, no-path JSON proof, retired symbol proof. | Any need to change T03 shared types or another P3 file; stop and escalate. |
| S05B bounded catalog cleanup | Sonnet | P3 | S05A complete | Run `git grep -l -E '"get_record"|GetRecord|"suggest_next_record"|SuggestNextRecord' -- drmcp/src`. Add only returned MCP registration or registration-test files that directly expose a retired operation, and record them before editing. Initial direct catalog candidates: `drmcp/src/internal/designrecordsmcp/tools.go`, `tools_call.go`, `tools_call_test.go`, `jsonrpc_test.go`, and `drmcp/src/cmd/design-records-mcp/main_test.go`. | Remove retired public catalog entries and their direct tests. | Unrelated server code, authoring operations, repository-wide cleanup, resolver test semantics. | Exact grep; targeted package test for every added package. | Pre/post symbol inventory and exact expanded boundary. | Any match whose ownership is unclear or whose removal affects authoring; stop for boundary review. |
| S05C mechanical verification | Haiku | P3 | S05A and S05B complete | Final recorded S05 boundary; read-only. | No changes. | Contract decisions or repairs. | Expanded `gofmt -d` on all surviving S05 Go files; targeted tests; full affected-package tests; exact grep; scoped Git check. | Raw outputs, zero retired public registrations, exact changed paths. | Any failure or unexpected match; escalate to Sonnet. |

## Done condition

- `list_records` accepts only the corrected compact filters.
- Default order is descending and default limit is 20.
- Limit outside 1 through 100 is rejected.
- Specs and future legacy records are absent from normal listing.
- `get_records` accepts `refs`, preserves first-occurrence order, deduplicates, and returns successful records only.
- Partial success uses top-level warnings and no failure placeholder records.
- Normal list and retrieval JSON contains no physical path.
- `get_record` and `suggest_next_record` are absent from the public catalog.
- Obsolete ID-range behavior is absent from list and validation request paths.
- W008 cases C11, C12, C16, R02-R05, R14, and R20 are covered.
- T03 shared types and P3 sibling files remain unchanged.
- Own-branch targeted acceptance is recorded separately from P3 integration acceptance.
- Full package PASS is required before T08 only after T05, T06, and T07 are integrated.

## Verification

Run from repository root:

```powershell
git grep -n -E '"get_record"|GetRecord|"suggest_next_record"|SuggestNextRecord' -- drmcp/src
$files = @(
  "drmcp/src/internal/designrecords/tools.go",
  "drmcp/src/internal/designrecords/id_range.go",
  "drmcp/src/internal/designrecords/list_records_test.go",
  "drmcp/src/internal/designrecords/get_records_test.go",
  "drmcp/src/internal/designrecords/get_record_test.go",
  "drmcp/src/internal/designrecords/suggest_next_record_test.go",
  "drmcp/src/internal/designrecordsmcp/tools.go",
  "drmcp/src/internal/designrecordsmcp/tools_call.go",
  "drmcp/src/internal/designrecordsmcp/tools_call_test.go",
  "drmcp/src/internal/designrecordsmcp/jsonrpc_test.go",
  "drmcp/src/cmd/design-records-mcp/main_test.go"
)
# Remove deleted files from $files after S05B records the final surviving boundary.
gofmt -d -- $files
go test ./drmcp/src/internal/designrecords -run 'Test(ListRecords|GetRecords|RetiredReadTool|NormalRead.*Path)' -count=1
go test ./drmcp/src/internal/designrecords -count=1
```

When S05B adds an MCP adapter package, run its complete package test as well and record the exact command.
Deleted files are verified by scoped Git evidence, not by `gofmt`.

## Evidence

Record:

- final exact file boundary after S05B inventory;
- accepted list and retrieval JSON examples;
- exact warning classifications;
- no-path serialization assertion;
- retired symbol and catalog inventory before and after;
- targeted and full affected-package outputs;
- confirmation that resolver and validation files were not changed.
- own-branch targeted acceptance and later P3 integration acceptance as separate evidence.

### Provisional implementation mapping

```yaml
implementation_mapping:
  status: provisional
  contract_refs:
    - DRMCP-REQ-MCP-001
    - DRMCP-ADR-MCP-001
    - DRMCP-WORK-MCP-004
    - DRMCP-WORK-MCP-006
    - DRMCP-WORK-MCP-008
  fixture_cases:
    - C11
    - C12
    - C16
    - R02
    - R03
    - R04
    - R05
    - R14
    - R20
  implementation:
    - path: drmcp/src/internal/designrecords/tools.go
      symbols: []
    - path: drmcp/src/internal/designrecordsmcp/tools.go
      symbols: []
    - path: drmcp/src/internal/designrecordsmcp/tools_call.go
      symbols: []
  verification:
    - path: drmcp/src/internal/designrecords/list_records_test.go
      tests: []
    - path: drmcp/src/internal/designrecords/get_records_test.go
      tests: []
    - path: drmcp/src/internal/designrecordsmcp/tools_call_test.go
      tests: []
    - path: drmcp/src/internal/designrecordsmcp/jsonrpc_test.go
      tests: []
    - path: drmcp/src/cmd/design-records-mcp/main_test.go
      tests: []
  future_canonicalization:
    internal_design_ref: pending
    bpdsl_ref: pending
```

Populate `symbols` and `tests` with real names before Task closure.
Remove deleted, unchanged, or non-contract-significant paths from the final mapping.
