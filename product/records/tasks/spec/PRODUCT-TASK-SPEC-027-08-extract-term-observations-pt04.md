# PRODUCT-TASK-SPEC-027-08: Extract term observations for PT04

- **id**: PRODUCT-TASK-SPEC-027-08
- **status**: done
- **date**: 2026-07-04
- **work_item**: PRODUCT-WORK-SPEC-027
- **task_type**: investigation
- **estimate**: 1d
- **depends_on**:
  - PRODUCT-TASK-SPEC-027-04
- **outputs**:
  - product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/PT04.observations.jsonl
  - product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/PT04.coverage.json

## Goal

Extract governance-sensitive observed-use evidence from exactly one assigned batch into its owned observations JSONL and coverage JSON.

## Work

- Use `product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/batches/PT04.json` as exact source authority and call `get_inventory_batch` for `PT04`.
- Read no batch-external source.
- Apply PRODUCT-REQ-SPEC-013 and schemas `bp-wide-term-observation-v1` and `bp-wide-term-batch-v1`.
- Record every source through `record_inventory_file_result` as `observations`, `no_candidate`, or `failed`.
- Treat `changed` as non-blocking and finish with `validate_inventory_batch`.
- Change only the two owned result files and this Task record.
- Do not normalize, merge meanings, judge correctness, define terms, aggregate, cluster, extract use cases, modify PRODUCT-INV-SPEC-011, stage, or commit.

## Done condition

Every `PT04` source has one disposition, both outputs conform, validation has no structural diagnostics, and no excluded semantic work occurred.

## Verification

Confirm batch `PT04`, prefix `PRODUCT-TASK-SPEC-027-08-O`, exact source count, schemas, output isolation, and no stage or commit.

## Evidence

### Authority and execution boundary

- PRODUCT-REQ-SPEC-013 owned the gathering criteria and schemas `bp-wide-term-observation-v1` and `bp-wide-term-batch-v1`.
- PRODUCT-TASK-SPEC-027-04 owned the `PT04` assignment and writer isolation.
- The user accepted per-batch lightweight Investigation Evidence because PRODUCT-TASK-SPEC-027-37 concludes PRODUCT-INV-SPEC-011.
- `get_inventory_batch` returned batch `PT04`, 27 exact sources, output prefix `PRODUCT-TASK-SPEC-027-08-O`, and the two owned result paths.
- Only assigned corpus sources were used for term observation extraction. Repository startup, governing Requirement, schema, and authoring-policy records were read only as execution authority.

### Extraction result

- All 27 assigned sources received exactly one `observations` disposition through `record_inventory_file_result`.
- `no_candidate`: 0 sources.
- `failed`: 0 sources.
- Observation count: 259.
- Observation ID range: `PRODUCT-TASK-SPEC-027-08-O0001` through `PRODUCT-TASK-SPEC-027-08-O0259`.
- Every recorded and current source snapshot status was `unchanged`.
- No term normalization, meaning merge, correctness judgment, preferred-term definition, aggregation, clustering, or use-case extraction was performed.

### Structural validation

- `validate_inventory_batch` returned `valid: true`.
- Batch source count: 27.
- Disposed source count: 27.
- Structural diagnostics: 0.
- Coverage reports `files_discovered: 27`, `files_scanned: 27`, `files_failed: 0`, and `observation_count: 259`.
- Output schemas are `bp-wide-term-observation-v1` and `bp-wide-term-batch-v1`.

### Write and release boundary

- Changed files are limited to `PT04.observations.jsonl`, `PT04.coverage.json`, and this Task record.
- Scoped Git whitespace inspection returned `pass`; LF-to-CRLF conversion notices were advisory only.
- No source Task, PRODUCT-INV-SPEC-011 record, sibling result, stage, or commit was changed.
