# PRODUCT-TASK-SPEC-027-13: Extract term observations for PN01

- **id**: PRODUCT-TASK-SPEC-027-13
- **status**: done
- **date**: 2026-07-04
- **work_item**: PRODUCT-WORK-SPEC-027
- **task_type**: investigation
- **estimate**: 1d
- **depends_on**:
  - PRODUCT-TASK-SPEC-027-04
- **outputs**:
  - product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/PN01.observations.jsonl
  - product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/PN01.coverage.json

## Goal

Extract governance-sensitive observed-use evidence from exactly one assigned batch into its owned observations JSONL and coverage JSON.

## Work

- Use `product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/batches/PN01.json` as exact source authority and call `get_inventory_batch` for `PN01`.
- Read no batch-external source.
- Apply PRODUCT-REQ-SPEC-013 and schemas `bp-wide-term-observation-v1` and `bp-wide-term-batch-v1`.
- Record every source through `record_inventory_file_result` as `observations`, `no_candidate`, or `failed`.
- Treat `changed` as non-blocking and finish with `validate_inventory_batch`.
- Change only the two owned result files and this Task record.
- Do not normalize, merge meanings, judge correctness, define terms, aggregate, cluster, extract use cases, modify PRODUCT-INV-SPEC-011, stage, or commit.

## Done condition

Every `PN01` source has one disposition, both outputs conform, validation has no structural diagnostics, and no excluded semantic work occurred.

## Verification

Confirm batch `PN01`, prefix `PRODUCT-TASK-SPEC-027-13-O`, exact source count, schemas, output isolation, and no stage or commit.

## Evidence

- PRODUCT-REQ-SPEC-013 owns gathering and schemas.
- T04 owns assignment and writer isolation.
- The user accepted per-batch lightweight Investigation Evidence because T37 concludes PRODUCT-INV-SPEC-011.
- `get_inventory_batch` confirmed batch `PN01`, bucket `product_non_tasks`, 28 assigned sources, and observation ID prefix `PRODUCT-TASK-SPEC-027-13-O`.
- All 28 sources received exactly one `observations` disposition. The batch contains 238 observations, 0 `no_candidate` dispositions, and 0 failures.
- The MCP wrote `results/PN01.observations.jsonl` using `bp-wide-term-observation-v1` and `results/PN01.coverage.json` using `bp-wide-term-batch-v1`.
- `validate_inventory_batch` returned `valid: true`, 28 of 28 disposed sources, 238 observations, and no structural diagnostics.
- `product/records/work-items/spec/PRODUCT-WORK-SPEC-027-inventory-brewprint-design-governance-terms.md` reported `changed` snapshot drift. The drift was non-blocking under this Task contract. All other source snapshots remained unchanged.
- Extraction preserved source-specific observed meanings and consequences. No normalization, definition, aggregation, clustering, use-case extraction, or vocabulary judgment occurred.
- Only the two batch-owned result files and this Task record changed. Stage and commit were not performed.
