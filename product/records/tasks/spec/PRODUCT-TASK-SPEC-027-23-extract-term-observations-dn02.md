# PRODUCT-TASK-SPEC-027-23: Extract term observations for DN02

- **id**: PRODUCT-TASK-SPEC-027-23
- **status**: done
- **date**: 2026-07-04
- **work_item**: PRODUCT-WORK-SPEC-027
- **task_type**: investigation
- **estimate**: 1d
- **depends_on**:
  - PRODUCT-TASK-SPEC-027-04
- **outputs**:
  - product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/DN02.observations.jsonl
  - product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/DN02.coverage.json

## Goal

Extract governance-sensitive observed-use evidence from exactly one assigned batch into its owned observations JSONL and coverage JSON.

## Work

- Use `product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/batches/DN02.json` as exact source authority and call `get_inventory_batch` for `DN02`.
- Read no batch-external source.
- Apply PRODUCT-REQ-SPEC-013 and schemas `bp-wide-term-observation-v1` and `bp-wide-term-batch-v1`.
- Record every source through `record_inventory_file_result` as `observations`, `no_candidate`, or `failed`.
- Treat `changed` as non-blocking and finish with `validate_inventory_batch`.
- Change only the two owned result files and this Task record.
- Do not normalize synonyms, merge meanings, judge correctness, define terms, aggregate, cluster, extract use cases, modify PRODUCT-INV-SPEC-011, stage, or commit.

## Done condition

Every `DN02` source has one disposition, both outputs conform, validation has no structural diagnostics, and no excluded semantic work occurred.

## Verification

Confirm batch `DN02`, prefix `PRODUCT-TASK-SPEC-027-23-O`, exact source count, schemas, output isolation, and no stage or commit.

## Evidence

- Startup read `prompt_chappy.md` first. `CLAUDE.md` and `AGENTS.md` were not read.
- DRMCP is non-operational. The Task record update used filesystem fallback after reading the Task authoring and agent authoring standards.
- `get_inventory_batch` confirmed batch `DN02`, prefix `PRODUCT-TASK-SPEC-027-23-O`, 24 sources, and the two owned result paths.
- All 24 sources were recorded through `record_inventory_file_result` with result `observations`.
- The batch contains 107 observations, zero failed files, and zero unrecorded sources.
- Every source snapshot remained `unchanged` during extraction.
- `DN02.observations.jsonl` conforms to `bp-wide-term-observation-v1`.
- `DN02.coverage.json` conforms to `bp-wide-term-batch-v1` and records 24 discovered and scanned files.
- `validate_inventory_batch` returned `valid: true` with no structural diagnostics.
- Extraction preserved observed use without synonym normalization, meaning merges, correctness judgments, definitions, aggregation, clustering, or use-case extraction.
- PRODUCT-REQ-SPEC-013 owns gathering and schemas.
- PRODUCT-TASK-SPEC-027-04 owns assignment and writer isolation.
- T37 concludes the shared PRODUCT-INV-SPEC-011 from all batch Evidence.
- No stage or commit occurred.
