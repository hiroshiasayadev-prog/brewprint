# DRCLI-TASK-CLI-003-18: Verify complete Windows and Linux CLI acceptance

- **status**: not_started
- **date**: 2026-10-08
- **work_item**: DRCLI-WORK-CLI-003
- **task_type**: verification
- **estimate**: 2d
- **depends_on**:
  - DRCLI-TASK-CLI-003-17
- **outputs**:
  - DRCLI-TASK-CLI-003-18

## Goal

Produce one objective cross-platform integrated acceptance gate.

## Work

### Read

- `DRCLI-TASK-CLI-003-18`.
- `drcli/tests/integration/integration_test.go`
- `drcli/tests/package/package_test.go`
- `drcli/scripts/package.ps1`
- `drcli/scripts/package.sh`

### Execute

- Run `go -C drcli test ./... -count=1` and packaged acceptance on copied Windows and Linux installations. On each native OS actually run root/group/command help, all 19 commands, UTF-8 JSON/status fixtures, ignore/no-ignore, validation, authoring, retry cache and proposal across processes. Compare expected versus actual per check and record binary hashes, OS version, command and exit. A skipped Linux execution is BLOCKED, never PASS. Only objective verification; no source repair or acceptance judgment.
- Success boundary: Each defined criterion records expected/actual and the single gate reports PASS, FAIL, or BLOCKED with reproducible commands.

### Stop

- Record BLOCKED on absent exact predecessor, missing OS verification, missing authority or unapproved changes. Do not guess or rewrite another owner's work.

### Prohibited operations

- No source Specification, PRODUCT, DRMCP, historical completed Task, unrelated code or other Task ownership changes; no stage/commit.

### Output

- Task-local actual results, evidence, exact changed outputs when applicable, and downstream handoff.

## Done condition

- Each defined criterion records expected/actual and the single gate reports PASS, FAIL, or BLOCKED with reproducible commands.

## Verification

- Check exact command logs and result parity on real Windows and Linux hosts.

## Evidence

TBD
