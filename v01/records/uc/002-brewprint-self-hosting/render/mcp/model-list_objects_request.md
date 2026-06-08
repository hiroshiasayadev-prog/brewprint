# list_objects_request

MCP tool `list_objects` のrequest model。
docs/spec/mcp/tools/list-objects.md §2 のInputに対応する。
`kind` は object-dependent filterであり、正本は docs/spec/mcp/tools/list-objects.md §2 と docs/spec/mcp/schema.md §1.1。
このYAMLはUC-002上のmodel表現であり、DATA DSLにdependent enumを導入するものではない。

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
| kind | str | 任意。object-dependent kind filter。値集合とobject省略時の扱いは docs/spec/mcp/tools/list-objects.md §2 と docs/spec/mcp/schema.md §1.1 を正本とする。 |
| module | str | 任意。module path。例: order, payment.webhooks。 |
| file | str | 任意。FileID。yaml/ からのslash正規化相対path。 |

