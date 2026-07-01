# PRODUCT-ADR-SPEC-011: Preserve Requirement and Work Item identity through completion-boundary continuity

- **status**: accepted
- **date**: 2026-07-01
- **depends_on**:
  - PRODUCT-ADR-SPEC-005
  - PRODUCT-ADR-SPEC-009
  - PRODUCT-ADR-SPEC-010
- **supersedes**: []
- **migrated_to_spec**: null

## Context

Design convergence may reveal that an originating Requirement or Work Item no longer matches the accepted decision.

Unrestricted in-place amendment can erase a materially different request or completion boundary.
Mechanical creation of new records can fragment one coherent resolution flow.

The workflow needs one shared identity rule for Requirement amendment and Work Item continuation.

## Decision

Preserve an existing Requirement identity when:

- the motivating problem and purpose remain the same;
- the change clarifies or adjusts a boundary within the same Required Outcome;
- the current Work Item completion meaning remains coherent.

Create a new Requirement when the motivating problem, Required Outcome, or material scope changes enough for independent acceptance or rejection.

Create a follow-up Requirement when the original remains completable and an adjacent additional need appears.
Reconsider the decision when the Requirement remains valid and the decision departed from it.

Preserve an existing Work Item identity when:

- it still resolves the same direct material sources;
- its Goal and Completion Conditions retain their meaning;
- only Tasks, dependencies, or ordering change inside the same delivery boundary.

Create or split another Work Item when:

- a new Requirement is introduced;
- the new scope has an independent completion judgment;
- ownership, release timing, or primary deliverables separate;
- an adjacent design topic can proceed independently.

Task count alone does not determine Work Item identity.

## Rationale

Requirement identity represents one stable need and acceptance boundary.
Work Item identity represents one stable resolution and completion boundary.

Using those identities prevents both historical rewriting and artificial fragmentation.
The same principle applies at both levels: preserve the record while its acceptance or completion meaning remains coherent.

## Rejected alternatives

| alternative | rejection reason |
|---|---|
| Amend every existing Requirement. | Materially different requests would lose independent identity and acceptance history. |
| Create a new Requirement for every clarification. | One stable need would fragment into unnecessary records. |
| Split a Work Item whenever Tasks are added. | Graph size does not create an independent completion boundary. |
| Keep all adjacent scope in one Work Item. | Independently releasable or completable topics would share a misleading lifecycle. |
| Always change the originating artifact to match the decision. | The decision may be the element that conflicts with valid authority. |

## Consequences

- `spec:product.design_records.authoring_standards.requirement_authoring` must define the amendment, replacement, and follow-up boundary.
- `spec:product.design_records.authoring_standards.work_item_authoring` must define continuation and split criteria.
- Reconciliation decisions must classify the originating-artifact disposition before authoring.
- Coordination may add Tasks without splitting the Work Item when completion identity remains unchanged.
- A new Requirement normally requires a distinct resolution Work Item.

## Evidence

- `PRODUCT-TASK-SPEC-018-01`: D-012 and D-013.
- `PRODUCT-TASK-SPEC-018-02`: B-003 routing boundary.
