# DRMCP-TASK-MCP-016-06: Author required DRMCP application-architecture ADRs

- **id**: DRMCP-TASK-MCP-016-06
- **status**: done
- **date**: 2026-07-04
- **work_item**: DRMCP-WORK-MCP-016
- **task_type**: authoring
- **estimate**: 1.0d
- **depends_on**:
  - DRMCP-TASK-MCP-016-05
- **outputs**:
  - DRMCP-ADR-MCP-002
  - DRMCP-ADR-MCP-007
  - DRMCP-ADR-MCP-008
  - DRMCP-ADR-MCP-009

## Goal

Author the complete bounded ADR set required by the accepted DRMCP application-architecture routing result.

## Work

- Consume only the exact ADR IDs, dispositions, decision IDs, and boundary partitions recorded by T05.
- Create `DRMCP-ADR-MCP-007` for the whole-application component model.
- Create `DRMCP-ADR-MCP-009` for request-scoped record state and application lifecycle.
- Make `DRMCP-ADR-MCP-009` supersede `DRMCP-ADR-MCP-002`.
- Preserve the retained request-scoped snapshot and composition-lifecycle choices from `DRMCP-ADR-MCP-002`.
- Replace the unconditional Legacy-loading rule with D-009 operation-specific capability selection.
- Create `DRMCP-ADR-MCP-008` for inward dependencies and responsibility ownership.
- Make `DRMCP-ADR-MCP-008` depend on `DRMCP-ADR-MCP-007` and `DRMCP-ADR-MCP-009`.
- Keep current normative Specification text out of ADR bodies.
- Keep Guidance source-path correction outside ADR authoring.
- Keep D-012 and D-013 outside the current ADR set.
- Do not create an ADR for read-before-authoring delivery sequencing.
- Record the exact authored ADR set in Task outputs and Evidence.

## Done condition

- `DRMCP-ADR-MCP-002` is superseded by `DRMCP-ADR-MCP-009`.
- `DRMCP-ADR-MCP-007` records the whole-application component model.
- `DRMCP-ADR-MCP-008` records inward dependencies and responsibility ownership.
- `DRMCP-ADR-MCP-009` records request-scoped record state and application lifecycle.
- Every authored ADR preserves the accepted T03 choice without adding new architecture judgment.
- ADR dependencies and supersession history match T05.
- ADR consequences identify the affected canonical architecture views.
- D-012 and D-013 remain outside the current ADR set.
- Exact ADR outputs are available for T07 and T09.

## Verification

- PASS: Authoring outputs are limited to `DRMCP-ADR-MCP-002`, `DRMCP-ADR-MCP-007`, `DRMCP-ADR-MCP-008`, and `DRMCP-ADR-MCP-009`.
- PASS: This Task changed only its lifecycle, Verification, and Evidence for completion recording.
- PASS: New ADR IDs match their filename prefixes and use the canonical ADR section shape.
- PASS: `DRMCP-ADR-MCP-009` supersedes `DRMCP-ADR-MCP-002`, whose status is `superseded`.
- PASS: `DRMCP-ADR-MCP-009` retains D-004 and D-017 while replacing unconditional Legacy loading through D-009.
- PASS: `DRMCP-ADR-MCP-008` depends on `DRMCP-ADR-MCP-007` and `DRMCP-ADR-MCP-009`.
- PASS: Each new ADR uses the exact T05 decision partition.
- PASS: No Specification, Work Item graph, review, finding, closure, implementation, or test artifact changed.
- PASS: D-002, D-012, D-013, and D-016 were not authored as ADR decisions.
- PASS: The Guidance physical source-path correction does not appear in the ADR set.
- PASS: Scoped Git diff inspection was complete and scoped whitespace inspection reported no findings.

## Evidence

- Startup: Read repository-root `prompt_chappy.md` before every other repository file.
- Startup exclusion: Did not read `CLAUDE.md` or `AGENTS.md`.
- Access mode: DRMCP is non-operational, so all Design Record reads and writes used filesystem fallback.
- Created: `DRMCP-ADR-MCP-007` for D-001, D-003, D-005, and D-006.
- Created: `DRMCP-ADR-MCP-009` for D-004, D-009, and D-017.
- Created: `DRMCP-ADR-MCP-008` for D-007, D-008, D-010, D-011, D-014, and D-015.
- Lifecycle: Updated `DRMCP-ADR-MCP-002` from `accepted` to `superseded` without rewriting its historical decision body.
- Supersession reason: D-009 replaces ADR-002's unconditional configured-Legacy loading with operation-specific Legacy capability selection.
- Retained authority: ADR-009 preserves fresh immutable request state, request-end disposal, trustworthy-snapshot failure, composition lifecycle, and immutable runtime configuration.
- Dependency: ADR-008 depends on ADR-007 and ADR-009. ADR-007 and ADR-009 retain only the existing ADR-001 baseline dependency.
- Excluded routing: D-002 delivery sequencing, D-012 proposal/body-cache design, D-013 write-transaction design, and D-016 view topology did not become ADR decisions.
- Excluded authoring: Did not author application-architecture Specifications or correct Guidance Specifications.
- Excluded workflow: Did not perform review, finding authoring or correction, closure synchronization, implementation, tests, stage, or commit.
- Scoped diff: Inspected the four ADR outputs and this Task. No writable-boundary violation was found.
- Whitespace: Scoped Git whitespace inspection returned `pass` with no findings. LF-to-CRLF notices were advisory only.
- Semantic validator: The standalone responsibility validator was unavailable because DRMCP is non-operational and no standalone invocation tool is exposed in this session.
