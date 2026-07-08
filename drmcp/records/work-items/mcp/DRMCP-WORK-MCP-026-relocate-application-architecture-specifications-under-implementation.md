# DRMCP-WORK-MCP-026: Relocate application architecture specifications under implementation

- **id**: DRMCP-WORK-MCP-026
- **status**: done
- **date**: 2026-07-08
- **source_refs**:
  - DRMCP-REQ-MCP-008
- **impact_refs**:
  - `spec:drmcp.application_architecture`
  - `spec:drmcp.application_architecture.application_boundary_and_components`
  - `spec:drmcp.application_architecture.dependency_and_responsibility`
  - `spec:drmcp.application_architecture.runtime_and_state`
  - `spec:drmcp.application_architecture.failure_and_evolution`
  - `spec:drmcp.implementation`
- **tasks**:
  - DRMCP-TASK-MCP-026-01
  - DRMCP-TASK-MCP-026-02
  - DRMCP-TASK-MCP-026-03
  - DRMCP-TASK-MCP-026-04
  - DRMCP-TASK-MCP-026-05
  - DRMCP-TASK-MCP-026-06
  - DRMCP-TASK-MCP-026-07

## Goal

Relocate or otherwise disposition the DRMCP application-architecture Specification tree under the DRMCP implementation Specification topology.

The relocation must preserve accepted architecture semantics while making the Specification placement match its implementation-facing role.

## Boundary

This Work Item owns:

- inventorying spec-internal references that mention or target `spec:drmcp.application_architecture`;
- limiting the initial reference search to `drmcp/records/spec/`;
- deciding the target implementation Specification topology;
- deciding canonical ref and path handling for the relocated architecture Specifications;
- deciding whether a compatibility note or transitional reference is needed;
- authoring the selected Specification moves and spec-internal reference synchronization;
- reviewing the final Specification topology and reference state.

This Work Item does not own:

- re-deciding the accepted application architecture;
- re-deciding W018 module contracts;
- changing module-contract content under `spec:drmcp.implementation.contracts` except for necessary references;
- editing ADR, Work Item, Task, Requirement, or Investigation history solely for placement cleanup;
- implementation package layout, Go symbols, tests, fixtures, or production implementation;
- implementation execution-graph planning.

## Impact Scope

Known source tree:

- `spec:drmcp.application_architecture`
- `spec:drmcp.application_architecture.application_boundary_and_components`
- `spec:drmcp.application_architecture.dependency_and_responsibility`
- `spec:drmcp.application_architecture.runtime_and_state`
- `spec:drmcp.application_architecture.failure_and_evolution`

Known target area:

- `spec:drmcp.implementation.application_architecture`

The selected physical target path is `drmcp/records/spec/implementation/application-architecture/`.

`spec:drmcp.implementation` root-role cleanup is deferred.
Spec-internal reference synchronization is in scope.
Repository-wide historical-reference cleanup is out of scope unless a later decision explicitly expands scope.
No compatibility stub remains under the old `spec:drmcp.application_architecture` tree.

## Task flow

```text
DRMCP-TASK-MCP-026-01 inventory spec references affected by application-architecture relocation
  -> DRMCP-TASK-MCP-026-02 decide relocation topology and migration route
     -> DRMCP-TASK-MCP-026-03 author relocation migration script
        -> DRMCP-TASK-MCP-026-04 review relocation migration script and dry-run
           -> DRMCP-TASK-MCP-026-05 apply application-architecture relocation
              -> DRMCP-TASK-MCP-026-06 review application-architecture relocation
                 -> DRMCP-TASK-MCP-026-07 synchronize closure
```

T01 through T07 are materialized.
T07 synchronized closure after T06 PASS.

## Task Candidates

| task | task type | responsibility | dependency |
|---|---|---|---|
| `DRMCP-TASK-MCP-026-01` | `investigation` | Inventory spec-internal references and relocation-sensitive metadata under `drmcp/records/spec/`. | none |
| `DRMCP-TASK-MCP-026-02` | `decision` | Decide target subtree, compatibility policy, ADR route, and migration route. | T01 |
| `DRMCP-TASK-MCP-026-03` | `implementation` | Author a deterministic migration script with dry-run support. | T02 |
| `DRMCP-TASK-MCP-026-04` | `review` | Independently review the migration script and dry-run output before apply mode. | T03 |
| `DRMCP-TASK-MCP-026-05` | `authoring` | Apply the reviewed migration to move files and synchronize spec refs. | T04 |
| `DRMCP-TASK-MCP-026-06` | `review` | Independently review the relocated Specification tree and reference synchronization. | T05 |
| `DRMCP-TASK-MCP-026-07` | `synchronization` | Synchronize W026 closure after review pass or closed findings. | T06 |

## Completion Condition

- Spec references affected by moving the application-architecture tree are inventoried.
- The target implementation Specification topology is decided.
- Required Specification moves or rewrites are authored.
- Spec-internal canonical refs, parent refs, topic tables, related-spec refs, and prose references are synchronized.
- Accepted architecture semantics are preserved.
- W018 module-contract semantics are preserved.
- One integrated independent review passes or every closure-blocking finding is independently closed.
- Closure state is synchronized.
- No production implementation planning is released by this Work Item.

## Evidence

- DRMCP-REQ-MCP-008 captures the need to relocate or disposition the current application-architecture Specification tree under the implementation Specification topology.
- W018 and ADR-013 intentionally left application-architecture relocation outside the module-contract decision.
- T01 is created as the initial bounded spec-reference inventory before topology or authoring decisions.
- T01 completed the spec-only reference inventory without creating a separate Investigation record.
- T02 selected `spec:drmcp.implementation.application_architecture` as the target subtree.
- T02 selected `drmcp/records/spec/implementation/application-architecture/` as the physical target path.
- T02 deferred `spec:drmcp.implementation` root-role cleanup.
- T02 selected no compatibility stub for old application-architecture refs.
- T02 selected no ADR for this relocation.
- T02 selected a migration-script review route before apply-mode migration.
- T03 authored `drmcp/scripts/relocate_application_architecture_specs.py`.
- T03 dry-run returned five planned moves, 12 planned in-place rewrite files, and 57 total ref replacements.
- T04 review returned `PASS` and allowed T05 apply.
- T05 applied the reviewed migration and reported five moved files and 12 rewritten files.
- T05 post-apply old-ref check reported no `spec:drmcp.application_architecture` matches under `drmcp/records/spec/`.
- T06 independent relocation review returned `PASS` with no findings.
- T07 synchronized W026 closure.
