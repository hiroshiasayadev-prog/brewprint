# DRMCP-ADR-MCP-008: Inward dependency and responsibility ownership

- **status**: superseded
- **date**: 2026-07-04
- **depends_on**:
  - DRMCP-ADR-MCP-007
  - DRMCP-ADR-MCP-009
- **supersedes**: []
- **migrated_to_spec**: null

## Context

The six-component application model requires one durable dependency direction and responsibility ownership model.
Without that model, transport, orchestration, domain policy, and concrete I/O can drift across component boundaries.

The architecture must also place PRODUCT standards, minimal Guidance behavior, and failure classification without adding new top-level components.
Downstream design needs a clear boundary between local refinement and architecture change.

## Decision

Dependencies point inward.

The MCP Inbound Adapter depends on Application Use Cases.
Application Use Cases do not depend on MCP protocol types.

Application Use Cases depend on Record Domain, Guidance Domain, and application-owned source ports.
Infrastructure I/O Adapters implement those source ports.

Record Domain and Guidance Domain do not depend on Application, MCP, filesystem, or runtime configuration.
Composition / Lifecycle is the only component that may depend on every concrete component for construction and wiring.

Application Use Cases contain two responsibility groups:

| group | responsibility |
|---|---|
| Operation-specific use cases | Public-operation policy, sequencing, and result semantics. |
| Shared application orchestration | Reusable request-scoped source loading and logical-state assembly. |

Public use cases do not call one another.
They depend on shared application orchestration and domain capabilities.
Shared orchestration is not a public use case or a seventh top-level component.

PRODUCT remains the external semantic authority for Design Records standards.
Record Domain owns domain-side rule contracts and applies the selected rules.
Composition / Lifecycle selects and injects the concrete rule implementation or catalog.

Infrastructure I/O Adapters load standards files only when downstream design uses external files.
Application Use Cases invoke domain behavior but do not interpret PRODUCT standards.
MCP Inbound Adapter and Guidance Domain do not own PRODUCT standard interpretation.

Guidance remains minimal.

- Guidance operations remain operation-specific Application Use Cases.
- A Guidance Source port provides directory enumeration and file reads.
- Infrastructure I/O Adapters implement the Guidance Source port.
- Guidance Domain owns path-independent guide identity and projection semantics.
- Guidance Domain owns exact ID lookup, title extraction, abstract extraction, and ASCII ordering semantics.
- Guidance does not use Record Domain state or the record snapshot.
- Guidance has no preload, cache, retained catalog, or background-refresh lifecycle.

Failures use three ownership classes:

| class | owner and handling |
|---|---|
| Request rejection | MCP owns protocol and schema decoding failures. Application owns operation-specific invalid requests. No normal operation response is returned. |
| Expected semantic outcome | Domain produces modeled states or findings. Application selects the operation-specific normal result, diagnostic, partial success, or declared error. |
| Execution failure | Infrastructure reports concrete access failure. Application aborts when a complete trustworthy result cannot be constructed. MCP only encodes the selected result or error. |

The MCP Inbound Adapter does not reclassify semantic outcomes or execution failures.

Downstream contract and detailed-design work may refine these concerns locally:

- package and file layout;
- concrete interfaces and data types;
- parser, resolver, and validator algorithms;
- adapter APIs and filesystem libraries;
- rule encoding and snapshot helper structure;
- error-code representation;
- caching that preserves accepted ownership and lifetime semantics.

Work returns to application-architecture decision when a change would:

- add or remove a top-level component;
- reverse or bypass an accepted dependency;
- move responsibility across component boundaries;
- introduce request-spanning mutable state or a new resource lifecycle;
- change source authority or Current-versus-Legacy separation;
- require cross-use-case coordination outside shared application orchestration;
- change trustworthy-result or failure semantics;
- make an external provider or process part of runtime behavior.

The following dependency and ownership patterns are prohibited:

- a use case depends directly on a concrete filesystem adapter;
- Domain performs I/O;
- an adapter calls a public MCP tool;
- adapters perform lateral orchestration;
- MCP reclassifies semantic results;
- Application Use Cases interpret PRODUCT standards independently;
- one public use case calls another public use case.

This ADR does not decide proposal state, write transactions, packages, interfaces, types, functions, methods, or algorithms.

## Rationale

Inward dependencies keep protocol and infrastructure concerns outside application and domain policy.
Application-owned ports let use cases define required source capabilities without importing concrete adapters.

Operation-specific use cases preserve public-operation ownership.
Shared orchestration permits reuse without coupling public use cases to one another.

PRODUCT authority remains external while Record Domain applies the selected rules.
This placement avoids a new standards component and prevents Application from redefining semantic authority.

Minimal Guidance placement preserves pure guide semantics without creating a stateful subsystem.

The three failure classes keep semantic invalidity available as trustworthy output.
They also prevent incomplete state from becoming a normal response.

Explicit return conditions protect the architecture while leaving local implementation choices downstream.

## Rejected alternatives

| alternative | rejection reason |
|---|---|
| Let use cases import concrete filesystem adapters | Application policy would depend on infrastructure details. |
| Let Domain perform source access | Domain behavior would depend on filesystem and runtime concerns. |
| Let adapters call public MCP tools or orchestrate laterally | Transport and infrastructure boundaries would become application control flow. |
| Let public use cases call one another | Operation policy and result contracts would become transitively coupled. |
| Let MCP reclassify semantic results | Protocol mapping would become the owner of operation meaning. |
| Let Application interpret PRODUCT standards | DRMCP would duplicate or redefine external semantic authority. |
| Make Guidance a preloaded or cached stateful subsystem | The accepted Guidance contract requires only minimal source-backed retrieval. |

## Consequences

- Concrete adapters depend on inward-owned source-port contracts.
- Application and Domain tests can run without MCP or filesystem setup.
- MCP handlers remain protocol mappers rather than semantic owners.
- PRODUCT standard changes enter through selected domain rule contracts or catalogs.
- Guidance remains independent from record snapshots and retained runtime state.
- Operation contracts must state their semantic-result and execution-failure projections.
- Downstream design may refine local structure only inside the accepted ownership model.
- Any listed architecture-return trigger requires new architecture decision work.

Affected design areas:

- `spec:drmcp.application_architecture.dependency_and_responsibility`;
- `spec:drmcp.application_architecture.runtime_and_state`;
- `spec:drmcp.application_architecture.failure_and_evolution`.

## Evidence

- Source Requirement: `DRMCP-REQ-MCP-005`.
- Source Work Item: `DRMCP-WORK-MCP-016`.
- Routed decisions: D-007, D-008, D-010, D-011, D-014, and D-015 in `DRMCP-TASK-MCP-016-03`.
- ADR routing authority: B-02 in `DRMCP-TASK-MCP-016-05`.
- Component-model dependency: `DRMCP-ADR-MCP-007`.
- Lifecycle dependency: `DRMCP-ADR-MCP-009`.
