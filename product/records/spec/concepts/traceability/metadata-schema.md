# Reference: Trace metadata schema

- **id**: `spec:product.concepts.traceability.metadata_schema`
- **status**: draft
- **date**: 2026-06-23
- **parent**: `spec:product.concepts.traceability`

## What this is

Defines the trace metadata fields for spec front matter and investigation reference metadata, including the canonical reference boundary and workflow artifact relation validation boundary for the traceability MVP.

## Trace metadata

Trace metadata is docs metadata for declaring and referencing canonical references. It is not brewprint DSL YAML, and it is not a target for `yaml:` semantic refs.

The metadata defined by this spec in the MVP is limited to:

- `semantic_refs` / `sections` in spec front matter
- Permitted forms and validation boundary for canonical references in investigation metadata
- Validation boundary for declared relations among complete requirement, work item, and task public IDs

Internal-design relation metadata, coverage mapping YAML, and relation entry schemas are not defined in the MVP.

## Common fields

MVP spec trace metadata may optionally carry the following fields:

| field | required | type | meaning |
|---|---:|---|---|
| `semantic_refs` | no | list<string> | document-level `spec:` semantic ref list |
| `sections` | no | map<string,string> | mapping between section-level `spec:` semantic refs and Markdown heading text |

`scope` / `status` / `last_updated` / `summary` / `depends_on` are owned by doc-policy.

## `semantic_refs`

`semantic_refs` is a list of document-level `spec:` refs owned by the spec document as a whole.

```yaml
semantic_refs:
  - spec:trace
```

Rules:

- Values must conform to the active `spec:` semantic ref grammar
- Duplicate values are invalid
- Document-level refs must be unique within the same project
- Physical paths must not be used as values

## `sections`

`sections` is a map representing the correspondence between section-level `spec:` semantic refs and Markdown heading text.

```yaml
sections:
  spec:trace.semantic-ref.definition: Semantic ref definition
  spec:trace.resolve: Resolve
```

Rules:

- Keys must conform to the active `spec:` semantic ref grammar
- Keys must be unique within the same project
- Values must match the actual heading text without the `#` heading marker
- When a heading is renamed or a section is moved, the semantic ref key is preserved; only the value is updated if needed

Writing `{#anchor}` directly on Markdown headings as canonical identity is not used.

## Internal design metadata boundary

`docs/internal-design/` continues to exist as an implementation-facing artifact layer, but the semantic trace MVP does not require internal design front matter to carry `internal-design:` refs or source `spec:` relation declarations.

If a concrete consumer arises that needs to resolve internal design as canonical references, the prefix, metadata fields, resolver rules, and validation contract will be decided in a subsequent spec.

## Coverage metadata boundary

The MVP does not define external coverage artifacts, `coverage:` refs, `COV-*` IDs, or coverage mapping YAML schemas.

If the need arises to hold gap / completeness / evidence / sign-off / audit / approved relation sets in external artifacts, their names and schemas will be decided in a subsequent spec.

## Workflow artifact metadata boundary

> **Drift guard:** PRODUCT owns concrete bidirectional integrity rule statements. DRMCP owns parser, response, and diagnostic details — do not add those here.

Requirement / work item / task are workflow record artifacts handled by Design Records MCP. The complete parser/response schema and diagnostic categories for workflow artifacts are owned by the Design Records MCP spec; this spec defines only the canonical relation boundary.

Workflow artifact inter-relations are declared via the following ID-as-ref fields:

| source artifact | field | canonical target |
|---|---|---|
| requirement | `work_items` | Complete work item public ID |
| work item | `source_requirement` | Complete requirement public ID |
| work item | `tasks` | Complete task public ID |
| task | `work_item` | Complete work item public ID |
| task | `source_requirement` | Complete requirement public ID |
| task | `depends_on` | Complete task public ID |

Workflow relations are read only from complete public ID-as-refs; parent relations are not inferred from physical paths or ID string structure. Bare kind forms are not accepted as external relation identities. `req:` / `work:` / `task:` semantic prefixes are also not introduced.

Existing issued records retain their legacy public IDs. New workflow artifacts use `spec:product.concepts.namespace_model.artifact_id_grammar`.

The MVP handles existence checks for the above fields and the following consistency checks for declared relations:

- Mutual relation between requirement and work item: `requirement.work_items` and `work_item.source_requirement`
- Mutual relation between work item and task: `work_item.tasks` and `task.work_item`
- Match between task's source requirement and parent work item's source requirement

Orphan diagnostics for disconnected artifacts, task-status-derived progress projection, workflow traversal queries, and task dependency cycle / execution order projection are not handled.

## Investigation reference metadata

> **Drift guard:** PRODUCT owns concrete bidirectional integrity rule statements. DRMCP owns parser, response, and diagnostic details — do not add those here.

The field composition, required/optional breakdown, status, lifecycle, and authoring format of investigation metadata are handled as authoring guidance by guide ID `investigation-authoring`.

What this spec owns is the canonical reference rule based on V01-ADR-087 / V01-ADR-088:

- `source_refs` must use record ID-as-refs or active `spec:` semantic refs; recorded values must be resolvable
- `follow_up_results`, when recorded, must use record ID-as-refs or active `spec:` semantic refs; recorded values must be resolvable
- When artifact references are recorded in `follow_up_candidates`, canonical form must be used. Since they may point to not-yet-created artifact candidates, unresolved is not an error — it is surfaced as an `info` diagnostic
- The workflow artifact ID-as-refs additionally usable by investigation metadata are limited to complete requirement and work item public IDs
- Complete task public IDs are supported as direct resolver inputs and workflow inter-relations, but are not included in investigation metadata canonical references. Appearance in `source_refs` / `follow_up_results` is an unsupported error; appearance in `follow_up_candidates` is unsupported info
- Physical paths are not canonical references. Appearance in `source_refs` / `follow_up_results` is an error diagnostic; appearance in `follow_up_candidates` is an `info` diagnostic indicating a noncanonical candidate
- Resolve/validation rules for `trigger` / `related_*` are defined in a subsequent contract

## Validation responsibility

> **Drift guard:** PRODUCT owns invalid conditions and scope-boundary clauses. DRMCP owns diagnostic category names, JSON shape, and tool response vocabulary — do not add those here.

MVP validation is canonical reference validation and integrity validation of ID-as-ref relations explicitly declared by workflow artifacts.

- Grammar and uniqueness of `semantic_refs` / `sections`
- Match between `sections` values and actual headings
- Canonicality and resolution of investigation `source_refs` / recorded `follow_up_results`
- Canonical form of artifact references in investigation `follow_up_candidates` and `info` diagnostic for unresolved candidates
- Requirement and work item public IDs permitted, and task public IDs unsupported, in investigation metadata
- Noncanonical diagnostic for physical paths: error for `source_refs` / `follow_up_results`, `info` for `follow_up_candidates`
- Canonical target kind / resolution / declared bidirectional consistency of workflow relation fields

Coverage mapping endpoints, semantic realization relations, `COV-*`, `internal-design:` resolution, workflow orphan diagnostics, progress projection, and traversal queries are outside MVP validation scope.

## Out of scope

- Brewprint DSL YAML schema / entity-level semantic refs
- Internal-design semantic ref / relation metadata schema
- Coverage mapping schema
- Workflow artifact orphan diagnostics / progress projection / traversal query / dependency cycle detection
- MCP writer tool request / response
- Fixture-level traceability
