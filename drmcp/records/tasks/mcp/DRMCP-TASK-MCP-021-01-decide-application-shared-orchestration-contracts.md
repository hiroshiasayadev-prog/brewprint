# DRMCP-TASK-MCP-021-01: Decide Application shared orchestration contracts

- **id**: DRMCP-TASK-MCP-021-01
- **status**: done
- **date**: 2026-07-06
- **work_item**: DRMCP-WORK-MCP-021
- **task_type**: decision
- **estimate**: 1.0d
- **depends_on**: []
- **outputs**:
  - DRMCP-TASK-MCP-021-01

## Goal

Decide the provisional Application shared-orchestration contract shape for Current Records Snapshot Assembly, Legacy exact lookup-map Assembly, Current Records Snapshot, and Legacy exact lookup map.

## Work

### Decision boundary

This Task owns a resumable decision ledger for W021.

It decides only provisional detailed-contract boundaries for:

- Current Records Snapshot Assembly;
- Legacy exact lookup-map Assembly;
- Current Records Snapshot;
- Legacy exact lookup map.

It does not author detailed Specification content.
It does not decide public operation request or response contracts.
It does not decide Domain parser, logical-tree, graph, resolver, or validator internals.
It does not decide Infrastructure concrete access fields beyond assembly-facing source contract needs.
It does not add Go package layout, signatures, structs, interfaces, functions, algorithms, fixtures, or tests.

### Authority inventory

| authority | accepted input for W021 |
|---|---|
| DRMCP-WORK-MCP-021 | W021 must decide and author a provisional detailed contract baseline for Application shared orchestration. |
| DRMCP-WORK-MCP-020 | W021 is the first child in the provisional detailed contract wave. |
| DRMCP-WORK-MCP-018 | W018 released downstream component-scoped detailed contract convergence but not production implementation planning. |
| DRMCP-ADR-MCP-013 | Application owns request-level assembly orchestration. Current Records Snapshot and Legacy Lookup State are handoff type or protocol contracts. |
| `spec:drmcp.application_architecture.runtime_and_state` | Each Read, Validation, or Guidance request builds one fresh immutable Current Records Snapshot and discards it after the request. |
| `spec:drmcp.application_architecture.dependency_and_responsibility` | Application Use Cases depend on Domain and Application-owned source-port contracts. Infrastructure implements those source-port contracts. |
| `spec:drmcp.application_architecture.failure_and_evolution` | Application selects execution failure when mandatory state is incomplete or untrustworthy. MCP does not reclassify the selected result. |
| `spec:drmcp.implementation.contracts.application_use_cases.contract_boundary` | Current Records Snapshot Assembly and Legacy Lookup State Assembly are Application shared-orchestration behavioral components. |
| `spec:drmcp.implementation.contracts.record_domain_logical_tree.contract_boundary` | Domain owns semantic structures inside the snapshot and exposes typed construction, lookup, resolution, graph, and validation outcomes. |
| `spec:drmcp.implementation.contracts.infrastructure_io_adapters.contract_boundary` | Current and Legacy source access remain separate and expose source-family states to Application-owned contracts. |
| `spec:drmcp.implementation.contracts.mcp_inbound_adapter.contract_boundary` | MCP passes transport-neutral requests and receives transport-neutral outputs or execution failures. MCP does not observe request-state internals. |
| `spec:drmcp.implementation.contracts.composition_lifecycle.contract_boundary` | Composition constructs validated dependencies. It does not own request-state assembly or retain request-scoped state. |

### Accepted constraints

| constraint | effect on W021 |
|---|---|
| Application-owned orchestration | Assembly components coordinate request-level source access, parsing, Domain construction, completeness handling, and state handoff. |
| Domain semantic ownership | W021 may define assembly-facing Domain construction results but must not move semantic identity, graph, lookup, or validation rules out of Domain. |
| Source-family split | Current and Legacy source inputs remain distinct. A merged generic compatibility state is outside W021. |
| Fresh immutable request state | Current Records Snapshot and Legacy exact lookup map are request-scoped, immutable after construction, and discarded after request end. |
| Visibility boundary | Application Use Cases and explicitly named Domain collaborators may read request state. MCP and Infrastructure do not observe internal structure. |
| Trustworthy-result rule | Incomplete or untrustworthy mandatory state cannot produce partial normal operation data. |
| Implementation neutrality | Contract decisions may define conceptual input, output, state, and failure boundaries, but not concrete Go signatures or structs. |
| Provisional wave boundary | W021 must produce a baseline that W022 can consume without guessing Application-owned assembly responsibilities. |

### Vocabulary mapping

| term | W021 meaning |
|---|---|
| Legacy lookup state | Existing-spec abstract term used by older accepted contracts. |
| Legacy exact lookup map | W021 concrete contract-surface wording for the separate read-only exact lookup map built from configured Legacy Archive roots and keyed by accepted legacy issued IDs. |

Source evidence: `spec:drmcp.design_records_mcp.namespace_scanning`, `spec:drmcp.application_architecture.runtime_and_state`, `spec:drmcp.design_records_mcp.resolver`, and `spec:drmcp.design_records_mcp.tools.validate_records`.

### Decision inventory

| ID | topic | status | depends on | decision summary | reason / evidence | canonical target | ADR route |
|---|---|---|---|---|---|---|---|
| D-001 | Detailed Specification target partition for W021 | `decided` | accepted W018/W020/W021 authority | Create four focused detailed contract specs under `application-use-cases/shared-orchestration/` for the two assembly components and two request-state transfer surfaces. W021 remains Application shared orchestration plus the external boundaries of the request-state handoff surfaces that shared orchestration produces. | User accepted the four focused spec route and corrected the placement to `application-use-cases/shared-orchestration/`. W021 completion names the four covered surfaces separately. W018 forbids flat catch-all type catalogs and requires focused topic files under the nearest owning subdomain. | Detailed specs under `spec:drmcp.implementation.contracts.application_use_cases.shared_orchestration` | not_required |
| D-002 | Current Records Snapshot Assembly input boundary | `decided` | D-001 | Current Records Snapshot Assembly receives only current-source access or current-source results, validated dependencies, and Domain construction capabilities needed to build one request-scoped Current Records Snapshot. It does not receive MCP request objects, public response policy, concrete Infrastructure adapters, Domain internal rule details, or the Legacy exact lookup map. | User accepted the simplified boundary. The accepted shape keeps Application shared orchestration as the assembly owner without moving Domain semantics, Infrastructure concretes, MCP transport, or Legacy exact lookup-map data into Current assembly. | W021 detailed Application shared-orchestration specs | not_required |
| D-003 | Current Records Snapshot Assembly output and failure boundary | `decided` | D-002 | Current Records Snapshot Assembly returns either one complete trustworthy Current Records Snapshot or an Application-selected execution failure. It must not return partial normal Current state from incomplete or untrustworthy mandatory Current sources. | User accepted the output and failure boundary. The accepted shape preserves fresh immutable request state and the trustworthy-result rule while keeping Domain semantic internals inside Domain-owned construction. | W021 detailed Application shared-orchestration specs | not_required |
| D-004 | Legacy exact lookup-map Assembly trigger and input boundary | `decided` | D-001 | Legacy exact lookup-map Assembly runs only when an operation-specific use case decides that Legacy compatibility lookup is required. Shared application orchestration may load the Legacy exact lookup map through the distinct Legacy Archive source port. The input boundary is configured Legacy Archive roots, the Legacy Archive source port, accepted legacy issued-ID lexical mapping, exact legacy lookup keys, and source provenance needed for diagnostics and repair. The Legacy exact lookup map remains separate from Current Records state and is not loaded for Guidance operations. Existing specs may call this surface `Legacy lookup state`. | User accepted the revised trigger and input boundary while noting partial conceptual understanding. The wording now uses `Legacy exact lookup map` as the concrete W021 contract-surface term. Namespace scanning defines the map as a separate read-only exact lookup map from configured Legacy Archive roots. Resolver and validation specs cite the same configured map for accepted legacy issued-ID lookup. | W021 detailed Application shared-orchestration specs | not_required |
| D-005 | Legacy exact lookup-map Assembly output and failure boundary | `decided` | D-004 | Legacy exact lookup-map Assembly returns one complete trustworthy fresh immutable Legacy exact lookup map when invoked. If the required map cannot be built trustworthily from configured Legacy Archive roots, accepted legacy issued-ID mapping, source state, and required provenance, Application selects an execution failure. The assembly must not return a partial normal Legacy exact lookup map. | User accepted that an unbuildable required Legacy exact lookup map returns only an error. Namespace scanning forbids partial legacy lookup-map construction from invalid configured legacy roots. Validation and resolver authority preserve lookup outcomes inside a built map but treat unavailable required lookup state as execution failure. | W021 detailed Application shared-orchestration specs | not_required |
| D-006 | Current Records Snapshot transfer type boundary | `decided` | D-003 | Current Records Snapshot is one fresh immutable request-scoped transfer object for derived Current Records state. It contains current canonical identity lookup, parsed metadata and headings, Domain logical tree inputs, relation graph inputs, validation subjects, duplicate conflict groups, and source provenance or source references needed for diagnostics and exact retrieval. It does not have to retain every record body in memory. When `get_records` needs body content, the operation may read body content through the current source contract using snapshot provenance or source references. It excludes MCP request and response policy, concrete Infrastructure adapter access, Legacy exact lookup map contents, mutable cache, watcher, and process-wide index. | User corrected the earlier broad wording. The accepted boundary keeps the snapshot immutable without requiring all record bodies to be retained in the snapshot, avoiding an unnecessarily heavy request object. | W021 detailed Application shared-orchestration specs | not_required |
| D-007 | Legacy exact lookup-map transfer type boundary | `decided` | D-005 | Legacy exact lookup map is one immutable request-scoped transfer object for exact Legacy lookup. It contains accepted legacy issued IDs mapped to lookup results, source provenance for diagnostics, and distinct outcomes for unique, absent, duplicate, and unreadable archived sources. It excludes Current Records Snapshot, active current index, MCP request and response policy, concrete filesystem access, and mutable cache. Existing specs may call this surface `Legacy lookup state`. | User accepted the concise transfer boundary. The accepted wording keeps the map exact, request-scoped, immutable, and separate from Current Records, MCP, Infrastructure, and cache concerns. | W021 detailed Application shared-orchestration specs | not_required |
| D-008 | Expected semantic states versus execution failures | `decided` | D-002, D-003, D-004, D-005, D-006, D-007 | When a use case can explain the failed lookup, validation, or retrieval item as result data, the operation returns a normal response with warnings or diagnostics. When the request cannot build the required Current Records Snapshot, required Legacy exact lookup map, source enumeration, index, tree, or trustworthy failure location, Application selects an operation error and returns no partial normal response. | User accepted the simplified split. Explainable item-level failures remain normal semantic outcomes. Broken request foundations remain operation errors. | W021 detailed Application shared-orchestration specs | not_required |
| D-009 | Downstream decision, review, and ADR route | `decided` | D-001 through D-008 | No ADR is required for W021 because D-001 through D-008 refine provisional detailed-contract boundaries without superseding accepted architecture decisions. The next work is a decision Task that derives minimum Current Records Snapshot and Legacy exact lookup-map requirements from public use-case contracts before any detailed Specification authoring. Production implementation planning remains blocked. | User accepted the no-ADR route, then clarified that authoring must not start until public-use-case-derived request-state requirements and existing-spec gaps are decided. The decision ledger is complete, and downstream work must now start with T02 requirement derivation. | W021 Task graph and exact canonical targets | not_required |

### Status definitions

- `open`: Known judgment not yet selected.
- `in_discussion`: The one current user judgment.
- `decided`: Explicit accepted outcome persisted.
- `blocked`: Cannot proceed until a named dependency or authority is available.
- `deferred`: Explicitly routed outside this Work Item.
- `superseded`: Replaced by a later decision item.

### Current cursor

- Decision: none.
- Loop state: `complete`.
- Next action: Start DRMCP-TASK-MCP-021-02 to investigate public-use-case request-state requirements before authoring, then run DRMCP-TASK-MCP-021-03 as the confirmation decision loop.

### Expected downstream route

```text
DRMCP-TASK-MCP-021-01 decision ledger
  -> DRMCP-TASK-MCP-021-02 investigate public-use-case request-state requirements
  -> DRMCP-TASK-MCP-021-03 confirm snapshot requirements through decision loop
  -> classify existing-spec coverage and concrete gaps
  -> graph coordination when downstream synchronization, authoring, review, and closure owners are exactly known
  -> narrow Specification synchronization or provisional Specification authoring
  -> integrated independent review
  -> conditional finding route when required
  -> closure synchronization
  -> W022 consumption of the provisional Application shared-orchestration baseline
```

No downstream Task is materialized by this decision Task.

## Done condition

- Every owned decision item is `decided`, `deferred`, or validly `blocked`.
- The provisional detailed Specification target partition is selected.
- Current Records Snapshot Assembly inputs and outputs are decided at detailed-contract boundary level.
- Legacy exact lookup-map Assembly inputs and outputs are decided at detailed-contract boundary level.
- Current Records Snapshot and Legacy exact lookup-map transfer type boundaries are decided at detailed-contract boundary level.
- Expected semantic states and execution-failure boundaries are separated.
- Request-state freshness, immutability, visibility, and discard obligations are preserved.
- The downstream authoring route and ADR route are known or validly blocked.
- No ADR, Specification, implementation Task, production implementation, Go contract-shape probe, test, or independent review is authored by this Task.

## Verification

- Decision inventory terminal check: pass.
- ADR route check: no ADR required.
- Downstream route check: T02 must investigate public-use-case request-state requirements and T03 must confirm them before graph coordination materializes synchronization, authoring, integrated independent review, and closure synchronization owners.
- Scope check: no ADR, Specification, implementation Task, production implementation, Go contract-shape probe, test, or independent review was authored by this Task.

## Evidence

- DRMCP is non-operational. Design Records MCP cannot author this record, so filesystem authoring is the required fallback.
- No DRMCP namespace guide exists at `drmcp/records/guides`; authoring used PRODUCT authoring standards and active design-convergence workflow rules.
- W021 was untracked before this Task was created in the current worktree inspection.
- The initial scoped whitespace inspection passed before W021 Task creation.
- D-001 is decided. The accepted placement is `application-use-cases/shared-orchestration/` under the Application Use Cases contract scope.
- D-002 is decided. Current Records Snapshot Assembly is an Application shared orchestration input boundary over Current source access/results, validated dependencies, and Domain construction capabilities only.
- D-003 is decided. Current Records Snapshot Assembly returns either one complete trustworthy Current Records Snapshot or an Application-selected execution failure, not partial normal Current state.
- D-004 is decided using concrete W021 contract-surface wording. Legacy exact lookup-map Assembly runs only when an operation-specific use case decides that Legacy compatibility lookup is required, loads through the distinct Legacy Archive source port, uses configured Legacy Archive roots, accepted legacy issued-ID lexical mapping, exact legacy lookup keys, and source provenance, remains separate from Current Records state, and is not loaded for Guidance operations.
- D-005 is decided. Legacy exact lookup-map Assembly returns one complete trustworthy fresh immutable Legacy exact lookup map when invoked, or Application selects an execution failure when the required map cannot be built trustworthily. It does not return a partial normal Legacy exact lookup map.
- D-006 is decided. Current Records Snapshot is one fresh immutable request-scoped transfer object for derived Current Records state. It contains current identity lookup, parsed metadata and headings, Domain logical tree inputs, relation graph inputs, validation subjects, duplicate conflict groups, and source provenance or source references. It does not have to retain every record body in memory; body content may be read through the current source contract when exact retrieval needs it.
- D-007 is decided. Legacy exact lookup map is one immutable request-scoped transfer object for exact Legacy lookup. It maps accepted legacy issued IDs to lookup results and diagnostic provenance. It excludes Current Records Snapshot, active current index, MCP request and response policy, concrete filesystem access, and mutable cache.
- D-008 is decided. Explainable lookup, validation, or retrieval item failures remain normal semantic outcomes with warnings or diagnostics. Broken request foundations become operation errors with no partial normal response.
- D-009 is decided. No ADR is required for W021. The next work is T02 public-use-case request-state requirement investigation, followed by T03 confirmation decision loop before authoring. Production implementation planning remains blocked.
