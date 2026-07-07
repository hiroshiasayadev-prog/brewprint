# PRODUCT-TASK-SPEC-027-16: Extract term observations for PN04

- **id**: PRODUCT-TASK-SPEC-027-16
- **status**: done
- **date**: 2026-07-04
- **work_item**: PRODUCT-WORK-SPEC-027
- **task_type**: investigation
- **estimate**: 1d
- **depends_on**:
  - PRODUCT-TASK-SPEC-027-04
- **outputs**:
  - product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/PN04.observations.jsonl
  - product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/PN04.coverage.json

## Goal

Extract governance-sensitive observed-use evidence from exactly one assigned batch into its owned observations JSONL and coverage JSON.

## Work

- Use `product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/batches/PN04.json` as exact source authority and call `get_inventory_batch` for `PN04`.
- Read no batch-external source.
- Apply PRODUCT-REQ-SPEC-013 and schemas `bp-wide-term-observation-v1` and `bp-wide-term-batch-v1`.
- Record every source through `record_inventory_file_result` as `observations`, `no_candidate`, or `failed`.
- Treat `changed` as non-blocking and finish with `validate_inventory_batch`.
- Change only the two owned result files and this Task record.
- Do not normalize, merge meanings, judge correctness, define terms, aggregate, cluster, extract use cases, modify PRODUCT-INV-SPEC-011, stage, or commit.

## Done condition

Every `PN04` source has one disposition, both outputs conform, validation has no structural diagnostics, and no excluded semantic work occurred.

## Verification

Confirm batch `PN04`, prefix `PRODUCT-TASK-SPEC-027-16-O`, exact source count, schemas, output isolation, and no stage or commit.

## Evidence

- PRODUCT-REQ-SPEC-013 owns gathering and schemas.
- T04 owns assignment and writer isolation.
- The user accepted per-batch lightweight Investigation Evidence because T37 concludes PRODUCT-INV-SPEC-011.
- `get_inventory_batch` confirmed batch `PN04`, 30 assigned sources, and observation prefix `PRODUCT-TASK-SPEC-027-16-O`.
- Every assigned source received exactly one disposition: 30 `observations`, 0 `no_candidate`, and 0 `failed`.
- `PN04.observations.jsonl` contains 351 records conforming to `bp-wide-term-observation-v1`.
- `PN04.coverage.json` conforms to `bp-wide-term-batch-v1` and records 30 discovered, 30 scanned, and 0 failed sources.
- `validate_inventory_batch` returned `valid: true` with 351 observations, no diagnostics, and all 30 source snapshots `unchanged`.
- Only the two owned result files and this Task record were changed. No normalization, definition, aggregation, clustering, use-case extraction, staging, or commit occurred.
