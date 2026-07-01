# PRODUCT-TASK-SPEC-017-06: Author required ADRs

- **id**: PRODUCT-TASK-SPEC-017-06
- **status**: done
- **date**: 2026-07-01
- **work_item**: PRODUCT-WORK-SPEC-017
- **source_requirement**: PRODUCT-REQ-SPEC-006
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-017-05
- **outputs**:
  - PRODUCT-ADR-SPEC-007
  - PRODUCT-ADR-SPEC-008

## Goal

Author the exact ADR set required by the accepted T05 routing result.

Record durable provenance and migration choices without writing current Specification text.

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
- Confirm no Specification or migration target file changed.

## Evidence

### Authoring authority

- `PRODUCT-TASK-SPEC-017-05` is `done` and requires exactly two new ADRs.
- DRMCP is non-operational under the current agent authoring policy.
- Filesystem authoring was used for the two exact T05 records.
- The ADR authoring contract is `spec:product.design_records.authoring_standards.adr_authoring`.
- No migration action is owned or performed by this Task.

### Authored ADRs

| ADR | status | decision boundary | source decisions | primary Specification targets |
|---|---|---|---|---|
| `PRODUCT-ADR-SPEC-007` | accepted | Work Item-only generic `source_refs`, Task provenance through `work_item`, direct Requirement reverse derivation, and provenance validity. | D-001 through D-012 and D-017 | Requirement, Work Item, and Task authoring; artifact boundary; responsibility matrix; traceability. |
| `PRODUCT-ADR-SPEC-008` | accepted | Staged migration with atomic per-record transitions, removal-only Task migration, and exact-match Requirement reverse-list removal. | D-013 through D-016 | Requirement, Work Item, and Task authoring; trace metadata and validation. |

### Covered and no-ADR routes

- D-018 remains covered by accepted `PRODUCT-ADR-SPEC-001`.
- D-019 remains outside ADR files as a completed writer-order fact.
- C-001 through C-026 and C-031 remain direct Requirement, Specification, or workflow projections.
- C-027 through C-030 remain downstream DRMCP design impacts.

### Historical ADR treatment

- `V01-ADR-091` and `V01-ADR-092` contain partial legacy relation conflicts.
- Neither historical ADR is amended or superseded as a whole.
- `PRODUCT-ADR-SPEC-007` records that current Specifications replace only the conflicting legacy relation sections.

### Verification result

- Both ADRs use the canonical PRODUCT SPEC domain placement.
- Both ADRs use the required H1, metadata, and body section shape.
- `PRODUCT-ADR-SPEC-007` depends on `PRODUCT-ADR-SPEC-001`.
- `PRODUCT-ADR-SPEC-008` depends on `PRODUCT-ADR-SPEC-007`.
- Neither ADR supersedes an accepted ADR.
- ADR content records durable choices and rationale rather than replacing current Specifications.
- No independent review was performed by T06.
- No Specification or migration target file was changed by T06.
- T07 has the exact accepted ADR set and direct no-ADR inputs.
- Bootstrap authority remains `PRODUCT-TASK-SPEC-016-01` Evidence, `BOOTSTRAP-001`.
- This Task remains authored under the current workflow bootstrap `source_requirement` and `work_item` metadata contract.
