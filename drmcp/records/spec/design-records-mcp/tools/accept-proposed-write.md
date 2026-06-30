# Contract: `accept_proposed_write`

- **id**: `spec:drmcp.design_records_mcp.tools.accept_proposed_write`
- **status**: draft
- **date**: 2026-06-30
- **parent**: `spec:drmcp.design_records_mcp.tools.overview`
- **contract_class**: `interface`

## What this is

`accept_proposed_write` applies a retained proposal after accept-time checks. This is the only Design Records MCP authoring tool that may write repository files.

## Request

```json
{
  "proposal_id": "pw_opaque"
}
```

| field | required | type | meaning |
|---|---:|---|---|
| `proposal_id` | yes | string | Proposal lookup key |

## Response

```json
{
  "proposal_id": "pw_opaque",
  "state": "accepted",
  "written": true,
  "files_written": [
    {
      "path": "v01/records/tasks/mcp/V01-TASK-MCP-008-04-mcp-tools-spec-reflection.md",
      "record_id": "V01-TASK-MCP-008-04",
      "record_kind": "task"
    }
  ],
  "validation": {
    "ok": true,
    "diagnostics": []
  },
  "repair_guidance": [],
  "diagnostics": []
}
```

| field | required | meaning |
|---|---:|---|
| `proposal_id` | yes | Accepted proposal ID |
| `state` | yes | `accepted` on successful write; otherwise the retained or rejection state |
| `written` | yes | Whether repository files were modified by this accept call |
| `files_written` | yes | Written file list using normalized repository-relative paths; empty when `written: false` |
| `validation` | yes | Post-accept validation result when applicable; otherwise current validation result |
| `repair_guidance` | yes | Actionable repair suggestions; empty when no repair is needed |
| `diagnostics` | yes | Accept diagnostics |

`written: false` is required when the proposal is unknown, expired, discarded, already accepted, stale, has a changed target, ID collision, unresolved target, invalid proposal, required-follow-up-not-satisfied, or cannot construct every affected target and required source-backed diagnostic location before writing. None of these outcomes may modify repository files, and `files_written` is empty.

`written: true` means repository files were modified. Each `files_written[].path` uses the same normalized repository-relative spelling as the corresponding proposal target, `diff.files[]` entry, and unified-diff operand. Absolute host paths and backslash-separated Windows paths are prohibited.

Current-format post-write validation integration is deferred to `DRMCP-REQ-MCP-002`.

This contract does not require the current authoring transaction to invoke the W011 snapshot architecture after writing. It does not treat YAML or V01-SPEC authoring semantics as integrated with current-format validation.

If post-write validation fails after writing, the response must still return `written: true`, include validation diagnostics, and provide repair guidance. MVP does not automatically roll back accepted writes after post-write validation failure. The caller should create a repair proposal. An implementation failure discovered after files were actually modified must not report `written: false` or erase the actual written-file state.

This contract retains `written`, `files_written`, post-write failure reporting, repair guidance, and no-automatic-rollback behavior. `spec:drmcp.implementation` does not own current authoring-transaction validation integration.

Force-accepting invalid proposals (proposals with pre-write error diagnostics) is outside MVP. Pre-write error diagnostics cause `written: false`.

## Errors

| code | condition |
|---|---|
| `proposal_not_found` | Requested proposal ID does not exist |
| `proposal_expired` | Proposal is past expiry |
| `proposal_discarded` | Proposal is already discarded |
| `proposal_already_accepted` | Proposal is already accepted |
| `proposal_stale` | Base state and current target state do not match |
| `target_changed` | Target record kind / path / identity differs from proposal creation time |
| `id_collision` | Create proposal's resolved ID was claimed before acceptance |
| `required_follow_up_not_satisfied` | Required follow-up updates are not yet satisfied |
