# PRODUCT-TASK-SPEC-018-14: Author Work Item decomposition Task type

- **id**: PRODUCT-TASK-SPEC-018-14
- **status**: done
- **date**: 2026-07-01
- **work_item**: PRODUCT-WORK-SPEC-018
- **task_type**: authoring
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-018-13
- **outputs**:
  - PRODUCT-TASK-SPEC-018-14
  - PRODUCT-ADR-SPEC-004
  - PRODUCT-ADR-SPEC-005
  - PRODUCT-ADR-SPEC-010
  - spec:product.design_records.authoring_standards.task_authoring
  - spec:product.design_records.authoring_standards.work_item_authoring
  - skills/design-convergence-workflow/SKILL.md
  - skills/design-convergence-workflow/graph-coordination.md
  - skills/design-convergence-workflow/convergence-routing.md
  - skills/design-convergence-workflow/work-item-decomposition.md

## Goal

Project the accepted `work_item_decomposition` Task type into the exact ADR, Specification, and workflow-support targets.

## Work

- Add `work_item_decomposition` to the closed Task taxonomy in ADR-004.
- Separate workflow-graph coordination from Work Item decomposition in ADR-005.
- Add the conditional Work Item decomposition phase to ADR-010.
- Add the new type and adjacent responsibility boundary to `task_authoring`.
- Assign child Work Item creation and split to the new type in `work_item_authoring`.
- Update the design-convergence workflow authority and routing companions.
- Create `work-item-decomposition.md` as the phase companion.
- Preserve every unrelated ADR, Specification, and workflow rule.

This Task must not:

- create or split a Work Item;
- migrate existing records;
- create another Task;
- change Requirement content;
- perform review, synchronization, implementation, stage, or commit work.

## Done condition

- ADR-004, ADR-005, and ADR-010 express the same Task-type boundary.
- `task_authoring` contains the new allowed type, outcome, completion judgment, and adjacent boundary.
- `work_item_authoring` assigns child Work Item creation and split to `work_item_decomposition`.
- The workflow skill routes Task-graph changes to `coordination` and child Work Item decomposition to the new type.
- No migration or follow-up Work Item is introduced.

## Verification

- Confirm every output serves the same Task-type addition.
- Confirm `coordination` no longer owns child Work Item creation or split.
- Confirm `work_item_decomposition` does not own Task-graph changes or child deliverables.
- Confirm ADR, Specification, and workflow-support wording agrees.
- Confirm no additional Task, Work Item, Requirement, or ADR was created.
- Confirm stage and commit were not performed.

## Evidence

- `PRODUCT-ADR-SPEC-004` now defines ten Task types and includes `work_item_decomposition`.
- `PRODUCT-ADR-SPEC-005` now separates coordination, Work Item decomposition, and synchronization.
- `PRODUCT-ADR-SPEC-010` now includes conditional Work Item decomposition as a typed responsibility phase.
- `task_authoring` and `work_item_authoring` now contain the accepted normative boundary.
- `SKILL.md`, `graph-coordination.md`, and `convergence-routing.md` now route the two responsibilities separately.
- `work-item-decomposition.md` defines the bounded phase contract.
- No migration, follow-up Work Item, review, synchronization, implementation, stage, or commit work was performed.
- Result: `PASS`.
