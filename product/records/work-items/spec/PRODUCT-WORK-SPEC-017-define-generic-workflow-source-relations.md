# PRODUCT-WORK-SPEC-017: Define generic workflow source relations

- **id**: PRODUCT-WORK-SPEC-017
- **status**: done
- **date**: 2026-07-01
- **source_requirement**: PRODUCT-REQ-SPEC-006
- **impact_refs**:
  - PRODUCT-REQ-SPEC-005
  - PRODUCT-REQ-SPEC-006
  - PRODUCT-WORK-SPEC-016
  - spec:product.design_records.authoring_standards.requirement_authoring
  - spec:product.design_records.authoring_standards.work_item_authoring
  - spec:product.design_records.authoring_standards.task_authoring
  - spec:product.design_records.authoring_standards.artifact_boundary
  - spec:product.design_records.artifact_model.artifact_responsibility_matrix
  - spec:product.design_records.traceability
  - spec:drmcp.design_records_mcp.schema.metadata_grammar
- **tasks**:
  - PRODUCT-TASK-SPEC-017-01
  - PRODUCT-TASK-SPEC-017-02
  - PRODUCT-TASK-SPEC-017-03
  - PRODUCT-TASK-SPEC-017-04
  - PRODUCT-TASK-SPEC-017-05
  - PRODUCT-TASK-SPEC-017-06
  - PRODUCT-TASK-SPEC-017-07
  - PRODUCT-TASK-SPEC-017-08
  - PRODUCT-TASK-SPEC-017-09
  - PRODUCT-TASK-SPEC-017-10
  - PRODUCT-TASK-SPEC-017-11

## Goal

Define and independently validate the canonical generic workflow source-relation contract required by `PRODUCT-REQ-SPEC-006`.

Replace Requirement-only provenance and persisted reverse membership without losing Task-to-Work-Item decomposition or workflow traceability.

## Boundary

This Work Item owns:

- a complete decision inventory for generic Work Item and Task source relations;
- an interactive one-decision-at-a-time loop for unresolved requirement-level choices;
- investigation of semantic, migration, and writer conflicts with current Specifications and `PRODUCT-REQ-SPEC-005`;
- a separate interactive decision loop for every material cross-contract conflict;
- ADR routing and authoring for durable provenance, ownership, and migration choices;
- synchronization of accepted source-relation semantics into canonical PRODUCT Specifications;
- removal of Requirement `work_items` from the canonical contract and definition of derived reverse relations;
- a final cross-requirement integration review covering REQ-005 and REQ-006 shared surfaces;
- conditional finding correction, independent re-review, and closure synchronization.

This Work Item does not own:

- DRMCP parser, validator, diagnostic, index, reverse-traversal, or tool-projection implementation;
- validator model selection;
- hub progress calculation;
- a new Hub Work Item artifact kind;
- unlimited Work Item nesting semantics;
- concrete DRMCP request and response schemas;
- Task-type semantics owned by `PRODUCT-WORK-SPEC-016`;
- bulk migration of existing records.

`BOOTSTRAP-001` reuses the user decision recorded in `PRODUCT-TASK-SPEC-016-01`.

The workflow Tasks retain the legacy `source_requirement` and `work_item` metadata shape under `BOOTSTRAP-001`.
This retained shape is a workflow-local exclusion, not the canonical Task metadata contract.
The workflow Tasks do not add unaccepted Task-type or `source_refs` fields.
The workflow Tasks require no migration action.

## Impact Scope

| ref or area | impact |
|---|---|
| `PRODUCT-REQ-SPEC-006` | Supplies the accepted generic source-relation requirement. |
| `PRODUCT-REQ-SPEC-005` | Constrains every future Task to one accepted primary type and one primary responsibility. |
| `PRODUCT-WORK-SPEC-016` | Owns Task-type semantics and writes shared Task-authoring content before this Work Item synchronizes source relations. |
| `spec:product.design_records.authoring_standards.requirement_authoring` | Remove canonical persisted `work_items` ownership and define derived reverse-relation expectations. |
| `spec:product.design_records.authoring_standards.work_item_authoring` | Replace `source_requirement` with generic `source_refs` and define downstream Work Item provenance. |
| `spec:product.design_records.authoring_standards.task_authoring` | Remove Task `source_requirement` without replacement. Preserve `work_item` provenance and the accepted REQ-005 type contract. Tasks persist no source field. |
| `spec:product.design_records.authoring_standards.artifact_boundary` | Clarify provenance references versus canonical artifact ownership. |
| `spec:product.design_records.artifact_model.artifact_responsibility_matrix` | Align Requirement, Work Item, and Task relation ownership with forward-only persistence. |
| `spec:product.design_records.traceability` | Define forward source relations and DRMCP-derived reverse relation views. |
| `spec:drmcp.design_records_mcp.schema.metadata_grammar` | Downstream DRMCP consumer of the accepted metadata and migration contract; PRODUCT does not implement it here. |

## Task flow

| phase | dependency | outcome |
|---|---|---|
| A. Requirement decision inventory | Accepted `PRODUCT-REQ-SPEC-006` | Identify unresolved source-ref, reverse-relation, decomposition, validation, and migration decisions. |
| B. Requirement decision loop | Phase A | Ask exactly one unresolved decision at a time and persist every explicit answer. |
| C. Specification-conflict investigation | Phase B and the current PRODUCT Specifications | Identify stale reciprocal relations, shared-file writer conflicts, REQ-005 overlap, and downstream DRMCP impacts. |
| D. Conflict decision loop | Phase C | Resolve every material semantic or migration conflict one at a time. |
| E. ADR routing and authoring | Phases B and D | Create, amend, or supersede ADRs only for durable provenance and migration choices requiring decision history. |
| F. Wait for W016 shared-file synchronization | W016 Specification synchronization | Consume the accepted Task-type contract before writing shared `task_authoring` or artifact-model surfaces. |
| G. PRODUCT Specification synchronization | Phase F | Write source-relation, reverse-view, decomposition, validation, and migration contracts into canonical PRODUCT Specifications. |
| H. Cross-requirement integrated review | W016 and W017 synchronization | `PRODUCT-TASK-SPEC-017-08` independently reviews both Requirements against the final shared Specification state. |
| I. Conditional correction and re-review | T08 `NEEDS REVISION` | W017 T09 corrects named findings. W017 T10 independently dispositions every required finding. |
| J. Separate closure synchronization | Applicable integrated-review acceptance gate | W016 T11 closes only W016. W017 T11 closes only W017. |

Requirement decision work for W016 and W017 may proceed in parallel when their decision ledgers and writers are separate.
The final flow is serialized:

```text
W016 Task-contract Specification writer
  -> W017 source-relation Specification writer
  -> W017 integrated review
  -> conditional W017 correction and re-review
  -> separate W016 and W017 closure synchronization
```

`PRODUCT-TASK-SPEC-017-08` is the only cross-requirement integrated review owner.
The reviewer must be independent of both Specification writers.

| T08 verdict | required route |
|---|---|
| `PASS` | W017 T09 and T10 do not execute. W016 T11 and W017 T11 consume accepted T08 Evidence. |
| `NEEDS REVISION` | W017 T09 completes correction. W017 T10 independently closes every required finding. Both T11 Tasks then consume accepted T10 dispositions. |

Synthetic no-op Tasks must not satisfy either branch.

## Task Candidates

The following Task graph is retained under the `BOOTSTRAP-001` workflow-local metadata exclusion:

| Task | primary outcome | dependency |
|---|---|---|
| `PRODUCT-TASK-SPEC-017-01` | Record the complete known REQ-006 decision inventory and dependency order. | None. |
| `PRODUCT-TASK-SPEC-017-02` | Persist one REQ-006 decision per user turn. | T01. |
| `PRODUCT-TASK-SPEC-017-03` | Inventory exact Specification and migration conflicts. | T02. |
| `PRODUCT-TASK-SPEC-017-04` | Persist one conflict disposition per user turn. | T03. |
| `PRODUCT-TASK-SPEC-017-05` | Classify ADR routing for all decided items. | T02 and T04. |
| `PRODUCT-TASK-SPEC-017-06` | Author only ADRs required by T05. | T05. |
| `PRODUCT-TASK-SPEC-017-07` | Synchronize REQ-006 Specifications after the W016 writer. | W016-07, T05, and T06. |
| `PRODUCT-TASK-SPEC-017-08` | Own the only integrated REQ-005 and REQ-006 review verdict. | W016-07 and T07. |
| `PRODUCT-TASK-SPEC-017-09` | Correct only named T08 findings after `NEEDS REVISION`. | T08. Not required after `PASS`. |
| `PRODUCT-TASK-SPEC-017-10` | Independently re-review corrected findings after T09. | T09. Not required after `PASS`. |
| `PRODUCT-TASK-SPEC-017-11` | Synchronize only W017 lifecycle and Evidence after the applicable acceptance gate. | Metadata predecessor is T08. T10 acceptance is additionally required only after `NEEDS REVISION`. |

Each Task owns one primary outcome and one completion judgment.
W016 T11 and W017 T11 remain separate closure owners.

## Completion Condition

- Every REQ-006 requirement-level decision is durably recorded or explicitly deferred or blocked.
- Every material conflict with current Specifications or REQ-005 has an explicit disposition.
- Work Items persist required non-empty `source_refs` as canonical workflow provenance.
- Tasks persist no source field and do not persist `source_refs`.
- Task provenance is reached through `work_item`.
- Legacy Task `source_requirement` migration is removal-only.
- Multiple material upstream sources, including Requirements, Tasks, investigations, decisions, and Specifications, are represented without dedicated parent or hub fields.
- Task `work_item` and Work Item `tasks` remain the explicit ownership relation.
- Requirement `work_items` is removed from the canonical persisted metadata contract.
- Requirement-to-Work-Item reverse views are defined as derived DRMCP projections.
- Existence, canonical identity, duplicate, and self-reference validation rules are canonical.
- Work Item `source_requirement` and Requirement `work_items` migration behavior is explicit.
- Bootstrap workflow Tasks retain the old metadata shape only as a workflow-local exclusion.
- The bootstrap exclusion is not a canonical metadata exception.
- Downstream Work Item creation from a Task is represented through forward source relations without hidden hierarchy semantics.
- Required ADRs are accepted and linked to the current Specification state.
- Canonical PRODUCT Specifications contain the final relation contract without stale contradictory text.
- Shared Task-authoring content preserves the accepted REQ-005 Task-type and single-responsibility contract.
- `PRODUCT-TASK-SPEC-017-08` provides the only cross-requirement integrated-review verdict.
- A direct `PASS` permits both closure Tasks without T09 or T10 execution.
- A `NEEDS REVISION` verdict requires T09 completion and accepted T10 finding dispositions before either closure Task.
- Required findings are independently closed.
- Existing record migration is explicitly routed but not performed by this Work Item.
- Downstream DRMCP implementation requirements are identified without implementation beginning here.

## Evidence

- `PRODUCT-REQ-SPEC-006`: accepted source Requirement.
- `PRODUCT-REQ-SPEC-005`: accepted related Requirement constraining future Task records.
- `PRODUCT-ADR-SPEC-007`: Work Item-only persisted `source_refs` and Task provenance through `work_item`.
- `PRODUCT-ADR-SPEC-008`: removal-only legacy Task source migration.
- `PRODUCT-TASK-SPEC-017-08`: only integrated review owner and initial `NEEDS REVISION` verdict.
- `PRODUCT-TASK-SPEC-017-09`: integrated finding-correction owner.
- `PRODUCT-TASK-SPEC-017-10`: independent finding-closure owner.
- W016 T11 and W017 T11 remain separate lifecycle and Evidence closure owners.
- `BOOTSTRAP-001` authority: `PRODUCT-TASK-SPEC-016-01`, `## Evidence`, `### Bootstrap disposition`.
- The released T01 through T11 graph uses legacy `source_requirement` and `work_item` metadata only.
- The old Task metadata shape is a workflow-local exclusion.
- The workflow-local exclusion does not alter canonical Task metadata semantics.
- The workflow Tasks require no migration action.
- Existing-record migration execution remains outside this Work Item.

### Closure result

- Authoritative integrated review: `PRODUCT-TASK-SPEC-017-08`.
- Accepted finding closure: `PRODUCT-TASK-SPEC-017-10`.
- Overall verdict: `PASS`.
- `F-MAJ-01`: `CLOSED`.
- `F-MAJ-02`: `CLOSED`.
- Accepted Requirement: `PRODUCT-REQ-SPEC-006`.
- Canonical ADR set: `PRODUCT-ADR-SPEC-001`, `PRODUCT-ADR-SPEC-007`, and `PRODUCT-ADR-SPEC-008`.
- Canonical Requirement and Specification synchronization: W017 T07 exact trace.
- C-027 through C-030 remain downstream DRMCP obligations.
- Migration semantics are design-complete.
- Existing-record migration execution remains outside W017.
- All W017 completion conditions: satisfied.
- W017 lifecycle closure owner: `PRODUCT-TASK-SPEC-017-11`.
- Work Item status synchronized to `done`.
