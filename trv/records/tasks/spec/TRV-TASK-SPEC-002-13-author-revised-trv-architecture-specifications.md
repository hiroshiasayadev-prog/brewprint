# TRV-TASK-SPEC-002-13: Author revised TRV architecture Specifications

- **id**: TRV-TASK-SPEC-002-13
- **status**: done
- **date**: 2026-07-02
- **work_item**: TRV-WORK-SPEC-002
- **task_type**: authoring
- **estimate**: 0.75d
- **depends_on**:
  - TRV-TASK-SPEC-002-12
- **outputs**:
  - spec:trv
  - spec:trv.application_architecture
  - spec:trv.application_architecture.component_model
  - spec:trv.application_architecture.dependency_model
  - spec:trv.application_architecture.validation_flow
  - spec:trv.application_architecture.boundary
  - spec:trv.model_runtime
  - TRV-TASK-SPEC-002-13

## Goal

Author the revised TRV architecture Specification set as one readable and internally consistent architecture document family.

## Work

- Update `trv/records/spec/application-architecture/index.md` as the `spec:trv.application_architecture` Overview.
- Keep the Overview limited to purpose and scope, one whole-system composition diagram, concise boundary notes, and authoritative `## Topics` navigation.
- Create `trv/records/spec/application-architecture/component-model.md` for `spec:trv.application_architecture.component_model` and the five top-level components plus application-core internal elements.
- Create `trv/records/spec/application-architecture/dependency-model.md` for `spec:trv.application_architecture.dependency_model`, static source dependencies, exact application-owned dependency targets, startup construction edges, and allowed and forbidden edge tables.
- Use Mermaid subgraphs for top-level components and the application-core internals.
- Create `trv/records/spec/application-architecture/validation-flow.md` for `spec:trv.application_architecture.validation_flow`, the runtime sequence, and stage ownership from MCP input through Task and checklist access, prompt construction, model evaluation, criterion validation, overall-result construction, and MCP projection.
- Create `trv/records/spec/application-architecture/boundary.md` for `spec:trv.application_architecture.boundary`, ADR-versus-Spec ownership, and W002/W003/W004 handoff rules.
- Update `trv/records/spec/model-runtime/index.md` for `spec:trv.model_runtime` with the amended model-port, Ollama-adapter, decoded-result, execution-failure, and core-validation boundary.
- Update `trv/records/spec/index.md` for `spec:trv` topic registration and summaries to match the new child structure.
- Keep the Overview, child Specifications, model-runtime Specification, and amended ADRs semantically consistent.
- Keep exact MCP fields, Go packages, symbols, interfaces, constructors, prompt wording, schemas, commands, and implementation mechanics outside this Task.

This Task must not:

- change a T09 decision or T10 routing outcome;
- add another architecture document without a decided responsibility;
- duplicate ADR alternatives and rationale as Specification body content;
- author W003 external contracts or W004 implementation-ready design;
- perform independent review, correction, synchronization, implementation, stage, or commit work.

## Done condition

- The architecture area has one Overview and four child Concept Specifications with non-overlapping responsibilities.
- The Overview provides orientation, a whole-system diagram, and authoritative child navigation without duplicating detailed child rules.
- Component, dependency, validation-flow, boundary, and runtime views describe one consistent architecture.
- `spec:trv` registers the revised architecture topics correctly.
- W003 and W004 can proceed without choosing component ownership or dependency direction.

## Verification

- Confirm every T09 D-001 through D-013 normative target is represented.
- Confirm every child `parent` matches one authoritative `## Topics` row.
- Confirm dependency diagrams distinguish source dependencies from startup construction edges.
- Confirm runtime flow is not used as a substitute for the static dependency model.
- Confirm ADR rationale is not copied into the Specifications.
- Confirm no W003 or W004 detail entered the document family.

## Evidence

- T09 fixed the document set and architecture views.
- T10 fixed ADR and Specification routing.
- T12 completed both ADR amendments before Specification authoring began.
- Updated `trv/records/spec/index.md` as the TRV topic registry.
- Replaced `trv/records/spec/application-architecture/index.md` with an Overview containing one whole-system composition diagram and authoritative child navigation.
- Created `component-model.md`, `dependency-model.md`, `validation-flow.md`, and `boundary.md` under the application-architecture area.
- Updated `trv/records/spec/model-runtime/index.md` with the completed-prompt, decoded-result, execution-failure, and external Ollama boundaries.
- Confirmed each child `parent` matches exactly one `## Topics` declaration.
- Confirmed the dependency view uses solid source-dependency edges and distinct dashed startup construction edges.
- Confirmed the validation-flow view validates Task readability and `task_type` before checklist selection.
- Confirmed the Specifications contain current normative state without copying ADR alternatives or rationale.
- Confirmed exact MCP fields remain W003-owned and exact packages, interfaces, prompt text, schemas, constructors, configuration names, and commands remain W004-owned.
- DRMCP authoring is non-operational, so filesystem authoring was used under the current agent authoring policy.
- No independent review, correction, synchronization, implementation, stage, or commit work occurred.
- Scoped whitespace inspection returned `PASS` before final Task Evidence update.
- Result: `PASS`.
