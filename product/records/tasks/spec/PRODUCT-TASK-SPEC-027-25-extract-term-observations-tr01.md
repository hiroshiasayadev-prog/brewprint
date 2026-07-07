# PRODUCT-TASK-SPEC-027-25: Extract term observations for TR01

- **id**: PRODUCT-TASK-SPEC-027-25
- **status**: done
- **date**: 2026-07-04
- **work_item**: PRODUCT-WORK-SPEC-027
- **task_type**: investigation
- **estimate**: 1d
- **depends_on**:
  - PRODUCT-TASK-SPEC-027-04
- **outputs**:
  - product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/TR01.observations.jsonl
  - product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/TR01.coverage.json

## Goal

Extract governance-sensitive observed-use evidence from exactly one assigned batch into its owned observations JSONL and coverage JSON.

## Work

- Treat `product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/batches/TR01.json` as exact scope authority.
- Call `get_inventory_batch` for `TR01`.
- Follow PRODUCT-REQ-SPEC-013, `bp-wide-term-observation-v1`, and `bp-wide-term-batch-v1`.
- Record each source through `record_inventory_file_result` as `observations`, `no_candidate`, or `failed`.
- Treat `changed` as non-blocking.
- Run `validate_inventory_batch` after all dispositions exist.
- Read no batch-external source and write no other batch result.
- Do not normalize, merge, judge, define, aggregate, cluster, extract use cases, modify PRODUCT-INV-SPEC-011, stage, or commit.

## Done condition

Every `TR01` source has one disposition, both owned outputs conform, and validation has no structural diagnostics.

## Verification

Confirm batch `TR01`, prefix `PRODUCT-TASK-SPEC-027-25-O`, exact source count, schemas, output isolation, and no stage or commit.

## Evidence

- PRODUCT-REQ-SPEC-013 owns gathering and schemas.
- PRODUCT-TASK-SPEC-027-04 owns assignment and writer isolation.
- T37 concludes PRODUCT-INV-SPEC-011 from all batch Evidence.
- `get_inventory_batch` returned batch `TR01`, bucket `trv_all`, 35 assigned sources, observation prefix `PRODUCT-TASK-SPEC-027-25-O`, and the two declared result paths.
- Every one of the 35 assigned sources received exactly one `observations` disposition through `record_inventory_file_result`; `files_failed` is 0.
- The owned JSONL contains 393 observations with IDs `PRODUCT-TASK-SPEC-027-25-O0001` through `PRODUCT-TASK-SPEC-027-25-O0393` using `bp-wide-term-observation-v1`.
- The owned coverage file uses `bp-wide-term-batch-v1` and records `files_discovered: 35`, `files_scanned: 35`, `files_failed: 0`, and `observation_count: 393`.
- Every recorded and current source snapshot status is `unchanged`.
- `validate_inventory_batch` returned `valid: true` with 35 assigned sources, 35 disposed sources, 393 observations, and no diagnostics.
- Head and tail inspection confirmed the required schema version, source path, observation prefix, and contiguous first and last observation IDs.
- Only `results/TR01.observations.jsonl` and `results/TR01.coverage.json` were written as batch outputs. No other batch result or PRODUCT-INV-SPEC-011 content was modified.
- No term normalization, meaning merge, correctness judgment, definition, aggregation, clustering, or use-case extraction was performed.
- DRMCP is non-operational, so filesystem fallback was used only to update this Task record.
- The active semantic validator authority is suspended. The retained deprecated common and `investigation` checklist assets were used only as an advisory responsibility-boundary check; the local evaluation returned all 23 criteria compliant, and direct Task-text comparison confirmed the result.
- Stage and commit were not performed.
- Result: `PASS`.
