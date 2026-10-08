# DRMCP-TASK-MCP-009-33: Run T05 focused verification

- **id**: DRMCP-TASK-MCP-009-33
- **status**: not_started
- **date**: 2026-06-30
- **work_item**: DRMCP-WORK-MCP-009
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 0.5d
- **depends_on**:
  - DRMCP-TASK-MCP-009-32
- **outputs**:
  - DRMCP-TASK-MCP-009-33

## Goal

Run only the T05 correction-focused Go verification commands.

Return exact command results without aggregate routing or finding judgment.

## Work

### Read

Read only this Task.

Do not read T15, sibling leaves, production source, or Go test source unless a command cannot start.
Do not explore the repository.

### Execute

Run from repository root:

```powershell
go test ./drmcp/src/internal/designrecords -run '^TestListRecordsCurrent' -count=1

go test ./drmcp/src/internal/designrecords -run '^(TestGetRecordsCurrent|TestNormalReadGetRecords)' -count=1

go test ./drmcp/src/internal/designrecordsmcp -run '^TestTools(CallSuccess|ListWorkflowKindEnums|CallToolErrors)$' -count=1
```

Run each command once.
Do not run the full package command.

### Record

For each command, record:

- exact command;
- exit status;
- exact failing top-level test names when non-zero;
- concise failure excerpt when non-zero;
- whether the command could start.

Result is `PASS` only when all three commands exit `0`.
Otherwise result is `FAIL` or `BLOCKED` with raw evidence.

Do not assign an owner, accept a mapping, close a finding, or determine T15 overall result.

### Prohibited operations

- Do not modify files.
- Do not run static boundary or mapping checks.
- Do not run the full package test.
- Do not update Design Records.
- Do not stage or commit.

### Output

1. Result: `PASS`, `FAIL`, or `BLOCKED`.
2. Three command results and exit status.
3. Exact failing test names when present.
4. Blocker.
5. T15 evidence readiness.

## Done condition

- Each command runs once or has an exact execution blocker.
- Every exit status is recorded.
- No full-package command runs.
- No ownership or finding judgment is made.
- No file changes occur.

## Verification

Use only the three commands in `## Work`.

## Evidence

Execution pending reviewed release by T32.
