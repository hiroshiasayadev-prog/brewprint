# PRODUCT-TASK-SPEC-019-14: Author validator architecture decisions

- **id**: PRODUCT-TASK-SPEC-019-14
- **status**: done
- **date**: 2026-07-01
- **work_item**: PRODUCT-WORK-SPEC-019
- **task_type**: authoring
- **estimate**: 1d
- **depends_on**:
  - PRODUCT-TASK-SPEC-019-13
- **outputs**:
  - PRODUCT-TASK-SPEC-019-14
  - PRODUCT-ADR-SPEC-015
  - PRODUCT-ADR-SPEC-016
  - PRODUCT-ADR-SPEC-017

## Goal

Author the three new ADRs selected by the completed W019 routing ledger.

## Work

- Create PRODUCT-ADR-SPEC-015 for the semantic validation contract.
- Create PRODUCT-ADR-SPEC-016 for the temporary standalone ownership boundary.
- Create PRODUCT-ADR-SPEC-017 for invocation timing and human-owned violation exceptions.
- Use the exact B-001 through B-003 boundaries from T12.
- Preserve the accepted alternatives, rationale, exclusions, dependencies, and consequence targets.
- Set ADR dependencies consistently:
  - ADR-015 depends on ADR-004 and ADR-005.
  - ADR-016 depends on ADR-001 and ADR-015.
  - ADR-017 depends on ADR-009, ADR-015, and ADR-016.
- Keep exact checklist wording, result-field names, runtime choices, and implementation details outside the ADRs.
- Record exact created artifacts and verification Evidence in this Task.

This Task must not:

- reopen T12 routing;
- amend existing ADRs;
- amend PRODUCT-REQ-SPEC-007;
- author the validator Specification or `task_authoring` rule;
- author checklist content or implementation;
- perform review, correction, synchronization, stage, or commit work.

## Done condition

- PRODUCT-ADR-SPEC-015 through PRODUCT-ADR-SPEC-017 exist.
- Each ADR satisfies the canonical ADR authoring contract.
- Each ADR owns only its routed durable question.
- ADR dependencies match the T12 route.
- No routed decision is omitted or duplicated across ADR boundaries.
- No out-of-scope implementation or checklist detail is introduced.

## Verification

- Compare each ADR against its T12 boundary and included decision IDs.
- Confirm ADR-015, ADR-016, and ADR-017 use unique IDs and valid metadata.
- Confirm all required ADR sections are substantive.
- Confirm no existing ADR, Requirement, Specification, Work Item, checklist, or implementation artifact changed.
- Confirm no independent review, correction, synchronization, stage, or commit occurred.

## Evidence

### Result

- Result: `PASS`.
- T12 selected three new ADRs and no amendments or supersessions.
- T13 materialized this single bounded ADR-authoring owner.

### Authoring access

- DRMCP is non-operational under `spec:product.design_records.authoring_standards.agent_authoring_policy`.
- Filesystem authoring was used as the required fallback.

### Created ADRs

- `PRODUCT-ADR-SPEC-015`: B-001 Task-local semantic responsibility-validation semantics.
- `PRODUCT-ADR-SPEC-016`: B-002 temporary standalone ownership and current DRMCP separation.
- `PRODUCT-ADR-SPEC-017`: B-003 two-point invocation and human-owned violation exceptions.

### Dependencies

- `PRODUCT-ADR-SPEC-015` depends on `PRODUCT-ADR-SPEC-004` and `PRODUCT-ADR-SPEC-005`.
- `PRODUCT-ADR-SPEC-016` depends on `PRODUCT-ADR-SPEC-001` and `PRODUCT-ADR-SPEC-015`.
- `PRODUCT-ADR-SPEC-017` depends on `PRODUCT-ADR-SPEC-009`, `PRODUCT-ADR-SPEC-015`, and `PRODUCT-ADR-SPEC-016`.

### Scoped verification

- File names, H1 public IDs, and routed ADR IDs match.
- All required ADR metadata fields are present and canonical.
- All required ADR sections are present and substantive.
- B-001, B-002, and B-003 remain non-overlapping.
- Every T12 included decision is represented in its routed ADR.
- Checklist wording, response field names, runtime choices, and implementation details remain excluded.
- The scoped Git patch is non-truncated.
- Scoped whitespace verification reports `PASS`.

### Change boundary

- Only this Task and the three routed ADRs were changed by T14.
- No existing ADR, Requirement, Specification, Work Item, checklist, DRMCP, or implementation artifact was changed.
- No independent review, correction, closure synchronization, stage, or commit occurred.
