# TRV-ADR-SPEC-002: Use external Ollama through a provider-neutral model port

- **status**: accepted
- **date**: 2026-07-02
- **depends_on**:
  - PRODUCT-ADR-SPEC-016
- **supersedes**: []
- **migrated_to_spec**: 2026-07-02

## Context

TRV requires semantic evaluation from a language model.
The first deployment uses Ollama on a separate server.

The application must not make Ollama runtime lifecycle part of the validation use case.
The selected provider can also change independently from the application core.

Embedding provider-specific HTTP and response behavior in the core would make future provider replacement a core rewrite.
Managing Ollama from TRV would add process, model-storage, and deployment ownership unrelated to validation.

## Decision

The application core depends on one application-owned model-evaluation port.
The port contains no Ollama-specific request, response, or error type.

The model-evaluation port accepts one complete prompt string from the application core.
The port returns one of:

- decoded criterion-result candidates;
- an execution failure.

The first provider adapter calls Ollama over HTTP using deployment configuration for:

- the Ollama base URL;
- the model name.

The Ollama adapter owns:

- provider-specific request translation;
- HTTP execution and timeout handling;
- provider-specific response translation;
- syntactic response decoding;
- provider and network failure reporting.

The application core owns:

- validation that every expected criterion has one result;
- validation that each result corresponds to one input criterion;
- rejection of incomplete evaluation output;
- separation of execution failure from semantic non-compliance;
- exact criterion-set validation;
- overall compliance construction through logical AND.

The Ollama adapter does not construct the validation prompt.
The Ollama adapter does not derive overall compliance.

Ollama remains an externally managed runtime on a separate server.
TRV does not:

- install, start, stop, update, or supervise Ollama;
- manage Ollama health as a process-lifecycle responsibility;
- download, store, update, or delete model files.

Network reachability, Ollama availability, and model availability are deployment preconditions.
Additional model providers require a later app-local design decision and a new adapter.
They do not change the provider-neutral application port by default.

## Rationale

The application-owned port keeps the validation use case independent from Ollama-specific types and protocol details.
The completed-prompt boundary avoids an unnecessary structured provider-request abstraction in the first release.

The Ollama adapter contains the unstable provider protocol, HTTP behavior, and syntactic decoding.
The application core retains semantic validation and aggregation.

External runtime ownership matches the actual deployment topology.
It also avoids expanding TRV into a model server or model-management application.

Supporting only Ollama initially keeps the first implementation narrow without making Ollama part of application identity.

## Rejected alternatives

| alternative | rejection reason |
|---|---|
| Import Ollama-specific request and response types into the application core. | Provider replacement would change application-owned contracts and orchestration. |
| Start and supervise Ollama from TRV. | TRV would acquire unrelated process and deployment lifecycle responsibilities. |
| Store or manage model files inside TRV. | Model lifecycle and storage are external deployment concerns. |
| Implement a generic multi-provider framework in the first release. | No accepted need justifies the additional configuration and adapter surface. |
| Default silently to a local Ollama endpoint. | The accepted deployment uses a separate server and must not depend on an accidental localhost runtime. |
| Return raw Ollama responses to the application core. | Ollama response types and protocol details would leak into the application boundary. |
| Let the Ollama adapter validate criterion completeness or derive overall compliance. | Provider infrastructure would own PRODUCT semantic rules. |
| Require a structured provider-independent evaluation request. | The first Ollama-only implementation does not justify the additional request model and rendering layer. |

## Consequences

- TRV requires an externally reachable Ollama endpoint for the first deployment.
- Runtime and model failures appear as adapter or deployment failures rather than semantic non-compliance.
- The application core sends one completed prompt to the model-evaluation port.
- The adapter returns decoded criterion-result candidates or an execution failure.
- The application core validates completeness and criterion correspondence before deriving overall compliance.
- Tests can replace the provider adapter or use a bounded HTTP stub.
- Future providers can be added behind the application-owned port after a separate decision.
- Detailed Specifications must define the current port, failure, runtime, and ownership boundaries.
- Exact HTTP schema, retry, timeout, prompt text, response fields, and configuration names remain later design work.
- `spec:trv.model_runtime` owns the current normative runtime boundary.

## Evidence

- TRV-TASK-SPEC-001-02 D-005 selected a provider-neutral port and a remote Ollama HTTP adapter.
- The user clarified that Ollama runs on a separate server.
- TRV-TASK-SPEC-002-02 routed D-005 into this independent ADR boundary.
- TRV-TASK-SPEC-002-09 D-010 clarified the model-port payload, decoding, failure, and semantic-validation boundary.
- TRV-TASK-SPEC-002-10 classified this clarification as a non-material amendment of this ADR.
- PRODUCT-ADR-SPEC-016 assigns app-local runtime and provider choices to TRV.
