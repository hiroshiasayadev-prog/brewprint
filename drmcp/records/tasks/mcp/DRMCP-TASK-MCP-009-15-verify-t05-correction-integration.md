# DRMCP-TASK-MCP-009-15: Verify T05 correction integration

- **id**: DRMCP-TASK-MCP-009-15
- **status**: not_started
- **date**: 2026-06-30
- **work_item**: DRMCP-WORK-MCP-009
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 0.5d
- **depends_on**:
  - DRMCP-TASK-MCP-009-33
  - DRMCP-TASK-MCP-009-34
  - DRMCP-TASK-MCP-009-35
  - DRMCP-TASK-MCP-009-36
- **outputs**:
  - DRMCP-TASK-MCP-009-15
  - DRMCP-TASK-MCP-009-33
  - DRMCP-TASK-MCP-009-34
  - DRMCP-TASK-MCP-009-35
  - DRMCP-TASK-MCP-009-36

## Goal

Synchronize the four released verification-leaf reports and determine the T05 correction aggregate result.

Route diagnostic package failures through the reviewed persistent owner map without executing commands or judging finding closure.

## Work

### Read

Read only:

- this Task;
- T23;
- T30;
- T33, T34, T35, and T36;
- the complete executor report for each leaf.

Do not read production source, Go tests, fixtures, T16, or T17.
Do not explore the repository.

### Start gate

Require:

- T32 `done` with T31 `PASS` and no blocking, major, or minor finding;
- one complete report for each of T33 through T36;
- T36 executed only after a complete T34 report;
- no report claims source modification, Design Record modification, stage, or commit.

Stop as `BLOCKED` when a required report is missing, ambiguous, or inconsistent with its Task contract.

### Synchronize leaf Evidence

For each leaf:

- copy the exact command or observation result into its Evidence;
- record exact exit status, failing test names, item observations, and limitations;
- set the leaf to `done` only when its own Done condition is satisfied;
- set the leaf to `blocked` when its required result cannot be established;
- do not repair a failure or reinterpret raw evidence.

T15 is the sole writer for T33 through T36 execution Evidence and lifecycle synchronization.

### Diagnostic routing

Use only the exact owner map in T30.
Do not infer ownership from failure text, feature name, prefix, regex, source-file guess, or category.

Classify T35 as follows:

| T35 result | condition | T15 diagnostic result |
|---|---|---|
| `PASS` | Full package exit `0`. | `PASS` |
| `COMPLETE_FAILURE_INVENTORY` | Every exact test maps only to T06, T07, or T08. | `ROUTED_DOWNSTREAM` |
| `COMPLETE_FAILURE_INVENTORY` | Any exact test maps to T05. | `BLOCKED` |
| `COMPLETE_FAILURE_INVENTORY` | Any exact test is unmapped. | `BLOCKED` |
| `BLOCKED_INVENTORY` | Any compile, setup, initialization, panic, unknown, incomplete-name, or mechanically unclassifiable condition exists. | `BLOCKED` |

A mixed set containing only T06, T07, and T08 names is `ROUTED_DOWNSTREAM`.
A mixed set containing T05 or an unmapped name is `BLOCKED`.

For `ROUTED_DOWNSTREAM`, record each exact failing test and its persistent owner Task.
Do not claim that the failure is accepted, fixed, or semantically correct.
Do not claim full-package PASS.

### Aggregate result

Set overall result `PASS` only when:

- T33 result is `PASS`;
- T34 result is `PASS`;
- T36 result is `PASS`;
- diagnostic result is `PASS` or `ROUTED_DOWNSTREAM`;
- no leaf Evidence is missing or contradictory.

Set overall result `BLOCKED` for every other state.

A `ROUTED_DOWNSTREAM` diagnostic does not prevent T15 overall `PASS`.
T08 retains the first accepted full-package PASS after T05, T06, and T07 integration.

### Mapping proposal

When overall result is `PASS`:

- copy the exact T36 frozen manifest into T15 Evidence as `final_implementation_mapping`;
- preserve every field and value without merge, supplementation, conversion, or discretionary reordering;
- keep manifest status `proposed_for_independent_acceptance`.

Do not accept the mapping.
T16 is the only mapping acceptance owner.

### Lifecycle

- Set T15 to `done` only when overall result is `PASS` and complete Evidence is persisted.
- Set T15 to `blocked` when overall result is `BLOCKED`.
- Do not update T05, T13, T14, T16, T17, T21, T22, T23, T08, T09, or W009.

### Allowed changes

- T33;
- T34;
- T35;
- T36;
- T15.

### Prohibited operations

- Do not run Go, Git, grep, formatter, generator, or filesystem inspection commands.
- Do not modify production source, Go tests, fixtures, schemas, manifests, ADRs, Requirements, or Specifications.
- Do not repair a leaf result.
- Do not assign ownership outside the T30 exact map.
- Do not judge B-01, M-01, or M-04 closure.
- Do not accept the mapping.
- Do not start T16 or T17.
- Do not stage or commit.

### Output

1. Overall result: `PASS` or `BLOCKED`.
2. Synchronized leaf statuses.
3. T33 focused result.
4. T34 boundary result.
5. T35 diagnostic result and exact routing.
6. T36 mapping observation result.
7. Exact proposed complete replacement manifest when eligible.
8. T16 readiness.
9. Blocker.

## Done condition

- T33 through T36 Evidence matches the supplied reports exactly.
- Every leaf has a lifecycle state consistent with its own Done condition.
- Diagnostic routing uses exact equality against the reviewed T30 owner map.
- T06, T07, and T08 failures do not decide T05 finding closure.
- T05-mapped, unmapped, non-test, and unclassifiable failures block T15.
- No command runs during T15.
- The proposed mapping is copied exactly from T36 only after all aggregate gates pass.
- T15 records overall `PASS` with complete Evidence.
- T16 receives an unambiguous review gate.

## Verification

- Compare each leaf Evidence with its executor report.
- Compare every T35 failing test name with T30 by exact equality.
- Confirm T36 manifest equals the T15 proposed manifest field-for-field and value-for-value.
- Confirm only T15 and T33 through T36 changed.
- Run no repository command during this Task.

## Evidence

The first pre-amendment T15 executor report is not accepted as T15 Evidence because it omitted exact failing test names and executed the defective aggregate contract.

T15 remains `not_started` pending T31 review, T32 release, and T33 through T36 execution.
