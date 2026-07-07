# PRODUCT-TASK-SPEC-027-09: Extract term observations for PT05

- **id**: PRODUCT-TASK-SPEC-027-09
- **status**: done
- **date**: 2026-07-04
- **work_item**: PRODUCT-WORK-SPEC-027
- **task_type**: investigation
- **estimate**: 1d
- **depends_on**:
  - PRODUCT-TASK-SPEC-027-04
- **outputs**:
  - product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/PT05.observations.jsonl
  - product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/PT05.coverage.json

## Goal

Extract governance-sensitive observed-use evidence from exactly one assigned batch into its owned observations JSONL and coverage JSON.

## Work

- Use `product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/batches/PT05.json` as exact source authority and call `get_inventory_batch` for `PT05`.
- Read no batch-external source.
- Apply PRODUCT-REQ-SPEC-013 and schemas `bp-wide-term-observation-v1` and `bp-wide-term-batch-v1`.
- Record every source through `record_inventory_file_result` as `observations`, `no_candidate`, or `failed`.
- Treat `changed` as non-blocking and finish with `validate_inventory_batch`.
- Change only the two owned result files and this Task record.
- Do not normalize, merge meanings, judge correctness, define terms, aggregate, cluster, extract use cases, modify PRODUCT-INV-SPEC-011, stage, or commit.

## Done condition

Every `PT05` source has one disposition, both outputs conform, validation has no structural diagnostics, and no excluded semantic work occurred.

## Verification

Confirm batch `PT05`, prefix `PRODUCT-TASK-SPEC-027-09-O`, exact source count, schemas, output isolation, and no stage or commit.

## Evidence

- PRODUCT-REQ-SPEC-013 owns gathering and schemas.
- T04 owns assignment and writer isolation.
- The user accepted per-batch lightweight Investigation Evidence because T37 concludes PRODUCT-INV-SPEC-011.
- `get_inventory_batch` returned batch `PT05` with 26 exact sources, output paths `results/PT05.observations.jsonl` and `results/PT05.coverage.json`, and observation prefix `PRODUCT-TASK-SPEC-027-09-O`.
- Every assigned source received exactly one `observations` disposition through `record_inventory_file_result`; no source required `no_candidate` or `failed`.
- The batch records 165 observations in `bp-wide-term-observation-v1` shape and complete `bp-wide-term-batch-v1` coverage for all 26 sources.
- `validate_inventory_batch` returned `valid: true`, `disposed_source_count: 26`, `observation_count: 165`, and no structural diagnostics.
- `PRODUCT-TASK-SPEC-027-02` reported source drift as `changed`; the other 25 sources reported `unchanged`. The accepted contract treats drift as non-blocking.
- Semantic normalization, meaning merge, correctness judgment, definition, aggregation, clustering, and use-case extraction were not performed.
- DRMCP is non-operational. Filesystem fallback updated only this Task record; the term-inventory MCP wrote only the two PT05-owned result files.
- Stage and commit were not performed.
