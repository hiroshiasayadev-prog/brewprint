# DRMCP-WORK-MCP-007: Validation Work Item disposition and rebaseline

- **id**: DRMCP-WORK-MCP-007
- **status**: in_progress
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
| `DRMCP-WORK-SPEC-001` | Candidate for retain, supersede, absorb, or close. |
| `DRMCP-WORK-SPEC-002` | Candidate for retain, supersede, absorb, or close. |
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
- Disposition, synchronization, validation, and independent-review evidence: pending Task execution.
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
  - Final disposition remains intentionally open for T02.
