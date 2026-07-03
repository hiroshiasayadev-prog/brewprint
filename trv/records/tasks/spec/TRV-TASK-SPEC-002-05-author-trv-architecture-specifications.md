# TRV-TASK-SPEC-002-05: Author TRV architecture Specifications

- **id**: TRV-TASK-SPEC-002-05
- **status**: done
- **date**: 2026-07-02
- **work_item**: TRV-WORK-SPEC-002
- **task_type**: authoring
- **estimate**: 1d
- **depends_on**:
  - TRV-TASK-SPEC-002-04
- **outputs**:
  - TRV-WORK-SPEC-002
  - TRV-TASK-SPEC-002-05
  - spec:trv
  - spec:trv.application_architecture
  - spec:trv.model_runtime

## Goal

Project the accepted TRV architecture ADRs into current normative Specifications and register them under `spec:trv`.

## Work

- Create `spec:trv.application_architecture` at `trv/records/spec/application-architecture/index.md`.
- Define the application core, ports, adapters, logical component ownership, inward dependency direction, and startup/dependency-wiring boundary.
- Create `spec:trv.model_runtime` at `trv/records/spec/model-runtime/index.md`.
- Define the provider-neutral model-evaluation port, remote Ollama HTTP adapter ownership, external runtime and model lifecycle, and deployment preconditions.
- Update `spec:trv` to register both architecture topics and replace pending architecture statements with current reviewed-target routing.
- Reference PRODUCT-owned semantic behavior rather than duplicating it.
- Preserve external MCP contract and exact implementation detail for later Work Items.

This Task must not:

- change accepted ADR decisions;
- author contract or implementation-ready detailed Specifications;
- make new design decisions;
- perform review, correction, synchronization, implementation, stage, or commit work.

## Done condition

- `spec:trv.application_architecture` and `spec:trv.model_runtime` exist with normative current-state content.
- `spec:trv` registers both child topics.
- Specification content agrees with TRV-ADR-SPEC-001 and TRV-ADR-SPEC-002.
- PRODUCT semantic ownership remains external and referenced.
- Contract and implementation-ready details remain outside the authored boundary.

## Verification

- Confirm spec IDs, paths, parent refs, topic registration, and document shapes follow active TRV SPEC rules.
- Confirm every ADR decision has an exact normative projection.
- Confirm no rationale is duplicated from ADRs.
- Confirm no contract or detailed-design requirement was introduced.
- Confirm only the declared outputs changed.

## Evidence

- T04 authored accepted TRV-ADR-SPEC-001 and TRV-ADR-SPEC-002.
- T03 materialized this Specification writer after the ADR writer.
- `spec:trv.application_architecture` was created as the normative component and dependency boundary.
- `spec:trv.model_runtime` was created as the normative provider-port and external Ollama runtime boundary.
- `spec:trv` now registers both topics and records that W002 independent review remains pending.
- TRV-WORK-SPEC-002 now lists the exact ADR and Specification impact refs.
- PRODUCT semantic behavior remains referenced through `spec:product.responsibility_boundary_validator`.
- External MCP contract and implementation-ready source design remain excluded.
- No review, correction, synchronization, contract, detailed-design, implementation, stage, or commit work occurred.
- Result: `PASS`.
