# DRMCP-REQ-MCP-008: Relocate application architecture specifications under implementation

- **id**: DRMCP-REQ-MCP-008
- **status**: accepted
- **date**: 2026-07-08
- **source_refs**:
  - DRMCP-REQ-MCP-007
  - DRMCP-WORK-MCP-018
  - DRMCP-ADR-MCP-013

## Requirement

DRMCP implementation-significant architecture Specifications must live under the DRMCP implementation Specification topic tree.

The current standalone `spec:drmcp.application_architecture` tree must be dispositioned because it defines implementation architecture boundaries, collaboration, lifecycle, state, and failure semantics.

## Evidence

- `DRMCP-WORK-MCP-018` kept application-architecture Specifications as current authority and left relocation under `implementation` as a separate topology or migration question.
- `DRMCP-ADR-MCP-013` moved only module contracts under `spec:drmcp.implementation.contracts` and did not move application-architecture Specifications.
- `DRMCP-REQ-MCP-007` requires implementation-ready detailed Specifications to derive from accepted architecture and contracts.
- Direction on 2026-07-08 states that the current application-architecture Specification tree is implementation-facing and should move under `implementation`.

## Required Outcome

- The standalone application-architecture Specification tree has an explicit relocation disposition.
- Spec-internal references affected by the relocation are inventoried before any canonical ref or path change.
- The target implementation Specification topology is decided before authoring changes.
- Canonical Specification refs, topic tables, parent refs, related-spec refs, and prose references are synchronized inside the selected spec scope.
- Relocation preserves the accepted architecture semantics unless a separate decision explicitly changes them.
- Historical ADR, Work Item, Task, and review records are not rewritten solely to hide the former placement.

## Explicitly Excluded Scope

- Re-deciding the DRMCP application architecture.
- Re-deciding W018 module-contract topology.
- Changing component responsibilities, dependency direction, lifecycle semantics, or failure semantics.
- Production implementation.
- Implementation Task authoring.
- Repository-wide historical-reference cleanup outside the selected spec scope.

## Boundary

This Requirement owns the need to relocate or otherwise disposition DRMCP application-architecture Specifications under the implementation Specification topology.

It does not decide the target topology, exact file moves, canonical refs, compatibility handling, or authoring sequence.
