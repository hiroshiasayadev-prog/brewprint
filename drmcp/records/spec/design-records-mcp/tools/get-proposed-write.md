# Contract: `get_proposed_write`

- **id**: `spec:drmcp.design_records_mcp.tools.get_proposed_write`
- **status**: draft
- **date**: 2026-06-17
- **parent**: `spec:drmcp.design_records_mcp.tools.overview`
- **contract_class**: `interface`

## What this is

`get_proposed_write` retrieves a retained proposal by proposal ID. It does not write repository files.

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

For a retained proposal, returns the proposal detail using the common proposal fields and current `state`. Always returns the full diff (`diff.text` included — `patch`-equivalent regardless of the original `diff_mode` used at proposal creation time).

## Errors

| code | condition |
|---|---|
| `proposal_not_found` | Requested proposal ID does not exist |
| `proposal_expired` | Requested proposal ID is past expiry |
