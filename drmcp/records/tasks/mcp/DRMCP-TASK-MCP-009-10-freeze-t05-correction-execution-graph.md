# DRMCP-TASK-MCP-009-10: Freeze T05 correction execution graph

- **id**: DRMCP-TASK-MCP-009-10
- **status**: done
- **date**: 2026-06-29
- **work_item**: DRMCP-WORK-MCP-009
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 0.5d
- **depends_on**:
  - DRMCP-TASK-MCP-009-04
- **outputs**:
  - DRMCP-WORK-MCP-009
  - DRMCP-TASK-MCP-009-05
  - DRMCP-TASK-MCP-009-08
  - DRMCP-TASK-MCP-009-09
  - DRMCP-TASK-MCP-009-10
  - DRMCP-TASK-MCP-009-11
  - DRMCP-TASK-MCP-009-12
  - DRMCP-TASK-MCP-009-13
  - DRMCP-TASK-MCP-009-14
  - DRMCP-TASK-MCP-009-15
  - DRMCP-TASK-MCP-009-16
  - DRMCP-TASK-MCP-009-17

## Goal

Persist the unresolved T05 correction as an executable Work Item-owned graph.

Keep production implementation, independent review, release, verification, finding closure, and lifecycle synchronization separate.

## Work

- Apply the execution-hub pattern to the unresolved T05 findings.
- Compare an in-W009 correction wave with a separate correction Work Item.
- Keep `DRMCP-REQ-MCP-001` as the single source Requirement.
- Preserve existing public Task IDs T01 through T09.
- Create exact coordinator, executor, aggregate, review, and closure Tasks.
- Freeze model routing, writer ownership, dependencies, tests, failure routing, and stop conditions.
- Keep T06 and T07 accepted ownership frozen.
- Preserve T08's existing T05 dependency and accepted ownership of authoring-test migration and the first full `designrecords` package PASS.
- Synchronize T08 wording only where required to receive a T15 `DEFERRED_TO_T08` route.
- Update T09 only to consume accepted correction evidence during final W009 review.
- Do not issue a production implementation prompt.

### Topology decision

Use an in-W009 correction wave.

| criterion | result |
|---|---|
| Source Requirement | Remains `DRMCP-REQ-MCP-001`. |
| W009 Goal and Boundary | Already own corrected list, exact retrieval, stale implementation removal, and implementation review. |
| Neighbor ownership | T06 and T07 remain frozen. T08 and T09 retain their lane roles. |
| Requirement resolution | A new Work Item would duplicate W009 ownership. |
| Public IDs | T01 through T09 remain unchanged. |
| Aggregate owner | T15 owns read-only broader verification. |
| Finding closure owner | T16 owns independent finding disposition. |
| Lifecycle owner | T17 owns correction evidence and T05 closure synchronization. |
| Pattern cost | Justified by two executor models, disjoint writers, an aggregate gate, and independent closure review. |

### Persistent correction graph

```text
T10 third correction revision
  -> T11 independent re-review
  -> T12 acceptance / release synchronization
  -> T13 adapter helper retirement ┐
  -> T14 package cleanup           ├─ parallel
                                   ┘
  -> T15 focused verification
     + exact diagnostic allowlist routing
     + complete replacement mapping manifest
  -> T16 finding and complete mapping acceptance
  -> T17 mechanical complete replacement synchronization
  -> T08 allowlisted authoring-test migration
     + first full designrecords package PASS
  -> T09 final W009 review and closure
```

### Exact Task artifacts

| ID | filename |
|---|---|
| DRMCP-TASK-MCP-009-10 | `DRMCP-TASK-MCP-009-10-freeze-t05-correction-execution-graph.md` |
| DRMCP-TASK-MCP-009-11 | `DRMCP-TASK-MCP-009-11-review-t05-correction-execution-graph.md` |
| DRMCP-TASK-MCP-009-12 | `DRMCP-TASK-MCP-009-12-synchronize-reviewed-t05-correction-release.md` |
| DRMCP-TASK-MCP-009-13 | `DRMCP-TASK-MCP-009-13-retire-unused-adapter-helper.md` |
| DRMCP-TASK-MCP-009-14 | `DRMCP-TASK-MCP-009-14-retire-legacy-list-get-and-correct-compact-projection.md` |
| DRMCP-TASK-MCP-009-15 | `DRMCP-TASK-MCP-009-15-verify-t05-correction-integration.md` |
| DRMCP-TASK-MCP-009-16 | `DRMCP-TASK-MCP-009-16-review-t05-finding-closure.md` |
| DRMCP-TASK-MCP-009-17 | `DRMCP-TASK-MCP-009-17-synchronize-t05-correction-closure.md` |

### Model routing

| Task | model route | reason |
|---|---|---|
| T10 | ChatGPT coordinator | Work Item topology and authority judgment. |
| T11 | Claude Code Opus or equivalent independent reviewer | Ownership and graph review require strong reasoning. |
| T12 | Claude Code Haiku | Exact review-result persistence and mechanical release state. |
| T13 | Claude Code Haiku | One file and one exact obsolete helper removal. |
| T14 | Claude Code Sonnet | Shared package cleanup spans production and tests and requires semantic preservation. |
| T15 | Claude Code Haiku | Read-only exact commands, deterministic failure routing, and observation of a pre-frozen mapping manifest. |
| T16 | Claude Code Opus or equivalent independent reviewer | Finding disposition, direct-regression judgment, and mapping-manifest acceptance. |
| T17 | Claude Code Haiku | Mechanical lifecycle, Evidence, and accepted-manifest synchronization. |

### Writer ownership

| phase | path or artifact | sole writer |
|---|---|---|
| Graph authoring | W009, T05, T08, T09, T10-T17 record definitions | T10 |
| Release synchronization | W009, T05, and T10 through T14 records | T12 after independent T11 output |
| Adapter correction | `drmcp/src/internal/designrecordsmcp/tools.go` | T13 |
| Package correction | `drmcp/src/internal/designrecords/tools.go` | T14 |
| Package correction | `drmcp/src/internal/designrecords/list_records_test.go` | T14 |
| Package correction | `drmcp/src/internal/designrecords/get_records_test.go` | T14 |
| Aggregate verification | No production or test writer | T15 read-only |
| Finding closure review | No production or test writer | T16 read-only |
| Correction closure persistence | T05, W009, T13-T17 Evidence and lifecycle | T17 |
| Final Work Item closure | T09, W009, and hub T09 | Existing T09 closure slice |

T13 and T14 may run in parallel after T12 because their writable files are disjoint.

## Done condition

- W009 owns T10 through T17 in metadata and Task flow.
- T05 is marked `in_progress` and delegates unresolved correction execution to T13 through T17.
- T09 consumes accepted correction closure evidence without replacing T16.
- Every new Task has an exact ID, filename, dependency, writer boundary, model route, and verification owner.
- T13 and T14 have no overlapping writable path.
- T15 is read-only, freezes an exact T08-owned test-name allowlist, classifies diagnostic failure mechanically, and does not require full `designrecords` PASS.
- T15 freezes the complete replacement implementation mapping for independent acceptance.
- T16 independently owns B-01, M-01, and M-04 disposition and complete replacement mapping-manifest acceptance.
- T17 is the only correction closure synchronization owner and performs no repository exploration, merge, supplementation, conversion, or significance judgment.
- No production source, Go test, fixture, schema, manifest, ADR, Requirement, or Specification changes.
- The graph is ready for independent review.

## Verification

- Confirm no T10 through T17 file existed before authoring.
- Confirm every new Task uses the W009 work sequence and `DRMCP-REQ-MCP-001`.
- Confirm W009 lists T01 through T17 exactly once.
- Confirm dependency edges are acyclic.
- Confirm T13 and T14 writable paths are disjoint.
- Confirm T08 still depends on T05.
- Confirm T09 remains the final W009 review and closure owner.
- Inspect only changed Design Record paths with scoped Git and whitespace tools.

## Evidence

- Topology decision: existing W009 correction wave.
- Separate correction Work Item: rejected because the Requirement, Goal, Boundary, and completion judgment remain W009-owned.
- New Task IDs T10 through T17 were unused in the bounded `drmcp/records/tasks/mcp` listing.
- The accepted T05 findings are B-01, M-01, and M-04.
- T13 owns the M-04 adapter-helper clause.
- T14 owns B-01, M-01, and the M-04 package-retirement clause.
- T15 owns correction-focused aggregate verification and no repair.
- Correction revision after T11 `NEEDS REVISION`: the first accepted full `designrecords` package PASS remains T08-owned; T15 may route only known authoring-test migration failures as `DEFERRED_TO_T08`.
- Correction revision after T11 `NEEDS REVISION`: T15 now produces an exact final mapping manifest, T16 accepts it independently, and T17 copies it mechanically.
- Correction revision after T11 `NEEDS REVISION`: T14 and T15 absence, Git, and whitespace checks are fixed to exact commands and paths.
- Second correction revision after the second T11 `NEEDS REVISION`: T15 freezes the exact T08-owned test-name allowlist from T04, T08, and the permitted authoring test declarations.
- Second correction revision: T15 classifies diagnostic full-package failures mechanically by exact allowlist membership; compile, setup, panic, unknown, mixed, and non-allowlisted failures are `BLOCKED`.
- Second correction revision: T15 provides a complete replacement mapping manifest with canonical record IDs, all nine T05 fixture cases, and preserved future canonicalization.
- Second correction revision: the mandatory adapter command executes `TestToolsCallSuccess`, `TestToolsListWorkflowKindEnums`, and `TestToolsCallToolErrors`.
- Second correction revision: T16 judges complete-replacement usability, and T17 performs field-for-field and value-for-value replacement without merge or supplementation.
- Third correction revision after the third T11 `NEEDS REVISION`: T15 evaluates unconditional compile, setup, initialization, panic, unknown, and mechanically unclassifiable blockers before allowlist comparison.
- Third correction revision: an allowlisted-test panic and an allowlisted but unclassifiable failure are explicitly `BLOCKED`.
- Third correction revision: diagnostic classification is total and ordered across exactly `PASS`, `DEFERRED_TO_T08`, and `BLOCKED`.
- Third correction revision: T15 Evidence invariants are synchronized with terminal result, exit code, non-test failure flags, allowlist comparison, and routed owner.
- T16 owns independent finding disposition and mapping acceptance.
- T17 owns persistence of accepted execution, verification, review, mapping, and T05 lifecycle evidence.

### Final graph review

```yaml
final_graph_review:
  task: DRMCP-TASK-MCP-009-11
  verdict: PASS

accepted_findings:
  F-BLOCK-01: CLOSED
  F-MAJ-01: CLOSED
  F-MIN-01: CLOSED

release_task: DRMCP-TASK-MCP-009-12
released_executors:
  - DRMCP-TASK-MCP-009-13
  - DRMCP-TASK-MCP-009-14
execution_mode: parallel
```

The accepted graph preserves the third correction revision and all earlier correction history.
- DRMCP authoring transaction tools were not present in the current tool surface.
- Filesystem authoring was used under the active `drmcp/records` namespace.
- Production implementation and implementation prompts were not started.
