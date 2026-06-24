# Overview: Traceability

- **id**: `spec:product.concepts.traceability`
- **status**: draft
- **date**: 2026-06-23
- **parent**: `root`

## What this is

Entry point for the brewprint traceability spec set. Defines the canonical reference resolution foundation MVP scope, out-of-scope boundary, and the set of child specs that own the detail contracts.

## Current contract

MVP covers `spec:` semantic ref canonical resolution, record ID-as-ref resolution, investigation reference validation, and workflow artifact declared relation integrity validation. The boundary is defined under `## MVP scope` and `## Out of MVP scope`. Child specs in `## Topics` own the detail rules.

## Purpose

This spec set defines the semantic trace MVP for brewprint docs.

Per V01-ADR-088 and V01-ADR-092, the goal of the MVP is not to pre-build a realization graph across artifacts, but to provide a **canonical reference resolution foundation** that lets design records, investigations, and workflow artifacts resolve and validate canonical references and declared relations without depending on physical paths.

## MVP scope

What the MVP covers:

- Declaration, stability, and document/section resolution of `spec:` semantic refs
- Resolution of complete public record ID-as-refs. Existing issued records retain legacy public IDs; new records use the namespace-aware v2 grammar
- Canonical reference resolution and unresolved errors for investigation `source_refs` and recorded `follow_up_results`. Additional workflow ID-as-ref support is limited to requirement and work item public IDs
- Canonical form checks for artifact references in `follow_up_candidates`. Additional workflow ID-as-ref support is limited to requirement and work item public IDs; unresolved candidates are not required to exist
- Existence check and bidirectional consistency check for declared ID-as-ref relations between workflow artifacts
- Boundary that excludes physical paths as canonical references

The active semantic ref prefix for the MVP is limited to:

```yaml
active_prefixes:
  - spec
```

`yaml:` is reserved for brewprint DSL YAML but is not activated.

## Out of MVP scope

The MVP does not treat the following as operational mechanisms:

- `internal-design:` semantic endpoint
- `coverage:` semantic endpoint
- `COV-*` mapping identity
- External relation / assurance artifacts and their placement
- `maps_to` / `covers` / `validates` relations
- `spec:` → `internal-design:` realization mapping
- YAML endpoint, fixture/golden traceability, coverage/evidence matrix
- Workflow artifact orphan diagnostics, task-status-derived progress projection, workflow-dedicated traversal queries, task dependency cycle / execution order projection
- Task public ID-as-ref support in investigation metadata

The `docs/internal-design/` artifact layer itself continues to exist. What is out of MVP scope is the contract for resolving and validating that layer as a semantic trace endpoint. External relation / assurance artifacts will be introduced with placement and responsibility decisions when the need is established; no directory is reserved in the MVP layout.

## Terms

### semantic ref

An identifier that stably references the concept represented by a brewprint docs artifact, independent of physical path, Markdown heading, or directory layout.

MVP example:

```text
spec:trace
spec:trace.semantic-ref
spec:trace.resolve-and-validation
```

### record ID-as-ref

A stable ID pointing to a record artifact handled by Design Records MCP.

```text
V01-ADR-088
DRMCP-INV-MCP-001
DRMCP-REQ-MCP-002
PRODUCT-WORK-SPEC-011
PRODUCT-TASK-SPEC-011-08
```

A record ID-as-ref always uses the complete public ID. Bare forms such as `REQ-...` are not external canonical references. New and migrated specs use path-derived `spec:` refs; legacy `SPEC-*` public IDs are compatibility-only.

### brewprint DSL YAML

The primary implementation source for the target system / design model expressed in brewprint DSL. Activating the `yaml:` semantic endpoint is a future decision.

### trace metadata

Metadata for declaring and referencing canonical references. Separate from brewprint DSL YAML responsibilities. The MVP focuses on canonical reference rules in spec front matter and investigation metadata.

## Topics

| title | kind | ref | summary |
|---|---|---|---|
| Semantic ref | Reference | `spec:product.concepts.traceability.semantic_ref` | `spec:` grammar, stability, document/section refs. |
| Artifact refs | Reference | `spec:product.concepts.traceability.artifact_refs` | Active/reserved/deferred semantic ref prefixes and ID-as-ref scope. |
| Metadata schema | Reference | `spec:product.concepts.traceability.metadata_schema` | Trace metadata fields, investigation reference boundary, workflow artifact relation boundary. |
| Coverage mapping | Reference | `spec:product.concepts.traceability.coverage_mapping` | Realization mapping and external coverage boundary; reintroduction triggers. |
| Resolve and validation | Reference | `spec:product.concepts.traceability.resolve_and_validation` | Canonical resolve and validation boundary; lookup sources. |
| Out of scope | Reference | `spec:product.concepts.traceability.out_of_scope` | MVP out-of-scope items and future extension triggers. |

## Source of truth boundary

The traceability spec owns the canonical reference model for docs artifacts. Per V01-ADR-087, Design Records MCP is the tool boundary that implements resolve and validation, but it does not own the traceability semantic model itself.

The responsibility boundary across the full artifact system is owned by `spec:product.concepts.project_artifact_model`.

## Sources

- V01-ADR-081: requirements layer and semantic traceability
- V01-ADR-083: project artifact boundary and YAML as primary implementation source
- V01-ADR-084: semantic trace MVP scope and artifact boundary
- V01-ADR-087: Design Records MCP investigation support and semantic ref resolve
- V01-ADR-088: Reduce semantic trace MVP to a canonical reference resolution foundation
- V01-ADR-091: Work item / task responsibility separation and legacy milestone migration
- V01-ADR-092: Design Records MCP workflow artifact record and relation boundary
