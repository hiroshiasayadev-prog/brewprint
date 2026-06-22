# Reference: Coverage mapping

- **id**: `spec:product.concepts.traceability.coverage_mapping`
- **status**: draft
- **date**: 2026-06-22
- **parent**: `spec:product.concepts.traceability`

## What this is

Defines the boundary that keeps semantic realization mapping and external coverage artifacts out of the traceability MVP, and the triggers for reintroducing them as future requirements.

## MVP realization mapping boundary

The semantic trace MVP does not treat semantic realization mapping as an operational contract.

The MVP is limited to canonical resolution of `spec:` semantic refs and record artifact ID-as-refs, and to investigation reference validation. It does not build implementation realization graphs, coverage matrices, or evidence matrices across artifacts.

Accordingly, the MVP does not define or require:

- External mapping artifacts holding semantic realization relations, and their placement
- `coverage:` semantic ref prefix
- `COV-*` individual mapping IDs
- Coverage mapping YAML schema
- Authoring / resolve / validation of `maps_to` / `covers` relations
- Mapping direction or endpoint constraints for `spec:` → `internal-design:`

## Internal design boundary

`docs/internal-design/` continues to exist as an implementation-facing design artifact layer, but `internal-design:` is not an active MVP endpoint.

The existence of internal design documents is not a reason to require semantic realization relations in the MVP. The MVP also does not define a schema requiring internal design metadata to declare source `spec:` relations, or a reverse-lookup graph from spec to internal design.

## Deferred mechanisms

The following are mechanisms whose necessity has not been denied but whose decisions are deferred until concrete requirements are confirmed.

| mechanism | MVP treatment |
|---|---|
| `internal-design:` semantic endpoint | deferred |
| `coverage:` semantic endpoint | deferred |
| `COV-*` mapping identity | deferred |
| external relation / assurance artifacts and their placement | deferred |
| `maps_to` semantic realization relation | deferred |
| `covers` semantic realization relation | deferred |
| `yaml:` semantic endpoint | reserved / inactive |

Moving `maps_to` from external artifacts to endpoint metadata is also not done in the MVP. When a relation is operationally introduced, its identity, direction, owner, schema, and validation will be decided together.

## Reintroduction triggers

If any of the following are confirmed as requirements, the necessity of endpoints, relations, and external artifacts will be reconsidered via a subsequent ADR / requirement / work item.

- Machine-driven navigation/impact analysis from spec is needed for multiple internal design documents
- An investigation / work item / MCP query needs to resolve and validate internal design artifacts as canonical references
- Activation of `yaml:` makes a realization chain or cross-layer validation between spec / internal design / YAML necessary
- Central management of gap / completeness / approved relation sets becomes necessary
- evidence / sign-off / audit snapshots / release baselines need to be bound to relations
- Relation entries themselves need stable identity / approval / lifecycle / history

## Validation boundary

What the MVP targets for validation are the rules governing the canonical reference foundation.

- Grammar / uniqueness / resolve of `spec:` semantic refs
- Resolution of record ID-as-refs
- Canonicality and resolution of investigation `source_refs` / recorded `follow_up_results`
- Canonical form of artifact references in investigation `follow_up_candidates`
- Physical paths not being canonical references

Coverage mapping schema and realization endpoint validation are outside MVP scope.

## Future artifact placement

If external relation artifacts are introduced in the future, whether to call them `coverage`, whether to separate semantic mapping from assurance matrices, and which directory to place them in will all be subject to re-decision. The MVP does not reserve a directory for external artifacts.

## Sources

V01-ADR-083, V01-ADR-084, V01-ADR-088; V01-INV-DOCS-002; V01-INV-DOCS-003
