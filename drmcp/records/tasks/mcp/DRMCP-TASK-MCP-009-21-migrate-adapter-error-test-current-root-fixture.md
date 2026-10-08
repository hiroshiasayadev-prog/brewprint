# DRMCP-TASK-MCP-009-21: Migrate adapter error test current-root fixture

- **id**: DRMCP-TASK-MCP-009-21
- **status**: done
- **date**: 2026-06-30
- **work_item**: DRMCP-WORK-MCP-009
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 0.5d
- **depends_on**:
  - DRMCP-TASK-MCP-009-20
- **outputs**:
  - drmcp/src/internal/designrecordsmcp/tools_call_test.go

## Goal

Migrate the stale `TestToolsCallToolErrors` server Config fixture to the accepted explicit current-root contract.

Preserve every tool-error case and expected error code.

## Work

### Read

- This Task.
- `drmcp/src/internal/designrecordsmcp/tools_call_test.go`.
- `drmcp/src/internal/designrecords/config.go` only for the accepted Config field contract.

### Change

Only change:

- `drmcp/src/internal/designrecordsmcp/tools_call_test.go`.

### Implement

In `TestToolsCallToolErrors`, replace only this stale server Config fixture:

```go
designrecords.Config{Root: "."}
```

Use a Config with:

- `Root: "."`;
- exactly one `RecordsRoots` entry;
- `AppNamespace: "drmcp"`;
- `RecordsRoot: "drmcp/records"`;
- `NamespacePrefix: "DRMCP-"`.

Do not change:

- the table cases;
- JSON-RPC request strings;
- expected error codes;
- the custom index builder;
- any other test or Config fixture.

### Test

Run from repository root:

```powershell
gofmt -w drmcp/src/internal/designrecordsmcp/tools_call_test.go

go test ./drmcp/src/internal/designrecordsmcp -run '^TestToolsCallToolErrors$' -count=1
```

The command must exit `0`.
T23 persists this focused result. T15 owns the combined adapter command and aggregate correction verification.

### Stop

Stop as `BLOCKED` without wider exploration when:

- another file must change;
- a production Config contract must change;
- an expected tool error code or request case must change;
- the focused command fails for a reason outside the exact fixture replacement;
- the exact command cannot run.

### Prohibited operations

- Do not modify production source.
- Do not modify other tests or helpers.
- Do not update Design Records.
- Do not perform unrelated cleanup or independent review.
- Do not stage or commit.
- Do not perform repository-wide traversal.

### Output

1. Result.
2. Changed file.
3. Exact fixture replacement.
4. Focused command and exit status.
5. Blocker.
6. T23 readiness.

## Done condition

- `TestToolsCallToolErrors` uses one non-empty explicit current root.
- Every table case and expected error code remains unchanged.
- No other file or test fixture changes.
- The focused command exits `0`.
- T21 is ready for T23 Evidence synchronization after T22 also completes.

## Verification

- Use the exact command in `## Work`.
- Inspect the scoped diff and whitespace for `tools_call_test.go` only.
- Confirm only the Config literal in `TestToolsCallToolErrors` changed.
- Do not infer repository-wide cleanliness.

## Evidence

Release state: `released` by `DRMCP-TASK-MCP-009-20`.

Execution mode: originally parallel with T22.
Before the valid retry recorded below, the Task status remained `not_started`.

A first executor report claimed `PASS`, but the reported Config replacement is absent from the worktree.
`TestToolsCallToolErrors` still contains `designrecords.Config{Root: "."}`.
The reported result is rejected and must not be synchronized by T23.

After rejection of the first report, T21 remained released under the existing contract and required a clean retry.
At that point, the next gate was T23 Evidence synchronization after a valid persisted T21 result and corrected T22 result became available.

Blocker source:

```text
TestToolsCallToolErrors/get_authoring_guidance_unknown_id
current roots must not be empty: auto-discovery has been removed
```

The failure is outside T13's source-only writer boundary.

### Valid retry execution Evidence

```yaml
result: PASS
changed_file: drmcp/src/internal/designrecordsmcp/tools_call_test.go
owned_test: TestToolsCallToolErrors
config_fixture:
  Root: "."
  RecordsRoots:
    - AppNamespace: "drmcp"
      RecordsRoot: "drmcp/records"
      NamespacePrefix: "DRMCP-"
records_roots_entry_count: 1
test_cases_preserved: 16
expected_codes_preserved:
  invalid_request: 15
  guide_not_found: 1
verification:
  gofmt_exit: 0
  focused_test_exit: 0
scoped_diff: PASS
whitespace: PASS
staged_changes: none
repository_wide_clean: not inspected and not inferred
blocker: none
t23_readiness: READY
```

The complete `tools_call_test.go` diff contains broader pre-existing T05 changes.
Only the `TestToolsCallToolErrors` Config fixture replacement is attributed to T21; T21 does not own the complete file diff.

T21 is complete. Its valid focused repair Evidence is persisted by `DRMCP-TASK-MCP-009-23`.
