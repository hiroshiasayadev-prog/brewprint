# Reference: Artifact refs

- **id**: `spec:product.concepts.traceability.artifact_refs`
- **status**: draft
- **date**: 2026-06-23
- **parent**: `spec:product.concepts.traceability`

## What this is

Defines which semantic ref prefixes and record ID-as-ref forms the traceability MVP resolves, which are reserved, and which are deferred.

## Purpose

This file defines the ref types that the traceability MVP resolves as a canonical reference resolution foundation.

The MVP does not make all artifact layers or realization relations first-class targets of a semantic trace graph. Active semantic refs are limited to `spec:` only; non-spec record artifacts are resolved by full public ID-as-ref.

## Active prefixes

The MVP active semantic ref prefixes are limited to:

```yaml
active_prefixes:
  - spec
```

### `spec:`

`spec:` represents the canonical document-level or section-level identity for a design spec.

```text
spec:product.concepts.traceability
spec:product.concepts.traceability.semantic_ref
spec:product.concepts.traceability.resolve_and_validation
```

New and migrated specs use path-derived `spec:` refs according to `spec:product.concepts.spec_format.spec_id_as_ref`.

Legacy `SPEC-*` public IDs may remain resolvable as compatibility inputs when an indexed legacy spec record carries one. They are not the canonical identity for new specs.

## Reserved prefixes

```yaml
reserved_prefixes:
  - yaml
```

### `yaml:`

`yaml:` is reserved for the future semantic endpoint for brewprint DSL definitions representing the target system or design model. It is not activated in the MVP and its resolve behavior is not fixed.

Trace metadata such as spec metadata and investigation metadata is not a target for `yaml:`.

## Deferred prefixes

The following are prefix candidates whose operational contracts as semantic trace endpoints are deferred to a future decision. Their existence as artifact layers is not denied.

```yaml
deferred_prefixes:
  - internal-design
  - coverage
```

### `internal-design:`

The internal-design artifact layer continues to exist. However, the MVP does not index, resolve, or validate internal design documents via semantic refs, and does not handle their realization relation with `spec:`.

### `coverage:`

The MVP does not adopt external coverage artifacts. Therefore `coverage:` is not an active or reserved mapping-set identity. The name will be reconsidered if external relation artifacts are introduced.

## ID-as-ref

A record ID-as-ref uses the complete public ID of the target record. Bare kind forms such as `REQ-...` or `WORK-...` are internal grammar fragments and are not canonical external references.

Public ID grammar is owned by the namespace model:

- Existing issued records retain their legacy public IDs under `spec:product.concepts.namespace_model.legacy_id_compatibility` and `spec:product.concepts.namespace_model.existing_artifacts`.
- New ADR, investigation, requirement, work item, and task records use `spec:product.concepts.namespace_model.artifact_id_grammar`.
- New and migrated specs use path-derived `spec:` refs rather than a new `SPEC-*` ID.

| record type | canonical new form | legacy compatibility example |
|---|---|---|
| decision | `<APP>-ADR-<DOMAIN>-<SEQUENCE>` | `V01-ADR-088` |
| spec | path-derived `spec:` ref | `V01-SPEC-design-records-mcp-schema` |
| investigation | `<APP>-INV-<DOMAIN>-<SEQUENCE>` | `V01-INV-DOCS-003` |
| requirement | `<APP>-REQ-<DOMAIN>-<SEQUENCE>` | `V01-REQ-MCP-003` |
| work item | `<APP>-WORK-<DOMAIN>-<SEQUENCE>` | `V01-WORK-MCP-003` |
| task | `<APP>-TASK-<DOMAIN>-<WORK-SEQUENCE>-<TASK-SEQUENCE>` | `V01-TASK-MCP-003-01` |

Examples of current app-aware artifact ID-as-refs include `DRMCP-REQ-MCP-002`, `PRODUCT-WORK-SPEC-011`, and `PRODUCT-TASK-SPEC-011-08`.

## Design record ID-as-ref boundary

ADR and investigation public IDs are direct resolver and validation targets.

Specs use canonical `spec:` refs. A legacy indexed `SPEC-*` public ID may remain a compatibility resolver target, but new relations must use the canonical path-derived `spec:` ref.

Investigation `source_refs` and recorded `follow_up_results` may use supported record ID-as-refs or active `spec:` refs according to the metadata boundary below.

## Workflow ID-as-ref boundary

Requirement, work item, and task relations use complete public ID-as-refs. Physical paths and semantic prefixes such as `requirement:` are not relation identities.

The workflow kinds additionally permitted in investigation `source_refs`, recorded `follow_up_results`, and `follow_up_candidates` are limited to requirement and work item public IDs.

Task public IDs are supported as workflow inter-relations and direct resolver inputs, but are unsupported in investigation canonical reference fields.

## `COV-*`

`COV-*` is not a canonical reference form in the MVP. It will be reconsidered if external coverage artifacts and individually identified mappings are introduced.

## Relation endpoint boundary

The MVP does not define endpoints for semantic realization relations. `spec:` to `internal-design:` mapping, `maps_to`, `covers`, and coverage mapping endpoint constraints are future scope.

## Scope-out prefixes

`fixture:` is neither an active nor a reserved prefix in the MVP. Fixture and golden assets are verification assets and are not included in the project-level canonical reference foundation.

`requirement:` / `work-item:` / `task:` prefixes are also not adopted. Workflow artifacts are resolved through their complete public IDs.

## Sources

- `spec:product.concepts.namespace_model.legacy_id_compatibility`: V01-* legacy ID families and retention policy.
- `spec:product.concepts.namespace_model.artifact_id_grammar`: canonical artifact ID grammar.
- `spec:product.concepts.spec_format.spec_id_as_ref`: canonical spec identity.
- V01-ADR-083 §8: separation of trace identity from physical path.
- V01-ADR-087: Design Records MCP resolve responsibility and investigation canonical reference rule.
- V01-ADR-088: canonical reference resolution foundation boundary.
- V01-ADR-091 and V01-ADR-092: workflow ID-as-ref relation boundary.
