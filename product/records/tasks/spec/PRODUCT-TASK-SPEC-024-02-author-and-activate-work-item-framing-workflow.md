# PRODUCT-TASK-SPEC-024-02: Author and activate Work Item framing workflow

- **id**: PRODUCT-TASK-SPEC-024-02
- **status**: done
- **date**: 2026-07-03
- **work_item**: PRODUCT-WORK-SPEC-024
- **task_type**: authoring
- **estimate**: 1d
- **depends_on**:
  - PRODUCT-TASK-SPEC-024-01
- **outputs**:
  - skills/work-item-framing/SKILL.md
  - skills/work-item-framing/interactive-framing-loop.md
  - skills/work-item-framing/framing-routing.md
  - prompt_chappy.md
  - skills/design-convergence-workflow/SKILL.md
  - skills/design-convergence-workflow/decision-ledger.md
  - skills/design-convergence-workflow/impact-investigation.md
  - skills/design-convergence-workflow/convergence-routing.md
  - skills/design-convergence-workflow/graph-coordination.md
  - skills/design-convergence-workflow/adr-routing.md
  - skills/design-convergence-workflow/design-authoring.md
  - skills/design-convergence-workflow/design-review-gate.md
  - skills/design-convergence-workflow/closure-synchronization.md
  - spec:product.design_records.authoring_standards.task_authoring
  - PRODUCT-TASK-SPEC-024-02
  - PRODUCT-WORK-SPEC-024

## Goal

Author and activate the accepted MVP Work Item framing workflow as one coherent authority update.

## Work

- Read PRODUCT-WORK-SPEC-024 and the completed T01 decision ledger.
- Create `skills/work-item-framing/SKILL.md`.
- Create `skills/work-item-framing/interactive-framing-loop.md`.
- Create `skills/work-item-framing/framing-routing.md`.
- Define Requirement-first entry and the no-Requirement capture route.
- Define the one-question framing decision loop.
- Define source dispositions, unknown routing, and the proceed contract.
- Define conditional same-Work-Item Task materialization by the active framing decision.
- Define the coordination and `work_item_decomposition` boundaries.
- Define direct simple framing closure.
- Activate framing in `prompt_chappy.md` before repository-persistent Work Item planning.
- Amend design convergence to consume an accepted framing route.
- Make design-convergence impact Investigation conditional across every directly affected companion.
- Preserve mandatory integrated independent review inside design convergence.
- Amend Task authoring authority with the bounded framing-decision materialization exception.
- Keep the exception limited to one active framing decision and one uniquely determined same-Work-Item route.
- Inspect the complete scoped diff for contradictory mandatory-Investigation or immediate-start wording.
- Record verification Evidence and directly close PRODUCT-WORK-SPEC-024 when every Done condition is satisfied.

This Task must not:

- introduce unresolved design judgment;
- add a `framing` Task type;
- migrate or repair existing unframed Work Items;
- define `cancelled` lifecycle semantics;
- require an independent framing review;
- perform downstream design or implementation;
- stage or commit changes.

## Done condition

- The three-file `work-item-framing` skill is substantive and internally consistent.
- `prompt_chappy.md` activates framing at the accepted boundary.
- Existing-Requirement and no-Requirement entry routes match T01.
- The interactive loop asks one unresolved framing judgment per user turn.
- Source disposition and proceed-contract routing match T01.
- Formal Investigation is conditional in framing and design convergence.
- The same-Work-Item direct materialization exception is canonical and bounded.
- Coordination remains required for independent graph judgment.
- `work_item_decomposition` remains the owner of downstream Work Item creation.
- Framing has no mandatory independent review or synchronization Task.
- Design convergence still requires its own integrated independent review.
- Existing Work Item migration and repair remain excluded.
- Every directly conflicting design-convergence statement is amended.
- PRODUCT-WORK-SPEC-024 is `done` with complete Evidence.

## Verification

- Inspect every declared output and the scoped Git diff.
- Search `skills/design-convergence-workflow/` for stale immediate-start and mandatory-Investigation statements.
- Confirm conditional wording is consistent across entry, routing, authoring, review, and closure companions.
- Confirm Task authoring allows only the accepted bounded framing exception.
- Confirm the new skill references canonical Requirement, Work Item, and Task authoring authority instead of duplicating their schemas.
- Confirm no existing Work Item or Task record is migrated.
- Confirm no review, coordination, Investigation, synchronization, implementation, stage, or commit work was absorbed.
- Invoke the standalone Task responsibility validator after final Evidence when available.

## Evidence

- PRODUCT-TASK-SPEC-024-01 is complete and fixes D-001 through D-021.
- Created `skills/work-item-framing/SKILL.md`.
- Created `skills/work-item-framing/interactive-framing-loop.md`.
- Created `skills/work-item-framing/framing-routing.md`.
- `prompt_chappy.md` now activates Work Item framing before repository-persistent downstream planning.
- Design convergence now starts from an accepted framing handoff.
- Formal impact Investigation is conditional across the design-convergence entry, routing, authoring, review, and closure authority.
- Mandatory integrated independent review remains unchanged inside design convergence.
- `spec:product.design_records.authoring_standards.task_authoring` now defines the bounded same-Work-Item framing-decision materialization exception.
- Coordination remains required for independent graph judgment.
- `work_item_decomposition` remains the owner of downstream Work Item creation or split.
- The framing skill permits direct simple Work Item closure without a synchronization Task.
- A scoped search found no stale immediate-start or mandatory-Investigation statements in active framing and design-convergence authority.
- A Qwen first-pass comparison of T01 D-001 through D-021 against the three framing skill files returned no findings.
- A Qwen first-pass review of the canonical Task materialization exception returned no findings.
- The complete scoped Git diff was inspected.
- Scoped whitespace validation passed with no findings.
- No existing Work Item or Task was migrated or repaired.
- No formal Investigation, coordination, independent review, synchronization Task, implementation, stage, or commit work occurred.
- PRODUCT-WORK-SPEC-024 was directly closed after every Completion Condition passed.
- Direct parent closure was a framing workflow action with no separate propagation outcome or completion judgment.
- The standalone Task responsibility validator evaluated all common and `authoring` criteria after final Evidence. Every criterion returned `true`.
- DRMCP is non-operational. Filesystem authoring was the required fallback.
