# TASK-MCP-005-04: Implementation and tests

- **id**: TASK-MCP-005-04
- **status**: done
- **date**: 2026-06-01
- **work_item**: WORK-MCP-005
- **source_requirement**: REQ-MCP-005
- **estimate**: 1d
- **depends_on**:
  - TASK-MCP-005-03
- **outputs**:
  - Design Records MCP authoring guidance tool implementation
  - automated tests for guide listing, retrieval, errors, and record-index boundary

## Goal

Implement the `list_authoring_guides` and `get_authoring_guidance` MCP tools according to the WORK-MCP-005 / REQ-MCP-005 contract without changing existing Design Records record retrieval, validation, or reference resolution behavior.

## Work

- Investigate existing Design Records MCP tool registration, request / response schema handling, Markdown parsing utilities, and tests.
- Implement guide discovery from `docs/guides/*.md`.
- Implement `list_authoring_guides` response with `id`, `title`, and `abstract` only.
- Implement `get_authoring_guidance` response with `id`, `title`, and raw Markdown `content` only.
- Add tool errors for missing guide ID and invalid requests.
- Add tests that guide files are not added to the Design Records record index.
- Run `gofmt` and Go tests.

## Done condition

- `list_authoring_guides` is registered and returns all guide Markdown files in ID lexical order.
- `get_authoring_guidance` is registered and returns raw Markdown content for exact guide IDs.
- Guide tool responses do not expose source paths.
- Unknown guide ID returns `guide_not_found`.
- Invalid guide requests return `invalid_request`.
- Existing record index / `get_record` / `get_records` / `validate_records` / `resolve_reference` behavior is not expanded to guide files.
- Automated tests cover the agreed contract.

## Verification

- Ran `gofmt` on the changed Design Records MCP implementation and test files.
- Ran targeted Design Records MCP tests:
  - `go test ./internal/designrecords ./internal/designrecordsmcp`
  - `go test ./cmd/design-records-mcp ./internal/designrecords ./internal/designrecordsmcp`
- Ran repository-wide tests:
  - `go test ./...`

## Evidence

- Implemented authoring guidance retrieval in `internal/designrecords/authoring_guidance.go`.
- Added public response/request structs and `guide_not_found` tool error code in `internal/designrecords/types.go`.
- Registered `list_authoring_guides` and `get_authoring_guidance` in `internal/designrecordsmcp/tools.go`.
- Wired MCP `tools/call` handling in `internal/designrecordsmcp/tools_call.go`.
- Added tests in `internal/designrecords/authoring_guidance_test.go` for:
  - listing all `docs/guides/*.md`
  - `id` lexical ordering
  - `id` / `title` / `abstract` only in list response
  - no source path exposure
  - raw Markdown content retrieval
  - `guide_not_found`
  - `invalid_request`
  - guide files not being added to the Design Records record index
- Added / updated MCP transport tests in `internal/designrecordsmcp/tools_call_test.go` and `internal/designrecordsmcp/jsonrpc_test.go` for tool registration, response shape, and tool errors.
- Test result: `go test ./internal/designrecords ./internal/designrecordsmcp` passed.
- Test result: `go test ./cmd/design-records-mcp ./internal/designrecords ./internal/designrecordsmcp` passed.
- Test result: `go test ./...` failed outside this task scope in render expectation tests:
  - `cmd/brewprint`: `TestRunRender` and `TestRunRenderCleanRemovesStaleFiles` expected a different rendered file count.
  - `internal/render/project`: `TestRenderUC001Manifest` saw additional `model-*.md` render outputs.
  - These failures are in renderer / fixture expectations and are not caused by the Design Records MCP guidance implementation files touched by this task.
