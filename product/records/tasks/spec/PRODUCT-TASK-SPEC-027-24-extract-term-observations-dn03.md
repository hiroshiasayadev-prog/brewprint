# PRODUCT-TASK-SPEC-027-24: Extract term observations for DN03

- **id**: PRODUCT-TASK-SPEC-027-24
- **status**: done
- **date**: 2026-07-04
- **work_item**: PRODUCT-WORK-SPEC-027
- **task_type**: investigation
- **estimate**: 1d
- **depends_on**:
  - PRODUCT-TASK-SPEC-027-04
- **outputs**:
  - product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/DN03.observations.jsonl
  - product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/DN03.coverage.json

## Goal

Extract governance-sensitive observed-use evidence from exactly one assigned batch into its owned observations JSONL and coverage JSON.

## Work

- Use `product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/batches/DN03.json` as exact source authority and call `get_inventory_batch` for `DN03`.
- Read no batch-external source.
- Apply PRODUCT-REQ-SPEC-013 and schemas `bp-wide-term-observation-v1` and `bp-wide-term-batch-v1`.
- Record every source through `record_inventory_file_result` as `observations`, `no_candidate`, or `failed`.
- Treat `changed` as non-blocking and finish with `validate_inventory_batch`.
- Change only the two owned result files and this Task record.
- Do not normalize synonyms, merge meanings, judge correctness, define terms, aggregate, cluster, extract use cases, modify PRODUCT-INV-SPEC-011, stage, or commit.

## Done condition

Every `DN03` source has one disposition, both outputs conform, validation has no structural diagnostics, and no excluded semantic work occurred.

## Verification

Confirm batch `DN03`, prefix `PRODUCT-TASK-SPEC-027-24-O`, exact source count, schemas, output isolation, and no stage or commit.

## Evidence

- Startup read `prompt_chappy.md` first. `CLAUDE.md` and `AGENTS.md` were not read.
- DRMCP is non-operational. The Task record update used filesystem fallback after reading the Task authoring and agent authoring standards.
- `get_inventory_batch` confirmed batch `DN03`, bucket `drmcp_non_tasks`, prefix `PRODUCT-TASK-SPEC-027-24-O`, 24 sources, and the two owned result paths.
- All 24 assigned sources were recorded exactly once through `record_inventory_file_result` with result `observations`.
- The owned JSONL contains 177 observations with IDs `PRODUCT-TASK-SPEC-027-24-O0001` through `PRODUCT-TASK-SPEC-027-24-O0177` and schema `bp-wide-term-observation-v1`.
- The owned coverage file uses `bp-wide-term-batch-v1` and records 24 discovered files, 24 scanned files, zero failed files, and 177 observations.
- Snapshot status was `changed` for `drmcp/records/spec/design-records-mcp/schema/authoring-guidance-source.md` and `drmcp/records/spec/design-records-mcp/tools/list-authoring-guides.md`; the remaining 22 sources were `unchanged`. The two changed snapshots were treated as non-blocking under this Task contract.
- `validate_inventory_batch` returned `valid: true` with 24 assigned sources, 24 disposed sources, 177 observations, and no structural diagnostics.
- Head and tail inspection confirmed schema version, source path, observation prefix, and IDs `PRODUCT-TASK-SPEC-027-24-O0001` and `PRODUCT-TASK-SPEC-027-24-O0177`.
- Extraction preserved observed use without synonym normalization, meaning merges, correctness judgments, definitions, aggregation, clustering, or use-case extraction.
- PRODUCT-REQ-SPEC-013 owns gathering and schemas.
- PRODUCT-TASK-SPEC-027-04 owns assignment and writer isolation.
- T37 concludes the shared PRODUCT-INV-SPEC-011 from all batch Evidence.
- Only the two owned result files and this Task record were changed in this execution scope.
- No stage or commit occurred.
- Result: `PASS`.
