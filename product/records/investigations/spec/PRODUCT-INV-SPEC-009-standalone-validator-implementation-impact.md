# PRODUCT-INV-SPEC-009: Standalone validator implementation impact

- **status**: concluded
- **date**: 2026-07-02
- **trigger**: PRODUCT-WORK-SPEC-021 standalone validator implementation impact Investigation
- **scope**: T02で確定したstandalone validator実装境界に対するrepository placement、implementation seam、runtime constraint、test surface、writer boundary、verification commandの調査。
- **non_scope**: 設計判断の変更、ADRまたはSpecification authoring、Task graph変更、sourceまたはtest実装、review、closure、stage、commit。
- **source_refs**:
  - PRODUCT-WORK-SPEC-021
  - PRODUCT-WORK-SPEC-020
  - spec:product.responsibility_boundary_validator
  - PRODUCT-ADR-SPEC-015
  - PRODUCT-ADR-SPEC-016
  - PRODUCT-ADR-SPEC-017
- **follow_up_candidates**:
  - executor-ready implementation graph coordination for PRODUCT-WORK-SPEC-021
- **related_work_items**:
  - PRODUCT-WORK-SPEC-021
  - PRODUCT-WORK-SPEC-020
- **related_adrs**:
  - PRODUCT-ADR-SPEC-015
  - PRODUCT-ADR-SPEC-016
  - PRODUCT-ADR-SPEC-017
- **related_specs**:
  - spec:product.responsibility_boundary_validator
  - spec:product.design_records.authoring_standards.task_authoring

## Investigation scope

This Investigation answers one bounded research question.

```text
Given the accepted standalone-validator contract and terminal implementation-boundary decisions, what repository placement, existing implementation seams, local runtime constraints, test surfaces, writer boundaries, and verification commands constrain an executor-ready implementation graph?
```

PRODUCT-TASK-SPEC-021-02 D-001 through D-019 are terminal fixed inputs.
The Investigation records repository facts, implementation candidates, uncertainty, and downstream ownership without reopening those decisions.

## Out of scope

- Selecting or naming the dedicated application namespace.
- Changing the accepted standalone validator behavior.
- Integrating the validator into current DRMCP.
- Choosing Python instead of Go.
- Authoring source, tests, fixtures, scripts, ADRs, Specifications, Tasks, or prompts.
- Materializing an implementation hub or executor graph.
- Running implementation-time `go test`, `go build`, or runtime smoke commands.
- Performing independent review, closure synchronization, stage, or commit work.

## Background

`spec:product.responsibility_boundary_validator` defines one-Task semantic evaluation, Task-local Evidence, automatic checklist selection, criterion-level results, logical-AND aggregation, separated structural and execution failures, and two workflow invocation points.

PRODUCT-WORK-SPEC-020 completed the common checklist and one type-specific checklist for every canonical Task type.
PRODUCT-WORK-SPEC-021 owns the separate temporary standalone implementation.
PRODUCT-ADR-SPEC-016 explicitly excludes current DRMCP integration and future code-reuse guarantees.

The implementation boundary is therefore stable at the semantic and runtime-contract levels.
The remaining Investigation concern is how that boundary maps onto the current repository and an executor-ready implementation graph.

## What was investigated

The inspection was limited to these repository areas.

- The canonical validator Specification and accepted ADRs.
- The evaluator instructions, common criteria, and all type-specific checklist files.
- The current Task authoring Specification and active Brewprint app namespace profile.
- The root Go module and current DRMCP command, stdio, JSON-RPC, tool, parser, configuration, and test patterns.
- `tools/grep-mcp`, `tools/ollama-mcp`, and `tools/repo-ops-mcp` startup and runtime patterns.
- The root `bin/` output location.
- Execution-hub and executor-ready Task design authority.

Current DRMCP source was inspected only as implementation-pattern Evidence.
No current DRMCP source or internal package is proposed as an accepted dependency.

## Findings

### 1. Repository placement and namespace inventory

| item | observed state | relation to T02 decisions | exact Evidence | implementation-graph effect | uncertainty / next owner |
| ---- | -------------- | ------------------------- | -------------- | --------------------------- | ------------------------ |
| Active app namespaces | The active profile names `DRMCP` and `BPDSL` as runtime app namespaces. `PRODUCT` is a product-level namespace rather than a runtime application. No standalone validator app namespace is assigned. | T02 requires a dedicated standalone application namespace and excludes generic `tools` and current DRMCP placement. | `product/records/spec/brewprint/namespaces/app-namespaces.md`, `## Current assignments` and `## Profile notes`. | An executor Task cannot yet name an exact top-level source directory or import path. | Blocker for executor-ready exact paths. T04 must route the missing namespace decision to an authorized decision owner rather than adopt a name. |
| Current source layout | Active application source exists under `drmcp/src/` and `bpdsl/`. Supporting local operations exist under `tools/`. | The standalone application must be separate from DRMCP and must not be reclassified as a generic supporting tool. | `product/records/spec/brewprint/layout/index.md`, `## Current structure`; repository root directory listing. | A new top-level app directory is consistent with the current app-first layout only after its canonical namespace is accepted. | Exact directory remains unresolved with the app namespace. |
| Root Go module | The repository root is one Go module: `github.com/hiroshiasayadev-prog/brewprint`, Go 1.22. | T02 fixes Go and permits an independently built application inside the root module. | `go.mod`. | A new top-level app package can be built without creating a second module. | No module split is required by observed state. |
| Executable output | `bin/` exists and contains current Windows executables. | T02 fixes `bin/task-responsibility-validator.exe` as the build output. | Root `bin/` directory listing. | The build target has a grounded output directory even though its source package path is unresolved. | Whether the executable is retained or only generated is outside this Investigation. |

### 2. Existing Go and MCP implementation seams

| item | observed state | relation to T02 decisions | exact Evidence | implementation-graph effect | uncertainty / next owner |
| ---- | -------------- | ------------------------- | -------------- | --------------------------- | ------------------------ |
| Testable command entry | Current DRMCP separates `main` from `run(args, stdin, stdout, stderr)`. `main` returns non-zero only when `run` fails. | This pattern supports T02 startup-failure and stdio-process semantics without importing DRMCP. | `drmcp/src/cmd/design-records-mcp/main.go`, `main` and `run`. | Candidate: use an equivalent command seam so startup configuration and stdio behavior are unit-testable. | Exact private function names remain implementation-local. |
| Stdio framing | Current DRMCP reads newline-delimited JSON-RPC with a bounded scanner and writes one response line. | T02 requires a dedicated stdio MCP server behind external `mcp-proxy`. | `drmcp/src/internal/designrecordsmcp/stdio.go`, `ServeJSONRPCLines`. | Candidate: a standalone stdio transport package can follow the same observable framing. | Scanner limit and protocol-version choice need an executor contract only if they become externally observable. |
| JSON-RPC and tool projection | Current DRMCP implements `initialize`, `tools/list`, and `tools/call`; tool schemas reject additional properties; tool argument decoding rejects unknown fields. | This is a repository-local pattern for a one-tool dedicated server. | `drmcp/src/internal/designrecordsmcp/jsonrpc.go`, `tools.go`, and `tools_call.go`. | Candidate: separate transport/tool decoding from the validator application service so handler tests do not require Ollama. | Exact server name, tool name, and protocol version are not fixed by T02. |
| Current parser boundary | Current DRMCP parses Task H1, metadata, headings, and raw body, but `taskMetadata`, `parseTaskMetadata`, and `TaskDetail` do not retain `task_type`. | T02 requires selection from the declared `task_type`; current DRMCP cannot provide that value through its current public internal model. | `drmcp/src/internal/designrecords/parser.go`, `taskMetadata` and `parseTaskMetadata`; `drmcp/src/internal/designrecords/types.go`, `TaskDetail`. | The standalone implementation needs its own minimal Task read/parse boundary or a newly accepted shared package. Current DRMCP modification is not a candidate. | Candidate: standalone minimal parsing of `task_type` plus preservation of the full Task text. |
| Go internal visibility | DRMCP packages are under `drmcp/src/internal/...`. A separate top-level app cannot import them across the Go `internal` boundary. | This reinforces the accepted no-current-DRMCP-dependency boundary. | Paths under `drmcp/src/internal/designrecords/` and `drmcp/src/internal/designrecordsmcp/`; Go internal package rule applied to the observed layout. | Executor Tasks must not assume direct package reuse. Copied concepts require standalone ownership and tests. | No blocker beyond the unresolved standalone app path. |

### 3. Task input and checklist loading seams

| item | observed state | relation to T02 decisions | exact Evidence | implementation-graph effect | uncertainty / next owner |
| ---- | -------------- | ------------------------- | -------------- | --------------------------- | ------------------------ |
| Relative-path safety patterns | `repo-ops-mcp` rejects absolute paths, drives, and `..`, resolves the target, and verifies repository containment. `grep-mcp` resolves under one configured root and rejects root escape. | T02 requires repository-root-relative input, absolute-path rejection, root-escape rejection, and `/` plus `\` normalization. | `tools/repo-ops-mcp/server.py`, `_normalize_pathspecs`; `tools/grep-mcp/server.py`, `_resolve_target`. | Candidate: one filesystem adapter owns normalization, pre-join absolute/drive/UNC rejection, `filepath.FromSlash`, `filepath.Clean`, `filepath.Abs`, and post-resolution root containment. | Exact Go symbols remain implementation-local. Symlink handling must preserve root containment. |
| Minimal Task parse | Task metadata is an H1-adjacent bullet block and `task_type` is a required scalar from a closed eleven-value set. | Structural failure must occur before semantic evaluation when `task_type` is missing or invalid. | `product/records/spec/design-records/authoring-standards/task-authoring.md`, `### Metadata schema` and allowed values. | Candidate: parse only enough Markdown to locate the metadata block, extract exactly one `task_type`, validate the closed set, and retain the complete Task text for the prompt. | The standalone parser must not silently inherit obsolete DRMCP metadata assumptions. |
| Checklist inventory | `evaluator-instructions.md`, `common.md`, and exactly eleven type-specific files exist. The eleven names match the canonical Task types. | This satisfies the fixed checklist source set and type mapping. | `skills/task-responsibility-boundary-validator/prompts/`; directory listing of `prompts/task-types/`; Task authoring allowed-value list. | Checklist loading can be deterministic: fixed evaluator path, fixed common path, and a closed `task_type` to filename map. | Missing or unreadable files are startup or structural failures according to the exact load timing selected by the executor contract. |
| Criterion identity | Common criteria use `C01` through `C14`. Type files use `T01` through their final type-specific item. | T02 requires exact criterion-ID set and count validation. | `prompts/common.md` and all eleven `prompts/task-types/*.md`. | Candidate: extract criterion IDs while loading each file, reject duplicate or malformed IDs, and retain the expected set for response validation. | The extraction grammar should be frozen in an executor Task before low-context implementation. |
| Prompt composition | The skill fixes evaluator instructions, common criteria, and declared-type criteria in that order. T02 adds the full Task as the fourth component. | Request composition order is fixed. Model result ordering is explicitly irrelevant under T02. | `skills/task-responsibility-boundary-validator/SKILL.md`, `## Prompt composition`; PRODUCT-TASK-SPEC-021-02 D-004 and D-005. | Application code must preserve prompt-fragment order but validate returned criterion results by ID set rather than array position. | The evaluator text says not to reorder criteria; this constrains model behavior but does not permit positional response validation. |

### 4. Ollama and structured-output seams

| item | observed state | relation to T02 decisions | exact Evidence | implementation-graph effect | uncertainty / next owner |
| ---- | -------------- | ------------------------- | -------------- | --------------------------- | ------------------------ |
| HTTP endpoint pattern | The existing local Ollama tool builds a base URL, posts to `/api/chat`, sends `stream: false`, supplies the model, and places JSON Schema in `format`. | This matches T02 local Ollama, non-streaming, and schema-constrained request decisions. | `tools/ollama-mcp/server.py`, `_endpoint` and `_delegate` request payload. | Candidate: a standalone Ollama adapter can use Go `net/http` and an injected HTTP client interface. | No shared Python implementation is proposed. |
| Runtime configuration | The existing tool reads Ollama URL and model from environment and validates required values at startup. | T02 fixes `TASK_VALIDATOR_OLLAMA_MODEL`, optional URL default, and process-environment-only configuration. | `tools/ollama-mcp/start.ps1` and `server.py` environment loading. | Configuration parsing should be a distinct startup seam with deterministic tests. | Exact Qwen tag is a deployment value, not a repository design decision. |
| Structured response handling | The existing tool decodes `message.content`, parses JSON, validates schema, and does not require model thinking text in the external result. | T02 requires strict criterion fields/types/non-empty values and forbids raw model response exposure. | `tools/ollama-mcp/server.py`, `_delegate` response handling. | Candidate: keep raw upstream response private; decode only the expected envelope and criterion array; then validate exact IDs, count, booleans, reasons, and sections. | Root `go.mod` currently has no JSON Schema library. Explicit Go validation may satisfy the strict contract without adding one; dependency choice remains implementation-local unless it affects writers. |
| Retry and timeout | T02 fixes 300 seconds per attempt, at most two identical attempts, retryable transport/timeout/5xx/malformed failures, and non-retryable 4xx. | These are mandatory adapter behaviors rather than investigation options. | PRODUCT-TASK-SPEC-021-02 D-008 and D-009. | Candidate: freeze request bytes before attempt one; each attempt receives its own context timeout; a retry-policy unit owns classification. | Caller timeout must remain at least 620 seconds; caller configuration is outside this repository implementation. |
| Test double seam | Current Go tests use dependency injection; Go standard library provides `httptest.Server`. | Ollama behavior can be tested without a running model. | `drmcp/src/internal/designrecordsmcp/server.go`, injected `IndexBuilder`; repository Go test style in `*_test.go`. | Candidate: inject an HTTP doer or base URL and cover success, 4xx, 5xx, timeout, malformed output, and same-request retry. | Exact interface name is private implementation detail. |

### 5. Configuration and startup constraints

| item | observed state | relation to T02 decisions | exact Evidence | implementation-graph effect | uncertainty / next owner |
| ---- | -------------- | ------------------------- | -------------- | --------------------------- | ------------------------ |
| Required environment | Repository root and model are mandatory. URL and attempt timeout have fixed defaults. | These values are fixed by T02 and must not become flags or config files. | PRODUCT-TASK-SPEC-021-02 D-014. | Startup tests must cover missing, empty, invalid, defaulted, and explicit values. | Environment variable names are temporary standalone compatibility only. |
| Process exit boundary | Current Go command exits non-zero when startup/server `run` fails. Tool-level errors remain protocol results. | T02 allows non-zero process exit only for startup invalidity or inability to continue serving. | `drmcp/src/cmd/design-records-mcp/main.go`; PRODUCT-TASK-SPEC-021-02 D-015. | MCP handler errors must not terminate the server process. Startup validation belongs before serving. | Exact error category strings require executor-contract freezing. |
| External proxy launch | Current repository tools launch stdio servers behind `mcp-proxy` with `--pass-environment`, repeated `--env`, `--cwd`, and a command after `--`. | T02 requires external `mcp-proxy` and process environment configuration. | `tools/grep-mcp/start.ps1`, `tools/ollama-mcp/start.ps1`, and `tools/repo-ops-mcp/start.ps1`. | Candidate: one Windows launcher or deployment guide owns environment forwarding and the Go executable command. | Exact launcher path and which repository-local `mcp-proxy` project supplies the binary remain unresolved. |
| Windows operation | Existing start scripts validate paths and parameters in PowerShell before launching the proxy. | T02 fixes Windows executable delivery. | The three `tools/*-mcp/start.ps1` files. | A launcher writer and a Windows build/smoke owner are required if a script is included. | Script authoring may be combined with command wiring only under one acceptance boundary. |

### 6. Production file candidates

The following is one non-adopted candidate layout.
`<standalone-app>` is a placeholder for the unresolved canonical app directory.

| item | observed state | relation to T02 decisions | exact Evidence | implementation-graph effect | uncertainty / next owner |
| ---- | -------------- | ------------------------- | -------------- | --------------------------- | ------------------------ |
| Candidate: command entry | `<standalone-app>/src/cmd/task-responsibility-validator/main.go` | Owns environment startup, stdio wiring, and process exit behavior. | Current pattern: `drmcp/src/cmd/design-records-mcp/main.go`. | Single writer must own command wiring and final executable build target. | Path cannot be released until the app namespace is accepted. |
| Candidate: core contract | `<standalone-app>/src/internal/taskvalidator/model.go` and `validator.go` | Owns outcome classes, criterion result validation, logical AND, and orchestration. | Validator Specification and T02 D-005 through D-007, D-010, D-011, D-018. | Core types become predecessor output for MCP and Ollama adapters if implementation is split. | File names are candidates, not accepted symbols. |
| Candidate: filesystem input | `<standalone-app>/src/internal/taskvalidator/taskfile.go` and `checklists.go` | Owns safe Task path resolution, minimal `task_type` parse, fixed checklist loading, criterion-ID extraction, and prompt composition. | Task authoring spec; validator skill; repository path-safety patterns. | This can be a parallel leaf only after shared core types are frozen. | Exact package split may remain one package to avoid shared API churn. |
| Candidate: Ollama adapter | `<standalone-app>/src/internal/taskvalidator/ollama.go` | Owns `/api/chat`, schema request, timeout, retry, response decode, and raw-response containment. | `tools/ollama-mcp/server.py`; T02 runtime decisions. | Adapter can be tested independently through a fake HTTP server after core request/result types are fixed. | Adding dependencies would create `go.mod` and `go.sum` shared writers. |
| Candidate: MCP transport | `<standalone-app>/src/internal/taskvalidatormcp/server.go`, `jsonrpc.go`, `stdio.go`, `tools.go`, and `tools_call.go` | Owns dedicated stdio MCP behavior and the one-path input schema. | Current DRMCP MCP pattern; T02 D-002 and D-003. | A separate MCP leaf is possible after the application service interface is frozen. | Server name, tool name, and protocol version remain unresolved. |
| Candidate: Windows launcher | `<standalone-app>/start.ps1` | Owns environment forwarding and external `mcp-proxy` invocation. | Existing `tools/*-mcp/start.ps1` patterns. | Must run after executable path and proxy dependency route are fixed. | T04 must decide whether this is source, deployment guidance, or a separate deliverable. |
| Shared module files | `go.mod` and `go.sum` | May change only when implementation adds a Go dependency. | Current root module contains only YAML and diff dependencies. | If touched, both files require one sole writer and must precede dependent leaves. | A standard-library-only implementation may avoid this collision. |

### 7. Test file and fixture candidates

| item | observed state | relation to T02 decisions | exact Evidence | implementation-graph effect | uncertainty / next owner |
| ---- | -------------- | ------------------------- | -------------- | --------------------------- | ------------------------ |
| Candidate: command tests | `<standalone-app>/src/cmd/task-responsibility-validator/main_test.go` | Covers startup config, stdio in-process behavior, process exit, and process smoke. | `drmcp/src/cmd/design-records-mcp/main_test.go`. | Command tests should be owned with command wiring or by a dependent integration test writer. | Exact process-smoke test name is a candidate. |
| Candidate: path and parse tests | `<standalone-app>/src/internal/taskvalidator/taskfile_test.go` | Covers `/` and `\`, absolute/drive/UNC rejection, `..`, symlink/root escape, unreadable Task, missing/invalid `task_type`, and full-text preservation. | T02 D-003; repository path-safety patterns; Task authoring spec. | Suitable focused leaf after core structural outcome types are fixed. | Symlink test behavior must use Windows-capable test setup or skip with explicit Evidence when unavailable. |
| Candidate: checklist tests | `<standalone-app>/src/internal/taskvalidator/checklists_test.go` | Covers all eleven mappings, fixed load order, criterion extraction, duplicates, missing files, and exact expected sets. | Actual prompt asset inventory and validator skill. | Can share read-only checklist fixtures but must have one fixture-family writer. | Checklist source files themselves remain protected inputs. |
| Candidate: validator tests | `<standalone-app>/src/internal/taskvalidator/validator_test.go` | Covers semantic success/false, structural failure, execution failure, strict response fields, criterion order independence, and logical AND. | Validator Specification and T02 D-005 through D-007 and D-018. | Core acceptance tests should precede adapter integration. | No model call is required. |
| Candidate: Ollama tests | `<standalone-app>/src/internal/taskvalidator/ollama_test.go` | Uses a fake HTTP server for request shape, non-streaming, schema field, timeout, two attempts, identical retry payload, 4xx/5xx, malformed response, and raw-response exclusion. | Existing Ollama HTTP pattern and Go dependency-injection test style. | Parallelizable after core adapter interface is frozen. | Timeout tests need injectable timing or short test contexts to avoid 300-second tests. |
| Candidate: MCP tests | `<standalone-app>/src/internal/taskvalidatormcp/jsonrpc_test.go`, `stdio_test.go`, `tools_test.go`, and `tools_call_test.go` | Covers initialize/list/call, exact one-path input schema, protocol versus tool errors, and server continuity. | Current DRMCP MCP tests. | Parallelizable after application service interface and external result schema are frozen. | Exact tool and server names must be accepted first. |
| Candidate: fixtures | `<standalone-app>/src/internal/taskvalidator/testdata/` | Holds Task Markdown and temporary checklist-tree fixtures, not copies of canonical production prompts. | Repository tests use `t.TempDir` and package-local test data patterns. | One fixture writer prevents parallel overwrite and hidden assumptions. | Exact fixture names belong in executor Tasks. |

### 8. Focused verification command candidates

No command in this section was executed because implementation paths do not yet exist.
Every command remains conditional on the accepted `<standalone-app>` path.

| item | observed state | relation to T02 decisions | exact Evidence | implementation-graph effect | uncertainty / next owner |
| ---- | -------------- | ------------------------- | -------------- | --------------------------- | ------------------------ |
| Candidate: core and filesystem | `go test ./<standalone-app>/src/internal/taskvalidator/...` | Verifies Task parsing, checklist loading, strict result handling, and Ollama adapter tests in the candidate single-core package. | Root Go module and existing colocated Go test pattern. | Focused implementation owner command. | Split package commands must be revised if T04 chooses finer package boundaries. |
| Candidate: MCP transport | `go test ./<standalone-app>/src/internal/taskvalidatormcp/...` | Verifies stdio, JSON-RPC, tools/list, and tools/call behavior. | Current DRMCP MCP package test pattern. | Focused MCP leaf command. | Tool/server names must be frozen. |
| Candidate: command package | `go test ./<standalone-app>/src/cmd/task-responsibility-validator/...` | Verifies startup and command wiring. | Current DRMCP command test pattern. | Focused command/integration owner command. | Depends on all imported candidate packages. |
| Candidate: process stdio smoke | `go test ./<standalone-app>/src/cmd/task-responsibility-validator -run TestProcessStdioSmoke -count=1` | Exercises the built process boundary without a live Ollama when the test injects or fakes the upstream endpoint. | `drmcp/src/cmd/design-records-mcp/main_test.go`, `TestProcessStdioSmoke`. | Runtime-smoke owner command after command wiring. | The exact test name and fake-upstream arrangement must be frozen in the implementation contract. |

### 9. Aggregate verification and smoke command candidates

| item | observed state | relation to T02 decisions | exact Evidence | implementation-graph effect | uncertainty / next owner |
| ---- | -------------- | ------------------------- | -------------- | --------------------------- | ------------------------ |
| Candidate: app aggregate | `go test ./<standalone-app>/...` | Runs every standalone package test under one gate. | Root Go module supports top-level package patterns. | Candidate objective verification owner, distinct from individual leaf verification when the graph is split. | Exact path waits for namespace resolution. |
| Candidate: repository aggregate | `go test ./...` | Detects cross-module compile and regression effects in the root module. | Existing root `go.mod`. | Final broader verification after app-focused tests pass. | This may be expensive but has one explicit owner rather than every leaf. |
| Candidate: Windows build | `go build -o bin/task-responsibility-validator.exe ./<standalone-app>/src/cmd/task-responsibility-validator` | Produces the T02-fixed executable path. | T02 D-015; existing `bin/`; current `drmcp/src/cmd/...` command layout. | Required integration gate before runtime smoke or release review. | Source path waits for namespace resolution. |
| Candidate: proxy smoke | Start the accepted PowerShell launcher, then call initialize, tools/list, and the validator tool through the external `mcp-proxy` endpoint with a fake or reachable Ollama endpoint. | Verifies the T02 external-proxy deployment boundary. | Existing `tools/*-mcp/start.ps1`; current DRMCP JSON-RPC process smoke. | Separate runtime smoke owner after build and launcher completion. | No exact command can be frozen until launcher path, port, proxy project, and tool name are accepted. |

### 10. Writer-boundary inventory

| item | observed state | relation to T02 decisions | exact Evidence | implementation-graph effect | uncertainty / next owner |
| ---- | -------------- | ------------------------- | -------------- | --------------------------- | ------------------------ |
| Production source writer candidate | One implementation owner can own the candidate core, adapters, transport, and command source when delivered as one bounded executable. | All source changes serve the one standalone validator outcome. | Candidate production inventory above; executor-ready Task authority permits one Task when ownership and acceptance are unified. | A single strong-model implementation Task is feasible after exact paths and contracts are frozen. | T04 decides whether coordination cost justifies decomposition. |
| Test writer candidate | Tests can remain with each source leaf, or one dependent test writer can own cross-package/process tests. | T02 requires package, adapter, timeout/retry, MCP, build, and smoke coverage. | Candidate test inventory above. | Parallel source/test leaves require disjoint files and frozen APIs. | T04 must assign every `_test.go` and fixture family to one writer. |
| Launcher or deployment writer candidate | One writer must own the PowerShell launcher or deployment guidance. | External `mcp-proxy` is part of the accepted runtime boundary. | Existing start-script patterns. | Must follow executable path and proxy route freeze. | Exact artifact type and path remain open. |
| Objective verification owner candidate | One verification owner should run app aggregate tests, repository tests, Windows build, and required smoke checks. | Broader verification must not be duplicated across parallel implementation leaves. | `skills/claude-code-token-budget/execution-hub-task-pattern.md`, `## Focused and aggregate verification`. | Separate gate becomes necessary when multiple leaves exist. | T04 graph coordination owner. |
| Independent review owner candidate | A reviewer independent of implementation must inspect the integrated diff, tests, and accepted contract. | Implementation and independent review have separate responsibility boundaries. | Task authoring type contract; executor-ready Task authority. | Review follows aggregate verification. | T04 graph coordination owner. |
| Closure synchronization owner candidate | One later synchronization owner updates accepted lifecycle and Evidence only after review. | Implementation, review, and lifecycle propagation must remain separate. | Task authoring synchronization contract; execution-hub authority. | Closure follows accepted review and does not repair implementation. | T04 graph coordination owner. |
| Shared-file collisions | `go.mod`, `go.sum`, candidate core types, command wiring, launcher, shared fixtures, and the built executable are shared or integration-sensitive artifacts. | T02 does not permit ambiguous concurrent writers. | Root module files; candidate inventory; execution-hub single-writer rules. | These files require one writer or sequencing. | Standard-library-only implementation may remove the module-file collision. |

### 11. Execution-hub trigger evaluation

| item | observed state | relation to T02 decisions | exact Evidence | implementation-graph effect | uncertainty / next owner |
| ---- | -------------- | ------------------------- | -------------- | --------------------------- | ------------------------ |
| Single-Task feasibility | Core logic, filesystem input, Ollama adapter, MCP transport, command wiring, and focused tests produce one executable and can share one implementation owner and one acceptance boundary. | T02 contains no requirement for multiple executors. | One executable contract in T02; execution-hub non-trigger rule for one self-contained leaf. | A single strong-model implementation Task remains valid if it owns every exact source/test path and focused verification. | Requires exact namespace, tool name, launcher treatment, and package paths first. |
| Hub trigger | The hub pattern applies when T04 splits multiple executors, uses multiple model classes, separates broader verification, creates dependent API output, or assigns shared writers. | These are the authoritative trigger conditions. | `skills/claude-code-token-budget/execution-hub-task-pattern.md`, `## Application trigger`. | Parallel core, Ollama, MCP, test, or launcher leaves would require a persistent reviewed and released graph. | T04 decides based on the final writer map, not file count. |
| Parallel leaf candidates | Candidate leaves are Task/checklist filesystem behavior, Ollama adapter behavior, and MCP transport behavior. | Each has distinct focused tests but consumes shared core result and error types. | Candidate production and test inventories. | Parallelism is safe only after one predecessor freezes the shared service interface and each writer set is disjoint. | Otherwise use one implementation owner or sequence leaves. |
| Integration gate | Aggregate app tests, repository tests, Windows build, process stdio smoke, and proxy smoke form one downstream gate. | T02 includes Windows executable and external proxy behavior beyond package-level tests. | Candidate aggregate commands and existing process/proxy patterns. | A separate objective verification Task is useful when implementation is split. | A single implementation Task may own focused build/smoke only if no separate release gate is needed. |
| Implementation Task authoring | Executor-ready leaves require exact files, symbols, behavior, tests, commands, dependencies, and owners. | T03 only supplies candidates and cannot materialize Tasks. | `skills/claude-code-token-budget/executor-ready-task-design.md`, `## Executor-ready gate`. | T04 may need an authoring or scope-freeze owner before implementation prompts are allowed. | Task IDs and graph changes remain T04 responsibility. |

### 12. Uncertainty and blockers

| item | observed state | relation to T02 decisions | exact Evidence | implementation-graph effect | uncertainty / next owner |
| ---- | -------------- | ------------------------- | -------------- | --------------------------- | ------------------------ |
| App namespace and directory | T02 requires a dedicated app namespace, but the active namespace profile contains no accepted validator namespace. | The Investigation must not invent one. | `product/records/spec/brewprint/namespaces/app-namespaces.md`. | Blocks exact source and test paths and therefore blocks executor release. | T04 must route a bounded decision before executor Task authoring if no existing authority is found. |
| MCP identity | Dedicated stdio MCP is fixed, but exact server name, tool name, and protocol version are not fixed in T02. | Future tool-name reuse is explicitly not guaranteed. | PRODUCT-TASK-SPEC-021-02 D-002 and D-016. | Blocks exact tools/list schema assertions and process-smoke requests. | T04 must route an implementation-boundary decision or freeze observable-equivalent names in an authorized executor contract. |
| Launcher ownership | External `mcp-proxy` use is fixed, but exact script path, port, and repository-local proxy dependency route are not fixed. | Environment-only configuration must still be preserved. | T02 D-014 and D-015; existing tool start scripts. | Blocks an exact proxy smoke command. | T04 graph coordination and, if necessary, a bounded decision owner. |
| Dependency writer | Root Go module has no MCP or JSON Schema library dependency. | Go and strict validation are fixed, but library choice is not. | `go.mod`. | Multiple implementation leaves could collide on `go.mod` and `go.sum`. | T04 must assign one writer if a dependency is required; implementation may avoid the collision with the standard library. |

The app namespace gap is a blocker for an executor-ready exact-path graph, not a blocker to concluding this Investigation.
All obtainable repository and runtime Evidence was recorded without adopting a namespace or implementation option.

## Cross-cutting observations

- The stable semantic contract is not the implementation bottleneck.
- The unresolved application identity controls source paths, test paths, import paths, launcher paths, and verification commands at once.
- Current DRMCP is useful as a stdio, JSON-RPC, strict-decoding, command, and testing pattern only.
- Current DRMCP cannot satisfy the standalone `task_type` seam through its existing Task model.
- The canonical checklist set is read-only input to W021 and has complete eleven-type coverage.
- Prompt order and response-result order are separate concerns. The former is fixed; the latter must be validated by criterion identity.
- The largest shared-writer risks are root module files, shared core types, command wiring, launcher ownership, and cross-package fixtures.
- Parallelism has value only after shared types and observable MCP identity are frozen.

## Follow-up judgment candidates

- Candidate: select and register the dedicated standalone app namespace and top-level directory.
- Candidate: freeze the MCP server name, tool name, tool input schema, and observable result envelope.
- Candidate: decide whether startup is a source-controlled PowerShell script or deployment guidance and identify its proxy dependency route.
- Candidate: choose one strong-model implementation Task when one owner can safely own the full source and test boundary.
- Candidate: apply an execution hub only when T04 creates multiple executors, shared writers, dependent outputs, or a separate aggregate verification owner.
- Candidate: assign one writer for `go.mod` and `go.sum` only if a new dependency is required.

These are downstream judgment candidates.
This Investigation does not select among them.

## Recommendation

T04 should first treat the missing app namespace, MCP identity, and launcher route as executor-readiness blockers.
It should route any unresolved observable choice to a bounded decision owner instead of embedding the choice in an implementation Task.

After those values are fixed, T04 should prefer one implementation Task when one owner can own all production and focused test files without shared-writer ambiguity.
If T04 splits filesystem, Ollama, MCP, command, or test work across executors, it should materialize the execution-hub pattern with:

- one shared-contract predecessor;
- disjoint writer maps;
- explicit producer and consumer order;
- focused verification per leaf;
- one aggregate verification and runtime-smoke owner;
- one independent integrated review owner;
- one closure synchronization owner.

## Follow-up artifact candidates

- An executor-ready implementation graph coordination result for PRODUCT-WORK-SPEC-021.
- A bounded decision record or decision Task for unresolved app namespace and observable MCP identity, when required by T04.
- Exact implementation Task contracts with production files, test files, symbols, behavior, commands, and stop conditions.
- An independent execution-graph review when the execution-hub trigger applies.
- A release synchronization result before implementation prompts are issued when the execution-hub trigger applies.
- A later objective verification result covering aggregate tests, Windows build, stdio process smoke, and external-proxy smoke.

No Task ID or artifact ID is reserved by this Investigation.

## Open questions

- What canonical app namespace and top-level directory identify the standalone validator?
- What exact MCP server name, tool name, and protocol version form the temporary external interface?
- Does W021 own a PowerShell launcher, deployment guidance, or both?
- Which repository-local dependency supplies `mcp-proxy` for the standalone launcher?
- Can the implementation remain standard-library-only, or will one writer need to update `go.mod` and `go.sum`?
- Does the final graph retain one implementation owner, or do real writer and verification boundaries justify an execution hub?
