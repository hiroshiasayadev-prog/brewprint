# DRMCP-TASK-MCP-011-02: Route rebuild read-runtime decisions to ADRs

- **id**: DRMCP-TASK-MCP-011-02
- **status**: done
- **date**: 2026-06-30
- **work_item**: DRMCP-WORK-MCP-011
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 1d
- **depends_on**:
  - DRMCP-TASK-MCP-011-01
- **outputs**:
  - DRMCP-TASK-MCP-011-02
  - DRMCP-WORK-MCP-011
  - DRMCP-ADR-MCP-002
  - DRMCP-ADR-MCP-003
  - DRMCP-ADR-MCP-004
  - DRMCP-ADR-MCP-005
  - DRMCP-ADR-MCP-006

## Goal

Route the accepted D-001 through D-009 architecture decisions to durable ADR authority.

Preserve existing ADR authority without duplicating or rewriting accepted history.

## Work

- Read `DRMCP-ADR-MCP-001` before assigning existing coverage.
- Classify every accepted decision as covered, amended, superseded, or requiring a new ADR.
- Group decisions only when they share one coherent decision boundary.
- Keep independently changeable runtime, contract, validation, and package choices in separate ADRs.
- Author every required ADR with explicit rationale, alternatives, consequences, source Requirement, and affected Specification targets.
- Record the final routing matrix in this Task.
- Synchronize the new Task and Phase B completion into `DRMCP-WORK-MCP-011`.
- Stop before Specification synchronization, independent review, or implementation planning.

## Done condition

- D-001 through D-009 each have one resolved routing outcome.
- Every `new ADR required` decision points to an authored ADR.
- Every `covered by existing ADR` decision points to an accepted, non-superseded ADR with sufficient scope.
- No existing ADR is amended or superseded without a material authority reason.
- New ADRs use the next available MCP-domain ADR sequence.
- New ADRs identify `DRMCP-REQ-MCP-001` and affected Specification targets.
- `DRMCP-WORK-MCP-011.tasks` lists this Task.
- `DRMCP-WORK-MCP-011` records Phase B completion and Phase C readiness.
- No Specification, production code, implementation Task, or independent-review Task changes.
- Scoped Git diff and whitespace inspection pass.

## Verification

- Compare the routing matrix with the D-001 through D-009 summaries in `DRMCP-TASK-MCP-011-01`.
- Compare D-005 with the complete decision and consequences in `DRMCP-ADR-MCP-001`.
- Confirm the MCP-domain ADR inventory contains only `DRMCP-ADR-MCP-001` before allocating 002 through 006.
- Confirm no new ADR overturns `DRMCP-ADR-MCP-001`.
- Confirm each new ADR has required metadata and canonical sections.
- Confirm the changed-file boundary contains this Task, `DRMCP-WORK-MCP-011`, and ADRs 002 through 006 only.
- Run scoped Git worktree and textual-diff inspection with whitespace checking.

## Evidence

Phase B ADR routing and authoring is complete.

The MCP-domain ADR inventory contained only `DRMCP-ADR-MCP-001` before this Task allocated ADRs 002 through 006.

### Authored ADRs

| ADR | decision boundary | decision IDs |
|---|---|---|
| `DRMCP-ADR-MCP-002` | Request-scoped snapshot and runtime lifecycle. | D-002; consumes the D-005 separation authority. |
| `DRMCP-ADR-MCP-003` | Layered application, use-case, port, and adapter boundaries. | D-001, D-004, D-007 |
| `DRMCP-ADR-MCP-004` | Internal state and operation-contract separation. | D-003, D-008 |
| `DRMCP-ADR-MCP-005` | Validation orchestration over fresh snapshots. | D-006 |
| `DRMCP-ADR-MCP-006` | Concrete Go package boundaries. | D-009 |

Every new ADR is `accepted`, uses date `2026-06-30`, identifies `DRMCP-REQ-MCP-001`, and lists affected Specification targets.
All new ADRs have `migrated_to_spec: null` because Phase C has not begun.

### ADR routing matrix

| decision | routing outcome | ADR | reason | affected Specification targets |
|---|---|---|---|---|
| D-001 | new ADR required | `DRMCP-ADR-MCP-003` | Dedicated use-case dispatch and the MCP adapter boundary are part of the layered application decision. | `spec:drmcp.design_records_mcp.responsibility_boundary`, `spec:drmcp.design_records_mcp.tools.overview` |
| D-002 | new ADR required | `DRMCP-ADR-MCP-002` | Per-invocation rebuilding, immutable snapshots, and startup versus operation failure ownership are a durable runtime lifecycle choice. | `spec:drmcp.design_records_mcp.overview`, `spec:drmcp.design_records_mcp.namespace_scanning` |
| D-003 | new ADR required | `DRMCP-ADR-MCP-004` | Internal state types and public operation projections require separate ownership from the component topology. | `spec:drmcp.design_records_mcp.schema.overview`, `spec:drmcp.design_records_mcp.schema.record_model` |
| D-004 | new ADR required | `DRMCP-ADR-MCP-003` | Narrow application-owned ports and outer adapters are inseparable from the accepted dependency direction. | `spec:drmcp.design_records_mcp.namespace_scanning`, `spec:drmcp.design_records_mcp.responsibility_boundary` |
| D-005 | covered by existing ADR | `DRMCP-ADR-MCP-001` | The accepted ADR already requires separate current and legacy indexes, active-only listing, exact legacy retrieval, and current-first fallback. | `spec:drmcp.design_records_mcp.namespace_scanning`, `spec:drmcp.design_records_mcp.resolver` |
| D-006 | new ADR required | `DRMCP-ADR-MCP-005` | Validation pass ordering, pure validators, snapshot ownership, and post-write reload form one durable orchestration choice. | `spec:drmcp.design_records_mcp.tools.validate_records`, `spec:drmcp.design_records_mcp.schema.diagnostics` |
| D-007 | new ADR required | `DRMCP-ADR-MCP-003` | The component model, inward dependency direction, and orchestration ownership are one coherent architecture boundary. | `spec:drmcp.design_records_mcp.overview`, `spec:drmcp.design_records_mcp.responsibility_boundary` |
| D-008 | new ADR required | `DRMCP-ADR-MCP-004` | Typed use-case contracts, result projection, expected semantic states, and execution errors share one operation-contract boundary. | `spec:drmcp.design_records_mcp.tools.overview`, operation-specific tool specs, `spec:drmcp.design_records_mcp.schema.diagnostics` |
| D-009 | new ADR required | `DRMCP-ADR-MCP-006` | Concrete Go package placement can change independently from the logical component architecture. | `spec:drmcp.design_records_mcp.overview`, `spec:drmcp.design_records_mcp.responsibility_boundary` |

### Existing ADR disposition

| ADR | disposition | scope assessment |
|---|---|---|
| `DRMCP-ADR-MCP-001` | covered; no amendment | Fully covers D-005. It does not define D-001 through D-004 or D-006 through D-009 at sufficient architecture detail. |

No existing ADR requires amendment or supersession.

### Changed-file boundary

Phase B changed only:

- this Task;
- `DRMCP-WORK-MCP-011`;
- `DRMCP-ADR-MCP-002` through `DRMCP-ADR-MCP-006`.

No Specification, production code, implementation Task, or independent-review Task changed.

### Verification result

- Scoped `git.inspect_diff` returned `pass` and included all seven Phase B files.
- Scoped `git.inspect_worktree` returned `pass`.
- Whitespace status was `pass` with no findings across all seven untracked files.
- LF-to-CRLF conversion messages were non-blocking working-copy warnings.
- Repository-wide cleanliness was not checked or inferred.
- This closure update changes the checked Task bytes. Final post-closure scoped Git inspection is supplied externally and is not written back into this file.
