# Concept: PuppyDSL events

- **id**: `spec:puppydsl.language.events`
- **status**: draft
- **date**: 2026-07-25
- **parent**: `spec:puppydsl.language`

## What this is

Defines declared event identity, typed synchronous dispatch, context invalidation ordering, and bounded complete-function restart.

## Concept model

```text
declared event
  -> generated typed callable
  -> synchronous occurrence recording
  -> matching context invalidation
  -> occurrence publication to restart observers
```

An event is design data with one stable qualified ID and one payload type.

## Event declaration

```yaml
events:
  repository.changed:
    detail: Reports that relevant repository content changed.
    payload: RepositoryChange
```

An event without data declares `payload: void`.

```yaml
events:
  runtime.shutdown:
    detail: Reports that the runtime begins orderly shutdown.
    payload: void
```

`never` is invalid as an event payload.

## Generated callable

Every event creates one callable with the same ID.

| payload | callable contract |
|---|---|
| `T` | `<event-id>(payload: T) -> void` |
| `void` | `<event-id>() -> void` |

The callable has no ordinary native, steps, or stub implementation declaration. The runtime event dispatcher owns execution.

Event callables share the callable namespace with ordinary functions. Identity conflicts are invalid.

## Dispatch

Event dispatch uses ordinary call syntax and ordinary argument type checking.

Dispatch is synchronous:

```text
validate call
  -> record occurrence
  -> invalidate matching context generations
  -> publish completed occurrence to restart observers
  -> return void
```

A caller returning from dispatch may rely on matching current generations already being stale.

Dispatch does not execute arbitrary subscribers, rebuild contexts, run destructors, or perform external I/O.

## Event producers

Runtime infrastructure and explicitly selected operations may call generated event callables.

Steps-defined functions have no subscription syntax. Hidden asynchronous control flow is outside the language.

## Context invalidation

A context lists event IDs through `invalidated_by`.

The event payload is ignored unless a separate typed consumer explicitly needs it. `invalidated_by` contains event IDs only.

Context invalidation semantics are defined by `spec:puppydsl.language.contexts`.

## Complete-function restart

A restartable entry function may declare:

```yaml
restart_on:
  - repository.changed
max_restarts: 1
```

The runner records an event cursor before opening invocation-owned contexts.

After the complete invocation finishes:

| observed state | action |
|---|---|
| No matching event after cursor | Return completed result. |
| Matching event within allowance | Discard result, close invocation contexts, rerun complete invocation. |
| Matching event after final allowance | Discard result and fail the invocation. |

Restart never resumes an internal step and never replaces context values during one active invocation.

## Restart safety

Only selected entry functions may declare `restart_on`.

A restartable function must not directly or transitively dispatch any event. This prevents self-restart loops without requiring idempotence analysis.

`restart_on` requires a finite `max_restarts`. The exact accepted integer domain remains pending and is not fixed by the current language contract.

## Occurrence tracking

The runtime maintains a monotonic event cursor or equivalent occurrence ordering.

Restart asks whether a relevant event occurred after the invocation cursor. Correctness does not depend on one exact physical notification count.

## Rules

- Every event declares exactly one payload type.
- Event IDs are statically resolved.
- Generated event callables are typed and synchronous.
- Context invalidation completes before restart observers see the occurrence.
- `invalidated_by` and `restart_on` list event IDs only.
- Restart applies to one complete entry-function invocation.
- Restart count is finite.
- Restartable functions cannot dispatch events.

## Boundary

| owned here | owned elsewhere |
|---|---|
| Event identity and payload type. | Runtime occurrence storage implementation. |
| Generated callable language contract. | Generated source or native dispatcher API. |
| Synchronous dispatch ordering. | Producer-specific external work before dispatch. |
| Complete-invocation restart semantics. | Host failure-envelope projection. |
| Restart safety constraints. | Diagnostic codes and CLI behavior. |

## Related specs

| ref | relation |
|---|---|
| `spec:puppydsl.language.functions` | Declares restart fields and callable contracts. |
| `spec:puppydsl.language.contexts` | Defines generation invalidation. |
| `spec:puppydsl.language.program_validity.flow_analysis` | Validates event references and restart call graphs. |
