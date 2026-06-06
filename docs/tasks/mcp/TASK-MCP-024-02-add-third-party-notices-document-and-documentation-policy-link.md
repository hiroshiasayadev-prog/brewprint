# TASK-MCP-024-02: Add third-party notices document and documentation policy link

- **id**: TASK-MCP-024-02
- **status**: done
- **date**: 2026-06-07
- **work_item**: WORK-MCP-024
- **source_requirement**: REQ-MCP-029
- **estimate**: 0.5d
- **depends_on**:
  - TASK-MCP-024-01
- **outputs**:
  - docs/third-party-notices.md
  - docs/doc-policy.md link update
  - REQ-MCP-029 close evidence

## Goal

Add the lightweight third-party notices document and link it from the documentation policy.

## Steps

- Create `docs/third-party-notices.md` as the repository location for third-party dependency notices.
- Add the `github.com/pmezard/go-difflib` notice entry with module path, purpose, license family, copyright notice, and redistribution notice expectations.
- Link `docs/third-party-notices.md` from `docs/doc-policy.md`.
- Update close evidence for this work item and REQ-MCP-029.

## Done condition

- `docs/third-party-notices.md` exists.
- `docs/doc-policy.md` clearly identifies `docs/third-party-notices.md` as the notice-management location.
- REQ-MCP-029 acceptance criteria are satisfied without expanding into SBOM automation, vendoring, or full legal review.

## Work

- Created `docs/third-party-notices.md` as the repository notice-management location.
- Added the `github.com/pmezard/go-difflib v1.0.0` notice entry with module path, imported package, purpose, license family, copyright notice, redistribution expectations, endorsement restriction, and upstream license pointer.
- Linked `docs/third-party-notices.md` from `docs/doc-policy.md` under a lightweight Third-party notices section.

## Verification

- `docs/third-party-notices.md` exists and contains the go-difflib notice entry.
- `docs/doc-policy.md` identifies `docs/third-party-notices.md` as the notice-management location and distinguishes `go.mod` / `go.sum` metadata from redistribution notice obligations.
- Scope remained limited to documentation; no SBOM automation, vendoring, or full legal review was added.

## Evidence

- Files updated: `docs/third-party-notices.md`, `docs/doc-policy.md`.
- REQ-MCP-029 acceptance criteria satisfied by the notice document and policy link.
