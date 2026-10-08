# DRMCP-TASK-MCP-009-17: Synchronize T05 correction closure

- **id**: DRMCP-TASK-MCP-009-17
- **status**: not_started
- **date**: 2026-06-30
- **work_item**: DRMCP-WORK-MCP-009
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 0.5d
- **depends_on**:
  - DRMCP-TASK-MCP-009-16
- **outputs**:
  - DRMCP-TASK-MCP-009-05
  - DRMCP-TASK-MCP-009-13
  - DRMCP-TASK-MCP-009-14
  - DRMCP-TASK-MCP-009-15
  - DRMCP-TASK-MCP-009-16
  - DRMCP-TASK-MCP-009-17
  - DRMCP-WORK-MCP-009

## Goal

Persist accepted T05 correction verification, finding disposition, and complete replacement mapping.

Close only T05 while leaving W009 and downstream T06, T07, T08, and T09 gates open.

## Work

### Read

Read only:

- W009;
- T05, T13, T14, T15, T16, and T17;
- T21, T22, and T23;
- T30 through T36.

Do not read production source or Go tests.
Do not explore the repository.

### Start gate

Require:

- T23 `done`;
- T31 `PASS` and T32 `done`;
- T33 through T36 complete with Evidence synchronized by T15;
- T15 `done` with overall `PASS`;
- T15 diagnostic result `PASS` or `ROUTED_DOWNSTREAM`;
- T16 verdict `PASS`;
- B-01, M-01, and M-04 `CLOSED`;
- no open blocking, major, or minor direct-regression finding;
- `mapping manifest accepted: yes`;
- exact accepted T15 mapping present and identical to T36.

Stop as `BLOCKED` without lifecycle changes when any prerequisite is missing or inconsistent.

### Synchronize

- Record the accepted T33 focused results in T13, T14, and T05 Evidence where relevant.
- Preserve the T21 and T22 focused Evidence already synchronized by T23.
- Record T34 static boundary results in T05 Evidence.
- Record T35 diagnostic result and exact downstream routing in T05 and W009 Evidence.
- Do not call `ROUTED_DOWNSTREAM` a full-package PASS or downstream acceptance.
- Record the T16 verdict, finding dispositions, and mapping acceptance.
- Set T13, T14, T15, and T16 to `done` only when their own Done conditions are satisfied.
- Replace T05 `implementation_mapping` with the exact T16-accepted T15 manifest.
- Copy every field and value field-for-field and value-for-value.
- Set T05 to `done` only after complete accepted Evidence and mapping replacement.
- Keep W009 `in_progress`.
- Record that the T05 dependency of T08 is satisfied.
- Record that T08 remains not start-eligible until T06 and T07 are also accepted.
- Set T17 to `done` after synchronization is verified.

### Allowed changes

- T05;
- T13 through T17;
- W009.

### Prohibited operations

- Do not modify T06, T07, T08, T09, T21 through T36, production source, Go tests, fixtures, schemas, manifests, ADRs, Requirements, or Specifications.
- Do not close W009.
- Do not claim T06, T07, T08, or T09 acceptance.
- Do not claim full-package PASS for `ROUTED_DOWNSTREAM`.
- Do not rerun implementation, verification, or review.
- Do not merge, supplement, convert, or reinterpret the accepted mapping.
- Do not stage or commit.

### Output

1. Result.
2. Updated Design Records.
3. Persisted verification and routing.
4. Finding disposition.
5. T05 status.
6. W009 status.
7. T08 dependency state.
8. Blocker.

## Done condition

- Accepted T13 through T16 Evidence is persisted.
- B-01, M-01, and M-04 are recorded closed.
- T05 contains the exact accepted complete replacement mapping.
- T05 is `done`.
- W009 remains `in_progress`.
- T08 has only its T05 dependency satisfied and remains gated by T06 and T07.
- No downstream acceptance or full-package PASS is falsely claimed.
- No production, test, fixture, or authority artifact changes.

## Verification

- Confirm all start-gate Evidence directly from named records.
- Confirm T05 mapping equals the accepted T15 manifest field-for-field and value-for-value.
- Confirm T13 through T17 lifecycle states match Evidence.
- Confirm W009 remains `in_progress`.
- Confirm T08 still depends on T05, T06, and T07.
- Confirm only allowed Design Records changed.
- Run scoped Git and whitespace inspection on allowed records.

## Evidence

Closure synchronization is pending T15 aggregate `PASS` and T16 independent `PASS`.

T17 remains the sole T05 correction closure writer.
T09 remains the sole final W009 closure owner.
