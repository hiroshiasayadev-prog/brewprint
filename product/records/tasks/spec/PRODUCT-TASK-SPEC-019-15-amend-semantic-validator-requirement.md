# PRODUCT-TASK-SPEC-019-15: Resolve semantic validator Requirement amendment

- **id**: PRODUCT-TASK-SPEC-019-15
- **status**: done
- **date**: 2026-07-01
- **work_item**: PRODUCT-WORK-SPEC-019
- **task_type**: authoring
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-019-14
- **outputs**:
  - PRODUCT-TASK-SPEC-019-15

## Goal

Resolve whether PRODUCT-REQ-SPEC-007 requires amendment after the accepted validator workflow decisions.

## Work

- Compare PRODUCT-REQ-SPEC-007 with T07 R-005, T07 R-006, and PRODUCT-ADR-SPEC-017.
- Distinguish the stable validation need from downstream workflow-policy decisions.
- Preserve PRODUCT-REQ-SPEC-007 unchanged when its motivating problem and Required Outcome remain accurate.
- Preserve PRODUCT-REQ-SPEC-005 unchanged.
- Record the no-amendment disposition and verification Evidence in this Task.

This Task must not:

- add downstream design decisions to PRODUCT-REQ-SPEC-007;
- change the Requirement identity or motivating problem;
- amend PRODUCT-REQ-SPEC-005;
- author ADR, Specification, checklist, or implementation content;
- modify `task_authoring`;
- perform review, correction, synchronization, stage, or commit work.

## Done condition

- PRODUCT-REQ-SPEC-007 is confirmed to remain accurate without amendment.
- Two-point invocation and human exception handling remain owned by PRODUCT-ADR-SPEC-017 and downstream Specifications.
- No downstream design decision is duplicated into the Requirement.
- PRODUCT-REQ-SPEC-005 is unchanged.
- PRODUCT-TASK-SPEC-019-16 is released.

## Verification

- Confirm PRODUCT-REQ-SPEC-007 already owns the stable need for semantic Task responsibility-boundary validation.
- Confirm its Required Outcome already supports use during Task authoring or before Task release.
- Confirm two mandatory invocation points and human exception handling are downstream design decisions.
- Confirm Requirement authoring policy prohibits recording design decisions in a Requirement.
- Confirm PRODUCT-REQ-SPEC-007 and PRODUCT-REQ-SPEC-005 are unchanged.
- Confirm no ADR, Specification, Work Item, checklist, implementation, review, correction, synchronization, stage, or commit change occurred.

## Evidence

### Result

- Result: `PASS`.
- PRODUCT-REQ-SPEC-007 requires semantic validation of one Task responsibility boundary.
- Its Required Outcome already supports validation during Task authoring or before Task release.
- The motivating problem and stable required outcome remain unchanged.

### No-amendment disposition

- T07 R-005 and R-006 selected two mandatory invocation points and human-owned violation judgment.
- T12 routed that durable workflow policy to PRODUCT-ADR-SPEC-017.
- PRODUCT-ADR-SPEC-017 now owns the accepted invocation and exception decision.
- Requirement authoring policy assigns design decisions to ADRs rather than Requirements.
- Adding the same workflow policy to PRODUCT-REQ-SPEC-007 would duplicate downstream design authority.
- PRODUCT-REQ-SPEC-007 therefore remains unchanged.

### Change boundary

- Only PRODUCT-TASK-SPEC-019-15 was changed.
- PRODUCT-REQ-SPEC-007 and PRODUCT-REQ-SPEC-005 were not changed.
- No ADR, Specification, Work Item, checklist, implementation, review, correction, synchronization, stage, or commit work occurred.
- PRODUCT-TASK-SPEC-019-16 is the exact next gate.
