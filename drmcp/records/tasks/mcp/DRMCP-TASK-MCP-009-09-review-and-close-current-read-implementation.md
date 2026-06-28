# DRMCP-TASK-MCP-009-09: Review and close current read implementation

- **id**: DRMCP-TASK-MCP-009-09
- **status**: not_started
- **date**: 2026-06-28
- **work_item**: DRMCP-WORK-MCP-009
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 1d
- **depends_on**:
  - DRMCP-TASK-MCP-009-08
- **outputs**:
  - DRMCP-WORK-MCP-009
  - DRMCP-TASK-MCP-001-09

## Goal

Independently review the complete current-format read implementation and close W009 and hub T09 only after acceptance.

Confirm contract traceability, fixture coverage, changed-file boundaries, test evidence, and exclusion of legacy and retained spec-validator behavior.

## Work

- Review T02 through T08 evidence and final merged implementation.
- Compare behavior with W003 through W008 and accepted current specifications.
- Verify every implementation file belongs to an accepted Task boundary.
- Verify public retirement of `get_record` and `suggest_next_record`.
- Verify current-only configuration, parsing, indexing, list, exact retrieval, resolver, validation, diagnostics, and path hiding.
- Verify complete current fixture integration and authoring non-regression.
- Verify every W009-owned contract and major production symbol has a provisional implementation mapping.
- Verify every production mapping has a corresponding verification mapping.
- Verify every fixture-owned behavior has a fixture case mapping.
- Verify every mapped path, Go symbol, and Go test function exists in the accepted implementation state.
- Verify no Internal Design or BPDSL canonical ID was guessed or invented.
- Accept `pending` only because the formal Internal Design or BPDSL foundation is not operational.
- Verify Task Evidence remains execution traceability and is not treated as current-state implementation-design authority.
- Build the final surviving Go-file inventory from T02 through T08 evidence and scoped Git inventory.
- Run final `gofmt -d` only on that surviving changed-file inventory.
- Classify findings as blocking, major, minor, or advisory.
- Apply corrections only through the owning Task boundary.
- Repeat review after every blocking or major correction.
- Synchronize T09, W009, and hub T09 to `done` only after a PASS with no blocking or major finding.

### Execution slices

| slice | owner model | parallel group | dependency | exact file boundary or inventory method | allowed changes | prohibited changes | commands | expected evidence | escalation condition |
|---|---|---|---|---|---|---|---|---|---|
| S09A independent review | Sonnet | P5 | T08 accepted | Read-only review of W009 Task-owned files and exact T05 catalog expansion. | No file changes. | Self-correction during review, new design authority, legacy implementation. | Full affected-package tests, final changed-boundary `gofmt -d`, exact retired-symbol grep, scoped Git and whitespace checks. | Verdict, finding list, previous-finding disposition, contract and fixture assessment, final format inventory. | Any blocking or major finding; return to the owning Task. |
| S09B correction | Sonnet | P5 | Explicit S09A finding | Only files already owned by the Task named in the finding. Boundary must be recorded before edit. | Minimal correction and focused tests. | Cross-Task opportunistic cleanup or authority changes. | Owning Task commands plus full affected-package tests. | Finding-to-correction trace and new outputs. | Correction requires an unowned file or contract decision; stop and create a reviewed boundary change. |
| S09C closure synchronization | Haiku | P5 | Independent PASS and all required commands accepted | `DRMCP-TASK-MCP-009-09`, W009, and hub T09 only. | Set final statuses and append exact accepted evidence pointers. | Go, tests, fixtures, other Tasks, ADRs, Requirements, Specifications. | Lifecycle shape check and scoped Git/whitespace check. | Any status mismatch, unresolved finding, or missing evidence; do not close and escalate to Sonnet. |

## Done condition

- T01 through T08 are `done`.
- All accepted W009 behavior is implemented and tested.
- All W009-owned current fixture cases are covered.
- Full affected-package tests pass.
- Final `gofmt -d` passes on every surviving Go file changed by T02 through T08.
- Authoring tests pass without authoring source changes.
- `get_record` and `suggest_next_record` are absent from the public catalog.
- No normal read projection exposes physical paths.
- No legacy fallback, legacy active index, legacy validation subject, or legacy authoring behavior is implemented.
- W-SPEC-001 and W-SPEC-002 behavior is not claimed.
- T02 through T08 provisional mappings contain accepted contract refs, owned fixture cases, real production paths and contract-significant symbols, and real verification paths and test functions.
- Future canonicalization remains possible without treating Task Evidence as the current-state design source of truth.
- No Internal Design or BPDSL canonical ID, schema, artifact, or sidecar manifest is invented.
- Independent review reports PASS with no blocking or major finding.
- T09, W009, and hub T09 are synchronized to `done` only after that PASS.
- Final evidence records exact commits or worktree state, commands, outputs, review verdict, corrections, and limitations.

## Verification

Run from repository root:

```powershell
$files = @(
  "drmcp/src/internal/designrecords/config.go",
  "drmcp/src/internal/designrecords/config_test.go",
  "drmcp/src/internal/designrecords/types.go",
  "drmcp/src/internal/designrecords/types_test.go",
  "drmcp/src/internal/designrecords/parser.go",
  "drmcp/src/internal/designrecords/parser_index_test.go",
  "drmcp/src/internal/designrecords/index.go",
  "drmcp/src/internal/designrecords/index_test.go",
  "drmcp/src/internal/designrecords/tools.go",
  "drmcp/src/internal/designrecords/id_range.go",
  "drmcp/src/internal/designrecords/list_records_test.go",
  "drmcp/src/internal/designrecords/get_records_test.go",
  "drmcp/src/internal/designrecords/get_record_test.go",
  "drmcp/src/internal/designrecords/suggest_next_record_test.go",
  "drmcp/src/internal/designrecords/resolver.go",
  "drmcp/src/internal/designrecords/resolve_reference_test.go",
  "drmcp/src/internal/designrecords/validation.go",
  "drmcp/src/internal/designrecords/validation_test.go",
  "drmcp/src/internal/designrecords/current_read_fixture_test.go",
  "drmcp/src/internal/designrecordsmcp/tools.go",
  "drmcp/src/internal/designrecordsmcp/tools_call.go",
  "drmcp/src/internal/designrecordsmcp/tools_call_test.go",
  "drmcp/src/internal/designrecordsmcp/jsonrpc_test.go",
  "drmcp/src/cmd/design-records-mcp/main_test.go"
)
# Replace this list with the final surviving changed-file inventory from T02-T08 evidence and scoped Git inventory.
gofmt -d -- $files
go test ./drmcp/src/internal/designrecords -count=1
git grep -n -E '"get_record"|GetRecord|"suggest_next_record"|SuggestNextRecord' -- drmcp/src
```

Run the full test command for every additional package in the accepted T05 catalog boundary.
Verify every T02 through T08 `implementation_mapping` path, symbol, and test function against the accepted files.
Verify each `contract_refs` entry names an accepted record and each `fixture_cases` entry remains inside that Task's ownership.
Treat `pending` future refs as valid only while the formal canonical foundation is unavailable.
Deleted files are excluded from `$files` and verified through scoped Git evidence.
Run `git.inspect_worktree` for only the W009 changed-file manifest.
Do not infer repository-wide cleanliness.

## Evidence

Required final review output:

1. Verdict: PASS / NEEDS REVISION
2. Previous-finding disposition
3. Blocking findings
4. Major findings
5. Minor findings
6. Advisories
7. Contract traceability assessment
8. File-boundary and parallel-merge assessment
9. Current fixture coverage assessment
10. Public tool-surface assessment
11. Path-hiding and diagnostic-location assessment
12. Legacy and retained-validator exclusion assessment
13. Test and command evidence assessment
14. Final changed-boundary format assessment
15. Provisional implementation mapping and future canonicalization assessment
16. T09 closure readiness
17. W009 and hub T09 closure readiness

Final evidence is pending Task execution.
