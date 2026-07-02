# PRODUCT-TASK-SPEC-020-03: Investigate checklist coverage and placement

- **id**: PRODUCT-TASK-SPEC-020-03
- **status**: done
- **date**: 2026-07-01
- **work_item**: PRODUCT-WORK-SPEC-020
- **task_type**: investigation
- **estimate**: 1d
- **depends_on**:
  - PRODUCT-TASK-SPEC-020-02
- **outputs**:
  - PRODUCT-TASK-SPEC-020-03
  - PRODUCT-INV-SPEC-008

## Goal

Produce one Investigation that verifies checklist coverage, authority alignment, placement effects, and authoring risks.

## Work

- Create PRODUCT-INV-SPEC-008 for one bounded checklist coverage question.
- Inspect every canonical Task type in `task_authoring`.
- Identify common criteria that apply to every Task type.
- Identify type-specific primary outcomes, completion judgments, and prohibited overlaps.
- Check the decided format and placement against repository authority.
- Identify missing coverage, semantic conflicts, shared writers, and graph-change candidates.
- Record uncertainty without selecting new design choices.

This Task must not:

- adopt new checklist contract decisions;
- author checklist artifacts;
- modify ADRs or Specifications;
- implement the validator;
- perform review, correction, synchronization, stage, or commit work.

## Done condition

- PRODUCT-INV-SPEC-008 covers all canonical Task types.
- Every common and type-specific source rule has an exact authority reference.
- Every placement or format conflict has evidence.
- Every missing judgment or graph candidate has a named next owner.
- No checklist wording is silently adopted.

## Verification

- Confirm all ten canonical Task types are covered.
- Confirm common and type-specific coverage remain distinct.
- Confirm the Investigation does not author checklist content.
- Confirm no repository-wide traversal was used when exact authority was known.
- Confirm every uncertainty has a named owner.

## Evidence

- Created `PRODUCT-INV-SPEC-008` with `status: concluded`.
- The Investigation answers one bounded checklist coverage and placement question.
- Common coverage includes one outcome, one completion judgment, declared-type alignment, section-role alignment, supporting-action boundaries, and all canonical split conditions.
- Type-specific coverage includes all ten canonical `task_type` values.
- Every coverage topic maps to an exact `task_authoring` section or table row.
- Structural format, metadata, lifecycle, prose, runtime, and external response-schema rules are excluded from semantic checklist projection.
- Task-local evaluation does not infer actual external artifact, graph, command, or reviewer state.
- The accepted `skills/task-responsibility-boundary-validator/` placement does not conflict with Design Record placement authority.
- Common and type-specific partitioning requires no additional writer or graph change.
- T04 can own the minimal skill, evaluator instructions, common checklist, and ten type-specific checklists as one bounded artifact set.
- The evaluator-internal structured JSON instruction remains separate from any external validator response contract.
- No unresolved candidate blocks T04. Conditional stop routes have named decision, coordination, implementation-contract, or finding-correction owners.
- Verification passed: all ten types are explicit; common and type coverage are separate; every adopted topic has authority; excluded rules are classified; no final criterion wording or IDs were authored; T02 was not reopened; no graph or writer candidate was applied; only the two permitted files changed; no stage or commit occurred.
