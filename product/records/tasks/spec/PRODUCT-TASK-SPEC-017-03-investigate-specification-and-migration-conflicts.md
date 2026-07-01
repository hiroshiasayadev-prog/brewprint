# PRODUCT-TASK-SPEC-017-03: Investigate Specification and migration conflicts

- **id**: PRODUCT-TASK-SPEC-017-03
- **status**: done
- **date**: 2026-07-01
- **work_item**: PRODUCT-WORK-SPEC-017
- **source_requirement**: PRODUCT-REQ-SPEC-006
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-017-02
- **outputs**:
  - PRODUCT-TASK-SPEC-017-03

## Goal

Identify every material conflict between decided REQ-006 semantics and current canonical Specifications.

Produce an exact semantic, migration, and shared-writer conflict inventory without selecting dispositions.

## Work

- Compare decided T02 entries with current Requirement, Work Item, Task, artifact-model, and traceability Specifications.
- Identify stale reciprocal relations and Requirement-only provenance rules.
- Identify migration conflicts for `source_requirement` and Requirement `work_items`.
- Identify shared-file overlap with REQ-005 and W016.
- Separate PRODUCT semantic conflicts from downstream DRMCP implementation impacts.

## Done condition

- Every material conflict has an exact canonical ref and affected section description.
- Every conflict has a declared owner: W017, W016 dependency, shared serialization, or downstream DRMCP.
- No conflict disposition or product decision is introduced.
- T04 can ask one conflict decision at a time.

## Verification

- Trace each conflict to a decided T02 entry or accepted Requirement.
- Confirm W016-owned Task-type semantics remain undecided by W017.
- Confirm no Specification, ADR, migration target, source, test, or fixture file changed.

## Evidence

### Investigation authority

- `PRODUCT-TASK-SPEC-017-01` and `PRODUCT-TASK-SPEC-017-02` are the authority for D-001 through D-019.
- T02 has `Loop status: decision_complete` and `Current decision: none`.
- `PRODUCT-TASK-SPEC-016-07` is `done` and is the authority for the current W16-written shared Specification state.
- `PRODUCT-REQ-SPEC-005` constrains Task type and single-responsibility wording.
- `PRODUCT-REQ-SPEC-006` supplies the accepted source-relation requirement, except for the D-001 correction target.
- DRMCP is non-operational. Filesystem reads were limited to the declared workflow records, Requirements, and nine canonical PRODUCT Specification refs.

### Routing classification summary

| classification | conflict IDs | result |
|---|---|---|
| Repository-resolvable correction | C-001 through C-026 and C-031 | Accepted authority determines the required semantic direction or migration exclusion. T04 user judgment is not required. |
| T04 user decision required | None | No materially different conflict disposition remains after D-001 through D-019. |
| Downstream DRMCP routing only | C-027 through C-030 | PRODUCT must define semantics. DRMCP later owns implementation and projection details. |
| No conflict | N-001 through N-008 | Existing wording is compatible or is an explicit workflow-only exclusion. |

No conflict disposition is selected by this Task.
The routing classification only states whether accepted authority already resolves the issue.
ADR routing remains owned by T05.

### Conflict inventory identity

| ID | topic | exact canonical ref | affected heading or section | conflict class | owner |
|---|---|---|---|---|---|
| C-001 | REQ-006 requires Task `source_refs` | `PRODUCT-REQ-SPEC-006` | `## Requirement`; `## Required Outcome`; `## Boundary` | requirement-correction | Requirement-correction |
| C-002 | Requirement persists reciprocal `work_items` | `spec:product.design_records.authoring_standards.requirement_authoring` | `### Metadata schema`; `### Kind-specific authoring rules`; `### Canonical reference policy`; `### Create`; `### Update` | reverse-relation | W017 |
| C-003 | Requirement `work_items` removal lacks exact-match migration guard | `spec:product.design_records.authoring_standards.requirement_authoring` | `### Metadata schema`; `### Update` | migration | W017 |
| C-004 | Work Item persists required scalar `source_requirement` | `spec:product.design_records.authoring_standards.work_item_authoring` | `### Metadata schema`; `### Canonical reference policy`; `### Create`; `### Update` | metadata-shape | W017 |
| C-005 | Work Item requires reciprocal Requirement linkage and Task source matching | `spec:product.design_records.authoring_standards.work_item_authoring` | `### Metadata schema`; `### Kind-specific authoring rules` | reverse-relation | W017 |
| C-006 | Work Item lacks required non-empty generic `source_refs` | `spec:product.design_records.authoring_standards.work_item_authoring` | `### Metadata schema`; `### Canonical reference policy`; `### Create`; `### Update` | metadata-shape | W017 |
| C-007 | Work Item lacks direct and material source-selection rules | `spec:product.design_records.authoring_standards.work_item_authoring` | `### Kind-specific authoring rules` | semantic | W017 |
| C-008 | Task-created Work Item provenance is not defined | `spec:product.design_records.authoring_standards.work_item_authoring` | `### Kind-specific authoring rules`; `Parent coordination boundary` | semantic | W017 |
| C-009 | Work Item atomic `source_requirement` conversion is not defined | `spec:product.design_records.authoring_standards.work_item_authoring` | `### Metadata schema`; `### Update` | migration | W017 |
| C-010 | Task persists required scalar `source_requirement` | `spec:product.design_records.authoring_standards.task_authoring` | `### Metadata schema`; `### Canonical reference policy`; `### Create`; `### Update` | metadata-shape | W017 |
| C-011 | Task provenance is required to match the Work Item Requirement | `spec:product.design_records.authoring_standards.task_authoring` | `### Metadata schema`; `#### General Task rules` | semantic | W017 |
| C-012 | Task removal-only migration is not defined | `spec:product.design_records.authoring_standards.task_authoring` | `### Metadata schema`; `### Update` | migration | W017 |
| C-013 | W17 must amend the W16-written Work Item rule section without weakening parent coordination boundaries | `spec:product.design_records.authoring_standards.work_item_authoring` | `### Kind-specific authoring rules`; `Parent coordination boundary`; `## Related specs` | shared-writer | shared-serialized |
| C-014 | W17 must amend W16-written Task metadata and relation wording without weakening `task_type` or single-responsibility rules | `spec:product.design_records.authoring_standards.task_authoring` | `### Metadata schema`; `### Kind-specific authoring rules`; `### Create`; `### Update`; `## Related specs` | shared-writer | shared-serialized |
| C-015 | Work Item purpose remains Requirement-only in the artifact selector | `spec:product.design_records.authoring_standards.artifact_boundary` | `## Covered artifact kinds`; `## Distinguishing adjacent artifacts` | semantic | W017 |
| C-016 | W17 must preserve the W16 decision-workflow projection while changing Work Item provenance wording | `spec:product.design_records.authoring_standards.artifact_boundary` | `## Distinguishing adjacent artifacts`; `## Decision workflow projection`; `## Related specs` | shared-writer | shared-serialized |
| C-017 | Work Item responsibility remains Requirement-only in the canonical matrix | `spec:product.design_records.artifact_model.artifact_responsibility_matrix` | `## Artifact responsibility matrix` | semantic | W017 |
| C-018 | Canonical matrix lacks the PRODUCT-semantics versus DRMCP-projection source-relation boundary | `spec:product.design_records.artifact_model.artifact_responsibility_matrix` | `## Artifact responsibility matrix`; `## Canonical design-state boundary` | semantic | W017 |
| C-019 | W17 must preserve W16 Task design-state wording while adding source-relation ownership | `spec:product.design_records.artifact_model.artifact_responsibility_matrix` | `## Artifact responsibility matrix`; `## Canonical design-state boundary`; `## Related specs` | shared-writer | shared-serialized |
| C-020 | Traceability overview declares the legacy reciprocal workflow field set | `spec:product.design_records.traceability` | `## Current contract`; `## Traceability model` | semantic | W017 |
| C-021 | Traceability overview lacks the direct Requirement reverse set and transitive traversal distinction | `spec:product.design_records.traceability` | `## Current contract`; `## Traceability model`; `## Non-goals` | reverse-relation | W017 |
| C-022 | Trace metadata schema declares legacy reciprocal fields and mismatch rules | `spec:product.design_records.traceability.metadata_schema` | `## Workflow relation metadata` | metadata-shape | W017 |
| C-023 | Trace metadata schema lacks unordered, duplicate-free, non-empty Work Item `source_refs` semantics | `spec:product.design_records.traceability.metadata_schema` | `## Workflow relation metadata`; `## Metadata boundary` | validation | W017 |
| C-024 | Trace metadata schema lacks staged repository migration and same-record old/new exclusion | `spec:product.design_records.traceability.metadata_schema` | `## Workflow relation metadata`; `## Metadata boundary` | migration | W017 |
| C-025 | Artifact refs declares the legacy workflow relation identity table | `spec:product.design_records.traceability.artifact_refs` | `## Workflow relation identity` | metadata-shape | W017 |
| C-026 | Resolve and validation defines reciprocity and Task Requirement checks instead of source-relation validation | `spec:product.design_records.traceability.resolve_and_validation` | `## Workflow relation validation`; `## Resolve and validation boundary` | validation | W017 |
| C-027 | DRMCP metadata parsing and writer schemas must consume the accepted field transition | `spec:product.design_records.traceability.metadata_schema` | `## Workflow relation metadata`; `## Metadata boundary` | downstream-drmcp | downstream-DRMCP |
| C-028 | DRMCP must derive direct Requirement reverse sets separately from transitive graph traversal | `spec:product.design_records.traceability` | `## Current contract`; `## Traceability model` | downstream-drmcp | downstream-DRMCP |
| C-029 | DRMCP must implement Task-owner resolution for provenance-cycle analysis without PRODUCT fixing the mechanism | `spec:product.design_records.traceability.resolve_and_validation` | `## Workflow relation validation`; `## Excluded implementation behavior`; `## Resolve and validation boundary` | downstream-drmcp | downstream-DRMCP |
| C-030 | DRMCP diagnostics, indexing, reverse lookup, traversal, and projection must follow PRODUCT semantics | `spec:product.design_records.artifact_model.artifact_responsibility_matrix` | `## Artifact responsibility matrix`; `## Canonical design-state boundary` | downstream-drmcp | downstream-DRMCP |
| C-031 | Workflow bootstrap Tasks must remain outside existing-record migration action | `PRODUCT-TASK-SPEC-017-01` | `### Bootstrap disposition` | workflow-bootstrap-exclusion | W017 |

### Conflict contract and routing

| ID | current wording or contract | conflicting authority | routing | likely ADR route | likely synchronization target |
|---|---|---|---|---|---|
| C-001 | Requirement text says Work Item and Task provenance both use `source_refs`. | D-001 and D-010 require `source_refs` only on Work Items. | repository-resolvable correction | `candidate` through D-001 | Correct the named REQ-006 sections before canonical closure. |
| C-002 | Requirement metadata requires persisted `work_items` and describes it as the resolution relation. | D-011 and D-012 require a derived direct reverse set. | repository-resolvable correction | `candidate` through D-011 | Replace persisted reverse ownership in `requirement_authoring`. |
| C-003 | No removal precondition exists for Requirement `work_items`. | D-015 requires exact unordered set equality before removal. | repository-resolvable correction | `candidate` through D-015 | Add the migration precondition to `requirement_authoring`. |
| C-004 | Work Item metadata, create, update, and references require one Requirement ID. | D-001 and D-006 require non-empty generic canonical refs. | repository-resolvable correction | `candidate` through D-001 | Replace the field contract in `work_item_authoring`. |
| C-005 | Requirement reciprocity and child Task Requirement equality are PRODUCT invariants. | D-010 through D-012 remove Task Requirement provenance and persisted Requirement reverse membership. | repository-resolvable correction | `candidate` through D-011 | Retain only Work Item `tasks` and Task `work_item` membership invariants. |
| C-006 | No Work Item `source_refs` field, minimum cardinality, or active-ref boundary exists. | D-001 and D-006 define the required field and allowed classes. | repository-resolvable correction | `candidate` through D-001 | Add field shape and reference policy to `work_item_authoring`. |
| C-007 | Work Items have exactly one Requirement source. No materiality rule exists. | D-007 requires all direct material sources and excludes incidental or merely transitive ancestors. | repository-resolvable correction | `not_required` | Add source-selection rules to `work_item_authoring`. |
| C-008 | Parent coordination may list child Work Items, but no direct provenance rule exists. | D-008 requires the exact source Task. D-009 forbids automatic owner Work Item inclusion. | repository-resolvable correction | `candidate` through D-008 | Add Task-decomposition provenance without adding parent or child metadata. |
| C-009 | No record-atomic conversion exists. | D-013 and D-016 require one-element conversion, same-update deletion, and no inferred sources. | repository-resolvable correction | `candidate` through D-013 and D-016 | Add Work Item migration rules. |
| C-010 | Task metadata, create, update, and references require `source_requirement`. | D-001, D-010, and D-014 require no Task source field. | repository-resolvable correction | `candidate` through D-001 and D-014 | Remove Task `source_requirement` while preserving `work_item`. |
| C-011 | Task Requirement provenance must match the parent Work Item Requirement. | D-010 makes Task provenance reachable only through `work_item`. | repository-resolvable correction | `not_required` | Remove the source-matching invariant. Preserve membership and Task-type contracts. |
| C-012 | No removal-only Task migration exists. | D-014 and D-016 require atomic removal with no replacement field. | repository-resolvable correction | `candidate` through D-014 and D-016 | Add Task migration rules. |
| C-013 | W16 added parent coordination boundaries in the same authoring section W17 must amend. | D-007 through D-009 define provenance. D-019 fixes writer order and preservation. | repository-resolvable correction | `not_required` | Add provenance rules after W16 wording. Do not reopen Task-type decisions. |
| C-014 | W16 added required `task_type`, type contracts, and single-responsibility rules around legacy source wording. | D-010 and D-014 remove only Task source provenance fields. D-019 requires preservation. | repository-resolvable correction | `not_required` | Make relation-only edits around the accepted W16 contract. |
| C-015 | Work Item purpose is described only as Requirement resolution. | D-007 through D-009 allow direct material sources including Tasks and other canonical refs. | repository-resolvable correction | `candidate` through D-008 | Generalize Work Item provenance purpose without changing artifact selection ownership. |
| C-016 | W16 decision workflow projection is accepted and shares the file. | D-019 requires W17 to preserve W16 semantics while updating relation wording. | repository-resolvable correction | `not_required` | Keep the complete W16 decision flow unchanged. |
| C-017 | The Work Item row owns Requirement resolution flow. | D-007 through D-009 require generic material source provenance and Task decomposition. | repository-resolvable correction | `candidate` through D-008 | Generalize the Work Item ownership row. |
| C-018 | No canonical row assigns source-relation semantics to PRODUCT and projection mechanics to DRMCP. | D-018 fixes that ownership split. | repository-resolvable correction | `not_required` | Add the ownership boundary to the matrix. |
| C-019 | W16 added Task workflow-state and canonical design-state boundaries in the same file. | D-019 requires preservation during W17 source-relation edits. | repository-resolvable correction | `not_required` | Add relation ownership without changing Task, ADR, or spec design-state ownership. |
| C-020 | Workflow relations are `requirement.work_items`, `work_item.source_requirement`, `work_item.tasks`, `task.work_item`, `task.source_requirement`, and `task.depends_on`. | D-001 and D-010 through D-012 replace only the provenance and reverse relation subset. | repository-resolvable correction | `candidate` through D-001 and D-011 | Rewrite the active workflow relation set. |
| C-021 | No canonical direct reverse projection exists. Traversal and cycle checks are broadly excluded. | D-011 and D-012 define the direct set. D-017 and D-018 separate semantics from mechanics. | repository-resolvable correction | `candidate` through D-011 and D-017 | Define direct reverse semantics and leave traversal algorithms to DRMCP. |
| C-022 | Metadata tables and invalid conditions encode reciprocal Requirement and Task source fields. | D-001 and D-010 through D-012 define forward Work Item provenance and persisted membership only. | repository-resolvable correction | `candidate` through D-001 and D-011 | Replace the workflow metadata table and mismatch rules. |
| C-023 | No set semantics, minimum cardinality, duplicate rule, self-reference rule, or reverse set rule exists. | D-001 through D-005 and D-011 through D-012 define those conditions. | repository-resolvable correction | `candidate` through D-001 and D-011 | Add normalized semantic rules without DRMCP diagnostic vocabulary. |
| C-024 | No staged migration or same-record mixed-field prohibition exists. | D-013 through D-016 define staged repository migration and atomic record transitions. | repository-resolvable correction | `candidate` through D-013 through D-016 | Add the migration-state boundary to `metadata_schema`. |
| C-025 | Workflow identity still names Requirement `work_items` and both `source_requirement` fields. | D-001, D-010, and D-012 replace those identities. | repository-resolvable correction | `candidate` through D-001 and D-011 | Keep active ref classes. Replace only the workflow relation table. |
| C-026 | Validation checks reciprocity and Task Requirement equality. It excludes cycle detection from PRODUCT traceability. | D-003 through D-005 require source-ref validation. D-017 makes semantic provenance cycles invalid. | repository-resolvable correction | `candidate` through D-017 | Replace legacy checks and define PRODUCT-owned invalid conditions. |
| C-027 | PRODUCT delegates parsing, writer behavior, and migration implementation to DRMCP while the accepted field set changes. | D-001 and D-013 through D-016 require parser and migration support downstream. | downstream DRMCP routing only | Follow T05 ADR results | Route field parsing, writer, and migration implementation to a later DRMCP Work Item. Current DRMCP state was not inspected. |
| C-028 | The accepted reverse view must be derived without reciprocal Requirement storage. | D-011 and D-012 require direct reverse lookup and separate transitive traversal. | downstream DRMCP routing only | Follow T05 ADR results | Route indexing, reverse lookup, traversal, and projection updates downstream. Current DRMCP state was not inspected. |
| C-029 | PRODUCT intentionally leaves Task-owner resolution mechanics unspecified. | D-017 defines invalidity and Task-to-owner normalization semantics. D-018 leaves the mechanism to DRMCP. | downstream DRMCP routing only | Follow T05 ADR results | Route owner resolution and graph-analysis implementation downstream. Current DRMCP state was not inspected. |
| C-030 | Diagnostics and response behavior are outside PRODUCT. | D-018 assigns diagnostics, response schemas, and projections to DRMCP. | downstream DRMCP routing only | `not_required` for the ownership split | Route exact diagnostics and tool projection contracts downstream after PRODUCT synchronization. |
| C-031 | A repository migration could accidentally include W16 and W17 bootstrap Tasks because those records still use the old metadata. | `BOOTSTRAP-001` in T01 excludes those workflow Tasks from migration action. | repository-resolvable correction | `not_required` | Preserve the exclusion in migration execution scope and Evidence. Do not create a generic canonical exception. |

### No-conflict findings

| ID | exact canonical ref or workflow record | assessed section | result |
|---|---|---|---|
| N-001 | `spec:product.design_records.authoring_standards.task_authoring` | `work_item` metadata and membership wording | Task `work_item` already represents the owning Work Item and must remain. |
| N-002 | `spec:product.design_records.authoring_standards.work_item_authoring` | `tasks` metadata and membership wording | Work Item `tasks` already represents owned Tasks and must remain. |
| N-003 | `spec:product.design_records.traceability.artifact_refs` | `## Active reference classes` | The active canonical classes already cover D-006, including Task IDs and resolver-supported legacy issued IDs. |
| N-004 | `spec:product.design_records.traceability` | PRODUCT and DRMCP ownership paragraph | The high-level ownership split already aligns with D-018. Relation-specific wording still needs C-021 correction. |
| N-005 | `spec:product.design_records.authoring_standards.task_authoring` | Task type, single responsibility, adjacent boundaries, and decision workflow Evidence | W16 T07 wording is accepted and is not a W17 decision target. |
| N-006 | W16 and W17 workflow Tasks | Bootstrap metadata and migration boundary | No source-relation semantic conflict exists. C-031 records the workflow-only migration exclusion. |
| N-007 | `PRODUCT-TASK-SPEC-016-07` | `### W017 boundary preservation` | The shared writer dependency is complete. W17 may now write relation wording. |
| N-008 | `PRODUCT-WORK-SPEC-017` | `## Evidence` | T02 lifecycle text is stale. The issue is outside this Task write boundary and is not a source-relation conflict. |

### T04 routing

No conflict currently requires a new user decision.

T04 must not reopen D-001 through D-019.
The T04 user-decision input set is empty under the current inventory.
T04 remains unexecuted and its lifecycle is unchanged by this Task.

### Migration findings

- Work Item migration is an atomic one-element conversion with no inferred extra sources.
- Task migration removes `source_requirement` and adds no replacement source field.
- Requirement migration requires exact direct reverse-set equality before removing `work_items`.
- Repository-wide staged migration is allowed.
- Each migrated record must use either old or new provenance metadata, never both.
- W16 and W17 workflow bootstrap Tasks require no migration action.
- This Task does not execute migration or inspect the full existing-record population.

### W16 shared-writer preservation

- Preserve the required scalar `task_type` and the closed nine-value set.
- Preserve every Task-type outcome, completion judgment, and prohibited overlap.
- Preserve mandatory single-responsibility and common-section alignment rules.
- Preserve parent coordination boundaries in `work_item_authoring`.
- Preserve Task checkpoint, ADR routing, and Specification synchronization boundaries in `artifact_boundary`.
- Preserve Task, ADR, and Specification design-state ownership in the responsibility matrix.
- Change only source-relation, reverse-relation, validation, migration, and ownership wording owned by W17.

### Verification result

- Only this Task is within the writable boundary.
- Every C-001 through C-031 entry has an exact canonical authority ref and affected section.
- Every product conflict traces to D-001 through D-019 or an accepted Requirement clause.
- W16-owned Task-type semantics are preserved and not re-decided.
- No conflict disposition, ADR route, Specification change, Requirement correction, migration action, or DRMCP implementation was performed.
- No repository-wide traversal or repository-wide clean-status inference was used.
- Bootstrap authority remains `PRODUCT-TASK-SPEC-016-01` Evidence, `BOOTSTRAP-001`.
- This Task remains authored under the current workflow bootstrap `source_requirement` and `work_item` metadata contract.
