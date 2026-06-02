# TASK-MCP-011-03: Implement authoring proposal fixes

- **id**: TASK-MCP-011-03
- **status**: done
- **date**: 2026-06-02
- **work_item**: WORK-MCP-011
- **source_requirement**: REQ-MCP-011
- **estimate**: 1d
- **depends_on**:
  - TASK-MCP-011-02
- **outputs**:
  - Design Records MCP authoring create input normalization implementation
  - Design Records MCP proposal validation scope isolation implementation
  - Implementation notes for regression tests and runtime smoke

## Goal

Implement the Design Records MCP authoring proposal behavior specified by `TASK-MCP-011-02`.

## Work

- Reject invalid create content-source combinations before proposal or body-cache creation.
- Use the resolved top-level target ID as the metadata ID source instead of requiring `fields.id`.
- Normalize or compare request domain and ID domain according to the updated public spec.
- Scope propose-time and accept-time pre-write validation to the affected record set.
- Preserve semantic reference sources when building hypothetical proposal indexes.
- Keep unrelated repository diagnostics out of proposal-local blocking diagnostics.
- Do not broaden into unrelated repository validation cleanup.

## Done condition

- Implementation follows the updated `SPEC-design-records-mcp-tools` and `SPEC-design-records-mcp-schema` contract.
- Observed TASK-MCP-011-01 failures have corresponding implementation changes or explicit notes for TASK-MCP-011-04 tests.
- Targeted authoring tests pass or failing test gaps are documented for TASK-MCP-011-04.
- No unrelated repository validation errors are fixed as part of this task.

## Verification

Recommended commands:

```powershell
go test ./internal/designrecords ./internal/designrecordsmcp -run Authoring

go test ./internal/designrecords ./internal/designrecordsmcp
```

## Evidence

2026-06-02: Implemented Design Records MCP authoring proposal fixes in `internal/designrecords/authoring.go`.

- Create input normalization now rejects `fields` with `body` / `body_cache_id` as `invalid_request` before proposal or body-cache creation, while preserving `body` + `body_cache_id` as `invalid_body_source`.
- Structured workflow create now uses the top-level target ID for rendered metadata and no longer requires `fields.id`.
- Supplied `fields.id` is treated as a duplicate consistency input: exact IDs must match after canonical ID normalization, and `new` placeholder creates reject supplied `fields.id`.
- Workflow create domain comparison is case-insensitive; canonical target domains remain uppercase and repository path domains remain lowercase.
- Proposal-time and accept-time validation now returns proposal-local diagnostics for the affected record set only: proposed target files and reciprocal related files actually modified by the same proposal.
- Hypothetical proposal indexes now preserve existing `SemanticRefSources` for unchanged records and replace only affected paths.
- Unrelated repository validation errors were not fixed.

Regression coverage was added in `internal/designrecords/authoring_test.go` for create content-source rejection, optional / matching / mismatching / placeholder `fields.id`, lowercase domain input, unrelated diagnostic isolation, and semantic ref source preservation.

Verification run:

```powershell
go test ./internal/designrecords ./internal/designrecordsmcp -run Authoring
go test ./internal/designrecords ./internal/designrecordsmcp
```

Both commands passed on 2026-06-02.

Remaining TASK-MCP-011-04-level confirmation points, without creating `TASK-MCP-011-04` in this task:

- runtime no-write smoke through the MCP JSON-RPC/tool surface for representative create/update proposals;
- broader response-shape review if `repository_health` or another separate health field is later added;
- additional edge-case tests for multi-file reciprocal proposals with affected-record validation diagnostics on related records.
