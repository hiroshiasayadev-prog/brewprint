# DRMCP-TASK-MCP-016-14: Author portable standards realignment authorities

- **id**: DRMCP-TASK-MCP-016-14
- **status**: done
- **date**: 2026-07-04
- **work_item**: DRMCP-WORK-MCP-016
- **task_type**: authoring
- **estimate**: 1.0d
- **depends_on**:
  - DRMCP-TASK-MCP-016-13
- **outputs**:
  - DRMCP-REQ-MCP-003
  - DRMCP-ADR-MCP-007
  - DRMCP-ADR-MCP-008
  - DRMCP-ADR-MCP-009
  - DRMCP-ADR-MCP-010
  - DRMCP-ADR-MCP-011
  - DRMCP-ADR-MCP-012
  - DRMCP-TASK-MCP-016-14

## Goal

Author the revised Requirement and ADR authority selected by T12 and routed by T13.

## Work

- Amend DRMCP-REQ-MCP-003 while preserving its portable-package requirement identity.
- Create ADR-010 for the five-component whole-application model.
- Create ADR-012 for unified Current Records state and lifecycle.
- Create ADR-011 for inward ownership and Guidance query aliases.
- Mark ADR-007, ADR-008, and ADR-009 `superseded`.
- Preserve ADR-001 unchanged as current upstream authority.
- Record affected Specification areas without authoring their normative text.

## Done condition

- DRMCP-REQ-MCP-003 requires normal Current Records treatment for the portable package.
- ADR-010, ADR-011, and ADR-012 are accepted and contain complete rationale and consequences.
- ADR-010 supersedes ADR-007.
- ADR-011 supersedes ADR-008.
- ADR-012 supersedes ADR-009.
- ADR-001 remains accepted and unchanged.
- T15 has complete authority for canonical Specification authoring.

## Verification

- Confirm Requirement identity, status, and portability intent remain intact.
- Confirm replacement ADR metadata and supersession chains are exact.
- Confirm old accepted ADR bodies remain historical and only status changes.
- Confirm ADR-011 depends on ADR-010 and ADR-012.
- Confirm no Specification, review verdict, graph, lifecycle closure, implementation, stage, or commit changed.

## Evidence

- The repository-root `prompt_chappy.md` was read first.
- `CLAUDE.md` and `AGENTS.md` were not read.
- DRMCP is non-operational. Design Record authoring used filesystem fallback.
- T12 D-018 through D-024 supplied the complete selected authority.
- T13 routed three replacement ADR boundaries and one Requirement amendment.
- DRMCP-REQ-MCP-003 was amended to treat the package as a configured spec-only Current Records source under `design_records`.
- ADR-010 was created and supersedes ADR-007.
- ADR-012 was created and supersedes ADR-009.
- ADR-011 was created and supersedes ADR-008.
- ADR-001 was not changed.
- ADR-007, ADR-008, and ADR-009 retain their historical bodies and are marked `superseded`.
- ADR-010, ADR-011, and ADR-012 are `accepted` with exact replacement metadata.
- ADR-011 depends on ADR-010 and ADR-012.
- The complete scoped authority diff was inspected without truncation.
- Scoped whitespace verification passed.
- No scoped file is staged.
- The standalone semantic responsibility validator was not executed because no operational invocation tool is available. No validator PASS was synthesized.
- No canonical Specification was authored by this Task.
- No independent review occurred.
- No stage or commit was performed.
