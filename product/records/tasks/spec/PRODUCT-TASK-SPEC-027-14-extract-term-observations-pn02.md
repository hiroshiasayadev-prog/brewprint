# PRODUCT-TASK-SPEC-027-14: Extract term observations for PN02

- **id**: PRODUCT-TASK-SPEC-027-14
- **status**: done
- **date**: 2026-07-04
- **work_item**: PRODUCT-WORK-SPEC-027
- **task_type**: investigation
- **estimate**: 1d
- **depends_on**:
  - PRODUCT-TASK-SPEC-027-04
- **outputs**:
  - product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/PN02.observations.jsonl
  - product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/PN02.coverage.json

## Goal

Extract governance-sensitive observed-use evidence from exactly one assigned batch into its owned observations JSONL and coverage JSON.

## Work

- Use `product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/batches/PN02.json` as exact source authority and call `get_inventory_batch` for `PN02`.
- Read no batch-external source.
- Apply PRODUCT-REQ-SPEC-013 and schemas `bp-wide-term-observation-v1` and `bp-wide-term-batch-v1`.
- Record every source through `record_inventory_file_result` as `observations`, `no_candidate`, or `failed`.
- Treat `changed` as non-blocking and finish with `validate_inventory_batch`.
- Change only the two owned result files and this Task record.
- Do not normalize, merge meanings, judge correctness, define terms, aggregate, cluster, extract use cases, modify PRODUCT-INV-SPEC-011, stage, or commit.

## Done condition

Every `PN02` source has one disposition, both outputs conform, validation has no structural diagnostics, and no excluded semantic work occurred.

## Verification

Confirm batch `PN02`, prefix `PRODUCT-TASK-SPEC-027-14-O`, exact source count, schemas, output isolation, and no stage or commit.

## Evidence

- PRODUCT-REQ-SPEC-013 owns gathering and schemas.
- T04 owns assignment and writer isolation.
- The user accepted per-batch lightweight Investigation Evidence because T37 concludes PRODUCT-INV-SPEC-011.
- `get_inventory_batch` confirmed `PN02`, 29 source files, and prefix `PRODUCT-TASK-SPEC-027-14-O`.
- All 29 sources received one `observations` disposition through `record_inventory_file_result`; none failed.
- The owned outputs contain 262 `bp-wide-term-observation-v1` observations and `bp-wide-term-batch-v1` coverage metadata.
- `validate_inventory_batch` returned `valid: true`, 29 of 29 sources disposed, 262 observations, and no structural diagnostics.
- All source snapshots remained `unchanged`.
- No normalization, definition, aggregation, clustering, or use-case extraction was performed.
- Only the two owned result files and this Task record were changed. No stage or commit was performed.
