# DRMCP-WORK-MCP-009: Current-format read implementation

- **id**: DRMCP-WORK-MCP-009
- **status**: in_progress
- **date**: 2026-06-30
- **source_requirement**: DRMCP-REQ-MCP-001
- **impact_refs**:
  - DRMCP-ADR-MCP-001
  - DRMCP-INV-MCP-002
  - DRMCP-WORK-MCP-001
  - DRMCP-WORK-MCP-003
  - DRMCP-WORK-MCP-004
  - DRMCP-WORK-MCP-005
  - DRMCP-WORK-MCP-006
  - DRMCP-WORK-MCP-007
  - DRMCP-WORK-MCP-008
  - DRMCP-TASK-MCP-001-09
  - spec:drmcp.design_records_mcp.overview
  - spec:drmcp.design_records_mcp.schema.overview
  - spec:drmcp.design_records_mcp.tools.overview
- **tasks**:
  - DRMCP-TASK-MCP-009-01
  - DRMCP-TASK-MCP-009-02
  - DRMCP-TASK-MCP-009-03
  - DRMCP-TASK-MCP-009-04
  - DRMCP-TASK-MCP-009-05
  - DRMCP-TASK-MCP-009-06
  - DRMCP-TASK-MCP-009-07
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
  - DRMCP-TASK-MCP-009-18
  - DRMCP-TASK-MCP-009-19
  - DRMCP-TASK-MCP-009-20
  - DRMCP-TASK-MCP-009-21
  - DRMCP-TASK-MCP-009-22
  - DRMCP-TASK-MCP-009-23
  - DRMCP-TASK-MCP-009-24
  - DRMCP-TASK-MCP-009-25
  - DRMCP-TASK-MCP-009-26
  - DRMCP-TASK-MCP-009-27
  - DRMCP-TASK-MCP-009-28
  - DRMCP-TASK-MCP-009-29
  - DRMCP-TASK-MCP-009-30
  - DRMCP-TASK-MCP-009-31
  - DRMCP-TASK-MCP-009-32
  - DRMCP-TASK-MCP-009-33
  - DRMCP-TASK-MCP-009-34
  - DRMCP-TASK-MCP-009-35
  - DRMCP-TASK-MCP-009-36

## Goal

Implement one current-only DRMCP read path against the corrected contracts and accepted current fixtures.

Remove legacy assumptions from active discovery, indexing, query, retrieval, resolution, validation, diagnostics, and normal responses.

## Boundary

This Work Item owns:

- configured current-root loading and app association;
- current sequential record parsing;
- H1-adjacent current spec parsing;
- path-derived canonical `spec:` ref handling;
- deterministic active-index construction and duplicate detection;
- an index structure that remains separate from any later legacy archive index;
- corrected `list_records` implementation;
- retirement of the public `get_record` tool;
- corrected `get_records` implementation for current canonical refs;
- current-side canonical resolution;
- current cross-namespace relation validation;
- current repository validation and diagnostics;
- physical-path hiding in normal listing and retrieval responses;
- current-only automated tests using `DRMCP-WORK-MCP-008` fixtures;
- current-only integration verification without configured legacy roots;
- removal or correction of stale implementation and test assumptions;
- implementation review and closure evidence;
- the persistent T05 correction graph for findings B-01, M-01, and M-04;
- independent review, release synchronization, aggregate verification, finding closure, and lifecycle synchronization for that correction wave;
- the persistent post-release repair amendment for T13 and T14 external stale-test blockers;
- the persistent T15 verification-ownership amendment and executor-leaf split.

The T05 correction remains inside this Work Item because it uses the same source Requirement and implements the existing W009 Goal and Boundary.
A separate correction Work Item would duplicate Requirement-resolution ownership.

This Work Item does not own:

- contract or semantic rule design;
- fixture design or fixture authoring;
- accepted legacy-family semantics;
- legacy-root loading, legacy parsing, or legacy index construction;
- exact legacy retrieval or fallback resolution;
- current-to-legacy relation validation implementation;
- legacy leakage tests;
- authoring transaction, guidance, create, or update behavior;
- portable standards-package behavior;
- retained future spec-format validator implementation owned through T07 disposition.

This Work Item may add extension seams required by later legacy fallback.
This Work Item must not implement legacy behavior through those seams.

## Impact Scope

| ref or area | impact |
|---|---|
| `DRMCP-REQ-MCP-001` | Source Requirement for the current-format read implementation. |
| `DRMCP-ADR-MCP-001` | Governs current-format-first implementation and removal of legacy active assumptions. |
| `DRMCP-INV-MCP-002` | Supplies implementation-drift and stale-test evidence. |
| `DRMCP-WORK-MCP-003` | Supplies discovery, parsing, identity, and active-index contracts. |
| `DRMCP-WORK-MCP-004` | Supplies listing and exact current retrieval contracts. |
| `DRMCP-WORK-MCP-005` | Supplies current-side resolver behavior. |
| `DRMCP-WORK-MCP-006` | Supplies validation, diagnostic, and path-exposure contracts. |
| `DRMCP-WORK-MCP-007` | Supplies validation-work non-overlap and rebaseline results. |
| `DRMCP-WORK-MCP-008` | Supplies accepted current fixtures. |
| DRMCP parser, index, resolver, validation, tool, and server code | Receive the current-only implementation changes. |
| Existing automated tests | Are rewritten to assert corrected contracts rather than legacy behavior. |

## Task flow

| parallel group | Task | dependency gate | outcome |
|---|---|---|---|
| P0 | T01 | W003-W008 accepted | Freeze the contract-to-code map, file ownership, execution slices, and review gates. |
| P1 | T02 | T01 accepted | Replace current-root auto-discovery with explicit configured current roots and defer the full-package gate to T04. |
| P1 | T03 | T01 accepted | Freeze shared current read types, implement exact current parsers, and defer the full-package gate to T04. |
| P2 | T04 | T02 and T03 targeted acceptance | Integrate T02/T03, build the active index, split public operation ownership, freeze shared APIs, and classify transitional failures for T05-T07 or T08 ownership. |
| P3 | T05 | T04 accepted | Implement compact listing, exact retrieval, path hiding, and retired-tool cleanup on the post-T04 `tools.go` boundary. |
| P3 | T06 | T04 accepted | Implement exact current canonical resolution on the post-T04 `resolver.go` boundary. |
| P3 | T07 | T04 accepted | Implement current validation and portable diagnostics on the post-T04 `validation.go` boundary. |
| P3C0 | T10 | Unresolved T05 findings identified | Freeze the T05 correction topology, executor contracts, writers, verification ownership, and closure ownership. |
| P3C1 | T11 | T10 correction revision complete | Independently re-review F-BLOCK-01, F-MAJ-01, F-MIN-01, and the revised correction graph without changing files. |
| P3C2 | T12 | T11 re-review PASS; all three findings CLOSED; no new blocking, major, or minor finding | Persist review acceptance and release T13 and T14. |
| P3C3 | T13 | T12 release | Remove only the unused adapter `idRangeSchema` helper. |
| P3C3 | T14 | T12 release | Remove obsolete package list/get behavior and legacy tests, then correct compact listing projection. |
| P3C3A | T18 | T13 and T14 external blocker evidence available | Persist the post-release blocker repair graph without reopening accepted T10 through T12 history. |
| P3C3B | T19 | T18 complete | Independently review repair ownership, temporary T06 test ownership, dependencies, and downstream verification. |
| P3C3C | T20 | T19 PASS with no blocking, major, or minor finding | Persist amendment acceptance and release T21 and T22. |
| P3C3D | T21 | T20 release | Migrate the adapter error-test Config fixture to one explicit current root. |
| P3C3D | T22 | T20 release | Remove retired list and single-get assertions from one resolver test while preserving semantic-ref resolution assertions. |
| P3C3F | T24 | First T21 and T22 executor results available | Reject the unpersisted T21 result and freeze the exact test-local current-root amendment for T22. |
| P3C3G | T25 | T24 complete | Independently review the amended T22 contract, shared-helper protection, and T06 preservation. |
| P3C3H | T26 | T25 PASS with no blocking, major, or minor finding | Persist review acceptance and re-release the bounded T22 correction. |
| P3C3I | T27 | Second T22 execution exposes unresolved semantic alias behavior | Replace the invalid current-root migration with a stale-call-only T22 contract and compile-only verification. |
| P3C3J | T28 | T27 complete | Independently review the third amendment against W003, W005, and T06 authority. |
| P3C3K | T29 | T28 PASS with no blocking, major, or minor finding | Persist review acceptance and re-release only the stale-call-only T22 correction. |
| P3C3E | T23 | Valid T21 and compile-only T22 execution results available | Persist focused repair Evidence and set T21 and T22 to done without running aggregate verification. |
| P3C3L | T30 | T23 complete and the first T15 execution exposes the aggregate ownership defect | Persist the amended verification graph, exact diagnostic owner map, writer boundaries, and model routing. |
| P3C3M | T31 | T30 complete | Independently review graph executability, exact owner maps, and separation of command, aggregate, mapping, review, and closure responsibility. |
| P3C3N | T32 | T31 PASS with no blocking, major, or minor finding | Synchronize review acceptance and release the read-only verification leaves. |
| P3C3O | T33 | T32 release | Run only the three T05 mandatory focused commands. |
| P3C3O | T34 | T32 release | Run only static absence, presence, diff, protected-boundary, and whitespace checks. |
| P3C3O | T35 | T32 release | Run one diagnostic full-package command and inventory exact failures without owner judgment. |
| P3C3P | T36 | T32 release and T34 report complete | Mechanically assemble the frozen mapping manifest from T34 item-level Evidence. |
| P3C4 | T15 | T33 through T36 reports complete | Synchronize leaf Evidence, route exact diagnostic names through T30, and determine aggregate T05 correction result without running commands. |
| P3C5 | T16 | T15 overall PASS | Independently judge B-01, M-01, and M-04 closure, direct regressions, diagnostic routing, and complete replacement mapping acceptance. |
| P3C6 | T17 | T16 PASS and complete replacement mapping accepted | Mechanically replace T05 mapping field-for-field and value-for-value, then close T05 without closing W009 or downstream lanes. |
| P4 | T08 | T05, T06, and T07 accepted and merged; their owned focused tests pass; any remaining mapped T08 authoring-test migration is identified | Migrate protected authoring test setup to the final current-only fixture/config shape, then own the first accepted full `designrecords` package PASS and current fixture integration. |
| P5 | T09 | T08 accepted | Run independent review, corrections, and synchronized closure. |

P1 and P3 may run in parallel only in separate branches or worktrees.
T02 and T03 individual branches require targeted acceptance, scoped format checks, scoped Git evidence, and package compile.
T02 and T03 do not require full-package PASS before individual acceptance.
T04 owns the serial foundation integration gate for the merged T02/T03 branches.
T04 must make the package compile and must record remaining full-package failures by failing test and assigned T05-T07 or T08 owner.
Protected authoring tests may remain transitional only when failure is caused by retired current-root, identity, or spec-format setup; T08 owns their test-only migration after P3 stabilizes the read API.
T04 must reject any other unresolved ownership, API mismatch, or stale-behavior workaround.
Post-T04 P3 file boundaries are disjoint.
T05 remains open while its correction wave executes.
T13 and T14 were released in parallel after T12 because their writable paths are disjoint.
Both executors completed their scoped edits and stopped on mandatory failures caused by stale tests outside their writer boundaries.
T18 through T20 own the reviewed repair amendment.
T21 and T22 were released in parallel after T20 because their writable paths are disjoint.
The first T21 report was rejected because its claimed Config replacement was absent from the worktree.
The first T22 execution completed the stale-call deletion but exposed stale shared test initialization outside its released contract.
T24 through T26 record the historical current-root amendment and re-release.
The second T22 execution followed that contract but exposed a contradiction with current spec and resolver authority.
T27 through T29 own the stale-call-only amendment, independent review, and third release synchronization.
T21 retry may run independently while T28 reviews the third amendment.
T22 has temporary test-only ownership only for removing stale T05 API calls from `resolve_reference_test.go`; T06 retains resolver assertions, `resolver.go`, and resolver semantic ownership.
T23 is the sole focused repair Evidence and lifecycle synchronization owner before the amended verification wave.
The first T15 execution report is amendment-trigger evidence only because it omitted exact failing test names and applied the defective aggregate contract.
T30 through T32 own amendment authoring, independent review, and release synchronization.
T33 owns only the T05 focused commands.
T34 owns only static correction-boundary checks.
T35 owns only the diagnostic full-package failure inventory and makes no owner judgment.
T36 owns only mechanical mapping observation after T34.
T15 runs no repository command. T15 synchronizes T33 through T36 Evidence, routes exact test names through the reviewed T30 map, and proposes the complete replacement manifest.
Mapped T06, T07, and T08 failures do not prevent T15 overall `PASS` and do not decide T05 finding closure.
T05-mapped, unmapped, compile, setup, initialization, panic, unknown, or mechanically unclassifiable failures block T15.
T16 is the only owner of B-01, M-01, and M-04 finding disposition and complete replacement mapping acceptance.
T17 is the only correction closure synchronization owner. T17 replaces T05 mapping field-for-field and value-for-value from the T16-accepted manifest without merge, supplementation, conversion, or judgment and must leave W009 `in_progress`.
T08 cannot start until T05, T06, and T07 are integrated and their owned focused tests pass. T08 then owns authoring test-only migration, fixture and adapter integration verification, and the first accepted full `designrecords` package PASS.
A downstream slice must escalate to the owning upstream Task instead of editing an upstream-owned shared file.
Configured legacy fallback proceeds only after this Work Item is accepted.

## Task Candidates

| Task | owner model | primary file boundary | dependency |
|---|---|---|---|
| T01 | Sonnet design; Haiku mechanical verification | W009 workflow records only | Accepted upstream Work Items. |
| T02 | Sonnet implementation; Haiku verification | `config.go`, `config_test.go` | T01. |
| T03 | Sonnet implementation; Haiku verification | `types.go`, `types_test.go`, `parser.go`, `parser_index_test.go` | T01. |
| T04 | Sonnet integration and ownership split; Haiku verification | `index.go`, `index_test.go`, `tools.go`, `resolver.go`, `validation.go`, and direct compile-adjustment tests only | T02-T03 targeted acceptance. |
| T05 | Aggregate list/get implementation lane; correction delegated through T13-T36 | existing T05 changed-file inventory; `id_range.go` protected from the correction wave | T04, then T10-T36 correction gate. |
| T06 | Sonnet implementation; Haiku verification | `resolver.go`, `resolve_reference_test.go` | T04. |
| T07 | Sonnet implementation; Haiku verification | `validation.go`, `validation_test.go` | T04. |
| T08 | Sonnet integration and test-only migration; Haiku complete verification | new `current_read_fixture_test.go`, `authoring_test.go`, and `authoring_guidance_test.go`; production authoring source and fixtures read-only | T05-T07. |
| T09 | Sonnet review/correction; Haiku closure sync | accepted Task boundaries, then three closure records | T08. |
| T10 | ChatGPT coordinator | W009, T05, T08, T09, and T10-T17 Design Records only | Unresolved T05 findings. |
| T11 | Opus or equivalent independent reviewer | W009 correction graph and exact leaf boundaries; read-only | T10. |
| T12 | Haiku release synchronization | W009, T05, and T10 through T14 Design Records only | T11 PASS. |
| T13 | Haiku executor | `drmcp/src/internal/designrecordsmcp/tools.go` | T12. |
| T14 | Sonnet executor | `tools.go`, `list_records_test.go`, and `get_records_test.go` in `internal/designrecords` | T12. |
| T15 | Sonnet or ChatGPT aggregate owner | T30 and T33-T36 Design Record Evidence only; no repository commands or source reads | T33-T36. |
| T16 | Opus or equivalent independent reviewer | B-01, M-01, and M-04 correction diff, T15 overall result, diagnostic routing, and complete replacement mapping manifest; read-only | T15 overall PASS. |
| T17 | Haiku closure synchronization | Exact T05, T13-T17, and W009 Design Records only | T16 PASS and complete mapping accepted. |
| T18 | ChatGPT coordinator | W009, T05, T13-T17, and T18-T23 Design Records only | T13 and T14 external blocker evidence. |
| T19 | Opus or equivalent independent reviewer | Amended graph, exact repair files, and current-root contract; read-only | T18. |
| T20 | Haiku release synchronization | W009, T05, T13-T14, and T18-T23 Design Records only | T19 PASS. |
| T21 | Haiku executor | `drmcp/src/internal/designrecordsmcp/tools_call_test.go` | T20. |
| T22 | Haiku executor | `drmcp/src/internal/designrecords/resolve_reference_test.go` | T20. |
| T23 | Haiku Evidence synchronization | T21, T22, and T23 Design Records only | Valid T21 and compile-only T22 execution results. |
| T24 | ChatGPT coordinator | W009, T05, T15-T16, and T21-T26 Design Records only | First T21 and T22 executor results. |
| T25 | Opus or equivalent independent reviewer | Amended T22 contract, exact test file, Config contract, and shared helper; read-only | T24. |
| T26 | Haiku release synchronization | W009, T05, T22-T26 Design Records only | T25 PASS. |
| T27 | ChatGPT coordinator | W009, T05, T15-T16, T22-T23, and T27-T29 Design Records only | Second T22 execution result. |
| T28 | Opus or equivalent independent reviewer | Third amendment, W003/W005/T06 authority, and exact T22 diff; read-only | T27. |
| T29 | Haiku release synchronization | W009, T05, T15-T16, T22-T23, and T27-T29 Design Records only | T28 PASS. |
| T30 | ChatGPT coordinator | W009, T15-T17, and T30-T36 Design Records only | T23 and the first T15 graph-defect report. |
| T31 | Opus or equivalent independent reviewer | W009, T06-T08, T15-T17, T23, T30-T36, and exact test declaration files; read-only | T30. |
| T32 | Haiku release synchronization | W009, T15, and T30-T36 Design Records only | T31 PASS. |
| T33 | Haiku focused verifier | Three exact T05 focused Go commands; read-only | T32. |
| T34 | Haiku static verifier | Six writer paths, seven protected paths, and T13/T14/T21-T23 Evidence; read-only | T32. |
| T35 | Sonnet diagnostic inventory | One JSON full-package command; no ownership judgment or writes | T32. |
| T36 | Haiku mapping observer | T05 provisional mapping plus complete T34 report; no source reads or commands | T32; complete T34 report start gate. |

Haiku never owns contract interpretation, shared-type design, parser behavior, index semantics, resolver behavior, diagnostic semantics, or review findings.

## Completion Condition

This Work Item is complete when all of the following are true:

- configured current roots build one deterministic active index;
- current sequential records and current specs use corrected parsing and identity behavior;
- duplicate current identity fails deterministically;
- current and future legacy index structures remain separate;
- corrected listing and exact current retrieval contracts are implemented;
- `get_record` is absent from the public tool surface;
- current canonical resolution and cross-namespace validation match accepted contracts;
- current repository validation and machine-readable diagnostics match accepted contracts;
- normal listing and retrieval responses do not expose physical paths;
- current-only operation passes with no configured legacy roots;
- legacy archive records and legacy active assumptions are absent from normal current behavior;
- accepted current fixtures have automated coverage;
- all relevant automated tests pass;
- T02 through T08 contain completed provisional `implementation_mapping` Evidence with accepted `contract_refs`, owned `fixture_cases`, actual implementation paths and contract-significant symbols, and actual verification paths and test functions;
- each mapping keeps `internal_design_ref` and `bpdsl_ref` as `pending` unless a formal canonical ID exists, and Task Evidence is not treated as current-state implementation-design authority;
- T05 correction Tasks T10 through T36 are complete;
- T13 and T14 scoped implementation results plus T21 and T22 focused repair verification are accepted, and T15 records overall `PASS`;
- T15 diagnostic routing uses the independently reviewed exact T05, T06, T07, and T08 owner map;
- T05-mapped, unmapped, compile, setup, initialization, panic, unknown, and mechanically unclassifiable failures block T15;
- exact T06, T07, and T08 failures may route downstream without deciding T05 closure;
- T16 independently records B-01, M-01, and M-04 as closed and accepts the complete replacement implementation mapping manifest;
- T17 mechanically replaces T05 mapping field-for-field and value-for-value from the accepted manifest;
- any diagnostic full `designrecords` package failure routed by T15 contains only exact T06, T07, or T08 owner-map entries;
- T08 records the first accepted full `designrecords` package PASS after migration;
- T05 correction and repair writers did not modify `types.go`, `id_range.go`, validation, `resolver.go`, config production, or authoring source;
- T22 changed `resolve_reference_test.go` only by deleting the two stale T05 API assertion blocks after fully rolling back the invalid second-amendment migration;
- implementation review reports no blocking or major findings;
- `DRMCP-REQ-MCP-001` lists this Work Item in `work_items`;
- final evidence records changed code, tests, commands, results, review verdict, and residual limitations.

## Evidence

- `DRMCP-ADR-MCP-001`: Accepted current-format-first implementation sequence.
- `DRMCP-INV-MCP-002`: Implementation and test drift baseline.
- `DRMCP-WORK-MCP-003` through `DRMCP-WORK-MCP-008`: Accepted contract, disposition, and fixture inputs.
- `DRMCP-TASK-MCP-001-09`: Hub lifecycle gate for this Work Item.
- `DRMCP-TASK-MCP-009-01`: Bounded contract-to-code correspondence, accepted-fixture allocation, parallel graph, model ownership, exact file boundaries, command plans, and escalation gates.
- `DRMCP-TASK-MCP-009-10`: Chose an in-W009 correction wave and froze T11-T17 ownership.
- Topology reason: the correction uses `DRMCP-REQ-MCP-001`, stays inside the existing W009 Goal and Boundary, preserves T01-T09 public IDs, and does not reopen T06 or T07.
- T13 and T14 are disjoint correction writers. T21 and T22 are disjoint test-only repair writers. T33 through T36 split focused commands, static boundaries, package inventory, and mapping observation. T15 owns only aggregate Evidence synchronization, exact reviewed routing, and the proposed complete replacement mapping manifest. T16 owns independent finding and mapping acceptance. T17 owns mechanical T05 correction closure synchronization.
- First correction revision: removed the T15/T08 full-package ownership deadlock. T08 retains authoring-test migration and the first accepted full `designrecords` package PASS.
- Second correction revision: froze the exact T08-owned test-name allowlist, mechanical diagnostic classification, canonical mapping record IDs, all nine T05 fixture cases, future canonicalization, and mandatory `TestToolsCallSuccess` execution.
- Second correction revision: T16 judges complete-replacement usability and T17 performs no manual merge or supplementation.

### Original T05 correction release synchronization

```text
P3C0 T10 graph freeze:
  accepted

P3C1 T11 independent review:
  PASS

P3C2 T12 release synchronization:
  done

P3C3 parallel correction executors:
  T13: released, not_started
  T14: released, not_started

P3C4 T15 aggregate correction verification:
  not_started
  blocked by T13 and T14

P3C5 T16 independent finding closure:
  not_started
  blocked by T15

P3C6 T17 T05 correction closure:
  not_started
  blocked by T16
```

- T13 and T14 have non-overlapping writer paths and may execute in parallel.
- T08 retains its accepted role and remains dependency-gated by T05, T06, and T07.
- T09 remains the sole final W009 closure owner.
- W009 remains `in_progress`.
- DRMCP authoring transactions were unavailable in the current tool surface, so filesystem authoring was used under the active `drmcp/records` namespace.
- Post-release blocker amendment: T18 authored T19 through T23 and revised T15 through T17 connectivity after T13 and T14 stopped on external stale-test blockers.
- T19 independent amendment review: `PASS`; no blocking, major, or minor finding.
- T20 release synchronization: `done`; T21 and T22 released in parallel.
- First repair execution results: T21 report rejected because no matching worktree change exists; T22 stale-call deletion accepted but focused verification blocked on stale shared test initialization.
- T24 second amendment: `done`; T25 independent review: `PASS`; T26 release synchronization: `done`.
- Second T22 execution followed the released contract but returned `unresolved` for the stale semantic alias assertion.
- T27 third amendment: `done`; T28 independent review: `done / PASS`; T29 release synchronization: `done`.
- T21 final retry result: `PASS`; focused command exit `0`; 16 cases preserved.
- T22 final stale-call-only result: `PASS`; compile-only exit `0`; stale-call grep exit `1`; no resolver semantic acceptance claimed.
- T21, T22, and T23 are `done` with scoped whitespace `PASS`.
- The first T15 execution returned `BLOCKED` under the defective pre-amendment graph. Its mandatory T05 focused commands passed, but its report omitted exact failing test names.
- T30 classifies that report as amendment-trigger evidence only and introduces T31-T36.
- Implementation, automated-test, and independent-review evidence remains pending for T02 through T09 and the reviewed T30-T36 amendment wave where not already recorded.

### T05 post-release blocker repair amendment

```text
T13 adapter source edit:
  status: blocked
  scoped edit: complete
  blocker path: drmcp/src/internal/designrecordsmcp/tools_call_test.go
  repair owner: T21 after reviewed release

T14 package cleanup:
  status: blocked
  scoped edits: complete
  blocker path: drmcp/src/internal/designrecords/resolve_reference_test.go
  repair owner: T22 after reviewed release

T18 graph amendment:
  status: done

T19 independent amendment review:
  status: done
  verdict: PASS
  blocking findings: none
  major findings: none
  minor findings: none

T20 repair release synchronization:
  status: done
  result: PASS

T25 independent T22 amendment review:
  status: done
  verdict: PASS
  blocking findings: none
  major findings: none
  minor findings: none

T26 T22 correction re-release synchronization:
  status: done

T27 stale-call-only amendment:
  status: done

T28 independent third-amendment review:
  status: done
  verdict: PASS

T29 stale-call-only re-release synchronization:
  status: done

T21 final retry:
  status: done
  focused test exit: 0
  cases preserved: 16

T22 final stale-call-only execution:
  status: done
  compile-only exit: 0
  stale-call grep exit: 1
  resolver semantic acceptance: not claimed

T23:
  status: done
  result: PASS

First T15 execution:
  persistent status: not_started
  mandatory focused commands: PASS
  diagnostic full package: BLOCKED under pre-amendment contract
  exact failing test names: missing from report
  accepted as T15 Evidence: false

T30 amendment:
  status: done
  next gate: T31 independent review

T15 amended aggregate verification:
  status: not_started
  blocked by: T31, T32, and T33-T36 reports
```

- T10 through T12 remain accepted and are not reopened.
- T19 independently accepted the amendment with no blocking, major, or minor finding.
- T20 released T21 and T22 for parallel execution.
- The rejected first T21 report remains historical evidence; the clean retry later passed and is persisted by T23.
- T24 froze the historical test-local current-root correction.
- T25 returned `PASS`, and T26 re-released only T22 under that historical second amendment.
- The second T22 execution exposed that the second amendment conflicted with W003, W005, and T06 authority.
- T27 superseded the execution contract with a stale-call-only rollback and compile-only verification.
- T28 returned `PASS`; T29 re-released T22 for stale-call-only execution.
- The final T22 execution passed compile-only verification and stale-call absence without a resolver semantic claim.
- T23 persisted valid T21 and T22 Evidence and set all three repair records to `done`.
- The first T15 execution exposed a verification ownership defect and is not synchronized as T15 Evidence.
- T30 through T36 split the amended verification graph before any T15 retry or T16 review.
- T16 and T17 retain independent review and closure responsibilities.

Planning lifecycle began on 2026-06-28.
The T01 planning gate was accepted on 2026-06-28 after an independent review returned PASS with no blocking, major, or minor finding.
T02 and T03 are ready to start in parallel under the accepted P1 boundaries.
No production Go source, existing Go test, or accepted fixture was changed during Task graph authoring or closure synchronization.
