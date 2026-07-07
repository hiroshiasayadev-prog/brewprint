# PRODUCT-TASK-SPEC-027-36: Extract term observations for LI02

- **id**: PRODUCT-TASK-SPEC-027-36
- **status**: done
- **date**: 2026-07-04
- **work_item**: PRODUCT-WORK-SPEC-027
- **task_type**: investigation
- **estimate**: 1d
- **depends_on**:
  - PRODUCT-TASK-SPEC-027-04
- **outputs**:
  - product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/LI02.observations.jsonl
  - product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/LI02.coverage.json

## Goal

Extract governance-sensitive observed-use evidence from exactly one assigned batch into its owned observations JSONL and coverage JSON.

## Work

- Treat `product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/batches/LI02.json` as the exact source authority.
- Call `get_inventory_batch` for `LI02`.
- Follow PRODUCT-REQ-SPEC-013 and collection schemas `bp-wide-term-observation-v1` and `bp-wide-term-batch-v1`.
- Use `record_inventory_file_result` to record each source as `observations`, `no_candidate`, or `failed`.
- Treat `changed` as a non-blocking diagnostic.
- Run `validate_inventory_batch` after every source has one disposition.
- Read no source outside LI02 and write no other batch result.
- Do not normalize synonyms, merge meanings, judge correctness, define terms, aggregate, cluster, extract use cases, modify PRODUCT-INV-SPEC-011, stage, or commit.

## Done condition

Every LI02 source has one disposition, both owned outputs conform, and validation has no structural diagnostics.

## Verification

Confirm batch `LI02`, prefix `PRODUCT-TASK-SPEC-027-36-O`, exact source count, schemas, output isolation, and no stage or commit.

## Evidence

- PRODUCT-REQ-SPEC-013 owns the gathering criterion and schemas.
- PRODUCT-TASK-SPEC-027-04 owns batch assignment and writer isolation.
- T37 concludes PRODUCT-INV-SPEC-011 from all batch Evidence.
- `get_inventory_batch` returned batch `LI02` with 6 sources and prefix `PRODUCT-TASK-SPEC-027-36-O`.
- `record_inventory_file_result` recorded one `observations` disposition for every assigned source.
- The owned outputs contain 121 observations across all 6 sources.
- `LI02.coverage.json` uses schema `bp-wide-term-batch-v1`.
- Every JSONL row uses schema version `bp-wide-term-observation-v1`.
- `validate_inventory_batch` returned `valid: true` with no structural diagnostics.
- Every source snapshot remained `unchanged` during validation.
- Output writes were isolated to `results/LI02.observations.jsonl` and `results/LI02.coverage.json`.
- No corpus source outside `LI02` was read for extraction.
- No synonym normalization, meaning merge, correctness judgment, definition, aggregation, clustering, or use-case extraction was performed.
- The deprecated responsibility-boundary validator skill was not invoked as a workflow gate, as required by its `TRV-ADR-SPEC-006` deprecation notice.
- DRMCP authoring is non-operational. The Task Evidence update used the required filesystem fallback.
- No stage or commit was performed.
