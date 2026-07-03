# TRV-TASK-SPEC-002-04: Author TRV architecture ADRs

- **id**: TRV-TASK-SPEC-002-04
- **status**: done
- **date**: 2026-07-02
- **work_item**: TRV-WORK-SPEC-002
- **task_type**: authoring
- **estimate**: 0.5d
- **depends_on**:
  - TRV-TASK-SPEC-002-03
- **outputs**:
  - TRV-WORK-SPEC-002
  - TRV-TASK-SPEC-002-04
  - TRV-ADR-SPEC-001
  - TRV-ADR-SPEC-002

## Goal

Author the two accepted TRV architecture ADR boundaries selected by T02.

## Work

- Create `TRV-ADR-SPEC-001` for ports and adapters, inward dependency direction, application-core ownership, and the five logical component areas.
- Create `TRV-ADR-SPEC-002` for the provider-neutral model port and externally managed remote Ollama HTTP runtime.
- Set both ADRs to `accepted`.
- Make both ADRs depend on PRODUCT-ADR-SPEC-016.
- Preserve T02 decision IDs and routed boundary ownership in ADR Evidence.
- Record rejected architecture alternatives and durable consequences without copying normative Specification text.
- Keep exact packages, files, symbols, interfaces, schemas, and implementation details outside the ADRs.

This Task must not:

- change ADR routing or reopen architecture decisions;
- author architecture Specifications;
- make contract or detailed-design decisions;
- perform review, correction, synchronization, implementation, stage, or commit work.

## Done condition

- TRV-ADR-SPEC-001 and TRV-ADR-SPEC-002 exist with status `accepted`.
- Each ADR contains substantive Context, Decision, Rationale, Consequences, and Evidence.
- ADR boundaries match T02 routing exactly.
- PRODUCT-ADR-SPEC-016 is recorded as the ownership dependency.
- No implementation-ready detail or normative Specification text is duplicated.

## Verification

- Confirm ADR IDs, H1s, metadata, and paths follow active TRV SPEC ADR rules.
- Confirm no accepted alternative or boundary changed from T02.
- Confirm rejected alternatives are decision-relevant and concise.
- Confirm `supersedes` is empty and `migrated_to_spec` remains null before reviewed Specification projection.
- Confirm only the declared outputs changed.

## Evidence

- TRV-TASK-SPEC-002-02 selected two `create` ADR boundaries.
- TRV-TASK-SPEC-002-03 materialized this authoring owner.
- TRV-ADR-SPEC-001 was created for ports and adapters, inward dependencies, logical component ownership, and no separate domain layer.
- TRV-ADR-SPEC-002 was created for the provider-neutral model port and externally managed remote Ollama runtime.
- Both ADRs are `accepted`, depend on PRODUCT-ADR-SPEC-016, supersede nothing, and retain `migrated_to_spec: null`.
- TRV-WORK-SPEC-002 was advanced to `in_progress` when architecture authoring began.
- No architecture Specification, contract, detailed-design, review, correction, synchronization, implementation, stage, or commit work occurred.
- Result: `PASS`.
