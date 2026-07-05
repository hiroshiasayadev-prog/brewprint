# DRMCP-TASK-MCP-016-05: Route DRMCP application-architecture decisions to ADRs

- **id**: DRMCP-TASK-MCP-016-05
- **status**: done
- **date**: 2026-07-04
- **work_item**: DRMCP-WORK-MCP-016
- **task_type**: decision
- **estimate**: 0.5d
- **depends_on**:
  - DRMCP-TASK-MCP-016-04
- **outputs**:
  - DRMCP-TASK-MCP-016-05
  - DRMCP-TASK-MCP-016-06
  - DRMCP-WORK-MCP-016

## Goal

Produce the complete ADR-routing result for the accepted DRMCP application-architecture decision set.

## Work

- Read T03 D-001 through D-017 without reopening their selected outcomes.
- Inspect accepted DRMCP ADRs that may cover or conflict with the decision set.
- Assign every T03 decision one routing outcome: `required`, `covered`, `not_required`, or `blocked`.
- Remove authoring sequencing and authoring-detail deferral from the ADR boundary.
- Partition durable decisions into three coherent boundaries:
  - whole-application component model;
  - inward dependency and responsibility ownership;
  - request-scoped record state and application lifecycle.
- Select `create`, `reuse`, or `supersede` for each boundary.
- Name exact ADR IDs and authoring order for T06.
- Route Guidance source-path correction as a Specification-only correction.
- Keep D-012 and D-013 outside current ADR authoring.
- Record affected canonical architecture Specifications for every routed decision.

## Done condition

- Every T03 decision has one resolved ADR-routing outcome.
- Every `required` decision belongs to one coherent ADR boundary.
- Every `covered` decision names one accepted non-superseded ADR.
- Every `not_required` decision records a reason and exact canonical target.
- No decision remains `blocked`.
- The final three ADR boundaries have exact IDs and dispositions.
- Authoring deferral is not an ADR boundary.
- Exact ADR authoring targets and dependencies are sufficient for T06.
- No ADR body or Specification content is authored.

## Verification

- Confirm all D-001 through D-017 appear in the routing Evidence.
- Confirm no durable architecture choice is omitted.
- Confirm Guidance path correction does not become a standalone architecture ADR.
- Confirm D-012 and D-013 do not become current architecture commitments.
- Confirm authoring sequencing is absent from the ADR set.
- Confirm the boundaries avoid one-ADR-per-row fragmentation and an omnibus whole-workflow ADR.
- Confirm T06 can execute without selecting a new disposition, ID, or authoring order.

## Evidence

### Existing ADR assessment

| ADR | assessment | routing effect |
|---|---|---|
| `DRMCP-ADR-MCP-001` | Accepted contract-baseline authority. It does not own the whole-application component graph. | Retain unchanged. Use as an upstream authority where applicable. |
| `DRMCP-ADR-MCP-002` | Accepted request-scoped snapshot and composition-lifecycle choice. It requires all four W011 operations to build Legacy state when configured. D-009 moves Legacy loading to operation-specific capability selection. | Supersede with a whole-application lifecycle ADR. Preserve the retained request-scoped and composition-lifecycle choices. |
| `DRMCP-ADR-MCP-003` | Accepted logical architecture for the four-operation W011 read-runtime slice. It excludes Guidance and other application operations. | Retain as narrower historical authority. Do not reuse it as the whole-application ADR. |
| `DRMCP-ADR-MCP-004` | Accepted internal-state and operation-contract separation for the W011 read-runtime slice. | Retain unchanged. Its detailed contract boundary remains downstream of W016. |
| `DRMCP-ADR-MCP-005` | Accepted validation orchestration for the W011 read-runtime slice. | Retain unchanged. Its validation-stage detail remains downstream of W016. |
| `DRMCP-ADR-MCP-006` | Accepted Go package boundary for the W011 read-runtime slice. | Retain unchanged. Package placement remains outside W016. |

### Final ADR boundaries

| boundary | ADR and disposition | included decisions | bounded question | dependency and authoring order |
|---|---|---|---|---|
| B-01 Whole-application component model | `DRMCP-ADR-MCP-007`, `create` | D-001, D-003, D-005, D-006 | Which stable application components and scope boundaries cover active DRMCP operations without deciding operation-specific internal contracts? | Independent of B-03. Author before B-02. |
| B-02 Inward dependency and responsibility ownership | `DRMCP-ADR-MCP-008`, `create` | D-007, D-008, D-010, D-011, D-014, D-015 | Which component owns each policy, orchestration, I/O, semantic result, and failure boundary under inward dependencies? | Depend on `DRMCP-ADR-MCP-007` and `DRMCP-ADR-MCP-009`. Author last. |
| B-03 Request-scoped record state and application lifecycle | `DRMCP-ADR-MCP-009`, `create`; supersede `DRMCP-ADR-MCP-002` | D-004, D-009, D-017 | How are Current and Legacy state, configuration, request lifetime, and server lifecycle kept explicit and separate? | Author before B-02. Preserve retained choices from `DRMCP-ADR-MCP-002`. |

B-03 requires supersession. `DRMCP-ADR-MCP-002` requires every W011 read or validation invocation to build Legacy state when configured. D-009 assigns Legacy loading to each operation-specific use case. The selected loading policy therefore changes materially.

`DRMCP-ADR-MCP-009` retains the fresh immutable request snapshot, trustworthy-snapshot failure rule, and composition lifecycle. It replaces the unconditional Legacy-loading rule and broadens the authority to the whole application. Guidance remains outside the record snapshot.

`DRMCP-ADR-MCP-003` is not superseded. Its four-operation read-runtime boundary remains compatible with the broader component and dependency architecture. The new ADRs own the whole-application authority.

### Decision routing

| decision | outcome | ADR boundary or reason | canonical target |
|---|---|---|---|
| D-001 | `required` | B-01 / `DRMCP-ADR-MCP-007` create | `spec:drmcp.application_architecture`; `spec:drmcp.application_architecture.application_boundary_and_components` |
| D-002 | `not_required` | Authoring deferral is delivery sequencing. The MCP-to-use-case seam is already owned by D-005 and D-007. | `spec:drmcp.application_architecture.application_boundary_and_components` as a current-scope exclusion only |
| D-003 | `required` | B-01 / `DRMCP-ADR-MCP-007` create | `spec:drmcp.application_architecture.application_boundary_and_components` |
| D-004 | `required` | B-03 / `DRMCP-ADR-MCP-009` retains the request-scoped immutable snapshot choice while superseding its former ADR owner. | `spec:drmcp.application_architecture.runtime_and_state` |
| D-005 | `required` | B-01 / `DRMCP-ADR-MCP-007` create | `spec:drmcp.application_architecture.application_boundary_and_components` |
| D-006 | `required` | B-01 / `DRMCP-ADR-MCP-007` create | `spec:drmcp.application_architecture.application_boundary_and_components`; `spec:drmcp.application_architecture.dependency_and_responsibility` |
| D-007 | `required` | B-02 / `DRMCP-ADR-MCP-008` create | `spec:drmcp.application_architecture.dependency_and_responsibility` |
| D-008 | `required` | B-02 / `DRMCP-ADR-MCP-008` create | `spec:drmcp.application_architecture.dependency_and_responsibility`; `spec:drmcp.application_architecture.runtime_and_state` |
| D-009 | `required` | B-03 / `DRMCP-ADR-MCP-009` create and supersede `DRMCP-ADR-MCP-002` | `spec:drmcp.application_architecture.runtime_and_state` |
| D-010 | `required` | B-02 / `DRMCP-ADR-MCP-008` create | `spec:drmcp.application_architecture.dependency_and_responsibility` |
| D-011 | `required` | B-02 owns Guidance placement, source-port direction, and minimal domain responsibility. The physical source-path correction is excluded from the ADR. | `spec:drmcp.application_architecture.dependency_and_responsibility`; `spec:drmcp.application_architecture.runtime_and_state` |
| D-012 | `not_required` | Deferred authoring-state details and the non-binding responsibility example require no current ADR. | `spec:drmcp.application_architecture.runtime_and_state`; `spec:drmcp.application_architecture.failure_and_evolution` as exclusions only |
| D-013 | `not_required` | Deferred write-transaction and post-write validation details require no current ADR. | `spec:drmcp.application_architecture.runtime_and_state`; `spec:drmcp.application_architecture.failure_and_evolution` as exclusions only |
| D-014 | `required` | B-02 / `DRMCP-ADR-MCP-008` create | `spec:drmcp.application_architecture.failure_and_evolution` |
| D-015 | `required` | B-02 / `DRMCP-ADR-MCP-008` create | `spec:drmcp.application_architecture.failure_and_evolution` |
| D-016 | `not_required` | The four-view partition and ADR routing are canonical-authoring and workflow structure, not a durable architecture trade-off. | `spec:drmcp.application_architecture` and its four child Specifications; this routing Task |
| D-017 | `required` | B-03 / `DRMCP-ADR-MCP-009` retains startup configuration, construction, wiring, startup, and shutdown under the replacement lifecycle authority. | `spec:drmcp.application_architecture.runtime_and_state` |

No decision is `blocked`.

### Non-ADR projections

- T08 corrects the Guidance source path in:
  - `spec:drmcp.design_records_mcp.schema.authoring_guidance_source`;
  - `spec:drmcp.design_records_mcp.tools.list_authoring_guides`;
  - `spec:drmcp.design_records_mcp.tools.get_authoring_guidance`.
- T07 projects the four-view topology selected by D-016.
- D-012 and D-013 remain deferred exclusions. Their non-binding responsibility example does not enter an ADR.
- The decision to implement read capabilities before authoring capabilities is delivery sequencing. It does not enter an ADR.

### Authoring handoff

T06 uses one writer and one completion judgment for this exact order:

1. Create `DRMCP-ADR-MCP-007` for B-01.
2. Create `DRMCP-ADR-MCP-009` for B-03 and supersede `DRMCP-ADR-MCP-002`.
3. Create `DRMCP-ADR-MCP-008` for B-02 with dependencies on `DRMCP-ADR-MCP-007` and `DRMCP-ADR-MCP-009`.

- Repository-root `prompt_chappy.md` was read before other repository files in this session.
- `CLAUDE.md` and `AGENTS.md` were not read.
- DRMCP is non-operational. All Design Record reads and writes used filesystem fallback.
- No ADR body, Specification body, Guidance correction, implementation, review verdict, finding repair, lifecycle closure, stage, or commit was performed.
- The standalone semantic responsibility validator was not invoked because current DRMCP does not integrate it and no operational invocation tool is available in this session.
