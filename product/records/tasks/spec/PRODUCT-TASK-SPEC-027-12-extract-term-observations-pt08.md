# PRODUCT-TASK-SPEC-027-12: Extract term observations for PT08

- **id**: PRODUCT-TASK-SPEC-027-12
- **status**: done
- **date**: 2026-07-04
- **work_item**: PRODUCT-WORK-SPEC-027
- **task_type**: investigation
- **estimate**: 1d
- **depends_on**:
  - PRODUCT-TASK-SPEC-027-04
- **outputs**:
  - product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/PT08.observations.jsonl
  - product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/PT08.coverage.json

## Goal

Extract governance-sensitive observed-use evidence from exactly one assigned batch into its owned observations JSONL and coverage JSON.

## Work

- Use `product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/batches/PT08.json` as exact source authority and call `get_inventory_batch` for `PT08`.
- Read no batch-external source.
- Apply PRODUCT-REQ-SPEC-013 and schemas `bp-wide-term-observation-v1` and `bp-wide-term-batch-v1`.
- Record every source through `record_inventory_file_result` as `observations`, `no_candidate`, or `failed`.
- Treat `changed` as non-blocking and finish with `validate_inventory_batch`.
- Change only the two owned result files and this Task record.
- Do not normalize, merge meanings, judge correctness, define terms, aggregate, cluster, extract use cases, modify PRODUCT-INV-SPEC-011, stage, or commit.

## Done condition

Every `PT08` source has one disposition, both outputs conform, validation has no structural diagnostics, and no excluded semantic work occurred.

## Verification

Confirm batch `PT08`, prefix `PRODUCT-TASK-SPEC-027-12-O`, exact source count, schemas, output isolation, and no stage or commit.

## Evidence

- PRODUCT-REQ-SPEC-013 owns gathering and schemas.
- T04 owns assignment and writer isolation.
- The user accepted per-batch lightweight Investigation Evidence because T37 concludes PRODUCT-INV-SPEC-011.
- Batch: `PT08`.
- Assignment schema: `bp-wide-term-batch-assignment-v1`.
- Observation schema: `bp-wide-term-observation-v1`.
- Coverage schema: `bp-wide-term-batch-v1`.
- Observation ID prefix: `PRODUCT-TASK-SPEC-027-12-O`.
- Source files discovered, scanned, and failed: `26`, `26`, and `0`.
- Source dispositions: `26` with observations, `0` no-candidate, and `0` failed.
- Observations recorded: `171`.
- Snapshot drift: all `26` sources remained unchanged.
- `validate_inventory_batch` result: `valid`; structural diagnostics: none.
- Outputs:
  - `product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/PT08.observations.jsonl`
  - `product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/PT08.coverage.json`
- Scope remained limited to PT08 source authority.
- Scoped Git inspection covered exactly the Task and two result files; all three are untracked and unstaged.
- Scoped whitespace inspection passed with no findings. LF-to-CRLF conversion warnings are non-blocking repository normalization advisories.
- No normalization, meaning merge, correctness judgment, definition, aggregation, clustering, use-case extraction, Investigation modification, staging, or commit occurred.
