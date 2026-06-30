# DRMCP-ADR-MCP-002: Request-scoped read-runtime snapshot and lifecycle

- **status**: accepted
- **date**: 2026-06-30
- **depends_on**:
  - DRMCP-ADR-MCP-001
- **supersedes**: []
- **migrated_to_spec**: 2026-06-30

## Context

`DRMCP-REQ-MCP-001` requires current filesystem state, configured current roots, optional configured legacy roots, and truthful failure behavior.

`DRMCP-ADR-MCP-001` already requires separate current and legacy indexes.
It does not decide how long those indexes live or when DRMCP rebuilds them.

A process-wide mutable index would require cache invalidation, watcher behavior, incremental update rules, and cross-request consistency ownership.
A startup-only index would not reflect filesystem fallback writes or external edits until restart.

## Decision

The composition root loads and validates runtime configuration when the server starts.
Invalid startup configuration prevents server start.

Each invocation of `list_records`, `get_records`, `resolve_reference`, or `validate_records` builds one fresh read snapshot:

- enumerate every configured current root;
- read and parse current sources;
- build one current active index;
- build one separate legacy lookup map when legacy roots are configured;
- freeze both structures for the duration of the operation invocation.

The snapshot is immutable within one invocation of those four operations.
Filesystem changes become visible on the next invocation of one of those operations.

This lifecycle does not define snapshot behavior for authoring-guidance or authoring-transaction operations.
Their internal architecture remains with their owning Requirement or Work Item.

DRMCP does not introduce:

- a process-wide mutable record index;
- incremental index patching;
- filesystem watchers;
- cache invalidation protocols;
- stale-snapshot reuse across invocations.

A post-start root, scan, read, or index failure fails only the current operation.
DRMCP does not return a partial normal response from an untrustworthy snapshot.

The composition root owns construction, startup, shutdown, and wiring.
It does not own parsing, indexing, resolution, validation, or operation semantics.

## Rationale

Per-invocation rebuilding gives each of the four W011 read-runtime operations a consistent filesystem view without introducing cache coordination.

An immutable snapshot prevents one operation from observing mixed index generations.
The next-invocation freshness rule also makes filesystem fallback writes observable without a separate invalidation channel.

Startup configuration failures and later filesystem failures have different owners.
Failing startup for invalid configuration avoids an unusable server.
Failing only the affected operation preserves availability after runtime filesystem changes.

## Rejected alternatives

| alternative | rejection reason |
|---|---|
| Build indexes once at startup | External edits and filesystem fallback writes would remain stale until restart. |
| Keep one process-wide mutable index | The design would require mutation ordering, invalidation, watcher, and concurrency contracts outside the accepted scope. |
| Incrementally patch indexes after writes | External edits would still require a second freshness mechanism. |
| Reuse a previous snapshot after scan failure | The response would conceal that current filesystem state could not be trusted. |
| Continue with valid roots when one mandatory root fails | Partial index operation would contradict the accepted mandatory-root contract. |

## Consequences

- Read latency includes source enumeration, parsing, and index construction for each invocation of the four W011 operations.
- Filesystem freshness is deterministic at those operation boundaries.
- One such invocation cannot observe a mid-operation filesystem change through its snapshot.
- The current active index and legacy lookup map remain separate under `DRMCP-ADR-MCP-001`.
- Any future integration that validates after persisted filesystem writes must discard candidate and pre-write state and rebuild from persisted files.
- Current authoring-transaction integration is deferred to `DRMCP-REQ-MCP-002`.
- Future performance work must preserve the accepted freshness and consistency semantics or supersede this ADR.

Affected Specification targets:

- `spec:drmcp.design_records_mcp.overview`;
- `spec:drmcp.design_records_mcp.namespace_scanning`;
- operation contracts that define execution failure behavior.

## Evidence

- Source Requirement: `DRMCP-REQ-MCP-001`.
- Source Work Item: `DRMCP-WORK-MCP-011`.
- Accepted decisions: D-002 and the request-snapshot use of D-005 in `DRMCP-TASK-MCP-011-01`.
- Existing separation authority: `DRMCP-ADR-MCP-001`.
