# PRODUCT-TASK-SPEC-016-06: Author required ADRs

- **id**: PRODUCT-TASK-SPEC-016-06
- **status**: done
- **date**: 2026-06-30
- **work_item**: PRODUCT-WORK-SPEC-016
- **source_requirement**: PRODUCT-REQ-SPEC-005
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-016-05
- **outputs**:
  - PRODUCT-ADR-SPEC-004
  - PRODUCT-ADR-SPEC-005
  - PRODUCT-ADR-SPEC-006

## Goal

Author the exact ADR set required by the accepted T05 routing result.

Record durable choices and rationale without writing current Specification text.

## Work

- Create, amend, or supersede only ADRs identified by T05.
- Preserve one coherent decision boundary per ADR.
- Record adopted choices, rationale, rejected alternatives, and consequences.
- Link each ADR to its originating T02 or T04 decision evidence.
- Record exact ADR refs for T07.

## Done condition

- Every ADR required by T05 is authored with a coherent decision boundary.
- Supersession is explicit where an accepted ADR is replaced.
- Decisions classified as no ADR remain outside ADR files.
- T07 has exact accepted ADR refs or an explicit no-ADR result.

## Verification

- Compare authored ADRs with the complete T05 routing set.
- Confirm ADRs contain decision history rather than current normative Specification text.
- Confirm no independent review is performed by this Task.
- Confirm no Specification file changed.

## Evidence

### Authoring authority

- `PRODUCT-TASK-SPEC-016-05` is `done` and requires exactly three new ADRs.
- DRMCP is non-operational under the current agent authoring policy.
- Filesystem authoring was used for the three exact T05 records.
- The ADR authoring contract is `spec:product.design_records.authoring_standards.adr_authoring`.
- No migration action is owned or required by this Task.

### Authored ADRs

| ADR | status | decision boundary | source decisions | Specification targets |
|---|---|---|---|---|
| `PRODUCT-ADR-SPEC-004` | accepted | Closed `task_type` field, nine-value taxonomy, owned outcome, completion judgment, and prohibited overlaps. | D-001 through D-011 | `task_authoring`, `work_item_authoring` |
| `PRODUCT-ADR-SPEC-005` | accepted | Mandatory single-responsibility cohesion and independent completion boundaries. | D-013 through D-018; C-003; C-005; C-008 through C-012 | `task_authoring`, `work_item_authoring` |
| `PRODUCT-ADR-SPEC-006` | accepted | Decision-workflow checkpoints, conditional ADR routing, and canonical design-state ownership. | D-012; C-006; C-007; C-013; C-014 | `task_authoring`, `artifact_boundary`, `artifact_responsibility_matrix` |

### No-ADR decisions

The following T05 routes remain outside ADR files:

- C-001: metadata, Create, and Update projection for `task_type`.
- C-002: type-contract table representation.
- C-004: conditional `## Implementation contract` placement, columns, and `TBD` rules.

T07 must synchronize those details directly from T02 and T04 together with the three accepted ADRs.

### Verification result

- All three ADRs use the canonical PRODUCT SPEC domain placement.
- All three ADRs use the required H1, metadata, and body section shape.
- `PRODUCT-ADR-SPEC-005` and `PRODUCT-ADR-SPEC-006` depend on `PRODUCT-ADR-SPEC-004`.
- No ADR supersedes an accepted decision.
- No independent review was performed by T06.
- No Specification file was changed.
- T07 has the exact accepted ADR set and direct no-ADR inputs.
