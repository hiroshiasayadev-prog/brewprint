# PRODUCT-WORK-SPEC-005: Existing spec format migration and restructuring

- **id**: PRODUCT-WORK-SPEC-005
- **status**: not_started
- **date**: 2026-06-10
- **requirement_refs**:
  - PRODUCT-REQ-SPEC-001
- **source_work_items**:
  - PRODUCT-WORK-SPEC-001
  - PRODUCT-WORK-SPEC-002
  - PRODUCT-WORK-SPEC-003
  - PRODUCT-WORK-SPEC-004
  - PRODUCT-WORK-SPEC-006
- **task_refs**:

## Summary

Migrate existing PRODUCT / DRMCP / BPDSL spec files to the accepted spec format after compatibility, authoring guide, ownership decisions, and temporary validation tooling are complete.

## Scope

| area | in scope |
|---|---|
| format migration | Add accepted H1 format, H1-adjacent metadata, required sections, and visible source/topic structures. |
| restructuring | Split or restructure specs that cannot be migrated directly. |
| ownership relocation | Execute accepted relocation plan from PRODUCT-WORK-SPEC-004. |
| alias / compatibility | Apply compatibility decisions from PRODUCT-WORK-SPEC-002. |
| temporary validation | Use PRODUCT-WORK-SPEC-006 tooling to validate target files during migration review. |
| validation hardening | Coordinate when validation severity can move from migration warning to error. |

## Non-scope

| area | owner |
|---|---|
| format design | PRODUCT-WORK-SPEC-001 |
| ID/ref compatibility design | PRODUCT-WORK-SPEC-002 |
| authoring guide design | PRODUCT-WORK-SPEC-003 |
| ownership decision | PRODUCT-WORK-SPEC-004 |
| temporary validator/tooling build | PRODUCT-WORK-SPEC-006 |
| DRMCP implementation | DRMCP-WORK-SPEC-001 / DRMCP-WORK-SPEC-002 |

## Dependencies

| dependency | reason |
|---|---|
| PRODUCT-WORK-SPEC-001 | Accepted format contract. |
| PRODUCT-WORK-SPEC-002 | Compatibility and alias behavior. |
| PRODUCT-WORK-SPEC-003 | Authoring guidance for migration edits. |
| PRODUCT-WORK-SPEC-004 | Ownership boundary and relocation plan. |
| PRODUCT-WORK-SPEC-006 | Temporary validation support for migration; this replaces any need to patch current DRMCP before migration. |

## Done condition

| item | done when |
|---|---|
| migrated specs | Target spec set follows the accepted format. |
| relocation complete | Accepted ownership relocation is executed. |
| temporary validation used | Migration batches are checked with PRODUCT-WORK-SPEC-006 tooling or an explicitly accepted equivalent. |
| validation state clean | Validation diagnostics match post-migration severity expectations. |
| stale aliases addressed | Alias / redirect mappings are present where needed and no untracked stale references remain. |
| review complete | Migration is reviewed in phases rather than as one opaque rewrite. |

## Source records

| ref | role |
|---|---|
| PRODUCT-REQ-SPEC-001 | Requirement for MCP-readable spec format and topic tree support. |
| PRODUCT-INV-SPEC-001 | Investigation evidence and migration classification basis. |
| PRODUCT-WORK-SPEC-001 | Format contract and follow-up split. |

## Evidence
