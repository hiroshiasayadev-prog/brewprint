# PRODUCT-TASK-SPEC-014-01: Correct legacy compatibility authority and pointers

- **id**: PRODUCT-TASK-SPEC-014-01
- **status**: done
- **date**: 2026-06-27
- **work_item**: PRODUCT-WORK-SPEC-014
- **source_requirement**: PRODUCT-REQ-SPEC-004
- **estimate**: 0.5d
- **depends_on**: []
- **outputs**:
  - spec:product.brewprint.compatibility.legacy_id_compatibility
  - spec:product.brewprint.compatibility
  - PRODUCT-REQ-SPEC-004
  - PRODUCT-WORK-SPEC-014

## Goal

Remove obsolete `V01-SPEC-*` compatibility authority from the active Brewprint compatibility profile.

Preserve the accepted V01 sequential families and the path-derived current spec identity authority.

## Work

- Use `DRMCP-ADR-MCP-001` and `PRODUCT-REQ-SPEC-004` as the accepted family boundary.
- Remove `V01-SPEC-*` from the accepted-family table.
- Narrow retention language to IDs in the accepted legacy families.
- Remove the compatibility-only legacy spec identity section.
- Keep `spec:product.design_records.spec_format.spec_id_as_ref` as the current spec identity pointer.
- Remove compatibility-only spec identity claims from the parent compatibility overview.
- Correct stale current-state wording in `PRODUCT-REQ-SPEC-004` without changing its required outcome.
- Search active PRODUCT and DRMCP records for remaining claims that `V01-SPEC-*` is accepted.
- Route stale DRMCP contract claims to their owning Work Items.
- Preserve historical evidence and explicit rejection statements.

This Task does not change DRMCP runtime contracts, fixtures, or implementation.

## Done condition

- The accepted-family table contains only `V01-ADR-*`, `V01-INV-*`, `V01-REQ-*`, `V01-WORK-*`, and `V01-TASK-*`.
- Retention language does not imply that unaccepted V01 families remain resolvable.
- The compatibility-only legacy spec identity section is absent.
- The parent compatibility overview no longer assigns compatibility-only V01 spec identity authority.
- Current spec identity still points to `spec:product.design_records.spec_format.spec_id_as_ref`.
- No active PRODUCT compatibility contract claims that `V01-SPEC-*` is an accepted fallback family.
- Stale DRMCP contract claims are identified and routed to `DRMCP-WORK-MCP-003` or `DRMCP-WORK-MCP-006`.
- Historical records and rejection contracts remain intact.

## Verification

- Inspect the changed PRODUCT records.
- Run a targeted search for accepted-family and compatibility-only identity claims.
- Run strict validation for every changed PRODUCT spec record.
- Confirm the changed set contains no DRMCP runtime specification or implementation file.

## Evidence

- Accepted boundary: `DRMCP-ADR-MCP-001`.
- Source Requirement: `PRODUCT-REQ-SPEC-004`.
- Changed `spec:product.brewprint.compatibility.legacy_id_compatibility`:
  - removed the `V01-SPEC-*` accepted-family row;
  - narrowed retention to IDs in accepted legacy families;
  - removed the compatibility-only spec identity section;
  - retained the path-derived current spec identity pointer.
- Changed `spec:product.brewprint.compatibility`:
  - removed the compatibility-only V01 spec identity ownership row;
  - narrowed the legacy compatibility topic summary.
- Changed `PRODUCT-REQ-SPEC-004` to preserve the opening-state evidence without claiming the stale state remains current.
- Corrected `PRODUCT-WORK-SPEC-014` task relations and impact scope.
- Strict validation:
  - Command: `python -X utf8 product/src/tools/validate_spec.py product/records/spec/brewprint/compatibility/index.md product/records/spec/brewprint/compatibility/legacy-id-compatibility.md --strict --no-color`
  - Exit code: 0
  - Result: `[strict]  All 2 file(s) OK.`
- PRODUCT compatibility authority search:
  - Command: `rg -n "V01-SPEC-\*|compatibility-only V01 spec identity|compatibility-only spec identity" product/records/spec/brewprint/compatibility/index.md product/records/spec/brewprint/compatibility/legacy-id-compatibility.md`
  - Exit code: 1.
  - Result: 0 matching lines across the two active PRODUCT compatibility specs.
- Routed stale DRMCP contract search:
  - Command: `rg -n "V01-SPEC-\*" drmcp/records/spec/design-records-mcp/schema/fields.md drmcp/records/spec/design-records-mcp/schema/metadata-grammar.md drmcp/records/spec/design-records-mcp/tools/validate-records.md`
  - Exit code: 0.
  - Result: exactly 3 matching lines.
  - `schema/fields.md:122` routes to `DRMCP-WORK-MCP-003`.
  - `schema/metadata-grammar.md:86` routes to `DRMCP-WORK-MCP-003`.
  - `tools/validate-records.md:11` routes to `DRMCP-WORK-MCP-006`.
- The former broad aggregate classification was removed because searching Task and Work Item evidence made the count self-referential.
- These DRMCP claims do not remain PRODUCT compatibility authority and do not block T01 closure.
- DRMCP runtime/spec/source/fixture change check returned no changed DRMCP runtime files.
