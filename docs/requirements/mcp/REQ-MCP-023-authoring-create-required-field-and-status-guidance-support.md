# REQ-MCP-023: authoring create required field and status guidance support

- **id**: REQ-MCP-023
- **status**: accepted
- **date**: 2026-06-05
- **source_refs**:
  - REQ-MCP-020
  - REQ-MCP-021
  - REQ-MCP-022
- **work_items**:

## Requirement

Design Records MCP authoring create operations MUST provide clear, kind-specific guidance for required fields, allowed status values, and required reciprocal follow-up behavior before callers reach repeated discard/retry cycles.

When `propose_record_create` receives incomplete or invalid fields for the target kind, diagnostics SHOULD identify all currently detectable issues in one response where possible, including missing required fields and invalid enum/status values.

For workflow artifact creation, invalid status diagnostics MUST include the allowed values for that artifact kind.

## Evidence

During task creation and close synchronization, the authoring flow required several discard/retry cycles:

```text
outputs フィールドが必須でした。body_cache_id を使って再試行しつつ、残りのYAML編集も並行で進めます。
```

```text
in_progress はタスクkindで無効でした。破棄して not_started で再起票します。
```

```text
not_started もタスクkindでは無効なようです。有効なステータスを確認します。
```

```text
有効なステータスは todo / doing / blocked / done でした。doing で再起票します。
```

The flow also reached accept-time failure for required follow-up behavior:

```text
required_follow_up_not_satisfied — include_required モードで再起票が必要です。破棄して作り直します。
```

These errors were individually discoverable only after multiple attempts. The authoring contract should guide callers toward a valid create request earlier.

## Required Outcome

`propose_record_create` provides kind-aware required field and status guidance sufficient for callers to repair common create errors without probing one failure at a time.

Acceptance criteria:

- Missing required fields for the requested kind are reported with field names and kind context.
- Invalid status diagnostics include the allowed status values for the requested kind.
- Task status diagnostics identify `todo`, `doing`, `blocked`, and `done` as allowed values.
- When a supplied status is a common non-canonical synonym or near-miss, diagnostics provide a repair suggestion when safe, for example `in_progress` to `doing`.
- Detectable required reciprocal follow-up issues are surfaced at proposal time where possible, rather than only at `accept_proposed_write` time.
- When `reciprocal_update_mode = include_required` is required for a safe accept, proposal diagnostics or repair guidance identify that requirement clearly.

## Explicitly Excluded Scope

- Changing canonical status vocabularies.
- Silently rewriting invalid status values without proposal-visible diagnostics.
- Removing `accept_proposed_write` guard checks.
- Automatically accepting proposals after repair suggestions are produced.

## Boundary
This requirement belongs to the Design Records MCP authoring create contract. It improves diagnostics and repair guidance for create-time inputs and reciprocal follow-up handling. It does not change DATA-domain task content or workflow semantics.

**Supersede note (2026-06-07)**: Accepted as superseded in practice via the following decomposition:

- **Criterion 3** (task allowed values `todo / doing / blocked / done`): stale. ADR-094 changed canonical values to `not_started / in_progress / blocked / done`.
- **Criterion 4** (synonym repair, e.g., `in_progress → doing`): resolved. `in_progress` is now canonical per ADR-094; no synonym mapping needed.
- **Criterion 2** (`allowed_values` in status diagnostics): transferred to REQ-MCP-024.
- **Criteria 1, 5, 6** (batch field validation, reciprocal follow-up surfacing, `include_required` mode clarity): transferred to REQ-MCP-028.
