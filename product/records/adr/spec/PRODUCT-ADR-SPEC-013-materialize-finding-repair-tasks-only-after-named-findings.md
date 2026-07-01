# PRODUCT-ADR-SPEC-013: Materialize finding repair Tasks only after named findings

- **status**: accepted
- **date**: 2026-07-01
- **depends_on**:
  - PRODUCT-ADR-SPEC-005
  - PRODUCT-ADR-SPEC-012
- **supersedes**: []
- **migrated_to_spec**: null

## Context

A design Work Item may need correction and finding-closure review after integrated review.
The exact repair targets and writable boundaries are unknown before findings exist.

Pre-creating placeholder Tasks forces speculative outputs, dependencies, and completion conditions.
Those contracts can become stale when the actual finding set differs from the prediction.

## Decision

Do not create correction or finding-closure review Tasks in the initial graph.

The Work Item records only the conditional route.

| review result | route |
|---|---|
| `PASS` | Proceed directly to closure synchronization. |
| `NEEDS REVISION` | Use coordination to create exact correction and independent finding-closure review Tasks. |

Derive each created Task from named findings.
The derived contract must define:

- exact finding IDs;
- exact writable artifacts;
- exact dependencies;
- one repair outcome;
- one independent closure-review boundary.

Do not create placeholder or synthetic no-op Tasks for branches that were not taken.

## Rationale

Named findings provide the first reliable basis for repair ownership and scope.

Delayed materialization prevents stale contracts and avoids Tasks whose only valid result is no work.
The conditional Work Item route remains sufficient before the review verdict exists.

Correction and finding closure retain separate owners and completion judgments.

## Rejected alternatives

| alternative | rejection reason |
|---|---|
| Create generic correction and re-review Tasks in every initial graph. | Their outputs and writable boundaries would be speculative. |
| Keep no conditional route until review completes. | The Work Item would omit required failure routing. |
| Use one Task for repair and closure. | The repair author must not independently close the same findings. |
| Mark unused branch Tasks done as no-op. | Synthetic completion creates misleading workflow history. |
| Let the review Task perform corrections. | Review and correction have different ownership and independence requirements. |

## Consequences

- Work Items must describe the `PASS` and `NEEDS REVISION` branches.
- Coordination creates repair Tasks only after the review records named findings.
- Correction Tasks repair only the named finding set.
- A later independent review decides each finding as `CLOSED` or `OPEN`.
- Closure synchronization depends on `PASS` or independently closed required findings.

## Evidence

- `PRODUCT-TASK-SPEC-018-01`: D-020.
- `PRODUCT-TASK-SPEC-018-02`: B-005 routing boundary.
- `PRODUCT-ADR-SPEC-012`: integrated review boundary.
