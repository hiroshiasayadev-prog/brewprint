# TRV-TASK-SPEC-002-10: Route revised TRV architecture ADRs

- **id**: TRV-TASK-SPEC-002-10
- **status**: done
- **date**: 2026-07-02
- **work_item**: TRV-WORK-SPEC-002
- **task_type**: decision
- **estimate**: 0.25d
- **depends_on**:
  - TRV-TASK-SPEC-002-09
- **outputs**:
  - TRV-TASK-SPEC-002-10

## Goal

Produce one complete ADR-routing ledger for the terminal T09 architecture-documentation, ownership, dependency, flow, and runtime decisions.

## Work

- Consume terminal T09 decisions without changing them.
- Compare each decision with TRV-ADR-SPEC-001 and TRV-ADR-SPEC-002.
- Classify each decision as `required`, `covered`, `not_required`, or validly `blocked`.
- Select `create`, `amend`, `reuse`, or `supersede` for every ADR-owned choice.
- Preserve independent supersession boundaries between application architecture and model-runtime architecture.
- Record exact normative Specification targets.
- Keep exact packages, files, symbols, concrete interfaces, schemas, and implementation choices in W004.

This Task must not:

- reopen or revise T09 decisions;
- author or amend an ADR or Specification body;
- change the Task graph;
- perform review, correction, synchronization, implementation, stage, or commit work.

### Routing ledger

| decision | outcome | disposition / ADR boundary | reason | normative targets |
|---|---|---|---|---|
| D-001 | `not_required` | Specification-only | Document partition and navigation do not preserve a durable architecture trade-off. | `spec:trv`, `spec:trv.application_architecture`, and its child Specifications. |
| D-002 | `not_required` | Specification-only | Overview contents are document-structure requirements, not an independently durable choice. | `spec:trv.application_architecture`. |
| D-003 | `required` | `amend` TRV-ADR-SPEC-001 | The selected ports-and-adapters alternative remains unchanged, but the five-component responsibility model needs precise current wording. | `spec:trv.application_architecture.component_model`. |
| D-004 | `required` | `amend` TRV-ADR-SPEC-001 | Exact inward dependency targets and startup-edge semantics materially constrain downstream design while preserving the original architecture. Mermaid notation itself is Specification-only. | `spec:trv.application_architecture.dependency_model`. |
| D-005 | `not_required` | Specification-only | The runtime sequence is mechanically derived from the accepted component and port ownership decisions. | `spec:trv.application_architecture.validation_flow`. |
| D-006 | `not_required` | Specification and Work Item boundary | The handoff among W002, W003, W004, ADRs, and PRODUCT authority records current scope ownership without a separate architecture alternative. | `spec:trv.application_architecture.boundary` and TRV-WORK-SPEC-002. |
| D-007 | `required` | `amend` TRV-ADR-SPEC-001 | The explicit inbound application-port boundary is a durable dependency choice inside the accepted ports-and-adapters architecture. | component and dependency Specifications. |
| D-008 | `required` | `amend` TRV-ADR-SPEC-001 | Splitting Task-record and checklist-catalog capabilities is a durable port-ownership choice. | component, dependency, and validation-flow Specifications. |
| D-009 | `required` | `amend` TRV-ADR-SPEC-001 | Core-owned complete-prompt construction is a durable ownership choice with a rejected structured-request alternative. | component and validation-flow Specifications. |
| D-010 | `required` | `amend` TRV-ADR-SPEC-002 | The model port now has a precise prompt, decoding, failure, and semantic-validation boundary while preserving the provider-neutral-port and external-Ollama choice. | `spec:trv.model_runtime`, dependency, and validation-flow Specifications. |
| D-011 | `required` | `amend` TRV-ADR-SPEC-001 | MCP-shape validation versus application-semantic validation is a durable adapter/core ownership boundary. | component, validation-flow, and boundary Specifications. |
| D-012 | `required` | `amend` TRV-ADR-SPEC-001 | Startup construction, wiring, and server-lifecycle ownership constrains the architecture while preserving the existing composition-root choice. | component and dependency Specifications. |
| D-013 | `covered` | TRV-ADR-SPEC-001 and TRV-ADR-SPEC-002 | The combined architecture is the consistency result of the two independently durable ADR boundaries; it does not require an omnibus ADR. | all revised architecture Specifications. |

### ADR boundaries

#### Boundary A — Application architecture ownership

- **ADR**: TRV-ADR-SPEC-001.
- **Disposition**: `amend`.
- **Included decisions**: D-003, D-004, D-007, D-008, D-009, D-011, D-012.
- **Bounded question**: Within the accepted ports-and-adapters architecture, which components own the inbound boundary, input ports, prompt construction, application outcomes, transport validation, and startup composition?
- **Materiality judgment**: Non-material amendment. The selected ports-and-adapters alternative, inward-dependency rationale, five top-level responsibility areas, and no-domain-layer choice remain valid. The amendment makes previously ambiguous ownership and consequences explicit; it does not reverse the ADR.
- **Dependency**: None beyond the accepted current ADR and T09.

#### Boundary B — Model-provider execution boundary

- **ADR**: TRV-ADR-SPEC-002.
- **Disposition**: `amend`.
- **Included decision**: D-010.
- **Bounded question**: What crosses the application-owned model-evaluation port, and which validation and failure responsibilities belong to the core versus the Ollama adapter?
- **Materiality judgment**: Non-material amendment. The provider-neutral application port, remote external Ollama runtime, and adapter-owned provider mechanics remain unchanged. The amendment clarifies the port payload and semantic-validation boundary.
- **Dependency**: Preserve Boundary A prompt ownership and application-result semantics.

No new ADR and no supersession are required.
The two ADR boundaries remain independently supersedable.

## Done condition

- Every terminal T09 decision has one resolved ADR-routing outcome.
- TRV-ADR-SPEC-001 and TRV-ADR-SPEC-002 have exact non-material amendment boundaries.
- Every Specification-only decision has a reason and exact normative target.
- No create or supersede route remains open.
- T11 can materialize exact writers and review dependencies without new routing judgment.

## Verification

- Confirmed D-001 through D-013 each have one route.
- Confirmed neither accepted ADR changes its selected alternative or core rationale.
- Confirmed application architecture and model-runtime remain independently supersedable.
- Confirmed no package, symbol, concrete Go interface, schema, or implementation mechanism entered the routing ledger.
- Confirmed no ADR or Specification body changed.

## Evidence

- T08 created this revised ADR-routing owner.
- T09 completed D-001 through D-013 before routing began.
- TRV-ADR-SPEC-001 already owns ports and adapters, five logical areas, inward dependencies, application orchestration, and startup wiring.
- TRV-ADR-SPEC-002 already owns the application model port, Ollama adapter, and external runtime boundary.
- The T09 decisions refine those two accepted alternatives rather than reversing them.
- The user explicitly waived a separate impact Investigation because TRV is a new application with no existing implementation, migration, or operational integration surface. T10 therefore used direct scoped comparison against the current ADRs and Specifications.
- Result: `PASS`.
