# DRMCP-TASK-MCP-009-23: Synchronize T05 repair execution evidence

- **id**: DRMCP-TASK-MCP-009-23
- **status**: done
- **date**: 2026-06-30
- **work_item**: DRMCP-WORK-MCP-009
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 0.5d
- **depends_on**:
  - DRMCP-TASK-MCP-009-21
  - DRMCP-TASK-MCP-009-22
- **outputs**:
  - DRMCP-TASK-MCP-009-21
  - DRMCP-TASK-MCP-009-22
  - DRMCP-TASK-MCP-009-23

## Goal

Persist the completed T21 and T22 focused repair evidence before T15 aggregate verification.

Keep implementation, aggregate verification, independent review, and final correction closure separate.

## Work

### Read

Read only:

- this Task;
- T21 and its executor result;
- T22 and its executor result.

Do not read production source or tests.
Do not rerun implementation or verification commands.

### Start gate

Require T21 to report and persist:

- only `drmcp/src/internal/designrecordsmcp/tools_call_test.go` changed;
- the exact current-root Config fixture replacement, including `AppNamespace: "drmcp"`;
- focused command exit `0`;
- scoped diff and whitespace result;
- no stage or commit.

Reject any summary whose claimed change is absent from the worktree.

Require corrected T22 to report:

- only `drmcp/src/internal/designrecords/resolve_reference_test.go` changed;
- the two named stale assertion blocks remain removed;
- the T24 through T26 current-root migration is fully rolled back;
- the fixture path is `records/spec/non-record.md`;
- the named test uses `idx := buildTestIndex(t, root)`;
- the expected document target path is `records/spec/non-record.md`;
- all surviving document and section resolver assertions are preserved without a T05 semantic claim;
- shared `buildTestIndex`, `validation_test.go`, and `resolver.go` remain unchanged;
- compile-only command `go test ./drmcp/src/internal/designrecords -run '^$' -count=1` exits `0`;
- stale-call absence check returns exit `1`;
- scoped diff and whitespace result;
- no stage or commit.

Do not require the stale resolver test to execute or pass.
T06 owns resolver assertion migration and behavioral verification.

Stop as `BLOCKED` when either result is missing, ambiguous, or inconsistent with its Task contract.

### Synchronize

- Record the exact T21 changed file, fixture replacement, commands, exit status, diff result, and whitespace result in T21 Evidence.
- Record the exact T22 changed file, removed blocks, rollback result, compile-only command, exit status, absence result, diff result, protected-path state, and whitespace result in T22 Evidence.
- Record explicitly that T22 makes no resolver semantic acceptance claim.
- Set T21 and T22 to `done` only when their own Done conditions are satisfied.
- Set T23 to `done` after both Task records and lifecycle states are verified.
- Do not update T13, T14, T15, T16, T17, T05, or W009.

### Allowed changes

- `DRMCP-TASK-MCP-009-21-migrate-adapter-error-test-current-root-fixture.md`.
- `DRMCP-TASK-MCP-009-22-remove-retired-list-get-assertions-from-resolver-test.md`.
- `DRMCP-TASK-MCP-009-23-synchronize-t05-repair-execution-evidence.md`.

### Prohibited operations

- Do not modify production source or Go tests.
- Do not rerun commands.
- Do not perform independent review or aggregate verification.
- Do not claim resolver behavioral acceptance for T22.
- Do not close T13, T14, T05, or W009.
- Do not modify T15, T16, T17, T08, or T09.
- Do not stage or commit.
- Do not perform repository-wide traversal.

### Output

1. Result.
2. Updated Design Records.
3. T21 status and evidence summary.
4. T22 status and compile-only evidence summary.
5. Resolver-ownership statement.
6. T23 status.
7. T15 readiness.
8. Blocker.

## Done condition

- T21 focused repair evidence is complete and its status is `done`.
- T22 stale-call-only repair evidence is complete and its status is `done`.
- T22 records no resolver behavior acceptance claim.
- T23 records no implementation, review, or aggregate-verification claim.
- Only T21 through T23 Design Records changed.
- T15 has persistent focused evidence for both repair leaves.

## Verification

- Confirm T21 and T22 Evidence matches the supplied executor results exactly.
- Confirm each Task Done condition is satisfied before its status changes.
- Confirm T22 Evidence uses compile-only verification and does not claim the named resolver test passed.
- Confirm T23 does not claim T15 aggregate acceptance or T16 finding closure.
- Confirm only T21 through T23 records changed.
- Run scoped Git and whitespace inspection on those three records.

## Evidence

T20 released T21 and T22 for parallel execution.

The first T21 report is rejected because its claimed Config replacement is absent from the worktree.

The first T22 execution completed the stale-call deletion but failed on shared stale initialization.

T24 through T26 released a current-root migration. The second T22 execution followed that contract but exposed a contradiction with accepted current spec and resolver authority.

```text
second T22 execution: BLOCKED
focused resolver test: exit 1
failure: spec:non-record.doc resolved as unresolved
classification: released-contract gap
```

T27 froze the stale-call-only repair contract.

```yaml
result: PASS

t21:
  task: DRMCP-TASK-MCP-009-21
  status: done
  changed_file: drmcp/src/internal/designrecordsmcp/tools_call_test.go
  config_fixture:
    Root: "."
    RecordsRoots:
      - AppNamespace: "drmcp"
        RecordsRoot: "drmcp/records"
        NamespacePrefix: "DRMCP-"
  focused_test:
    command: go test ./drmcp/src/internal/designrecordsmcp -run '^TestToolsCallToolErrors$' -count=1
    exit_status: 0
  whitespace: PASS
  staged_changes: none

t22:
  task: DRMCP-TASK-MCP-009-22
  status: done
  changed_file: drmcp/src/internal/designrecords/resolve_reference_test.go
  final_diff:
    - ListRecords assertion-block deletion
    - GetRecord assertion-block deletion
  rollback_complete: true
  compile_only:
    command: go test ./drmcp/src/internal/designrecords -run '^$' -count=1
    exit_status: 0
  stale_call_grep:
    exit_status: 1
  resolver_semantic_acceptance_claim: none
  resolver_behavior_owner: DRMCP-TASK-MCP-009-06
  protected_paths_unchanged:
    - drmcp/src/internal/designrecords/resolver.go
    - drmcp/src/internal/designrecords/validation_test.go
  whitespace: PASS
  staged_changes: none

synchronization:
  implementation_performed: false
  commands_rerun: false
  independent_review_performed: false
  aggregate_verification_performed: false
  stage_or_commit: false

next_gate:
  task: DRMCP-TASK-MCP-009-15
  readiness: READY
```

```text
T21 focused repair Evidence: persisted
T22 compile-only repair Evidence: persisted
T22 resolver behavior acceptance: not claimed
T21 status: done
T22 status: done
T23 status: done
T15 aggregate verification: READY
```

T23 is the sole pre-T15 repair Evidence and lifecycle synchronization owner.
T15 has not run or passed; it is only READY through the completed T23 dependency.
