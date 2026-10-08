# Runtime context model

## Status

Working contract for structured context ownership, lazy dependency construction, context-instance-local caching, cleanup, ordered multi-context opening, and static availability validation.

## Declaration detail

Every context declaration includes a non-empty `detail` field that states the lifecycle responsibility and values owned by that context.
The field is bootstrap-validated design authority and is not executable provider or cleanup logic.

## Context definition

A context definition declares one named context and the values available while one instance of that context is open.

The declaration key is the context ID. A context ID is not a separate scope ID. A context scope is the runtime interval during which one context instance remains active.

Context declarations are grouped beneath the `contexts` root mapping, and provider or destructor values reuse the ordinary function-call syntax:

```yaml
contexts:
  operation:
    detail: Owns repository-derived values cached for one operation invocation.
    provides:
      current_records:
        provider: >-
          runtime.current_records.collect(
            source_access: runtime.repository_source_access,
            namespaces: runtime.configured_namespaces
          )

      current_identity_index:
        provider: >-
          runtime.current_records.build_identity_index(
            records: operation.current_records
          )

      repository_watcher:
        provider: runtime.repository_watcher.start()
        destructor: runtime.repository_watcher.stop(self)
```

Each provided value has:

- one stable dependency ID formed as `<context-id>.<provided-key>`;
- one provider call;
- an output type inferred from the provider function signature;
- an optional destructor call.

A separate `type` field is not declared because it would duplicate the provider's authoritative output type.
A provider is an ordinary declared function call whose output type must produce a cacheable value; `void` and `never` are invalid provider outputs.
Every provider input declared by its signature must be bound exactly once by the call arguments.
A zero-input provider still uses an explicit empty call `()`.
A provider must not rely on function-level `requires`; all of its dependencies are expressed by explicit call arguments.

A provider argument may reference only a value provided by the same context or an active ancestor context.
It is not a nested function call or an arbitrary expression.
If another function must derive the argument value, that derived value is declared as its own provided context value and referenced by dependency ID.

Provider argument names must exactly match the provider signature. Each argument value must be assignment-compatible with the declared parameter type; exact type identity is required except for direct assignment of a declared union variant to that union type.
Unknown, missing, duplicate, or incompatible provider bindings are bootstrap errors.

Provider references form a dependency graph over provided values. That graph must be acyclic across direct and indirect dependencies, including dependencies that cross active ancestor contexts. Mutual reference and longer cycles are bootstrap errors. Declaration order does not establish provider construction order; the dependency graph does.

A provider executes lazily when the value is first required in one context instance.
The result is cached in that context instance and reused for the remainder of its lifetime.

A destructor is an ordinary declared function call whose output type is `void`.
A zero-input destructor uses an explicit empty call `()`.
Destructor arguments may reference only `self`, which denotes the value cached for that provided-value entry.
A destructor must not reference another value from the same context or from an ancestor context, and it must not declare function-level `requires`.

A destructor runs only when its provider completed successfully and the context generation closes.
Constructed values are destroyed in reverse dependency-construction order.

## Lifetime ownership

A context definition does not declare a fixed lifetime category such as process, operation, or generation lifetime.
The component that opens the context owns its lifetime.
A design-facing function may own the context for its complete invocation, while the runtime host may own a context across multiple entry operations.

To obtain a particular lifetime, the design defines or selects the function or host boundary whose scope matches that lifetime.
An `invalidated_by` declaration affects replacement of retained generations but does not itself create or extend the context lifetime.

## Function-owned context opening

A function may open zero or more declared contexts for its complete invocation through an ordered `opens` list:

```yaml
functions:
  operation.validation.validate_records:
    detail: Validates selected records using one operation-scoped dependency set.
    opens:
      - repository_snapshot
      - operation
    signature: >-
      (selectors: list<RecordSelector>) -> ValidationResult
    implementation:
      native: operation.validation.validate_records
```

The `opens` list contains context IDs. The first listed context becomes the outermost newly opened context and the last listed context becomes the innermost:

```text
inherit active caller contexts
  -> open repository_snapshot
  -> open operation
  -> resolve and execute the complete function body
  -> allow nested calls to inherit the complete active context chain
  -> close operation after return or failure
  -> close repository_snapshot after return or failure
```

Contexts close in reverse open order. This is a structured lifecycle equivalent to nested Python `with` blocks or nested Go `defer Close()` boundaries.

A function without `opens` opens no context and inherits the active context chain from its caller.

## Context consumption

A function declares context-provided values through `requires`:

```yaml
functions:
  artifact.base.reference.resolve_current:
    detail: Resolves one artifact reference against the current identity index.
    requires:
      - operation.current_identity_index
    signature: >-
      (reference: ArtifactReference) -> ReferenceResolution
    implementation:
      native: artifact.base.reference.resolve_current
```

The runtime resolves the requirement from the active context chain.
If the value has not been constructed in the matching context instance, its provider runs and the result is cached before the consumer function executes.

A function may require values from:

- a context opened by that same function through `opens`;
- an inherited context opened by an ancestor function or the runtime host.

YAML does not receive a general context object and cannot perform dynamic key lookup.

## Structured boundary

Context opening and closing are function-level operations only.
The initial language prohibits opening or closing a context in:

- an individual step;
- an `if` branch;
- a `for` body;
- an arbitrary expression.

If a smaller lifecycle is required, the author defines a separate function with the appropriate `opens` list and calls that function.
This keeps resource lifetime visible in the function call graph and prevents partially executed steps from leaking context instances.

## Static availability validation

Runtime bootstrap must prove that every context requirement is satisfiable before an entry function can execute.

For each call edge, validation carries the ordered active context chain:

```text
caller active context chain
  + contexts opened by callee in declared order
  -> contexts available inside callee body
```

A call is invalid when the callee or any transitive nested call requires a dependency whose context is absent from that chain.

Bootstrap validation must also prove that each newly opened context can construct every reachable provided value from values in the same context or contexts already active outside it. An earlier `opens` entry cannot depend on a later entry because the later context is not active when the earlier context is opened.

Examples:

```text
entry function opens operation
  -> nested resolve_current requires operation.current_identity_index
  -> valid
```

```text
entry function opens no operation context
  -> nested resolve_current requires operation.current_identity_index
  -> invalid at bootstrap
```

This validation applies to all declared entry routes, including MCP operations, CLI operations, tests, and artifact-module capability slots.
A missing context must not be deferred until the consumer function happens to run.

## Host-owned contexts

A process-lifetime context may be opened and closed by the runtime host rather than one design-facing function.
For example:

```text
runtime process starts
  -> open runtime context
  -> execute MCP operations
  -> close runtime context
```

An operation function may then open an `operation` context while the host-owned `runtime` context remains active.
The operation body can require values from both contexts.

## Dynamic structured parent chain

A context definition does not declare a fixed `parent` context.
The active parent chain is formed by the structured scopes that are open at a concrete entry route and call site.

For example, one route may form:

```text
runtime
  -> repository_snapshot
  -> operation
```

while another may form:

```text
runtime
  -> operation
```

The same context definition may therefore participate in different valid active chains.
A provider may reference a value from the same context or from an active ancestor context, but bootstrap validation must prove that the referenced ancestor exists on every entry route that can construct that provided value.
A missing required ancestor is an invalid entry route rather than a reason to add a fixed parent declaration to the context.

## Same-context reopening

Opening a context whose context ID is already active in the current context chain is prohibited.
The initial runtime has no reentrant context instances of the same ID.

Nested reopening would create a second cache and cleanup boundary with the same semantic identity while the first remains active.
If a genuinely separate lifecycle is required, it must use a separately named context and a separately declared function boundary.

## Context validation

Bootstrap validation must reject:

- an unknown context ID in a function `opens` list;
- duplicate context IDs within one `opens` list;
- reopening a context ID already active on any entry route;
- an `opens` order that makes a provider depend on a context opened later in the same list;
- an unknown provided-value ID in `requires` or provider arguments;
- provider arguments that reference a context unavailable on a reachable construction route;
- direct or indirect cycles in the provided-value dependency graph;
- a context requirement unavailable from the complete active chain of any declared entry route;
- a provider whose output type is `void` or `never`;
- a destructor with an invalid signature, dependency, or non-`void` output.

Design-closure validation should additionally report:

- a declared context that no host or function ever opens;
- a provided value unreachable from every selected entry route;
- a provided value never required by any reachable function or provider.

These closure findings do not replace bootstrap correctness checks. They identify orphaned design rather than unsafe execution.

## Long-lived context invalidation

Invalidation is separate from nested reopening and from destructor execution.
It applies only when a context generation is intentionally retained across more than one operation.

A long-lived context declares the events that invalidate its current generation:

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

A filesystem watcher or an explicit repository write may dispatch `repository.changed`.
The context manager is the sole owner of the generic invalidation operation.
Ordinary YAML functions do not call an arbitrary context-specific invalidation function.

Before an operation body begins, the operation runner pins each long-lived context generation required by that operation's transitive `requires` set.
All context requirements resolved during that operation use the pinned generation, including requirements first encountered after an invalidating event.
Provided values inside the pinned generation remain lazily constructed.

The required invalidation sequence is:

```text
dispatch declared event
  -> record the event occurrence
  -> find contexts whose invalidated_by includes that event
  -> atomically mark each current generation stale
  -> reject that generation for later acquisitions
  -> preserve it for consumers that already acquired it
  -> lazily construct a replacement generation on later acquisition
  -> close the stale generation after its final active consumer exits
```

Invalidation does not immediately destroy or clear a generation that is in use.
A stale generation remains usable by operations that pinned it before invalidation, but it cannot be acquired by a later operation.
This preserves one consistent context generation for the complete consuming operation and prevents one operation from mixing old and replacement generations.

A generation close discards its cached values.
If a cached value has a destructor, the context manager invokes that destructor during close.
Therefore:

```text
invalidation
  -> changes generation eligibility

destructor
  -> performs optional resource cleanup when a generation finally closes
```

The invalidation mechanism requires the context manager to track at least:

```text
context ID
generation ID
fresh or stale state
active consumer count
cached provided values
construction order
```

Event identity, producer authority, event occurrence tracking, and complete-operation restart are defined in `event-model.md`.
The same declared event may invalidate a long-lived context and trigger bounded restart of selected restartable operations.

A context whose lifetime is only one function invocation requires no invalidation event because it is closed and discarded when that invocation ends.
