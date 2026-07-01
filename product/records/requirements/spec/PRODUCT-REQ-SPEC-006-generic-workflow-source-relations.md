# PRODUCT-REQ-SPEC-006: Generic workflow source relations

- **id**: PRODUCT-REQ-SPEC-006
- **status**: accepted
- **date**: 2026-07-01
- **source_refs**:
  - spec:product.design_records.authoring_standards.requirement_authoring
  - spec:product.design_records.authoring_standards.work_item_authoring
  - spec:product.design_records.authoring_standards.task_authoring
  - spec:product.design_records.artifact_model.artifact_responsibility_matrix
  - spec:product.design_records.traceability
- **work_items**:
  - PRODUCT-WORK-SPEC-017

## Requirement

Work Item provenance must use generic source relations instead of a Requirement-only source field.

Task provenance must be reached through the owning `work_item` relation and must not be duplicated in a Task source field.

Requirement records must not persist reciprocal `work_items` lists for relations derivable from downstream source relations.

The relation model must support Work Items created from Requirements, Tasks, investigations, decisions, specifications, or other accepted workflow sources.

The relation model must support Task-to-Work-Item decomposition without a dedicated Hub Work Item kind or parent-child field.

DRMCP must derive and expose reverse relation views without requiring reciprocal relation storage in Requirement records.

## Evidence

- Work Item `source_requirement` restricts provenance to one artifact kind, while Task source metadata duplicates provenance already reachable through `work_item`.
- A Task may create a downstream Work Item with its own Task graph and completion boundary.
- A Work Item may be motivated by multiple accepted records rather than one Requirement alone.
- Requirement `work_items` duplicates a relation that can be derived from downstream source references.
- Reciprocal persisted relations increase synchronization failures and stale graph state.
- Hub behavior can emerge from Task outputs and downstream Work Item source relations.

## Required Outcome

- Replace Work Item `source_requirement` with generic `source_refs`.
- Remove Task `source_requirement` without replacement and rely on Task `work_item` for provenance traversal.
- Permit `source_refs` to contain canonical Design Record public IDs and active semantic refs.
- Permit multiple source refs when multiple upstream artifacts materially motivate the record.
- Preserve Task `work_item` as the owning Work Item membership relation.
- Preserve Work Item `tasks` as the explicit membership list for Tasks owned by that Work Item.
- Remove Requirement `work_items` from the canonical Requirement metadata contract.
- Derive Requirement-to-Work-Item reverse relations from Work Item `source_refs`.
- Allow a downstream Work Item to reference the Task that created or decomposed it.
- Keep each downstream Work Item responsible for its own Task graph and completion condition.
- Avoid dedicated Hub Work Item, parent Work Item, and child Work Item fields unless later evidence proves them necessary.
- Define source relation validation for existence, canonical identity, duplicate refs, and self-reference.
- Define migration behavior for existing `source_requirement` and Requirement `work_items` metadata.
- Keep reverse relation discovery and visualization within DRMCP projections and tools.

## Explicitly Excluded Scope

- Validator model selection or LLM-assisted validation.
- Hub progress calculation.
- New Hub Work Item artifact kinds or role metadata.
- Unlimited Work Item nesting semantics.
- Task responsibility types and single-responsibility rules.
- User-interface design beyond derived reverse relation availability.
- Concrete DRMCP request and response schemas.

## Boundary

PRODUCT owns canonical provenance semantics and persisted relation ownership.

Work Item records persist forward source relations through `source_refs`.

Task records persist no source field; Task provenance is reached through the owning `work_item` relation.

Requirement records do not persist derived reverse Work Item membership.

DRMCP owns relation indexing, reverse traversal, graph validation, and user-visible relation projections.
