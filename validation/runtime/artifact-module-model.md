# Artifact module and catalog model

## Status

Working contract for artifact providers, module manifests, capability-slot binding, bootstrap validation, and runtime routing.

## Declaration detail

Every artifact module manifest includes a non-empty `detail` field that states the artifact provider responsibility and capability boundary.
The field is bootstrap-validated design authority and is not a duplicate of the manifest registrations.

## Core distinction

An artifact module is not a new language-level interface construct.

The model has four separate concepts:

```text
artifact module manifest
  registration declaration

artifact capability contract
  expected function signature for one standardized slot

artifact module registration
  validated runtime object containing resolved function handles

ArtifactModuleCatalog
  runtime lookup indexes for registrations
```

## Generic function registry beneath the module model

All declared functions are first registered in the generic runtime function registry:

```text
FunctionID -> Callable
```

The artifact module builder resolves function IDs from a manifest against this registry.
The artifact module catalog does not replace the generic function registry.

## Artifact module manifest

An artifact module manifest is an application-owned bound manifest delivered through `manifest-binding-model.md` rather than a `.puppy.yaml` declaration root.
Puppygen discovers its exact `.manifest.yaml` source, validates its `#bind:` header, parses its mapping root, and exposes the resulting `manifest.Document` from the generated `puppygen/manifest` package.

The artifact-module subsystem, not the generic manifest binder, owns the document schema.
Module entries are grouped beneath the manifest document's `modules` root mapping, and each module identity is the mapping key.
The application-owned bind name selects the generated manifest document; it does not replace the module IDs declared inside that document.

Conceptual content:

```yaml
#bind: artifact_modules

modules:
  artifact.task:
    detail: Provides Task reference, fact, identity, and validation capabilities.
    artifact:
      kind: task
      structure: ArtifactStructure.sequential
    reference:
      discriminator:
        kind_segment: "TASK"
    provides:
      reference.parse: artifact.task.reference.parse
      reference.canonical_form: artifact.task.reference.canonical_form
      facts.extract: artifact.task.facts.extract
      identity.determine: artifact.task.identity.determine
      validation.run: artifact.task.validation.run
```

The example bind name is application-owned; the artifact-module content contract does not require one globally fixed bind name.
The manifest does not duplicate function input or output declarations.
Each referenced function remains an ordinary function declaration in the generic registry.

## Capability slots

Each `provides` key names a standardized artifact capability slot.
Each slot has a base-owned signature contract.

Initial slot families are:

```text
reference.parse
reference.canonical_form
facts.extract
identity.determine
validation.run
```

This slot set may be refined as artifact/base fact inventory continues.
A slot is retained only when generic runtime flow genuinely needs to invoke it without artifact-specific branching.

The module as a whole does not satisfy one monolithic interface.
Each provided function satisfies the contract of its assigned slot.

## Structure-specific reference parsing

Shared reference parsing returns `SharedRoutingOutcome`. Only `RoutedBaseReference` proceeds to module selection, carrying one `BaseReference` whose active variant type is structure-specific:

```text
TreeBaseReference
SequentialBaseReference
```

`UnrouteableReference` terminates shared routing before catalog lookup.
The active `BaseReference` variant type is runtime type information and is not exposed as a YAML field.
The selected module's `reference.parse` slot applies the artifact-specific grammar and returns `ArtifactReferenceParseOutcome`.

Conceptually:

```text
tree module:
  TreeBaseReference -> ArtifactReferenceParseOutcome

sequential module:
  SequentialBaseReference -> ArtifactReferenceParseOutcome
```

At a module-bound invocation, the runtime verifies that the module registration structure agrees with the active `BaseReference` variant type and narrows the value before invoking the resolved slot function.
The target function may be a steps-defined YAML function whose input is the concrete variant type and whose output is `ArtifactReferenceParseOutcome = ArtifactReference | InvalidArtifactReference`.
A successful sequential parser constructs an opaque `SequentialArtifactReference` from the original `SequentialBaseReference` plus a native `map<string, string>` of artifact-specific named components.
A successful tree parser constructs an opaque `TreeArtifactReference` containing only the original `TreeBaseReference`; no additional tree-specific field map is required.
Both native forms cross the design-facing boundary as the single opaque type `ArtifactReference`.
Only the `ArtifactReference` variant may proceed to canonical formatting, module lookup from confirmed reference, or current-state resolution.
The caller does not need to author an explicit type switch for this module-bound call; module-bound dispatch remains one declared narrowing boundary.
Elsewhere, design-facing YAML may consume a union through the exhaustive `switch` construct defined by the control-flow model.
The module structure, union declaration, and slot signature establish compatibility for the implicit module-bound narrowing.

## Module builder

The artifact module builder is an artifact-specific runtime facility built on the generic function registry.

Conceptually:

```text
generated manifest.Document
  -> decode modules mapping
  -> BuildArtifactModule(manifest, function_registry)
     -> ArtifactModuleRegistration
     | configuration failure
```

The builder:

1. decodes and validates the artifact-module manifest structure;
2. resolves every provided function ID;
3. compares each function signature with its capability-slot contract;
4. constructs typed or slot-specific resolved handles;
5. validates module-level invariants;
6. returns one immutable registration.

Normal execution must not repeatedly resolve manifest strings.

## Artifact module registration

A successful registration contains at least:

```text
module ID
artifact kind
artifact structure
reference discriminator
resolved capability-slot handles
```

The registration is an opaque runtime value to design YAML.
YAML may invoke a declared module-bound capability but may not inspect or mutate the registration.

## ArtifactModuleCatalog

The runtime catalog indexes validated registrations.

Required conceptual indexes include:

```text
module ID -> ArtifactModuleRegistration
artifact kind -> ArtifactModuleRegistration
sequential discriminator -> ArtifactModuleRegistration
tree discriminator -> ArtifactModuleRegistration
```

The catalog is frozen before operation execution begins.
Runtime mutation or late registration is outside the initial design target.

## Routing operations

Generic processing requires conceptual lookup operations equivalent to:

```text
from_base_reference(BaseReference) -> ArtifactModule
from_reference(ArtifactReference) -> ArtifactModule
from_kind(ArtifactKind) -> ArtifactModule
```

These operations may become named artifact/base functions or direct runtime facilities.
Their design-facing semantics must remain independent of concrete artifact kinds.

### From base reference

The base reference discriminator selects the module whose artifact parser should confirm the full grammar.

```text
TASK discriminator -> artifact.task
spec tree prefix -> artifact.spec
```

The discriminator is routing evidence, not a confirmed artifact kind.

### From confirmed reference

After module-bound parsing succeeds, the native `ArtifactReference` retains its structure-specific base reference.
The base discriminator identifies the registered module for later canonical formatting or artifact behavior:

```text
TreeArtifactReference.base.record_kind
SequentialArtifactReference.base.artifact_kind
```

No duplicate top-level `ArtifactReference.kind` field is required.
YAML cannot inspect either native representation directly.

### From expected kind

Source-corpus processing already knows the corpus-selected artifact kind.
It may obtain the module directly without parsing a reference discriminator.

## Bootstrap invariants

Runtime bootstrap must reject configuration when any of these conditions exists:

- duplicate module ID;
- duplicate artifact kind;
- conflicting tree discriminator;
- conflicting sequential discriminator;
- unsupported artifact structure;
- missing required capability slot;
- provided function ID does not exist;
- provided function signature is incompatible with its slot contract;
- module structure is incompatible with a structure-specific slot;
- a module-bound narrowing contract does not map the module structure to exactly one declared union variant type;
- module-bound invocation receives a union value whose active variant type disagrees with the selected module structure;
- a sequential `ArtifactReference` success value lacks the original `SequentialBaseReference` or its artifact-specific field map is inconsistent with the module's canonical segment interpretation;
- a tree `ArtifactReference` success value contains anything other than the original `TreeBaseReference` as its complete identity representation;
- the retained base discriminator disagrees with the registered artifact module.

These are runtime configuration failures, not record validation findings.

## Artifact extension rule

Adding a supported artifact kind should require:

- one artifact module manifest;
- functions that satisfy the required capability slots;
- artifact-owned declarations and validation recipes;
- successful bootstrap registration.

It should not require adding an artifact-kind switch to generic YAML processing or operation code.

## Initial exclusions

The initial module model does not provide:

- unrestricted third-party runtime plugins;
- dynamic module installation during operation execution;
- version negotiation between modules;
- module inheritance;
- design-facing reflection over provided capability names;
- optional fallback to another artifact module after artifact-specific parsing fails.
