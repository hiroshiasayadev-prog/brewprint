# DRMCP-INV-MCP-901: Current read fixture observations

- **status**: concluded
- **date**: 2026-06-28
- **trigger**: DRMCP-REQ-MCP-901
- **scope**: Observe the structure needed for accepted current read fixtures.
- **non_scope**: Runtime parser, index, resolver, and validator behavior.
- **source_refs**:
  - PRODUCT-ADR-SPEC-901
- **follow_up_candidates**: []

## Investigation scope

Inspect current fixture placement, identity, and relation inputs across two app roots.

## Out of scope

Runtime operation results and legacy archive behavior are excluded.

## Background

Current implementation tests need stable shared source material.

## What was investigated

The investigation covered app-aware IDs, path-derived spec refs, and exact cross-app relations.

## Findings

Separate PRODUCT and DRMCP roots can represent the accepted current-only arrangement without legacy input.

## Cross-cutting observations

Canonical refs remain independent of fixture file locations.

## Follow-up judgment candidates

None.

## Recommendation

The shared fixture arrangement appears suitable for later runtime tests.

## Follow-up artifact candidates

None.

## Open questions

None within the T02 fixture boundary.
