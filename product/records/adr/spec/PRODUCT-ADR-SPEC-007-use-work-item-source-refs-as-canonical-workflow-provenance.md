# PRODUCT-ADR-SPEC-007: Use Work Item source refs as canonical workflow provenance

- **status**: accepted
- **date**: 2026-07-01
- **depends_on**:
  - PRODUCT-ADR-SPEC-001
- **supersedes**: []
- **migrated_to_spec**: null

## Context

The current workflow relation model persists Requirement-only provenance in Work Items and Tasks.
It also persists reciprocal Requirement `work_items` metadata.

The model cannot represent Work Items created from Tasks, investigations, decisions, Specifications, or several material sources.
Task source metadata also duplicates provenance already reachable through the owning `work_item` relation.
Reciprocal forward and reverse fields create synchronization failures and stale graph state.

`V01-ADR-091` and `V01-ADR-092` record the earlier workflow relation model.
Both ADRs also contain unrelated decisions that remain useful.
This decision replaces only their legacy source-relation and reciprocal-validation sections through current Specifications.
It does not supersede either historical ADR as a whole.

This decision defines PRODUCT provenance semantics and validity rules.
It does not define DRMCP parser, index, traversal, diagnostic, response, or projection mechanisms.

## Decision

### Canonical forward provenance

A Work Item persists required non-empty `source_refs`.

`source_refs` is an unordered set of active canonical references.
The accepted reference classes are those defined by the canonical artifact-reference contract.
The set may contain Specification refs and supported Design Record public IDs.

A Work Item lists every direct upstream artifact that materially motivates the Work Item.
A Work Item omits incidental context.
A Work Item also omits transitively reachable ancestors unless they are independently material.

### Task provenance

A Task persists no source field.

Task provenance is reached through the Task `work_item` relation.
A Task does not inherit or repeat the owning Work Item `source_refs`.

A downstream Work Item created or decomposed from a Task cites that exact Task ID in `source_refs`.
The downstream Work Item does not automatically copy the source Task owner or its upstream sources.
The owner Work Item is included only when it independently and directly motivates the downstream Work Item.

### Requirement reverse relation

Requirement records do not persist canonical `work_items` metadata.

The direct Requirement reverse relation is derived from Work Items whose `source_refs` directly contain the Requirement ID.
The derived relation is unordered and duplicate-free.
Transitive descendants are excluded from this direct reverse relation.
Transitive traversal remains a separate DRMCP capability.

### Validity

The following Work Item `source_refs` states are invalid:

- an empty set;
- duplicate canonical refs;
- the Work Item's own canonical identity;
- an unresolved ref;
- an unrecognized or noncanonical ref form.

Every cycle in the semantic Work Item provenance graph is invalid.
A Task ref is normalized to the Task's owning Work Item for cycle semantics.
The owner-resolution mechanism belongs to DRMCP.

Work Item `tasks`, Task `work_item`, and Task `depends_on` remain separate persisted membership or dependency relations.
They are excluded from semantic provenance-cycle edges unless a Work Item source ref points to a Task.

### Ownership boundary

PRODUCT owns persisted provenance semantics and invalid conditions.

DRMCP owns:

- parsing and indexing;
- Task-owner resolution mechanics;
- direct reverse lookup;
- transitive traversal;
- cycle-analysis algorithms;
- diagnostics;
- response schemas;
- user-visible projections.

## Rationale

Work Item provenance must describe why the Work Item exists.
Work Items are the stable workflow units that own resolution flows and Task graphs.

Keeping source metadata only on Work Items avoids duplicated Task provenance.
The `work_item` relation already supplies the Task ownership path.

Direct material references preserve exact provenance without copying the whole reachable graph.
Exact Task refs preserve Task-to-Work-Item decomposition without introducing parent or Hub fields.

A derived Requirement reverse relation removes reciprocal state.
The direct-only boundary also keeps reverse membership distinct from graph traversal.

Unordered set semantics match provenance meaning.
Duplicate, self, unresolved, and cyclic refs would make that meaning ambiguous or invalid.

Reusing the active canonical reference classes avoids a second Work Item-only reference taxonomy.

## Rejected alternatives

| alternative | rejection reason |
|---|---|
| Keep Work Item `source_requirement` | A scalar Requirement-only field cannot represent generic or multiple material sources. |
| Keep persisted Requirement `work_items` | Reciprocal persistence duplicates derivable state and can become stale. |
| Add `source_refs` to both Work Items and Tasks | Task source metadata would duplicate provenance available through `work_item`. |
| Copy a source Task owner and all upstream sources | Transitive copying creates stale and misleading direct provenance. |
| Treat transitive descendants as direct Requirement reverse membership | Direct membership and graph traversal have different semantics. |
| Permit unresolved refs as future candidates | Persisted provenance must identify existing canonical sources. |
| Permit semantic provenance cycles | Cycles prevent a valid upstream provenance interpretation. |
| Define Work Item-only reference forms | A second taxonomy would duplicate the canonical artifact-reference contract. |
| Add Hub, parent Work Item, or child Work Item fields | Exact Task-to-Work-Item provenance already represents decomposition. |

## Consequences

- `PRODUCT-REQ-SPEC-006` reflects Work Item-only `source_refs` and Task provenance through `work_item`.
- Requirement authoring must remove canonical persisted `work_items` ownership.
- Work Item authoring must define required generic `source_refs` and direct material source selection.
- Task authoring must remove Task source metadata while preserving W016 Task-type rules.
- Traceability Specifications must define direct reverse relations and provenance validity.
- `PRODUCT-ADR-SPEC-008` defines migration from the legacy persisted fields.
- Existing record migration remains outside `PRODUCT-WORK-SPEC-017` execution scope.
- DRMCP needs later app-local design and implementation work for mechanisms and diagnostics.
- Current Specifications replace conflicting relation sections in `V01-ADR-091` and `V01-ADR-092` without superseding their unrelated decisions.

## Evidence

- `PRODUCT-REQ-SPEC-006`: accepted generic workflow source-relation requirement.
- `PRODUCT-TASK-SPEC-017-02`: D-001 through D-012 and D-017.
- `PRODUCT-TASK-SPEC-017-04`: C-001 through C-026 conflict dispositions.
- `PRODUCT-TASK-SPEC-017-05`: ADR routing and this coherent decision boundary.
- `PRODUCT-ADR-SPEC-001`: PRODUCT semantic ownership and app-local implementation boundary.
- `PRODUCT-ADR-SPEC-004` through `PRODUCT-ADR-SPEC-006`: accepted W016 Task contract preserved by this decision.
- `V01-ADR-091`: historical workflow relation and Work Item / Task responsibility context.
- `V01-ADR-092`: historical reciprocal workflow relation validation context.
