# DRMCP-TASK-MCP-001-01: Define child Work Item split and ownership gates

- **id**: DRMCP-TASK-MCP-001-01
- **status**: done
- **date**: 2026-06-26
- **work_item**: DRMCP-WORK-MCP-001
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 1d
- **depends_on**: []
- **outputs**:
  - DRMCP-REQ-MCP-001
  - DRMCP-WORK-MCP-001
  - PRODUCT-REQ-SPEC-001
  - PRODUCT-REQ-SPEC-004
  - PRODUCT-WORK-SPEC-014
  - PRODUCT-WORK-SPEC-015
  - DRMCP-WORK-MCP-003
  - DRMCP-WORK-MCP-004
  - DRMCP-WORK-MCP-005
  - DRMCP-WORK-MCP-006
  - DRMCP-WORK-MCP-007
  - DRMCP-WORK-MCP-008
  - DRMCP-WORK-MCP-009
  - DRMCP-WORK-MCP-010
  - DRMCP-TASK-MCP-001-05
  - DRMCP-TASK-MCP-001-06
  - DRMCP-TASK-MCP-001-07
  - DRMCP-TASK-MCP-001-08
  - DRMCP-TASK-MCP-001-09
  - DRMCP-TASK-MCP-001-10

## Goal

Define the executable child Work Items and lifecycle gates for the corrected DRMCP read baseline.

## Work

- Inventory the contract, fixture, implementation, compatibility, and validation-disposition workstreams.
- Classify each workstream by owning app namespace and source Requirement.
- Reuse a suitable existing Work Item when its boundary and source Requirement match.
- Create a child Work Item when an independently executable workstream has no suitable owner.
- Keep detailed Tasks, execution evidence, and local review inside each child Work Item.
- Update Tasks T02-T10 with the exact child Work Item ID selected for each gate.
- Record an explicit disposition when a candidate requires no separate child Work Item.

This Task defines and creates the execution structure.
This Task does not perform contract correction, fixture authoring, or implementation.

## Done condition

- Every T02-T10 workstream has an exact child Work Item ID or an explicit accepted no-child disposition.
- Each selected child Work Item has the correct owner and source Requirement.
- Each child Work Item boundary excludes hub-level lifecycle tracking.
- `DRMCP-WORK-MCP-001` reflects the accepted child Work Item graph and dependencies.
- No independently executable implementation work remains assigned directly to a hub Task.

## Verification

- Compare each child Work Item against `DRMCP-REQ-MCP-001` and `DRMCP-ADR-MCP-001`.
- Confirm reciprocal Requirement and Task relations required by the authoring standards.
- Review the graph for duplicated scope, missing gates, and cross-owner leakage.

## Evidence

2026-06-26 accepted child Work Item graph:

| hub Task | owner chain | accepted child Work Item boundary |
|---|---|---|
| T02 | `PRODUCT-REQ-SPEC-004` | `PRODUCT-WORK-SPEC-014`: Remove obsolete `V01-SPEC-*` compatibility authority. |
| T03 | `DRMCP-REQ-MCP-001` | `DRMCP-WORK-MCP-003`: Current discovery, current spec parsing, and active-index contracts. |
| T04 | `DRMCP-REQ-MCP-001` | `DRMCP-WORK-MCP-004`: Compact active listing and exact batch retrieval contracts. |
| T05 | `DRMCP-REQ-MCP-001` | `DRMCP-WORK-MCP-005`: Current-first resolver and configured legacy-fallback contracts. |
| T06 | `DRMCP-REQ-MCP-001` | `DRMCP-WORK-MCP-006`: Validation execution, diagnostics, and path-exposure contracts. |
| T07 | `DRMCP-REQ-MCP-001` and `PRODUCT-REQ-SPEC-001` | `DRMCP-WORK-MCP-007`: Validation Work Item disposition. `PRODUCT-WORK-SPEC-015`: PRODUCT owner-pointer synchronization. |
| T08 | `DRMCP-REQ-MCP-001` | `DRMCP-WORK-MCP-008`: Current and legacy read fixtures. |
| T09 | `DRMCP-REQ-MCP-001` | `DRMCP-WORK-MCP-009`: Current-format read implementation. |
| T10 | `DRMCP-REQ-MCP-001` | `DRMCP-WORK-MCP-010`: Configured legacy archive fallback implementation. |

Verification results:

- Every T02-T10 lifecycle Task has an exact child Work Item target.
- T07 uses two owner-local Work Items instead of one cross-owner implementation container.
- `DRMCP-WORK-SPEC-001/002` remain disposition candidates and were not reused as T07 containers.
- `DRMCP-REQ-MCP-001` lists `DRMCP-WORK-MCP-003` through `DRMCP-WORK-MCP-010` plus the hub Work Item.
- `PRODUCT-REQ-SPEC-001` lists `PRODUCT-WORK-SPEC-015`.
- `DRMCP-WORK-MCP-001` records the accepted child graph and exact lifecycle gates.
- Contract, fixture, and implementation evidence remain owned by the selected child Work Items.
- No independently executable implementation work remains assigned directly to a hub Task.
