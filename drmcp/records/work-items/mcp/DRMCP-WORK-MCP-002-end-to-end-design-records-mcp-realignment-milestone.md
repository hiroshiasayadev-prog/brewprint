# DRMCP-WORK-MCP-002: End-to-end Design Records MCP realignment milestone

- **id**: DRMCP-WORK-MCP-002
- **status**: in_progress
- **date**: 2026-06-26
- **source_requirement**: DRMCP-REQ-MCP-004
- **impact_refs**:
  - DRMCP-ADR-MCP-001
  - DRMCP-REQ-MCP-001
  - DRMCP-REQ-MCP-002
  - DRMCP-REQ-MCP-003
  - PRODUCT-REQ-SPEC-003
  - DRMCP-WORK-MCP-001
  - PRODUCT-WORK-SPEC-013
  - DRMCP-WORK-SPEC-001
  - DRMCP-WORK-SPEC-002
  - spec:product.brewprint.compatibility.legacy_id_compatibility
  - spec:product.design_records.spec_format.validation_policy
- **tasks**:
  - DRMCP-TASK-MCP-002-02

## Goal

Coordinate the full Design Records MCP realignment from child Work Item creation through child Work Item completion.

Make the dependency graph, hard gates, parallel work, and integrated milestone closure visible while leaving detailed contracts, tasks, implementation, and evidence inside the source-specific child Work Items.

## Boundary

This Work Item owns:

- milestone-level sequencing across child Work Items;
- confirmation or creation of each required child Work Item;
- verification that each child Work Item uses the correct source Requirement;
- coarse lifecycle tracking from child Work Item opening through `done`;
- cross-Work-Item gate acceptance;
- identification of additional REQ-003 follow-up Work Items beyond P0 when required;
- integrated milestone review and closure evidence.

This Work Item does not own:

- read-contract correction or implementation details;
- portable-package projection, generation, or release details;
- package loader or load-time validation details;
- guidance-projection or authoring-integration details;
- authoring transaction implementation details;
- child task status replication;
- child Requirement closure decisions;
- changes to a child Work Item's source Requirement.

A milestone task closes only when its named child Work Item is `done`, or when an explicit accepted disposition states that the child Work Item is no longer required.
The milestone records only the gate result and evidence pointer, not the child's internal execution log.

Future topic sessions use this Work Item as the stable entrypoint, then read the selected child Work Item and only its unfinished or directly relevant Tasks.
Child Work Item and Task metadata remain the canonical status sources; the milestone must not duplicate their live progress.

## Impact Scope

| ref or workstream | milestone impact |
|---|---|
| `DRMCP-WORK-MCP-001` | Existing child Work Item for the current read baseline, legacy fallback, fixtures, implementation, and review. |
| `PRODUCT-WORK-SPEC-013` | Existing child Work Item for portable package production, drift detection, release validation, and producer handoff. |
| `DRMCP-REQ-MCP-003` | Requires a P0 child Work Item for package consumer contract, loader, configuration, load-time validation, and minimum portability fixtures. |
| `DRMCP-REQ-MCP-003` later phases | May require one or more child Work Items for guidance projection, authoring integration, proposal reproducibility, and end-to-end package capability exposure. |
| `DRMCP-REQ-MCP-002` | Requires a child Work Item for workflow artifact authoring, spec authoring, investigation authoring, fixtures, implementation, and review. |
| `DRMCP-WORK-SPEC-001/002` | Existing validation Work Items whose disposition is gated by PRODUCT validation-policy owner pointers. |
| Brewprint legacy compatibility | Must reject `V01-SPEC-*` before legacy fallback implementation is accepted. |
| Integrated release surface | Must prove that read, package, and authoring capabilities form one coherent portable runtime. |

## Task flow

```text
A. Complete the current read-baseline child Work Item
   DRMCP-WORK-MCP-001
           ║
           ║ may proceed in parallel
           ║
B. Complete the portable package producer child Work Item
   PRODUCT-WORK-SPEC-013
           │
           ├── producer contract and reviewable package fixture
           ▼
C. Create and complete the DRMCP package-consumer P0 child Work Item
   source: DRMCP-REQ-MCP-003
           │
           ├── loader, configuration, load-time validation
           ▼
D. Create and complete later DRMCP package-integration child Work Item(s), when required
   source: DRMCP-REQ-MCP-003
           │
           ▼
E. Create and complete the package-backed authoring child Work Item
   source: DRMCP-REQ-MCP-002
           │
           ▼
F. Run integrated milestone review and close
```

Dependency rules:

- A and B may execute in parallel.
- C planning may begin after the producer interface is reviewable; C acceptance requires a reviewable or released package fixture.
- D exists only for REQ-003 scope not completed by P0. T04 decides whether follow-up Work Items are required; each created Work Item receives a separate milestone lifecycle Task rather than being bundled into T04.
- E contract design may begin earlier, but E runtime implementation and acceptance require C completion and the package mappings needed by the affected authoring phase.
- F requires every required child Work Item to be `done` and every accepted deferral or residual limitation to be explicit.

## Task Candidates

| candidate | milestone task scope | completion gate |
|---|---|---|
| T01 | Track `DRMCP-WORK-MCP-001` from its current state through completion. All detailed read-contract, fixture, implementation, and review work remains inside that Work Item. | `DRMCP-WORK-MCP-001` is `done`; REQ-001 completion evidence and independent review are recorded. |
| T02 | Track `PRODUCT-WORK-SPEC-013` from its current state through package release and completion. All package projection, generation, drift, fixture, and release details remain inside that Work Item. | `PRODUCT-WORK-SPEC-013` is `done`; package contract, generated artifact, release validation, and producer handoff evidence are recorded. |
| T03 | Verify `DRMCP-WORK-MCP-001` T11 readiness, create the P0 Work Item sourced from `DRMCP-REQ-MCP-003`, verify its consumer-only boundary, and follow it through completion. Creation is a discrete initial step; the Task completes only when the child Work Item is `done`. | P0 Work Item is `done`; package loader, configuration, load-time validation, minimum indexing, and portability fixtures are accepted. |
| T04 | Evaluate REQ-003 P0 evidence and decide whether additional REQ-003 Work Items are required. Create each required Work Item and add a separate milestone lifecycle Task for it; otherwise record an explicit accepted no-follow-up disposition. | The post-P0 Work Item decision is recorded, every required Work Item exists with a corresponding milestone lifecycle Task, or an accepted disposition proves none is required. |
| T05 | Verify `DRMCP-WORK-MCP-001` T12 readiness, create the Work Item sourced from `DRMCP-REQ-MCP-002`, verify that its runtime tasks are package-backed, and follow it through workflow, spec, and investigation authoring completion. Creation is a discrete initial step; the Task completes only when the child Work Item is `done`. | REQ-002 Work Item is `done`; workflow, spec, and investigation authoring phases and their fixtures/reviews are accepted. |
| T06 | Run an integrated milestone review across all completed child Work Items, apply or route required findings to the owning Work Item, and close the milestone. | No blocking or major cross-workstream findings remain; all child Work Items and milestone gates are reflected in final evidence. |

These candidate tasks intentionally use a larger granularity than implementation tasks.
They exceed the canonical 0.5–3 day Task scope by design because each tracks a child Work Item lifecycle.
They must still use the canonical Task structure and one observable Done condition without splitting into per-day units.
For T03 and T05, child Work Item creation is an initial Work step; the sole Done condition is completion of the named child Work Item.
Each candidate tracks a complete child Work Item lifecycle rather than reproducing that Work Item's internal tasks.
Any post-P0 Work Item created by T04 receives its own milestone lifecycle Task.

## Completion Condition

This Work Item is complete when all of the following are true:

- `DRMCP-WORK-MCP-001` is `done` and resolves the accepted current read baseline;
- `PRODUCT-WORK-SPEC-013` is `done` and provides a released, self-contained portable standards package with producer evidence;
- a P0 Work Item sourced from `DRMCP-REQ-MCP-003` is `done` and provides package loading, configuration, load-time validation, minimum package indexing, and portability fixtures;
- all additional Work Items required under `DRMCP-REQ-MCP-003` are `done`, or their absence has an explicit accepted disposition;
- a Work Item sourced from `DRMCP-REQ-MCP-002` is `done` and provides package-backed workflow, spec, and investigation authoring;
- legacy fallback accepts only the approved exact V01 sequential families and does not expose legacy specs, fuzzy repair, or automatic archive discovery;
- existing validation Work Item disposition and PRODUCT validation-policy owner pointers are consistent;
- integrated review confirms that `PRODUCT-WORK-SPEC-013` release-validation evidence preserves PRODUCT ownership and excludes Brewprint compatibility, DRMCP API, and BPDSL runtime authority;
- DRMCP package and authoring runtime do not depend on the host repository's `product` namespace or physical Brewprint layout;
- authoring runtime consumes validated package mappings and does not hard-code PRODUCT semantics;
- path hiding, canonical identity, proposal, validation, and write boundaries are coherent across the integrated surface;
- an independent integrated review reports no blocking or major findings;
- every required child Work Item ID, source Requirement, final status, review result, and completion evidence pointer is recorded here;
- residual limitations and explicitly deferred capabilities are listed without being claimed complete;
- `DRMCP-REQ-MCP-004` continues to list this Work Item in `work_items`.

## Evidence

- `DRMCP-ADR-MCP-001`: Accepted realignment authority and phased delivery direction.
- `DRMCP-REQ-MCP-001`: Current read-baseline requirement.
- `PRODUCT-REQ-SPEC-003`: Portable package producer requirement.
- `DRMCP-REQ-MCP-003`: Portable package consumer and integration requirement.
- `DRMCP-REQ-MCP-002`: Package-backed authoring transaction requirement.
- `DRMCP-WORK-MCP-001`: Existing read-baseline child Work Item.
- `PRODUCT-WORK-SPEC-013`: Existing package-producer child Work Item.
- Future REQ-003 and REQ-002 child Work Item IDs, completion evidence, integrated review, and milestone closure evidence: pending execution.
- `PRODUCT-TASK-SPEC-013-08` completed the producer handoff for this milestone. The independent re-review returned `PASS`; T08 F-MIN-01 is `CLOSED`; no blocking, major, or minor findings remain; and `PRODUCT-WORK-SPEC-013` is `done`.
- Producer authority: `PRODUCT-REQ-SPEC-003` and `PRODUCT-WORK-SPEC-013`.
- Accepted producer roots: source `product/records/spec/design-records/`, generated location `bin/design-records/`, logical root `design-records/`, package root `spec:design_records`, and guidance root `spec:design_records.authoring_standards`.
- Accepted transformation: rewrite only canonical refs under `spec:product.design_records` into `spec:design_records`; preserve ordinary prose, public IDs, physical paths, external refs, noncanonical lookalikes, and binary bytes.
- Accepted producer behavior: whole-tree rebuild, warning-only exit `0`, operational-failure exit `1`, derived non-authoritative output, and no process-current-directory, host registry, DRMCP runtime, network, or third-party dependency.
- Accepted warning boundary: duplicate canonical ref, unresolved internal ref, external canonical ref, unrewritten source ref, and source-authoring boundary finding are non-blocking. The producer does not repair, filter, delete, generalize, or reinterpret those findings.
- Accepted release evidence: 34 generated files, 79 non-blocking warnings, generator exit `0`, 35 passing tests with exit `0`, `scripts\\verify.bat` exit `0`, generated-artifact ignore exit `0`, scoped whitespace pass, and T07 independent re-review `PASS`.
- Known producer residuals: no manifest, version negotiation, remote registry, incremental generation, or concurrent-generation support. Generated output remains non-authoritative, and source correction may require separate ownership.
- Consumer-owned remaining scope: package-root configuration, loader and availability behavior, recursive discovery, localized indexes and diagnostics, guidance projection, authoring and validation integration, proposal reproducibility, and package-dependent capability exposure.
- No P0 Work Item sourced from `DRMCP-REQ-MCP-003` exists in the scoped consumer Work Item set. Producer closure is accepted. The next consumer-side action is to create that P0 Work Item through milestone T03. This milestone does not treat producer closure as consumer implementation completion.

### T02 producer package gate acceptance

`DRMCP-TASK-MCP-002-02` formally accepts the completed PRODUCT producer handoff as the milestone producer gate.

| evidence | accepted result |
|---|---|
| `PRODUCT-WORK-SPEC-013` | `done` |
| `PRODUCT-TASK-SPEC-013-08` | `done` |
| producer final re-review | `PASS` |
| T08 F-MIN-01 | `CLOSED` |
| blocking / major / minor findings | none |
| producer gate | accepted |

Release-evidence pointer:

- 34 generated files;
- 79 non-blocking semantic warnings;
- generator exit `0`;
- 35 tests passed with test exit `0`;
- `scripts\\verify.bat` exit `0`;
- generated-artifact ignore exit `0`;
- scoped whitespace pass.

The 79 warning messages remain in producer execution output and are not copied into this milestone.

Producer gate acceptance means the PRODUCT-owned contract, generated package, and release evidence are available to the milestone.
The acceptance also means the producer input for future T03 execution is ready.

Producer gate acceptance does not mean a P0 Work Item exists or any consumer loader, configuration, indexing, guidance, validation, or authoring integration is implemented.
Consumer runtime completion and milestone completion remain open.

The next gate remains the read-baseline readiness flow:

| record | current status |
|---|---|
| `DRMCP-WORK-MCP-001` | `in_progress` |
| `DRMCP-TASK-MCP-001-11` | `not_started` |
| `DRMCP-TASK-MCP-001-12` | `not_started` |

While this state remains, T03 must not execute and no Work Item sourced from `DRMCP-REQ-MCP-003` may be created or assigned an inferred ID.
The next milestone action is T03 execution after T11 and T12 readiness is accepted.
The independent T02 review returned `PASS` with no blocking, major, or minor findings. The producer package gate remains accepted.
This Work Item remains `in_progress` because later milestone Tasks remain.
