# DRMCP-WORK-SPEC-002: Index Topics graph validation

- **id**: DRMCP-WORK-SPEC-002
- **status**: not_started
- **date**: 2026-06-10
- **requirement_refs**:
  - PRODUCT-REQ-SPEC-001
- **source_work_items**:
  - PRODUCT-WORK-SPEC-001
  - PRODUCT-WORK-SPEC-002
  - DRMCP-WORK-SPEC-001
- **task_refs**:

## Summary

Implement cross-file graph validation for spec topic trees declared through `## Topics` tables in a future Design Records MCP implementation phase.

This is not a prerequisite for PRODUCT spec-format stabilization. PRODUCT migration validation should use PRODUCT-WORK-SPEC-006 or an explicitly accepted equivalent before this DRMCP implementation work begins.

## Scope

| area | in scope |
|---|---|
| Topics table shape | Validate required columns `title/kind/parent/file/summary`. |
| child target resolution | Resolve `file` column targets to spec files. |
| parent grammar | Validate `parent` values against accepted `spec:` ID-as-ref grammar. |
| declaring parent consistency | Validate that row `parent` matches the declaring spec `id` unless compatibility rules allow an exception. |
| duplicate parent | Detect duplicate authoritative parents for the same child. |
| cycle detection | Detect topic graph cycles after parent/child resolution. |
| diagnostics | Emit actionable diagnostics for graph validation failures. |

## Non-scope

| area | owner |
|---|---|
| per-file H1 / metadata / section validation | DRMCP-WORK-SPEC-001 |
| PRODUCT compatibility design | PRODUCT-WORK-SPEC-002 |
| temporary PRODUCT validation tooling | PRODUCT-WORK-SPEC-006 |
| current DRMCP monkey-patch work | Explicitly excluded until DRMCP redesign/reimplementation begins. |
| migration execution | PRODUCT-WORK-SPEC-005 |
| UI topic tree rendering | future UI work |

## Dependencies

| dependency | reason |
|---|---|
| PRODUCT-WORK-SPEC-001 | Defines `## Topics` table contract and parent grammar. |
| PRODUCT-WORK-SPEC-002 | Must settle alias / redirect / derived topic compatibility before graph exception rules harden. |
| DRMCP-WORK-SPEC-001 | Per-file validation should run before graph validation. |
| PRODUCT-WORK-SPEC-006 | Temporary migration tooling may provide graph-validation cases for the later DRMCP implementation. |
| DRMCP redesign/reimplementation plan | Required before implementing this work in DRMCP. |

## Done condition

| item | done when |
|---|---|
| graph validator implemented | DRMCP validates Topics edges across files. |
| diagnostics implemented | Duplicate parent, unresolved child target, invalid parent, row.parent mismatch, and cycle diagnostics exist. |
| tests added | Fixtures cover valid graph, missing child, duplicate parent, invalid parent, row.parent mismatch, and cycle. |
| no per-file duplication | Per-file validation remains in DRMCP-WORK-SPEC-001. |
| phase boundary preserved | Work starts only in the DRMCP implementation/reimplementation phase, after PRODUCT spec-format stabilization is no longer blocked on it. |

## Source records

| ref | role |
|---|---|
| PRODUCT-REQ-SPEC-001 | Requirement for MCP-readable spec format and topic tree support. |
| PRODUCT-WORK-SPEC-001 | Defines topic table and graph validation follow-up split. |
| PRODUCT-WORK-SPEC-002 | Provides compatibility rules needed by graph validation. |

## Evidence
