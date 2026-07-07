# PRODUCT-TASK-SPEC-027-29: Extract term observations for SK01

- **id**: PRODUCT-TASK-SPEC-027-29
- **status**: done
- **date**: 2026-07-04
- **work_item**: PRODUCT-WORK-SPEC-027
- **task_type**: investigation
- **estimate**: 1d
- **depends_on**:
  - PRODUCT-TASK-SPEC-027-04
- **outputs**:
  - product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/SK01.observations.jsonl
  - product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/SK01.coverage.json

## Goal

Extract governance-sensitive observed-use evidence from exactly one assigned batch into its owned observations JSONL and coverage JSON.

## Work

- Use `product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/batches/SK01.json` as the exact scope authority.
- Call `get_inventory_batch` for `SK01`.
- Follow PRODUCT-REQ-SPEC-013 and both accepted collection schemas.
- Use `record_inventory_file_result` for every source disposition.
- Record `observations`, `no_candidate`, or `failed`; treat `changed` as non-blocking.
- Run `validate_inventory_batch` after every source has one disposition.
- Read no source outside SK01 and write no result outside the two owned paths.
- Do not normalize synonyms, merge meanings, judge correctness, define terms, aggregate, cluster, extract use cases, modify PRODUCT-INV-SPEC-011, stage, or commit.

## Done condition

Every SK01 source has one disposition, both outputs conform to `bp-wide-term-observation-v1` and `bp-wide-term-batch-v1`, and validation has no structural diagnostics.

## Verification

Confirm prefix `PRODUCT-TASK-SPEC-027-29-O`, exact source count, output isolation, and no stage or commit.

## Evidence

- PRODUCT-REQ-SPEC-013 owns gathering and schemas.
- PRODUCT-TASK-SPEC-027-04 owns assignment and writer isolation.
- `get_inventory_batch` resolved SK01 with exactly 40 assigned sources and prefix `PRODUCT-TASK-SPEC-027-29-O`.
- `record_inventory_file_result` recorded one `observations` disposition for every assigned source.
- The result contains 401 observations with IDs `PRODUCT-TASK-SPEC-027-29-O0001` through `PRODUCT-TASK-SPEC-027-29-O0401`.
- `SK01.observations.jsonl` uses `bp-wide-term-observation-v1`.
- `SK01.coverage.json` uses `bp-wide-term-batch-v1` and records 40 discovered, 40 scanned, and 0 failed sources.
- `validate_inventory_batch` returned `valid: true` with no structural diagnostics.
- Every source snapshot was `unchanged` at validation time.
- Source-corpus reads were limited to the 40 SK01 paths. Governing contract and authoring records were read separately.
- Writes were limited to the two declared outputs and this Task lifecycle Evidence update. PRODUCT-INV-SPEC-011 was not modified.
- Scoped Git inspection found the two outputs and this Task unstaged. No stage or commit was performed.
- T37 concludes PRODUCT-INV-SPEC-011 from all batch Evidence.
