# PRODUCT-REQ-SPEC-008: Review-local mechanical correction

- **id**: PRODUCT-REQ-SPEC-008
- **status**: accepted
- **date**: 2026-07-02
- **source_refs**:
  - PRODUCT-TASK-SPEC-021-12
  - PRODUCT-ADR-SPEC-013
  - spec:product.design_records.authoring_standards.task_authoring

## Requirement

Design review workflows must avoid separate correction and closure-review Tasks for defects whose repair is mechanical and uniquely determined by accepted authority.

A reviewer must be able to repair an eligible defect within the active review and then evaluate the corrected final state.

The exception must preserve reviewer independence and must not allow the reviewer to make or approve new design judgment.

## Evidence

- PRODUCT-TASK-SPEC-021-12 found one metadata-to-prose impact relation mismatch.
- The required repair followed directly from the accepted Work Item authoring contract.
- The repair changed no design decision, semantic contract, ownership boundary, completion condition, or Task graph.
- The existing workflow required finding-specific coordination, correction, and independent closure review for the mechanical repair.
- The additional workflow records and review cycle outweighed the risk and complexity of the defect.
- PRODUCT-ADR-SPEC-013 currently routes every formal `NEEDS REVISION` finding through separate correction and closure-review Tasks.
- The current Task authoring contract prohibits review and correction overlap without a bounded mechanical exception.

## Required Outcome

- Define a bounded review-local correction exception.
- Permit the active reviewer to repair only defects with one mechanically derivable outcome.
- Require accepted authority to determine the repair without new user or design judgment.
- Limit eligible repairs to non-semantic projection, reference, metadata, formatting, or equivalent consistency defects.
- Require the reviewer to inspect the corrected final state before issuing the review verdict.
- Require review Evidence to record the defect, exact repair, authority, changed artifacts, and post-repair verification.
- Preserve the normal `NEEDS REVISION` route for Blocking and Major findings.
- Preserve the normal route when several materially valid repairs exist.
- Preserve the normal route when the repair changes Requirement intent, ADR decisions, normative behavior, ownership, completion boundaries, Task graph, dependencies, lifecycle, or release conditions.
- Prevent a reviewer from repairing and self-closing a finding after the finding and verdict have been finalized as workflow history.
- Keep correction and independent finding-closure review available for every defect outside the bounded exception.
- Define consistent rules across review workflow authority and Task responsibility contracts.

## Explicitly Excluded Scope

- Allowing reviewers to make new design decisions.
- Allowing reviewers to change accepted alternatives or rationale.
- Allowing review Tasks to perform unrestricted authoring or correction.
- Allowing self-closure of already finalized findings.
- Weakening independent review for semantic, architectural, ownership, lifecycle, or release changes.
- Defining one concrete tool or automated implementation for applying corrections.
- Retroactively rewriting completed review history.

## Boundary

PRODUCT owns the workflow requirement for bounded review-local mechanical correction and preservation of review independence.

Subsequent design work owns eligibility rules, verdict timing, Evidence shape, ADR disposition, Specification changes, skill changes, migration, and compatibility with completed review records.
