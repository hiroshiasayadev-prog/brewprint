# Reference: PuppyDSL bounded iteration

- **id**: `spec:puppydsl.language.control_flow.iteration`
- **status**: draft
- **date**: 2026-07-25
- **parent**: `spec:puppydsl.language.control_flow`

## What this is

Defines bounded iteration over one finite immutable `list<T>` or `dict<K, V>` value.

## List iteration

```yaml
results:
  for:
    item: source
    in: sources
    steps:
      result: example.process(source)
    return: result
```

`item` is required. `key` and `value` are prohibited.

The item binding has type `T`.

## Dictionary iteration

```yaml
results:
  for:
    key: name
    value: entry
    in: entries
    steps:
      result: example.project(name, entry)
    return: result
```

`key` and `value` are both required. `item` is prohibited.

The key binding has type `K`. The value binding has type `V`.

## Input

`in` names one previously bound immutable collection.

Inline calls, access chains, constructors, and other expressions are invalid in `in`.

The runtime resolves the input once before iteration.

## Order

List iteration preserves list order.

Dictionary iteration uses deterministic canonical key order:

| key type | order |
|---|---|
| `string` | Ascending exact Unicode scalar sequence. |
| `integer` | Ascending numeric order. |
| `boolean` | `false`, then `true`. |
| Named scalar | Backing primitive order while preserving nominal type. |
| Enum | Declaration order. |

Dictionary order exists for deterministic execution. It does not expose insertion order.

## Result

| body form | complete result |
|---|---|
| No `return` | `void`; enclosing step uses a discard label. |
| `return` names `U` | `list<U>`; enclosing step uses an ordinary local. |

Every contributing iteration contributes exactly one value when `return` is present.

When the per-iteration result is `IterationResultContributionAbsent`, the current iteration contributes no element and execution proceeds normally with the next item. This state is not `void` and does not execute `continue`.

An empty input produces an empty typed `list<U>` or `void` according to the declared form.

Returned lists are not flattened. Optional values are not filtered implicitly.

## Continue

Explicit `continue` is allowed only as the terminal action of an `if` branch or switch case inside a bounded iteration.

An information-bearing `if` used as the per-iteration result may instead produce `IterationResultContributionAbsent` from an omitted branch. That state omits one collected element but does not terminate the current path early. Any later reachable iteration-body steps still execute subject to normal binding-visibility rules.

`continue`:

- stops the current iteration;
- contributes no output element;
- begins the next item;
- targets the nearest enclosing `for`.

`break`, labelled continue, and outer-targeted continue are not supported.

A `never` call terminates the complete invocation, not only the current iteration.

## Scope

Each iteration body is one fresh child scope.

Item, key, value, and body locals do not escape. Only the complete information-bearing `for` result enters the enclosing scope.

## Validation rules

- Reject `in` when the binding is unknown or not a list or dictionary.
- Reject inline expressions in `in`.
- Reject a binding form that does not match the collection type.
- Reject iteration bindings that shadow active ancestors.
- Reject body-local references outside the iteration.
- Reject `continue` outside its permitted terminal branch or case position.
- Reject a later step after `continue` in the same path.
- Reject a returned iteration path whose return binding is absent for any reason other than `IterationResultContributionAbsent`, explicit `continue`, or termination through `never`.
- Reject incompatible returned element types.

## Related specs

| ref | relation |
|---|---|
| `spec:puppydsl.language.control_flow` | Parent control-flow overview. |
| `spec:puppydsl.language.control_flow.conditional` | Defines `IterationResultContributionAbsent` for an omitted information-bearing branch. |
| `spec:puppydsl.language.functions.steps_and_bindings` | Defines iteration lexical scope. |
| `spec:puppydsl.language.type_system` | Defines list and dictionary values. |
