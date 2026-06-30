# DRMCP-TASK-MCP-001-13: Amend read-baseline execution graph for rebuild

- **id**: DRMCP-TASK-MCP-001-13
- **status**: done
- **date**: 2026-06-30
- **work_item**: DRMCP-WORK-MCP-001
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 1d
- **depends_on**:
  - DRMCP-TASK-MCP-001-01
  - DRMCP-TASK-MCP-011-07
- **outputs**:
  - DRMCP-TASK-MCP-001-13
  - DRMCP-TASK-MCP-001-14
  - DRMCP-TASK-MCP-001-15
  - DRMCP-TASK-MCP-001-16
  - DRMCP-TASK-MCP-001-17
  - DRMCP-WORK-MCP-001
  - DRMCP-TASK-MCP-001-09
  - DRMCP-TASK-MCP-001-10
  - DRMCP-TASK-MCP-001-11
  - DRMCP-REQ-MCP-001
  - DRMCP-WORK-MCP-012
  - DRMCP-TASK-MCP-012-01
  - DRMCP-WORK-MCP-010
  - DRMCP-WORK-SPEC-001
  - DRMCP-WORK-SPEC-002

## Goal

Amend the W001 execution graph after W011 architecture closure.

The amendment must:

- replace the retired W009 implementation gate with a new W012 replacement-runtime gate;
- preserve W009 as historical retired implementation planning;
- sequence the retained W-SPEC-001 and W-SPEC-002 implementation owners;
- require W010 graph rebaseline against the completed replacement runtime;
- add independent graph review and release synchronization;
- prevent production implementation until both the W001 amendment and W012 execution graph are independently reviewed and released.

This Task performs graph and workflow authoring only.
It does not implement production code.

## Work

### Responsibility

```text
graph amendment and workflow authoring
```

### Accepted authority

Use the following accepted state:

- `DRMCP-WORK-MCP-011` is `done`.
- D-001 through D-009 are `recorded`.
- T06 verdict is `PASS`.
- F-MAJ-01 and F-MAJ-02 are `CLOSED`.
- No unresolved W011 architecture decision remains.
- `spec:drmcp.implementation` is the implementation-architecture authority for the four W011 read-runtime operations.
- W009 is replaced and retired for the rebuild line.
- W010 is retained and blocked pending execution-graph rebaseline.
- W-SPEC-001 and W-SPEC-002 are retained, separate, and `not_started`.

### Future ID allocation

Use these exact IDs:

```text
DRMCP-TASK-MCP-001-14
DRMCP-TASK-MCP-001-15
DRMCP-TASK-MCP-001-16
DRMCP-TASK-MCP-001-17
DRMCP-WORK-MCP-012
DRMCP-TASK-MCP-012-01
```

Do not reserve other IDs.

### Create replacement Work Item W012

Create:

```text
drmcp/records/work-items/mcp/
DRMCP-WORK-MCP-012-rebuild-current-read-runtime-implementation.md
```

Use this metadata:

```text
id: DRMCP-WORK-MCP-012
status: not_started
date: 2026-06-30
source_requirement: DRMCP-REQ-MCP-001
tasks:
  - DRMCP-TASK-MCP-012-01
```

W012 must consume:

- `DRMCP-WORK-MCP-011`;
- `DRMCP-ADR-MCP-002`;
- `DRMCP-ADR-MCP-003`;
- `DRMCP-ADR-MCP-004`;
- `DRMCP-ADR-MCP-005`;
- `DRMCP-ADR-MCP-006`;
- `spec:drmcp.implementation`;
- `DRMCP-WORK-MCP-003`;
- `DRMCP-WORK-MCP-004`;
- `DRMCP-WORK-MCP-005`;
- `DRMCP-WORK-MCP-006`;
- `DRMCP-WORK-MCP-008`.

W012 must own clean implementation of:

- `list_records`;
- `get_records`;
- `resolve_reference`;
- `validate_records`.

W012 must own:

- composition-root wiring for the four operations;
- application use cases;
- request-scoped snapshot orchestration;
- core record, index, resolving, and validation infrastructure;
- MCP, filesystem, and configuration adapters;
- general current-read automated tests;
- current-only integration against W008 fixtures;
- aggregate implementation verification;
- independent implementation review and closure.

W012 must not own:

- authoring-guidance application architecture;
- authoring transaction behavior;
- portable package loading;
- configured legacy archive fallback implementation;
- W-SPEC-001 per-file detector implementation;
- W-SPEC-002 Topics graph implementation;
- PRODUCT semantic rules;
- fixture authoring;
- reuse of W009 production structure, Task graph, writer allocation, or extension seams.

W012 may inspect old code only as non-authoritative evidence.
W012 must not preserve old code merely to reduce implementation effort.

### Create W012 graph-authoring Task

Create:

```text
drmcp/records/tasks/mcp/
DRMCP-TASK-MCP-012-01-author-rebuild-current-read-runtime-execution-graph.md
```

Use this metadata:

```text
id: DRMCP-TASK-MCP-012-01
status: not_started
date: 2026-06-30
work_item: DRMCP-WORK-MCP-012
source_requirement: DRMCP-REQ-MCP-001
depends_on:
  - DRMCP-TASK-MCP-001-15
```

Responsibility:

```text
execution graph authoring and scope freeze
```

W012 T01 must:

- inspect the current source and test inventory inside the exact DRMCP implementation boundary;
- map the accepted W011 package architecture to actual replacement files and symbols;
- define exact executor Tasks;
- define model routing;
- define one writer for every production, test, fixture-consumer, helper, and lifecycle path;
- define predecessor and consumer dependencies;
- define focused verification owners;
- define one aggregate integration and verification owner;
- define independent implementation-review ownership;
- define release-synchronization ownership;
- create the persistent execution graph;
- stop before production implementation.

W012 T01 must apply the execution-hub Task pattern.

No implementation prompt may be issued from W012 until this sequence is complete:

```text
W012 graph authoring
  -> independent graph review PASS
  -> release synchronization
```

### Replace the W009 lifecycle gate

Amend `DRMCP-TASK-MCP-001-09`.

Retain its public ID and current-read lifecycle role.
Replace its selected child Work Item with `DRMCP-WORK-MCP-012`.
Update its outputs accordingly.

Set T09 to `blocked` with this exact blocker:

```text
Awaiting T15 reviewed graph release and the start of the W012
execution-graph authoring sequence.
```

T09 must state:

- W009 is replaced and retired for the rebuild line;
- W009 is not a completion gate;
- W012 is the sole replacement current-read implementation gate;
- T09 performs lifecycle tracking only;
- T09 does not implement production code;
- T09 changes to `in_progress` only when the released W012 workflow begins;
- T09 reaches `done` only when W012 is reviewed and `done`.

Do not alter W009 or its child Tasks.

### Preserve and rebaseline W010 sequencing

Amend `DRMCP-TASK-MCP-001-10` and `DRMCP-WORK-MCP-010`.

T10 continues to track W010.
T10 must depend on:

```text
DRMCP-TASK-MCP-001-02
DRMCP-TASK-MCP-001-05
DRMCP-TASK-MCP-001-08
DRMCP-TASK-MCP-001-09
DRMCP-TASK-MCP-001-15
```

W010 must add exact dependency and impact refs for:

- `DRMCP-WORK-MCP-012`;
- `DRMCP-WORK-MCP-011`;
- `spec:drmcp.implementation`.

W010 remains `blocked`.
W010 must state:

- its optional configured legacy-fallback scope remains accepted;
- W010 consumes the completed W012 current-runtime output;
- the retired W009 extension seam is not authority;
- W010 requires a new execution-graph authoring, independent review, and release sequence after W012 completion;
- existing Task Candidates are not released implementation contracts;
- no W010 production Task may start before rebaseline;
- W010 public behavior and fixture contracts are unchanged.

Do not author W010 executor Tasks.

### Add retained validation lifecycle gates

Create:

```text
drmcp/records/tasks/mcp/
DRMCP-TASK-MCP-001-16-track-parser-aware-spec-validation-implementation.md
DRMCP-TASK-MCP-001-17-track-topics-graph-validation-implementation.md
```

T16 metadata and output:

```text
status: not_started
depends_on:
  - DRMCP-TASK-MCP-001-09
  - DRMCP-TASK-MCP-001-15
outputs:
  - DRMCP-WORK-SPEC-001
```

T16 must:

- track W-SPEC-001 through Task graph authoring, implementation, review, and `done`;
- treat W-SPEC-001 as the retained per-file detector owner;
- consume W012 runtime boundaries;
- avoid absorbing per-file validation into W012 or W010;
- perform no implementation itself.

T17 metadata and output:

```text
status: not_started
depends_on:
  - DRMCP-TASK-MCP-001-16
outputs:
  - DRMCP-WORK-SPEC-002
```

T17 must:

- track W-SPEC-002 through Task graph authoring, implementation, review, and `done`;
- require the accepted W-SPEC-001 detector-result boundary;
- treat W-SPEC-002 as the retained Topics graph owner;
- avoid duplicating W-SPEC-001 detectors;
- perform no implementation itself.

Amend W-SPEC-001 and W-SPEC-002 only to replace the retired W009 implementation dependency with W012.

Preserve:

- their accepted `retain` disposition;
- their `not_started` status;
- their semantic boundaries;
- their separate owner identities;
- their W-SPEC-001 to W-SPEC-002 dependency;
- PRODUCT semantic ownership;
- W008 fixture ownership.

Do not create implementation Tasks inside W-SPEC-001 or W-SPEC-002.

### Amend integrated verification

Amend `DRMCP-TASK-MCP-001-11`.

T11 must depend on:

```text
DRMCP-TASK-MCP-001-07
DRMCP-TASK-MCP-001-08
DRMCP-TASK-MCP-001-09
DRMCP-TASK-MCP-001-10
DRMCP-TASK-MCP-001-16
DRMCP-TASK-MCP-001-17
```

T11 remains the sole integrated verification owner.
T11 must verify:

- W012 current runtime;
- W-SPEC-001 per-file detectors;
- W-SPEC-002 Topics graph validation;
- W010 configured legacy fallback;
- current-only behavior;
- configured-fallback behavior;
- non-leakage behavior;
- accepted diagnostic and path-exposure contracts;
- full affected-package tests;
- independent integrated review.

T11 must not repair implementation failures.
Every failure must route to its owning Work Item.

T12 remains dependent on T11.
Change T12 only when a stale W009 reference exists.

### Add graph review and release Tasks

Create:

```text
drmcp/records/tasks/mcp/
DRMCP-TASK-MCP-001-14-independently-review-rebuild-read-baseline-graph-amendment.md
DRMCP-TASK-MCP-001-15-synchronize-reviewed-rebuild-read-baseline-graph-release.md
```

T14 contract:

```text
status: not_started
depends_on:
  - DRMCP-TASK-MCP-001-13
responsibility:
  independent execution-graph review
```

T14 is read-only.
T14 must verify:

- W012 consumes W011 without inheriting W009 structure;
- W012 T01 owns graph authoring, not implementation;
- T09 tracks W012 rather than W009;
- W010 waits for W012 and requires rebaseline;
- T16 and T17 preserve W-SPEC ownership;
- writer ownership is not assigned concurrently;
- the dependency graph is acyclic;
- T11 remains the aggregate owner;
- production implementation remains unreleased;
- no hidden architecture decision is delegated to an executor;
- W001, Requirement, Work Item, and Task relations are coherent.

T14 verdicts:

```text
PASS
NEEDS REVISION
NOT READY
BLOCKED
```

T15 contract:

```text
status: not_started
depends_on:
  - DRMCP-TASK-MCP-001-14
responsibility:
  reviewed graph release synchronization
```

T15 may execute only after T14 `PASS` with no blocking or major finding.
T15 must release only `DRMCP-TASK-MCP-012-01`.

T15 must not release:

- W012 production implementation;
- W010 production implementation;
- W-SPEC-001 implementation;
- W-SPEC-002 implementation;
- T11 integrated verification.

T15 must state:

- W012 T01 may author the replacement execution graph;
- production implementation remains blocked until W012's own graph review and release complete;
- T09 may begin lifecycle tracking when W012 T01 begins;
- T16, T17, and T10 retain their predecessor gates;
- executor statuses remain unchanged until their own execution begins.

### Amend W001

Add these Tasks to W001:

```text
DRMCP-TASK-MCP-001-14
DRMCP-TASK-MCP-001-15
DRMCP-TASK-MCP-001-16
DRMCP-TASK-MCP-001-17
```

Preserve T01 through T13.

Represent this flow:

```text
Architecture closure:
W011 done

Graph amendment:
T13 -> T14 -> T15

Replacement current runtime:
T09 tracks W012

Retained validation:
T16 tracks W-SPEC-001
  -> T17 tracks W-SPEC-002

Legacy fallback:
T10 tracks rebaselined W010 after T09

Integrated verification:
T11 waits for T09, T10, T16, and T17

Downstream handoff:
T12 waits for T11
```

Allowed parallelism after T09 completes:

```text
T10 and T16 may proceed in parallel after their own graphs are
reviewed and released.

T17 begins only after T16 completes.
```

W001 remains `in_progress`.

### Update Requirement relation

Append `DRMCP-WORK-MCP-012` to `DRMCP-REQ-MCP-001.work_items`.
Do not remove or reorder existing Work Item IDs.

### Future writable boundary

Restrict T13 execution to these files.

New files:

```text
drmcp/records/work-items/mcp/
DRMCP-WORK-MCP-012-rebuild-current-read-runtime-implementation.md

drmcp/records/tasks/mcp/
DRMCP-TASK-MCP-012-01-author-rebuild-current-read-runtime-execution-graph.md
DRMCP-TASK-MCP-001-14-independently-review-rebuild-read-baseline-graph-amendment.md
DRMCP-TASK-MCP-001-15-synchronize-reviewed-rebuild-read-baseline-graph-release.md
DRMCP-TASK-MCP-001-16-track-parser-aware-spec-validation-implementation.md
DRMCP-TASK-MCP-001-17-track-topics-graph-validation-implementation.md
```

Updated files:

```text
drmcp/records/tasks/mcp/
DRMCP-TASK-MCP-001-09-track-current-read-implementation.md
DRMCP-TASK-MCP-001-10-track-configured-legacy-fallback-implementation.md
DRMCP-TASK-MCP-001-11-run-integrated-read-baseline-validation-and-review.md
DRMCP-TASK-MCP-001-13-amend-read-baseline-execution-graph-for-rebuild.md

drmcp/records/work-items/mcp/
DRMCP-WORK-MCP-001-current-read-baseline-and-realignment-coordination.md
DRMCP-WORK-MCP-010-configured-legacy-archive-fallback-implementation.md

drmcp/records/work-items/spec/
DRMCP-WORK-SPEC-001-parser-aware-spec-format-validation.md
DRMCP-WORK-SPEC-002-index-topics-graph-validation.md

drmcp/records/requirements/mcp/
DRMCP-REQ-MCP-001-multi-root-multi-namespace-mcp-tool-contract.md
```

No other file may change during T13 execution.

### Prohibited operations

Do not:

- implement production code;
- implement tests;
- author fixtures;
- modify W009 or a W009 Task;
- modify W011;
- modify an ADR or Specification;
- author W012 executor cards beyond W012 T01;
- author W010 executor Tasks;
- author W-SPEC-001 implementation Tasks;
- author W-SPEC-002 implementation Tasks;
- perform independent graph review;
- release the graph;
- execute T14 or T15;
- stage or commit;
- perform repository-wide traversal;
- claim repository-wide cleanliness.

## Done condition

- W012 exists with source Requirement `DRMCP-REQ-MCP-001`.
- W012 consumes W011 and `spec:drmcp.implementation`.
- W012 excludes W009 structure and code reuse as authority.
- W012 T01 exists and owns execution-graph authoring only.
- T09 tracks W012 and no longer tracks W009.
- W009 remains unchanged and retired.
- W010 remains blocked with an exact W012 rebaseline dependency.
- T16 tracks W-SPEC-001.
- T17 tracks W-SPEC-002 after T16.
- W-SPEC-001 and W-SPEC-002 retain their accepted boundaries.
- T11 waits for W012, W010, W-SPEC-001, and W-SPEC-002 completion.
- T14 owns independent graph review.
- T15 owns reviewed graph release.
- T15 releases only W012 T01.
- Production implementation remains unreleased.
- W001 remains `in_progress`.
- REQ-MCP-001 lists W012.
- No production, test, or fixture file changes occur.
- Scoped Git inspection passes.

## Verification

Before authoring:

- confirm the exact file inventory;
- confirm T14 through T17 do not exist;
- confirm W012 and W012 T01 do not exist.

After authoring:

- confirm the exact new and modified file manifest;
- confirm `DRMCP-REQ-MCP-001` links W012;
- confirm W012 links W012 T01;
- confirm W001 links T13 through T17;
- confirm T09 links W012;
- confirm T16 links W-SPEC-001;
- confirm T17 links W-SPEC-002;
- confirm dependency acyclicity;
- confirm T11 is the sole aggregate owner;
- confirm no production release occurs before T15;
- confirm T15 does not release production implementation;
- inspect the future writable boundary with scoped `git.inspect_worktree`;
- inspect the complete textual patch with scoped `git.inspect_diff`;
- confirm whitespace passes;
- confirm no staged files;
- treat LF-to-CRLF warnings as advisory;
- do not check or infer repository-wide cleanliness.

## Evidence

### Authoring result

- Created `DRMCP-WORK-MCP-012` and `DRMCP-TASK-MCP-012-01`.
- Created T14 through T17 with the exact graph-review, release, and lifecycle responsibilities.
- Updated T09 to track W012 and recorded the exact blocker.
- Updated T10 and W010 for post-W012 execution-graph rebaseline.
- Updated W-SPEC-001 and W-SPEC-002 only to replace W009 implementation dependency with W012.
- Updated T11 as the sole integrated verification owner.
- Updated W001 and REQ-MCP-001 relations.
- W009 and its child Tasks were not changed.

### Verification result

- Exact writable-boundary manifest confirmed.
- W012 links W012 T01.
- W012 T01 depends on T15.
- T14 depends on T13.
- T15 depends on T14 and releases only W012 T01.
- T16 depends on T09 and T15.
- T17 depends on T16.
- T10 depends on T09 and T15.
- T11 depends on T09, T10, T16, and T17.
- Dependency graph is acyclic.
- T11 remains the sole integrated read-baseline verification owner.
- Production implementation remains unreleased.
- Scoped textual diff and whitespace verification completed with no staged files.
- LF-to-CRLF warnings are advisory only.
- Repository-wide cleanliness was not checked or inferred.

### Post-execution workflow defect

A coordinator review found that the authored graph moved from package architecture directly to implementation execution-graph authoring.
The graph omitted two required design stages:

- responsibility-level contract decision workflows;
- function-level internal-specification decision workflows.

Correction disposition:

- T14 is `blocked` and cannot produce a releaseable verdict.
- T15 is `blocked` with an empty release set.
- W013 owns responsibility-contract design.
- W014 owns function-level internal-specification design after W013.
- W012 and W012 T01 remain `blocked` until W013 and W014 complete reviewed closure.
- T18 and T19 track the two design hubs from W001.

Current authoring transaction support is non-operational.
Filesystem authoring was used under the current agent-authoring policy.

### Exact next gate

```text
Review the W013 and W014 hub authoring and blocked W012 disposition.

Do not execute T14, T15, W012 T01, or production implementation.
```
