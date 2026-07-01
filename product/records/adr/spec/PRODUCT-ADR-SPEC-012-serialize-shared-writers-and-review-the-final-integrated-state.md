# PRODUCT-ADR-SPEC-012: Serialize shared writers and review the final integrated state

- **status**: accepted
- **date**: 2026-07-01
- **depends_on**:
  - PRODUCT-ADR-SPEC-005
  - PRODUCT-ADR-SPEC-009
  - PRODUCT-ADR-SPEC-010
- **supersedes**: []
- **migrated_to_spec**: null

## Context

Several authoring Tasks may need to change the same artifact or section.
Parallel writes can overwrite accepted semantics or produce inconsistent combined text.

Independent partial reviews also cannot prove that the final combined state remains coherent.
The workflow needs one writer-order rule and one final review boundary.

## Decision

Serialize all Tasks that write the same artifact or section.

Coordination must persist a deterministic writer order.
Each later writer must read and preserve accepted semantics from earlier writers.
A later writer must return to investigation and decision when it needs to weaken, remove, or reinterpret earlier accepted semantics.

Place one integrated independent review after the final writer.
Use one integrated review Task per Work Item.

The integrated review inspects:

- the originating Requirement;
- the Work Item Goal, Boundary, and Completion Conditions;
- decided outcomes and supporting Investigation records;
- required ADRs;
- the final Specification and originating-artifact state;
- added or amended Task graph structure;
- the combined state after every shared writer.

Split another Work Item when a design boundary is independently closable.
Each resulting Work Item receives its own integrated review.

## Rationale

Serialization prevents lost updates and preserves accepted meaning across shared artifacts.

A final integrated review validates the state that downstream work will actually consume.
Reviewing intermediate fragments cannot establish consistency after later writers change the same artifact.

One review per completion boundary also avoids duplicated or competing verdicts.

## Rejected alternatives

| alternative | rejection reason |
|---|---|
| Allow concurrent writers with later merge resolution. | Merge resolution may silently choose between accepted semantics. |
| Review each writer independently and omit final review. | No review would inspect the final combined state. |
| Create several partial integrated reviews inside one Work Item. | Competing verdicts would obscure the Work Item completion authority. |
| Let a later writer reinterpret earlier decisions. | Authoring must not introduce new design judgment. |
| Use one repository-wide review for several independent Work Items. | Independent completion boundaries need independent verdicts. |

## Consequences

- Work Item Task flows must identify shared-writer ordering.
- Coordination owns writer-order changes.
- Downstream authoring depends on the earlier writer when targets overlap.
- Integrated review begins only after all required authoring is complete.
- Review findings evaluate the final combined state rather than isolated diffs.

## Evidence

- `PRODUCT-TASK-SPEC-018-01`: D-015 and D-016.
- `PRODUCT-TASK-SPEC-018-02`: B-004 routing boundary.
- `PRODUCT-ADR-SPEC-005`: independent Task completion boundary.
