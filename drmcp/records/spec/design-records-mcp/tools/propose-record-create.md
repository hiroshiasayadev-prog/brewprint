# Contract: `propose_record_create`

- **id**: `spec:drmcp.design_records_mcp.tools.propose_record_create`
- **status**: draft
- **date**: 2026-06-28
- **parent**: `spec:drmcp.design_records_mcp.tools.overview`
- **contract_class**: `interface`

## What this is

`propose_record_create` creates a retained proposal for a new record or workflow artifact. It does not write repository files.

MVP create support: `decision`, `requirement`, `work_item`, `task`. Spec skeleton creation and investigation creation are outside MVP.

## Request

```json
{
  "kind": "task",
  "id": "V01-TASK-MCP-008-new",
  "domain": "MCP",
  "parent_id": "V01-WORK-MCP-008",
  "title": "MCP tools spec reflection",
  "fields": {
    "status": "not_started",
    "date": "2026-06-01",
    "source_requirement": "V01-REQ-MCP-008",
    "estimate": "1d-2d",
    "depends_on": ["V01-TASK-MCP-008-03"],
    "outputs": ["Updated V01-SPEC-design-records-mcp-tools"]
  },
  "body": "## Goal\n\nReflect the accepted authoring transaction contract in the Design Records MCP tools spec.\n\n## Work\n\n- Update the public tool contract.\n- Record verification evidence.\n\n## Done condition\n\nThe spec documents the current request and response contract.\n\n## Verification\n\n- go test ./internal/designrecords ./internal/designrecordsmcp\n\n## Evidence\n",
  "reciprocal_update_mode": "include_required"
}
```

| field | required | type | meaning |
|---|---:|---|---|
| `kind` | yes | string | Create target kind |
| `id` | yes | string | Exact public ID or allowed `new` placeholder ID |
| `domain` | conditional | string | Domain for domain-scoped workflow create. Must case-insensitively match the ID domain when the ID contains it. |
| `parent_id` | conditional | string | Required for task create. Parent work item public ID. |
| `title` | yes | string | H1 title |
| `fields` | yes | object | Kind-specific structured authoring fields. MCP generates H1 and metadata from these. |
| `body` | conditional | string | Section-only content body. Valid only when combined with `fields`. |
| `body_cache_id` | conditional | string | Cached section-only body lookup key for `fields + body` retry. |
| `reciprocal_update_mode` | no | string | Workflow reciprocal metadata handling mode |
| `diff_mode` | no | string | Diff return format in proposal response: `summary` (default) / `patch` / `none` |

Content input combinations:

| mode | required content input | forbidden content input |
|---|---|---|
| Structured metadata only | `fields` | `body`, `body_cache_id` |
| Structured metadata plus content sections | `fields` plus section-only `body` | `body_cache_id` |
| Retry with cached content sections | `fields` plus `body_cache_id` | `body` |
| Body-only create | — | invalid; `fields` is required |
| Cache-only create | — | invalid; `fields` is required |

In all modes, top-level `kind`, `id`, `domain`, `parent_id`, and `title` remain request-level target inputs.

When `fields` is present, MCP renders the record H1 and metadata block from top-level `id`, `title`, server-resolved target identity, and `fields`. When `body` is also present, it is appended after the generated metadata block as content sections only. `body` must start at the first content section (e.g. `## Goal`, `## Requirement`, `## 背景`) and must not include a leading H1, bullet metadata block, YAML metadata, metadata `id`, or a guessed resolved ID. For `new` placeholder IDs, the generated H1 and metadata must use `target.resolved_id`, not the literal `new` token.

`fields.id` is not required for create. When supplied with an exact top-level ID, it must exactly match the top-level ID after the same canonical normalization. When supplied with a `new` placeholder top-level ID, the request is invalid (the final ID is unknown to the caller before server-side resolution). Mismatch or placeholder-time `fields.id` produces `invalid_request` and creates no proposal.

For domain-scoped workflow creates, `domain` is compared with the ID domain case-insensitively. Canonical IDs keep the uppercase domain segment. Repository paths use the lowercase normalized domain directory (e.g. `domain: "mcp"` + `id: "V01-REQ-MCP-011"` → canonical domain `MCP`, path domain `mcp`).

Allowed `id` forms (public ID with namespace_prefix; MVP: `namespace_prefix = V01-`):

| kind | allowed create ID forms |
|---|---|
| `decision` | exact `V01-ADR-NNN` or `V01-ADR-new` |
| `spec` | not supported for create in MVP |
| `requirement` | exact `V01-REQ-<DOMAIN>-NNN` or `V01-REQ-<DOMAIN>-new` |
| `work_item` | exact `V01-WORK-<DOMAIN>-NNN` or `V01-WORK-<DOMAIN>-new` |
| `task` | exact `V01-TASK-<DOMAIN>-<WORK-SEQUENCE>-NN` or `V01-TASK-<DOMAIN>-<WORK-SEQUENCE>-new` |

`new` is the only accepted placeholder token in the sequence position. Any other token in the sequence position is rejected as `invalid_request`. `new` is valid only for create operations. MCP resolves the final ID using the current record index; gaps are not filled unless a later spec changes this rule.

Exact ID create is allowed for workflow artifacts but may return a non-blocking `exact_id_sequence_gap` info diagnostic in `diagnostics` (not `validation.diagnostics`) when the requested ID would create a sequence gap. `new` placeholder creates never emit `exact_id_sequence_gap`. Gap-fill creates (filling an existing gap) do not emit `exact_id_sequence_gap`; existing duplicate ID checks still reject an already-indexed exact ID.

Task create requires `parent_id`. The parent work item must resolve to an indexed `work_item` record. For task placeholder IDs, `<DOMAIN>` and `<WORK-SEQUENCE>` must match the parent work item ID. Task parent relation must be written from explicit metadata, not inferred from ID shape.

`reciprocal_update_mode` values:

| value | meaning |
|---|---|
| `include_required` | Default. Include required reciprocal workflow metadata updates in the same proposal. Returns `reciprocal_update_included` info diagnostic for each included update. `required_follow_up_updates` remains empty. |
| `report_required_follow_up` | Do not include reciprocal file updates; return explicit required follow-up updates and reject acceptance until they are satisfied. |

`include_required` may create a multi-file proposal but only for required reciprocal metadata updates (e.g. adding a new work item to `REQ.work_items` or a new task to `WORK.tasks`). It is not a general-purpose multi-record atomic transaction. Unsafe or ambiguous cases remain blocking or report-only.

> Source: V01-REQ-MCP-011, V01-TASK-MCP-011-01

## Response

```json
{
  "proposal_id": "pw_opaque",
  "state": "proposed",
  "operation": "create",
  "target_kind": "task",
  "target": {
    "requested_id": "V01-TASK-MCP-008-new",
    "resolved_id": "V01-TASK-MCP-008-04",
    "kind": "task",
    "domain": "MCP",
    "parent_id": "V01-WORK-MCP-008",
    "path": "v01/records/tasks/mcp/V01-TASK-MCP-008-04-mcp-tools-spec-reflection.md"
  },
  "expires_at": "2026-06-05T00:00:00Z",
  "retention_days": 3,
  "diff": {
    "format": "unified",
    "files": [
      {
        "path": "v01/records/tasks/mcp/V01-TASK-MCP-008-04-mcp-tools-spec-reflection.md",
        "change": "create",
        "record_id": "V01-TASK-MCP-008-04",
        "record_kind": "task"
      }
    ],
    "text": "diff --git a/v01/records/tasks/mcp/V01-TASK-MCP-008-04-mcp-tools-spec-reflection.md b/v01/records/tasks/mcp/V01-TASK-MCP-008-04-mcp-tools-spec-reflection.md\nnew file mode 100644\n--- /dev/null\n+++ b/v01/records/tasks/mcp/V01-TASK-MCP-008-04-mcp-tools-spec-reflection.md\n..."
  },
  "validation": {
    "ok": true,
    "diagnostics": []
  },
  "diagnostics": [],
  "note": "No repository files have been written. Call accept_proposed_write with this proposal_id to apply the diff."
}
```

When `reciprocal_update_mode: "include_required"` applies a reciprocal update, `diagnostics` includes a `reciprocal_update_included` info diagnostic:

```json
{
  "diagnostics": [
    {
      "category": "reciprocal_update_included",
      "severity": "info",
      "message": "reciprocal update included: V01-WORK-MCP-008.tasks will receive V01-TASK-MCP-008-04",
      "subject": {
        "type": "record",
        "ref": "V01-WORK-MCP-008",
        "record_kind": "work_item"
      },
      "field": {
        "name": "tasks"
      },
      "value": {
        "actual": "V01-TASK-MCP-008-04"
      }
    }
  ]
}
```

In this case, `diff.files` includes both the new record create entry and the parent record modify entry. `required_follow_up_updates` is empty.
The diagnostic uses the shared envelope and structured associations from `spec:drmcp.design_records_mcp.schema.diagnostics`; its trigger and proposal behavior remain owned by this operation.

When `reciprocal_update_mode: "report_required_follow_up"` is used and required follow-ups exist, `required_follow_up_updates` is non-empty and acceptance is rejected with `written: false` until they are satisfied.

Exact ID sequence-gap warning example:

```json
{
  "diagnostics": [
    {
      "category": "exact_id_sequence_gap",
      "severity": "info",
      "message": "V01-REQ-MCP-020 skips the next available sequence V01-REQ-MCP-019; prefer V01-REQ-MCP-new unless this ID is intentional"
    }
  ]
}
```

## Errors

| code | condition |
|---|---|
| `invalid_request` | Missing `fields`, `fields.id` mismatch, `new` placeholder with `fields.id`, domain/ID domain mismatch, or other invalid request |
| `invalid_body_source` | Both `body` and `body_cache_id` supplied, or required body source rule violated |
| `body_cache_not_found` | `body_cache_id` does not exist |
| `body_cache_expired` | `body_cache_id` is past expiry |
| `proposal_preparation_failed` | Proposal preparation failed before persistence |
