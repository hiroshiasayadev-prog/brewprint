# V01-REQ-MCP-024: authoring dry-run repair guidance completeness

- **id**: V01-REQ-MCP-024
- **status**: accepted
- **date**: 2026-06-05
- **source_refs**:
  - V01-REQ-MCP-023
- **work_items**:
  - V01-WORK-MCP-025

## Requirement

Design Records MCP authoring proposal responses MUST provide complete, machine-readable repair guidance for common validation and guard failures so callers can construct the next valid request without trial-and-error probing.

When a proposal cannot be created or accepted due to known input problems, diagnostics and repair guidance SHOULD include enough structured information to identify the failing field, valid alternatives, and the minimal safe repair where such a repair is deterministic.

## Evidence

Recent authoring operations required repeated discard/retry cycles because each failure revealed only the next missing or invalid input:

```text
outputs フィールドが必須でした。body_cache_id を使って再試行しつつ、残りのYAML編集も並行で進めます。
```

```text
in_progress はタスクkindで無効でした。
```

```text
not_started もタスクkindでは無効なようです。有効なステータスを確認します。
```

```text
required_follow_up_not_satisfied — include_required モードで再起票が必要です。
```

These should be repairable from a single structured response where the server has enough information to determine the allowed values or required mode.

## Required Outcome

Authoring proposal responses include structured repair guidance for common validation categories.

Acceptance criteria:

- Invalid enum or status diagnostics include `allowed_values` when the server knows the allowed set.
- Missing required field diagnostics identify the required field, target kind, and whether an empty default is valid.
- Deterministic repairs include a machine-readable `repair_suggestion` or equivalent structured field.
- Required reciprocal follow-up failures identify the required mode or missing reciprocal update target when known.
- Proposal-time and accept-time guard failures use consistent diagnostic categories and repair guidance shapes where practical.
- The response remains safe: repair suggestions are advisory and do not bypass validation or accept guards.

Example desired diagnostic shape:

```json
{
  "category": "invalid_metadata_value",
  "field": "status",
  "value": "in_progress",
  "allowed_values": ["todo", "doing", "blocked", "done"],
  "repair_suggestion": {
    "metadata": {
      "status": "doing"
    }
  }
}
```

## Explicitly Excluded Scope

- Automatically applying repair suggestions without an explicit proposal/accept flow.
- Inferring non-deterministic business intent.
- Changing canonical artifact schemas or status vocabularies.
- Replacing human review of proposal diffs.

## Boundary

This requirement belongs to the Design Records MCP authoring diagnostics and repair guidance contract. It complements create/update validation behavior by making failures actionable for tool clients and LLM callers.
