# Concept: TRV model runtime

- **id**: `spec:trv.model_runtime`
- **status**: draft
- **date**: 2026-07-02
- **parent**: `spec:trv`

## What this is

Defines the application-owned model-evaluation port, the first Ollama adapter, and the ownership of the external Ollama runtime.
It excludes exact provider schemas and configuration names.

## Concept model

| element | contract |
|---|---|
| Model-evaluation port | Application-owned boundary with no Ollama-specific request, response, or error type. |
| Port input | One complete prompt string built by the application core. |
| Port output | Decoded criterion-result candidates or an execution failure. |
| First provider adapter | TRV-owned Ollama HTTP adapter. |
| Provider configuration | Deployment supplies the Ollama base URL and model name. Exact configuration names remain W004-owned. |
| Ollama runtime | Externally managed service on a separate server. |
| Model files | Externally managed deployment assets. |
| Availability | Network reachability, Ollama availability, and model availability are deployment preconditions. |

### Responsibility split

| responsibility | application core | Ollama adapter |
|---|---:|---:|
| Build the complete validation prompt | yes | no |
| Select and compose validation criteria | yes | no |
| Translate the prompt into an Ollama request | no | yes |
| Execute HTTP and handle provider timeout | no | yes |
| Decode provider-specific response syntax | no | yes |
| Return decoded criterion-result candidates | consumes | produces |
| Validate one-to-one criterion correspondence | yes | no |
| Reject incomplete criterion output | yes | no |
| Distinguish semantic false from execution failure | yes | reports execution failure only |
| Derive overall compliance through logical AND | yes | no |

## Rules

- The application core must depend on the application-owned model-evaluation port.
- The application core must not import Ollama-specific request, response, HTTP, or error types.
- The model-evaluation port must accept one complete prompt string.
- The model-evaluation port must return decoded criterion-result candidates or an execution failure.
- The Ollama adapter must own provider request translation, HTTP execution, timeout handling, response translation, syntactic decoding, and provider failure reporting.
- The Ollama adapter must not construct the validation prompt.
- The Ollama adapter must not select criteria, validate criterion completeness, or derive overall compliance.
- The application core must validate result completeness and criterion correspondence after adapter decoding.
- Incomplete or uncorrelatable model output must become execution failure, not semantic non-compliance.
- TRV must not install, start, stop, update, or supervise Ollama.
- TRV must not download, store, update, or delete model files.
- TRV must not treat Ollama process health as an application lifecycle responsibility.
- The first implementation supports Ollama only.
- Adding another provider requires a later app-local design decision and another adapter.
- A later provider must preserve the application-owned port unless another decision changes the application contract.

## Non-goals

- Exact Ollama HTTP path, request schema, response schema, or decode algorithm.
- Exact retry, timeout, backoff, logging, metrics, or health-check behavior.
- Exact prompt wording, criterion keys, response fields, or configuration names.
- Ollama installation, process supervision, model storage, or deployment commands.

## Boundary

| concern | owner |
|---|---|
| Prompt meaning, criterion composition, completeness validation, and aggregation | Application core under `spec:trv.application_architecture`. |
| Provider HTTP mechanics and syntactic decoding | Ollama adapter under this Specification. |
| Ollama process and model lifecycle | External deployment environment. |
| Exact provider protocol, interfaces, configuration, and failure types | TRV-WORK-SPEC-004 and its Specifications. |
| External MCP projection of execution failure | TRV-WORK-SPEC-003 and its Specifications. |

## Related specs

| ref | relation |
|---|---|
| `spec:trv` | Parent TRV overview. |
| `spec:trv.application_architecture` | Application architecture Overview. |
| `spec:trv.application_architecture.component_model` | Defines the application core and model-provider adapter components. |
| `spec:trv.application_architecture.dependency_model` | Defines the inward model-port dependency. |
| `spec:trv.application_architecture.validation_flow` | Defines runtime provider invocation and result handling. |
| `spec:product.responsibility_boundary_validator` | Defines semantic evaluation and failure meaning. |
| TRV-ADR-SPEC-002 | Durable model-port and external Ollama runtime decision. |
