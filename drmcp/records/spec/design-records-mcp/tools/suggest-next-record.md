# Contract: `suggest_next_record`

- **id**: `spec:drmcp.design_records_mcp.tools.suggest_next_record`
- **status**: draft
- **date**: 2026-06-17
- **parent**: `spec:drmcp.design_records_mcp.tools.overview`
- **contract_class**: `interface`

## What this is

`suggest_next_record` is a P1 read-only tool that assists with creating a new ADR. It suggests the next ADR ID and a recommended filename path from the existing record index. It does not create files.

`next_number` is the maximum existing `decision` record number plus 1. Gaps in the sequence are not filled.

> Source: V01-ADR-077 §suggest_next_record の責務

## Request

```json
{
  "kind": "decision",
  "title": "Design Records MCP implementation package layout"
}
```

| field | required | type | meaning |
|---|---:|---|---|
| `kind` | yes | string | MVP: `decision` only |
| `title` | yes | string | New ADR title (used for slug generation) |

Spec new-record path suggestion is outside MVP scope.

## Response

```json
{
  "kind": "decision",
  "title": "Design Records MCP implementation package layout",
  "next_id": "V01-ADR-078",
  "next_number": 78,
  "suggested_path": "v01/records/adr/V01-ADR-078-design-records-mcp-implementation-package-layout.md",
  "existing_max_id": "V01-ADR-077"
}
```

`suggested_path` is a suggestion only; no file is created.

Filename slug generation rules (applied to `title`):

- Lowercase ASCII alphanumeric characters.
- Non-alphanumeric ASCII character sequences are replaced with `-`.
- Consecutive `-` are collapsed to one.
- Leading and trailing `-` are stripped.
- Non-ASCII characters are treated as `-`.

When the slug would be empty, `suggested_path` may use `v01/records/adr/V01-ADR-{NNN}.md`. The `suggested_path` is advisory; it may be overridden by the person authoring the record.

## Errors

| code | condition |
|---|---|
| `unsupported_kind` | `kind` value other than `decision` |
