# DRMCP-TASK-MCP-018-02: Decide architecture-derived module contracts

- **id**: DRMCP-TASK-MCP-018-02
- **status**: done
- **date**: 2026-07-05
- **work_item**: DRMCP-WORK-MCP-018
- **task_type**: decision
- **estimate**: 2.0d
- **depends_on**:
  - DRMCP-TASK-MCP-018-01
- **outputs**:
  - DRMCP-TASK-MCP-018-02

## Goal

Decide a coherent contract partition and the module or collaboration boundaries required by the accepted DRMCP application architecture.

## Work

### Working definition

`Contract-level component` means one named logical unit inside an accepted architecture component. It owns one cohesive behavior and exposes an explicit contract to collaborators. It does not imply one Go package, type, interface, process, or Specification file.

- Inventory accepted application-architecture authority and existing operation contracts.
- Separate repository-resolvable facts from user design judgments.
- Classify existing coverage and unresolved module-contract gaps.
- Maintain one resumable decision inventory with explicit dependencies and canonical targets.
- Derive a smaller contract-level component model inside the accepted application architecture before selecting a contract partition.
- Give every active Application Use Case at least one explicit component.
- Add shared components only when they own distinct reusable orchestration, domain behavior, source access, state, or failure obligations.
- Decide responsibility, collaboration, input, output, state, failure, and invariant obligations.
- Decide canonical Specification topology and ADR-routing inputs without authoring canonical artifacts.
- Ask exactly one unresolved user judgment at a time and persist each answer before advancing.
- Stop and route through graph coordination when another Task owner becomes necessary.

### Repository inventory

#### Accepted architecture facts

- DRMCP has five application-level responsibility components.
- Dependencies point inward from MCP and Infrastructure toward Application and Domain contracts.
- Public use cases do not call one another.
- Operation-specific use cases reuse shared application orchestration and Domain capabilities.
- Read, Validation, and Guidance use one fresh immutable request-scoped Current Records snapshot.
- Legacy Archive remains a separate source capability and separate request state.
- Guidance is a fixed-scope Application projection over normal Current Records.
- Request rejection, expected semantic outcomes, and execution failures have distinct owners.
- The trustworthy-result rule prohibits partial normal results from incomplete mandatory state.
- Authoring transaction internals and request-spanning mutable proposal state remain architecture-return concerns.

#### Existing contract coverage

| boundary | established coverage | unresolved module-contract gap |
|---|---|---|
| Composition / Lifecycle | Startup configuration, source registration, application wiring, server start, immutable server-lifetime configuration, and shutdown ownership. | Lifecycle collaboration contract, dependency inputs, resource obligations, and failure propagation. |
| MCP Inbound Adapter | Public request and response meaning, protocol decoding boundary, application invocation, and no semantic reclassification. | Shared adapter-to-use-case contract and transport-independent application input/output boundary. |
| Application Use Cases | Operation policy, sequencing, source selection, response projection, and shared orchestration responsibility. | Use-case and shared-orchestration partition, collaboration rules, and cross-operation invariants. |
| Record Domain / Logical Tree | Current identity, parsed record state, logical lookup, retrieval, resolution, validation findings, headings, sections, and body representation. | Cohesive domain capability contracts, accepted inputs and outputs, invariants, and modeled failures. |
| Infrastructure I/O Adapters | Current and Legacy enumeration, reading, source separation, mandatory-root behavior, and access failure reporting. | Inward-owned source-port contracts, adapter obligations, source-result shapes, and completeness guarantees. |
| Cross-component collaboration | Dependency direction, request-state lifetime, Current-versus-Legacy separation, failure classes, and architecture-return triggers. | Exact seams, state visibility, preconditions, postconditions, and propagation rules. |
| Deferred Authoring extension | Public operation names and the MCP-to-use-case seam remain recognized. | Internals remain excluded. Retained mutable state or write transactions require architecture return. |

#### Existing lower-level design input

`spec:drmcp.implementation` already defines a partial lower-level structure for the four W011 read-runtime operations:

- operation-specific use cases;
- request-snapshot orchestration;
- records and parsing;
- current indexing and legacy exact lookup;
- reference resolution;
- validation;
- filesystem and configuration adapters.

That Specification is not a complete authority for W018 because it excludes both Guidance use cases and is scoped to the former four-operation W011 slice. Its conceptual separation is reusable evidence; its exact Go package tree is outside W018.

#### Inventory conclusions

- Existing operation contracts define public behavior more precisely than internal collaboration.
- The architecture defines responsibility placement but does not select contract artifacts.
- The five components do not imply five contracts or five Work Items.
- Contract-level components must be smaller than the five architecture components.
- Every active Application Use Case requires at least one explicit component before contract partitioning.
- Shared orchestration, Domain capability, and source-access components may be added when they own distinct reusable behavior.
- Contract partitioning follows component decomposition; component count does not predetermine Specification count.
- No formal Investigation is currently required.

### Decision inventory

| ID | topic | status | depends on | decision summary | reason / evidence | canonical target | ADR route |
|---|---|---|---|---|---|---|---|
| D-001 | Component-first derivation order | `decided` | none | First derive a contract-level component model below the accepted architecture. Give every active Application Use Case at least one explicit component. Add smaller shared components only for distinct reusable responsibilities. Select the contract partition after this decomposition. | The architecture components are too coarse for module contracts. The earlier `responsibility-and-seam` wording was undefined and not safely resumable. | W018 component model and contract partition | candidate |
| D-002 | Active use-case component decomposition | `decided` | D-001 | Create six explicit contract-level components: `List Records Use Case`, `Get Records Use Case`, `Resolve Reference Use Case`, `Validate Records Use Case`, `List Authoring Guides Use Case`, and `Get Authoring Guidance Use Case`. | The user requires at least one component per Application Use Case. Deferred Authoring internals remain outside W018 because current architecture requires an architecture return before those internals are designed. | W018 component model | candidate |
| D-003 | Shared component decomposition | `decided` | D-002 | Separate `Current Records Snapshot Assembly` and `Legacy Lookup State Assembly`. Keep app namespace, domain namespace, canonical identity, and current addressability inside one `Current Records Logical Tree` responsibility rather than creating a separate namespace component. The logical tree has two current identity branches: a sequential-artifact branch organized by app, domain, kind, and record identity; and a Spec branch organized by app and path-derived topic segments. Add one `Record Relation Graph` component over logical-tree nodes for relations such as `depends_on`, parent, tasks, work items, and source refs. W018 fixes the graph component seat and collaboration boundary only; graph storage shape, adjacency representation, indexing strategy, and traversal algorithm remain downstream detailed-design freedom. Add one separate `Record Parser` component as the only boundary from untrusted raw Markdown plus provenance into typed domain parse results. The parser output models valid values, missing values, invalid present values, document structure, and parse findings explicitly. `Current Records Logical Tree` accepts only typed parser results and does not parse or reinterpret raw Markdown. Add one separate `Reference Resolution` component that owns exact supplied-value classification, current-first lookup sequencing, accepted Legacy eligibility, Legacy fallback lookup, and transport-neutral resolution outcomes. It depends on Current Records Logical Tree and optional Legacy Lookup State but does not own public response projection or exact retrieval. Split validation into two contract-level components: `Local Record Validation` for checks decidable from one typed source or record, and `Relation Graph Validation` for checks requiring logical-tree, relation-graph, Current target, Legacy target, reciprocity, or Topics-graph context. Treat both as stable validation containers with internal orchestration. Individual validation rules or handlers receive stable rule IDs and remain downstream detailed-Specification units, not additional contract-level components. `Validate Records Use Case` remains the application orchestrator that selects subjects, invokes validation capabilities, aggregates findings, suppresses semantic duplicates, orders results, and projects the public response. Legacy exact lookup remains outside the Current Records Logical Tree. Split infrastructure source access by source family into `Current Records Source Access` and `Legacy Archive Source Access`. Each component owns its family-specific enumeration, reading, provenance, boundary enforcement, and access-failure reporting. Enumeration and reading remain internal operations or downstream port refinements rather than additional contract-level components. | Accepted architecture already assigns identity and immutable logical structures to Record Domain / Logical Tree. Existing identity contracts define sequential and path-derived Spec forms, so namespace handling is a refinement of the logical tree rather than an independent owner. Relation data can cross branches and be many-to-many, so it requires an explicit graph responsibility but not an implementation representation at contract-convergence time. Separating parsing prevents raw, partially trustworthy source representation from leaking into logical-tree and use-case contracts while preserving invalid sources as explicit typed states for validation. Reference resolution has independent current/legacy grammar, ordering, and outcome semantics and is too cohesive and complex to bury inside logical-tree lookup. Validation rules change more frequently than the stable local-versus-cross-record dependency boundary, so the two-component split preserves a stable contract while allowing detailed specifications to assign evolving rules internally. Stable rule IDs preserve traceability from PRODUCT authority to checks and findings without forcing every rule change to reopen the component model. Current and Legacy source access have different authority, consumers, provenance, validation participation, and failure meaning, while enumeration and reading share one lifecycle and source-family boundary. | W018 component model | candidate |
| D-004 | Coherent contract partition | `decided` | D-002, D-003 | Organize module contracts as a recursive hierarchical canonical topic tree. Use the five accepted architecture components as the exact top-level contract domains: `composition-lifecycle`, `mcp-inbound-adapter`, `application-use-cases`, `record-domain-logical-tree`, and `infrastructure-io-adapters`. Under each top-level domain, place the contract-level components decided by D-002 and D-003 as semantic subdomains; where no smaller component was selected, use focused topic files directly under the architecture domain. Every domain directory contains an `index.md` limited to Overview or navigation-first Index content. Substantive obligations always live in focused topic files, and another nested subdomain is created when a coherent responsibility boundary exists below the current domain. Do not use `index.md` as the main all-in-one contract and do not flatten every topic into one directory. | The accepted architecture supplies the stable top-level ownership map, while D-002 and D-003 supply the smaller contract-level decomposition. PRODUCT repository-layout and Spec-format authorities require a canonical topic tree and limit Index content to navigation. This partition keeps architecture traceability explicit while allowing recursively deeper semantic grouping without coupling file count to component count. | Contract Specification topology | candidate |
| D-005 | Responsibility and collaboration boundaries | `decided` | D-004 | Public use cases do not call one another. MCP invokes only the matching Application Use Case. Concrete Infrastructure adapters are reached only through inward-owned source contracts. Domain performs no I/O. `Reference Resolution` directly consumes exact lookup capabilities from `Current Records Logical Tree` and optional `Legacy Lookup State`; `Resolve Reference Use Case` does not mediate each Current and Legacy lookup step or reproduce resolution sequencing. `Relation Graph Validation` directly consumes `Record Relation Graph`, `Current Records Logical Tree`, and optional `Legacy Lookup State`; `Validate Records Use Case` does not perform or pass each relation-target lookup. `Local Record Validation` receives one already-selected typed source or record plus provenance and returns findings without querying the tree, graph, Legacy state, parser, or Infrastructure. `Current Records Snapshot`, `Legacy Lookup State`, typed parser results, validation findings, and resolver outcomes are treated as component handoff type or protocol contracts, not automatically as behavioral components. At this stage W018 enumerates those type/protocol surfaces; their recursive non-flat Specification placement is decided later. | Direct specialist collaboration preserves one owner for resolution and relation-validation semantics while keeping request-level sequencing and response projection in Application. Local validation stays per-subject because its rules are decidable from one typed subject. The user's correction separates behavioral components from handoff objects: request-state and outcome objects exist, but their contract is a type/protocol boundary rather than another component boundary. | Contract Specifications | candidate |
| D-006 | Conceptual input and output boundaries | `decided` | D-004, D-005 | Establish the initial conceptual handoff type/protocol surface inventory without deciding field-level shape or physical Specification placement. The inventory includes: public operation request, public operation response, Current source input or source-access result, Legacy source input or legacy-access result, typed source, typed record, parse finding, Current Records Snapshot, Legacy Lookup State, logical-tree lookup result, relation-graph query result, resolution outcome, validation finding, aggregated validation result, source-backed location reference, and Guidance projection input. Each surface is defined by its producer component, consumer component, exchanged meaning, and purpose; field schemas and concrete implementation types remain downstream. | Contracts require accepted inputs and outputs without implementation-language signatures. D-005 established that some W018 artifacts are handoff object contracts rather than behavioral components, so D-006 first lists the conceptual type surfaces that later Specifications will define. The accepted list is deliberately a type/protocol inventory, not a flat Specification topology. | Contract Specifications | candidate |
| D-007 | State ownership and visibility | `decided` | D-004, D-005, D-006 | Current Records Snapshot and Legacy Lookup State are fresh immutable request state. Application Use Cases and explicitly named Domain collaborators may read them. MCP Inbound Adapter and Infrastructure I/O Adapters do not observe their internal structure. No component mutates them after construction. No Current or Legacy request state is retained across requests, cached as server-wide record state, refreshed in the background, or reused after source failure. Server-lifetime configuration remains owned by Composition / Lifecycle and is visible to active use cases only as validated dependencies. | This preserves the accepted runtime-and-state contract: every Read, Validation, and Guidance request builds one fresh immutable Current snapshot, optional Legacy state remains separate, and request state is discarded when the request ends. Visibility is limited to Application and Domain so MCP cannot reclassify semantics and Infrastructure cannot depend on domain state. | Contract Specifications | candidate |
| D-008 | Failure ownership and propagation | `decided` | D-005, D-006, D-007 | Expected invalidity, lookup miss, unsupported value, duplicate conflict, fallback disabled, unreadable indexed target, zero match, parse finding, and validation finding are modeled as typed semantic states or transport-neutral findings returned by Domain or source contracts. Application Use Cases project those states or findings into operation-specific normal responses, diagnostics, warnings, request errors, or operation errors. Technical failures that prevent complete trustworthy mandatory state remain execution failures selected by Application and encoded by MCP without normal result wrapping. MCP Inbound Adapter does not reclassify semantic outcomes or execution failures. | This preserves the accepted distinction between request rejection, expected semantic outcome, and execution failure. It prevents broad exception-to-diagnostic conversion while allowing Application Use Cases to project Domain outcomes into operation-specific results. It also preserves the trustworthy-result rule: incomplete or untrustworthy mandatory state cannot produce partial normal data. | Contract Specifications | candidate |
| D-009 | Invariants, preconditions, and postconditions | `decided` | D-005 through D-008 | Each behavioral component and each handoff type/protocol contract records observable preconditions, invariants, postconditions, and forbidden bypasses. Contracts do not record algorithms, storage layout, concrete Go signatures, concrete struct fields, or physical file placement. Cross-cutting invariants include dependency direction, no public use-case calls, raw Markdown crossing only through parser outputs, Current and Legacy state separation, fresh immutable discarded request state, expected semantic outcomes as typed states or findings, and untrustworthy mandatory-state failure preventing normal results. | Observable obligations are sufficient for downstream detailed design without freezing implementation choices. This shape converts D-005 through D-008 into enforceable contract constraints while preserving algorithm, storage, type, signature, and file-layout freedom for later Specifications. | Contract Specifications | candidate |
| D-010 | Canonical Specification topology | `decided` | D-004 through D-009 | Place W018 module contracts under the DRMCP implementation Specification area as `spec:drmcp.implementation.contracts` or an equivalent `implementation/contracts` topic subtree. Inside that contract tree, use the five accepted architecture components as the first semantic subdomains: `composition-lifecycle`, `mcp-inbound-adapter`, `application-use-cases`, `record-domain-logical-tree`, and `infrastructure-io-adapters`. Place behavioral component contracts and handoff type/protocol contracts under the nearest owning subdomain. Keep every `index.md` navigation-only and avoid a flat catch-all component list or type catalog. Existing `application-architecture` specs remain the current architecture authority for W018; relocating them under `implementation` is a separate topology or migration question outside this decision unless explicitly routed later. | Repository layout requires Specifications to be organized by canonical topic tree under `records/spec/`. D-004 already rejected flat placement, and D-009 fixed the obligation shape. The user corrected the contract-tree home from `design-records-mcp/contracts` toward `implementation/contracts`, which better matches module-contract ownership because these contracts refine implementation architecture rather than public MCP tool behavior. Existing `spec:drmcp.implementation` is W011-scoped and cannot be adopted unchanged, so W018 requires a new contracts subtree rather than simply extending the current W011 implementation text. | `spec:drmcp.implementation.contracts` topic tree | candidate |
| D-011 | ADR routing and architecture-return disposition | `decided` | D-004 through D-010 | Treat D-001 through D-010 as ADR-required as one coherent module-contract ownership route because the combined component decomposition, collaboration boundaries, type/protocol inventory, state and failure semantics, obligation shape, and `implementation/contracts` topology establish durable module ownership boundaries and constrain several downstream Specifications and implementation waves. Treat the existing `application-architecture` placement concern as outside W018 module-contract authoring: do not move it here, but record a follow-up topology or migration question if desired. Treat changes that add or remove top-level architecture components, reverse inward dependencies, merge Current and Legacy state, introduce request-spanning mutable state, or change trustworthy-result semantics as architecture-return triggers rather than W018-local contract work. | ADR routing is required because these choices have durable ownership, lifecycle, failure, and topology consequences that are not mechanically derived from one existing ADR. The architecture-return boundary keeps W018 as a module-contract refinement of the accepted architecture rather than an architecture rewrite or repository-wide topology migration. ADR routing is separate from ADR authoring; this decision classifies the route and boundary without writing ADR content. | ADR-routing owner and application-architecture route | required |
| D-012 | Detailed-specification handoff | `decided` | D-004 through D-011 | W018 hands downstream authoring a fixed component model, fixed collaboration rules, fixed type/protocol surface inventory, fixed request-state visibility and retention rules, fixed semantic-vs-execution failure ownership, fixed obligation shape, fixed `implementation/contracts` topology, and a required ADR route. W018 closure does not release production implementation planning. It releases downstream ADR authoring, canonical module-contract Specification authoring, and follow-up detailed contract convergence split at least by component or coherent subdomain. That detailed contract convergence still does not release implementation planning by itself; it only enables operation or feature behavior Specification work written against the completed component contracts. Downstream detailed contract work may choose exact ADR boundaries, exact spec files inside the topology, field schemas, Go type names, interface signatures, package layout, storage representation, algorithms, rule-handler decomposition, and test fixtures only when those choices remain observable-equivalent under D-001 through D-011. Production implementation planning remains blocked until required ADRs, module-contract Specifications, component-level detailed contracts, operation or feature behavior Specifications, integrated review, and closure synchronization complete. | W018 defines fixed contract state and downstream local freedom. This final handoff prevents hidden design judgment from leaking into detailed contract authoring or implementation while avoiding the false claim that W018-level module boundaries, or even one coarse detailed-contract pass, are sufficient for direct implementation. | W018 closure and downstream detailed contract framing | not_required |

### Status definitions

- `open`: Known judgment not yet selected.
- `in_discussion`: The one current user judgment.
- `decided`: Explicit accepted outcome persisted.
- `blocked`: Cannot proceed until a named dependency or authority is available.
- `deferred`: Explicitly routed outside this Work Item.
- `superseded`: Replaced by a later decision item.

### Current cursor

- Decision: none
- Loop state: `decision_complete`
- Next action: Proceed to downstream graph coordination, ADR routing, canonical module-contract Specification authoring, integrated review, closure synchronization, and then component-scoped detailed contract convergence. Do not start implementation planning until component-level detailed contracts and operation or feature behavior Specifications close.

### Expected downstream route

```text
active contract decision ledger
  -> conditional Investigation route when a durable bounded question appears
  -> decision ledger completion
     -> later graph coordination
     -> ADR routing
     -> canonical module-contract Specification authoring
     -> integrated independent review
     -> conditional finding-specific correction and closure review
     -> closure synchronization
     -> downstream detailed contract convergence split at least by component or coherent subdomain
     -> operation or feature behavior Specification work based on completed component contracts
     -> implementation planning only after relevant behavior Specifications close
```

No downstream Task is materialized by this decision Task.

## Done condition

- Every owned decision item is `decided`, `deferred`, or validly `blocked`.
- The contract-level component model is explicit and smaller than the accepted architecture model.
- Every active Application Use Case has at least one explicit component.
- The partition is coherent, non-overlapping, architecture-derived, and traceable.
- Each required boundary has responsibility, collaboration, input, output, state, failure, and invariant decisions.
- Exact canonical targets and ADR-routing inputs are identified.
- The detailed-specification handoff and architecture-return boundary are explicit.
- No ADR, Specification, implementation Task, production implementation, or independent review is authored.

## Verification

- Confirmed the inventory uses current accepted five-component architecture.
- Confirmed all six active operations and the deferred Authoring seam are represented.
- Confirmed no decision item remains `in_discussion`.
- Confirmed every explicit answer was persisted before the cursor advanced.
- Confirmed architecture component count does not determine contract-level component count or Specification count.
- Confirmed implementation detail remains undecided and production implementation planning remains blocked until component-level detailed contracts and operation or feature behavior Specifications close.

## Evidence

- DRMCP-WORK-MCP-018 supplies the accepted Goal, Boundary, Completion Condition, Non-goals, and unknown-handling route.
- `spec:drmcp.application_architecture` and its four views supply the current five-component architecture baseline.
- Active operation contracts cover `list_records`, `get_records`, `resolve_reference`, `validate_records`, `list_authoring_guides`, and `get_authoring_guidance`.
- Existing contracts define public behavior, source semantics, and architecture-level ownership, but not one coherent internal contract partition.
- The initial inventory first identified an undefined contract-unit criterion and was corrected after user review.
- D-001 records component-first derivation before contract partitioning.
- D-002 fixes six explicit active use-case components, one for each active public Application Use Case.
- D-003 records the accepted separation of Current Records snapshot assembly from Legacy lookup-state assembly.
- D-003 also assigns app/domain namespace and canonical current identity to one Current Records Logical Tree with sequential and Spec identity branches.
- D-003 reserves a separate Record Relation Graph component over logical-tree nodes while deferring its internal storage and traversal representation to downstream detailed design.
- D-003 separates Record Parser from Current Records Logical Tree. Raw Markdown and provenance cross into the Domain only through typed parse results that preserve valid, missing, invalid-present, structural, and finding states explicitly.
- D-003 separates Reference Resolution from Current Records Logical Tree lookup because current-first and Legacy fallback sequencing form a distinct reusable semantic capability.
- D-003 splits validation into Local Record Validation and Relation Graph Validation, with Validate Records Use Case retaining selection, orchestration, aggregation, duplicate suppression, ordering, and response projection.
- The split criterion is a stable dependency and responsibility boundary. Evolving PRODUCT rules such as new task types or lifecycle statuses are assigned to stable-ID internal validation rules by downstream detailed Specification without reopening the contract-level component model.
- Internal validation rules or handlers are not contract-level components.
- D-008 now records, without deciding, the revealed candidate that components expose expected invalidity as typed states or findings while untrustworthy technical failures remain execution failures.
- D-003 splits Infrastructure I/O by source family into Current Records Source Access and Legacy Archive Source Access. Enumeration and reading remain internal operations or downstream port refinements.
- D-003 is decided. The accepted component model now covers six public use cases, request-state assembly, Current logical structure, relation graph, parsing, resolution, two validation containers, and two source-family access components.
- The existing W011 implementation Specification supplies a partial conceptual decomposition candidate, but it cannot be adopted unchanged because it excludes Guidance and contains implementation-specific package decisions.
- D-004 fixes the partition shape as a recursive domain tree: each domain has an Overview or navigation-only `index.md`, substantive obligations live in focused topic files, and nested subdomain directories are used when another coherent responsibility area exists below the parent.
- D-004 fixes the five accepted architecture components as exact top-level contract domains and places D-002/D-003 contract-level components beneath them as semantic subdomains.
- D-004 is decided. The earlier flat Overview-plus-fixed-contract-groups proposal and the later optional-child-file interpretation are discarded.
- D-005 accepts direct specialist collaboration for Reference Resolution: it queries Current Records Logical Tree and optional Legacy Lookup State itself, while Resolve Reference Use Case owns request-level orchestration and response projection.
- D-005 also accepts direct specialist collaboration for Relation Graph Validation: it queries Record Relation Graph, Current Records Logical Tree, and optional Legacy Lookup State itself, while Validate Records Use Case owns subject selection, validator invocation, aggregation, ordering, and response projection.
- D-005 fixes Local Record Validation as a pure per-subject validator: it receives one already-selected typed source or record plus provenance and does not query the tree, graph, Legacy state, parser, or Infrastructure.
- D-005 records the correction that Current Records Snapshot, Legacy Lookup State, typed parser results, validation findings, and resolver outcomes are handoff type or protocol contracts, not automatically behavioral components.
- D-005 is decided.
- D-006 fixes the initial conceptual handoff type/protocol surface inventory: public operation request, public operation response, Current source input or source-access result, Legacy source input or legacy-access result, typed source, typed record, parse finding, Current Records Snapshot, Legacy Lookup State, logical-tree lookup result, relation-graph query result, resolution outcome, validation finding, aggregated validation result, source-backed location reference, and Guidance projection input.
- D-006 records that the list is a type/protocol inventory only. Field schemas, implementation types, and physical Specification placement remain downstream.
- D-007 fixes Current Records Snapshot and Legacy Lookup State as fresh immutable request state. Application Use Cases and explicitly named Domain collaborators may read them; MCP and Infrastructure do not observe their internal structure; no component mutates them; no Current or Legacy request state is retained across requests or reused after source failure.
- D-008 fixes expected invalidity, lookup miss, unsupported value, duplicate conflict, fallback disabled, unreadable indexed target, zero match, parse finding, and validation finding as typed semantic states or transport-neutral findings returned by Domain or source contracts. Application projects those outcomes into operation-specific results. Technical failures that prevent complete trustworthy mandatory state remain execution failures selected by Application and encoded by MCP without normal result wrapping.
- D-009 fixes the contract-obligation shape: behavioral component contracts and handoff type/protocol contracts record observable preconditions, invariants, postconditions, and forbidden bypasses, but not algorithms, storage layout, concrete Go signatures, concrete struct fields, or physical file placement.
- D-010 fixes the canonical non-flat Specification topology: W018 module contracts belong under an `implementation/contracts` topic subtree such as `spec:drmcp.implementation.contracts`, not under `design-records-mcp/contracts`. Behavioral component contracts and handoff type/protocol contracts are placed under the nearest owning architecture-component subdomain. Existing `application-architecture` remains current architecture authority for W018 unless a separate topology or migration decision explicitly moves it.
- D-011 fixes ADR routing and architecture-return disposition: D-001 through D-010 require a coherent ADR route for durable module-contract ownership; `application-architecture` placement migration remains a separate follow-up; architecture-return triggers remain outside W018-local contract work.
- D-012 fixes the detailed-specification handoff boundary: W018 closure releases ADR authoring, canonical module-contract Specification authoring, and follow-up detailed contract convergence split at least by component or coherent subdomain.
- D-012 records the correction from user feedback: neither W018-level module boundaries nor one coarse detailed-contract pass are sufficient for direct implementation planning. Component-level detailed contracts first enable operation or feature behavior Specifications; implementation planning remains blocked until those behavior Specifications close.
- All D-001 through D-012 decision items are terminal and decided. The decision loop state is `decision_complete`.
- The Task Done condition is satisfied for the decision ledger. No canonical contract content, ADR, Specification, implementation Task, production implementation, or review verdict was authored.
- DRMCP is non-operational. Design Records MCP cannot author this record, so filesystem authoring is the required fallback.
