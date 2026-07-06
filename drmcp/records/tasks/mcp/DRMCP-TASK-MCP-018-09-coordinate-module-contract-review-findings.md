# DRMCP-TASK-MCP-018-09: Coordinate module-contract review findings

- **id**: DRMCP-TASK-MCP-018-09
- **status**: done
- **date**: 2026-07-06
- **work_item**: DRMCP-WORK-MCP-018
- **task_type**: coordination
- **estimate**: 0.5d
- **depends_on**:
  - DRMCP-TASK-MCP-018-07
- **outputs**:
  - DRMCP-TASK-MCP-018-09
  - DRMCP-WORK-MCP-018
  - DRMCP-TASK-MCP-018-10
  - DRMCP-TASK-MCP-018-11
  - DRMCP-TASK-MCP-018-08

## Goal

Materialize the finding-specific correction and closure-review route required by the W018 integrated review verdict.

## Work

- Read the T07 independent review verdict and findings.
- Group findings that share one correction owner and completion judgment.
- Create one correction Task for F-MAJ-W018-07-01 and F-MIN-W018-07-01.
- Create one independent finding-closure review Task after the correction Task.
- Update W018 Task flow and Task Candidates.
- Block T08 closure synchronization until the finding-closure review completes.
- Do not repair findings in this coordination Task.
- Do not perform independent finding closure.
- Do not synchronize W018 closure.

## Done condition

- F-MAJ-W018-07-01 and F-MIN-W018-07-01 have one correction owner.
- A later independent finding-closure review owner exists.
- T08 depends on the finding-closure review, not directly on the failed integrated review.
- No speculative correction or review Task exists outside the named findings.
- No canonical content was corrected by this coordination Task.

## Verification

- Confirmed T10 owns correction for F-MAJ-W018-07-01 and F-MIN-W018-07-01.
- Confirmed T11 owns independent finding-closure review.
- Confirmed T08 was rerouted to wait for T11.
- Confirmed no implementation Task was created.
- Confirmed no production implementation planning was released.

## Evidence

- Source review Task: DRMCP-TASK-MCP-018-07.
- T07 verdict: NEEDS REVISION.
- F-MAJ-W018-07-01 requires Specification projection of Current Records Snapshot Assembly and Legacy Lookup State Assembly as Application shared-orchestration behavioral components.
- F-MIN-W018-07-01 requires W018 Impact Scope to distinguish final W018 target state from downstream detailed-contract target selection.
- The user confirmed that Assembly belongs in Application shared orchestration, while Domain owns semantic structures and lookup or validation semantics.
- DRMCP is non-operational. Filesystem authoring was used as the required fallback.
