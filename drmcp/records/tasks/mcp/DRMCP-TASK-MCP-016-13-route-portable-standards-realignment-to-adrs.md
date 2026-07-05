# DRMCP-TASK-MCP-016-13: Route portable standards realignment to ADRs

- **id**: DRMCP-TASK-MCP-016-13
- **status**: done
- **date**: 2026-07-04
- **work_item**: DRMCP-WORK-MCP-016
- **task_type**: decision
- **estimate**: 0.5d
- **depends_on**:
  - DRMCP-TASK-MCP-016-12
- **outputs**:
  - DRMCP-TASK-MCP-016-13
  - DRMCP-TASK-MCP-016-14

## Goal

Resolve ADR routing and supersession for D-018 through D-024.

## Work

- Assess ADR-001 and ADR-007 through ADR-009 against the revised decisions.
- Assign every D-018 through D-024 decision one ADR-routing outcome.
- Partition required durable choices into coherent replacement boundaries.
- Select exact ADR IDs, dependencies, supersession targets, and authoring order.
- Identify non-ADR Requirement and Specification projections.
- Preserve completed T03, T05, and T06 records as historical Evidence.

## Done condition

- Every D-018 through D-024 decision has one resolved routing outcome.
- Every required decision belongs to one coherent ADR boundary.
- ADR-001 has an explicit retained disposition.
- ADR-007 through ADR-009 have explicit supersession dispositions.
- Exact replacement ADR IDs and dependencies are fixed.
- T14 can author every required authority without new judgment.

## Verification

- Confirm no routing item remains `blocked`.
- Confirm no accepted ADR body is rewritten to conceal a reversal.
- Confirm the replacement boundaries avoid one ADR per decision row.
- Confirm Requirement amendment and Specification authoring remain separate from ADR routing.
- Confirm no ADR or Specification body is authored by this Task.

## Evidence

### Existing ADR assessment

| ADR | assessment | disposition |
|---|---|---|
| `DRMCP-ADR-MCP-001` | Portable fixed-namespace standards package and indexed-Spec Guidance projection remain valid. | `reuse` unchanged. |
| `DRMCP-ADR-MCP-007` | The six-component model includes a Guidance Domain that the revised decision removes. | Supersede with `DRMCP-ADR-MCP-010`. |
| `DRMCP-ADR-MCP-008` | Guidance Domain and Guidance Source ownership change materially to normal record-query orchestration. | Supersede with `DRMCP-ADR-MCP-011`. |
| `DRMCP-ADR-MCP-009` | Guidance changes from a separate source outside record snapshots to a normal Current Records source inside the request snapshot. | Supersede with `DRMCP-ADR-MCP-012`. |

### Final ADR boundaries

| boundary | ADR and disposition | included decisions | bounded question | dependency and authoring order |
|---|---|---|---|---|
| B-04 Five-component whole-application model | `DRMCP-ADR-MCP-010`, `create`, superseding ADR-007 | D-018, D-020 | Which top-level components cover the application when portable standards are normal Current Records? | Depends on ADR-001. Author first. |
| B-05 Inward ownership with Guidance query aliases | `DRMCP-ADR-MCP-011`, `create`, superseding ADR-008 | D-021, D-022 | Which layer owns Guidance scope, lookup, projection, and ordering without a Guidance Domain or source port? | Depends on ADR-010 and ADR-012. Author last. |
| B-06 Unified Current Records state and lifecycle | `DRMCP-ADR-MCP-012`, `create`, superseding ADR-009 | D-019, D-023 | How does a configured portable spec-tree source join normal Current Records state and lifecycle? | Depends on ADR-001. Author before ADR-011. |

### Decision routing

| decision | outcome | ADR boundary or reason | canonical targets |
|---|---|---|---|
| D-018 | `required` | B-04 establishes package treatment within the component model. | Requirement, component and runtime Specifications. |
| D-019 | `required` | B-06 establishes the spec-tree source mapping and lifecycle boundary. | Requirement, namespace scanning, discovery, identity, source Specifications. |
| D-020 | `required` | B-04 replaces the six-component model with five components. | Component ADR and architecture Specifications. |
| D-021 | `required` | B-05 replaces Guidance Domain and source-port ownership. | Dependency ADR and Guidance operation Specifications. |
| D-022 | `required` | B-05 fixes durable canonical identity and projection ownership. | Guidance schema and operation Specifications. |
| D-023 | `required` | B-06 changes Guidance request state and package lifecycle materially. | Lifecycle ADR and runtime Specification. |
| D-024 | `not_required` | Artifact disposition is workflow and authoring routing. | DRMCP-REQ-MCP-003, ADR lifecycle metadata, T14 and T15. |

### Authoring handoff

T14 uses one writer and one authority completion judgment in this order:

1. Amend DRMCP-REQ-MCP-003.
2. Create DRMCP-ADR-MCP-010 and mark ADR-007 `superseded`.
3. Create DRMCP-ADR-MCP-012 and mark ADR-009 `superseded`.
4. Create DRMCP-ADR-MCP-011 and mark ADR-008 `superseded`.

T15 follows T14 and authors the revised canonical Specification state.

- The repository-root `prompt_chappy.md` was read first.
- `CLAUDE.md` and `AGENTS.md` were not read.
- DRMCP is non-operational. Design Record access used filesystem fallback.
- No ADR body, Requirement body, Specification body, review verdict, stage, or commit was produced by this Task.
