# Contract: `accept_proposed_write`

- **id**: `spec:drmcp.design_records_mcp.tools.accept_proposed_write`
- **status**: draft
- **date**: 2026-06-17
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
| `files_written` | yes | Written file list; empty when `written: false` |
| `validation` | yes | Post-accept validation result when applicable; otherwise current validation result |
| `repair_guidance` | yes | Actionable repair suggestions; empty when no repair is needed |
| `diagnostics` | yes | Accept diagnostics |

`written: false` is required when the proposal is unknown, expired, discarded, already accepted, stale, has a changed target, ID collision, unresolved target, invalid proposal, or required-follow-up-not-satisfied. None of these outcomes may modify repository files.

`written: true` means repository files were modified. If post-write validation fails after writing, the response must still return `written: true`, include validation diagnostics, and provide repair guidance. MVP does not automatically roll back accepted writes after post-write validation failure. The caller should create a repair proposal.

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
