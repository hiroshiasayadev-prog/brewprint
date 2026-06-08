# V01-TASK-MCP-005-06: Verification and close evidence

- **id**: V01-TASK-MCP-005-06
- **status**: done
- **date**: 2026-06-01
- **work_item**: V01-WORK-MCP-005
- **source_requirement**: V01-REQ-MCP-005
- **estimate**: 0.5d
- **depends_on**:
  - V01-TASK-MCP-005-05
- **outputs**:
  - close verification evidence for V01-WORK-MCP-005
  - V01-WORK-MCP-005 status synchronization
  - V01-REQ-MCP-005 status synchronization if appropriate

## Goal

Verify that V01-WORK-MCP-005 satisfies the authoring guidance discovery / retrieval requirement and record close evidence before marking the work item done.

## Work

- Review V01-REQ-MCP-005, V01-WORK-MCP-005, and V01-TASK-MCP-005-01 through V01-TASK-MCP-005-05 for stale wording or relation mismatch.
- Confirm `docs/guides/*.md` guide source model is documented and implemented.
- Confirm `list_authoring_guides` exposes only `id`, `title`, and `abstract`.
- Confirm `get_authoring_guidance` exposes only `id`, `title`, and raw Markdown `content`.
- Confirm source path is not exposed as public tool response contract.
- Confirm guides are not Design Records record kinds and are not included in record index / resolver / validation targets.
- Confirm doc-policy and README-style entry points now refer to guide IDs where appropriate.
- Record targeted test evidence.
- Record known out-of-scope repository-wide test failures, if still present.
- Update V01-WORK-MCP-005 close evidence and status if close conditions are satisfied.
- Update V01-REQ-MCP-005 status if the requirement is satisfied.

## Done condition

- Verification evidence is recorded in this task.
- V01-WORK-MCP-005 has close evidence.
- V01-WORK-MCP-005 status is updated consistently with the verification result.
- V01-REQ-MCP-005 status is updated if appropriate.
- Any remaining follow-up is explicitly identified and not left ambiguous.

## Verification

- Reviewed `V01-REQ-MCP-005`; no stale section / scope retrieval wording remains.
- Reviewed `V01-REQ-MCP-005`; guide metadata outcome is `id` / `title` / `abstract` without source path exposure.
- Reviewed `V01-WORK-MCP-005`; all tasks `V01-TASK-MCP-005-01` through `V01-TASK-MCP-005-06` are materialized.
- Reviewed `docs/doc-policy.md`; it is now a thin startup policy and delegates artifact-specific rules to authoring guidance tools.
- Reviewed `docs/guides/*.md`; all initial guide files use `## Migration Note` instead of `## Source`.
- Confirmed guide migration notes explicitly state they are not instructions to read original files and are not part of the public retrieval contract.
- Confirmed `list_authoring_guides` contract exposes only `id`, `title`, and `abstract`.
- Confirmed `get_authoring_guidance` contract exposes `id`, `title`, and raw Markdown `content` by guide ID.
- Confirmed source path is not exposed as public tool response contract.
- Confirmed guides are not Design Records record kinds and are not resolver / validation targets.
- Confirmed targeted implementation tests were reported passing for guidance tooling.
- Confirmed repository-wide `go test ./...` has known render / manifest failures outside V01-WORK-MCP-005 scope.

## Evidence

- Guide source files verified:
  - `docs/guides/adr-authoring.md`
  - `docs/guides/spec-authoring.md`
  - `docs/guides/requirement-authoring.md`
  - `docs/guides/work-item-authoring.md`
  - `docs/guides/task-authoring.md`
  - `docs/guides/investigation-authoring.md`
  - `docs/guides/artifact-boundary.md`
- `docs/doc-policy.md` now points startup authoring flow to:
  - `list_authoring_guides`
  - `get_authoring_guidance`
  - guide IDs such as `spec-authoring`, `adr-authoring`, `requirement-authoring`, `work-item-authoring`, `task-authoring`, `investigation-authoring`, and `artifact-boundary`
- Targeted implementation test evidence reported by implementation session:
  - `gofmt -w ...`: passed
  - `go test ./internal/designrecords ./internal/designrecordsmcp`: passed
  - `go test ./cmd/design-records-mcp ./internal/designrecords ./internal/designrecordsmcp`: passed
- Known out-of-scope failure reported by implementation session:
  - `go test ./...`: failed in render / manifest expectations outside V01-WORK-MCP-005 scope.
- Close result:
  - V01-WORK-MCP-005 can be marked `done`.
  - V01-REQ-MCP-005 can be marked `accepted`.
  - No remaining V01-WORK-MCP-005 follow-up is required for phase 1.
