# V01-REQ-MCP-031: authoring proposal diff response output mode with summary default

- **id**: V01-REQ-MCP-031
- **status**: accepted
- **date**: 2026-06-07
- **source_refs**:
  - V01-REQ-MCP-027
  - V01-WORK-MCP-026
- **work_items**:
  - V01-WORK-MCP-028

## Requirement

Design Records MCP authoring proposal tools should support a request-controlled diff response output mode so agents can avoid receiving large unified diff text unless needed.

The default diff response mode should be `summary`, not full patch text.

Current proposal responses include unified diff text in `diff.text`. For create proposals, this includes the full newly generated Markdown file body as `+` lines. This makes proposal JSON large even when the request used `body_cache_id` to avoid resending long body content.

## Evidence

Observed `propose_record_create` responses for `V01-TASK-MCP-026-01..05` included full create diffs. The diffs were semantically correct, but the JSON response was large because every new task body appeared inside `diff.text`.

This is useful for human review, but inefficient for agent workflows that only need changed-file metadata, validation, diagnostics, and proposal IDs before accepting or fetching details separately.

## Required Outcome

- `propose_record_create` accepts `diff_mode` with values `summary`, `patch`, and `none`.
- `propose_record_update` accepts `diff_mode` with values `summary`, `patch`, and `none`.
- When `diff_mode` is omitted, the default is `summary`.
- `summary` mode omits unified diff text and returns changed file metadata plus concise operation summaries.
- `patch` mode preserves the current unified diff response, including `diff.text`.
- `none` mode omits diff details and returns an explicit omitted marker while still returning proposal ID, target, validation, diagnostics, and body cache metadata when available.
- Validation, diagnostics, proposal retention, body cache behavior, and accept behavior are independent of `diff_mode`.
- Invalid `diff_mode` values return `invalid_request` or an equivalent blocking diagnostic.

## Explicitly Excluded Scope

- Do not remove the ability to request full unified diff text.
- Do not change the retained proposal content or accept behavior based on response diff mode.
- Do not use `diff_mode` to hide validation or diagnostics.
- Do not require agents to fetch full diff text before `accept_proposed_write` if validation is clean and the caller intentionally selected summary/none mode.

## Boundary

This requirement concerns response shaping for Design Records MCP authoring proposal tools. It does not define git diff generation semantics themselves, no-op detection semantics, or update operation ordering beyond what existing requirements and specs cover.
