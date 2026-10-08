# Executable implementation design working notes

## Status

- Working notes created on 2026-07-17.
- This directory is not a canonical Design Record area.
- It preserves the current proposal that declarative YAML becomes the canonical implementation-design authority and that human-readable implementation Specifications are generated from those declarations.
- Existing public operation, diagnostic, artifact, and Product Specifications remain authoritative for externally observable behavior and artifact semantics.
- When these notes conflict with a later accepted ADR or Specification, the accepted canonical artifact wins.

## Problem

The current design flow has a missing layer between external operation Specifications and executable implementation.

```text
public operation and diagnostic Specifications
  -> missing implementation-realization layer
  -> executable YAML declarations
  -> narrow native Go implementations
```

Filling that gap with separately authored implementation Specifications would duplicate the same design in prose and YAML:

```text
external Specification
  -> handwritten implementation Specification
  -> YAML transcription
  -> Go implementation
  -> manual consistency checking
```

That model creates three predictable failures:

- the prose design and executable declaration drift;
- implementation feasibility is discovered only after the prose design is accepted;
- function, type, context, and native-binding coverage must be checked manually.

## Proposed direction

Treat YAML declarations as executable implementation design rather than configuration.

```text
external operation and artifact Specifications
  -> top-down YAML implementation design
  -> generated human-readable implementation projection
  -> native Go leaves
```

The YAML design begins at one public operation and is refined top down:

```text
operation entry function
  -> coarse semantic processing DAG
  -> narrower composite functions
  -> typed values and expected outcome unions
  -> native primitive leaves
```

A design slice closes when every referenced declaration exists, every call is type-correct, every expected union is handled exhaustively, every required context is available, and every leaf function has either a declarative body or a valid native binding.

## Authority boundary

| concern | authority |
|---|---|
| Public request, response, ordering, limits, warnings, diagnostics, and operation outcomes | Existing DRMCP operation and diagnostic Specifications. |
| Artifact identity, structure, source, and validation semantics | Existing Product and DRMCP artifact Specifications. |
| Runtime declaration language and bootstrap rules | Validation runtime working specification set, later routed into canonical authority. |
| Operation realization, semantic processing DAG, typed handoffs, context use, and native boundaries | Executable YAML declarations. |
| Human-readable function, type, DAG, and coverage documentation | Generated projection from YAML; not independently edited authority. |
| Concrete filesystem, parsing, comparison, collection, and other primitive behavior | Native Go binding implementations constrained by their YAML declarations. |

Generated Markdown must not redefine or override either the external Specifications or the YAML declarations from which it was produced.

## Core traceability graph

Every design-facing declaration participates in a traceability graph:

```text
external requirement or Specification responsibility
  -> YAML function, type, context, event, or module declaration
  -> declarative implementation or native binding
  -> generated implementation-design projection
  -> executable acceptance evidence where available
```

The proposed `implements` relation makes this graph queryable and permits USDM-like coverage reporting at declaration granularity.

Coverage proves an explicit implementation-design relationship. It does not alone prove correct behavior. Behavioral evidence requires bootstrap validation, executable scenarios, or tests.

## Documents

| document | responsibility |
|---|---|
| `design-workflow.md` | Top-down YAML-first design process, refinement loop, closure conditions, and review outputs. |
| `traceability-format.md` | Proposed declaration-level `implements` metadata, reference rules, coverage directions, and validation behavior. |
| `generated-projection.md` | Human-readable Markdown generated from operation, function, type, context, and coverage declarations. |
| `partial-design-validation.md` | Authoring, design-closure, and activation validation profiles; declaration-required authoring; typed stubs; context validation; and partial-generation rules. |
| `open-questions.md` | Decisions intentionally left unresolved until the first operation is designed through native leaves. |

## Immediate consequence

The current ArtifactReference resolution discussion should not continue by inventing isolated result types.

The first operation slice should instead declare a top-level operation function with a coarse semantic DAG. Reference parsing, current-record resolution, tree-node resolution, source access, diagnostic projection, and any required types are then introduced only when demanded by an actual call edge in that DAG.

This makes the design itself executable and exposes missing functions, types, contexts, outcome branches, and Specification coverage while the design is still being refined.
