# DRMCP-WORK-MCP-009: Current-format read implementation

- **id**: DRMCP-WORK-MCP-009
- **status**: blocked
- **date**: 2026-06-28
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
  - DRMCP-WORK-MCP-011
  - DRMCP-TASK-MCP-001-09
  - spec:drmcp.implementation
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
- implementation review and closure evidence.

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
| P2 | T04 | T02 and T03 targeted acceptance | Integrate T02/T03, build the active index, split public operation ownership, freeze shared APIs, and run the serial foundation gate. |
| P3 | T05 | T04 accepted | Implement compact listing, exact retrieval, path hiding, and retired-tool cleanup on the post-T04 `tools.go` boundary. |
| P3 | T06 | T04 accepted | Implement exact current canonical resolution on the post-T04 `resolver.go` boundary. |
| P3 | T07 | T04 accepted | Implement current validation and portable diagnostics on the post-T04 `validation.go` boundary. |
| P4 | T08 | T05, T06, and T07 accepted, merged, and passing the P3 integration gate | Run accepted current fixture integration with no legacy roots. |
| P5 | T09 | T08 accepted | Run independent review, corrections, and synchronized closure. |

P1 and P3 may run in parallel only in separate branches or worktrees.
T02 and T03 individual branches require targeted acceptance, scoped format checks, scoped Git evidence, and package compile.
T02 and T03 do not require full-package PASS before individual acceptance.
T04 owns the serial foundation integration gate for the merged T02/T03 branches.
T04 must make the package compile and must record any remaining full-package failures by failing test and T05-T07 owner.
T04 must reject unresolved ownership, API mismatch, or stale-behavior workarounds.
Post-T04 P3 file boundaries are disjoint.
T08 cannot start until T05, T06, and T07 are integrated and their combined state reaches full-package PASS.
A downstream slice must escalate to the owning upstream Task instead of editing an upstream-owned shared file.
Configured legacy fallback proceeds only after this Work Item is accepted.

## Task Candidates

| Task | owner model | primary file boundary | dependency |
|---|---|---|---|
| T01 | Sonnet design; Haiku mechanical verification | W009 workflow records only | Accepted upstream Work Items. |
| T02 | Sonnet implementation; Haiku verification | `config.go`, `config_test.go` | T01. |
| T03 | Sonnet implementation; Haiku verification | `types.go`, `types_test.go`, `parser.go`, `parser_index_test.go` | T01. |
| T04 | Sonnet integration and ownership split; Haiku verification | `index.go`, `index_test.go`, `tools.go`, `resolver.go`, `validation.go`, and direct compile-adjustment tests only | T02-T03 targeted acceptance. |
| T05 | Sonnet implementation and catalog inventory; Haiku verification | `tools.go`, `id_range.go`, list/get tests, retired public catalog files | T04. |
| T06 | Sonnet implementation; Haiku verification | `resolver.go`, `resolve_reference_test.go` | T04. |
| T07 | Sonnet implementation; Haiku verification | `validation.go`, `validation_test.go` | T04. |
| T08 | Sonnet integration test; Haiku complete verification | new `current_read_fixture_test.go`; fixtures read-only | T05-T07. |
| T09 | Sonnet review/correction; Haiku closure sync | accepted Task boundaries, then three closure records | T08. |

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
- implementation review reports no blocking or major findings;
- `DRMCP-REQ-MCP-001` lists this Work Item in `work_items`;
- final evidence records changed code, tests, commands, results, review verdict, and residual limitations.

## Evidence

- 2026-06-30: The implementation plan in this Work Item is retired for the `brewprint-rebuild` line. Production code produced under the old implementation plan must not be reused. Existing architecture may be consulted only as non-authoritative reference because the rebuild may adopt different application routing and structural boundaries.

- `DRMCP-ADR-MCP-001`: Accepted current-format-first implementation sequence.
- `DRMCP-INV-MCP-002`: Implementation and test drift baseline.
- `DRMCP-WORK-MCP-003` through `DRMCP-WORK-MCP-008`: Accepted contract, disposition, and fixture inputs.
- `DRMCP-TASK-MCP-001-09`: Hub lifecycle gate for this Work Item.
- `DRMCP-TASK-MCP-009-01`: Bounded contract-to-code correspondence, accepted-fixture allocation, parallel graph, model ownership, exact file boundaries, command plans, and escalation gates.
- Implementation, automated-test, and independent-review evidence: pending T02 through T09 execution.

Planning lifecycle began on 2026-06-28.
The T01 planning gate was accepted on 2026-06-28 after an independent review returned PASS with no blocking, major, or minor finding.
T02 and T03 are ready to start in parallel under the accepted P1 boundaries.
No production Go source, existing Go test, or accepted fixture was changed during Task graph authoring or closure synchronization.

### Rebuild-line disposition

Disposition: replaced and retired for the rebuild line.

- W009's former production structure, Task graph, file ownership, and extension seam are not execution authority.
- Do not resume implementation Tasks under W009.
- Do not carry production code created under W009 into the replacement implementation.
- Accepted public contracts and fixtures remain valid.
- A new implementation Work Item must consume the architecture reviewed under `DRMCP-WORK-MCP-011`.
- W009 remains `blocked` until the replacement Work Item exists.
- W009 is not `done`.
- No unsupported `canceled` status is introduced.
- W009 Task graph, Task statuses, production source, and tests remain unchanged by this closure.
