# PRODUCT-WORK-SPEC-006: Temporary spec format validator tooling

- **id**: PRODUCT-WORK-SPEC-006
- **status**: done
- **date**: 2026-06-16
- **requirement_refs**:
  - PRODUCT-REQ-SPEC-001
- **source_work_items**:
  - PRODUCT-WORK-SPEC-001
  - PRODUCT-WORK-SPEC-002
- **task_refs**:
  - PRODUCT-TASK-SPEC-006-01

## Summary

Create temporary PRODUCT-level tooling that can resolve a spec file path to the expected path-derived `spec:` ID and validate the accepted spec format contract before existing specs are migrated.

This is a bridge for PRODUCT-WORK-SPEC-005. It is not a DRMCP reimplementation and must not add more patches to the current DRMCP codebase.

## Scope

| area | in scope |
|---|---|
| path-derived spec ID resolver | Given a spec file path, compute the default `spec:` ID using the PRODUCT format contract. |
| parser-aware file validation | Validate real ATX H1 count while ignoring YAML front matter and fenced code blocks. |
| H1 / metadata checks | Validate H1 kind/title format and H1-adjacent `id/status/date/parent` markers. |
| required section checks | Validate required / prohibited sections by spec kind with migration-aware severity. |
| Topics table checks | Validate required `title/kind/parent/file/summary` columns enough to support migration review. |
| output form | Provide either a lightweight script or temporary MCP surface suitable for PRODUCT migration work. |

## Non-scope

| area | owner |
|---|---|
| DRMCP production implementation or reimplementation | DRMCP-WORK-SPEC-001 / DRMCP-WORK-SPEC-002, after DRMCP redesign. |
| Current DRMCP monkey-patch work | Explicitly excluded. |
| Stable ID alias / redirect compatibility design | PRODUCT-WORK-SPEC-002. |
| Existing spec migration edits | PRODUCT-WORK-SPEC-005. |
| App namespace MCP contract redesign | Existing DRMCP-REQ-MCP-001 / DRMCP-INV-MCP-001 and later concrete WORK. |

## Dependencies

| dependency | reason |
|---|---|
| PRODUCT-WORK-SPEC-001 | Defines the accepted spec format and path-derived ID rules. |
| PRODUCT-WORK-SPEC-002 | Needed before hardening alias / redirect / mismatch behavior. |

## Done condition

| item | done when |
|---|---|
| resolver available | A path-derived spec ID resolver can be run against target spec paths. |
| validator available | Temporary tooling validates the accepted spec-format checks needed by migration. |
| diagnostics usable | Output is actionable for PRODUCT-WORK-SPEC-005 migration review. |
| boundary clean | No current DRMCP implementation code is changed. |

## Source records

| ref | role |
|---|---|
| PRODUCT-REQ-SPEC-001 | Requirement for MCP-readable spec format and topic tree support. |
| PRODUCT-WORK-SPEC-001 | Defines the format contract and this temporary tooling follow-up. |
| PRODUCT-WORK-SPEC-002 | Defines compatibility behavior needed before strict mismatch handling. |

## Evidence

- Output form decided: standalone Python script (not MCP surface). Rationale: temporary bridge; MCP surface implies stable contract; script is disposable when DRMCP reimplementation lands.
- Created `product/src/tools/validate_spec.py`.
- Inventory mode: 10 pre-migration spec files produce 49 warnings, 0 errors. New-format specs (10 files) pass clean.
- Strict mode: pre-migration diagnostics escalate to errors; new-format specs still clean.
- ID derivation spot-checked against examples in `spec:product.concepts.spec_format.spec_id_as_ref`.
- No DRMCP implementation code changed.
- No `v01/records/**` files changed.
- Done condition satisfied: resolver available, validator available, diagnostics actionable for WORK-SPEC-005.
