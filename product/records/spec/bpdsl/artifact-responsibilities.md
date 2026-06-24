# Reference: Artifact responsibility preservation

- **id**: `spec:product.bpdsl.artifact_responsibilities`
- **status**: draft
- **date**: 2026-06-24
- **parent**: `spec:product.bpdsl`

## What this is

Preservation-only staging for BPDSL and implementation-flow responsibility rows removed from the generic artifact responsibility matrix.
It preserves existing meaning without accepting canonical PRODUCT ownership.

## Preservation status

| rule | status |
|---|---|
| Ownership claim | No PRODUCT canonical ownership claim. |
| Expected final owner | `bpdsl/records/spec/**` for BPDSL-internal content after migration review. |
| Integration claim | Historical; not adopted by PRODUCT. Retained as migration-review evidence only. Historical model differs from current app-local BPDSL contracts. |
| Migration review obligation | Review all retained statements during BPDSL migration or when an explicit integration requirement is accepted. |

## Preserved responsibility rows

| artifact | preserved owns statement | preserved does-not-own statement | status |
|---|---|---|---|
| internal design | Wiring or route from spec semantics to implementation. | Canonical authority for spec semantics or primary source for target model. | Previous description; not adopted by PRODUCT. No current app-local BPDSL owner identified for an internal design artifact class. Historical context; retained as migration-review evidence only. |
| brewprint DSL definition | Target implementation source when the app's DSL pipeline is operational. | Decision history or responsibility boundaries of Design Records. | Preservation-only BPDSL staging. |
| source implementation | Generated realization of DSL definitions, or handwritten bootstrap implementation while DSL support is insufficient. | Canonical authority for design contracts. | Preservation-only implementation-flow staging. |
| render | Views derived from DSL definitions. | Editable source of truth. | Preservation-only BPDSL staging. |
| target implementation | Executable or deployable realization built from source implementation. | Canonical authority for design contracts. | Preservation-only implementation-flow staging. |
| impl note | Handover and review notes for completed implementation. | Current specs or future tasks. | Preservation-only implementation-flow staging. |

## Deferred integration disposition

Primary disposition lives in `spec:product.design_records.artifact_model`.
This BPDSL staging file keeps only the minimum pointer needed to preserve the old matrix context.

| preserved statement | disposition |
|---|---|
| External relation or assurance artifacts are not active artifacts in the MVP. | Existing evidence owner: V01-INV-DOCS-002 and V01-ADR-088. |

## Source mapping

| source file | source sections | disposition |
|---|---|---|
| `product/records/spec/concepts/project-artifact-model/artifact-responsibility-matrix.md` | BPDSL, source implementation, render, target implementation, internal design, and impl note rows. External relation rows are pointer-only to the primary T05 location. | BPDSL and implementation-flow rows preserved here; external relation material is not duplicated as primary text. |
| `product/records/spec/concepts/project-artifact-model/index.md` | Design and implementation artifact rows. | Preserved here and in `spec:product.bpdsl.design_flow`. |

## Related specs

| ref | relation |
|---|---|
| `spec:product.bpdsl` | Temporary staging parent. |
| `spec:product.design_records.artifact_model.artifact_responsibility_matrix` | Generic Design Records responsibility matrix after extraction. |
