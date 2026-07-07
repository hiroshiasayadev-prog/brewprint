# PRODUCT-TASK-SPEC-027-01: Decide term-inventory execution

- **id**: PRODUCT-TASK-SPEC-027-01
- **status**: done
- **date**: 2026-07-04
- **work_item**: PRODUCT-WORK-SPEC-027
- **task_type**: decision
- **estimate**: 0.5d
- **depends_on**: []
- **outputs**:
  - PRODUCT-WORK-SPEC-027
  - PRODUCT-TASK-SPEC-027-01
  - PRODUCT-TASK-SPEC-027-02
  - PRODUCT-TASK-SPEC-027-03
  - PRODUCT-TASK-SPEC-027-04

## Goal

Fix one executable, parallel-safe investigation route for producing the JSONL term-observation corpus required by PRODUCT-REQ-SPEC-013.

## Work

- Decide the exact active corpus scope and exclusions.
- Decide raw JSONL and batch-coverage output placement.
- Decide the durable research-artifact and extraction Task responsibility model.
- Partition the selected corpus into complete, non-overlapping batches.
- Fix writer ownership so parallel sessions do not share output files.
- Decide per-file no-candidate, failed-file, and coverage accounting representation.
- Materialize only the implementation, extraction, and final coverage Tasks uniquely required by the accepted decisions.

This Task does not scan source files, author term observations, aggregate observations, implement helper tools, classify terms, or define vocabulary.

### Decision ledger

| ID | topic | status | decision summary | reason | dependency | downstream owner or target |
|---|---|---|---|---|---|---|
| D-001 | Corpus scope | `decided` | Include `product/records/**/*.md`, `drmcp/records/**/*.md`, `bpdsl/records/**/*.md`, `trv/records/**/*.md`, `skills/**/*.md`, `v01/records/adr/**/*.md`, and `v01/records/investigations/**/*.md`. Exclude all other `v01/records/**`, implementation source, scripts, generated outputs, build artifacts, this workflow's generated T03–T38 records, and PRODUCT-INV-SPEC-011 itself. Keep T01, T02, and PRODUCT-WORK-SPEC-027 in scope because they precede and define the execution snapshot. Treat scope as a physical path set rather than requiring a prior definition of Design Records. | Active records and workflow instructions contain current governance semantics. The current workflow's generated shell, coordination, extraction, conclusion, and closure artifacts must not become self-observing inventory sources. Legacy ADRs and Investigations preserve high-value provenance without doubling the corpus with historical Tasks, Work Items, and migrated Specifications. | none | Work Item Boundary and extraction batches |
| D-002 | Inventory-support MCP | `decided` | Add one lightweight implementation Task before extraction. Implement manifest creation, deterministic batch partitioning, batch retrieval, result recording, and structural validation under `tools/term-inventory-mcp/`. Treat source hashes and manifest state as best-effort snapshots because concurrent sessions may change source files. | Direct JSONL writing still needs a constrained authoring interface and validator to avoid unusable observations. Source immutability is not achievable during parallel repository work. | D-001 | implementation Task and extraction prerequisites |
| D-003 | Output placement | `decided` | Reserve `PRODUCT-INV-SPEC-011` as the durable Investigation record. Store machine-readable Evidence under `product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/`, with `manifest.json`, `batches/<batch_id>.json`, `results/<batch_id>.observations.jsonl`, and `results/<batch_id>.coverage.json`. Do not create a merged master JSONL in this Work Item. | The Investigation record owns the research method and concluded facts, while per-batch data files preserve machine-readable source Evidence and independent writer ownership. | D-001, D-002 | PRODUCT-INV-SPEC-011 and batch outputs |
| D-004 | Research artifact and Task responsibility | `decided` | One authoring Task creates the Investigation shell and data directory contract. Each extraction Task owns exactly one batch scope and exactly two result files: one observations JSONL and one coverage JSON. The Investigation remains `investigating` until all batches and final coverage reconciliation complete. | This separates durable research narration from parallel Evidence production without making aggregation or semantic analysis part of the inventory outcome. | D-001 through D-003 | Investigation authoring Task, extraction Tasks, and final coverage Task |
| D-005 | Batch partitioning | `decided` | Adopt the measured `coherent-32` profile: `PT01`–`PT08`, `PN01`–`PN04`, `DT01`–`DT05`, `DN01`–`DN03`, `TR01`–`TR02`, `BP01`–`BP02`, `SK01`, `LA01`–`LA05`, and `LI01`–`LI02`. The exact source files come from the persisted batch JSON files, not prose duplication in Task records. | The measured profile covers 718 files and 6,539,750 bytes with zero unreadable UTF-8 files. Each batch is approximately 39k–61k estimated input tokens and remains within one coherent source family. | D-001 through D-004 | manifest, batch files, and extraction Task graph |
| D-006 | Parallel writer ownership | `decided` | Each extraction Task owns exactly one batch ID, one `results/<batch_id>.observations.jsonl`, and one `results/<batch_id>.coverage.json`. No extraction Task may write another batch's files. T03 creates the Investigation shell. The final conclusion Task updates the Investigation only after every extraction Task completes. | One-writer ownership prevents parallel conflicts while allowing sequential Investigation shell creation and conclusion. | D-005 | extraction Tasks and Investigation writers |
| D-007 | Coverage accounting | `decided` | Every coverage file uses `bp-wide-term-batch-v1`, preserves the Requirement summary fields, and includes one per-source result with `source_path`, `result` (`observations`, `no_candidate`, or `failed`), `observation_count`, `snapshot_status` (`unchanged`, `changed`, `missing`, or `unreadable`), and optional `failure`. Summary counts must reconcile with the per-source results. `changed` is diagnostic and does not itself fail the batch. | Complete per-file disposition is required to prove scope coverage while tolerating best-effort source snapshots. | D-001 through D-006 | collection MCP, batch coverage files, and final validation |
| D-008 | Task graph | `decided` | Materialize T03 to create PRODUCT-INV-SPEC-011 and the data contract, then T04 to generate the manifest and coherent-32 batch files and materialize T05–T36 as parallel `investigation` Tasks. Reserve T05–T12 for PT01–PT08, T13–T16 for PN01–PN04, T17–T21 for DT01–DT05, T22–T24 for DN01–DN03, T25–T26 for TR01–TR02, T27–T28 for BP01–BP02, T29 for SK01, T30–T34 for LA01–LA05, and T35–T36 for LI01–LI02. Reserve T37 for final structural validation and Investigation conclusion, and T38 for lifecycle closure synchronization. T02 and T03 may run in parallel; T04 depends on both. | The route separates tooling, research-shell authoring, graph materialization, parallel extraction, final validation, and lifecycle closure into distinct completion judgments. | D-001 through D-007 | PRODUCT-WORK-SPEC-027 Task graph |

### Current cursor

- Decision: none
- Loop state: `decision_complete`
- Every owned decision is terminal.

## Done condition

- D-001 through D-008 are `decided`, `deferred`, or validly `blocked`.
- PRODUCT-TASK-SPEC-027-03 and PRODUCT-TASK-SPEC-027-04 are materialized from the accepted route.
- The selected corpus has explicit inclusions and exclusions.
- Every output has one writer owner.
- Batch scopes are non-overlapping and collectively cover the selected corpus.
- No-candidate and failed-file accounting are fixed.
- The exact required Task graph is materialized without speculative semantic aggregation, classification, use-case extraction, or vocabulary-definition Tasks.

## Verification

- Confirm every accepted user answer appears once in the ledger.
- Confirm the selected corpus can be enumerated mechanically from the persisted scope.
- Confirm each source file belongs to at most one extraction batch.
- Confirm each output file has exactly one writer Task.
- Confirm Work Item Completion Condition remains limited to observation gathering and coverage.

## Evidence

- PRODUCT-REQ-SPEC-013 fixes the gathering criterion and JSONL schemas but delegates exact corpus scope and execution planning to this Work Item.
- PRODUCT-TASK-SPEC-026-01 selects a downstream decision-first route.
- The user requires separate sessions to author assigned JSONL directly.
- The user originally excluded helper tooling, then revised the boundary after concluding that constrained result writing and structural validation are necessary to prevent unusable JSONL evidence.
- Aggregation, semantic clustering, use-case extraction, and semantic-analysis tooling remain excluded.
- The user accepted the active five-family corpus plus targeted legacy ADR and Investigation provenance.
- The user accepted a lightweight `tools/term-inventory-mcp/` implementation with structural validation before extraction; PRODUCT-TASK-SPEC-027-02 is materialized as its implementation owner.
- Source file drift is accepted as unavoidable; manifests and hashes are best-effort snapshots rather than immutable authority.
- The user ran the exploratory 16-batch plan over 718 files and 6,534,001 bytes with zero unreadable UTF-8 files.
- The exploratory plan balanced every batch to approximately 408 KB, but each batch mixed many source families and carried approximately 102,000 estimated input tokens before instructions and generated observations.
- `tools/term-inventory-mcp/plan_batches.py` now provides a `coherent-32` profile that preserves PRODUCT Tasks, PRODUCT non-Tasks, DRMCP Tasks, DRMCP non-Tasks, TRV, BPDSL, skills, legacy ADRs, and legacy Investigations as separate bucket families.
- The user accepted PRODUCT-INV-SPEC-011 as the durable Investigation owner and `product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/` as the machine-readable Evidence root.
- Each extraction Task will own one batch and separate observations and coverage files; no merged master JSONL is required.
- The measured coherent-32 plan covers all 718 files, ranges from approximately 38,899 to 60,852 estimated input tokens per batch, and has zero unreadable UTF-8 files.
- D-001 through D-008 are decided.
- Post-materialization review found that the first T04 snapshot had admitted a partial set of this workflow's own generated records. D-001 was clarified to exclude T03–T38 and PRODUCT-INV-SPEC-011 while retaining T01, T02, and PRODUCT-WORK-SPEC-027.
- T03 and T04 are materialized from the accepted execution route.
- DRMCP is non-operational. Filesystem authoring is the required fallback.
