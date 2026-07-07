# PRODUCT-TASK-SPEC-027-28: Extract term observations for BP02

- **id**: PRODUCT-TASK-SPEC-027-28
- **status**: done
- **date**: 2026-07-04
- **work_item**: PRODUCT-WORK-SPEC-027
- **task_type**: investigation
- **estimate**: 1d
- **depends_on**:
  - PRODUCT-TASK-SPEC-027-04
- **outputs**:
  - product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/BP02.observations.jsonl
  - product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/BP02.coverage.json

## Goal

Extract governance-sensitive observed-use evidence from exactly one assigned batch into its owned observations JSONL and coverage JSON.

## Work

- Use `product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/batches/BP02.json` as the exact scope authority.
- Call `get_inventory_batch` for `BP02`.
- Follow PRODUCT-REQ-SPEC-013 and both accepted collection schemas.
- Use `record_inventory_file_result` for every source disposition.
- Record `observations`, `no_candidate`, or `failed`; treat `changed` as non-blocking.
- Run `validate_inventory_batch` after every source has one disposition.
- Read no source outside BP02 and write no result outside the two owned paths.
- Do not normalize synonyms, merge meanings, judge correctness, define terms, aggregate, cluster, extract use cases, modify PRODUCT-INV-SPEC-011, stage, or commit.

## Done condition

Every BP02 source has one disposition, both outputs conform to `bp-wide-term-observation-v1` and `bp-wide-term-batch-v1`, and validation has no structural diagnostics.

## Verification

Confirm prefix `PRODUCT-TASK-SPEC-027-28-O`, exact source count, output isolation, and no stage or commit.

## Evidence

### Result

- `PASS`.
- `get_inventory_batch` returned BP02 with 19 assigned sources.
- `record_inventory_file_result` recorded one `observations` disposition for every assigned source.
- The batch contains 132 observations and zero failed sources.
- Every recorded source snapshot remained `unchanged`.

### Outputs

- `product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/BP02.observations.jsonl`
- `product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/BP02.coverage.json`

### Verification

- `validate_inventory_batch` returned `valid: true` with zero diagnostics.
- Coverage records 19 discovered files, 19 scanned files, and 0 failed files.
- Observation IDs use prefix `PRODUCT-TASK-SPEC-027-28-O` from `O0001` through `O0132`.
- Both outputs use the assigned BP02 paths and accepted schemas.
- No source outside BP02 was used for extraction.
- No result was written outside the two batch-owned output paths.

### Responsibility boundary

- The Task produced one bounded BP02 research result and one completion judgment.
- The Task did not normalize synonyms, merge meanings, judge correctness, define terms, aggregate, cluster, or extract use cases.
- T37 remains the owner of cross-batch validation and PRODUCT-INV-SPEC-011 conclusion.
- The abandoned automated responsibility-validator assets under `skills/DEPRECATED-task-responsibility-boundary-validator/` were not used as authority.
- A Task-local responsibility check against `spec:product.design_records.authoring_standards.task_authoring` returned `PASS`.

### Execution route

- PRODUCT-REQ-SPEC-013 supplied the gathering criterion and accepted schemas.
- PRODUCT-TASK-SPEC-027-04 supplied exact assignment and writer isolation.
- The term-inventory MCP performed all batch result writes and structural validation.
- DRMCP is non-operational. Filesystem authoring updated only this Task lifecycle and Evidence.
- Stage and commit were not performed.
