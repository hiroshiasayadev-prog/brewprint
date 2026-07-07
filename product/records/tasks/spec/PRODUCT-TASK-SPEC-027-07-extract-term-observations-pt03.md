# PRODUCT-TASK-SPEC-027-07: Extract term observations for PT03

- **id**: PRODUCT-TASK-SPEC-027-07
- **status**: done
- **date**: 2026-07-04
- **work_item**: PRODUCT-WORK-SPEC-027
- **task_type**: investigation
- **estimate**: 1d
- **depends_on**:
  - PRODUCT-TASK-SPEC-027-04
- **outputs**:
  - product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/PT03.observations.jsonl
  - product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/PT03.coverage.json

## Goal

Extract governance-sensitive observed-use evidence from exactly one assigned batch into its owned observations JSONL and coverage JSON.

## Work

- Use `product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/batches/PT03.json` as exact source authority and call `get_inventory_batch` for `PT03`.
- Read no batch-external source.
- Apply PRODUCT-REQ-SPEC-013 and schemas `bp-wide-term-observation-v1` and `bp-wide-term-batch-v1`.
- Record every source through `record_inventory_file_result` as `observations`, `no_candidate`, or `failed`.
- Treat `changed` as non-blocking and finish with `validate_inventory_batch`.
- Change only the two owned result files and this Task record.
- Do not normalize, merge meanings, judge correctness, define terms, aggregate, cluster, extract use cases, modify PRODUCT-INV-SPEC-011, stage, or commit.

## Done condition

Every `PT03` source has one disposition, both outputs conform, validation has no structural diagnostics, and no excluded semantic work occurred.

## Verification

Confirm batch `PT03`, prefix `PRODUCT-TASK-SPEC-027-07-O`, exact source count, schemas, output isolation, and no stage or commit.

## Evidence

- PRODUCT-REQ-SPEC-013 owns gathering and schemas.
- T04 owns assignment and writer isolation.
- The user accepted per-batch lightweight Investigation Evidence because T37 concludes PRODUCT-INV-SPEC-011.
- Access mode: DRMCP is non-operational. The Task record used filesystem fallback; batch outputs used the completed term-inventory MCP.
- `get_inventory_batch` confirmed batch `PT03`, 27 exact sources, and observation prefix `PRODUCT-TASK-SPEC-027-07-O`.
- Every assigned source received one `observations` disposition through `record_inventory_file_result`.
- `PT03.observations.jsonl` contains 145 `bp-wide-term-observation-v1` records with IDs `PRODUCT-TASK-SPEC-027-07-O0001` through `PRODUCT-TASK-SPEC-027-07-O0145`.
- `PT03.coverage.json` conforms to `bp-wide-term-batch-v1`: 27 discovered, 27 scanned, 0 failed, and 145 observations.
- `validate_inventory_batch` returned `valid: true` with no structural diagnostics. All 27 source snapshots remained `unchanged`.
- Changed files are limited to the two PT03 result files and this Task record.
- No normalization, meaning merge, correctness judgment, definition, aggregation, clustering, use-case extraction, staging, or commit occurred.
