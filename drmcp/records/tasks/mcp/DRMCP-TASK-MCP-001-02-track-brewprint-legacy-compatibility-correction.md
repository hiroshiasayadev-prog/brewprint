# DRMCP-TASK-MCP-001-02: Track Brewprint legacy compatibility correction

- **id**: DRMCP-TASK-MCP-001-02
- **status**: done
- **date**: 2026-06-27
- **work_item**: DRMCP-WORK-MCP-001
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 0.5d coordination
- **depends_on**:
  - DRMCP-TASK-MCP-001-01
- **outputs**:
  - PRODUCT-WORK-SPEC-014
  - DRMCP-WORK-MCP-003

## Goal

Accept the Brewprint compatibility gate required before DRMCP legacy fallback implementation.

## Work

- Track `PRODUCT-WORK-SPEC-014` as the selected child Work Item.
- Confirm that `PRODUCT-WORK-SPEC-014` remains sourced from `PRODUCT-REQ-SPEC-004`.
- Delegate removal of `V01-SPEC-*` compatibility authority to `PRODUCT-WORK-SPEC-014`.
- Track the child Work Item through implementation, validation, review, and `done`.
- Record the child Work Item ID and accepted completion evidence here.

This Task does not edit the Brewprint compatibility spec.
All specification changes and corrective work belong to the selected child Work Item.

## Done condition

- `PRODUCT-WORK-SPEC-014` is `done`.
- `spec:product.brewprint.compatibility.legacy_id_compatibility` no longer accepts `V01-SPEC-*`.
- Compatibility-only legacy spec identity rules are removed.
- The child review has no blocking or major findings.
- The `PRODUCT-WORK-SPEC-014` evidence pointer is recorded here.

## Verification

- Read the final compatibility spec and child Work Item evidence.
- Confirm that accepted V01 sequential families remain unchanged unless separately decided.
- Confirm that this Task contains no direct compatibility implementation evidence.

## Evidence

- Selected child Work Item: `PRODUCT-WORK-SPEC-014`.
- Child source Requirement: `PRODUCT-REQ-SPEC-004`.
- Ownership: PRODUCT owns `spec:product.brewprint.compatibility.legacy_id_compatibility`; DRMCP consumes the corrected policy.
- Child completion: `PRODUCT-WORK-SPEC-014` is `done`.
- PRODUCT compatibility result:
  - accepted legacy families remain `V01-ADR-*`, `V01-INV-*`, `V01-REQ-*`, `V01-WORK-*`, and `V01-TASK-*`;
  - `V01-SPEC-*` is absent from the active PRODUCT compatibility authority;
  - the compatibility-only legacy spec identity rule is removed;
  - current spec identity remains owned by `spec:product.design_records.spec_format.spec_id_as_ref`.
- Validation: changed PRODUCT compatibility specs passed strict validation with `[strict]  All 2 file(s) OK.`
- Review: initial major finding on non-reproducible search counts was corrected; re-review verdict `PASS` with no remaining findings.
- Routed DRMCP follow-up:
  - `schema.fields` and `schema.metadata_grammar` remain owned by `DRMCP-WORK-MCP-003`;
  - `tools.validate_records` remains owned by `DRMCP-WORK-MCP-006`.
- `DRMCP-WORK-MCP-003` compatibility blocker is resolved.
