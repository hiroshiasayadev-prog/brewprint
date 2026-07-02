# PRODUCT-TASK-SPEC-019-17: Author Task validator usage rule

- **id**: PRODUCT-TASK-SPEC-019-17
- **status**: done
- **date**: 2026-07-01
- **work_item**: PRODUCT-WORK-SPEC-019
- **task_type**: authoring
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-019-16
  - PRODUCT-TASK-SPEC-018-19
- **outputs**:
  - PRODUCT-TASK-SPEC-019-17
  - spec:product.design_records.authoring_standards.task_authoring

## Goal

Add the narrow semantic-validator ownership and usage rule to the canonical Task authoring Specification.

## Work

- Read and preserve the accepted W018 `task_type`, coordination, and Work Item decomposition semantics.
- State that the dedicated validator Specification owns semantic validator behavior.
- State that current DRMCP retains structural parsing, validation, diagnostics, indexing, and current tool projections.
- Require the authoring workflow to invoke the validator immediately after Task authoring.
- Require the completion or release workflow to invoke the same validator after final Evidence is written.
- Keep workflow enforcement and human violation acceptance outside the validator.
- Reference the dedicated validator semantic ref instead of duplicating its contract.
- Preserve all unrelated Task authoring rules.
- Record exact changed sections and verification Evidence in this Task.

This Task must not:

- redefine validator criteria, result shape, aggregation, or failure semantics;
- change the accepted Task-type taxonomy;
- change coordination or Work Item decomposition ownership;
- author ADR, Requirement, checklist, or implementation content;
- perform review, correction, synchronization, stage, or commit work.

## Done condition

- `task_authoring` points to the dedicated validator Specification as the semantic behavior owner.
- The two required invocation points are explicit.
- Current DRMCP structural ownership remains distinct from standalone semantic evaluation.
- Accepted W018 Task-type and decomposition semantics remain unchanged.
- No validator behavior is duplicated into `task_authoring`.

## Verification

- Compare the update against PRODUCT-ADR-SPEC-016, PRODUCT-ADR-SPEC-017, and the validator Specification.
- Confirm PRODUCT-TASK-SPEC-018-19 closed F-BLK-01 and F-MAJ-01 before this writer runs.
- Confirm the shared-writer update preserves current full text from the accepted W018 repair.
- Confirm no unrelated `task_authoring` section changed.
- Confirm no ADR, Requirement, checklist, implementation, review, correction, synchronization, stage, or commit work occurred.

## Evidence

### Result

- Result: `PASS`.
- `task_authoring` now points to `spec:product.responsibility_boundary_validator` as the semantic behavior owner.
- The authoring workflow must invoke the standalone validator after Task authoring.
- The Task completion or release workflow must invoke the same validator after final Task Evidence.
- Semantic violations route to explicit human acceptance or rejection.
- The validator does not enforce correction, completion, release, acceptance, or rejection.

### Shared-writer gate

- T07 R-004 fixed the external shared-writer gate.
- PRODUCT-TASK-SPEC-018-19 is `done` with result `PASS`.
- PRODUCT-TASK-SPEC-018-19 closes F-BLK-01 and F-MAJ-01.
- The accepted W018 Task-type, coordination, and Work Item decomposition semantics were preserved.

### Changed sections

- `spec:product.design_records.authoring_standards.task_authoring`:
  - added `### Semantic responsibility validation usage` under `## Rules`;
  - added PRODUCT-ADR-SPEC-016, PRODUCT-ADR-SPEC-017, and `spec:product.responsibility_boundary_validator` to `## Related specs`.
- `PRODUCT-TASK-SPEC-019-17`:
  - changed lifecycle state to `done`;
  - recorded authoring and verification Evidence.

### Boundary verification

- Validator criteria, checklist composition, result shape, aggregation, and failure semantics were not duplicated.
- Current DRMCP structural ownership remains explicit and unchanged.
- No accepted Task-type value or type contract row changed.
- No coordination, Work Item decomposition, continuation, or reconvergence rule changed.
- Scoped Git diff inspection was complete and non-truncated.
- Scoped whitespace verification returned `PASS`.
- No ADR, Requirement, checklist, implementation, independent review, correction, synchronization, stage, or commit work occurred.
- DRMCP is non-operational, so filesystem authoring was used.
