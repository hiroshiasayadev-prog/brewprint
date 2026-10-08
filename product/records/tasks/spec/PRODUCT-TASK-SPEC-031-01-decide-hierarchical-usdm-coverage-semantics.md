# PRODUCT-TASK-SPEC-031-01: Decide hierarchical USDM coverage semantics

- **status**: done
- **date**: 2026-09-30
- **work_item**: PRODUCT-WORK-SPEC-031
- **task_type**: decision
- **estimate**: 0.25d
- **depends_on**:
- **outputs**:

## Goal

Fix the durable hierarchy and effective coverage semantics for PRODUCT-REQ-SPEC-016.

## Work

- Resolve the current flat row-ID and direct-coverage contracts.
- Fix pure decomposition semantics for hierarchical row IDs.
- Fix recursive effective coverage and uncovered-descendant warning behavior.
- Keep richer relation semantics outside the current boundary.

## Done condition

Every material hierarchy and coverage judgment is terminal and can be authored without interpretation.

## Verification

The decision set determines row identity, parent derivation, effective coverage, warning behavior, and excluded relation semantics.

## Evidence

| decision | state | selected outcome | reason |
|---|---|---|---|
| D001 | decided | Row IDs use `RNNN(-NN)*`; each non-top-level row's parent is the ID with the final `-NN` removed. | The ID carries the decomposition relation without duplicate parent metadata. |
| D002 | decided | Every non-top-level row requires its immediate parent in the same USDM record. | A hierarchical ID without its parent would not form a valid decomposition tree. |
| D003 | decided | Effective coverage is direct coverage OR, for a non-leaf, effective coverage of every direct child. | A parent absent from implementation specs can be satisfied through complete decomposition coverage. |
| D004 | decided | Direct parent coverage is sufficient regardless of descendant coverage. | The parent contract is already represented by a Specification. |
| D005 | decided | An uncovered descendant under a directly covered ancestor is non-blocking but remains visible as a warning. | Refinement visibility is preserved without invalidating a directly covered parent. |
| D006 | decided | An uncovered row with no directly covered ancestor is blocking. | No accepted coverage path satisfies that requirement. |
| D007 | decided | No explicit optional, OR, conditional, weighted, or cardinality semantics are added. | No current example requires them; warning behavior provides the needed non-blocking refinement behavior. |
| D008 | decided | Existing flat `RNNN` IDs and top-level compact range syntax remain valid. | The change is backward-compatible and does not require migration or a new range grammar. |
