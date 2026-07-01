# PRODUCT-TASK-SPEC-017-02: Run REQ-006 interactive decision loop

- **id**: PRODUCT-TASK-SPEC-017-02
- **status**: done
- **date**: 2026-07-01
- **work_item**: PRODUCT-WORK-SPEC-017
- **source_requirement**: PRODUCT-REQ-SPEC-006
- **estimate**: 1d
- **depends_on**:
  - PRODUCT-TASK-SPEC-017-01
- **outputs**:
  - PRODUCT-TASK-SPEC-017-02

## Goal

Persist every unresolved REQ-006 decision one at a time.

Finish with each required decision decided, deferred, or blocked for a named reason.

## Work

- Resume from the current decision cursor.
- Ask exactly one decision per user turn.
- Check each answer against accepted Requirements and prior decisions.
- Persist each explicit answer before advancing.
- Update status, summary, dependencies, ADR route, canonical target, and cursor.
- Keep W016-owned Task-type decisions outside this loop.

## Done condition

- Every required decision has status `decided`, `deferred`, or validly `blocked`.
- Every explicit user answer is durably recorded.
- No more than one decision is `in_discussion`.
- No accepted constraint is reopened without an explicit contradiction.
- ADR routing and Specification targets are ready for downstream Tasks.

## Verification

- Inspect the scoped T02 diff after every persisted answer.
- Confirm the current decision ID and cursor match.
- Confirm no unrelated decision row changed.
- Confirm no decision is marked `recorded` before canonical synchronization.
- Confirm no ADR, Specification, review, migration execution, or DRMCP implementation begins in T02.

## Evidence

### Bootstrap disposition

`BOOTSTRAP-001` reuses the user decision recorded in `PRODUCT-TASK-SPEC-016-01`.

- Use the current canonical `source_requirement` and `work_item` metadata contract.
- Do not add an unaccepted primary Task type or `source_refs` field.
- Do not require migration of this workflow Task.
- Treat the bootstrap choice as workflow authoring policy only.

### Loop state

Loop status: decision_complete

Current decision: none

Inventory source: PRODUCT-TASK-SPEC-017-01

### Decision confirmation loop

| ID | Topic | Status | Depends on | Decision summary | ADR route | Canonical target |
|---|---|---|---|---|---|---|
| D-001 | `source_refs` ownership and minimum cardinality | decided | — | Work Item `source_refs` is required and contains at least one ref. Task has no `source_refs` field. Task provenance is reached through its owning `work_item`. PRODUCT-REQ-SPEC-006 must be corrected because its accepted Required Outcome currently requires Task `source_refs`. | candidate | `PRODUCT-REQ-SPEC-006`; `spec:product.design_records.authoring_standards.work_item_authoring`; `spec:product.design_records.authoring_standards.task_authoring` |
| D-002 | Source-ref ordering semantics | decided | — | `source_refs` order has no semantic meaning. The field is interpreted as an unordered set of canonical refs. Reordering alone does not change provenance semantics. | not_required | `spec:product.design_records.traceability.metadata_schema` |
| D-003 | Duplicate source refs | decided | — | Duplicate canonical refs are invalid. Validation reports an error; persistence does not silently deduplicate or normalize them. | not_required | `spec:product.design_records.traceability.resolve_and_validation` |
| D-004 | Self-reference | decided | — | A Work Item must not include its own canonical identity in `source_refs`. Self-reference is a validation error. | not_required | `spec:product.design_records.traceability.resolve_and_validation` |
| D-005 | Existence and unknown refs | decided | D-006 | Every persisted Work Item `source_refs` entry must resolve at validation time. Unresolved-but-well-formed refs and unrecognized ref forms are validation errors. `source_refs` is not used for future candidates. | not_required | `spec:product.design_records.traceability.resolve_and_validation` |
| D-006 | Allowed public-ID and semantic-ref boundary | decided | — | Work Item `source_refs` accepts every active canonical reference class defined by `artifact_refs`: spec refs; ADR, investigation, requirement, work item, and task public IDs; and resolver-supported legacy issued IDs. | not_required | `spec:product.design_records.traceability.artifact_refs` |
| D-007 | Multiple-source materiality | decided | D-001, D-006 | Every direct upstream artifact that materially motivates the Work Item must be listed. Incidental context and transitively reachable ancestors are omitted unless independently material. | not_required | `spec:product.design_records.authoring_standards.work_item_authoring` |
| D-008 | Work Item sourced from a Task | decided | — | A downstream Work Item created or decomposed from a Task must cite that exact Task ID in `source_refs`. The Task is the direct provenance edge; the Task owner's upstream sources are not copied merely because they are transitively reachable. | candidate | `spec:product.design_records.authoring_standards.work_item_authoring` |
| D-009 | Parent Work Item source inclusion | decided | D-008 | A downstream Work Item does not automatically list the upstream Task's owning Work Item. It lists that Work Item only when the Work Item independently and directly materially motivates the downstream Work Item under D-007. | not_required | `spec:product.design_records.authoring_standards.work_item_authoring` |
| D-010 | Task source inheritance | decided | D-001 | Task does not inherit or repeat Work Item source refs because Task has no `source_refs` field. The `work_item` relation provides provenance traversal. | not_required | `spec:product.design_records.authoring_standards.task_authoring` |
| D-011 | Requirement reverse derivation depth | decided | — | Requirement reverse relations include only Work Items that directly cite the Requirement in `source_refs`. Transitively related descendants remain available only through separate DRMCP graph traversal. | candidate | `spec:product.design_records.traceability.metadata_schema` |
| D-012 | Requirement reverse projection semantics | decided | D-002, D-011 | The canonical reverse projection is the unordered, duplicate-free set of Work Items whose `source_refs` directly contain the Requirement ID. Transitive descendants are excluded from this relation. | not_required | `spec:product.design_records.authoring_standards.requirement_authoring`; `spec:product.design_records.traceability.metadata_schema` |
| D-013 | Work Item `source_requirement` migration | decided | D-001, D-006, D-016 | Migration copies the existing `source_requirement` value into a one-element `source_refs` list, removes `source_requirement` in the same record update, and does not infer additional sources. | candidate | `spec:product.design_records.authoring_standards.work_item_authoring` |
| D-014 | Task `source_requirement` migration | decided | D-016 | Migration removes Task `source_requirement` without replacement, preserves `work_item`, and relies exclusively on the owning Work Item for provenance. | candidate | `spec:product.design_records.authoring_standards.task_authoring` |
| D-015 | Requirement `work_items` migration | decided | D-011, D-012, D-016 | Migration may remove Requirement `work_items` only after the old list exactly matches the unordered, duplicate-free set of Work Items that directly cite the Requirement in `source_refs`. Mismatches block migration and are not silently repaired. | candidate | `spec:product.design_records.authoring_standards.requirement_authoring` |
| D-016 | Mixed old and new metadata period | decided | — | Repository-wide staged migration is allowed, but each migrated record switches atomically. A Work Item must not contain both `source_requirement` and `source_refs`; a migrated Requirement removes `work_items` in its record update. | candidate | `spec:product.design_records.traceability.metadata_schema` |
| D-017 | Source-relation cycle validation | decided | D-008, D-009, D-010, D-011 | Every cycle in the Work Item semantic provenance graph is invalid. Task refs are normalized to their owning Work Item for cycle analysis; the owner-resolution mechanism is not fixed by this decision. Work Item `tasks` and derived reverse projections are excluded. | candidate | `spec:product.design_records.traceability.resolve_and_validation` |
| D-018 | PRODUCT semantics versus DRMCP projection | decided | D-011, D-012, D-017 | PRODUCT defines canonical persisted provenance semantics and validity rules. DRMCP owns indexing, traversal algorithms, reverse lookup, Task-owner resolution mechanics, diagnostics, response schemas, and user-visible projections. | not_required | `spec:product.design_records.traceability`; `spec:product.design_records.artifact_model.artifact_responsibility_matrix` |
| D-019 | Shared `task_authoring` integration | decided | PRODUCT-TASK-SPEC-016-07, D-010, D-014 | PRODUCT-TASK-SPEC-016-07 completed on 2026-07-01. W017 may now write source-relation rules after the accepted REQ-005 Task contract while preserving D-010 and D-014. | not_required | `spec:product.design_records.authoring_standards.task_authoring` |

### Decided D-018 result

The user accepted the following result on 2026-06-30:

- PRODUCT defines canonical persisted provenance semantics and validity rules.
- DRMCP owns indexing, traversal algorithms, reverse lookup, and Task-owner resolution mechanics.
- DRMCP also owns diagnostics, response schemas, and user-visible projections.

### Decided D-019 result

The dependency cleared on 2026-07-01.

- PRODUCT-TASK-SPEC-016-07 is `done`.
- W017 may now update the shared authoring Specifications.
- D-010 and D-014 remain binding during that update.
- No additional user decision is required.

### Decided D-017 result

The user accepted the following result on 2026-06-30:

- Every cycle in the Work Item semantic provenance graph is invalid.
- Task refs are normalized to their owning Work Item for cycle analysis.
- The Task owner-resolution mechanism is not fixed by this decision.
- Work Item `tasks` and derived reverse projections are excluded from cycle analysis.

### Decided D-015 result

The user accepted the following result on 2026-06-30:

- Requirement `work_items` may be removed only after exact-match verification.
- The old list is compared as an unordered, duplicate-free set against Work Items that directly cite the Requirement in `source_refs`.
- A mismatch blocks migration.
- Migration does not silently delete, add, or repair relations.

### Decided D-014 result

The user accepted the following result on 2026-06-30:

- Migration removes Task `source_requirement` without replacement.
- Task `work_item` is preserved.
- Task provenance relies exclusively on the owning Work Item.

### Decided D-013 result

The user accepted the following result on 2026-06-30:

- Migration copies the existing Work Item `source_requirement` value into a one-element `source_refs` list.
- `source_requirement` is removed in the same record update.
- Migration does not infer or append additional sources.

### Decided D-016 result

The user accepted the following result on 2026-06-30:

- Repository-wide staged migration is allowed.
- Each migrated record switches atomically.
- A Work Item must not contain both `source_requirement` and `source_refs`.
- A migrated Requirement removes `work_items` in the same record update.

### Decided D-011 and derived D-012 result

The user accepted the following D-011 result on 2026-06-30:

- Requirement reverse relations include only Work Items that directly cite the Requirement in `source_refs`.
- Transitively related descendants are excluded from this canonical reverse relation.
- Transitive descendants remain available through separate DRMCP graph traversal.

Together with D-002, this directly determines D-012:

- The canonical reverse projection is an unordered, duplicate-free set.
- It contains exactly the Work Items whose `source_refs` directly contain the Requirement ID.

### Decided D-009 result

The user accepted the following result on 2026-06-30:

- A downstream Work Item does not automatically list the upstream Task's owning Work Item.
- The owning Work Item is listed only when it independently and directly materially motivates the downstream Work Item under D-007.
- Mere transitive reachability through the source Task is insufficient.

### Decided D-008 result

The user accepted the following result on 2026-06-30:

- A downstream Work Item created or decomposed from a Task must cite that exact Task ID in `source_refs`.
- The Task is the direct provenance edge.
- The Task owner's upstream sources are not copied merely because they are transitively reachable.

### Decided D-007 result

The user accepted the following result on 2026-06-30:

- Every direct upstream artifact that materially motivates the Work Item must be listed.
- Incidental context is omitted.
- Transitively reachable ancestors are omitted unless they are independently material to the Work Item.

### Decided D-005 result

The user accepted the following result on 2026-06-30:

- Every persisted Work Item `source_refs` entry must resolve at validation time.
- Unresolved-but-well-formed refs are validation errors.
- Unrecognized ref forms are validation errors.
- `source_refs` is not used for future artifact candidates.

### Decided D-006 result

The user accepted the following result on 2026-06-30:

- Work Item `source_refs` accepts every active canonical reference class defined by `artifact_refs`.
- This includes `spec:` refs; ADR, investigation, requirement, work item, and task public IDs; and resolver-supported legacy issued IDs.
- W017 does not create a narrower Work Item-only reference taxonomy.

### Decided D-004 result

The user accepted the following result on 2026-06-30:

- A Work Item must not include its own canonical identity in `source_refs`.
- Self-reference is a validation error.

### Decided D-003 result

The user accepted the following result on 2026-06-30:

- Duplicate canonical refs are invalid.
- Validation reports an error.
- Persistence does not silently deduplicate or normalize duplicate refs.

### Decided D-002 result

The user accepted the following result on 2026-06-30:

- `source_refs` order has no semantic meaning.
- The field is interpreted as an unordered set of canonical refs.
- Reordering alone does not change provenance semantics.

### Decided D-001 result

The user accepted the following result on 2026-06-30:

- Work Item `source_refs` is required and contains at least one ref.
- Task has no `source_refs` field.
- Task provenance is reached through its owning `work_item` relation.
- PRODUCT-REQ-SPEC-006 is a required correction target because its accepted Required Outcome currently requires Task `source_refs`.

The decision avoids duplicated Task provenance and stale copied refs.
D-010 is also decided as a direct consequence.
D-014 now owns removal of Task `source_requirement` without replacement by Task `source_refs`.

### Authority-resolved scope

- Accepted PRODUCT-REQ-SPEC-006 currently requires Work Item and Task provenance to use generic `source_refs`.
- Requirement `work_items` will not remain persisted canonical metadata.
- Task `work_item` and Work Item `tasks` remain explicit membership relations.
- Dedicated Hub, parent Work Item, and child Work Item fields remain excluded.
- Concrete DRMCP request, response, diagnostic, and implementation behavior remains outside W017.
- Task-type semantics remain owned by W016.
