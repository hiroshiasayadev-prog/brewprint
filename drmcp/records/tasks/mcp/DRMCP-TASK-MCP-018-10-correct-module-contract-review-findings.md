# DRMCP-TASK-MCP-018-10: Correct module-contract review findings

- **id**: DRMCP-TASK-MCP-018-10
- **status**: done
- **date**: 2026-07-06
- **work_item**: DRMCP-WORK-MCP-018
- **task_type**: correction
- **estimate**: 0.5d
- **depends_on**:
  - DRMCP-TASK-MCP-018-09
- **outputs**:
  - DRMCP-TASK-MCP-018-10
  - `spec:drmcp.implementation.contracts.application_use_cases.contract_boundary`
  - `spec:drmcp.implementation.contracts.record_domain_logical_tree.contract_boundary`
  - DRMCP-WORK-MCP-018

## Goal

Correct the W018 integrated review findings F-MAJ-W018-07-01 and F-MIN-W018-07-01 without reopening accepted decisions.

## Work

- Project Current Records Snapshot Assembly and Legacy Lookup State Assembly into the Application Use Cases contract boundary as shared-orchestration behavioral components.
- Preserve Current Records Snapshot and Legacy Lookup State as handoff state surfaces.
- Clarify that Domain owns semantic structures inside snapshots, not request-level source-access orchestration.
- Correct W018 Impact Scope to distinguish current W018 canonical targets from downstream detailed-contract target selection.
- Do not add field schemas, Go signatures, package layout, algorithms, fixtures, or operation behavior Specifications.
- Do not close findings.
- Do not synchronize W018 closure.

## Done condition

- F-MAJ-W018-07-01 is repaired by projecting the two Assembly behavioral components at baseline level.
- F-MIN-W018-07-01 is repaired by updating W018 Impact Scope.
- The correction preserves D-001 through D-012 and ADR-013.
- The correction does not release implementation planning.
- The correction is ready for independent finding-closure review.

## Verification

- Confirmed Application Use Cases contract boundary now names Current Records Snapshot Assembly and Legacy Lookup State Assembly as shared-orchestration components.
- Confirmed the same Specification still treats Current Records Snapshot and Legacy Lookup State as handoff state surfaces.
- Confirmed Record Domain / Logical Tree contract boundary states that Domain owns semantic structures and does not own request-level source-access orchestration.
- Confirmed W018 Impact Scope no longer says W018 contract Specification targets remain undecided.
- Confirmed no implementation Task, production implementation, or implementation planning was added.

## Evidence

- Source review Task: DRMCP-TASK-MCP-018-07.
- Source coordination Task: DRMCP-TASK-MCP-018-09.
- Corrected finding F-MAJ-W018-07-01 in `spec:drmcp.implementation.contracts.application_use_cases.contract_boundary` and direct Domain-boundary consistency text.
- Corrected finding F-MIN-W018-07-01 in DRMCP-WORK-MCP-018 Impact Scope.
- The user confirmed the placement rule: Assembly belongs in Application shared orchestration; Domain owns semantic structures and lookup or validation semantics.
- Finding closure remains owned by DRMCP-TASK-MCP-018-11.
- DRMCP is non-operational. Filesystem authoring was used as the required fallback.
