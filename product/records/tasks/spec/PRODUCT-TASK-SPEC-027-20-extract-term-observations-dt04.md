# PRODUCT-TASK-SPEC-027-20: Extract term observations for DT04

- **id**: PRODUCT-TASK-SPEC-027-20
- **status**: done
- **date**: 2026-07-04
- **work_item**: PRODUCT-WORK-SPEC-027
- **task_type**: investigation
- **estimate**: 1d
- **depends_on**:
  - PRODUCT-TASK-SPEC-027-04
- **outputs**:
  - product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/DT04.observations.jsonl
  - product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/DT04.coverage.json

## Goal

Extract governance-sensitive observed-use evidence from exactly one assigned batch into its owned observations JSONL and coverage JSON.

## Work

- Use `product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/batches/DT04.json` as exact source authority and call `get_inventory_batch` for `DT04`.
- Read no batch-external source.
- Apply PRODUCT-REQ-SPEC-013 and schemas `bp-wide-term-observation-v1` and `bp-wide-term-batch-v1`.
- Record every source through `record_inventory_file_result` as `observations`, `no_candidate`, or `failed`.
- Treat `changed` as non-blocking and finish with `validate_inventory_batch`.
- Change only the two owned result files and this Task record.
- Do not normalize synonyms, merge meanings, judge correctness, define terms, aggregate, cluster, extract use cases, modify PRODUCT-INV-SPEC-011, stage, or commit.

## Done condition

Every `DT04` source has one disposition, both outputs conform, validation has no structural diagnostics, and no excluded semantic work occurred.

## Verification

Confirm batch `DT04`, prefix `PRODUCT-TASK-SPEC-027-20-O`, exact source count, schemas, output isolation, and no stage or commit.

## Evidence

- Authority: `PRODUCT-REQ-SPEC-013` supplied the gathering criterion and schemas.
- Assignment: `PRODUCT-TASK-SPEC-027-04` supplied batch `DT04`, writer isolation, and prefix `PRODUCT-TASK-SPEC-027-20-O`.
- Access: Used `get_inventory_batch` for the exact `DT04` assignment.
- Source scope: Read only the 16 source files assigned by `DT04.json`.
- Coverage: Recorded one disposition for all 16 sources.
- Result: All 16 dispositions are `observations`; no source is `no_candidate` or `failed`.
- Observations: Wrote 194 unnormalized observed-use records using `bp-wide-term-observation-v1`.
- Coverage output: Wrote `DT04.coverage.json` using `bp-wide-term-batch-v1`.
- Validation: `validate_inventory_batch` returned `valid: true` with zero structural diagnostics.
- Drift: All 16 source snapshots remained `unchanged`.
- Output isolation: The inventory writer changed only `DT04.observations.jsonl` and `DT04.coverage.json`.
- Semantic boundary: No synonym normalization, meaning merge, correctness judgment, definition, aggregation, clustering, or use-case extraction occurred.
- Semantic validator: The standalone Task responsibility validator was unavailable in this session. No semantic validator pass is claimed.
- Git boundary: No stage or commit operation was performed.
- Follow-up: T37 concludes the shared `PRODUCT-INV-SPEC-011` from all batch Evidence.
