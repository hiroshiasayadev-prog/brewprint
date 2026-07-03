# Overview: Task Responsibility Validator

- **id**: `spec:trv`
- **status**: draft
- **date**: 2026-07-03
- **parent**: root

## What this is

Entry point for TRV-owned design records.
It identifies the application and routes readers to app-local architecture and runtime Specifications.

## Current contract

| area | current contract |
|---|---|
| App namespace | `TRV`. |
| Formal name | Task Responsibility Validator. |
| Records root | `trv/records/`. |
| PRODUCT-owned semantics | `spec:product.responsibility_boundary_validator`. |
| TRV-owned scope | App-local Requirement, ADR, Specification, and later implementation work. |
| Application architecture | Ports and adapters around one application core, with five top-level components and inward dependencies. |
| Validation ownership | The application core owns checklist selection, complete prompt construction, result completeness validation, and overall-result construction. |
| Model runtime | An application-owned model-evaluation port calls an externally managed remote Ollama runtime through a provider adapter. |
| Current design state | TRV delivery is suspended by TRV-ADR-SPEC-006 pending controlled responsibility vocabulary and renewed semantic-feasibility evidence. The reviewed architecture remains preserved. |

## Topics

| title | kind | ref | summary |
|---|---|---|---|
| TRV application architecture | Overview | `spec:trv.application_architecture` | Whole-system composition and navigation to component, dependency, validation-flow, and boundary views. |
| TRV model runtime | Concept | `spec:trv.model_runtime` | Application-owned model port, Ollama adapter responsibilities, and external runtime boundary. |

## Topic map

| topic | route |
|---|---|
| Cross-app semantic validator contract | `spec:product.responsibility_boundary_validator`. |
| Application component responsibilities | `spec:trv.application_architecture.component_model`. |
| Static dependency direction | `spec:trv.application_architecture.dependency_model`. |
| Runtime validation sequence | `spec:trv.application_architecture.validation_flow`. |
| Architecture handoff rules | `spec:trv.application_architecture.boundary`. |
| Suspended semantic-validator delivery | TRV-ADR-SPEC-006. |
| Remaining app-local contract design | TRV-WORK-SPEC-005, currently blocked. |
| Implementation-ready detailed design | TRV-WORK-SPEC-004, currently blocked. |
| Future DRMCP integration | Separate Requirement or Work Item. |

## Non-goals

- External MCP tool names, request fields, response fields, or caller workflow contracts not yet closed by W003.
- Exact package, file, symbol, interface, schema, constructor, fixture, command, or implementation design owned by W004.
- Current DRMCP integration.
- Production implementation planning or execution.
- Duplication of PRODUCT-owned semantic validation rules.

## Boundary

| content | owner |
|---|---|
| Cross-app semantic validator behavior and invocation policy | PRODUCT. |
| App-local architecture, contract design, detailed design, and later implementation | TRV. |
| Current DRMCP structural Design Record operations | DRMCP. |
| Future DRMCP integration | Separate Requirement or Work Item. |

## Related specs

| ref | relation |
|---|---|
| `spec:product.responsibility_boundary_validator` | PRODUCT-owned semantic contract consumed by TRV. |
| `spec:trv.application_architecture` | TRV architecture Overview and authoritative child-topic parent. |
| `spec:trv.model_runtime` | Application-owned model port and external Ollama runtime boundary. |
| TRV-ADR-SPEC-001 | Selects ports and adapters with explicit application ownership. |
| TRV-ADR-SPEC-002 | Selects the model-evaluation port and external Ollama runtime. |
| TRV-ADR-SPEC-006 | Suspends semantic-validator delivery and defines restart prerequisites. |
| `spec:product.brewprint.namespaces.app_namespaces` | Active TRV app namespace assignment. |
| `spec:product.brewprint.namespaces.domain_catalog` | Active TRV / SPEC domain assignment. |
| `spec:product.design_records.repository_layout` | App-independent placement contract for `trv/records/`. |
