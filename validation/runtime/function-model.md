# Runtime function model

## Status

Working contract for callable units used by design-facing YAML.
The signature, context opening, generic declaration, native binding, stub authoring, generated caller facade, generated native binder, uniform invoker, initial registry-backed call forms, and opaque nominal-wrapper use in generated signatures are confirmed at the working-contract level. Execution-context extensions to the initial call boundary and non-opaque value-carrier mappings remain pending.

## Declaration detail

Every function declaration includes a non-empty `detail` field that states the function's design responsibility and result meaning.
The field is bootstrap-validated design authority and is not executable implementation logic.
It must not merely restate the signature or native binding.

## Function identity

Every callable function has one stable FunctionID. Ordinary functions use a qualified ID; the designated entry function uses the reserved unqualified ID `main`.

Examples:

```text
common.text.trim
common.graph.find_directed_cycles
artifact.task.reference.parse
operation.validation.check.exactly_one
```

Every namespace segment and ordinary terminal segment uses the lower-snake identifier grammar from `index.md`. An implementation-oriented terminal segment may begin with one leading underscore, such as `operation.validation._compose_record_checks`. Hyphens are prohibited in function IDs. The designated entry function is the reserved unqualified ID `main`.

Namespace ownership expresses the function's design meaning.
It does not express whether the implementation is native Go or a steps-defined composition.

## Required declaration

Every function declaration provides at least:

```text
function ID
optional generic type parameters
one signature string
implementation form
required execution-context dependencies
```

The signature string declares all named inputs and exactly one output type:

```text
(<input-name>: <type>, ...) -> <output-type>
```

Single-line signature content:

```yaml
functions:
  common.text.is_empty:
    detail: Determines whether a string contains no characters.
    signature: >-
      (value: string) -> boolean
    implementation:
      native: text.is_empty
```

Multiline signature content:

```yaml
functions:
  common.collection.compare_sets<T>:
    detail: Compares two collections as sets while preserving asymmetric differences.
    generic:
      T: equatable
    signature: >-
      (
        left: list<T>,
        right: list<T>
      ) -> SetComparison<T>
    implementation:
      native: collection.compare_sets
```

Every `signature` is authored as a YAML folded block scalar using `>-`, including signatures whose content fits on one line.
This avoids YAML plain-scalar ambiguity around the `: ` tokens inside named parameters while presenting one normalized signature string to the signature parser.

Every function signature declares exactly one output type.
A normal non-`void`, non-`never` function returns one typed value. `void` completes normally without an information-bearing value, while `never` has no normally continuing result path.
When several related values must be returned, the output type is one named transparent record rather than a language-level multiple return.

Generic type parameters are part of the function model because reusable functions such as collection comparison must preserve their supplied element type:

```text
list<T> + list<T> -> SetComparison<T>
```

A generic function lists its ordered type parameters in the declaration key and uses them in the signature:

```yaml
functions:
  common.collection.compare_sets<T>:
    detail: Compares two collections as sets while preserving asymmetric differences.
    generic:
      T: equatable
    signature: >-
      (
        left: list<T>,
        right: list<T>
      ) -> SetComparison<T>
    implementation:
      native: collection.compare_sets
```

The `generic` field is optional.
It is omitted when the function has no type parameters or when none of its declared parameters requires a constraint.
When present, it maps only constrained parameters directly to their constraint names.
Unlisted declared parameters are unconstrained.
The initial supported constraint vocabulary contains only `equatable`.
A `generic` entry whose name is absent from the function's `<...>` parameter list, an unsupported constraint, or a concrete call whose inferred type argument violates the constraint is a bootstrap error.

Input and output declarations are authoritative for static validation and module capability-slot matching.

## Implementation forms

The runtime supports three implementation forms:

```text
native
  implementation supplied by a precompiled Go binding

steps
  implementation composed from other declared functions

stub
  typed authoring placeholder with no executable implementation
```

All forms expose the same typed function contract.
A caller does not change its call syntax based on implementation form.

A stub is permitted only by the authoring validation profile. It must declare the complete function ID, `detail`, signature, context opening and requirements, and traceability metadata required for the declaration. Its `stub` value records the accepted implementation direction rather than a boolean presence marker.

The closed stub-intent vocabulary is:

```text
native
  The final implementation form is decided as native, but its binding and Go implementation are not yet authored.

steps
  The final implementation form is decided as steps-defined, but its PuppyDSL body is not yet authored.

undecided
  The function responsibility and signature are fixed, but native-versus-steps ownership still requires a design decision.
```

`stub: native` and `stub: steps` are valid without warning under authoring validation. `stub: undecided` is valid under authoring validation but emits one unresolved-implementation-direction warning. Every stub is non-executable and is rejected by design-closure and activation validation regardless of its intent value.

A native function is appropriate for parsing, Markdown AST handling, graph algorithms, repository lookup, or other operations whose detailed implementation should remain hidden.

A steps-defined function is appropriate when named semantic functions can express the processing without exposing low-level implementation operators.

The earlier term `graph function` described one possible derived representation rather than the authoring surface.
Human-authored YAML uses `steps`.
Steps execute in source order; a derived dependency graph may be retained for inspection, context analysis, or debugging but does not determine execution order.

## Steps-defined function behavior

A steps-defined function binds declared inputs and evaluates named steps in source order. It then returns one declared output value, completes normally as `void`, or terminates on every path as `never` according to its signature.

The declaration form is:

```yaml
functions:
  artifact.base.content.is_substantive:
    detail: Determines whether artifact content is neither empty nor a placeholder.
    signature: >-
      (body: string) -> boolean
    implementation:
      steps:
        normalized: common.text.trim(body)
        empty: common.text.is_empty(normalized)
        placeholder: artifact.base.content.is_placeholder(normalized)
        result: common.boolean.neither(empty, placeholder)
      return: result
```

`implementation.steps` is an ordered YAML mapping.
Each value is one permitted step form: a function call, binding-origin access chain, dictionary construction, transparent-record construction, `if`, exhaustive `switch` over a declared union, enum, or optional value, or bounded `for`.

A step whose static result type is neither `void` nor `never` uses an ordinary key and declares one immutable local binding.
A step whose static result type is `void` or `never` uses a discard key of the form `_<completion-name>`.
The discard key records the completed action for source order, diagnostics, and derived execution graphs, but it creates no local binding and must not be represented as a value asset.
The exact key `_` is invalid because a discard step still requires a non-empty semantic completion name.
Authors use a completion-oriented name such as `_served`, `_validated`, `_stopped`, or `_failed`; bootstrap validates the underscore/result-type contract and non-empty suffix rather than natural-language tense.

A step may reference function inputs, injected context values, active ancestor bindings, and ordinary locals already bound earlier in its current lexical scope.
A discard key is never a visible binding and cannot be referenced or returned.

For a steps-defined function whose output is neither `void` nor `never`, `implementation.return` is required and names exactly one input, injected context value, or already bound local.
Inside an `if` branch, switch case, or bounded iteration body, the corresponding branch or body `return` may name any binding visible in that child scope. In a union switch case, this includes the switch target under its case-narrowed static type. In an enum switch case, the target retains its declared enum type. In an optional `present` case, the target may be returned as `T`; in `absent`, the target is not a visible value binding.
Inline calls, projections, or other expressions are not permitted in `return`; the value must first be bound by a step.
The returned value must be assignment-compatible with the function signature's output type. Exact type identity is required except that a direct variant of a declared union may be returned where that union output type is declared.

For a `void` function, `implementation.return` is omitted and successful completion of the final step completes the function:

```yaml
functions:
  runtime.watcher.stop:
    detail: Stops one repository watcher and releases its native resources.
    signature: >-
      (watcher: RepositoryWatcher) -> void
    implementation:
      steps:
        _stopped: runtime.watcher.stop_native(watcher)
```

The `_stopped` key is a discard step label rather than a `void` local.
Execution continues after the call, but no design-facing binding is added to the lexical namespace.

A function with output type `never` has no normally continuing result path.
Its invocation is a terminal discard step in the current lexical path and uses a completion-oriented key such as `_failed`.
No local value is established, and no reachable step may follow it in that lexical path.
A steps-defined `never` function omits `implementation.return`, and every reachable path through its steps must terminate through another `never` call.
The initial native function `runtime.throw(message: string) -> never` raises one execution failure with the supplied message.
PuppyDSL provides no catch, recovery, or inspection construct for that failure.

`implementation.steps` must contain at least one step.
A declaration must not combine the steps form with `implementation.native` or `implementation.stub`.

Stub declarations use one accepted intent value:

```yaml
functions:
  current_records.discover_repository_inventory:
    detail: Discovers the registered repository paths used by one operation.
    signature: >-
      () -> RepositoryInventory
    implementation:
      stub: native
```

```yaml
functions:
  operation.validation._compose_record_checks:
    detail: Produces one validation result from declared semantic check functions.
    signature: >-
      (record: CurrentRecord) -> RecordValidationResult
    implementation:
      stub: steps
```

```yaml
functions:
  operation.search._prepare_query:
    detail: Prepares one exact search query for operation execution.
    signature: >-
      (query: SearchQuery) -> PreparedSearchQuery
    implementation:
      stub: undecided
```

`implementation.stub` must be exactly one of `native`, `steps`, or `undecided`. Boolean values, arbitrary strings, and other YAML scalar forms are invalid. A stub has no `steps`, `return`, or `native` field. The intent is authoring metadata only; it does not register a native binding, create an empty steps body, or make the function executable.

## Context ownership and dependencies

A function may declare an ordered list of context IDs that it opens for its complete invocation:

```yaml
functions:
  operation.validation.validate_records:
    detail: Validates selected records within one operation-scoped dependency set.
    opens:
      - repository_snapshot
      - operation
    signature: >-
      (selectors: list<RecordSelector>) -> ValidationResult
    implementation:
      native: operation.validation.validate_records
```

Before the function body begins, the runtime opens each listed context in source order. The first listed context is outermost and the last is innermost.
The function body and every nested function call inherit the resulting active context chain.
The runtime closes newly opened contexts in reverse order after return or failure, invoking registered destructors and discarding each context-instance-local cache.

A function without `opens` opens no context and inherits the contexts already active at its call site.
Context opening and closing are prohibited inside individual steps, conditions, and iterations.
If a smaller lifecycle is required, it must be represented by a separate function with its own `opens` list.

A restartable entry function may additionally declare runtime events that restart its complete invocation after completion:

```yaml
restart_on:
  - repository.changed
max_restarts: 1
```

This event-driven restart contract is defined in `event-model.md`.
It never restarts an individual step or mutates an active context generation.

A function explicitly declares any context-provided value that is not supplied as an ordinary design input:

```yaml
requires:
  - operation.current_identity_index
```

This is the runtime's dependency-injection boundary.
A required value may be supplied by a context opened by the current function or by an inherited ancestor context.
Hidden access to arbitrary global state is prohibited.

Possible context-provided values include:

- configured current-record state;
- current identity index;
- selected artifact tree;
- repository source access;
- artifact module catalog.

Context ownership and dependencies are indexed as design data so tooling can answer questions such as:

```text
which functions open the operation context?
which functions require operation.current_identity_index?
which entry functions transitively depend on runtime.repository_source_access?
```

Transitive requirements are derived from called functions, including recursive function groups.

## Invocation and output assignment

A call whose output is neither `void` nor `never` produces one typed value and uses an ordinary step key that becomes its local name.
A `void` or `never` call uses a discard key beginning with `_` and creates no local binding.

Conceptual forms:

```yaml
normalized: common.text.trim(body)
_served: drmcp_runtime.mcp.serve()
_failed: runtime.throw(message)
```

A visible transparent-record or dictionary value permits the statically typed binding-origin access chains defined by `invocation-model.md`, including record field projection and required-key dictionary lookup.
An opaque output may only be passed to other compatible functions.
A derived dataflow graph may retain `_served` or `_failed` as an action or terminal node, but it must not emit a corresponding value asset.

## Native bindings

A native implementation declaration resolves one native binding ID registered by the Go runtime:

```yaml
functions:
  common.text.trim:
    detail: Removes leading and trailing whitespace before semantic interpretation.
    signature: >-
      (value: string) -> string
    implementation:
      native: text.trim
```

The design-facing function ID and native binding ID need not be identical.
The `implementation.native` value is a registry ID, not a Go import path, package name, or source symbol.
Design YAML cannot import arbitrary Go packages or identify executable code directly.

Exactly one implementation form is permitted.
A declaration must not combine `implementation.native` with `implementation.steps` or `implementation.return`.

During bootstrap, the runtime resolves the binding ID from the activated native binding registry and verifies that the binding's declared parameter and result types exactly match the design-facing function signature, including generic constraints when present.
An unresolved binding, duplicate binding registration, or incompatible signature is a bootstrap error.

The native binding registry contains functions compiled into an activated native runtime artifact.
Native Go implementations live in source modules separate from YAML declarations and may own ordinary `go.mod` and `go.sum` files.
External libraries are dependencies of those native extension modules.
The accepted automated build-and-activation direction is defined in `native-extension-model.md`.

### Generated Go caller and binder

For every declared function, the build tool generates a typed Go caller facade in the package derived from the function namespace:

```text
FunctionID: operation.discovery_and_listing.discover_record_scopes
Go package: <go-native-module>/puppygen/operation/discovery_and_listing
Go symbol:  DiscoverRecordScopes
```

The reserved entry function `main` generates `Main` in the root `<go-native-module>/puppygen` package.

The root package exports `Functions map[string]any`, containing every non-generic declared FunctionID exactly once and mapping it to its generated typed caller facade. The generated source literal orders entries by FunctionID; Go map iteration order remains unspecified. The catalog includes non-generic native, steps-defined, stub, and `main` declarations; it indexes declared callable contracts and does not determine executability, visibility, native binding ownership, or transport publication. Generic functions are excluded because Go requires concrete type arguments before a generic function can become a function value. Consumers may use each catalogued concrete facade function for Go reflection or typed adapter construction.

The caller facade exists for `native`, `steps`, and `stub` declarations. It represents the declared contract rather than implementation availability. Calling a stub facade is valid Go source, but design-closure and activation validation still reject a reachable stub before execution.

For `implementation.native`, the same generated package additionally provides:

```go
type ServeImplementation func(native.Call) error
func BindServe(implementation ServeImplementation) native.Binding
```

The generated `Bind<FunctionName>` function fixes the design-facing FunctionID, native binding ID, and expected Go implementation signature from the PuppyDSL declaration. Passing a Go function with an incompatible concrete signature is a Go compile error.

When a function signature contains an opaque type, the generated caller and implementation signature use that opaque TypeID's distinct generated nominal wrapper.
They do not substitute the raw bound Go representation.
Two opaque types bound to the same representation therefore remain different Go parameter and result types, and a handwritten implementation cannot exchange them without an explicit TypeID-specific generated access boundary.
Final function generation is prohibited until every referenced opaque TypeID has exactly one valid native representation binding.

The binder wraps the typed implementation in one uniform native invoker with the runtime shape `func(Call, []any) (any, error)`. The generated wrapper validates the runtime argument count, converts each argument to its declared generated Go type, invokes the typed implementation, and returns its result through the uniform carrier. Internal uniform transport does not permit generated opaque carriers to use `any`; each opaque value reaches this boundary through its exact nominal wrapper. A nil typed implementation produces a binding with a nil invoker so registry construction rejects it before activation. Mixed steps/native execution preserves the current runtime `Call` boundary when dispatching into the invoker, allowing generated facade calls made by native code to re-enter the same steps-aware execution state.

A native implementation may place its binding declaration immediately beside the implementation:

```go
func Serve(call native.Call) error {
    return nil
}

var serveBinding = puppy_mcp.BindServe(Serve)
```

The exact generated-file split and binding aggregation mechanism are build-tool implementation details. A handwritten central registry and import-side-effect registration through `init()` are not required by the language contract.

## Function registry

The runtime maintains one generic function registry:

```text
FunctionID -> declared callable function
```

The registry contains common, artifact/base, artifact-specific, and operation functions.
Declared events also contribute generated typed callables to the same callable namespace.
A generated event callable has no ordinary implementation body and dispatches through the runtime event system.
Event IDs and ordinary function IDs must therefore be unique across the shared callable namespace.

Manifest function IDs are resolved once during bootstrap.
Normal execution uses resolved handles and must not repeat dynamic string-based lookup for each call.

## Visibility convention

The initial declaration contract has no `public` or `private` field.
A function ID segment beginning with `_` marks an implementation-oriented function by convention:

```text
artifact.task.reference._parse_sequence_segments
```

The initial runtime does not need to enforce access restrictions.
A later enforcement rule may use the existing naming convention without changing every function declaration.

## Static and bootstrap validation

Before operation execution, the runtime must reject:

- function IDs, context IDs, event IDs, module IDs, fields, inputs, locals, or enum members that violate the identifier grammar in `index.md`;
- duplicate function IDs or conflicts between ordinary function IDs and generated event-callable IDs;
- malformed signatures;
- duplicate parameter names;
- missing output type declarations;
- unknown type references;
- final generation or activation with an unbound, multiply bound, dynamically bound, or `any`-backed opaque TypeID;
- a generated native caller or binder that substitutes one opaque TypeID's raw representation or another opaque TypeID for its nominal wrapper;
- `never` used anywhere except a function output;
- unknown context IDs in `opens`;
- duplicate context IDs within one `opens` list;
- reopening a context ID already active on a reachable entry route;
- an `opens` order incompatible with provider ancestor dependencies;
- unknown context-provided dependency IDs;
- steps that call a missing function under every validation profile;
- an invalid or incomplete stub declaration;
- a stub reachable under design-closure or activation validation;
- unknown input names;
- incompatible argument types after applying direct variant-to-declared-union assignment compatibility;
- invalid dictionary constructors;
- invalid binding-origin access chains, including incompatible dictionary keys, non-dictionary bracket access, and undeclared transparent fields;
- invalid transparent-record constructors;
- a switch whose target is not a previously bound local of a declared union, enum, or optional type;
- a switch with missing, unknown, duplicate, or default cases relative to the target union, enum, or optional case set;
- a union switch case whose named type is not a variant of the target union;
- an enum switch case whose member is unqualified, unknown, or declared by another enum;
- an optional switch whose cases are not exactly `present` and `absent`, or whose `absent` case uses the target as a value;
- incompatible result types among continuing switch cases;
- references to unbound locals, including forward references to later steps or child-scope locals after their scope ends;
- duplicate step keys within one lexical scope;
- a `void` or `never` step whose key does not begin with `_`;
- a step whose key begins with `_` when its static result type is neither `void` nor `never`;
- the exact discard key `_`, a reference to a discard key, or a `return` naming a discard key;
- a function call whose required context is unavailable from the caller's active context chain or contexts opened by the callee;
- an entry function whose transitive call graph contains an unsatisfied context requirement;
- unknown event IDs in an event-driven restart declaration;
- unbounded or unsafe restart declarations;
- a steps-defined `never` function with any reachable path that does not terminate through a `never` call;
- a reachable step after a `never` call in the same lexical path;
- returned values incompatible with the signature's single output type after applying direct variant-to-declared-union assignment compatibility;
- direct or indirect cycles in context provided-value dependencies reachable from the selected entry routes.

## Binding and recursion boundary

A step may reference function inputs, injected context values from the active context chain, active ancestor bindings, and ordinary locals bound by earlier steps in its current lexical scope.
Discard step keys beginning with `_` do not enter that binding namespace.
Forward references are invalid as unbound-local references, so a separate local-step cycle rule is unnecessary.
Each lexical scope has one binding namespace. Sibling control-flow scopes may reuse a local name because neither binding is visible in the other scope.
Shadowing or reassignment of an active input, context binding, iteration binding, switch target, condition binding, or local is prohibited.

The initial language profile prohibits cycles in the YAML-defined function call graph.
Direct recursion and mutual recursion are bootstrap errors.
Native Go implementations may use recursion internally as an implementation detail.

This restriction is a bootstrap policy rather than a syntax dependency.
A future runtime could permit the same existing call syntax recursively by replacing cycle rejection with a dedicated recursion-safety policy, without changing ordinary function declarations or invocations.
The current profile intentionally keeps recursive tree traversal and similar algorithms in native Go because recursive YAML composition is harder to review and no accepted consumer requires it.

## Failure behavior

Steps-defined YAML functions describe successful design-facing data flow only.
They do not declare `raise`, `catch`, `except`, `try`, or generic error-value branching.

An expected business, lookup, parsing, or validation outcome remains a declared typed value and may participate in ordinary semantic branching.
A closed expected outcome may be declared as a union and consumed by exhaustive type switch without outcome-specific native predicate or extraction functions.
Examples include malformed input, unresolved reference, conflicted identity, absent optional data, and validation findings when those outcomes are expected parts of the operation contract. Declared unions use exhaustive variant switching, while `optional<T>` uses exhaustive `present` / `absent` switching and exposes `T` only in `present`.

Detailed recovery behavior for execution failures belongs inside a native Go implementation when required.
A native function may retry, translate implementation-specific errors, select a fallback, or perform compensating cleanup before either returning its declared successful value or reporting execution failure to the runtime.
If a composed workflow requires such behavior, it calls one semantic native wrapper that owns the complete recovery boundary rather than reproducing error plumbing in YAML.

An unhandled execution failure automatically aborts the current steps-defined invocation and propagates through callers to the entry boundary.
A call to a native `never` function deliberately initiates the same propagation path and has no normally continuing result.
The runtime owns only generic propagation and structured context closing; it does not expose catch or recovery to PuppyDSL.
A consuming host may report the terminal failure after `main` exits unsuccessfully, but the failure is not encoded as a generic YAML-visible `Result<T, error>` that every graph must inspect.

Native implementations must not convert unexpected execution failures into ordinary validation findings or other successful domain outcomes.

## Excluded language features

The initial function model does not include:

- method declaration on YAML-defined types;
- class inheritance;
- arbitrary operator or comparison expressions;
- non-exhaustive, defaulted, or open-ended switches;
- switching over strings, integers, named scalars, or other open-ended values;
- arbitrary runtime type assertions;
- lambda expressions;
- dynamic function construction;
- repeated string-based dispatch during normal execution;
- implicit coercion between named types and their primitive representations;
- raw native-representation substitution for an opaque parameter or result.
