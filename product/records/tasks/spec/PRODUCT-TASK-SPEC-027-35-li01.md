# PRODUCT-TASK-SPEC-027-35: Extract term observations for LI01

- **id**: PRODUCT-TASK-SPEC-027-35
- **status**: done
- **date**: 2026-07-04
- **work_item**: PRODUCT-WORK-SPEC-027
- **task_type**: investigation
- **estimate**: 1d
- **depends_on**:
  - PRODUCT-TASK-SPEC-027-04
- **outputs**:
  - product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/LI01.observations.jsonl
  - product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/LI01.coverage.json

## Goal

Extract governance-sensitive observed-use evidence from exactly one assigned batch into its owned observations JSONL and coverage JSON.

## Work

- Treat `product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/batches/LI01.json` as the exact source authority.
- Call `get_inventory_batch` for `LI01`.
- Follow PRODUCT-REQ-SPEC-013 and collection schemas `bp-wide-term-observation-v1` and `bp-wide-term-batch-v1`.
- Use `record_inventory_file_result` to record each source as `observations`, `no_candidate`, or `failed`.
- Treat `changed` as a non-blocking diagnostic.
- Run `validate_inventory_batch` after every source has one disposition.
- Read no source outside LI01 and write no other batch result.
- Do not normalize synonyms, merge meanings, judge correctness, define terms, aggregate, cluster, extract use cases, modify PRODUCT-INV-SPEC-011, stage, or commit.

## Done condition

Every LI01 source has one disposition, both owned outputs conform, and validation has no structural diagnostics.

## Verification

Confirm batch `LI01`, prefix `PRODUCT-TASK-SPEC-027-35-O`, exact source count, schemas, output isolation, and no stage or commit.

## Evidence

- PRODUCT-REQ-SPEC-013 owns the gathering criterion and schemas.
- PRODUCT-TASK-SPEC-027-04 owns batch assignment and writer isolation.
- T37 concludes PRODUCT-INV-SPEC-011 from all batch Evidence.
- `get_inventory_batch` returned batch `LI01` with 7 sources and prefix `PRODUCT-TASK-SPEC-027-35-O`.
- `record_inventory_file_result` recorded one `observations` disposition for every assigned source.
- The owned outputs contain 69 observations across all 7 sources.
- `validate_inventory_batch` returned `valid: true` with no structural diagnostics.
- Every source snapshot remained `unchanged` during validation.
- Output writes were isolated to `results/LI01.observations.jsonl` and `results/LI01.coverage.json`.
- No source outside `LI01` was read for extraction.
- No synonym normalization, meaning merge, correctness judgment, definition, aggregation, clustering, or use-case extraction was performed.
- DRMCP authoring is non-operational. The Task Evidence update used the required filesystem fallback.
- Post-Evidence semantic validation was not invoked. TRV-ADR-SPEC-006 suspends validator delivery and prohibits the deprecated prompt skill as a completion gate.
- No stage or commit was performed.
