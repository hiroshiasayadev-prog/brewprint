# DRMCP-TASK-MCP-026-01: Inventory spec references for application-architecture relocation

- **id**: DRMCP-TASK-MCP-026-01
- **status**: done
- **date**: 2026-07-08
- **work_item**: DRMCP-WORK-MCP-026
- **task_type**: investigation
- **estimate**: 0.5d
- **depends_on**: []
- **outputs**:
  - DRMCP-TASK-MCP-026-01

## Goal

Inventory spec-internal references affected by relocating the DRMCP application-architecture Specification tree under the implementation Specification topology.

## Work

Before starting, read:

1. `prompt_chappy.md`
2. `product/records/spec/design-records/authoring-standards/writing-standard.md`
3. `product/records/spec/design-records/authoring-standards/task-authoring.md`
4. `skills/design-convergence-workflow/SKILL.md`
5. `skills/design-convergence-workflow/impact-investigation.md`
6. `drmcp/records/requirements/mcp/DRMCP-REQ-MCP-008-relocate-application-architecture-specifications-under-implementation.md`
7. `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-026-relocate-application-architecture-specifications-under-implementation.md`
8. `drmcp/records/spec/application-architecture/index.md`
9. `drmcp/records/spec/implementation/index.md`

Then search only under:

- `drmcp/records/spec/`

Inventory these relocation-sensitive references:

| class | include |
|---|---|
| Canonical spec refs | Backtick refs and metadata values that start with `spec:drmcp.application_architecture`. |
| Parent refs | H1-adjacent `parent` values that identify the old application-architecture tree. |
| Topic tables | `## Topics` rows that point at the old application-architecture refs. |
| Related specs | Related-spec rows that point at the old application-architecture refs. |
| Prose references | Normative prose that names the old application-architecture ref or old physical tree. |
| Mermaid or diagrams | Diagram labels or links only when the old placement or ref is part of the diagram meaning. |
| Physical paths | Mentions of `drmcp/records/spec/application-architecture/` or child files. |

Classify each affected reference with this table shape:

| source file | line or section | current reference | reference class | likely action | notes |
|---|---|---|---|---|

Use `likely action` values:

- `must_update_after_target_decision`
- `maybe_update_after_target_decision`
- `retain_as_historical_context`
- `no_change`

Do not decide the target topology in this Task.
Do not move files in this Task.
Do not edit Specifications in this Task.
Do not search ADR, Requirement, Work Item, Task, or Investigation records unless the Task records a stop reason and asks for scope expansion.

## Done condition

- The spec-only reference inventory is complete for `drmcp/records/spec/`.
- Each affected reference is classified with a likely action.
- The Task identifies the minimum next decision boundary for target topology and compatibility handling.
- No canonical Specification authoring or file move is performed.

## Verification

- Required startup files were read.
- Search scope stayed inside `drmcp/records/spec/`.
- No Investigation record was created.
- No Specification was edited.
- No file move was performed.
- No ADR, Requirement, Work Item, or Task history cleanup was performed.

## Evidence

User judgment on 2026-07-08 selected the investigation-Task lightweight Evidence exception.
A separate Investigation record is disproportionate for this bounded spec-reference inventory.

### Search scope

Only `drmcp/records/spec/` was searched.
No scope expansion was performed.

Search terms and results:

| search | result |
|---|---|
| `spec:drmcp.application_architecture` | 57 matches in 17 files. |
| `drmcp/records/spec/application-architecture/` | 0 matches. |
| `application-architecture`, `application architecture`, `application_architecture` | 65 matches in 18 files. |

### Reference inventory

| source file | line or section | current reference | reference class | likely action | notes |
|---|---|---|---|---|---|
| `drmcp/records/spec/application-architecture/index.md` | metadata `id` | `spec:drmcp.application_architecture` | Canonical spec ref | `must_update_after_target_decision` | Root ref changes if the tree moves under `implementation`. |
| `drmcp/records/spec/application-architecture/index.md` | `## Topics`, lines 66-69 | Four child refs under `spec:drmcp.application_architecture.*` | Topic tables | `must_update_after_target_decision` | Child refs must match the selected target subtree. |
| `drmcp/records/spec/application-architecture/index.md` | `## Boundary`, line 75 | `spec:drmcp.application_architecture.failure_and_evolution` | Prose reference | `must_update_after_target_decision` | Architecture-return pointer must follow the relocated failure view. |
| `drmcp/records/spec/application-architecture/index.md` | `## Related specs`, lines 81-84 | Four child refs under `spec:drmcp.application_architecture.*` | Related specs | `must_update_after_target_decision` | Related-spec rows must match the selected target subtree. |
| `drmcp/records/spec/application-architecture/index.md` | title and topic map prose | `DRMCP application architecture`; `application architecture` | Prose reference | `maybe_update_after_target_decision` | Concept label can remain if topology decision keeps the term as a subtopic name. |
| `drmcp/records/spec/application-architecture/application-boundary-and-components.md` | metadata `id`, `parent`, lines 3 and 6 | `spec:drmcp.application_architecture.application_boundary_and_components`; `spec:drmcp.application_architecture` | Canonical spec refs; parent refs | `must_update_after_target_decision` | Child metadata must follow relocated parent and child refs. |
| `drmcp/records/spec/application-architecture/application-boundary-and-components.md` | line 78 | `spec:drmcp.application_architecture.dependency_and_responsibility` | Prose reference | `must_update_after_target_decision` | Cross-view pointer must follow relocated dependency view. |
| `drmcp/records/spec/application-architecture/application-boundary-and-components.md` | `## Related specs`, lines 117-120 | Parent and three sibling refs under `spec:drmcp.application_architecture.*` | Related specs | `must_update_after_target_decision` | Related-spec rows must follow relocated root and siblings. |
| `drmcp/records/spec/application-architecture/dependency-and-responsibility.md` | metadata `id`, `parent`, lines 3 and 6 | `spec:drmcp.application_architecture.dependency_and_responsibility`; `spec:drmcp.application_architecture` | Canonical spec refs; parent refs | `must_update_after_target_decision` | Child metadata must follow relocated parent and child refs. |
| `drmcp/records/spec/application-architecture/dependency-and-responsibility.md` | lines 48, 59, 108, 115 | Sibling refs under `spec:drmcp.application_architecture.*` | Prose references | `must_update_after_target_decision` | Cross-view pointers must follow relocated sibling refs. |
| `drmcp/records/spec/application-architecture/dependency-and-responsibility.md` | `## Related specs`, lines 129-132 | Parent and three sibling refs under `spec:drmcp.application_architecture.*` | Related specs | `must_update_after_target_decision` | Related-spec rows must follow relocated root and siblings. |
| `drmcp/records/spec/application-architecture/runtime-and-state.md` | metadata `id`, `parent`, lines 3 and 6 | `spec:drmcp.application_architecture.runtime_and_state`; `spec:drmcp.application_architecture` | Canonical spec refs; parent refs | `must_update_after_target_decision` | Child metadata must follow relocated parent and child refs. |
| `drmcp/records/spec/application-architecture/runtime-and-state.md` | lines 27 and 116 | Sibling refs under `spec:drmcp.application_architecture.*` | Prose references | `must_update_after_target_decision` | Cross-view pointers must follow relocated sibling refs. |
| `drmcp/records/spec/application-architecture/runtime-and-state.md` | lines 123-124 | `application-architecture decision work` | Prose reference | `maybe_update_after_target_decision` | The meaning may remain architecture-return work even after path relocation. |
| `drmcp/records/spec/application-architecture/runtime-and-state.md` | `## Related specs`, lines 162-165 | Parent and three sibling refs under `spec:drmcp.application_architecture.*` | Related specs | `must_update_after_target_decision` | Related-spec rows must follow relocated root and siblings. |
| `drmcp/records/spec/application-architecture/failure-and-evolution.md` | metadata `id`, `parent`, lines 3 and 6 | `spec:drmcp.application_architecture.failure_and_evolution`; `spec:drmcp.application_architecture` | Canonical spec refs; parent refs | `must_update_after_target_decision` | Child metadata must follow relocated parent and child refs. |
| `drmcp/records/spec/application-architecture/failure-and-evolution.md` | line 50 | `application-architecture decision` | Prose reference | `maybe_update_after_target_decision` | The decision-route name may remain even after path relocation. |
| `drmcp/records/spec/application-architecture/failure-and-evolution.md` | line 63 | `spec:drmcp.application_architecture.runtime_and_state` | Prose reference | `must_update_after_target_decision` | Cross-view pointer must follow relocated runtime view. |
| `drmcp/records/spec/application-architecture/failure-and-evolution.md` | `## Related specs`, lines 86-89 | Parent and three sibling refs under `spec:drmcp.application_architecture.*` | Related specs | `must_update_after_target_decision` | Related-spec rows must follow relocated root and siblings. |
| `drmcp/records/spec/design-records-mcp/namespace-scanning.md` | lines 228-229 | `spec:drmcp.application_architecture.runtime_and_state`; `spec:drmcp.application_architecture.dependency_and_responsibility` | Prose references | `must_update_after_target_decision` | External spec pointers to architecture authority must follow relocated refs. |
| `drmcp/records/spec/design-records-mcp/namespace-scanning.md` | `## Related specs`, lines 264-265 | `spec:drmcp.application_architecture.runtime_and_state`; `spec:drmcp.application_architecture.dependency_and_responsibility` | Related specs | `must_update_after_target_decision` | External related-spec rows must follow relocated refs. |
| `drmcp/records/spec/design-records-mcp/tools/overview.md` | line 39 | `application architecture` | Prose reference | `no_change` | The text does not name the old ref or physical tree. |
| `drmcp/records/spec/implementation/index.md` | whole document | Existing W011 read-runtime implementation root has no `spec:drmcp.application_architecture` refs. | Target-area observation | `maybe_update_after_target_decision` | Target topology decision must decide whether this root remains W011-specific, becomes a navigation root, or receives child topics. |
| `drmcp/records/spec/implementation/contracts/index.md` | line 52 | `spec:drmcp.application_architecture` | Related specs | `must_update_after_target_decision` | Module-contract root must point to relocated architecture authority. |
| `drmcp/records/spec/implementation/contracts/index.md` | line 10 | `application architecture` | Prose reference | `maybe_update_after_target_decision` | Wording can remain if the relocated subtree keeps the architecture label. |
| `drmcp/records/spec/implementation/contracts/application-use-cases/contract-boundary.md` | line 159 | `spec:drmcp.application_architecture.runtime_and_state` | Related specs | `must_update_after_target_decision` | Contract boundary must point to relocated runtime authority. |
| `drmcp/records/spec/implementation/contracts/application-use-cases/index.md` | line 23 | `spec:drmcp.application_architecture.runtime_and_state` | Related specs | `must_update_after_target_decision` | Component index must point to relocated runtime authority. |
| `drmcp/records/spec/implementation/contracts/composition-lifecycle/contract-boundary.md` | line 87 | `spec:drmcp.application_architecture.runtime_and_state` | Related specs | `must_update_after_target_decision` | Contract boundary must point to relocated runtime authority. |
| `drmcp/records/spec/implementation/contracts/composition-lifecycle/index.md` | line 23 | `spec:drmcp.application_architecture.runtime_and_state` | Related specs | `must_update_after_target_decision` | Component index must point to relocated runtime authority. |
| `drmcp/records/spec/implementation/contracts/infrastructure-io-adapters/contract-boundary.md` | line 102 | `spec:drmcp.application_architecture.dependency_and_responsibility` | Related specs | `must_update_after_target_decision` | Contract boundary must point to relocated dependency authority. |
| `drmcp/records/spec/implementation/contracts/infrastructure-io-adapters/index.md` | line 23 | `spec:drmcp.application_architecture.dependency_and_responsibility` | Related specs | `must_update_after_target_decision` | Component index must point to relocated dependency authority. |
| `drmcp/records/spec/implementation/contracts/mcp-inbound-adapter/contract-boundary.md` | line 87 | `spec:drmcp.application_architecture.dependency_and_responsibility` | Related specs | `must_update_after_target_decision` | Contract boundary must point to relocated dependency authority. |
| `drmcp/records/spec/implementation/contracts/mcp-inbound-adapter/index.md` | line 23 | `spec:drmcp.application_architecture.dependency_and_responsibility` | Related specs | `must_update_after_target_decision` | Component index must point to relocated dependency authority. |
| `drmcp/records/spec/implementation/contracts/record-domain-logical-tree/contract-boundary.md` | line 125 | `spec:drmcp.application_architecture.application_boundary_and_components` | Related specs | `must_update_after_target_decision` | Contract boundary must point to relocated component authority. |
| `drmcp/records/spec/implementation/contracts/record-domain-logical-tree/index.md` | line 23 | `spec:drmcp.application_architecture.application_boundary_and_components` | Related specs | `must_update_after_target_decision` | Component index must point to relocated component authority. |

### Summary by likely action

| likely action | count | meaning |
|---|---:|---|
| `must_update_after_target_decision` | 28 inventory rows | Canonical refs, parent refs, topic rows, related-spec rows, or cross-view pointers must change after target refs are decided. |
| `maybe_update_after_target_decision` | 5 inventory rows | Wording or root topology may change depending on selected target topology. |
| `no_change` | 1 inventory row | Generic prose does not bind to old placement. |
| `retain_as_historical_context` | 0 inventory rows | No historical-only spec-internal reference was found inside the scoped search. |

### Physical path result

No text mention of `drmcp/records/spec/application-architecture/` was found under `drmcp/records/spec/`.

### Minimum next decision boundary

The next Task should be a decision Task.
It should decide:

| decision topic | reason |
|---|---|
| Target subtree | The inventory does not decide whether the target ref is `spec:drmcp.implementation.application_architecture`, `spec:drmcp.implementation.architecture`, or another implementation child. |
| Root role of `spec:drmcp.implementation` | The current root still describes the historical W011 read-runtime slice. It may need to become a broader implementation overview or route the old content to a child topic. |
| Compatibility policy | The Work Item must decide whether to leave a transitional note, a non-canonical pointer, or no compatibility stub for the old `spec:drmcp.application_architecture` refs. |
| Writer sequence | File moves, metadata ref edits, topic-table edits, and implementation-root updates may share files. Writer order must be fixed before authoring. |

### Boundary confirmation

No canonical Specification authoring was performed.
No file move was performed.
No ADR was created.
No Requirement, Work Item, Task, or Investigation history cleanup was performed.
No production implementation planning was performed.
