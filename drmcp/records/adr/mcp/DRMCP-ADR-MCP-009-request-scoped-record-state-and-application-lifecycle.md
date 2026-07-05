# DRMCP-ADR-MCP-009: Request-scoped record state and application lifecycle

- **status**: superseded
- **date**: 2026-07-04
- **depends_on**:
  - DRMCP-ADR-MCP-001
- **supersedes**:
  - DRMCP-ADR-MCP-002
- **migrated_to_spec**: null

## Context

`DRMCP-ADR-MCP-002` established fresh immutable snapshots and composition-owned lifecycle for four read-runtime operations.
It also required every operation to build Legacy state whenever legacy roots were configured.

The whole-application architecture distinguishes Current Records, Legacy Archive, and Guidance as different source capabilities.
Their state has different semantics, consumers, and lifetimes.

A replacement lifecycle authority must retain request freshness and trustworthy-failure behavior.
It must remove unconditional Legacy loading and cover the whole application.

## Decision

Composition / Lifecycle owns runtime configuration, construction, wiring, server startup, and ordered shutdown.
Runtime configuration remains immutable for the server lifetime.

Each Read or Validation request builds one fresh immutable Current Records snapshot.
The request uses the snapshot from start to finish and discards it afterward.

The Current Records snapshot provides:

- the active logical tree and index;
- current record identity and source state;
- validation subjects and relation inputs.

Current Records and Legacy Archive use separate source capabilities, separate source ports, and separate logical state.
They are not merged into one repository abstraction or generic index.

Legacy Archive provides exact issued-ID lookup state only.
Each operation-specific use case decides whether Legacy state is required.
Shared application orchestration may load Legacy state through its distinct source port when required.

Guidance operations use their own source capability.
They do not build or consume a record snapshot.

DRMCP does not introduce:

- a process-wide mutable record index;
- a shared record cache across requests;
- filesystem watchers;
- background refresh;
- incremental index patching;
- stale-snapshot reuse after a source failure.

Filesystem changes become visible to the next request that builds the affected state.

An operation fails when mandatory source state cannot produce a complete trustworthy result.
DRMCP does not return a partial normal response from an incomplete or untrustworthy snapshot.

This ADR supersedes the unconditional Legacy-loading rule in `DRMCP-ADR-MCP-002`.
It preserves that ADR's fresh snapshot, request immutability, failure, and composition-lifecycle decisions.

## Rationale

Fresh request state makes external filesystem changes visible without cache invalidation or watcher contracts.
Immutability prevents one operation from observing mixed source generations.

Separate Current and Legacy capabilities preserve their different semantics.
Current Records define active state and validation subjects.
Legacy Archive remains a read-only compatibility lookup.

Operation-specific Legacy loading avoids constructing state that a use case does not need.
The rule also prevents Guidance from inheriting record-snapshot cost or lifecycle.

Composition-owned lifecycle keeps configuration and concrete resource construction outside active use cases.

## Rejected alternatives

| alternative | rejection reason |
|---|---|
| Build record indexes once at startup | External edits and filesystem fallback writes would remain stale until restart. |
| Keep one process-wide mutable index or shared cache | Mutation, invalidation, watcher, and concurrency contracts would become application requirements. |
| Merge Current Records and Legacy Archive into one repository or index | Legacy compatibility state would leak into active listing, validation, and identity semantics. |
| Load Legacy state for every record operation whenever configured | Legacy capability needs differ by operation and must remain operation-specific. |
| Route Guidance through the record snapshot | Guidance has a separate source capability and no record-state dependency. |
| Reuse stale state after mandatory source failure | A normal response would conceal that the promised result is not trustworthy. |

## Consequences

- Read and Validation requests pay the cost of fresh Current Records state construction.
- Operations that require Legacy lookup pay the additional cost of separate Legacy state construction.
- Operations that do not require Legacy state do not load it.
- Guidance requests avoid record snapshot construction.
- Current and Legacy source ports remain distinct in downstream contracts.
- Runtime configuration changes require a new server lifetime.
- Performance refinements must preserve request freshness, immutable request state, and source separation.
- Request-spanning mutable record state requires a new application-architecture decision.

Affected design areas:

- `spec:drmcp.application_architecture.runtime_and_state`;
- `spec:drmcp.application_architecture.failure_and_evolution`;
- operation contracts that select Current or Legacy source capabilities.

## Evidence

- Source Requirement: `DRMCP-REQ-MCP-005`.
- Source Work Item: `DRMCP-WORK-MCP-016`.
- Routed decisions: D-004, D-009, and D-017 in `DRMCP-TASK-MCP-016-03`.
- ADR routing authority: B-03 in `DRMCP-TASK-MCP-016-05`.
- Superseded lifecycle authority: `DRMCP-ADR-MCP-002`.
- Contract baseline authority: `DRMCP-ADR-MCP-001`.
