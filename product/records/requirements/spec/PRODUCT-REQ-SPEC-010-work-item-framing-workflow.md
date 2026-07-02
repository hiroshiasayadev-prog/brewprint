# PRODUCT-REQ-SPEC-010: Work Item framing workflow

- **id**: PRODUCT-REQ-SPEC-010
- **status**: accepted
- **date**: 2026-07-03
- **source_refs**:
  - PRODUCT-WORK-SPEC-018
  - spec:product.design_records.authoring_standards.work_item_authoring
  - spec:product.design_records.authoring_standards.task_authoring

## Requirement

Repository-persistent work must establish the intended handling of a Requirement before downstream design or execution planning begins.

The workflow must align the user's desired outcome with the Requirement's Required Outcome.

The workflow must conclude with an explicit source disposition and, when proceeding, an accepted Work Item contract.

## Evidence

- The current design-convergence workflow starts when a design topic is raised.
- The current workflow assumes that the Work Item Goal and downstream design boundary are already understood.
- User and agent understanding may differ about the concrete work required by the same Requirement.
- A formal Investigation is unnecessary when framing can conclude from accepted authority and direct user judgment.
- A speculative full Task graph creates workflow overhead for branches that may never be required.
- Requirement rejection, deferral, formal Investigation, limited downstream research, and direct progression are all valid framing conclusions.

## Required Outcome

- Establish a reusable Work Item framing workflow before downstream design or execution planning.
- Use an existing Requirement as the source for repository-persistent framing.
- When no Requirement exists, identify the Problem and Desired Outcome sufficiently to capture one before framing continues.
- Treat an existing Requirement's Problem as accepted input unless a material conflict is discovered.
- Confirm that the Desired Outcome matches the existing Required Outcome.
- Route a mismatch through Requirement amendment, split, follow-up creation, replacement, rejection, deferral, Investigation, or blocking as appropriate.
- Use one-question-at-a-time interaction for unresolved framing judgments.
- Persist framing judgments in one `decision` Task.
- Do not introduce a `framing` Task type.
- Start the framing Task graph with exactly one `decision` Task.
- Create no speculative Investigation, authoring, decomposition, coordination, review, or synchronization Task.
- Permit the active framing `decision` Task to create and register same-Work-Item Tasks when the selected route uniquely fixes their type, outcome, and dependency.
- Require `coordination` only when graph repair or a separate graph judgment remains.
- Treat formal Investigation as conditional.
- Permit limited downstream research without a formal Investigation when the accepted route remains coherent.
- When proceeding, fix the Work Item Goal, Boundary, Completion Condition, unknown handling, and initial downstream route.
- Create a downstream Work Item through `work_item_decomposition` when a separate Work Item is required.
- Do not require an independent framing review.
- Do not require a synchronization Task for simple framing closure.
- Allow the framing workflow to close its Work Item directly after all required framing outcomes and materialized Tasks are complete.
- Make design convergence one possible downstream route selected by framing.
- Remove mandatory impact Investigation from design convergence when no formal Investigation is required.

## Explicitly Excluded Scope

- Migrating or retroactively legitimizing existing Work Items that did not use framing.
- Repairing an existing unframed Work Item.
- Defining the `cancelled` Work Item or Task lifecycle.
- Performing the downstream design, implementation, or execution selected by framing.
- Requiring an independent review for framing.
- Defining production implementation for an automated framing tool.

## Boundary

PRODUCT owns the reusable framing workflow, its Task ownership, its conditional routing, and its connection to design convergence.

Subsequent design work owns the skill text, canonical Task exception, activation instructions, design-convergence amendments, and implementation choices.
