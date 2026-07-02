# PRODUCT-TASK-SPEC-021-08: Route bootstrap decisions to ADR boundaries

- **id**: PRODUCT-TASK-SPEC-021-08
- **status**: done
- **date**: 2026-07-02
- **work_item**: PRODUCT-WORK-SPEC-021
- **task_type**: decision
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-021-07
- **outputs**:
  - PRODUCT-TASK-SPEC-021-08

## Goal

Decide the complete ADR route and coherent ADR boundaries for T06 decisions D-001 through D-005.

## Work

- Read T06 without changing its completed decision ledger.
- Read PRODUCT-INV-SPEC-009 as historical impact Evidence.
- Check accepted ADR coverage, including PRODUCT-ADR-SPEC-005, PRODUCT-ADR-SPEC-011, and PRODUCT-ADR-SPEC-015 through PRODUCT-ADR-SPEC-017.
- Classify every T06 decision as `required`, `covered`, `not_required`, or `blocked`.
- Select `create`, `amend`, `reuse`, or `supersede` for every routed ADR boundary.
- Partition decisions only when alternatives, rationale, ownership, and consequences share one boundary.
- Identify exact PRODUCT and TRV profile or Specification targets.
- Define exact ADR-authoring Task boundaries for T09 when authoring is required.

This Task must not:

- author or modify an ADR, Specification, Work Item, namespace profile, or skill;
- change T06 decisions;
- create downstream Tasks;
- perform review, synchronization, implementation, stage, or commit work.

## Done condition

- D-001 through D-005 each have one terminal ADR routing result.
- Every `required` decision belongs to one coherent ADR boundary.
- Every `covered` decision names an accepted non-superseded ADR.
- Every `not_required` decision records one reason and exact canonical target.
- Every `blocked` decision records one exact missing input.
- Create, amend, reuse, and supersede dispositions are explicit.
- Exact ADR-authoring boundaries and writers are sufficient for T09 graph coordination.

## Verification

- Confirm exactly five decision rows.
- Confirm every reused ADR is accepted and not superseded.
- Confirm no ADR body or canonical Specification content changed.
- Confirm no boundary combines independently changeable choices.
- Confirm T06 remains unchanged.
- Confirm T09 can materialize exact authoring owners without new design judgment.

## Evidence

### Routing authority

- Completed decision ledger: `PRODUCT-TASK-SPEC-021-06`, D-001 through D-005.
- Historical impact Investigation: `PRODUCT-INV-SPEC-009`.
- Active routing authority: `skills/design-convergence-workflow/adr-routing.md`.
- Active ADR authoring authority: `spec:product.design_records.authoring_standards.adr_authoring`.
- DRMCP is non-operational, so this Task used filesystem read and write operations.

### Existing ADR assessment

| ADR | status | routing treatment | result |
|---|---|---|---|
| `PRODUCT-ADR-SPEC-001` | accepted | reuse | Covers PRODUCT semantic ownership and prohibits executable implementation in the PRODUCT validator Specification area. |
| `PRODUCT-ADR-SPEC-005` | accepted | reuse | Covers independent completion boundaries, Work Item decomposition, and optional Work Item execution tracking. |
| `PRODUCT-ADR-SPEC-009` | accepted | reuse | Covers reviewed design closure and excludes production implementation from design convergence. |
| `PRODUCT-ADR-SPEC-010` | accepted | reuse | Covers typed routing through ADR authoring, review, decomposition, and synchronization. |
| `PRODUCT-ADR-SPEC-011` | accepted | reuse | Covers normal Work Item continuation and split criteria. W021's immediate correction is workflow-record repair, not a new architecture decision. |
| `PRODUCT-ADR-SPEC-012` | accepted | reuse | Covers one integrated review per Work Item and independent review for separately closable boundaries. |
| `PRODUCT-ADR-SPEC-015` | accepted | reuse | Covers the PRODUCT-fixed semantic evaluation and failure contract. |
| `PRODUCT-ADR-SPEC-016` | accepted | amend | Its core standalone-versus-DRMCP decision remains valid, but its W021 implementation-delivery consequence is stale. |
| `PRODUCT-ADR-SPEC-017` | accepted | reuse | Covers two-point invocation and caller-owned human exception handling. |

No accepted ADR is superseded by D-001 through D-005.

### Required ADR boundaries

#### B-001 — Explicit TRV app-local ownership clarification

- **ADR**: `PRODUCT-ADR-SPEC-016`.
- **Disposition**: `amend`.
- **Materiality**: `non-material`.
- **Decision IDs**: D-003. D-002 supplies the resolved owner identifier but requires no separate ADR rationale.
- **Bounded question**: How is the already accepted standalone-versus-current-DRMCP ownership split expressed after the standalone application receives the `TRV` namespace?
- **Required amendment**:
  - preserve PRODUCT ownership of the semantic validator contract;
  - preserve separation from current DRMCP;
  - identify TRV as the app-local design and later implementation owner;
  - remove the stale consequence that W021 owns concrete implementation delivery;
  - state that W021 owns reviewed PRODUCT conceptual design only.
- **Why amend**: The selected architecture and rationale remain unchanged. The amendment clarifies the concrete app owner and removes one stale Work Item consequence without reversing the standalone decision.
- **Dependency**: None beyond completed T08 routing.
- **Canonical targets**: `spec:product.responsibility_boundary_validator`, `spec:product.brewprint.namespaces.app_namespaces`, `spec:trv`, and W021.

### Decision routing results

| decision | routing result | ADR boundary | disposition | existing ADR coverage | reason | canonical targets | blocker |
|---|---|---|---|---|---|---|---|
| D-001 | `not_required` | — | direct workflow correction | ADR-011 remains the general continuity authority. | The W021 rescope repairs an incorrectly expanded workflow artifact before downstream delivery. The correction history belongs in T06 and W021 Evidence, not in an ADR. | W021 | — |
| D-002 | `not_required` | — | direct projection | — | `TRV`, its formal name, and `trv/` are resolved registry and placement identifiers. No independent trade-off needs ADR history. | `spec:product.brewprint.namespaces.app_namespaces`; `spec:product.brewprint.namespaces.domain_catalog`; `spec:product.brewprint.layout`; `spec:trv` | — |
| D-003 | `required` | B-001 / `PRODUCT-ADR-SPEC-016` | `amend` | ADR-001, ADR-015, ADR-016, and ADR-017 preserve the semantic and workflow contract. | ADR-016 still assigns concrete implementation delivery to W021. A non-material amendment must identify TRV and remove the stale consequence. | ADR-016; `spec:product.responsibility_boundary_validator`; `spec:trv`; W021 | — |
| D-004 | `covered` | existing coverage | `reuse` | ADR-005, ADR-009, ADR-011, and ADR-012. | The accepted rules already require a separate Work Item and integrated review for an independently closable app-local design boundary. The exact ID and title are decomposition inputs. | `TRV-WORK-SPEC-001` | — |
| D-005 | `covered` | existing coverage | `reuse` | ADR-005, ADR-009, ADR-010, and ADR-012. | Existing workflow authority already covers ADR routing, serialized authoring, integrated review, decomposition, closure synchronization, and omission of `work_item_execution` when the parent does not wait for child completion. | W021 Task graph and closure state | — |

### ADR authoring handoff

T09 must materialize one ADR amendment Task.

#### Authoring boundary A-001

- **Primary outcome**: Amend `PRODUCT-ADR-SPEC-016` for B-001.
- **Writer**: one PRODUCT ADR author independent of the routing decision.
- **Dependency**: T09 only.
- **Inputs**: T06 D-002 and D-003, this routing result, ADR-001, ADR-015, ADR-016, and ADR-017.
- **Completion judgment**: The core standalone separation remains unchanged, TRV app-local ownership is explicit, and the stale W021 implementation consequence is removed.
- **Prohibited scope**: New ADR creation, Specification edits, namespace-profile edits, review, or synchronization.

T09 must place A-001 before T10.

### Completion verification result

- Decision rows: 5.
- Required ADR boundaries: 1.
- New ADRs: 0.
- Existing ADR amendments: 1.
- Covered decisions: 2.
- Not-required decisions: 2.
- Blocked decisions: 0.
- Supersessions: 0.
- T06 remained unchanged.
- No ADR, Specification, namespace profile, Work Item, graph, review, synchronization, implementation, stage, or commit operation was performed.
