# Reference: PuppyDSL steps and bindings

- **id**: `spec:puppydsl.language.functions.steps_and_bindings`
- **status**: draft
- **date**: 2026-07-25
- **parent**: `spec:puppydsl.language.functions`

## What this is

Defines ordered step execution, immutable lexical bindings, discard labels, return rules, and child-scope visibility for steps-defined functions.

## Ordered steps

`implementation.steps` is a non-empty YAML mapping. Entries execute in source order.

Each step value is one supported expression or structured control-flow node.

| static result | step key | lexical result |
|---|---|---|
| Information-bearing type | Ordinary lower-snake name | One immutable local binding. |
| `void` | `_<completion-name>` | No local binding. Execution continues. |
| `never` | `_<completion-name>` | No local binding. The lexical path terminates. |

The exact key `_` is invalid.

## Ordinary bindings

An ordinary step binds its result once.

```yaml
normalized: common.text.trim(body)
empty: common.text.is_empty(normalized)
```

A local may reference:

- function inputs;
- required context values;
- active ancestor-scope bindings;
- locals bound by earlier steps in the same scope;
- current iteration bindings;
- a switch target under case-local narrowing.

Forward references are invalid.

## Discard labels

A discard label records one completed action for ordering and diagnostics.

```yaml
_served: server.serve()
_failed: runtime.throw(message)
```

A discard label:

- does not enter the binding namespace;
- cannot be referenced or returned;
- must have a non-empty suffix;
- must match a `void` or `never` result.

A reachable step after a `never` discard is invalid.

## Return rules

A steps-defined function whose output is neither `void` nor `never` declares `implementation.return`.

```yaml
implementation:
  steps:
    normalized: common.text.trim(body)
  return: normalized
```

`return` names one visible binding. It does not contain a call, projection, constructor, or other expression.

| declared output | return rule |
|---|---|
| Information-bearing type | Required. Returned binding must be assignable. |
| `void` | Prohibited. Completion follows the final continuing step. |
| `never` | Prohibited. Every reachable path must terminate. |

## Lexical scopes

Each function body is one lexical scope. Each `if` branch, switch case, and iteration body creates one child scope.

| rule | contract |
|---|---|
| Read ancestor binding | Allowed. |
| Export child local directly | Prohibited. |
| Export complete control-flow result | Allowed through the enclosing step key. |
| Reassign active binding | Prohibited. |
| Shadow active ancestor binding | Prohibited. |
| Reuse a name in sibling scopes | Allowed because the bindings do not coexist. |

Iteration bindings and narrowed switch targets follow the same child-scope boundary.

## Source order and call graph

Steps execute in source order. A derived dependency graph may support validation or inspection, but it does not replace source-order execution.

YAML-defined direct and mutual recursion are invalid in the initial language profile. Native implementations may use recursion internally.

## Validation rules

- Reject an empty steps mapping.
- Reject duplicate step keys in one lexical scope.
- Reject forward references and references to child-scope locals after scope exit.
- Reject reassignment or shadowing of an active binding.
- Reject an ordinary key for a `void` or `never` result.
- Reject an underscore-prefixed key for an information-bearing result.
- Reject the exact discard key `_`.
- Reject references or returns that name a discard label.
- Reject a reachable step after a `never` result.
- Reject a missing, prohibited, or incompatible `return`.
- Reject direct or mutual YAML-defined recursion.

## Related specs

| ref | relation |
|---|---|
| `spec:puppydsl.language.functions` | Parent function overview. |
| `spec:puppydsl.language.functions.implementation_forms` | Defines the steps implementation form. |
| `spec:puppydsl.language.expressions` | Defines permitted step expressions. |
| `spec:puppydsl.language.control_flow` | Defines child-scope control-flow nodes. |
| `spec:puppydsl.language.program_validity.flow_analysis` | Applies reachability and call-graph validation. |
