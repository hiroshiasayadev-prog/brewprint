# PRODUCT-TASK-SPEC-018-05: Author and activate design convergence workflow

- **id**: PRODUCT-TASK-SPEC-018-05
- **status**: done
- **date**: 2026-07-01
- **work_item**: PRODUCT-WORK-SPEC-018
- **task_type**: authoring
- **estimate**: 1.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-018-04
- **outputs**:
  - PRODUCT-TASK-SPEC-018-05
  - skills/design-convergence-workflow/SKILL.md
  - skills/design-convergence-workflow/interactive-decision-loop.md
  - skills/design-convergence-workflow/impact-investigation.md
  - skills/design-convergence-workflow/convergence-routing.md
  - skills/design-convergence-workflow/graph-coordination.md
  - skills/design-convergence-workflow/adr-routing.md
  - skills/design-convergence-workflow/design-authoring.md
  - skills/design-convergence-workflow/design-review-gate.md
  - skills/design-convergence-workflow/closure-synchronization.md
  - prompt_chappy.md
  - skills/design-decision-workflow/

## Goal

Author and activate the complete design-convergence workflow skill.

## Work

- Author the successor `SKILL.md` as the entry and phase-routing authority.
- Author one companion file per workflow judgment authority.
- Preserve accepted decisions from `PRODUCT-TASK-SPEC-018-01` and ADRs 009 through 014.
- Point `prompt_chappy.md` to the successor and its phase companions.
- Remove `skills/design-decision-workflow/` only after successor readiness is verified.
- Keep canonical Specification synchronization outside this Task.

## Done condition

- The successor skill defines design-topic intake through reviewed closure.
- Each companion owns one coherent workflow authority.
- The instruction pointer references only the successor skill.
- The old skill directory is removed without a deprecated stub.
- The successor contains no unresolved design judgment.

## Verification

- Confirm all nine successor files exist and contain substantive content.
- Confirm phase routing matches ADRs 009 through 014.
- Confirm decision, coordination, authoring, review, correction, and synchronization remain separate.
- Confirm `prompt_chappy.md` contains no active reference to `skills/design-decision-workflow/`.
- Confirm the old skill directory no longer exists.
- Confirm no canonical Specification file changed.

## Evidence

- Authoring inputs: `PRODUCT-TASK-SPEC-018-01` through `PRODUCT-TASK-SPEC-018-04`.
- Durable authority: `PRODUCT-ADR-SPEC-004` through `PRODUCT-ADR-SPEC-006` and `PRODUCT-ADR-SPEC-009` through `PRODUCT-ADR-SPEC-014`.
- Successor files authored: 9.
- `prompt_chappy.md` now points to `skills/design-convergence-workflow/SKILL.md` and all phase companions.
- `skills/design-decision-workflow/` was removed from the repository path after successor readiness verification.
- The previous directory was moved to ignored local `memory/retired-design-decision-workflow/` because the available filesystem tool has no delete operation; it is not repository authority or a deprecated stub.
- Canonical Specification files were not changed.
- Independent review, closure synchronization, stage, and commit were not performed.
