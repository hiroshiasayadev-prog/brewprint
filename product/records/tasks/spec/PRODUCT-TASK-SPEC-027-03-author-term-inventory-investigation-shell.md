# PRODUCT-TASK-SPEC-027-03: Author term-inventory Investigation shell

- **id**: PRODUCT-TASK-SPEC-027-03
- **status**: done
- **date**: 2026-07-04
- **work_item**: PRODUCT-WORK-SPEC-027
- **task_type**: authoring
- **estimate**: 0.25d
- **depends_on**:
  - PRODUCT-TASK-SPEC-027-01
- **outputs**:
  - PRODUCT-INV-SPEC-011

## Goal

Create the durable Investigation shell that owns the Brewprint design-governance term inventory method, scope, machine-readable Evidence contract, and later concluded findings.

## Work

- Create `PRODUCT-INV-SPEC-011` using the canonical Investigation shape.
- Set status to `investigating`.
- Use PRODUCT-REQ-SPEC-013 as the direct research source and PRODUCT-WORK-SPEC-027 as the related Work Item.
- Record the exact included and excluded path sets decided by PRODUCT-TASK-SPEC-027-01.
- Record `bp-wide-term-observation-v1` and `bp-wide-term-batch-v1` as the collection schemas.
- Record the machine-readable Evidence root:
  - `product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/manifest.json`;
  - `product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/batches/<batch_id>.json`;
  - `product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/<batch_id>.observations.jsonl`;
  - `product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/<batch_id>.coverage.json`.
- State that source hashes and modification metadata are best-effort snapshots rather than immutable authority.
- Leave findings and recommendation as `TBD` until all extraction batches and final structural validation complete.

This Task must not create the manifest, assign source files to batches, extract observations, merge JSONL, classify terms, or conclude the Investigation.

## Done condition

- PRODUCT-INV-SPEC-011 exists with status `investigating`.
- Its scope and non-scope match PRODUCT-TASK-SPEC-027-01 D-001.
- Its method records the gathering criterion and schemas from PRODUCT-REQ-SPEC-013.
- Its Evidence contract matches PRODUCT-TASK-SPEC-027-01 D-003 and D-004.
- No research finding, semantic classification, or use-case conclusion is claimed prematurely.

## Verification

- Confirm every canonical Investigation section is present.
- Confirm Investigation metadata contains no Task ID in `source_refs`, `follow_up_candidates`, or `follow_up_results`.
- Confirm PRODUCT-REQ-SPEC-013 is a direct source and PRODUCT-WORK-SPEC-027 is recorded as related work.
- Confirm the Investigation remains `investigating`.
- Confirm no batch result file is authored by this Task.

## Evidence

- PRODUCT-TASK-SPEC-027-01 D-003 reserves PRODUCT-INV-SPEC-011 and its machine-readable Evidence root.
- PRODUCT-TASK-SPEC-027-01 D-004 assigns Investigation-shell authoring to one bounded authoring Task.
- Final findings and conclusion remain owned by the post-extraction conclusion Task.
- DRMCP is non-operational. Filesystem authoring is the required fallback.
- `product/records/investigations/spec/PRODUCT-INV-SPEC-011-brewprint-design-governance-term-inventory.md` now contains the Investigation shell.
- The Investigation metadata omits an explicit `id` field and contains no Task ID in `source_refs`, `follow_up_candidates`, or `follow_up_results`.
- All ten canonical Investigation sections are present.
- The Investigation remains `investigating`; findings, recommendation, and follow-up content remain deferred until extraction completion.
- The shell records the selected physical path set, gathering criterion, collection schemas, and machine-readable Evidence contract.
- This Task created no manifest, batch assignment, observations JSONL, coverage JSON, merged JSONL, semantic aggregation, or vocabulary classification.
