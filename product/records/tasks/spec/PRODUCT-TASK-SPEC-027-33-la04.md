# PRODUCT-TASK-SPEC-027-33: Extract term observations for LA04

- **id**: PRODUCT-TASK-SPEC-027-33
- **status**: done
- **date**: 2026-07-04
- **work_item**: PRODUCT-WORK-SPEC-027
- **task_type**: investigation
- **estimate**: 1d
- **depends_on**:
  - PRODUCT-TASK-SPEC-027-04
- **outputs**:
  - product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/LA04.observations.jsonl
  - product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/LA04.coverage.json

## Goal

Extract governance-sensitive observed-use evidence from exactly one assigned batch into its owned observations JSONL and coverage JSON.

## Work

- Treat `product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/batches/LA04.json` as the exact source authority.
- Call `get_inventory_batch` for `LA04`.
- Follow PRODUCT-REQ-SPEC-013 and collection schemas `bp-wide-term-observation-v1` and `bp-wide-term-batch-v1`.
- Use `record_inventory_file_result` to record each source as `observations`, `no_candidate`, or `failed`.
- Treat `changed` as a non-blocking diagnostic.
- Run `validate_inventory_batch` after every source has one disposition.
- Read no source outside LA04 and write no other batch result.
- Do not normalize synonyms, merge meanings, judge correctness, define terms, aggregate, cluster, extract use cases, modify PRODUCT-INV-SPEC-011, stage, or commit.

## Done condition

Every LA04 source has one disposition, both owned outputs conform, and validation has no structural diagnostics.

## Verification

Confirm batch `LA04`, prefix `PRODUCT-TASK-SPEC-027-33-O`, exact source count, schemas, output isolation, and no stage or commit.

## Evidence

- PRODUCT-REQ-SPEC-013 owns the gathering criterion and schemas.
- PRODUCT-TASK-SPEC-027-04 owns batch assignment and writer isolation.
- T37 concludes PRODUCT-INV-SPEC-011 from all batch Evidence.
- DRMCP is non-operational. Filesystem fallback was used for Design Record reads and this Task update.
- `get_inventory_batch` confirmed `LA04`, 20 sources, prefix `PRODUCT-TASK-SPEC-027-33-O`, and the two owned result paths.
- All 20 sources received `observations`; no source received `no_candidate` or `failed`.
- The JSONL contains 79 `bp-wide-term-observation-v1` records with IDs `PRODUCT-TASK-SPEC-027-33-O0001` through `PRODUCT-TASK-SPEC-027-33-O0079`.
- Coverage uses `bp-wide-term-batch-v1`: 20 discovered, 20 scanned, 0 failed, and 79 observations.
- `validate_inventory_batch` returned `valid: true`, 20 disposed sources, 79 observations, and no diagnostics.
- All 20 source snapshots were `unchanged`.
- Only the LA04 observations, coverage, and this Task record were written.
- No synonym normalization, meaning merge, correctness judgment, term definition, aggregation, clustering, use-case extraction, Investigation modification, staging, or commit occurred.
