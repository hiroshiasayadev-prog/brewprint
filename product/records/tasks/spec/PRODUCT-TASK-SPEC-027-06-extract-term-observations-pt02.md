# PRODUCT-TASK-SPEC-027-06: Extract term observations for PT02

- **id**: PRODUCT-TASK-SPEC-027-06
- **status**: done
- **date**: 2026-07-04
- **work_item**: PRODUCT-WORK-SPEC-027
- **task_type**: investigation
- **estimate**: 1d
- **depends_on**:
  - PRODUCT-TASK-SPEC-027-04
- **outputs**:
  - product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/PT02.observations.jsonl
  - product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/PT02.coverage.json

## Goal

Extract governance-sensitive observed-use evidence from exactly one assigned batch into its owned observations JSONL and coverage JSON.

## Work

- Treat `product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/batches/PT02.json` as the exact source authority.
- Call `get_inventory_batch` with batch ID `PT02` and `batches_directory: batches`.
- Read only sources returned for `PT02`.
- Apply the gathering criterion in PRODUCT-REQ-SPEC-013.
- Author observations through `record_inventory_file_result`.
- Use `bp-wide-term-observation-v1` and `bp-wide-term-batch-v1`.
- Record every assigned source as `observations`, `no_candidate`, or `failed`.
- Treat `changed` as non-blocking.
- Run `validate_inventory_batch` for `PT02` after all dispositions exist.
- Update only this Task record for its own status and Evidence.

This Task must not read outside `PT02`, write another batch, normalize synonyms, merge meanings, judge correctness, define terms, aggregate, cluster, extract use cases, modify PRODUCT-INV-SPEC-011, stage, or commit.

## Done condition

- Every `PT02` source has exactly one disposition.
- Both owned result files conform structurally.
- `validate_inventory_batch` reports no structural diagnostics.
- No semantic normalization, classification, aggregation, definition, or use-case extraction occurred.

## Verification

- Confirm batch ID `PT02` and prefix `PRODUCT-TASK-SPEC-027-06-O`.
- Confirm coverage count equals the exact batch source count.
- Confirm schemas and exact output paths.
- Confirm only both batch-owned results and this Task record changed.
- Confirm no stage or commit occurred.

## Evidence

- PRODUCT-REQ-SPEC-013 owns the gathering criterion and schemas.
- PRODUCT-TASK-SPEC-027-04 owns the batch assignment and writer map.
- The user accepted per-batch Investigation Tasks under the lightweight Evidence exception because T37 concludes the shared Investigation.
- `get_inventory_batch` returned batch `PT02` with 24 assigned sources and observation prefix `PRODUCT-TASK-SPEC-027-06-O`.
- All 24 assigned sources were read and dispositioned through `record_inventory_file_result`; every source produced `observations`, no source failed, and every recorded snapshot was `unchanged`.
- Wrote 119 `bp-wide-term-observation-v1` observations to `product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/PT02.observations.jsonl`.
- Wrote `bp-wide-term-batch-v1` coverage to `product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/PT02.coverage.json`: 24 discovered, 24 scanned, 0 failed, and 119 observations.
- `validate_inventory_batch` returned `valid: true`, 24 of 24 sources disposed, and no structural diagnostics.
- Observation IDs span `PRODUCT-TASK-SPEC-027-06-O0001` through `PRODUCT-TASK-SPEC-027-06-O0119`.
- Gathering remained source-observational: no synonym normalization, meaning merge, correctness judgment, term definition, aggregation, clustering, or use-case extraction occurred.
- Only the two PT02 result files and this Task record were changed. No stage or commit occurred.
