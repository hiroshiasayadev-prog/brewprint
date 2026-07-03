# Concept: TRV component model

- **id**: `spec:trv.application_architecture.component_model`
- **status**: draft
- **date**: 2026-07-02
- **parent**: `spec:trv.application_architecture`

## What this is

Defines the stable component responsibilities in the TRV application architecture.
It does not define source packages or constructor structure.

## Concept model

### Top-level components

| component | owns | does not own |
|---|---|---|
| MCP adapter | MCP and JSON shape validation, conversion to application input, inbound-port invocation, and projection of application outcomes. | Validation orchestration, Task access, checklist selection, prompt construction, model execution, or overall-result construction. |
| Application core | Validation use case, application-owned ports and models, Task and checklist orchestration, prompt construction, result completeness checks, and overall-result construction. | MCP, filesystem, checklist-storage, Ollama HTTP, environment-variable, or external process-lifecycle mechanics. |
| Record and checklist adapter | Reading the Task source and checklist catalog through application-owned contracts. | Checklist-selection policy, prompt construction, semantic evaluation, or MCP projection. |
| Model-provider adapter | Provider-specific request translation, HTTP execution, timeout handling, response translation, syntactic decoding, and provider failure reporting. | Prompt meaning, criterion selection, semantic completeness validation, or overall-result construction. |
| Startup and composition root | Configuration loading, required-configuration validation, top-level construction, port wiring, and MCP server startup and shutdown. | Validation behavior, adapter-local protocol behavior, Task access, checklist access, prompt construction, or criterion evaluation. |

### Application-core elements

| element | responsibility |
|---|---|
| Inbound validation-use-case port | Application-owned entry boundary used by the MCP adapter. |
| Validation use case | Implements the inbound validation-use-case port and orchestrates one complete Task responsibility-boundary evaluation. |
| Task-record source port | Requests Task content and declared `task_type`. |
| Checklist-catalog port | Requests common criteria and criteria for the declared `task_type`. |
| Model-evaluation port | Sends one completed prompt and receives decoded criterion-result candidates or execution failure. |
| Application models | Carry application input, Task and checklist data, model-evaluation candidates, criterion results, and application outcomes. |

### Application outcomes

| outcome | meaning |
|---|---|
| Structural precondition failure | The Task cannot be read or parsed, or the declared `task_type` is missing or invalid. |
| Completed semantic evaluation | Every applied criterion has one validated result and one mechanically derived overall result. |
| Execution failure | Model execution or response handling prevented complete semantic evaluation. |

## Rules

- Use exactly the five top-level components in this architecture view.
- Keep the validation use case, application ports, and application models inside the application core.
- The validation use case must implement the inbound validation-use-case port.
- The MCP adapter must call the inbound application port, not the concrete validation-use-case implementation.
- The record and checklist adapter may implement both input ports.
- The application core must select the checklist from the declared `task_type`.
- The application core must construct the completed model prompt.
- The model-provider adapter must not redefine evaluation meaning or criterion selection.
- The application core must keep structural failure, semantic evaluation, and execution failure distinct.
- Do not introduce a separate domain layer without a later architecture decision.

## Non-goals

- Exact interface names, method signatures, structs, package names, or file paths.
- Exact MCP request and response contracts.
- Exact prompt text or criterion-correlation representation.
- Exact provider protocol, retry, timeout, and configuration contracts.

## Boundary

| concern | owner |
|---|---|
| Cross-app semantic validation rules | `spec:product.responsibility_boundary_validator`. |
| Static dependency direction | `spec:trv.application_architecture.dependency_model`. |
| Runtime interaction sequence | `spec:trv.application_architecture.validation_flow`. |
| Cross-Work-Item ownership handoff | `spec:trv.application_architecture.boundary`. |
| Model-provider and external runtime boundary | `spec:trv.model_runtime`. |

## Related specs

| ref | relation |
|---|---|
| `spec:trv.application_architecture` | Parent architecture Overview. |
| `spec:trv.application_architecture.dependency_model` | Defines allowed and forbidden source dependencies. |
| `spec:trv.application_architecture.validation_flow` | Defines runtime stage ownership. |
| `spec:trv.model_runtime` | Defines model-port and Ollama runtime ownership. |
| TRV-ADR-SPEC-001 | Durable application component and ownership decision. |
