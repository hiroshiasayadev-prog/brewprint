# list_objects_request

MCP tool `list_objects` のrequest model。
docs/spec/mcp/tools/list-objects.md §2 のInputに対応する。
`kind` は object種別ごとのkind filterであり、object-dependent語彙は REQ-DATA-007 / WORK-DATA-014 スコープのためnoteで保持する。

## Public model

| property | value |
|---|---|
| kind | struct |
| visibility | public |
| source | yaml/mcp/model/list_objects_request.yaml |

### Fields

| field | type | note |
|---|---|---|
| object | list_objects_object_filter | 任意。省略時は全object種別を対象にする。 |
| kind | str | 任意。task / model / api_table / transition / field 等。object種別ごとのkind filter。 |
| module | str | 任意。module path。例: order, payment.webhooks。 |
| file | str | 任意。FileID。yaml/ からのslash正規化相対path。 |

