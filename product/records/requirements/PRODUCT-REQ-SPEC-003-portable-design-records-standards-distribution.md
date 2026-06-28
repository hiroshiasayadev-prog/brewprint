# PRODUCT-REQ-SPEC-003: Portable Design Records standards distribution

- **id**: PRODUCT-REQ-SPEC-003
- **status**: accepted
- **date**: 2026-06-26
- **source_refs**:
  - PRODUCT-ADR-SPEC-001
  - DRMCP-ADR-MCP-001
  - DRMCP-INV-MCP-002
  - spec:product.design_records.authoring_standards
  - spec:product.design_records.namespace_model
  - spec:product.design_records.repository_layout
  - spec:product.design_records.spec_format
  - spec:product.design_records.traceability
  - spec:product.design_records.artifact_model
- **work_items**:
  - PRODUCT-WORK-SPEC-013

## Requirement

PRODUCT-owned Design Records semantics must be distributable as a portable, operationally standalone standards package.

The package must preserve PRODUCT as the semantic authority while allowing consumers such as DRMCP to load the app-independent Design Records spec tree outside the Brewprint repository.

Package production must not create a manually maintained second standard, select or rewrite semantic content, or require a host repository to expose the `product` namespace.

For the first package, operationally standalone means package files can be distributed and discovered outside Brewprint, loading does not require a host repository or host `product` namespace, package refs are not resolved through a host app registry, and root resolution does not depend on the process working directory.

Operational self-containment does not mean that producer warnings for external, unresolved, duplicate, or unrewritten refs block generation or package use. The source-authoring target remains an app-independent and semantically coherent `design-records/` tree.

## Evidence

- Current authoring standards depend on sibling namespace, repository-layout, spec-format, traceability, and artifact-model contracts.
- Copying only `authoring-standards/` would leave unresolved transitive references and an incomplete runtime contract.
- `DRMCP-INV-MCP-002` found inconsistent guidance sources and duplicated semantic authority across PRODUCT and DRMCP records.
- `DRMCP-ADR-MCP-001` requires a portable, fixed-namespace distribution of PRODUCT semantics before portable authoring is complete.
- `DRMCP-REQ-MCP-003` consumes the package through operational loading, localized indexing behavior, guidance projection, authoring, and semantic validation without taking ownership of source correction.

## Required Outcome

### Authoritative source tree

The authoritative source for the first portable Design Records standards package is the whole tree:

```text
product/records/spec/design-records/
```

This tree owns app-independent Design Records semantics. If Brewprint wiring, DRMCP runtime facts, BPDSL rules, or other app-local material appears in this tree, the defect is in source authoring or placement. The package producer must not remove, generalize, or reinterpret that material during package creation.

Source cleanup remains a source-authoring responsibility. Semantic source defects may be reported as producer warnings, but they are not package-generation blockers and do not authorize package-time filtering, repair, or reinterpretation.

### Package production contract

The first package is a deterministic whole-tree copy:

```text
source: product/records/spec/design-records/
destination: <exe-dir>/design-records/
```

The only semantic transformation performed during package production is canonical spec-ref prefix rewriting:

```text
spec:product.design_records
  -> spec:design_records

spec:product.design_records.<suffix>
  -> spec:design_records.<suffix>
```

The rewrite applies only to canonical spec refs, including visible spec `id` values, Topics refs, Related specs refs, body tokens recognized as canonical `spec:` refs, and other canonical spec-ref fields.

The rewrite must not alter ordinary prose words, Design Record public IDs such as `PRODUCT-REQ-*`, physical path strings, non-canonical code or example strings, or any `spec:` namespace outside `spec:product.design_records`.

Full-text string replacement is prohibited.

The package-facing roots are ordinary spec refs:

- package spec tree root: `spec:design_records`
- authoring guidance root: `spec:design_records.authoring_standards`

`<exe-dir>/design-records/index.md` is the copied and ref-rewritten root overview from the source tree. It is not a package-specific metadata document.

### Generation and synchronization

Define a deterministic generation or synchronization mechanism from PRODUCT authorities to the portable package.

The mechanism must:

- prevent manual package edits from becoming independent semantic authority;
- produce reproducible package contents from the source tree;
- apply only the allowed canonical ref-prefix rewrite;
- inspect for unrewritten source-prefix refs, duplicate package refs, unresolved package refs, external canonical refs, and apparent app-local material;
- emit semantic warnings during generation or check execution when those findings exist;
- record generation and review evidence;
- keep source cleanup and package generation separately reviewable;
- continue whole-tree generation when only semantic warnings exist.

### Producer warning and operational-error boundary

Semantic findings are producer warnings rather than generation or release blockers.

Producer warnings include:

- canonical `spec:` refs outside `spec:design_records`;
- unresolved package refs;
- duplicate canonical spec refs;
- unrewritten `spec:product.design_records` refs;
- apparent app-local, wiring, migration, or project-tracking material.

Operational generation failures remain producer errors. These include unreadable source directory, unwritable or uncreatable destination, failed copy execution, failed prefix-rewrite execution, or inability to produce the package tree.

Warnings do not authorize producer-side filtering, repair, removal, generalization, or reinterpretation of source semantics. The producer emits warnings during generation or check execution, but does not own persistent warning evidence, warning disposition, or source-authoring follow-up. When a human reviewer decides that tracked source correction is required, the relevant warning output may be copied into a manually created source-authoring Requirement.

### Removed first-version scope

The first package does not define:

- a separate package ID;
- version negotiation or compatibility semantics;
- exposed package feature metadata;
- source snapshot metadata;
- a package-specific manifest;
- a package-specific entry document;
- section-level package selection;
- package-time prose generalization;
- package-time app-local semantic filtering.

### Consumer handoff

The package contract must expose enough stable roots and ref rules for DRMCP and other consumers to load and index the package as a normal spec tree with localized handling of semantic defects.

Runtime loading, configuration, indexing, guidance projection, proposal staleness, and consumer diagnostics remain owned by consumer requirements such as `DRMCP-REQ-MCP-003`. Consumers do not own source correction, do not perform package-wide semantic closure validation before use, and do not auto-repair package refs.

## Explicitly Excluded Scope

- DRMCP package-root configuration or loading implementation.
- DRMCP guidance request and response contracts.
- DRMCP authoring transaction behavior.
- Brewprint legacy-ID compatibility or archive behavior.
- General-purpose network registries or remote package distribution.
- Relative-reference, namespace-remapping, or package-local alias design for the first package release.
- BPDSL design or migration.
- Bulk migration of existing Design Records.

## Boundary

PRODUCT owns:

- authoritative source contracts;
- the whole-tree source boundary;
- fixed package namespace;
- deterministic package generation or synchronization;
- source-prefix rewrite rules;
- package warning emission and operational generation failure handling;
- package generation and review evidence.

The portable package owns no independent semantics.
It is a copied and ref-rewritten distribution artifact derived from PRODUCT authority.

DRMCP and other consumers own:

- package configuration and loading;
- operational package availability checks and localized indexing;
- runtime indexing;
- guidance projection;
- authoring and validation integration;
- consumer-side diagnostics and staleness guards.
