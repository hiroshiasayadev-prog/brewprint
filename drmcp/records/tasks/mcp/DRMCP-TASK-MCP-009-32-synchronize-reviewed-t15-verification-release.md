# DRMCP-TASK-MCP-009-32: Synchronize reviewed T15 verification release

- **id**: DRMCP-TASK-MCP-009-32
- **status**: not_started
- **date**: 2026-06-30
- **work_item**: DRMCP-WORK-MCP-009
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 0.5d
- **depends_on**:
  - DRMCP-TASK-MCP-009-31
- **outputs**:
  - DRMCP-WORK-MCP-009
  - DRMCP-TASK-MCP-009-15
  - DRMCP-TASK-MCP-009-30
  - DRMCP-TASK-MCP-009-31
  - DRMCP-TASK-MCP-009-32
  - DRMCP-TASK-MCP-009-33
  - DRMCP-TASK-MCP-009-34
  - DRMCP-TASK-MCP-009-35
  - DRMCP-TASK-MCP-009-36

## Goal

Synchronize independent acceptance of the amended T15 graph.

Release only the four read-only verification leaves.

## Work

### Start gate

Require T31 to record:

- verdict `PASS`;
- no blocking, major, or minor finding;
- graph-defect classification accepted;
- exact owner maps accepted;
- writer ownership accepted;
- dependency graph accepted;
- T33 through T36 executor readiness accepted.

### Synchronize

- Persist the T31 verdict and finding disposition.
- Set T31 to `done` after its Evidence is complete.
- Release T33, T34, T35, and T36.
- Keep each released leaf `not_started` until its own execution begins.
- Record that T33, T34, and T35 may run in parallel.
- Record that T36 remains execution-gated by the complete T34 report while lifecycle synchronization remains with T15.
- Keep T15 `not_started` and blocked by all four leaf reports.
- Keep T16 and T17 `not_started`.
- Set T32 to `done` after release synchronization is verified.

### Allowed changes

- W009;
- T15;
- T30 through T36.

### Prohibited operations

- Do not run T33 through T36.
- Do not run T15.
- Do not start T16 or T17.
- Do not correct review findings.
- Do not modify production source, Go tests, fixtures, schemas, manifests, ADRs, Requirements, or Specifications.
- Do not stage or commit.

### Output

1. Result.
2. Accepted review.
3. Released Tasks.
4. Parallelism and the T36 report gate.
5. Deferred Tasks and blockers.
6. T15 readiness.

## Done condition

- T31 is `done` with `PASS` and no material finding.
- T33 through T36 are explicitly released.
- T36 remains execution-gated until the complete T34 report exists.
- T15 remains unstarted until all four reports exist.
- T16 and T17 remain unstarted.
- No implementation or verification command runs.

## Verification

- Confirm T31 Evidence satisfies the release gate.
- Confirm T33 through T36 depend on T32 and T36 has a complete-T34-report start gate without a lifecycle dependency cycle.
- Confirm T15 depends on all four leaves.
- Confirm no other Task is released.
- Confirm only allowed Design Records changed.
- Run scoped Git and whitespace inspection on allowed records.

## Evidence

Release synchronization pending T31 `PASS`.
