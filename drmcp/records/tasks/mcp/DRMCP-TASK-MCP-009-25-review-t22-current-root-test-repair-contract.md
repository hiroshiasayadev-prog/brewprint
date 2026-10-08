# DRMCP-TASK-MCP-009-25: Review T22 current-root test repair contract

- **id**: DRMCP-TASK-MCP-009-25
- **status**: done
- **date**: 2026-06-30
- **work_item**: DRMCP-WORK-MCP-009
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 0.5d
- **depends_on**:
  - DRMCP-TASK-MCP-009-24
- **outputs**:
  - DRMCP-TASK-MCP-009-25

## Goal

Independently judge whether the amended T22 contract can make the focused resolver test pass without reopening T06 or changing shared test infrastructure.

## Work

### Read

Read only:

```text
drmcp/records/work-items/mcp/DRMCP-WORK-MCP-009-current-format-read-implementation.md
drmcp/records/tasks/mcp/DRMCP-TASK-MCP-009-05-implement-current-list-and-exact-retrieval.md
drmcp/records/tasks/mcp/DRMCP-TASK-MCP-009-21-migrate-adapter-error-test-current-root-fixture.md
drmcp/records/tasks/mcp/DRMCP-TASK-MCP-009-22-remove-retired-list-get-assertions-from-resolver-test.md
drmcp/records/tasks/mcp/DRMCP-TASK-MCP-009-23-synchronize-t05-repair-execution-evidence.md
drmcp/records/tasks/mcp/DRMCP-TASK-MCP-009-24-amend-t22-current-root-test-repair-contract.md
drmcp/records/tasks/mcp/DRMCP-TASK-MCP-009-25-review-t22-current-root-test-repair-contract.md
drmcp/records/tasks/mcp/DRMCP-TASK-MCP-009-26-synchronize-reviewed-t22-current-root-repair-release.md
drmcp/records/tasks/mcp/DRMCP-TASK-MCP-009-15-verify-t05-correction-integration.md
drmcp/records/tasks/mcp/DRMCP-TASK-MCP-009-16-review-t05-finding-closure.md
drmcp/records/tasks/mcp/DRMCP-TASK-MCP-009-17-synchronize-t05-correction-closure.md
drmcp/src/internal/designrecords/resolve_reference_test.go
drmcp/src/internal/designrecords/config.go
drmcp/src/internal/designrecords/index.go
drmcp/src/internal/designrecords/validation_test.go
```

Use `validation_test.go` only to inspect the shared `buildTestIndex` definition.
Do not review unrelated validation tests.

### Review

Confirm:

- the original T22 stale-call deletion is complete and correct;
- the focused failure occurs before resolver assertions;
- shared `buildTestIndex` uses stale `NewConfig(root, "records")` setup;
- changing shared `buildTestIndex` would reopen a wider test-migration lane;
- test-local `NewConfig(root, "drmcp/records")` is consistent with the current Config contract;
- the amended fixture path matches that current root;
- the expected document target path changes mechanically with the fixture path;
- the section semantic-ref assertion remains unchanged;
- no `resolver.go` or resolver semantic change is required;
- no shared helper, validation test, or production writer overlap exists;
- the amended T22 leaf is Haiku-ready;
- T23, T15, T16, and T17 ownership remains valid;
- T21 retry remains independent from this amendment.

### Prohibited operations

- Do not modify files.
- Do not correct T22.
- Do not re-run T21.
- Do not release T22.
- Do not update lifecycle state.
- Do not run Go tests, formatter, or generator.
- Do not perform repository-wide traversal.
- Do not stage or commit.

### Output

1. Verdict: `PASS`, `NEEDS REVISION`, or `BLOCKED`.
2. Reviewed files.
3. T22 blocker classification.
4. Blocking findings.
5. Major findings.
6. Minor findings.
7. Advisories.
8. Writer and dependency assessment.
9. Haiku readiness.
10. Downstream ownership assessment.
11. T26 release readiness.

## Done condition

- The amended T22 contract receives an independent verdict.
- Every material finding has a stable ID and severity.
- The review confirms whether the correction is bounded to one named test.
- The review confirms whether shared test infrastructure and T06 remain closed.
- No file is changed.

## Verification

- Inspect the scoped T22 diff directly.
- Inspect the exact `NewConfig` and `BuildIndex` contracts.
- Inspect only the shared `buildTestIndex` function in `validation_test.go`.
- Confirm the amended paths, setup, expected path, and preserved assertions are exact.
- Confirm T22 depends on T26 and T23 still depends on T21 and T22.
- Confirm no production or Go test change is present in the T24 authoring diff.

## Evidence

```yaml
review_task: DRMCP-TASK-MCP-009-25
verdict: PASS

findings:
  blocking: none
  major: none
  minor: none
  advisories: none

t22_blocker_classification: graph-contract gap
executor_mistake: false

amended_contract:
  executable: true
  new_config_root: accepted
  build_index_call: accepted
  fixture_path_migration: accepted
  expected_document_path_migration: accepted
  section_semantic_ref_preserved: true
  section_name_assertion_preserved: true
  shared_build_test_index_change_required: false
  production_change_required: false

ownership:
  t22_writer_unique: true
  t21_writer_disjoint: true
  t06_production_ownership_preserved: true
  dependency_graph_acyclic: true
  t23_ownership_preserved: true
  t15_ownership_preserved: true
  t16_ownership_preserved: true
  t17_ownership_preserved: true
  downstream_ownership_preserved: true

model_routing:
  t22_haiku_readiness: accepted

release_readiness:
  t26: READY

review_execution:
  read_only: true
  go_tests_run: false
```

The amended T22 contract is executable without changing shared `buildTestIndex`, `validation_test.go`, `resolver.go`, or T06 production ownership.
The review found no blocking, major, minor, or advisory finding.
T25 does not replace T16 finding closure or T09 final W009 review.
