# Concept: Follow-up boundary

- **id**: `spec:product.design_records.spec_format.follow_up_boundary`
- **status**: accepted
- **date**: 2026-06-14
- **parent**: `spec:product.design_records.spec_format`

## What this is

This spec records the ownership boundary around the spec-format contract and its follow-up work. It separates PRODUCT specification stabilization, temporary PRODUCT migration tooling, later DRMCP implementation work, and historical `v01/` snapshots.

The boundary is part of the spec-format contract because future agents could otherwise misread validation rows as permission to patch current DRMCP code or migrate unrelated specs during format cleanup.

## Concept model

| follow-up | owns |
|---|---|
| PRODUCT-WORK-SPEC-002 | Done. Canonical ref contract, ref-first `## Topics`, and explicit compatibility exception decisions are recorded in `spec:product.design_records.spec_format.spec_id_as_ref`. |
| PRODUCT-WORK-SPEC-003 | Spec authoring guide update. |
| PRODUCT-INV-SPEC-002 | Artifact / traceability ownership boundary investigation before migration. |
| PRODUCT-WORK-SPEC-004 | Ownership boundary decision and relocation plan. |
| PRODUCT-WORK-SPEC-006 | Temporary PRODUCT-level spec format validator / resolver tooling before migration. |
| PRODUCT-WORK-SPEC-005 | Existing spec format migration and restructuring using temporary validation tooling. |
| PRODUCT-WORK-SPEC-008 | ADR `target_specs` traceability investigation and update planning. |
| DRMCP-REQ-MCP-001 / DRMCP-INV-MCP-001 | Existing app namespace / multi-root MCP contract redesign input after spec restructuring. |
| DRMCP-WORK-SPEC-001 | Later DRMCP implementation-phase parser-aware spec format validation; not prerequisite for PRODUCT spec-format stabilization. |
| DRMCP-WORK-SPEC-002 | Later DRMCP implementation-phase Index Topics graph validation; not prerequisite for PRODUCT spec-format stabilization. |

## Rules

Work in the spec-format area must not silently cross ownership boundaries. PRODUCT spec-format cleanup can define and split the format contract, but it does not reopen completed work items, patch current DRMCP implementation code, migrate existing specs, or edit historical `v01/` records unless a later accepted work item explicitly changes that scope.

Specs should not carry reverse traceability tables. Semantic spec changes are governed by ADR decisions, and ADRs may later carry `target_specs` metadata after PRODUCT-WORK-SPEC-008 determines the update plan. WORK/TASK `## Evidence` may list edited files, but that evidence is work history rather than the spec traceability source of truth.

## Boundary

| not owned here | owner / reason |
|---|---|
| Current DRMCP implementation | DRMCP-WORK-SPEC-001 / 002 are later implementation-phase follow-ups after PRODUCT stabilization and DRMCP redesign; this spec does not require patching current DRMCP. |
| Temporary validation tooling | PRODUCT-WORK-SPEC-006 owns the bridge tooling needed before migration. |
| Actual migration of existing specs | PRODUCT-WORK-SPEC-005. |
| Path-derived canonical ref compatibility exceptions | PRODUCT-WORK-SPEC-002. |
| ADR `target_specs` traceability investigation | PRODUCT-WORK-SPEC-008. |
| Authoring guide text | PRODUCT-WORK-SPEC-003. |
| App namespace MCP contract redesign | Existing DRMCP-REQ-MCP-001 / DRMCP-INV-MCP-001 capture this follow-up; this spec does not decide the redesign. |
| UI topic tree display | Future UI work, not needed for this contract. |
| v01 snapshot updates | `v01/` is historical and must not be modified by this contract. |

## Related specs

| ref | relation |
|---|---|
| `spec:product.design_records.spec_format` | Parent Index for this concept. |
| `spec:product.design_records.spec_format.validation_policy` | Defines validation ownership and temporary tooling boundary. |
| `spec:product.design_records.spec_format.spec_id_as_ref` | Defers alias, redirect, stale-ref, and derived-topic compatibility exceptions to PRODUCT-WORK-SPEC-002. |
