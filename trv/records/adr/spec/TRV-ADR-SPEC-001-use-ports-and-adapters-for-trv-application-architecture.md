# TRV-ADR-SPEC-001: Use ports and adapters for TRV application architecture

- **status**: accepted
- **date**: 2026-07-02
- **depends_on**:
  - PRODUCT-ADR-SPEC-016
- **supersedes**: []
- **migrated_to_spec**: 2026-07-02

## Context

TRV must realize one PRODUCT-owned semantic validation contract.
The first delivery uses MCP, repository files, stored checklists, and one model provider.

Those mechanisms can change independently.
Future DRMCP integration may replace the transport and record-access adapters.
A transport-first or provider-first structure would couple the validation use case to temporary delivery choices.

The application also needs clear ownership without adding layers that serve no independent domain model.

## Decision

Use a ports-and-adapters architecture.

Use five logical ownership areas:

1. application core;
2. record and checklist adapter;
3. model-provider adapter;
4. MCP adapter;
5. startup and composition root.

The application core owns:

- the inbound validation-use-case port;
- validation orchestration;
- application input, output, and outcome models;
- application outcomes for structural precondition failure, completed semantic evaluation, and execution failure;
- one Task-record source port that returns Task content and the declared `task_type`;
- one checklist-catalog port that returns common criteria and criteria for the declared `task_type`;
- complete prompt construction;
- one-to-one correspondence between input criteria and returned criterion results;
- application and semantic precondition validation;
- criterion-result completeness validation;
- exact criterion-set validation;
- mechanical overall-result construction under `spec:product.responsibility_boundary_validator`.

The MCP adapter owns:

- MCP and JSON shape validation;
- conversion from MCP input into application-owned input;
- invocation through the inbound application port;
- projection of application outcomes into the MCP contract.

The MCP adapter must not:

- depend on the concrete validation-use-case implementation;
- recalculate overall compliance;
- synthesize missing criterion results;
- convert execution failure into semantic non-compliance.

The record and checklist adapter may implement both outbound input ports.
The adapter does not select validation behavior or construct the model prompt.

The model-provider adapter implements the model-evaluation port.
Provider-specific execution remains outside the application core.

The startup and composition root owns:

- configuration loading and required-configuration validation;
- construction of the application core and concrete adapters;
- wiring adapter implementations to application-owned ports;
- MCP server startup and shutdown.

The startup and composition root does not own validation behavior.

Apply these dependency rules:

- adapters depend on application-owned ports and models;
- the MCP adapter depends on the inbound port, not the concrete use case;
- the application core depends only on application-owned ports and models;
- adapters do not call each other directly;
- startup may know top-level concrete components for construction and wiring;
- the application core has no MCP, filesystem, checklist-storage, HTTP, environment-variable, or process-lifecycle knowledge.

Do not introduce a separate domain layer.
Exact packages, source paths, symbols, interface declarations, constructor shapes, prompt text, and criterion-key mechanics remain detailed-design decisions.

## Rationale

The validation use case is the stable application concern.
MCP, filesystem access, checklist storage, and Ollama are replaceable mechanisms.

Inward dependencies preserve the application use case across adapter replacement.
The inbound port keeps MCP separate from the concrete use-case implementation.
Separate Task-record and checklist-catalog ports preserve distinct responsibilities and failure boundaries.

Core-owned prompt construction keeps validator meaning outside the Ollama adapter.
The provider adapter remains responsible only for provider execution and translation.

Transport-shape validation belongs in the MCP adapter.
Application and semantic validation belong in the validation use case.

Explicit logical areas prevent unrelated infrastructure responsibilities from collapsing into one package.
A separate domain layer would add ceremony without owning an independent model or policy boundary.

## Rejected alternatives

| alternative | rejection reason |
|---|---|
| Put the validation use case inside the MCP server layer. | Transport replacement would require moving or rewriting application behavior. |
| Use one application core plus one generic infrastructure area. | Task access, checklist loading, provider calls, and transport projection would share unclear ownership. |
| Use domain, application, infrastructure, and transport layers. | The separate domain layer would not own an independent model and would add avoidable indirection. |
| Let adapters call each other directly. | Adapter ordering and coupling would bypass application-owned orchestration. |
| Let the MCP adapter call the concrete validation use case. | The transport layer would depend on one application implementation rather than an application-owned boundary. |
| Use one generic input port for Task and checklist access. | Task access and checklist access have different responsibilities and failure modes. |
| Let the model-provider adapter construct the validation prompt. | PRODUCT validation semantics and checklist composition would leak into provider infrastructure. |
| Pass a structured provider-neutral evaluation request instead of a completed prompt. | The first Ollama-only implementation does not justify the added abstraction and mapping surface. |

## Consequences

- Application behavior is testable without MCP, filesystem, or Ollama.
- MCP and record access can change during future DRMCP integration without changing the application core contract.
- Provider-specific retry, timeout, HTTP, and decoding stay outside the core.
- The MCP adapter validates MCP and JSON shape before application invocation.
- The validation use case owns application meaning, prompt construction, result completeness, and overall-result construction.
- Task access and checklist access use separate application-owned ports.
- Startup code owns configuration validation, top-level construction, wiring, and MCP server lifecycle.
- Detailed Specifications must define current component, dependency, flow, and boundary views before implementation begins.
- Exact packages, interfaces, types, prompt text, schemas, and wiring mechanics remain later design work.
- `spec:trv.application_architecture` and its child Specifications own the current normative projection.

## Evidence

- TRV-TASK-SPEC-001-02 D-017 selected ports and adapters with inward dependencies.
- TRV-TASK-SPEC-001-02 D-011 selected the five logical ownership areas and no separate domain layer.
- TRV-TASK-SPEC-002-02 routed D-017 and D-011 into this ADR boundary.
- TRV-TASK-SPEC-002-09 D-003, D-004, D-007, D-008, D-009, D-011, and D-012 clarified the component and responsibility model.
- TRV-TASK-SPEC-002-10 classified these clarifications as a non-material amendment of this ADR.
- PRODUCT-ADR-SPEC-016 assigns app-local design and later implementation ownership to TRV.
