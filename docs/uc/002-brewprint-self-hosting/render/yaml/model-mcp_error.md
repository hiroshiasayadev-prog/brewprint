# mcp_error

MCP tool実行自体が成立しない場合に返すtool error payload。
JSON-RPC envelopeそのものではなく、brewprint MCP error bodyを表す。

## Public model

| property | value |
|---|---|
| kind | struct |
| visibility | public |
| source | yaml/yaml/mcp/model/mcp_error.yaml |

### Fields

| field | type | note |
|---|---|---|
| code | mcp_error_code | error code。 |
| message | str | human-readable error message。必須 |
| selector | object_selector | 入力selector。selector解決に関係するerrorの場合のみ |
| diagnostics | diagnostic_list | 関連diagnostic。任意 |

