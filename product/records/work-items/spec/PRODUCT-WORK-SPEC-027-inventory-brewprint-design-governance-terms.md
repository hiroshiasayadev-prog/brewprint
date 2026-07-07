# PRODUCT-WORK-SPEC-027: Inventory Brewprint design-governance terms

- **id**: PRODUCT-WORK-SPEC-027
- **status**: done
- **date**: 2026-07-04
- **source_refs**:
  - PRODUCT-REQ-SPEC-013
  - PRODUCT-TASK-SPEC-026-03
- **impact_refs**: []
- **tasks**:
  - PRODUCT-TASK-SPEC-027-01
  - PRODUCT-TASK-SPEC-027-02
  - PRODUCT-TASK-SPEC-027-03
  - PRODUCT-TASK-SPEC-027-04
  - PRODUCT-TASK-SPEC-027-05
  - PRODUCT-TASK-SPEC-027-06
  - PRODUCT-TASK-SPEC-027-07
  - PRODUCT-TASK-SPEC-027-08
  - PRODUCT-TASK-SPEC-027-09
  - PRODUCT-TASK-SPEC-027-10
  - PRODUCT-TASK-SPEC-027-11
  - PRODUCT-TASK-SPEC-027-12
  - PRODUCT-TASK-SPEC-027-13
  - PRODUCT-TASK-SPEC-027-14
  - PRODUCT-TASK-SPEC-027-15
  - PRODUCT-TASK-SPEC-027-16
  - PRODUCT-TASK-SPEC-027-17
  - PRODUCT-TASK-SPEC-027-18
  - PRODUCT-TASK-SPEC-027-19
  - PRODUCT-TASK-SPEC-027-20
  - PRODUCT-TASK-SPEC-027-21
  - PRODUCT-TASK-SPEC-027-22
  - PRODUCT-TASK-SPEC-027-23
  - PRODUCT-TASK-SPEC-027-24
  - PRODUCT-TASK-SPEC-027-25
  - PRODUCT-TASK-SPEC-027-26
  - PRODUCT-TASK-SPEC-027-27
  - PRODUCT-TASK-SPEC-027-28
  - PRODUCT-TASK-SPEC-027-29
  - PRODUCT-TASK-SPEC-027-30
  - PRODUCT-TASK-SPEC-027-31
  - PRODUCT-TASK-SPEC-027-32
  - PRODUCT-TASK-SPEC-027-33
  - PRODUCT-TASK-SPEC-027-34
  - PRODUCT-TASK-SPEC-027-35
  - PRODUCT-TASK-SPEC-027-36
  - PRODUCT-TASK-SPEC-027-37
  - PRODUCT-TASK-SPEC-027-38

## Goal

Produce one coverage-accountable JSONL corpus of observed Brewprint design-governance term usage across the corpus scope selected by this Work Item's decision loop.

## Boundary

This Work Item owns:

- selection of the exact active corpus scope, including exclusion of this workflow's generated T03–T38 records and PRODUCT-INV-SPEC-011 itself;
- placement and ownership of raw JSONL observations and batch coverage metadata;
- partitioning of the selected corpus into non-overlapping extraction batches;
- one lightweight inventory-support MCP under `tools/term-inventory-mcp/` for manifest, batch, constrained result recording, and structural validation;
- parallel source scanning and direct JSONL observation authoring through the accepted interface;
- per-file coverage accountability, including explicit no-candidate results;
- the Task graph required to complete the inventory.

This Work Item does not own:

- aggregation or clustering of observations;
- use-case extraction from the resulting corpus;
- semantic-analysis, aggregation, clustering, or use-case extraction tooling;
- term classification, definition, normalization, approval, prohibition, replacement, or consolidation;
- Specification, skill, authoring-guide, or validator projection;
- bulk rewriting of source artifacts;
- Task-type use-case vocabulary.

## Impact Scope

| target | impact |
|---|---|
| selected active Brewprint artifact corpus | Read-only source scan under the gathering criterion in PRODUCT-REQ-SPEC-013. |
| PRODUCT-INV-SPEC-011 | Durable Investigation record for scope, method, concluded findings, and follow-up candidates. |
| `product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/manifest.json` | Best-effort source manifest for the selected corpus snapshot. |
| `product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/batches/<batch_id>.json` | Exact non-overlapping source assignment for one extraction batch. |
| `product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/<batch_id>.observations.jsonl` | Batch-owned observations using `bp-wide-term-observation-v1`. |
| `product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/<batch_id>.coverage.json` | Batch-owned coverage metadata using `bp-wide-term-batch-v1`. |

## Task flow

```text
PRODUCT-TASK-SPEC-027-01 decide execution
  -> PRODUCT-TASK-SPEC-027-02 implement inventory MCP
  -> PRODUCT-TASK-SPEC-027-03 author Investigation shell

PRODUCT-TASK-SPEC-027-02 + PRODUCT-TASK-SPEC-027-03
  -> PRODUCT-TASK-SPEC-027-04 persist manifest, batches, and graph
     -> PRODUCT-TASK-SPEC-027-05 .. PRODUCT-TASK-SPEC-027-36
        parallel, writer-disjoint extraction
        -> PRODUCT-TASK-SPEC-027-37 validate coverage and conclude Investigation
           -> PRODUCT-TASK-SPEC-027-38 synchronize closure
```

T02 and T03 may run in parallel after T01.
T05 through T36 may run in parallel after T04 because every Task owns a different batch and result pair.
T37 depends on all and only T05 through T36.
T38 depends only on T37.

### Batch and Task mapping

| Task | batch | bucket |
|---|---|---|
| PRODUCT-TASK-SPEC-027-05 | PT01 | product_tasks |
| PRODUCT-TASK-SPEC-027-06 | PT02 | product_tasks |
| PRODUCT-TASK-SPEC-027-07 | PT03 | product_tasks |
| PRODUCT-TASK-SPEC-027-08 | PT04 | product_tasks |
| PRODUCT-TASK-SPEC-027-09 | PT05 | product_tasks |
| PRODUCT-TASK-SPEC-027-10 | PT06 | product_tasks |
| PRODUCT-TASK-SPEC-027-11 | PT07 | product_tasks |
| PRODUCT-TASK-SPEC-027-12 | PT08 | product_tasks |
| PRODUCT-TASK-SPEC-027-13 | PN01 | product_non_tasks |
| PRODUCT-TASK-SPEC-027-14 | PN02 | product_non_tasks |
| PRODUCT-TASK-SPEC-027-15 | PN03 | product_non_tasks |
| PRODUCT-TASK-SPEC-027-16 | PN04 | product_non_tasks |
| PRODUCT-TASK-SPEC-027-17 | DT01 | drmcp_tasks |
| PRODUCT-TASK-SPEC-027-18 | DT02 | drmcp_tasks |
| PRODUCT-TASK-SPEC-027-19 | DT03 | drmcp_tasks |
| PRODUCT-TASK-SPEC-027-20 | DT04 | drmcp_tasks |
| PRODUCT-TASK-SPEC-027-21 | DT05 | drmcp_tasks |
| PRODUCT-TASK-SPEC-027-22 | DN01 | drmcp_non_tasks |
| PRODUCT-TASK-SPEC-027-23 | DN02 | drmcp_non_tasks |
| PRODUCT-TASK-SPEC-027-24 | DN03 | drmcp_non_tasks |
| PRODUCT-TASK-SPEC-027-25 | TR01 | trv_all |
| PRODUCT-TASK-SPEC-027-26 | TR02 | trv_all |
| PRODUCT-TASK-SPEC-027-27 | BP01 | bpdsl_all |
| PRODUCT-TASK-SPEC-027-28 | BP02 | bpdsl_all |
| PRODUCT-TASK-SPEC-027-29 | SK01 | skills_all |
| PRODUCT-TASK-SPEC-027-30 | LA01 | legacy_adrs |
| PRODUCT-TASK-SPEC-027-31 | LA02 | legacy_adrs |
| PRODUCT-TASK-SPEC-027-32 | LA03 | legacy_adrs |
| PRODUCT-TASK-SPEC-027-33 | LA04 | legacy_adrs |
| PRODUCT-TASK-SPEC-027-34 | LA05 | legacy_adrs |
| PRODUCT-TASK-SPEC-027-35 | LI01 | legacy_investigations |
| PRODUCT-TASK-SPEC-027-36 | LI02 | legacy_investigations |

## Task Candidates

| task | task type | responsibility | dependency |
|---|---|---|---|
| PRODUCT-TASK-SPEC-027-01 | decision | Fix corpus scope, output representation and placement, batch partitioning, extraction responsibility, coverage accounting, and the exact extraction Task graph. | none |
| PRODUCT-TASK-SPEC-027-02 | implementation | Implement the lightweight manifest, batching, constrained record-writing, and structural-validation MCP required before extraction. | T01 |
| PRODUCT-TASK-SPEC-027-03 | authoring | Create PRODUCT-INV-SPEC-011 as the investigating research shell and record the machine-readable Evidence contract. | T01 |
| PRODUCT-TASK-SPEC-027-04 | coordination | Persist the current manifest and coherent-32 assignments, then materialize T05–T38 with exact batch ownership and dependencies. | T02, T03 |
| PRODUCT-TASK-SPEC-027-05 .. 027-36 | investigation | Extract one assigned batch into one observations JSONL and one coverage JSON. | T04; parallel and writer-disjoint |
| PRODUCT-TASK-SPEC-027-37 | investigation | Structurally validate all batches, record factual totals, and conclude PRODUCT-INV-SPEC-011 without semantic aggregation. | exactly T05–T36 |
| PRODUCT-TASK-SPEC-027-38 | synchronization | Synchronize Requirement and Work Item lifecycle and closure Evidence after Investigation conclusion. | exactly T37 |

## Completion Condition

- Every source file in the selected corpus scope is accounted for by conforming observations or an explicit no-candidate result.
- Every observation conforms to `bp-wide-term-observation-v1`.
- Every extraction batch records `bp-wide-term-batch-v1` coverage metadata.
- Failed or unreadable files are recorded explicitly and do not disappear from coverage accounting.
- Observations remain unclassified and unnormalized source Evidence.
- The inventory-support MCP can create manifests, partitions, batch views, constrained observation records, and structural validation results.
- Aggregation, semantic clustering, use-case extraction, vocabulary definition, and vocabulary projection remain outside this Work Item.

## Evidence

- PRODUCT-REQ-SPEC-013 defines the gathering criterion, JSONL observation schema, coverage requirement, and exclusions.
- PRODUCT-TASK-SPEC-026-01 fixed the downstream Goal, Boundary, Completion Condition, direct source, unknown handling, and initial decision route.
- PRODUCT-TASK-SPEC-026-02 separated later aggregation from the inventory outcome; this Work Item later admitted only the constrained collection interface needed for valid JSONL evidence.
- PRODUCT-TASK-SPEC-026-03 created this Work Item from the accepted framing contract.
- The selected corpus includes active PRODUCT, DRMCP, BPDSL, TRV, and skill Markdown plus targeted legacy ADR and Investigation provenance. All other legacy and implementation sources are excluded, as are this workflow's generated T03–T38 records and PRODUCT-INV-SPEC-011 itself. T01, T02, and this Work Item remain in scope as pre-execution authority.
- The user accepted a lightweight collection MCP and structural validator because malformed JSONL would invalidate the investigation evidence.
- PRODUCT-TASK-SPEC-027-02 owns that bounded implementation; semantic aggregation tooling remains excluded.
- `tools/term-inventory-mcp/plan_batches.py` measured the accepted coherent-32 profile over 718 files and 6,539,750 bytes with zero unreadable UTF-8 files.
- PRODUCT-INV-SPEC-011 is reserved as the durable research artifact. Its batch Evidence root is `product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/`.
- Each extraction Task will own exactly one batch JSONL and one batch coverage file; this Work Item will not create a merged master JSONL.
- PRODUCT-TASK-SPEC-027-01 completed D-001 through D-008 and materialized T03 and T04.
- T04 materialized the 32 extraction Tasks, final validation and conclusion Task, and closure synchronization Task.
- DRMCP is non-operational. Filesystem authoring was used for Design Record reads and writes.
- The first T04 manifest at `2026-07-04T02:14:33+09:00` contained 736 files and 6,600,094 bytes but was superseded after self-generated workflow records were found in scope.
- The corrected manifest was persisted at `2026-07-04T03:14:40+09:00` and records 733 files, 6,642,319 bytes, and 0 unreadable UTF-8 files.
- The earlier planner values remain comparison Evidence only.
- The corrected coherent-32 partition contains exactly 32 accepted batch IDs and assigns all 733 manifest sources.
- MCP partition checks found no duplicate, missing, extra, or invalid-batch assignment.
- Every extraction Task owns a unique observations JSONL and coverage JSON path.
- T05 through T36 are parallel and writer-disjoint.
- T37 depends on all and only T05 through T36. T38 depends only on T37.
- The correction re-partitioned the canonical assignments without modifying any batch-owned result file. T25–T36 remained executable while PRODUCT and DRMCP assignments were corrected.
- Live validation found no coverage-scope mismatch in the sampled concurrent BPDSL, skill, TRV, legacy ADR, and legacy Investigation results.
- T04 and its correction created no semantic observations, merged JSONL, aggregation, clustering, use-case extraction, or Investigation conclusion.
- Stage and commit were not performed.

### Closure synchronization

- PRODUCT-TASK-SPEC-027-37 is `done`.
- PRODUCT-INV-SPEC-011 is `concluded`.
- PRODUCT-REQ-SPEC-013 remains `accepted`.
- This Work Item owns the completion judgment for PRODUCT-REQ-SPEC-013.

| Completion Condition | result | Evidence |
|---|---|---|
| Every selected source is accounted for. | PASS | 733 of 733 corrected-manifest sources have a disposition. |
| Every observation conforms to `bp-wide-term-observation-v1`. | PASS | All 32 batches passed structural validation. |
| Every batch records `bp-wide-term-batch-v1` coverage metadata. | PASS | All 32 coverage files passed structural validation. |
| Failed or unreadable files remain explicit. | PASS | `failed`, `missing`, and `unreadable` totals are all 0. |
| Observations remain unclassified and unnormalized. | PASS | T37 and PRODUCT-INV-SPEC-011 record no semantic classification or normalization. |
| The inventory-support MCP provides the accepted five operations. | PASS | PRODUCT-TASK-SPEC-027-02 is `done`; T04 and T37 used the implemented interface. |
| Semantic aggregation and vocabulary projection remain excluded. | PASS | No merged JSONL, clustering, use-case extraction, definition, or projection was produced. |

- The Work Item lists exactly PRODUCT-TASK-SPEC-027-01 through PRODUCT-TASK-SPEC-027-38.
- All 38 Task records point to PRODUCT-WORK-SPEC-027 through `work_item`.
- PRODUCT-TASK-SPEC-027-38 depends only on PRODUCT-TASK-SPEC-027-37.
- The Work Item lifecycle changed from `in_progress` to `done`.
- No Task graph, observation JSONL, coverage JSON, Investigation finding, or canonical design content changed during closure.
