# PRODUCT-TASK-SPEC-017-05: Classify ADR routing

- **id**: PRODUCT-TASK-SPEC-017-05
- **status**: done
- **date**: 2026-07-01
- **work_item**: PRODUCT-WORK-SPEC-017
- **source_requirement**: PRODUCT-REQ-SPEC-006
- **estimate**: 0.25d
- **depends_on**:
  - PRODUCT-TASK-SPEC-017-02
  - PRODUCT-TASK-SPEC-017-04
- **outputs**:
  - PRODUCT-TASK-SPEC-017-05

## Goal

Classify the ADR disposition for every decided REQ-006 and conflict decision.

Produce one complete routing judgment without authoring ADR content.

## Work

- Read the ADR-routing workflow authority for this phase.
- Check existing ADR coverage for each decided item.
- Classify each item as no ADR, new ADR, amendment, or supersession.
- Group only decisions that form one coherent durable choice.
- Record exact ADR authoring inputs for T06.

## Done condition

- Every decided item has one supported ADR disposition.
- Existing ADR coverage and supersession needs are explicit.
- T06 has a complete bounded authoring set.
- No ADR file is created or changed.

## Verification

- Confirm every T02 and T04 decided item appears once in the routing result.
- Confirm editorial Specification details are not promoted to ADRs.
- Confirm durable provenance or migration trade-offs are not omitted.
- Confirm no Specification or ADR file changed.

## Evidence

### Routing authority

- `skills/design-decision-workflow/adr-routing.md` defines the routing criteria.
- `PRODUCT-TASK-SPEC-017-02` supplies D-001 through D-019.
- `PRODUCT-TASK-SPEC-017-04` supplies C-001 through C-026 and C-031 dispositions plus C-027 through C-030 downstream routes.
- `PRODUCT-ADR-SPEC-001` is accepted and owns the generic PRODUCT semantic ownership boundary.
- `PRODUCT-ADR-SPEC-004` through `PRODUCT-ADR-SPEC-006` constrain shared W16 wording but do not own REQ-006 source-relation choices.
- `V01-ADR-091` and `V01-ADR-092` are historical snapshots. Their legacy reciprocal relation decisions are not current canonical Specification authority.

### Routing summary

| result | items | action |
|---|---|---|
| New ADR required | D-001 through D-012 and D-017 | Author `PRODUCT-ADR-SPEC-007` for the canonical workflow source-relation model. |
| New ADR required | D-013 through D-016 | Author `PRODUCT-ADR-SPEC-008` for staged atomic migration from legacy workflow relation metadata. |
| Covered | D-018 | Reuse accepted `PRODUCT-ADR-SPEC-001`. No amendment is needed. |
| ADR not required | D-019 | Writer serialization and preservation are workflow facts derived from completed W16 T07. |
| ADR not required | C-001 through C-026 and C-031 | Each conflict disposition is a projection of routed D-items, accepted W16 wording, or workflow-only migration scope. |
| Downstream only | C-027 through C-030 | No PRODUCT ADR is added for concrete DRMCP parser, index, diagnostic, traversal, or projection mechanisms. |

No routing item is blocked.
No existing PRODUCT ADR requires amendment or supersession.

### T02 decision routing

| ID | routing result | ADR | reason | affected targets |
|---|---|---|---|---|
| D-001 | required | `PRODUCT-ADR-SPEC-007` | Work Item-only generic provenance and Task provenance through membership establish a durable workflow relation boundary. | REQ-006; Work Item and Task authoring; traceability. |
| D-002 | required | `PRODUCT-ADR-SPEC-007` | Unordered-set semantics are part of the chosen canonical relation model and constrain persistence and projection. | Trace metadata schema. |
| D-003 | required | `PRODUCT-ADR-SPEC-007` | Duplicate rejection is a durable validity and failure-handling choice. | Resolve and validation. |
| D-004 | required | `PRODUCT-ADR-SPEC-007` | Self-reference rejection is a durable graph-validity choice. | Resolve and validation. |
| D-005 | required | `PRODUCT-ADR-SPEC-007` | Mandatory resolution and rejection of future-candidate refs define the persisted provenance boundary. | Resolve and validation. |
| D-006 | required | `PRODUCT-ADR-SPEC-007` | Reusing every active canonical ref class rejects a narrower Work Item-only taxonomy. | Artifact refs; Work Item authoring. |
| D-007 | required | `PRODUCT-ADR-SPEC-007` | Direct material source selection and transitive omission define the meaning of a provenance edge. | Work Item authoring. |
| D-008 | required | `PRODUCT-ADR-SPEC-007` | Exact Task provenance for Task-created Work Items defines cross-Work-Item decomposition. | Work Item authoring. |
| D-009 | required | `PRODUCT-ADR-SPEC-007` | Prohibiting automatic owner Work Item copying rejects hidden hierarchy and transitive duplication. | Work Item authoring. |
| D-010 | required | `PRODUCT-ADR-SPEC-007` | Task source non-inheritance is inseparable from Work Item-only provenance. | Task authoring. |
| D-011 | required | `PRODUCT-ADR-SPEC-007` | Direct-only Requirement reverse relations reject persisted reciprocity and transitive membership semantics. | Requirement authoring; traceability. |
| D-012 | required | `PRODUCT-ADR-SPEC-007` | The unordered duplicate-free direct reverse projection completes the selected relation model. | Requirement authoring; trace metadata schema. |
| D-013 | required | `PRODUCT-ADR-SPEC-008` | One-element Work Item conversion with no inferred sources is a durable migration choice. | Work Item migration contract. |
| D-014 | required | `PRODUCT-ADR-SPEC-008` | Task removal-only migration rejects replacement source metadata and preserves membership. | Task migration contract. |
| D-015 | required | `PRODUCT-ADR-SPEC-008` | Exact-match verification before Requirement reverse-list removal defines migration failure handling. | Requirement migration contract. |
| D-016 | required | `PRODUCT-ADR-SPEC-008` | Repository-wide staging with atomic per-record transitions rejects dual-field coexistence and big-bang-only migration. | Trace metadata migration boundary. |
| D-017 | required | `PRODUCT-ADR-SPEC-007` | Semantic provenance-cycle invalidity and Task-owner normalization establish a durable graph-validity boundary. | Resolve and validation. |
| D-018 | covered | `PRODUCT-ADR-SPEC-001` | PRODUCT semantic ownership and app-local implementation ownership are already accepted. REQ-006 applies that boundary without changing it. | Traceability; responsibility matrix. |
| D-019 | not_required | — | Completed writer order and preservation follow W16 T07 and contain no durable alternative or architecture choice. | Shared authoring Specifications. |

### T04 conflict routing

| conflict set | routing result | reason |
|---|---|---|
| C-001 through C-012 | not_required | Requirement, Work Item, and Task corrections are normative projections of D-001 through D-016. |
| C-013, C-014, C-016, C-019 | not_required | Shared-writer preservation is fixed by D-019 and accepted W16 Specifications. |
| C-015, C-017 | not_required | Work Item wording changes project D-007 through D-009 and add no independent trade-off. |
| C-018 | covered | The PRODUCT versus DRMCP ownership split is covered by `PRODUCT-ADR-SPEC-001`. |
| C-020 through C-026 | not_required | Traceability table, validation, and migration wording project the routed D-items into canonical Specifications. |
| C-031 | not_required | Bootstrap exclusion is execution scope for this workflow, not a durable canonical metadata exception. |
| C-027 through C-030 | downstream only | Concrete DRMCP mechanism choices remain outside PRODUCT and require later app-local design work when implementation begins. |

### Existing ADR coverage and supersession

| ADR | disposition | result |
|---|---|---|
| `PRODUCT-ADR-SPEC-001` | covered and retained | Covers D-018. `PRODUCT-ADR-SPEC-007` must depend on it rather than restating the generic ownership rationale. |
| `PRODUCT-ADR-SPEC-002` | unrelated | Portable package generation does not cover workflow provenance. |
| `PRODUCT-ADR-SPEC-003` | unrelated | Package operational checks do not cover workflow provenance. |
| `PRODUCT-ADR-SPEC-004` | retained constraint | Task taxonomy remains valid and must be preserved during source-field removal. |
| `PRODUCT-ADR-SPEC-005` | retained constraint | Single-responsibility and independent completion boundaries remain valid. |
| `PRODUCT-ADR-SPEC-006` | retained constraint | Decision-workflow checkpoint ownership remains valid. |
| `V01-ADR-091` | historical partial conflict | Its Work Item and Task ownership model remains useful, but its explicit `source_requirement` relation set is obsolete. Do not supersede the whole omnibus snapshot. Cite the changed relation boundary in ADR-007 context. |
| `V01-ADR-092` | historical partial conflict | Its reciprocal workflow validation contract is obsolete, but its record retrieval and ID-as-ref decisions remain historical context. Do not supersede the whole omnibus snapshot. Cite the changed relation boundary in ADR-007 context. |

No in-place amendment is appropriate.
No whole-record supersession is appropriate because both V01 ADRs contain unrelated decisions that remain compatible.
The new ADRs must record that current Specifications replace the conflicting legacy relation sections.

### T06 authoring set

#### PRODUCT-ADR-SPEC-007

Title:

`Use Work Item source refs as canonical workflow provenance`

Metadata inputs:

| field | value |
|---|---|
| status | `accepted` only after the repository-required authoring and review process; T06 must use the lifecycle state allowed by its exact contract. |
| date | `2026-07-01` |
| depends_on | `PRODUCT-ADR-SPEC-001` |
| supersedes | empty |
| migrated_to_spec | `null` |

Coherent decision boundary:

- Work Items persist required non-empty generic `source_refs`.
- Tasks persist no source field and expose provenance through `work_item`.
- `source_refs` is an unordered set of active canonical refs.
- Direct material sources are listed; incidental and merely transitive ancestors are omitted.
- A Task-created Work Item cites the exact source Task and does not automatically copy its owner Work Item.
- Requirement reverse membership is derived from direct Work Item refs and excludes transitive descendants.
- Duplicate, self, unresolved, and noncanonical source refs are invalid.
- Every semantic Work Item provenance cycle is invalid.
- Task refs normalize to the owning Work Item for cycle semantics; DRMCP owns the mechanism.
- Work Item `tasks`, Task `work_item`, and Task `depends_on` remain separate persisted membership or dependency relations.

Material rejected alternatives:

- Keep `source_requirement` and persisted Requirement `work_items`.
- Add `source_refs` to both Work Items and Tasks.
- Copy transitive ancestors or the source Task owner automatically.
- Treat transitive descendants as direct Requirement reverse membership.
- Permit unresolved future-candidate refs or semantic provenance cycles.
- Define a narrower Work Item-only reference taxonomy.

Primary Specification targets:

- `spec:product.design_records.authoring_standards.requirement_authoring`
- `spec:product.design_records.authoring_standards.work_item_authoring`
- `spec:product.design_records.authoring_standards.task_authoring`
- `spec:product.design_records.authoring_standards.artifact_boundary`
- `spec:product.design_records.artifact_model.artifact_responsibility_matrix`
- `spec:product.design_records.traceability`
- `spec:product.design_records.traceability.metadata_schema`
- `spec:product.design_records.traceability.artifact_refs`
- `spec:product.design_records.traceability.resolve_and_validation`

#### PRODUCT-ADR-SPEC-008

Title:

`Migrate legacy workflow source relations through atomic record transitions`

Metadata inputs:

| field | value |
|---|---|
| status | `accepted` only after the repository-required authoring and review process; T06 must use the lifecycle state allowed by its exact contract. |
| date | `2026-07-01` |
| depends_on | `PRODUCT-ADR-SPEC-007` |
| supersedes | empty |
| migrated_to_spec | `null` |

Coherent decision boundary:

- Repository-wide staged migration is allowed.
- Each record transitions atomically and never persists old and new provenance fields together.
- A Work Item converts `source_requirement` into a one-element `source_refs` list and infers no additional source.
- A Task removes `source_requirement` without replacement and preserves `work_item`.
- A Requirement removes `work_items` only after exact unordered, duplicate-free equality with the derived direct reverse set.
- Any Requirement mismatch blocks migration and is not silently repaired.

Material rejected alternatives:

- Require one repository-wide big-bang transaction.
- Permit a dual-field compatibility period within one record.
- Infer or append extra Work Item sources during conversion.
- Replace Task `source_requirement` with Task `source_refs`.
- Remove or repair Requirement reverse relations without exact-match verification.

Primary Specification targets:

- `spec:product.design_records.authoring_standards.requirement_authoring`
- `spec:product.design_records.authoring_standards.work_item_authoring`
- `spec:product.design_records.authoring_standards.task_authoring`
- `spec:product.design_records.traceability.metadata_schema`
- `spec:product.design_records.traceability.resolve_and_validation`

### Requirement correction writer advisory

C-001 requires correction of `PRODUCT-REQ-SPEC-006`.
The current graph does not explicitly assign a Requirement writer.
T05 does not amend the graph because coordination is outside this Task.
The writer gap does not block T06 ADR authoring.
It must be resolved before Requirement correction or T07 closure claims complete synchronization.

### Closure verification

- D-001 through D-019 each have exactly one resolved ADR route.
- C-001 through C-026 and C-031 each have exactly one resolved ADR route.
- C-027 through C-030 remain downstream-only impacts.
- Two coherent new ADR boundaries are identified.
- No accepted PRODUCT ADR is duplicated, amended, or superseded unnecessarily.
- Historical V01 partial conflicts are recorded without falsely superseding unrelated decisions.
- No ADR, Requirement, Specification, migration target, source, test, or fixture file changed.
- T05 is complete and T06 has an exact bounded authoring set.
- Bootstrap authority remains `PRODUCT-TASK-SPEC-016-01` Evidence, `BOOTSTRAP-001`.
- This Task remains authored under the current workflow bootstrap `source_requirement` and `work_item` metadata contract.
- No migration action is owned or performed by this workflow Task.
