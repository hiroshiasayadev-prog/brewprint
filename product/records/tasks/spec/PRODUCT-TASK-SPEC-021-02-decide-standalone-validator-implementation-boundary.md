# PRODUCT-TASK-SPEC-021-02: Decide standalone validator implementation boundary

- **id**: PRODUCT-TASK-SPEC-021-02
- **status**: done
- **date**: 2026-07-02
- **work_item**: PRODUCT-WORK-SPEC-021
- **task_type**: decision
- **estimate**: 1d
- **depends_on**:
  - PRODUCT-TASK-SPEC-021-01
- **outputs**:
  - PRODUCT-TASK-SPEC-021-02

## Goal

Produce one implementation-boundary decision ledger for the temporary standalone validator.

## Work

- Run an interactive decision loop with the user.
- Ask one decision question at a time.
- Persist each accepted answer before advancing the cursor.
- Keep options and conclusions open until the corresponding question is answered.
- Record a concise reason, boundary, canonical target, and routing state for each terminal item.
- Stop downstream progression when an item remains materially ambiguous.

This Task must not:

- perform the focused implementation impact Investigation;
- author ADR, Requirement, Specification, checklist, or implementation content;
- inspect repository implementation seams as a substitute for T03;
- materialize executor-ready implementation Tasks;
- issue a Claude Code implementation prompt;
- perform independent review, correction, synchronization, stage, or commit work.

### Decision inventory

| ID | Topic | Status | Depends on |
|---|---|---|---|
| D-001 | App namespace and source ownership | `decided` | none |
| D-002 | Executable interface | `decided` | D-001 |
| D-003 | Task record input method | `decided` | D-002 |
| D-004 | Checklist discovery and prompt composition | `decided` | D-001, D-003 |
| D-005 | Local model, provider, and runtime | `decided` | D-001 |
| D-006 | Model invocation and structured-output handling | `decided` | D-005 |
| D-007 | Criterion result parsing | `decided` | D-004, D-006 |
| D-008 | Semantic non-compliance and execution-failure separation | `decided` | D-006, D-007 |
| D-009 | Overall logical-AND aggregation | `decided` | D-007, D-008 |
| D-010 | Retry policy | `decided` | D-005, D-006, D-008 |
| D-011 | Timeout policy | `decided` | D-005, D-006, D-008 |
| D-012 | Malformed response policy | `decided` | D-006, D-007, D-008 |
| D-013 | Human violation acceptance and rejection interaction | `decided` | D-008, D-009 |
| D-014 | Post-authoring invocation | `decided` | D-002, D-003, D-013 |
| D-015 | Post-Evidence invocation | `decided` | D-002, D-003, D-013 |
| D-016 | Configuration and secrets boundary | `decided` | D-005, D-006 |
| D-017 | Executable packaging and launch command | `decided` | D-001, D-002, D-005, D-016 |
| D-018 | Output and exit-status behavior | `decided` | D-002, D-008, D-009, D-012, D-013 |
| D-019 | Implementation compatibility boundary | `decided` | D-001 through D-018 |

### Decision records

#### D-001: App namespace and source ownership

- **Status**: `decided`
- **Decision summary**: Use a dedicated standalone application namespace. Do not place the validator under a generic repository-local `tools` namespace.
- **Source ownership**: The standalone app owns its entrypoint, validation use case, criterion result model, Task and checklist adapters, and local-model adapter.
- **Migration boundary**: Keep DRMCP transport, record access, and routing outside the standalone app. Future DRMCP integration replaces outer adapters without redefining the validation use case.
- **Reason**: The validator has enough orchestration and failure handling to exceed a small tool boundary. Independent app ownership avoids later separation work during DRMCP migration.
- **Canonical target**: W021 implementation app boundary and the executor graph materialized by T04.
- **ADR route**: `candidate`; final routing remains downstream of T03.

#### D-002: Executable interface

- **Status**: `decided`
- **Decision summary**: Expose the standalone validator through a dedicated MCP server.
- **Interface boundary**: The MCP adapter is the required external interface for AI agents. A thin CLI may exist only for development and smoke verification.
- **DRMCP boundary**: Do not integrate the validator into DRMCP in W021. Future migration may replace the standalone MCP adapter while preserving the application use case.
- **Reason**: A CLI-only interface is unavailable to ChatGPT sessions without arbitrary command execution. A dedicated MCP server provides direct tool access without coupling the validator to current DRMCP.
- **Canonical target**: W021 executable interface and the executor graph materialized by T04.
- **ADR route**: `candidate`; final routing remains downstream of T03.

#### D-003: Task record input method

- **Status**: `decided`
- **Decision summary**: Accept exactly one repository-root-relative Task file path per MCP invocation.
- **Input boundary**: The standalone validator reads the Task file from the supplied path and obtains the declared `task_type` from that record.
- **Excluded inputs**: Do not accept Task body content as the primary external input. Do not resolve a Task public ID through DRMCP or another record index.
- **Reason**: A path-only input keeps the temporary MCP contract thin and avoids importing DRMCP record-resolution responsibilities.
- **Canonical target**: W021 MCP tool input contract and the executor graph materialized by T04.
- **ADR route**: `candidate`; final routing remains downstream of T03.

#### D-004: Checklist discovery and prompt composition

- **Status**: `decided`
- **Decision summary**: Load checklist fragments from fixed repository-root-relative paths and compose them in the accepted order.
- **Checklist paths**:
  - `skills/task-responsibility-boundary-validator/prompts/evaluator-instructions.md`
  - `skills/task-responsibility-boundary-validator/prompts/common.md`
  - `skills/task-responsibility-boundary-validator/prompts/task-types/<task_type>.md`
- **Composition boundary**: Validate `task_type` against the canonical closed value set. Append the full Task record after the three checklist fragments with an explicit separator.
- **Excluded mechanism**: Do not add a manifest, caller-selected criteria, absolute paths, or runtime criterion reordering.
- **Reason**: Repository-root-relative fixed paths preserve deterministic discovery without coupling the implementation to one machine path.
- **Canonical target**: W021 checklist adapter and prompt-composition contract.
- **ADR route**: `candidate`; final routing remains downstream of T03.

#### D-005: Local model, provider, and runtime

- **Status**: `decided`
- **Decision summary**: Use a local Qwen model through an externally managed Ollama runtime.
- **Runtime boundary**: The standalone validator calls the local Ollama API. It does not start, stop, install, or manage Ollama or model files.
- **Model selection**: Supply the concrete Ollama model name through configuration rather than hard-coding one model tag.
- **Provider boundary**: W021 supports Ollama only. Direct llama.cpp integration and remote hosted providers remain outside scope.
- **Reason**: Ollama keeps model lifecycle separate from the MCP server and permits local model replacement without changing the validation use case.
- **Canonical target**: W021 local-model adapter and runtime configuration contract.
- **ADR route**: `candidate`; final routing remains downstream of T03.

#### D-006: Model invocation and structured-output handling

- **Status**: `decided`
- **Decision summary**: Evaluate one Task with one non-streaming Ollama chat request constrained by a JSON Schema.
- **Invocation boundary**: Use `stream: false` and `temperature: 0`. Supply the composed evaluator prompt and full Task record in one request.
- **Structured output**: Require one JSON array whose elements contain `id`, `result`, `reason`, and `section` fields with schema-constrained types.
- **Application validation boundary**: Treat schema decoding as provider-level structure only. The application validates criterion identity, completeness, and duplication under D-007.
- **Excluded invocation**: Do not invoke the model once per criterion and do not accept free-form prose outside the structured result.
- **Reason**: One schema-constrained request minimizes inference overhead while preserving application-owned completeness checks.
- **Canonical target**: W021 Ollama adapter and structured-response contract.
- **ADR route**: `candidate`; final routing remains downstream of T03.

#### D-007: Criterion result parsing

- **Status**: `decided`
- **Decision summary**: Require the returned criterion set to match the expected checklist criterion set exactly, without requiring response order to match checklist order.
- **Required checks**:
  - result count equals the expected criterion count;
  - every expected criterion ID appears exactly once;
  - no unknown, duplicate, or missing criterion ID exists;
  - `result` is boolean;
  - `reason` and `section` are non-empty strings;
  - no undeclared result field exists.
- **Order boundary**: Accept any criterion-result order. The application may normalize results into checklist order after successful set validation.
- **Failure boundary**: Do not repair IDs, synthesize missing results, or discard extra results. Any mismatch prevents complete semantic evaluation and becomes an execution failure.
- **Reason**: Criterion identity and completeness affect semantic validity. Array order does not affect the criterion judgments or logical-AND result.
- **Canonical target**: W021 application result parser and completeness validator.
- **ADR route**: `candidate`; final routing remains downstream of T03.

#### D-008: Semantic non-compliance and execution-failure separation

- **Status**: `decided`
- **Decision summary**: Represent structural precondition failure, completed semantic evaluation, and execution failure as distinct tagged outcomes.
- **Semantic evaluation**: Return criterion results and one derived `overall_compliant` value. One or more false criteria remain a successfully completed semantic evaluation.
- **Structural precondition failure**: Use when the Task cannot be read or parsed, or when `task_type` is missing or invalid. Return no criterion results and no overall compliance value.
- **Execution failure**: Use when Ollama access, timeout, HTTP handling, JSON decoding, or D-007 completeness validation prevents full evaluation. Return no criterion results and no overall compliance value.
- **Boundary rule**: Never convert structural or execution failures into semantic non-compliance. Never synthesize missing criterion judgments.
- **Reason**: Semantic false is an evaluated judgment. Structural and runtime failures provide no valid semantic judgment.
- **Canonical target**: W021 application outcome model and MCP response contract.
- **ADR route**: `candidate`; final routing remains downstream of T03.

#### D-009: Overall logical-AND aggregation

- **Status**: `decided`
- **Decision summary**: Derive `overall_compliant` mechanically in the application through logical AND across every validated criterion result.
- **Aggregation rule**: Return `true` only when every criterion result is `true`. Return `false` when one or more criterion results are `false`.
- **Precondition**: Run D-007 criterion-set validation before aggregation.
- **Failure boundary**: Return no overall compliance value for structural or execution failures. Treat an empty expected criterion set as an implementation execution failure rather than vacuous compliance.
- **Model boundary**: Do not ask the model for a separate overall verdict and do not accept one as authority.
- **Reason**: Mechanical aggregation preserves the accepted contract and prevents a second model judgment from contradicting criterion results.
- **Canonical target**: W021 application aggregation and MCP response contract.
- **ADR route**: `candidate`; final routing remains downstream of T03.

#### D-010: Retry policy

- **Status**: `decided`
- **Decision summary**: Retry at most once after the initial model attempt, for a maximum of two attempts per invocation.
- **Retryable failures**: Retry transient Ollama connection failure, HTTP 5xx, timeout, JSON decode failure, and D-007 criterion-set validation failure.
- **Non-retryable outcomes**: Do not retry structural precondition failures, configuration errors, HTTP 4xx, or completed semantic evaluations with false criteria.
- **Retry prompt boundary**: Reuse the same request unchanged. Do not inject the prior failure reason into the prompt.
- **Final failure**: After the second failed attempt, return `execution_failure` with attempt count and the final failure category.
- **Reason**: One unchanged retry covers transient failure without introducing a failure-taxonomy-dependent prompt-repair protocol.
- **Canonical target**: W021 Ollama invocation policy and execution-failure response.
- **ADR route**: `candidate`; final routing remains downstream of T03.

#### D-011: Timeout policy

- **Status**: `decided`
- **Decision summary**: Allow up to 300 seconds per model attempt and up to two attempts per MCP invocation.
- **Attempt timeout**: Cancel one Ollama attempt after 300 seconds.
- **Invocation budget**: Allow up to 600 seconds for model attempts. Require the MCP caller timeout to be at least 620 seconds to preserve response-handling margin.
- **Retry boundary**: A timed-out first attempt may use the one retry allowed by D-010. A timed-out second attempt returns `execution_failure`.
- **Configuration boundary**: Expose the timeout values through configuration while retaining 300 seconds per attempt and 620 seconds caller timeout as the operational defaults.
- **Reason**: Observed Qwen reasoning can exceed three minutes for long prompts. A five-minute attempt budget preserves useful inference while keeping bounded MCP execution.
- **Canonical target**: W021 Ollama timeout configuration and MCP deployment guidance.
- **ADR route**: `candidate`; final routing remains downstream of T03.

#### D-012: Malformed response policy

- **Status**: `decided`
- **Decision summary**: Classify all unusable model outputs under one `malformed_model_response` execution-failure category.
- **Covered failures**: Include invalid JSON, schema mismatch, missing, duplicate, or unknown criterion IDs, invalid or empty fields, and undeclared fields.
- **Retry behavior**: Apply D-010 unchanged. Retry the same request once and return `execution_failure` after the second malformed response.
- **Response boundary**: Return the attempt count and one concise error reason. Do not include the full raw model response in the MCP result.
- **Repair boundary**: Do not repair JSON, correct IDs, discard extra fields, or accept partial criterion results.
- **Reason**: One category preserves enough failure meaning without introducing a detailed malformed-response taxonomy or repair protocol.
- **Canonical target**: W021 execution-failure model and MCP response contract.
- **ADR route**: `candidate`; final routing remains downstream of T03.

#### D-013: Human violation acceptance and rejection interaction

- **Status**: `decided`
- **Decision summary**: Keep human acceptance, rejection, and Task Evidence persistence outside the standalone validator and under the calling workflow.
- **Validator response**: A non-compliant semantic evaluation returns the violated criteria and their reasons and indicates that human action is required.
- **Caller responsibility**: The calling agent presents the violations to the human and obtains either acceptance with a required reason or rejection.
- **Acceptance persistence**: The caller records each violated criterion, its validator reason, the acceptance decision, and the acceptance reason in Task Evidence.
- **Rejection route**: The caller returns work to Task correction or responsibility-boundary reconsideration based on the current workflow state.
- **Validator boundary**: Do not add acceptance or rejection MCP tools. Do not persist human judgment, rewrite the Task, or select the return route automatically.
- **Reason**: Accepted authority assigns enforcement and exception persistence to the caller while keeping the validator read-only.
- **Canonical target**: W021 MCP result contract and caller integration guidance.
- **ADR route**: `reuse PRODUCT-ADR-SPEC-017`.

#### D-014: Post-authoring invocation

- **Status**: `decided`
- **Decision summary**: Invoke the validator synchronously after the Task file is written and before authoring completion or downstream release.
- **Trigger**: Run after Task creation and after any substantive authoring update that can change responsibility semantics.
- **Caller behavior**: The authoring caller supplies the repository-root-relative Task path and waits for the validator result.
- **Continuation rule**: Continue only after a compliant result or an explicit human acceptance under D-013.
- **Failure rule**: Stop authoring completion on structural or execution failure and report the failure.
- **Excluded mechanism**: Do not add a filesystem watcher, repository scan, or validator-owned Task discovery.
- **Reason**: D-003 requires a persisted path. Post-write, pre-release validation preserves that contract without adding temporary-file or direct-body input handling.
- **Canonical target**: W021 caller integration guidance for Task authoring workflows.
- **ADR route**: `reuse PRODUCT-ADR-SPEC-017`.

#### D-015: Post-Evidence invocation

- **Status**: `decided`
- **Decision summary**: Invoke the same validator synchronously after final Task Evidence is written and before the Task becomes `done` or is released.
- **Trigger**: Run once the execution result and final Evidence are present in the Task record.
- **Caller behavior**: The completion or release workflow supplies the same repository-root-relative Task path and waits for the result.
- **Continuation rule**: Permit `done` or release only after a compliant result or explicit human acceptance under D-013.
- **Failure rule**: Stop completion on structural or execution failure. A rejected semantic violation returns to correction or responsibility-boundary reconsideration.
- **Reason**: Final Evidence can reveal work that differs from the authored responsibility boundary. Pre-completion validation evaluates the actual recorded outcome.
- **Canonical target**: W021 caller integration guidance for Task completion and release workflows.
- **ADR route**: `reuse PRODUCT-ADR-SPEC-017`.

#### D-016: Configuration and secrets boundary

- **Status**: `decided`
- **Decision summary**: Configure the standalone validator through process environment variables only.
- **Required variables**:
  - `BREWPRINT_REPO_ROOT`;
  - `TASK_VALIDATOR_OLLAMA_MODEL`.
- **Optional variables with defaults**:
  - `TASK_VALIDATOR_OLLAMA_URL`, default `http://127.0.0.1:11434`;
  - `TASK_VALIDATOR_ATTEMPT_TIMEOUT_SEC`, default `300`.
- **Fixed values**: Keep retry count and checklist paths fixed by accepted decisions rather than exposing them as runtime configuration.
- **Secrets boundary**: Do not require or manage API keys for local Ollama. Do not add `.env` loading, repository YAML configuration, or secret-store integration.
- **Failure behavior**: Reject missing or invalid required configuration during server startup, print the reason to stderr, and exit non-zero.
- **Reason**: Environment variables fit the existing PowerShell and `mcp-proxy` launch workflow and avoid a separate configuration parser and file contract.
- **Canonical target**: W021 process configuration and deployment guidance.
- **ADR route**: `candidate`; final routing remains downstream of T03.

#### D-017: Executable packaging and launch command

- **Status**: `decided`
- **Decision summary**: Implement the standalone validator in Go and build one Windows executable that serves MCP over stdio.
- **Application boundary**: Keep the standalone app outside the DRMCP package tree and do not depend on DRMCP internal packages.
- **Migration boundary**: Isolate the validation core, filesystem adapter, and Ollama adapter from the standalone MCP adapter so future DRMCP integration can reuse or move the core without retaining the standalone transport.
- **Build contract**: Build the app with `go build` into `bin/task-responsibility-validator.exe`.
- **Launch contract**: Run the executable behind external `mcp-proxy`. The validator does not own an HTTP server, port binding, or proxy lifecycle.
- **Runtime boundary**: Do not require Python or Node.js at runtime. Do not package a global install or single-purpose installer.
- **Reason**: Go matches the future DRMCP implementation language and avoids rewriting the validator core during migration while preserving current standalone ownership.
- **Canonical target**: W021 source layout, build command, and deployment guidance.
- **ADR route**: `candidate`; final routing remains downstream of T03.

#### D-018: Output and exit-status behavior

- **Status**: `decided`
- **Decision summary**: Return every validation outcome as one structured MCP tool result and keep the server process running after per-invocation failures.
- **Semantic result**: Return `outcome: semantic_evaluation`, the Task path, all criterion results, derived `overall_compliant`, and `human_action_required`.
- **Structural result**: Return `outcome: structural_precondition_failure`, the Task path, one failure category, and a concise message.
- **Execution result**: Return `outcome: execution_failure`, the Task path, one failure category, a concise message, and attempt count when model invocation began.
- **Protocol boundary**: A semantic violation is a successful MCP tool call. Structural and execution failures remain structured domain results rather than process termination.
- **Process exit boundary**: Exit non-zero only when startup validation fails or the stdio MCP server cannot continue. Print the startup failure reason to stderr.
- **Configuration correction**: Missing or invalid required environment variables cause startup failure rather than an invocation-level execution result.
- **Reason**: Per-Task failures must remain inspectable without destroying the long-lived MCP server. Invalid startup configuration cannot produce a useful server.
- **Canonical target**: W021 MCP response schema, startup validation, and process lifecycle contract.
- **ADR route**: `candidate`; final routing remains downstream of T03.

#### D-019: Implementation compatibility boundary

- **Status**: `decided`
- **Decision summary**: Preserve the accepted semantic validator behavior across future DRMCP integration, but do not guarantee source-code reuse or standalone interface compatibility.
- **Standalone scope**: W021 supports current canonical Task files, the closed 11-value `task_type` set, repository-root-relative Task paths, and Windows execution.
- **Path safety**: Normalize `/` and `\` separators. Reject absolute paths and any path that escapes `BREWPRINT_REPO_ROOT`.
- **Future migration**: A later DRMCP integration may replace path input with ID-as-ref resolution and may rewrite the implementation. The migration must preserve criterion results, logical-AND aggregation, outcome separation, caller-owned human judgment, and the two invocation points.
- **Non-goals**: Do not support `v01/`, legacy Task formats, alternate checklist placement, permanent standalone tool-name compatibility, or permanent environment-variable compatibility.
- **Reason**: The standalone app is temporary. Future DRMCP record identity and transport differ from the current path-based MCP boundary, so only semantic behavior warrants continuity.
- **Canonical target**: W021 standalone implementation boundary and a future DRMCP integration Requirement or Work Item.
- **ADR route**: `candidate`; final routing remains downstream of T03.

### Current cursor

- Current item: none.
- Loop state: `decision_complete`.
- Question policy: ask exactly one unresolved item at a time during execution.

### Fixed inputs

- The accepted W019 semantic validator contract remains unchanged.
- The accepted W020 checklist artifact set remains unchanged.
- One invocation evaluates one Task record.
- The common checklist and exactly one declared-`task_type` checklist are composed in the accepted order.
- Every criterion returns one boolean result and one concise Task-local reason.
- Overall compliance is the logical AND of criterion results.
- Structural precondition failure, semantic evaluation, and execution failure remain distinct.
- Semantic violations require explicit human acceptance or rejection.
- The same semantic contract applies after Task authoring and after final Evidence.
- Current DRMCP integration and future DRMCP integration remain outside W021.

## Done condition

- D-001 through D-019 are `decided`, `deferred`, or validly `blocked`.
- Every terminal item records its selected outcome and concise reason.
- The ledger is sufficient to bound T03 repository investigation.
- No Investigation, implementation, or canonical authoring is performed.

## Verification

- Confirm at most one item is `in_discussion`.
- Confirm every terminal item has a concise reason and boundary.
- Confirm every open dependency is resolved before its dependent item becomes terminal.
- Confirm no option or conclusion was invented by T01.
- Confirm the final ledger preserves the accepted W019 and W020 contracts.
- Confirm no implementation or downstream graph materialization occurred.

## Evidence

- PRODUCT-TASK-SPEC-021-01 created this decision owner without selecting any implementation option.
- W019 supplies fixed semantic behavior.
- W020 supplies the accepted checklist artifact set.
- D-001 selected a dedicated standalone application namespace with DRMCP-independent core ownership.
- D-002 selected a dedicated MCP server as the required external interface.
- D-003 selected one repository-root-relative Task path as the invocation input.
- D-004 selected fixed repository-root-relative checklist paths and deterministic composition order.
- D-005 selected configured local Qwen inference through an externally managed Ollama runtime.
- D-006 selected one non-streaming, schema-constrained Ollama request per Task.
- D-007 selected strict criterion-set validation without response-order enforcement.
- D-008 selected three distinct tagged outcomes for semantic evaluation, structural failure, and execution failure.
- D-009 selected application-owned logical-AND aggregation after complete criterion validation.
- D-010 selected one unchanged retry for retryable execution failures.
- D-011 selected a 300-second attempt timeout, two-attempt budget, and 620-second minimum caller timeout.
- D-012 selected one `malformed_model_response` category without automatic repair or raw-response exposure.
- D-013 kept human acceptance, rejection, and Evidence persistence under the calling workflow.
- D-014 selected synchronous post-write, pre-release validation for Task authoring.
- D-015 selected synchronous post-Evidence validation before Task completion or release.
- D-016 selected process environment variables only, with no secret or config-file subsystem.
- D-017 selected a Go Windows executable serving stdio MCP behind external `mcp-proxy`.
- D-018 selected structured per-invocation results and non-zero exit only for startup or unrecoverable server failure.
- D-019 preserved semantic compatibility only and explicitly allowed future ID-as-ref input and implementation replacement.
- D-001 through D-019 are all `decided`.
- Verification confirmed no more than one item was ever `in_discussion`, every terminal item has a reason and boundary, and no Investigation or implementation work occurred.
