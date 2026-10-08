# Runtime event model

## Status

Working contract for declared runtime events, context invalidation, and bounded operation restart.

## Declaration detail

Every event declaration includes a non-empty `detail` field that states the runtime condition represented by the event.
The field is bootstrap-validated design authority and is not dispatch or consumer logic.

## Event identity

Every runtime event has one stable fully qualified event ID.

Examples:

```text
repository.changed
configuration.changed
artifact_modules.changed
native_extensions.activated
runtime.shutdown
```

An event ID is declared design data rather than an arbitrary string literal.
Bootstrap validation rejects unknown event references.

Event declarations are grouped beneath the `events` root mapping:

```yaml
events:
  repository.changed:
    detail: Reports that repository content relevant to current-record state changed.
    payload: RepositoryChange
```

Every event declares exactly one payload type.
When identity and occurrence are sufficient, the event declares `void`:

```yaml
events:
  runtime.shutdown:
    detail: Reports that the runtime host is beginning orderly shutdown.
    payload: void
```

`void` means that dispatch carries no event data.
It is not `null`, `nil`, or an absent optional value.
A non-void payload type must be a declared runtime type that can exist as a value. `never` is prohibited as an event payload.

## Generated event callable

Every event declaration generates one typed callable with the same stable ID as the event.

```yaml
events:
  repository.changed:
    detail: Reports that repository content relevant to current-record state changed.
    payload: RepositoryChange
```

generates the callable contract:

```text
repository.changed(payload: RepositoryChange) -> void
```

A void-payload event generates a zero-input callable:

```yaml
events:
  runtime.shutdown:
    detail: Reports that the runtime host is beginning orderly shutdown.
    payload: void
```

```text
runtime.shutdown() -> void
```

Design functions dispatch an event by calling this generated callable through the ordinary function-call syntax.
The callable has no YAML-defined implementation body; execution is routed to the runtime event dispatcher.
Event payload presence and type are therefore validated by the ordinary typed-call rules without introducing a variadic `runtime.event.invoke` intrinsic.

Generated event callables participate in call-graph analysis so bootstrap validation can detect direct and transitive event dispatch.
An event ID must not conflict with a declared ordinary function ID.

## Event producers

Runtime infrastructure and explicitly authorized operations may dispatch declared events.
A producer dispatching a non-void event calls the generated event callable with exactly one value of the declared payload type.
A producer dispatching a void event calls the generated event callable with no argument.
Examples include:

- a filesystem watcher dispatching `repository.changed`;
- a successful repository write dispatching `repository.changed` synchronously;
- native extension activation dispatching `native_extensions.activated`;
- the runtime host dispatching `runtime.shutdown`.

Ordinary design functions do not subscribe to or dispatch events directly.
This prevents hidden asynchronous control flow inside steps-defined functions.

## Synchronous dispatch

A generated event callable executes synchronously.
It returns only after the runtime has completed the event's required state transition:

```text
validate the generated callable invocation
  -> assign and record one event occurrence
  -> invalidate every matching current context generation
  -> make the completed occurrence visible to restart observers
  -> return void
```

A caller returning from an event callable may therefore rely on matching context generations already being stale.
This prevents a later operation from acquiring an old generation after a successful repository write has dispatched its change event.

The runtime does not queue an event for deferred invalidation and then return early.
A watcher may call the event callable from its own goroutine, but that goroutine waits until dispatch completes.

Invalidation and occurrence publication are ordered so that a restart observer cannot observe an event whose required context invalidation remains incomplete.
After successful bootstrap, a correctly typed generated event-callable invocation is infallible and returns `void`.
Dispatch performs only in-memory runtime state transitions: occurrence recording, matching context-generation invalidation, and restart-observer publication.
It does not execute arbitrary subscribers, rebuild contexts, run destructors, read external sources, or perform other external I/O.
Failures in a producer before dispatch, in later context reconstruction, or in resource cleanup remain in those operations' existing failure paths.
An impossible dispatcher invariant violation is a runtime defect rather than a YAML-visible recoverable outcome.

A runtime consumer binds the payload only when it needs to inspect that data.
Consumers that react only to event identity declare no payload binding.
Therefore `invalidated_by` and `restart_on` contain event IDs only:

```yaml
invalidated_by:
  - repository.changed

restart_on:
  - repository.changed
```

The absence of a consumer payload binding means that the payload is ignored; it does not change the event declaration or dispatch contract.

## Context invalidation

A long-lived context declares events that invalidate its current cached generation:

```yaml
contexts:
  repository_snapshot:
    detail: Owns one reusable generation of repository-derived snapshot values.
    invalidated_by:
      - repository.changed
    provides:
      current_records:
        provider: runtime.current_records.collect()
```

Event dispatch and context invalidation form one ordered runtime transaction:

```text
record event occurrence
  -> mark every matching current context generation stale
  -> publish the completed occurrence to restart observers
```

A restart observer must not observe an occurrence whose required context invalidation has not completed.
When a matching event occurs, the context manager rejects the stale generation for later acquisitions while preserving it for operations that already acquired it.
A later acquisition constructs a new generation lazily.
The stale generation closes after its final active consumer exits.

Context invalidation does not reopen a nested instance of the same context ID.
Opening a context ID already active in the current chain remains prohibited.

## Operation restart

A restartable function may declare events that require its complete invocation to be rerun:

```yaml
functions:
  operation.retrieval.get_records:
    detail: Retrieves selected records against one stable operation snapshot.
    opens:
      - operation
    restart_on:
      - repository.changed
    max_restarts: 1
    signature: >-
      (selectors: list<RecordSelector>) -> GetRecordsResult
    implementation:
      native: operation.retrieval.get_records
```

The operation runner records an event cursor before opening the function-owned contexts.
The function executes against one stable context generation.
After completion, the runner checks whether any declared restart event occurred after that cursor.

```text
no matching event
  -> return the completed result

matching event occurred
  -> discard the completed result
  -> close the completed invocation contexts
  -> rerun the complete function invocation
```

A restart never resumes from an internal step and never replaces context values during an active invocation.

Restart count is bounded by `max_restarts`.
If another matching event occurs after the final permitted restart, the runtime discards that completed result and reports an `operation_restart_exhausted` execution failure at the operation boundary.

Restart exhaustion uses the runtime execution-failure envelope:

```text
code: operation_restart_exhausted
function: <fully-qualified FunctionID>
message: <non-empty human-readable explanation>
context:
  max_restarts: <declared limit>
  restarts_performed: <actual restart count>
```

`function` identifies the function invocation whose execution failed. For restart exhaustion, this is the restartable entry function whose invocation exhausted its restart allowance, not an internal function that happened to be executing when an event was emitted.
`message` is required for human inspection, but exact wording is not a stable branching contract.
`context` contains failure-specific structured data. For `operation_restart_exhausted`, it contains exactly `max_restarts` and `restarts_performed`, both integers.
This failure is internal runtime data and is not a YAML-visible normal value.

Projection into a consuming operation's external error or diagnostic contract is owned by that operation and adapter layer.
The shared field names do not make the runtime envelope dependent on DRMCP's diagnostic catalog, placement, or response rules.
The runtime does not return the last completed result because doing so would violate the operation's declared restart consistency contract.
Restart exhaustion is not a validation finding, advisory, or expected domain outcome.

## Restart safety

Only restartable entry operations may use `restart_on`.
An operation that declares `restart_on` must not directly or transitively dispatch any event.
This initial prohibition avoids self-restart loops and removes the need to prove event-emitting operations idempotent.

The runtime supports event-driven operation restart as a general mechanism, but an individual DRMCP read operation is not required to use it unless its operation contract explicitly declares `restart_on`.
A fast read may instead complete against its pinned context generation without restart.

Bootstrap validation rejects:

- unknown event IDs in `restart_on`;
- `restart_on` without a finite `max_restarts`;
- event-driven restart on an internal step or ordinary nested function;
- direct or transitive event dispatch by an operation that declares `restart_on`;
- event subscription or dispatch syntax inside a steps-defined implementation.

## Event occurrence tracking

The runtime maintains a monotonic occurrence sequence or equivalent cursor for declared events.
Operation restart asks whether a relevant event occurred after the invocation cursor; it does not depend on a resource-specific field such as `consistency.repository`.

Repeated physical notifications may be debounced or deduplicated by an event producer, but correctness must not depend on one exact notification count.

## Relationship to context lifetime

Events and active context instances have separate responsibilities:

```text
context scope
  -> runtime interval in which one context instance is active
  -> defines cache, validity, and cleanup lifetime

event
  -> reports that a runtime condition occurred

invalidated_by
  -> maps an event to context-generation invalidation

restart_on
  -> maps an event to bounded complete-function restart
```

The same event may therefore invalidate one context and restart selected operations without introducing a resource-specific consistency sublanguage.
