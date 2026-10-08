# Generated implementation-design projection

## Status

This document defines the intended human-readable projection generated from executable YAML declarations.
It does not define the final generator implementation or exact Markdown syntax.

## Purpose

Provide reviewable implementation-design documents without maintaining a second handwritten design authority.

Generated Markdown must let a reviewer understand:

- what one operation or function does;
- which functions it calls;
- how data flows through its DAG;
- which types cross its direct boundary;
- which runtime contexts it requires;
- which external responsibilities it implements;
- which native leaves remain;
- which declarations or coverage links are unresolved.

## Authority rule

Generated files are reproducible projections.

They must contain a visible generated-source marker and must not be manually edited.
A mismatch between generated Markdown and YAML is a generator defect or stale output, not an ambiguity between two equal authorities.

The generator should support a validation mode that fails when committed generated output differs from a fresh generation.

## Operation and function page

Each operation entry function and reviewable composite function should produce one section or page with the following information.

### Identity and responsibility

- function ID;
- `detail`;
- implementation kind: declarative, native, or stub;
- direct `implements` refs;
- direct callers;
- runtime entry-point status when applicable.

### Signature and runtime dependencies

- complete signature;
- ordered context IDs opened by the function, when any;
- required context values;
- restart and event behavior, when applicable;
- native binding ID for a native function.

### Execution DAG

The declarative body is rendered as a source-ordered graph or equivalent structured flow.

The graph should preserve:

- call order;
- typed local names;
- field projections;
- `if`, exhaustive union `switch`, and bounded `for` boundaries;
- branch-local continuation or return;
- final return value.

The generator must not invent algorithmic edges that are hidden inside one native function.

### Called-function table

At minimum:

| function | detail | implementation | direct authority |
|---|---|---|---|
| called function ID | semantic detail | YAML or native | direct `implements` refs |

This table is generated from direct call edges only.
A separate reachability view may show transitive dependencies.

### Direct type surface

For each input, output, direct local boundary, expected union, or projected transparent record used by the function, show:

- type ID;
- type kind;
- top-level fields or union variants;
- direct authority refs;
- link to the canonical generated type page.

Transparent records are expanded one level only.
Nested record fields show their type names rather than recursively expanding the complete object graph.

Example:

```text
ResolvedArtifactReference
  reference: ArtifactReference
  record: CurrentRecord
```

The page does not recursively print every field of `CurrentRecord`.

### Outcome handling

For every consumed union, show:

- all declared variants;
- the function or switch case handling each variant;
- whether a variant becomes a normal value, item-level result, operation diagnostic, or continued iteration;
- whether any handling remains unresolved.

Execution failures are shown separately from expected union variants.

### Native leaf inventory

List every native function directly called by the function and optionally every transitively reachable native leaf.

For each leaf show:

- YAML function ID;
- native binding ID;
- signature;
- `detail`;
- authority refs;
- activation or implementation status.

### Design gaps

A partial design page visibly reports:

- reachable stub functions;
- unresolved native bindings;
- unavailable contexts;
- invalid `opens` order or same-context reopening;
- provider dependency cycles;
- non-exhaustive unions;
- invalid projections;
- missing required traceability;
- uncovered required external responsibilities.

Unknown calls and types are validation errors and do not become supported partial-design nodes.

Generation should remain possible for partial designs so the gaps can be reviewed before closure.

## Type page

Each named type should produce one canonical page containing:

- type ID and `detail`;
- scalar, enum, transparent record, opaque, or union kind;
- generic parameters and constraints;
- complete top-level fields or variant list;
- direct `implements` refs;
- constructing functions;
- consuming functions;
- types referenced directly by its fields;
- native conformance status where applicable.

Opaque types show no hidden native fields.
The page describes only the design-facing contract.

## Context, event, and module pages

Generated pages should also cover:

### Context

- provided values;
- providers and destructors;
- active-context-chain relationships and provider dependencies;
- invalidating events;
- functions opening or requiring the context.

### Event

- payload type;
- invalidated contexts;
- restartable entry functions;
- direct authority refs.

### Artifact module

- module identity and artifact kind;
- registered capability slots;
- bound functions and signatures;
- structure-specific narrowing guarantees;
- direct authority refs.

## Coverage projections

The generator should produce at least two complementary views.

### Authority-first view

```text
authority responsibility
  -> direct YAML declarations
  -> operation entries reaching them
  -> scenarios or tests providing evidence
```

### Declaration-first view

```text
declaration
  -> direct authority refs
  -> callers and callees
  -> generated document location
  -> declarative body or native binding
```

Coverage pages should distinguish direct implementation, transitive reachability, and executable evidence rather than collapsing them into one boolean.

## Proposed generated layout

The physical output location is not decided, but the logical projection is:

```text
generated/
  operations/
  functions/
  types/
  contexts/
  events/
  modules/
  coverage/
  incomplete/
```

A later canonical Design Record layout may group or split these differently.
The generator model should not depend on one Markdown directory layout.

## Determinism

Generation must be deterministic with respect to:

- declaration ordering;
- call and type-edge ordering;
- authority-ref ordering;
- field and variant ordering;
- generated filenames or refs;
- gap-report ordering.

Filesystem enumeration order must not affect generated output.

## Non-goals

- Replacing public operation or artifact Specifications.
- Expanding opaque native representations.
- Generating implementation rationale that was never declared.
- Treating a pretty DAG as proof of behavioral correctness.
- Manually editing generated function or type pages.
