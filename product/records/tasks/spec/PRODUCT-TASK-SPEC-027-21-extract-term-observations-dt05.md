# PRODUCT-TASK-SPEC-027-21: Extract term observations for DT05

- **id**: PRODUCT-TASK-SPEC-027-21
- **status**: done
- **date**: 2026-07-04
- **work_item**: PRODUCT-WORK-SPEC-027
- **task_type**: investigation
- **estimate**: 1d
- **depends_on**:
  - PRODUCT-TASK-SPEC-027-04
- **outputs**:
  - product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/DT05.observations.jsonl
  - product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/DT05.coverage.json

## Goal

Extract governance-sensitive observed-use evidence from exactly one assigned batch into its owned observations JSONL and coverage JSON.

## Work

- Use `product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/batches/DT05.json` as exact source authority and call `get_inventory_batch` for `DT05`.
- Read no batch-external source.
- Apply PRODUCT-REQ-SPEC-013 and schemas `bp-wide-term-observation-v1` and `bp-wide-term-batch-v1`.
- Record every source through `record_inventory_file_result` as `observations`, `no_candidate`, or `failed`.
- Treat `changed` as non-blocking and finish with `validate_inventory_batch`.
- Change only the two owned result files and this Task record.
- Do not normalize synonyms, merge meanings, judge correctness, define terms, aggregate, cluster, extract use cases, modify PRODUCT-INV-SPEC-011, stage, or commit.

## Done condition

Every `DT05` source has one disposition, both outputs conform, validation has no structural diagnostics, and no excluded semantic work occurred.

## Verification

Confirm batch `DT05`, prefix `PRODUCT-TASK-SPEC-027-21-O`, exact source count, schemas, output isolation, and no stage or commit.

## Evidence

- PRODUCT-REQ-SPEC-013 owns gathering and schemas.
- PRODUCT-TASK-SPEC-027-04 owns assignment and writer isolation.
- T37 concludes the shared PRODUCT-INV-SPEC-011 from all batch Evidence.
- Batch authority was `product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/batches/DT05.json`: batch `DT05`, observation prefix `PRODUCT-TASK-SPEC-027-21-O`, and 16 assigned sources.
- `get_inventory_batch` was invoked twice for `DT05`, but both calls were blocked by the platform safety check. The exact batch authority file was therefore read through filesystem fallback. No batch-external corpus source was read.
- Every assigned source was disposed exactly once through `record_inventory_file_result`; all 16 dispositions are `observations`.
- `DT05.observations.jsonl` contains 159 `bp-wide-term-observation-v1` observations with IDs `PRODUCT-TASK-SPEC-027-21-O0001` through `PRODUCT-TASK-SPEC-027-21-O0159`.
- `DT05.coverage.json` uses `bp-wide-term-batch-v1` and records 16 discovered, 16 scanned, 0 failed, and 159 observations.
- `validate_inventory_batch` returned `valid: true`, 16 of 16 disposed sources, no structural diagnostics, and `unchanged` current and recorded snapshot status for every source.
- No synonym normalization, meaning merge, correctness judgment, term definition, aggregation, clustering, use-case extraction, or shared Investigation modification occurred.
- The changed-file boundary is limited to this Task and its two owned result files. No file was staged or committed.
