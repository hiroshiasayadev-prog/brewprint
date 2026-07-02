# PRODUCT-WORK-SPEC-024: Establish Work Item framing workflow

- **id**: PRODUCT-WORK-SPEC-024
- **status**: done
- **date**: 2026-07-03
- **source_refs**:
  - PRODUCT-REQ-SPEC-010
- **impact_refs**:
  - spec:product.design_records.authoring_standards.requirement_authoring
  - spec:product.design_records.authoring_standards.work_item_authoring
  - spec:product.design_records.authoring_standards.task_authoring
- **tasks**:
  - PRODUCT-TASK-SPEC-024-01
  - PRODUCT-TASK-SPEC-024-02

## Goal

Establish and activate one reusable Work Item framing workflow.

The workflow must align Requirement intent, decide source disposition, and fix an actionable Work Item contract before downstream routing.

## Boundary

This Work Item owns:

- the MVP framing workflow contract;
- one-question framing interaction;
- Requirement-to-framing entry rules;
- source disposition and Work Item contract decisions;
- conditional unknown and Investigation routing;
- conditional same-Work-Item Task materialization by the active framing decision;
- the boundary between direct materialization and coordination;
- direct simple framing closure;
- activation of the framing skill;
- the framing-to-design-convergence handoff;
- removal of mandatory impact Investigation from design convergence.

This Work Item does not own:

- migration or repair of existing unframed Work Items;
- `cancelled` lifecycle design;
- downstream design or implementation selected by framing;
- a production framing tool;
- mandatory independent framing review;
- stage or commit work.

## Impact Scope

| target | impact |
|---|---|
| `skills/work-item-framing/` | Add the framing entry skill, interactive loop, and routing authority. |
| `prompt_chappy.md` | Require framing before repository-persistent Work Item planning and route design convergence after framing. |
| `skills/design-convergence-workflow/` | Start from an accepted framing route and make formal impact Investigation conditional. |
| `spec:product.design_records.authoring_standards.task_authoring` | Add the bounded framing-decision Task materialization exception. |
| Requirement and Work Item authoring authority | Reuse existing Problem, Required Outcome, Goal, Boundary, and Completion Condition ownership. |

## Task flow

```text
PRODUCT-TASK-SPEC-024-01 framing workflow decision
  -> PRODUCT-TASK-SPEC-024-02 author and activate framing workflow
  -> direct Work Item closure after T02 verification
```

T01 directly materialized T02 because the accepted decision fixed one authoring outcome and one dependency.

No formal Investigation, independent review, coordination, or closure synchronization Task is required by the accepted MVP route.

## Task Candidates

| task | task type | responsibility | dependency |
|---|---|---|---|
| `PRODUCT-TASK-SPEC-024-01` | `decision` | Fix the MVP framing contract and exact conditional authoring route. | none |
| `PRODUCT-TASK-SPEC-024-02` | `authoring` | Author and activate the accepted framing workflow and required authority amendments. | T01 |

No additional Task is materialized speculatively.

## Completion Condition

- A reusable `work-item-framing` skill exists and is active.
- Existing-Requirement and no-Requirement entry routes are explicit.
- The framing decision loop owns source disposition and conditional Work Item contract decisions.
- The initial framing graph contains one decision Task only.
- Additional Tasks are created only when the active decision determines they are required.
- The bounded same-Work-Item direct materialization exception is canonical.
- Formal Investigation and independent review are not mandatory framing phases.
- Simple framing closure does not require a synchronization Task.
- Design convergence begins only when framing selects that route.
- Design-convergence impact Investigation is conditional.
- Existing Work Item migration and repair remain excluded.
- T02 verifies the complete authoring boundary and directly closes this Work Item.

## Evidence

- PRODUCT-REQ-SPEC-010 is accepted and defines the framing workflow requirement.
- The user accepted a separate framing skill before design convergence.
- The user accepted Requirement-first repository persistence.
- The user accepted one decision Task as the initial graph.
- The user accepted decision-driven conditional Task materialization.
- The user accepted no mandatory formal Investigation, independent review, or closure synchronization Task.
- The user excluded existing Work Item migration and repair.
- T01 records the complete accepted MVP decision set.
- T01 directly materialized T02 and updated this Work Item.
- `skills/work-item-framing/` now contains the active three-file MVP workflow.
- `prompt_chappy.md` activates framing before repository-persistent downstream Work Item planning.
- Design convergence now starts from an accepted framing handoff.
- Design-convergence impact Investigation is conditional.
- Integrated independent review remains mandatory for design convergence.
- The bounded same-Work-Item direct materialization exception is canonical in Task authoring authority.
- T02 completed the full authoring boundary and passed scoped diff and whitespace inspection.
- Qwen first-pass consistency checks returned no findings.
- Existing Work Item migration and repair remained excluded.
- No independent framing review or closure synchronization Task was created.
- Every Completion Condition passed on 2026-07-03.
- DRMCP is non-operational. Filesystem authoring was the required fallback.
