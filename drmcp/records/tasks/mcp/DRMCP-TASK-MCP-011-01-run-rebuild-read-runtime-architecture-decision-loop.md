# DRMCP-TASK-MCP-011-01: Run rebuild read-runtime architecture decision loop

- **id**: DRMCP-TASK-MCP-011-01
- **status**: done
- **date**: 2026-06-30
- **work_item**: DRMCP-WORK-MCP-011
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 3d
- **depends_on**: []
- **outputs**:
  - DRMCP-TASK-MCP-011-01
  - DRMCP-WORK-MCP-011

## Goal

Resolve the architecture decisions required before planning a clean implementation of the DRMCP read runtime.

Keep every accepted answer recoverable from this Task instead of relying on chat history.

## Work

- Treat W003 through W008 as accepted semantic and fixture inputs.
- Treat W009 production structure, Task graph, file ownership, and extension seams as retired.
- Resolve repository-answerable facts without asking the user.
- Ask exactly one unresolved architecture decision per user turn.
- Explain the consequence and provide a recommendation when useful.
- Persist every explicit answer in the decision register before advancing.
- Keep ambiguous answers on the same decision ID.
- Detect conflicts with accepted ADRs, Requirements, Specifications, or earlier decisions.
- Update dependency states and the current cursor after each accepted answer.
- Route each decided item as `required`, `covered`, `not_required`, or `blocked` for ADR handling.
- Identify current Specification targets without inventing unsupported refs or file paths.
- Stop before production implementation, implementation Task authoring, or independent review.

## Done condition

- Every known required decision has one decision-register row.
- At most one decision is `in_discussion`.
- Every explicit user answer is persisted before the cursor advances.
- Every required decision is `decided`, `deferred`, or validly `blocked`.
- Every decided item has an ADR route or an explicit ADR-routing blocker.
- Expected Specification targets are identified or blocked for a named reason.
- No accepted answer exists only in chat.
- The loop status is `decision_complete`.
- No production implementation or implementation Task authoring has begun.

## Verification

- Compare the decision register with `DRMCP-WORK-MCP-011` scope and completion conditions.
- Compare accepted answers with `DRMCP-ADR-MCP-001` and `DRMCP-REQ-MCP-001`.
- Confirm that W003 through W008 boundaries are consumed without reopening accepted operation behavior.
- Confirm that W009 and W010 implementation assumptions are not treated as authority.
- Confirm that no more than one decision is `in_discussion`.
- Confirm that every `decided` item contains a concise summary, dependencies, ADR route, and expected canonical target.
- Re-read the written decision-register section after every update.
- Confirm that the current cursor selects the first unblocked non-terminal decision.

## Evidence

### Scope baseline

| input | disposition |
|---|---|
| `DRMCP-ADR-MCP-001` | Accepted semantic and compatibility direction. Architecture-specific internal boundaries remain open. |
| `DRMCP-REQ-MCP-001` | Source Requirement and required runtime outcomes. |
| `DRMCP-WORK-MCP-003` through `DRMCP-WORK-MCP-006` | Accepted discovery, query, retrieval, resolver, validation, diagnostic, and path-exposure contracts. |
| `DRMCP-WORK-MCP-007` | Accepted per-file and Topics graph validation ownership split. |
| `DRMCP-WORK-MCP-008` | Accepted fixture meanings and coverage. Physical fixture placement is non-authoritative. |
| `DRMCP-WORK-MCP-009` | Retired implementation plan. Production code and architecture are not reused. |
| `DRMCP-WORK-MCP-010` | Not started. Requires replacement current-runtime architecture before rebaseline. |
| `DRMCP-WORK-SPEC-001/002` | Retained semantic owners. Concrete implementation integration remains open. |

### Decision confirmation loop

Loop status: `decision_complete`

Current decision: `none`

| ID | Topic | Status | Depends on | Decision summary | ADR route | Canonical target |
|---|---|---|---|---|---|---|
| D-001 | MCP tool routing and application use-case boundary | recorded | — | MCP transport keeps the public request and response schemas for `list_records`, `get_records`, `resolve_reference`, and `validate_records`. Each of those four operations invokes a dedicated transport-neutral application use case. The MCP adapter maps between MCP schemas and application types. Authoring-guidance and authoring-transaction operations are outside this decision. No generic App Router is introduced. | required | `spec:drmcp.design_records_mcp.responsibility_boundary`, `spec:drmcp.design_records_mcp.tools.overview` |
| D-002 | Composition root and runtime lifecycle | recorded | D-001 | The composition root loads and validates configuration at server startup. Each invocation of `list_records`, `get_records`, `resolve_reference`, or `validate_records` rebuilds current and configured legacy indexes from the filesystem and uses one immutable request-scoped snapshot. Filesystem changes are therefore visible on the next invocation of one of those operations. Startup-invalid configuration prevents server start; post-start scan or index failures fail only the current operation. | required | `spec:drmcp.design_records_mcp.overview`, `spec:drmcp.design_records_mcp.namespace_scanning` |
| D-003 | Domain, parsed-source, index-state, diagnostic, and public-projection separation | recorded | D-001 | Raw or parsed source state, canonical index state, diagnostics, and public operation projections use separate internal types. Duplicate conflicts and invalid-but-addressable sources are explicit index states. No single nullable Record type spans all layers. | required | `spec:drmcp.design_records_mcp.schema.overview`, `spec:drmcp.design_records_mcp.schema.record_model` |
| D-004 | Configuration, filesystem, source loading, and index ports and adapters | recorded | D-001, D-003 | Application use cases depend on narrow ports for configuration access, source enumeration, and source reading. Filesystem and configuration behavior live in outer adapters. Parsers and index builders remain pure where possible. No generic all-purpose repository interface is introduced. | required | `spec:drmcp.design_records_mcp.namespace_scanning`, `spec:drmcp.design_records_mcp.responsibility_boundary` |
| D-005 | Current-index and optional legacy-index runtime composition | recorded | D-002, D-004 | One request-scoped snapshot contains a separate current active index and optional legacy exact lookup map. Current operations query only the current index. Legacy exact retrieval queries only the legacy map. Resolver orchestration consults both in accepted current-first order. No generic merged index hides source-family differences. | covered | `spec:drmcp.design_records_mcp.namespace_scanning`, `spec:drmcp.design_records_mcp.resolver` |
| D-006 | Per-file and Topics graph validation pipeline integration | recorded | D-003, D-005 | Standalone `validate_records` reads the configured filesystem and builds a fresh request-scoped snapshot. Individual validators perform no filesystem I/O. Per-source validation runs before relation and Topics graph validation. Any future persisted-write caller must rebuild validation input from persisted files instead of reusing candidate or pre-write state. Current authoring-transaction integration is deferred to `DRMCP-REQ-MCP-002`. Validators return findings as data; aggregation, ordering, deduplication, and MCP formatting occur outside validators. | required | `spec:drmcp.design_records_mcp.tools.validate_records`, `spec:drmcp.design_records_mcp.schema.diagnostics` |
| D-007 | Overall component structure and dependency direction | recorded | D-001, D-002, D-003, D-004, D-005, D-006 | The W011 read-runtime slice for `list_records`, `get_records`, `resolve_reference`, and `validate_records` uses a composition root, one inbound MCP adapter, an application layer with operation use cases and request-snapshot orchestration, a transport- and infrastructure-independent core, and outbound filesystem and configuration adapters. Dependencies point inward. The composition root wires components and owns no business behavior. | required | `spec:drmcp.design_records_mcp.overview`, `spec:drmcp.design_records_mcp.responsibility_boundary` |
| D-008 | Operation input, output, error, and diagnostic ownership | recorded | D-007 | Each application use case owns its typed operation input and output contract and projects core values into the accepted public response meaning. Core owns domain values, lookup states, and validation findings, not complete tool responses. Expected semantic states remain result data; failures that prevent a trustworthy result are execution errors. MCP owns protocol wrappers, structural JSON validation, decode, invocation, and encode without redesigning operation semantics. | required | `spec:drmcp.design_records_mcp.tools.overview`, operation-specific tool specs, `spec:drmcp.design_records_mcp.schema.diagnostics` |
| D-009 | Concrete Go package boundaries and placement | recorded | D-007, D-008 | The W011 read-runtime slice uses separate `app`, `core`, and `adapters` package trees under `internal`, with one application package for each of `list_records`, `get_records`, `resolve_reference`, and `validate_records`, shared request-snapshot orchestration, core packages split by semantic responsibility, and distinct MCP, filesystem, and config adapters. The package tree is complete for those four operations, not for the complete Design Records MCP server. Authoring-guidance and authoring-transaction package placement remain outside W011. The composition root remains in `cmd/design-records-mcp`. No compatibility wrapper preserves the retired `internal/designrecords` API. | required | `spec:drmcp.design_records_mcp.overview`, `spec:drmcp.design_records_mcp.responsibility_boundary` |

### Canonical recording

| decision | ADR authority | primary Specification reflection |
|---|---|---|
| D-001 | `DRMCP-ADR-MCP-003` | `spec:drmcp.implementation`, `spec:drmcp.design_records_mcp.responsibility_boundary`, `spec:drmcp.design_records_mcp.tools.overview` |
| D-002 | `DRMCP-ADR-MCP-002` | `spec:drmcp.implementation`, `spec:drmcp.design_records_mcp.overview`, `spec:drmcp.design_records_mcp.namespace_scanning` |
| D-003 | `DRMCP-ADR-MCP-004` | `spec:drmcp.implementation`, `spec:drmcp.design_records_mcp.schema.overview`, `spec:drmcp.design_records_mcp.schema.record_model` |
| D-004 | `DRMCP-ADR-MCP-003` | `spec:drmcp.implementation`, `spec:drmcp.design_records_mcp.responsibility_boundary`, `spec:drmcp.design_records_mcp.namespace_scanning` |
| D-005 | `DRMCP-ADR-MCP-001` | `spec:drmcp.implementation`, `spec:drmcp.design_records_mcp.namespace_scanning`, `spec:drmcp.design_records_mcp.resolver` |
| D-006 | `DRMCP-ADR-MCP-005` | `spec:drmcp.implementation`, `spec:drmcp.design_records_mcp.tools.validate_records`, `spec:drmcp.design_records_mcp.schema.diagnostics` |
| D-007 | `DRMCP-ADR-MCP-003` | `spec:drmcp.implementation`, `spec:drmcp.design_records_mcp.overview`, `spec:drmcp.design_records_mcp.responsibility_boundary` |
| D-008 | `DRMCP-ADR-MCP-004` | `spec:drmcp.implementation`, `spec:drmcp.design_records_mcp.tools.overview`, `spec:drmcp.design_records_mcp.schema.diagnostics` |
| D-009 | `DRMCP-ADR-MCP-006` | `spec:drmcp.implementation` |

- T06 independently confirmed the ADR routing and Specification reflection for D-001 through D-009.
- D-001 through D-009 are `recorded`.
- No unresolved W011 architecture decision remains.
- Current authoring-transaction integration remains deferred to `DRMCP-REQ-MCP-002`.

### D-001 accepted decision

- MCP transport keeps the public request and response schemas required by `list_records`, `get_records`, `resolve_reference`, and `validate_records`.
- Each of those four read-runtime operations invokes one dedicated application use case.
- Application input and output types remain independent from MCP transport types.
- The MCP adapter maps MCP requests to application inputs and application results to MCP responses for those four operations.
- Authoring-guidance and authoring-transaction operations are outside D-001.
- No generic App Router is introduced for the four read-runtime operations.
- MCP handlers for those operations do not call domain or infrastructure adapters directly.

ADR route: `required`.

Reason:

- the choice establishes a durable application boundary and dependency direction;
- a generic router would duplicate the MCP SDK dispatch without a separate responsibility;
- transport-neutral application types prevent MCP schemas from becoming domain authority.

### D-002 accepted decision

- One composition root loads and validates configuration at server startup.
- The composition root wires scanners, parsers, index builders, use cases, and the MCP adapter.
- Each invocation of `list_records`, `get_records`, `resolve_reference`, or `validate_records` rebuilds the current active index from configured current roots.
- Each invocation of one of those operations also rebuilds the configured legacy lookup map when legacy roots are enabled.
- One invocation of one of those operations uses one immutable request-scoped index snapshot from start to finish.
- Filesystem changes become visible on the next invocation of one of those operations.
- Tool execution does not mutate or incrementally patch a shared process-wide index.
- Configuration that is invalid at startup prevents the server from starting.
- A root, scan, or index failure caused after startup fails the current operation without returning a partial normal response.
- Shutdown releases only resources owned by the composition root.

ADR route: `required`.

Reason:

- request-scoped rebuilding preserves filesystem freshness without a watcher or cache-invalidation protocol;
- one immutable snapshot keeps each operation internally consistent;
- separating startup configuration failure from later filesystem failure preserves truthful failure ownership.

### D-003 accepted decision

- Raw or parsed source state uses dedicated internal types.
- Canonical index entries and index conflicts use dedicated index-state types.
- Invalid-but-addressable sources remain explicit index states instead of malformed successful records.
- Diagnostics remain separate from source and index structures.
- `list_records`, `get_records`, and resolver outputs are application projections built from internal state.
- No single nullable `Record` type is shared across transport, application, domain, parsing, indexing, and validation layers.

ADR route: `required`.

Reason:

- each type has one authority and lifecycle;
- explicit index states preserve accepted conflict and invalid-source behavior;
- public response changes do not force parser or index model changes.

### D-004 accepted decision

- Application use cases depend on narrow ports for configuration access, source enumeration, and source reading.
- Filesystem traversal and file reading live in outer adapters.
- Concrete configuration loading lives in an outer adapter.
- Parsers and index builders remain pure internal services where practical.
- Application use cases orchestrate ports and pure services without importing concrete filesystem packages.
- No generic repository interface combines unrelated source, configuration, indexing, and validation operations.

ADR route: `required`.

Reason:

- narrow ports isolate infrastructure changes;
- pure parsing and indexing remain easy to test;
- a future database-backed adapter can implement the same source access contracts without changing application use cases;
- source identity, path diagnostics, and freshness semantics still require explicit adapter contracts.

### D-005 accepted decision

- One request-scoped snapshot contains a current active index and an optional legacy exact lookup map.
- The two structures remain separate.
- Current list and current exact retrieval query only the current index.
- Legacy exact retrieval queries only the legacy lookup map.
- Resolver orchestration consults current first and legacy second under the accepted fallback rules.
- No generic merged index API hides current and legacy behavioral differences.

ADR route: `covered` by `DRMCP-ADR-MCP-001` unless later architecture decisions introduce a conflicting runtime rule.

Reason:

- the accepted contract already requires current and legacy separation;
- separate structures preserve source-specific parsing, identity, and failure behavior;
- the request snapshot still provides one consistent operation boundary.

### D-006 accepted decision

- Standalone `validate_records` reads the configured filesystem and builds a fresh request-scoped snapshot.
- Source enumeration and source reading happen before validation passes start.
- Individual validators perform no filesystem I/O and do not rescan independently.
- Per-source syntax, metadata, identity, and document-shape detectors consume retained source state from the snapshot.
- Relation and Topics graph validation run afterward against the complete current index and optional legacy lookup map.
- Validators return findings as data and do not format MCP responses.
- Validators do not call public tools or invoke one another through transport boundaries.
- Any future caller that validates after persisted filesystem writes discards candidate and pre-write state.
- That caller rebuilds validation input from persisted files before using this validation architecture.
- Current authoring-transaction timing and current-format integration are deferred to `DRMCP-REQ-MCP-002`.
- Filesystem writes become visible to the next invocation of a W011 read-runtime operation.
- Diagnostic aggregation, stable ordering, deduplication, and MCP response formatting occur outside individual validators.

ADR route: `required`.

Reason:

- filesystem freshness belongs to orchestration;
- validation rules remain pure and independently testable;
- one snapshot preserves consistency inside an invocation;
- persisted-write reloading, when integrated by a future caller, validates actual persisted state rather than intended state.

### Decision graph correction

The previous D-007 attempted to decide error and response ownership before the overall architecture existed.
That ordering was too abstract.

D-001 through D-006 remain accepted architectural constraints.
They do not complete the overall component model.

The remaining order is now:

1. D-007 defines the complete component structure and dependency direction.
2. D-008 assigns operation input, output, error, and diagnostic ownership inside that structure.
3. D-009 maps the accepted structure to concrete Go packages.

### D-007 accepted decision

This component structure applies to the W011 read-runtime slice for `list_records`, `get_records`, `resolve_reference`, and `validate_records`.
It does not decide the internal architecture of authoring-guidance or authoring-transaction operations.

Component structure:

```text
composition root
├─ inbound adapter: MCP
├─ application
│  ├─ operation use cases
│  └─ request-snapshot orchestration
├─ core
│  ├─ source and parsed models
│  ├─ parsers
│  ├─ current index
│  ├─ legacy lookup
│  ├─ resolver
│  └─ validation
└─ outbound adapters
   ├─ filesystem
   └─ configuration
```

Dependency direction:

```text
MCP adapter ───────┐
filesystem adapter ├──> application ───> core
config adapter ────┘
composition root wires every component
core depends on no outer component
```

Accepted rules:

- the composition root owns construction and wiring only;
- the MCP adapter is the only inbound adapter for the four W011 read-runtime operations;
- application use cases for those four operations own operation sequencing and request-snapshot orchestration;
- core owns parsing, indexing, resolution, and validation logic without MCP, filesystem, or config dependencies;
- application-owned narrow ports are implemented by filesystem and configuration adapters;
- outbound adapters do not call use cases;
- core does not import application or adapter packages;
- operation schema and error ownership remain undecided until D-008;
- concrete Go package names and directory placement remain undecided until D-009.

ADR route: `required`.

Reason:

- the structure assigns orchestration, pure logic, and external effects to distinct owners;
- inward dependencies preserve replaceability and testability;
- delaying operation contracts and package names prevents premature coupling.

### D-008 accepted decision

- Each application use case owns one typed input and one typed output contract for its operation.
- The output contract follows the accepted public tool response meaning, including records, resolver status, warnings, validation diagnostics, and summaries.
- Core returns domain values, lookup states, and validation findings but does not define complete tool responses.
- Application use cases project core results into operation outputs.
- Semantic invalidity and expected not-found or conflict states remain typed operation result data where the accepted tool contract requires them.
- Failures that prevent the operation from constructing a trustworthy result return an application execution error.
- The MCP adapter derives or declares JSON schemas from application operation contracts and handles decode, structural validation, invocation, and encode.
- The MCP adapter does not redesign fields, statuses, diagnostics, or error semantics.
- MCP SDK and protocol wrappers remain adapter-owned and do not enter application or core types.

ADR route: `required`.

Reason:

- operation meaning belongs with orchestration rather than transport;
- core stays reusable without becoming aware of public tool envelopes;
- MCP remains a thin protocol adapter while still exposing the accepted public schema.

### D-009 accepted decision

Repository fact:

- the retired implementation placed most read logic in one `internal/designrecords` package and MCP transport in `internal/designrecordsmcp`;
- that package structure is non-authoritative and is not preserved.

Accepted package structure for the four W011 read-runtime operations:

```text
drmcp/src/
├─ cmd/design-records-mcp/
│  └─ main.go
└─ internal/
   ├─ app/
   │  ├─ snapshot/
   │  ├─ listrecords/
   │  ├─ getrecords/
   │  ├─ resolvereference/
   │  └─ validaterecords/
   ├─ core/
   │  ├─ records/
   │  ├─ indexing/
   │  ├─ resolving/
   │  └─ validation/
   └─ adapters/
      ├─ mcp/
      ├─ filesystem/
      └─ config/
```

Accepted rules:

- this package tree is the complete W011 inventory for `list_records`, `get_records`, `resolve_reference`, and `validate_records`;
- this package tree is not a complete inventory for the Design Records MCP server;
- authoring-guidance and authoring-transaction package placement remain outside W011;
- `cmd/design-records-mcp` is the composition root and contains only startup, wiring, run, and shutdown behavior;
- each of the four W011 operations has its own application package with its input, output, use case, and focused tests;
- request-snapshot orchestration is a shared application package rather than duplicated across use cases;
- core is split by semantic responsibility, not by every individual type or file;
- parsers and source models live with `core/records` unless implementation evidence later justifies a separate parsing package;
- current and legacy index structures live together under `core/indexing` while remaining distinct types;
- resolver logic lives under `core/resolving`;
- per-source and graph validators live under `core/validation` and may use core record and index types;
- MCP, filesystem, and configuration implementations live under separate adapter packages;
- no compatibility wrapper preserves the retired `internal/designrecords` package API;
- package cycles are prohibited and imports follow the D-007 dependency direction.

ADR route: `required`.

Reason:

- package boundaries now mirror the accepted component boundaries;
- operation packages keep public orchestration contracts focused;
- core packages remain cohesive without fragmenting every type into its own package;
- the retired monolithic package does not constrain the rebuild.

### Decision-loop completion

- D-001 through D-009 are all `recorded`.
- No decision remains `open`, `in_discussion`, `decided`, or blocked.
- Every accepted user answer is persisted in this Task.
- ADR routing and Specification reflection are complete and independently confirmed by T06.
- Current authoring-transaction integration remains deferred to `DRMCP-REQ-MCP-002`.
- No production implementation or implementation Task authoring has begun.

### Downstream phase ownership

| phase | future writer |
|---|---|
| ADR routing and authoring | Dedicated ADR Task after this loop reaches `decision_complete`. |
| Specification synchronization | Dedicated Specification Task after required ADRs are ready. |
| Independent design review | Read-only reviewer after Specification synchronization. |
| Finding correction | Dedicated correction writer for named findings only. |
| Finding closure re-review | Independent reviewer. |
| Lifecycle closure | Dedicated closure-synchronization Task. |

### Creation evidence

- The exact `drmcp/records/work-items/mcp/` inventory contained W001 through W010 and no W011.
- The exact `drmcp/records/tasks/mcp/` inventory contained no `DRMCP-TASK-MCP-011-*` record.
- W011 and T01 therefore use the next valid Work Item and Task sequences.
- DRMCP authoring transactions are non-operational under the current agent-authoring policy.
- Filesystem authoring is the required fallback.
