# PRODUCT-WORK-SPEC-016: Define typed single-responsibility Task contract

- **id**: PRODUCT-WORK-SPEC-016
- **status**: done
- **date**: 2026-07-01
- **source_requirement**: PRODUCT-REQ-SPEC-005
- **impact_refs**:
  - PRODUCT-REQ-SPEC-005
  - PRODUCT-REQ-SPEC-006
  - PRODUCT-WORK-SPEC-017
  - spec:product.design_records.authoring_standards.task_authoring
  - spec:product.design_records.authoring_standards.work_item_authoring
  - spec:product.design_records.authoring_standards.artifact_boundary
  - spec:product.design_records.artifact_model.artifact_responsibility_matrix
  - spec:product.design_records.traceability
  - spec:drmcp.design_records_mcp.schema.metadata_grammar
- **tasks**:
  - PRODUCT-TASK-SPEC-016-01
  - PRODUCT-TASK-SPEC-016-02
  - PRODUCT-TASK-SPEC-016-03
  - PRODUCT-TASK-SPEC-016-04
  - PRODUCT-TASK-SPEC-016-05
  - PRODUCT-TASK-SPEC-016-06
  - PRODUCT-TASK-SPEC-016-07
  - PRODUCT-TASK-SPEC-016-08
  - PRODUCT-TASK-SPEC-016-09
  - PRODUCT-TASK-SPEC-016-10
  - PRODUCT-TASK-SPEC-016-11

## Goal

Define and independently validate the canonical typed single-responsibility Task contract required by `PRODUCT-REQ-SPEC-005`.

Resolve Task-type semantics before existing or new Task graphs are migrated to the contract.

## Boundary

This Work Item owns:

- a complete decision inventory for the Task-type and single-responsibility requirement;
- an interactive one-decision-at-a-time loop for unresolved requirement-level choices;
- investigation of semantic and writer conflicts with current Specifications and `PRODUCT-REQ-SPEC-006`;
- a separate interactive decision loop for every material cross-contract conflict;
- ADR routing and authoring for durable choices that require decision history;
- synchronization of accepted Task-type semantics into canonical PRODUCT Specifications;
- consumption of the authoritative integrated-review Evidence from `PRODUCT-TASK-SPEC-017-08`;
- W016 lifecycle and Evidence closure through `PRODUCT-TASK-SPEC-016-11` after the applicable acceptance gate.

`PRODUCT-TASK-SPEC-017-08` owns the only cross-requirement integrated review.
`PRODUCT-TASK-SPEC-017-09` and `PRODUCT-TASK-SPEC-017-10` own integrated finding correction and closure re-review when required.

This Work Item does not own:

- a separate W016 integrated review, correction, or re-review execution route;
- validator implementation;
- checklist or diagnostic schema design;
- agent model routing;
- bulk migration of existing Tasks;
- DRMCP parser, validator, index, or projection implementation;
- generic source-relation semantics owned by `PRODUCT-WORK-SPEC-017`;
- implementation planning for unrelated product or DRMCP Work Items.

The child Task graph retains the legacy pre-REQ-005/006 metadata shape under `BOOTSTRAP-001`.
The Tasks use `source_requirement` and `work_item` without an unaccepted primary-type field or `source_refs`.
This workflow-local exclusion does not define canonical Task metadata.
These workflow Tasks require no later migration.

## Impact Scope

| ref or area | impact |
|---|---|
| `PRODUCT-REQ-SPEC-005` | Supplies the accepted typed single-responsibility requirement. |
| `PRODUCT-REQ-SPEC-006` | Shares Task metadata and Task-authoring surfaces and may constrain workflow provenance wording. |
| `PRODUCT-WORK-SPEC-017` | Owns the generic source-relation decision and Specification workflow. Its shared-file synchronization must not write concurrently with this Work Item. |
| `spec:product.design_records.authoring_standards.task_authoring` | Primary canonical target for Task type, one-responsibility, completion-judgment, and prohibited-overlap rules. |
| `spec:product.design_records.authoring_standards.work_item_authoring` | May require coordination-boundary wording when Work Items delegate typed Tasks. |
| `spec:product.design_records.authoring_standards.artifact_boundary` | May require clarification of Task versus ADR, Specification, review, correction, coordination, and synchronization ownership. |
| `spec:product.design_records.artifact_model.artifact_responsibility_matrix` | May require the canonical Task ownership row to match the accepted type model. |
| `spec:product.design_records.traceability` | Must remain consistent with typed Task responsibility and evidence ownership. |
| `spec:drmcp.design_records_mcp.schema.metadata_grammar` | Downstream DRMCP consumer of the accepted metadata contract; PRODUCT does not implement it here. |

## Task flow

| phase | dependency | outcome |
|---|---|---|
| A. Requirement decision inventory | Accepted `PRODUCT-REQ-SPEC-005` | Identify every unresolved Task-type, responsibility, completion, overlap, and migration-boundary decision. |
| B. Requirement decision loop | Phase A | Ask exactly one unresolved decision at a time and persist every explicit answer. |
| C. Specification-conflict investigation | Phase B and the current PRODUCT Specifications | Identify stale contracts, shared-file writer conflicts, semantic overlap with REQ-006, and downstream DRMCP impacts. |
| D. Conflict decision loop | Phase C | Resolve every material conflict one at a time without silently choosing precedence. |
| E. ADR routing and authoring | Phases B and D | Create, amend, or supersede ADRs only for durable choices that require decision history. |
| F. PRODUCT Specification synchronization | Phase E | Write the accepted Task-type contract into the canonical PRODUCT Specifications. This writer completes before W017 writes shared Specification files. |
| G. Cross-requirement integration | W017 Specification synchronization | Consume the final REQ-006 source-relation changes and verify that shared Task-authoring content preserves the accepted REQ-005 contract. |
| H. Authoritative integrated review | Phase G | Consume the cross-requirement verdict and findings from `PRODUCT-TASK-SPEC-017-08`. |
| I. Conditional integrated correction and re-review | W017 T08 `NEEDS REVISION` | W017 T09 corrects named findings. W017 T10 independently dispositions every required finding. |
| J. W016 closure synchronization | Applicable integrated-review acceptance gate | `PRODUCT-TASK-SPEC-016-11` synchronizes only W016 lifecycle and Evidence. |

Requirement decision work for W016 and W017 may proceed in parallel when their decision ledgers and writers are separate.
The final flow is serialized:

```text
W016 Task-contract Specification writer
  -> W017 source-relation Specification writer
  -> W017 integrated review
  -> conditional W017 correction and re-review
  -> separate W016 and W017 closure synchronization
```

| T08 verdict | W016 route |
|---|---|
| `PASS` | W016 T11 consumes accepted T08 Evidence. W017 T09 and T10 do not execute. |
| `NEEDS REVISION` | W017 T09 completes correction. W017 T10 independently closes every required finding. W016 T11 then consumes the accepted T10 dispositions. |

## Task Candidates

The following released Tasks form the W016 workflow graph.
They do not declare an unaccepted primary Task type.

| Task | primary outcome | dependency |
|---|---|---|
| `PRODUCT-TASK-SPEC-016-01` | Complete known REQ-005 decision inventory and dependency order. | None. |
| `PRODUCT-TASK-SPEC-016-02` | Persist one REQ-005 decision per user turn. | T01. |
| `PRODUCT-TASK-SPEC-016-03` | Exact Specification-conflict and shared-writer inventory. | T02. |
| `PRODUCT-TASK-SPEC-016-04` | Persist one Specification-conflict disposition per user turn. | T03. |
| `PRODUCT-TASK-SPEC-016-05` | Complete ADR routing classification. | T02 and T04. |
| `PRODUCT-TASK-SPEC-016-06` | Required ADR set authored or explicit no-ADR result. | T05. |
| `PRODUCT-TASK-SPEC-016-07` | Accepted REQ-005 contract synchronized to PRODUCT Specifications. | T05 and T06. |
| `PRODUCT-TASK-SPEC-016-08` | Retained historical planned review route. Superseded by W017 T08 and not required for execution. | Historical dependency remains T07. |
| `PRODUCT-TASK-SPEC-016-09` | Retained historical planned correction route. Superseded by W017 T09 and not required for execution. | Historical dependency remains W016 T08. |
| `PRODUCT-TASK-SPEC-016-10` | Retained historical planned re-review route. Superseded by W017 T10 and not required for execution. | Historical dependency remains W016 T09. |
| `PRODUCT-TASK-SPEC-016-11` | Synchronize only W016 lifecycle and Evidence after the applicable integrated-review acceptance gate. | Metadata predecessor is W017 T08. T10 acceptance is additionally required only after `NEEDS REVISION`. |

W016 T08, T09, and T10 remain listed for traceability.
They are non-executable superseded routes and do not require synthetic no-op completion.

## Completion Condition

- Every REQ-005 requirement-level decision is durably recorded or explicitly deferred or blocked.
- Every material conflict with current Specifications or REQ-006 has an explicit disposition.
- The allowed primary Task-type set is closed and canonical.
- Every Task type has one owned outcome, one completion judgment, and explicit prohibited overlaps.
- Authoring, independent review, correction, finding-closure review, coordination, synchronization, implementation, and verification boundaries are unambiguous.
- `Goal`, `Work`, `Done condition`, and `Verification` alignment rules are canonical.
- Required ADRs are accepted and linked to the current Specification state.
- Canonical PRODUCT Specifications contain the final Task contract without stale contradictory text.
- W017 source-relation synchronization does not weaken or overwrite the accepted REQ-005 contract.
- `PRODUCT-TASK-SPEC-017-08` provides the only accepted cross-requirement integrated-review verdict.
- A direct `PASS` permits W016 T11 without W017 T09 or T10 execution.
- A `NEEDS REVISION` verdict requires W017 T09 completion and accepted W017 T10 finding dispositions before W016 T11.
- Required findings are independently closed.
- W016 T11 synchronizes only W016 lifecycle and Evidence.
- Existing Task migration is explicitly routed but not performed by this Work Item.
- Downstream DRMCP implementation requirements are identified without implementation beginning here.

## Evidence

- `PRODUCT-REQ-SPEC-005`: accepted source Requirement.
- `PRODUCT-REQ-SPEC-006`: accepted related Requirement sharing Task metadata and authoring surfaces.
- `PRODUCT-TASK-SPEC-017-08`: authoritative integrated review owner and current `NEEDS REVISION` verdict.
- `PRODUCT-TASK-SPEC-017-09`: authoritative correction owner for the integrated findings.
- `PRODUCT-TASK-SPEC-017-10`: authoritative independent finding-closure owner.
- W016 T08, T09, and T10 remain historical planned records.
- The three W016 records are superseded, not required, and non-executable.
- W016 T11 remains the separate W016 lifecycle and Evidence closure owner.
- `BOOTSTRAP-001`: user decided on 2026-06-30 to retain the then-current legacy metadata shape for these workflow Tasks.
- Bootstrap metadata uses `source_requirement` and `work_item` only.
- Bootstrap metadata does not add `type`, `task_type`, `primary_type`, or `source_refs`.
- The user explicitly excluded migration requirements for these workflow Tasks.
- The bootstrap decision is workflow Evidence and is not the canonical REQ-005 product contract.

### Closure result

- Authoritative integrated review: `PRODUCT-TASK-SPEC-017-08`.
- Accepted finding closure: `PRODUCT-TASK-SPEC-017-10`.
- Overall verdict: `PASS`.
- `F-MAJ-01`: `CLOSED`.
- `F-MAJ-02`: `CLOSED`.
- Canonical ADR set: `PRODUCT-ADR-SPEC-004`, `PRODUCT-ADR-SPEC-005`, and `PRODUCT-ADR-SPEC-006`.
- Canonical Specification synchronization: exact `PRODUCT-TASK-SPEC-016-07` outputs.
- All W016 completion conditions: satisfied.
- W016 lifecycle closure owner: `PRODUCT-TASK-SPEC-016-11`.
- Existing Task migration: not performed.
- Downstream DRMCP work: remains separate.
- Work Item status: synchronized to `done`.
