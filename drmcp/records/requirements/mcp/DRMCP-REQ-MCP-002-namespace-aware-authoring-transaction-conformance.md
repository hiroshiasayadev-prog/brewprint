# DRMCP-REQ-MCP-002: Namespace-aware authoring transaction conformance

- **id**: DRMCP-REQ-MCP-002
- **status**: captured
- **date**: 2026-06-23
- **source_refs**:
  - spec:product.concepts.authoring_standards
  - spec:product.concepts.authoring_standards.adr_authoring
  - spec:product.concepts.authoring_standards.investigation_authoring
  - spec:product.concepts.namespace_model.v2_grammar
  - spec:product.concepts.repository_layout
  - PRODUCT-WORK-SPEC-011
- **work_items**:

## Requirement

DRMCP authoring transactions must conform to PRODUCT-owned authoring standards, namespace semantics, and repository layout rules.

Current DRMCP contracts retain assumptions from the legacy domain-first and single-root model. The contracts do not consistently separate create input, partial update input, and persisted record requirements.

The redesigned contract must support namespace-aware IDs, generated placement, semantic metadata, canonical English headings, and per-kind authoring rules without redefining PRODUCT semantics.

## Evidence

- `PRODUCT-WORK-SPEC-011` defines shared and per-artifact authoring standards as the author-facing contract.
- `spec:product.concepts.authoring_standards.adr_authoring` separates author-supplied values from generated H1 and file path values.
- The ADR standard distinguishes create, partial update, and persisted metadata requirements.
- The ADR standard defines `date` as the decision-validity date, not an automatic modification timestamp.
- `spec:product.concepts.authoring_standards.investigation_authoring` defines conclusion readiness for Investigation records.
- `spec:product.concepts.namespace_model.v2_grammar` defines namespace-aware artifact IDs.
- `DRMCP-REQ-MCP-001` explicitly excludes authoring routing and `suggest_next_record` replacement from its scope.

## Required Outcome

### Contract inventory

Identify every DRMCP tool, schema, parser, validator, and write guard affected by namespace-aware authoring.

The inventory must distinguish:

- PRODUCT-owned semantic inputs;
- DRMCP-owned request and response contracts;
- DRMCP-owned parsing and normalization;
- DRMCP-owned diagnostics;
- implementation-only behavior.

### ID and sequence resolution

Define authoring support for exact IDs and unresolved `new` placeholders.

For non-task sequential artifacts, the sequence allocation scope is:

```text
app namespace + artifact kind + domain namespace
```

Examples:

- `DRMCP-ADR-MCP-new` resolves within DRMCP ADRs in the MCP domain.
- DRMCP requirements or work items do not affect the ADR sequence.
- ADRs in another DRMCP domain use an independent sequence.

TASK allocation follows the parent work item. TASK sequence allocation must not use the non-task scope above.

The contract must define:

- accepted exact-ID and `new` forms by artifact kind;
- sequence lookup inputs;
- collision handling;
- exact-ID sequence-gap behavior;
- concurrency and accept-time staleness guards;
- generated resolved-ID output.

The canonical sequence rule belongs to the PRODUCT namespace model. DRMCP must consume the rule without redefining it.

### Placement and generated values

Define placement resolution from app namespace, artifact kind, and domain namespace.

The contract must distinguish author-supplied and generated values.

Author-supplied values include:

- namespace inputs;
- artifact title;
- author-controlled metadata;
- body sections;
- exact ID or `new` placeholder.

Generated values include:

- resolved sequence when `new` is used;
- public ID;
- H1;
- file name;
- repository path.

For new ADRs, generated placement must use:

```text
<app>/records/adr/<domain>/<APP>-ADR-<DOMAIN>-<NNN>-<slug>.md
```

Discovery must support both the new domain-subdirectory path and existing flat ADR compatibility records. Flat ADR placement is not canonical for new records.

### Metadata state model

Define separate contracts for:

- create input;
- partial update input;
- persisted record state.

The contract must not infer that every recognized metadata field is required in every update request.

The contract must preserve PRODUCT-defined field semantics. In particular:

- semantic dates remain author-controlled;
- editorial updates do not automatically change semantic dates;
- omitted partial-update fields remain unchanged;
- generated values are not accepted as ordinary metadata replacements.

### Body and heading contract

Define body input and named-section update behavior using canonical English headings from per-artifact authoring standards.

The contract must define:

- body-only create input;
- exclusion of generated H1 and metadata from body input;
- canonical section selectors;
- behavior for legacy non-English headings;
- missing, duplicate, or ambiguous heading diagnostics.

DRMCP must consume PRODUCT-defined status-gated narrative rules without redefining them.

For an Investigation with `status: concluded`:

- `## Investigation scope` must exist and contain substantive content;
- `## Findings` must exist and contain substantive content;
- placeholder-only content such as `TBD` must not satisfy the gate;
- validation must report an error for an invalid persisted concluded Investigation;
- authoring proposals that would persist an invalid concluded Investigation must fail.

This requirement is limited to the PRODUCT-defined Investigation conclusion-readiness rule. Broader narrative quality assessment remains excluded.

### Kind support matrix

Define create and update support independently for each artifact kind.

Required create support:

| artifact kind | required create behavior |
|---|---|
| ADR | Accept exact IDs and `<APP>-ADR-<DOMAIN>-new`. Resolve `new` within the app + kind + domain scope. |
| requirement | Accept exact IDs and `<APP>-REQ-<DOMAIN>-new`. Resolve `new` within the app + kind + domain scope. |
| work item | Accept exact IDs and `<APP>-WORK-<DOMAIN>-new`. Resolve `new` within the app + kind + domain scope. |
| task | Accept exact IDs and `<APP>-TASK-<DOMAIN>-<WORK-SEQUENCE>-new`. Resolve `new` within the parent work item. Require an explicit parent work item relation. |
| investigation | Accept exact IDs and `<APP>-INV-<DOMAIN>-new`. Resolve `new` within the app + kind + domain scope. |
| spec | Create specs from PRODUCT-owned path-derived identity rules. Do not use a numeric `new` placeholder. |

SPEC create must support:

- app namespace input;
- parent spec reference;
- target child segment or equivalent placement input;
- spec kind and title;
- required H1-adjacent metadata;
- body sections;
- generated canonical `spec:` ref;
- generated repository path and H1;
- path / ref collision checks;
- parent existence checks;
- spec-kind section validation;
- `contract_class` validation for `Contract` specs;
- rejection of prohibited YAML front matter.

TASK `new` is a required conformance behavior even when an earlier implementation lacks support. The parent work item determines the work sequence and task sequence scope.

Create support for SPEC and investigation is required by the redesigned contract even though the earlier MVP excluded those operations.

Unsupported update operations must be explicit. Tool support must not be inferred from parser support.

### Diagnostics and compatibility

Define DRMCP-owned diagnostic categories and response fields for:

- invalid namespace inputs;
- unsupported artifact kinds or operations;
- unresolved placement;
- sequence collision;
- stale sequence resolution;
- invalid generated-value input;
- metadata state violations;
- canonical heading violations;
- accepted compatibility input.

Compatibility behavior must not make legacy forms canonical for new records.

### Follow-up tracking

Create investigation, work item, and task records for contract redesign and implementation.

The follow-up artifacts must reference the applicable PRODUCT authoring standards and namespace specs.

## Explicitly Excluded Scope

- Defining artifact meaning, responsibility boundaries, or required prose content.
- Defining canonical ID grammar or sequence semantics in DRMCP specs.
- Authoring guidance discovery, composition, or retrieval tools.
- Query, resolve, and cross-namespace validation concerns owned by DRMCP-REQ-MCP-001.
- Bulk migration of existing records.
- Rewriting existing authoring standards.
- UI behavior.

## Boundary

PRODUCT owns:

- artifact semantics;
- authoring standards;
- app and domain namespace semantics;
- canonical ID grammar;
- sequence allocation scope;
- repository layout semantics;
- metadata meaning and author-facing requiredness;
- canonical section headings.

DRMCP owns:

- tool request and response schemas;
- placeholder resolution;
- placement resolution implementation;
- generated file values;
- parser and normalization behavior;
- write guards;
- diagnostics;
- operation support by artifact kind.

A DRMCP contract must cite PRODUCT semantic sources rather than duplicate their normative definitions.
