# DRMCP-TASK-MCP-009-12: Synchronize reviewed T05 correction release

- **id**: DRMCP-TASK-MCP-009-12
- **status**: done
- **date**: 2026-06-29
- **work_item**: DRMCP-WORK-MCP-009
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 0.5d
- **depends_on**:
  - DRMCP-TASK-MCP-009-11
- **outputs**:
  - DRMCP-WORK-MCP-009
  - DRMCP-TASK-MCP-009-05
  - DRMCP-TASK-MCP-009-10
  - DRMCP-TASK-MCP-009-11
  - DRMCP-TASK-MCP-009-12
  - DRMCP-TASK-MCP-009-13
  - DRMCP-TASK-MCP-009-14

## Goal

Record accepted graph-review evidence and make T13 and T14 start eligibility explicit.

## Work

### Release gate

Do not execute this Task until an independent T11 re-review records all of the following:

- verdict: `PASS`;
- `F-BLOCK-01`: `CLOSED`;
- `F-MAJ-01`: `CLOSED`;
- `F-MIN-01`: `CLOSED`;
- no new blocking, major, or minor finding;
- T13 executor contract: ready;
- T14 executor contract: ready;
- T15 mechanically executable: yes;
- T15 classification is total and ordered;
- T15 non-test and unclassifiable blockers have precedence over allowlist routing;
- an allowlisted-test panic results in `BLOCKED`;
- an allowlisted but unclassifiable test failure results in `BLOCKED`;
- `DEFERRED_TO_T08` is possible only when every exact failing test is allowlisted and every unconditional blocker is false;
- T15 exact diagnostic allowlist and `DEFERRED_TO_T08` route: correctly defined;
- T16 mapping gate accepted: yes;
- T17 complete replacement synchronization accepted: yes.

### Synchronize accepted release

- Record the accepted T11 re-review verdict and exact finding disposition in T11 Evidence.
- Record T13 and T14 as the only released executor leaves.
- Set T11 to `done` only after its accepted re-review is persisted.
- Set T12 to `done` only after release synchronization is verified.
- Keep executor statuses unchanged until each executor starts.
- Keep T15 deferred until T13 and T14 complete.
- Keep T16 deferred until T15 records overall `PASS`; diagnostic `DEFERRED_TO_T08` may accompany that PASS.
- Keep T17 deferred until T16 returns `PASS` and `mapping manifest accepted: yes`.

Allowed changes:

- `DRMCP-WORK-MCP-009-current-format-read-implementation.md`.
- `DRMCP-TASK-MCP-009-05-implement-current-list-and-exact-retrieval.md`.
- `DRMCP-TASK-MCP-009-10-freeze-t05-correction-execution-graph.md`.
- `DRMCP-TASK-MCP-009-11-review-t05-correction-execution-graph.md`.
- `DRMCP-TASK-MCP-009-12-synchronize-reviewed-t05-correction-release.md`.
- `DRMCP-TASK-MCP-009-13-retire-unused-adapter-helper.md`.
- `DRMCP-TASK-MCP-009-14-retire-legacy-list-get-and-correct-compact-projection.md`.

Prohibited operations:

- Do not redesign the graph or correct findings.
- Do not modify T15, T16, T17, T08, or T09.
- Do not modify production source, tests, fixtures, schemas, or manifests.
- Do not implement production work.
- Do not stage or commit.

## Done condition

- T11 independent re-review records `PASS`.
- F-BLOCK-01, F-MAJ-01, and F-MIN-01 are all recorded as `CLOSED`.
- No new blocking, major, or minor finding exists.
- T13 and T14 readiness is accepted.
- T15 is mechanically executable with total ordered classification, unconditional blocker precedence, and exact allowlist routing only after every unconditional blocker is false.
- An allowlisted-test panic and an allowlisted but unclassifiable test failure are accepted only as `BLOCKED`.
- T16 complete replacement mapping gate is accepted.
- T17 complete replacement synchronization contract is accepted.
- T13 and T14 are the only released executor leaves.
- Parallel eligibility and deferred Tasks are explicit.
- T11 and T12 are `done` with accepted evidence.
- No implementation or executor lifecycle state changes occur during synchronization.

## Verification

- Confirm T11 verdict is `PASS`.
- Confirm F-BLOCK-01, F-MAJ-01, and F-MIN-01 are `CLOSED`.
- Confirm no new blocking, major, or minor finding exists.
- Confirm T13 and T14 writable paths are disjoint and both executor contracts are ready.
- Confirm T15 does not require a full `designrecords` package PASS.
- Confirm T15 classification is total and ordered.
- Confirm unconditional compile, setup, initialization, panic, unknown, and mechanically unclassifiable blockers are evaluated before allowlist comparison.
- Confirm an allowlisted-test panic and an allowlisted but unclassifiable test failure are `BLOCKED`.
- Confirm T15 uses only exact frozen test-name allowlist membership for diagnostic deferral after every unconditional blocker is false.
- Confirm T16 accepts or rejects the complete replacement T15 mapping manifest independently.
- Confirm T17 is mechanical, requires an accepted complete manifest, and performs no merge or supplementation.
- Confirm T15, T16, and T17 remain dependency-gated.
- Inspect Git and whitespace for the writable W009, T05, and T10 through T14 Design Records only.

## Evidence

### Prior review history

- The first T11 review returned `NEEDS REVISION` with F-BLOCK-01, F-MAJ-01, and F-MIN-01.
- The second T11 review returned `NEEDS REVISION`; second correction authoring was `CORRECTED_PENDING_REVIEW`.
- The third T11 review returned `NEEDS REVISION`; F-BLOCK-01 remained `OPEN`, while F-MAJ-01 and F-MIN-01 were `CLOSED`.
- Third correction authoring was `CORRECTED_PENDING_REVIEW`.

### Accepted release synchronization

```yaml
accepted_review_task: DRMCP-TASK-MCP-009-11
accepted_verdict: PASS

accepted_finding_disposition:
  F-BLOCK-01: CLOSED
  F-MAJ-01: CLOSED
  F-MIN-01: CLOSED

new_findings:
  blocking: none
  major: none
  minor: none

released_executors:
  - DRMCP-TASK-MCP-009-13
  - DRMCP-TASK-MCP-009-14

release_mode: parallel

writer_ownership:
  T13:
    - drmcp/src/internal/designrecordsmcp/tools.go

  T14:
    - drmcp/src/internal/designrecords/tools.go
    - drmcp/src/internal/designrecords/list_records_test.go
    - drmcp/src/internal/designrecords/get_records_test.go

post_release_gate:
  requirement: T13 and T14 must both complete before DRMCP-TASK-MCP-009-15 may start

deferred_tasks:
  T15: dependency-gated
  T16: dependency-gated
  T17: dependency-gated
  T08: dependency-gated by T05, T06, T07
  T09: final W009 closure gate
```

- T13 and T14 writer paths are non-overlapping.
- T13 and T14 remain `not_started` until their own executor sessions begin.
- T15 remains blocked until both T13 and T14 complete.
- T16 remains blocked until T15 records overall `PASS`.
- T17 remains blocked until T16 returns `PASS` and accepts the complete replacement mapping manifest.
- Production implementation was not started during release synchronization.
