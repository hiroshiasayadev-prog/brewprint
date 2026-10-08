# DRMCP-TASK-MCP-009-29: Synchronize reviewed T22 stale-call-only repair release

- **id**: DRMCP-TASK-MCP-009-29
- **status**: done
- **date**: 2026-06-30
- **work_item**: DRMCP-WORK-MCP-009
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 0.5d
- **depends_on**:
  - DRMCP-TASK-MCP-009-28
- **outputs**:
  - DRMCP-WORK-MCP-009
  - DRMCP-TASK-MCP-009-05
  - DRMCP-TASK-MCP-009-15
  - DRMCP-TASK-MCP-009-16
  - DRMCP-TASK-MCP-009-22
  - DRMCP-TASK-MCP-009-23
  - DRMCP-TASK-MCP-009-27
  - DRMCP-TASK-MCP-009-28
  - DRMCP-TASK-MCP-009-29

## Goal

Record independent acceptance of the stale-call-only T22 contract.

Re-release only T22 for rollback of the invalid current-root migration and compile-only verification.

## Work

### Release gate

Do not execute this Task until T28 records all of the following:

- verdict: `PASS`;
- no blocking, major, or minor finding;
- second T22 failure: released-contract gap;
- T25 review-scope omission: confirmed;
- rollback contract: accepted;
- compile-only verification: accepted;
- remaining resolver assertions: T06-owned;
- T22 writer boundary: unique;
- corrected T22 executor readiness: accepted;
- T23, T15, T16, and T17 ownership: preserved.

### Synchronize accepted release

- Record the accepted T28 verdict and finding disposition.
- Set T28 to `done` only after its review Evidence is persisted.
- Record T22 as the only newly re-released correction leaf.
- Keep T21 retry eligibility unchanged.
- Keep T22 status `blocked` until T23 persists a successful compile-only T22 result.
- Record that T22 execution is eligible despite the retained historical blocked status.
- Keep T23 deferred until valid T21 and T22 results are available.
- Keep T15, T16, and T17 dependency-gated.
- Set T29 to `done` only after release synchronization is verified.

Allowed changes:

```text
drmcp/records/work-items/mcp/DRMCP-WORK-MCP-009-current-format-read-implementation.md
drmcp/records/tasks/mcp/DRMCP-TASK-MCP-009-05-implement-current-list-and-exact-retrieval.md
drmcp/records/tasks/mcp/DRMCP-TASK-MCP-009-15-verify-t05-correction-integration.md
drmcp/records/tasks/mcp/DRMCP-TASK-MCP-009-16-review-t05-finding-closure.md
drmcp/records/tasks/mcp/DRMCP-TASK-MCP-009-22-remove-retired-list-get-assertions-from-resolver-test.md
drmcp/records/tasks/mcp/DRMCP-TASK-MCP-009-23-synchronize-t05-repair-execution-evidence.md
drmcp/records/tasks/mcp/DRMCP-TASK-MCP-009-27-amend-t22-stale-call-only-repair-contract.md
drmcp/records/tasks/mcp/DRMCP-TASK-MCP-009-28-review-t22-stale-call-only-repair-contract.md
drmcp/records/tasks/mcp/DRMCP-TASK-MCP-009-29-synchronize-reviewed-t22-stale-call-only-repair-release.md
```

### Prohibited operations

- Do not modify T22 implementation.
- Do not execute T21.
- Do not modify T17, T08, T09, or T10 through T21.
- Do not rewrite T24 through T26 history.
- Do not modify production source, Go tests, fixtures, schemas, or manifests.
- Do not stage or commit.

## Done condition

- T28 records `PASS` with no blocking, major, or minor finding.
- The stale-call-only T22 contract is accepted.
- T22 is the only newly re-released correction leaf.
- T21 retry remains independently eligible.
- T23 remains blocked until both valid executor results exist.
- T28 and T29 are `done` with accepted Evidence.
- No implementation or T22 lifecycle closure occurs.

## Verification

- Confirm the T28 verdict and finding disposition.
- Confirm T22 depends on T29.
- Confirm T22 writer path remains unique.
- Confirm remaining resolver assertions and `resolver.go` remain T06-owned.
- Confirm T23 still depends on T21 and T22.
- Confirm T15, T16, and T17 dependencies remain unchanged.
- Inspect Git and whitespace for only the allowed Design Records.

## Evidence

```yaml
result: PASS

accepted_review:
  task: DRMCP-TASK-MCP-009-28
  verdict: PASS
  blocking_findings: none
  major_findings: none
  minor_findings: none
  advisories: none

released_executor_tasks:
  - DRMCP-TASK-MCP-009-22

t22:
  lifecycle_status: blocked
  correction_release_state: re_released_stale_call_only
  correction_execution_eligible: true

t21:
  retry_eligibility: unchanged

deferred:
  DRMCP-TASK-MCP-009-23:
    blocker:
      - valid T21 result
      - compile-only T22 result
  DRMCP-TASK-MCP-009-15:
    blocker: DRMCP-TASK-MCP-009-23
  DRMCP-TASK-MCP-009-16:
    blocker: DRMCP-TASK-MCP-009-15
  DRMCP-TASK-MCP-009-17:
    blocker: DRMCP-TASK-MCP-009-16

implementation_started: false
tests_run: false
stage_or_commit: false
```

T22 is the only executor Task released by T29.
