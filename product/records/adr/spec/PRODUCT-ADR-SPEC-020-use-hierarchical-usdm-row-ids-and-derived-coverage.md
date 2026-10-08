# PRODUCT-ADR-SPEC-020: Use hierarchical USDM row IDs and derived coverage

- **status**: accepted
- **date**: 2026-09-30
- **depends_on**:
- **supersedes**:
- **migrated_to_spec**: 2026-09-30

## Context

MVP USDM requirement rows use flat `RNNN` row IDs.
Coverage is direct when an implementation Specification lists a full requirement ID in `usdm_covers`.

Detailed requirement decomposition needs stable parent-child identity without adding duplicate parent metadata.
The coverage model also needs to distinguish an uncovered omission from an uncovered refinement beneath an already covered parent.

Requiring every descendant to be covered before a directly covered parent passes would make later refinement block an existing parent contract.
Silently treating parent coverage as descendant coverage would hide useful refinement gaps.

## Decision

Use row-local requirement IDs with this grammar:

```text
RNNN(-NN)*
```

A top-level row such as `R001` has no parent.
For a non-top-level row, remove the final `-NN` segment to derive the immediate parent.
The immediate parent must exist in the same USDM requirement record.
The requirement decomposition tree never crosses into another USDM requirement record or app namespace.

Each sibling set owns its own numeric sequence. A newly authored child sibling set starts at `01`. New sibling IDs increase monotonically. Gaps after reviewed removal are allowed, and a removed sibling ID is never reused for a different requirement.

The hierarchy represents pure requirement decomposition.
The hierarchy does not encode optional, OR, conditional, weighted, or cardinality relations.

### Effective coverage

Keep direct coverage as the explicit relation from an implementation Specification to a USDM requirement row.

Define effective coverage recursively:

```text
effective_covered(row) =
  direct_covered(row)
  OR
  (has_children(row) AND every direct child is effective_covered)
```

A leaf therefore requires direct coverage for its own effective coverage.

Direct coverage of a parent is sufficient for that parent's effective coverage.
Direct parent coverage does not create direct coverage relations for descendants.

Cross-app implementation responsibility is expressed through direct coverage rather than cross-app row parentage. A Specification may list a full USDM requirement ID owned by another app namespace in `usdm_covers`; that creates a direct coverage relation to the named row while leaving the target row's decomposition tree local to its owning record.

### Uncovered descendant handling

An effectively uncovered row is blocking when it has no directly covered ancestor.

An effectively uncovered row is non-blocking when a directly covered ancestor exists.
Coverage checks report that row as a warning so the refinement remains visible.

The warning state does not add an `optional` attribute and does not redefine the row as semantically optional.
It provides the required non-blocking refinement behavior until a richer relation case is demonstrated.

### Compatibility boundary

Existing flat `RNNN` row IDs remain valid top-level IDs.

Existing top-level compact range syntax such as `R001-R005` remains unchanged.
Hierarchical rows may be referenced directly or as comma-separated row tokens.
This decision does not introduce hierarchical range shorthand.

No cross-record `refines`, `derived_from`, or equivalent row-to-row relation is introduced by this decision.

Existing rows are not migrated or renumbered by this decision.

## Rationale

Encoding decomposition in the row ID keeps the relation stable and visible without adding a second parent field that could disagree with the ID.

Recursive all-child coverage lets a parent absent from implementation Specifications be satisfied by its complete decomposition.

Direct coverage remains authoritative for the exact requirement a Specification claims to cover.
Treating uncovered descendants as warnings preserves refinement visibility without turning an already covered parent into a failure.

Deferring richer relation semantics keeps the current model limited to cases demonstrated by the product work.
A later requirement can add explicit relation types when a real OR, conditional, or optional case appears.

## Rejected alternatives

| alternative | rejection reason |
|---|---|
| Keep flat row IDs and add a separate parent field. | ID and parent metadata could disagree and would duplicate a relation derivable from the hierarchical ID. |
| Require all descendants to be effectively covered even when the parent is directly covered. | Later refinement would invalidate an already represented parent contract. |
| Treat direct parent coverage as direct coverage of every descendant. | The tool would hide uncovered refinement requirements and invent Specification-to-row relations. |
| Add an explicit `optional` field now. | No current requirement needs independent optional semantics; non-blocking descendant warnings cover the demonstrated case. |
| Add OR or conditional relation operators now. | No current example requires those semantics, and they would complicate satisfaction rules before a concrete need exists. |
| Allow row parent-child links across USDM records or app namespaces. | Recursive closure would depend on repository-wide child discovery and could change when another app adds or removes rows. Cross-app `usdm_covers` already represents implementation responsibility without distributing one decomposition tree. |
| Add a cross-record `refines` or `derived_from` relation now. | No current traceability need requires row-to-row cross-record refinement; introduce it only when a concrete use case appears. |

## Consequences

- USDM row parsing and validation must accept `RNNN(-NN)*`.
- Every non-top-level row must resolve an immediate parent in the same USDM requirement record.
- Initial numbering is monotonic within each sibling set; removed IDs remain unavailable for reuse.
- Coverage tooling must distinguish direct, derived, blocking-uncovered, and warning-uncovered states.
- Existing direct-coverage response surfaces remain backward-compatible where practical; new effective-coverage states are additive.
- `usdm_covered_by` reports direct covering Specifications only; derived coverage never fabricates `covered_by` entries.
- Requirement decomposition remains record-local; cross-app responsibility uses full-ID `usdm_covers` references.
- A directly covered parent may have uncovered descendants without making the parent uncovered.
- Warning-only uncovered descendants do not make the coverage check fail.
- Tool implementation remains a separate downstream Work Item.
- Existing flat rows and current top-level range expressions remain valid.

## Evidence

- PRODUCT-REQ-SPEC-016 requires hierarchical decomposition and derived coverage.
- PRODUCT-TASK-SPEC-031-01 records decisions D001-D008.
- PRODUCT-TASK-SPEC-031-02 routes the original decision set to this ADR.
- PRODUCT-TASK-SPEC-031-07 records D009-D012 for sibling allocation, additive response shape, record-local hierarchy, and cross-app coverage responsibility.
- `spec:product.design_records.usdm.artifact_format` defines the resulting row hierarchy.
- `spec:product.design_records.usdm.coverage_format` defines direct and effective coverage semantics.
- `spec:product.design_records.usdm.coverage_tools` defines the static reporting contract.
- `spec:product.design_records.authoring_standards.usdm_authoring` defines authoring and stability rules.
