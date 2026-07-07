# PRODUCT-TASK-SPEC-027-10: Extract term observations for PT06

- **id**: PRODUCT-TASK-SPEC-027-10
- **status**: done
- **date**: 2026-07-04
- **work_item**: PRODUCT-WORK-SPEC-027
- **task_type**: investigation
- **estimate**: 1d
- **depends_on**:
  - PRODUCT-TASK-SPEC-027-04
- **outputs**:
  - product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/PT06.observations.jsonl
  - product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/PT06.coverage.json

## Goal

Extract governance-sensitive observed-use evidence from exactly one assigned batch into its owned observations JSONL and coverage JSON.

## Work

- Use `product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/batches/PT06.json` as exact source authority and call `get_inventory_batch` for `PT06`.
- Read no batch-external source.
- Apply PRODUCT-REQ-SPEC-013 and schemas `bp-wide-term-observation-v1` and `bp-wide-term-batch-v1`.
- Record every source through `record_inventory_file_result` as `observations`, `no_candidate`, or `failed`.
- Treat `changed` as non-blocking and finish with `validate_inventory_batch`.
- Change only the two owned result files and this Task record.
- Do not normalize, merge meanings, judge correctness, define terms, aggregate, cluster, extract use cases, modify PRODUCT-INV-SPEC-011, stage, or commit.

## Done condition

Every `PT06` source has one disposition, both outputs conform, validation has no structural diagnostics, and no excluded semantic work occurred.

## Verification

Confirm batch `PT06`, prefix `PRODUCT-TASK-SPEC-027-10-O`, exact source count, schemas, output isolation, and no stage or commit.

## Evidence

- PRODUCT-REQ-SPEC-013 owns gathering and schemas.
- T04 owns assignment and writer isolation.
- The user accepted per-batch lightweight Investigation Evidence because T37 concludes PRODUCT-INV-SPEC-011.
- DRMCP is non-operational under the current agent authoring policy. Filesystem fallback was used only to update this Task record.
- `get_inventory_batch` returned batch `PT06` with 27 exact sources, output ownership for `PT06.observations.jsonl` and `PT06.coverage.json`, and observation prefix `PRODUCT-TASK-SPEC-027-10-O`.
- Every assigned source received exactly one `observations` disposition through `record_inventory_file_result`.
- The batch produced 167 unnormalized observations under `bp-wide-term-observation-v1` and complete coverage under `bp-wide-term-batch-v1`.
- `validate_inventory_batch` returned `valid: true`, 27 disposed sources, 167 observations, and no structural diagnostics.
- All 27 source snapshots remained `unchanged`; no drift exception was required.
- No batch-external source was read for semantic extraction.
- No term normalization, meaning merge, correctness judgment, definition, aggregation, clustering, use-case extraction, stage, or commit occurred.
