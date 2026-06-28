# DRMCP-REQ-MCP-002: Namespace-aware authoring transaction conformance

- **id**: DRMCP-REQ-MCP-002
- **status**: accepted
- **date**: 2026-06-25
- **source_refs**:
  - DRMCP-INV-MCP-002
  - DRMCP-ADR-MCP-001
  - PRODUCT-ADR-SPEC-001
  - PRODUCT-REQ-SPEC-003
  - DRMCP-REQ-MCP-003
  - spec:product.design_records.authoring_standards
  - spec:product.design_records.authoring_standards.adr_authoring
  - spec:product.design_records.authoring_standards.requirement_authoring
  - spec:product.design_records.authoring_standards.work_item_authoring
  - spec:product.design_records.authoring_standards.task_authoring
  - spec:product.design_records.authoring_standards.spec_authoring
  - spec:product.design_records.authoring_standards.investigation_authoring
  - spec:product.design_records.namespace_model.artifact_id_grammar
  - spec:product.design_records.repository_layout
  - spec:product.design_records.spec_format
  - PRODUCT-WORK-SPEC-011
- **work_items**: []

## Requirement

DRMCP authoring transactions must conform to current PRODUCT-owned authoring, identity, placement, and format semantics.

The authoring surface must use canonical artifact identity as the public input boundary.
It must generate H1, metadata serialization, filename, and physical placement internally.

The redesigned contract must preserve the useful proposal transaction mechanics while removing legacy spec-format assumptions, normal path exposure, and duplicated semantic authority.

## Evidence

- `DRMCP-INV-MCP-002` found a conflict between the read-only MVP and authoring trial, stale YAML spec updates, incomplete artifact-kind support, and duplicated PRODUCT semantics.
- `DRMCP-ADR-MCP-001` accepted the current delivery sequence and retained the propose-then-accept transaction boundary.
- PRODUCT authoring standards distinguish author inputs, generated values, partial updates, persisted state, canonical headings, and lifecycle gates.
- PRODUCT namespace and repository-layout specs own sequence and placement semantics.
- PRODUCT spec-format contracts require H1-adjacent metadata and path-derived `spec:` identity.

## Required Outcome

### Delivery sequence

Authoring support must be delivered in this order:

1. ADR, requirement, work-item, and task authoring after the corrected current read baseline.
2. Spec authoring.
3. Investigation authoring.

A later artifact-kind phase may depend on an earlier phase.
An implementation must not claim support for a kind or operation before its corrected contract and fixtures are complete.

### Shared proposal transaction

Authoring writes must use a propose-then-accept model.

The corrected contract must retain:

- proposal creation without repository writes;
- explicit accept-only filesystem writes;
- proposal retention and expiry;
- accept-time target and staleness checks;
- body cache and retry behavior;
- patch, summary, and omitted diff modes;
- affected-record candidate-state validation;
- exact named-section replacement;
- machine-readable diagnostics.

Proposal mechanics inherited from the authoring trial are design inputs under `DRMCP-ADR-MCP-001`.
They are not independent authority where they conflict with current PRODUCT semantics.

### Create, update, and persisted-state separation

Each supported artifact kind must define separate schemas for:

- create input;
- partial update input;
- persisted record state.

The authoring contract must not require every persisted field in a partial update.
Omitted update fields remain unchanged.
Generated identity and placement values are not ordinary metadata update targets.

Semantic dates remain author-controlled.
Editorial updates must not change semantic dates automatically.

### Sequential artifact identity

ADR, requirement, work-item, task, and investigation creation must support exact IDs and the applicable `new` placeholder.

For non-task sequential artifacts, server-side allocation uses the PRODUCT-owned scope:

```text
app namespace + artifact kind + domain namespace
```

Task allocation follows the explicit parent work item.
Task creation must not infer its parent only from ID shape.

The contract must define:

- exact-ID and `new` input forms by kind;
- collision and sequence-gap behavior;
- task parent validation;
- concurrency and accept-time stale-allocation guards;
- generated resolved-ID output;
- reciprocal workflow relation updates when required.

`suggest_next_record` must not remain as a tool or compatibility surface.

### Generated placement and path boundary

DRMCP must generate physical placement from canonical identity and PRODUCT-owned layout rules.

Authors do not supply generated:

- public IDs resolved from `new`;
- H1 headings;
- serialized metadata blocks;
- filenames;
- repository paths.

Normal authoring requests and proposal-summary responses must not use physical paths as their primary contract.

Physical paths may appear only in:

- explicit patch output;
- diagnostics that require a source location;
- debug or emergency inspection output.

### Workflow artifact authoring

The first authoring phase must support create and supported updates for:

- ADR;
- requirement;
- work item;
- task.

The contract must consume the PRODUCT per-kind authoring standards for:

- metadata meaning and requiredness;
- lifecycle values and gates;
- canonical H2 sections;
- public ID grammar;
- relationship semantics;
- generated placement.

DRMCP must define request schemas, response schemas, parser mappings, proposal behavior, write guards, and diagnostics.
It must not restate PRODUCT rules as independent DRMCP semantics.

### Portable package dependency

Authoring contract design may proceed before the portable standards package is implemented.
Runtime authoring implementation must consume a validated package produced under `PRODUCT-REQ-SPEC-003` and loaded under `DRMCP-REQ-MCP-003`.
Runtime implementation must not hard-code PRODUCT semantics or assume the Brewprint repository layout.

### Current spec authoring

Spec authoring must use the current H1-adjacent format only.
YAML front matter is prohibited for spec create and update.

Spec creation must accept a logical create selector:

| selector | target | persisted canonical ref |
|---|---|---|
| `spec:<segments>` | Leaf spec | `spec:<segments>` |
| `spec:<segments>.index` | Topic index spec | `spec:<segments>` |

Without `.index`, the selector always targets a leaf spec.
DRMCP must not infer leaf or topic placement from repository state.

The `.index` suffix is create-only and must not persist in the canonical ref.

The contract must define:

- app namespace validation;
- logical selector parsing;
- generated current spec path and H1;
- parent existence and parent-ref validation;
- current spec kind and `contract_class` handling;
- current required-section validation;
- H1-adjacent metadata serialization;
- rejection of YAML front matter;
- target canonical-ref collision checks;
- physical-path collision checks;
- accept-time target staleness guards.

A leaf and topic index must not share one canonical ref.
The contract must reject both:

- `<topic>.md` when `<topic>/index.md` exists;
- `<topic>/index.md` when `<topic>.md` exists.

A topic index may coexist with child specs below that topic.
A create proposal must never overwrite an existing spec path or canonical ref.

Spec metadata updates must target current H1-adjacent metadata.
Legacy YAML metadata replacement must not remain as a compatibility update operation.

### Investigation authoring

Investigation authoring must be delivered after spec authoring.

The contract must support exact IDs and `<APP>-INV-<DOMAIN>-new`.
It must define create, metadata update, named-section update, and persisted-state validation behavior.

DRMCP must consume the PRODUCT investigation conclusion-readiness rule.
For `status: concluded`, required sections must exist and contain substantive content.
Placeholder-only content must not satisfy the gate.

### Body source and section replacement

Create requests must use structured fields plus optional section-only body content.

The contract must define:

- `fields` as required create input;
- exactly one of `body` or `body_cache_id` when section content is supplied;
- rejection of body-only or cache-only create;
- exclusion of H1 and metadata from body input;
- cache preservation on retryable preparation failure;
- canonical named-section selectors;
- zero-match, duplicate, and ambiguous heading diagnostics;
- no-op update behavior;
- candidate headings when selector repair is possible.

Legacy non-English headings must not silently become canonical headings for new records.
Any compatibility handling must be explicit and update-only.

### Validation and write guards

Proposal-local validation must cover only the affected record set in candidate state.
Unrelated repository diagnostics must not block acceptance.

Accept must re-check:

- proposal lifecycle state;
- target file state;
- resolved identity availability;
- target kind and identity;
- sequence or selector collision;
- required reciprocal updates;
- affected-record validation.

A failed accept check must return `written: false` and must not modify repository files.

### Diagnostics and fixtures

DRMCP must define machine-readable diagnostics for at least:

- invalid app or domain namespace;
- unsupported artifact kind or operation;
- invalid exact ID or `new` placeholder;
- unresolved or invalid task parent;
- sequence collision or stale allocation;
- invalid generated-value input;
- invalid logical spec selector;
- leaf/topic collision;
- existing canonical ref or path;
- prohibited YAML front matter;
- metadata-state violation;
- canonical-heading violation;
- proposal staleness;
- affected-record validation failure.

Fixtures must cover each delivery phase independently.
Legacy implementation fixtures must not define the corrected contract.

### Follow-up tracking

A coordinating DRMCP Work Item must split workflow authoring, spec authoring, investigation authoring, fixtures, implementation, and independent review into separate reviewable tasks.

## Explicitly Excluded Scope

- Query, exact retrieval, resolver, and current/legacy index behavior owned by `DRMCP-REQ-MCP-001`.
- Portable standards-package production owned by `PRODUCT-REQ-SPEC-003`.
- Portable package loading, validation, and guidance projection owned by `DRMCP-REQ-MCP-003`.
- Defining artifact meaning, lifecycle semantics, ID grammar, or canonical section requirements in DRMCP.
- Bulk migration of existing records.
- Migration or authoring of legacy YAML-front-matter specs.
- Reintroduction of `suggest_next_record`.
- Arbitrary multi-record atomic transactions with rollback semantics.
- Natural-language quality assessment beyond PRODUCT-defined structural gates.
- UI behavior.
- BPDSL design or migration.

## Boundary

PRODUCT owns:

- artifact semantics and lifecycle rules;
- authoring standards;
- app and domain namespace semantics;
- canonical ID and sequence semantics;
- repository layout semantics;
- metadata meaning and author-facing requiredness;
- canonical section headings;
- current spec format and logical identity semantics.

PRODUCT also owns portable standards-package production under `PRODUCT-REQ-SPEC-003`.
DRMCP package loading and runtime semantic access are owned separately by `DRMCP-REQ-MCP-003`.

DRMCP owns:

- authoring request and response schemas;
- exact-ID and placeholder processing;
- logical selector processing;
- generated placement implementation;
- proposal lifecycle and caches;
- diff representation;
- write guards and accept behavior;
- parser and serialization mappings;
- diagnostics;
- operation support by artifact kind.

DRMCP contracts must cite PRODUCT semantic sources instead of duplicating their normative definitions.
