# PRODUCT-REQ-SPEC-004: Brewprint legacy spec compatibility correction

- **id**: PRODUCT-REQ-SPEC-004
- **status**: accepted
- **date**: 2026-06-26
- **source_refs**:
  - DRMCP-ADR-MCP-001
  - DRMCP-REQ-MCP-001
  - DRMCP-INV-MCP-002
  - spec:product.brewprint.compatibility.legacy_id_compatibility
  - spec:product.design_records.spec_format.spec_id_as_ref
- **work_items**:
  - PRODUCT-WORK-SPEC-014

## Requirement

Brewprint legacy compatibility must stop treating `V01-SPEC-*` as an accepted resolvable family.

The compatibility profile must preserve issued V01 sequential workflow artifact families while removing the compatibility-only legacy spec identity rule.

## Evidence

- `DRMCP-ADR-MCP-001` rejects `V01-SPEC-*` from the configured legacy fallback surface.
- `DRMCP-REQ-MCP-001` requires current specs to use H1-adjacent metadata and path-derived canonical `spec:` refs.
- `DRMCP-INV-MCP-002` found obsolete spec identity behavior in the active compatibility baseline.
- At Work Item start, `spec:product.brewprint.compatibility.legacy_id_compatibility` listed `V01-SPEC-*` as accepted and defined a compatibility-only spec identity note.

## Required Outcome

- Remove `V01-SPEC-*` from the accepted legacy families.
- Remove the compatibility-only legacy spec identity rule.
- Preserve `V01-ADR-*`, `V01-INV-*`, `V01-REQ-*`, `V01-WORK-*`, and `V01-TASK-*` unless a separate accepted decision changes them.
- Preserve issued legacy IDs without renaming or migration.
- Keep canonical current spec identity owned by `spec:product.design_records.spec_format.spec_id_as_ref`.
- Synchronize affected pointers and review the corrected compatibility boundary.

## Explicitly Excluded Scope

- DRMCP resolver or legacy-index implementation.
- Migration of V01 specs into the active spec tree.
- Renaming issued V01 artifacts.
- Changes to accepted V01 sequential workflow artifact families.
- General redesign of Brewprint archive policy.

## Boundary

PRODUCT owns the Brewprint compatibility profile and its accepted legacy families.

DRMCP owns runtime configuration, exact legacy grammar matching, fallback lookup, and diagnostics.
