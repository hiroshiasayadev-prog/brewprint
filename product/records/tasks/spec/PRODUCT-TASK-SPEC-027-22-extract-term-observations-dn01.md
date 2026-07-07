# PRODUCT-TASK-SPEC-027-22: Extract term observations for DN01

- **id**: PRODUCT-TASK-SPEC-027-22
- **status**: done
- **date**: 2026-07-04
- **work_item**: PRODUCT-WORK-SPEC-027
- **task_type**: investigation
- **estimate**: 1d
- **depends_on**:
  - PRODUCT-TASK-SPEC-027-04
- **outputs**:
  - product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/DN01.observations.jsonl
  - product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/DN01.coverage.json

## Goal

Extract governance-sensitive observed-use evidence from exactly one assigned batch into its owned observations JSONL and coverage JSON.

## Work

- Use `product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/batches/DN01.json` as exact source authority and call `get_inventory_batch` for `DN01`.
- Read no batch-external source.
- Apply PRODUCT-REQ-SPEC-013 and schemas `bp-wide-term-observation-v1` and `bp-wide-term-batch-v1`.
- Record every source through `record_inventory_file_result` as `observations`, `no_candidate`, or `failed`.
- Treat `changed` as non-blocking and finish with `validate_inventory_batch`.
- Change only the two owned result files and this Task record.
- Do not normalize synonyms, merge meanings, judge correctness, define terms, aggregate, cluster, extract use cases, modify PRODUCT-INV-SPEC-011, stage, or commit.

## Done condition

Every `DN01` source has one disposition, both outputs conform, validation has no structural diagnostics, and no excluded semantic work occurred.

## Verification

Confirm batch `DN01`, prefix `PRODUCT-TASK-SPEC-027-22-O`, exact source count, schemas, output isolation, and no stage or commit.

## Evidence

- PRODUCT-REQ-SPEC-013 owns gathering and schemas.
- PRODUCT-TASK-SPEC-027-04 owns assignment and writer isolation.
- T37 concludes the shared PRODUCT-INV-SPEC-011 from all batch Evidence.
- `get_inventory_batch` returned batch `DN01`, 24 exact sources, and observation prefix `PRODUCT-TASK-SPEC-027-22-O`.
- All 24 sources were recorded exactly once through `record_inventory_file_result`; every disposition is `observations`.
- The owned outputs contain 151 observations and complete coverage metadata.
- `validate_inventory_batch` returned `valid: true`, `disposed_source_count: 24`, `batch_source_count: 24`, and no structural diagnostics.
- `drmcp/records/spec/design-records-mcp/tools/get-authoring-guidance.md` was reported as `changed`; the accepted batch contract treats this drift as non-blocking.
- Observation and coverage outputs conform to `bp-wide-term-observation-v1` and `bp-wide-term-batch-v1` through batch validation.
- Extraction used only the 24 DN01 corpus sources. PRODUCT-REQ-SPEC-013 and authoring policy were read only as governing contracts.
- No synonym normalization, meaning merge, correctness judgment, definition, aggregation, clustering, or use-case extraction occurred.
- DRMCP authoring transactions are non-operational. The Task record was updated through the filesystem fallback under the current agent-authoring policy.
- No stage or commit was performed.
