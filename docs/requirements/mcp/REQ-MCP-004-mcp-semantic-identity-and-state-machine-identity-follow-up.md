# REQ-MCP-004: MCP semantic identity and state-machine identity follow-up

- **id**: REQ-MCP-004
- **status**: captured
- **date**: 2026-05-31
- **source_refs**:
  - INV-DATA-001
  - INV-DATA-002
  - REQ-DATA-001
  - WORK-DATA-001
  - REQ-RESOLVE-001
  - WORK-RESOLVE-001
  - ADR-078
  - ADR-079
  - ADR-080
- **work_items**:
  - WORK-MCP-004

## Requirement

The MCP semantic identity follow-up from ADR-078 / ADR-079 / ADR-080 needs to be captured outside the M15 data-layer release and outside the completed resolver file-private bug fix.

The project needs an explicit workflow artifact for semantic-anchor synthetic IDs, transition ID policy, and the state-machine semantic object / file-path-free scenario reference direction, including a split decision before implementation.

## Evidence

`REQ-DATA-001` / `WORK-DATA-001` explicitly excluded ADR-078 / ADR-079 / ADR-080 from the M15 / `v1.1.0-spec` boundary. `INV-DATA-001` and `INV-DATA-002` classify these issues as MCP semantic identity / state-machine identity rather than data-layer expressiveness.

`REQ-RESOLVE-001` / `WORK-RESOLVE-001` resolved the resolver-side file-private sub node identity bug. This requirement does not reopen that bugfix. It captures the remaining MCP public identity and state-machine semantic object policy work.

## Required Outcome

- The ADR-078 / ADR-079 / ADR-080 follow-up is split into an implementable MCP / state identity boundary.
- The work distinguishes accepted policy, proposed policy, and implementation migration work.
- MCP private object exposure / ObjectRef schema migration is either included in a clearly bounded sub-flow or split into a separate requirement / work item.
- Existing closed resolver and M15 scopes remain closed.

## Explicitly Excluded Scope

- The resolver file-private sub node identity bug already resolved by `REQ-RESOLVE-001`.
- M15 / `v1.1.0-spec` reopening.
- Data-layer helper model or tagged union migration.
- Implementing all ADR-078 / 079 / 080 consequences as one undifferentiated task.

## Boundary

This requirement captures a follow-up need and the split decision required before implementation. It does not by itself accept proposed ADRs, finalize transition ID syntax, update MCP schema, or migrate fixtures. Those belong to the linked work item and future tasks.

