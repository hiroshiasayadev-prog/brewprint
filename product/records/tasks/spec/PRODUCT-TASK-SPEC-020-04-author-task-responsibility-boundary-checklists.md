# PRODUCT-TASK-SPEC-020-04: Author Task responsibility-boundary checklists

- **id**: PRODUCT-TASK-SPEC-020-04
- **status**: done
- **date**: 2026-07-02
- **work_item**: PRODUCT-WORK-SPEC-020
- **task_type**: authoring
- **estimate**: 2d
- **depends_on**:
  - PRODUCT-TASK-SPEC-020-03
  - PRODUCT-TASK-SPEC-022-03
- **outputs**:
  - PRODUCT-TASK-SPEC-020-04
  - skills/task-responsibility-boundary-validator/SKILL.md
  - skills/task-responsibility-boundary-validator/prompts/evaluator-instructions.md
  - skills/task-responsibility-boundary-validator/prompts/common.md
  - skills/task-responsibility-boundary-validator/prompts/task-types/investigation.md
  - skills/task-responsibility-boundary-validator/prompts/task-types/decision.md
  - skills/task-responsibility-boundary-validator/prompts/task-types/authoring.md
  - skills/task-responsibility-boundary-validator/prompts/task-types/implementation.md
  - skills/task-responsibility-boundary-validator/prompts/task-types/review.md
  - skills/task-responsibility-boundary-validator/prompts/task-types/correction.md
  - skills/task-responsibility-boundary-validator/prompts/task-types/verification.md
  - skills/task-responsibility-boundary-validator/prompts/task-types/coordination.md
  - skills/task-responsibility-boundary-validator/prompts/task-types/work_item_decomposition.md
  - skills/task-responsibility-boundary-validator/prompts/task-types/work_item_execution.md
  - skills/task-responsibility-boundary-validator/prompts/task-types/synchronization.md

## Goal

Author one complete evaluator prompt asset set for every canonical Task responsibility type.

## Work

- Preserve the accepted Markdown placement and common-plus-type composition.
- Preserve every approved common and existing type checklist baseline.
- Add the canonical `work_item_decomposition`, `work_item_execution`, and `synchronization` projections.
- Add a minimal skill wrapper with canonical-authority precedence and deterministic prompt composition.
- Keep every criterion as one affirmative binary semantic claim.
- Keep every criterion judgeable from one Task record.
- Keep structural validation, runtime policy, external response fields, and implementation outside the asset set.
- Record criterion counts, compactness results, authority mapping, and aggregate prompt overhead.

The user authorized completing the remaining artifact set on 2026-07-02 without the earlier item-by-item approval pause.

This Task must not:

- introduce or reinterpret canonical Task rules;
- implement the validator;
- modify current DRMCP artifacts;
- issue an independent review verdict;
- perform finding correction, lifecycle synchronization, stage, or commit work.

## Done condition

- The common checklist exists in the accepted format and placement.
- All eleven canonical Task types have one exact checklist projection.
- Every criterion expresses one binary responsibility-boundary judgment.
- Every criterion is grounded in accepted Task-authoring authority.
- The minimal skill selects evaluator instructions, common criteria, and the declared-type criteria deterministically.
- The complete asset set is directly consumable without content inference.
- Aggregate prompt overhead is recorded.
- No implementation or external validator response contract is introduced.

## Verification

- Confirm all eleven canonical Task types are present.
- Confirm common criteria are not repeated into type files without a type-specific reason.
- Confirm every criterion has one semantic claim and Task-local evidence boundary.
- Confirm type files preserve the canonical primary outcome, completion judgment, and prohibited overlaps.
- Confirm evaluator instructions require a reason and Task section for every result.
- Confirm the skill points to canonical Task authoring authority and defines conflict precedence.
- Confirm no implementation source changed.
- Confirm no independent review, correction, synchronization, stage, or commit occurred.

## Evidence

### Result

`PASS`.

The complete derived evaluator asset set exists for all eleven canonical Task types.
The earlier `work_item_decomposition` blocker is resolved by the accepted `work_item_execution` addition without changing decomposition semantics.
PRODUCT-TASK-SPEC-022-03 released this existing authoring owner after canonical authority was available.

### Artifact status and compactness

| item | criteria | lines | words | characters | approximate tokens |
|---|---:|---:|---:|---:|---:|
| evaluator instructions | n/a | 5 | 62 | 377 | 94 |
| common | 14 | 14 | 169 | 1069 | 267 |
| investigation | 9 | 9 | 84 | 551 | 138 |
| decision | 9 | 9 | 88 | 554 | 139 |
| authoring | 9 | 9 | 82 | 526 | 132 |
| implementation | 8 | 8 | 72 | 479 | 120 |
| review | 9 | 9 | 86 | 544 | 136 |
| correction | 7 | 7 | 66 | 405 | 101 |
| verification | 9 | 9 | 85 | 534 | 134 |
| coordination | 7 | 7 | 68 | 433 | 108 |
| work_item_decomposition | 7 | 7 | 104 | 702 | 176 |
| work_item_execution | 8 | 8 | 111 | 774 | 194 |
| synchronization | 8 | 8 | 97 | 647 | 162 |

Character counts exclude each file's terminal newline.
Token counts are character-based estimates.
The largest runtime composition is evaluator instructions plus common plus `work_item_execution`:

- 27 lines;
- 342 words;
- 2222 characters;
- approximately 556 tokens.

### New type authority mapping

| checklist | projected authority |
|---|---|
| `work_item_decomposition` | The Task record owns only child Work Item boundary creation or split, coarse parent routing, and decomposition completion; it does not absorb child execution or Task-graph coordination. |
| `work_item_execution` | The Task record references one existing child through `work_item_ref`, describes only its coarse parent-graph role, and uses child completion as its sole completion boundary without duplicating child-owned work. |
| `synchronization` | One mechanically derived accepted-state propagation; exact targets and values; no judgment, graph change, substantive authoring, repair, implementation, or review. |

### Completion checks

- Existing approved evaluator, common, investigation, decision, authoring, implementation, review, correction, verification, and coordination baselines remain present.
- The evaluator instruction now includes the Task section required by `spec:product.responsibility_boundary_validator`.
- `SKILL.md` identifies the assets as derived prompts and makes `task_authoring` authoritative on conflict.
- The type directory contains one file for each of the eleven canonical `task_type` values.
- No checklist requires external repository inference.
- The user explicitly approved Task-local responsibility refinements for `work_item_decomposition` and `work_item_execution` after rejecting criteria that depended on inspecting child or parent Work Item content.
- Direct read-back confirmed the approved decomposition and execution checklist text.
- The earlier PRODUCT-TASK-SPEC-022-04 review predates these two refinements and is historical Evidence only; the current complete checklist set remains subject to PRODUCT-TASK-SPEC-020-05 integrated review.
- No validator implementation, DRMCP integration, independent review, correction, synchronization, stage, or commit work occurred.
- DRMCP is non-operational, so filesystem authoring was used.
