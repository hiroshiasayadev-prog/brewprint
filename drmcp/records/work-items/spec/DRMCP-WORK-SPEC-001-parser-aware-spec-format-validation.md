# DRMCP-WORK-SPEC-001: Parser-aware spec format validation

- **id**: DRMCP-WORK-SPEC-001
- **status**: not_started
- **date**: 2026-06-10
- **requirement_refs**:
  - PRODUCT-REQ-SPEC-001
- **source_work_items**:
  - PRODUCT-WORK-SPEC-001
- **task_refs**:

## Summary

Implement parser-aware per-file validation for the PRODUCT-owned spec format contract in a future Design Records MCP implementation phase.

This is not a prerequisite for PRODUCT spec-format stabilization and must not be used as justification to add more patches to the current DRMCP codebase before PRODUCT-WORK-SPEC-005 migration work.

## Scope

| area | in scope |
|---|---|
| H1 validation | Validate exactly one real ATX H1 outside YAML front matter and fenced code blocks. |
| H1 format | Validate `# <SpecKind>: <Title>` and accepted spec kinds. |
| H1-adjacent metadata | Validate required `id/status/date/parent` markers. |
| required sections | Validate required / prohibited sections by spec kind with migration-aware severity. |
| front matter policy | Warn for existing front matter and error for new/migrated specs under the new format. |
| spec ID mismatch | Warn when visible `id` does not match the path-derived default spec ID unless compatibility rules allow it. |
| parser strategy | Reuse or extend parser behavior to ignore fenced headings and front matter when counting headings. |

## Non-scope

| area | owner |
|---|---|
| Index Topics graph validation | DRMCP-WORK-SPEC-002 |
| PRODUCT format design | PRODUCT-WORK-SPEC-001 |
| ID/ref compatibility design | PRODUCT-WORK-SPEC-002 |
| temporary PRODUCT validation tooling | PRODUCT-WORK-SPEC-006 |
| current DRMCP monkey-patch work | Explicitly excluded until DRMCP redesign/reimplementation begins. |
| migration execution | PRODUCT-WORK-SPEC-005 |

## Dependencies

| dependency | reason |
|---|---|
| PRODUCT-WORK-SPEC-001 | Defines the spec format contract. |
| PRODUCT-WORK-SPEC-002 | Desirable before finalizing mismatch severity and alias behavior. |
| PRODUCT-WORK-SPEC-006 | Temporary migration tooling may provide fixtures and diagnostic expectations for the later DRMCP implementation. |
| DRMCP redesign/reimplementation plan | Required before implementing this work in DRMCP. |

## Done condition

| item | done when |
|---|---|
| diagnostics implemented | DRMCP emits parser-aware diagnostics for per-file spec format violations. |
| tests added | Tests cover real H1s versus fenced-code headings, H1-adjacent metadata, front matter, and required sections. |
| migration severity | Existing specs can warn without blocking migration prematurely. |
| no graph scope creep | Cross-file Topics graph validation is deferred to DRMCP-WORK-SPEC-002. |
| phase boundary preserved | Work starts only in the DRMCP implementation/reimplementation phase, after PRODUCT spec-format stabilization is no longer blocked on it. |

## Source records

| ref | role |
|---|---|
| PRODUCT-REQ-SPEC-001 | Requirement for MCP-readable spec format and topic tree support. |
| PRODUCT-WORK-SPEC-001 | Defines the format contract and validation split. |

## Evidence
