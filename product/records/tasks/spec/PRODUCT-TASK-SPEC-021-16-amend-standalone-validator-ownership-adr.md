# PRODUCT-TASK-SPEC-021-16: Amend standalone validator ownership ADR

- **id**: PRODUCT-TASK-SPEC-021-16
- **status**: done
- **date**: 2026-07-02
- **work_item**: PRODUCT-WORK-SPEC-021
- **task_type**: authoring
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-021-09
- **outputs**:
  - PRODUCT-TASK-SPEC-021-16
  - PRODUCT-ADR-SPEC-016

## Goal

Amend PRODUCT-ADR-SPEC-016 with the non-material TRV ownership clarification selected by corrected T08 B-001.

## Work

- Preserve PRODUCT ownership of the semantic validator contract.
- Preserve separation from current DRMCP.
- Identify `TRV` as the app-local design and later implementation owner.
- Remove the stale consequence that W021 owns concrete standalone implementation delivery.
- State that W021 owns reviewed PRODUCT conceptual design only.
- Preserve future DRMCP integration as separately decidable work.
- Preserve the existing selected alternative, core architecture, and rationale.
- Treat the W021 correction history as workflow Evidence, not ADR authority.
- Keep the existing ADR status and non-superseded lifecycle.

This Task must not:

- create another ADR;
- reverse or supersede `PRODUCT-ADR-SPEC-016`;
- alter PRODUCT semantic validation behavior;
- author or modify a Specification, namespace profile, Work Item, or Task graph;
- perform review, synchronization, implementation, stage, or commit work.

## Done condition

- `PRODUCT-ADR-SPEC-016` remains `accepted` and non-superseded.
- The core standalone-versus-current-DRMCP decision remains unchanged.
- TRV app-local ownership is explicit.
- The stale W021 implementation-delivery consequence is removed.
- W021 conceptual-design ownership is explicit.
- No circular ADR dependency is introduced.
- No prohibited artifact changed.

## Verification

- Confirm the amendment follows the ADR authoring standard for a non-material change.
- Confirm the ADR date changes only when required by the accepted authoring standard.
- Confirm `supersedes` remains empty.
- Confirm no new ADR or dependency was introduced for the W021 correction history.
- Confirm current DRMCP ownership and future-integration exclusions remain intact.
- Confirm no Specification, namespace, Work Item, Task graph, or implementation artifact changed.

## Evidence

- Corrected T08 B-001 selected `PRODUCT-ADR-SPEC-016` with disposition `amend` and materiality `non-material`.
- T09 is the direct graph predecessor.
- DRMCP is non-operational, so filesystem authoring was used.
- Updated `PRODUCT-ADR-SPEC-016` only.
- Preserved the accepted standalone-versus-current-DRMCP decision.
- Preserved PRODUCT ownership of `spec:product.responsibility_boundary_validator`.
- Added `TRV` as the app-local design and later implementation owner.
- Replaced the stale W021 implementation-delivery consequence with PRODUCT conceptual-design and namespace-bootstrap ownership.
- Kept future DRMCP integration separately decidable.
- Kept `status: accepted` and `supersedes: []`.
- Updated the ADR date to 2026-07-02 because the documented ownership consequence changed.
- No new ADR, dependency, Specification, namespace profile, Work Item, Task graph, implementation, review, synchronization, stage, or commit work was performed.

### Verification result

- Selected alternative: unchanged.
- Core architecture: unchanged.
- Core rationale: unchanged.
- TRV ownership: explicit.
- Stale W021 implementation consequence: removed.
- Circular ADR dependency: none.
- Prohibited artifact changes: none.
