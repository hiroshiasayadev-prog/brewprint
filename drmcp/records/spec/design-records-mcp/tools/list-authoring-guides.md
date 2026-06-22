# Contract: `list_authoring_guides`

- **id**: `spec:drmcp.design_records_mcp.tools.list_authoring_guides`
- **status**: draft
- **date**: 2026-06-17
- **parent**: `spec:drmcp.design_records_mcp.tools.overview`
- **contract_class**: `interface`

## What this is

`list_authoring_guides` returns the project authoring guidance catalog so that an AI assistant can discover the correct authoring guide by guide ID for a given authoring target. This tool does not return the Design Records record index. Guides are not a record kind — they are a separate authoring guidance retrieval surface.

## Request

```json
{}
```

No request fields. Implementations may reject requests containing unknown fields with `invalid_request`.

## Response

```json
{
  "guides": [
    {
      "id": "adr-authoring",
      "title": "ADR Authoring Guide",
      "abstract": "ADR を起票・レビュー・更新するときの実践ルールを定める。ADR は設計判断の履歴を所有し、現行仕様本文や作業 checklist を所有しない。"
    }
  ]
}
```

`guides[]` fields:

| field | required | meaning |
|---|---:|---|
| `id` | yes | Guide ID derived from the filename stem of `docs/guides/<id>.md` |
| `title` | yes | First H1 text from the guide file |
| `abstract` | yes | Content of the `## Abstract` section from the guide file |

Guide source file path MUST NOT be included in the response.

`guides[]` is ordered by `id` in ASCII lexical order.

## Errors

No tool-level errors are defined for this tool beyond `invalid_request` for malformed requests.
