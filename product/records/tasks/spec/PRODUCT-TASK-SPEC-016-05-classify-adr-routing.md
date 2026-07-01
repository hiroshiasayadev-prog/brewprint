# PRODUCT-TASK-SPEC-016-05: Classify ADR routing

- **id**: PRODUCT-TASK-SPEC-016-05
- **status**: done
- **date**: 2026-06-30
- **work_item**: PRODUCT-WORK-SPEC-016
- **source_requirement**: PRODUCT-REQ-SPEC-005
- **estimate**: 0.25d
- **depends_on**:
  - PRODUCT-TASK-SPEC-016-02
  - PRODUCT-TASK-SPEC-016-04
- **outputs**:
  - PRODUCT-TASK-SPEC-016-05

## Goal

Classify the ADR disposition for every decided REQ-005 and conflict decision.

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
- Confirm durable trade-offs are not omitted to reduce artifact count.
- Confirm no Specification or ADR file changed.

## Evidence

### Routing authority

- `PRODUCT-TASK-SPEC-016-02` is `done` with D-001 through D-018 decided.
- `PRODUCT-TASK-SPEC-016-04` is `done` with C-001 through C-014 decided.
- `skills/design-decision-workflow/adr-routing.md` defines the routing criteria.
- DRMCP is non-operational. Filesystem authoring is the required fallback.
- This Task owns routing only. No ADR or Specification file was changed.
- Existing Task migration remains outside PRODUCT-WORK-SPEC-016.

### Existing ADR coverage

| ADR | status | coverage result |
|---|---|---|
| `PRODUCT-ADR-SPEC-001` | accepted | Defines PRODUCT Specification semantic ownership and placement. It does not define Task types or workflow responsibility boundaries. |
| `PRODUCT-ADR-SPEC-002` | accepted | Defines portable package generation. It is unrelated to the REQ-005 Task contract. |
| `PRODUCT-ADR-SPEC-003` | accepted | Defines package-generation checks and warnings. It is unrelated to the REQ-005 Task contract. |
| `V01-ADR-091` | accepted snapshot | Provides compatible historical context for Work Item and Task separation and short-term Task granularity. It does not define `task_type`, the closed taxonomy, or the accepted type boundaries. |
| `V01-ADR-092` | accepted snapshot | Defines DRMCP workflow-record and relation support. It does not own PRODUCT Task responsibility semantics. |
| `V01-ADR-094` | accepted snapshot | Defines workflow status vocabulary. It does not own Task responsibility or completion boundaries. |

No accepted ADR fully covers the decided REQ-005 contract.
No accepted ADR requires amendment or supersession.

### Routing result

| item | routing result | ADR | reason | affected Specification targets |
|---|---|---|---|---|
| D-001 | required | `PRODUCT-ADR-SPEC-004` | The persisted field name is inseparable from the durable typed taxonomy and downstream parser contract. | `task_authoring` |
| D-002 | required | `PRODUCT-ADR-SPEC-004` | The closed value set constrains all future Task authoring and validation. | `task_authoring` |
| D-003 | required | `PRODUCT-ADR-SPEC-004` | The investigation type establishes one durable owned outcome and completion boundary. | `task_authoring` |
| D-004 | required | `PRODUCT-ADR-SPEC-004` | The decision type establishes durable workflow ownership. | `task_authoring` |
| D-005 | required | `PRODUCT-ADR-SPEC-004` | The authoring type establishes the canonical-writing boundary. | `task_authoring` |
| D-006 | required | `PRODUCT-ADR-SPEC-004` | The implementation type establishes a durable implementation acceptance boundary. | `task_authoring` |
| D-007 | required | `PRODUCT-ADR-SPEC-004` | The review type establishes independent verdict ownership. | `task_authoring` |
| D-008 | required | `PRODUCT-ADR-SPEC-004` | The correction type establishes named-finding repair ownership. | `task_authoring` |
| D-009 | required | `PRODUCT-ADR-SPEC-004` | The verification type establishes objective acceptance-gate ownership. | `task_authoring` |
| D-010 | required | `PRODUCT-ADR-SPEC-004` | The coordination type establishes parent graph-boundary ownership. | `task_authoring`, `work_item_authoring` |
| D-011 | required | `PRODUCT-ADR-SPEC-004` | The synchronization type establishes accepted-state propagation ownership. | `task_authoring` |
| D-012 | required | `PRODUCT-ADR-SPEC-006` | The decision-to-authoring handoff changes cross-artifact ownership and workflow sequencing. | `task_authoring`, `artifact_boundary`, `artifact_responsibility_matrix` |
| D-013 | required | `PRODUCT-ADR-SPEC-005` | Review and verification have distinct independent completion judgments. | `task_authoring` |
| D-014 | required | `PRODUCT-ADR-SPEC-005` | Correction and finding closure require separate owners and judgments. | `task_authoring` |
| D-015 | required | `PRODUCT-ADR-SPEC-005` | Coordination and synchronization have distinct graph and propagation responsibilities. | `task_authoring`, `work_item_authoring` |
| D-016 | required | `PRODUCT-ADR-SPEC-005` | The implementation judgment boundary constrains future design and implementation work. | `task_authoring` |
| D-017 | required | `PRODUCT-ADR-SPEC-005` | Cross-section alignment is the normative Task conformance test. | `task_authoring` |
| D-018 | required | `PRODUCT-ADR-SPEC-005` | The outcome-based multi-file rule rejects file count as a responsibility test. | `task_authoring` |
| C-001 | not_required | — | Adding `task_type` to create, update, and metadata tables is direct Specification projection of D-001 and D-002. | `task_authoring` |
| C-002 | not_required | — | A type-contract table is a Specification representation choice for D-003 through D-011. | `task_authoring` |
| C-003 | required | `PRODUCT-ADR-SPEC-005` | Making cohesion mandatory instead of advisory is a durable responsibility policy. | `task_authoring` |
| C-004 | not_required | — | Heading placement, table columns, and `TBD` handling are bounded authoring-format details derived from D-006 and D-017. | `task_authoring` |
| C-005 | required | `PRODUCT-ADR-SPEC-005` | The exact Goal, Work, Done condition, and Verification roles define durable conformance semantics. | `task_authoring` |
| C-006 | required | `PRODUCT-ADR-SPEC-006` | Conditional ADR routing replaces an absolute every-decision-to-ADR rule. | `task_authoring`, `artifact_boundary` |
| C-007 | required | `PRODUCT-ADR-SPEC-006` | The concise workflow rationale versus canonical ADR rationale boundary is durable artifact ownership. | `task_authoring`, `artifact_responsibility_matrix` |
| C-008 | required | `PRODUCT-ADR-SPEC-005` | Supporting checks and independent verification gates require a stable classification rule. | `task_authoring` |
| C-009 | required | `PRODUCT-ADR-SPEC-005` | Finding repair and finding closure require durable independent ownership. | `task_authoring` |
| C-010 | required | `PRODUCT-ADR-SPEC-005` | Graph judgment and mechanical propagation require durable stop conditions. | `task_authoring` |
| C-011 | required | `PRODUCT-ADR-SPEC-005` | Parent coordination must not duplicate child Work Item internals. | `work_item_authoring` |
| C-012 | required | `PRODUCT-ADR-SPEC-005` | Contract-affecting implementation choices must return to decision work. | `task_authoring` |
| C-013 | required | `PRODUCT-ADR-SPEC-006` | The canonical design-state owner must remain distinct from Task workflow state. | `artifact_responsibility_matrix` |
| C-014 | required | `PRODUCT-ADR-SPEC-006` | Task checkpoints and durable ADR recording require separate artifact-selection rules. | `artifact_boundary` |

### Routing refinement

The inventory routes were provisional.
T05 groups inseparable durable choices instead of preserving one preliminary route per row.

- D-001 and D-018 move from `not_required` to `required` because each is part of a coherent durable Task-contract choice.
- C-006 and C-007 move from `not_required` to `required` because they join D-012, C-013, and C-014 in one cross-artifact ownership decision.
- C-001, C-002, and C-004 remain direct Specification projection details and require no ADR.

### T06 authoring set

The next available PRODUCT SPEC ADR sequence is `004` through `006`.
These IDs are reserved by this routing result for PRODUCT-WORK-SPEC-016.

#### PRODUCT-ADR-SPEC-004

| input | value |
|---|---|
| Title | Define the closed typed Task responsibility taxonomy |
| Status | `accepted` |
| Date | `2026-07-01` |
| Depends on | `[]` |
| Supersedes | `[]` |
| Migrated to spec | `null` |
| Decision IDs | D-001 through D-011 |
| Context | REQ-005 requires one primary Task type and responsibility. Current authoring rules define neither a persisted type nor closed type semantics. |
| Adopted choice | Use required scalar `task_type`. Use the nine-value closed set. Define one owned outcome, one completion judgment, and prohibited overlaps for every value. |
| Material alternatives | Infer type from prose; allow free-form values; allow multiple primary types; use one generic Task class; define labels without type contracts. |
| Decisive rationale | A closed semantic taxonomy makes responsibility reviewable, authorable, and later validatable. Labels without owned outcomes do not prevent mixed responsibility. |
| Consequences | Update Task authoring and the coordination-facing Work Item boundary. Route parser, validation, and diagnostics to downstream DRMCP. Perform no existing-Task migration in W016. |
| Specification targets | `task_authoring`, `work_item_authoring` |

#### PRODUCT-ADR-SPEC-005

| input | value |
|---|---|
| Title | Enforce single-responsibility and independent Task completion boundaries |
| Status | `accepted` |
| Date | `2026-07-01` |
| Depends on | `PRODUCT-ADR-SPEC-004` |
| Supersedes | `[]` |
| Migrated to spec | `null` |
| Decision IDs | D-013 through D-018; C-003; C-005; C-008 through C-012 |
| Context | A type label does not enforce one responsibility. Current rules permit ambiguous section alignment, self-closure, graph judgment during synchronization, and contract judgment during implementation. |
| Adopted choice | Require one type-aligned outcome and completion judgment. Separate review from objective verification, correction from finding closure, and coordination from synchronization. Limit implementation to observable-equivalent local choices. Permit multiple files only under one outcome and acceptance boundary. Keep parent coordination free of child-internal Task graphs. |
| Material alternatives | Split by file count or effort alone; allow correction to close its own findings; require a separate verification Task for every command; allow synchronization to make graph decisions; allow implementation to resolve contract choices. |
| Decisive rationale | Independent outcomes require independent owners and judgments. Outcome-based cohesion preserves legitimate multi-file work without permitting mixed responsibility. |
| Consequences | Add mandatory cohesion and stop conditions to Task authoring. Clarify the parent Work Item boundary. Route validator and diagnostic behavior to downstream DRMCP. |
| Specification targets | `task_authoring`, `work_item_authoring` |

#### PRODUCT-ADR-SPEC-006

| input | value |
|---|---|
| Title | Separate decision workflow checkpoints from canonical design state |
| Status | `accepted` |
| Date | `2026-07-01` |
| Depends on | `PRODUCT-ADR-SPEC-004` |
| Supersedes | `[]` |
| Migrated to spec | `null` |
| Decision IDs | D-012; C-006; C-007; C-013; C-014 |
| Context | Current authoring rules imply that every decision immediately belongs in an ADR and that a Task cannot retain decision state. The accepted decision workflow requires resumable Task checkpoints before conditional ADR routing. |
| Adopted choice | A decision Task owns temporary and historical workflow state. It does not own canonical design state. Decision work ends when the choice, concise reason, and canonical target are fixed. Authoring then writes the canonical ADR or Specification. ADR routing is conditional. ADRs own durable rationale. Specifications own current normative state. |
| Material alternatives | Create an ADR for every checkpoint; prohibit all rationale in Tasks; treat Task ledgers as canonical design sources; let decision Tasks author ADR body content. |
| Decisive rationale | Workflow persistence and canonical design ownership solve different problems. Separating them preserves resumability without duplicating or weakening canonical artifacts. |
| Consequences | Revise Task authoring, artifact selection, and the responsibility matrix. Preserve Task history after ADR authoring and link the canonical ADR. |
| Specification targets | `task_authoring`, `artifact_boundary`, `artifact_responsibility_matrix` |

### Verification result

- D-001 through D-018 appear exactly once in the routing table.
- C-001 through C-014 appear exactly once in the routing table.
- Three coherent new ADRs are required.
- Three Specification projection details are `not_required`.
- No item is `covered`, `blocked`, amended, or superseded.
- No ADR or Specification file was created or changed.
- T06 has an exact bounded authoring set.
