# Validation runtime working specification set

## Status

- Working specification set created on 2026-07-15.
- The PuppyDSL core language contract has moved to `spec:puppydsl.language` under `puppydsl/records/spec/language/`.
- This directory is no longer the current authority for source, type, expression, function, control-flow, context, event, or program-validity semantics.
- `bootstrap-validation-inventory.md` remains implementation-planning input until canonical validator specifications replace it.
- Generator, native-integration, manifest-binding, and application-specific notes remain working inputs for separately owned specification areas.
- Do not amend this directory to change the canonical PuppyDSL language contract.

## Purpose

This set preserves the working material that produced the canonical PuppyDSL language specification and the remaining validator and integration design inputs.

The runtime must let YAML describe meaningful processing steps such as:

```yaml
empty: common.text.is_empty(value)
parse_outcome: module.reference.parse(base_ref)
```

It must not require YAML to express low-level implementation conditions such as:

```yaml
empty: value == ""
if base_ref is SequentialBaseReference:
```

## Working architecture

```text
source
  -> shared source-structure parse
  -> artifact facts and identity determination
  -> candidate formation and repository-wide admission
  -> listing, retrieval, navigation, or search
  -> artifact conformance validation
```

Full artifact validation does not gate listing or retrieval.
Only identity availability and expected-context agreement gate candidate formation.

Reference processing uses artifact-module routing:

```text
exact raw reference
  -> artifact.base.reference.parse(raw)
     -> SharedRoutingOutcome
        -> UnrouteableReference
        | RoutedBaseReference.reference -> BaseReference union
  -> ArtifactModuleCatalog lookup
  -> module-bound narrowing and artifact-specific canonical parse
     -> ArtifactReferenceParseOutcome
        -> InvalidArtifactReference
        | opaque ArtifactReference
  -> canonical formatting or current-state resolution
```

## PuppyDSL source files

A design-facing PuppyDSL source file uses the exact lowercase compound extension `.puppy.yaml`.
The runtime bootstrap source discovery treats only files whose names end with that exact suffix as PuppyDSL source files.
The shorter `.yaml` extension, the `.yml` extension, and the compound `.puppy.yml` extension do not identify PuppyDSL source.

Files with ordinary YAML extensions remain ordinary application data even when they are located beneath the selected PuppyDSL application root.
Their shape is owned by the native function or application component that reads them; the PuppyDSL declaration loader must not parse them as type, function, context, or event declarations merely because of their location or root keys.

Application manifests that must be compiled into the generated Go application use the distinct exact suffix `.manifest.yaml` and the physical-line-one bind header defined by `manifest-binding-model.md`.
Puppygen parses those files during generation and emits their mapping-root `yaml.Node` values beneath `<go-native-module>/puppygen/manifest`.
Bound manifests remain separate from `.puppy.yaml` declarations and are not added to the PuppyDSL declaration registry.

The designated entry declaration follows the same source-file rule and is named `main.puppy.yaml`.
It remains an ordinary declaration file using the `functions` root rather than introducing a separate entry-file grammar or root key.
A runtime invocation selects one PuppyDSL application root, discovers its `.puppy.yaml` declaration files, and resolves the entry function declared by `main.puppy.yaml` after declaration registration and bootstrap validation.

## Identifier grammar

PuppyDSL declaration identifiers use a Go-mappable spelling so generated Go APIs do not require lossy normalization or collision-prone alias rules.

The initial identifier forms are:

```text
type ID
  [A-Z][A-Za-z0-9]*

lower-snake identifier
  [a-z][a-z0-9]*(?:_[a-z0-9]+)*

implementation-oriented function segment
  _[a-z][a-z0-9]*(?:_[a-z0-9]+)*
```

Type IDs use the type form. Function namespace segments, ordinary function terminal segments, context IDs, event IDs, module IDs, provided-value names, signature input names, local bindings, transparent-record fields, and enum members use the lower-snake form. An implementation-oriented function may use the leading-underscore form for its terminal segment only.

A qualified function ID has the form:

```text
<namespace-segment>(.<namespace-segment>)*.<function-segment>
```

The designated entry function is the reserved unqualified ID `main`.

The hyphen character is prohibited in these declaration identifiers. This prohibition does not alter canonical Design Record IDs, artifact references, filesystem paths carried as application data, quoted string literals, or ordinary prose.

Application-owned PuppyDSL module directory segments and `.puppy.yaml` file stems use lower-snake spelling. Declaration-kind directories such as `types`, `functions`, and `contexts` organize source files but do not add namespace segments. Source ownership directories do not define FunctionID namespaces; for example, `lib/operations/` may own functions in the `operation.*` namespace.

## Generated Go API

A PuppyDSL application receives a generated typed API beneath its application-owned native Go root:

```text
<puppydsl-root>/go_native/puppygen
```

The Go import path is derived from the Go module that contains `<puppydsl-root>/go_native`.

Global PuppyDSL types are imported from:

```text
<go-native-module>/puppygen/types
```

Each opaque TypeID is emitted there as its own nominal wrapper around one application-bound compile-time Go representation.
Opaque types are never emitted as raw representation aliases or `any`-backed placeholders.
Final type and function generation requires every opaque TypeID to have exactly one valid native representation binding; the preliminary binding API and exact staged-generation command remain build-tool details.

For a qualified function ID `<namespace>.<function_name>`, the generated function package is imported from:

```text
<go-native-module>/puppygen/<namespace as path segments>
```

The reserved entry function `main` is exported as `Main` from the root `<go-native-module>/puppygen` package.

The generated Go function name is the PascalCase form of `<function_name>`. Every declared function receives a typed caller facade regardless of whether its implementation is `native`, `steps`, or `stub`. A function with `implementation.native` additionally receives `Bind<FunctionName>` in the same generated package.

The root `<go-native-module>/puppygen` package exports `Functions map[string]any`. It contains every non-generic declared FunctionID exactly once and maps each ID to its generated typed caller facade. The generated source literal is emitted in deterministic FunctionID order, while Go map iteration order remains unspecified. Generic functions are excluded because Go requires concrete type arguments before they can become function values. This is generic generated Go metadata for consumers that need to locate and inspect concrete facades; it does not select public functions or introduce transport semantics.

Bound application manifests are imported from:

```text
<go-native-module>/puppygen/manifest
```

That package exposes one generated `manifest.Document` variable per bind name, exact bind-name lookup through `ByBindName`, and a manifest source-set fingerprint as defined by `manifest-binding-model.md`.

Generated file names, file counts, private helpers, descriptor splitting, and manifest-node literal rendering are implementation details. The package import paths and exported symbols above are the stable application-facing contract.

## Declaration files

Related declarations may be grouped in one `.puppy.yaml` file. Each declaration file uses exactly one declaration-kind root mapping from this vocabulary:

```text
types
functions
contexts
events
```

Every declaration file begins with this required source header before its declaration root:

```yaml
# responsibility: Defines artifact-wide identity and namespace types shared across artifact kinds and operations.
# excludes: Runtime configuration paths, module registration types, and operation-specific request or result types.

types:
  ArtifactKind:
    detail: Identifies one artifact kind registered by an artifact module.
    scalar: string
```

The first non-empty source line must be one `# responsibility:` comment with a non-empty single-line responsibility statement. The next non-empty source line must be one `# excludes:` comment with a non-empty single-line exclusion statement. Exactly one empty line separates the required header from the declaration root. Additional source comments may appear only after this required header.

The header defines why the declarations belong in the same file and which adjacent concerns must remain elsewhere. It is source-level design authority rather than decoded YAML data. The declaration loader validates the raw source envelope before ordinary YAML declaration decoding and retains enough source information to report missing, malformed, or misplaced header comments.

The declaration identity is the mapping key beneath the root. YAML multi-document separators are not used.

Every type, function, context, and event declaration contains a non-empty `detail` field. `detail` is bootstrap-validated design authority describing one declaration's semantic responsibility; it is not executable implementation logic and must not merely repeat the declaration shape. The file header owns the declaration-set boundary, while `detail` owns the individual declaration boundary.

A file groups declarations of one kind only. Declaration placement follows semantic ownership rather than the set of consumers:

```text
common
  artifact-independent reusable values and functions

runtime
  generic interpreter, registry, bootstrap, and execution-infrastructure declarations

<app>_runtime
  application-owned launch, configuration, transport, and runtime-composition declarations; for example `drmcp_runtime`

artifact/base
  Design Record semantics shared across artifact kinds

artifact/<kind>
  semantics owned by one artifact kind
```

Artifact-owned declarations are located beneath the directory of the owning artifact module rather than being collected into one unrelated global file. Subdirectories narrower than these ownership boundaries are introduced only when a concrete declaration set makes the additional boundary useful.

## Documents

| document | responsibility |
|---|---|
| `value-model.md` | Runtime values, named types, immutable dictionaries, transparent records, opaque values, unions, explicit source/evidence fields, and failure boundaries. |
| `function-model.md` | Typed function declarations, native and steps-defined implementations, runtime dependencies, and execution behavior. |
| `context-model.md` | Function-owned context scopes, lazy providers, scope-local caches, cleanup, and static context availability validation. |
| `event-model.md` | Declared runtime events, context invalidation, and bounded complete-operation restart. |
| `native-extension-model.md` | Separate Go extension modules, external dependencies, staged opaque representation binding, native function binding registration, and automated build and activation. |
| `native-value-binding-model.md` | YAML-to-Go value conformance, opaque nominal wrappers, native representation bindings, and remaining value-carrier decisions. |
| `manifest-binding-model.md` | Build-time discovery, bind headers, YAML parsing, generated manifest-package access, fingerprints, and the generic consumer boundary. |
| `artifact-module-model.md` | Artifact module manifests, capability slots, bootstrap validation, registrations, catalog indexes, and routing. |
| `reference-parsing-boundary.md` | Product- and DRMCP-derived boundary among shared routing, artifact canonical parsing, malformed input, and current-state resolution. |
| `execution-model.md` | Source parsing, artifact facts, identity gating, admission, operation consumption, and validation execution. |
| `invocation-model.md` | Design-facing local binding, binding-origin field and dictionary access chains, dictionary and record construction, function-call model, module-bound invocation, and restrictions on dynamic access. |
| `control-flow.md` | Boolean `if`, exhaustive `switch` over declared unions, enums, and optional values, bounded list and dictionary `for`, and their shared lexical-scope and output rules. |
| `open-questions.md` | Remaining decisions that must be resolved from concrete consumers rather than speculative language design. |

## Consistency rule

A future change to one of these concerns must update every affected document in this set in the same design slice:

- PuppyDSL source-file identification or entry-file naming;
- bound-manifest identification, bind-header grammar, or generated manifest-package mapping;
- identifier grammar, module-path spelling, or generated Go package mapping;
- runtime value or type shape;
- opaque native representation binding, nominal-wrapper generation, or staged-generation behavior;
- function signature rules;
- artifact module capability slots;
- reference routing;
- invocation semantics;
- control-flow semantics;
- source, admission, retrieval, or validation boundaries.

A change must not be applied only to an example YAML snippet while leaving the type, function, or module contract inconsistent.

## Relationship to current authority

`spec:puppydsl.language` owns current PuppyDSL language semantics.

`../runtime-execution-model.md` records the discussion path that produced this working set.

`../atomic-component-inventory.md` defines currently accepted callable capabilities. Those capabilities must conform to the canonical PuppyDSL language specifications and to their separately owned application contracts.
