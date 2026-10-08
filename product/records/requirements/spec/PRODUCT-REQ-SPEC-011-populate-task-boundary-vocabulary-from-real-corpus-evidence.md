# PRODUCT-REQ-SPEC-011: Populate task-boundary-vocabulary from real corpus evidence

- **id**: PRODUCT-REQ-SPEC-011
- **status**: accepted
- **date**: 2026-07-03
- **source_refs**:
  - TRV-INV-SPEC-004
  - TRV-INV-SPEC-005
  - PRODUCT-TASK-SPEC-021-13

## Requirement

Populate `skills/task-boundary-vocabulary/` task_type files with real, corpus-verified boundary-violation vocabulary entries, beyond the currently populated `decision.md`.

## Evidence

- TRV-INV-SPEC-004 and TRV-INV-SPEC-005 found that qwen3.6-27b, gpt-oss:20b, and Claude Haiku 4.5 consistently missed T04 (Specification-authoring impersonation) and T06 (independent-review impersonation) boundary violations, regardless of reasoning effort or denial-text presence.
- Five independent interventions (grouping, desensitization, denial-override correction, thinking-on, and others) converged on the same noise-versus-detection trade-off.
- `skills/DEPRECATED-task-responsibility-boundary-validator/` records the abandoned automated-judgment route and retains its checklist as reference material.
- `skills/task-boundary-vocabulary/decision.md` already contains corpus-verified entries for the `decision` task_type only.

## Required Outcome

- At least one additional task_type file in `skills/task-boundary-vocabulary/` (beyond `decision.md`) contains corpus-verified boundary-vocabulary entries.
- Each entry traces to a real Task ID in `product/records/tasks/spec/`.
- Low-confidence entries are isolated in an `## Open questions` section rather than mixed into confirmed content.
- `spec:product.design_records.authoring_standards.task_authoring` permits a lightweight Evidence-only route for an `investigation` Task when a full Investigation record is disproportionate and an explicit user judgment accepts the trade-off.
- A general, per-task_type vocabulary reference (normal in-type phrasing, not boundary violations) exists, built from the same corpus scan, as a separate artifact from the boundary-violation dictionary.

## Explicitly Excluded Scope

- Rewriting existing Task records.
- Reimplementing or re-evaluating TRV as an automated validator.
- Full ASD-STE100 adoption.
- Production lint implementation.
- A fixed 12-field collection schema.
- Formal Investigation records authored in parallel per responsibility cluster.
- Reconciliation of canonical term definitions (deferred to a later Work Item after roughly 30 log entries accumulate).

## Boundary

This Requirement owns delivery of corpus-verified boundary-vocabulary entries and the lightweight investigation-Evidence authoring-standard exception that enables collecting them without disproportionate overhead.

This Requirement does not own TRV application delivery (TRV-REQ-SPEC-001) or final canonical-term reconciliation.
