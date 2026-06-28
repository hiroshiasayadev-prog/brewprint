# PRODUCT-TASK-SPEC-014-02: Validate, review, and close legacy compatibility correction

- **id**: PRODUCT-TASK-SPEC-014-02
- **status**: done
- **date**: 2026-06-27
- **work_item**: PRODUCT-WORK-SPEC-014
- **source_requirement**: PRODUCT-REQ-SPEC-004
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-014-01
- **outputs**:
  - PRODUCT-WORK-SPEC-014
  - DRMCP-TASK-MCP-001-02
  - DRMCP-WORK-MCP-003
  - DRMCP-WORK-MCP-006

## Goal

Verify the corrected Brewprint compatibility boundary and close the cross-owner gate.

## Work

- Review the T01 diff against `PRODUCT-REQ-SPEC-004` and `DRMCP-ADR-MCP-001`.
- Run strict validation for changed PRODUCT specs.
- Inspect changed Requirement, Work Item, and Task metadata and reciprocal relations.
- Run targeted searches for accepted `V01-SPEC-*` and compatibility-only spec identity claims.
- Confirm that stale DRMCP contract claims are routed to `DRMCP-WORK-MCP-003` or `DRMCP-WORK-MCP-006` rather than corrected inside this PRODUCT Work Item.
- Obtain an independent review of the PRODUCT-owned correction boundary.
- Apply required corrections from blocking or major findings.
- Record final changed files, validation results, review verdict, and residual limitations.
- Mark `PRODUCT-WORK-SPEC-014` done after its completion conditions are satisfied.
- Update `DRMCP-TASK-MCP-001-02` with accepted child completion evidence and close the coordination Task.
- Remove the resolved compatibility blocker from current `DRMCP-WORK-MCP-003` evidence.

This Task does not perform DRMCP contract correction or runtime implementation.

## Done condition

- T01 is done with complete evidence.
- Every changed PRODUCT spec passes strict validation.
- Changed Requirement, Work Item, and Task records have valid metadata and reciprocal relations.
- Targeted searches find no active PRODUCT compatibility contract that accepts `V01-SPEC-*`.
- Remaining stale DRMCP contract claims are classified and routed to their owning Work Items.
- Independent review reports no blocking or major findings.
- `PRODUCT-WORK-SPEC-014` records final evidence and is done.
- `DRMCP-TASK-MCP-001-02` records the accepted child evidence and is done.
- `DRMCP-WORK-MCP-003` no longer reports the corrected compatibility profile as an active blocker.

## Verification

- Re-run the accepted validation and targeted-search commands after corrections.
- Inspect final statuses, reciprocal relations, and evidence pointers.
- Confirm that accepted V01 sequential families remain unchanged.
- Confirm that no normative DRMCP contract, fixture, or implementation change entered this Work Item.

## Evidence

- T01 correction and targeted-search evidence are recorded in `PRODUCT-TASK-SPEC-014-01`.
- PRODUCT spec validation baseline: `[strict]  All 2 file(s) OK.`
- Routed stale DRMCP contracts: `schema.fields` and `schema.metadata_grammar` to `DRMCP-WORK-MCP-003`; `tools.validate_records` to `DRMCP-WORK-MCP-006`.
- Independent review verdict: `NEEDS REVISION` with one major finding on non-reproducible T01 search classification.
- Correction applied: replaced the self-referential broad aggregate with two fixed-file searches recording 0 PRODUCT authority matches and exactly 3 routed DRMCP matches.
- Re-review verdict: `PASS`.
- Previous major finding: closed.
- Remaining findings: none.
- T01 closure assessment: justified.
- T02 closure readiness: ready for final synchronization.
- Final synchronization completed:
  - `PRODUCT-WORK-SPEC-014` marked `done` with final evidence;
  - `DRMCP-TASK-MCP-001-02` marked `done` with accepted child evidence;
  - `DRMCP-WORK-MCP-003` compatibility blocker removed and status advanced to `in_progress`.
