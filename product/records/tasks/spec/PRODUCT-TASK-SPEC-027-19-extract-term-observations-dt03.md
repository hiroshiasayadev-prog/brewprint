# PRODUCT-TASK-SPEC-027-19: Extract term observations for DT03

- **id**: PRODUCT-TASK-SPEC-027-19
- **status**: done
- **date**: 2026-07-04
- **work_item**: PRODUCT-WORK-SPEC-027
- **task_type**: investigation
- **estimate**: 1d
- **depends_on**:
  - PRODUCT-TASK-SPEC-027-04
- **outputs**:
  - product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/DT03.observations.jsonl
  - product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/DT03.coverage.json

## Goal

Extract governance-sensitive observed-use evidence from exactly one assigned batch into its owned observations JSONL and coverage JSON.

## Work

- Use `product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/batches/DT03.json` as exact source authority and call `get_inventory_batch` for `DT03`.
- Read no batch-external source.
- Apply PRODUCT-REQ-SPEC-013 and schemas `bp-wide-term-observation-v1` and `bp-wide-term-batch-v1`.
- Record every source through `record_inventory_file_result` as `observations`, `no_candidate`, or `failed`.
- Treat `changed` as non-blocking and finish with `validate_inventory_batch`.
- Change only the two owned result files and this Task record.
- Do not normalize, merge meanings, judge correctness, define terms, aggregate, cluster, extract use cases, modify PRODUCT-INV-SPEC-011, stage, or commit.

## Done condition

Every `DT03` source has one disposition, both outputs conform, validation has no structural diagnostics, and no excluded semantic work occurred.

## Verification

Confirm batch `DT03`, prefix `PRODUCT-TASK-SPEC-027-19-O`, exact source count, schemas, output isolation, and no stage or commit.

## Evidence

- PRODUCT-REQ-SPEC-013 owns gathering and schemas.
- T04 owns assignment and writer isolation.
- The user accepted per-batch lightweight Investigation Evidence because T37 concludes PRODUCT-INV-SPEC-011.
- `get_inventory_batch` confirmed batch `DT03`, 16 exact sources, owned result paths, and observation prefix `PRODUCT-TASK-SPEC-027-19-O`.
- Every source received exactly one `observations` disposition through `record_inventory_file_result`; no source was recorded as `no_candidate` or `failed`.
- The owned JSONL contains 180 observations with IDs `PRODUCT-TASK-SPEC-027-19-O0001` through `PRODUCT-TASK-SPEC-027-19-O0180` and schema version `bp-wide-term-observation-v1`.
- The owned coverage file uses schema `bp-wide-term-batch-v1` and records 16 discovered, 16 scanned, 0 failed, and 180 observations.
- `validate_inventory_batch` returned `valid: true`, 16 disposed sources, 180 observations, and no structural diagnostics.
- Every current and recorded source snapshot was `unchanged`.
- Changes are limited to the two owned result files and this Task record.
- No normalization, meaning merge, correctness judgment, term definition, aggregation, clustering, use-case extraction, PRODUCT-INV-SPEC-011 modification, stage, or commit occurred.
- DRMCP is non-operational; filesystem authoring was used for this Task status and Evidence update.
