# DRMCP-WORK-MCP-007: Validation Work Item disposition and rebaseline

- **id**: DRMCP-WORK-MCP-007
- **status**: done
- **date**: 2026-06-26
- **source_requirement**: DRMCP-REQ-MCP-001
- **impact_refs**:
  - DRMCP-ADR-MCP-001
  - DRMCP-WORK-MCP-001
  - DRMCP-WORK-MCP-003
  - DRMCP-WORK-MCP-004
  - DRMCP-WORK-MCP-005
  - DRMCP-WORK-MCP-006
  - DRMCP-WORK-SPEC-001
  - DRMCP-WORK-SPEC-002
  - DRMCP-TASK-MCP-001-07
  - PRODUCT-WORK-SPEC-015
  - spec:product.design_records.spec_format.validation_policy
  - spec:product.design_records.spec_format.follow_up_boundary
- **tasks**:
  - DRMCP-TASK-MCP-007-01
  - DRMCP-TASK-MCP-007-02
  - DRMCP-TASK-MCP-007-03
  - DRMCP-TASK-MCP-007-04

## Goal

Resolve the lifecycle and ownership disposition of the existing DRMCP spec-validation Work Items.

Rebaseline retained or replacement work without duplicating the current read-contract and implementation Work Items.

## Boundary

This Work Item owns:

- disposition of `DRMCP-WORK-SPEC-001`;
- disposition of `DRMCP-WORK-SPEC-002`;
- one explicit `retain`, `supersede`, `absorb`, or `close` result for each Work Item;
- comparison against `DRMCP-WORK-MCP-003` through `DRMCP-WORK-MCP-006`;
- correction of stale workflow metadata and obsolete implementation-phase assumptions where required by the disposition;
- definition of the retained or replacement DRMCP validation boundary;
- source-Requirement and dependency correction for any retained or replacement Work Item;
- DRMCP-side record synchronization, scoped validation, and independent review.

This Work Item does not own:

- PRODUCT validation rules or severity policy;
- PRODUCT validation-policy or follow-up-boundary pointer edits;
- parser-aware spec validation or Topics graph implementation;
- current read-contract correction owned by `DRMCP-WORK-MCP-003` through `DRMCP-WORK-MCP-006`;
- current read implementation owned by the later implementation Work Item;
- fixture authoring;
- temporary PRODUCT validation tooling.

`PRODUCT-WORK-SPEC-015` owns PRODUCT-side pointer synchronization after this Work Item accepts the disposition.

## Impact Scope

| ref or area | impact |
|---|---|
| `DRMCP-REQ-MCP-001` | Source Requirement for the read-baseline reimplementation phase. |
| `DRMCP-ADR-MCP-001` | Requires coordinated disposition before replacing validation Work Items. |
| `DRMCP-WORK-SPEC-001` | Retained per-file validation implementation owner rebaselined by T03. |
| `DRMCP-WORK-SPEC-002` | Retained Topics graph validation implementation owner rebaselined by T03. |
| `DRMCP-WORK-MCP-003` through `DRMCP-WORK-MCP-006` | Define boundaries that must not be duplicated by retained validation work. |
| `PRODUCT-WORK-SPEC-015` | Consumes the accepted disposition and synchronizes PRODUCT owner pointers. |
| PRODUCT validation-policy specs | Remain semantic and ownership authorities on the PRODUCT side. |

## Task flow

| phase | dependency | outcome |
|---|---|---|
| A. Existing Work Item audit | Accepted W003-W006 boundaries | Confirm current scope, stale assumptions, metadata shape, and overlap for W-SPEC-001/002. |
| B. Disposition decision | Phase A | Select retain, supersede, absorb, or close for each candidate. |
| C. Rebaseline and DRMCP synchronization | Phase B | Define retained or replacement boundaries and correct DRMCP records. |
| D. PRODUCT handoff | Phase C | Provide an accepted disposition for `PRODUCT-WORK-SPEC-015`. |
| E. Validation and review | Phases C-D | Validate DRMCP changes, run independent review, apply corrections, and close. |

PRODUCT-side pointer edits occur only in `PRODUCT-WORK-SPEC-015`.

## Task Candidates

| candidate | scope | dependency |
|---|---|---|
| T01 | Audit W-SPEC-001/002 against current metadata rules and W003-W006 boundaries. | W003-W006 boundaries available. |
| T02 | Decide retain, supersede, absorb, or close for each existing Work Item. | T01. |
| T03 | Rebaseline retained or replacement DRMCP Work Items and synchronize DRMCP records. | T02. |
| T04 | Record the PRODUCT handoff, validate, run independent review, correct, and close. | T03. |

## Completion Condition

This Work Item is complete when all of the following are true:

- both existing validation Work Items have explicit accepted dispositions;
- no retained or replacement Work Item duplicates W003-W006 authority;
- retained or replacement Work Items use the correct source Requirement and dependencies;
- stale workflow metadata and obsolete phase assumptions are resolved or explicitly preserved with reason;
- parser-aware spec and Topics graph implementation remain outside current read-contract correction unless explicitly retained;
- `PRODUCT-WORK-SPEC-015` receives an exact accepted pointer target set;
- all DRMCP-side changed records pass scoped validation;
- independent review reports no blocking or major findings;
- `DRMCP-REQ-MCP-001` lists this Work Item in `work_items`;
- final evidence records both dispositions, changed records, validation results, review verdict, and residual limitations.

## Evidence

- `DRMCP-ADR-MCP-001`: Requires coordinated disposition and matching PRODUCT pointer updates.
- `DRMCP-WORK-SPEC-001`: Existing parser-aware spec-format validation candidate.
- `DRMCP-WORK-SPEC-002`: Existing Topics graph-validation candidate.
- `DRMCP-WORK-MCP-003` through `DRMCP-WORK-MCP-006`: Accepted non-overlap boundaries.
- `DRMCP-TASK-MCP-001-07`: Hub lifecycle gate for this cross-owner disposition.
- T02 disposition and review evidence is complete. T03 rebaseline and review are complete. T04 handoff, correction, limited re-review, and W007 closure are complete.
- `DRMCP-TASK-MCP-007-01` opened on 2026-06-28.
  - Exact Task inventory confirmed that no `DRMCP-TASK-MCP-007-*` record existed before creation.
  - T01 audits `DRMCP-WORK-SPEC-001` and `DRMCP-WORK-SPEC-002` without accepting `retain`, `supersede`, `absorb`, or `close`.
  - The audit records current metadata conformance, stale metadata, obsolete phase assumptions, W003-W006 overlap, residual implementation scope, PRODUCT-WORK-SPEC-015 handoff requirements, candidate manifests, scoped verification commands, and an independent review prompt.
  - Neither candidate Work Item nor any PRODUCT record changes during T01.
  - Hub `DRMCP-TASK-MCP-001-07` remains `not_started` because `PRODUCT-WORK-SPEC-015` has not begun and the hub requires both owner-side Work Items to begin.
  - Scoped whitespace verification before closure synchronization completed with `tracked_exit=0` and `untracked_exit=1`; no whitespace error or exit code `2` or greater was reported.
  - The untracked exit code `1` is the expected difference result against `NUL`.
  - Independent baseline review returned `PASS` with no blocking, major, or minor finding.
  - Review advisories were limited to independently unexecuted Git commands and the inability to reconstruct the pre-creation Task inventory from current filesystem state; neither advisory changes the audit baseline.
  - T02 decision-input readiness and T01 closure readiness were both assessed as `READY`.
  - T01 is `done`.
  - Final post-closure whitespace verification must run against the changed T01 and W007 bytes; its result is not written back into these files.
- `DRMCP-TASK-MCP-007-02` opened on 2026-06-28.
  - Exact Task inventory confirmed T01 as the only existing `DRMCP-TASK-MCP-007-*` record before T02 creation.
  - `DRMCP-WORK-SPEC-001` accepted disposition: `retain`.
  - `DRMCP-WORK-SPEC-002` accepted disposition: `retain`.
  - Both records require full T03 rebaseline but preserve their existing implementation-owner identities.
  - Per-file validation and Topics graph validation remain separate Work Items.
  - W-SPEC-002 depends on the accepted W-SPEC-001 detector boundary and must not duplicate per-file checks.
  - W008 owns shared fixture authoring and fixture-local structural checks.
  - W-SPEC-001 and W-SPEC-002 own their detector and graph automated implementation tests.
  - W009 retains general current-read implementation tests and does not absorb retained spec-validator implementation.
  - Both retained Work Items will use `DRMCP-REQ-MCP-001` as `source_requirement` after T03 rebaseline.
  - PRODUCT semantic rules remain authoritative through current PRODUCT spec refs.
  - The exact PRODUCT pointer targets remain `DRMCP-WORK-SPEC-001` for per-file rules and `DRMCP-WORK-SPEC-002` for Topics graph rules.
  - T02 changes only the T02 Task and this Work Item.
  - Candidate Work Items, PRODUCT records, DRMCP Requirements, W003-W006, W008-W009, hub T07, and owner pointers remain unchanged.
  - T03 is not created or started.
  - Independent T02 review verdict: `PASS`.
  - Blocking findings: none.
  - Major findings: none.
  - Minor findings: none.
  - The review accepted both `retain` dispositions, scope separation, dependency, fixture and test ownership, source-Requirement policy, PRODUCT handoff target, and T03 candidate manifest.
  - Review advisory A-01 records that repository-local Git commands were not independently executed by the reviewer.
  - Review advisory A-02 records that T03 absence was confirmed from workflow records rather than an additional directory listing.
  - External scoped whitespace verification reported W007 as tracked and modified and T02 as untracked.
  - External whitespace results were `tracked_exit=0` and `untracked_exit=1`.
  - No whitespace error or exit code `2` or greater was reported.
  - LF-to-CRLF working-copy warnings for both files were non-blocking.
  - T02 changed to `done` on 2026-06-28.
  - This closure synchronization changes the checked T02 and W007 bytes.
  - One post-closure scoped whitespace check remains external and must not be written back into these files.
  - At T02 closure, W007 remained `in_progress` and T03 had not started.
- `DRMCP-TASK-MCP-007-03` opened on 2026-06-28.
  - Exact pre-creation inventory found T01 and T02 only; no T03 record existed.
  - `DRMCP-WORK-SPEC-001` kept its ID and was fully rebaselined as the per-file validation implementation owner.
  - `DRMCP-WORK-SPEC-002` kept its ID and was fully rebaselined as the Topics graph validation implementation owner.
  - Both retained Work Items now use `DRMCP-REQ-MCP-001` as `source_requirement`.
  - W-SPEC-001 consumes W003 parsed state and canonical identity and W006 validation and diagnostic contracts.
  - W-SPEC-002 consumes the accepted W-SPEC-001 detector result boundary, W003 active-index state, and W006 validation and diagnostic contracts.
  - W-SPEC-002 does not duplicate H1, metadata, spec-kind, contract-class, required-section, or local Topics table-shape detectors.
  - The current Topics row is `title/kind/ref/summary`.
  - The declaring Index or Overview remains the authoritative parent.
  - Child H1-adjacent `parent` consistency, duplicate authoritative-parent detection, and cycle detection remain W-SPEC-002 scope.
  - Obsolete `file`, row-level `parent`, alias-waiting, temporary-tooling, and generic redesign assumptions were removed.
  - PRODUCT semantic authorities remain referenced through `impact_refs` and Boundary prose without copied ownership.
  - `DRMCP-REQ-MCP-001.work_items` now includes both retained Work Item IDs without removing or reordering existing entries.
  - W008 required no change because it already owns shared fixture files, manifests, and fixture-local checks while excluding production behavior.
  - W009 required no change because it already excludes retained validator implementation and owns only general current-read implementation and tests.
  - W-SPEC-001 owns per-file detector and runtime integration tests.
  - W-SPEC-002 owns graph algorithm and runtime integration tests.
  - The PRODUCT handoff target set remains W-SPEC-001 for per-file rows and W-SPEC-002 for Topics and graph rows.
  - PRODUCT records, W003-W006, W008, W009, hub T07, implementation source, fixtures, and T04 remain unchanged.
  - T03 changed-file manifest contains the new T03 Task, both retained Work Items, `DRMCP-REQ-MCP-001`, and this Work Item.
  - Task and Work Item records are outside the strict spec-format validator scope.
  - No exact applicable repository validator command was identified for the changed Requirement.
  - Repository-local commands were not executed by this assistant.
  - The user executed the targeted five-path status command.
  - Targeted status matched the changed-file manifest: four tracked modified files and one untracked T03 Task.
  - External whitespace results were `tracked_exit=0` and `untracked_exit=1`.
  - No whitespace error or exit code `2` or greater was reported.
  - LF-to-CRLF working-copy warnings for all five files were non-blocking.
  - The untracked exit code `1` is the expected difference result against `NUL`.
  - This Evidence synchronization changes the checked T03 Task and W007 bytes.
  - A final pre-review targeted whitespace check repeated `tracked_exit=0` and `untracked_exit=1` with no whitespace error or exit code `2` or greater.
  - Independent T03 review verdict: `PASS`.
  - Blocking findings: none.
  - Major findings: none.
  - Minor findings: none.
  - Review advisories were limited to independently unexecuted Git commands and no additional directory inventory for T04 absence.
  - Both advisories were non-blocking and did not change closure readiness.
  - T03 closure readiness was assessed as `READY`.
  - T04 handoff readiness was assessed as `READY AFTER T03 CLOSURE`.
  - T03 changed to `done` on 2026-06-28.
  - This closure synchronization changes the checked T03 Task and W007 bytes.
  - One post-closure targeted whitespace check remains external and must not be written back into these files.
  - W007 remains `in_progress`.
  - `PRODUCT-WORK-SPEC-015` remains `not_started`.
  - Hub `DRMCP-TASK-MCP-001-07` remains `not_started`.
  - T04 is not created or started.
  - Independent T03 review and finding closure are complete.
  - T04 may begin after the required post-closure targeted whitespace check.
- `DRMCP-TASK-MCP-007-04` opened on 2026-06-28.
  - Exact pre-creation inventory found T01, T02, and T03 only; no T04 record existed.
  - The accepted PRODUCT handoff target set is `DRMCP-WORK-SPEC-001` for per-file validation and `DRMCP-WORK-SPEC-002` for Topics and graph validation.
  - Both retained owner identities remain unchanged.
  - W-SPEC-002 retains its dependency on the accepted W-SPEC-001 detector result boundary.
  - PRODUCT remains the semantic rule and severity-policy authority.
  - PRODUCT follow-up-boundary already names W-SPEC-001 and W-SPEC-002 as the durable implementation owners.
  - PRODUCT validation-policy is not fully aligned at rule granularity.
  - Its local `## Topics` table column-shape row currently points to W-SPEC-002.
  - The accepted T03 boundary assigns that local detector to W-SPEC-001.
  - T04 does not edit the PRODUCT record because pointer synchronization belongs to `PRODUCT-WORK-SPEC-015`.
  - `PRODUCT-WORK-SPEC-015` must change the known local Topics column-shape owner from W-SPEC-002 to W-SPEC-001 while preserving PRODUCT rule text and severity.
  - `PRODUCT-WORK-SPEC-015` also owns repeated-pointer inventory, any additional synchronization, explicit no-change evidence, scoped validation, cross-owner review, and closure.
  - T04 does not assume that every PRODUCT pointer-only record has already been inventoried.
  - W008 remains the shared fixture owner.
  - W009 remains the general current-read implementation and test owner outside the retained validators.
  - W-SPEC-001 and W-SPEC-002 remain `not_started`.
  - `PRODUCT-WORK-SPEC-015` remains `not_started`.
  - Hub `DRMCP-TASK-MCP-001-07` remains `not_started`.
  - T04 changed-file manifest contains only the new T04 Task and this Work Item.
  - Repository-local commands were not executed by this assistant.
  - The user executed the targeted two-path status command.
  - Targeted status matched the changed-file manifest: W007 tracked and modified, and T04 untracked.
  - External whitespace results were `tracked_exit=0` and `untracked_exit=1`.
  - No whitespace error or exit code `2` or greater was reported.
  - LF-to-CRLF working-copy warnings for both files were non-blocking.
  - The untracked exit code `1` is the expected difference result against `NUL`.
  - Initial independent T04 review verdict: `NEEDS REVISION`.
  - Blocking findings: none.
  - Major finding F-MAJ-01 identified the omitted local Topics column-shape owner mismatch.
  - Minor findings: none.
  - F-MAJ-01 correction removes the full-alignment claim and records the exact PRODUCT synchronization target.
  - The correction preserves T04's no-PRODUCT-edit boundary and delegates the pointer update to `PRODUCT-WORK-SPEC-015`.
  - The correction changes the checked T04 Task and W007 bytes.
  - One targeted whitespace check remains external before independent limited re-review.
  - The limited re-review prompt is recorded in T04 Evidence.
  - The targeted whitespace result supplied to the limited reviewer was `tracked_exit=0` and `untracked_exit=1`, with no whitespace error or exit code `2` or greater.
  - Independent limited re-review verdict: `PASS`.
  - F-MAJ-01: `CLOSED`.
  - New blocking findings: none.
  - New major findings: none.
  - New minor findings: none.
  - PRODUCT pointer correction assessment: `PASS`.
  - PRODUCT-WORK-SPEC-015 handoff correction assessment: `PASS`.
  - Regression assessment: `PASS`.
  - T04 closure readiness: `READY`.
  - W007 closure readiness: `READY AFTER T04 CLOSURE SYNCHRONIZATION`.
  - PRODUCT-WORK-SPEC-015 start readiness: `READY AFTER W007 IS done`.
  - T04 changed to `done` on 2026-06-28.
  - W007 changed to `done` on 2026-06-28.
  - `PRODUCT-WORK-SPEC-015` remains `not_started` and may now begin.
  - Hub `DRMCP-TASK-MCP-001-07` remains `not_started`.
  - This closure synchronization changes the checked T04 Task and W007 bytes.
  - One post-closure targeted whitespace check remains external and must not be written back into these files.
