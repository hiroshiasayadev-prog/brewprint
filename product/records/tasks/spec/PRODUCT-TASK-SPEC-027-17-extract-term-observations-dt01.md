# PRODUCT-TASK-SPEC-027-17: Extract term observations for DT01

- **id**: PRODUCT-TASK-SPEC-027-17
- **status**: done
- **date**: 2026-07-04
- **work_item**: PRODUCT-WORK-SPEC-027
- **task_type**: investigation
- **estimate**: 1d
- **depends_on**:
  - PRODUCT-TASK-SPEC-027-04
- **outputs**:
  - product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/DT01.observations.jsonl
  - product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/DT01.coverage.json

## Goal

Extract governance-sensitive observed-use evidence from exactly one assigned batch into its owned observations JSONL and coverage JSON.

## Work

- Use `product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/batches/DT01.json` as exact source authority and call `get_inventory_batch` for `DT01`.
- Read no batch-external source.
- Apply PRODUCT-REQ-SPEC-013 and schemas `bp-wide-term-observation-v1` and `bp-wide-term-batch-v1`.
- Record every source through `record_inventory_file_result` as `observations`, `no_candidate`, or `failed`.
- Treat `changed` as non-blocking and finish with `validate_inventory_batch`.
- Change only the two owned result files and this Task record.
- Do not normalize, merge meanings, judge correctness, define terms, aggregate, cluster, extract use cases, modify PRODUCT-INV-SPEC-011, stage, or commit.

## Done condition

Every `DT01` source has one disposition, both outputs conform, validation has no structural diagnostics, and no excluded semantic work occurred.

## Verification

Confirm batch `DT01`, prefix `PRODUCT-TASK-SPEC-027-17-O`, exact source count, schemas, output isolation, and no stage or commit.

## Evidence

- PRODUCT-REQ-SPEC-013 owns gathering and schemas.
- T04 owns assignment and writer isolation.
- The user accepted per-batch lightweight Investigation Evidence because T37 concludes PRODUCT-INV-SPEC-011.
- Repository-root `prompt_chappy.md` was read before other repository files. `CLAUDE.md` and `AGENTS.md` were not read.
- `get_inventory_batch` returned batch `DT01`, bucket `drmcp_tasks`, 16 exact sources, and observation prefix `PRODUCT-TASK-SPEC-027-17-O`.
- Every assigned source received exactly one `observations` disposition through `record_inventory_file_result`; no source received `no_candidate` or `failed`.
- The owned JSONL contains 123 observations with IDs `PRODUCT-TASK-SPEC-027-17-O0001` through `PRODUCT-TASK-SPEC-027-17-O0123` and schema version `bp-wide-term-observation-v1`.
- The owned coverage file uses schema `bp-wide-term-batch-v1` and records `files_discovered: 16`, `files_scanned: 16`, `files_failed: 0`, and `observation_count: 123`.
- `validate_inventory_batch` returned `valid: true`, `disposed_source_count: 16`, no structural diagnostics, and `unchanged` snapshot status for all 16 sources.
- Only `DT01.observations.jsonl`, `DT01.coverage.json`, and this Task record were changed.
- No normalization, meaning merge, correctness judgment, term definition, aggregation, clustering, use-case extraction, PRODUCT-INV-SPEC-011 modification, stage, or commit was performed.
- DRMCP authoring transactions are non-operational, so this Task lifecycle and Evidence update used filesystem fallback under the current agent-authoring policy.
