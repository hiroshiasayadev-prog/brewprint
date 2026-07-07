# PRODUCT-TASK-SPEC-027-32: Extract term observations for LA03

- **id**: PRODUCT-TASK-SPEC-027-32
- **status**: done
- **date**: 2026-07-04
- **work_item**: PRODUCT-WORK-SPEC-027
- **task_type**: investigation
- **estimate**: 1d
- **depends_on**:
  - PRODUCT-TASK-SPEC-027-04
- **outputs**:
  - product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/LA03.observations.jsonl
  - product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/LA03.coverage.json

## Goal

Extract governance-sensitive observed-use evidence from exactly one assigned batch into its owned observations JSONL and coverage JSON.

## Work

- Treat `product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/batches/LA03.json` as the exact source authority.
- Call `get_inventory_batch` for `LA03`.
- Follow PRODUCT-REQ-SPEC-013 and collection schemas `bp-wide-term-observation-v1` and `bp-wide-term-batch-v1`.
- Use `record_inventory_file_result` to record each source as `observations`, `no_candidate`, or `failed`.
- Treat `changed` as a non-blocking diagnostic.
- Run `validate_inventory_batch` after every source has one disposition.
- Read no source outside LA03 and write no other batch result.
- Do not normalize synonyms, merge meanings, judge correctness, define terms, aggregate, cluster, extract use cases, modify PRODUCT-INV-SPEC-011, stage, or commit.

## Done condition

Every LA03 source has one disposition, both owned outputs conform, and validation has no structural diagnostics.

## Verification

Confirm batch `LA03`, prefix `PRODUCT-TASK-SPEC-027-32-O`, exact source count, schemas, output isolation, and no stage or commit.

## Evidence

- PRODUCT-REQ-SPEC-013 owns the gathering criterion and schemas.
- PRODUCT-TASK-SPEC-027-04 owns batch assignment and writer isolation.
- T37 concludes PRODUCT-INV-SPEC-011 from all batch Evidence.
- `get_inventory_batch` resolved LA03 to 20 legacy ADR sources, 202,012 bytes, and observation prefix `PRODUCT-TASK-SPEC-027-32-O`.
- Every assigned source received exactly one `observations` disposition through `record_inventory_file_result`.
- The batch produced 151 observations with IDs `PRODUCT-TASK-SPEC-027-32-O0001` through `PRODUCT-TASK-SPEC-027-32-O0151`.
- All 20 source snapshots remained `unchanged`; no source failed and no source used `no_candidate`.
- `validate_inventory_batch` returned `valid: true`, 20 discovered sources, 20 dispositions, 151 observations, and no diagnostics.
- `LA03.coverage.json` uses `bp-wide-term-batch-v1`; every JSONL entry uses `bp-wide-term-observation-v1`.
- Writes were limited to the two LA03 result files and this Task Evidence and lifecycle update.
- PRODUCT-INV-SPEC-011, other batches, aggregation artifacts, stage state, and commit history were not modified.
- Post-Evidence responsibility validation applied all 14 common and 9 `investigation` criteria to this Task; all 23 criterion results were compliant.
