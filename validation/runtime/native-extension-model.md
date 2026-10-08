# Native extension model

## Status

Working contract confirmed on 2026-07-18.
This document records the accepted compiled native-extension direction and generated Go application API, including the generated complete function catalog.
Exact build commands, cache layout, activation protocol, trust policy, and automated binding aggregation mechanism remain pending.

## Objective

Native Go implementations and design-facing YAML declarations are separate source sets.
A native implementation can use ordinary Go modules and external Go dependencies without making the YAML interpreter import Go packages dynamically.

The runtime provides an automated build-and-activation workflow so adding a native binding does not require a user to manually edit or rebuild the interpreter host.

## Source separation

The source layout separates:

```text
runtime SDK and interpreter host
native Go extension modules
YAML type and function declarations
application-owned bound YAML manifests
```

Bound manifests use the generic build-time discovery and generated `puppygen/manifest` package defined by `manifest-binding-model.md`.
They are compiled application data rather than native Go binding registrations or PuppyDSL declaration files.

Each native extension module has its own valid Go module boundary and may contain:

```text
go.mod
go.sum
Go packages
binding registration
unit tests
```

Each extension must be independently testable and buildable with normal Go tooling.

For the DRMCP application, application-owned native code will live beneath `drmcp/puppydsl/go_native/`. The generic PuppyDSL runtime SDK and interpreter host remain outside that application-owned module.

## Generated Go application API

The build tool generates one typed application API beneath:

```text
<puppydsl-root>/go_native/puppygen
```

The generated import path is derived from the Go module that contains the application-owned `go_native` directory.

Global PuppyDSL types are exported from:

```text
<go-native-module>/puppygen/types
```

A PuppyDSL function ID maps to a generated Go package and symbol as follows:

```text
operation.discovery_and_listing.discover_record_scopes
  -> <go-native-module>/puppygen/operation/discovery_and_listing
  -> DiscoverRecordScopes
```

The package path is derived from the function namespace. The exported symbol is the PascalCase form of the terminal function segment. The reserved unqualified entry function `main` is exported as `Main` from the root `<go-native-module>/puppygen` package. Generated file names and internal source layout may vary without changing this import contract.

The root `<go-native-module>/puppygen` package also exports:

```go
var Functions map[string]any
```

`Functions` contains every non-generic declared FunctionID exactly once and maps it to that function's generated typed caller facade. The generated source literal is emitted in deterministic FunctionID order; ordinary Go map iteration order remains unspecified. Native, steps-defined, stub, and `main` declarations participate when non-generic because the catalog indexes declared callable contracts rather than executable availability or transport publication. Generic functions are excluded from this reflection catalog because Go requires concrete type arguments before a generic function can become a function value. A consumer may inspect the concrete facade function type with ordinary Go reflection, but the catalog does not expose native binding IDs, choose public functions, or alter runtime invocation semantics.

Every declared function receives a typed caller facade. This is true for `native`, `steps`, and `stub` declarations because the facade represents the declared contract. Stub reachability remains prohibited during design-closure and activation validation.

Every transparent record generates one named Go struct beneath `<go-native-module>/puppygen/types`.
A zero-field transparent record generates a distinct named zero-field struct rather than a shared unit carrier. For example:

```go
type CandidateNotFormed struct{}
```

Its value is constructed with the ordinary named zero value:

```go
CandidateNotFormed{}
```

Generated caller and native-binder signatures use that exact named type. Puppygen must not substitute unnamed `struct{}`, a shared generated unit type, or the internal representation of `void`, because any of those substitutions would erase the PuppyDSL TypeID boundary.
A zero-field transparent record requires no application-owned representation binding.

Every opaque TypeID also participates in staged generation.
Puppygen first emits or exposes a representation-binding API that does not depend on the final generated nominal type.
Application-owned native source binds that TypeID to one compile-time Go representation type.
After resolving all opaque bindings, puppygen emits the final distinct nominal wrappers beneath `<go-native-module>/puppygen/types` and uses those wrappers in generated caller and native-binder signatures.
The exact bootstrap package path and binding-source collection mechanism remain pending, but final generated opaque types must never fall back to `any`, `interface{}`, or raw representation aliases.

Application-owned bound manifests are generated separately beneath:

```text
<go-native-module>/puppygen/manifest
```

That package exposes parsed `yaml.Node` documents and a distinct manifest source-set fingerprint. Manifest-specific decoding and registration remain outside the native extension API.

## Binding relationship

A design-facing function declaration references one native binding ID:

```yaml
functions:
  common.text.trim:
    detail: Removes leading and trailing whitespace before semantic interpretation.
    signature: >-
      (value: string) -> string
    implementation:
      native: text.trim
```

The native extension registers the binding ID against a compiled implementation.
The function ID and native binding ID are separate identities and need not match.

For every function using `implementation.native`, the generated function package also exports a typed binder named `Bind<FunctionName>`. For example, `drmcp_runtime.mcp.serve() -> void` conceptually generates:

```go
type ServeImplementation func(native.Call) error
func BindServe(implementation ServeImplementation) native.Binding
```

The generated binder fixes the FunctionID, native binding ID, and expected concrete Go signature from the PuppyDSL declaration. A Go implementation with an incompatible parameter or result shape fails Go compilation when passed to the binder.

Each binder adapts its typed Go implementation to one uniform runtime invoker:

```go
type Invoker func(call Call, arguments []any) (any, error)
```

The generated adapter validates argument count, converts each argument to its declared generated Go type, invokes the typed implementation, and returns the typed result through the uniform carrier. Runtime registry construction rejects an empty FunctionID, empty signature, empty binding ID, surrounding identifier or signature whitespace, nil invoker, duplicate FunctionID, or duplicate binding ID.

The initial registry indexes bindings by both FunctionID and binding ID. Registry invocation verifies that the requested generated signature exactly matches the registered descriptor before calling the uniform invoker. The registry itself satisfies `Call` for native-only call chains. For mixed steps/native execution, registry dispatch preserves the outer runtime `Call` boundary when entering the native invoker, so a native implementation that invokes another generated facade returns through the same steps-aware execution state rather than bypassing it.

The native implementation and its binding declaration may be colocated:

```go
func Serve(call native.Call) error {
    return nil
}

var serveBinding = puppy_mcp.BindServe(Serve)
```

The binding declaration does not bind the function to MCP or another transport. It binds the Go implementation to the `implementation.native` entry declared by the PuppyDSL function.

Runtime bootstrap verifies that:

- the referenced native binding exists;
- its runtime input and output contract satisfies the declared signature;
- every directly required runtime dependency is available;
- duplicate native binding IDs do not exist.

## External dependencies

A native extension may import external Go packages.
Those dependencies are owned by the extension module's `go.mod` and `go.sum`.
YAML does not declare, download, or import Go packages.

Changing an extension's Go dependencies changes the compiled extension build rather than the design-facing function identity.

## Automated build and activation

The runtime or an associated build coordinator performs the conceptual workflow:

```text
PuppyDSL declaration, bound manifest, native source, or module dependency change
  -> load and statically validate PuppyDSL declarations
  -> discover opaque TypeIDs and emit or expose representation-binding APIs
  -> compile-check and collect application-owned opaque TypeID-to-Go-type bindings
  -> reject missing, duplicate, unknown, non-opaque, or dynamic opaque bindings
  -> discover, bind-header validate, and parse .manifest.yaml documents
  -> generate nominal opaque wrappers, other global Go types, typed caller facades, native binders, and package manifest
  -> collect or generate native function-binding aggregation
  -> run Go validation, tests, and build
  -> produce a versioned native runtime artifact
  -> validate generated declaration, opaque-binding, and manifest fingerprints
  -> validate YAML function declarations against registered bindings
  -> activate the new artifact for subsequent execution
```

The exact process boundary remains pending.
The initial direction may rebuild and restart one native host process rather than dynamically load a Go shared object into an already running process.

## Plugin characterization

The accepted model is a compiled native extension system with automated build and activation.
It provides plugin-like extensibility from the user's perspective, but it does not require Go's standard in-process `plugin` package.

The baseline must remain usable on Windows.

## Trust boundary

Native extension source is trusted executable code.
Building and activating an extension grants it the authority of the native runtime process unless a later process-isolation contract narrows that authority.

Untrusted downloaded Go source must not be compiled and executed automatically without a separate trust, review, and sandboxing design.

## Build-time and bootstrap guarantees

The generated API and native extension build must reject before process activation:

- a missing generated package or exported facade for a declared function;
- an opaque TypeID with no representation binding, more than one binding, a binding targeting an unknown or non-opaque type, or an `any` / `interface{}` representation;
- a final generated opaque carrier implemented as a raw representation alias or an `any`-backed wrapper;
- a generated zero-field transparent record that is not emitted as its own named `struct{}` type, or a caller or binder that substitutes unnamed `struct{}`, a shared unit type, or the internal `void` carrier for that declared type;
- a generated caller or native binder whose opaque parameter or result does not use the exact TypeID-specific nominal wrapper;
- a stale generated API whose declaration or opaque-binding fingerprint differs from the selected PuppyDSL source and binding set;
- a stale generated manifest package whose manifest fingerprint differs from the selected bound-manifest source set;
- a native Go implementation passed to `Bind<FunctionName>` with an incompatible concrete signature;
- a native binding declaration for a function that is not `implementation.native`;
- more than one registered implementation for one native binding ID;
- an `implementation.native` declaration with no activated binding.

Bootstrap repeats declaration, descriptor, binding-ID, and signature-fingerprint checks so a stale binary cannot be paired with newer PuppyDSL declarations. Normal operation execution must not discover a missing or incompatible binding for the first time.

## Pending decisions

- canonical repository layout for the generic runtime SDK and interpreter host;
- one combined native host versus isolated extension workers;
- exact bootstrap package path and source format for opaque representation bindings;
- exact build-time collection and generated aggregation mechanism for opaque representation bindings and implementation-adjacent native function bindings;
- build cache and invalidation keys;
- exact test and validation commands before activation;
- versioned artifact location and rollback;
- process restart or handoff protocol;
- extension enablement manifest;
- trust and allow-list policy.
