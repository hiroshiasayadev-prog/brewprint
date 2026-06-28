# DRMCP-REQ-MCP-001: Multi-root and multi-namespace read contract realignment

- **id**: DRMCP-REQ-MCP-001
- **status**: accepted
- **date**: 2026-06-25
- **source_refs**:
  - V01-REQ-MCP-033
  - DRMCP-INV-MCP-001
  - DRMCP-INV-MCP-002
  - DRMCP-ADR-MCP-001
  - PRODUCT-ADR-SPEC-001
  - spec:product.design_records.namespace_model
  - spec:product.design_records.repository_layout
  - spec:product.design_records.traceability
  - spec:product.design_records.spec_format
  - spec:product.brewprint.compatibility.legacy_id_compatibility
- **work_items**:
  - DRMCP-WORK-MCP-001
  - DRMCP-WORK-MCP-003
  - DRMCP-WORK-MCP-004
  - DRMCP-WORK-MCP-005
  - DRMCP-WORK-MCP-006
  - DRMCP-WORK-MCP-007
  - DRMCP-WORK-MCP-008
  - DRMCP-WORK-MCP-009
  - DRMCP-WORK-MCP-010
  - DRMCP-WORK-SPEC-001
  - DRMCP-WORK-SPEC-002

## Requirement

DRMCP read-oriented operations must provide one current-format contract across multiple active app roots and optional legacy archive roots.

The contract must align discovery, indexing, query, exact retrieval, reference resolution, and validation with `DRMCP-ADR-MCP-001`.
It must not preserve legacy behavior merely because the current implementation or tests encode it.

Current records and legacy archive records require separate indexes and separate operational scopes.
Legacy compatibility must remain explicit, read-only, and configuration-gated.

## Evidence

- `DRMCP-INV-MCP-001` identified unresolved multi-root query, namespace, identity, and validation behavior.
- `DRMCP-INV-MCP-002` found stale namespace assumptions, obsolete spec identity behavior, mixed compatibility fixtures, and incomplete resolver contracts.
- `DRMCP-ADR-MCP-001` established current-format-first resolution, a separate legacy archive index, accepted legacy sequential-ID families, and normal path hiding.
- PRODUCT specs own canonical namespace, identity, layout, traceability, and spec-format semantics.
- Brewprint compatibility specs own project-specific accepted legacy families.

## Required Outcome

### Current active index

DRMCP must define one active index over configured current records roots.

The active index must:

- discover current records under configured app roots;
- index decision, spec, investigation, requirement, work-item, and task records where supported by corrected contracts;
- derive current spec identity from the current path-derived `spec:` format;
- reject YAML front matter as an active spec metadata source;
- reject duplicate canonical identity across current roots;
- preserve fully qualified app-aware public IDs and active `spec:` refs;
- exclude legacy archive roots from normal active discovery.

Duplicate identity behavior must be deterministic and diagnostic.
The contract must not select one duplicate record by filesystem order.

### Query and namespace scope

Read-oriented tools must define explicit namespace and root scope.

The corrected contracts must specify:

- supported namespace filters;
- supported artifact-kind and status filters;
- default behavior when namespace filters are omitted;
- cross-namespace result ordering;
- duplicate and partial-result behavior;
- range behavior for sequential artifact IDs;
- rejection of unsupported range forms for specs and other non-sequential identities.

Normal listing must query the active index only.
Legacy archive records must not appear in normal listing results.

### Exact retrieval

Exact retrieval must remain distinct from reference resolution.

`get_record` and `get_records` must define:

- accepted current exact IDs and active `spec:` refs;
- exact accepted legacy-ID retrieval when legacy fallback is enabled;
- partial-result behavior for batch retrieval;
- duplicate requested-input behavior;
- stable response ordering;
- unsupported and unresolved item behavior.

Exact batch retrieval must not perform fuzzy normalization or infer a missing app prefix.

### Current-first reference resolution

Reference resolution must use this order:

1. Parse and resolve the input through current canonical grammar and the active index.
2. If unresolved, test the input against an accepted legacy grammar.
3. Query the legacy archive index only after an exact accepted legacy grammar match.
4. Return unresolved or unsupported when neither grammar accepts the input.

Accepted legacy fallback families are limited to:

- `V01-ADR-*`;
- `V01-INV-*`;
- `V01-REQ-*`;
- `V01-WORK-*`;
- `V01-TASK-*`.

The resolver must reject:

- `V01-SPEC-*`;
- app-prefixless bare IDs;
- physical-path inputs as canonical references;
- fuzzy prefix repair;
- legacy YAML semantic-ref aliases.

Legacy resolution preserves the issued legacy ID.
It must not translate a legacy ID into a current ID.

### Legacy archive configuration and index

Legacy fallback must be disabled unless configuration declares one or more `legacy_roots`.
DRMCP must not auto-discover `v01/` or any other archive directory.

The legacy archive index must remain separate from the active index.
Legacy archive records must be:

- read-only;
- excluded from normal listing;
- excluded from current repository-wide validation;
- excluded from create and update targets;
- available only for exact accepted legacy retrieval and fallback resolution.

The configuration contract must define invalid, missing, duplicate, and overlapping current/legacy root diagnostics.

### Cross-namespace and legacy relation validation

Current records may reference current records in another configured app namespace.
Current records may also reference an accepted legacy sequential ID when legacy fallback is enabled.

Validation must define:

- current cross-namespace relation resolution;
- accepted current-to-legacy relation resolution;
- unresolved and unsupported reference diagnostics;
- duplicate identity diagnostics;
- behavior when a legacy root is not configured;
- distinction between semantic invalidity and DRMCP diagnostic representation.

Legacy archive records themselves are not part of normal current repository validation.

### Current spec behavior

Current specs must be discovered, retrieved, resolved, and validated through:

- H1-adjacent metadata;
- path-derived canonical `spec:` refs;
- current PRODUCT-owned spec kinds and section contracts.

Legacy YAML-front-matter specs and `V01-SPEC-*` aliases must not remain runtime compatibility inputs.

### Path exposure boundary

Normal listing and retrieval responses must use canonical identity rather than physical paths.

Physical paths may appear only where operationally required, including:

- explicit patch output;
- diagnostics that require a source location;
- debug or emergency inspection output.

The corrected tool contracts must define this exception narrowly.

### Diagnostics and fixtures

DRMCP must define machine-readable diagnostics for at least:

- invalid current root configuration;
- invalid legacy root configuration;
- duplicate canonical identity;
- unsupported input grammar;
- unresolved current reference;
- unresolved accepted legacy reference;
- legacy fallback disabled;
- unsupported legacy family;
- cross-namespace relation failure;
- current spec source-format violation.

Fixtures must distinguish:

- current app-aware sequential IDs;
- current path-derived spec refs;
- accepted V01 sequential IDs;
- rejected `V01-SPEC-*` inputs;
- rejected app-prefixless bare IDs;
- current-only operation without legacy roots;
- configured legacy fallback behavior.

### Follow-up tracking

A coordinating DRMCP Work Item must split contract correction, fixtures, implementation, and independent review into reviewable tasks.

The Work Item must coordinate any replacement of `DRMCP-WORK-SPEC-001` or `DRMCP-WORK-SPEC-002` with the PRODUCT validation-policy owner pointers.

## Explicitly Excluded Scope

- Authoring transaction create or update behavior owned by `DRMCP-REQ-MCP-002`.
- Portable standards-package loading or guidance projection owned by `DRMCP-REQ-MCP-003`.
- Bulk migration or renaming of issued legacy sequential IDs.
- Migration of legacy YAML specs into the active spec tree.
- Reintroduction of `V01-SPEC-*` aliases.
- Fuzzy reference repair.
- UI behavior.
- BPDSL design or migration.

## Boundary

PRODUCT owns:

- canonical identity and reference semantics;
- app and domain namespace semantics;
- repository layout semantics;
- traceability semantics;
- current spec format and invalid states.

Brewprint compatibility specs own:

- accepted project-specific legacy families;
- archive and migration facts.

DRMCP owns:

- configured root loading;
- discovery and index construction;
- query and exact retrieval contracts;
- resolution order and lookup implementation;
- validation execution;
- diagnostic categories and response representation;
- path-exposure behavior in MCP responses.

DRMCP specs must cite the semantic authorities instead of restating their rules as independent authority.
