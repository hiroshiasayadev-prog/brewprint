# Reference: Artifact refs

- **id**: `spec:product.concepts.traceability.artifact_refs`
- **status**: draft
- **date**: 2026-06-22
- **parent**: `spec:product.concepts.traceability`

## What this is

Defines which semantic ref prefixes and record ID-as-ref forms the traceability MVP resolves, which are reserved, and which are deferred.

## Purpose

This file defines the ref types that the traceability MVP resolves as a canonical reference resolution foundation.

The MVP does not make all artifact layers or realization relations first-class targets of a semantic trace graph. Active semantic refs are limited to `spec:` only; record artifacts are resolved by ID-as-ref.

## Active prefixes

The MVP active semantic ref prefixes are limited to:

```yaml
active_prefixes:
  - spec
```

### `spec:`

`spec:` represents a document-level or section-level semantic ref for a design spec.

```text
spec:trace
spec:trace.semantic-ref
spec:trace.resolve-and-validation
```

`spec:` section ref resolution uses the `sections` mapping in spec front matter. Physical heading anchors are not treated as canonical identity.

## Reserved prefixes

```yaml
reserved_prefixes:
  - yaml
```

### `yaml:`

`yaml:` is reserved for the future semantic endpoint for brewprint DSL YAML representing the target system / design model. It is not activated in the MVP and its resolve behavior is not fixed.

Trace metadata such as spec front matter and investigation metadata is not a target for `yaml:`.

## Deferred prefixes

The following are prefix candidates whose operational contracts as semantic trace endpoints are deferred to a future decision. Their existence as artifact layers is not denied.

```yaml
deferred_prefixes:
  - internal-design
  - coverage
```

### `internal-design:`

`docs/internal-design/` continues to exist as an implementation-facing documentation layer. However, the MVP does not index, resolve, or validate internal design documents via semantic refs, and does not handle the realization relation with `spec:`.

### `coverage:`

The MVP does not adopt external coverage artifacts. Therefore `coverage:` is not an active/reserved contract as a mapping set identity. The name will be reconsidered when reintroducing external artifacts.

## ID-as-ref

The ID-as-refs treated as canonical references by the MVP are at minimum the record artifact IDs of Design Records MCP.

| record type | ID format |
|---|---|
| decision | `ADR-NNN` |
| spec record | `SPEC-<slug>` |
| investigation | `INV-<DOMAIN>-NNN` |
| requirement | `REQ-<DOMAIN>-NNN` |
| work item | `WORK-<DOMAIN>-NNN` |
| task | `TASK-<DOMAIN>-<WORK-SEQUENCE>-<TASK-SEQUENCE>` |

### `ADR-*` / `SPEC-*` / `INV-*`

`ADR-*` / `SPEC-*` / `INV-*` are record artifact IDs targeted for index/query/validation by Design Records MCP. Investigation `source_refs` and recorded `follow_up_results` may use these ID-as-refs or active `spec:` refs as canonical references.

### `REQ-*` / `WORK-*` / `TASK-*`

`REQ-*` / `WORK-*` / `TASK-*` are workflow artifact ID-as-refs targeted for index/query/validation by Design Records MCP. Canonical relations between workflow artifacts use these ID-as-refs; physical paths and semantic prefixes are not used as relation identities.

The workflow artifact ID-as-refs that investigation metadata `source_refs` / recorded `follow_up_results` / `follow_up_candidates` may additionally use are limited to `REQ-*` / `WORK-*`. `TASK-*` is supported as a workflow artifact inter-relation and direct resolver input, but is unsupported as an investigation metadata canonical reference.

### `COV-*`

`COV-*` is not a canonical reference form in the MVP. It will be reconsidered if the need for external coverage artifacts and individual mappings arises.

## Relation endpoint boundary

The MVP does not define endpoints for semantic realization relations. `spec:` → `internal-design:` mapping, `maps_to`, `covers`, and coverage mapping endpoint constraints are future scope.

## Scope-out prefixes

`fixture:` is neither an active nor a reserved prefix in the MVP. Fixture/golden assets are verification assets for processors/renderers/validators and are not included in the project-level canonical reference foundation.

`requirement:` / `work-item:` / `task:` prefixes are also not adopted in the MVP. Workflow artifacts are resolved by `REQ-*` / `WORK-*` / `TASK-*` ID-as-refs.

## Sources

- V01-ADR-081 §5: requirement IDs do not combine with ADR numbers
- V01-ADR-083 §8: principle of separating trace identity from physical path
- V01-ADR-087: Design Records MCP resolve responsibility and investigation canonical reference rule
- V01-ADR-088: scope reduced to canonical reference resolution foundation by deferring realization endpoints and external coverage artifacts to MVP-out
- V01-ADR-091: ID-as-ref boundary for workflow artifact inter-relations
- V01-ADR-092: workflow artifact record / resolve / relation validation boundary
