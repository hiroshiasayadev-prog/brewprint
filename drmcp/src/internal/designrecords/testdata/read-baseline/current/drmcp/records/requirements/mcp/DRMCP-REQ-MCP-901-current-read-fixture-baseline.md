# DRMCP-REQ-MCP-901: Current read fixture baseline

- **id**: DRMCP-REQ-MCP-901
- **status**: accepted
- **date**: 2026-06-28
- **source_refs**:
  - PRODUCT-ADR-SPEC-901
- **work_items**:
  - DRMCP-WORK-MCP-901

## Requirement

Shared current read fixtures must preserve canonical identity across separate configured app roots.

## Evidence

`PRODUCT-ADR-SPEC-901` provides the cross-app source relation for this fixture.

## Required Outcome

The fixture provides valid current records, exact refs, and no legacy archive source.

## Explicitly Excluded Scope

Runtime behavior and legacy fallback are excluded.

## Boundary

This fixture records source structure only.
