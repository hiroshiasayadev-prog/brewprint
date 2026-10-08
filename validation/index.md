# Validation design working notes

## Status

- Working notes created on 2026-07-15.
- This directory is not a canonical Design Record area.
- The documents preserve the completed artifact-constraint investigation, user-confirmed gap decisions, and the current validation capability direction until they are routed into ADRs and Specifications.
- When these notes conflict with a later accepted ADR or Specification, the accepted canonical artifact wins.

## Purpose

The immediate goal is to avoid losing decisions that currently exist only in conversation while the validation component model is still being designed.

These notes cover:

- the six currently modeled artifact kinds and their validation surfaces;
- gaps found while comparing DRMCP artifact Specifications with Product authority;
- the accepted resolution of each gap;
- the decision to compose validation from reusable atomic components and declarative YAML recipes;
- the proposed capability namespaces and directory structure;
- boundaries between generic facts, artifact meaning, generic validation checks, findings, and operation diagnostics;
- the proposal to treat YAML declarations as executable implementation-design authority and generate human-readable implementation projections and coverage from them;
- known current-spec mismatches and intentionally deferred engine details.

## Documents

| document | content |
|---|---|
| `artifact-constraint-inventory.md` | Investigation scope, source inventory, artifact-by-artifact constraint summary, and GAP-001 through GAP-016 resolutions. |
| `capability-architecture.md` | YAML composition direction, capability layering, proposed directories, component boundaries, and confirmed versus provisional runtime decisions. |
| `atomic-component-inventory.md` | User-confirmed atomic capability inventory, category boundaries, and rejected or merged candidates. |
| `runtime-execution-model.md` | Discussion-derived overview of the runtime and artifact-module direction. |
| `runtime/index.md` | Working specification set and current consistency baseline for runtime values, functions, artifact modules, execution, invocation, and control flow. |
| `executable-design/index.md` | YAML-first executable implementation-design proposal, declaration-level traceability, generated documentation, and coverage direction. |
| `next-steps.md` | Canonical synchronization needs, next design sequence, and details intentionally deferred until the component inventory is complete. |

## Current high-level conclusion

Validation should not be implemented as one artifact-specific monolith or as hard-coded operation logic.

The current direction is:

```text
artifact source and repository state
  -> reusable facts and transforms
  -> artifact-owned validation recipe
  -> generic validation checks
  -> validation finding projection
  -> validate_scope / validate_records operation projection
```

The artifact contract owns what a value means and whether a condition is a `contract_violation` or an `advisory`.
Generic components only extract, transform, compare, resolve, enumerate, or analyze values.
Validation operations select subjects, execute the applicable validation contracts, aggregate results, and project the externally defined result shapes.

## Important current mismatch

The latest accepted direction makes a missing `index.md` on every tree directory node a contract violation while preserving the node and subtree as navigable current-state structure.

`drmcp/records/spec/design-records-mcp/current-record-model/current-record-scopes.md` currently classifies the same absence as an advisory.
That Specification is therefore stale relative to these working decisions and requires an explicit canonical update route before implementation planning.
