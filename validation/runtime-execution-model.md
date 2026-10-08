# Validation runtime execution model

## Status

- Working design note created on 2026-07-15.
- This document is not a canonical Design Record.
- It records the runtime and artifact-module direction confirmed while building the atomic component inventory.
- Exact YAML syntax, Go types, and final function signatures remain provisional until confirmed one topic at a time.
- The structured working baseline now lives under `runtime/`; this document remains a discussion-derived overview and source note.

## Design objective

The YAML surface represents concrete design semantics and processing steps.
It should call named functions such as `common.text.is-empty` rather than expose lower-level implementation expressions such as `value == ""`.

A named function may have either:

- a native Go implementation; or
- a graph implementation composed from other declared functions.

Both forms expose the same typed function contract.
Implementation form does not determine namespace ownership.

## Source, candidate, and validation boundary

The runtime preserves this processing boundary:

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
Invalid but structurally available projection values remain retrievable; structurally unavailable projections produce an unavailable or null result.

## Runtime function registry

The runtime has one generic function registry:

```text
fully-qualified function ID -> declared callable function
```

The registry contains common, artifact/base, artifact-specific, and operation functions.
A function declaration supplies at least:

- one stable fully-qualified function ID;
- typed inputs;
- typed outputs;
- one implementation form;
- implementation data for either native or graph execution.

Function IDs stored in manifests are resolved during runtime bootstrap.
Normal execution uses resolved function handles rather than repeated dynamic string lookup.

## Reference processing model

Reference processing has two interpretation stages.

### Base reference parsing

Shared base parsing splits a raw reference into one structure-specific base value without establishing complete artifact grammar conformance:

```text
BaseReference
  = TreeBaseReference
  | SequentialBaseReference
```

A base reference retains:

- the exact raw string;
- the structure form;
- a routing discriminator;
- raw segments.

Any source location or evidence associated with the raw reference belongs to an explicitly declared surrounding record; it is not hidden metadata propagated by the runtime.

The discriminator identifies an artifact-module candidate for routing.
It is not yet the confirmed artifact kind.

### Artifact-specific reference parsing

The selected artifact module applies its artifact grammar to the base reference and returns one confirmed artifact reference:

```text
ArtifactReference
  kind: ArtifactKind
  value:
    TreeArtifactReference
    | SequentialArtifactReference
```

Structure-specific confirmed values avoid one reference type per artifact:

```text
TreeArtifactReference
  app_namespace
  path_segments

SequentialArtifactReference
  app_namespace
  domain_namespace
  sequence_segments
```

Artifact differences are represented by `ArtifactReference.kind` plus artifact declarations and rules rather than separate Task, Requirement, Decision, Investigation, or Work Item reference classes.

The tagged-union representation is a runtime type model.
Design YAML does not inspect union variants or branch on runtime types.

## Artifact module manifest

An artifact module manifest is a registration declaration, not a new module-level interface construct.
It identifies one artifact provider and assigns declared functions to standardized artifact capability slots.

Conceptual content:

```yaml
module: artifact.task

artifact:
  kind: task
  structure: sequential

reference:
  discriminator:
    kind_segment: TASK

provides:
  reference.parse: artifact.task.reference.parse
  reference.canonical-form: artifact.task.reference.canonical-form
  facts.extract: artifact.task.facts.extract
  identity.determine: artifact.task.identity.determine
  validation.run: artifact.task.validation.run
```

Each `provides` target is an ordinary declared function.
The runtime checks the function against the capability contract for its slot.
The manifest does not duplicate the function's input and output declaration.

## Artifact module registration and catalog

Runtime bootstrap converts a valid manifest into an `ArtifactModuleRegistration` containing:

- module ID;
- artifact kind;
- record structure;
- reference discriminator;
- resolved function handles for required capability slots.

The registration is stored in an artifact-specific catalog with indexes such as:

```text
module ID -> ArtifactModuleRegistration
artifact kind -> ArtifactModuleRegistration
sequential discriminator -> ArtifactModuleRegistration
tree discriminator -> ArtifactModuleRegistration
```

The function registry is generic runtime infrastructure.
The artifact module builder and catalog are artifact-specific runtime facilities built on top of that registry.

## Generic module routing

A raw reference follows this conceptual flow:

```yaml
base_ref: artifact.base.reference.parse(raw_ref)
module: artifact.base.module.from-base-reference(base_ref)
reference: module.reference.parse(base_ref)
canonical_ref: module.reference.canonical-form(reference)
resolution: artifact.base.reference.resolve-current(reference)
```

Routing uses the base reference discriminator to obtain one artifact module.
The module-bound `reference.parse` function confirms artifact grammar and returns an `ArtifactReference` whose kind agrees with the module registration.

After an `ArtifactReference` exists, generic processing can retrieve the same module by confirmed artifact kind without writing an artifact-kind switch in YAML.

YAML does not contain runtime type checks such as:

```yaml
if reference is SequentialArtifactReference:
```

Runtime type dispatch remains internal.
YAML `if` is reserved for meaningful design-rule branching.

## Bootstrap validation

Runtime bootstrap must reject the configuration before operation execution when any of these conditions exists:

- duplicate module ID;
- duplicate artifact kind;
- conflicting tree or sequential discriminator;
- missing required capability slot;
- referenced function ID does not exist;
- function input or output types do not satisfy its slot contract;
- module structure disagrees with a structure-specific slot signature;
- artifact reference parser returns a kind other than the module's registered kind;
- artifact reference parser returns a tree/sequential value inconsistent with the module structure.

These are runtime configuration failures, not validation findings about a record.

## Emerging runtime requirements

The current design establishes probable requirements for:

- named user-defined record types;
- tagged-union runtime values hidden from YAML type branching;
- typed function declarations;
- native and graph function implementations;
- opaque local values;
- assignment of function results to local names;
- module-bound function invocation;
- startup function-signature validation;
- a generic function registry;
- an artifact module manifest, builder, registration object, and catalog;
- bounded iteration for repeated artifact facts;
- conditional execution only for design-rule branching.

The runtime should not require design YAML to perform:

- arbitrary reflection;
- dynamic type assertions;
- union variant switches;
- repeated dynamic string-based function lookup;
- low-level comparisons that should be hidden behind named semantic functions.

## Immediate design sequence

The next runtime slice is intentionally handled before the remaining artifact capability inventory because it now constrains every function input and output.

One topic is confirmed at a time in this order:

1. runtime value and type model;
2. function declaration contract;
3. artifact module manifest and capability slots;
4. module bootstrap, catalog, and routing;
5. design-facing invocation syntax;
6. minimum `if` and bounded `for` semantics.

After this slice, continue the atomic inventory with artifact/base facts, generic validation checks, and artifact-specific facts or custom components.
