# PRODUCT-WORK-SPEC-015: Synchronize validation owner pointers

- **id**: PRODUCT-WORK-SPEC-015
- **status**: in_progress
- **date**: 2026-06-26
- **source_requirement**: PRODUCT-REQ-SPEC-001
- **impact_refs**:
  - PRODUCT-REQ-SPEC-001
  - PRODUCT-WORK-SPEC-001
  - DRMCP-ADR-MCP-001
  - DRMCP-WORK-MCP-007
  - DRMCP-WORK-SPEC-001
  - DRMCP-WORK-SPEC-002
  - DRMCP-TASK-MCP-001-07
  - spec:product.design_records.spec_format.validation_policy
  - spec:product.design_records.spec_format.follow_up_boundary
- **tasks**:
  - PRODUCT-TASK-SPEC-015-01

## Goal

Synchronize PRODUCT validation-owner pointers with the accepted DRMCP validation Work Item disposition.

Keep PRODUCT semantic rules authoritative while pointing implementation ownership to current DRMCP Work Items.

## Boundary

This Work Item owns:

- consumption of the accepted `DRMCP-WORK-MCP-007` disposition;
- owner-pointer updates in `spec:product.design_records.spec_format.validation_policy`;
- follow-up pointer updates in `spec:product.design_records.spec_format.follow_up_boundary`;
- removal of obsolete Work Item IDs as current implementation owners;
- synchronization of other PRODUCT records only when they repeat the same owner pointer;
- preservation of PRODUCT-owned validation rules and severity semantics;
- scoped validation and independent cross-owner review.

This Work Item does not own:

- disposition of `DRMCP-WORK-SPEC-001` or `DRMCP-WORK-SPEC-002`;
- DRMCP Work Item creation or lifecycle state;
- parser-aware validation or Topics graph implementation;
- validation rule or severity redesign;
- temporary PRODUCT validator implementation;
- DRMCP read-contract, diagnostic, fixture, or runtime implementation;
- broad PRODUCT spec cleanup unrelated to validation-owner pointers.

`DRMCP-WORK-MCP-007` decides the DRMCP target set.
This Work Item updates PRODUCT pointers without redefining that disposition.

## Impact Scope

| ref or area | impact |
|---|---|
| `PRODUCT-REQ-SPEC-001` | Source Requirement for MCP-readable spec validation and topic-tree support. |
| `DRMCP-ADR-MCP-001` | Requires matching PRODUCT pointer changes when validation Work Items are replaced or absorbed. |
| `DRMCP-WORK-MCP-007` | Supplies the accepted retain, supersede, absorb, or close disposition. |
| `DRMCP-WORK-SPEC-001` | Existing parser-aware validation pointer candidate. |
| `DRMCP-WORK-SPEC-002` | Existing Topics graph-validation pointer candidate. |
| `spec:product.design_records.spec_format.validation_policy` | Receives current implementation-owner pointers. |
| `spec:product.design_records.spec_format.follow_up_boundary` | Receives synchronized follow-up ownership text. |
| Other PRODUCT pointer-only records | Change only when they repeat an obsolete owner ID. |

## Task flow

| phase | dependency | outcome |
|---|---|---|
| A. Disposition intake | Accepted `DRMCP-WORK-MCP-007` result | Confirm exact current DRMCP owner targets and affected PRODUCT pointers. |
| B. Pointer synchronization | Phase A | Update validation-policy, follow-up-boundary, and exact repeated pointers. |
| C. Cross-owner consistency review | Phase B | Confirm PRODUCT semantics remain unchanged and DRMCP targets resolve. |
| D. Validation and closure | Phases B-C | Validate changed PRODUCT records, run independent review, correct, and close. |

This Work Item starts only after the DRMCP disposition is accepted.

## Task Candidates

| candidate | scope | dependency |
|---|---|---|
| T01 | Read the accepted DRMCP disposition and inventory affected PRODUCT owner pointers. | `DRMCP-WORK-MCP-007` done. |
| T02 | Synchronize validation-policy, follow-up-boundary, and exact repeated owner pointers. | T01. |
| T03 | Run scoped validation and cross-owner independent review, apply corrections, and close. | T02. |

## Completion Condition

This Work Item is complete when all of the following are true:

- every affected PRODUCT validation-owner pointer matches the accepted DRMCP disposition;
- obsolete Work Item IDs do not remain as current implementation owners;
- PRODUCT validation rules and severity semantics remain unchanged unless a separate decision authorizes change;
- PRODUCT records point to canonical DRMCP owners rather than lifecycle Tasks;
- no duplicated implementation authority remains across retained and replacement Work Items;
- all changed PRODUCT records pass scoped validation;
- cross-owner independent review reports no blocking or major findings;
- `PRODUCT-REQ-SPEC-001` lists this Work Item in `work_items`;
- final evidence records the accepted target set, changed files, validation results, review verdict, and residual limitations.

## Evidence

- `PRODUCT-REQ-SPEC-001`: Source Requirement.
- `DRMCP-ADR-MCP-001`: Requires coordinated pointer synchronization.
- `DRMCP-WORK-MCP-007`: DRMCP disposition owner and upstream gate.
- `DRMCP-TASK-MCP-001-07`: Hub lifecycle gate tracking both owner-side Work Items.
- `PRODUCT-TASK-SPEC-015-01` opened on 2026-06-28.
  - Exact Task inventory confirmed that no `PRODUCT-TASK-SPEC-015-*` record existed before creation.
  - Scoped PRODUCT pointer inventory identified two owner-only T02 changes in `spec:product.design_records.spec_format.validation_policy`.
  - The inventory also identified one W015 wording-only normalization from pointer-candidate wording to retained-owner wording.
  - The known local Topics column-shape row must move from W-SPEC-002 to W-SPEC-001.
  - The `parent grammar violation` row is an additional stale pointer and must move from W-SPEC-002 to W-SPEC-001.
  - PRODUCT rule text and severity remain unchanged.
  - `spec:product.design_records.spec_format.follow_up_boundary` requires no change from the T01 inventory.
  - Targeted status matched the three-file changed manifest.
  - External whitespace results were `tracked_exit=0` and `untracked_exit=1`.
  - No whitespace error or exit code `2` or greater was reported.
  - LF-to-CRLF warnings were non-blocking.
  - Final pre-review whitespace results were `tracked_exit=0` and `untracked_exit=1`.
  - Independent T01 review returned `PASS` with no blocking, major, or minor finding.
  - T01 is closed as `done`.
  - Final post-closure whitespace verification remains pending.
- PRODUCT pointer synchronization, validation, and independent-review closure evidence: pending later Tasks.
