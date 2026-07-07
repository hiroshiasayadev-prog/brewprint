# PRODUCT-TASK-SPEC-027-11: Extract term observations for PT07

- **id**: PRODUCT-TASK-SPEC-027-11
- **status**: done
- **date**: 2026-07-04
- **work_item**: PRODUCT-WORK-SPEC-027
- **task_type**: investigation
- **estimate**: 1d
- **depends_on**:
  - PRODUCT-TASK-SPEC-027-04
- **outputs**:
  - product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/PT07.observations.jsonl
  - product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/PT07.coverage.json

## Goal

Extract governance-sensitive observed-use evidence from exactly one assigned batch into its owned observations JSONL and coverage JSON.

## Work

- Use `product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/batches/PT07.json` as exact source authority and call `get_inventory_batch` for `PT07`.
- Read no batch-external source.
- Apply PRODUCT-REQ-SPEC-013 and schemas `bp-wide-term-observation-v1` and `bp-wide-term-batch-v1`.
- Record every source through `record_inventory_file_result` as `observations`, `no_candidate`, or `failed`.
- Treat `changed` as non-blocking and finish with `validate_inventory_batch`.
- Change only the two owned result files and this Task record.
- Do not normalize, merge meanings, judge correctness, define terms, aggregate, cluster, extract use cases, modify PRODUCT-INV-SPEC-011, stage, or commit.

## Done condition

Every `PT07` source has one disposition, both outputs conform, validation has no structural diagnostics, and no excluded semantic work occurred.

## Verification

Confirm batch `PT07`, prefix `PRODUCT-TASK-SPEC-027-11-O`, exact source count, schemas, output isolation, and no stage or commit.

## Evidence

- PRODUCT-REQ-SPEC-013 owns gathering and schemas.
- T04 owns assignment and writer isolation.
- The user accepted per-batch lightweight Investigation Evidence because T37 concludes PRODUCT-INV-SPEC-011.
### Result

`PASS`.

### Batch execution

- Batch: `PT07`.
- Assignment schema: `bp-wide-term-batch-assignment-v1`.
- Source count: 27.
- Observation ID prefix: `PRODUCT-TASK-SPEC-027-11-O`.
- Observations recorded: 173.
- Source dispositions: 27 `observations`, 0 `no_candidate`, 0 `failed`.
- Observation schema: `bp-wide-term-observation-v1`.
- Coverage schema: `bp-wide-term-batch-v1`.

### Outputs

- `product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/PT07.observations.jsonl`
- `product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/PT07.coverage.json`

### Validation

- `validate_inventory_batch` returned `valid: true`.
- Disposed sources: 27 of 27.
- Structural diagnostics: 0.
- Coverage summary: 27 discovered, 27 scanned, 0 failed, 173 observations.
- `PRODUCT-TASK-SPEC-027-01` reported `snapshot_status: changed`.
- The changed snapshot is non-blocking under the accepted batch contract.

### Boundary confirmation

- Read only the 27 sources assigned by `PT07.json`.
- Wrote only the two PT07 result files and this Task record.
- Performed no normalization, meaning merge, correctness judgment, definition, aggregation, clustering, use-case extraction, stage, or commit.
