# Overview: Repository layout model

- **id**: `spec:product.concepts.repository_layout`
- **status**: draft
- **date**: 2026-06-22
- **parent**: `root`

## What this is

Defines the app namespace-first directory structure of the brewprint repository. Distinct from BPDSL project layout (`docs/spec/project-layout.md`); this spec covers the repository management structure.

PRODUCT namespace owns this spec as cross-app governance.

## Current contract

### Ownership boundary

Owned by this spec:

- App namespace-first top-level directory structure definition
- Meaning and purpose of `records/`, `dsl/`, `src/` per app namespace
- `dsl/ → src/` generation pattern intent
- `records/` internal subdirectory structure policy
- Migration target declarations for current `docs/` and `internal/` directories

Not owned by this spec:

- Actual file/directory migration (V01-WORK-PRODUCT-003)
- BPDSL DSL generation implementation
- MCP API path change handling
- Namespace catalog formal schema or machine-readable registry format
- BPDSL project layout (`docs/spec/project-layout.md`)

### Repository directory model

Top-level directories are cut by app namespace.

```
<app-namespace>/
  records/   # decision history (ADR, spec) and workflow (REQ, WORK, TASK, INV)
  dsl/       # BPDSL YAML definitions (source of truth for domain model)
  src/       # implementation code (generated from dsl/ in principle)
```

Active app namespace directories:

| directory | app namespace | nature |
|---|---|---|
| `drmcp/` | Design Records MCP | design records management MCP server |
| `bpdsl/` | Brewprint DSL | brewprint design description DSL |
| `product/` | cross-app / repository-wide | cross-app policy and governance |

`product/` has no runtime component and therefore holds `records/` only — no `dsl/` or `src/`.

### Sub-directory definitions

**records/**: Stores decision history and workflow artifacts.

- ADR (Architecture Decision Records)
- spec (specifications)
- requirement / work item / task / investigation (workflow artifacts)

Current `docs/` design records map to this directory.

**dsl/**: Stores BPDSL YAML definitions. Source of truth for domain model within the app namespace. When present, `src/` implementation is generated from `dsl/` in principle.

**src/**: Stores implementation code. Generated from `dsl/` in principle; manual code is permitted until BPDSL support catches up.

Current `internal/` Go implementation maps to this directory.

All three subdirectories are optional — place only when the app namespace has that concern.

### dsl → src generation pattern

The `dsl/ → src/` generation pattern is the long-term target for all app namespaces.

Current state:

- `bpdsl/`: has `dsl/` with BPDSL YAML definitions; `src/` is the generation target.
- `drmcp/`: has `src/` only (current `internal/designrecords/`). Will add `dsl/` and migrate to the generation pattern when BPDSL support is available.

This intent assigns meaning to directory structure: `drmcp/src/` handwritten implementation explicitly indicates a temporary state pending BPDSL support.

External precedent: the Protobuf/gRPC `proto/<domain>/` → `gen/go/` pattern (V01-ADR-097).

### records/ internal structure

`records/` internal structure retains the current kind-first layout.

```
records/
  adr/
  spec/
  requirements/
  work-items/
  tasks/
  investigations/
```

Reorganization of `records/` internal structure from kind-first to namespace-first is out of scope for this spec.

### Current state mapping

Migration targets from current `docs/` and `internal/` to target layout:

| current | target | app namespace |
|---|---|---|
| `docs/` DRMCP-related records | `drmcp/records/` | DRMCP |
| `docs/` BPDSL-related records | `bpdsl/records/` | BPDSL |
| `docs/requirements/product/` and other cross-app records | `product/records/` | PRODUCT |
| `internal/designrecords/` | `drmcp/src/` | DRMCP |

Actual migration execution, compatibility policy, and timing are governed by V01-WORK-PRODUCT-003.

## Topics

| title | kind | ref | summary |
|---|---|---|---|
| Record discovery paths | Reference | `spec:product.concepts.repository_layout.record_discovery_paths` | Kind-level path patterns for discovering record files within a records_root. |

> Source: V01-ADR-097, V01-ADR-095, V01-ADR-096, V01-REQ-PRODUCT-003, V01-WORK-PRODUCT-003
