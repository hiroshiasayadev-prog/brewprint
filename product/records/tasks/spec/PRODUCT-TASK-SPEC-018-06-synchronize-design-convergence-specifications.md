# PRODUCT-TASK-SPEC-018-06: Synchronize design convergence Specifications

- **id**: PRODUCT-TASK-SPEC-018-06
- **status**: done
- **date**: 2026-07-01
- **work_item**: PRODUCT-WORK-SPEC-018
- **task_type**: authoring
- **estimate**: 1d
- **depends_on**:
  - PRODUCT-TASK-SPEC-018-04
- **outputs**:
  - PRODUCT-TASK-SPEC-018-06
  - spec:product.design_records.authoring_standards.requirement_authoring
  - spec:product.design_records.authoring_standards.work_item_authoring
  - spec:product.design_records.authoring_standards.task_authoring
  - spec:product.design_records.authoring_standards.artifact_boundary
  - spec:product.design_records.artifact_model.artifact_responsibility_matrix

## Goal

Project the accepted design-convergence workflow rules into the exact canonical Specification targets.

## Work

- Project the accepted Requirement identity boundary into Requirement authoring rules.
- Project Work Item continuation, shared-writer, integrated-review, and conditional Task rules into Work Item authoring rules.
- Project completed-decision preservation, finding routing, conditional Task materialization, and synchronization boundaries into Task authoring rules.
- Project ADR routing and completed-record ownership into the artifact boundary.
- Project historical Evidence and closure-write ownership into the artifact responsibility matrix.
- Preserve existing artifact taxonomies, lifecycle states, and single-responsibility rules.

This Task does not modify ADRs, skills, reviews, closure state, production implementation, stage, or commit state.

## Done condition

- All accepted rules assigned to the five canonical Specification targets are normative and non-duplicative.
- The stale completed-decision ADR-reference writeback rule is removed.
- Completed-record preservation and finding-driven Task materialization are explicit.
- Shared-writer, integrated-review, and closure-synchronization boundaries are explicit.
- No unresolved design judgment is introduced.

## Verification

- Confirm this Task ID, H1, file name, parent Work Item, and dependency agree.
- Confirm the parent Work Item lists this Task.
- Confirm exactly the five declared Specification targets changed.
- Confirm no completed-decision writeback or `decided` to `recorded` transition remains.
- Confirm correction, finding closure, synchronization, and graph coordination remain separate responsibilities.
- Confirm scoped Git inspection reports no whitespace findings.

## Evidence

- Materialized from the reserved T06 contract in `PRODUCT-WORK-SPEC-018`.
- Authoring dependency: `PRODUCT-TASK-SPEC-018-04`.
- Changed Specification files:
  - `product/records/spec/design-records/authoring-standards/requirement-authoring.md`
  - `product/records/spec/design-records/authoring-standards/work-item-authoring.md`
  - `product/records/spec/design-records/authoring-standards/task-authoring.md`
  - `product/records/spec/design-records/authoring-standards/artifact-boundary.md`
  - `product/records/spec/design-records/artifact-model/artifact-responsibility-matrix.md`
- `requirement-authoring.md` projects D-012 and `PRODUCT-ADR-SPEC-011`.
- `work-item-authoring.md` projects D-013, D-015, D-016, D-020, D-022, and ADRs 011 through 014.
- `task-authoring.md` projects D-010, D-014, D-017 through D-023, and ADRs 006, 013, and 014.
- `artifact-boundary.md` projects D-019, D-021, D-023, `PRODUCT-ADR-SPEC-006`, and `PRODUCT-ADR-SPEC-014`.
- `artifact-responsibility-matrix.md` projects D-019, D-022, D-023, and ADRs 013 and 014.
- Removed the stale rule that required ADR-reference writeback into a completed decision Task.
- Preserved completed decision, authoring, and review records as historical Evidence without downstream progress writeback.
- Added finding-driven conditional correction and independent finding-closure Task materialization.
- Added shared-writer serialization and one final integrated review per Work Item.
- Limited closure synchronization to exact mechanically derivable lifecycle, Evidence, completion-result, and relation propagation.
- No ADR, skill, review, closure, production, stage, or commit change was performed.
- Scoped `git.inspect_diff` reported only this Task, the parent Work Item, and the five declared Specification targets.
- Scoped `git.inspect_worktree` reported no whitespace findings.
