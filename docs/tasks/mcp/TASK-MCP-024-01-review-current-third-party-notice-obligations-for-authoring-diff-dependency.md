# TASK-MCP-024-01: Review current third-party notice obligations for authoring diff dependency

- **id**: TASK-MCP-024-01
- **status**: done
- **date**: 2026-06-07
- **work_item**: WORK-MCP-024
- **source_requirement**: REQ-MCP-029
- **estimate**: 0.5d
- **depends_on**:
- **outputs**:
  - current dependency notice inventory
  - go-difflib notice metadata confirmation

## Goal

Confirm the current third-party dependency notice scope created by the authoring diff implementation.

## Steps

- Inspect `go.mod` for the external dependency added by REQ-MCP-027 / TASK-MCP-023-03.
- Confirm the dependency purpose and required notice metadata from REQ-MCP-029 evidence.
- Keep the scope limited to the current dependency and the lightweight notice-management path.

## Done condition

- The current dependency notice target is identified as `github.com/pmezard/go-difflib` / `github.com/pmezard/go-difflib/difflib`.
- The notice metadata needed by TASK-MCP-024-02 is known.
- No SBOM automation, vendoring, or full historical license audit is introduced.

## Work

- Inspected `go.mod` and confirmed `github.com/pmezard/go-difflib v1.0.0` is the external dependency added for authoring diff behavior.
- Confirmed REQ-MCP-029 evidence identifies the imported package as `github.com/pmezard/go-difflib/difflib` and the license family as BSD 3-Clause style.
- Kept scope limited to notice-management documentation; no SBOM automation, vendoring, replacement implementation, or full historical license audit was introduced.

## Verification

- `go.mod` contains `github.com/pmezard/go-difflib v1.0.0`.
- REQ-MCP-029 acceptance criteria identify the notice metadata needed for TASK-MCP-024-02.

## Evidence

- Current dependency notice target: `github.com/pmezard/go-difflib` / `github.com/pmezard/go-difflib/difflib`.
- Notice metadata needed downstream: module path, imported package, purpose, license family, notice requirement summary, and upstream license pointer.
