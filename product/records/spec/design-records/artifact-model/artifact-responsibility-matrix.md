# Reference: Artifact responsibility matrix

- **id**: `spec:product.design_records.artifact_model.artifact_responsibility_matrix`
- **status**: draft
- **date**: 2026-07-01
- **parent**: `spec:product.design_records.artifact_model`

## What this is

Defines the app-independent responsibility boundary for Design Records artifacts.
It does not define BPDSL artifacts, implementation source artifacts, render artifacts, or target implementation artifacts.

## Artifact responsibility matrix

| artifact | owns | does not own |
|---|---|---|
| ADR | Accepted design decisions, rationale for adoption, and rejected alternatives. | Current specifications, exploration logs, progress management, or implementation state. |
| investigation | Research results, evidence, uncertainties, options, and follow-up artifact candidates. | Decisions, current specs, cross-cutting progress, or completion state. |
| requirement | The need, gap, request, and stable requirement identity. | Research process, implementation procedure, design decision, or persisted derived Work Item membership. |
| work item | Direct material provenance, bounded resolution flow, goal state, cross-cutting progress, layer-by-layer impact tracking, Task graph, and Task-originated downstream decomposition. | Spec body, decision history, research report, canonical subordinate Task status, or implicit parent-child Work Item hierarchy. |
| task | Concrete closeable work, completion conditions, individual status, verification Evidence, and temporary or historical decision-workflow state. | Canonical requirements, canonical design state, durable decision rationale, current Specifications, Work Item goal state, or duplicated source provenance. |
| spec | Current specifications, scope, and currently valid contracts. | Decision history, exploration logs, progress state, or implementation authority. |

## Canonical design-state boundary

| artifact | design-state responsibility |
|---|---|
| Task | Owns resumable decision-workflow state and historical checkpoints. Does not own canonical design state. |
| ADR | Owns durable choice, alternatives, rationale, consequences, and supersession history. |
| spec | Owns the current normative behavior, structure, boundary, and constraint. |

A decision Task may retain questions, options, answers, concise reasons, routing state, and cursor position.
The Task ledger must not replace an ADR or Specification.

## Completed workflow record boundary

Completed decision, authoring, and review Tasks remain historical Evidence.
Later design judgment uses a new decision, authoring, and integrated review Task chain.
Completed records do not change their outcome, Evidence, verdict, or finding set to represent later work.

Correction owns repair of named findings.
A later independent review owns finding closure.
The correction owner must not close the repaired findings.

## Synchronization ownership boundary

Synchronization owns only exact mechanically derivable propagation named by its contract:

- lifecycle state;
- closure Evidence;
- completion-result propagation;
- relation propagation.

Synchronization does not own canonical authoring, correction, Task graph changes, or new judgment.
Synchronization must not rewrite a completed record's outcome, Evidence, verdict, or finding set.
Synchronization stops when required state cannot be derived mechanically.

## Source-relation ownership boundary

| owner | responsibility |
|---|---|
| PRODUCT | Persisted provenance semantics, canonical relation meaning, direct material source selection, invalid conditions, and migration semantics. |
| DRMCP | Parsing, indexing, Task-owner resolution mechanics, direct reverse lookup, transitive traversal, cycle-analysis algorithms, diagnostics, response schemas, and user-visible projections. |

Work Item `source_refs` is the only persisted workflow provenance field.
Task provenance is reached through Task `work_item`.
Requirement-to-Work-Item reverse membership is derived from direct Work Item refs.

The ownership split does not change Task workflow-state ownership or ADR and Specification design-state ownership.

## Extracted implementation responsibilities

The previous matrix included implementation-flow rows.
Those rows are not generic Design Records responsibilities.
They are preserved under temporary BPDSL staging for later BPDSL migration review.

| previous row | disposition | preserved location |
|---|---|---|
| `internal design` | Not adopted by PRODUCT as a current generic Design Records artifact class; preserved only in temporary BPDSL staging for T09 classification. | `spec:product.bpdsl.artifact_responsibilities`. |
| `brewprint DSL definition` | Stage BPDSL. | `spec:product.bpdsl.artifact_responsibilities`. |
| `source implementation` | Stage BPDSL. | `spec:product.bpdsl.artifact_responsibilities`. |
| `render` | Stage BPDSL. | `spec:product.bpdsl.artifact_responsibilities`. |
| `target implementation` | Stage BPDSL. | `spec:product.bpdsl.artifact_responsibilities`. |
| `impl note` | Stage implementation-flow material. | `spec:product.bpdsl.artifact_responsibilities`. |

## Deferred material disposition

Primary disposition lives in `spec:product.design_records.artifact_model`.
This file keeps only the matrix-local pointer.

| source material | final disposition |
|---|---|
| External relation or assurance artifacts for completeness, evidence, sign-off, approval, central matrices, or relation lifecycle. | Existing evidence owner: V01-INV-DOCS-002 and V01-ADR-088. |
| A standalone `internal-design:` trace endpoint. | Existing evidence owner: V01-INV-DOCS-003 and V01-ADR-088. |

## Related specs

| ref | relation |
|---|---|
| `spec:product.design_records.artifact_model` | Parent artifact-model overview. |
| `spec:product.bpdsl.artifact_responsibilities` | Temporary preservation of extracted BPDSL and implementation-flow responsibility rows. |
| PRODUCT-REQ-SPEC-005 | Typed single-responsibility Task contract. |
| PRODUCT-REQ-SPEC-006 | Generic workflow source-relation requirement. |
| PRODUCT-ADR-SPEC-001 | PRODUCT semantic ownership and app-local implementation boundary. |
| PRODUCT-ADR-SPEC-006 | Decision checkpoints and canonical design-state boundary. |
| PRODUCT-ADR-SPEC-007 | Canonical workflow provenance and validity semantics. |
| PRODUCT-ADR-SPEC-008 | Atomic workflow relation migration semantics. |
| PRODUCT-ADR-SPEC-013 | Finding-driven correction and closure-review materialization. |
| PRODUCT-ADR-SPEC-014 | Completed-record preservation and synchronization write ownership. |
