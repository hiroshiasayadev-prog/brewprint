# WORK-MCP-015: Preserve named section replacement spacing

- **id**: WORK-MCP-015
- **status**: done
- **date**: 2026-06-03
- **source_requirement**: REQ-MCP-016
- **impact_refs**:
  - SPEC-design-records-mcp-tools
- **tasks**:
  - TASK-MCP-015-01
  - TASK-MCP-015-02
  - TASK-MCP-015-03
  - TASK-MCP-015-04

## Goal

Fix `propose_record_update` `named_section_replace` formatting so a replaced section remains separated from the next same-level heading by the canonical Markdown blank line.

This work item completes `REQ-MCP-016` by ensuring section-only update output preserves readable spacing independent of whether the caller-supplied replacement body ends with a newline.

## Boundary

This work item owns only `named_section_replace` section replacement formatting for `propose_record_update`.

In scope:

- current splice / replacement behavior inventory
- canonical separator decision for replaced section content followed by the next heading
- implementation and regression tests for body ending with no newline, one newline, and multiple newlines
- runtime smoke evidence for `propose_record_update` `named_section_replace`
- close evidence synchronization for `REQ-MCP-016`

Out of scope:

- Markdown whole-document formatter introduction
- `propose_record_create` contract changes
- metadata block replacement semantics changes
- section selector matching semantics changes
- unrelated design record content rewrite

## Impact Scope

- `SPEC-design-records-mcp-tools` if the public `propose_record_update` formatting contract needs explicit documentation.
- `internal/designrecords` authoring replacement / render logic.
- `internal/designrecords` and/or `internal/designrecordsmcp` regression tests.

## Task flow

1. Inventory current `named_section_replace` formatting and reproduce the spacing loss.
2. Decide and document the minimal canonical separator rule.
3. Implement the replacement formatting fix and regression tests.
4. Run runtime smoke and synchronize close evidence.

## Task Candidates

- Inventory current named section replacement spacing behavior.
- Decide spacing contract and update spec or guidance if needed.
- Implement named section replacement spacing normalization and regression tests.
- Complete runtime smoke and close synchronization.

## Completion Condition

- `named_section_replace` keeps exactly the intended readable separator between replaced section content and the next heading regardless of replacement body trailing newline shape.
- Regression tests cover no trailing newline, one trailing newline, and already-separated replacement bodies.
- Runtime smoke verifies the MCP `propose_record_update` path, not only unit-level helper behavior.
- `REQ-MCP-016` and this work item have close evidence and valid metadata after task completion.

## Evidence

- `TASK-MCP-015-01`: done. Identified `replaceNamedSection` as the responsible splice path and recorded the spacing loss conditions.
- `TASK-MCP-015-02`: done. Defined the minimal canonical blank-line separator rule and confirmed no SPEC / guidance update was needed.
- `TASK-MCP-015-03`: done. Implemented the `replaceNamedSection` spacing normalization and regression tests in `internal/designrecords/authoring.go` / `internal/designrecords/authoring_test.go`.
- `TASK-MCP-015-04`: done. Confirmed the fix through the MCP runtime `propose_record_update` public tool path.
- Verification:
  - `go test ./internal/designrecords -run TestReplaceNamedSectionSpacing -v`: pass.
  - `go test ./internal/designrecords ./internal/designrecordsmcp`: pass.
  - `go test ./...`: pass.
- Runtime smoke confirmed all three replacement body shapes (`no trailing newline`, `one trailing newline`, `already separated`) produced a blank line before the next heading in proposal diff output:

```text
+Runtime smoke replacement ...
+
+## Evidence
```

- Smoke proposals were discarded and no repository files were accepted for the smoke.
