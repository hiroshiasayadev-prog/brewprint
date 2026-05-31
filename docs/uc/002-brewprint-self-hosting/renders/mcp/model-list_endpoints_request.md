# list_endpoints_request

MCP tool `list_endpoints` のrequest model。
docs/spec/mcp/tools/list-endpoints.md §2 のInputに対応する。
API Tableが複数存在し、api_table_idが省略された場合は、全API Tableをresponse.tables[]に分けて返す。
このtool自体はMCP toolでありHTTP endpointではないため、対応taskに endpoint: true は付けない。

## Public model

| property | value |
|---|---|
| kind | struct |
| visibility | public |
| source | yaml/mcp/model/list_endpoints_request.yaml |

### Fields

| field | type | note |
|---|---|---|
| api_table_id | str | 任意。API Table view ID。省略時はproject内の全API Tableを返す。 |

