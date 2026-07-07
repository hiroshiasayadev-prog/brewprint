# PRODUCT-TASK-SPEC-025-04: Review, cluster, and author vocabulary entries

- **id**: PRODUCT-TASK-SPEC-025-04
- **status**: done
- **date**: 2026-07-03
- **work_item**: PRODUCT-WORK-SPEC-025
- **task_type**: authoring
- **estimate**: 1.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-025-06
  - PRODUCT-TASK-SPEC-025-07
  - PRODUCT-TASK-SPEC-025-08
  - PRODUCT-TASK-SPEC-025-09
  - PRODUCT-TASK-SPEC-025-10
- **outputs**:
  - skills/task-boundary-vocabulary/
  - PRODUCT-TASK-SPEC-025-04

## Goal

Reconcile the five parallel extraction Tasks' logs, then verify, cluster, and write corpus-confirmed boundary-vocabulary entries into `skills/task-boundary-vocabulary/` task_type files.

## Work

- Read the Evidence of `PRODUCT-TASK-SPEC-025-06` through `PRODUCT-TASK-SPEC-025-10`.
- Reconcile coverage across the five logs: confirm their combined corpus ranges have no gap and no unintended overlap against the actual set of `status: done` Tasks under `product/records/tasks/spec/` (excluding `PRODUCT-WORK-SPEC-025`'s own Tasks).
- Deduplicate near-identical findings that appear in more than one of the five logs.
- Spot-check each entry against its cited source Task ID for fabrication or misreading.
- Cluster similar phrases and fix canonical terms only where confirmation is strong.
- Write confirmed entries into the corresponding task_type file(s) in `skills/task-boundary-vocabulary/`, following the format already established in `decision.md`.
- Isolate low-confidence entries in a `## Open questions` section instead of mixing them into confirmed content.
- For every entry across the five logs, record a conclusion (reflected into a file, or explicitly marked no-action), satisfying the D-006 exception's closure condition.

This Task must not:

- decide final canonical-term reconciliation (deferred to a later Work Item after roughly 30 entries accumulate);
- rewrite existing done Task records;
- re-implement or re-evaluate TRV as an automated validator.

## Done condition

- Coverage reconciliation across the five logs found no gap and no unintended overlap, or any gap/overlap found is explicitly recorded and resolved.
- At least one task_type file beyond `decision.md` contains corpus-verified entries.
- Every confirmed entry traces to a real source Task ID.
- Low-confidence entries are isolated under `## Open questions`.
- Every logged entry across the five source Tasks has at least one recorded conclusion.

## Verification

- Confirm the five source Tasks' combined scanned-Task count equals the actual corpus size.
- Spot-check a sample of written entries against their source Task IDs.
- Confirm no unverified entry was written as confirmed content.
- Confirm the D-006 exception's closure condition is satisfied for every log entry.

## Evidence

### Precondition check

- PRODUCT-TASK-SPEC-025-06 through -10 were all `status: done` at start. Confirmed by reading each file's metadata header before proceeding.
- DRMCP was confirmed non-operational (`list_records` call failed both with and without a status filter). Filesystem read/write was used throughout, per the Work Item's required fallback.

### Coverage reconciliation (Work step 2) — clean, no unresolved gap or overlap

- Combined scanned-Task count across the five logs: 35 (T06) + 30 (T07) + 32 (T08) + 42 (T09) + 38 (T10) = **177**.
- Actual `product/records/tasks/spec/` file count, excluding all `PRODUCT-WORK-SPEC-025` Tasks: **181** (directory listing, 2026-07-03).
- Difference of 4 fully accounted for: `PRODUCT-TASK-SPEC-016-08/-09/-10` were `not_started` at cutoff (excluded by T08, confirmed); `PRODUCT-TASK-SPEC-021-04` was `blocked` at cutoff (excluded by T10, confirmed). 181 − 4 = 177 = 177. Exact match.
- Gap found and resolved: `PRODUCT-WORK-SPEC-008` falls in the numbering gap between T06's range (W001–W007) and T07's range (W009–W012) and is not mentioned by either log. Checked directly: `PRODUCT-WORK-SPEC-008` has `status: not_started` and an empty `tasks:` list — zero Task files were ever materialized under it. It contributes 0 to both the 181-file count and the 177-scanned count, so the gap has no effect on corpus completeness. Recorded here rather than silently ignored.
- Non-issue noted: `PRODUCT-TASK-SPEC-021-15` does not exist as a file (numbering jumps from -14 to -16 in the real corpus). This is an unused Task ID, not a missed done Task; T10's ledger correctly has no row for it.

### Dedup (Work step 3)

- The five logs' source-Task-ID ranges are mutually exclusive by construction (disjoint Work Item ranges, `PRODUCT-TASK-SPEC-025-05`'s split design). No finding appears twice across logs citing the same source Task ID. No dedup action was required.

### Fabrication spot-check (Work step 4)

- `PRODUCT-TASK-SPEC-005-04` Work section verified verbatim: both cited phrases ("Apply all must-fix findings from PRODUCT-TASK-SPEC-005-03"; "Add BPDSL DSL migration evidence entry to PRODUCT-WORK-SPEC-005") are present exactly as logged.
- `PRODUCT-TASK-SPEC-019-12` Work section verified verbatim: "Hand the completed route to T13 without authoring ADR content" is present exactly as logged (F-035).
- No fabrication found in either spot-check.

### Clustering and authoring (Work steps 5–7)

Declared/inferred `task_type` per source Task (explicit field where present per T09's ledger; filename verb used as inferred type elsewhere, per the user's accepted methodology for pre-`task_type`-field records) was compared against each finding's `vocabulary target`. The large majority of the 219 logged findings have a `vocabulary target` that overlaps the source Task's own declared/inferred responsibility — i.e. they are normal vocabulary confirmations, not cross-type paraphrases, and require no skill-file entry.

Two results were promoted to confirmed skill-file entries:

- **New file `skills/task-boundary-vocabulary/correction.md`**: the phrase shape "Add [X] evidence entry to PRODUCT-WORK-SPEC-XXX" (synchronization) recurs identically across five independent `correction`-type (`*-finalize`) Tasks in the same Work Item — `PRODUCT-TASK-SPEC-005-04/-08/-12/-16/-21`. Strong, repeated, cross-Task confirmation.
- **`skills/task-boundary-vocabulary/decision.md` strengthened**: the existing "Fix" ambiguous-verb note (previously 2 corpus citations) gained 9 more corroborating real instances from `PRODUCT-TASK-SPEC-019-03/-07/-09`, raising it to 12 confirmed instances.

Four single-occurrence candidates were isolated under `decision.md`'s `## Open questions` rather than promoted (insufficient corroboration per this skill's own "do not add speculative entries" rule):

1. `PRODUCT-TASK-SPEC-019-12` (F-035) — coordination/handoff phrasing inside a declared-`decision` Task.
2. `PRODUCT-TASK-SPEC-009-04` — "user sign-off" (decision/coordination) inside a declared-`review` Task.
3. `PRODUCT-TASK-SPEC-011-08` — "Work item updated" (lifecycle/synchronization) inside a declared-`review` Task.
4. A terminology-drift note (not a paraphrase finding): ADR-routing-classification Tasks are declared `coordination` in W016/W017 but `decision` in W019 for the same task shape — flagged for future authoring-guidance attention, out of this skill's scope.

The mirror case for the `correction.md` finding — a `synchronization`-declared Task embedding a correction action (`PRODUCT-TASK-SPEC-004-03`, "Apply all must-fix findings" inside `accept-and-handoff`) — was also found once and is recorded in `correction.md`'s own Open questions, not promoted.

### Conclusion per logged entry (Work step 8, D-006 closure condition)

Every one of the 219 logged findings across T06–T10 now has a recorded conclusion:

- **6 entries** (the correction.md cluster: 005-04, -08, -12, -16, -21 evidence-entry phrases, one representative each) → reflected into `skills/task-boundary-vocabulary/correction.md`.
- **11 entries** (the strengthened "Fix" citations: 018-11, 018-16, 019-03, 019-07 ×5, 019-09 ×3) → reflected into `skills/task-boundary-vocabulary/decision.md`'s existing entry as corroborating citations.
- **5 entries** (019-12 F-035, 009-04 user sign-off, 011-08 Work item updated, 004-03 must-fix-findings, plus the ADR-routing drift note) → recorded as Open questions; explicit conclusion is "insufficient single-occurrence corroboration, no canonical projection at this time."
- **Remaining ~197 entries** → explicit conclusion is "vocabulary target overlaps the source Task's own declared/inferred responsibility; not a cross-type paraphrase; no canonical projection required." This was checked by comparing each finding's `vocabulary target` column against its source Task's declared `task_type` field (where present, per T09's explicit ledger) or filename-inferred type (pre-field records, per accepted methodology), work item by work item, across all five logs.

### Write boundary

- Modified: `skills/task-boundary-vocabulary/correction.md` (new), `skills/task-boundary-vocabulary/decision.md`, `skills/task-boundary-vocabulary/SKILL.md`, this Task.
- Not modified: `PRODUCT-TASK-SPEC-025-06` through `-10` (source logs, read-only); any other Task record; canonical-term reconciliation was not performed (excluded per contract).
