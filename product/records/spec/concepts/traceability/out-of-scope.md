# Reference: Traceability out of scope

- **id**: `spec:product.concepts.traceability.out_of_scope`
- **status**: draft
- **date**: 2026-06-22
- **parent**: `spec:product.concepts.traceability`

## What this is

Catalogs what is explicitly out of scope for the traceability MVP and the triggers for reintroducing each item as a future requirement.

## MVP out-of-scope boundary

The semantic trace MVP is limited to the canonical reference resolution foundation. Items outside scope are not deemed unnecessary — they will be decided via subsequent ADRs / requirements / work items when concrete requirements are confirmed.

## Internal-design semantic endpoint and realization relation

The `docs/internal-design/` artifact layer continues to exist, but the MVP does not handle:

- Activation of the `internal-design:` semantic ref
- Canonical ref resolve/validation of internal design documents
- Source `spec:` relation declarations via internal design metadata
- `spec:` → `internal-design:` realization mapping
- Reverse graph / impact queries from spec to internal design

Reintroduction triggers:

- Machine-driven navigation/impact analysis from spec becomes necessary for multiple internal design documents
- An investigation / work item / MCP query needs canonical resolve of internal design artifacts
- Cross-layer realization chain involving YAML trace becomes necessary

## External coverage artifact and relation vocabulary

The MVP does not handle external relation / assurance artifacts, their placement, `coverage:` refs, `COV-*` mapping identities, or coverage mapping YAML schemas.

Adopting `maps_to` / `covers` / `validates` as an operational semantic realization relation vocabulary is also deferred. Validation itself remains in the MVP as tool behavior / metadata contract, not as a relation.

Reintroduction triggers:

- Central management of gap / completeness / approved relation sets becomes necessary
- evidence / sign-off / audit snapshots / release baselines need to be bound to relations
- Relation entries themselves need stable identity / approval / lifecycle / history

If external artifacts are introduced in the future, whether to call them `coverage`, whether to separate semantic mapping from assurance matrices, and which directory to place them in will all be subject to re-decision. The MVP does not reserve a directory for external artifacts.

## Brewprint DSL YAML entity-level refs

The MVP does not activate the `yaml:` prefix.

Out of scope:

- File-level / entity-level semantic refs for brewprint DSL YAML
- Semantic refs per node / edge / view / model / task / asset unit
- Anchor and logical unit resolver rules within YAML
- Realization chains between spec / internal design / YAML

Activating `yaml:` will be decided when self-hosting / UC-002 reconstruction, or concrete requirements for YAML entity refs / resolver rules, are established.

## Fixture and golden traceability

The MVP does not define the `fixture:` prefix and does not include fixture/golden in the project-level canonical reference foundation.

Out of scope:

- Fixture semantic refs
- Project-level integration of fixture-local coverage
- Correspondence between golden outputs and spec semantic refs
- Render expected comparison semantics
- Test harness schema / golden update workflow

## Workflow semantic prefixes and derived operations

Requirement / work item / task become public record / resolver / declared relation validation targets of Design Records MCP via `REQ-*` / `WORK-*` / `TASK-*` ID-as-refs. However, the MVP does not define `requirement:` / `work-item:` / `task:` semantic prefixes.

The following are outside MVP scope for workflow artifact support:

- Orphan diagnostics for requirement / work item / task
- Projection deriving work item progress from task status
- Workflow-dedicated traversal / tree / graph query tools
- Task dependency cycle detection / execution order projection
- `TASK-*` canonical references in investigation metadata

These are not prerequisites for declared relation integrity validation or direct ID-as-ref resolution, and will be decided when concrete operational requirements are confirmed.

## Full MCP writer tools

The MVP does not define request/response schemas for MCP writer tools.

Future candidates:

- Create requirement / work item
- Register `spec:` semantic refs
- Update section mappings
- Update investigation reference metadata

If writer tools are introduced, dry-run diff, user confirmation, conflict handling, format preservation, and write permission boundaries will be defined in a separate spec / ADR.

## Future extensions

| extension | trigger |
|---|---|
| `internal-design:` activation | Internal design canonical navigation/validation becomes a concrete requirement |
| realization relation | Machine-readable relation between spec and implementation-facing artifacts becomes necessary |
| external relation artifact | Central management of gap / evidence / sign-off / lifecycle becomes necessary |
| `yaml:` activation | YAML entity refs / cross-layer trace becomes necessary |
| fixture-level traceability | Long-term management of golden fixture and docs/spec correspondence becomes necessary |
| workflow semantic prefix | Section-level addressing requirement emerges that `REQ-*` / `WORK-*` / `TASK-*` ID-as-refs cannot satisfy |
| workflow orphan / progress / traversal capability | Concrete requirements for disconnected artifact diagnostics, status aggregation views, or dedicated graph traversal emerge |
| MCP resolve contract refinement | Additional refinement of the adopted canonical reference / workflow relation validation contract becomes necessary |
| MCP writer tools | Generating or updating canonical metadata via tools becomes necessary |

## Follow-up artifact placement

When future extensions are needed, the owners of requirements, decisions, progress, and contracts follow the existing artifact boundary.

- Requirements / gaps / requests: `product/records/requirements/`
- Decisions: `product/records/adr/`
- Cross-cutting progress: `product/records/work-items/`
- Concrete work: `product/records/tasks/`
- Canonical reference / tool-independent trace contract: `product/records/spec/concepts/traceability/`
- Internal design artifacts *(planned)*: no canonical path yet — placement will be decided when `internal-design:` activation is triggered
- Implementation follow-up notes *(planned)*: `docs/impl/` is the current convention; canonical `<namespace>/records/` path not yet assigned

## Sources

V01-ADR-082, V01-ADR-083, V01-ADR-084, V01-ADR-087, V01-ADR-088, V01-ADR-091, V01-ADR-092; V01-INV-DOCS-002; V01-INV-DOCS-003
