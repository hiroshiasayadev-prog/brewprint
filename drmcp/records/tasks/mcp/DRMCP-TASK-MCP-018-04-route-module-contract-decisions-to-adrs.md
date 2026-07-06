# DRMCP-TASK-MCP-018-04: Route module-contract decisions to ADRs

- **id**: DRMCP-TASK-MCP-018-04
- **status**: done
- **date**: 2026-07-06
- **work_item**: DRMCP-WORK-MCP-018
- **task_type**: coordination
- **estimate**: 0.5d
- **depends_on**:
  - DRMCP-TASK-MCP-018-03
- **outputs**:
  - DRMCP-TASK-MCP-018-04

## Goal

Resolve the ADR route for the terminal W018 module-contract decisions without authoring ADR body content.

## Work

- Classify D-001 through D-012 from DRMCP-TASK-MCP-018-02.
- Partition required durable decisions into coherent ADR boundaries.
- Identify create, amend, reuse, or supersede disposition.
- Identify affected Specification targets.
- Record the authoring handoff for DRMCP-TASK-MCP-018-05.

## Done condition

- Every terminal W018 decision has one routing outcome.
- Every required decision belongs to one ADR boundary.
- Every not-required decision has a reason and target.
- The ADR authoring Task has exact routed input.
- No ADR body content is written by this Task.

## Verification

- Confirmed D-001 through D-010 are routed to ADR boundary B-01.
- Confirmed B-01 disposition is create.
- Confirmed the new ADR target is DRMCP-ADR-MCP-013.
- Confirmed D-012 is not ADR-required and is projected to downstream route text.
- Confirmed no existing accepted ADR fully covers the W018 module-contract ownership route.

## Evidence

### Decision routing

| decision | outcome | boundary | disposition | target | reason |
|---|---|---|---|---|---|
| D-001 | required | B-01 | create | DRMCP-ADR-MCP-013 | Component-first derivation constrains module ownership. |
| D-002 | required | B-01 | create | DRMCP-ADR-MCP-013 | Active use-case components constrain downstream contracts. |
| D-003 | required | B-01 | create | DRMCP-ADR-MCP-013 | Shared component decomposition selects durable collaboration owners. |
| D-004 | required | B-01 | create | DRMCP-ADR-MCP-013 | Contract partition topology affects Specification ownership. |
| D-005 | required | B-01 | create | DRMCP-ADR-MCP-013 | Collaboration boundaries constrain allowed dependencies. |
| D-006 | required | B-01 | create | DRMCP-ADR-MCP-013 | Handoff surface inventory constrains downstream type contracts. |
| D-007 | required | B-01 | create | DRMCP-ADR-MCP-013 | Request-state visibility and retention constrain lifecycle. |
| D-008 | required | B-01 | create | DRMCP-ADR-MCP-013 | Failure ownership constrains results and error propagation. |
| D-009 | required | B-01 | create | DRMCP-ADR-MCP-013 | Obligation shape constrains later contract authoring. |
| D-010 | required | B-01 | create | DRMCP-ADR-MCP-013 | `implementation/contracts` topology constrains canonical placement. |
| D-011 | covered | B-01 | create | DRMCP-ADR-MCP-013 | D-011 is the routing decision that creates B-01. |
| D-012 | not_required | none | none | DRMCP-WORK-MCP-018 and downstream graph | Handoff sequencing needs workflow projection, not a separate durable ADR. |

### ADR boundary B-01

| field | value |
|---|---|
| Included decisions | D-001 through D-011. |
| Bounded question | How DRMCP derives module contracts from the accepted application architecture. |
| Disposition | Create one new ADR. |
| New ADR | DRMCP-ADR-MCP-013. |
| Depends on | DRMCP-ADR-MCP-010, DRMCP-ADR-MCP-011, DRMCP-ADR-MCP-012. |
| Affected Specifications | `spec:drmcp.implementation`, `spec:drmcp.implementation.contracts`, and child module-contract specs. |
| Authoring owner | DRMCP-TASK-MCP-018-05. |

### Existing ADR coverage

DRMCP-ADR-MCP-010, DRMCP-ADR-MCP-011, and DRMCP-ADR-MCP-012 establish the whole-application architecture.
They do not select the W018 contract-level component model, module-contract topology, or handoff type/protocol inventory.

### Excluded route

Application-architecture placement migration remains outside W018.
Changes that add top-level architecture components, reverse dependency direction, merge Current and Legacy state, introduce request-spanning mutable state, or change trustworthy-result semantics return to architecture design.

DRMCP is non-operational. Filesystem authoring was used as the required fallback.
