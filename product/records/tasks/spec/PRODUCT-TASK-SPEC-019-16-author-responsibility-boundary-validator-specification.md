# PRODUCT-TASK-SPEC-019-16: Author responsibility-boundary validator Specification

- **id**: PRODUCT-TASK-SPEC-019-16
- **status**: done
- **date**: 2026-07-01
- **work_item**: PRODUCT-WORK-SPEC-019
- **task_type**: authoring
- **estimate**: 1d
- **depends_on**:
  - PRODUCT-TASK-SPEC-019-15
- **outputs**:
  - PRODUCT-TASK-SPEC-019-16
  - PRODUCT-ADR-SPEC-001
  - spec:product.responsibility_boundary_validator
  - spec:product

## Goal

Author the canonical standalone PRODUCT Specification for semantic Task responsibility-boundary validation and align its direct PRODUCT-root ownership.

## Work

- Amend PRODUCT-ADR-SPEC-001 to recognize `responsibility-boundary-validator/` as a direct standalone PRODUCT spec area without assigning it to `design-records/`.
- Create `spec:product.responsibility_boundary_validator` at `product/records/spec/responsibility-boundary-validator/index.md`.
- Register the new topic directly under `spec:product`.
- Project PRODUCT-ADR-SPEC-015 through PRODUCT-ADR-SPEC-017 into current normative rules.
- Define one-Task input and Task-local evidence boundaries.
- Define common plus `task_type`-specific checklist composition and automatic selection.
- Define criterion-level binary results, concise reasons, section references, and optional excerpts.
- Define logical-AND overall aggregation.
- Define structural precondition failure, semantic non-compliance, and execution failure as separate outcomes.
- Define the temporary standalone ownership boundary and exclude current DRMCP integration.
- Define post-authoring and post-Evidence invocation semantics plus human violation judgment.
- State that exact checklist wording and storage belong to PRODUCT-WORK-SPEC-020.
- State that concrete implementation belongs to PRODUCT-WORK-SPEC-021.
- Preserve runtime, provider, retry, timeout, concrete field names, and interaction mechanisms outside this Specification.
- Record exact changed sections and verification Evidence in this Task.

This Task must not:

- author checklist content;
- implement the validator;
- modify current DRMCP Specifications or source;
- modify `task_authoring`;
- reopen PRODUCT-ADR-SPEC-015 through PRODUCT-ADR-SPEC-017 or Requirement decisions;
- alter the ownership of existing `design-records/`, `brewprint/`, or `bpdsl/` areas;
- perform review, correction, synchronization, stage, or commit work.

## Done condition

- PRODUCT-ADR-SPEC-001 recognizes the standalone validator as a direct PRODUCT-root area.
- The validator Specification exists at the accepted semantic ref and physical target.
- The parent PRODUCT topic map registers the child exactly once.
- `spec:product.design_records` does not register or own the validator.
- Every retained T12 canonical projection after the T15 no-amendment disposition reaches one normative owner.
- ADR rationale and Specification current-state text remain distinct.
- The Specification separates semantic result, structural precondition failure, and execution failure.
- W020 and W021 boundaries are explicit.
- No exact checklist wording or implementation choice is introduced.

## Verification

- Trace T01, T03, T07, T09, ADR-015 through ADR-017, and the T15 no-amendment disposition into the new Specification.
- Confirm D-008 and D-011 direct projections are present without creating extra ADR semantics.
- Confirm the PRODUCT-ADR-SPEC-001 amendment preserves existing area ownership while adding the standalone validator area.
- Confirm the document shape and PRODUCT-parent registration satisfy the canonical Spec format.
- Confirm `spec:product.design_records` does not register the validator.
- Confirm current DRMCP artifacts and `task_authoring` are unchanged.
- Confirm no checklist, implementation, review, correction, synchronization, stage, or commit work occurred.

## Evidence

### Result

- Result: `PASS`.
- PRODUCT-ADR-SPEC-001 now recognizes `responsibility-boundary-validator/` as a direct PRODUCT-root area.
- `spec:product.responsibility_boundary_validator` exists at `product/records/spec/responsibility-boundary-validator/index.md`.
- `spec:product` registers the validator exactly once.
- `spec:product.design_records` does not register or own the validator.

### Canonical projection

- PRODUCT-ADR-SPEC-015 supplies Task-local Evidence, checklist composition, binary criterion results, logical-AND aggregation, and outcome separation.
- PRODUCT-ADR-SPEC-016 supplies standalone PRODUCT ownership and excludes current DRMCP integration.
- PRODUCT-ADR-SPEC-017 supplies both invocation points and human-owned violation disposition.
- T01 D-008 is projected as required section references, optional excerpts, and no required line numbers.
- T01 D-011 is projected by omitting required checklist revision and stable criterion identifiers from the external result contract.
- The T15 no-amendment disposition preserves PRODUCT-REQ-SPEC-007 unchanged.

### Changed artifacts

- `PRODUCT-ADR-SPEC-001`: metadata, top-level PRODUCT areas, root routing, ownership, dependency direction, rationale, alternatives, consequences, and Evidence.
- `spec:product`: metadata, current area contract, Topics, placement rules, dependency direction, and authoring boundary.
- `spec:product.responsibility_boundary_validator`: new standalone validator Overview and normative contract.
- `PRODUCT-TASK-SPEC-019-16`: lifecycle and authoring Evidence.

### Verification

- The new Spec ID matches its path-derived canonical ref.
- The child `parent` marker matches the authoritative `spec:product` Topics row.
- The direct validator ref appears once in the parent Topics table and once as the child document ID.
- No old `spec:product.design_records.responsibility_boundary_validator` ref or Design Records physical target remains under `product/records`.
- The validator Specification separates structural precondition failure, semantic evaluation, and execution failure.
- PRODUCT-WORK-SPEC-020 and PRODUCT-WORK-SPEC-021 retain checklist and implementation ownership.
- Current DRMCP artifacts and `task_authoring` were not changed.
- No checklist content, implementation, independent review, correction, synchronization, stage, or commit work occurred.
- Scoped Git diff inspection was complete and non-truncated.
- Scoped whitespace verification returned `PASS`.
- DRMCP is non-operational, so filesystem authoring was used.
