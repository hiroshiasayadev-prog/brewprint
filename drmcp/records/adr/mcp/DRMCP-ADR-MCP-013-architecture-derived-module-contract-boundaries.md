# DRMCP-ADR-MCP-013: Architecture-derived module contract boundaries

- **status**: accepted
- **date**: 2026-07-06
- **depends_on**:
  - DRMCP-ADR-MCP-010
  - DRMCP-ADR-MCP-011
  - DRMCP-ADR-MCP-012
- **supersedes**: []
- **migrated_to_spec**: 2026-07-06

## Context

DRMCP-WORK-MCP-016 established the current application architecture.
That architecture defines five application-level components and inward dependency direction.

The architecture does not define module contracts.
It does not define contract-level components, collaboration surfaces, handoff type surfaces, or contract Specification topology.

W018 must refine the architecture into a module-contract baseline.
The baseline must be smaller than the five application-level components.
It must also avoid implementation-ready detail.

The former W011 implementation Specification is partial evidence.
It covers four read-runtime operations and excludes Guidance operations.
It also includes package choices that are outside the W018 contract boundary.

## Decision

DRMCP derives module contracts from the accepted application architecture.

The derivation uses these rules:

| rule | decision |
|---|---|
| Component-first derivation | Derive contract-level components before selecting Specification partition. |
| Use-case coverage | Give each active Application Use Case at least one explicit component. |
| Shared component admission | Add shared components only for distinct reusable responsibility. |
| Architecture alignment | Keep the five accepted application-level components as the top contract domains. |
| Implementation neutrality | Do not bind the contract model to Go packages, interfaces, structs, signatures, storage, algorithms, or test fixtures. |

The active Application Use Case components are:

| component | responsibility |
|---|---|
| List Records Use Case | Owns list-records request policy, sequencing, and response projection. |
| Get Records Use Case | Owns exact retrieval request policy, sequencing, and response projection. |
| Resolve Reference Use Case | Owns public resolve-reference request policy and response projection. |
| Validate Records Use Case | Owns validation subject selection, validator orchestration, aggregation, ordering, and response projection. |
| List Authoring Guides Use Case | Owns fixed Guidance list scope, ordering, and response projection. |
| Get Authoring Guidance Use Case | Owns exact Guidance detail scope and response projection. |

The shared contract-level components are:

| component | responsibility |
|---|---|
| Current Records Snapshot Assembly | Builds one request-scoped Current Records state from configured current sources. |
| Legacy Lookup State Assembly | Builds optional request-scoped Legacy Archive exact lookup state. |
| Current Records Logical Tree | Owns app namespace, domain namespace, canonical current identity, and current addressability. |
| Record Relation Graph | Owns relation graph capability over logical-tree nodes. |
| Record Parser | Owns the only boundary from raw Markdown and provenance into typed parse results. |
| Reference Resolution | Owns exact supplied-value classification, current-first lookup sequencing, accepted Legacy eligibility, fallback lookup, and resolution outcomes. |
| Local Record Validation | Owns checks decidable from one already-selected typed source or record. |
| Relation Graph Validation | Owns checks that require logical tree, relation graph, current targets, Legacy targets, reciprocity, or Topics graph context. |
| Current Records Source Access | Owns Current Records source enumeration, reading, provenance, boundary enforcement, and access-failure reporting. |
| Legacy Archive Source Access | Owns Legacy Archive enumeration, reading, provenance, boundary enforcement, and access-failure reporting. |

The Current Records Logical Tree has two current identity branches:

| branch | identity rule |
|---|---|
| Sequential-artifact branch | App namespace, domain namespace, kind, and record identity. |
| Spec branch | App namespace and path-derived topic segments. |

Public use cases do not call one another.
The MCP Inbound Adapter invokes only the matching Application Use Case.
Infrastructure I/O Adapters are reached only through inward-owned source contracts.
Domain components perform no I/O.

Specialist components may collaborate directly when they own the shared semantic rule:

| specialist | allowed collaboration |
|---|---|
| Reference Resolution | Consumes Current Records Logical Tree and optional Legacy Lookup State directly. |
| Relation Graph Validation | Consumes Record Relation Graph, Current Records Logical Tree, and optional Legacy Lookup State directly. |
| Local Record Validation | Receives one already-selected typed source or record plus provenance. It does not query other state. |

Current Records Snapshot, Legacy Lookup State, typed parser results, validation findings, and resolver outcomes are handoff type or protocol contracts.
They are not automatically behavioral components.

The initial handoff surface inventory is:

| surface | role |
|---|---|
| Public operation request | Transport-neutral input from MCP to Application. |
| Public operation response | Transport-neutral result from Application to MCP. |
| Current source input or source-access result | Current source data, provenance, and access state. |
| Legacy source input or legacy-access result | Legacy source data, issued-ID candidate, provenance, and access state. |
| Typed source | Parser output for source-level validation and snapshot construction. |
| Typed record | Parser output for addressable record semantics. |
| Parse finding | Parser-produced finding for validation projection. |
| Current Records Snapshot | Fresh immutable request state for Current Records. |
| Legacy Lookup State | Separate request state for Legacy exact lookup. |
| Logical-tree lookup result | Exact current lookup outcome. |
| Relation-graph query result | Graph query outcome over relation edges. |
| Resolution outcome | Current-first and optional Legacy fallback outcome. |
| Validation finding | Transport-neutral diagnostic candidate. |
| Aggregated validation result | Application-level validation projection. |
| Source-backed location reference | Stable diagnostic location material. |
| Guidance projection input | Current Records projection input for Guidance use cases. |

Current Records Snapshot and Legacy Lookup State are fresh immutable request state.
Application Use Cases and explicitly named Domain collaborators may read them.
MCP and Infrastructure do not observe their internal structure.
No component mutates them after construction.
No request state is retained across requests.

Expected invalidity, lookup miss, unsupported value, duplicate conflict, fallback disabled, unreadable indexed target, zero match, parse finding, and validation finding are typed semantic states or transport-neutral findings.
Application Use Cases project those states into operation-specific results.
Technical failures that prevent complete trustworthy mandatory state remain execution failures.
MCP does not reclassify semantic outcomes or execution failures.

Each behavioral component contract and each handoff type or protocol contract records:

- preconditions;
- invariants;
- postconditions;
- forbidden bypasses.

Contracts do not record algorithms, storage layout, concrete Go signatures, concrete struct fields, or physical file placement.

Canonical module-contract Specifications live under `spec:drmcp.implementation.contracts`.
The first subdomains are:

| subdomain | owner |
|---|---|
| `composition-lifecycle` | Composition / Lifecycle contracts. |
| `mcp-inbound-adapter` | MCP Inbound Adapter contracts. |
| `application-use-cases` | Application Use Case contracts. |
| `record-domain-logical-tree` | Record Domain / Logical Tree contracts. |
| `infrastructure-io-adapters` | Infrastructure I/O Adapter contracts. |

Every `index.md` in that tree is navigation-only.
Substantive obligations live in focused topic files.
Behavioral component contracts and handoff type or protocol contracts are placed under the nearest owning subdomain.
A flat component list or catch-all type catalog is not allowed.

W018 closure does not release production implementation planning.
It releases ADR authoring, canonical module-contract Specification authoring, and later detailed contract convergence.
Detailed contract convergence must be split at least by component or coherent subdomain.
Operation or feature behavior Specifications follow completed component contracts.
Implementation planning remains blocked until the relevant behavior Specifications close.

## Rationale

The five application-level components are too coarse for module contracts.
A smaller component model makes use-case ownership and reusable Domain behavior explicit.

A component-first derivation prevents Specification layout from driving responsibility boundaries.
It also prevents one flat contract file from becoming a hidden architecture substitute.

The accepted architecture already owns top-level components, dependency direction, request freshness, Current-versus-Legacy separation, and trustworthy-result semantics.
W018 refines those rules into contract seams without reopening the architecture.

Separating parser, logical tree, graph, resolution, and validation preserves one owner for each semantic rule.
It also keeps invalid source states available for validation without letting raw Markdown leak past the parser.

The type surface inventory gives downstream authors a stable handoff map.
It avoids treating every state object as another behavioral component.

The `implementation/contracts` topology fits the scope because these contracts refine implementation architecture.
They are not public MCP tool contracts.

## Rejected alternatives

| alternative | rejection reason |
|---|---|
| Use the five architecture components as the complete contract model | The boundaries are too coarse for use-case and collaborator contracts. |
| Select the Specification layout before component derivation | File layout would hide responsibility decisions. |
| Create one component per handoff object | State and outcome objects can be type or protocol contracts without behavior ownership. |
| Put module contracts under `design-records-mcp/contracts` | That path suggests public MCP tool contracts rather than implementation module contracts. |
| Use a flat `components.md` and `types.md` catalog | A flat catalog loses recursive ownership and parent-child traceability. |
| Start implementation planning after W018 closure | W018 does not define component-level field contracts or operation behavior Specifications. |
| Treat one detailed-contract pass as implementation-ready | Detailed contract convergence must be split by component or coherent subdomain. |

## Consequences

- W018 creates a module-contract baseline below the accepted application architecture.
- Existing application-architecture ADRs remain current authority.
- No application-architecture Specification is moved by this ADR.
- `spec:drmcp.implementation.contracts` becomes the canonical module-contract home.
- The existing W011 implementation Specification cannot be adopted unchanged.
- Downstream Specs must preserve the W018 component model and collaboration rules.
- Downstream detailed contracts may choose field schemas, Go type names, signatures, package layout, storage representation, algorithms, rule-handler decomposition, and test fixtures only when observable behavior remains equivalent under this ADR.
- Production implementation planning remains blocked until relevant behavior Specifications close.
- Changes that add or remove top-level architecture components, reverse inward dependencies, merge Current and Legacy state, introduce request-spanning mutable state, or change trustworthy-result semantics return to application-architecture work.

Affected design areas:

- `spec:drmcp.implementation`;
- `spec:drmcp.implementation.contracts`;
- `spec:drmcp.application_architecture`;
- downstream component-level detailed contract work;
- downstream operation or feature behavior Specification work.

## Evidence

- Source Requirement: DRMCP-REQ-MCP-006.
- Source Work Item: DRMCP-WORK-MCP-018.
- Source architecture Work Item: DRMCP-WORK-MCP-016.
- Accepted architecture ADRs: DRMCP-ADR-MCP-010, DRMCP-ADR-MCP-011, and DRMCP-ADR-MCP-012.
- Decision ledger: DRMCP-TASK-MCP-018-02.
- ADR routing: DRMCP-TASK-MCP-018-04.
- The decision ledger records D-001 through D-012.
