# PRODUCT-TASK-SPEC-027-31: Extract term observations for LA02

- **id**: PRODUCT-TASK-SPEC-027-31
- **status**: done
- **date**: 2026-07-04
- **work_item**: PRODUCT-WORK-SPEC-027
- **task_type**: investigation
- **estimate**: 1d
- **depends_on**:
  - PRODUCT-TASK-SPEC-027-04
- **outputs**:
  - product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/LA02.observations.jsonl
  - product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/LA02.coverage.json

## Goal

Extract governance-sensitive observed-use evidence from exactly one assigned batch into its owned observations JSONL and coverage JSON.

## Work

- Treat `product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/batches/LA02.json` as the exact source authority.
- Call `get_inventory_batch` for `LA02`.
- Follow PRODUCT-REQ-SPEC-013 and collection schemas `bp-wide-term-observation-v1` and `bp-wide-term-batch-v1`.
- Use `record_inventory_file_result` to record each source as `observations`, `no_candidate`, or `failed`.
- Treat `changed` as a non-blocking diagnostic.
- Run `validate_inventory_batch` after every source has one disposition.
- Read no source outside LA02 and write no other batch result.
- Do not normalize synonyms, merge meanings, judge correctness, define terms, aggregate, cluster, extract use cases, modify PRODUCT-INV-SPEC-011, stage, or commit.

## Done condition

Every LA02 source has one disposition, both owned outputs conform, and validation has no structural diagnostics.

## Verification

Confirm batch `LA02`, prefix `PRODUCT-TASK-SPEC-027-31-O`, exact source count, schemas, output isolation, and no stage or commit.

## Evidence

- Result: PASS.
- Batch: `LA02`.
- Source count: 19 discovered, 19 scanned, 0 failed.
- Dispositions: 19 `observations`, 0 `no_candidate`, 0 `failed`.
- Observation count: 111.
- Observation schema: `bp-wide-term-observation-v1`.
- Coverage schema: `bp-wide-term-batch-v1`.
- Observation ID prefix: `PRODUCT-TASK-SPEC-027-31-O`.
- Outputs:
  - `product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/LA02.observations.jsonl`
  - `product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/LA02.coverage.json`
- `validate_inventory_batch` returned `valid: true` with no structural diagnostics.
- All 19 source snapshots were `unchanged`.
- No source outside `LA02` was scanned for observation extraction.
- No other batch result, PRODUCT-INV-SPEC-011 content, stage, or commit was modified.
- PRODUCT-REQ-SPEC-013 owns the gathering criterion and schemas.
- PRODUCT-TASK-SPEC-027-04 owns batch assignment and writer isolation.
- T37 concludes PRODUCT-INV-SPEC-011 from all batch Evidence.
- Current DRMCP is non-operational. The Task lifecycle update used filesystem fallback after reading the Task authoring and agent authoring standards.
- Post-Evidence responsibility-boundary validation: PASS, 23 of 23 criteria compliant.
