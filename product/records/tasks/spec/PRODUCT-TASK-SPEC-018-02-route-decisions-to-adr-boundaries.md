# PRODUCT-TASK-SPEC-018-02: Route decisions to ADR boundaries

- **id**: PRODUCT-TASK-SPEC-018-02
- **status**: done
- **date**: 2026-07-01
- **work_item**: PRODUCT-WORK-SPEC-018
- **task_type**: decision
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-018-01
- **outputs**:
  - PRODUCT-TASK-SPEC-018-02

## Goal

Decide the complete ADR route and coherent ADR-boundary partition for D-001 through D-023.

## Work

- Read the completed decision ledger without changing it.
- Check accepted ADR coverage and supersession state.
- Classify every decision as `required`, `covered`, `not_required`, or `blocked`.
- Separate ADR routing from ADR authoring.
- Group decisions only when alternatives, rationale, responsibility, and consequences share one boundary.
- Select `create`, `amend`, `reuse`, or `supersede` for every routed boundary.
- Identify direct Specification and skill targets.
- Derive the exact next authoring Task boundaries.

No ADR, Specification, skill, review, lifecycle, production, stage, or commit work belongs to this Task.

## Done condition

- D-001 through D-023 each have one complete routing result.
- Required decisions map to coherent ADR boundaries.
- Covered decisions map to exact accepted, non-superseded ADRs.
- Not-required decisions have an exact reason and direct target.
- Blocked decisions have an exact blocker.
- ADR create, amend, reuse, and supersede dispositions are explicit.
- The next ADR-authoring Tasks have one responsibility and one completion judgment each.

## Verification

- Count exactly 23 decision rows.
- Confirm every accepted ADR used for coverage has `status: accepted` and no superseding ADR.
- Confirm no row uses one-ADR-per-decision as a default.
- Confirm no proposed ADR combines independently changeable workflow choices.
- Confirm `PRODUCT-ADR-SPEC-006` can preserve its core decision.
- Confirm the completed decision ledger remains unchanged.
- Confirm no ADR, Specification, or skill file changed.

## Evidence

### Routing authority

- Completed decisions: `PRODUCT-TASK-SPEC-018-01`, D-001 through D-023.
- Temporary source copy: `skills/design-convergence-workflow/decision-ledger.md`.
- Accepted Task taxonomy: `PRODUCT-ADR-SPEC-004`.
- Accepted responsibility boundaries: `PRODUCT-ADR-SPEC-005`.
- Accepted checkpoint and canonical-state boundary: `PRODUCT-ADR-SPEC-006`.
- Canonical Task and artifact guides were checked at their active `spec:` refs.

### Existing ADR disposition summary

| ADR | status | routing disposition | result |
|---|---|---|---|
| `PRODUCT-ADR-SPEC-004` | accepted | reuse | Covers Investigation ownership and the generic decision responsibility boundary. |
| `PRODUCT-ADR-SPEC-005` | accepted | reuse | Covers single responsibility, graph coordination, Task splitting, and default verification boundaries. |
| `PRODUCT-ADR-SPEC-006` | accepted | amend | Preserve the core checkpoint-versus-canonical-state decision. Remove stale decision-Task writeback and add routing-boundary ownership. |

No accepted ADR requires full or partial supersession.

### Proposed ADR boundaries

| boundary | ADR | decision IDs | disposition | coherent decision boundary |
|---|---|---|---|---|
| B-001 | `PRODUCT-ADR-SPEC-009` | D-001, D-004, D-005 | create | The successor owns one end-to-end workflow from topic intake through reviewed closure, excluding production implementation. |
| B-002 | `PRODUCT-ADR-SPEC-010` | D-006, D-011 | create | The workflow uses explicit responsibility phases and mismatch classes to select decision, coordination, or authoring routes. |
| B-003 | `PRODUCT-ADR-SPEC-011` | D-012, D-013 | create | Requirement and Work Item identity continues only while motivating outcome and completion identity remain coherent. |
| B-004 | `PRODUCT-ADR-SPEC-012` | D-015, D-016 | create | Shared writers serialize, and one integrated review evaluates the final combined Work Item state. |
| B-005 | `PRODUCT-ADR-SPEC-013` | D-020 | create | Correction and finding-closure Tasks materialize only after named findings exist. |
| B-006 | `PRODUCT-ADR-SPEC-014` | D-010, D-017, D-019, D-022, D-023 | create | Completed workflow records remain historical. New judgment re-enters convergence through new Tasks, and closure cannot rewrite completed evidence owners. |
| B-007 | `PRODUCT-ADR-SPEC-006` | D-019, D-021 | amend | Preserve conditional routing and canonical ownership while removing downstream decision-Task writeback and adding routing partition responsibility. |

### Decision-to-route table

| decision | routing result | ADR boundary | disposition | existing ADR | reason | affected Specification or skill targets | blocker |
|---|---|---|---|---|---|---|---|
| D-001 | required | B-001 / `PRODUCT-ADR-SPEC-009` | create | — | End-to-end workflow ownership is a durable repository governance boundary. | successor `SKILL.md`; `prompt_chappy.md` trigger text | — |
| D-002 | not_required | — | direct projection | — | Removing the replaced internal skill without a stub is a bounded repository transition, not a separate architecture choice. | successor activation Task; old skill directory removal; `prompt_chappy.md` | — |
| D-003 | not_required | — | direct projection | — | The successor path is a local naming and placement choice with no durable trade-off. | `skills/design-convergence-workflow/`; `prompt_chappy.md` | — |
| D-004 | required | B-001 / `PRODUCT-ADR-SPEC-009` | create | — | Starting before inventory fixes the workflow authority boundary and prevents pre-work from remaining outside governance. | successor `SKILL.md` workflow entry section | — |
| D-005 | required | B-001 / `PRODUCT-ADR-SPEC-009` | create | — | Reviewed closure and synchronization define the workflow terminal authority. | successor `SKILL.md`; `design-review-gate.md`; Work Item completion contract | — |
| D-006 | required | B-002 / `PRODUCT-ADR-SPEC-010` | create | `PRODUCT-ADR-SPEC-004`; `PRODUCT-ADR-SPEC-005` | Existing ADRs define generic Task types and cohesion. The exact convergence-phase architecture is a new durable workflow choice. | successor `SKILL.md`; Work Item Task flow guidance | — |
| D-007 | covered | existing coverage | reuse | `PRODUCT-ADR-SPEC-004` | The accepted Investigation Task contract already owns one Investigation record for one bounded research question. | successor `SKILL.md` impact-investigation phase | — |
| D-008 | covered | existing coverage | reuse | `PRODUCT-ADR-SPEC-004`; `PRODUCT-ADR-SPEC-005` | The accepted decision type owns selected outcomes. Existing responsibility rules prohibit graph and canonical authoring overlap. | successor conflict-reconciliation section | — |
| D-009 | covered | existing coverage | reuse | `PRODUCT-ADR-SPEC-005` | The accepted coordination boundary already owns graph, owner, dependency, blocker, and release changes. | successor coordination route; Work Item Task flow | — |
| D-010 | required | B-006 / `PRODUCT-ADR-SPEC-014` | create | `PRODUCT-ADR-SPEC-005`; `PRODUCT-ADR-SPEC-006` | Returning unresolved judgment without reopening completed records adds an append-only convergence rule. | successor `SKILL.md`; `interactive-decision-loop.md`; `design-review-gate.md`; `task_authoring` | — |
| D-011 | required | B-002 / `PRODUCT-ADR-SPEC-010` | create | `PRODUCT-ADR-SPEC-004`; `PRODUCT-ADR-SPEC-005` | The four mismatch classes select different responsibility routes and preserve a durable conflict-handling policy. | successor `SKILL.md`; conflict-investigation and reconciliation guidance | — |
| D-012 | required | B-003 / `PRODUCT-ADR-SPEC-011` | create | — | Requirement amendment, replacement, and follow-up splitting depend on stable need and independently acceptable outcome identity. | `requirement_authoring`; successor reconciliation guidance | — |
| D-013 | required | B-003 / `PRODUCT-ADR-SPEC-011` | create | — | Work Item continuation versus split depends on resolution identity and completion boundary, not Task count. | `work_item_authoring`; successor reconciliation guidance | — |
| D-014 | covered | existing coverage | reuse | `PRODUCT-ADR-SPEC-005` | Same-type, same-outcome extension and split-on-independent-responsibility are direct applications of the accepted single-responsibility decision. | `task_authoring`; successor coordination guidance | — |
| D-015 | required | B-004 / `PRODUCT-ADR-SPEC-012` | create | `PRODUCT-ADR-SPEC-005` | Deterministic shared-writer serialization and preservation obligations are new cross-Task authoring policy. | successor `SKILL.md`; `work_item_authoring`; shared-writer guidance | — |
| D-016 | required | B-004 / `PRODUCT-ADR-SPEC-012` | create | `PRODUCT-ADR-SPEC-005` | One final integrated review per Work Item is coupled to reviewing the final serialized state. | successor `design-review-gate.md`; `work_item_authoring` | — |
| D-017 | required | B-006 / `PRODUCT-ADR-SPEC-014` | create | `PRODUCT-ADR-SPEC-005` | Finding correction and renewed decision work require different return routes after review. | successor `design-review-gate.md`; `task_authoring` | — |
| D-018 | covered | existing coverage | reuse | `PRODUCT-ADR-SPEC-005` | The accepted review-versus-verification decision already makes separate verification conditional on independent gate ownership. | successor pre-authoring gate; `task_authoring` | — |
| D-019 | required | B-006 and B-007 | create and amend | `PRODUCT-ADR-SPEC-006` | The ADR core remains valid. Its post-authoring Task-reference consequence is stale, and completed-decision revision needs an append-only route. | `PRODUCT-ADR-SPEC-006`; `task_authoring`; `artifact_boundary`; `artifact_responsibility_matrix`; successor workflow files | — |
| D-020 | required | B-005 / `PRODUCT-ADR-SPEC-013` | create | `PRODUCT-ADR-SPEC-005` | Delayed Task materialization rejects speculative no-op contracts and preserves finding-specific writer boundaries. | successor `design-review-gate.md`; `work_item_authoring`; `task_authoring` | — |
| D-021 | required | B-007 / `PRODUCT-ADR-SPEC-006` | amend | `PRODUCT-ADR-SPEC-006` | Routing remains conditional, but its accepted scope must explicitly include ADR partitioning and create, amend, reuse, or supersede disposition. | `PRODUCT-ADR-SPEC-006`; successor `adr-routing.md`; `artifact_boundary`; `task_authoring` | — |
| D-022 | required | B-006 / `PRODUCT-ADR-SPEC-014` | create | `PRODUCT-ADR-SPEC-005`; `PRODUCT-ADR-SPEC-006` | New post-review judgment requires a new decision-authoring-review chain while completed Tasks remain historical Evidence. | successor `design-review-gate.md`; `task_authoring`; `work_item_authoring` | — |
| D-023 | required | B-006 / `PRODUCT-ADR-SPEC-014` | create | `PRODUCT-ADR-SPEC-005` | The exact closure write boundary extends generic synchronization rules by protecting completed decision, authoring, and review owners. | successor `design-review-gate.md`; `task_authoring`; `artifact_responsibility_matrix` | — |

### PRODUCT-ADR-SPEC-006 conflict disposition

The core decision remains valid:

- a decision Task owns resumable workflow checkpoints;
- an ADR owns durable rationale and supersession history;
- a Specification owns current normative state;
- ADR routing remains conditional and separate from canonical authoring.

The following current consequence is stale:

```text
After ADR authoring, the Task retains its workflow history and references the ADR.
```

D-019 prohibits downstream progress or ADR references from being written back into the completed decision Task.
The ADR authoring, review, and closure Tasks must own their own references and Evidence.

Disposition:

- preserve the ADR status as `accepted`;
- perform a clarification amendment;
- remove only the stale writeback consequence;
- add D-021 routing-boundary partition responsibility;
- do not supersede the ADR;
- do not rewrite the original checkpoint-versus-canonical-state choice.

### Direct Specification projections

| target | projected decisions |
|---|---|
| `spec:product.design_records.authoring_standards.requirement_authoring` | D-012 |
| `spec:product.design_records.authoring_standards.work_item_authoring` | D-013, D-015, D-016, D-020, D-022 |
| `spec:product.design_records.authoring_standards.task_authoring` | D-010, D-014, D-017 through D-023 |
| `spec:product.design_records.authoring_standards.artifact_boundary` | D-019, D-021, D-023 |
| `spec:product.design_records.artifact_model.artifact_responsibility_matrix` | D-019, D-022, D-023 |

Current target assessment:

- `task_authoring` contains an explicit stale post-authoring ADR-reference writeback and requires amendment.
- `artifact_boundary` is compatible with D-019 but should state downstream Evidence ownership when projecting D-021 and D-023.
- `artifact_responsibility_matrix` is compatible with checkpoint ownership but should state the completed-record write boundary.

### Direct skill projections

| target | projected decisions |
|---|---|
| successor `SKILL.md` | D-001, D-004 through D-020, D-022, D-023 |
| successor `interactive-decision-loop.md` | D-010, D-019, D-022 |
| successor `adr-routing.md` | D-021 and every route in this Task |
| successor `design-review-gate.md` | D-005, D-016, D-017, D-020, D-022, D-023 |
| `prompt_chappy.md` | D-002, D-003 after successor readiness |
| old skill directory | D-002 after successor activation |

### Exact next authoring Tasks

#### PRODUCT-TASK-SPEC-018-03

- Task type: `authoring`
- Responsibility: clarify `PRODUCT-ADR-SPEC-006` only.
- Inputs: B-007, D-019, D-021, accepted ADR authoring policy.
- Writable artifact: `PRODUCT-ADR-SPEC-006`.
- Done judgment: the core decision remains unchanged, stale decision-Task writeback is removed, and routing partition responsibility is explicit.
- Prohibited: new ADR creation, Specification edits, skill edits, review, lifecycle synchronization.

#### PRODUCT-TASK-SPEC-018-04

- Task type: `authoring`
- Responsibility: author the six new ADRs selected by B-001 through B-006.
- Inputs: T02 routing Evidence and the amended `PRODUCT-ADR-SPEC-006` from T03.
- Writable artifacts:
  - `PRODUCT-ADR-SPEC-009`
  - `PRODUCT-ADR-SPEC-010`
  - `PRODUCT-ADR-SPEC-011`
  - `PRODUCT-ADR-SPEC-012`
  - `PRODUCT-ADR-SPEC-013`
  - `PRODUCT-ADR-SPEC-014`
- Done judgment: every boundary is authored once, no ADR mixes independent choices, and no routed decision is omitted.
- Prohibited: Specification edits, skill edits, independent review, lifecycle synchronization.

The six files share one writer, one routed input set, and one later integrated review boundary.
One authoring Task therefore avoids mechanical one-Task-per-ADR fragmentation.
The existing-ADR amendment remains separate because it has a distinct historical-preservation judgment.

### Blocked items

None.

### Completion verification result

- Decision rows: 23.
- New ADR boundaries: 6.
- Existing ADR amendments: 1.
- Existing ADR reuses: 2 primary accepted ADRs plus preserved ADR-006 core.
- Supersessions: 0.
- Not-required decisions: D-002 and D-003.
- Blocked decisions: 0.
- ADR, Specification, skill, production, stage, and commit operations: not performed.
