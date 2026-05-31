# TASK-DATA-008-01: Reproduce and localize UC-002 duplicate task QID issue

- **id**: TASK-DATA-008-01
- **status**: done
- **date**: 2026-06-01
- **work_item**: WORK-DATA-008
- **source_requirement**: REQ-DATA-002
- **estimate**: 0.5d-1d
- **depends_on**:
- **outputs**:
  - UC-002 duplicate task QID / unresolved flow task reproduction evidence
  - Root-cause class recommendation for follow-up correction

## Goal

Reproduce and localize the pre-existing UC-002 duplicate task QID / unresolved flow task issue before any YAML, resolver, validator, fixture, renderer, or test correction is attempted.

This task does not fix the issue. It establishes the evidence needed to decide whether the next correction should target YAML identity, validation behavior, resolver behavior, fixture drift, diagnostic cascade handling, or another localized cause.

## Work

- Read the repo instructions and the current WORK-DATA-008 boundary before running commands.
- Reproduce the current UC-002 duplicate task QID / unresolved flow task diagnostic using the smallest command set that exposes the issue.
- Capture the exact command, exit result, and diagnostic output needed to make the problem reproducible.
- Identify the smallest file / fixture / resolver / validation path involved in the issue.
- Classify the likely root-cause bucket:
  - YAML identity problem.
  - Validation behavior problem.
  - Resolver behavior problem.
  - Fixture drift.
  - Diagnostic cascade.
  - Other localized cause.
- Recommend the next task boundary for deciding and applying the correction.

## Done Condition

- The issue is reproducible from a clean repo checkout using documented commands.
- The observed diagnostics are recorded without being normalized away or mixed with unrelated DATA expressiveness debt.
- The likely root-cause class is identified enough to decide the boundary of the next task.
- No implementation, YAML, fixture, golden, renderer, validator, parser, or MCP public contract behavior is changed by this task.

## Verification

- Run the minimal reproduction command(s) and record the exact diagnostic output.
- If available, run the smallest relevant `go test` target that covers the same path and record whether it passes or fails.
- Confirm the reproduction does not require reopening M15, WORK-DATA-001, WORK-DATA-002, WORK-DATA-003, or WORK-DATA-004.
- Confirm ADR-073 tagged union support, ADR-074 DAG asset TypeRef hint, ADR-078 / ADR-079 / ADR-080 MCP identity, broad UC-002 notes retreat cleanup, and helper model / model-file render redesign remain outside this task.

## Evidence

Completed on 2026-06-01.

### Current HEAD / working tree result

The UC-002 duplicate task QID / unresolved flow task issue does not reproduce on the current working tree.

Commands executed by the implementation agent:

```text
go run ./cmd/brewprint validate --yaml-root docs\uc\002-brewprint-self-hosting\yaml
```

Result:

```text
ok
```

```text
go run ./cmd/brewprint validate --yaml-root docs\uc\002-brewprint-self-hosting\yaml --format json
```

Result:

```json
{"diagnostics":null,"error_count":0,"warning_count":0}
```

```text
go run ./cmd/brewprint render --yaml-root docs\uc\002-brewprint-self-hosting\yaml --out $env:TEMP\brewprint-uc002-data008-render --clean
```

Result:

```text
rendered 40 file(s)
```

```text
go test ./cmd/brewprint ./internal/resolve ./internal/render/placement ./internal/render/dag
```

Result:

```text
ok  	github.com/hiroshiasayadev-prog/brewprint/cmd/brewprint	(cached)
ok  	github.com/hiroshiasayadev-prog/brewprint/internal/resolve	(cached)
ok  	github.com/hiroshiasayadev-prog/brewprint/internal/render/placement	(cached)
ok  	github.com/hiroshiasayadev-prog/brewprint/internal/render/dag	(cached)
```

```text
go test ./...
```

Result: exit 0; all packages passed.

### Historical reproduction

The historical issue reproduces from clean snapshot `fe12ef6` with the smallest command:

```text
go run ./cmd/brewprint validate --yaml-root docs\uc\002-brewprint-self-hosting\yaml
```

The historical diagnostic was duplicate public QIDs for repeated file-local task helper IDs, followed by unresolved flow steps in the same files. A representative pattern was:

```text
error duplicate_node mcp/task/get_reference_tree.yaml: duplicate node qid: mcp.task.build_response
error duplicate_node mcp/task/get_reference_tree.yaml: duplicate node qid: mcp.task.query_service
error duplicate_node mcp/task/get_reference_tree.yaml: duplicate node qid: mcp.task.validate_request
error unresolved_flow_task mcp/task/get_reference_tree.yaml: unresolved flow step: build_response
error unresolved_flow_task mcp/task/get_reference_tree.yaml: unresolved flow step: query_service
error unresolved_flow_task mcp/task/get_reference_tree.yaml: unresolved flow step: validate_request
```

The same pattern appeared across UC-002 MCP task files such as:

- `mcp/task/get_reference_tree.yaml`
- `mcp/task/get_references.yaml`
- `mcp/task/get_signature.yaml`
- `mcp/task/get_source.yaml`
- `mcp/task/inspect.yaml`
- `mcp/task/list_endpoints.yaml`
- `mcp/task/list_objects.yaml`

The historical run ended with:

```text
validation failed: 42 error(s), 0 warning(s)
exit status 1
```

### Localization

Smallest historical reproducing command:

```text
go run ./cmd/brewprint validate --yaml-root docs\uc\002-brewprint-self-hosting\yaml
```

Render is not the smaller reproducer because render validates first and aborts before project rendering.

Minimal affected paths:

- UC-002 YAML: `docs/uc/002-brewprint-self-hosting/yaml/mcp/task/*.yaml`.
- Repeated file-local task helper IDs: `validate_request`, `query_service`, `build_response`.
- Resolver path: `internal/resolve/builder.go` -> `internal/resolve/symbols.go` -> `internal/resolve/flow.go`.
- Validation / CLI path: `cmd/brewprint/main.go`.
- Render cascade path: render aborts after validation before `internal/render/project.Render`.

### Root-cause classification

Primary bucket: resolver behavior problem.

Secondary bucket: diagnostic cascade.

Not YAML identity problem: ADR-058 and the current specs say file-private subnodes may share local IDs across files. The UC-002 YAML pattern is valid.

Not fixture drift as the root cause: render fixture state may be dirty, but the duplicate QID diagnostic came from resolver indexing, not golden mismatch.

The historical failure happened because non-main sub tasks were registered as public `module.kind.id` QIDs such as `mcp.task.validate_request`. Once a second task file reused the same helper IDs, `symbols.addNode` emitted `duplicate_node` and skipped adding those nodes to `NodesByFile`, which cascaded into same-file `flow.step` resolution as `unresolved_flow_task`.

Current code fixes this by using file-private internal IDs such as `mcp/task/get_signature.yaml#validate_request`, excluding private subnodes from `NodesByQID`, and resolving bare flow IDs through same-file private nodes first.

### Conflict / stale follow-up result

`WORK-DATA-008` and this task were created from stale M15 close evidence that still treated the UC-002 duplicate task QID / unresolved flow task issue as unresolved.

Current evidence shows the issue was already resolved by the later `WORK-RESOLVE-001` / ADR-058-aligned resolver behavior. Therefore no DATA YAML, resolver, renderer, validator, fixture, golden, MCP contract, ADR-073, ADR-074, ADR-078, ADR-079, or ADR-080 work is needed for this blocker.

### Verification result

- Current UC-002 validation passes.
- Current UC-002 JSON validation output has no diagnostics.
- Current UC-002 temp render succeeds.
- Current targeted Go tests pass.
- Current full `go test ./...` passes.
- Historical reproduction confirms the old issue and localizes it to resolver behavior plus diagnostic cascade.
- No repository files were modified by the implementation agent during reproduction; only temporary archive / extract / render directories under `%TEMP%` were used.
