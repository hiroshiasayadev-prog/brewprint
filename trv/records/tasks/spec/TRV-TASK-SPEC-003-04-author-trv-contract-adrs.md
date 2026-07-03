# TRV-TASK-SPEC-003-04: Author TRV contract ADRs

- **id**: TRV-TASK-SPEC-003-04
- **status**: done
- **date**: 2026-07-02
- **work_item**: TRV-WORK-SPEC-003
- **task_type**: authoring
- **estimate**: 0.5d
- **depends_on**:
  - TRV-TASK-SPEC-003-03
- **outputs**:
  - TRV-TASK-SPEC-003-04
  - TRV-ADR-SPEC-003
  - TRV-ADR-SPEC-004
  - TRV-ADR-SPEC-005

## Goal

Author the three TRV contract ADR boundaries selected by T02.

## Work

- Create TRV-ADR-SPEC-003 for the standalone stdio MCP interface and tagged outcomes.
- Create TRV-ADR-SPEC-004 for repository-root-relative Task-path input.
- Create TRV-ADR-SPEC-005 for semantic-only compatibility across future DRMCP integration.
- Set all three ADRs to `accepted`.
- Preserve PRODUCT-ADR-SPEC-017 as reused caller-workflow authority.
- Keep concrete schema encoding, Go types, path mechanics, and library choices outside the ADRs.

This Task must not change routing, author contract Specifications, perform review or closure, or start detailed design or implementation.

## Done condition

- All three ADRs exist with accepted lifecycle state and complete rationale.
- ADR boundaries match T02 exactly.
- W004-owned detail remains excluded.

## Verification

- Confirm IDs, metadata, paths, dependencies, and sections follow active ADR rules.
- Confirm no accepted contract decision changed.
- Confirm only declared outputs changed.

## Evidence

- Consumed TRV-TASK-SPEC-001-02 decisions D-002, D-003, D-009, D-010, and D-014.
- Preserved the T02 routing boundary without modification.
- Created `trv/records/adr/spec/TRV-ADR-SPEC-003-use-standalone-stdio-mcp-interface-and-tagged-outcomes.md` with `accepted` status.
- Created `trv/records/adr/spec/TRV-ADR-SPEC-004-use-repository-root-relative-task-path-input.md` with `accepted` status.
- Created `trv/records/adr/spec/TRV-ADR-SPEC-005-preserve-semantic-compatibility-across-future-drmcp-integration.md` with `accepted` status.
- TRV-ADR-SPEC-003 projects D-002 and D-009.
- TRV-ADR-SPEC-004 projects D-003.
- TRV-ADR-SPEC-005 projects D-014.
- PRODUCT-ADR-SPEC-017 remains reused authority for D-010; no new human-judgment ADR was created.
- Preserved the W002 five-component architecture, application-port ownership, dependency direction, orchestration ownership, prompt ownership, provider boundary, and application-outcome ownership.
- The MCP adapter only projects application outcomes and does not recalculate, complete, or reinterpret them.
- Excluded exact JSON Schema syntax, Go types, normalization, separators, case handling, symlink behavior, filesystem APIs, error encoding, MCP library, serialization, retry, timeout, provider schema, packages, symbols, constructors, commands, and implementation planning.
- DRMCP is non-operational under the current agent authoring policy, so filesystem authoring was used.
- Scoped Git inspection covered only the three ADR paths and this Task.
- Scoped diff inspection confirmed the declared file set and complete ADR content.
- Scoped whitespace inspection returned `pass`; LF-to-CRLF warnings were advisory only.
- No Specification authoring, review, finding correction, lifecycle synchronization, parent synchronization, detailed design, implementation, stage, or commit occurred.
- Post-authoring inspection found that W003 and T05 materialize only external MCP, Task-input, caller, and compatibility topics.
- The current graph does not materialize architecture-derived application contracts for the inbound use-case port, outbound ports, application models, outcome invariants, or boundary-to-boundary failure projection.
- The created ADRs satisfy T04's routed ADR-authoring boundary but do not complete W003 or decide the later Specification partition.
- Result: `PASS` for the bounded T04 authoring outcome.
