# Reference: Artifact responsibility matrix

- **id**: `spec:product.design_records.artifact_model.artifact_responsibility_matrix`
- **status**: draft
- **date**: 2026-06-24
- **parent**: `spec:product.design_records.artifact_model`

## What this is

Defines the app-independent responsibility boundary for Design Records artifacts.
It does not define BPDSL artifacts, implementation source artifacts, render artifacts, or target implementation artifacts.

## Artifact responsibility matrix

| artifact | owns | does not own |
|---|---|---|
| ADR | Accepted design decisions, rationale for adoption, and rejected alternatives. | Current specifications, exploration logs, progress management, or implementation state. |
| investigation | Research results, evidence, uncertainties, options, and follow-up artifact candidates. | Decisions, current specs, cross-cutting progress, or completion state. |
| requirement | The need, gap, request, and stable requirement identity. | Research process, implementation procedure, or design decision. |
| work item | Requirement resolution flow, goal state, cross-cutting progress, layer-by-layer impact tracking, and task graph. | Spec body, decision history, research report, or canonical subordinate task status. |
| task | Concrete closeable work, completion conditions, individual status, and verification evidence. | Canonical requirements, design decisions, current specs, or work item goal state. |
| spec | Current specifications, scope, and currently valid contracts. | Decision history, exploration logs, progress state, or implementation authority. |

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
