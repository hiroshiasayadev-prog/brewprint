# TRV-TASK-SPEC-001-02: Decide TRV app-local design boundary

- **id**: TRV-TASK-SPEC-001-02
- **status**: done
- **date**: 2026-07-02
- **work_item**: TRV-WORK-SPEC-001
- **task_type**: decision
- **estimate**: 2d
- **depends_on**:
  - TRV-TASK-SPEC-001-01
- **outputs**:
  - TRV-TASK-SPEC-001-02

## Goal

Produce one bounded decision ledger for the TRV app-local design boundary.

## Work

- Run one interactive decision loop.
- Ask exactly one unresolved decision per user turn.
- Persist each explicit answer before advancing the cursor.
- Keep PRODUCT-owned semantics fixed and outside the decision inventory.
- Treat PRODUCT-TASK-SPEC-021-02 and PRODUCT-INV-SPEC-009 as historical candidate inputs only.
- Decide how TRV realizes the fixed semantic contract as an application.
- Record one concise outcome, reason, canonical target, ADR-routing state, and downstream owner for every terminal item.
- Stop downstream progression when an item remains materially ambiguous.

This Task must not:

- author a Requirement, Investigation, ADR, or Specification;
- change checklist wording or PRODUCT-owned semantic behavior;
- perform ADR routing or canonical authoring;
- materialize implementation Tasks or an executor prompt;
- design current DRMCP integration;
- perform implementation, review, correction, synchronization, stage, or commit work.

### Decision inventory

| ID | Stage | Topic | Status | Depends on | Expected canonical target | ADR route |
|---|---|---|---|---|---|---|
| D-001 | Requirement | TRV application Requirement identity and Required Outcome | `decided` | none | TRV-REQ-SPEC-001 | `unknown` |
| D-016 | Decomposition | Architecture, contract, detailed Specification, and implementation Work Item boundaries | `decided` | D-001 | TRV-WORK-SPEC-001 and child Work Item boundaries | `unknown` |
| D-017 | Architecture | Application architecture style and component responsibility boundary | `decided` | D-016 | TRV architecture Specification and possible ADR | `candidate` |
| D-005 | Architecture | Model, provider, and runtime ownership | `decided` | D-017 | TRV runtime Specification and possible ADR | `candidate` |
| D-011 | Architecture | Logical component boundary, dependency policy, and source ownership | `decided` | D-017, D-005 | TRV architecture Specification and possible ADR | `candidate` |
| D-002 | Contract | External application interface boundary | `decided` | D-011 | TRV interface Specification and possible ADR | `candidate` |
| D-003 | Contract | Task input, repository root, path safety, and Task parsing boundary | `decided` | D-002 | TRV input contract Specification | `candidate` |
| D-009 | Contract | MCP server identity, tool identity, and input/output envelope | `decided` | D-002, D-003 | TRV transport Specification and possible ADR | `candidate` |
| D-010 | Contract | Human-action indication and caller interaction boundary | `decided` | D-009 | TRV caller-integration Specification | `reuse PRODUCT-ADR-SPEC-017` |
| D-014 | Contract | Compatibility and future DRMCP migration boundary | `decided` | D-002, D-003, D-009, D-010 | TRV compatibility Specification and possible ADR | `candidate` |
| D-004 | Detailed Specification | Checklist discovery, criterion identity, and prompt composition | `decided` | D-014 | TRV checklist adapter Specification | `candidate` |
| D-006 | Detailed Specification | Model invocation, structured output, and result validation | `decided` | D-004, D-005, D-009 | TRV evaluation runtime Specification | `candidate` |
| D-007 | Detailed Specification | Retry, timeout, malformed response, and concrete failure handling | `decided` | D-006, D-009 | TRV failure-handling Specification | `candidate` |
| D-008 | Detailed Specification | Configuration and startup validation | `decided` | D-005, D-007 | TRV configuration Specification | `candidate` |
| D-012 | Detailed Specification | Launcher, external `mcp-proxy`, and Windows deployment boundary | `decided` | D-008, D-009, D-011 | TRV deployment Specification | `candidate` |
| D-013 | Detailed Specification | Test surface, fake model runtime, process smoke, and verification boundary | `decided` | D-003, D-004, D-005, D-006, D-007, D-008, D-009, D-010, D-011, D-012, D-014 | TRV verification Specification | `candidate` |
| D-015 | Handoff | Later TRV implementation Work Item identity and handoff | `decided` | D-001, D-002, D-003, D-004, D-005, D-006, D-007, D-008, D-009, D-010, D-011, D-012, D-013, D-014, D-016, D-017 | Later TRV implementation-readiness gate and decomposition authority | `not_applicable` |

### Decision records

#### D-001: TRV application Requirement identity and Required Outcome

- **Status**: `decided`
- **Selected outcome**: Create `TRV-REQ-SPEC-001` as the technology-neutral Requirement for delivering the Task Responsibility Validator application.
- **Required Outcome boundary**:
  - provide a TRV application that realizes `spec:product.responsibility_boundary_validator`;
  - require reviewed application architecture before contract authoring;
  - require a reviewed external and application contract before detailed Specification authoring;
  - require reviewed implementation-ready detailed Specifications before production implementation;
  - require production implementation and verification through a later implementation Work Item;
  - exclude current DRMCP integration.
- **Technology boundary**: Do not fix standalone MCP, Go, Ollama, exact source layout, or deployment mechanics in the Requirement.
- **Reason**: The stable need is delivery of the TRV application. Architecture, contract, detailed design, and implementation are resolution stages rather than independent needs.
- **Canonical target**: `TRV-REQ-SPEC-001`.
- **ADR route**: `unknown`; final routing remains downstream of T04.
- **Downstream owner**: T04 for Requirement-authoring and Work Item decomposition routing.

#### D-016: Architecture, contract, detailed Specification, and implementation Work Item boundaries

- **Status**: `decided`
- **Selected outcome**: Resolve one TRV application Requirement through three sequential design Work Items and one later implementation Work Item.
- **Design sequence**:
  1. application architecture;
  2. external and application contract;
  3. implementation-ready detailed Specifications.
- **Implementation sequence**: Create a separate implementation Work Item only after the three design boundaries are reviewed and closed.
- **Parent boundary**: Retain `TRV-WORK-SPEC-001` as the coordinating app-local design boundary. Child Work Item identities and execution relations are materialized downstream.
- **Child ownership**: Each child Work Item owns its internal Task graph, canonical artifacts, integrated review, and closure.
- **Reason**: Architecture, contract, and detailed Specifications have different canonical outputs and completion judgments. Separate Work Items prevent detailed authoring from starting on an unstable architecture or contract.
- **Canonical target**: `TRV-WORK-SPEC-001` plus child Work Items created through a later decomposition owner.
- **ADR route**: `unknown`; final routing remains downstream of T04.
- **Downstream owner**: T04 for exact child Work Item decomposition, dependencies, execution tracking, and review order.

#### D-017: Application architecture style and component responsibility boundary

- **Status**: `decided`
- **Selected outcome**: Use a ports-and-adapters architecture with an application core independent of MCP, filesystem access, checklist storage, and the model provider.
- **Application core ownership**:
  - validation orchestration;
  - application-owned ports and outcome models;
  - criterion-result completeness validation;
  - mechanical overall-result construction under the PRODUCT semantic contract.
- **Adapter ownership**:
  - one inbound adapter owns MCP transport and tool projection;
  - outbound adapters own Task file access, checklist loading, and model-provider calls;
  - startup composition owns configuration validation and adapter wiring.
- **Dependency direction**: Adapters depend on application-owned ports and models. The application core does not import transport, filesystem, or provider-specific packages.
- **Mechanics boundary**: Provider HTTP, retry, timeout, and decoding remain outside the application core.
- **Reason**: MCP and the first model provider are replaceable delivery choices. Keeping them outside the core preserves the validation use case across standalone operation and later integration.
- **Canonical target**: TRV architecture Specification and a possible architecture ADR.
- **ADR route**: `candidate`; final routing remains downstream of T04.
- **Downstream owner**: T04 for architecture Work Item decomposition and later ADR routing.

#### D-005: Model, provider, and runtime ownership

- **Status**: `decided`
- **Selected outcome**: TRV owns a provider-neutral model-evaluation port and one Ollama HTTP adapter. The Ollama runtime and model files remain externally managed on a separate server.
- **Runtime boundary**:
  - TRV calls a configured Ollama base URL over HTTP;
  - TRV does not install, start, stop, update, supervise, or health-manage Ollama as a process;
  - TRV does not download, store, update, or delete model files;
  - network reachability and Ollama availability are deployment preconditions rather than TRV lifecycle responsibilities.
- **Configuration boundary**: The Ollama base URL and concrete model name are deployment configuration, not application identity.
- **Provider boundary**: The first implementation supports Ollama only. Additional providers require a later app-local design decision but do not change the application core port.
- **Prior-decision treatment**: This preserves PRODUCT-TASK-SPEC-021-02 D-005 and D-016 while correcting the ambiguous word `local` to the actual remote-server deployment topology.
- **Reason**: The same runtime ownership boundary was already accepted in W021. TRV only needs to adopt it under the correct app-local authority and record the separate-server clarification.
- **Canonical target**: TRV runtime architecture Specification and a possible architecture ADR.
- **ADR route**: `candidate`; final routing remains downstream of T04.
- **Downstream owner**: T04 for architecture Work Item decomposition and later ADR routing.

#### D-011: Logical component boundary, dependency policy, and source ownership

- **Status**: `decided`
- **Selected outcome**: Use five logical ownership areas without introducing a separate domain layer.
- **Logical areas**:
  1. **application core** — validation use case, ports, outcome types, criterion-set validation, and logical-AND aggregation;
  2. **record/checklist adapter** — safe Task reading, minimal Task metadata parsing, checklist loading, and prompt-input assembly;
  3. **Ollama HTTP adapter** — remote Ollama request/response translation and provider-specific execution mechanics;
  4. **MCP adapter** — external tool projection and transport translation only;
  5. **startup and dependency wiring** — configuration loading, startup validation, concrete dependency construction, server wiring, and process lifecycle.
- **Dependency rule**:
  - adapters and startup wiring depend inward on application-owned contracts;
  - adapters do not call each other directly;
  - the core contains no MCP, filesystem, HTTP, environment-variable, or process-lifecycle knowledge.
- **Detailed-design boundary**: Exact Go package names, directories, files, symbols, and constructor shapes remain for the detailed Specification Work Item.
- **Reason**: This preserves explicit adapter ownership without adding a ceremonial domain layer or mixing unrelated infrastructure responsibilities.
- **Canonical target**: TRV architecture Specification and a possible architecture ADR.
- **ADR route**: `candidate`; final routing remains downstream of T04.
- **Downstream owner**: T04 for architecture Work Item decomposition and later ADR routing.

#### D-002: External application interface boundary

- **Status**: `decided`
- **Selected outcome**: Preserve the W021 decision to expose TRV through a dedicated MCP server.
- **Interface boundary**:
  - MCP is the required external interface for AI agents;
  - the application serves MCP over stdio;
  - an optional thin CLI may exist only for development and smoke verification;
  - TRV does not own an HTTP server or current DRMCP integration.
- **Prior-decision treatment**: PRODUCT-TASK-SPEC-021-02 D-002 is adopted under TRV authority because the new architecture does not conflict with it.
- **Reason**: The MCP adapter cleanly projects the application use case while remaining replaceable and independent of the core.
- **Canonical target**: TRV external-interface contract Specification and a possible interface ADR.
- **ADR route**: `candidate`; final routing remains downstream of T04.
- **Downstream owner**: T04 for contract Work Item decomposition and later ADR routing.

#### D-003: Task input, repository root, path safety, and Task parsing boundary

- **Status**: `decided`
- **Selected outcome**: Preserve the W021 decision that one invocation accepts exactly one repository-root-relative Task file path.
- **Input boundary**:
  - the caller supplies one Task path relative to the configured repository root;
  - TRV reads the persisted Task and obtains its declared `task_type` from the record;
  - Task body content and public record ID are not alternative primary inputs;
  - current DRMCP record resolution remains outside scope.
- **Safety boundary**: Absolute paths and repository-root escape are rejected. Exact normalization and symlink handling belong to the detailed Specification.
- **Prior-decision treatment**: PRODUCT-TASK-SPEC-021-02 D-003 and D-019 are adopted under TRV authority because the app-local architecture does not conflict with them.
- **Reason**: A single path input keeps the temporary standalone contract narrow and avoids importing record-index responsibilities.
- **Canonical target**: TRV input contract Specification.
- **ADR route**: `candidate`; final routing remains downstream of T04.
- **Downstream owner**: T04 for contract Work Item decomposition and later ADR routing.

#### D-009: MCP server identity, tool identity, input schema, and tagged output envelope

- **Status**: `decided`
- **Selected outcome**:
  - server name: `task-responsibility-validator`;
  - tool name: `validate_task_responsibility`;
  - input: one required string field, `task_path`, with no additional properties.
- **Output envelope**: Use one `outcome` discriminator with exactly three outcome classes:
  - `semantic_evaluation` with criterion results, derived `overall_compliant`, and `human_action_required`;
  - `structural_precondition_failure` with one failure category and concise message;
  - `execution_failure` with one failure category, concise message, and attempt count when model invocation began.
- **Process boundary**: Per-Task failures remain structured tool results and do not terminate the MCP process.
- **Protocol boundary**: MCP protocol version is negotiated through normal initialization rather than fixed as TRV application identity.
- **Reason**: The explicit names describe the narrow responsibility boundary and preserve one stable interface for both validation phases.
- **Canonical target**: TRV MCP transport and result-contract Specification plus a possible interface ADR.
- **ADR route**: `candidate`; final routing remains downstream of T04.
- **Downstream owner**: T04 for contract Work Item decomposition and later ADR routing.

#### D-010: Human-action indication and caller interaction boundary

- **Status**: `decided`
- **Selected outcome**: Preserve the W021 caller-owned human judgment boundary.
- **TRV responsibility**: A non-compliant semantic result returns violated criteria, concise reasons, and `human_action_required: true`.
- **Caller responsibility**:
  - present violations to the human;
  - obtain acceptance with a reason or rejection;
  - persist accepted violations and reasons in Task Evidence;
  - route rejected work back to correction or responsibility-boundary reconsideration.
- **Excluded TRV behavior**: No acceptance/rejection tool, Evidence mutation, Task rewrite, release decision, or automatic return-route selection.
- **Prior-decision treatment**: PRODUCT-TASK-SPEC-021-02 D-013 through D-015 are adopted under TRV authority.
- **Reason**: Human authorization is workflow authority, not validator behavior.
- **Canonical target**: TRV caller-integration contract Specification.
- **ADR route**: `reuse PRODUCT-ADR-SPEC-017`.
- **Downstream owner**: T04 for contract Work Item decomposition.

#### D-014: Compatibility and future DRMCP migration boundary

- **Status**: `decided`
- **Selected outcome**: Preserve semantic compatibility across future DRMCP integration without guaranteeing standalone source, tool-name, path-input, environment-variable, or transport compatibility.
- **Stable behavior**:
  - criterion-level boolean results and concise reasons;
  - exact criterion-set validation;
  - logical-AND overall result;
  - distinct semantic, structural, and execution outcomes;
  - caller-owned human judgment;
  - post-authoring and post-Evidence invocation points.
- **Replaceable boundary**: Future DRMCP integration may replace MCP transport, Task-path input, configuration shape, adapters, and the implementation itself.
- **Prior-decision treatment**: PRODUCT-TASK-SPEC-021-02 D-019 is adopted under TRV authority.
- **Reason**: The standalone app is temporary, while its semantic behavior is the durable contract.
- **Canonical target**: TRV compatibility contract Specification and a possible compatibility ADR.
- **ADR route**: `candidate`; final routing remains downstream of T04.
- **Downstream owner**: T04 for contract Work Item decomposition and later ADR routing.

#### D-004: Checklist discovery, criterion identity, and prompt composition

- **Status**: `decided`
- **Selected outcome**: Preserve the fixed repository-root-relative checklist discovery and composition contract from W021.
- **Composition order**:
  1. evaluator instructions;
  2. common checklist;
  3. the declared `task_type` checklist;
  4. the full Task record after an explicit separator.
- **Discovery boundary**: Use the accepted fixed paths under `skills/task-responsibility-boundary-validator/prompts/` and the closed canonical Task-type set.
- **Criterion boundary**: Extract and retain the expected criterion IDs from the selected common and type-specific checklist inputs.
- **Excluded mechanisms**: No caller-selected criteria, manifest, absolute checklist path, dynamic criterion reordering, or checklist mutation.
- **Prior-decision treatment**: PRODUCT-TASK-SPEC-021-02 D-004 is adopted under TRV authority.
- **Reason**: Deterministic discovery and prompt composition are already accepted and unaffected by the app-local architecture split.
- **Canonical target**: TRV checklist-adapter detailed Specification.
- **ADR route**: `candidate`; final routing remains downstream of T04.
- **Downstream owner**: T04 for detailed-Specification Work Item decomposition.

#### D-006: Model invocation, structured output, and result validation

- **Status**: `decided`
- **Selected outcome**: Preserve one non-streaming Ollama chat request per Task with JSON-Schema-constrained output.
- **Invocation boundary**: Use `stream: false`, deterministic temperature, the composed prompt, and the configured model.
- **Provider boundary**: The Ollama adapter decodes provider structure. The application core validates exact criterion identity, completeness, uniqueness, field validity, and absence of undeclared results.
- **Order boundary**: Returned criterion order need not match checklist order; the validated set may be normalized afterward.
- **Aggregation boundary**: The application derives `overall_compliant`; no model-supplied overall verdict is authoritative.
- **Prior-decision treatment**: PRODUCT-TASK-SPEC-021-02 D-006 through D-009 are adopted under TRV authority.
- **Reason**: These behaviors define the accepted semantic-to-provider seam and remain compatible with remote Ollama over HTTP.
- **Canonical target**: TRV evaluation-runtime detailed Specification.
- **ADR route**: `candidate`; final routing remains downstream of T04.
- **Downstream owner**: T04 for detailed-Specification Work Item decomposition.

#### D-007: Retry, timeout, malformed response, and concrete failure handling

- **Status**: `decided`
- **Selected outcome**: Preserve the W021 execution policy.
- **Attempt policy**:
  - maximum two identical attempts;
  - 300-second timeout per attempt;
  - no prompt repair between attempts.
- **Retryable failures**: Connection failure, timeout, HTTP 5xx, JSON decode failure, and incomplete or invalid criterion result sets.
- **Non-retryable cases**: Structural precondition failure, configuration error, HTTP 4xx, and completed semantic evaluation with false criteria.
- **Malformed-response policy**: Use one `malformed_model_response` execution-failure category; do not repair JSON, synthesize criteria, discard extras, or expose the full raw response.
- **Final failure**: Return `execution_failure` with the final category, concise reason, and attempt count.
- **Prior-decision treatment**: PRODUCT-TASK-SPEC-021-02 D-010 through D-012 are adopted under TRV authority.
- **Reason**: Remote deployment changes the endpoint location, not the accepted bounded retry and failure semantics.
- **Canonical target**: TRV failure-handling detailed Specification.
- **ADR route**: `candidate`; final routing remains downstream of T04.
- **Downstream owner**: T04 for detailed-Specification Work Item decomposition.

#### D-008: Configuration and startup validation

- **Status**: `decided`
- **Selected outcome**: Require the repository root, remote Ollama base URL, and model name explicitly through process environment variables.
- **Required variables**:
  - `BREWPRINT_REPO_ROOT`;
  - `TASK_VALIDATOR_OLLAMA_URL`;
  - `TASK_VALIDATOR_OLLAMA_MODEL`.
- **Optional variable**:
  - `TASK_VALIDATOR_ATTEMPT_TIMEOUT_SEC`, default `300`.
- **Startup validation**:
  - validate required values and absolute HTTP or HTTPS URL shape before serving MCP;
  - do not perform an Ollama network health check during startup;
  - print one concise stderr reason and exit non-zero for missing or invalid startup configuration.
- **Excluded configuration**: No localhost URL default, `.env` loading, repository YAML, command-line configuration flags, API-key management, or secret-store integration.
- **Reason**: The accepted deployment uses a separate Ollama server. Requiring the URL prevents an accidental localhost fallback while keeping runtime availability an invocation-time concern.
- **Canonical target**: TRV configuration and startup detailed Specification.
- **ADR route**: `candidate`; final routing remains downstream of T04.
- **Downstream owner**: T04 for detailed-Specification Work Item decomposition and later ADR routing.

#### D-012: Windows launcher, external `mcp-proxy`, and HTTP exposure boundary

- **Status**: `decided`
- **Selected outcome**: Provide one source-controlled `trv/start.ps1` that launches the built stdio MCP executable behind an externally provided `mcp-proxy` command.
- **Launcher boundary**:
  - launcher path: `trv/start.ps1`;
  - built executable: `bin/task-responsibility-validator.exe`;
  - default bind address: `127.0.0.1`;
  - default port: `8933`;
  - launcher parameters may override host, port, repository root, Ollama URL, model, timeout, and stateless mode;
  - launcher validates required paths, port, URL, executable presence, and proxy-command availability before launch;
  - launcher forwards the accepted environment variables and sets the proxy working directory to the repository root.
- **Ownership boundary**: `mcp-proxy` remains an external deployment dependency. TRV does not vendor, install, supervise, or expose a lifecycle API for it.
- **Excluded deployment**: No TRV-owned HTTP listener, installer, Windows service registration, global installation, or bundled proxy runtime.
- **Reason**: A source-controlled launcher makes deployment repeatable while preserving stdio MCP as the application boundary and externalizing HTTP translation.
- **Canonical target**: TRV Windows deployment and launcher detailed Specification.
- **ADR route**: `candidate`; final routing remains downstream of T04.
- **Downstream owner**: T04 for detailed-Specification Work Item decomposition and later ADR routing.

#### D-013: Required automated tests, fake Ollama boundary, process smoke, and live-runtime verification

- **Status**: `decided`
- **Selected outcome**: Use deterministic required tests with a small in-process HTTP stub for Ollama behavior. Keep live remote-Ollama verification optional and non-blocking for source acceptance.
- **Required unit and component coverage**:
  - Task-path normalization, root containment, parsing, and invalid `task_type`;
  - checklist mappings, composition order, and criterion-ID extraction;
  - application-core success, semantic false, structural failure, execution failure, exact criterion-set validation, and logical AND;
  - Ollama HTTP request shape, JSON decoding, 4xx, 5xx, timeout, identical retry, malformed output, and raw-response containment through `httptest.Server` or an equivalent in-process HTTP stub;
  - MCP initialize, tools/list, tools/call, strict `task_path` input, tagged output mapping, and server continuity;
  - startup configuration and launcher input validation.
- **Required integration gates**:
  - focused `go test` for the TRV tree;
  - repository-wide `go test ./...`;
  - Windows executable build;
  - stdio MCP process smoke against the HTTP stub;
  - `mcp-proxy` launcher smoke against the HTTP stub.
- **Live-runtime boundary**: A manually invoked remote-Ollama smoke may verify reachability and one successful semantic result, but external availability or model nondeterminism does not fail source acceptance.
- **Detailed-design boundary**: Exact test files, commands, fixtures, timeout controls, and skip conditions belong to the detailed Specification Work Item.
- **Reason**: A tiny HTTP stub makes success and failure paths cheap and reproducible without implementing an Ollama clone or coupling acceptance to an external server.
- **Canonical target**: TRV verification detailed Specification.
- **ADR route**: `candidate`; final routing remains downstream of T04.
- **Downstream owner**: T04 for detailed-Specification Work Item decomposition and later ADR routing.

#### D-015: Implementation-readiness gate and future decomposition authority

- **Status**: `decided`
- **Selected outcome**: Do not select or materialize the implementation Work Item shape in this Task. Defer implementation decomposition until the detailed-Specification Work Item is independently reviewed and closed.
- **Current-state clarification**: This ledger is not implementation-ready design. It fixes architecture and contract choices and defines what later detailed Specifications must make exact.
- **Required readiness gate**:
  - reviewed and closed architecture Specification;
  - reviewed and closed external/application contract Specification;
  - reviewed and closed detailed Specifications containing exact packages, source paths, interfaces, types, constructors, dependency wiring, request and response models, parsing and validation behavior, retry and error behavior, file writers, fixtures, commands, and smoke procedures.
- **Future decomposition authority**:
  - a later decomposition or `work_item_execution` Task evaluates the final writer map and implementation scope;
  - that later Task decides whether one implementation Work Item is sufficient or multiple Work Items and an execution hub are justified;
  - no implementation Work Item ID, executor Task, or implementation prompt is created by current T02 or downstream graph-coordination T04.
- **Design protection**: Implementation work may not invent missing design details or reopen accepted decisions without an explicit design-change route.
- **Reason**: Implementation structure must be derived from reviewed detailed Specifications and actual writer boundaries rather than guessed from conceptual design.
- **Canonical target**: Future TRV implementation-readiness gate and decomposition Task after detailed-Specification closure.
- **ADR route**: `not_applicable`.
- **Downstream owner**: A later decomposition or `work_item_execution` Task created after reviewed detailed-Specification closure.

### Current cursor

- Current item: none.
- Loop state: `decision_complete`.
- No item remains `in_discussion`.

### Status definitions

| status | meaning |
|---|---|
| `open` | The item is known but not currently being asked. |
| `in_discussion` | The item is the single active user judgment. |
| `decided` | The selected outcome and reason are persisted. |
| `blocked` | A named external input prevents judgment. |
| `deferred` | The user explicitly moved the item outside this Work Item. |
| `superseded` | A later workflow artifact replaced the item. |

### Fixed PRODUCT inputs

- One invocation evaluates one Task.
- Semantic Evidence comes only from Task-local content.
- The declared `task_type` selects the checklist automatically.
- The common checklist and one type-specific checklist both apply.
- Every criterion returns one boolean and one concise reason.
- Overall compliance is the logical AND of all criterion results.
- Structural precondition failure, semantic evaluation, and execution failure remain distinct.
- Human acceptance or rejection belongs to the caller.
- Post-authoring and post-Evidence validation use the same contract.
- The validator does not rewrite, correct, complete, release, accept, or reject a Task.
- Current DRMCP integration remains outside TRV-WORK-SPEC-001.

### Historical candidate treatment

- PRODUCT-TASK-SPEC-021-02 D-001 through D-019 are prior accepted inputs whose ownership was misplaced under PRODUCT.
- TRV adopts an unchanged prior decision without re-questioning when it does not conflict with the new TRV architecture or user clarification.
- A prior decision is reopened only when the ownership move, architecture split, or current deployment facts materially change its meaning.
- PRODUCT-INV-SPEC-009 implementation layouts and seams remain non-adopted candidates.
- Historical status labels are not copied blindly; each adopted decision records explicit TRV provenance and canonical target.
- The reviewed checklist assets remain read-only inputs.

### Expected downstream route

```text
terminal T02 decisions
  -> T03 graph amendment removing the unnecessary Investigation route
  -> T04 bounded repository alignment and downstream graph coordination
  -> conditional reconciliation, Requirement authoring, ADR routing, ADR authoring, Specification authoring, integrated review, and closure synchronization
```

Implementation and current DRMCP integration remain outside this route.

## Done condition

- D-001 through D-017 are `decided`, `deferred`, or validly `blocked`.
- Every terminal item records one selected outcome and concise reason.
- Every terminal item identifies its canonical target and provisional ADR route.
- The ledger is sufficient to bound T03 impact and conflict investigation.
- PRODUCT-owned semantics remain unchanged.
- No Requirement, Investigation, ADR, Specification, implementation Task, or executor prompt is produced.

## Verification

- Confirm at most one item is `in_discussion`.
- Confirm every explicit user answer is persisted before the cursor advances.
- Confirm dependent items do not become active before their prerequisites are terminal.
- Confirm historical candidates are not treated as canonical TRV decisions.
- Confirm PRODUCT-fixed inputs do not appear as decision topics.
- Confirm no downstream authoring, implementation, review, or synchronization occurs.

## Evidence

- TRV-TASK-SPEC-001-01 created this decision owner and dependency route.
- The initial inventory contains the minimum app-local design topics required by TRV-WORK-SPEC-001.
- The user selected one technology-neutral TRV application Requirement.
- D-001 is `decided` with `TRV-REQ-SPEC-001` as its canonical target.
- The user selected architecture, contract, detailed Specification, and implementation as separate Work Item boundaries.
- D-016 is `decided` and preserves `TRV-WORK-SPEC-001` as the coordinating design boundary.
- The user selected ports and adapters for the TRV application architecture.
- D-017 is `decided` with an application core independent of transport, filesystem, checklist storage, and provider-specific mechanics.
- PRODUCT-TASK-SPEC-021-02 had already selected an externally managed Ollama runtime and configurable Ollama endpoint.
- The user clarified that Ollama runs on a separate server and TRV calls it over HTTP.
- D-005 is `decided`; TRV owns only the Ollama HTTP adapter, not runtime or model lifecycle.
- The user selected the five logical ownership areas and clarified `composition root` as startup and dependency wiring.
- D-011 is `decided` without fixing exact packages or paths.
- Unchanged W021 interface and Task-path decisions were migrated without redundant re-questioning.
- D-002 and D-003 are `decided` under TRV authority.
- The user selected the explicit MCP server name, tool name, path input schema, and tagged result envelope.
- D-009 is `decided`.
- Unchanged W021 caller, migration, checklist, model-invocation, result-validation, retry, timeout, and malformed-response decisions were migrated under TRV authority.
- D-010, D-014, D-004, D-006, and D-007 are `decided`.
- The separate-server Ollama clarification conflicts with W021's optional localhost URL default, so configuration was revised rather than migrated blindly.
- The user selected an explicit required remote Ollama URL with environment-only configuration and no startup health check.
- D-008 is `decided`.
- The user selected `trv/start.ps1`, external `mcp-proxy`, the fixed executable path, localhost-only default binding, and no bundled proxy or installer.
- D-012 is `decided`.
- The user selected deterministic required tests using a small in-process Ollama HTTP stub and optional live remote-Ollama smoke.
- D-013 is `decided`.
- The user identified that the current decision ledger is not implementation-ready detailed design.
- D-015 was reframed and accepted as a future implementation-readiness gate and decomposition-authority decision.
- D-015 is `decided`; no implementation Work Item or executor graph is materialized before detailed-Specification closure.
- D-001 through D-017 are terminal.
- T02 is `done` and the decision loop is complete.
- No Requirement, child Work Item, or other downstream deliverable was created.
- PRODUCT-TASK-SPEC-021-02 and PRODUCT-INV-SPEC-009 were read as candidate inputs only.
- DRMCP is non-operational, so filesystem authoring was used.
