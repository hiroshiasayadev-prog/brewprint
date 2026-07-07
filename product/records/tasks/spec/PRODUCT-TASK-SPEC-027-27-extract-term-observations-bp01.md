# PRODUCT-TASK-SPEC-027-27: Extract term observations for BP01

- **id**: PRODUCT-TASK-SPEC-027-27
- **status**: done
- **date**: 2026-07-04
- **work_item**: PRODUCT-WORK-SPEC-027
- **task_type**: investigation
- **estimate**: 1d
- **depends_on**:
  - PRODUCT-TASK-SPEC-027-04
- **outputs**:
  - product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/BP01.observations.jsonl
  - product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/BP01.coverage.json

## Goal

Extract governance-sensitive observed-use evidence from exactly one assigned batch into its owned observations JSONL and coverage JSON.

## Work

- Treat `product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/batches/BP01.json` as exact scope authority.
- Call `get_inventory_batch` for `BP01`.
- Follow PRODUCT-REQ-SPEC-013, `bp-wide-term-observation-v1`, and `bp-wide-term-batch-v1`.
- Record each source through `record_inventory_file_result` as `observations`, `no_candidate`, or `failed`.
- Treat `changed` as non-blocking and run `validate_inventory_batch` after all dispositions exist.
- Read no batch-external source and write no other batch result.
- Do not normalize, merge, judge, define, aggregate, cluster, extract use cases, modify PRODUCT-INV-SPEC-011, stage, or commit.

## Done condition

Every `BP01` source has one disposition, both owned outputs conform, and validation has no structural diagnostics.

## Verification

Confirm batch `BP01`, prefix `PRODUCT-TASK-SPEC-027-27-O`, exact source count, schemas, output isolation, and no stage or commit.

## Evidence

- PRODUCT-REQ-SPEC-013 owns gathering and schemas.
- PRODUCT-TASK-SPEC-027-04 owns assignment and writer isolation.
- T37 concludes PRODUCT-INV-SPEC-011 from all batch Evidence.
- DRMCP is non-operational. Filesystem fallback was used for Task retrieval and Evidence authoring.
- `get_inventory_batch` returned batch `BP01`, 18 assigned sources, prefix `PRODUCT-TASK-SPEC-027-27-O`, and the two Task-owned output paths.
- All 18 assigned sources received exactly one `observations` disposition. No source received `no_candidate` or `failed`.
- `BP01.observations.jsonl` contains 147 observations using schema `bp-wide-term-observation-v1` and IDs `PRODUCT-TASK-SPEC-027-27-O0001` through `PRODUCT-TASK-SPEC-027-27-O0147`.
- `BP01.coverage.json` uses schema `bp-wide-term-batch-v1` and records 18 discovered, 18 scanned, 0 failed, and 147 observations.
- `validate_inventory_batch` returned `valid: true`, 18 disposed sources, 147 observations, and no structural diagnostics.
- All 18 current source snapshots and recorded snapshots were `unchanged`.
- Writes were isolated to `results/BP01.observations.jsonl`, `results/BP01.coverage.json`, and this Task Evidence.
- No normalization, merging, semantic judgment, definition, aggregation, clustering, use-case extraction, PRODUCT-INV-SPEC-011 modification, stage, or commit occurred.
