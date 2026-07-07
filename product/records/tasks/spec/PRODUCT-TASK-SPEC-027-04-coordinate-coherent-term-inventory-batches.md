# PRODUCT-TASK-SPEC-027-04: Coordinate coherent term-inventory batches

- **id**: PRODUCT-TASK-SPEC-027-04
- **status**: done
- **date**: 2026-07-04
- **work_item**: PRODUCT-WORK-SPEC-027
- **task_type**: coordination
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-027-02
  - PRODUCT-TASK-SPEC-027-03
- **outputs**:
  - PRODUCT-WORK-SPEC-027
  - product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/manifest.json
  - product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/batches/
  - PRODUCT-TASK-SPEC-027-05 through PRODUCT-TASK-SPEC-027-38

## Goal

Persist the current best-effort corpus manifest, materialize the accepted coherent-32 batch assignments, and create the exact parallel extraction and completion Task graph.

## Work

- Use the completed `tools/term-inventory-mcp/` implementation rather than the temporary planner as the authoritative collection interface for this execution snapshot.
- Build a manifest over exactly these path sets:
  - `product/records/**/*.md`;
  - `drmcp/records/**/*.md`;
  - `bpdsl/records/**/*.md`;
  - `trv/records/**/*.md`;
  - `skills/**/*.md`;
  - `v01/records/adr/**/*.md`;
  - `v01/records/investigations/**/*.md`.
- Exclude all other `v01/records/**`, implementation source, scripts, generated outputs, build artifacts, PRODUCT-TASK-SPEC-027-03 through PRODUCT-TASK-SPEC-027-38, and PRODUCT-INV-SPEC-011 itself.
- Retain PRODUCT-TASK-SPEC-027-01, PRODUCT-TASK-SPEC-027-02, and PRODUCT-WORK-SPEC-027 because they precede and define the execution snapshot.
- Persist the manifest at `product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/manifest.json`.
- Partition the current manifest using the accepted coherent-32 profile and persist one exact batch file per batch under `batches/<batch_id>.json`.
- Preserve these batch IDs and Task mappings:
  - T05–T12: `PT01`–`PT08`;
  - T13–T16: `PN01`–`PN04`;
  - T17–T21: `DT01`–`DT05`;
  - T22–T24: `DN01`–`DN03`;
  - T25–T26: `TR01`–`TR02`;
  - T27–T28: `BP01`–`BP02`;
  - T29: `SK01`;
  - T30–T34: `LA01`–`LA05`;
  - T35–T36: `LI01`–`LI02`.
- Create T05–T36 as parallel `investigation` Tasks. Each Task must:
  - depend on T04;
  - own exactly one batch ID;
  - read only its batch assignment and source files;
  - write only `results/<batch_id>.observations.jsonl` and `results/<batch_id>.coverage.json`;
  - use the gathering criterion and schemas from PRODUCT-REQ-SPEC-013;
  - avoid term classification, normalization, definition, aggregation, and use-case extraction.
- Create T37 as the post-extraction `investigation` Task that depends on T05–T36, runs structural validation over every batch, records factual corpus and coverage totals, and concludes PRODUCT-INV-SPEC-011 without producing a merged master JSONL or semantic classification.
- Create T38 as the final `synchronization` Task that depends on T37 and updates Requirement and Work Item lifecycle and closure Evidence mechanically.
- Update PRODUCT-WORK-SPEC-027 with the materialized Task list, exact dependency flow, and current manifest facts.

## Drift handling

- The persisted manifest is a best-effort execution snapshot.
- Record the actual file count, total bytes, unreadable-file count, and generation timestamp produced by the completed MCP.
- Do not require those values to equal the earlier planner result of 718 files and 6,539,750 bytes.
- A source change after manifest creation is a `snapshot_status: changed` diagnostic during extraction, not an automatic batch failure.
- A source missing or unreadable at extraction time must remain an explicit per-file coverage result.
- Do not silently add new files discovered after manifest persistence. A later scope refresh requires an explicit coordination decision.

## Done condition

- One manifest and exactly 32 non-overlapping batch JSON files exist.
- Every manifest source appears in exactly one batch.
- Every batch uses one accepted batch ID and has one extraction Task owner.
- T05–T36, T37, and T38 exist with the accepted dependencies and responsibilities.
- PRODUCT-WORK-SPEC-027 records the complete materialized graph.
- No source observation, semantic classification, merged JSONL, or final Investigation conclusion is authored by this Task.

## Verification

- Run the MCP's structural checks over the manifest and batch assignments.
- Confirm all selected sources are assigned once and only once.
- Confirm all 32 expected batch IDs exist and no additional batch ID exists.
- Confirm every extraction Task has exactly one batch, one JSONL output, and one coverage output.
- Confirm no two extraction Tasks share a writable result path.
- Confirm T37 depends on all and only T05–T36.
- Confirm T38 depends on T37.
- Inspect the scoped Git worktree for the Work Item, manifest, batch files, and newly materialized Tasks.

## Evidence

- PRODUCT-TASK-SPEC-027-01 D-005 accepts the measured coherent-32 family-preserving split.
- PRODUCT-TASK-SPEC-027-01 D-006 fixes one-writer ownership per batch.
- PRODUCT-TASK-SPEC-027-01 D-007 fixes per-file coverage and source-drift representation.
- PRODUCT-TASK-SPEC-027-01 D-008 fixes the Task ID ranges and dependency route.
- The earlier planner measured zero unreadable UTF-8 files and approximately 38,899–60,852 estimated input tokens per coherent batch. These values are planning Evidence, not immutable manifest requirements.
- DRMCP is non-operational. Filesystem authoring was used for Design Record reads and writes because the supported DRMCP authoring path is unavailable.
- The operational term-inventory MCP created `manifest.json` and the coherent-32 batch assignments.
- The first persisted manifest at `2026-07-04T02:14:33+09:00` contained 736 files and 6,600,094 bytes, but post-materialization review found that it admitted only the then-existing subset of this workflow's generated records. It is superseded as extraction authority.
- The corrected manifest was persisted at `2026-07-04T03:14:40+09:00` with explicit self-exclusions.
- Corrected manifest facts: 733 files, 6,642,319 bytes, and 0 unreadable UTF-8 files.
- The earlier 718-file and 6,539,750-byte planner values remain comparison Evidence only; the corrected manifest is a current best-effort snapshot rather than a forced reconstruction of the planner snapshot.
- `partition_inventory_manifest` returned 32 batches, 733 manifest sources, and 733 assigned sources.
- The MCP partition implementation rejected duplicate, missing, extra, and invalid-batch assignments before returning success.
- The persisted `batches/` directory contains exactly PT01–PT08, PN01–PN04, DT01–DT05, DN01–DN03, TR01–TR02, BP01–BP02, SK01, LA01–LA05, and LI01–LI02.
- PT01–PT08, PN01–PN04, DT01–DT05, and DN01–DN03 were re-partitioned from the corrected manifest before their extraction Tasks began.
- TR01–TR02, BP01–BP02, SK01, LA01–LA05, and LI01–LI02 retained the same file counts and byte totals as the first partition. Their canonical assignment files were atomically rewritten with unchanged batch IDs and output ownership while T25–T36 were allowed to continue; no result file was modified or deleted by this correction.
- Live validation passed for completed BP01, BP02, LA01, and LI01 results. In-progress SK01 and TR01 reported only missing dispositions and no coverage-scope mismatch, confirming alignment with the corrected canonical assignments.
- `get_inventory_batch` returned every one of the 32 accepted batch assignments after bounded retries.
- Every batch has the expected bucket, exact source list, Task-owned observation prefix, and exactly two output paths.
- T05–T36 were created as parallel batch extraction Tasks. T37 depends on exactly T05–T36. T38 depends only on T37.
- Every extraction Task owns unique `results/<batch_id>.observations.jsonl` and `results/<batch_id>.coverage.json` paths. No writer conflict exists.
- T04 and this correction did not create or modify any observation JSONL or coverage result. T25–T36 may independently create their owned result files after the user started those Tasks.
- No semantic extraction, observation authoring, aggregation, clustering, use-case extraction, Investigation conclusion, lifecycle closure, stage, or commit was performed by T04 or this correction.
