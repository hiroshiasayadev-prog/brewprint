# DRMCP-TASK-MCP-009-05: Implement current list and exact retrieval

- **id**: DRMCP-TASK-MCP-009-05
- **status**: in_progress
- **date**: 2026-06-30
- **work_item**: DRMCP-WORK-MCP-009
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 2d
- **depends_on**:
  - DRMCP-TASK-MCP-009-04
- **outputs**:
  - drmcp/src/internal/designrecords/tools.go
  - drmcp/src/internal/designrecords/list_records_test.go
  - drmcp/src/internal/designrecords/get_records_test.go
  - drmcp/src/internal/designrecords/get_record_test.go
  - drmcp/src/internal/designrecords/suggest_next_record_test.go
  - drmcp/src/internal/designrecordsmcp/tools.go
  - drmcp/src/internal/designrecordsmcp/tools_call.go
  - drmcp/src/internal/designrecordsmcp/tools_call_test.go
  - drmcp/src/internal/designrecordsmcp/jsonrpc_test.go
  - drmcp/src/cmd/design-records-mcp/main_test.go

## Goal

Implement compact active-current listing and ordered exact current retrieval.

Retire obsolete public and package-local read behavior while preserving the protected compatibility surface.

## Work

- Implement `list_records` with app namespace, kind, domain, optional status, descending default order, default limit 20, and limit range 1 through 100.
- Return compact `ref`, `title`, `status`, and `date` fields plus `has_more` and warnings.
- Exclude specs from normal listing.
- Implement `get_records` with `refs`, ordered deduplication, 1 through 20 inputs, partial success, successful records only, and top-level warnings.
- Return normalized metadata, headings, and optional body without physical paths.
- Classify malformed, unsupported, unresolved, and duplicate requested refs without repair.
- Remove obsolete package-local legacy list and get behavior.
- Preserve `id_range.go` and the protected IDRange compatibility surface.
- Keep resolver semantics in T06 and validation semantics in T07.

### Current accepted implementation state

The T05 worktree contains uncommitted implementation changes.
The following test files are already deleted as part of T05B correction:

- `drmcp/src/internal/designrecords/get_record_test.go`.
- `drmcp/src/internal/designrecords/suggest_next_record_test.go`.

| finding | remaining correction |
|---|---|
| B-01 | Remove obsolete legacy package list and get behavior. |
| M-01 | Omit `kind` from compact current listing JSON. |
| M-04 | Remove obsolete package behavior and the unused adapter `idRangeSchema` helper. |

### Correction execution graph

The remaining correction is not executed directly from this aggregate Task.

| Task | responsibility | writer boundary | model route |
|---|---|---|---|
| T10 | Freeze the persistent correction graph. | Design Records only. | ChatGPT coordinator. |
| T11 | Independently review the correction graph. | Read-only. | Opus or equivalent reviewer. |
| T12 | Persist review acceptance and release executors. | W009, T05, and T10 through T14 Design Records only. | Haiku. |
| T13 | Remove only `idRangeSchema`. | `drmcp/src/internal/designrecordsmcp/tools.go`. | Haiku. |
| T14 | Remove legacy package list/get symbols and tests; omit `kind` from compact listing JSON. | `tools.go`, `list_records_test.go`, and `get_records_test.go` in `internal/designrecords`. | Sonnet. |
| T15 | Run correction-owned focused, absence, boundary, complete replacement mapping-manifest, diff, and whitespace checks. Route the diagnostic full package mechanically through the frozen exact T08 allowlist. | Read-only. | Haiku. |
| T16 | Independently disposition B-01, M-01, and M-04 and accept or reject the complete replacement mapping manifest. | Read-only. | Opus or equivalent reviewer. |
| T17 | Mechanically replace the provisional mapping with the accepted complete manifest, then close T05. | Exact T05, W009, and T13-T17 records only. | Haiku. |
| T18 | Persist the post-release blocker repair amendment. | W009, T05, T13-T17, and T18-T23 Design Records only. | ChatGPT coordinator. |
| T19 | Independently review the repair amendment. | Read-only graph and exact repair boundaries. | Opus or equivalent reviewer. |
| T20 | Persist amendment acceptance and release T21 and T22. | W009, T05, T13-T14, and T18-T23 Design Records only. | Haiku. |
| T21 | Migrate the stale adapter error-test Config fixture. | `drmcp/src/internal/designrecordsmcp/tools_call_test.go`. | Haiku. |
| T22 | Remove retired list and get assertions from one resolver test. | `drmcp/src/internal/designrecords/resolve_reference_test.go`. | Haiku. |
| T23 | Persist valid T21 and compile-only T22 focused Evidence and lifecycle state. | T21, T22, and T23 Design Records only. | Haiku. |
| T24 | Freeze the second repair amendment after the first executor results. | W009, T05, T15-T16, and T21-T26 Design Records only. | ChatGPT coordinator. |
| T25 | Independently review the amended T22 current-root contract. | Read-only exact Task and test boundaries. | Opus or equivalent reviewer. |
| T26 | Persist review acceptance and re-release the bounded T22 correction. | W009, T05, and T22-T26 Design Records only. | Haiku. |
| T27 | Replace the invalid T22 current-root migration with a stale-call-only repair contract. | W009, T05, T15-T16, T22-T23, and T27-T29 Design Records only. | ChatGPT coordinator. |
| T28 | Independently review the third T22 amendment against W003, W005, and T06 authority. | Read-only exact contract and code boundary. | Opus or equivalent reviewer. |
| T29 | Persist review acceptance and re-release only the stale-call-only T22 correction. | W009, T05, T15-T16, T22-T23, and T27-T29 Design Records only. | Haiku. |

T13 and T14 executed in parallel after T12 and stopped on stale tests outside their writer boundaries.
T18 through T20 own the reviewed repair amendment.
T21 and T22 were released in parallel after T20.
The first T21 report is rejected because its claimed Config replacement is absent from the worktree.
The first T22 execution completed the stale-call deletion but stopped on stale shared test initialization outside its contract.
T24 through T26 record the historical current-root amendment and re-release.
The second T22 execution followed that contract but exposed a contradiction with accepted current spec and resolver authority.
T27 through T29 own the stale-call-only amendment, independent review, and third release synchronization.
T21 retry may run while T28 reviews the third amendment.
T23 persists valid T21 and compile-only T22 repair results before aggregate verification.
T15, T16, and T17 execute sequentially after T23 completes.
T05 remains `in_progress` until T17 completes.

### Protected correction boundary

The correction wave must not change:

- `drmcp/src/internal/designrecords/types.go`.
- `drmcp/src/internal/designrecords/id_range.go`.
- `drmcp/src/internal/designrecords/validation.go`.
- `drmcp/src/internal/designrecords/validation_test.go`.
- `drmcp/src/internal/designrecords/authoring.go`.
- `drmcp/src/internal/designrecords/resolver.go`.
- `drmcp/src/internal/designrecords/resolve_reference_test.go`, except for the exact T22 temporary test-only repair boundary.
- fixtures, schemas, manifests, ADRs, Requirements, or Specifications.

T06 and T07 accepted ownership remains frozen.

## Done condition

- `list_records` accepts only the corrected compact filters.
- Default order is descending and default limit is 20.
- Limit outside 1 through 100 is rejected.
- Specs and future legacy records are absent from normal listing.
- `get_records` accepts `refs`, preserves first-occurrence order, deduplicates, and returns successful records only.
- Partial success uses top-level warnings and no failure placeholder records.
- Normal list and retrieval JSON contains no physical path.
- Compact current listing JSON contains no `kind` key.
- `get_record` and `suggest_next_record` are absent from the public catalog.
- Obsolete legacy list and get package behavior is absent.
- Obsolete ID-range behavior is absent from the current public list request path.
- `id_range.go` and its protected compatibility surface remain unchanged.
- W008 cases C11, C12, C16, R02-R05, R14, and R20 are covered.
- T03 shared types and P3 sibling files remain unchanged.
- T10 through T29 are complete.
- T13 and T14 scoped implementation results are accepted.
- T21 and T22 focused repair verification is accepted.
- T15 correction-owned aggregate verification records overall `PASS`.
- The optional diagnostic result is `PASS`, `DEFERRED_TO_T08`, or not run; diagnostic `BLOCKED` prevents T15 completion.
- B-01, M-01, and M-04 are independently recorded as closed by T16.
- T16 records `mapping manifest accepted: yes`.
- T17 mechanically records accepted correction evidence and the accepted manifest before setting T05 to `done`.
- A full `designrecords` package PASS is not required to close T05 before T08.
- T08 retains allowlisted authoring-test migration and the first accepted full `designrecords` package PASS.

## Verification

Focused correction verification:

```powershell
go test ./drmcp/src/internal/designrecordsmcp -run '^TestTools(CallSuccess|ListWorkflowKindEnums|CallToolErrors)$' -count=1
go test ./drmcp/src/internal/designrecords -run '^TestListRecordsCurrent' -count=1
go test ./drmcp/src/internal/designrecords -run '^(TestGetRecordsCurrent|TestNormalReadGetRecords)' -count=1
```

T15 mandatory acceptance uses only the three focused commands above plus exact symbol absence, compact-projection assertion, frozen complete replacement mapping-manifest observation, protected-boundary, scoped diff, changed-path, and whitespace checks.

T15 may run the following command only as a diagnostic:

```powershell
go test ./drmcp/src/internal/designrecords -count=1
```

Diagnostic exit `0` is not required for T05 correction closure.
T15 records diagnostic `DEFERRED_TO_T08` only when every failing test name exactly matches its frozen allowlist and no compile, setup, panic, unknown, mixed, or non-allowlisted failure exists.
Any other diagnostic failure is `BLOCKED` and prevents T15 completion.
T16 owns independent finding closure and complete replacement manifest acceptance.
T17 replaces the provisional mapping field-for-field and value-for-value with the accepted manifest.
T17 performs no merge, reference conversion, fixture-case supplementation, future-canonicalization addition, or significance judgment.
Deleted files are verified by scoped Git evidence.

## Evidence

Record:

- final exact T05 changed-file boundary;
- accepted list and retrieval JSON examples;
- warning classifications and no-path or no-`kind` assertions;
- retired symbol and catalog inventory before and after;
- focused outputs and any diagnostic full-package output;
- proof that `types.go`, `id_range.go`, `resolver.go`, validation, config production, and authoring source were not changed by the correction and repair wave;
- proof that `resolve_reference_test.go` changed only through removal of the two stale T05 API assertion blocks, with the invalid second-amendment migration fully rolled back;
- T13 and T14 execution and external-blocker evidence;
- T21 focused repair evidence and T22 compile-only stale-call-removal evidence;
- T15 correction-owned aggregate verification evidence and exact complete replacement mapping manifest;
- T16 B-01, M-01, and M-04 disposition plus `mapping manifest accepted: yes / no`;
- T17 mechanical closure synchronization evidence.

### Correction coordination state

- Topology: correction wave inside W009.
- Coordinator Tasks: T10, T11, and T12.
- Original executors: T13 and T14.
- Repair amendment: T18, T19, and T20.
- Repair executors: T21 and T22.
- Repair Evidence synchronization: T23.
- Second T22 repair amendment: T24, T25, and T26.
- Third T22 repair amendment: T27, T28, and T29.
- Aggregate verification: T15.
- Independent finding closure: T16.
- Correction closure synchronization: T17.
- T13 and T14 execution became eligible after T11 PASS and T12 release.
- Both stopped on external stale-test blockers after completing their scoped edits.
- T19 independently accepted the first repair amendment with no blocking, major, or minor finding.
- T20 released T21 and T22 for parallel execution.
- The first T21 result is rejected because no matching worktree change exists.
- T22 is blocked after completing the exact stale-call deletion.
- T24 froze a test-local current-root correction without changing shared test infrastructure.
- T25 independently accepted the second amended T22 contract and T26 re-released only T22.
- The second T22 execution exposed a contradiction with W003, W005, and T06 authority.
- T27 replaces the execution contract with a stale-call-only rollback and compile-only verification.
- T28 review is `done` with `PASS`, and T29 release synchronization is `done`.
- T22 remains `blocked` with third-release state `re_released_stale_call_only` and correction execution eligibility `true`.
- T21 retry eligibility is unchanged.
- T23, T15, T16, and T17 remain dependency-gated.

### Original correction release synchronization

```yaml
correction_graph_review: accepted
review_task: DRMCP-TASK-MCP-009-11
release_synchronization: DRMCP-TASK-MCP-009-12

finding_disposition:
  F-BLOCK-01: CLOSED
  F-MAJ-01: CLOSED
  F-MIN-01: CLOSED

released_correction_executors:
  - DRMCP-TASK-MCP-009-13
  - DRMCP-TASK-MCP-009-14

execution_mode: parallel
next_gate: both T13 and T14 must complete before T15
closure_remains_deferred_to: DRMCP-TASK-MCP-009-17
```

T05 remains `in_progress`.
Original implementation and final mapping remain owned by T13 through T17. Repair Evidence is owned by T21 through T23.
The block above records the accepted T12 release history. Current start eligibility is governed by the post-release repair amendment below.

### Post-release blocker repair amendment

```yaml
amendment_authoring: DRMCP-TASK-MCP-009-18
amendment_authoring_status: done
independent_review: DRMCP-TASK-MCP-009-19
independent_review_status: done
independent_review_verdict: PASS
blocking_findings: none
major_findings: none
minor_findings: none
release_synchronization: DRMCP-TASK-MCP-009-20
release_synchronization_status: done

blocked_original_executors:
  DRMCP-TASK-MCP-009-13:
    scoped_edit: complete
    external_blocker: drmcp/src/internal/designrecordsmcp/tools_call_test.go
    repair_task: DRMCP-TASK-MCP-009-21
  DRMCP-TASK-MCP-009-14:
    scoped_edits: complete
    external_blocker: drmcp/src/internal/designrecords/resolve_reference_test.go
    repair_task: DRMCP-TASK-MCP-009-22

repair_executors:
  - DRMCP-TASK-MCP-009-21
  - DRMCP-TASK-MCP-009-22

repair_release_state: released
released_executor_tasks:
  - DRMCP-TASK-MCP-009-21
  - DRMCP-TASK-MCP-009-22
execution_mode_after_release: parallel
first_execution_results:
  DRMCP-TASK-MCP-009-21: rejected_no_persisted_change
  DRMCP-TASK-MCP-009-22: blocked_after_stale_call_deletion
second_amendment_authoring: DRMCP-TASK-MCP-009-24
second_amendment_authoring_status: done

second_amendment_review: DRMCP-TASK-MCP-009-25
second_amendment_review_status: done
second_amendment_review_verdict: PASS
second_amendment_review_blocking_findings: none
second_amendment_review_major_findings: none
second_amendment_review_minor_findings: none

second_amendment_release: DRMCP-TASK-MCP-009-26
second_amendment_release_status: done
second_execution_result: unresolved_semantic_alias
second_execution_classification: released_contract_gap

third_amendment_authoring: DRMCP-TASK-MCP-009-27
third_amendment_authoring_status: done
third_amendment_review: DRMCP-TASK-MCP-009-28
third_amendment_review_status: done
third_amendment_review_verdict: PASS
third_amendment_release: DRMCP-TASK-MCP-009-29
third_amendment_release_status: done

t22_third_release_state: re_released_stale_call_only
released_correction_leaf: DRMCP-TASK-MCP-009-22

t21_retry_eligibility: unchanged
t22_correction_execution_eligible: true

next_gate:
  - valid T21 retry result
  - compile-only T22 execution result
  - T23 Evidence synchronization

repair_evidence_synchronization: DRMCP-TASK-MCP-009-23
aggregate_gate: DRMCP-TASK-MCP-009-15
```

T10 through T12 remain accepted and are not reopened.
T28 records `done / PASS`; T29 release synchronization is `done`.
T22 is re-released for stale-call-only execution while its lifecycle status remains `blocked`.
T15 begins only after T23 persists valid T21 and compile-only T22 focused Evidence.
T16 and T17 retain their accepted roles.

### Provisional implementation mapping

```yaml
implementation_mapping:
  status: provisional
  contract_refs:
    - DRMCP-REQ-MCP-001
    - DRMCP-ADR-MCP-001
    - DRMCP-WORK-MCP-004
    - DRMCP-WORK-MCP-006
    - DRMCP-WORK-MCP-008
  fixture_cases:
    - C11
    - C12
    - C16
    - R02
    - R03
    - R04
    - R05
    - R14
    - R20
  implementation:
    - path: drmcp/src/internal/designrecords/tools.go
      symbols: []
    - path: drmcp/src/internal/designrecordsmcp/tools.go
      symbols: []
    - path: drmcp/src/internal/designrecordsmcp/tools_call.go
      symbols: []
  verification:
    - path: drmcp/src/internal/designrecords/list_records_test.go
      tests: []
    - path: drmcp/src/internal/designrecords/get_records_test.go
      tests: []
    - path: drmcp/src/internal/designrecordsmcp/tools_call_test.go
      tests: []
    - path: drmcp/src/internal/designrecordsmcp/jsonrpc_test.go
      tests: []
    - path: drmcp/src/cmd/design-records-mcp/main_test.go
      tests: []
  future_canonicalization:
    internal_design_ref: pending
    bpdsl_ref: pending
```

T15 produces the exact complete replacement `final_implementation_mapping` manifest defined in its Task.
The manifest preserves canonical contract refs, all nine fixture cases, implementation, verification, removed items, protected paths, exclusions, and future canonicalization.
T16 independently accepts or rejects that exact manifest as a complete replacement for this provisional block.
T17 replaces this provisional block field-for-field and value-for-value with the accepted manifest.
T17 must not explore the repository, merge with this provisional block, supplement missing information, convert references, or add, remove, reorder, or reinterpret any manifest item.
