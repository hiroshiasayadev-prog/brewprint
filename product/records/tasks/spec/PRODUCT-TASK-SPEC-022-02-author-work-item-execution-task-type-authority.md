# PRODUCT-TASK-SPEC-022-02: Author work_item_execution Task-type authority

- **id**: PRODUCT-TASK-SPEC-022-02
- **status**: done
- **date**: 2026-07-02
- **work_item**: PRODUCT-WORK-SPEC-022
- **task_type**: authoring
- **estimate**: 1d
- **depends_on**:
  - PRODUCT-TASK-SPEC-022-01
- **outputs**:
  - PRODUCT-TASK-SPEC-022-02
  - PRODUCT-ADR-SPEC-004
  - PRODUCT-ADR-SPEC-005
  - PRODUCT-ADR-SPEC-010
  - spec:product.design_records.authoring_standards.task_authoring
  - spec:product.design_records.authoring_standards.work_item_authoring
  - skills/design-convergence-workflow/SKILL.md
  - skills/design-convergence-workflow/graph-coordination.md
  - skills/design-convergence-workflow/convergence-routing.md
  - skills/design-convergence-workflow/work-item-decomposition.md
  - skills/design-convergence-workflow/work-item-execution.md
  - prompt_chappy.md

## Goal

Project the decided `work_item_execution` contract into one bounded canonical authority set.

## Work

- Amend ADR-004 with the eleventh closed Task type.
- Amend ADR-005 with the execution relation and adjacent responsibility boundaries.
- Amend ADR-010 with conditional child Work Item execution tracking.
- Add conditional `work_item_ref` metadata and the type contract to Task authoring.
- Add parent hub and child ownership rules to Work Item authoring.
- Add the new phase, routing, graph handoff, and decomposition handoff to design-convergence guidance.
- Create `work-item-execution.md` as the bounded workflow companion.
- Activate the companion from `prompt_chappy.md`.
- Preserve every unrelated accepted rule and existing-record lifecycle.

This Task must not:

- change any T01 decision;
- replace `work_item_decomposition`;
- author W020 checklist assets;
- migrate existing records;
- perform independent review, correction, synchronization, implementation, stage, or commit work.

## Done condition

- Every output expresses the same type name, outcome, completion judgment, and relation.
- `work_item_ref` is required only for `work_item_execution`.
- The child Work Item must be `done` before the execution Task becomes `done`.
- Parent completion remains independently evaluated.
- Decomposition, coordination, execution, and synchronization boundaries do not overlap.
- Child internals remain child-owned.
- No existing-record migration or supersession is introduced.

## Verification

- Compare all outputs with T01 D-001 through D-017.
- Confirm ADR-004, ADR-005, and ADR-010 use the accepted in-place amendment route.
- Confirm the allowed Task-type set contains eleven values.
- Confirm no reverse child relation was added.
- Confirm workflow activation points to the new companion.
- Confirm no checklist, review, correction, synchronization, implementation, stage, or commit work occurred.

## Evidence

### Result

`PASS`.

### Canonical projection

- ADR-004 defines `work_item_execution` as the eleventh closed Task type.
- ADR-005 defines `work_item_ref`, child-status effects, and the non-overlap with decomposition, coordination, and synchronization.
- ADR-010 adds conditional child Work Item execution tracking to the typed design-convergence phases.
- Task authoring defines `work_item_ref` metadata, relation invariants, type semantics, completion rules, and create/update inputs.
- Work Item authoring defines parent ownership, child execution targeting, no reverse field, no automatic parent closure, and no child-internal duplication.
- Design-convergence guidance routes child creation to decomposition and parent-graph child tracking to execution.
- `work-item-execution.md` defines the execution phase contract and stop routes.
- `prompt_chappy.md` activates the new companion.

### Boundary verification

- `work_item_decomposition` still creates or splits child Work Items.
- `coordination` still owns graph changes and may only create or release the execution Task.
- `work_item_execution` owns one existing child completion boundary and no child work.
- `synchronization` still owns only accepted-state propagation.
- No `canceled` lifecycle state was added.
- No existing record was migrated.
- No W020 checklist, independent review, correction, synchronization, implementation, stage, or commit work occurred.
- DRMCP is non-operational, so filesystem authoring was used.
