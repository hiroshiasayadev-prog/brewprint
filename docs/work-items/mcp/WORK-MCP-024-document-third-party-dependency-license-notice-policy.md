# WORK-MCP-024: Document third-party dependency license notice policy

- **id**: WORK-MCP-024
- **status**: done
- **date**: 2026-06-07
- **source_requirement**: REQ-MCP-029
- **impact_refs**:
  - REQ-MCP-029
- **tasks**:
  - TASK-MCP-024-01
  - TASK-MCP-024-02

## Goal

Define a lightweight repository policy and location for third-party dependency license notices triggered by MCP authoring diff dependencies.

## Scope

- Add a reusable repository notice document for third-party dependency notices.
- Link the notice document from `docs/doc-policy.md`.
- Capture the current `github.com/pmezard/go-difflib` dependency notice metadata.
- Close REQ-MCP-029 without introducing SBOM automation, vendoring, license replacement, or full historical dependency audit.

## Boundary

This work item belongs to the MCP/project maintenance boundary. The immediate trigger is the MCP authoring diff implementation, but the resulting notice path is repo-level documentation hygiene.

## Done condition

- `docs/third-party-notices.md` exists and covers `github.com/pmezard/go-difflib` at minimum.
- `docs/doc-policy.md` points to `docs/third-party-notices.md` as the notice-management location.
- REQ-MCP-029 is accepted with close evidence.

## Evidence

- TASK-MCP-024-01 completed the current notice target review for `github.com/pmezard/go-difflib` / `github.com/pmezard/go-difflib/difflib`.
- TASK-MCP-024-02 created `docs/third-party-notices.md` and linked it from `docs/doc-policy.md`.
- Notice policy distinguishes `go.mod` / `go.sum` module metadata from source / binary redistribution notice obligations.
- Scope remained limited to lightweight notice documentation; no SBOM automation, vendoring, dependency replacement, or full historical license audit was introduced.
