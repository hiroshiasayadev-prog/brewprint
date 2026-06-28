# PRODUCT-WORK-SPEC-014: Remove V01 spec compatibility authority

- **id**: PRODUCT-WORK-SPEC-014
- **status**: done
- **date**: 2026-06-27
- **source_requirement**: PRODUCT-REQ-SPEC-004
- **impact_refs**:
  - DRMCP-ADR-MCP-001
  - DRMCP-REQ-MCP-001
  - DRMCP-WORK-MCP-001
  - DRMCP-WORK-MCP-003
  - DRMCP-WORK-MCP-006
  - DRMCP-TASK-MCP-001-02
  - spec:product.brewprint.compatibility
  - spec:product.brewprint.compatibility.legacy_id_compatibility
  - spec:product.design_records.spec_format.spec_id_as_ref
- **tasks**:
  - PRODUCT-TASK-SPEC-014-01
  - PRODUCT-TASK-SPEC-014-02

## Goal

Correct the Brewprint legacy compatibility profile so only accepted V01 sequential workflow artifact families remain resolvable.

Remove `V01-SPEC-*` authority before DRMCP legacy fallback implementation begins.

## Boundary

This Work Item owns:

- removal of `V01-SPEC-*` from the accepted legacy families;
- removal of the compatibility-only legacy spec identity note;
- confirmation that current spec identity remains path-derived;
- synchronization of directly affected PRODUCT pointers and cross-owner DRMCP handoff pointers;
- validation and independent boundary review.

This Work Item does not own:

- normative DRMCP contract correction, including field, metadata-grammar, validation, resolver, index, configuration, or diagnostic contracts;
- DRMCP fixture or implementation changes;
- migration or rewriting of V01 spec files;
- renaming issued V01 artifacts;
- changes to accepted V01 sequential workflow artifact families;
- broad archive-policy redesign;
- lifecycle tracking for `DRMCP-WORK-MCP-001`.

## Impact Scope

| ref or area | impact |
|---|---|
| `PRODUCT-REQ-SPEC-004` | Source Requirement resolved by this Work Item. |
| `spec:product.brewprint.compatibility` | Remove the parent overview pointer to compatibility-only V01 spec identity. |
| `spec:product.brewprint.compatibility.legacy_id_compatibility` | Remove `V01-SPEC-*` and the compatibility-only spec identity rule. |
| `spec:product.design_records.spec_format.spec_id_as_ref` | Remains the canonical current spec identity authority. |
| `DRMCP-ADR-MCP-001` | Supplies the accepted legacy-family boundary. |
| `DRMCP-REQ-MCP-001` | Consumes the corrected compatibility policy. |
| `DRMCP-TASK-MCP-001-02` | Tracks this Work Item as the cross-owner completion gate. |
| `DRMCP-WORK-MCP-003` | Consume the corrected PRODUCT authority and correct stale field and metadata-grammar contracts. |
| `DRMCP-WORK-MCP-006` | Correct stale `validate_records` contract claims after upstream DRMCP boundaries are available. |

## Task flow

| phase | dependency | outcome |
|---|---|---|
| A. Compatibility correction | `PRODUCT-REQ-SPEC-004` | Remove obsolete `V01-SPEC-*` authority while preserving accepted sequential families. |
| B. Pointer synchronization and routing | Phase A | Update PRODUCT pointers and route stale DRMCP contract claims to their owning Work Items. |
| C. Validation and review | Phases A-B | Validate the corrected boundary and close review findings. |

## Task Candidates

| candidate | scope | dependency |
|---|---|---|
| T01 | Update the Brewprint legacy compatibility authority, synchronize PRODUCT pointers, and route stale DRMCP claims. | None. |
| T02 | Validate the corrected compatibility boundary, run independent review, apply required corrections, and close the Work Item. | T01. |

Detailed DRMCP runtime work remains delegated to DRMCP Work Items sourced from `DRMCP-REQ-MCP-001`.

## Completion Condition

- `V01-SPEC-*` is absent from the accepted legacy families.
- The compatibility-only legacy spec identity note is removed.
- Accepted V01 sequential workflow artifact families remain unchanged.
- Current spec identity still points to `spec:product.design_records.spec_format.spec_id_as_ref`.
- No active PRODUCT compatibility contract claims that `V01-SPEC-*` is an accepted fallback family.
- Remaining stale DRMCP contract claims are identified and routed to `DRMCP-WORK-MCP-003` or `DRMCP-WORK-MCP-006` without treating them as PRODUCT authority.
- Changed PRODUCT specs pass strict validation.
- Changed Requirement, Work Item, and Task records have valid metadata and reciprocal relations.
- Independent review reports no blocking or major findings.
- `PRODUCT-REQ-SPEC-004` lists this Work Item in `work_items`.
- Final evidence is available to `DRMCP-TASK-MCP-001-02`.

## Evidence

- `DRMCP-ADR-MCP-001`: Accepted decision excluding `V01-SPEC-*` from legacy fallback.
- `DRMCP-REQ-MCP-001`: Consumer requirement for the corrected compatibility boundary.
- `PRODUCT-REQ-SPEC-004`: PRODUCT-owned source Requirement.
- `DRMCP-TASK-MCP-001-02`: Hub lifecycle gate.
- T01 correction, strict PRODUCT spec validation, targeted-search classification, and DRMCP owner routing are recorded in `PRODUCT-TASK-SPEC-014-01`.
- T02 independent review initially returned `NEEDS REVISION` for non-reproducible aggregate search counts.
- The aggregate was replaced with fixed-file searches showing 0 active PRODUCT compatibility matches and exactly 3 routed stale DRMCP matches.
- Re-review verdict: `PASS`; previous major finding closed; no remaining findings.
- `PRODUCT-TASK-SPEC-014-01` and `PRODUCT-TASK-SPEC-014-02` are `done`.
- `DRMCP-TASK-MCP-001-02` accepted this completion evidence and is `done`.
- `DRMCP-WORK-MCP-003` no longer reports this compatibility correction as an active blocker.
