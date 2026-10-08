# Minimum control-flow model

## Status

Working requirements derived from current validation consumers.
The initial `if`, exhaustive `switch` over declared unions, enums, and optional values, and bounded `for` forms are defined here.

## Boundary

Control flow exists only to express meaningful design-rule behavior.
It includes exhaustive branching over statically declared unions and enums and over the closed present/absent states of `optional<T>`, but does not expose reflection, arbitrary runtime type assertions, active variant types as values, open-ended scalar switching, null values, or general-purpose programming.

## Common lexical scope

`if`, `for`, and `switch` use one common child-scope rule.

- A branch, iteration body, or switch case may read every active binding from its enclosing scope, except that an optional switch target is unavailable as a value inside its own `absent` case.
- A local created inside that child scope is unavailable after the branch, iteration, or case completes.
- Only an ordinary local bound to a non-`void`, non-`never` control-flow result becomes visible to later steps in the enclosing scope.
- A control-flow result of `void` or `never` uses a discard key of the form `_<completion-name>` and creates no enclosing-scope binding.
- Sibling child scopes may reuse the same local names because their bindings never coexist.
- A child binding must not shadow or reassign an active ancestor binding.
- Bindings remain immutable.

Control-flow-specific bindings follow the same boundary. A list `for` item binding, or dictionary `for` key and value bindings, exist only in that iteration body. A union switch target keeps the same local name, but its static type is narrowed to the selected variant only inside that case; outside the case it retains its declared union type. An enum switch target retains its declared enum type in every case, while the selected member is known only inside that case. An optional switch target keeps the same local name and narrows to `T` inside `present`; inside `absent`, the target is unavailable in value positions because no `T` value exists. Outside the optional switch, the target retains `optional<T>`.

## Conditional execution

A minimum conditional construct is likely required for rules such as:

- conditional metadata requiredness;
- Required Outcome applicability based on Requirement status;
- Task output interpretation as canonical reference or literal;
- status-dependent or task-type-dependent validation;
- suppressing a dependent finding when its prerequisite value is missing or invalid.

A condition is exactly one boolean local produced by a named function or earlier rule evaluation.
The conditional grammar does not contain inline comparison, arithmetic, arbitrary field traversal beyond statically typed binding-origin access chains, boolean composition, or type-test expressions.

The authoring form is:

```yaml
<result-step>:
  if:
    condition: <boolean-local>
    then:
      steps:
        <ordered branch steps>
      return: <branch-local>
    else:
      steps:
        <ordered branch steps>
      return: <branch-local>
```

`<result-step>` is an ordinary local name when the complete `if` returns an information-bearing value and a discard key when the complete result is `void` or `never`.
`else` and each branch `return` are optional subject to the output rules below.
The condition is exactly one previously produced boolean local.

Conditional output is determined by branch completion:

```text
no branch returns a value and at least one path continues normally
  -> if result is void
  -> the complete if step uses a discard key and creates no binding
  -> else may be omitted and the false path is a no-op

all possible branches terminate with a never call
  -> if result is never
  -> the complete if step uses a terminal discard key and creates no binding

one branch returns U and an explicit sibling branch continues
  -> every continuing branch must return the same U

one branch returns U and else is omitted inside bounded for
  -> omitted else is an implicit continue of the nearest enclosing for

branch terminates with explicit or implicit continue
  -> that branch returns no value
  -> remaining continuing branches determine the if result type

branch terminates with a never call
  -> that branch has no normally continuing path
  -> remaining continuing branches determine the if result type
```

A branch that reaches the surrounding execution after the `if` must establish the result local whenever the `if` has non-void type.
A non-void `if` with omitted `else` is therefore valid only inside a bounded `for`; outside iteration it is a bootstrap error.
The runtime does not implicitly wrap a missing branch value in `optional<U>`.

## Conditional lexical scope

The `then` and `else` branches are independent child scopes governed by the common lexical-scope rule.
The condition local remains available with its original type. An information-bearing complete `if` result is exported through its ordinary local, while a `void` or `never` result exports no binding.

## Prohibited use of conditional execution

YAML must not use `if` to implement runtime type routing or inline computation.

Prohibited:

```yaml
if base_ref is SequentialBaseReference:
  ...

if (a.ccc.ddd * 100 / pi) >= b.eee:
  ...
```

A declared union, enum, or optional is consumed through exhaustive `switch`, while module capability routing may continue to use the module-bound dispatch boundary:

```yaml
module: artifact.base.module.from_base_reference(base_ref)
parse_outcome: module.reference.parse(base_ref)
```

## Exhaustive switch

A `switch` consumes exactly one previously bound local whose static type is a declared union, a declared enum, or `optional<T>`.
The target type determines whether case keys name union variants, enum members, or the exact optional cases `present` and `absent`.
The switch target must be one previously bound local; inline calls, projections, and expressions are prohibited.

The common authoring form is:

```yaml
<result-step>:
  switch: <union-enum-or-optional-local>
  cases:
    <variant-type-enum-member-or-optional-case>:
      steps:
        <ordered case steps>
      return: <case-local>
```

`<result-step>` is an ordinary local name for an information-bearing switch result and a discard key for a `void` or `never` result.
`steps` and `return` are independently optional subject to the output and completion rules below.

### Union target

For a declared union, each case key names one variant type from that union declaration.
Inside a case, the switch target local keeps its original name and is statically narrowed to the named variant type.
No separate tag or payload binding exists.

For the declared union:

```yaml
types:
  BaseReference:
    detail: Carries one structure-specific base reference for module routing.
    union:
      - TreeBaseReference
      - SequentialBaseReference
```

a valid union switch is:

```yaml
raw:
  switch: reference
  cases:
    TreeBaseReference:
      steps:
        value: reference.raw
      return: value
    SequentialBaseReference:
      steps:
        value: reference.raw
      return: value
```

Within the first case, `reference` has static type `TreeBaseReference`. Within the second, it has static type `SequentialBaseReference`. After the switch, the outer `reference` binding remains `BaseReference`.

The `cases` mapping must contain every variant type declared by the union exactly once.
Unknown, missing, or duplicate variant cases are bootstrap errors.
The active variant type is used only by the interpreter to select a statically validated case. It cannot be projected, compared, returned, or passed as a value.

### Enum target

For a declared enum, each case key is one qualified member literal from that enum declaration.
Every declared member must appear exactly once.

For the declared enum:

```yaml
types:
  RecordStructure:
    detail: Identifies the closed source-structure model used by an artifact module.
    enum:
      - sequential
      - tree
```

a valid enum switch is shown below. Both called functions must declare the same output type because switch output does not infer a common union from different concrete case results.

```yaml
artifact_root:
  switch: structure
  cases:
    RecordStructure.sequential:
      steps:
        discovered: operation.discovery._discover_sequential_root(records_root, module)
      return: discovered
    RecordStructure.tree:
      steps:
        discovered: operation.discovery._discover_tree_root(records_root, module)
      return: discovered
```

Inside either case, `structure` retains static type `RecordStructure`; enum members are values of one enum type rather than separate member-specific types.
The selected member is known only within its case and does not create a new binding or singleton type.
Unknown members, unqualified members, members from another enum, missing members, and duplicate member cases are bootstrap errors.

An enum switch is permitted because the target type declares one closed member set that bootstrap can exhaustively validate.
It does not permit switching over arbitrary strings, integers, named scalars, or other open-ended values.

### Optional target

For `optional<T>`, the `cases` mapping contains exactly these two unqualified case keys:

```text
present
absent
```

A valid optional switch is:

```yaml
maybe_record: >-
  stdlib.dictionary.get(
    dictionary: records,
    key: node_ref
  )

selected:
  switch: maybe_record
  cases:
    present:
      steps:
        projected: operation.tree.project_record(maybe_record)
      return: projected
    absent:
      steps:
        projected: operation.tree.project_missing_record()
      return: projected
```

Inside `present`, the switch target keeps its original binding name and is statically narrowed from `optional<T>` to `T`. No separate payload binding is introduced.
Inside `absent`, the runtime has established that no `T` value exists. The target name does not resolve as a value binding in that case and cannot be projected, passed, returned, or otherwise used in a value position.
After the switch, the outer target retains its original `optional<T>` type and value.

The optional case words are control-flow syntax rather than enum members, types, literals, or constructible values.
The invocation grammar does not gain `none`, `some`, `null`, or another optional constructor from this switch form.
An optional value may be passed or returned without switching, but any access to its contained `T` requires an exhaustive optional switch.

A derived control-flow or Mermaid DAG may label the two outgoing case edges `present` and `absent`.
The `absent` edge represents control state only and must not create a null, none, or absent-value asset node.

### Exhaustiveness and case order

Case order has no semantic effect, but authors should normally follow the union variant or enum member declaration order and write `present` before `absent` for optional switches.
A `default` case is prohibited for union, enum, and optional switches.
Every possible union variant or enum member must be represented explicitly, and an optional switch must contain exactly one `present` case and one `absent` case.
This construct is not reflection, `isinstance`, an open-ended type test, or arbitrary value switching.

### Switch output

Switch output follows these rules:

```text
no continuing case returns a value and at least one case continues normally
  -> switch result is void
  -> the complete switch step uses a discard key and creates no binding

all cases terminate with a never call
  -> switch result is never
  -> the complete switch step uses a terminal discard key and creates no binding

one continuing case returns U
  -> every other continuing case must return the same U

case terminates with continue inside bounded for
  -> that case contributes no value
  -> remaining continuing cases determine the switch result type

case terminates with a never call
  -> that case has no normally continuing path
  -> remaining continuing cases determine the switch result type
```

A non-`void`, non-`never` switch requires every case that reaches surrounding execution to return the same exact type.
A case ending in a `never` call does not reach surrounding execution and therefore does not return or participate in result-type agreement.
The runtime does not infer a common supertype, implicitly construct another union, or wrap divergent case values in `optional<T>`.
A case `return` may name one local bound earlier in that case. It may also name the switch target when that target's case-local static type is assignment-compatible with the required output type.

### Switch lexical scope

Each case is an independent child scope governed by the common lexical-scope rule.
A union case narrows the switch target to its selected variant type only inside that case.
An enum case leaves the switch target's static type unchanged while establishing its selected member only inside that case.
An optional `present` case narrows the target to `T`; an optional `absent` case establishes absence and removes the target from value binding resolution for that case.
Outside the switch, the target retains its original declared type and value.

## Bounded iteration

A bounded iteration construct is required by concrete consumers including:

- trimming and validating every metadata list item;
- parsing and resolving every persisted reference;
- traversing configured applications or artifact roots stored in dictionaries;
- generating findings for duplicate groups;
- validating every physical direct child in a Specification tree;
- constructing normalized relation edges;
- generating findings for detected cyclic regions.

List iteration uses one item binding:

```yaml
<result-step>:
  for:
    item: <item-binding>
    in: <list-local>
    steps:
      <ordered body steps>
    return: <body-local>
```

Dictionary iteration uses separate key and value bindings:

```yaml
<result-step>:
  for:
    key: <key-binding>
    value: <value-binding>
    in: <dictionary-local>
    steps:
      <ordered body steps>
    return: <body-local>
```

Exactly one binding form is permitted. List iteration requires `item` and prohibits `key` and `value`; dictionary iteration requires both `key` and `value` and prohibits `item`.

`return` is optional. The enclosing `<result-step>` is always required so the `for` construct remains one step in the surrounding ordered body.
A `for` with `return` produces `list<U>` and uses an ordinary result local. A `for` without `return` produces `void`, uses a discard key, and creates no local binding.

Example over a list:

```yaml
routing_outcomes:
  for:
    item: raw_ref
    in: raw_refs
    steps:
      parsed: artifact.base.reference.parse(raw_ref)
    return: parsed
```

Example over a dictionary:

```yaml
app_scopes:
  for:
    key: app_namespace
    value: app_discovery
    in: discovered_apps
    steps:
      scope: >-
        operation.discovery_and_listing._project_app_scope(
          app_namespace: app_namespace,
          discovery: app_discovery
        )
    return: scope
```

Example without collection:

```yaml
_changes_dispatched:
  for:
    item: change
    in: changes
    steps:
      _change_dispatched: repository.changed(change)
```

## Iteration input evaluation

The `in` field references one previously bound local of type `list<T>` or `dict<K, V>`.
It does not contain an inline function call, access chain, constructor, or arbitrary expression.

For `list<T>`, the declared `item` binding has type `T`.
For `dict<K, V>`, the declared `key` binding has type `K` and the declared `value` binding has type `V`.

The runtime resolves the input local once when the `for` begins and iterates over that fixed immutable collection value.
Element or entry count, ordering, and membership do not change during the iteration, even if an inherited long-lived context generation becomes stale while the loop is running.
List iteration preserves list order. Dictionary iteration uses the canonical key order defined by `value-model.md`; dictionary insertion or source declaration order is not observable.
An empty input collection executes no iteration body:

```text
for with return
  -> empty list<U>

for without return
  -> void
```

The runtime need not physically copy the collection when the value model already guarantees immutability.

## Iteration requirements

The minimum iteration model must define:

- input collection type `list<T>` or `dict<K, V>`;
- one statically typed item binding for list iteration;
- statically typed key and value bindings for dictionary iteration;
- deterministic list order or canonical dictionary key order;
- lexical scope for iteration locals;
- propagation of runtime execution failure;
- prohibition of mutation to the input collection.

Iteration output is determined mechanically by the presence of `return`:

```text
no return
  -> void

return of U
  -> list<U>
```

Each iteration contributes exactly one returned value when `return` is present. Output order matches list order for `list<T>` input and canonical key order for `dict<K, V>` input.
For a `for` that declares `return: value`, every reachable path in one iteration must either define `value` and reach the iteration return, terminate that iteration through explicit or implicit `continue`, or terminate the complete runtime invocation through a `never` call.
This is validated statically before execution.

The type `U` is determined statically from the declared type of the returned body local, not inferred from runtime elements.
The `for` result therefore remains `list<U>` when the input list is empty or every iteration continues without contributing a value.
An empty runtime result is a typed empty list rather than an untyped collection.

A direct `void` step creates no local binding, so it cannot be named by the iteration `return` and bounded iteration does not collect direct `void` completions into `list<void>`.
A separately declared function may still return an information-bearing composite type such as `list<void>` if that type is required by another contract; that complete composite value is bound normally because its output type is not `void` itself.
The runtime does not implicitly filter optional values or flatten returned lists.

## Continue

A bounded iteration may skip the current item with `continue`.
Explicit `continue` is permitted only as the terminal action of an `if` branch or switch case lexically contained within that `for` body.
A call to a `never` function is likewise terminal in any branch, case, or ordered step body, but it terminates the complete runtime invocation rather than only the current iteration.
A non-void `if` inside bounded iteration also has implicit-continue semantics when its `else` branch is omitted.
`continue` cannot appear as an ordinary callable step, outside a bounded iteration, or with later steps in the same branch.

When explicit or implicit `continue` is executed, it:

- stops the current iteration immediately;
- contributes no value to the enclosing `for` result, even when that `for` declares `return`;
- begins the next input item without mutating the input collection.

The collected result therefore preserves the relative input order of only the iterations that reach `return`.
This provides explicit filtering without inserting `null`, wrapping every item in `optional<T>`, or adding implicit collection filtering.

`break` is not included in the initial model because no current concrete consumer requires prefix termination.

In nested bounded iterations, `continue` always targets the nearest lexically enclosing `for`.
Labelled `continue` and targeting an outer iteration directly are not included in the initial model.
An outer iteration must decide its own skip condition in its own body.

## Iteration lexical scope

Each iteration body is a fresh child scope governed by the common lexical-scope rule.
The list item binding, or the dictionary key and value bindings, and ordinary body locals exist only for that iteration. Outer bindings remain readable, and only a non-`void`, non-`never` complete `for` result is exported to the enclosing scope.

## Control-flow nesting

`if`, exhaustive `switch`, and bounded `for` may be nested in any combination.
Representative forms include:

```text
switch inside for
for inside switch
if inside switch
switch inside if
for inside if
if inside for
for inside for
```

Each nested construct creates and applies its own lexical child-scope rules recursively.
No special-case visibility or shadowing exception is introduced for nested control flow.

Control-flow nesting does not consume a generic per-node fuel budget.
A bounded `for` terminates because its `in` value is an already materialized finite `list<T>`, while unbounded loop constructs are absent.
Runtime cancellation or deadlines may still stop the complete invocation, but they do not change the language-level result semantics.

## Collection functions versus control flow

The runtime should not add convenience capabilities such as `trim_all`, `parse_all`, or `resolve_all` solely to avoid iteration.
Scalar functions remain scalar when the design meaning is per item.

Reusable collection algorithms such as duplicate discovery or set comparison remain named functions because they perform one independently meaningful collection operation.

## Domain-outcome branching

Expected domain outcomes such as parsed versus unrouteable, resolved versus unresolved, or successful versus expected failure may be declared as unions and consumed directly through exhaustive `switch`. Optional presence versus absence is consumed through the same switch construct using `present` and `absent`.
This avoids outcome-specific native predicate and extraction functions whose only purpose would be to reproduce union elimination.

Conceptually:

```yaml
handled:
  switch: routing_outcome
  cases:
    RoutedBaseReference:
      steps:
        reference: routing_outcome.reference
        module: artifact.base.module.from_base_reference(reference)
        parse_outcome: module.reference.parse(reference)
        handled_reference:
          switch: parse_outcome
          cases:
            ArtifactReference:
              steps:
                resolved: operation.reference.resolve(parse_outcome)
            InvalidArtifactReference:
              steps:
                diagnostic: operation.reference.invalid_artifact_reference(parse_outcome)
    UnrouteableReference:
      steps:
        diagnostic: operation.reference.routing_failure(routing_outcome)
```

The outcome declaration owns the closed variant-type list.
The switch narrows `routing_outcome` itself inside each case; it does not expose union representation, reflection, arbitrary type assertions, or execution-failure handling.
Unexpected execution failure continues to propagate outside ordinary YAML-visible outcomes.

## Excluded initial features

The initial control-flow model does not include:

- unbounded loops;
- `while`;
- recursion through control syntax;
- `break`;
- `continue` outside the terminal position of an `if` branch or switch case inside bounded iteration;
- non-exhaustive, defaulted, or open-ended switching;
- switching over strings, integers, named scalars, or other open-ended values;
- optional case names other than `present` and `absent`;
- arbitrary runtime type assertions;
- exception handling or recovery in YAML;
- mutable variables;
- parallel iteration;
- arbitrary boolean, comparison, field-traversal beyond statically typed binding-origin access chains, or arithmetic expression grammar.

Long-lived, waiting, or potentially unbounded loops such as filesystem watchers, event loops, channel receive loops, and server accept loops belong to native Go implementations or the runtime host.
YAML may declare and invoke their semantic boundaries but does not express their loop mechanics.
