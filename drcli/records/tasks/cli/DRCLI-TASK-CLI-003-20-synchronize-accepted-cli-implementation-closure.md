# DRCLI-TASK-CLI-003-20: Synchronize implementation review and Work Item closure

- **status**: not_started
- **date**: 2026-10-08
- **work_item**: DRCLI-WORK-CLI-003
- **task_type**: synchronization
- **estimate**: 0.5d
- **depends_on**:
  - DRCLI-TASK-CLI-003-19
- **outputs**:
  - DRCLI-TASK-CLI-003-20
  - DRCLI-WORK-CLI-003

## Goal

Synchronize final accepted implementation outcome into parent Work Item lifecycle.

## Work

### Read

- `DRCLI-TASK-CLI-003-20`.
- `drcli/records/tasks/cli/DRCLI-TASK-CLI-003-18-verify-complete-portable-cli-integration.md`
- `drcli/records/tasks/cli/DRCLI-TASK-CLI-003-19-independently-review-drcli-implementation.md`
- `drcli/records/work-items/cli/DRCLI-WORK-CLI-003-implement-portable-design-records-cli.md`

### Execute

- Only after all T04-T19 Tasks are DONE, T18 records true both-OS PASS, T19 independently returns PASS and all later required finding-closure reviews are CLOSED, propagate already accepted results into this Task Evidence and parent Work Item Evidence/status DONE. Include final binary and verification provenance and each Work Item Completion Condition. On NEEDS REVISION, fail, missing OS or open finding leave parent open, route fresh correction/review through coordination; do not pre-create them. No code or new design judgment.
- Success boundary: Recorded lifecycle and completion evidence exactly match accepted upstream outcomes; status DONE never precedes findings closure.

### Stop

- Record BLOCKED on absent exact predecessor, missing OS verification, missing authority or unapproved changes. Do not guess or rewrite another owner's work.

### Prohibited operations

- No source Specification, PRODUCT, DRMCP, historical completed Task, unrelated code or other Task ownership changes; no stage/commit.

### Output

- Task-local actual results, evidence, exact changed outputs when applicable, and downstream handoff.

## Done condition

- Recorded lifecycle and completion evidence exactly match accepted upstream outcomes; status DONE never precedes findings closure.

## Verification

- Cross-check T19 and any closure-review verdict before parent status change.

## Evidence

TBD
