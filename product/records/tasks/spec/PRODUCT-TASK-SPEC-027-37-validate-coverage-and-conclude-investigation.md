# PRODUCT-TASK-SPEC-027-37: Validate term-inventory coverage and conclude Investigation

- **id**: PRODUCT-TASK-SPEC-027-37
- **status**: done
- **date**: 2026-07-04
- **work_item**: PRODUCT-WORK-SPEC-027
- **task_type**: investigation
- **estimate**: 0.5d
- **depends_on**:
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
- **outputs**:
  - PRODUCT-INV-SPEC-011

## Goal

Validate complete corpus coverage, record factual extraction results, and conclude PRODUCT-INV-SPEC-011 without semantic aggregation.

## Work

- Run `validate_inventory_batch` for all 32 accepted batch IDs.
- Review every structural diagnostic and preserve unresolved failures as facts.
- Reconcile the corrected 733-source manifest total with the combined batch coverage totals.
- Aggregate factual counts for observations, `no_candidate`, `failed`, `changed`, `missing`, and `unreadable`.
- Record factual results in PRODUCT-INV-SPEC-011 Findings and related conclusion sections.
- Change PRODUCT-INV-SPEC-011 status from `investigating` to `concluded` only after structural reconciliation completes.
- Record follow-up candidates only when directly supported by collected Evidence.

This Task must not:

- create a merged master JSONL;
- classify terms;
- normalize synonyms;
- merge meanings into semantic clusters;
- extract domain use cases;
- define formal vocabulary;
- change Requirement or Work Item lifecycle;
- stage or commit changes.

## Done condition

- All 32 batch results have been structurally validated.
- Manifest and coverage totals reconcile without hidden or duplicate sources.
- Factual result counts and unresolved diagnostics are recorded in PRODUCT-INV-SPEC-011.
- PRODUCT-INV-SPEC-011 is `concluded`.
- No merged JSONL or semantic classification was produced.

## Verification

- Confirm `validate_inventory_batch` ran for every accepted batch ID.
- Confirm combined coverage accounts for exactly 733 corrected-manifest sources.
- Confirm observation, no-candidate, failed, drift, missing, and unreadable totals reconcile with per-source coverage.
- Confirm Investigation findings contain facts rather than vocabulary decisions.
- Confirm no Requirement, Work Item, result JSONL, stage, or commit change occurred outside this Task boundary.

## Evidence

- PRODUCT-INV-SPEC-011 owns the bounded research conclusion.
- PRODUCT-TASK-SPEC-027-05 through PRODUCT-TASK-SPEC-027-36 own independent batch Evidence.
- DRMCP is non-operational. Filesystem authoring was used for Design Record reads and writes.
- The term-inventory MCP ran `validate_inventory_batch` for every accepted batch ID.
- All 32 batches returned `valid: true` with zero structural diagnostics.
- Every batch reported `disposed_source_count` equal to `batch_source_count`.

| batch | sources | observation records | structural diagnostics | changed snapshots |
|---|---:|---:|---:|---:|
| PT01 | 22 | 135 | 0 | 0 |
| PT02 | 24 | 119 | 0 | 0 |
| PT03 | 27 | 145 | 0 | 0 |
| PT04 | 27 | 259 | 0 | 0 |
| PT05 | 26 | 165 | 0 | 1 |
| PT06 | 27 | 167 | 0 | 0 |
| PT07 | 27 | 173 | 0 | 1 |
| PT08 | 26 | 171 | 0 | 0 |
| PN01 | 28 | 238 | 0 | 1 |
| PN02 | 29 | 262 | 0 | 0 |
| PN03 | 30 | 314 | 0 | 0 |
| PN04 | 30 | 351 | 0 | 0 |
| DT01 | 16 | 123 | 0 | 0 |
| DT02 | 16 | 139 | 0 | 1 |
| DT03 | 16 | 180 | 0 | 0 |
| DT04 | 16 | 194 | 0 | 0 |
| DT05 | 16 | 159 | 0 | 0 |
| DN01 | 24 | 151 | 0 | 1 |
| DN02 | 24 | 107 | 0 | 0 |
| DN03 | 24 | 177 | 0 | 2 |
| TR01 | 35 | 393 | 0 | 0 |
| TR02 | 34 | 198 | 0 | 0 |
| BP01 | 18 | 147 | 0 | 0 |
| BP02 | 19 | 132 | 0 | 0 |
| SK01 | 40 | 401 | 0 | 0 |
| LA01 | 20 | 95 | 0 | 0 |
| LA02 | 19 | 111 | 0 | 0 |
| LA03 | 20 | 151 | 0 | 0 |
| LA04 | 20 | 79 | 0 | 0 |
| LA05 | 20 | 73 | 0 | 0 |
| LI01 | 7 | 69 | 0 | 0 |
| LI02 | 6 | 121 | 0 | 0 |
| **Total** | **733** | **5,699** | **0** | **7** |

- The corrected manifest contains 733 sources.
- The 32 accepted batches assign and dispose all 733 sources exactly once.
- The first 736-source manifest was superseded by T04 before this Task executed.
- All 733 sources have an `observations` disposition.
- `no_candidate` dispositions: 0.
- `failed` dispositions: 0.
- `changed` snapshots: 7.
- `missing` snapshots: 0.
- `unreadable` snapshots: 0.
- The seven changed sources remained readable and produced observations.
- PRODUCT-INV-SPEC-011 now records the factual totals and changed-source paths.
- PRODUCT-INV-SPEC-011 status changed from `investigating` to `concluded`.
- No merged JSONL, semantic classification, normalization, clustering, or use-case extraction was produced.
- No Requirement, Work Item, observation JSONL, coverage JSON, stage, or commit change occurred.
