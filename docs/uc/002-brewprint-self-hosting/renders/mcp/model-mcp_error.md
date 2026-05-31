# mcp_error

MCP tool実行自体が成立しない場合に返すtool error payload。
JSON-RPC envelopeそのものではなく、brewprint MCP error bodyを表す。
codeのenum値集合はv1 modelでは型制約できないためnoteに列挙する。

## Public model

| property | value |
|---|---|
| kind | struct |
| visibility | public |
| source | yaml/mcp/model/mcp_error.yaml |

### Fields

| field | type | note |
|---|---|---|
| code | str | error code。<br/>値集合: project_invalid / invalid_args / invalid_selector / invalid_change_payload / not_found / kind_mismatch / ambiguous / unsupported_object / unsupported_detail / unsupported_direction / invalid_depth / internal_error。 |
| message | str | human-readable error message。必須 |
| selector | object_selector | 入力selector。selector解決に関係するerrorの場合のみ |
| diagnostics | diagnostic_list | 関連diagnostic。任意 |

