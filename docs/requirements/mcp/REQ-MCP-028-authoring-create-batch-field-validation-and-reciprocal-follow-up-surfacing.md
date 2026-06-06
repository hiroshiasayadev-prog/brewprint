# REQ-MCP-028: authoring create batch field validation and reciprocal follow-up surfacing

- **id**: REQ-MCP-028
- **status**: captured
- **date**: 2026-06-07
- **source_refs**:
  - REQ-MCP-023
- **work_items**:

## Requirement

`propose_record_create` MUST surface all currently detectable input problems in a single response. When `fields` is incomplete or a reciprocal follow-up constraint is unmet, callers MUST be able to repair the request without probing one failure at a time.

This requirement covers three criteria carried forward from REQ-MCP-023 that were not resolved by the vocabulary alignment work (ADR-094) and are not covered by REQ-MCP-024:

1. **Batch required field validation**: When `fields` is missing one or more required fields for the target kind, all missing field names MUST be reported together with kind context in a single diagnostic response, rather than failing on the first missing field only.

2. **Reciprocal follow-up surfacing at proposal time**: When a create operation requires reciprocal metadata updates (e.g., adding a new work item to `REQ.work_items`, or a new task to `WORK.tasks`), proposal-time diagnostics MUST surface this requirement clearly, rather than surfacing it only at `accept_proposed_write` time.

3. **`include_required` mode clarity**: When `include_required` mode is necessary for a safe accept, proposal diagnostics or repair guidance MUST identify that requirement clearly so callers can retry with the correct mode without an unexpected accept-time failure.

## Evidence

During TASK-DATA-011 creation, each failure revealed only the next problem:

```text
outputs フィールドが必須でした。body_cache_id を使って再試行しつつ、残りのYAML編集も並行で進めます。
```

```text
required_follow_up_not_satisfied — include_required モードで再起票が必要です。破棄して作り直します。
```

The full failure sequence is documented in REQ-MCP-023. This requirement carries forward only the criteria not resolved by vocabulary alignment (ADR-094) or covered by structured repair guidance format (REQ-MCP-024).

## Required Outcome

`propose_record_create` reports all detectable field and reciprocal follow-up problems in a single proposal response.

Acceptance criteria:

- When `fields` is missing one or more required fields for the target kind, a single response identifies all missing field names and the target kind.
- Detectable required reciprocal follow-up issues are included in proposal-level diagnostics where the server can determine them at proposal time.
- When `reciprocal_update_mode = include_required` is required for a safe accept, proposal diagnostics or repair guidance state that requirement explicitly.
- Diagnostic shapes for missing fields are consistent with the structured diagnostic format defined by REQ-MCP-024.

## Explicitly Excluded Scope

- Status vocabulary or enum definition (addressed by ADR-094 and REQ-MCP-026).
- Synonym repair suggestion mapping (resolved by ADR-094; canonical vocabulary is now `not_started` / `in_progress`).
- `allowed_values` structured format for invalid status diagnostics (covered by REQ-MCP-024).
- Removing `accept_proposed_write` guard checks.
- Automatically accepting proposals after repair guidance is produced.

## Boundary

This requirement belongs to the Design Records MCP authoring create contract. It is the narrowed successor to REQ-MCP-023, carrying forward only the criteria not resolved by ADR-094 (criterion 4: synonym repair; criterion 3: stale token list) or transferred to REQ-MCP-024 (criterion 2: `allowed_values` diagnostic format). It complements REQ-MCP-024 by addressing field-level batch validation and reciprocal follow-up surfacing.
