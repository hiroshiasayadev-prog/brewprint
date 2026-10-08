# PRODUCT-REQ-SPEC-016: Hierarchical USDM requirement decomposition and derived coverage

- **status**: accepted
- **date**: 2026-09-30
- **source_refs**:
  - PRODUCT-REQ-SPEC-015
  - `spec:product.design_records.usdm.artifact_format`
  - `spec:product.design_records.usdm.coverage_format`

## Requirement

Brewprint USDM must represent requirement decomposition without introducing a separate parent relation field.

USDM coverage must distinguish direct Specification coverage from coverage derived through a requirement decomposition tree.

A directly covered parent requirement must remain covered even when a more detailed descendant is not separately covered.
The uncovered descendant must remain visible as a non-blocking warning.

## Evidence

- The current row grammar is flat `RNNN`, so detailed functional decomposition cannot preserve parent-child identity.
- The current coverage contract treats only direct `usdm_covers` declarations as coverage.
- Detailed requirement work needs a stable way to add one level or several levels below an existing requirement without renumbering the existing row.
- Requiring every refinement to be covered before a directly covered parent can pass would make refinement discovery block an already covered parent contract.
- Hiding uncovered refinements when a parent is directly covered would remove useful follow-up visibility.

## Required Outcome

- Row-local requirement IDs support `RNNN(-NN)*`.
- Removing the final `-NN` segment from a non-top-level row ID identifies its immediate parent.
- Every non-top-level row has an existing immediate parent in the same USDM requirement record.
- Requirement row decomposition does not cross USDM requirement-record or app-namespace boundaries.
- An implementation Specification may directly cover a USDM requirement row owned by another app namespace by using its full USDM requirement ID in `usdm_covers`.
- A row with no children is a leaf requirement.
- A row with children is effectively covered when it is directly covered or every direct child is effectively covered.
- A leaf is effectively covered only by direct Specification coverage.
- Direct coverage of a parent does not directly cover its descendants.
- An effectively uncovered descendant beneath a directly covered ancestor is reported as a non-blocking warning.
- An effectively uncovered row without a directly covered ancestor remains a blocking uncovered requirement.
- The hierarchy represents pure decomposition only.
- No explicit optional, OR, or conditional relation semantics are introduced by this requirement.
- Existing flat `RNNN` IDs remain valid without migration.

## Explicitly Excluded Scope

- Adding an explicit `optional` property.
- Adding OR, alternative, conditional, weighted, or cardinality-based requirement relations.
- Migrating existing flat requirement IDs.
- Changing USDM record-to-record parent semantics.
- Introducing section-level `usdm_covers`.
- Implementing the standalone or MCP tooling in this design-only change.

## Boundary

This requirement owns hierarchical USDM row identity and effective coverage semantics.

Tool implementation, existing-record migration, and richer relation semantics remain separate downstream work.
