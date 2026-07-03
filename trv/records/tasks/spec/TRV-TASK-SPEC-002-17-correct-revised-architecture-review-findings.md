# TRV-TASK-SPEC-002-17: Correct revised architecture review findings

- **id**: TRV-TASK-SPEC-002-17
- **status**: done
- **date**: 2026-07-02
- **work_item**: TRV-WORK-SPEC-002
- **task_type**: correction
- **estimate**: 0.5d
- **depends_on**:
  - TRV-TASK-SPEC-002-16
- **outputs**:
  - spec:trv.application_architecture.component_model
  - spec:trv.application_architecture.dependency_model
  - spec:trv.application_architecture.validation_flow
  - TRV-TASK-SPEC-002-17

## Goal

Repair T14 F-MAJ-01, F-MAJ-02, and F-MIN-01 without changing the accepted TRV architecture decisions.

## Work

### F-MAJ-01 — Record and checklist adapter in the runtime view

Update `trv/records/spec/application-architecture/validation-flow.md` so the runtime view represents the record and checklist adapter behind both outbound input ports.

Required correction:

- Keep the validation use case as orchestrator.
- Keep checklist selection in the application core.
- Show the Task-record source port delegating actual Task access to the record and checklist adapter.
- Show the checklist-catalog port delegating actual checklist access to the same adapter.
- Show application-owned Task and checklist data returning through the ports to the validation use case.
- Update stage ownership so actual Task and checklist access belongs to the adapter, while the validation use case owns orchestration and selection.
- Preserve the model-provider adapter path and the five-component architecture.

### F-MAJ-02 — PRODUCT semantic authority

Update `validation-flow.md` so TRV owns only stage order and component responsibility.

Required correction:

- Remove TRV-local normative restatement of one-Task scope, Task-only semantic Evidence, common-plus-type-specific checklist composition, criterion-result semantics, overall AND semantics, and the model-verdict prohibition.
- Replace duplicated semantic rules with a direct reference to `spec:product.responsibility_boundary_validator`.
- Retain only TRV-owned architecture rules for transport validation, application orchestration, adapter delegation, completed-prompt construction, provider execution and decode, application outcome construction, and MCP projection.
- Do not weaken or reinterpret the PRODUCT semantic authority.

### F-MIN-01 — Inbound-port implementation relation

Update `component-model.md` and `dependency-model.md`.

Required correction:

- State explicitly that the validation use case implements the inbound validation-use-case port.
- Remove `validationUseCase --> inboundPort` from the solid source-dependency graph.
- State that implementation or conformance is not represented as a normal source-dependency edge.
- Preserve the rule that the MCP adapter depends on the inbound port rather than the concrete validation use case.
- Do not introduce exact Go interface or method declarations.

This Task may change only:

```text
trv/records/spec/application-architecture/component-model.md
trv/records/spec/application-architecture/dependency-model.md
trv/records/spec/application-architecture/validation-flow.md
trv/records/tasks/spec/TRV-TASK-SPEC-002-17-correct-revised-architecture-review-findings.md
```

This Task must not:

- change T09 decisions, T10 routing, ADRs, the Overview, boundary Specification, model-runtime Specification, or `spec:trv`;
- introduce a new architecture judgment;
- close its own findings;
- change the Task graph or lifecycle closure state;
- design W003 or W004;
- perform implementation, stage, or commit work.

## Done condition

- The validation-flow view shows both outbound ports and the record and checklist adapter with unambiguous stage ownership.
- The validation-flow rules no longer duplicate PRODUCT-owned semantic authority.
- The validation use case is explicitly identified as the inbound-port implementation without a misleading source-dependency edge.
- F-MAJ-01, F-MAJ-02, and F-MIN-01 have complete correction Evidence ready for independent review.

## Verification

- Compare the corrected files directly with each T14 finding and required outcome.
- Confirm the five components, port ownership, prompt ownership, provider boundary, and application outcome ownership remain unchanged.
- Confirm no PRODUCT semantic rule is redefined by TRV.
- Confirm the dependency diagram uses solid arrows only for source dependencies and distinct dashed edges only for startup construction.
- Inspect the scoped diff and run scoped whitespace checks.

## Evidence

- T14 is the finding authority.
- T16 grouped the three projection defects into this correction Task.
- F-MAJ-01 correction: `validation-flow.md` now shows the Task-record source port and checklist-catalog port delegating actual access to one record and checklist adapter, with application-owned data returning through both ports.
- F-MAJ-01 correction: stage ownership now separates validation-use-case orchestration and checklist selection from adapter-owned Task and checklist source access.
- F-MAJ-02 correction: the validation-flow sequence, stage table, outcome table, and Rules now reference `spec:product.responsibility_boundary_validator` for semantic selection, provider-result validation, application outcomes, aggregation, and related behavior.
- F-MAJ-02 correction: TRV-local Rules now retain only transport validation, orchestration, adapter delegation, prompt construction, provider execution and decoding, application outcome construction, and MCP projection ownership.
- F-MIN-01 correction: `component-model.md` now states that the validation use case implements the inbound validation-use-case port.
- F-MIN-01 correction: `dependency-model.md` removes `validationUseCase --> inboundPort`, states that implementation or conformance is not a source-dependency edge, and preserves MCP dependency on the inbound port.
- No T09 decision, T10 routing outcome, ADR, Overview, boundary Specification, model-runtime Specification, `spec:trv`, Task graph, or lifecycle closure state changed.
- DRMCP authoring is non-operational under the current agent authoring policy. Filesystem fallback was used.
- Scoped diff inspection covered exactly the three corrected Specifications and this Task.
- No finding was self-closed; T18 owns independent finding closure.
- Result: `PASS`.
