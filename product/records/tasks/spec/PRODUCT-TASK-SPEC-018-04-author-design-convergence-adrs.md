# PRODUCT-TASK-SPEC-018-04: Author design convergence ADRs

- **id**: PRODUCT-TASK-SPEC-018-04
- **status**: done
- **date**: 2026-07-01
- **work_item**: PRODUCT-WORK-SPEC-018
- **task_type**: authoring
- **estimate**: 1d
- **depends_on**:
  - PRODUCT-TASK-SPEC-018-03
- **outputs**:
  - PRODUCT-TASK-SPEC-018-04
  - PRODUCT-ADR-SPEC-009
  - PRODUCT-ADR-SPEC-010
  - PRODUCT-ADR-SPEC-011
  - PRODUCT-ADR-SPEC-012
  - PRODUCT-ADR-SPEC-013
  - PRODUCT-ADR-SPEC-014

## Goal

Author the six new ADRs selected by routing boundaries B-001 through B-006.

## Work

- Author one coherent ADR for each routed boundary.
- Preserve accepted decisions D-001 through D-023 without reopening them.
- Record material alternatives, rationale, consequences, and dependencies.
- Avoid both omnibus ADRs and one-row-per-ADR fragmentation.
- Leave Specification and skill authoring to downstream Tasks.

## Done condition

- `PRODUCT-ADR-SPEC-009` through `PRODUCT-ADR-SPEC-014` exist and are `accepted`.
- Every B-001 through B-006 boundary is represented exactly once.
- Every routed decision appears in the intended ADR boundary.
- No ADR mixes independently changeable decisions.
- No Specification or skill file is changed by this Task.

## Verification

- Confirm six new ADR files exist.
- Confirm every ADR has the required metadata and canonical sections.
- Confirm `supersedes` is empty for all six ADRs.
- Confirm `migrated_to_spec` remains `null` pending Specification synchronization.
- Confirm ADR dependencies reference accepted ADRs.
- Confirm T03 completed before this Task.

## Evidence

| ADR | boundary | decision IDs | result |
|---|---|---|---|
| `PRODUCT-ADR-SPEC-009` | B-001 | D-001, D-004, D-005 | Defines the end-to-end workflow boundary. |
| `PRODUCT-ADR-SPEC-010` | B-002 | D-006, D-011 | Defines typed phases and mismatch routing. |
| `PRODUCT-ADR-SPEC-011` | B-003 | D-012, D-013 | Defines Requirement and Work Item identity continuity. |
| `PRODUCT-ADR-SPEC-012` | B-004 | D-015, D-016 | Defines shared-writer serialization and final integrated review. |
| `PRODUCT-ADR-SPEC-013` | B-005 | D-020 | Defines finding-driven Task materialization. |
| `PRODUCT-ADR-SPEC-014` | B-006 | D-010, D-017, D-019, D-022, D-023 | Defines append-only reconvergence and closure write boundaries. |

- Routing authority: `PRODUCT-TASK-SPEC-018-02`.
- Existing-ADR amendment prerequisite: `PRODUCT-TASK-SPEC-018-03`.
- ADR authoring completed without Specification, skill, review, closure, stage, or commit work.
