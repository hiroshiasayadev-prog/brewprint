# Overview: Traceability

- **id**: `spec:product.design_records.traceability`
- **status**: draft
- **date**: 2026-06-24
- **parent**: `spec:product.design_records`

## What this is

Entry point for the Design Records traceability spec set. It defines the active canonical reference foundation and separates PRODUCT-owned reference semantics from DRMCP-owned implementation behavior.

## Current contract

The active traceability contract covers:

| area | active contract |
|---|---|
| Spec refs | New and migrated specs use path-derived document-level `spec:` refs recorded in the H1-adjacent `id`. |
| Record ID-as-ref | ADR, investigation, requirement, work item, and task records use complete public IDs. Legacy issued IDs remain compatibility inputs through Brewprint compatibility specs. |
| Investigation references | `source_refs` and recorded `follow_up_results` must use canonical, resolvable refs. `follow_up_candidates` must use canonical form when an artifact ref is written, but not-yet-created candidates may be unresolved. |
| Workflow relations | Requirement, work item, and task relations use declared complete public IDs. Relation consistency is validated from metadata fields, not inferred from paths or ID string shape. |
| Physical paths | Physical paths are repository locations, not canonical relation or reference values. |

PRODUCT owns the canonical reference classes, lookup-source boundary, and invalid semantic conditions. DRMCP owns request shape, response shape, diagnostic names, parser behavior, persistence, indexing, UI, tool APIs, and writer behavior.

## Non-goals

- Activating `yaml:`, `internal-design:`, `coverage:`, `fixture:`, or workflow semantic prefixes.
- Defining `maps_to`, `covers`, or `validates` as trace relations.
- Defining external relation artifacts, assurance matrices, coverage schemas, mapping groups, or `COV-*` IDs.
- Defining orphan diagnostics, progress projection, workflow traversal, dependency-cycle checks, or execution-order projection.
- Defining MCP writer request or response behavior.
- Defining canonical section refs before a visible-table contract exists.

The `docs/internal-design/` artifact layer may continue to exist. No internal-design canonical trace endpoint is currently defined.

## Traceability model

| class | current rule | owner |
|---|---|---|
| `spec:` ref | Path-derived canonical document ref for new and migrated specs. | `spec:product.design_records.spec_format.spec_id_as_ref`. |
| Spec section ref | Not an active contract. A later visible-table contract is required before section refs become canonical. | Future PRODUCT decision. |
| Record ID-as-ref | Complete public record IDs only. Bare `REQ-*`, `WORK-*`, or similar grammar fragments are not external refs. | `spec:product.design_records.namespace_model.artifact_id_grammar` and Brewprint compatibility specs. |
| Investigation canonical refs | `source_refs`, recorded `follow_up_results`, and artifact refs in `follow_up_candidates` follow the investigation boundary. | This traceability area plus investigation authoring standards. |
| Workflow declared relations | `requirement.work_items`, `work_item.source_requirement`, `work_item.tasks`, `task.work_item`, `task.source_requirement`, and `task.depends_on`. | This traceability area plus workflow authoring standards. |

## Topics

| title | kind | ref | summary |
|---|---|---|---|
| Semantic ref | Reference | `spec:product.design_records.traceability.semantic_ref` | Current `spec:` ref class and dispositions for obsolete semantic-ref assumptions. |
| Artifact refs | Reference | `spec:product.design_records.traceability.artifact_refs` | Active canonical spec refs, record ID-as-ref boundaries, investigation refs, and workflow relation identity. |
| Metadata schema | Reference | `spec:product.design_records.traceability.metadata_schema` | Visible spec metadata pointers, investigation reference boundary, and workflow relation metadata boundary. |
| Resolve and validation | Reference | `spec:product.design_records.traceability.resolve_and_validation` | Canonical lookup sources and PRODUCT-owned invalid conditions without DRMCP response vocabulary. |

## Sources

| source | use |
|---|---|
| `spec:product.design_records.spec_format.spec_id_as_ref` | Path-derived canonical spec identity and no current front-matter section-ref contract. |
| `spec:product.brewprint.compatibility` | Legacy issued-ID compatibility pointer. |
| V01-ADR-087 | Investigation canonical reference and resolver responsibility evidence. |
| V01-ADR-088 | Canonical reference resolution foundation and removal of `internal-design:` / `coverage:` from MVP active scope. |
| V01-ADR-092 | Workflow record ID-as-ref and declared-relation integrity boundary. |
| V01-INV-DOCS-002 | Coverage artifact evidence and reintroduction trigger evidence. |
| V01-INV-DOCS-003 | Internal-design endpoint and relation deferral evidence. |
| PRODUCT-ADR-SPEC-001 | PRODUCT semantic ownership and temporary BPDSL staging boundary. |
