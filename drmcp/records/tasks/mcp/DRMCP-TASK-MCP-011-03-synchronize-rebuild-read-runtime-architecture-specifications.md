# DRMCP-TASK-MCP-011-03: Synchronize rebuild read-runtime architecture Specifications

- **id**: DRMCP-TASK-MCP-011-03
- **status**: done
- **date**: 2026-06-30
- **work_item**: DRMCP-WORK-MCP-011
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 2d
- **depends_on**:
  - DRMCP-TASK-MCP-011-02
- **outputs**:
  - DRMCP-TASK-MCP-011-03
  - DRMCP-WORK-MCP-011
  - DRMCP-ADR-MCP-002
  - DRMCP-ADR-MCP-003
  - DRMCP-ADR-MCP-004
  - DRMCP-ADR-MCP-005
  - DRMCP-ADR-MCP-006
  - spec:drmcp.implementation
  - spec:drmcp.design_records_mcp.overview
  - spec:drmcp.design_records_mcp.responsibility_boundary
  - spec:drmcp.design_records_mcp.namespace_scanning
  - spec:drmcp.design_records_mcp.schema.overview
  - spec:drmcp.design_records_mcp.schema.record_model
  - spec:drmcp.design_records_mcp.schema.diagnostics
  - spec:drmcp.design_records_mcp.tools.overview
  - spec:drmcp.design_records_mcp.tools.validate_records
  - spec:drmcp.design_records_mcp.tools.authoring_transaction_model
  - spec:drmcp.design_records_mcp.tools.accept_proposed_write

## Goal

Record the accepted rebuild read-runtime architecture as current DRMCP Specification authority.

Keep public read-operation behavior unchanged while separating implementation architecture from semantic and interface contracts.

## Work

- Create one top-level implementation Specification for the accepted runtime architecture.
- Reflect D-001 through D-009 from ADR-001 through ADR-006 without copying ADR rationale.
- Keep the complete implementation architecture in `spec:drmcp.implementation`.
- Add narrow ownership and relationship pointers to existing DRMCP Specifications.
- Record request-scoped freshness where namespace and validation contracts depend on it.
- Record internal state separation where the record-model contract depends on it.
- Record use-case and MCP-adapter ownership where tool contracts depend on it.
- Record validation pass order and pure-detector boundaries without changing validation outcomes.
- Record post-write fresh-snapshot validation without changing authoring transaction behavior.
- Update `migrated_to_spec` for ADR-002 through ADR-006.
- Recheck operation-specific read contracts and resolver contracts for direct contradictions.
- Stop before independent review, finding correction, lifecycle closure, or production implementation.

## Done condition

- `spec:drmcp.implementation` exists with the accepted runtime lifecycle, layers, state boundaries, ports, validation flow, result ownership, and Go package layout.
- Existing Specifications point to one implementation architecture authority instead of duplicating it.
- Each changed existing Specification retains its previous public behavior.
- D-005 remains consistent with `DRMCP-ADR-MCP-001`, namespace scanning, and resolver contracts.
- ADR-002 through ADR-006 record Specification migration on 2026-06-30.
- W011 lists this Task and records Phase C completion.
- No production source, fixture, Requirement, implementation Task, review Task, or downstream lifecycle changes.
- Scoped textual diff and whitespace inspection pass.

## Verification

- Trace D-001 through D-009 from T01 to ADR-001 through ADR-006 and then to changed Specifications.
- Confirm `spec:drmcp.implementation` is path-derived from `drmcp/records/spec/implementation/index.md`.
- Confirm the new Specification has valid Overview metadata and required sections.
- Confirm existing operation request, response, status, warning, diagnostic, and error meanings remain unchanged.
- Confirm current and legacy indexes remain distinct.
- Confirm no generic App Router, merged index, all-purpose repository interface, or retired package compatibility wrapper is introduced.
- Confirm validation detectors perform no filesystem I/O and post-write validation rebuilds persisted state.
- Confirm all modified spec dates are `2026-06-30`.
- Run scoped Git worktree and textual-diff inspection with whitespace checking.

## Evidence

Phase C Specification synchronization is complete.

DRMCP authoring transactions remain non-operational. Filesystem authoring was used under the current agent-authoring policy.

### Canonical architecture Specification

Created `spec:drmcp.implementation` at `drmcp/records/spec/implementation/index.md`.

The Specification owns:

- startup configuration and request-scoped runtime lifecycle;
- one fresh immutable snapshot per MCP invocation;
- separate current active-index and optional legacy lookup state;
- composition-root, MCP-adapter, application, core, and outbound-adapter boundaries;
- narrow configuration, source-enumeration, and source-reading ports;
- internal source, parsed, index, conflict, finding, and operation-projection separation;
- expected-result versus execution-error ownership;
- validation pass ordering and pure-detector boundaries;
- fresh post-write validation against persisted files;
- concrete Go package layout under `drmcp/src/`;
- rejection of a generic App Router, merged index, all-purpose repository, and retired package compatibility wrapper.

### Decision-to-Specification trace

| decision | ADR authority | Specification reflection |
|---|---|---|
| D-001 | `DRMCP-ADR-MCP-003` | `spec:drmcp.implementation`, responsibility boundary, tools overview |
| D-002 | `DRMCP-ADR-MCP-002` | `spec:drmcp.implementation`, Design Records MCP overview, namespace scanning |
| D-003 | `DRMCP-ADR-MCP-004` | `spec:drmcp.implementation`, schema overview, record model |
| D-004 | `DRMCP-ADR-MCP-003` | `spec:drmcp.implementation`, responsibility boundary, namespace scanning |
| D-005 | `DRMCP-ADR-MCP-001` | Existing namespace-scanning and resolver contracts remain authoritative; no resolver behavior change was required. |
| D-006 | `DRMCP-ADR-MCP-005` | `spec:drmcp.implementation`, validation contract, diagnostics, authoring post-write validation |
| D-007 | `DRMCP-ADR-MCP-003` | `spec:drmcp.implementation`, Design Records MCP overview, responsibility boundary |
| D-008 | `DRMCP-ADR-MCP-004` | `spec:drmcp.implementation`, tools overview, diagnostics |
| D-009 | `DRMCP-ADR-MCP-006` | `spec:drmcp.implementation` |

### Changed Specifications

Created:

- `spec:drmcp.implementation`.

Updated:

- `spec:drmcp.design_records_mcp.overview`;
- `spec:drmcp.design_records_mcp.responsibility_boundary`;
- `spec:drmcp.design_records_mcp.namespace_scanning`;
- `spec:drmcp.design_records_mcp.schema.overview`;
- `spec:drmcp.design_records_mcp.schema.record_model`;
- `spec:drmcp.design_records_mcp.schema.diagnostics`;
- `spec:drmcp.design_records_mcp.tools.overview`;
- `spec:drmcp.design_records_mcp.tools.validate_records`;
- `spec:drmcp.design_records_mcp.tools.authoring_transaction_model`;
- `spec:drmcp.design_records_mcp.tools.accept_proposed_write`.

Every substantively changed Specification uses date `2026-06-30`.

### Rechecked without change

The following contracts were rechecked and required no Phase C edit:

- `spec:drmcp.design_records_mcp.resolver`;
- `spec:drmcp.design_records_mcp.mvp_scope`;
- `spec:drmcp.design_records_mcp.tools.list_records`;
- `spec:drmcp.design_records_mcp.tools.get_records`;
- `spec:drmcp.design_records_mcp.tools.resolve_reference`.

Their existing request, response, status, warning, diagnostic, and error meanings remain unchanged.

### ADR synchronization

`DRMCP-ADR-MCP-002` through `DRMCP-ADR-MCP-006` now use:

```text
migrated_to_spec: 2026-06-30
```

ADR decision dates and decision meanings were not changed.

### Changed-file boundary

Phase C changed only:

- this Task;
- `DRMCP-WORK-MCP-011`;
- ADR-002 through ADR-006 migration metadata;
- the new implementation Specification;
- the ten existing Specifications listed above.

No Requirement, fixture, production source, implementation Task, review Task, or downstream lifecycle record changed.

### Verification result

- Manual trace confirmed D-001 through D-009 have ADR and Specification coverage.
- The path `drmcp/records/spec/implementation/index.md` derives `spec:drmcp.implementation`.
- The new Specification uses valid Overview metadata and required sections.
- Scoped `git.inspect_diff` returned `pass` across all 18 Phase C files.
- Scoped `git.inspect_worktree` returned `pass`.
- Whitespace status was `pass` with no findings.
- LF-to-CRLF conversion messages were non-blocking working-copy warnings.
- Repository-wide cleanliness was not checked or inferred.
- Repository-local strict Specification validation was not run because this session has no arbitrary command-execution boundary for the Windows repository.
- This closure synchronization changes the checked Task and Work Item bytes. Final post-closure scoped Git inspection is supplied externally and is not written back into these files.
