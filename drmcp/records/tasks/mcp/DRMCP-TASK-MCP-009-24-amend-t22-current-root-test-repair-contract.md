# DRMCP-TASK-MCP-009-24: Amend T22 current-root test repair contract

- **id**: DRMCP-TASK-MCP-009-24
- **status**: done
- **date**: 2026-06-30
- **work_item**: DRMCP-WORK-MCP-009
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 0.5d
- **depends_on**:
  - DRMCP-TASK-MCP-009-20
- **outputs**:
  - DRMCP-WORK-MCP-009
  - DRMCP-TASK-MCP-009-05
  - DRMCP-TASK-MCP-009-15
  - DRMCP-TASK-MCP-009-16
  - DRMCP-TASK-MCP-009-21
  - DRMCP-TASK-MCP-009-22
  - DRMCP-TASK-MCP-009-23
  - DRMCP-TASK-MCP-009-24
  - DRMCP-TASK-MCP-009-25
  - DRMCP-TASK-MCP-009-26

## Goal

Persist the second post-release repair amendment after the first T21 and T22 executor results.

Keep T21 retry ownership separate from the bounded T22 current-root test migration.

## Work

### Result classification

Classify the supplied executor results as follows.

| Task | accepted observation | disposition |
|---|---|---|
| T21 | The executor reported `PASS`, but the reported Config replacement is absent from the worktree. `TestToolsCallToolErrors` still uses `designrecords.Config{Root: "."}`. | Reject the result. Keep T21 released and `not_started`. Re-run the existing T21 contract. |
| T22 | The two retired assertion blocks are absent. The stale-call grep passed. The focused test fails before resolver assertions because shared `buildTestIndex` calls `NewConfig(root, "records")`. | Keep the scoped deletion. Mark T22 blocked pending an amended test-local current-root contract. |

T21 is an executor-result persistence failure, not a graph-contract defect.
T22 is a released-contract gap because its original contract preserved stale test initialization.

### T22 amended implementation contract

Keep the T22 writer path unchanged:

```text
drmcp/src/internal/designrecords/resolve_reference_test.go
```

Inside `TestResolveReferenceUsesSemanticRefsFromNonRecordSpec`, T22 correction owns only these additional changes:

1. Change the fixture path from:

```text
records/spec/non-record.md
```

To:

```text
drmcp/records/spec/non-record.md
```

2. Replace only this setup call:

```go
idx := buildTestIndex(t, root)
```

With test-local current-root setup using:

```go
cfg, err := NewConfig(root, "drmcp/records")
```

Then build the index with:

```go
idx, err := BuildIndex(context.Background(), cfg)
```

Both errors must use existing test-style `t.Fatalf` handling.

3. Change only the expected document target path from:

```text
records/spec/non-record.md
```

To:

```text
drmcp/records/spec/non-record.md
```

Preserve:

- the test function name;
- the semantic refs and section mapping;
- the document target-type assertion;
- the section target-type and section-name assertions;
- the prior T22 deletion of the retired `ListRecords` and `GetRecord` assertion blocks;
- `resolver.go` and resolver semantics.

Do not change:

- shared `buildTestIndex`;
- `validation_test.go`;
- another resolver test;
- production source;
- another fixture;
- T06 production ownership.

### Execution graph amendment

```text
T20 accepted repair release
  -> T21 reported result rejected; existing released T21 must be re-run
  -> T22 scoped stale-call deletion complete but focused test BLOCKED
  -> T24 exact T22 current-root contract amendment
  -> T25 independent amendment review
  -> T26 T22 correction re-release
  -> T22 bounded correction execution
  -> T23 valid T21 and T22 Evidence synchronization
  -> T15 aggregate verification
  -> T16 independent finding closure
  -> T17 correction closure synchronization
```

T21 retry may run while T25 reviews the T22 amendment.
T22 correction may start only after T25 `PASS` and T26 release synchronization.
T23 remains blocked until valid T21 and T22 results both exist.

### Model routing

| Task | model route | reason |
|---|---|---|
| T24 | ChatGPT coordinator | Persistent graph and exact test-contract amendment. |
| T25 | Claude Code Opus | Independent ownership and executability review. |
| T26 | Claude Code Haiku | Mechanical acceptance and T22 re-release. |
| T21 retry | Claude Code Haiku | Existing exact Config literal replacement. |
| T22 correction | Claude Code Haiku | One named test, one current-root setup, and one expected path update. |

## Done condition

- The T21 reported `PASS` is rejected because no matching worktree change exists.
- T21 remains released and eligible for a clean retry under its existing contract.
- T22 records the completed stale-call deletion and the exact setup blocker.
- T22 owns a complete test-local current-root migration contract.
- Shared `buildTestIndex`, `validation_test.go`, and `resolver.go` remain protected.
- T25 and T26 exist with exact review and release ownership.
- T15 and T16 recognize the expanded exact T22 boundary.
- T23 remains blocked until valid T21 and T22 results exist.
- No production source or Go test changes during graph authoring.

## Verification

- Confirm the reported T21 Config replacement is absent from `TestToolsCallToolErrors`.
- Confirm the T22 diff removes only the two retired assertion blocks so far.
- Confirm the focused T22 failure names `NewConfig(root, "records")` through shared `buildTestIndex`.
- Confirm the amended T22 contract changes only one named test in one file.
- Confirm T25 depends on T24 and T26 depends on T25.
- Confirm T22 depends on T26 before correction execution.
- Confirm T23 still depends on T21 and T22.
- Inspect only amended Design Records with scoped Git and whitespace checks.

## Evidence

- T21 report claimed a Config containing `RecordsRoots`, but the current worktree still contains the stale single-field Config literal in `TestToolsCallToolErrors`.
- The reported T21 result is not accepted as implementation evidence.
- T22 removed the exact retired `ListRecords` and `GetRecord` assertion blocks.
- T22 stale-call absence check returned exit code `1`.
- T22 focused test failed with `records_root "records" does not match required shape <app_namespace>/records`.
- The failure occurs before resolver assertions and is caused by shared stale test initialization.
- The minimum bounded correction uses test-local `drmcp/records` setup and does not change the shared helper.
- Filesystem authoring was used because DRMCP authoring transactions were unavailable.
- No production source, Go test, fixture, schema, manifest, ADR, Requirement, or Specification changed during T24.

Release history:

```text
T25 independent review: PASS
blocking findings: none
major findings: none
minor findings: none
T22 amended contract: accepted
T26 release synchronization: done
```
