# PRODUCT-TASK-SPEC-027-26: Extract term observations for TR02

- **id**: PRODUCT-TASK-SPEC-027-26
- **status**: done
- **date**: 2026-07-04
- **work_item**: PRODUCT-WORK-SPEC-027
- **task_type**: investigation
- **estimate**: 1d
- **depends_on**:
  - PRODUCT-TASK-SPEC-027-04
- **outputs**:
  - product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/TR02.observations.jsonl
  - product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/TR02.coverage.json

## Goal

Extract governance-sensitive observed-use evidence from exactly one assigned batch into its owned observations JSONL and coverage JSON.

## Work

- Treat `product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/batches/TR02.json` as exact scope authority.
- Call `get_inventory_batch` for `TR02`.
- Follow PRODUCT-REQ-SPEC-013, `bp-wide-term-observation-v1`, and `bp-wide-term-batch-v1`.
- Record each source through `record_inventory_file_result` as `observations`, `no_candidate`, or `failed`.
- Treat `changed` as non-blocking and run `validate_inventory_batch` after all dispositions exist.
- Read no batch-external source and write no other batch result.
- Do not normalize, merge, judge, define, aggregate, cluster, extract use cases, modify PRODUCT-INV-SPEC-011, stage, or commit.

## Done condition

Every `TR02` source has one disposition, both owned outputs conform, and validation has no structural diagnostics.

## Verification

Confirm batch `TR02`, prefix `PRODUCT-TASK-SPEC-027-26-O`, exact source count, schemas, output isolation, and no stage or commit.

## Evidence

- PRODUCT-REQ-SPEC-013 owns gathering and schemas.
- PRODUCT-TASK-SPEC-027-04 owns assignment and writer isolation.
- T37 concludes PRODUCT-INV-SPEC-011 from all batch Evidence.
- `get_inventory_batch` returned batch `TR02`, bucket `trv_all`, 34 assigned sources, observation prefix `PRODUCT-TASK-SPEC-027-26-O`, and the two declared result paths.
- Every one of the 34 assigned sources received exactly one `observations` disposition through `record_inventory_file_result`; `files_failed` is 0.
- The owned JSONL contains 198 observations with IDs `PRODUCT-TASK-SPEC-027-26-O0001` through `PRODUCT-TASK-SPEC-027-26-O0198` using `bp-wide-term-observation-v1`.
- The owned coverage file uses `bp-wide-term-batch-v1` and records `files_discovered: 34`, `files_scanned: 34`, `files_failed: 0`, and `observation_count: 198`.
- Every recorded and current source snapshot status is `unchanged`.
- `validate_inventory_batch` returned `valid: true` with 34 assigned sources, 34 disposed sources, 198 observations, and no diagnostics.
- Head and tail inspection confirmed the required schema version, source path, observation prefix, and contiguous first and last observation IDs.
- Only `results/TR02.observations.jsonl` and `results/TR02.coverage.json` were written as batch outputs. No other batch result or PRODUCT-INV-SPEC-011 content was modified.
- No term normalization, meaning merge, correctness judgment, definition, aggregation, clustering, or use-case extraction was performed.
- DRMCP is non-operational, so filesystem fallback was used only to update this Task record.
- Stage and commit were not performed.
- Result: `PASS`.
