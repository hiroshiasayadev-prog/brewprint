# Native value binding model

## Status

Working design for binding design-facing YAML types to Go native representations.

D-001 and D-006 are currently in discussion. The cross-cutting opaque nominal-binding boundary recorded below is confirmed; remaining decision-inventory topics must not be treated as confirmed contract until their own status becomes `decided`.

## Scope

This document owns:

- native type-binding identity and registration;
- YAML type ID to Go representation conformance;
- native function input and output boundary validation;
- transparent, opaque, union, and generic native value representation;
- bootstrap-time and execution-time failure boundaries;
- the initial external Go module registration boundary;
- the connection to later automatic compilation.

This document does not own:

- current-reference resolution;
- sequential artifact component declaration;
- bootstrap fixture inventory;
- artifact capability-slot inventory;
- operation projection contracts;
- a general plugin ABI or cross-language FFI.

## Fixed inputs

The design starts from these accepted constraints:

- YAML type declarations are the design-facing authority.
- Go structs do not define or extend the YAML-visible schema.
- Native and steps-defined functions share the same declared signature.
- Design-facing `dict<K, V>` is an immutable declared composite; dynamic `any`, reflection, implicit conversion, and untyped or schema-inferred maps are prohibited.
- Opaque native representation remains hidden from YAML.
- Every opaque TypeID requires one application-owned compile-time Go representation binding before final generation and activation.
- Final generated Go APIs preserve each opaque TypeID as a distinct nominal wrapper even when several opaque types use the same native representation.
- `any`, `interface{}`, and a generated `struct { value any }` fallback are prohibited for opaque representation binding and final opaque carriers.
- A declared union has one interpreter-owned active variant type and no user-defined tag or payload field.
- Native binding IDs are registry identities and are not Go package paths or source symbols.
- Generated Go caller facades and `Bind<FunctionName>` binders are required application APIs derived from PuppyDSL declarations.
- Global generated Go types are imported from `<go-native-module>/puppygen/types`, physically generated beneath `<puppydsl-root>/go_native/puppygen/types`; opaque nominal wrappers follow the confirmed boundary below, while other exact value-carrier mappings remain owned by this document's open decisions.
- Bootstrap configuration failures and operation execution failures remain separate.
- `ArtifactReference` is one design-facing opaque type whose native representation is either `TreeArtifactReference` or `SequentialArtifactReference`.
- `SequentialArtifactReference.fields: map<string, string>` is native-internal data and does not become a design-facing `dict<string, string>` merely because dictionaries are supported elsewhere.

## Decision inventory

| ID | Topic | Status | Depends on |
|---|---|---|---|
| D-001 | Native binding foundation | in_discussion | — |
| D-002 | Type-binding registry identity and duplicate handling | open | D-001 |
| D-003 | Primitive, named scalar, and enum representations | open | D-001, D-002 |
| D-004 | Transparent-record field mapping and undeclared Go fields | open | D-001, D-002 |
| D-005 | Pointer, value, nil, and zero-value rules | open | D-003, D-004 |
| D-006 | Opaque type conformance and `ArtifactReference` representation | in_discussion | D-001, D-002, D-005 |
| D-007 | Runtime union wrapper and active-variant storage | open | D-001, D-002 |
| D-008 | `list<T>`, `optional<T>`, `dict<K, V>`, and generic instantiation descriptors | open | D-001, D-003, D-004, D-005 |
| D-009 | Native function input/output validation | open | D-003 through D-008 |
| D-010 | Bootstrap-only versus execution-time validation | open | D-009 |
| D-011 | External Go module registration and automatic compilation seam | open | D-001, D-002, D-009, D-010 |
| D-012 | Binding mismatch failure classification | open | D-009 through D-011 |

## Confirmed opaque nominal-binding boundary

The following boundary is decided even though D-006 remains open for representation-specific validity, nil, zero-value, and `ArtifactReference` details.

Every `opaque: true` declaration requires exactly one application-owned native representation binding.
The binding associates the opaque TypeID with one compile-time Go representation type and must be available before puppygen emits the final application API.
An unbound opaque TypeID, duplicate binding, binding for an unknown or non-opaque TypeID, or binding to `any` or `interface{}` is a generation failure.

Puppygen must not emit an opaque type as an alias for its bound representation and must not use an `any`-backed placeholder carrier.
It emits one distinct nominal wrapper per opaque TypeID whose hidden field has the exact bound Go representation type.
For example, these bindings are permitted:

```text
ArtifactReference -> *yaml.Node
ManifestDocument  -> *yaml.Node
```

but the final generated types remain distinct:

```go
type ArtifactReference struct {
    value *yaml.Node
}

type ManifestDocument struct {
    value *yaml.Node
}
```

The wrappers are not mutually assignable.
The shared native representation does not weaken PuppyDSL nominal identity.

Generated caller facades and native implementation binders use the generated nominal wrappers in function signatures rather than substituting the raw native representation.
A handwritten implementation that accepts `ManifestDocument` where `ArtifactReference` is declared therefore fails Go compilation even when both wrappers contain `*yaml.Node`.

Application-owned native code receives generated typed construction and representation-access operations, or an equivalent compile-time API, for each wrapper.
Those operations expose only that opaque TypeID's bound representation and require no reflection, `any`, or dynamic cast.
Exact operation names remain a generated API detail until the binding source format is finalized.

The generation contract is conceptually staged:

```text
PuppyDSL declarations
  -> emit or expose opaque representation-binding APIs
  -> compile-check application-owned TypeID-to-Go-type bindings
  -> resolve every opaque binding
  -> emit final nominal wrappers, caller facades, and native binders
  -> compile application-owned native implementations
```

The bootstrap binding API must not depend on the final generated nominal wrappers; otherwise the binding source and final generated API would form an import cycle.
The exact command sequence, generated bootstrap package path, binding-source discovery, and aggregation mechanism remain pending implementation details.

## D-001 — Native binding foundation

### Concrete consumer

`artifact.base.reference.parse(raw: string) -> SharedRoutingOutcome` is a native function.
Its Go implementation must return a value conforming to a YAML-declared union whose variants contain transparent records and named scalar fields.
Later native functions must also accept or return opaque `ArtifactReference`, generic records such as `SetComparison<T>`, and composites such as `list<T>`, `optional<T>`, and `dict<K, V>`.

The runtime therefore needs one binding foundation that can:

- preserve YAML type identity independently of Go symbol names;
- validate Go representations before operation execution where possible;
- support hidden opaque representations;
- represent interpreter-owned union variants;
- handle concrete generic instantiations;
- work with separately compiled Go modules.

### Options

#### A. Go reflection-based registration

A registration supplies Go values or functions, and the runtime derives their shape from `reflect.Type`.
Struct fields, parameter types, result types, pointers, and interfaces are inspected directly.

Benefits:

- low initial adapter authoring;
- direct inspection of ordinary nongeneric Go function signatures;
- straightforward external-module registration after compilation.

Costs:

- risks making Go structure the effective schema authority;
- YAML field names still require tags or naming conventions;
- Go reflection cannot instantiate an uninstantiated generic type or generic function from runtime type arguments;
- opaque and union semantics require extra runtime metadata that reflection cannot infer;
- reflection-derived errors tend to expose implementation-oriented Go names;
- normal invocation through `reflect.Call` adds runtime checks and weaker compile-time guarantees.

#### B. Explicit native type descriptors

Each native binding explicitly registers its YAML type identity, Go representation identity, field mapping or opacity contract, and construction or validation operations.
The runtime compares these descriptors with parsed YAML declarations during bootstrap.

Benefits:

- keeps YAML declarations authoritative;
- supports explicit hidden opaque contracts and runtime-owned union wrappers;
- permits field mapping that is independent of Go field names and tags;
- produces design-facing mismatch diagnostics;
- can represent generic type constructors and concrete instantiation descriptors without pretending that Go reflection can instantiate generics;
- works across separately compiled modules through ordinary registration calls.

Costs:

- descriptor authoring is more verbose;
- a manual descriptor can disagree with implementation code unless bootstrap reflection or generated checks verify the claimed Go types;
- manually authored invocation adapters may still contain incorrect casts that only value-boundary validation detects.

#### C. Generated adapters

A build step reads YAML declarations and Go binding declarations, then generates typed registration and invocation adapters.
Generated code performs conversion, field access, generic specialization, and output wrapping.

Benefits:

- strongest compile-time checking of concrete Go function signatures;
- no repeated reflective invocation overhead;
- good support for concrete generic instantiations known at build time;
- generated mismatch locations can point to both YAML declarations and Go registrations.

Costs:

- makes code generation a prerequisite for every native extension build;
- requires a stable Go annotation or registration-source format before the runtime model itself is settled;
- increases automatic-compilation complexity and generated-source lifecycle requirements;
- does not remove the need for runtime descriptors because unions, optional absence, opaque identities, and loaded YAML declarations still require runtime type metadata;
- slows early iteration for little benefit when the initial binding set is small.

#### D. Reflection plus generated-adapter hybrid

Reflection validates discoverable Go declarations during development or bootstrap, while generated adapters perform normal invocation and generic specialization.

Benefits:

- combines readable bootstrap diagnostics with typed execution;
- can reduce handwritten descriptor and adapter drift;
- supports a later optimized extension toolchain.

Costs:

- highest initial implementation complexity;
- contains two conformance mechanisms that must agree;
- still requires an explicit runtime type model for opaque, union, optional, and generic descriptors;
- premature before the descriptor and failure contracts are stable.

### Comparison

| Axis | A. Reflection | B. Explicit descriptors | C. Generated adapters | D. Reflection + generated |
|---|---|---|---|---|
| YAML remains authority | weak to medium | strong | strong if generator consumes YAML | strong if both consume YAML |
| Bootstrap validation range | good for ordinary concrete Go types | strong for declared contracts; implementation claims need checks | strongest for generated concrete bindings | strongest but duplicated |
| Generic support | poor without concrete manual instantiation | explicit runtime instantiation model | strong for build-known instantiations | strong |
| Opaque support | requires added metadata | direct | direct through generated wrappers | direct |
| Union support | cannot infer semantics | direct runtime descriptor | generated wrapping still needs runtime descriptor | direct but complex |
| External Go modules | simple registration | simple registration | requires generator integration | requires generator and reflection integration |
| Authoring load | low initially | medium | low per binding after tooling, high tooling cost | highest tooling cost |
| Runtime overhead | highest if reflective invocation | low to medium depending on adapter | low | low |
| Error readability | Go-oriented unless translated | design-facing | potentially best | potentially best |
| Automatic compilation fit | acceptable | good stable seam | good after toolchain exists | good later |
| Initial implementation fit | tempting but incomplete | best fit | premature | premature |

### Recommendation

Adopt **B. Explicit native type descriptors** as the binding contract.

Use Go reflection only as an implementation-side bootstrap verifier for concrete registered Go types and ordinary function signatures. Reflection does not define type identity, infer YAML-visible fields, or become the execution contract.

Require generated typed caller facades and `Bind<FunctionName>` binders as the application-facing Go API. These generated APIs enforce concrete native implementation signatures and preserve stable FunctionID and binding-ID descriptors.

For opaque declarations, the confirmed nominal-binding boundary above additionally requires generated nominal wrappers and typed representation access. An explicit descriptor may record and validate the binding, but it may not replace the generated wrapper with a raw Go alias or dynamic carrier.

Do not yet require every runtime value conversion, transparent-field accessor, union wrapper, or generic specialization to be generated. Preserve the explicit descriptor contract as the authority for those unresolved value-boundary decisions until their mappings stabilize.

### Implementation concept

Conceptually, bootstrap consumes two independently owned inputs:

```text
parsed YAML type declarations
activated native binding registrations
```

A native registration exposes stable registry data rather than Go source identity:

```go
type NativeTypeBinding struct {
    TypeID         DeclaredTypeID
    Representation NativeRepresentationDescriptor
}

type NativeRepresentationDescriptor interface {
    DeclaredKind() NativeRepresentationKind
    GoType() reflect.Type
}
```

The exact descriptor variants remain later decisions. For example, a transparent-record descriptor will explicitly map YAML field names to native accessors or Go fields. An opaque descriptor records the TypeID and bound Go representation used to generate its nominal wrapper; representation-specific validity and zero-value rules remain part of D-005 and D-006.

Bootstrap compares the registration descriptor with the YAML declaration and produces one immutable resolved type binding. Normal execution uses that resolved binding and does not repeat string lookup or schema inference.

Reflection may verify claims such as:

```text
registered Go type exists
registered field or accessor has the stated Go type
registered native function has the stated concrete parameter/result carrier
registered opaque carrier satisfies the bound interface or concrete type
```

Reflection must not add undeclared fields or infer design type identity from package paths, names, tags, or struct shape.

### Future extension impact

This choice preserves:

- mandatory generated caller facades and native binders without making generated Go shape the YAML schema authority;
- generated value descriptors and conversion adapters as a later implementation choice where useful, except that opaque nominal wrappers and typed representation access are required;
- per-module automatic compilation that activates typed binding declarations;
- richer static checking without changing YAML declarations;
- replacement of reflective invocation with generated typed invocation while retaining the same registry identity and mismatch taxonomy.

It does not commit the runtime to a stable binary plugin ABI.

### Not required now

- dynamic `.so` loading;
- cross-language FFI;
- package-path-based type identity;
- Go annotation syntax for generation;
- mandatory generated source checked into the repository rather than produced by the build workflow;
- generic reflection capable of instantiating arbitrary Go generic declarations;
- zero-allocation adapters.
