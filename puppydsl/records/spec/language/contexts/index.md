# Concept: PuppyDSL contexts

- **id**: `spec:puppydsl.language.contexts`
- **status**: draft
- **date**: 2026-07-25
- **parent**: `spec:puppydsl.language`

## What this is

Defines structured context lifetimes, lazy provided values, context-local caching, cleanup, function requirements, and long-lived generation invalidation.

## Concept model

```text
function or host opens context instance
  -> provider runs lazily when required
  -> value is cached in that instance
  -> nested calls inherit active context chain
  -> context closes in reverse structured order
  -> destructors run for successfully constructed values
```

A context ID names one declaration. A context scope is the runtime interval during which one instance remains active.

## Context declaration

```yaml
contexts:
  operation:
    detail: Owns values cached for one operation invocation.
    provides:
      current_state:
        provider: runtime.state.collect()
      index:
        provider: runtime.state.build_index(operation.current_state)
      watcher:
        provider: runtime.watcher.start()
        destructor: runtime.watcher.stop(self)
```

Each provided value has one dependency ID:

```text
<context-id>.<provided-key>
```

The provider output type is inferred from the declared provider function signature. A separate type field is prohibited.

## Provider rules

| subject | contract |
|---|---|
| Provider | One ordinary declared function call. |
| Output | Cacheable information-bearing value; `void` and `never` are invalid. |
| Inputs | Every signature input bound exactly once. |
| Dependencies | Same-context or active-ancestor provided values only. |
| Hidden requirements | Provider function-level `requires` is prohibited. |
| Evaluation | Lazy on first requirement in one context instance. |
| Cache | Reused for the remainder of that context instance. |

Provider references form an acyclic dependency graph. Declaration order does not define construction order.

## Destructor rules

A destructor is one ordinary declared function call returning `void`.

Destructor arguments may reference only `self`, which denotes the cached value of that provided entry.

A destructor:

- runs only after successful provider construction;
- runs when the context generation closes;
- cannot depend on another provided value;
- cannot use function-level `requires`.

Successfully constructed values are destroyed in reverse dependency-construction order.

## Function-owned opening

A function opens contexts through an ordered `opens` list.

```yaml
opens:
  - repository_snapshot
  - operation
```

The first context is outermost. Contexts close in reverse order after return or failure.

A function without `opens` inherits the caller's active context chain.

Opening and closing contexts inside a step, branch, case, iteration, or expression is prohibited.

## Context consumption

A function declares context-provided dependencies through `requires`.

```yaml
requires:
  - operation.index
```

The runtime resolves the value from the active context chain. If the value is not cached, its provider runs before the consumer function body.

PuppyDSL exposes no general context object or dynamic context-key lookup.

## Dynamic parent chain

A context declaration has no fixed parent field.

The active parent chain is formed by structured function and host scopes. One context may participate in different valid chains.

Every provider ancestor reference must be available on every entry route capable of constructing that provider.

## Same-context reopening

Opening a context ID already active in the current chain is prohibited.

A separate semantic lifetime requires a separately named context.

## Host-owned contexts

A runtime host may open a context across several entry invocations. Host ownership does not change provider, requirement, cache, or cleanup semantics.

## Long-lived invalidation

A retained context may declare invalidating events.

```yaml
invalidated_by:
  - repository.changed
```

Invalidation marks the current generation stale for later acquisition. Existing consumers keep their pinned generation until completion.

```text
event dispatch
  -> mark matching current generation stale
  -> reject stale generation for later acquisition
  -> preserve it for active consumers
  -> lazily create a replacement on later acquisition
  -> close stale generation after its final consumer exits
```

Invalidation changes generation eligibility. Destructors run only when a generation closes.

## Rules

- Context opening is structured and function- or host-owned.
- Context instances cache provided values independently.
- Contexts close in reverse open order.
- Requirements must be statically satisfiable on every selected entry route.
- An earlier context in one `opens` list cannot depend on a later context.
- Provider dependency cycles are invalid.
- Long-lived invalidation never mutates values already pinned by an active invocation.

## Boundary

| owned here | owned elsewhere |
|---|---|
| Context declaration and dependency semantics. | Runtime context-manager implementation. |
| Function `opens` and `requires` meaning. | Host entry-route selection. |
| Lazy cache and cleanup semantics. | Native resource representation. |
| Generation invalidation semantics. | Event declaration and dispatch ordering. |
| Static satisfiability requirements. | Diagnostic code and output format. |

## Related specs

| ref | relation |
|---|---|
| `spec:puppydsl.language.functions` | Declares `opens` and `requires`. |
| `spec:puppydsl.language.events` | Defines invalidating event identity and dispatch. |
| `spec:puppydsl.language.program_validity.flow_analysis` | Proves route and provider dependency validity. |
