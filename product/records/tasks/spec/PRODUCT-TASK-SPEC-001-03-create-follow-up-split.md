# PRODUCT-TASK-SPEC-001-03: Create follow-up split

- **id**: PRODUCT-TASK-SPEC-001-03
- **status**: done
- **date**: 2026-06-10
- **work_item**: PRODUCT-WORK-SPEC-001
- **source_requirement**: PRODUCT-REQ-SPEC-001
- **source_investigation**: PRODUCT-INV-SPEC-001
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-001-02
- **outputs**:
  - PRODUCT-WORK-SPEC-002
  - PRODUCT-WORK-SPEC-003
  - PRODUCT-INV-SPEC-002
  - PRODUCT-WORK-SPEC-004
  - PRODUCT-WORK-SPEC-005
  - PRODUCT-WORK-SPEC-006
  - DRMCP-REQ-MCP-001 / DRMCP-INV-MCP-001
  - DRMCP-WORK-SPEC-001
  - DRMCP-WORK-SPEC-002

## Goal

Create or explicitly queue the follow-up artifacts required after the spec format contract is reviewed.

## Work

| follow-up artifact | kind | purpose | dependency / gate |
|---|---|---|---|
| PRODUCT-WORK-SPEC-002 | WORK | Stable `spec:` ID-as-ref and derived topic ref compatibility design. | Requires accepted format contract. |
| PRODUCT-WORK-SPEC-003 | WORK | Spec authoring guide format update. | Requires accepted format contract. |
| PRODUCT-INV-SPEC-002 | INV | Artifact / traceability ownership boundary investigation. | Required before migration / relocation work. |
| PRODUCT-WORK-SPEC-004 | WORK | Ownership boundary decision and relocation plan. | Requires PRODUCT-INV-SPEC-002. |
| PRODUCT-WORK-SPEC-006 | WORK | Temporary PRODUCT-level spec format validator / resolver tooling. | Required before migration validation; not a DRMCP reimplementation. |
| PRODUCT-WORK-SPEC-005 | WORK | Existing spec format migration and restructuring. | Requires format, compatibility, guide, ownership decisions, and temporary validation tooling. |
| DRMCP-REQ-MCP-001 / DRMCP-INV-MCP-001 | REQ / INV | Existing app namespace / multi-root MCP contract redesign input. | Revisit after PRODUCT spec restructuring; concrete WORK can be created later. |
| DRMCP-WORK-SPEC-001 | WORK | Parser-aware spec format validation. | Implementation-phase follow-up; not required before PRODUCT spec-format stabilization. |
| DRMCP-WORK-SPEC-002 | WORK | Index Topics graph validation. | Implementation-phase follow-up after compatibility and DRMCP redesign/reimplementation planning. |

## Done condition

| item | done when |
|---|---|
| split created | Follow-up artifacts are created or explicitly queued with rationale. |
| ordering captured | Dependencies and gates are written into follow-up artifacts or this task evidence. |
| ownership INV | PRODUCT-INV-SPEC-002 is created or explicitly queued before migration. |
| no migration | Existing specs are not migrated in this task. |
| no implementation | DRMCP implementation is not changed in this task. |

## Verification

- Follow-up artifact paths or queued IDs are listed in Evidence.
- PRODUCT-WORK-SPEC-001 follow-up list is synchronized if artifacts are created.

## Evidence

- Applied Opus non-blocking concerns NB-1 through NB-4 to `product/records/spec/concepts/spec-format/index.md`.
  - `## Topic map` is now explicitly a human navigation hint with no parser-validated structure.
  - Reference-kind table requirement now uses a validator-friendly `body H2 containing at least one Markdown table` rule.
  - `## Topics` row `parent` should match the declaring spec `id`; exceptions are deferred to PRODUCT-WORK-SPEC-002.
  - Related specs now notes current `spec:trace.semantic-ref` and path-derived `spec:product.concepts.traceability.semantic_ref` coexistence pending PRODUCT-WORK-SPEC-002.
- Created follow-up artifacts:
  - `product/records/work-items/spec/PRODUCT-WORK-SPEC-002-stable-spec-id-as-ref-and-derived-topic-compatibility.md`
  - `product/records/work-items/spec/PRODUCT-WORK-SPEC-003-spec-authoring-guide-format-update.md`
  - `product/records/investigations/spec/PRODUCT-INV-SPEC-002-artifact-traceability-ownership-boundary.md`
  - `product/records/work-items/spec/PRODUCT-WORK-SPEC-004-ownership-boundary-decision-and-relocation-plan.md`
  - `product/records/work-items/spec/PRODUCT-WORK-SPEC-005-existing-spec-format-migration-and-restructuring.md`
  - `product/records/work-items/spec/PRODUCT-WORK-SPEC-006-temporary-spec-format-validator-tooling.md`
  - `drmcp/records/work-items/spec/DRMCP-WORK-SPEC-001-parser-aware-spec-format-validation.md`
  - `drmcp/records/work-items/spec/DRMCP-WORK-SPEC-002-index-topics-graph-validation.md`
- Existing app namespace redesign refs found and reused instead of creating a thick new scaffold:
  - `drmcp/records/requirements/mcp/DRMCP-REQ-MCP-001-multi-root-multi-namespace-mcp-tool-contract.md`
  - `drmcp/records/investigations/mcp/DRMCP-INV-MCP-001-multi-root-multi-namespace-mcp-tool-contract-investigation.md`
- Corrected follow-up positioning:
  - PRODUCT-WORK-SPEC-006 is the temporary validation bridge for PRODUCT-WORK-SPEC-005.
  - DRMCP-WORK-SPEC-001/002 are implementation-phase follow-ups and are not prerequisites for PRODUCT spec-format stabilization.
- Captured ordering and dependency gates in the follow-up scaffolds.
- PRODUCT-INV-SPEC-002 is created as the ownership boundary investigation required before migration/relocation.
- No existing spec migration was performed.
- No DRMCP implementation code was changed.
- No `v01/records/**` files were changed.
