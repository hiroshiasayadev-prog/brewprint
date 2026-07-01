# PRODUCT-TASK-SPEC-017-01: Inventory REQ-006 source-relation decisions

- **id**: PRODUCT-TASK-SPEC-017-01
- **status**: done
- **date**: 2026-07-01
- **work_item**: PRODUCT-WORK-SPEC-017
- **source_requirement**: PRODUCT-REQ-SPEC-006
- **estimate**: 0.5d
- **depends_on**: []
- **outputs**:
  - PRODUCT-TASK-SPEC-017-01

## Goal

Create the complete known REQ-006 decision inventory and dependency order.

Separate accepted constraints, open source-relation choices, and W016-owned shared-surface blockers.

## Work

- Read REQ-006 and the directly affected PRODUCT Specifications.
- Resolve facts already fixed by accepted authority without asking the user.
- Identify unresolved provenance, decomposition, reverse-relation, validation, and migration decisions.
- Record dependency edges, ADR routes, and likely canonical targets.
- Route Task-type semantics to W016 instead of deciding them in W017.
- Select the first unblocked decision for T02.

## Done condition

- Every currently known REQ-006 decision has a stable local ID.
- Accepted constraints are separated from user decisions.
- Dependencies and W016-owned blockers are explicit.
- Bootstrap workflow Tasks are excluded from migration action.
- One first unblocked decision is selected for T02.

## Verification

- Compare the inventory with every Required Outcome and Explicitly Excluded Scope item in PRODUCT-REQ-SPEC-006.
- Confirm Task-type semantics remain owned by W016.
- Confirm no accepted REQ-006 constraint is restated as an open user decision.
- Confirm no more than one decision is selected for discussion.

## Evidence

### Bootstrap disposition

`BOOTSTRAP-001` reuses the user decision recorded in `PRODUCT-TASK-SPEC-016-01`.

- Use the current canonical Task metadata contract.
- Use `source_requirement` and `work_item`.
- Do not add `type`, `task_type`, `primary_type`, or `source_refs` before Specification acceptance.
- Do not require migration of the W016 or W017 workflow Tasks.
- Treat the bootstrap choice as workflow authoring policy, not as the canonical REQ-006 product contract.

### Accepted authority constraints

The following points are fixed by PRODUCT-REQ-SPEC-006 and are not user decisions in T02.

| constraint | accepted result |
|---|---|
| Persisted forward relation | PRODUCT-REQ-SPEC-006 currently says Work Item and Task provenance use generic `source_refs`; D-001 requires correction so only Work Item keeps `source_refs`, while Task provenance is reached through `work_item`. |
| Allowed identity families | `source_refs` may contain canonical Design Record public IDs and active semantic refs. |
| Multiple upstream sources | Multiple refs are allowed when multiple upstream artifacts materially motivate the record. |
| Task membership | Task `work_item` remains the owning Work Item relation. |
| Work Item membership list | Work Item `tasks` remains the explicit list of owned Tasks. |
| Requirement reverse storage | Requirement `work_items` will leave the canonical persisted Requirement contract. |
| Requirement reverse view | Requirement-to-Work-Item relations will be derived from Work Item `source_refs`. |
| Downstream decomposition | A downstream Work Item may reference the Task that created or decomposed it. |
| Downstream ownership | Each downstream Work Item owns its Task graph and completion condition. |
| Hierarchy fields | Dedicated Hub, parent Work Item, and child Work Item fields are not introduced. |
| Validation scope | Existence, canonical identity, duplicates, and self-reference require canonical rules. |
| Migration scope | Existing `source_requirement` and Requirement `work_items` require an explicit migration contract. |
| PRODUCT boundary | PRODUCT owns canonical provenance semantics and persisted relation ownership. |
| DRMCP boundary | DRMCP owns indexing, reverse traversal, graph validation, and user-visible projections. |
| Task types | Task-type and single-responsibility semantics remain outside W017 and are owned by W016. |
| Concrete tool schema | Concrete DRMCP request and response schemas remain outside W017. |

### Decision inventory

| ID | Topic | Status | Depends on | Decision summary | ADR route | Canonical target |
|---|---|---|---|---|---|---|
| D-001 | `source_refs` ownership and minimum cardinality | decided | — | Work Item `source_refs` is required with at least one ref. Task has no `source_refs` field; provenance is reached through `work_item`. PRODUCT-REQ-SPEC-006 requires correction. | candidate | `PRODUCT-REQ-SPEC-006`; `spec:product.design_records.authoring_standards.work_item_authoring`; `spec:product.design_records.authoring_standards.task_authoring` |
| D-002 | Source-ref ordering semantics | decided | — | `source_refs` order has no semantic meaning and is interpreted as an unordered set of canonical refs. | not_required | `spec:product.design_records.traceability.metadata_schema` |
| D-003 | Duplicate source refs | decided | — | Duplicate canonical refs are invalid. Validation reports an error; persistence does not silently deduplicate or normalize them. | not_required | `spec:product.design_records.traceability.resolve_and_validation` |
| D-004 | Self-reference | decided | — | A Work Item must not include its own canonical identity in `source_refs`. Self-reference is a validation error. | not_required | `spec:product.design_records.traceability.resolve_and_validation` |
| D-005 | Existence and unknown refs | decided | D-006 | Every persisted Work Item `source_refs` entry must resolve at validation time. Unresolved-but-well-formed refs and unrecognized ref forms are validation errors. `source_refs` is not used for future candidates. | not_required | `spec:product.design_records.traceability.resolve_and_validation` |
| D-006 | Allowed public-ID and semantic-ref boundary | decided | — | Work Item `source_refs` accepts every active canonical reference class defined by `artifact_refs`: spec refs; ADR, investigation, requirement, work item, and task public IDs; and resolver-supported legacy issued IDs. | not_required | `spec:product.design_records.traceability.artifact_refs` |
| D-007 | Multiple-source materiality | decided | D-001, D-006 | Every direct upstream artifact that materially motivates the Work Item must be listed. Incidental context and transitively reachable ancestors are omitted unless independently material. | not_required | `spec:product.design_records.authoring_standards.work_item_authoring` |
| D-008 | Work Item sourced from a Task | decided | — | A downstream Work Item created or decomposed from a Task must cite that exact Task ID in `source_refs`. The Task is the direct provenance edge; the Task owner's upstream sources are not copied merely because they are transitively reachable. | candidate | `spec:product.design_records.authoring_standards.work_item_authoring` |
| D-009 | Parent Work Item source inclusion | decided | D-008 | A downstream Work Item does not automatically list the upstream Task's owning Work Item. It lists that Work Item only when the Work Item independently and directly materially motivates the downstream Work Item under D-007. | not_required | `spec:product.design_records.authoring_standards.work_item_authoring` |
| D-010 | Task source inheritance | decided | D-001 | Task neither inherits nor repeats Work Item source refs because Task has no `source_refs` field. Provenance is traversed through `work_item`. | not_required | `spec:product.design_records.authoring_standards.task_authoring` |
| D-011 | Requirement reverse derivation depth | decided | — | Requirement reverse relations include only Work Items that directly cite the Requirement in `source_refs`. Transitively related descendants remain available only through separate DRMCP graph traversal. | candidate | `spec:product.design_records.traceability.metadata_schema` |
| D-012 | Requirement reverse projection semantics | decided | D-002, D-011 | The canonical reverse projection is the unordered, duplicate-free set of Work Items whose `source_refs` directly contain the Requirement ID. Transitive descendants are excluded from this relation. | not_required | `spec:product.design_records.authoring_standards.requirement_authoring`; `spec:product.design_records.traceability.metadata_schema` |
| D-013 | Work Item `source_requirement` migration | decided | D-001, D-006, D-016 | Migration copies the existing `source_requirement` value into a one-element `source_refs` list, removes `source_requirement` in the same record update, and does not infer additional sources. | candidate | `spec:product.design_records.authoring_standards.work_item_authoring` |
| D-014 | Task `source_requirement` migration | decided | D-016 | Migration removes Task `source_requirement` without replacement, preserves `work_item`, and relies exclusively on the owning Work Item for provenance. | candidate | `spec:product.design_records.authoring_standards.task_authoring` |
| D-015 | Requirement `work_items` migration | decided | D-011, D-012, D-016 | Migration may remove Requirement `work_items` only after the old list exactly matches the unordered, duplicate-free set of Work Items that directly cite the Requirement in `source_refs`. Mismatches block migration and are not silently repaired. | candidate | `spec:product.design_records.authoring_standards.requirement_authoring` |
| D-016 | Mixed old and new metadata period | decided | — | Repository-wide staged migration is allowed, but each migrated record switches atomically. A Work Item must not contain both `source_requirement` and `source_refs`; a migrated Requirement removes `work_items` in its record update. | candidate | `spec:product.design_records.traceability.metadata_schema` |
| D-017 | Source-relation cycle validation | decided | D-008, D-009, D-010, D-011 | Every cycle in the Work Item semantic provenance graph is invalid. Task refs are normalized to their owning Work Item for cycle analysis; the owner-resolution mechanism is not fixed by this decision. Work Item `tasks` and derived reverse projections are excluded. | candidate | `spec:product.design_records.traceability.resolve_and_validation` |
| D-018 | PRODUCT semantics versus DRMCP projection | decided | D-011, D-012, D-017 | PRODUCT defines canonical persisted provenance semantics and validity rules. DRMCP owns indexing, traversal algorithms, reverse lookup, Task-owner resolution mechanics, diagnostics, response schemas, and user-visible projections. | not_required | `spec:product.design_records.traceability`; `spec:product.design_records.artifact_model.artifact_responsibility_matrix` |
| D-019 | Shared `task_authoring` integration | decided | PRODUCT-TASK-SPEC-016-07, D-010, D-014 | PRODUCT-TASK-SPEC-016-07 completed on 2026-07-01. W017 may now write source-relation rules after the accepted REQ-005 Task contract, while preserving D-010 and D-014. No additional user decision is required. | not_required | `spec:product.design_records.authoring_standards.task_authoring` |

### Authority-resolved ownership

- Task `work_item` and Work Item `tasks` remain persisted membership relations.
- Their concrete writer transaction and diagnostic schema remain DRMCP-owned.
- Task-to-Work-Item decomposition uses source refs and independent downstream Work Item ownership.
- No parent, child, or Hub Work Item metadata is introduced.
- W016 writes shared Task-type content before W017 writes source-relation content.

### Scoped verification result

The inventory covers every required REQ-006 decision area named by the W017 startup contract.

No repository-wide traversal was used.
There is no current user-decision cursor. D-001 through D-019 are decided. PRODUCT-TASK-SPEC-016-07 completed on 2026-07-01, so the shared-file writer dependency is cleared.
