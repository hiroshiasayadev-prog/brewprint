# PRODUCT-TASK-SPEC-026-03: Create term-inventory Work Item

- **id**: PRODUCT-TASK-SPEC-026-03
- **status**: done
- **date**: 2026-07-04
- **work_item**: PRODUCT-WORK-SPEC-026
- **task_type**: work_item_decomposition
- **estimate**: 0.25d
- **depends_on**:
  - PRODUCT-TASK-SPEC-026-02
- **outputs**:
  - PRODUCT-WORK-SPEC-027
  - PRODUCT-TASK-SPEC-027-01

## Goal

Create one downstream Work Item for the Brewprint design-governance term inventory from the accepted framing contract.

## Work

- Create one downstream Work Item sourced directly from PRODUCT-REQ-SPEC-013.
- Use the Goal, Boundary, Completion Condition, unknown handling, and initial route fixed by PRODUCT-TASK-SPEC-026-01.
- Materialize exactly one initial `decision` Task in the child Work Item.
- Leave exact corpus scope, output placement, batch partitioning, parallel extraction ownership, and coverage accounting to that decision Task.
- Do not create extraction Tasks before those decisions are fixed.

## Done condition

One downstream inventory Work Item exists with one initial decision Task and no speculative extraction, aggregation, classification, or helper-tool Tasks.

## Verification

- Confirm the child Work Item cites PRODUCT-REQ-SPEC-013 as its direct source.
- Confirm the child Goal and Completion Condition end at JSONL observation gathering and coverage.
- Confirm aggregation, helper tooling, validation tooling, classification, definition, and use-case extraction are excluded.
- Confirm the child Task graph contains exactly one initial decision Task.

## Evidence

- The accepted framing contract and Requirement amendment are complete.
- Created PRODUCT-WORK-SPEC-027 with direct sources PRODUCT-REQ-SPEC-013 and PRODUCT-TASK-SPEC-026-03.
- Created PRODUCT-TASK-SPEC-027-01 as the child's only initial Task.
- The child Work Item owns JSONL observation gathering and coverage only; aggregation, helper tooling, validation tooling, classification, definition, and use-case extraction remain excluded.
- No extraction Task was materialized.
- DRMCP is non-operational. Filesystem authoring is the required fallback.
