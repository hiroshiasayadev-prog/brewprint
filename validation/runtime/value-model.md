# Runtime value model

## Status

Working contract for values exchanged by native and steps-defined functions.
Transparent-record declaration syntax, zero-field transparent-record semantics and Go representation, and the opaque nominal-wrapper boundary are confirmed; other exact Go value-carrier mappings remain pending.

## Declaration detail

Every type declaration includes a non-empty `detail` field that states the design meaning and responsibility of the declared type.
The field is design authority and bootstrap-validated declaration data; it is not executable implementation logic.
It must describe the type's semantic role without restating its field list or native representation.

## Principles

- The runtime uses typed values rather than dynamic `any` objects.
- Native and steps-defined functions exchange the same declared types.
- There is no implicit type conversion except direct assignment of a declared union variant to that union type.
- YAML may inspect only fields explicitly exposed by transparent record types.
- A union's active variant type is not an ordinary projectable value; YAML may consume a declared union only through the exhaustive type-switch construct or a declared runtime dispatch boundary.
- An `optional<T>` value is consumed through exhaustive `present` / `absent` switching; absence is control state rather than a projectable value or null literal.
- Expected domain outcomes are typed values; runtime execution failures are not normal function outputs.

## Primitive and control-result types

The initial primitive set is:

```text
string
boolean
integer
```

The design-facing signature language also includes:

```text
void
never
```

`void` means that a function completes normally without producing an information-bearing design value.
It is used for lifecycle operations such as stopping a watcher or closing a native resource.
`void` is not `null`, `nil`, or an absent `optional<T>` value.
The runtime may represent successful `void` completion with one internal unit-like value, but `unit` is not part of the design-facing YAML vocabulary and no direct `void` step creates a design-facing local binding.
Because `void` remains a valid type argument, a declared function boundary may still expose a composite value such as `list<void>`; the composite output is information-bearing as a whole even though its elements carry only successful-completion positions.
Bounded iteration cannot construct that list by returning direct `void` step results because discard steps create no returnable local.

`never` means that a function call has no normally continuing result path.
A function returning `never` terminates the current runtime invocation by propagating an execution failure and never produces a bindable value.
A `never` call is terminal within its current lexical path and does not participate in continuing-branch result-type agreement.
`never` is not `void`, is not a constructible value, is not a valid generic type argument, and cannot appear as a transparent-record field, union variant, context-provided value, event payload, or ordinary function input.
Its initial concrete consumer is `runtime.throw(message: string) -> never`.

`float`, generic `number`, and raw `bytes` are not currently required.

## Named scalar types

A named scalar preserves a distinct meaning while using one primitive representation.

Type declarations are grouped beneath the `types` root mapping. The declaration identity is the mapping key:

```yaml
types:
  ArtifactKind:
    detail: Identifies one artifact kind registered by an artifact module.
    scalar: string

  AppNamespace:
    detail: Identifies one configured application namespace.
    scalar: string
```

The `scalar` value names exactly one supported primitive representation.
The declared named scalar is a distinct type from that primitive and from every other named scalar using the same representation.
There is no implicit conversion in either direction.

A function expecting `string` cannot receive `ArtifactKind`, and a function expecting `ArtifactKind` cannot receive a raw `string`.
When a conversion is meaningful, one explicitly declared function must own it.
Named scalars expose no fields and do not close the set of permitted runtime values.

Current candidates include:

```text
ArtifactKind
ArtifactModuleID
FunctionID
AppNamespace
DomainNamespace
```

`ArtifactKind` is an open string-backed named scalar.
Artifact-module catalog registration determines which artifact kinds are currently supported.
Adding an artifact kind must not require changing a runtime-core enum.

Unknown primitive representations, generic parameters on a named scalar, and declarations that combine `scalar` with another type body are bootstrap errors.
The YAML declaration is authoritative; a native Go named type or wrapper must conform to it.

## Enumerations

An enum is used only for genuinely closed runtime vocabulary.

The declaration form is:

```yaml
types:
  ArtifactStructure:
    detail: Identifies the closed record-structure model used by an artifact module.
    enum:
      - tree
      - sequential
```

One enum declaration creates one type.
Every declared member is a value of that enum type rather than a separate member-specific type.
Enum values are referenced with a type-qualified literal:

```yaml
structure: ArtifactStructure.tree
```

For example:

```text
ArtifactStructure.tree
ArtifactStructure.sequential
```

are both values of type `ArtifactStructure`.
A different enum remains a distinct type even when it declares a member with the same spelling.
An enum is also distinct from `string`; raw strings and enum values are not implicitly interchangeable.

Enums are nongeneric and expose no fields.

A declared enum may be consumed by the exhaustive `switch` construct defined by the control-flow model.
Each case key uses one qualified member literal such as `ArtifactStructure.tree`, and every member declared by the target enum must appear exactly once.
A `default` case, an unqualified member, an unknown member, or a member from another enum is invalid.
Inside an enum case, the switch target retains its enum type; selecting one member does not create a member-specific type or a new binding.
This closed enum switch does not permit switching over strings, integers, named scalars, or other open-ended values.

An empty member list, duplicate member, unknown member reference, generic parameter list, or declaration combining `enum` with another type body is a bootstrap error.

Current candidates include:

```text
ArtifactStructure
  tree
  sequential

FindingClassification
  contract_violation
  advisory
```

Artifact-specific status or task-type vocabularies do not belong to runtime core unless they are intentionally declared as their own closed artifact-owned enum types.

## Generic composite types

The initial generic composites are:

```text
list<T>
optional<T>
dict<K, V>
```

`list<T>` preserves source or semantic ordering unless the consuming function explicitly treats it as a set.

`optional<T>` distinguishes absence from an empty or invalid value.
A runtime optional contains either one present value of exact type `T` or absence; it does not contain an untyped `null` value.
Design-facing YAML consumes an optional only through the exhaustive `switch` form defined by `control-flow.md`.
Inside the `present` case, the switch target keeps its binding name and is statically narrowed from `optional<T>` to `T`.
Inside the `absent` case, absence is statically established and the switch target is unavailable in value positions because no `T` value exists.
After the switch, the original outer binding remains `optional<T>`.

The initial invocation grammar provides no `none`, `some`, `null`, or other optional literal or constructor.
Optional values are produced by declared functions, context providers, generated runtime facilities, or other typed boundaries and may be passed or returned without elimination.
Runtime-internal logic must not use a general untyped `null` as a substitute for optionality.
External operation projection may still render an absent or unavailable value as JSON `null` when the operation contract requires it.

`dict<K, V>` preserves one immutable association from each unique key of type `K` to one value of type `V`.
Its key restrictions, equality, and construction rules are defined below.

## Dictionaries

A dictionary type is written as `dict<K, V>`.
The key type `K` must resolve to exactly one of:

- a primitive scalar: `string`, `integer`, or `boolean`;
- a named scalar whose declared primitive representation is `string`, `integer`, or `boolean`;
- a declared enum type.

A union, list, optional, dictionary, transparent record, opaque type, `void`, or `never` is not a valid dictionary key type.
This is a built-in restriction of `dict<K, V>` rather than a user-authored generic constraint.

Dictionary key equality preserves declared type identity as well as value identity.
A named scalar key is distinct from its primitive representation and from every other named scalar using the same representation.
Values from different enum types are likewise distinct even when their member spellings match.
No implicit conversion is applied when a dictionary key is supplied.

A dictionary contains at most one entry for each key.
Entry order has no semantic meaning and must not be used as source order.
A dictionary value is immutable after construction; the initial language profile provides no insertion, replacement, deletion, or mutation expression.

The invocation model provides required-key lookup through a statically typed access chain such as `inventory.apps[app_namespace]`.
A lookup key must have exactly the dictionary's declared key type, and a missing runtime key terminates the current invocation as an execution failure rather than producing `optional<V>` or another implicit absence value.
The control-flow model provides bounded dictionary iteration with one statically typed key binding and one statically typed value binding.
Higher-level search, aggregation, transformation, or optional lookup remains the responsibility of declared functions when a concrete consumer requires it.

Although dictionary entry order has no semantic meaning, bounded iteration uses one canonical deterministic key order:

- `string` keys use ascending lexical order by exact Unicode scalar-value sequence;
- `integer` keys use ascending numeric order;
- `boolean` keys use `false` before `true`;
- named scalar keys use the ordering of their declared primitive representation while preserving their named type identity;
- enum keys use declaration order.

This canonical order exists only for deterministic execution and output construction. It does not make insertion order or source declaration order observable dictionary data.

Design-facing YAML may construct one dictionary value with the initializer syntax defined by `invocation-model.md`:

```yaml
directories: >-
  dict<PhysicalDirectoryName, PhysicalDirectoryTree>{
    [directory_name]: directory_tree
  }
```

The explicit `dict<K, V>` target gives an empty initializer such as `dict<K, V>{}` a complete static type.
Each key and value is checked against the declared `K` and `V` types.
A duplicate key must never silently replace an earlier entry: statically evident duplicates are bootstrap errors, while equality discovered only from runtime binding values is an execution failure during construction.

The normal Go representation may use `map[K]V`, but Go representation does not define the PuppyDSL type identity, key eligibility, visibility, or mutability contract.
A native-internal Go map is not automatically a design-facing `dict<K, V>` value.

## Transparent records

A transparent record exposes a statically declared field set to YAML.

A transparent record may declare zero fields with `record: {}`:

```yaml
types:
  CandidateNotFormed:
    detail: Indicates that one readable Markdown source did not form a listing candidate under the selected artifact module contract.
    record: {}
```

A zero-field transparent record is an information-bearing nominal singleton type. Its declared type identity is the complete design value, all instances of that exact type are observationally equivalent, and no field projection is available.
It is distinct from `void`, from absence in `optional<T>`, and from every other zero-field transparent record type.
It may be bound, passed, returned, stored in a composite value, or used as a union variant wherever an ordinary transparent-record value is accepted.
The runtime and generated APIs must not collapse it into a shared unit value or a no-result completion marker.

A transparent-record declaration graph may be recursive only when every type-reference cycle crosses at least one `list<T>` or `dict<K, V>` type constructor.
A record field that directly names its own record type is invalid, and a group of records that refer to each other only through direct record fields is likewise invalid.
`optional<T>` does not establish an allowed recursive boundary in the initial language profile.

For example, this recursive shape is valid:

```yaml
types:
  PhysicalDirectoryTree:
    detail: Preserves one nested physical directory structure.
    record:
      directories: dict<PhysicalDirectoryName, PhysicalDirectoryTree>
      markdown_sources: list<MarkdownSourceFileName>
```

These shapes are invalid:

```yaml
types:
  DirectNode:
    detail: Invalidly embeds itself as one direct field.
    record:
      child: DirectNode

  OptionalNode:
    detail: Invalidly uses optionality as its only recursive boundary.
    record:
      child: optional<OptionalNode>
```

The language exposes no pointer type, nullable reference, address identity, or arbitrary reference-graph construction.
A design value requiring pointer-owned topology, shared node identity, cycles outside collection containment, or another native reference model must be declared as an opaque type and manipulated through declared functions.

A generic transparent record declares its ordered type parameters in the declaration key:

```yaml
types:
  SetComparison<T>:
    detail: Preserves asymmetric differences and equality from a set comparison.
    generic:
      T: equatable
    record:
      left_only: list<T>
      right_only: list<T>
      equal: boolean
```

The `generic` field is omitted when no parameter requires a constraint.

A nongeneric transparent record omits `generic` and declares no type-parameter list:

```yaml
types:
  SourceLocation:
    detail: Identifies the physical source position of one represented occurrence.
    record:
      path: string
      line: integer
      column: integer
```

The declaration key identifies the type and declares the ordered generic-parameter list when present.
A declaration without type parameters omits both the `<...>` suffix and the `generic` field.
When a parameter requires a capability constraint, `generic` maps that parameter directly to the constraint name:

```yaml
generic:
  T: equatable
```

Parameters without constraints are omitted from `generic`, even when other parameters are constrained.
The initial constraint vocabulary contains only `equatable`.
Every `record` entry declares one field name and one type expression.
Unknown types, unknown generic parameters, duplicate fields, unsupported constraints, and type arguments that violate a declared constraint are bootstrap errors.

Design YAML may project a declared field:

```yaml
comparison: common.collection.compare_sets(declared, physical)
missing: comparison.right_only
```

An unknown field reference is a bootstrap or static recipe error.
Transparent access is not runtime reflection.

Field projection does not expose undeclared implementation fields, and the runtime does not infer additional fields from a native Go struct.
The YAML type declaration is the design-facing authority for the transparent shape.

Design-facing YAML may construct a transparent-record value with an explicit record constructor:

```yaml
domain: >-
  SequentialDomainDiscovery{
    domain_namespace: domain_namespace,
    source_root: source_root
  }
```

A zero-field transparent record uses the same constructor with an empty field set:

```yaml
not_formed: >-
  CandidateNotFormed{}
```

The constructor target must resolve to one concrete transparent-record type. A generic transparent record requires one concrete type instantiation.
Every constructor is authored as one YAML folded scalar using `>-`; the dedicated expression parser interprets the type name, braces, field bindings, commas, literals, and binding references inside that scalar.

A constructor must provide every declared field exactly once and no undeclared field. Field order has no semantic effect, though authors should normally follow declaration order. A trailing comma is optional.
Each field value is exactly one literal or one already visible binding reference. Inline calls, field projections, nested constructors, and arbitrary expressions are prohibited; the caller must bind such a value in an earlier step.
Each supplied value must be assignment-compatible with the declared field type. Exact type identity is required except that a direct variant of a declared union may be assigned to a field of that union type.
The constructed local has the exact transparent-record type named by the constructor.

Named scalars, enums, unions, opaque types, primitives, and generic composite types do not use record-constructor syntax.

Likely transparent records include:

- `SourceLocation`;
- `HeadingOccurrence`;
- `SectionOccurrence`;
- `DuplicateGroup<T>`;
- `SetComparison<T>`;
- `CandidateNotFormed`.

## Opaque named types

An opaque type may be passed to declared functions but its internal fields are unavailable to YAML.

The declaration form is:

```yaml
types:
  MarkdownDocument:
    detail: Represents parsed Markdown whose internal syntax tree is hidden from YAML.
    opaque: true
```

An opaque type is a distinct nongeneric type.
Design-facing YAML cannot project fields, construct a literal value, or build an instance through record syntax.
Opaque values are produced by native functions, context providers, generated runtime facilities, or other declared functions and may be passed only where the exact opaque type is accepted.
There is no implicit conversion between opaque types or between an opaque type and any scalar, enum, or record type.

Every opaque declaration requires exactly one application-owned native representation binding before final Go API generation and activation.
The binding selects one compile-time Go representation type for that opaque TypeID.
An absent binding, duplicate binding, binding for an unknown or non-opaque TypeID, or fallback to `any` or `interface{}` is a generation failure.

Puppygen preserves opaque TypeID identity by generating one distinct nominal Go wrapper for each opaque declaration.
The wrapper contains the bound representation with its exact compile-time Go type; it is not a type alias for that representation and must not contain an `any` field.
Two opaque declarations may bind the same native representation while remaining non-interchangeable:

```text
ArtifactReference -> *yaml.Node
ManifestDocument  -> *yaml.Node
```

Conceptually, final generation still produces distinct types:

```go
type ArtifactReference struct {
    value *yaml.Node
}

type ManifestDocument struct {
    value *yaml.Node
}
```

A value of one generated wrapper cannot be passed where the other wrapper is required even when both contain the same native representation.
Generated native caller and binder signatures use the nominal opaque wrapper rather than the raw bound representation, so Go compilation preserves PuppyDSL opaque identity across handwritten native implementations.
Generated typed construction and representation-access operations, or an equivalent compile-time API, permit application-owned native code to wrap and access the bound representation without reflection or dynamic casts.
The exact exported operation names and generated-file split remain build-tool details.

A declaration that combines `opaque: true` with `scalar`, `enum`, `record`, a generic parameter list, or another type body is a bootstrap error.
The YAML declaration establishes the design-facing identity and visibility boundary; the application-owned representation binding selects the hidden Go representation, and puppygen generates the nominal carrier that joins those two contracts.

Current opaque candidates include:

```text
MarkdownDocument
MarkdownFragment
ArtifactReference
ArtifactModule
CurrentRecordIndex
ArtifactTree
```

Allowed:

```yaml
document: artifact.base.source.parse(source)
headings: common.markdown.heading_occurrences(document)

module: artifact.base.module.from_base_reference(base_ref)
parse_outcome: module.reference.parse(base_ref)
```

Prohibited:

```yaml
kind_token: base_ref.discriminator
kind: reference.kind
variant: reference.value
```

A meaningful named function must expose any operation that design YAML needs to perform on an opaque value.

## Unions

The initial structure-specific base-reference variants are transparent routing-evidence records:

```yaml
types:
  TreeBaseReference:
    detail: Preserves tree-reference routing evidence before artifact-specific grammar confirmation or current-tree resolution.
    record:
      raw: string
      record_kind: TreeReferenceDiscriminator
      app_namespace: string
      path_segments: list<string>

  SequentialBaseReference:
    detail: Preserves sequential-reference routing evidence while leaving artifact-specific segment meaning to the selected artifact module.
    record:
      raw: string
      app_namespace: string
      artifact_kind: SequentialReferenceDiscriminator
      domain_namespace: string
      artifact_segments: list<string>
```

`artifact_segments` preserves the source order of every token after the sequential domain namespace. Shared routing does not assign artifact-specific names or validate their count, width, role, or meaning.

`path_segments` preserves canonical tree identity segments after the app namespace. It does not identify a physical final file or distinguish a directory-backed node from a non-index-file-backed node.

A union declares one closed list of variant types:

```yaml
types:
  BaseReference:
    detail: Carries one structure-specific base reference for module routing.
    union:
      - TreeBaseReference
      - SequentialBaseReference
```

Each list item names one separately declared variant type. The list is nonempty and contains no duplicate type.
A runtime union value has exactly one active variant type and one value of that type. The active variant type is interpreter-owned runtime type information rather than a separately declared tag or projectable field.

A union is a distinct type from every listed variant.
A value whose static type is one directly listed variant is assignment-compatible with that union wherever the expected target type is explicitly known. This includes function inputs, transparent-record fields, declared function returns, context-provider arguments, and other statically typed value positions. The runtime preserves the supplied value and records its variant type as the union's active variant.

This variant-to-union rule is the only implicit union assignment. It does not permit union-to-variant downcast, conversion between unrelated unions, or assignment through a transitive variant relation. It also does not infer a union merely because independent control-flow branches or list elements have different variant types; those constructs still follow their own result-type rules unless an explicit union-typed target exists.
Design-facing YAML cannot construct a union literal, project the active variant type as data, use `isinstance`, or perform an arbitrary runtime type assertion.
A boolean type predicate alone cannot narrow the static type of a local.

A declared union may be consumed by the language-level exhaustive type-switch construct.
Each switch case names one declared variant type. Inside that case, the switch target local itself is statically narrowed to the named variant type; no separate payload binding is introduced.
Every declared variant type must appear exactly once. Missing, unknown, or duplicate cases are bootstrap errors.
No `default` case is permitted because exhaustiveness is checked against the closed union declaration.

A variant type may itself be a transparent record, opaque type, enum, or named scalar and may appear normally in declared function signatures.
A trusted runtime boundary may also narrow a union when that boundary's declaration proves the required variant.
The initial runtime-bound example is module-bound artifact dispatch: an artifact module registered as `tree` or `sequential` may receive the corresponding `TreeBaseReference` or `SequentialBaseReference` variant from a `BaseReference` value.
This narrowing is interpreter behavior driven by the validated module structure and capability-slot signature and does not require artifact-specific Go branching.

Shared reference routing uses a second union so success/failure branching remains separate from tree/sequential branching:

```yaml
RoutedBaseReference:
  record:
    reference: BaseReference

UnrouteableReference:
  record:
    raw: string

SharedRoutingOutcome:
  union:
    - RoutedBaseReference
    - UnrouteableReference
```

`UnrouteableReference` needs no separate reason value. Its variant identity means that the exact input matches neither supported routing envelope.

Artifact-specific canonical parsing uses another concrete union:

```yaml
ArtifactReference:
  opaque: true

InvalidArtifactReference:
  record:
    raw: string

ArtifactReferenceParseOutcome:
  union:
    - ArtifactReference
    - InvalidArtifactReference
```

`InvalidArtifactReference` means that routing succeeded but the exact input does not satisfy the selected artifact module's complete canonical reference contract. It likewise needs no finer reason value.

Unknown variant types, duplicate variant types, an empty union, generic parameters, or declarations that combine `union` with another type body are bootstrap errors.
The initial type-switch feature does not change the current prohibition on generic union declarations; that question remains separate from union consumption.

### ArtifactReference native representation

`ArtifactReference` remains one design-facing opaque type. Its native representation is structure-specific:

```text
ArtifactReference
  = TreeArtifactReference
  | SequentialArtifactReference

TreeArtifactReference
  base: TreeBaseReference

SequentialArtifactReference
  base: SequentialBaseReference
  fields: map<string, string>
```

`SequentialArtifactReference.fields` contains only artifact-specific named components that are not represented as named fields by `SequentialBaseReference`.
Its keys are the artifact-owned sequential component names declared by the module, such as `sequence`, `work_sequence`, or `task_sequence`, and each value preserves the corresponding canonical segment text.
A successful parser assigns every `base.artifact_segments` entry to exactly one declared field in source order and produces no missing or additional field.
For example, a Task parser may produce:

```text
base.artifact_segments: ["016", "01"]
fields:
  work_sequence: "016"
  task_sequence: "01"
```

The map is a native implementation structure, not the design-facing `dict<K, V>` value declared above. YAML cannot inspect or construct this opaque field merely because dictionaries exist elsewhere in the language.

`TreeArtifactReference` contains only `base: TreeBaseReference` because the base value already represents the complete canonical tree identity through `record_kind`, `app_namespace`, and `path_segments`.
It does not contain an empty artifact-fields map or any additional tree payload.

The base discriminator remains sufficient for artifact-module lookup. `ArtifactReference` therefore requires no duplicate top-level artifact-kind or structure field.
Ordinary design YAML continues to consume `ArtifactReference` only through declared functions.

## Source and evidence fields

The runtime does not attach hidden provenance metadata to every value and does not automatically propagate source traceability through function calls.
When a capability needs source locations, evidence, occurrence identity, or derivation data, that information is represented explicitly in its declared input and output types.

For example, an occurrence record may contain both its semantic value and a `SourceLocation` field.
A function accepting only `string` receives only that string; `common.text.trim(string) -> string` does not implicitly retain any location owned by a surrounding occurrence record.
The caller must carry or reconstruct the containing declared record when later processing needs those fields.

## Domain outcomes and execution failures

Expected domain outcomes are represented as typed values.
Examples include:

- malformed reference input;
- unresolved current reference;
- conflicted current reference;
- optional value absence;
- failed validation evaluation.

Execution and configuration failures are not ordinary union variants that every YAML function must inspect.
Examples include:

- a referenced function is missing;
- a function receives an incompatible runtime type despite successful bootstrap;
- a native implementation fails unexpectedly;
- configured current state cannot be constructed reliably.

These failures follow the runtime or operation failure path.

## Excluded from the initial model

The initial runtime does not include:

- recovery, catch, or value-level inspection of `never` termination;
- mutable dictionary operations or implicit mutation of `dict<K, V>` values;
- dynamic object or `map<string, any>`;
- design-facing or generated opaque-carrier `any`;
- `any` or `interface{}` as an opaque native representation binding;
- class inheritance;
- runtime reflection;
- arbitrary YAML type assertions outside exhaustive union or optional switching;
- implicit conversions other than direct assignment of a declared union variant to that union type;
- optional literals or constructors such as `none`, `some(value)`, or `null`;
- a general-purpose untyped `null`;
- design-facing `unit` syntax; no-result functions use `void`.
