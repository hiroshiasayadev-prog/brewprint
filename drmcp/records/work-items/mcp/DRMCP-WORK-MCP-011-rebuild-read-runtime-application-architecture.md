# DRMCP-WORK-MCP-011: Rebuild read-runtime application architecture

- **id**: DRMCP-WORK-MCP-011
- **status**: done
- **date**: 2026-06-30
- **source_requirement**: DRMCP-REQ-MCP-001
- **impact_refs**:
  - DRMCP-ADR-MCP-001
  - DRMCP-ADR-MCP-002
  - DRMCP-ADR-MCP-003
  - DRMCP-ADR-MCP-004
  - DRMCP-ADR-MCP-005
  - DRMCP-ADR-MCP-006
  - DRMCP-WORK-MCP-001
  - DRMCP-WORK-MCP-002
  - DRMCP-WORK-MCP-003
  - DRMCP-WORK-MCP-004
  - DRMCP-WORK-MCP-005
  - DRMCP-WORK-MCP-006
  - DRMCP-WORK-MCP-007
  - DRMCP-WORK-MCP-008
  - DRMCP-WORK-MCP-009
  - DRMCP-WORK-MCP-010
  - DRMCP-WORK-SPEC-001
  - DRMCP-WORK-SPEC-002
  - spec:drmcp.implementation
  - spec:drmcp.design_records_mcp.overview
  - spec:drmcp.design_records_mcp.responsibility_boundary
  - spec:drmcp.design_records_mcp.mvp_scope
  - spec:drmcp.design_records_mcp.namespace_scanning
  - spec:drmcp.design_records_mcp.resolver
  - spec:drmcp.design_records_mcp.schema.overview
  - spec:drmcp.design_records_mcp.schema.diagnostics
  - spec:drmcp.design_records_mcp.tools.overview
- **tasks**:
  - DRMCP-TASK-MCP-011-01
  - DRMCP-TASK-MCP-011-02
  - DRMCP-TASK-MCP-011-03
  - DRMCP-TASK-MCP-011-04
  - DRMCP-TASK-MCP-011-05
  - DRMCP-TASK-MCP-011-06
  - DRMCP-TASK-MCP-011-07

## Goal

Define and review the application architecture for a clean implementation of the `DRMCP-REQ-MCP-001` read runtime.

Preserve accepted read contracts and fixtures without inheriting the retired W009 production structure or implementation graph.

## Boundary

This Work Item owns:

- the decision inventory for the rebuild read-runtime architecture;
- MCP transport and tool-routing boundaries for `list_records`, `get_records`, `resolve_reference`, and `validate_records`;
- application use-case boundaries and orchestration ownership for those four operations;
- composition-root and runtime-lifecycle ownership for those four operations;
- domain, parsed-source, index-state, diagnostic, and public-projection separation;
- ports and adapters for configuration, filesystem access, indexing, and source loading;
- current-index and optional legacy-index runtime composition;
- validation detector and graph-validation integration boundaries;
- application error and diagnostic mapping to MCP operation responses;
- Go package dependency direction and package responsibility boundaries;
- ADR routing for durable architecture decisions;
- synchronization of accepted architecture decisions into current DRMCP Specifications;
- independent design review, finding correction, re-review, and design closure.

This Work Item does not own:

- production implementation;
- reuse or correction of production code created under W009;
- the retired W009 Task graph, file ownership, model routing, or extension seams;
- legacy fallback implementation under W010;
- implementation Tasks for WORK-SPEC-001 or WORK-SPEC-002;
- authoring-guidance application, snapshot, and package architecture owned by its contracts;
- authoring transaction runtime and current-format integration owned by `DRMCP-REQ-MCP-002`;
- portable standards-package consumption owned by `DRMCP-REQ-MCP-003`;
- changes to accepted read-operation behavior from W003 through W006;
- fixture redesign unless an architecture decision exposes a direct fixture-contract defect;
- W001 or W002 lifecycle synchronization before design closure.

Accepted contracts constrain externally observable behavior.
They do not prescribe the internal application structure selected by this Work Item.

## Impact Scope

| ref or area | impact |
|---|---|
| `DRMCP-REQ-MCP-001` | Source Requirement for the read runtime. |
| `DRMCP-ADR-MCP-001` | Supplies current-first authority and complete D-005 coverage for current and legacy separation. |
| `DRMCP-ADR-MCP-002` through `DRMCP-ADR-MCP-006` | Supply accepted runtime lifecycle, layered architecture, internal state and operation-contract, validation-orchestration, and Go package decisions. |
| `DRMCP-WORK-MCP-003` through `DRMCP-WORK-MCP-006` | Supply accepted discovery, read-operation, resolver, validation, diagnostic, and path-exposure contracts. |
| `DRMCP-WORK-MCP-007` | Supplies the retained ownership split for per-file and Topics graph validation. |
| `DRMCP-WORK-MCP-008` | Supplies accepted fixture meanings and coverage. Fixture placement remains non-authoritative. |
| `DRMCP-WORK-MCP-009` | Retired implementation plan. No production structure or Task graph is inherited. |
| `DRMCP-WORK-MCP-010` | Remains blocked from execution until the replacement current-runtime architecture is accepted. |
| `DRMCP-WORK-SPEC-001/002` | Retain semantic ownership but require implementation-boundary rebaselining after architecture closure. |
| DRMCP overview and responsibility specs | Receive the accepted internal architecture and ownership boundary where the current contract lacks it. |
| Future implementation Work Item | Consumes only reviewed architecture and accepted contracts. |

## Task flow

| phase | task or owner | dependency | outcome |
|---|---|---|---|
| A. Decision inventory and interactive loop | `DRMCP-TASK-MCP-011-01` | Accepted contracts and rebuild direction | `done`: D-001 through D-009 decided and persisted. |
| B. ADR routing and authoring | `DRMCP-TASK-MCP-011-02` | Phase A complete | Accepted ADR coverage for D-001 through D-009; Specifications remain unchanged. |
| C. Specification synchronization | `DRMCP-TASK-MCP-011-03` | Phase B | `done`: accepted implementation architecture and affected DRMCP Specifications synchronized. |
| D. Independent design review | `DRMCP-TASK-MCP-011-04` | Phase C | `done: NEEDS REVISION`. |
| E. Finding correction and closure re-review | `DRMCP-TASK-MCP-011-05`, then `DRMCP-TASK-MCP-011-06` | Phase D `NEEDS REVISION` | `done`: T06 `PASS`; F-MAJ-01 and F-MAJ-02 `CLOSED`. |
| F. Design closure synchronization | `DRMCP-TASK-MCP-011-07` | Phase E `PASS` | `done`: design closure and downstream disposition synchronized. |

Production implementation planning begins only after Phase F.

## Task Candidates

| candidate | scope | dependency |
|---|---|---|
| T01 | Maintain the architecture decision inventory and ask one decision at a time. | None. |
| T02 | Route decided items to existing, amended, superseding, or new ADRs. | T01 decision complete. |
| T03 | Synchronize accepted architecture into current DRMCP Specifications. | T02. |
| T04 | Perform independent design review. | T03. |
| T05 | Correct only named review findings. | T04 `NEEDS REVISION`. |
| T06 | Independently re-review finding closure. | T05. |
| T07 | Synchronize final design closure and downstream Work Item dispositions. | T04 `PASS` or T06 closure. |

Only T01 exists at Work Item creation.
Later Tasks are created after their exact scope and writer ownership become executable.

## Completion Condition

This Work Item is complete when all of the following are true:

- every required architecture decision is `recorded`, `deferred`, or explicitly blocked with a destination;
- tool routing and application use-case ownership are explicit for `list_records`, `get_records`, `resolve_reference`, and `validate_records`;
- composition-root and runtime-lifecycle ownership are explicit;
- domain, parsed-source, index, diagnostic, and public-projection boundaries are explicit;
- configuration, filesystem, indexing, and source loading use an accepted dependency direction;
- current and legacy index composition preserves accepted operational separation;
- per-file and Topics graph validation integration preserves W007 ownership;
- MCP transport types do not become accidental domain authority;
- durable architecture choices have accepted ADR coverage where required;
- current DRMCP Specifications state the accepted architecture without contradicting W003 through W006;
- independent design review reports `PASS` with no open blocking or major finding;
- W009 and W010 receive an explicit replace, rebaseline, cancel, or supersession disposition;
- W001, W002, WORK-SPEC-001, and WORK-SPEC-002 receive explicit downstream handling;
- no implementation Task depends on an unresolved architecture decision;
- completion Evidence identifies the accepted ADRs, Specifications, review result, and deferred scope.

## Evidence

- 2026-06-30 rebuild direction: production code created under the retired W009 plan is not reused.
- Existing architecture is non-authoritative reference material only.
- `DRMCP-ADR-MCP-001` and `DRMCP-REQ-MCP-001` remain accepted semantic inputs.
- W003 through W007 remain accepted contract and ownership inputs.
- W008 remains an accepted fixture-meaning input with non-authoritative physical placement and an unexecuted final structural verifier limitation.
- W009 is blocked and its implementation graph is retired.
- W010 was `not_started` before closure and is now `blocked` pending implementation-graph rebaseline.
- `DRMCP-TASK-MCP-011-01` completed the architecture decision loop with D-001 through D-009 decided.
- Accepted W011 architecture for the four operations now covers transport boundaries, runtime lifecycle, internal model separation, ports and adapters, current/legacy composition, validation orchestration, overall component structure, operation contract ownership, and concrete Go package placement.
- `DRMCP-TASK-MCP-011-02` routed every accepted decision to durable ADR authority.
- D-005 remains covered by `DRMCP-ADR-MCP-001`; no amendment or supersession was required.
- `DRMCP-ADR-MCP-002` through `DRMCP-ADR-MCP-006` record runtime lifecycle, layered application and adapter boundaries, internal state and operation-contract separation, validation orchestration, and concrete Go package boundaries.
- Phase B is complete.
- `DRMCP-TASK-MCP-011-03` completed Phase C Specification synchronization.
- `spec:drmcp.implementation` is the current normative implementation-architecture authority for runtime lifecycle, layers, ports, state separation, validation orchestration, and Go package layout for the four W011 read-runtime operations.
- Existing Design Records MCP Specifications retain authority for public request, response, status, warning, diagnostic, and error behavior.
- D-005 remains governed by `DRMCP-ADR-MCP-001`, namespace scanning, and resolver contracts without behavior change.
- `DRMCP-ADR-MCP-002` through `DRMCP-ADR-MCP-006` record `migrated_to_spec: 2026-06-30`.
- Phase C changed one new implementation Specification and ten existing Specifications. No Requirement, fixture, production source, implementation Task, review Task, or downstream lifecycle record changed.
- T04 independent design review completed.
- Verdict: NEEDS REVISION.
- Blocking findings: none.
- Major findings: F-MAJ-01 and F-MAJ-02.
- Minor findings: none.
- Advisories: none.
- Implementation planning remains blocked pending correction and independent finding-closure re-review.
- `DRMCP-TASK-MCP-011-05` corrected F-MAJ-01 and F-MAJ-02 only.
- W011 architecture is now limited to `list_records`, `get_records`, `resolve_reference`, and `validate_records`.
- The package tree is complete for that W011 read-runtime slice, not for the complete Design Records MCP server.
- Authoring-guidance and authoring-transaction architecture remain outside W011.
- D-006's persisted-state freshness rule is preserved as a future integration constraint.
- Current authoring-transaction integration is deferred to `DRMCP-REQ-MCP-002`.
- The correction author does not claim either finding is closed.
- T06 independent finding-closure re-review completed with verdict `PASS`.
- F-MAJ-01 and F-MAJ-02 are `CLOSED`.
- No direct regression finding was reported.
- No blocking, major, or required minor finding remains.
- Implementation-planning readiness is `READY FOR DESIGN CLOSURE`.
- DRMCP authoring transactions are non-operational. Filesystem authoring is used under the current agent-authoring policy.

### Review closure

- T06 independent re-review verdict: `PASS`.
- F-MAJ-01: `CLOSED`.
- F-MAJ-02: `CLOSED`.
- Direct regression findings: none.
- Blocking, major, and required minor findings: none.
- Implementation-planning readiness: `READY FOR DESIGN CLOSURE`.

### Canonical architecture

- D-001 through D-009 are `recorded`.
- D-005 remains covered by `DRMCP-ADR-MCP-001`.
- D-001, D-004, and D-007 are governed by `DRMCP-ADR-MCP-003`.
- D-002 is governed by `DRMCP-ADR-MCP-002`.
- D-003 and D-008 are governed by `DRMCP-ADR-MCP-004`.
- D-006 is governed by `DRMCP-ADR-MCP-005`.
- D-009 is governed by `DRMCP-ADR-MCP-006`.
- `spec:drmcp.implementation` is the implementation-architecture authority for the four W011 read-runtime operations.
- Public operation behavior remains owned by Design Records MCP child Specifications.
- Current authoring-transaction integration remains deferred to `DRMCP-REQ-MCP-002`.

### Downstream handling

- W009: replaced and retired for the rebuild line; status remains `blocked`.
- W010: retained scope; status `blocked` pending implementation-graph rebaseline.
- W001: remains `in_progress`; requires a dedicated graph amendment.
- W002: remains `in_progress`; no consumer-side gate is released.
- W-SPEC-001: retained and `not_started`.
- W-SPEC-002: retained and `not_started`.

### Retained validation Work Items

```text
DRMCP-WORK-SPEC-001:
retained, not_started, per-file validation implementation owner.

DRMCP-WORK-SPEC-002:
retained, not_started, Topics graph validation implementation owner.
```

- W007's accepted `retain` disposition remains unchanged.
- W-SPEC-001 and W-SPEC-002 are not absorbed into W009.
- Both Work Items must be sequenced in the replacement implementation graph that consumes W011.
- Exact implementation Task authoring begins only after the W001 graph amendment.
- This closure does not change their status, Task graph, or scope.

### Completion assessment

| condition | result |
|---|---|
| All decisions are `recorded`. | pass |
| ADR coverage is accepted. | pass |
| Specification reflection is complete. | pass |
| T06 reports `PASS`. | pass |
| No blocking or major finding remains. | pass |
| W009 and W010 disposition is explicit. | pass |
| W001 and W002 handling is explicit. | pass |
| W-SPEC-001 and W-SPEC-002 handling is explicit. | pass |
| No implementation Task depends on an unresolved W011 architecture decision. | pass |
| Deferred authoring scope has an exact owner. | pass: `DRMCP-REQ-MCP-002` |
| No production implementation began. | pass |

All W011 completion conditions are satisfied.
W011 is `done`.
The exact next gate is a dedicated W001 graph-amendment Task.
