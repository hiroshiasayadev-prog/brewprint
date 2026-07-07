# PRODUCT-TASK-SPEC-027-18: Extract term observations for DT02

- **id**: PRODUCT-TASK-SPEC-027-18
- **status**: done
- **date**: 2026-07-04
- **work_item**: PRODUCT-WORK-SPEC-027
- **task_type**: investigation
- **estimate**: 1d
- **depends_on**:
  - PRODUCT-TASK-SPEC-027-04
- **outputs**:
  - product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/DT02.observations.jsonl
  - product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/DT02.coverage.json

## Goal

Extract governance-sensitive observed-use evidence from exactly one assigned batch into its owned observations JSONL and coverage JSON.

## Work

- Use `product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/batches/DT02.json` as exact source authority and call `get_inventory_batch` for `DT02`.
- Read no batch-external source.
- Apply PRODUCT-REQ-SPEC-013 and schemas `bp-wide-term-observation-v1` and `bp-wide-term-batch-v1`.
- Record every source through `record_inventory_file_result` as `observations`, `no_candidate`, or `failed`.
- Treat `changed` as non-blocking and finish with `validate_inventory_batch`.
- Change only the two owned result files and this Task record.
- Do not normalize, merge meanings, judge correctness, define terms, aggregate, cluster, extract use cases, modify PRODUCT-INV-SPEC-011, stage, or commit.

## Done condition

Every `DT02` source has one disposition, both outputs conform, validation has no structural diagnostics, and no excluded semantic work occurred.

## Verification

Confirm batch `DT02`, prefix `PRODUCT-TASK-SPEC-027-18-O`, exact source count, schemas, output isolation, and no stage or commit.

## Evidence

- PRODUCT-REQ-SPEC-013 owns gathering and schemas.
- T04 owns assignment and writer isolation.
- The user accepted per-batch lightweight Investigation Evidence because T37 concludes PRODUCT-INV-SPEC-011.
- Startup instruction: `prompt_chappy.md` was the first repository file read.
- `CLAUDE.md` and `AGENTS.md` were not read.
- Inventory interface: completed through the operational term-inventory MCP.
- Batch: `DT02`.
- Batch schema: `bp-wide-term-batch-assignment-v1` with profile `coherent-32`.
- Observation ID prefix: `PRODUCT-TASK-SPEC-027-18-O`.
- Exact assigned source count: 16.
- Source dispositions: 16 `observations`, 0 `no_candidate`, 0 `failed`.
- Observation schema: `bp-wide-term-observation-v1`.
- Coverage schema: `bp-wide-term-batch-v1`.
- Observation count: 139.
- Observations output: `product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/DT02.observations.jsonl`.
- Coverage output: `product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/DT02.coverage.json`.
- Final `validate_inventory_batch` result: `valid: true`.
- Structural diagnostics: none.
- Snapshot result: 15 sources `unchanged`; `DRMCP-TASK-MCP-016-08` reported `changed` and remained non-blocking under the Task contract.
- No normalization, meaning merge, correctness judgment, definition, aggregation, clustering, use-case extraction, or PRODUCT-INV-SPEC-011 modification occurred.
- Changed files are limited to the two owned result files and this Task record.
- Scoped Git inspection found exactly those three files, all unstaged and untracked in the current worktree scope.
- Scoped whitespace inspection passed; LF-to-CRLF notices were advisory only.
- The standalone semantic responsibility validator was not executed because no operational invocation tool is available in this session. It was not treated as PASS.
- Stage and commit were not performed.
