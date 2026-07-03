# TRV-TASK-SPEC-002-09: Decide TRV architecture documentation and dependency model

- **id**: TRV-TASK-SPEC-002-09
- **status**: done
- **date**: 2026-07-02
- **work_item**: TRV-WORK-SPEC-002
- **task_type**: decision
- **estimate**: 0.5d
- **depends_on**:
  - TRV-TASK-SPEC-002-08
- **outputs**:
  - TRV-TASK-SPEC-002-09

## Goal

Produce one bounded decision ledger that defines the TRV architecture document structure, architectural views, and unambiguous module dependency model.

## Work

- Preserve T02 architecture decisions and accepted ADR semantics unless an explicit revision is required.
- Ask exactly one unresolved architecture decision per user turn.
- Persist each explicit answer before advancing the cursor.
- Decide the architecture document set, each document's normative responsibility, and the required architectural views.
- Decide architecture-level ownership, interaction stages, and dependency edges.
- Keep exact Go packages, files, symbols, interface declarations, schemas, constructors, and implementation mechanics in W004.
- Record one concise outcome, reason, canonical target, and ADR-routing state for every terminal item.
- Stop W003 release while any required item remains unresolved.

This Task must not:

- author or amend an ADR or Specification;
- perform ADR routing;
- make external MCP envelope or caller-workflow decisions owned by W003;
- choose exact source layout or Go declarations owned by W004;
- perform review, correction, synchronization, implementation, stage, or commit work.

### Decision inventory

| ID | Topic | Status | Depends on | Expected canonical target | ADR route |
|---|---|---|---|---|---|
| D-001 | Architecture Specification document set and canonical responsibility of each document | `decided` | none | `spec:trv.application_architecture` and child Specifications | `candidate` |
| D-002 | Architecture Overview content, whole-system composition diagram, and navigation boundary | `decided` | D-001 | `spec:trv.application_architecture` | `candidate` |
| D-003 | Component model document and logical component responsibility model | `decided` | D-001 | candidate `spec:trv.application_architecture.component_model` | `candidate` |
| D-004 | Dependency model document, allowed edges, forbidden edges, and diagram notation | `decided` | D-001, D-003 | candidate `spec:trv.application_architecture.dependency_model` and TRV-ADR-SPEC-001 | `candidate` |
| D-005 | Main validation-flow document, interaction stages, and stage ownership | `decided` | D-001, D-003, D-004 | candidate `spec:trv.application_architecture.validation_flow` | `candidate` |
| D-006 | Architecture-boundary document, ADR-versus-Spec responsibility, and W003/W004 handoff | `decided` | D-001 | candidate `spec:trv.application_architecture.boundary` | `candidate` |
| D-007 | Inbound application port and MCP-to-core call boundary | `decided` | D-003, D-004 | component and dependency model Specifications and TRV-ADR-SPEC-001 | `candidate` |
| D-008 | Outbound port partition for Task, checklist, and prompt-input capabilities | `decided` | D-003, D-004, D-007 | component and dependency model Specifications and TRV-ADR-SPEC-001 | `candidate` |
| D-009 | Prompt composition and expected criterion-identity ownership | `decided` | D-005, D-008 | component and validation-flow Specifications | `candidate` |
| D-010 | Model-evaluation port responsibility and provider-independent data boundary | `decided` | D-004, D-005, D-008, D-009 | dependency and validation-flow Specifications, `spec:trv.model_runtime`, and TRV-ADR-SPEC-002 | `candidate` |
| D-011 | Application outcome-model ownership and MCP projection direction | `decided` | D-005, D-007, D-010 | component and validation-flow Specifications | `candidate` |
| D-012 | Startup composition graph and concrete wiring responsibility | `decided` | D-003, D-004, D-007, D-008, D-010 | component and dependency model Specifications | `candidate` |
| D-013 | Complete architecture consistency across component, dependency, flow, boundary, and runtime views | `decided` | D-002 through D-012 | all revised architecture Specifications | `candidate` |

### Decision records

#### D-001 — Architecture Specification document set

- **Decision**: Use `spec:trv.application_architecture` as an Overview with four child Concept Specifications: `component_model`, `dependency_model`, `validation_flow`, and `boundary`.
- **Decision**: Keep `spec:trv.model_runtime` as a sibling under `spec:trv`, not as an application-architecture child.
- **Reason**: Each architectural view has a distinct normative responsibility. Model-runtime ownership can evolve independently from the application-architecture document tree.
- **ADR route**: Candidate. T10 decides whether the document partition itself requires ADR treatment.

#### D-002 — Architecture Overview boundary

- **Decision**: Limit the Overview to a short purpose and scope statement, one whole-system composition diagram, authoritative `## Topics` navigation, and concise non-goals or boundary notes.
- **Decision**: Keep component responsibilities, dependency rules, validation stages, and detailed runtime rules in child Specifications.
- **Reason**: The Overview must provide orientation and navigation without becoming another duplicated architecture contract.
- **ADR route**: Candidate. Likely Specification-only unless T10 identifies durable architectural rationale.

#### D-003 — Component model

- **Decision**: Use five top-level components: MCP adapter, application core, record/checklist adapter, model-provider adapter, and startup/composition root.
- **Decision**: Show validation use case, application ports, and application models as internal elements of the application core.
- **Decision**: Do not split the record/checklist adapter into smaller top-level components at this stage.
- **Reason**: The five components are the stable architectural responsibility units. Port partitioning and prompt ownership remain separate later decisions.
- **ADR route**: Candidate. TRV-ADR-SPEC-001 may require amendment if the exact component model is durable rationale.

#### D-004 — Dependency model and diagram notation

- **Decision**: Make `dependency_model` the static source/module dependency view. Keep runtime call order in `validation_flow`.
- **Decision**: Use Mermaid subgraphs for top-level components. Show inbound ports, validation use case, outbound ports, and application models inside the application-core subgraph.
- **Decision**: Draw adapter dependencies to the exact application-owned port or model rather than to an undifferentiated core box.
- **Decision**: Draw startup/composition-root construction edges only to top-level component boundaries in the architecture diagram.
- **Decision**: Do not expose constructor, factory, use-case, or internal port wiring from startup in this view; exact construction targets belong to W004.
- **Decision**: Use a distinct edge style for startup construction and wiring because those edges are not normal runtime invocation edges.
- **Decision**: Record allowed and forbidden dependency edges in a table beside the diagram.
- **Reason**: Adapter edges must expose their exact application-owned dependency targets. Startup edges only express top-level construction ownership; internal construction detail would overstate the architecture view and duplicate W004.
- **ADR route**: Candidate. TRV-ADR-SPEC-001 likely requires amendment for the precise inward-dependency model.

#### D-005 — Main validation flow

- **Decision**: Define the main runtime sequence as MCP adapter to inbound application port, validation use case, record/checklist capability, validation use case, model-evaluation capability, validation use case, criterion validation and overall-result construction, then MCP projection.
- **Decision**: Keep prompt-composition ownership open for D-009 rather than embedding it in the flow decision.
- **Decision**: Use `validation_flow` for runtime interaction order and stage ownership, not static source dependencies.
- **Reason**: The sequence fixes the architectural collaboration path while preserving later decisions about port partitioning and prompt ownership.
- **ADR route**: Candidate. Likely Specification projection unless T10 identifies durable architectural rationale.

#### D-006 — Architecture boundary and handoff

- **Decision**: Use `boundary` to define ownership and handoff among W002 architecture, W003 external/application contract, W004 implementation-ready detail, ADR rationale, and PRODUCT semantic authority.
- **Decision**: Keep architecture style, component ownership, port ownership, dependency direction, and main flow in W002.
- **Decision**: Keep MCP tool names, request and response contracts, and external failure envelopes in W003.
- **Decision**: Keep Go packages, exact interfaces, structs, schemas, constructors, configuration names, and verification commands in W004.
- **Decision**: Keep rationale, alternatives, and trade-offs in ADRs; keep current normative component, dependency, and flow state in Specifications.
- **Decision**: Require downstream work to return to W002 when it needs to change architecture ownership or dependency direction.
- **Reason**: The boundary prevents contract and detailed-design work from silently making architecture decisions or duplicating ADR rationale.
- **ADR route**: Candidate. Likely Specification-only unless T10 identifies a durable handoff decision requiring ADR treatment.

#### D-007 — Inbound application boundary

- **Decision**: The MCP adapter depends only on an application-owned inbound use-case port and application-owned input/output models.
- **Decision**: The MCP adapter does not depend on the concrete validation-use-case implementation.
- **Decision**: The validation use case implements the inbound port inside the application core.
- **Reason**: This preserves the inward dependency rule and keeps transport adapters replaceable without coupling them to one concrete use-case implementation.
- **ADR route**: Candidate. TRV-ADR-SPEC-001 likely requires amendment for the explicit inbound-port boundary.

#### D-008 — Outbound input ports

- **Decision**: Use two application-owned outbound ports: one Task-record source port and one checklist-catalog port.
- **Decision**: The Task-record source returns the evaluated Task content and declared `task_type` as application-owned data.
- **Decision**: The checklist-catalog port returns the common criteria and the criteria for the declared `task_type`.
- **Decision**: One record/checklist adapter may implement both ports without becoming two top-level components.
- **Decision**: Keep prompt construction outside both ports and decide its owner separately in D-009.
- **Reason**: Task retrieval and checklist retrieval have distinct responsibilities and failure modes. Keeping selection orchestration in the application core preserves the PRODUCT-owned automatic checklist-selection semantics.
- **ADR route**: Candidate. TRV-ADR-SPEC-001 likely requires amendment for the explicit outbound-port partition.

#### D-009 — Prompt composition and criterion-result correspondence

- **Decision**: The application core builds the complete prompt text passed to the model-evaluation port.
- **Decision**: The provider adapter sends that prompt and handles provider-specific HTTP translation and response decoding only.
- **Decision**: The application core owns the requirement that every returned criterion result corresponds one-to-one with an input criterion.
- **Decision**: The exact correspondence mechanism, such as temporary keys, preserved order, or repeated criterion text, remains W004-owned.
- **Reason**: A structured provider-neutral request adds unnecessary abstraction for the first Ollama-only implementation. Core-owned prompt construction keeps the implementation bounded while preserving semantic ownership outside the provider adapter.
- **ADR route**: Candidate. TRV-ADR-SPEC-001 may require amendment for prompt ownership; exact prompt and response shape remain W004-owned.

#### D-010 — Model-evaluation port boundary

- **Decision**: The model-evaluation port accepts a complete prompt string from the application core.
- **Decision**: The model-provider adapter owns provider-specific HTTP execution, timeout handling, request and response translation, and syntactic decoding.
- **Decision**: The port returns decoded criterion-result candidates or an execution failure; it does not return Ollama-specific response types to the core.
- **Decision**: The application core validates result completeness and one-to-one criterion correspondence, distinguishes incomplete evaluation from semantic false results, and derives overall compliance by logical AND.
- **Reason**: This keeps provider mechanics outside the core while preserving PRODUCT-owned semantic validation and aggregation inside the application boundary.
- **ADR route**: Candidate. TRV-ADR-SPEC-002 likely requires amendment for the exact provider-port responsibility boundary.

#### D-011 — Application outcome and MCP projection

- **Decision**: The application core returns application-owned outcomes for structural precondition failure, completed semantic evaluation, and execution failure.
- **Decision**: The MCP adapter validates MCP and JSON transport shape, including required fields, field types, and transport-level errors, before invoking the application port.
- **Decision**: The validation use case validates application meaning and semantic preconditions after transport conversion.
- **Decision**: The MCP adapter maps application outcomes to the MCP contract without recalculating overall compliance, synthesizing missing criterion results, or converting execution failure into semantic non-compliance.
- **Decision**: The application core does not depend on MCP-specific request, response, or error types.
- **Reason**: Transport-format validity belongs at the transport boundary; application and semantic validity belongs in the use case. Keeping those checks separate prevents protocol concerns from leaking into core behavior.
- **ADR route**: Candidate. The ownership boundary may require TRV-ADR-SPEC-001 amendment; concrete MCP projection remains W003-owned.

#### D-012 — Startup and composition responsibility

- **Decision**: Startup/composition root owns configuration loading, required-configuration validation, construction of the application core and concrete adapters, wiring adapter implementations to application-owned ports, and MCP server startup and shutdown.
- **Decision**: Startup does not read Tasks or checklists, compose prompts, invoke the model, evaluate criteria, or transform application outcomes.
- **Decision**: Architecture diagrams show startup construction and wiring only at top-level component boundaries. Exact constructors and wiring mechanics remain W004-owned.
- **Reason**: Startup is the composition boundary, not a business or adapter component. Limiting the architecture view to top-level construction avoids duplicating detailed implementation design.
- **ADR route**: Candidate. Likely Specification projection unless T10 identifies durable rationale requiring ADR treatment.

#### D-013 — Complete architecture consistency

- **Decision**: Accept one coherent architecture composed of the MCP adapter, application core, record/checklist adapter, model-provider adapter, and startup/composition root.
- **Decision**: The MCP adapter validates MCP and JSON shape, then invokes the application-owned inbound port.
- **Decision**: The application core reads the Task and checklist through separate outbound ports, builds the complete prompt, invokes the model-evaluation port, validates criterion-result completeness, derives overall compliance, and returns application-owned outcomes.
- **Decision**: The record/checklist adapter may implement both input ports. The model-provider adapter handles Ollama-specific execution and decoding only.
- **Decision**: Startup constructs and wires top-level components and owns MCP server lifecycle without owning validation behavior.
- **Decision**: The Overview, component, dependency, validation-flow, boundary, and model-runtime Specifications must describe the same component set, dependency direction, runtime sequence, and ownership boundaries.
- **Reason**: These decisions form one complete architecture without requiring W003 or W004 to choose new component ownership or dependency direction.
- **ADR route**: Candidate. T10 must route the combined durable choices into coherent ADR amendments or reuse outcomes.

### Current cursor

- Current item: none.
- Loop state: `decision_complete`.
- No item is in discussion.

### Status definitions

| status | meaning |
|---|---|
| `open` | The item is known but not currently being asked. |
| `in_discussion` | The item is the single active user judgment. |
| `decided` | The selected outcome and reason are persisted. |
| `blocked` | A named external input prevents judgment. |
| `deferred` | The user explicitly moved the item outside this Work Item. |
| `superseded` | A later workflow artifact replaced the item. |

### Fixed inputs

- TRV uses ports and adapters.
- The application core owns validation orchestration and application-owned models.
- MCP, filesystem, checklist storage, HTTP, environment variables, and process lifecycle remain outside the core.
- Adapters depend inward on application-owned contracts.
- Adapters do not call each other directly.
- Startup wiring owns concrete dependency construction.
- The first provider adapter calls an externally managed Ollama runtime over HTTP.
- Exact source and interface declarations remain W004-owned.

### Expected downstream route

```text
terminal T09 decisions
  -> T10 revised architecture ADR routing
  -> T11 exact writer and review-order coordination
  -> revised ADR and Specification authoring
  -> new integrated independent architecture review
  -> new closure synchronization
```

## Done condition

- D-001 through D-013 are `decided`, `deferred`, or validly `blocked`.
- Every terminal item records one selected architecture outcome and concise reason.
- Every terminal item identifies its canonical target and provisional ADR route.
- The resulting document set has one clear normative responsibility per document.
- The component, dependency, validation-flow, boundary, and runtime views are mutually consistent.
- The resulting module dependency graph has one interpretation at the architecture level.
- W003 can define application and external contracts without making a new module-ownership decision.

## Verification

- Confirm at most one item is `in_discussion`.
- Confirm every explicit answer is persisted before the cursor advances.
- Confirm dependent items do not become active before prerequisites are terminal.
- Confirm no package, file, symbol, concrete Go interface, or schema design enters the ledger.
- Confirm no required architecture decision exists only in chat.

## Evidence

- T08 created this decision owner after the current architecture was found non-unique.
- The current artifacts permit both direct concrete-service calls and inbound-port calls from the MCP adapter.
- The current artifacts permit several outbound port partitions for Task, checklist, and prompt-input capabilities.
- The current `spec:trv.application_architecture` is a single Concept that combines area entry, component summary, dependency rules, and boundary notes.
- The current architecture Specifications largely restate ADR outcomes and lack a whole-system composition view, dependency view, and validation-flow view.
- The user requested separate architecture documents and an Overview focused on the whole-system diagram and index.
- The user accepted the proposed Overview plus four child Concepts, with model runtime remaining a sibling Specification.
- D-001 through D-013 are decided.
- The decision loop is complete and ready for T10 ADR routing.
- Filesystem fallback is in use because DRMCP authoring is not operational for this workflow.
