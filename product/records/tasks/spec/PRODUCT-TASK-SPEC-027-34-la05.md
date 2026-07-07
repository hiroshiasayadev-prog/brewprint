# PRODUCT-TASK-SPEC-027-34: Extract term observations for LA05

- **id**: PRODUCT-TASK-SPEC-027-34
- **status**: done
- **date**: 2026-07-04
- **work_item**: PRODUCT-WORK-SPEC-027
- **task_type**: investigation
- **estimate**: 1d
- **depends_on**:
  - PRODUCT-TASK-SPEC-027-04
- **outputs**:
  - product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/LA05.observations.jsonl
  - product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/LA05.coverage.json

## Goal

Extract governance-sensitive observed-use evidence from exactly one assigned batch into its owned observations JSONL and coverage JSON.

## Work

- Treat `product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/batches/LA05.json` as the exact source authority.
- Call `get_inventory_batch` for `LA05`.
- Follow PRODUCT-REQ-SPEC-013 and collection schemas `bp-wide-term-observation-v1` and `bp-wide-term-batch-v1`.
- Use `record_inventory_file_result` to record each source as `observations`, `no_candidate`, or `failed`.
- Treat `changed` as a non-blocking diagnostic.
- Run `validate_inventory_batch` after every source has one disposition.
- Read no source outside LA05 and write no other batch result.
- Do not normalize synonyms, merge meanings, judge correctness, define terms, aggregate, cluster, extract use cases, modify PRODUCT-INV-SPEC-011, stage, or commit.

## Done condition

Every LA05 source has one disposition, both owned outputs conform, and validation has no structural diagnostics.

## Verification

Confirm batch `LA05`, prefix `PRODUCT-TASK-SPEC-027-34-O`, exact source count, schemas, output isolation, and no stage or commit.

## Evidence

- PRODUCT-REQ-SPEC-013 owns the gathering criterion and schemas.
- PRODUCT-TASK-SPEC-027-04 owns batch assignment and writer isolation.
- T37 concludes PRODUCT-INV-SPEC-011 from all batch Evidence.
- DRMCP is non-operational. Filesystem access was used for Design Record reads and this Task update.
- `get_inventory_batch` returned `LA05`, bucket `legacy_adrs`, 20 sources, and prefix `PRODUCT-TASK-SPEC-027-34-O`.
- Every assigned source received exactly one `observations` disposition through `record_inventory_file_result`.
- `LA05.observations.jsonl` contains 73 observations using `bp-wide-term-observation-v1` and IDs `PRODUCT-TASK-SPEC-027-34-O0001` through `PRODUCT-TASK-SPEC-027-34-O0073`.
- `LA05.coverage.json` uses `bp-wide-term-batch-v1` and records 20 discovered files, 20 scanned files, 0 failed files, and 73 observations.
- `validate_inventory_batch` returned `valid: true`, 20 disposed sources, 73 observations, and no structural diagnostics.
- Every source snapshot remained `unchanged`.
- TRV-ADR-SPEC-006 suspends semantic-validator delivery and prohibits the deprecated prompt validator from acting as a completion gate. No semantic-validator invocation occurred.
- This execution wrote only the two LA05 result files and this Task record. PRODUCT-INV-SPEC-011 was not modified.
- No normalization, synonym merging, meaning consolidation, correctness judgment, definition, aggregation, clustering, use-case extraction, staging, or commit occurred.
