# PRODUCT-TASK-SPEC-027-15: Extract term observations for PN03

- **id**: PRODUCT-TASK-SPEC-027-15
- **status**: done
- **date**: 2026-07-04
- **work_item**: PRODUCT-WORK-SPEC-027
- **task_type**: investigation
- **estimate**: 1d
- **depends_on**:
  - PRODUCT-TASK-SPEC-027-04
- **outputs**:
  - product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/PN03.observations.jsonl
  - product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/PN03.coverage.json

## Goal

Extract governance-sensitive observed-use evidence from exactly one assigned batch into its owned observations JSONL and coverage JSON.

## Work

- Use `product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/batches/PN03.json` as exact source authority and call `get_inventory_batch` for `PN03`.
- Read no batch-external source.
- Apply PRODUCT-REQ-SPEC-013 and schemas `bp-wide-term-observation-v1` and `bp-wide-term-batch-v1`.
- Record every source through `record_inventory_file_result` as `observations`, `no_candidate`, or `failed`.
- Treat `changed` as non-blocking and finish with `validate_inventory_batch`.
- Change only the two owned result files and this Task record.
- Do not normalize, merge meanings, judge correctness, define terms, aggregate, cluster, extract use cases, modify PRODUCT-INV-SPEC-011, stage, or commit.

## Done condition

Every `PN03` source has one disposition, both outputs conform, validation has no structural diagnostics, and no excluded semantic work occurred.

## Verification

Confirm batch `PN03`, prefix `PRODUCT-TASK-SPEC-027-15-O`, exact source count, schemas, output isolation, and no stage or commit.

## Evidence

- PRODUCT-REQ-SPEC-013 owns gathering and schemas.
- T04 owns assignment and writer isolation.
- The user accepted per-batch lightweight Investigation Evidence because T37 concludes PRODUCT-INV-SPEC-011.
- `get_inventory_batch` confirmed batch `PN03`, 30 assigned sources, observation prefix `PRODUCT-TASK-SPEC-027-15-O`, and the two owned result paths.
- Every assigned source received exactly one disposition through `record_inventory_file_result`.
- All 30 source snapshots remained `unchanged`; no batch-external corpus source was read.
- `PN03.observations.jsonl` contains 314 unnormalized `bp-wide-term-observation-v1` observations.
- `PN03.coverage.json` records 30 discovered and scanned sources, 0 failed sources, and 314 observations under `bp-wide-term-batch-v1`.
- `validate_inventory_batch` returned `valid: true` with no structural diagnostics.
- Extraction remained limited to observed-use evidence. No normalization, classification, definition, aggregation, clustering, use-case extraction, staging, or commit occurred.
