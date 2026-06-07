# REQ-MCP-032: Expose workflow status-gated required sections in authoring guidance

- **id**: REQ-MCP-032
- **status**: accepted
- **date**: 2026-06-07
- **source_refs**:
  - SPEC-design-records-mcp-tools
- **work_items**:
  - WORK-MCP-029

## Requirement

Authoring guidance for workflow artifacts must explicitly expose the MCP validation contract for valid statuses and status-gated required narrative sections.

The current MCP tools spec defines status-gated required narrative sections, but practical authoring guides do not consistently present them as authoring-time rules. In particular, work item guidance does not include `## Evidence` in the format even though `work_item` records with `status: done` require non-empty `Goal`, `Boundary`, and `Evidence` sections.

## Evidence

- `SPEC-design-records-mcp-tools` defines `Required narrative section policy` for `work_item done`, `task done`, and `requirement accepted`.
- Before this change, `docs/guides/work-item-authoring.md` listed the work item format without `## Evidence`, creating a stale guide relative to MCP validation.
- `docs/guides/task-authoring.md` lists required sections, but does not clearly state that they become MCP validation gates when `status: done`.
- `docs/guides/requirement-authoring.md` lists requirement sections, but does not clearly state that `Requirement` and `Required Outcome` become MCP validation gates when `status: accepted`.

## Required Outcome

- Work item authoring guidance includes `## Evidence` in the canonical format.
- Work item authoring guidance states that `status: done` requires non-empty `Goal`, `Boundary`, and `Evidence` sections.
- Task authoring guidance states that `status: done` requires non-empty `Goal`, `Work`, `Done condition`, `Verification`, and `Evidence` sections.
- Requirement authoring guidance states that `status: accepted` requires non-empty `Requirement` and `Required Outcome` sections.
- The MCP tools spec exposes a kind-by-kind valid status matrix, or otherwise points readers from authoring behavior to the existing validation policy.

## Explicitly Excluded Scope

- No change to the validation semantics themselves.
- No change to record status vocabularies.
- No new authoring transaction operation type.
- No implementation change unless tests or docs reveal an existing mismatch.

## Boundary

This requirement is a documentation and authoring guidance visibility fix. The MCP validation behavior remains owned by `SPEC-design-records-mcp-tools` and the implementation.
