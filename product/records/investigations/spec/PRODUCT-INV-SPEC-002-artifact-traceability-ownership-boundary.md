# PRODUCT-INV-SPEC-002: Artifact and traceability ownership boundary

- **id**: PRODUCT-INV-SPEC-002
- **status**: not_started
- **date**: 2026-06-10
- **requirement_refs**:
  - PRODUCT-REQ-SPEC-001
- **source_work_items**:
  - PRODUCT-WORK-SPEC-001

## Question

Which artifact and traceability spec sections are PRODUCT-owned cross-app semantics, DRMCP-owned tool contracts, or hybrid content that needs split or relocation before spec migration?

## Scope

| area | in scope |
|---|---|
| traceability specs | Inspect `product/records/spec/concepts/traceability/**` section by section. |
| project artifact model | Inspect `product/records/spec/concepts/project-artifact-model/index.md`. |
| ownership classification | Classify sections as PRODUCT-owned semantics, DRMCP-owned tool contract, or hybrid. |
| relocation candidates | Identify files or sections that may need relocation before migration. |
| migration prerequisites | Produce recommendations needed before PRODUCT-WORK-SPEC-004 and PRODUCT-WORK-SPEC-005. |

## Non-scope

| area | owner |
|---|---|
| actual relocation | PRODUCT-WORK-SPEC-004 / PRODUCT-WORK-SPEC-005 |
| format contract changes | PRODUCT-WORK-SPEC-001 or successor work |
| DRMCP implementation | DRMCP work items |
| bulk migration | PRODUCT-WORK-SPEC-005 |

## Expected evidence

| evidence | purpose |
|---|---|
| section-level classification table | Shows PRODUCT / DRMCP / hybrid ownership. |
| relocation candidate list | Feeds PRODUCT-WORK-SPEC-004. |
| ambiguity list | Captures unresolved ownership questions. |
| migration risk notes | Prevents migration before boundary decisions are made. |

## Done condition

| item | done when |
|---|---|
| classification complete | Target files/sections are classified. |
| relocation candidates identified | Candidate files/sections and reasons are listed. |
| decision handoff ready | PRODUCT-WORK-SPEC-004 can decide relocation plan from the investigation. |
| no implementation | No implementation or bulk migration is performed. |

## Source records

| ref | role |
|---|---|
| PRODUCT-REQ-SPEC-001 | Requirement for MCP-readable spec format and topic tree support. |
| PRODUCT-INV-SPEC-001 | Found ownership-sensitive traceability concepts as migration dependency. |
| PRODUCT-WORK-SPEC-001 | Defines this investigation as a migration gate. |

## Evidence

