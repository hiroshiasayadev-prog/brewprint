# Reference: PuppyDSL flow analysis

- **id**: `spec:puppydsl.language.program_validity.flow_analysis`
- **status**: draft
- **date**: 2026-07-25
- **parent**: `spec:puppydsl.language.program_validity`

## What this is

Defines whole-body and whole-program analysis for lexical visibility, reachability, path completion, function calls, context availability, provider dependencies, event dispatch, and restart safety.

## Lexical analysis

Each function body and control-flow child scope has one binding namespace.

Validation proves:

- every reference resolves to an active visible binding;
- every local is declared before use;
- no active binding is reassigned or shadowed;
- child locals do not escape;
- discard labels never enter the binding namespace;
- switch narrowing is limited to its case;
- optional absence exposes no target value.

## Reachability

A `never` result terminates the current lexical path.

Explicit `continue` terminates the current iteration path.

Any later step in the same terminated path is invalid.

## Function completion

| declared output | required completion proof |
|---|---|
| Information-bearing type | Every normally completing path reaches one assignable return binding. |
| `void` | Every normally completing path may finish without a return. |
| `never` | No reachable path completes normally. |

## Control-flow completion

Validation computes the completion state of every `if`, switch, and `for` node.

The analysis distinguishes:

- continuing path with a typed value;
- continuing path without a value;
- continuing path with `IterationResultContributionAbsent`;
- explicit iteration `continue`;
- invocation termination through `never`.

`IterationResultContributionAbsent` is a control-flow result state, not `void`, a declared union type, or an implicit `continue`. It preserves normal reachability while indicating that the current bounded iteration contributes no collected element.

Value-producing branch and case results must use one exact type. A `for` with `return` must establish its return binding on every contributing iteration path; `IterationResultContributionAbsent`, explicit `continue`, and `never` paths do not contribute an element.

## Function call graph

The source-defined call graph includes:

- direct calls from steps;
- calls used by context providers and destructors;
- generated event callable dispatch;
- calls nested in control-flow nodes.

Direct and mutual recursion among PuppyDSL-defined functions are invalid in the initial language profile.

Native implementation internals are outside this graph.

## Context-route analysis

For each selected entry route, validation carries the ordered active context chain.

```text
caller active contexts
  + callee opens in declared order
  -> contexts available inside callee
```

Validation proves:

- every direct and transitive `requires` value is available;
- a context is not reopened while already active;
- an earlier `opens` entry does not depend on a later entry;
- every provider ancestor dependency exists on each route that can construct it;
- host-provided initial contexts are included in the selected route contract.

A resolved provided-value identity is insufficient when its owning context is unavailable on the route.

## Provider dependency graph

Provider arguments create edges between provided values.

The complete reachable provider graph must be acyclic. The rule applies to same-context and cross-ancestor dependencies.

Destructors do not create provider dependencies because they may reference only `self`.

## Event and restart analysis

Validation proves:

- every `invalidated_by` and `restart_on` event exists;
- `restart_on` has a finite `max_restarts`;
- only selected entry functions declare restart;
- a restartable function does not directly or transitively dispatch any event;
- event callable dispatch uses the declared payload contract.

## Profile reachability

Design-closure and activation profiles apply executable-readiness checks to declarations reachable from selected entry routes.

Unreachable declarations may receive separate closure findings, but they do not excuse invalid syntax, unresolved identities, or invalid declaration shapes.

## Validation rules

- Reject unbound, forward, escaped, reassigned, or shadowing bindings.
- Reject a reachable step after `never` or `continue`.
- Reject incomplete function or control-flow paths, except an accepted `IterationResultContributionAbsent` path at a bounded iteration contribution boundary.
- Reject incompatible value-producing result types.
- Reject direct or mutual PuppyDSL function-call cycles.
- Reject unavailable context requirements on any selected route.
- Reject same-context reopening and invalid open order.
- Reject provider dependency cycles.
- Reject unsafe restart declarations and restartable event-dispatch paths.
- Evaluate one graph proof only after its referenced declarations resolve; independent graph regions remain analyzable.

## Related specs

| ref | relation |
|---|---|
| `spec:puppydsl.language.program_validity` | Parent validity overview. |
| `spec:puppydsl.language.functions.steps_and_bindings` | Defines lexical binding and completion inputs. |
| `spec:puppydsl.language.control_flow` | Defines structured path forms. |
| `spec:puppydsl.language.contexts` | Defines route and provider dependencies. |
| `spec:puppydsl.language.events` | Defines event dispatch and restart safety. |
