# PRODUCT-TASK-SPEC-021-10: Author TRV namespace profile

- **id**: PRODUCT-TASK-SPEC-021-10
- **status**: done
- **date**: 2026-07-02
- **work_item**: PRODUCT-WORK-SPEC-021
- **task_type**: authoring
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-021-09
  - PRODUCT-TASK-SPEC-021-16
- **outputs**:
  - PRODUCT-TASK-SPEC-021-10
  - spec:product.brewprint.namespaces.app_namespaces
  - spec:product.brewprint.namespaces.domain_catalog
  - spec:product.brewprint.layout
  - spec:trv

## Goal

Activate the TRV app namespace and SPEC domain through one bounded Brewprint profile and namespace-overview authoring set.

## Work

- Register app namespace `TRV` with formal name `Task Responsibility Validator`.
- Register the active `TRV` / `SPEC` domain assignment.
- Create the minimal `spec:trv` overview at `trv/records/spec/index.md`.
- Keep `spec:trv` limited to app identity, ownership boundary, and topic routing.
- Update the current Brewprint layout only after the `trv/records/` path exists.
- Preserve the generic repository-layout contract without amendment.
- Consume every required routed ADR from T08 and T09.

This Task must not:

- design the TRV interface, runtime, transport, packaging, model, provider, or implementation;
- create TRV Requirements, ADRs, Work Items, or Tasks;
- modify the PRODUCT validator semantic contract;
- perform independent review, decomposition, synchronization, implementation, stage, or commit work.

## Done condition

- `TRV` is listed as an active Brewprint app namespace.
- `TRV` / `SPEC` is listed as an active domain assignment.
- `spec:trv` exists under `trv/records/spec/index.md`.
- `spec:trv` contains no app-local design decision.
- `spec:product.brewprint.layout` records the observed `trv/records/` namespace path.
- Every routed ADR prerequisite is reflected without contradiction.
- No generic namespace or repository-layout rule is changed.

## Verification

- Confirm the four declared Specification outputs exist and satisfy their format contracts.
- Confirm the app namespace, formal name, directory, and domain match T06.
- Confirm `trv/records/` follows the active repository-layout contract.
- Confirm no TRV-local Requirement, ADR, Work Item, Task, or implementation artifact exists.
- Confirm only declared outputs changed.

## Evidence

- T06 selected `TRV`, `Task Responsibility Validator`, and `trv/`.
- T06 selected `TRV-WORK-SPEC-001`, which requires the `TRV` / `SPEC` assignment.
- T07 created this namespace-profile authoring owner.
- T09 materialized T16 as the complete required ADR-authoring route.
- T16 completed the non-material `PRODUCT-ADR-SPEC-016` ownership amendment.
- DRMCP is non-operational, so filesystem authoring was used.
- No namespace-local guide existed before TRV activation. The PRODUCT spec authoring standard and writing standard were used.
- Created `trv/records/spec/index.md` as `spec:trv`.
- Registered `TRV` as an active app namespace with formal name `Task Responsibility Validator`.
- Registered `TRV` / `SPEC` as an active domain assignment.
- Updated `spec:product.brewprint.layout` after `trv/records/` existed.
- Preserved `spec:product.design_records.repository_layout` without amendment.
- No TRV Requirement, ADR, Work Item, Task, interface, runtime, model, provider, packaging, deployment, implementation, review, synchronization, stage, or commit work was performed.

### Changed artifacts

- `spec:product.brewprint.namespaces.app_namespaces`
- `spec:product.brewprint.namespaces.domain_catalog`
- `spec:product.brewprint.layout`
- `spec:trv`
- `PRODUCT-TASK-SPEC-021-10`

### Verification result

- App namespace: `TRV`.
- Formal name: `Task Responsibility Validator`.
- Domain assignment: `TRV` / `SPEC` active.
- Records root: `trv/records/`.
- Canonical overview ref: `spec:trv`.
- Generic repository-layout contract changed: no.
- App-local design decision introduced: no.
- Prohibited artifact changes: none.
