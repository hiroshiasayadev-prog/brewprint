# PRODUCT-TASK-SPEC-027-05: Extract term observations for PT01

- **id**: PRODUCT-TASK-SPEC-027-05
- **status**: done
- **date**: 2026-07-04
- **work_item**: PRODUCT-WORK-SPEC-027
- **task_type**: investigation
- **estimate**: 1d
- **depends_on**:
  - PRODUCT-TASK-SPEC-027-04
- **outputs**:
  - product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/PT01.observations.jsonl
  - product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/PT01.coverage.json

## Goal

Extract governance-sensitive observed-use evidence from exactly one assigned batch into its owned observations JSONL and coverage JSON.

## Work

- Treat `product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/batches/PT01.json` as the exact source authority.
- Call `get_inventory_batch` with batch ID `PT01` and `batches_directory: batches`.
- Read only sources returned for `PT01`.
- Apply the gathering criterion in PRODUCT-REQ-SPEC-013.
- Author observations through `record_inventory_file_result`.
- Use `bp-wide-term-observation-v1` for observations.
- Use `bp-wide-term-batch-v1` for coverage.
- Record every assigned source as `observations`, `no_candidate`, or `failed`.
- Treat `changed` as a non-blocking snapshot diagnostic.
- Run `validate_inventory_batch` for `PT01` after all assigned sources have a disposition.
- Update only this Task record for its own status and Evidence.

This Task must not:

- read a source outside batch `PT01`;
- write another batch's result files;
- normalize synonyms;
- merge meanings;
- judge correctness;
- define terms;
- aggregate or cluster observations;
- extract use cases;
- modify PRODUCT-INV-SPEC-011;
- stage or commit changes.

## Done condition

- Every source in batch `PT01` has exactly one recorded disposition.
- The owned observations JSONL and coverage JSON use the required schemas and paths.
- `validate_inventory_batch` reports no structural diagnostics.
- No semantic normalization, classification, aggregation, definition, or use-case extraction occurred.

## Verification

- Confirm `get_inventory_batch` returns batch ID `PT01` and observation prefix `PRODUCT-TASK-SPEC-027-05-O`.
- Confirm coverage source count equals the exact batch source count.
- Confirm all observations use `bp-wide-term-observation-v1`.
- Confirm coverage uses `bp-wide-term-batch-v1`.
- Confirm only the two batch-owned result paths and this Task record changed.
- Confirm no stage or commit occurred.

## Evidence

- PRODUCT-REQ-SPEC-013 owns the gathering criterion and observation schema.
- PRODUCT-TASK-SPEC-027-04 owns the persisted batch assignment and writer map.
- The user explicitly accepted per-batch Investigation Tasks without separate Investigation records because T37 concludes the shared PRODUCT-INV-SPEC-011 from all batch Evidence.

### Execution result

| item | result |
|---|---|
| Batch | `PT01` |
| Observation ID prefix | `PRODUCT-TASK-SPEC-027-05-O` |
| Assigned sources | 22 |
| Recorded dispositions | 22 |
| `observations` dispositions | 22 |
| `no_candidate` dispositions | 0 |
| `failed` dispositions | 0 |
| Observation count | 135 |
| Structural diagnostics | 0 |
| Snapshot drift | 0 changed sources |
| Result | `PASS` |

### Output evidence

- `get_inventory_batch` returned batch `PT01` with 22 exact source paths.
- The batch returned observation prefix `PRODUCT-TASK-SPEC-027-05-O`.
- Only the 22 assigned sources were read.
- Every source received exactly one disposition through `record_inventory_file_result`.
- The MCP generated `bp-wide-term-observation-v1` observations.
- The MCP generated `bp-wide-term-batch-v1` coverage metadata.
- Observations were written to `product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/PT01.observations.jsonl`.
- Coverage was written to `product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/PT01.coverage.json`.
- `validate_inventory_batch` returned `valid: true`.
- Validation returned no structural diagnostics.
- All 22 current and recorded snapshots were `unchanged`.

### Boundary evidence

- Extraction preserved source-local phrases, observed meanings, and semantic consequences.
- No synonym normalization, meaning merge, correctness judgment, definition, aggregation, clustering, or use-case extraction occurred.
- PRODUCT-INV-SPEC-011 was not modified.
- No other batch result was written.
- No stage or commit occurred.
- DRMCP is non-operational. Filesystem authoring was the required Task-record fallback.
- The standalone Task responsibility validator was unavailable in the current tool set.
- `prompt_chappy.md` was read first. `CLAUDE.md` and `AGENTS.md` were not read.
