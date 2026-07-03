# Concept: TRV architecture boundary

- **id**: `spec:trv.application_architecture.boundary`
- **status**: draft
- **date**: 2026-07-02
- **parent**: `spec:trv.application_architecture`

## What this is

Defines which architecture concerns belong to W002 and which concerns must remain in PRODUCT, W003, W004, ADRs, or Specifications.
It prevents downstream contract and detailed-design work from silently changing the architecture.

## Concept model

| concern | canonical owner | boundary |
|---|---|---|
| Cross-app validator semantics | `spec:product.responsibility_boundary_validator` | Owns Task-local Evidence, checklist composition, criterion results, aggregation, outcome classes, and workflow-use semantics. |
| TRV application architecture | TRV-WORK-SPEC-002 and its Specifications | Owns architecture style, components, application ports, dependency direction, stage ownership, and architecture handoff rules. |
| External MCP and application contract | TRV-WORK-SPEC-003 and its Specifications | Owns tool names, external request and response contracts, caller workflow, and external failure envelope. |
| Implementation-ready TRV design | TRV-WORK-SPEC-004 and its Specifications | Owns Go packages, files, declarations, structs, schemas, constructors, configuration names, commands, and concrete verification. |
| Durable architecture rationale | TRV ADRs | Own alternatives, rationale, trade-offs, consequences, amendment, and supersession history. |
| Current normative architecture | TRV Specifications | Own the currently valid component, dependency, flow, runtime, and boundary rules. |
| Current DRMCP integration | Separate future Requirement or Work Item | W002 does not establish a current DRMCP dependency. |

## Rules

### W002-owned architecture

W002 owns:

- ports-and-adapters architecture;
- five top-level component boundaries;
- application-core internal responsibility elements;
- inbound and outbound port ownership;
- source dependency direction;
- startup composition responsibility;
- main validation sequence and stage ownership;
- application-versus-adapter validation boundaries;
- model-provider and external-runtime ownership at architecture level.

### W003 handoff

W003 may define:

- MCP tool names;
- externally visible request and response fields;
- transport-level error and failure projection;
- caller-visible workflow and invocation contract.

W003 must not:

- move validation orchestration into the MCP adapter;
- bypass the inbound application port;
- make the MCP adapter select criteria or construct the prompt;
- recalculate application outcomes or overall compliance;
- change component ownership or dependency direction.

### W004 handoff

W004 may define:

- package and file layout;
- exact interface and type declarations;
- exact prompt text and criterion-correlation mechanism;
- exact request and response schemas;
- constructors and composition mechanics;
- configuration keys and validation rules;
- retry, timeout, logging, testing, and command details.

W004 must not:

- merge or split architecture components without returning to W002;
- replace an application-owned port with a concrete adapter dependency;
- move prompt meaning, criterion completeness, or aggregation into provider infrastructure;
- change the main stage owner or dependency direction.

### ADR and Specification boundary

- ADRs explain durable choices, rejected alternatives, rationale, and consequences.
- Specifications state current normative architecture.
- Specifications must not copy full ADR rationale or rejected alternatives.
- ADRs must not replace detailed current architecture views.

### Return rule

Downstream work must return to a W002 architecture decision route when it needs to change:

- the top-level component set;
- application-core ownership;
- port ownership or direction;
- adapter-to-core dependency direction;
- startup composition ownership;
- main validation stage ownership;
- model-provider versus application-core responsibility.

Local details may remain in W003 or W004 when they preserve all W002 rules.

## Non-goals

- Exact W003 contract fields or error encoding.
- Exact W004 packages, interfaces, schemas, constructors, or commands.
- Implementation sequencing or production implementation.
- Current DRMCP integration design.

## Boundary

This Specification owns architecture handoff rules only.
It does not duplicate the detailed component, dependency, validation-flow, or runtime contracts.

## Related specs

| ref | relation |
|---|---|
| `spec:trv.application_architecture` | Parent architecture Overview. |
| `spec:trv.application_architecture.component_model` | Defines component and application-core ownership. |
| `spec:trv.application_architecture.dependency_model` | Defines static source dependency direction. |
| `spec:trv.application_architecture.validation_flow` | Defines runtime sequence and stage ownership. |
| `spec:trv.model_runtime` | Defines model-provider and external-runtime ownership. |
| `spec:product.responsibility_boundary_validator` | PRODUCT-owned semantic authority consumed by TRV. |
| TRV-ADR-SPEC-001 | Durable application architecture rationale. |
| TRV-ADR-SPEC-002 | Durable model-port and runtime rationale. |
