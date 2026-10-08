# DRMCP-TASK-MCP-009-28: Review T22 stale-call-only repair contract

- **id**: DRMCP-TASK-MCP-009-28
- **status**: done
- **date**: 2026-06-30
- **work_item**: DRMCP-WORK-MCP-009
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 0.5d
- **depends_on**:
  - DRMCP-TASK-MCP-009-27
- **outputs**:
  - DRMCP-TASK-MCP-009-28

## Goal

Independently judge whether the third T22 amendment restores the accepted current resolver boundary.

Confirm that T22 can complete as a stale-call-only repair without executing or redefining T06-owned resolver behavior.

## Work

### Read

Read only:

```text
drmcp/records/work-items/mcp/DRMCP-WORK-MCP-003-current-discovery-and-active-index-contract-realignment.md
drmcp/records/work-items/mcp/DRMCP-WORK-MCP-005-resolver-and-configured-legacy-fallback-contract-realignment.md
drmcp/records/work-items/mcp/DRMCP-WORK-MCP-009-current-format-read-implementation.md
drmcp/records/tasks/mcp/DRMCP-TASK-MCP-009-05-implement-current-list-and-exact-retrieval.md
drmcp/records/tasks/mcp/DRMCP-TASK-MCP-009-06-implement-current-reference-resolution.md
drmcp/records/tasks/mcp/DRMCP-TASK-MCP-009-15-verify-t05-correction-integration.md
drmcp/records/tasks/mcp/DRMCP-TASK-MCP-009-16-review-t05-finding-closure.md
drmcp/records/tasks/mcp/DRMCP-TASK-MCP-009-22-remove-retired-list-get-assertions-from-resolver-test.md
drmcp/records/tasks/mcp/DRMCP-TASK-MCP-009-23-synchronize-t05-repair-execution-evidence.md
drmcp/records/tasks/mcp/DRMCP-TASK-MCP-009-24-amend-t22-current-root-test-repair-contract.md
drmcp/records/tasks/mcp/DRMCP-TASK-MCP-009-25-review-t22-current-root-test-repair-contract.md
drmcp/records/tasks/mcp/DRMCP-TASK-MCP-009-26-synchronize-reviewed-t22-current-root-repair-release.md
drmcp/records/tasks/mcp/DRMCP-TASK-MCP-009-27-amend-t22-stale-call-only-repair-contract.md
drmcp/records/tasks/mcp/DRMCP-TASK-MCP-009-28-review-t22-stale-call-only-repair-contract.md
drmcp/records/tasks/mcp/DRMCP-TASK-MCP-009-29-synchronize-reviewed-t22-stale-call-only-repair-release.md
drmcp/src/internal/designrecords/resolve_reference_test.go
drmcp/src/internal/designrecords/parser.go
drmcp/src/internal/designrecords/resolver.go
```

Use `parser.go` only to inspect current spec front-matter rejection, path-derived identity, and semantic-ref population.

Use `resolver.go` only to inspect current semantic lookup behavior needed to understand the recorded focused failure.

Do not review unrelated parser, resolver, or test behavior.

### Review

Confirm:

- the second T22 execution followed the T24-T26 contract;
- the focused test reached resolver execution and returned `unresolved`;
- current spec discovery rejects YAML front matter;
- current spec identity is path-derived;
- accepted current resolution prohibits front-matter semantic aliases and section aliases;
- the T24-T26 current-root migration could not preserve the stale test's original semantic behavior;
- the failure is a released-contract gap and not an executor mistake;
- T25 omitted controlling W003, W005, and T06 authority from its judgment;
- rolling back only the second-amendment migration restores the pre-existing T06-owned test body;
- keeping the two stale T05 API assertion deletions is correct;
- compile-only package verification proves retired API references are absent without asserting resolver behavior;
- the remaining resolver assertions stay T06-owned;
- T23 requires only compile-only and stale-call absence Evidence for T22;
- T15 and T16 no longer treat current-root migration as a T05 acceptance condition;
- T21 retry remains independent;
- the corrected T22 leaf is executor-ready;
- T29 release synchronization can release only T22.

### Prohibited operations

- Do not modify files.
- Do not correct T22.
- Do not execute T21 or T22.
- Do not update lifecycle state.
- Do not run Go tests, formatter, or generator.
- Do not perform repository-wide traversal.
- Do not stage or commit.

### Output

1. Verdict: `PASS`, `NEEDS REVISION`, or `BLOCKED`.
2. Reviewed files.
3. Second T22 failure classification.
4. T25 review-scope assessment.
5. Blocking findings.
6. Major findings.
7. Minor findings.
8. Advisories.
9. Rollback-contract assessment.
10. Compile-only verification assessment.
11. Writer and dependency assessment.
12. T23, T15, T16, and T17 ownership assessment.
13. Executor readiness.
14. T29 release readiness.

## Done condition

- The third amendment receives an independent verdict.
- Every material finding has a stable ID and severity.
- The review confirms whether T22 is limited to stale-call-only repair.
- The review confirms whether resolver semantics and assertion migration remain T06-owned.
- The review confirms whether compile-only verification is sufficient for T22.
- No file is changed.

## Verification

- Inspect the complete scoped T22 diff directly.
- Inspect the recorded focused failure.
- Inspect current spec front-matter rejection and path-derived identity.
- Inspect W005 and T06 accepted removal of semantic and section aliases.
- Confirm T22 depends on T29 and T23 still depends on T21 and T22.
- Confirm no production or Go test change is present in the T27 authoring diff.

## Evidence

```yaml
review_task: DRMCP-TASK-MCP-009-28
verdict: PASS

findings:
  blocking: none
  major: none
  minor: none
  advisories: none

second_t22_failure:
  classification: released_contract_gap
  executor_mistake: false

t25_review_scope:
  omission_confirmed: true

rollback_contract:
  accepted: true
  final_t22_diff:
    - remove ListRecords call and zero-record assertion
    - remove GetRecord call and expected-failure assertion
  rollback:
    - restore fixture path records/spec/non-record.md
    - restore idx := buildTestIndex(t, root)
    - restore expected path records/spec/non-record.md

verification_contract:
  compile_only: accepted
  command: go test ./drmcp/src/internal/designrecords -run '^$' -count=1
  expected_exit_code: 0
  stale_call_grep_expected_exit_code: 1
  resolver_behavior_claim: none

ownership:
  remaining_resolver_assertions: DRMCP-TASK-MCP-009-06
  t22_writer_unique: true
  t21_t22_writer_overlap: false
  downstream_ownership_preserved: true
  dependency_graph_acyclic: true

readiness:
  t22_executor: accepted
  t29_release: READY

review_execution:
  read_only: true
  tests_run: false
  files_changed: none
```

```text
verdict: PASS
blocking findings: none
major findings: none
minor findings: none
advisories: none
second failure: released-contract gap
T25 scope omission: confirmed
rollback contract: accepted
compile-only verification: accepted
remaining resolver assertions: T06-owned
T22 executor readiness: accepted
T29 release readiness: READY
tests run: false
files changed by review: none
```

T28 does not claim that T22 implementation or verification completed.
T28 does not replace T16 finding closure or T09 final W009 review.
