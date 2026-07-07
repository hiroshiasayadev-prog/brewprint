# PRODUCT-TASK-SPEC-023-09: Author cancelled lifecycle canonical contract

- **id**: PRODUCT-TASK-SPEC-023-09
- **status**: done
- **date**: 2026-07-03
- **work_item**: PRODUCT-WORK-SPEC-023
- **task_type**: authoring
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-023-08
- **outputs**:
  - spec:product.design_records.authoring_standards.task_authoring
  - spec:product.design_records.authoring_standards.work_item_authoring
  - spec:product.responsibility_boundary_validator
  - skills/design-convergence-workflow/work-item-execution.md
  - skills/design-convergence-workflow/graph-coordination.md
  - skills/design-convergence-workflow/closure-synchronization.md
  - PRODUCT-TASK-SPEC-023-09

## Goal

Project the accepted cancellation decisions and ADR authority into one coherent canonical and workflow-support contract.

## Work

- Update Task authoring lifecycle, section-readiness, dependency, Evidence, `work_item_execution`, continuation, and validator-invocation rules.
- Update Work Item authoring lifecycle, section-readiness, propagation, Evidence, and continuation rules.
- Update the responsibility-boundary validator so its final invocation applies before `done`, not cancellation.
- Update Work Item execution and graph coordination support for cancelled child and prerequisite behavior.
- Narrow closure synchronization wording so reviewed-success closure means `done`, not cancellation.
- Preserve checklist artifacts, Requirement content, existing records, and concrete DRMCP mechanics unchanged.
- Record exact changed sections and verification Evidence.

This Task must not:

- reopen accepted decisions or ADR routes;
- amend ADR files;
- change the Task graph;
- modify checklist assets;
- implement cancellation mechanics;
- perform independent review, correction, synchronization, stage, or commit work.

## Done condition

- All T01 and T04 decisions are expressed in the correct current normative owner.
- ADR and Specification roles remain distinct.
- Workflow-support files agree with the canonical lifecycle contract.
- No stale four-status-only or child-cancellation contradiction remains in the writable boundary.
- T10 has one exact integrated review set.

## Verification

- Confirm only the three routed Specifications, three workflow-support files, and this Task changed.
- Confirm `cancelled` is terminal and irreversible in both authoring Specifications.
- Confirm atomic lifecycle semantics are stated without concrete command or DRMCP design.
- Confirm no checklist, Requirement, existing workflow record, implementation, stage, or commit changed.
- Confirm scoped whitespace and diff inspection pass.

## Evidence

- Updated Task authoring with `cancelled` lifecycle, terminal section readiness, dependency effects, atomic cancellation ownership, new-record resumption, cancelled-child execution behavior, and successful-completion-only post-Evidence validation.
- Updated Work Item authoring with `cancelled` lifecycle, terminal section readiness, owned-Task propagation, atomic cancellation ownership, child-execution effects, and new Work Item resumption.
- Updated the responsibility-boundary validator so its final invocation applies before `done` and not before `cancelled`.
- Updated Work Item execution for the child-cancelled terminal branch.
- Updated graph coordination to preserve cancelled records and route child cancellation without owning lifecycle writes.
- Updated closure synchronization so reviewed-success closure sets `done` and never performs cancellation.
- Checklist assets, Requirement content, existing workflow records outside W023, and DRMCP implementation remained unchanged.
- Scoped Git whitespace inspection passed for all T08 and T09 targets. LF-to-CRLF advisories are non-blocking repository normalization warnings.
- No independent review, correction, lifecycle synchronization, implementation, stage, or commit work occurred.
- DRMCP is non-operational. Filesystem authoring was the required fallback.
- Result: `PASS`.
