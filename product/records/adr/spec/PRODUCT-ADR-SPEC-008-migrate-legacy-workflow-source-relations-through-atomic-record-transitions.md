# PRODUCT-ADR-SPEC-008: Migrate legacy workflow source relations through atomic record transitions

- **status**: accepted
- **date**: 2026-07-01
- **depends_on**:
  - PRODUCT-ADR-SPEC-007
- **supersedes**: []
- **migrated_to_spec**: null

## Context

`PRODUCT-ADR-SPEC-007` replaces the legacy workflow source-relation model.

Existing records still use three legacy metadata shapes:

- Work Item `source_requirement`;
- Task `source_requirement`;
- Requirement `work_items`.

A repository-wide instant switch would create a large release boundary.
A prolonged dual-field period within one record would create ambiguous canonical state.
Migration also risks inventing provenance or silently deleting mismatched reverse relations.

This decision defines migration semantics only.
It does not execute record migration.
It does not define concrete DRMCP migration commands, diagnostics, or transaction APIs.

## Decision

### Migration sequencing

Repository-wide staged migration is allowed.

Each record transitions atomically.
A record must not persist its legacy and replacement provenance fields together.

The repository may contain migrated and unmigrated records during the staged migration period.
Consumers must use the migration state and applicable record shape without treating same-record dual fields as valid.

### Work Item transition

For each migrated Work Item:

1. Copy the existing `source_requirement` value into a one-element `source_refs` list.
2. Remove `source_requirement` in the same record update.
3. Preserve all unrelated metadata and body content.

Migration does not infer or append additional sources.
Later authoring may add another direct material source under the accepted Work Item contract.

### Task transition

For each migrated Task:

1. Remove `source_requirement` without replacement.
2. Preserve `work_item`.
3. Preserve `depends_on` and all unrelated metadata and body content.

Migration must not create Task `source_refs`.
Task provenance remains reachable through the owning Work Item.

### Requirement transition

Before removing Requirement `work_items`, migration derives the direct reverse set from Work Item `source_refs`.

The legacy list and derived set are compared as unordered, duplicate-free sets.
Migration may remove `work_items` only when both sets match exactly.

A mismatch blocks that Requirement transition.
Migration must not silently add, delete, or repair relations to force a match.

### Scope boundary

This ADR defines the accepted migration result and failure boundary.

DRMCP or another execution owner must separately define:

- migration command and transaction mechanics;
- repository scan order;
- diagnostics and reporting;
- rollback behavior;
- resume behavior;
- release and compatibility gates.

`PRODUCT-WORK-SPEC-017` does not perform bulk migration.
The W016 and W017 workflow bootstrap Tasks remain outside migration execution scope.

## Rationale

Staged migration limits change size and allows reviewable batches.
Atomic per-record transitions prevent ambiguous source-of-truth fields.

A one-element Work Item conversion preserves existing explicit provenance without fabricating new edges.
Removing Task source metadata avoids carrying the rejected duplicated model into the new shape.

Exact-match verification protects Requirement reverse relations from silent information loss.
Blocking on mismatch exposes stale or inconsistent graph state for explicit repair.

Separating migration semantics from execution mechanics preserves the PRODUCT and DRMCP ownership boundary.

## Rejected alternatives

| alternative | rejection reason |
|---|---|
| Require one repository-wide big-bang transaction | The release boundary would be unnecessarily large and difficult to review. |
| Permit legacy and replacement fields in the same record | Dual fields create ambiguous canonical provenance and synchronization rules. |
| Infer additional Work Item sources during conversion | Migration lacks authority to invent direct material provenance. |
| Replace Task `source_requirement` with Task `source_refs` | The accepted model stores Task provenance only through `work_item`. |
| Remove Requirement `work_items` without verification | Silent deletion could hide stale or missing direct relations. |
| Repair mismatches automatically | Automatic repair would make an unreviewed provenance decision. |
| Define migration tooling in PRODUCT | Concrete commands, transactions, and diagnostics belong to the execution owner. |

## Consequences

- Canonical Specifications must define staged migration and atomic record transitions.
- Work Item migration preserves exactly one existing source and infers none.
- Task migration removes duplicated source metadata.
- Requirement migration requires exact reverse-set agreement.
- Mismatched Requirement relations require explicit investigation or correction before migration.
- Downstream DRMCP work must define operational migration and diagnostic contracts.
- Existing records are not changed by this ADR or by `PRODUCT-WORK-SPEC-017`.
- The W016 and W017 workflow Tasks require no migration action.

## Evidence

- `PRODUCT-ADR-SPEC-007`: accepted canonical workflow provenance model.
- `PRODUCT-REQ-SPEC-006`: accepted requirement for explicit migration behavior.
- `PRODUCT-TASK-SPEC-017-02`: D-013 through D-016.
- `PRODUCT-TASK-SPEC-017-04`: C-003, C-009, C-012, and C-024 migration conflict dispositions.
- `PRODUCT-TASK-SPEC-017-05`: ADR routing and this coherent migration boundary.
- `V01-ADR-091` and `V01-ADR-092`: historical legacy relation shapes requiring migration.
