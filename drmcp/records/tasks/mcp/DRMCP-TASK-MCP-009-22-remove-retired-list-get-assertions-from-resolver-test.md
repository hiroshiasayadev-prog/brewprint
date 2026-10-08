# DRMCP-TASK-MCP-009-22: Remove retired list and get assertions from resolver test

- **id**: DRMCP-TASK-MCP-009-22
- **status**: done
- **date**: 2026-06-30
- **work_item**: DRMCP-WORK-MCP-009
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 0.5d
- **depends_on**:
  - DRMCP-TASK-MCP-009-20
  - DRMCP-TASK-MCP-009-26
  - DRMCP-TASK-MCP-009-29
- **outputs**:
  - drmcp/src/internal/designrecords/resolve_reference_test.go

## Goal

Remove stale assertions that call retired package list and single-get APIs from one resolver test.

Keep all surviving resolver assertions unchanged and under T06 ownership.

## Work

### Read

- This Task.
- `drmcp/src/internal/designrecords/resolve_reference_test.go`.

### Change

Only change:

- `drmcp/src/internal/designrecords/resolve_reference_test.go`.

This Task has temporary test-only ownership of that path for stale T05 API-call removal.
T06 retains ownership of `resolver.go`, resolver behavior, and every surviving resolver assertion.

### Implement

The first execution completed these deletions inside `TestResolveReferenceUsesSemanticRefsFromNonRecordSpec`:

- the `ListRecords` call and its zero-record assertion;
- the `GetRecord` call and its expected-failure assertion.

The second execution followed T24 through T26 and added a current-root migration. Focused execution proved that migration conflicts with the accepted current resolver contract.

After T28 `PASS` and T29 re-release, complete only this rollback:

1. Change the fixture path from `drmcp/records/spec/non-record.md` back to `records/spec/non-record.md`.
2. Replace the test-local `NewConfig` and `BuildIndex` setup with `idx := buildTestIndex(t, root)`.
3. Change the expected document target path from `drmcp/records/spec/non-record.md` back to `records/spec/non-record.md`.

Preserve exactly:

- the test function name;
- the fixture front matter and Markdown body;
- the semantic refs and section mapping;
- document semantic-ref resolution;
- the document target-type assertion;
- section semantic-ref resolution;
- the section target-type and section-name assertions;
- the prior deletion of the retired `ListRecords` and `GetRecord` assertion blocks.

Do not replace the removed assertions with current list or get calls.
Do not execute or reinterpret the surviving resolver assertions.
Do not change shared `buildTestIndex`, `validation_test.go`, or `resolver.go`.

### Test

Run from repository root:

```powershell
gofmt -w drmcp/src/internal/designrecords/resolve_reference_test.go

go test ./drmcp/src/internal/designrecords -run '^$' -count=1
```

The compile-only command must exit `0`.
It proves that the package tests no longer reference retired T05 APIs.
It does not assert resolver semantic behavior.

Run this scoped absence check:

```powershell
git grep -n -E '\b(ListRecords|GetRecord)\b' -- drmcp/src/internal/designrecords/resolve_reference_test.go
```

Exit semantics for the grep:

- `1`: no stale call remains; `PASS`;
- `0`: stale call remains; `FAIL`;
- `2` or greater: command failure; `BLOCKED`.

T23 persists this compile-only result. T15 owns the T05-focused combined package commands and aggregate correction verification. T06 owns resolver behavior verification.

### Stop

Stop as `BLOCKED` without wider exploration when:

- another file must change;
- another test function must change;
- `resolver.go` must change;
- shared `buildTestIndex` or `validation_test.go` must change;
- the test function name, fixture body, semantic refs, section mapping, or surviving resolver assertions must change;
- a replacement list or get assertion appears necessary;
- the compile-only command fails for a reason outside the exact stale assertion removal and migration rollback;
- the exact commands cannot run.

### Prohibited operations

- Do not modify `resolver.go`.
- Do not change resolver semantics.
- Do not modify shared `buildTestIndex` or `validation_test.go`.
- Do not add current list or get coverage to this test.
- Do not execute the named resolver test.
- Do not modify another test, production file, fixture, or Design Record.
- Do not perform unrelated cleanup or independent review.
- Do not stage or commit.
- Do not perform repository-wide traversal.

### Output

1. Result.
2. Changed file.
3. Preserved stale-call deletions.
4. Rolled-back current-root migration.
5. Compile-only command and exit status.
6. Absence-check result.
7. Protected-path state.
8. Blocker.
9. T23 readiness.

## Done condition

- `TestResolveReferenceUsesSemanticRefsFromNonRecordSpec` contains no `ListRecords` or `GetRecord` call.
- The fixture path is restored to `records/spec/non-record.md`.
- The named test uses its original shared `buildTestIndex(t, root)` setup.
- The expected document target path is restored to `records/spec/non-record.md`.
- The surviving document and section resolver assertions are unchanged and remain T06-owned.
- Shared `buildTestIndex`, `validation_test.go`, `resolver.go`, and all other files remain unchanged.
- The compile-only command exits `0`.
- The scoped absence check passes.
- T22 is ready for T23 Evidence synchronization after T21 also completes.

## Verification

- Use the exact commands in `## Work`.
- Inspect the scoped diff and whitespace for `resolve_reference_test.go` only.
- Confirm the complete T22 diff contains only the two stale assertion-block deletions.
- Confirm the second-amendment current-root migration is fully rolled back.
- Confirm `resolver.go` and `validation_test.go` have no scoped change.
- Do not infer repository-wide cleanliness.

## Evidence

First release: `DRMCP-TASK-MCP-009-20`.

First execution result: `BLOCKED` after the exact stale-call deletion completed.

```text
changed file: drmcp/src/internal/designrecords/resolve_reference_test.go
removed blocks: ListRecords zero-record assertion and GetRecord expected-failure assertion
stale-call absence check: PASS, grep exit code 1
focused test: FAIL, exit code 1
failure: NewConfig: records_root "records" does not match required shape <app_namespace>/records
failure location: shared buildTestIndex setup before resolver assertions
stage or commit: none
```

Second amendment history:

```yaml
authoring:
  task: DRMCP-TASK-MCP-009-24
  status: done
independent_review:
  task: DRMCP-TASK-MCP-009-25
  status: done
  verdict: PASS
release_synchronization:
  task: DRMCP-TASK-MCP-009-26
  status: done
```

Second execution result: `BLOCKED`.

```text
gofmt: PASS, exit code 0
focused resolver test: FAIL, exit code 1
failure: spec:non-record.doc resolved as unresolved
stale-call grep: not run after mandatory stop
scoped diff: PASS
whitespace: PASS
protected resolver.go: unchanged
protected validation_test.go: unchanged
```

The second execution followed the released contract. The focused failure exposed that the T24 through T26 contract conflicted with accepted current spec and resolver authority.

Third amendment:

```yaml
authoring:
  task: DRMCP-TASK-MCP-009-27
  status: done

independent_review:
  task: DRMCP-TASK-MCP-009-28
  status: done
  verdict: PASS
  blocking_findings: none
  major_findings: none
  minor_findings: none
  advisories: none

release_synchronization:
  task: DRMCP-TASK-MCP-009-29
  status: done

correction_release_state: re_released_stale_call_only
correction_execution_eligible: true
lifecycle_status: blocked
```

### Final execution Evidence

```yaml
result: PASS
changed_file: drmcp/src/internal/designrecords/resolve_reference_test.go
stale_assertion_blocks_removed:
  - ListRecords zero-record assertion
  - GetRecord expected-failure assertion
rollback:
  fixture_path_restored: records/spec/non-record.md
  index_setup_restored: idx := buildTestIndex(t, root)
  expected_path_restored: records/spec/non-record.md
verification:
  gofmt_exit: 0
  compile_only_test_exit: 0
  stale_call_grep_exit: 1
final_diff:
  - ListRecords assertion-block deletion
  - GetRecord assertion-block deletion
protected_paths:
  resolver.go: unchanged
  validation_test.go: unchanged
  shared_buildTestIndex: unchanged
surviving_resolver_assertions: unchanged
resolver_semantic_acceptance_claim: none
resolver_behavior_owner: DRMCP-TASK-MCP-009-06
named_resolver_test_executed: false
scoped_diff: PASS
whitespace: PASS
staged_changes: none
repository_wide_clean: not inspected and not inferred
blocker: none
t23_readiness: READY
```

T22 is complete. Its focused compile-only Evidence is persisted by `DRMCP-TASK-MCP-009-23`.
No resolver behavioral acceptance is claimed, and the named resolver test was not executed.
