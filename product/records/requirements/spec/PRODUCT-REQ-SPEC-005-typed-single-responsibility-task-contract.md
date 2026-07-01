# PRODUCT-REQ-SPEC-005: Typed single-responsibility Task contract

- **id**: PRODUCT-REQ-SPEC-005
- **status**: accepted
- **date**: 2026-06-30
- **source_refs**:
  - spec:product.design_records.authoring_standards.task_authoring
  - spec:product.design_records.authoring_standards.artifact_boundary
  - spec:product.design_records.artifact_model.artifact_responsibility_matrix
- **work_items**:
  - PRODUCT-WORK-SPEC-016

## Requirement

Every Task must declare one primary Task type.

Every Task must own one primary responsibility that matches the declared Task type.

A Task must remain closeable through one completion judgment after its declared dependencies are satisfied.

A Task must not combine independent responsibilities such as decision, authoring, implementation, review, correction, verification, coordination, or synchronization.

Task responsibility is determined by owned outcomes and completion judgment, not by changed-file count.

## Evidence

- The current Task authoring contract does not define a structured primary Task type.
- The current contract only recommends splitting Tasks that contain multiple independent concerns.
- Mixed-responsibility Tasks blur writer ownership, review independence, verification ownership, and completion authority.
- Review and correction require different writers and different completion judgments.
- Coordination and synchronization manage workflow state but must not absorb child execution responsibilities.
- A single implementation responsibility may legitimately modify production code, focused tests, and fixtures together.

## Required Outcome

- Define one required primary Task type field for Task records.
- Define a closed set of allowed primary Task type values.
- Permit exactly one primary Task type per Task.
- Define each Task type by its owned outcome and prohibited responsibility overlaps.
- Require `Goal`, `Work`, `Done condition`, and `Verification` to match the declared Task type.
- Require each Task to express one primary outcome and one completion judgment.
- Split independent responsibilities into separate Tasks with explicit dependencies.
- Prohibit authoring and independent review within the same Task.
- Prohibit correction and independent finding closure within the same Task.
- Prohibit unresolved design decisions inside implementation Tasks.
- Prohibit coordination Tasks from producing child Task or child Work Item deliverables directly.
- Prohibit synchronization Tasks from introducing new design, implementation, review, or correction work.
- Allow multiple changed files when all changes serve one primary outcome and verification boundary.
- Preserve Task `work_item` as the membership relation to its owning Work Item.

## Explicitly Excluded Scope

- Validator implementation or model selection.
- Validation checklist format or diagnostic schema.
- Agent model routing.
- Existing Task migration sequencing.
- Exact type-specific metadata beyond the primary Task type.
- Work Item hierarchy or workflow provenance relations.

## Boundary

PRODUCT owns the Task responsibility model, primary Task type semantics, and Task authoring contract.

DRMCP owns parsing, validation, diagnostics, indexing, and tool projections for the accepted contract.

Work Items own workflow graphs and dependency ordering. Tasks own only their declared concrete responsibility and evidence.
