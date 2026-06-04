# WORK-MCP-019: section heading canonicalization tolerant authoring update support

- **id**: WORK-MCP-019
- **status**: implementation_pending
- **date**: 2026-06-05
- **source_requirement**: REQ-MCP-021
- **impact_refs**:
  - SPEC-design-records-mcp-tools
  - ADR-093
- **tasks**:
  - TASK-MCP-019-01
  - TASK-MCP-019-02
  - TASK-MCP-019-03
  - TASK-MCP-019-04

## Goal

REQ-MCP-021 の required section heading case mismatch support を実現する。

## Boundary

この work item は Design Records MCP authoring update と workflow artifact validation の境界を扱う。

対象外:

- fuzzy heading matching
- canonical section name 自体の変更
- proposal / accept を経由しない mutation
- required-section validation の緩和
- DATA-domain task semantics の変更

## Impact Scope

- `SPEC-design-records-mcp-tools`
- `ADR-093`
- `internal/designrecords` authoring / validation implementation
- `internal/designrecordsmcp` tool call boundary and tests

## Task flow

```text
TASK-MCP-019-01 current gap review
  -> TASK-MCP-019-02 spec contract update
  -> TASK-MCP-019-03 regression tests
  -> TASK-MCP-019-04 implementation
  -> TASK-MCP-019-05 runtime smoke and close synchronization
```

## Task Candidates

- TASK-MCP-019-01: Review current heading selector, required-section validation, and authoring update behavior.
- TASK-MCP-019-02: Update MCP tools spec for safe case-only heading canonicalization proposal behavior.
- TASK-MCP-019-03: Add regression tests for case-only repair, ambiguous failure, non-case mismatch behavior, validation repair diagnostics, and proposal diff canonicalization.
- TASK-MCP-019-04: Implement heading canonicalization support in the authoring update path.
- TASK-MCP-019-05: Run runtime smoke, record evidence, and synchronize close status.

## Completion Condition

- REQ-MCP-021 acceptance criteria are satisfied.
- Case-only required-section heading mismatches can be repaired through MCP proposal / accept flow.
- Ambiguous heading matches fail safely with candidate headings.
- Non-case heading differences remain governed by existing selector and validation rules.
- Tests and runtime smoke evidence are recorded.
- REQ-MCP-021, this work item, and child tasks are status-synchronized.
