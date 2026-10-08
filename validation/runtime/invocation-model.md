# Design-facing invocation model

## Status

Working requirements for how design-facing YAML names values and calls functions.
The enclosing YAML scalar forms and invocation token meanings are confirmed; the complete invocation-parser grammar remains pending.

## Design objective

The YAML surface records concrete design semantics and processing order.
It should remain readable without exposing implementation-level comparison, parsing, or runtime-type operations.

Preferred semantic form:

```yaml
normalized: common.text.trim(body)
empty: common.text.is_empty(normalized)
```

Avoided low-level form:

```yaml
empty: common.value.equals(normalized, "")
```

when a meaningful named operation such as `is_empty` exists.

## Local assignment and discard steps

A step whose static result type is neither `void` nor `never` binds that result to one statically named local value.
A step whose static result type is `void` or `never` instead uses a discard key of the form `_<completion-name>` and creates no local value.

Conceptual syntax:

```yaml
<local-name>: <single-line-value-producing-expression>
<local-name>: >-
  <folded-value-producing-expression>
_<completion-name>: <void-or-never-expression>
```

Examples:

```yaml
normalized: common.text.trim(body)

directories: >-
  dict<PhysicalDirectoryName, PhysicalDirectoryTree>{
    [directory_name]: directory_tree
  }

_served: mcp.serve()
_failed: runtime.throw(message)
```

The plain-scalar and folded-scalar forms have the same binding semantics.
`>-` is a YAML authoring form for one expression string; it does not declare a mutable variable, introduce a separate initialization operation, or permit reassignment.
The expression parser determines the expression's static result type after YAML scalar decoding, and the ordinary step key binds that one result immutably.
Constructors that require folded authoring retain their constructor-specific `>-` requirement.

The exact key `_` is invalid; the suffix must be non-empty and should describe the completed action.
An ordinary key is invalid for a `void` or `never` result, and an underscore-prefixed key is invalid for an information-bearing result.

Ordinary local names are unique within one lexical scope.
Sibling control-flow scopes may reuse a name, while shadowing an active ancestor binding remains prohibited.
A local value is immutable after assignment.
Discard keys remain step labels for source order and diagnostics but are absent from the lexical binding namespace, cannot be referenced or returned, and do not become value assets in a derived dataflow graph.
No reachable step may follow a `never` discard step in the same lexical path.

## Function calls

A call names one fully qualified function ID or one statically resolved module-bound capability.

Conceptual direct call:

```yaml
normalized: common.text.trim(body)
```

Conceptual module-bound call:

```yaml
parse_outcome: module.reference.parse(base_ref)
```

The runtime statically knows the called function's input and output types.
A `void` output requires a non-terminal discard step, while a `never` output requires a terminal discard step.
Neither creates a local binding.

## Arguments

Function calls use a parenthesized call form.

The YAML parser first reads each call as one scalar value. A short call may be a YAML plain scalar, while a long call uses a folded block scalar. The resulting call text is then parsed by the dedicated invocation lexer and parser.

```text
YAML parser
  -> one scalar containing invocation text
  -> invocation lexer and parser
  -> function target, argument bindings, literals, and binding references
```

Literal and reference meaning is therefore defined by the invocation grammar inside that scalar, not by the YAML scalar resolver.

The invocation parser recognizes these argument tokens:

- a double-quoted or single-quoted token is a `string` literal;
- an unquoted identifier or qualified identifier is an input, local, context-provided value, or special binding reference;
- an unquoted decimal integer token is an `integer` literal;
- the exact unquoted tokens `true` and `false` are `boolean` literals;
- a type-qualified member such as `ArtifactStructure.tree` is an enum literal when the qualifier resolves to a declared enum type;
- no null or optional literal exists; `null`, `Null`, `NULL`, `~`, `none`, `some(...)`, and equivalent spellings are rejected in invocation text.

Example:

```yaml
result: >-
  example.evaluate(
    body: body,
    placeholder: "TBD",
    level: 2,
    required: true,
    structure: ArtifactStructure.tree
  )
```

Inside the invocation text, `body` is a binding reference while `"TBD"`, `2`, `true`, and `ArtifactStructure.tree` are literals of their respective declared types.
A quoted token is always a string literal even when its contents match a binding name, integer, boolean, or enum spelling.
A binding reference must resolve in the active lexical or context namespace; an unbound reference is a bootstrap error.

The language supports either positional or named argument binding for one call.
A single call must not mix positional and named arguments.
The invocation model requires:

- deterministic binding to declared input fields;
- bootstrap rejection of missing, duplicate, or unknown arguments;
- exact type identity except for direct assignment of a declared union variant to that union input type;
- readable calls for functions with several arguments.

A one-argument call will usually use positional binding.
Calls with several arguments, especially arguments of the same type, should normally use named binding.
This is an authoring recommendation rather than an initial runtime restriction.

A short call may use one YAML plain scalar:

```yaml
normalized: common.text.trim(body)
```

A long call uses a YAML folded scalar so its parenthesized invocation text remains readable:

```yaml
comparison: >-
  common.collection.compare_sets(
    left: declared_refs,
    right: physical_refs
  )
```

The signature determines positional order and the valid named-argument set.

## Dictionary construction

A step may construct one `dict<K, V>` value with this expression form:

```text
dict<K, V>{[<key>]: <value>, ...}
```

The initializer is authored as one YAML folded scalar:

```yaml
directories: >-
  dict<PhysicalDirectoryName, PhysicalDirectoryTree>{
    [directory_name]: directory_tree,
    [other_name]: other_tree
  }
```

`K` must satisfy the dictionary-key restriction defined by `value-model.md`.
Each `<key>` is exactly one primitive literal, enum literal, or already visible binding reference under the same token rules used by function-call arguments.
A named scalar has no implicit literal conversion from its primitive representation, so a named-scalar key is normally supplied through an already typed binding.
Each `<value>` is exactly one literal or one already visible binding reference.

Every key must be assignment-compatible with `K`, and every value must be assignment-compatible with `V`.
Exact type identity is required except that a value of one directly listed union variant may be supplied where `V` is that union type.
The bracketed key form distinguishes a dictionary key expression from a transparent-record field name.
Entry order is semantically irrelevant, and a trailing comma is optional.
An empty initializer `dict<K, V>{}` is valid because the explicit type arguments determine its complete static type.

Inline function calls, field projections, record constructors, nested dictionary constructors, and arbitrary expressions are prohibited in keys and values.
Such a value must first be bound by an earlier step.
A statically evident duplicate key is a bootstrap error.
When two distinct binding references evaluate to equal keys only during execution, dictionary construction fails rather than replacing an earlier entry.
The constructor produces one immutable value of the exact type `dict<K, V>`.

Dictionary construction itself introduces no mutation.
Required-key lookup is expressed through the binding-origin access chains defined below, and bounded key/value iteration is defined by `control-flow.md`.
Higher-level optional lookup, search, aggregation, or transformation remains the responsibility of declared functions when required by a concrete consumer.

## Transparent record construction

A step may construct one declared transparent-record value with this expression form:

```text
<RecordType>{<field-name>: <value>, ...}
```

The constructor is authored as one YAML folded scalar:

```yaml
domain: >-
  SequentialDomainDiscovery{
    domain_namespace: domain_namespace,
    source_root: source_root
  }
```

A declared zero-field transparent record is constructed with an empty body:

```yaml
not_formed: >-
  CandidateNotFormed{}
```

The empty constructor produces a bindable information-bearing value of exact type `CandidateNotFormed`; it is not a `void` completion and therefore uses an ordinary local key.

`<RecordType>` must resolve to one concrete transparent-record type. Each `<value>` is exactly one literal or one already visible binding reference under the same token rules used by function-call arguments.
Every declared field must appear exactly once, and no unknown field may appear. For a zero-field transparent record, the empty constructor body is the complete matching field set and any supplied field is unknown. Field order is semantically irrelevant; a trailing comma is optional.

A field value must be assignment-compatible with its declared field type. Exact type identity is required except that a value of one directly listed union variant may be supplied where that union type is declared.
The constructor produces one value of the exact named transparent-record type.

Inline function calls, field projections, nested constructors, and arbitrary expressions are prohibited inside constructor fields. Such a value must first be bound by an earlier step.
Record-constructor syntax is invalid for primitives, named scalars, enums, unions, opaque types, and generic composite types.

## Access chains and output access

A value-producing step may use one statically typed access chain beginning with a visible binding:

```text
<binding>(.<field> | [<key>])*
```

Each `.<field>` projects one field declared by the current transparent-record type.
Each `[<key>]` performs required-key lookup on the current `dict<K, V>` type and produces `V`.
The key is exactly one primitive literal, enum literal, or visible binding reference under the ordinary invocation-token rules, and its static type must be exactly `K`.

These access chains are valid when every intermediate type supports the next postfix operation:

```yaml
missing: comparison.right_only
app: inventory.apps[app_namespace]
artifact: inventory.apps[app_namespace].artifact_roots[artifact_kind]
source_root: inventory.apps[app_namespace].artifact_roots[artifact_kind].source_root
```

An access chain must begin with a visible binding rather than a function call, constructor, or arbitrary expression.
The key position likewise admits no inline function call, projection, constructor, arithmetic, or other nested expression.
A runtime dictionary lookup whose key is absent terminates the current invocation as an execution failure; bracket lookup does not produce `optional<V>` and performs no fallback or implicit insertion.

Union values permit no direct projection of their active variant type.
Inside an exhaustive union `switch` case, the switch target local itself is statically narrowed to the named variant type. A separately declared runtime dispatch boundary may perform the same narrowing implicitly.

An `optional<T>` value permits no direct projection or implicit unwrap.
Inside its `present` switch case, the target local itself is statically narrowed to `T` under the same binding name.
Inside its `absent` case, the target is unavailable in value positions because no contained `T` exists.

Opaque values permit no design-facing field access. For example, an `ArtifactModule` returned here remains opaque:

```yaml
module: artifact.base.module.from_base_reference(base_ref)
parse_outcome: module.reference.parse(base_ref)
```

`parse_outcome` is an `ArtifactReferenceParseOutcome`; inside its `ArtifactReference` switch case, the same local is narrowed to the opaque `ArtifactReference` type.

These are invalid:

```yaml
app: load_inventory().apps[app_namespace]
app: inventory.apps[common.text.trim(raw)]
item: list_items[index]
discriminator: base_ref.discriminator
kind: reference.kind
```

Bracket access is defined only for dictionaries; the initial profile provides no list indexing.
A named semantic function must provide any operation required on an opaque value.

## Module-bound dispatch

`module.reference.parse(base_ref)` is design-facing shorthand for invoking the resolved function handle stored in the module registration's `reference.parse` capability slot.

The runtime does not perform a string lookup on every invocation.
The manifest's function ID is resolved during bootstrap and stored in the immutable registration.

A capability slot may declare a structure-specific narrowing contract.
For `reference.parse`, the design-facing argument is `BaseReference`, while the resolved tree or sequential slot function accepts the corresponding declared variant type and returns `ArtifactReferenceParseOutcome`.
The interpreter verifies the union's active variant type against the selected module registration, narrows the input value, and invokes that resolved function.
This is a statically declared runtime dispatch rule rather than general subtyping or an implicit conversion available to ordinary calls.

YAML does not inspect the module's capability map or the union's active variant type and does not choose a slot through a dynamic string.

## Authoring surface and internal representation

A verbose object form such as:

```yaml
id: reference
use: artifact.task.reference.parse
with:
  base_ref: base_ref
```

may remain useful as an internal normalized representation.
It is not currently preferred as the human-authored design surface.

After YAML scalar decoding, the invocation parser may normalize compact step syntax into an internal expression node containing:

```text
ordinary local output name or discard step label
expression kind: call, access chain, record constructor, or dictionary constructor
resolved function ID, module slot, access path, or constructor target
bound arguments, access keys, fields, or dictionary entries
static output type and continuing or terminal completion
```

## Restricted expression surface

The initial authoring model does not permit arbitrary expressions such as:

```text
variant != null
a && b
x + 1
(a.ccc.ddd * 100 / pi) >= b.eee
list_items[index]
inline lambda
runtime type assertion
```

An exhaustive `switch` is a control-flow node rather than an invocation expression.
Its `switch` field references one already bound union, enum, or optional local.
Union cases are named by the variant types declared by that union, enum cases are named by qualified members declared by that enum, and optional cases are exactly `present` and `absent`.

Conditional execution consumes one previously produced boolean local rather than evaluating an inline expression.

Meaningful operations must be named functions, binding-origin access chains, dictionary constructors, transparent-record constructors, or limited control constructs defined by the runtime contract.

## Static validation

Before execution, the runtime must reject:

- unknown function IDs;
- unknown module capability slots;
- calls whose receiver is not an artifact module value;
- missing, extra, duplicate, or mixed-mode arguments;
- incompatible argument types, except for direct variant-to-declared-union assignment or a capability slot's explicitly declared and statically validated union-narrowing boundary;
- a module-bound union variant type that disagrees with the selected module registration;
- a dictionary whose key type is not a primitive scalar, primitive-backed named scalar, or enum;
- a dictionary constructor whose target types are unknown, unresolved, or invalid;
- a dictionary-constructor key or value whose type is not assignment-compatible with `K` or `V`;
- a statically evident duplicate dictionary key;
- a dictionary-constructor key or value containing an inline call, projection, nested constructor, or arbitrary expression;
- an access chain that does not begin with one visible binding;
- a field postfix whose receiver is not a transparent record or whose field is undeclared;
- a bracket postfix whose receiver is not `dict<K, V>`;
- a dictionary lookup key whose type is not exactly `K`;
- a dictionary lookup key containing an inline call, projection, constructor, arithmetic, or arbitrary expression;
- list indexing or bracket access on any non-dictionary value;
- a record constructor whose target is not one concrete transparent-record type;
- a record constructor with a missing, unknown, or duplicate field;
- a record-constructor field whose value is not assignment-compatible with the declared field type;
- a record-constructor field containing an inline call, projection, nested constructor, or arbitrary expression;
- invalid transparent field projections;
- direct projection of a union's active variant type;
- direct projection or implicit unwrap of `optional<T>` outside its `present` switch case;
- a switch whose target is not a previously bound local of a declared union, enum, or optional type;
- missing, unknown, duplicate, or default cases relative to the target union, enum, or optional case set;
- an enum switch case that uses an unqualified member or a member from another enum;
- an optional switch case other than `present` or `absent`, or any value-position reference to the target inside `absent`;
- access to an opaque field;
- duplicate step keys within one lexical scope, reassignment, or shadowing of an active ancestor binding;
- a `void` or `never` step whose key does not begin with `_`;
- a step whose key begins with `_` when its static result type is neither `void` nor `never`;
- the exact discard key `_`, or any reference or return that names a discard key;
- references to unbound locals, including forward references to later steps and child-scope locals after their scope ends;
- a reachable step after a `never` discard step in the same lexical path;
- a return value incompatible with the declaring function signature's single output type after applying direct variant-to-declared-union assignment compatibility.
