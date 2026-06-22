# PRODUCT-WORK-SPEC-004: Ownership boundary decision and relocation plan

- **id**: PRODUCT-WORK-SPEC-004
- **status**: done
- **date**: 2026-06-10
- **requirement_refs**:
  - PRODUCT-REQ-SPEC-001
- **investigation_refs**:
  - PRODUCT-INV-SPEC-002
  - PRODUCT-INV-SPEC-004
- **source_work_items**:
  - PRODUCT-WORK-SPEC-001
- **task_refs**:
  - PRODUCT-TASK-SPEC-004-01
  - PRODUCT-TASK-SPEC-004-02
  - PRODUCT-TASK-SPEC-004-03

## Summary

Turn PRODUCT-INV-SPEC-002 and PRODUCT-INV-SPEC-004 ownership findings into one accepted ownership boundary decision and relocation plan.

## Scope

| area | in scope |
|---|---|
| ownership decision | Decide PRODUCT-owned, DRMCP-owned, and hybrid boundaries for artifact / traceability / namespace-model specs. |
| relocation plan | Define source files/sections, destination paths, and move/split order. |
| dependency update plan | Identify records and references that must be updated when relocation happens. |
| migration handoff | Produce prerequisites for PRODUCT-WORK-SPEC-005, and identify whether PRODUCT-WORK-SPEC-009's format-only migration scope needs a new target file (`namespace-model/index.md`). |

## Non-scope

| area | owner |
|---|---|
| investigation | PRODUCT-INV-SPEC-002 / PRODUCT-INV-SPEC-004 |
| actual relocation execution | PRODUCT-WORK-SPEC-005 |
| format-only migration execution | PRODUCT-WORK-SPEC-009 |
| compatibility rules | PRODUCT-WORK-SPEC-002 |
| DRMCP validation implementation | DRMCP-WORK-SPEC-001 / DRMCP-WORK-SPEC-002 |

## Dependencies

| dependency | reason |
|---|---|
| PRODUCT-INV-SPEC-002 | Provides section-level ownership classification for traceability / project-artifact-model specs. |
| PRODUCT-INV-SPEC-004 | Provides section-level ownership classification for namespace-model vs. DRMCP namespace/ID-grammar content. |
| PRODUCT-WORK-SPEC-001 | Defines the spec format and migration boundary. |

## Done condition

| item | done when |
|---|---|
| boundary accepted | Ownership boundary decision is accepted by user and/or ADR-style decision artifact. |
| relocation plan complete | Candidate source/destination list and ordering are documented. |
| migration handoff ready | PRODUCT-WORK-SPEC-005 can execute without re-deciding ownership. |
| no bulk migration | No existing specs are moved or rewritten in this work. |

## Source records

| ref | role |
|---|---|
| PRODUCT-REQ-SPEC-001 | Requirement motivating spec format and topic tree support. |
| PRODUCT-WORK-SPEC-001 | Defines ownership boundary as migration gate. |
| PRODUCT-INV-SPEC-002 | Required investigation input — traceability / project-artifact-model classification. |
| PRODUCT-INV-SPEC-004 | Required investigation input — namespace-model vs. DRMCP classification. |

## Evidence

- PRODUCT-INV-SPEC-004 found that `namespace-model/index.md` is missing the currently-implemented (v1) namespace_prefix derivation algorithm, multi-root scan behavior, and public/bare record-ID grammar — these are presently described only inside DRMCP's `overview.md` / `schema.md`, discovered while planning PRODUCT-TASK-SPEC-005-13..-16. Recorded here so the boundary decision (PRODUCT-TASK-SPEC-004-01) treats both investigations as joint input.

