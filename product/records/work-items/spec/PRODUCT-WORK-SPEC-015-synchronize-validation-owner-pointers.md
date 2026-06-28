# PRODUCT-WORK-SPEC-015: Synchronize validation owner pointers

- **id**: PRODUCT-WORK-SPEC-015
- **status**: done
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
  - PRODUCT-TASK-SPEC-015-02
  - PRODUCT-TASK-SPEC-015-03

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
| `DRMCP-WORK-SPEC-001` | Retained per-file parser-aware validation implementation owner. |
| `DRMCP-WORK-SPEC-002` | Retained cross-file Topics graph-validation implementation owner. |
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
- `PRODUCT-TASK-SPEC-015-02` opened on 2026-06-28.
  - The exact synchronization set contains two validation-policy owner cells.
  - The local Topics column-shape row changes from W-SPEC-002 to W-SPEC-001.
  - The `parent grammar violation` row changes from W-SPEC-002 to W-SPEC-001.
  - Rule text, severity, row order, and table formatting remain unchanged.
  - Unresolved child refs, duplicate authoritative parents, parent consistency, and cycles remain W-SPEC-002 concerns.
  - W-SPEC-001 and W-SPEC-002 Impact Scope descriptions now identify retained implementation ownership without changing either ID.
  - `spec:product.design_records.spec_format.follow_up_boundary` remains unchanged.
  - W-SPEC-001 and W-SPEC-002 remain `not_started`.
  - T02 does not close this Work Item or hub T07.
  - The changed-file manifest contains only the new T02 Task, validation-policy, and this Work Item.
  - The strict spec-format validator applies only to validation-policy among the three changed files.
  - Repository-local commands were not available through the filesystem MCP and were not executed by this assistant.
  - User-executed strict validation returned `[strict]  All 1 file(s) OK.` and `validator_exit=0`.
  - User-executed targeted status matched the three-file manifest: validation-policy and this Work Item were modified, and the new T02 Task was untracked.
  - User-executed whitespace checks returned `tracked_exit=0` and `untracked_exit=1`.
  - No whitespace error or exit code `2` or greater was reported.
  - LF-to-CRLF warnings for the three files were non-blocking.
  - Initial independent T02 review returned `NEEDS REVISION` with no blocking or major finding and one minor finding, F-MIN-01.
  - F-MIN-01 identified an impermissible copy of the Task's current status in this Work Item Evidence.
  - The correction replaced the copied status with pending-evidence wording.
  - User-executed post-F-MIN-01-correction whitespace checks returned `tracked_exit=0` and `untracked_exit=1`.
  - No whitespace error or exit code `2` or greater was reported.
  - The first limited re-review closed F-MIN-01 and returned `NEEDS REVISION` with no blocking or major finding and one minor finding, F-MIN-02.
  - F-MIN-02 identified that the post-correction whitespace result had not yet been recorded in T02 and this Work Item Evidence.
  - The F-MIN-02 correction recorded that result without changing the synchronization set or lifecycle boundary.
  - The final pre-review whitespace check returned `tracked_exit=0` and `untracked_exit=1`, with no whitespace error or exit code `2` or greater.
  - The final independent review returned `PASS` and closed F-MIN-01, F-MIN-02, and F-MIN-03.
  - No blocking, major, or minor finding remains.
  - The reviewer confirmed readiness for the next Task after closure.
  - The accepted final post-closure whitespace check returned `tracked_exit=0` and `untracked_exit=1`.
  - No whitespace error or exit code `2` or greater was reported.
  - The independent review and limited re-review prompts are recorded in T02 Evidence.
- `PRODUCT-TASK-SPEC-015-03` opened on 2026-06-28.
  - The exact Task directory inventory contained T01 and T02 but no T03 or equivalent validate-review-close Task.
  - The stale T02 post-closure pending note was synchronized to the accepted final whitespace result.
  - Static inspection confirms the two synchronized validation-policy owner rows match the accepted DRMCP disposition.
  - Static inspection confirms W-SPEC-001 retains per-file parser-aware validation ownership.
  - Static inspection confirms W-SPEC-002 retains cross-file Topics graph-validation ownership.
  - `spec:product.design_records.spec_format.follow_up_boundary` requires no edit.
  - `PRODUCT-REQ-SPEC-001` already lists this Work Item in `work_items`; no Requirement correction is required.
  - The initial T03 changed-file manifest contains only the new T03 Task and this Work Item.
  - The cumulative final-review target includes T01, T02, T03, this Work Item, validation-policy, and the explicit recheck-only records.
  - The strict validator applies to validation-policy; Task and Work Item records remain outside that validator scope.
  - Repository-local commands are unavailable through the filesystem MCP and were not executed by this assistant.
  - External targeted status matched the initial T03 changed-file manifest.
  - External strict validation returned `[strict]  All 1 file(s) OK.` and `validator_exit=0`.
  - Initial external whitespace checks returned `tracked_exit=0` and `untracked_exit=1`.
  - No whitespace error or exit code `2` or greater was reported.
  - LF-to-CRLF warnings were non-blocking.
  - This Evidence synchronization changed the T03 Task and this Work Item bytes after that check.
  - Final pre-review whitespace for the current bytes returned `tracked_exit=0` and `untracked_exit=1`.
  - Independent cross-owner final review returned `PASS`.
  - No blocking, major, or minor finding remains.
  - Every Completion Condition is satisfied within the recorded validation and review boundary.
  - Final closure-time changed files are T03, this Work Item, and hub T07.
  - Repository-local commands were not independently rerun by the reviewer.
  - Byte-level Git diff comparison against pre-T02 bytes was not independently performed.
  - Repository-wide clean status remains unknown.
  - This Work Item closes as `done`.
  - Hub T07 records the PRODUCT owner-side handoff and remains `in_progress`; PRODUCT does not close the hub.
