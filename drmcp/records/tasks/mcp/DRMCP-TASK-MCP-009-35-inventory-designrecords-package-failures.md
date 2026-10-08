# DRMCP-TASK-MCP-009-35: Inventory designrecords package failures

- **id**: DRMCP-TASK-MCP-009-35
- **status**: not_started
- **date**: 2026-06-30
- **work_item**: DRMCP-WORK-MCP-009
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 0.5d
- **depends_on**:
  - DRMCP-TASK-MCP-009-32
- **outputs**:
  - DRMCP-TASK-MCP-009-35

## Goal

Run one diagnostic full-package command and produce a complete mechanical failure inventory.

Do not assign ownership or decide T05 acceptance.

## Work

### Read

Read only this Task.

Do not read T15, owner maps, sibling Tasks, production source, or Go test source.
Do not explore the repository.

### Execute

Run once from repository root:

```powershell
go test -json ./drmcp/src/internal/designrecords -count=1
```

Do not rerun the command to improve or complete output.
Do not run focused commands.

### Inventory

Record:

- command exit status;
- every exact failing top-level Go test function name;
- whether compile failure exists;
- whether package setup failure exists;
- whether package initialization failure exists;
- whether panic exists;
- whether unknown or mechanically unclassifiable failure exists;
- a concise excerpt for every non-test blocker;
- whether the captured output is sufficient to enumerate all failing top-level tests.

For JSON test events, treat an exact failed `Test` value without `/` as a top-level test name.
Nested `/` test values do not replace the required top-level failure event.
If a failing top-level name cannot be recovered mechanically from the single run, mark output completeness false.

Do not map any test to T05, T06, T07, or T08.
Do not classify failure content by feature area.

Use one result:

| result | condition |
|---|---|
| `PASS` | Command exit `0`. |
| `COMPLETE_FAILURE_INVENTORY` | Exit is non-zero, every failing top-level test name is exact, and every non-test blocker flag is false. |
| `BLOCKED_INVENTORY` | A non-test blocker exists, exact failing names are incomplete, or output is mechanically unclassifiable. |

An allowlisted test panic remains a panic.
Do not hide a panic behind a test name.

### Prohibited operations

- Do not modify files.
- Do not assign failure ownership.
- Do not judge findings or mappings.
- Do not update Design Records.
- Do not run Git inspection, grep, formatter, generator, or another Go command.
- Do not stage or commit.

### Output

1. Inventory result.
2. Command and exit status.
3. Exact failing top-level tests.
4. Non-test failure flags.
5. Output completeness.
6. Blocker excerpt.
7. T15 routing readiness.

## Done condition

- The command runs exactly once or has an exact startup blocker.
- Exit status and all failure flags are recorded.
- Every failing top-level test name is recorded when mechanically available.
- No ownership judgment is made.
- No file changes occur.

## Verification

Use only the command in `## Work`.

## Evidence

Execution pending reviewed release by T32.

The incomplete first T15 report is not a substitute for this inventory.
