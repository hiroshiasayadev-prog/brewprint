# PRODUCT-REQ-SPEC-002: Migrate artifact authoring guides to product namespace

- **id**: PRODUCT-REQ-SPEC-002
- **status**: accepted
- **date**: 2026-06-23
- **source_refs**:
  - PRODUCT-INV-SPEC-003
- **work_items**:
  - PRODUCT-WORK-SPEC-010
  - PRODUCT-WORK-SPEC-011

## Requirement

Artifact authoring guides in `v01/records/guides/` use the v01 format and predate the v0.2.0 namespace model. Product namespace records have no authoritative authoring guides. Each authoring guide must be migrated to the product namespace, adapted for the product record format, and extended to cover v0.2.0 namespace concepts — including the app namespace — where authoring rules change.

## Evidence

- `v01/records/guides/work-item-authoring.md` defines `source_requirement` and `impact_refs`/`tasks`; product work items in practice use fields such as `requirement_refs`, `source_work_items`, `investigation_refs`, and `task_refs`. No authoritative product guide documents this format; the format is inferred from existing files.
- v01 guides predate the v0.2.0 namespace model (product / drmcp / bpdsl / app). Authoring rules that are namespace-sensitive have no product-side documentation.
- `v01/records/guides/` is a read-only snapshot and cannot be updated to reflect product namespace rules.
- The app namespace is a v0.2.0 concept with no v01 counterpart; authoring rules for app-namespace artifacts are undocumented.
- PRODUCT-WORK-SPEC-010 (writing standard) completed.
- PRODUCT-WORK-SPEC-011 (per-artifact authoring guides) completed. Published guides for ADR, requirement, work item, task, and investigation under `spec:product.concepts.authoring_standards`; spec-authoring coverage at `product/records/guides/spec-authoring.md` verified complete.

## Required Outcome

- Product namespace authoring guides exist for: spec, ADR, investigation, requirement, work-item, task.
- Each guide documents the product record format (field schema, section structure, validation rules).
- Each guide covers v0.2.0 namespace concepts where authoring rules differ from v01.
- App namespace authoring rules are documented where they differ from other namespaces.
- Spec authoring guide coverage from PRODUCT-WORK-SPEC-003 is verified complete, or remaining gaps are addressed.

## Explicitly Excluded Scope

- Artifact boundary definitions (which artifacts belong in which namespace).
- `v01/records/guides/artifact-boundary.md` migration.
- Bulk rewrite of existing records to match updated guides.
- Authoring guide tooling or validation implementation.

## Boundary

This REQ owns the requirement for product-namespace authoring guides scoped to authoring rules only. Namespace boundary definitions, artifact placement decisions, and tooling are not in scope.
