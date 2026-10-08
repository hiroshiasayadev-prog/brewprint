# DRMCP-TASK-MCP-009-13: Retire unused adapter helper

- **id**: DRMCP-TASK-MCP-009-13
- **status**: blocked
- **date**: 2026-06-29
- **work_item**: DRMCP-WORK-MCP-009
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 0.5d
- **depends_on**:
  - DRMCP-TASK-MCP-009-12
- **outputs**:
  - drmcp/src/internal/designrecordsmcp/tools.go

## Goal

Remove the unused adapter helper that remains after obsolete range schema call sites were removed.

## Work

### Read

- This Task.
- `drmcp/src/internal/designrecordsmcp/tools.go`.

### Change

Only change `drmcp/src/internal/designrecordsmcp/tools.go`.

### Implement

- Delete the complete `idRangeSchema` function.
- Preserve every other tool definition, schema, description, helper, and ordering.

### Test

Run:

```powershell
gofmt -w drmcp/src/internal/designrecordsmcp/tools.go
go test ./drmcp/src/internal/designrecordsmcp -run '^TestTools(ListWorkflowKindEnums|CallToolErrors)$' -count=1
git grep -n 'func idRangeSchema' -- drmcp/src/internal/designrecordsmcp/tools.go
```

The test must pass. The grep must find no occurrence. T15 owns full package verification.

The source edit and symbol check completed. The focused test is blocked by a stale Config fixture in `tools_call_test.go`, outside this Task's writer boundary. T19 accepted the repair graph and T20 released T21 as the test-only repair owner.

### Stop

Stop as `BLOCKED` when another file or another schema must change, the failure is outside this helper removal, or the commands cannot run.

### Prohibited operations

- Do not modify tests or other helpers.
- Do not modify Design Records.
- Do not perform unrelated cleanup or independent review.
- Do not stage or commit.

### Output

1. Result.
2. Changed file.
3. Command results.
4. Blocker.
5. T15 readiness.

## Done condition

- `idRangeSchema` is absent.
- No other behavior or file changes.
- Focused adapter tests pass.
- T13 scoped implementation is ready for T15 after T21, T22, and T23 complete.

## Verification

Run the exact commands in `## Work`.
Inspect the scoped diff and whitespace for `tools.go` only.

## Evidence

Execution result: `BLOCKED` after the scoped source edit completed.

```text
release state: released
released by: DRMCP-TASK-MCP-009-12
execution mode: parallel
status: blocked
writer ownership:
  - drmcp/src/internal/designrecordsmcp/tools.go
parallel writer overlap: none
source edit: complete
idRangeSchema absence: PASS
gofmt: PASS
scoped whitespace: PASS
focused test: FAIL
stage or commit: none
blocker owner: DRMCP-TASK-MCP-009-21
repair release state: released
released by: DRMCP-TASK-MCP-009-20
post-repair gate: T15 waits for T23 completion
```

Focused failure:

```text
TestToolsCallToolErrors/get_authoring_guidance_unknown_id
current roots must not be empty: auto-discovery has been removed
```

The stale fixture is in `drmcp/src/internal/designrecordsmcp/tools_call_test.go` and uses `designrecords.Config{Root: "."}`.
That path is outside the T13 writer boundary.
T13 did not modify tests, Design Records, stage, or commit.

T14 owns only the three `internal/designrecords` writer paths.
T13 does not consume T14 output and may execute in parallel after release.

Finding ownership: M-04 adapter-helper retirement.
