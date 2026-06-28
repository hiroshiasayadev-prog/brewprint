# DRMCP-REQ-MCP-003: Portable authoring standards package loading and guidance projection

- **id**: DRMCP-REQ-MCP-003
- **status**: accepted
- **date**: 2026-06-26
- **source_refs**:
  - DRMCP-INV-MCP-002
  - DRMCP-ADR-MCP-001
  - PRODUCT-ADR-SPEC-001
  - PRODUCT-REQ-SPEC-003
  - spec:product.design_records.authoring_standards
  - spec:product.design_records.namespace_model
  - spec:product.design_records.repository_layout
  - spec:product.design_records.spec_format
  - spec:product.design_records.traceability
  - spec:product.design_records.artifact_model
  - spec:drmcp.design_records_mcp.schema.authoring_guidance_source
- **work_items**: []

## Requirement

DRMCP authoring and semantic validation must load and consume a portable, operationally standalone Design Records standards package produced under `PRODUCT-REQ-SPEC-003`.

The runtime package boundary must work outside the Brewprint repository.
It must not depend on a host repository's `product` namespace, Brewprint compatibility history, or legacy guide directory.
It must not resolve package refs through a host app registry, and root resolution must not depend on the process working directory.

For this requirement, operationally standalone means package files can be distributed and discovered outside Brewprint and loaded without a host repository or host `product` namespace.
It does not mean that producer warnings for external, unresolved, duplicate, or unrewritten refs block generation or package use.

DRMCP must load the selected package root operationally and build localized indexes over readable package specs.
Authoring guidance tools must project indexed package specs rather than read a separate legacy guide source.

## Evidence

- `DRMCP-INV-MCP-002` found three incompatible guidance sources: declared `docs/guides/*.md`, implemented `records/guides/*.md`, and current PRODUCT authoring-standard spec records.
- `DRMCP-ADR-MCP-001` requires a portable, fixed-namespace distribution of PRODUCT semantics before portable authoring is complete.
- `PRODUCT-REQ-SPEC-003` owns the whole-tree source boundary, package generation or synchronization, ref-prefix rewrite rules, producer warning emission, operational generation failures, and generation and review evidence.
- DRMCP still requires consumer-side configuration, operational loading, localized indexing, capability exposure, guidance projection, and authoring integration.
- `spec:drmcp.design_records_mcp.schema.authoring_guidance_source` is cited as the legacy guidance contract being replaced, not as current semantic authority.

## Required Outcome

### Package consumption boundary

DRMCP must consume the fixed package namespace, bundled or configured physical root, package spec tree root, guidance root, prefix-rewritten canonical refs, and package-internal refs defined by `PRODUCT-REQ-SPEC-003`.

The DRMCP consumer contract must define:

- bundled and configured package-root behavior;
- deterministic root and file resolution independent of process working directory;
- fixed-namespace indexing for readable package specs;
- package-internal ref indexing with localized unresolved behavior;
- localized handling of unrewritten `spec:product.design_records` refs in package content;
- separation from host records indexes and app registries;
- package unavailability when the selected root is absent, unreadable as a directory, or cannot be enumerated.

DRMCP must not select or redefine the package namespace.
It must not rewrite package refs into the host project's namespace.

### Package configuration

DRMCP configuration must identify the standards package root.

The configuration contract must define:

- default bundled-package behavior, when provided by a distribution;
- explicit package-root override behavior;
- deterministic path resolution independent of process working directory;
- missing or unreadable package behavior;
- separation from active records roots and legacy archive roots;
- localized file warning behavior for unreadable or unparseable package files.

When an explicit configured root is supplied, DRMCP must use only that configured root and must not silently fall back to the bundled package.

When no configured root is supplied, the bundled default is `<exe-dir>/design-records/`.

If the selected package root is absent, unreadable as a directory, or cannot be enumerated, the package is unavailable. This requirement does not choose startup exit behavior, process failure behavior, or capability-degradation protocol.

Any read-only raw retrieval retained when the package is unavailable must be explicitly separated from standards-dependent operations.

### Operational loading and localized indexing

DRMCP must discover package Markdown files recursively:

```text
<package-root>/**/*.md
```

Discovery must not require Topics-graph reachability.

For an individual Markdown file that cannot be read or parsed, DRMCP must emit a warning identifying that file, exclude only that file from the package index, and continue loading and using other readable files.

`<package-root>/index.md` has no special parse-success gate. It is an ordinary root spec when readable and indexable.

A readable spec with a valid canonical ID under `spec:design_records` or `spec:design_records.*` is eligible for the canonical package index.

DRMCP must maintain package indexes separately from host records indexes and must not resolve package refs through a host app registry.

Operational loading consumes the released package contract. It does not replace PRODUCT-side generation, warning emission, or source-authoring correction.

### Duplicate, unresolved, and unrewritten-ref behavior

When multiple readable documents declare the same canonical package ref, DRMCP must not use first-wins or last-wins. It must mark only that ref ambiguous. Exact get or resolve for that ref cannot select a document, unrelated unique refs remain usable, and the whole package is not rejected.

When a readable document contains an unresolved canonical ref, the document remains readable. List and get remain available where they do not require that resolution. Only resolution of that ref returns unresolved. The document and package are not rejected.

DRMCP must not repair unrewritten source-prefix refs. As a body ref, an unrewritten `spec:product.design_records.*` ref remains unresolved. As a visible document ID, an unrewritten `spec:product.design_records.*` document is not entered into the canonical `spec:design_records` package index. Other documents remain usable and the package remains loaded.

### Authoring guidance projection

Authoring guidance must resolve from indexed package specs.

The corrected guidance contracts must define:

- package specs under `spec:design_records.authoring_standards` that qualify for authoring-guidance projection;
- list projection fields such as canonical ref, title, summary, artifact kind, and applicable operation;
- detail retrieval by exact canonical package ref;
- ordering and filtering behavior;
- unreadable file, ambiguous ref, unresolved ref, invalid, or unsupported guidance diagnostics;
- path hiding in normal guidance responses.

The guidance list scope covers readable, indexable child specs under `<package-root>/authoring-standards/`, equivalently `spec:design_records.authoring_standards.*`.

The guidance root itself, `spec:design_records.authoring_standards`, is excluded from normal guidance listing and remains available through explicit exact get.

Guidance get accepts only an exact canonical ref. The first contract does not support basename lookup, filename lookup, physical-path lookup, title lookup, fuzzy lookup, aliases, or inferred candidates.

Legacy `docs/guides/*.md` and `records/guides/*.md` files must not remain canonical guidance sources.
DRMCP must not keep a separate legacy guide parser as the current guidance implementation.

Existing guidance tool names may be retained only when their corrected behavior is a projection over indexed package specs.
Tool retention does not preserve legacy guide IDs or filename-derived identity.

### Integration with authoring and validation

`DRMCP-REQ-MCP-002` runtime authoring operations must consume the loaded and indexed package for semantic mappings.

The integration contract must define:

- unsupported operation behavior when required package contracts are unavailable;
- accept-time handling when the package changes after proposal creation;
- diagnostic ownership between package loading/indexing and record validation;
- caching and invalidation behavior without weakening correctness.

A proposal created against one loaded and indexed package state must not be silently accepted against changed semantics.
Runtime implementation must not hard-code PRODUCT semantics or assume the Brewprint repository layout.

### Distribution consumption and portability verification

A DRMCP distribution that supports package-dependent authoring or validation must include or obtain a package released under `PRODUCT-REQ-SPEC-003`.

Verification must cover:

- loading outside the Brewprint repository;
- loading with a different process working directory;
- host records with no `product` namespace;
- package root at `<exe-dir>/design-records/` when bundled;
- package root `spec:design_records`;
- guidance root `spec:design_records.authoring_standards`;
- recursive Markdown discovery without Topics reachability;
- localized warning and exclusion for unreadable or unparseable individual files;
- duplicate canonical refs marked ambiguous without whole-package rejection;
- unresolved refs kept local to the affected resolution;
- unrewritten source-prefix refs not repaired by the consumer;
- guidance list and detail projection;
- authoring and validation using package-provided mappings;
- package change between proposal and acceptance.

### Follow-up tracking

The first DRMCP Work Item for this requirement is P0 and covers only:

- consumer package contract correction;
- loader and configuration contracts;
- operational loading and localized indexing;
- minimum fixtures and portability verification.

Later Work Items may be required for:

- full guidance projection correction;
- authoring and validation integration;
- proposal reproducibility guards;
- broader package-dependent capability exposure;
- independent end-to-end review.

PRODUCT whole-tree source boundary definition, package generation or synchronization, producer warning emission, and operational generation failure handling must be tracked under `PRODUCT-REQ-SPEC-003` and must not become tasks of a DRMCP Work Item.

## Explicitly Excluded Scope

- Defining or changing the PRODUCT whole-tree source boundary.
- Selecting the fixed package namespace.
- Producing or synchronizing the portable package from PRODUCT sources.
- PRODUCT-side source/package drift detection, producer checks, or warning emission tooling.
- Copying Brewprint compatibility rules into the portable package.
- Defining DRMCP tool schemas inside the portable package.
- Replacing PRODUCT specs as semantic authority.
- Defining query, resolver, and legacy archive behavior owned by `DRMCP-REQ-MCP-001`.
- Defining artifact authoring transaction behavior owned by `DRMCP-REQ-MCP-002`.
- General-purpose package registries or network package distribution.
- Relative refs, namespace remapping, or package-local alias design for the first package release.
- UI behavior.
- BPDSL design or migration.

## Boundary

PRODUCT owns under `PRODUCT-REQ-SPEC-003`:

- authoritative whole-tree source boundary;
- producer warning emission during generation or check execution;
- fixed package namespace;
- bundled physical root;
- ref-prefix rewrite rules;
- deterministic package generation or synchronization;
- operational generation failure handling;
- generation and review evidence.

The portable package owns no independent semantics.
It is a copied and ref-rewritten distribution artifact derived from PRODUCT authority.

DRMCP owns:

- package configuration and loading;
- operational package availability checks;
- package record indexing;
- package-dependent capability exposure;
- guidance projection request and response contracts;
- authoring and validation integration;
- proposal and validation reproducibility guards;
- consumer-side package diagnostics.

Brewprint profile specs own project-specific registry, compatibility, archive, and migration facts.
Those facts remain outside the portable package and DRMCP's portable consumer contract.
