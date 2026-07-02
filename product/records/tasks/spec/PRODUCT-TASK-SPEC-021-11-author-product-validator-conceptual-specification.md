# PRODUCT-TASK-SPEC-021-11: Author PRODUCT validator conceptual Specification

- **id**: PRODUCT-TASK-SPEC-021-11
- **status**: done
- **date**: 2026-07-02
- **work_item**: PRODUCT-WORK-SPEC-021
- **task_type**: authoring
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-021-10
- **outputs**:
  - PRODUCT-TASK-SPEC-021-11
  - spec:product.responsibility_boundary_validator

## Goal

Project the T06 PRODUCT-owned semantic contract and corrected TRV ownership boundary into the canonical validator Specification.

## Work

- Preserve the accepted one-Task evaluation boundary.
- Preserve automatic checklist selection from declared `task_type`.
- Preserve criterion-level boolean results and Task-local reasons.
- Preserve logical-AND aggregation and outcome-class separation.
- Preserve post-authoring and post-Evidence invocation.
- Preserve caller-owned human acceptance or rejection.
- Replace stale text that assigns standalone implementation delivery to W021.
- Assign concrete app-local design to the TRV namespace without defining that design.
- Keep current DRMCP integration excluded.
- Consume every required routed ADR and the active TRV namespace profile.

This Task must not:

- choose concrete TRV interface, runtime, transport, packaging, model, provider, or implementation details;
- create TRV Requirement, ADR, Work Item, Task, or implementation content;
- change checklist wording or placement;
- change PRODUCT-REQ-SPEC-007;
- perform independent review, decomposition, synchronization, implementation, stage, or commit work.

## Done condition

- `spec:product.responsibility_boundary_validator` contains every PRODUCT-fixed semantic rule from D-003.
- The Specification assigns concrete app-local design to TRV.
- The Specification no longer assigns implementation delivery to W021.
- T02 concrete technology choices remain non-canonical historical inputs.
- Current DRMCP integration remains excluded.
- No TRV-local design choice is introduced.
- No stale ownership contradiction remains within the Specification.

## Verification

- Trace every D-003 PRODUCT-fixed semantic item to one normative Specification location.
- Confirm every D-003 app-local item remains excluded or delegated to TRV.
- Confirm PRODUCT-ADR-SPEC-015 through PRODUCT-ADR-SPEC-017 remain consistent.
- Confirm the active namespace identity matches T10.
- Confirm only the declared Specification and this Task changed.

## Evidence

- T06 fixed the PRODUCT and TRV authority partition.
- T07 created this canonical authoring owner after namespace activation.
- T10 established active `TRV`, `TRV` / `SPEC`, `trv/records/`, and `spec:trv` profile state.
- T16 completed the non-material `PRODUCT-ADR-SPEC-016` ownership amendment.
- DRMCP is non-operational, so filesystem authoring was used.
- Updated only `spec:product.responsibility_boundary_validator` and this Task.
- Updated Specification date to 2026-07-02 for the substantive ownership contract change.
- Removed stale W021 implementation-delivery ownership.
- Added PRODUCT as semantic contract owner and TRV as application owner.
- Delegated app-local Requirement, ADR, Specification, interface, runtime, tests, and concrete interaction to TRV.
- Preserved current DRMCP separation and future-integration deferral.
- Kept historical technology candidates non-canonical until TRV-local design decisions exist.
- No PRODUCT Requirement, checklist, TRV record, Work Item graph, review, synchronization, implementation, stage, or commit work was performed.

### D-003 trace

| PRODUCT-fixed semantic item | normative location |
|---|---|
| One Task per evaluation | `Current contract` and `Rules / Evaluation boundary`. |
| Automatic checklist selection from `task_type` | `Current contract` and `Rules / Checklist selection and composition`. |
| Boolean result and concise Task-local reason per criterion | `Current contract` and `Rules / Criterion results and aggregation`. |
| Mechanical logical AND | `Current contract` and `Rules / Criterion results and aggregation`. |
| Semantic, structural, and execution outcome separation | `Current contract` and `Rules / Outcome separation`. |
| Caller-owned human acceptance or rejection | `Current contract` and `Rules / Workflow invocation and exceptions`. |
| Post-authoring and post-Evidence invocation | `Current contract` and `Rules / Workflow invocation and exceptions`. |
| Current DRMCP integration excluded | `Non-goals`, `Boundary`, and `Related specs`. |

### App-local exclusion trace

| app-local input | treatment |
|---|---|
| Interface and transport | Delegated to TRV. |
| Runtime, model, and provider | Delegated to TRV. |
| Checklist loading and prompt composition | Delegated to TRV implementation design. |
| Response schema and decode policy | Delegated to TRV. |
| Retry, timeout, and configuration | Delegated to TRV. |
| Packaging, launcher, build, tests, and deployment | Delegated to TRV. |
| Historical T02 technology choices | Non-canonical until TRV-local decisions. |

### Verification result

- D-003 PRODUCT-fixed items represented: 8 of 8.
- Stale W021 implementation ownership: removed.
- TRV app-local design choice introduced: no.
- PRODUCT-ADR-SPEC-015 consistency: preserved.
- PRODUCT-ADR-SPEC-016 consistency: preserved.
- PRODUCT-ADR-SPEC-017 consistency: preserved.
- Current DRMCP integration introduced: no.
- Prohibited artifact changes: none.
