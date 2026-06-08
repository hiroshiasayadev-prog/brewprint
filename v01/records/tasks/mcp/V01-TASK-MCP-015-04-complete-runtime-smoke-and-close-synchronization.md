# V01-TASK-MCP-015-04: Complete runtime smoke and close synchronization

- **id**: V01-TASK-MCP-015-04
- **status**: done
- **date**: 2026-06-03
- **work_item**: V01-WORK-MCP-015
- **source_requirement**: V01-REQ-MCP-016
- **estimate**: 0.5d
- **depends_on**:
  - V01-TASK-MCP-015-03
- **outputs**:
  - Runtime smoke evidence for MCP propose_record_update named_section_replace spacing
  - Close evidence and status synchronization for TASK-MCP-015, V01-WORK-MCP-015, and V01-REQ-MCP-016

## Goal

Verify the corrected behavior through the MCP runtime path and close the `V01-REQ-MCP-016` workflow artifacts when implementation evidence is complete.

## Work

- Run runtime smoke through the Design Records MCP server for `propose_record_update` `named_section_replace`.
- Include smoke cases that exercise replacement body trailing newline variants through the public tool path.
- Record verification commands, inputs, and observed output shape.
- Update task evidence, work item evidence, and requirement close evidence / status when all checks pass.
- Validate the affected design records after close synchronization.

## Done condition

- Runtime smoke confirms section content and the next heading remain separated by the intended blank line.
- Relevant Go tests pass after the implementation change.
- `V01-TASK-MCP-015-01..04`, `V01-WORK-MCP-015`, and `V01-REQ-MCP-016` are synchronized with evidence and final statuses.
- Design Records MCP validation for the affected records reports no diagnostics.

## Verification

- Run the MCP runtime smoke command and record the result.
- Run relevant package tests and, if reasonable, `go test ./...`.
- Run Design Records MCP validation for the affected REQ / WORK / TASK range.

## Evidence

### Test results

- `go test ./internal/designrecords -run TestReplaceNamedSectionSpacing -v`: pass.
  - `TestReplaceNamedSectionSpacingPreservation/no_trailing_newline`: pass.
  - `TestReplaceNamedSectionSpacingPreservation/one_trailing_newline`: pass.
  - `TestReplaceNamedSectionSpacingPreservation/already_separated`: pass.
  - `TestReplaceNamedSectionSpacingLastSection/no_trailing_newline`: pass.
  - `TestReplaceNamedSectionSpacingLastSection/one_trailing_newline`: pass.
- `go test ./internal/designrecords ./internal/designrecordsmcp`: pass.
- `go test ./...`: pass.

### Runtime smoke

- Command path: `go run ./cmd/design-records-mcp --root .`.
- Protocol path: stdio JSON-RPC `initialize` / `notifications/initialized` / `tools/call`.
- Public tool call: `propose_record_update` with `kind=requirement`, `id=V01-REQ-MCP-016`, `update.type=named_section_replace`, `section_selector.heading=Requirement`.
- Repository files were not accepted or written by the smoke. Each retained proposal was discarded with `discard_proposed_write`.
- Smoke cases:
  - no trailing newline body: `Runtime smoke replacement no trailing newline.`
  - one trailing newline body: `Runtime smoke replacement one trailing newline.\n`
  - already separated body: `Runtime smoke replacement already separated.\n\n`

Confirmed proposal diff shape for all cases:

```text
+Runtime smoke replacement ...
+
+## Evidence
```

The failing shape was not present:

```text
+Runtime smoke replacement ...
+## Evidence
```

Proposal IDs observed and discarded in the smoke process:

- `pw_000002`: no trailing newline, discarded, no write.
- `pw_000004`: one trailing newline, discarded, no write.
- `pw_000006`: already separated, discarded, no write.
