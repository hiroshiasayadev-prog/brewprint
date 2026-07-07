# PRODUCT-TASK-SPEC-027-30: Extract term observations for LA01

- **id**: PRODUCT-TASK-SPEC-027-30
- **status**: done
- **date**: 2026-07-04
- **work_item**: PRODUCT-WORK-SPEC-027
- **task_type**: investigation
- **estimate**: 1d
- **depends_on**:
  - PRODUCT-TASK-SPEC-027-04
- **outputs**:
  - product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/LA01.observations.jsonl
  - product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/LA01.coverage.json

## Goal

Extract governance-sensitive observed-use evidence from exactly one assigned batch into its owned observations JSONL and coverage JSON.

## Work

- Use `product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/batches/LA01.json` as exact scope authority and call `get_inventory_batch` for `LA01`.
- Follow PRODUCT-REQ-SPEC-013, `bp-wide-term-observation-v1`, and `bp-wide-term-batch-v1`.
- Record every source through `record_inventory_file_result` as `observations`, `no_candidate`, or `failed`.
- Treat `changed` as non-blocking and run `validate_inventory_batch` after all dispositions exist.
- Read no batch-external source and write only the two owned result files plus this Task Evidence.
- Do not normalize, merge, judge, define, aggregate, cluster, extract use cases, modify PRODUCT-INV-SPEC-011, stage, or commit.

## Done condition

Every LA01 source has one disposition, both outputs conform, and validation has no structural diagnostics.

## Verification

Confirm prefix `PRODUCT-TASK-SPEC-027-30-O`, exact source count, schema use, output isolation, and no stage or commit.

## Evidence

- PRODUCT-REQ-SPEC-013 owns gathering and schemas.
- PRODUCT-TASK-SPEC-027-04 owns assignment and writer isolation.
- T37 concludes PRODUCT-INV-SPEC-011 from all batch Evidence.
- `get_inventory_batch` confirmed exact LA01 scope ownership: 20 assigned legacy ADR sources, observation prefix `PRODUCT-TASK-SPEC-027-30-O`, and the two Task-owned result paths.
- Every assigned source was read and disposed exactly once through `record_inventory_file_result` as `observations`.
- The batch produced 95 `bp-wide-term-observation-v1` observations with IDs `PRODUCT-TASK-SPEC-027-30-O0001` through `PRODUCT-TASK-SPEC-027-30-O0095`.
- `LA01.coverage.json` records `files_discovered: 20`, `files_scanned: 20`, `files_failed: 0`, and `observation_count: 95` under `bp-wide-term-batch-v1`.
- All 20 source snapshots remained `unchanged`; no source drift diagnostic was produced.
- `validate_inventory_batch` returned `valid: true`, 20 disposed sources, 95 observations, and zero structural diagnostics.
- Writes were limited to `LA01.observations.jsonl`, `LA01.coverage.json`, and this Task Evidence/status. No aggregation, normalization, classification, definition, use-case extraction, stage, or commit occurred.
