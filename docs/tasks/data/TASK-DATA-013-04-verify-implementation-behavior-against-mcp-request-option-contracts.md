# TASK-DATA-013-04: Verify implementation behavior against MCP request option contracts

- **id**: TASK-DATA-013-04
- **status**: done
- **date**: 2026-06-07
- **work_item**: WORK-DATA-013
- **source_requirement**: REQ-DATA-006
- **estimate**: 1d-2d
- **depends_on**:
  - TASK-DATA-013-02
- **outputs**:
  - Implementation and test gap report for get_reference_tree, get_references, get_source, inspect, and list_endpoints request option behavior
  - Patch or follow-up recommendation if implementation behavior diverges from the updated MCP tool specs
  - Verification command results

## Goal

Verify whether the Go implementation behavior matches the MCP request option and response behavior contracts clarified by `TASK-DATA-013-02`.

This task is for implementation verification and small corrective patches only if the current behavior diverges from the accepted specs.

## Work

- Inspect implementation and tests for the following behavior:
  - `get_reference_tree.depth` rejects values outside `0..4` with `invalid_depth`.
  - `get_reference_tree.direction` rejects unknown values with `unsupported_direction`.
  - `get_references.direction` defaults to `out` when omitted and rejects unknown values with `unsupported_direction`.
  - `get_source.fallback` defaults to `file`, supports `file` / `error`, rejects unknown values with `invalid_args`, and returns `source_range_unavailable` when `fallback=error` and object range is unavailable.
  - `inspect.detail` defaults to `normal` and rejects unknown values with `unsupported_detail`.
  - `list_endpoints.api_table_id` omitted behavior returns all API Tables grouped by `tables[]`.
- Add or adjust tests only where the behavior is supposed to be enforced by current implementation scope.
- Keep any implementation patch limited to matching the accepted MCP tool contract.
- Do not introduce DATA DSL numeric range/default/fallback syntax.

## Done condition

- The implementation behavior is either confirmed to match the updated specs, or the divergence is patched or recorded as a follow-up.
- Relevant tests or runtime smoke commands are run and recorded.
- No unrelated MCP identity, selector matrix, tagged union, recursive structure, DAG TypeRef hint, or DATA DSL extension work is introduced.

## Verification

- Run targeted Go tests for the MCP/query implementation area identified during inspection.
- Run broader relevant tests if implementation files are changed.
- Record command results and any remaining divergence in Evidence.

## Evidence
Verdict: PASS after a small MCP error-code patch and targeted test additions.

Files updated:

- `internal/query/source.go`
  - Unknown `get_source.fallback` values now surface as invalid fallback input instead of generic unsupported behavior.
  - `fallback=error` with unavailable object-level source range now returns an error message carrying `source_range_unavailable`.
- `internal/mcp/server.go`
  - MCP error-code mapping now returns `invalid_args` for invalid fallback values.
  - MCP error-code mapping now returns `source_range_unavailable` for `get_source(fallback=error)` unavailable-range failures.
- `internal/mcp/server_test.go`
  - Added/expanded tests for the WORK-DATA-013 request option behavior contracts.

Behavior verification:

- `get_reference_tree.depth` outside `0..4` -> `invalid_depth`: verified for `depth=5` and `depth=-1`.
- `get_reference_tree.direction` unknown -> `unsupported_direction`: verified.
- `get_references.direction` omitted -> `out`: verified by response `direction`.
- `get_references.direction` unknown -> `unsupported_direction`: verified.
- `get_source.fallback` omitted -> `file` behavior when object range is unavailable: verified by whole-file fallback response with `fallback: "file"` and `source_range_unavailable` warning diagnostic.
- `get_source.fallback=file/error` behavior: verified for omitted/file-equivalent fallback response and `fallback=error` error branch.
- `get_source.fallback` unknown -> `invalid_args`: patched and verified.
- `get_source(fallback=error)` with unavailable object range -> `source_range_unavailable`: patched and verified.
- `inspect.detail` omitted -> normal-compatible response: verified as accepted omitted behavior; MCP v1 strict detail output differences remain implementation-optional per spec.
- `inspect.detail` unknown -> `unsupported_detail`: verified.
- `list_endpoints.api_table_id` omitted -> all API Tables grouped by `tables[]`: verified.

Verification performed:

- `gofmt -w internal/query/source.go internal/mcp/server.go internal/mcp/server_test.go` -> PASS.
- `go test ./internal/mcp ./internal/query` -> PASS.
- `go test ./...` -> PASS.

No DATA DSL numeric range/default/fallback syntax was introduced, and no unrelated MCP identity, selector matrix, tagged union, recursive structure, DAG TypeRef hint, or unrelated DATA work was reopened.

Remaining follow-up: none for the checked behavior contracts.
