# Contract: `discard_proposed_write`

- **id**: `spec:drmcp.design_records_mcp.tools.discard_proposed_write`
- **status**: draft
- **date**: 2026-06-17
- **parent**: `spec:drmcp.design_records_mcp.tools.overview`
- **contract_class**: `interface`

## What this is

`discard_proposed_write` discards a retained proposal, preventing future acceptance. It does not write repository files.

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
  "state": "discarded",
  "discarded": true,
  "written": false,
  "diagnostics": []
}
```

Discarding an already-accepted proposal must not undo the accepted write.

When the proposal ID is unknown or expired, the response returns diagnostics and `discarded: false`.

## Errors

| code | condition |
|---|---|
| `proposal_not_found` | Requested proposal ID does not exist |
| `proposal_expired` | Proposal is past expiry |
