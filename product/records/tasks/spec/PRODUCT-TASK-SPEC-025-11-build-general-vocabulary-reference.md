# PRODUCT-TASK-SPEC-025-11: Build general per-task_type vocabulary reference

- **id**: PRODUCT-TASK-SPEC-025-11
- **status**: done
- **date**: 2026-07-03
- **work_item**: PRODUCT-WORK-SPEC-025
- **task_type**: authoring
- **estimate**: 1d
- **depends_on**:
  - PRODUCT-TASK-SPEC-025-06
  - PRODUCT-TASK-SPEC-025-07
  - PRODUCT-TASK-SPEC-025-08
  - PRODUCT-TASK-SPEC-025-09
  - PRODUCT-TASK-SPEC-025-10
- **outputs**:
  - skills/task-vocabulary-reference/
  - PRODUCT-TASK-SPEC-025-11

## Goal

Reorganize the full raw finding logs of `PRODUCT-TASK-SPEC-025-06` through `-10` (219 entries, not just the boundary-violation subset `PRODUCT-TASK-SPEC-025-04` promoted) into a general, per-task_type vocabulary reference: what real phrasing each task_type actually uses in this corpus, independent of whether it crosses a type boundary.

This is a separate artifact from `skills/task-boundary-vocabulary/`, which stays scoped to cross-type disguise patterns only.

## Work

- Read the full Finding log tables of `PRODUCT-TASK-SPEC-025-06` through `-10` (no new source-Task reading required; this reorganizes already-collected data).
- Group entries by the source Task's declared or inferred `task_type` cluster (decision / authoring / coordination; review / verification / correction; investigation / decision / design; graph / lifecycle / execution), not by the "vocabulary target" column T04 used.
- For each cluster, list the distinct real phrases used, each with at least one citing source Task ID, without filtering out same-type usage.
- Do not exclude an entry because it overlaps its own source Task's declared responsibility — that overlap is exactly what this reference documents.
- Keep `PRODUCT-TASK-SPEC-025-04`'s cross-type findings (`correction.md`, the "Fix" entries) out of scope here; this reference is about normal in-type vocabulary breadth, not disguise detection.
- Create `skills/task-vocabulary-reference/` with one file per task_type cluster (or a single index file if the volume in a cluster is small), following plain declarative style consistent with `skills/task-boundary-vocabulary/decision.md`.

This Task must not:

- re-read the 177 source Tasks directly (use the existing T06–T10 logs as the source);
- decide canonical term definitions or flag boundary violations (that remains `skills/task-boundary-vocabulary/`'s scope);
- modify `skills/task-boundary-vocabulary/` files.

## Done condition

- `skills/task-vocabulary-reference/` exists with per-cluster content covering all five source logs.
- Every listed phrase cites at least one real source Task ID from T06–T10's logs.
- No entry from `skills/task-boundary-vocabulary/` was altered.

## Verification

- Confirm every phrase traces to an actual row in T06–T10's Finding log tables (no invented phrases).
- Confirm clusters cover all task_type groupings used in T06–T10.
- Confirm `skills/task-boundary-vocabulary/` was not modified by this Task.

## Evidence

- Read the full Finding log tables of T06 (56 findings), T07 (35 findings), T08 (49 findings), T09 (49 findings), and T10 (30 findings) — 219 entries total, matching T04's reconciliation count. No source Task under `product/records/tasks/spec/` was re-read; only the five logs already in the repository were used.
- Grouping deviated slightly from this Task's own Work instruction: T09 and T10 carry an explicit `declared task type` column per source Task, and T07 carries an `inferred responsibility cluster` column, but T06 and T08 do not carry any per-entry source-type field. Re-deriving T06/T08's source Task types would require re-reading the 177 source Tasks, which this Task must not do. Entries were grouped by the `vocabulary target` field instead, which in the large majority of rows across all five logs already reads as one of the four cluster labels those Tasks used verbatim (`decision / authoring / coordination`, `review / verification / correction`, `investigation / decision / design`, `graph / lifecycle / execution`). This is a pragmatic substitution, not the literal instruction as written; flagging it here rather than silently deviating.
- Created `skills/task-vocabulary-reference/` with `SKILL.md` and four cluster files (`decision-authoring-coordination.md`, `review-verification-correction.md`, `investigation-decision-design.md`, `graph-lifecycle-execution.md`).
- Identical or near-identical phrasing repeated across multiple source Tasks (e.g. `Apply all must-fix findings` ×6, `Guide published` ×4, `Fix` ×12, `Persist each explicit answer before advancing` ×3) was deduplicated to one entry with representative citations, per this Task's own Done condition ("every listed phrase cites at least one real source Task ID" — not every occurrence). This trades exact frequency counts for a readable reference; the source logs remain the frequency-accurate record if needed later.
- `skills/task-boundary-vocabulary/` was not modified by this Task.
- DRMCP is non-operational. Filesystem authoring was the required fallback.
