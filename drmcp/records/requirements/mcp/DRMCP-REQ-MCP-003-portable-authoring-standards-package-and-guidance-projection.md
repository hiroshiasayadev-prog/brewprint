# DRMCP-REQ-MCP-003: Portable authoring standards package loading and guidance projection

- **id**: DRMCP-REQ-MCP-003
- **status**: accepted
- **date**: 2026-07-04
- **source_refs**:
  - DRMCP-INV-MCP-002
  - DRMCP-ADR-MCP-001
  - DRMCP-TASK-MCP-016-09
  - DRMCP-TASK-MCP-016-12
  - PRODUCT-ADR-SPEC-001
  - PRODUCT-REQ-SPEC-003
  - spec:product.design_records.authoring_standards
  - spec:product.design_records.namespace_model
  - spec:product.design_records.repository_layout
  - spec:product.design_records.spec_format
  - spec:product.design_records.traceability
  - spec:product.design_records.artifact_model

## Requirement

DRMCP authoring and semantic validation must consume the portable Design Records standards package produced under `PRODUCT-REQ-SPEC-003`.

DRMCP must register the selected package spec tree as ordinary Current Records under the reserved `design_records` app namespace.
The package must use the normal current Spec discovery, identity, logical-tree, active-index, retrieval, resolution, and validation semantics.

The runtime boundary must work outside the Brewprint repository.
It must not depend on a host repository's `product` namespace, Brewprint compatibility history, or legacy guide directory.
DRMCP must not rewrite package refs at runtime.
Root resolution must not depend on the process working directory.

Authoring guidance tools must project indexed Current Records under `spec:design_records.authoring_standards`.
They must not read a separate guide source or maintain a package-specific record model.

## Evidence

- `DRMCP-INV-MCP-002` found incompatible legacy guide and PRODUCT Spec sources.
- `DRMCP-ADR-MCP-001` requires a portable fixed-namespace distribution of PRODUCT semantics.
- `PRODUCT-REQ-SPEC-003` produces a whole-tree copy at `<exe-dir>/design-records/` and rewrites refs to `spec:design_records.*`.
- The package files follow the same current Spec format used by ordinary Current Records.
- T09 found that the direct PRODUCT-directory Guidance model contradicted accepted package authority.
- T12 selected normal Current Records treatment and fixed-scope Guidance aliases.
- A second package index, parser, logical tree, or Guidance source would duplicate the current record model.

## Required Outcome

### Current Records source registration

DRMCP configuration must select one portable standards spec tree and associate it with `app_namespace: design_records`.

The selected source may be:

- the bundled default `<exe-dir>/design-records/`; or
- one explicit configured override.

An explicit override disables silent fallback to the bundled default.
Root resolution must be deterministic and independent of the process working directory.
The exact configuration serialization remains a downstream contract.

The selected package source is mandatory when package-dependent capabilities are enabled.
It follows the same source-availability and trustworthy-index rules as other mandatory Current Records sources.

### Normal current Spec processing

The package root is a configured spec-tree source.
DRMCP must discover Markdown files recursively under that root.

```text
<package-root>/**/*.md
```

The package root itself is the effective current Spec tree root.
`<package-root>/index.md` derives `spec:design_records`.
Child paths derive `spec:design_records.<suffix>` through the normal current Spec identity rules.

Package Specs must enter the same active Current Records index as other current records.
The index retains the explicit `design_records` app association.
DRMCP must not create a package-specific index, logical tree, parser, resolver, validator, or record model.

Normal Current Records behavior governs:

- invalid but addressable Specs;
- duplicate canonical identity;
- unresolved refs;
- exact retrieval;
- reference resolution;
- validation inputs and findings;
- physical-path hiding in normal responses.

Legacy Archive remains a separate compatibility state and does not absorb package Specs.

### Fixed package namespace

The producer supplies visible IDs and refs under `spec:design_records`.
DRMCP must not select, remap, or rewrite that namespace at runtime.

An unrewritten `spec:product.design_records.*` value receives normal current Spec mismatch or unresolved-ref handling.
It does not create a package alias or trigger consumer repair.

### Authoring guidance aliases

Authoring guidance operations must use shared record-query orchestration over the Current Records snapshot.
They must not call another public use case.

The fixed list scope is:

- app namespace: `design_records`;
- kind: `spec`;
- canonical subtree: `spec:design_records.authoring_standards.*`;
- excluded list root: `spec:design_records.authoring_standards`.

The fixed detail scope accepts an exact canonical ref under that child subtree.

Guidance projection is:

| field | rule |
|---|---|
| `id` | Canonical package Spec ref. |
| `title` | First H1 text. |
| `abstract` | Body of the `## What this is` section. |
| `content` | Complete source Markdown verbatim. |

List ordering uses canonical-ref ASCII lexical order.
Normal Guidance responses do not expose physical paths.
Legacy guide files, filename-stem IDs, title lookup, fuzzy lookup, and inferred candidates are not current Guidance authority.

### Integration with authoring and validation

`DRMCP-REQ-MCP-002` authoring operations and package-dependent validation must consume the same request-scoped Current Records state.

A proposal created against one package state must not be silently accepted against changed semantics.
Detailed staleness, caching, and accept-time checks remain downstream authoring contracts.
Runtime implementation must not hard-code PRODUCT semantics or assume the Brewprint repository layout.

### Distribution consumption and portability verification

A package-dependent DRMCP distribution must include or obtain a package released under `PRODUCT-REQ-SPEC-003`.

Verification must cover:

- loading outside the Brewprint repository;
- loading with a different process working directory;
- a host with no `product` app namespace;
- bundled and explicit package roots;
- explicit `design_records` app association;
- recursive current Spec discovery;
- normal active-index and invalid-source behavior;
- no runtime namespace rewrite;
- exact retrieval and resolution of package Specs;
- Guidance list and detail aliases;
- authoring and validation use of package-provided semantics.

### Follow-up tracking

Downstream Work Items may still be required for:

- concrete current-source configuration serialization;
- implementation and fixtures;
- authoring and validation integration;
- proposal reproducibility guards;
- package-dependent capability exposure;
- end-to-end portability review.

PRODUCT whole-tree source selection, package production, ref rewriting, producer warnings, and generation failures remain owned by `PRODUCT-REQ-SPEC-003`.

## Explicitly Excluded Scope

- Changing the PRODUCT whole-tree source boundary.
- Selecting or rewriting the fixed package namespace at runtime.
- Producing or synchronizing the portable package.
- PRODUCT-side source checks or warning tooling.
- Copying Brewprint compatibility rules into the package.
- Defining DRMCP tool schemas inside the package.
- Replacing PRODUCT specs as semantic authority.
- Redesigning generic query, resolver, validation, or Legacy Archive semantics.
- Defining authoring transaction behavior owned by `DRMCP-REQ-MCP-002`.
- A package-specific parser, logical tree, active index, resolver, validator, or Guidance domain.
- General-purpose package registries or network distribution.
- Relative refs, namespace remapping, or package-local aliases.
- UI behavior.
- BPDSL design or migration.

## Boundary

PRODUCT owns under `PRODUCT-REQ-SPEC-003`:

- the authoritative whole-tree source;
- the fixed `design_records` package namespace;
- the bundled physical root;
- ref-prefix rewriting;
- deterministic package generation;
- producer warnings and generation failures;
- generation and review evidence.

The portable package owns no independent semantics.
It is a copied and ref-rewritten distribution of PRODUCT authority.

DRMCP owns:

- selecting and registering the package spec-tree source;
- explicit `design_records` app association;
- normal Current Records loading and active-index participation;
- package-dependent capability exposure;
- Guidance alias request and response contracts;
- authoring and validation integration;
- proposal and validation reproducibility guards;
- consumer-side diagnostics.

Brewprint profile specs own project-specific registry, compatibility, archive, and migration facts.
Those facts remain outside the package and the portable consumer contract.
